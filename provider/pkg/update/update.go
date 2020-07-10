package update

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/pkg/errors"
	"github.com/sanathkr/yaml"
	uuid "github.com/satori/go.uuid"

	"github.com/pulumi/pulumi-cloudformation/provider/pkg/get"
)

type ResourceOption func(*ResourceData)

func Metadata(metadata map[string]interface{}) ResourceOption {
	return func(r *ResourceData) {
		r.Metadata = metadata
	}
}

func CreationPolicy(policy map[string]interface{}) ResourceOption {
	return func(r *ResourceData) {
		r.CreationPolicy = policy
	}
}

func DeletionPolicy(policy string) ResourceOption {
	return func(r *ResourceData) {
		r.DeletionPolicy = policy
	}
}

// nolint: golint
func UpdatePolicy(policy map[string]interface{}) ResourceOption {
	return func(r *ResourceData) {
		r.UpdatePolicy = policy
	}
}

// nolint: golint
func UpdateReplacePolicy(policy string) ResourceOption {
	return func(r *ResourceData) {
		r.UpdateReplacePolicy = policy
	}
}

type ResourceData struct {
	LogicalID           string
	PhysicalResourceID  string
	Type                string
	Metadata            map[string]interface{}
	Properties          map[string]interface{}
	CreationPolicy      map[string]interface{}
	DeletionPolicy      string
	UpdatePolicy        map[string]interface{}
	UpdateReplacePolicy string
	Attributes          map[string]interface{}
}

type resourceResponse struct {
	data *ResourceData
	err  error
}

type resourceRequest struct {
	isRead   bool
	resource *ResourceData
	response chan<- resourceResponse
}

type StackUpdate struct {
	client *cloudformation.CloudFormation
	getter *get.Getter

	stackName string
	stackID   string

	pendingRequests      []resourceRequest
	pendingRequestsMutex sync.Mutex
	pendingRequestsCond  *sync.Cond
}

func StartStackUpdate(ctx context.Context, configProvider client.ConfigProvider, stackName string) (*StackUpdate, error) {
	updater := &StackUpdate{
		client:    cloudformation.New(configProvider),
		getter:    get.NewGetter(configProvider),
		stackName: stackName,
	}
	updater.pendingRequestsCond = sync.NewCond(&updater.pendingRequestsMutex)

	stackID, err := updater.startUpdate(ctx)
	if err != nil {
		return nil, err
	}

	updater.stackID = stackID
	return updater, nil
}

func (su *StackUpdate) StackName() string {
	return su.stackName
}

func (su *StackUpdate) StackID() string {
	return su.stackID
}

func unmarshalCloudformationTemplate(body string) (map[string]interface{}, error) {
	asJSON, err := yaml.YAMLToJSON([]byte(body))
	if err != nil {
		return nil, errors.Wrap(err, "invalid YAML")
	}

	var template map[string]interface{}
	if err = json.Unmarshal(asJSON, &template); err != nil {
		return nil, errors.Wrap(err, "invalid YAML")
	}

	return template, nil
}

func marshalCloudformationTemplate(template map[string]interface{}) (string, error) {
	asJSON, err := json.Marshal(template)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal as JSON")
	}

	body, err := yaml.JSONToYAML(asJSON)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal as YAML")
	}

	return string(body), nil
}

func (su *StackUpdate) waitForStack(ctx context.Context) error {
	for {
		describeResp, err := su.client.DescribeStacksWithContext(ctx, &cloudformation.DescribeStacksInput{
			StackName: aws.String(su.stackName),
		})
		if err != nil {
			return errors.Wrap(err, "failed to describe stack")
		}
		if len(describeResp.Stacks) != 1 {
			return errors.New("failed to describe stack")
		}

		// TODO: look through stack events for status reasons.

		status := aws.StringValue(describeResp.Stacks[0].StackStatus)
		switch status {
		case "CREATE_COMPLETE", "UPDATE_COMPLETE":
			return nil
		case "CREATE_FAILED", "ROLLBACK_FAILED", "ROLLBACK_COMPLETE", "UPDATE_ROLLBACK_FAILED", "UPDATE_ROLLBACK_COMPLETE":
			return errors.Errorf("update failed: %v", aws.StringValue(describeResp.Stacks[0].StackStatusReason))
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

var allCapabilities = []*string{aws.String("CAPABILITY_IAM"), aws.String("CAPABILITY_NAMED_IAM"), aws.String("CAPABILITY_AUTO_EXPAND")}

func (su *StackUpdate) updateStack(ctx context.Context, requests []resourceRequest) error {
	// Fetch the ground state for the stack's resources.
	existingStates := make(map[string]*cloudformation.StackResourceSummary)
	err := su.client.ListStackResourcesPagesWithContext(ctx, &cloudformation.ListStackResourcesInput{
		StackName: aws.String(su.stackName),
	}, func(resp *cloudformation.ListStackResourcesOutput, _ bool) bool {
		for _, r := range resp.StackResourceSummaries {
			existingStates[aws.StringValue(r.LogicalResourceId)] = r
		}
		return true
	})
	if err != nil {
		return errors.Wrap(err, "failed to get existing resource states")
	}

	// Fetch the latest template.
	templateResp, err := su.client.GetTemplateWithContext(ctx, &cloudformation.GetTemplateInput{
		StackName: aws.String(su.stackName),
	})
	if err != nil {
		return errors.Wrap(err, "failed to fetch template")
	}

	// Parse the template.
	template, err := unmarshalCloudformationTemplate(*templateResp.TemplateBody)
	if err != nil {
		return err
	}

	resourcesV, ok := template["Resources"]
	if !ok {
		return errors.New("invalid template: missing 'Resources' section")
	}
	resources, ok := resourcesV.(map[string]interface{})
	if !ok {
		return errors.New("invalid template: malformed 'Resources' section")
	}

	// Remove the null resource and the never condition if they are present.
	if _, ok := resources["NullResource"]; ok {
		delete(resources, "NullResource")

		conditionsV, ok := template["Conditions"]
		if ok {
			conditions, ok := conditionsV.(map[string]interface{})
			if ok {
				delete(conditions, "Never")
			}
		}
	}

	// Patch the resources in the template.
	hasUpdates, missingReads, missingDeletes := false, make(map[*ResourceData]bool), make(map[*ResourceData]bool)
	for _, r := range requests {
		// Look for the existing state and definition of this resource, if any.
		definition, hasDefinition := resources[r.resource.LogicalID].(map[string]interface{})
		definitionType, _ := definition["Type"].(string)
		state, hasState := existingStates[r.resource.LogicalID]

		// Check for a read.
		if r.isRead {
			if !hasDefinition || r.resource.Type != definitionType {
				missingReads[r.resource] = true
				continue
			}
			r.resource.Metadata, _ = definition["Metadata"].(map[string]interface{})
			r.resource.Properties, _ = definition["Properties"].(map[string]interface{})
			r.resource.CreationPolicy, _ = definition["CreationPolicy"].(map[string]interface{})
			r.resource.DeletionPolicy, _ = definition["DeletionPolicy"].(string)
			r.resource.UpdatePolicy, _ = definition["UpdatePolicy"].(map[string]interface{})
			r.resource.UpdateReplacePolicy, _ = definition["UpdateReplacePolicy"].(string)
			continue
		}

		hasUpdates = true

		// Check for a delete.
		if r.resource.Type == "" {
			if !hasDefinition || !hasState {
				missingDeletes[r.resource] = true
				continue
			}
			if aws.StringValue(state.PhysicalResourceId) != r.resource.PhysicalResourceID {
				// Ignore this: it is likely a resource that was replaced in an earlier change.
				continue
			}

			delete(resources, r.resource.LogicalID)
			continue
		}

		resource := map[string]interface{}{
			"Type": r.resource.Type,
		}
		if r.resource.Metadata != nil {
			resource["Metadata"] = r.resource.Metadata
		}
		if r.resource.Properties != nil {
			resource["Properties"] = r.resource.Properties
		}
		if r.resource.CreationPolicy != nil {
			resource["CreationPolicy"] = r.resource.CreationPolicy
		}
		if r.resource.DeletionPolicy != "" {
			resource["DeletionPolicy"] = r.resource.DeletionPolicy
		}
		if r.resource.UpdatePolicy != nil {
			resource["UpdatePolicy"] = r.resource.UpdatePolicy
		}
		if r.resource.UpdateReplacePolicy != "" {
			resource["UpdateReplacePolicy"] = r.resource.UpdateReplacePolicy
		}

		resources[r.resource.LogicalID] = resource
	}

	// Marshal the template to YAML.
	var body string
	if len(resources) == 0 {
		body = emptyStackBody
	} else {
		body, err = marshalCloudformationTemplate(template)
		if err != nil {
			return err
		}
	}

	// Generate an update token.
	token := "pulumi-" + uuid.NewV4().String()

	// Create or update the stack.
	newStates := existingStates
	if hasUpdates {
		_, err = su.client.UpdateStackWithContext(ctx, &cloudformation.UpdateStackInput{
			ClientRequestToken: aws.String(token),
			StackName:          aws.String(su.stackName),
			TemplateBody:       aws.String(body),
			Capabilities:       allCapabilities,
		})
		if err != nil {
			// We may have failed because there's nothing to do. Ignore this failure and skip fetching new resource states.
			awsErr, ok := err.(awserr.Error)
			if !ok || awsErr.Code() != "ValidationError" || awsErr.Message() != "No updates are to be performed." {
				return errors.Wrap(err, "failed to start update")
			}
		} else {
			// Wait for the update to complete.
			if err = su.waitForStack(ctx); err != nil {
				return err
			}

			// Fetch the latest state of the stack's resources.
			newStates = make(map[string]*cloudformation.StackResourceSummary)
			err = su.client.ListStackResourcesPagesWithContext(ctx, &cloudformation.ListStackResourcesInput{
				StackName: aws.String(su.stackName),
			}, func(resp *cloudformation.ListStackResourcesOutput, _ bool) bool {
				for _, r := range resp.StackResourceSummaries {
					newStates[aws.StringValue(r.LogicalResourceId)] = r
				}
				return true
			})
			if err != nil {
				return errors.Wrap(err, "failed to get new resource states")
			}
		}
	}

	getData := func(r resourceRequest) (*ResourceData, error) {
		// Check for a delete
		if r.resource.Type == "" {
			if missingDeletes[r.resource] {
				return nil, errors.Errorf("delete failed: could not find resource %v with ID %v", r.resource.LogicalID, r.resource.PhysicalResourceID)
			}
			return nil, nil
		}

		// Check for a missing read
		if r.isRead && missingReads[r.resource] {
			r.resource.PhysicalResourceID = ""
			return r.resource, nil
		}

		state, ok := newStates[r.resource.LogicalID]
		if !ok && !r.isRead {
			return nil, errors.New("failed to get attributes: no resource summary")
		}

		reason := aws.StringValue(state.ResourceStatusReason)
		switch aws.StringValue(state.ResourceStatus) {
		case "CREATE_FAILED":
			return nil, errors.Errorf("create failed: %v", reason)
		case "DELETE_FAILED":
			return nil, errors.Errorf("delete failed: %v", reason)
		case "UPDATE_FAILED":
			return nil, errors.Errorf("update failed: %v", reason)
		}

		physicalResourceID := aws.StringValue(state.PhysicalResourceId)
		if physicalResourceID != "" {
			attrs, err := su.getter.GetResourceAttributes(ctx, r.resource.Type, physicalResourceID)
			if err != nil {
				if get.IsNotFound(err) && r.isRead {
					physicalResourceID = ""
				} else {
					return nil, errors.Wrap(err, "failed to get attributes")
				}
			}
			r.resource.Attributes = attrs
		}
		r.resource.PhysicalResourceID = physicalResourceID
		return r.resource, nil
	}

	// Get the attributes for each requested resource.
	for _, r := range requests {
		data, err := getData(r)
		r.response <- resourceResponse{data: data, err: err}
	}

	return nil
}

const emptyStackBody = `Conditions:
  Never: !Equals [ a, b ]
Resources:
  NullResource:
    Type: Custom::Null
    Condition: Never
`

func (su *StackUpdate) startUpdate(ctx context.Context) (string, error) {
	var stackID string
	describeResp, err := su.client.DescribeStacksWithContext(ctx, &cloudformation.DescribeStacksInput{
		StackName: aws.String(su.stackName),
	})
	switch {
	case err != nil:
		awsErr, ok := err.(awserr.Error)
		if !ok || awsErr.Code() != "ValidationError" {
			return "", errors.Wrap(err, "failed to describe stack")
		}

		createResp, err := su.client.CreateStackWithContext(ctx, &cloudformation.CreateStackInput{
			StackName:    aws.String(su.stackName),
			TemplateBody: aws.String(emptyStackBody),
			Capabilities: allCapabilities,
		})
		if err != nil {
			return "", errors.Wrap(err, "failed to create stack")
		}
		stackID = aws.StringValue(createResp.StackId)

		if err = su.waitForStack(ctx); err != nil {
			return "", err
		}
	case len(describeResp.Stacks) != 1:
		return "", errors.New("failed to describe stack")
	default:
		stackID = aws.StringValue(describeResp.Stacks[0].StackId)
	}

	go func() {
		for {
			su.pendingRequestsCond.L.Lock()
			for len(su.pendingRequests) == 0 {
				if ctx.Err() != nil {
					return
				}
				su.pendingRequestsCond.Wait()
			}
			reqs := su.pendingRequests
			su.pendingRequests = nil
			su.pendingRequestsCond.L.Unlock()

			if err := su.updateStack(ctx, reqs); err != nil {
				for _, r := range reqs {
					r.response <- resourceResponse{err: err}
				}
			}

			time.Sleep(1 * time.Second)
		}
	}()

	return stackID, nil
}

func (su *StackUpdate) handleRequest(ctx context.Context, resource *ResourceData, isRead bool) (*ResourceData, error) {
	responseChan := make(chan resourceResponse)
	req := resourceRequest{
		isRead:   isRead,
		resource: resource,
		response: responseChan,
	}

	su.pendingRequestsCond.L.Lock()
	su.pendingRequests = append(su.pendingRequests, req)
	su.pendingRequestsCond.L.Unlock()
	su.pendingRequestsCond.Signal()

	select {
	case response := <-responseChan:
		return response.data, response.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (su *StackUpdate) registerResource(ctx context.Context, logicalID, physicalResourceID, resourceType string, properties map[string]interface{}, options ...ResourceOption) (*ResourceData, error) {
	if resourceType == "" {
		return nil, errors.New("resourceType is required")
	}

	resource := &ResourceData{
		LogicalID:          logicalID,
		PhysicalResourceID: physicalResourceID,
		Type:               resourceType,
		Properties:         properties,
	}
	for _, option := range options {
		option(resource)
	}

	return su.handleRequest(ctx, resource, false)
}

func (su *StackUpdate) CreateResource(ctx context.Context, logicalID, resourceType string, properties map[string]interface{}, options ...ResourceOption) (*ResourceData, error) {
	if logicalID == "" {
		return nil, errors.New("logicalID is required")
	}
	return su.registerResource(ctx, logicalID, "", resourceType, properties, options...)
}

func (su *StackUpdate) UpdateResource(ctx context.Context, logicalID, physicalResourceID, resourceType string, properties map[string]interface{}, options ...ResourceOption) (*ResourceData, error) {
	if logicalID == "" {
		return nil, errors.New("logicalID is required")
	}
	if physicalResourceID == "" {
		return nil, errors.New("physicalResourceID is required")
	}
	return su.registerResource(ctx, logicalID, physicalResourceID, resourceType, properties, options...)
}

func (su *StackUpdate) DeleteResource(ctx context.Context, logicalID, physicalResourceID string) error {
	if logicalID == "" {
		return errors.New("logicalID is required")
	}
	if physicalResourceID == "" {
		return errors.New("physicalResourceID is required")
	}
	_, err := su.handleRequest(ctx, &ResourceData{LogicalID: logicalID, PhysicalResourceID: physicalResourceID}, false)
	return err
}

func (su *StackUpdate) ReadResource(ctx context.Context, logicalID, physicalResourceID, resourceType string) (*ResourceData, error) {
	if logicalID == "" {
		return nil, errors.New("logicalID is required")
	}
	if resourceType == "" {
		return nil, errors.New("resourceType is required")
	}
	return su.handleRequest(ctx, &ResourceData{
		LogicalID:          logicalID,
		PhysicalResourceID: physicalResourceID,
		Type:               resourceType,
	}, true)
}
