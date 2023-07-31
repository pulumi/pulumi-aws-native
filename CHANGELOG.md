## HEAD (Unreleased)

### Breaking changes:
- Resource "aws-native:omics:AnnotationStore" input "storeOptions" type changed from "#/types/aws-native:omics:AnnotationStoreStoreOptions" to "#/types/aws-native:omics:StoreOptionsProperties"
- Resource "aws-native:omics:AnnotationStore" output "storeOptions" type changed from "#/types/aws-native:omics:AnnotationStoreStoreOptions" to "#/types/aws-native:omics:StoreOptionsProperties"
- Resource "aws-native:appstream:Application" input "tags" items type changed from "#/types/aws-native:appstream:ApplicationTag" to "oneOf"
- Resource "aws-native:appstream:Application" output "tags" items type changed from "#/types/aws-native:appstream:ApplicationTag" to "oneOf"
- Resource "aws-native:appstream:AppBlock" input "tags" items type changed from "#/types/aws-native:appstream:AppBlockTag" to "oneOf"
- Resource "aws-native:appstream:AppBlock" output "tags" items type changed from "#/types/aws-native:appstream:AppBlockTag" to "oneOf"
- Resource "aws-native:mediatailor:PlaybackConfiguration" input "configurationAliases" type changed from "#/types/aws-native:mediatailor:PlaybackConfigurationConfigurationAliases" to "pulumi.json#/Any"
- Resource "aws-native:mediatailor:PlaybackConfiguration" output "configurationAliases" type changed from "#/types/aws-native:mediatailor:PlaybackConfigurationConfigurationAliases" to "pulumi.json#/Any"
- Resource "aws-native:m2:Application" input "definition" type changed from "#/types/aws-native:m2:ApplicationDefinition" to "oneOf"
- Resource "aws-native:m2:Application" output "definition" type changed from "#/types/aws-native:m2:ApplicationDefinition" to "oneOf"
- Resource "aws-native:nimblestudio:StudioComponent" input "configuration" type changed from "#/types/aws-native:nimblestudio:StudioComponentConfiguration" to "oneOf"
- Resource "aws-native:nimblestudio:StudioComponent" output "configuration" type changed from "#/types/aws-native:nimblestudio:StudioComponentConfiguration" to "oneOf"
- Resource "aws-native:verifiedpermissions:Policy" input "definition" type changed from "#/types/aws-native:verifiedpermissions:PolicyDefinition" to "oneOf"
- Resource "aws-native:verifiedpermissions:Policy" output "definition" type changed from "#/types/aws-native:verifiedpermissions:PolicyDefinition" to "oneOf"
- Function "aws-native:nimblestudio:getStudioComponent" output "configuration" type changed from "#/types/aws-native:nimblestudio:StudioComponentConfiguration" to "oneOf"
- Function "aws-native:verifiedpermissions:getPolicy" output "definition" type changed from "#/types/aws-native:verifiedpermissions:PolicyDefinition" to "oneOf"
- Function "aws-native:mediatailor:getPlaybackConfiguration" output "configurationAliases" type changed from "#/types/aws-native:mediatailor:PlaybackConfigurationConfigurationAliases" to "pulumi.json#/Any"
- Type "aws-native:securityhub:AutomationRuleNoteUpdate" input "updatedBy" type changed from "#/types/aws-native:securityhub:AutomationRulearnOrId" to "string"
- Type "aws-native:s3outposts:BucketRuleFilterProperties" input "andOperator" type changed from "#/types/aws-native:s3outposts:BucketFilterAndOperator" to "#/types/aws-native:s3outposts:FilterAndOperatorProperties"
- Type "aws-native:pipes:PipeSourceManagedStreamingKafkaParameters" input "credentials" type changed from "#/types/aws-native:pipes:PipeMSKAccessCredentials" to "oneOf"
- Type "aws-native:pipes:PipeSourceActiveMQBrokerParameters" input "credentials" type changed from "#/types/aws-native:pipes:PipeMQBrokerAccessCredentials" to "#/types/aws-native:pipes:MQBrokerAccessCredentialsProperties"
- Type "aws-native:pipes:PipeSourceSelfManagedKafkaParameters" input "credentials" type changed from "#/types/aws-native:pipes:PipeSelfManagedKafkaAccessConfigurationCredentials" to "oneOf"
- Type "aws-native:cleanrooms:ConfiguredTableAnalysisRulePolicy" input "v1" type changed from "#/types/aws-native:cleanrooms:ConfiguredTableAnalysisRulePolicyV1" to "oneOf"
- Type "aws-native:rolesanywhere:TrustAnchorSource" input "sourceData" type changed from "#/types/aws-native:rolesanywhere:TrustAnchorSourceData" to "oneOf"
- Type "aws-native:pipes:PipeSourceRabbitMQBrokerParameters" input "credentials" type changed from "#/types/aws-native:pipes:PipeMQBrokerAccessCredentials" to "#/types/aws-native:pipes:MQBrokerAccessCredentialsProperties"
- Type "aws-native:securityhub:AutomationRuleRelatedFinding" input "id" type changed from "#/types/aws-native:securityhub:AutomationRulearnOrId" to "string"
- Type "aws-native:amplifyuibuilder:FormStyle" input "verticalGap" type changed from "#/types/aws-native:amplifyuibuilder:FormStyleConfig" to "oneOf"
- Type "aws-native:amplifyuibuilder:FormStyle" input "horizontalGap" type changed from "#/types/aws-native:amplifyuibuilder:FormStyleConfig" to "oneOf"
- Type "aws-native:amplifyuibuilder:FormStyle" input "outerPadding" type changed from "#/types/aws-native:amplifyuibuilder:FormStyleConfig" to "oneOf"
- Type "aws-native:amplifyuibuilder:FormButton" input "position" type changed from "#/types/aws-native:amplifyuibuilder:FormFieldPosition" to "oneOf"

## 0.70.0

### Breaking changes:
- Resource "aws-native:cloudwatch:MetricStream" output "creationDate" type changed from "Any" to "string"
- Resource "aws-native:cloudwatch:MetricStream" output "lastUpdateDate" type changed from "Any" to "string"
- Function "aws-native:cloudwatch:getMetricStream" output "lastUpdateDate" type changed from "Any" to "string"
- Function "aws-native:cloudwatch:getMetricStream" output "creationDate" type changed from "Any" to "string"

## 0.69.0 (July 19, 2023)

### Breaking changes:
- Function "aws-native:s3:getBucket" missing output "accessControl"
- Function "aws-native:cloudformation:getStackSet" missing output "stackInstancesGroup"
- Function "aws-native:applicationinsights:getApplication" missing output "customComponents"
- Function "aws-native:applicationinsights:getApplication" missing output "groupingType"
- Function "aws-native:applicationinsights:getApplication" missing output "logPatternSets"
- Function "aws-native:applicationinsights:getApplication" missing output "opsItemSNSTopicArn"
- Function "aws-native:applicationinsights:getApplication" missing output "componentMonitoringSettings"
- Type "aws-native:ecs:TaskSetLoadBalancer" missing property "loadBalancerName"

#### New resources:

- `connect.Queue`
- `connect.RoutingProfile`
- `ec2.EIPAssociation`
- `iam.GroupPolicy`
- `iam.RolePolicy`
- `iam.UserPolicy`

#### New functions:

- `connect.getQueue`
- `connect.getRoutingProfile`
- `ec2.getEIPAssociation`
- `iam.getGroupPolicy`
- `iam.getRolePolicy`
- `iam.getUserPolicy`


## 0.68.0 (July 7, 2023)

2 Breaking Changes:
- Function "aws-native:ec2:getKeyPair" missing output "tags"
- Function "aws-native:ec2:getKeyPair" missing output "publicKeyMaterial"

New resources:
- apigatewayv2.ApiMapping
- comprehend.DocumentClassifier
- sns.TopicPolicy
New functions:
- apigatewayv2.getApiMapping
- comprehend.getDocumentClassifier
- sns.getTopicPolicy

- Upgrade AWS SDK dependencies to fix SSO issues (#954)
- PR #952 fixes an issue in schema generation where separate types would occasionally have the same
  name and would overwrite each other's metadata. This caused the following renames:
  - Resource `aws-native:eks:Cluster` input `logging` type changed from `#/types/aws-native:eks:ClusterLogging` to `#/types/aws-native:eks:Logging`
  - Resource `aws-native:eks:Cluster` output `logging` type changed from `#/types/aws-native:eks:ClusterLogging` to `#/types/aws-native:eks:Logging`
  - Function `aws-native:eks:getCluster` output `logging` type changed from `#/types/aws-native:eks:ClusterLogging` to `#/types/aws-native:eks:Logging`
  - Type `aws-native:ec2:LaunchTemplateData` input `tagSpecifications` items type changed from `#/types/aws-native:ec2:LaunchTemplateTagSpecification` to `#/types/aws-native:ec2:TagSpecification`
  - Type `aws-native:iotanalytics:DatastorePartition` input `partition` type changed from `#/types/aws-native:iotanalytics:DatastorePartition` to `#/types/aws-native:iotanalytics:Partition`
  - Type `aws-native:eks:ClusterLogging` was removed

## 0.67.0 (June 27, 2023)
### Does the PR have any schema changes?

Found 40 breaking changes:
- Resource aws-native:sagemaker:ModelPackage removed input "createdBy"
- Resource "aws-native:sagemaker:ModelPackage" removed input "modelPackageStatusItem"
- Resource "aws-native:sagemaker:ModelPackage" removed input "lastModifiedBy"
- Resource "aws-native:sagemaker:ModelPackage" removed input "additionalInferenceSpecificationDefinition"
- Resource "aws-native:sagemaker:ModelPackage" removed input "environment"
- Resource "aws-native:sagemaker:ModelPackage" removed output "createdBy"
- Resource "aws-native:sagemaker:ModelPackage" removed output "modelPackageStatusItem"
- Resource "aws-native:sagemaker:ModelPackage" removed output "additionalInferenceSpecificationDefinition"
- Resource "aws-native:sagemaker:ModelPackage" removed output "environment"
- Resource "aws-native:sagemaker:ModelPackage" removed output "lastModifiedBy"
- Function "aws-native:iam:getServiceLinkedRole" removed input "id"
- Function "aws-native:iam:getServiceLinkedRole" removed output "id"
- Function "aws-native:ec2:getVolumeAttachment" removed
- Function "aws-native:appstream:getApplication" removed output "attributesToDelete"
- Function "aws-native:sagemaker:getModelPackage" removed output "modelPackageStatusItem"
- Function "aws-native:sagemaker:getModelPackage" removed output "additionalInferenceSpecificationDefinition"
- Function "aws-native:sagemaker:getModelPackage" removed output "environment"
- Function "aws-native:sagemaker:getModelPackage" removed output "lastModifiedBy"
- Function "aws-native:sagemaker:getModelPackage" removed output "additionalInferenceSpecificationsToAdd"
- Function "aws-native:sagemaker:getModelPackage" removed output "createdBy"
- Type "aws-native:sagemaker:ModelPackageStatusDetails" removed property "imageScanStatuses"
- Type "aws-native:sagemaker:ModelPackageContainerDefinition" removed property "productId"
- Type "aws-native:sagemaker:ModelPackageCreatedBy" removed
- Type "aws-native:quicksight:DashboardColumnConfiguration" removed property "colorsConfiguration"
- Type "aws-native:appstream:AppBlockTag" removed property "tagValue"
- Type "aws-native:appstream:AppBlockTag" removed property "tagKey"
- Type "aws-native:quicksight:DashboardSpecialValue" removed
- Type "aws-native:sagemaker:ModelPackageLastModifiedBy" removed
- Type "aws-native:quicksight:TemplateSpecialValue" removed
- Type "aws-native:appstream:ApplicationTag" removed property "tagKey"
- Type "aws-native:appstream:ApplicationTag" removed property "tagValue"
- Type "aws-native:quicksight:AnalysisColumnConfiguration" removed property "colorsConfiguration"
- Type "aws-native:quicksight:TemplateColumnConfiguration" removed property "colorsConfiguration"
- Type "aws-native:quicksight:DashboardCustomColor" removed
- Type "aws-native:quicksight:TemplateCustomColor" removed
- Type "aws-native:quicksight:AnalysisSpecialValue" removed
- Type "aws-native:quicksight:AnalysisColorsConfiguration" removed
- Type "aws-native:quicksight:DashboardColorsConfiguration" removed
- Type "aws-native:quicksight:AnalysisCustomColor" removed
- Type "aws-native:quicksight:TemplateColorsConfiguration" removed

#### New resources:

- `apprunner.AutoScalingConfiguration`
- `appstream.AppBlockBuilder`
- `ec2.LaunchTemplate`
- `organizations.Organization`
- `verifiedpermissions.IdentitySource`
- `verifiedpermissions.Policy`
- `verifiedpermissions.PolicyStore`
- `verifiedpermissions.PolicyTemplate`

#### New functions:

- `apprunner.getAutoScalingConfiguration`
- `appstream.getAppBlockBuilder`
- `ec2.getLaunchTemplate`
- `organizations.getOrganization`
- `verifiedpermissions.getIdentitySource`
- `verifiedpermissions.getPolicy`
- `verifiedpermissions.getPolicyStore`
- `verifiedpermissions.getPolicyTemplate`

## 0.66.0 (June 16, 2023)

- Panorama ApplicationInstance `deviceId` and `statusFilter` fields removed in upstream spec (type definitions are still there but unused)
- Panorama PackageVersion `updatedLatestPatchVersion` is now a write-only property so not outputted
- Lambda LayerVersion `id` replaced with `layerVersionArn`
- SSM Document `updateMethod` and `attachments` now write-only
- go to uses `errors` rather than `github.com/pkg/errors`
- python requires `>=3.7`

## 0.65.0 (June 7, 2023)

#### Upstream breaking changes:
- Resource "aws-native:ec2:IPAM" removed input `defaultResourceDiscoveryAssociationId`
- Resource "aws-native:ec2:IPAM" removed input `defaultResourceDiscoveryId`
- Function "aws-native:applicationautoscaling:getScalableTarget" removed input `id`
- Function "aws-native:applicationautoscaling:getScalableTarget" removed output `roleARN`

#### New resources:
- `athena.CapacityReservation`
- `iam.ServiceLinkedRole`
- `lambda.LayerVersion`
- `lambda.LayerVersionPermission`

#### New functions:
- `athena.getCapacityReservation`
- `iam.getServiceLinkedRole`
- `lambda.getLayerVersion`
- `lambda.getLayerVersionPermission`

## 0.64.0 (June 1, 2023)

#### Upstream breaking changes:
- Resource `aws-native:apigatewayv2:RouteResponse` input `responseParameters` type changed from `pulumi.json#/Any` to `#/types/aws-native:apigatewayv2:RouteResponseRouteParameters`
- Resource `aws-native:apigatewayv2:RouteResponse` output `responseParameters` type changed from `pulumi.json#/Any` to `#/types/aws-native:apigatewayv2:RouteResponseRouteParameters`
- Function `aws-native:apigatewayv2:getRouteResponse` removed input `id`
- Function `aws-native:apigatewayv2:getRouteResponse` removed output `id`
- Function `aws-native:apigatewayv2:getRouteResponse` output `responseParameters` type changed from `pulumi.json#/Any` to `#/types/aws-native:apigatewayv2:RouteResponseRouteParameters`

#### New resources:

- `cognito.IdentityPoolPrincipalTag`
- `detective.OrganizationAdmin`

#### New functions:

- `cognito.getIdentityPoolPrincipalTag`
- `detective.getOrganizationAdmin`
- `ses.getDedicatedIpPool`

## 0.63.0 (May 25, 2023)

#### Upstream breaking changes:

- Function `aws-native:sagemaker:getApp` removed output `resourceSpec`

#### New resources:

- `appsync.SourceApiAssociation`
- `shield.DRTAccess`
- `shield.ProactiveEngagement`
- `shield.Protection`
- `shield.ProtectionGroup`

#### New functions:

- `appsync.getSourceApiAssociation`
- `shield.getDRTAccess`
- `shield.getProactiveEngagement`
- `shield.getProtection`
- `shield.getProtectionGroup`

## 0.62.0 (May 18, 2023)

Upstream breaking changes:

- Resource "aws-native:ec2:IPAM" removed input "resourceDiscoveryAssociationCount"
- Function "aws-native:simspaceweaver:getSimulation" removed output "roleArn"
- Function "aws-native:simspaceweaver:getSimulation" removed output "schemaS3Location"
- Function "aws-native:rds:getDBInstance" removed output "useDefaultProcessorFeatures"
- Function "aws-native:synthetics:getCanary" removed output "visualReference"
- Function "aws-native:synthetics:getCanary" removed output "deleteLambdaResourcesOnCanaryDeletion"
- Function "aws-native:synthetics:getCanary" removed output "startCanaryAfterCreation"
- Type "aws-native:appflow:FlowTriggerConfig" removed property "activateFlowOnCreate"

#### New resources:

- `connect.Prompt`
- `ec2.SubnetCidrBlock`
- `quicksight.Topic`

#### New functions:

- `connect.getPrompt`
- `ec2.getSubnetCidrBlock`
- `quicksight.getTopic`

## 0.61.0 (May 9, 2023)
Upstream breaking changes:

- `s3objectlambda.AccessPoint` `AliasProperties` and `PolicyStatusProperties` changed.
- `lambda.EventInvokeConfig` changed identifier.
- `lightsail.Instance` `userData` now write-only.
- `secretsmanager.Secret` `generateSecretString` and `secretString` are now write-only.
- `servicecatalogappregistry` resources have additional properties marked as create-only.
- `kinesisfirehose.DeliveryStream` removed properties.

## 0.60.0 (May 1, 2023)

 Upstream breaking changes:
 - `PolicyStatus` was removed from S3 `AccessPoint` upstream, resulting in
   - Resource "aws-native:s3:AccessPoint" missing input `policyStatus`
   - Resource "aws-native:s3:AccessPoint" missing output `policyStatus`
   - Function "aws-native:s3:getAccessPoint" missing output `policyStatus`
   - Type "aws-native:s3:AccessPointPolicyStatusPropertiesIsPublic" missing

 New resources:
 - `apigatewayv2.IntegrationResponse`
 - `connect.EvaluationForm`
 - `datasync.StorageSystem`
 - `ec2.VPCEndpointServicePermissions`
 - `lambda.EventInvokeConfig`
 - `msk.ClusterPolicy`
 - `msk.VpcConnection`
 - `quicksight.VPCConnection`

 New functions:
 - `apigatewayv2.getIntegrationResponse`
 - `connect.getEvaluationForm`
 - `datasync.getStorageSystem`
 - `ec2.getVPCEndpointServicePermissions`
 - `lambda.getEventInvokeConfig`
 - `msk.getClusterPolicy`
 - `msk.getVpcConnection`
 - `quicksight.getVPCConnection`

## 0.58.0 (April 14, 2023)

Upstream breaking changes:
- Function "aws-native:licensemanager:getGrant" missing output "status"
  - This is due to upstream CCAPI specifications now marking `status` for `licencemanager` as `WriteOnly`

#### New resources:

- `appconfig.Extension`
- `appconfig.ExtensionAssociation`
- `quicksight.RefreshSchedule`

#### New functions:

- `appconfig.getExtension`
- `appconfig.getExtensionAssociation`

## 0.57.0 (April 7, 2023)
### Does the PR have any schema changes?

Upstream breaking changes:
- Resource "aws-native:macie:FindingsFilter" missing output "findingsFilterListItems"
- Function "aws-native:macie:getFindingsFilter" missing output "findingsFilterListItems"
- Function "aws-native:gamelift:getGameServerGroup" missing output "maxSize"
- Function "aws-native:gamelift:getGameServerGroup" missing output "tags"
- Function "aws-native:gamelift:getGameServerGroup" missing output "autoScalingPolicy"
- Function "aws-native:gamelift:getGameServerGroup" missing output "launchTemplate"
- Function "aws-native:gamelift:getGameServerGroup" missing output "minSize"
- Function "aws-native:gamelift:getGameServerGroup" missing output "vpcSubnets"
- Type "aws-native:macie:FindingsFilterListItem" missing

#### New resources:

- `neptune.DBCluster`
- `ssmcontacts.Plan`
- `ssmcontacts.Rotation`

#### New functions:

- `neptune.getDBCluster`
- `ssmcontacts.getPlan`
- `ssmcontacts.getRotation`

## 0.56.0 (April 3, 2023)

New resources:
- `devopsguru.LogAnomalyDetectionIntegration`

New functions:
- `devopsguru.getLogAnomalyDetectionIntegration`
- `logs.getSubscriptionFilter`

Enhancements:

- Max retries are set to 25 by default, max rate limit disabled by default
  [#862](https://github.com/pulumi/pulumi-aws-native/pull/862)


## 0.55.0 (March 28, 2023)

Upstream breaking changes
  - Function `aws-native:iot:getCACertificate` missing output `removeAutoRegistration`

New resources:
  - `apigatewayv2.Route`


## 0.54.0 (March 21, 2023)

Upstream breaking changes
  - Resource `aws-native:quicksight:Analysis` removed input `errors`
  - Function `aws-native:athena:getWorkGroup` removed output `recursiveDeleteOption`

New resources:
  - `chatbot.MicrosoftTeamsChannelConfiguration`
  - `comprehend.Flywheel`
  - `ec2.VPCEndpoint`
  - `sagemaker.InferenceExperiment`


## 0.53.0 (March 14, 2023)

Bug Fixes:
  - Adds semantics around FIFO queue naming conventions as expected by AWS (#794)

Upstream breaking changes
- Resource `aws-native:appintegrations:EventIntegration` removed output `associations`.
- Resource `aws-native:organizations:Policy` `content` type changed from `string` to `pulumi.json#/Any`
- Function `aws-native:organizations:getPolicy` output `content` type changed from `string` to `pulumi.json#/Any`
- Function `aws-native:detective:getMemberInvitation` removed outputs `disableEmailNotification` and `message`. These are write-only properties.
- Function `aws-native:appintegrations:getEventIntegration` removed output `associations`.
- Function `aws-native:ec2:getVPCDHCPOptionsAssociation` removed. The Id properter was removed so the only remaining properties are the primary identifier which are the inputs to the function, so there's no output remaining, so it no longer makes sense to fetch this resource.

## 0.52.0 (March 1, 2023)

Breaking Changes:
  - Resource "aws-native:apigatewayv2:RouteResponse" missing output "routeResponseId"
  - Resource "aws-native:route53recoverycontrol:Cluster" missing input "clusterEndpoints"
  - Function "aws-native:rds:getDBInstance" missing output "deleteAutomatedBackups"
  - Function "aws-native:rds:getDBInstance" missing output "allowMajorVersionUpgrade"
  - Function "aws-native:redshift:getClusterParameterGroup" missing output "parameterGroupName"
  - Function "aws-native:apigatewayv2:getRouteResponse" missing input "apiId"
  - Function "aws-native:apigatewayv2:getRouteResponse" missing input "routeId"
  - Function "aws-native:apigatewayv2:getRouteResponse" missing input "routeResponseId"
  - Function "aws-native:apigatewayv2:getRouteResponse" missing output "routeResponseId"
  - Function "aws-native:rds:getDBProxyTargetGroup" missing input "dBProxyName"
  - Function "aws-native:rds:getDBProxyTargetGroup" missing output "dBProxyName"
  - Type "aws-native:iottwinmaker:ComponentTypeStatus" input "error" type changed from "#/types/aws-native:iottwinmaker:ComponentTypeStatusErrorProperties" to ""
  - Type "aws-native:connectcampaigns:CampaignOutboundCallConfigAnswerMachineDetectionConfigProperties" missing
  - Type "aws-native:iottwinmaker:EntityStatus" input "error" type changed from "#/types/aws-native:iottwinmaker:EntityStatusErrorProperties" to ""
  - Type "aws-native:connectcampaigns:CampaignOutboundCallConfig" input "answerMachineDetectionConfig" type changed from "#/types/aws-native:connectcampaigns:CampaignOutboundCallConfigAnswerMachineDetectionConfigProperties" to "#/types/aws-native:connectcampaigns:CampaignAnswerMachineDetectionConfig"

New resources:

    apigateway.VpcLink
    ec2.VPCEndpointService
    internetmonitor.Monitor
    ivschat.LoggingConfiguration
    ivschat.Room
    networkmanager.TransitGatewayRouteTableAttachment

New functions:

    apigateway.getVpcLink
    ec2.getVPCEndpointService
    internetmonitor.getMonitor
    ivschat.getLoggingConfiguration
    ivschat.getRoom
    networkmanager.getTransitGatewayRouteTableAttachment

## 0.51.0 (February 16, 2023)

Breaking Changes:
  - Function "aws-native:applicationautoscaling:getScalableTarget" missing input "resourceId"
  - Function "aws-native:applicationautoscaling:getScalableTarget" missing input "scalableDimension"
  - Function "aws-native:applicationautoscaling:getScalableTarget" missing input "serviceNamespace"

New resources:

    apigatewayv2.RouteResponse
    ec2.LocalGatewayRouteTable
    ec2.LocalGatewayRouteTableVirtualInterfaceGroupAssociation
    fms.ResourceSet
    networkmanager.TransitGatewayPeering
    organizations.ResourcePolicy
    systemsmanagersap.Application

New functions:

    apigatewayv2.getRouteResponse
    ec2.getLocalGatewayRouteTable
    ec2.getLocalGatewayRouteTableVirtualInterfaceGroupAssociation
    fms.getResourceSet
    networkmanager.getTransitGatewayPeering
    organizations.getResourcePolicy
    systemsmanagersap.getApplication

## 0.50.0 (February 09, 2023)

Breaking changes:
  - Resource "aws-native:rds:DBProxyEndpoint" missing input "targetRole"
  - Resource "aws-native:redshift:EndpointAccess" missing input "vpcSecurityGroups"
  - Resource "aws-native:redshift:EndpointAccess" missing input "vpcEndpoint"
  - Resource "aws-native:amplifyuibuilder:Theme" missing output "createdAt"
  - Resource "aws-native:amplifyuibuilder:Theme" missing output "modifiedAt"
  - Resource "aws-native:redshiftserverless:Workgroup" missing input "workgroup"
  - Resource "aws-native:redshiftserverless:Namespace" missing input "namespace"
  - Function "aws-native:opsworkscm:getServer" missing output "id"
  - Function "aws-native:opsworkscm:getServer" missing output "tags"
  - Function "aws-native:amplifyuibuilder:getTheme" missing output "createdAt"
  - Function "aws-native:amplifyuibuilder:getTheme" missing output "modifiedAt"
  - Function "aws-native:ssmcontacts:getContactChannel" missing output "deferActivation"
  - Function "aws-native:rds:getDBProxyTargetGroup" missing input "targetGroupArn"
  - Type "aws-native:wafv2:RuleGroupRuleActionBlockProperties" missing
  - Type "aws-native:rds:DBProxyAuthFormat" missing property "userName"
  - Type "aws-native:wafv2:RuleGroupRuleActionChallengeProperties" missing
  - Type "aws-native:wafv2:RuleGroupRuleAction" input "allow" type changed from "#/types/aws-native:wafv2:RuleGroupRuleActionAllowProperties" to "#/types/aws-native:wafv2:RuleGroupAllowAction"
  - Type "aws-native:wafv2:RuleGroupRuleAction" input "block" type changed from "#/types/aws-native:wafv2:RuleGroupRuleActionBlockProperties" to "#/types/aws-native:wafv2:RuleGroupBlockAction"
  - Type "aws-native:wafv2:RuleGroupRuleAction" input "captcha" type changed from "#/types/aws-native:wafv2:RuleGroupRuleActionCaptchaProperties" to "#/types/aws-native:wafv2:RuleGroupCaptchaAction"
  - Type "aws-native:wafv2:RuleGroupRuleAction" input "challenge" type changed from "#/types/aws-native:wafv2:RuleGroupRuleActionChallengeProperties" to "#/types/aws-native:wafv2:RuleGroupChallengeAction"
  - Type "aws-native:wafv2:RuleGroupRuleAction" input "count" type changed from "#/types/aws-native:wafv2:RuleGroupRuleActionCountProperties" to "#/types/aws-native:wafv2:RuleGroupCountAction"

  New resources:
  - `cloudtrail.Channel`
  - `cloudtrail.ResourcePolicy`
  - `ec2.IPAMPoolCidr`
  - `ec2.IPAMResourceDiscovery`
  - `ec2.IPAMResourceDiscoveryAssociation`
  - `omics.AnnotationStore`
  - `omics.ReferenceStore`
  - `omics.RunGroup`
  - `omics.SequenceStore`
  - `omics.VariantStore`
  - `omics.Workflow`
  - `sagemaker.ModelCard`
  - `sagemaker.Space`
  - `simspaceweaver.Simulation`

New functions:
  - `cloudtrail.getChannel`
  - `cloudtrail.getResourcePolicy`
  - `ec2.getIPAMPoolCidr`
  - `ec2.getIPAMResourceDiscovery`
  - `ec2.getIPAMResourceDiscoveryAssociation`
  - `omics.getAnnotationStore`
  - `omics.getReferenceStore`
  - `omics.getRunGroup`
  - `omics.getSequenceStore`
  - `omics.getVariantStore`
  - `omics.getWorkflow`
  - `sagemaker.getModelCard`
  - `sagemaker.getSpace`
  - `simspaceweaver.getSimulation`


## 0.49.0 (January 26, 2023)

Breaking changes:
- Function `aws-native:robomaker:getRobotApplication` missing output `sources`
- Function `aws-native:robomaker:getSimulationApplication` missing output `renderingEngine`
- Function `aws-native:robomaker:getSimulationApplication` missing output `sources`

New resources:
- `applicationautoscaling.ScalableTarget`
- `connect.ApprovedOrigin`
- `connect.IntegrationAssociation`
- `connect.SecurityKey`
- `ec2.VPNConnectionRoute`

New functions:
- `applicationautoscaling.getScalableTarget`
- `connect.getIntegrationAssociation`
- `connect.getSecurityKey`

## 0.48.0 (January 20, 2023)

Breaking changes:
  - Resource "aws-native:ec2:PlacementGroup" missing input "groupName"
  - Resource "aws-native:cloudwatch:MetricStream" output "creationDate" type changed from "string" to "any"
  - Resource "aws-native:cloudwatch:MetricStream" output "lastUpdateDate" type changed from "string" to "any"
  - Function "aws-native:cloudwatch:getMetricStream" output "lastUpdateDate" type changed from "string" to "any"
  - Function "aws-native:cloudwatch:getMetricStream" output "creationDate" type changed from "string" to "any"
  - Type "aws-native:transfer:WorkflowStepCopyStepDetailsProperties" input "destinationFileLocation" type changed from "#/types/aws-native:transfer:WorkflowInputFileLocation" to "#/types/aws-native:transfer:WorkflowS3FileLocation"

New resources:
  - `gamecast.Application`
  - `kendraranking.ExecutionPlan`
  - `secretsmanager.Secret`

New functions:
  - `gamecast.getApplication`
  - `kendraranking.getExecutionPlan`
  - `secretsmanager.getSecret`


## 0.47.0 (January 13, 2023)

Breaking changes:
  - Renamed broken type ClusterLogging to ClusterLoggingOuter, allowing the real ClusterLogging to be generated (#644)
  - Resource `aws-native:apigateway:RestApi` missing output `restApiId`
  - Resource `aws-native:datasync:Task` missing output `errorDetail`
  - Resource `aws-native:datasync:Task` missing output `errorCode`
  - Function `aws-native:apigateway:getRestApi` missing input `id`
  - Function `aws-native:apigateway:getRestApi` missing output `body`
  - Function `aws-native:apigateway:getRestApi` missing output `parameters`
  - Function `aws-native:apigateway:getRestApi` missing output `mode`
  - Function `aws-native:apigateway:getRestApi` missing output `cloneFrom`
  - Function `aws-native:apigateway:getRestApi` missing output `id`
  - Function `aws-native:apigateway:getRestApi` missing output `bodyS3Location`
  - Function `aws-native:apigateway:getRestApi` missing output `failOnWarnings`
  - Function `aws-native:apigateway:getRestApi` missing input `restApiId`
  - Function `aws-native:apigateway:getRestApi` missing output `restApiId`
  - Function `aws-native:datasync:getTask` missing output `errorCode`
  - Function `aws-native:datasync:getTask` missing output `errorDetail`
  - Function `aws-native:ec2:getPlacementGroup` missing output `groupName`
  - Function `aws-native:ec2:getVolume` missing input `id`
  - Function `aws-native:ec2:getVolume` missing output `id`
  - Function `aws-native:iotsitewise:getProject` missing output `assetIds`
  - Function `aws-native:logs:getSubscriptionFilter` missing
  - Function `aws-native:redshift:getClusterSubnetGroup` missing output `tags`
  - Function `aws-native:redshift:getEventSubscription` missing output `tags`
  - Function `aws-native:redshift:getClusterParameterGroup` missing output `tags`
  - Function `aws-native:redshiftserverless:getNamespace` missing output `finalSnapshotName`
  - Function `aws-native:redshiftserverless:getNamespace` missing output `finalSnapshotRetentionPeriod`

New resources:
  - `appflow.Connector`
  - `directoryservice.SimpleAD`
  - `gamecast.StreamGroup`
  - `gamelift.Build`

New functions:
  - `appflow.getConnector`
  - `directoryservice.getSimpleAD`
  - `gamecast.getStreamGroup`
  - `gamelift.getBuild`


## 0.46.0 (January 09, 2023)

Breaking changes:
- Resource "aws-native:apigateway:RestApi" missing output "restApiId"
- Function "aws-native:redshift:getClusterSubnetGroup" missing output "tags"
- Function "aws-native:ec2:getPlacementGroup" missing output "groupName"
- Function "aws-native:redshiftserverless:getNamespace" missing output "finalSnapshotName"
- Function "aws-native:redshiftserverless:getNamespace" missing output "finalSnapshotRetentionPeriod"
- Function "aws-native:iotsitewise:getProject" missing output "assetIds"
- Function "aws-native:logs:getSubscriptionFilter" missing
- Function "aws-native:apigateway:getRestApi" missing input "restApiId"
- Function "aws-native:apigateway:getRestApi" missing output "restApiId"
- Function "aws-native:redshift:getEventSubscription" missing output "tags"
- Function "aws-native:ec2:getVolume" missing input "id"
- Function "aws-native:ec2:getVolume" missing output "id"
- Function "aws-native:redshift:getClusterParameterGroup" missing output "tags"

New resources:
- `appflow.Connector`
- `directoryservice.SimpleAD`
- `gamelift.Build`

New functions:
- `appflow.getConnector`
- `directoryservice.getSimpleAD`
- `gamelift.getBuild`


## 0.45.0 (December 16, 2022)

Breaking changes:
- Resource "aws-native:backup:Framework" output "creationTime" type changed from "number" to "string"
- Function "aws-native:configuration:getAggregationAuthorization" missing input "aggregationAuthorizationArn"
- Function "aws-native:elasticloadbalancingv2:getTargetGroup" missing input "id"
- Function "aws-native:elasticloadbalancingv2:getTargetGroup" missing output "id"
- Function "aws-native:ecs:getCluster" missing output "serviceConnectDefaults"
- Function "aws-native:backup:getFramework" output "creationTime" type changed from "number" to "string"
- Type "aws-native:backup:BackupVaultLockConfigurationType" input "maxRetentionDays" type changed from "number" to "integer"
- Type "aws-native:backup:BackupVaultLockConfigurationType" input "minRetentionDays" type changed from "number" to "integer"
- Type "aws-native:backup:BackupVaultLockConfigurationType" input "changeableForDays" type changed from "number" to "integer"

New resources:

- `codepipeline.CustomActionType`
- `docdbelastic.Cluster`
- `elasticbeanstalk.ConfigurationTemplate`
- `grafana.Workspace`
- `iot.Thing`
- `iottwinmaker.SyncJob`
- `opensearchserverless.AccessPolicy`
- `opensearchserverless.Collection`
- `opensearchserverless.SecurityConfig`
- `opensearchserverless.SecurityPolicy`
- `opensearchserverless.VpcEndpoint`

New functions:

- `codepipeline.getCustomActionType`
- `docdbelastic.getCluster`
- `elasticbeanstalk.getConfigurationTemplate`
- `grafana.getWorkspace`
- `iot.getThing`
- `iottwinmaker.getSyncJob`
- `opensearchserverless.getAccessPolicy`
- `opensearchserverless.getCollection`
- `opensearchserverless.getSecurityConfig`
- `opensearchserverless.getSecurityPolicy`
- `opensearchserverless.getVpcEndpoint`

## 0.44.0 (December 7, 2022)

Breaking changes:

- Resource "aws-native:ec2:Volume" removed output "volumeId"
- Resource "aws-native:elasticloadbalancingv2:TargetGroup" removed output "targetGroupArn"
- Function "aws-native:batch:getComputeEnvironment" removed output "updatePolicy"
- Function "aws-native:elasticloadbalancingv2:getTargetGroup" removed output "targetGroupArn"
- Function "aws-native:ec2:getVolume" removed input "volumeId"
- Function "aws-native:ec2:getVolume" removed output "volumeId"

New resources:

- `apigateway.RestApi`
- `connect.Rule`
- `ec2.NetworkPerformanceMetricSubscription`
- `pipes.Pipe`

New functions:

- `apigateway.getRestApi`
- `connect.getRule`
- `pipes.getPipe`

## 0.43.0 (November 30, 2022)

- Add Lambda Functions `SnapStart`

Breaking changes:

- Resource `aws-native:elasticloadbalancingv2:getTargetGroup` no longer has input `targetGroupArn`

New resources:

- `gamelift.Location`
- `oam.Link`
- `oam.Sink`


## 0.42.0 (November 23, 2022)

Breaking Changes:

- Resource `aws-native:iot:JobTemplate` no longer has `jobExecutionsRetryConfig`
- Resource `aws-native:sagemaker:ModelPackage` no longer has in- and output "tag"
- Function `aws-native:apigatewayv2:getAuthorizer` no longer has in- and output "id"

New resources:

- `cloudfront.ContinuousDeploymentPolicy`
- `elasticloadbalancingv2.TargetGroup`

New functions:

- `cloudfront.getContinuousDeploymentPolicy`
- `elasticloadbalancingv2.getTargetGroup`

## 0.41.0 (November 17, 2022)

Breaking Changes:

- Function `aws-native:apigateway:getBasePathMapping` no longer has output "id"

New resources:

- `amplifyuibuilder.Form`
- `apigatewayv2.Authorizer`
- `elasticbeanstalk.Environment`
- `organizations.Account`
- `organizations.OrganizationalUnit`
- `organizations.Policy`
- `resourceexplorer2.DefaultViewAssociation`
- `resourceexplorer2.Index`
- `resourceexplorer2.View`
- `scheduler.Schedule`
- `scheduler.ScheduleGroup`
- `ses.VdmAttributes`
- `ssm.ResourcePolicy`
- `xray.ResourcePolicy`

New functions:

- `amplifyuibuilder.getForm`
- `apigatewayv2.getAuthorizer`
- `elasticbeanstalk.getEnvironment`
- `organizations.getAccount`
- `organizations.getOrganizationalUnit`
- `organizations.getPolicy`
- `resourceexplorer2.getDefaultViewAssociation`
- `resourceexplorer2.getIndex`
- `resourceexplorer2.getView`
- `sagemaker.getFeatureGroup`
- `scheduler.getSchedule`
- `scheduler.getScheduleGroup`
- `ses.getVdmAttributes`
- `ssm.getResourcePolicy`
- `xray.getResourcePolicy`

## 0.40.2 (November 3, 2022)
- Republish due to `nuget push` silently failing

## 0.40.1 (November 3, 2022)
- Fix reading resources which have been deleted in the cloud [#682](https://github.com/pulumi/pulumi-aws-native/pull/682)

## 0.40.0 (November 3, 2022)

- Fix updates to write-only properties [#678](https://github.com/pulumi/pulumi-aws-native/pull/678).
- Breaking changes to some RDS get functions due to upstream spec changes ([see #679](https://github.com/pulumi/pulumi-aws-native/pull/679))

#### New resources:

- `apigatewayv2.Api`
- `apigatewayv2.Deployment`
- `apprunner.VpcIngressConnection`
- `ec2.Volume`
- `fsx.DataRepositoryAssociation`
- `supportapp.SlackWorkspaceConfiguration`

#### New functions:

- `apigatewayv2.getApi`
- `apigatewayv2.getDeployment`
- `apprunner.getVpcIngressConnection`
- `ec2.getVolume`
- `fsx.getDataRepositoryAssociation`

## 0.39.0 (October 24, 2022)

- Update to Pulumi SDK v3.43.1
- Update to include the latest resource definitions

#### New resources:

- `aws-native:autoscaling:ScheduledAction`
- `aws-native:datapipeline:Pipeline`

#### New functions:

- `aws-native:autoscaling:getScheduledAction`
- `aws-native:datapipeline:getPipeline`

## 0.38.0 (October 17, 2022)

- Update to Pulumi SDK v3.42.0
- Update to include the latest resource definitions

#### New resources:

- aws-native:apigatewayv2:Model
- aws-native:cloudfront:MonitoringSubscription
- aws-native:codedeploy:Application
- aws-native:codedeploy:DeploymentConfig
- aws-native:ec2:EIP
- aws-native:ec2:VPNConnection
- aws-native:elasticache:SubnetGroup
- aws-native:elasticbeanstalk:ApplicationVersion
- aws-native:emr:SecurityConfiguration
- aws-native:greengrassv2:Deployment
- aws-native:identitystore:Group
- aws-native:identitystore:GroupMembership
- aws-native:logs:Destination
- aws-native:m2:Application
- aws-native:transfer:Agreement
- aws-native:transfer:Certificate
- aws-native:transfer:Connector
- aws-native:transfer:Profile


## 0.29.0 (September 8, 2022)

* Update to Pulumi SDK v3.39.3
* Update to include the latest resource definitions

#### New resources:

- `aws-native:m2:Environment`
- `aws-native:redshift:ClusterSubnetGroup`

#### New functions:

- `aws-native:m2:getEnvironment`
- `aws-native:redshift:getClusterSubnetGroup`


## 0.28.0 (September 7, 2022)

* Update to include the latest resource definitions

#### New resources:

- `aws-native:logs:LogStream`


## 0.27.0 (September 5, 2022)

* Update to include the latest resource definitions

#### New resources:

- `aws-native:cloudfront:OriginAccessControl`

#### New functions:

- `aws-native:cloudfront:getOriginAccessControl`


## 0.26.0 (September 2, 2022)

* Update to include the latest resource definitions

## 0.25.0 (September 1, 2022)

* Update to include the latest resource definitions

#### New resources:

- `aws-native:redshift:ClusterParameterGroup`

#### New functions:

- `aws-native:redshift:getClusterParameterGroup`


## 0.24.0 (August 31, 2022)

* Update to include the latest resource definitions

#### New resources:

- `aws-native:controltower:EnabledControl`


## 0.23.0 (August 30, 2022)

* Update to include the latest resource definitions

#### New resources:

- `aws-native:connect:InstanceStorageConfig`
- `aws-native:ec2:CustomerGateway`

#### New functions:

- `aws-native:connect:getInstanceStorageConfig`
- `aws-native:ec2:getCustomerGateway`


## 0.22.0 (August 26, 2022)

* Update to include the latest resource definitions
  **PLEASE NOTE:** `aws-native.ec2.CustomerGateway` and `aws-native.ec2.getCustomerGateway` have been removed upstream

#### New resources:

- `aws-native:rds:DBParameterGroup`

#### New functions:

- `aws-native:rds:getDBParameterGroup`


## 0.21.0 (August 24, 2022)

#### New resources:

- `aws-native:apigateway:DocumentationPart`
- `aws-native:appstream:DirectoryConfig`
- `aws-native:appstream:ImageBuilder`
- `aws-native:autoscaling:ScalingPolicy`
- `aws-native:cloudtrail:EventDataStore`
- `aws-native:connect:Instance`
- `aws-native:datasync:LocationFSxONTAP`
- `aws-native:dynamodb:Table`
- `aws-native:ec2:CapacityReservation`
- `aws-native:elasticbeanstalk:Application`
- `aws-native:evidently:Segment`
- `aws-native:iot:CACertificate`
- `aws-native:iot:Policy`
- `aws-native:lakeformation:DataCellsFilter`
- `aws-native:lakeformation:PrincipalPermissions`
- `aws-native:lakeformation:Tag`
- `aws-native:lakeformation:TagAssociation`
- `aws-native:logs:MetricFilter`
- `aws-native:logs:SubscriptionFilter`
- `aws-native:msk:ServerlessCluster`
- `aws-native:networkmanager:ConnectAttachment`
- `aws-native:networkmanager:ConnectPeer`
- `aws-native:networkmanager:CoreNetwork`
- `aws-native:networkmanager:SiteToSiteVpnAttachment`
- `aws-native:networkmanager:VpcAttachment`
- `aws-native:rds:DBCluster`
- `aws-native:rds:DBInstance`
- `aws-native:rds:OptionGroup`
- `aws-native:redshiftserverless:Workgroup`
- `aws-native:rolesanywhere:CRL`
- `aws-native:rolesanywhere:Profile`
- `aws-native:rolesanywhere:TrustAnchor`
- `aws-native:ses:DedicatedIpPool`
- `aws-native:ses:EmailIdentity`
- `aws-native:sns:Topic`
- `aws-native:supportapp:AccountAlias`
- `aws-native:supportapp:SlackChannelConfiguration`
- `aws-native:synthetics:Group`

#### New functions:

- `aws-native:apigateway:getDocumentationPart`
- `aws-native:appstream:getDirectoryConfig`
- `aws-native:appstream:getImageBuilder`
- `aws-native:autoscaling:getScalingPolicy`
- `aws-native:cloudtrail:getEventDataStore`
- `aws-native:connect:getInstance`
- `aws-native:datasync:getLocationFSxONTAP`
- `aws-native:dynamodb:getTable`
- `aws-native:ec2:getCapacityReservation`
- `aws-native:elasticbeanstalk:getApplication`
- `aws-native:evidently:getSegment`
- `aws-native:iot:getCACertificate`
- `aws-native:iot:getPolicy`
- `aws-native:lakeformation:getPrincipalPermissions`
- `aws-native:lakeformation:getTag`
- `aws-native:lakeformation:getTagAssociation`
- `aws-native:logs:getMetricFilter`
- `aws-native:logs:getSubscriptionFilter`
- `aws-native:msk:getServerlessCluster`
- `aws-native:networkmanager:getConnectAttachment`
- `aws-native:networkmanager:getConnectPeer`
- `aws-native:networkmanager:getCoreNetwork`
- `aws-native:networkmanager:getSiteToSiteVpnAttachment`
- `aws-native:networkmanager:getVpcAttachment`
- `aws-native:rds:getDBCluster`
- `aws-native:rds:getDBInstance`
- `aws-native:rds:getOptionGroup`
- `aws-native:redshiftserverless:getWorkgroup`
- `aws-native:rolesanywhere:getCRL`
- `aws-native:rolesanywhere:getProfile`
- `aws-native:rolesanywhere:getTrustAnchor`
- `aws-native:ses:getEmailIdentity`
- `aws-native:sns:getTopic`
- `aws-native:supportapp:getAccountAlias`
- `aws-native:supportapp:getSlackChannelConfiguration`
- `aws-native:synthetics:getGroup`

## 0.20.0 (August 23, 2022)

- Update to include the latest resource definitions
- Add support for allowedAccountIds and forbiddenAccountIds [#508](https://github.com/pulumi/pulumi-aws-native/pull/508)

## 0.19.0 (June 8, 2022)
- Update to include the latest resource definitions

## 0.18.0 (May 26, 2022)
- Update to include the latest resource definitions
- Add support to pass AWS IAM Role to Cloud Control API [#493](https://github.com/pulumi/pulumi-aws-native/pull/493)
- Revert use sequence numbers to generate deterministic autonames.
  [#341](https://github.com/pulumi/pulumi-aws-native/pull/341)

## 0.17.0 (May 23, 2022)
- Update to include the latest resource definitions
- Exclude functions with empty outputs [#499](https://github.com/pulumi/pulumi-aws-native/pull/499)

## 0.16.1 (Apr 29, 2022)
- Fix getAzs [#472](https://github.com/pulumi/pulumi-aws-native/pull/472)

## 0.16.0 (Apr 25, 2022)
- Implement support for getUrlSuffix.
  [#456](https://github.com/pulumi/pulumi-aws-native/pull/456)
- Ensure that all create-only properties are captured as input properties.
  [#466](https://github.com/pulumi/pulumi-aws-native/pull/466)
- Upgrade to Pulumi/Pulumi v3.30.0
  [#458](https://github.com/pulumi/pulumi-aws-native/pull/458)

---

## 0.15.0 (Apr 06, 2022)

- Update pulumi codegen dependency to fix secret property handling in Go SDK
- Update to include the latest resource definitions

## 0.14.0 (Mar 25, 2022)

- Update pulumi codegen dependency to fix secret property handling in Go SDK
- Update to include the latest resource definitions

## 0.13.0 (Feb 23, 2022)

- Update to include the latest resource definitions

## 0.12.0 (Feb 8, 2022)

- Implement get functions for all resources
  [#346](https://github.com/pulumi/pulumi-aws-native/pull/346)
- Update Pulumi dependencies to v3.23.0
- Use sequence numbers to generate deterministic autonames.
  [#341](https://github.com/pulumi/pulumi-aws-native/pull/341)
- Update to include the latest resource definitions


## 0.11.0 (Jan 24, 2022)

- Update to include the latest resource definitions

## 0.10.0 (Jan 10, 2022)

- Update to include the latest resource definitions
- Serve decompressed schema from GetSchema API
  [#287](https://github.com/pulumi/pulumi-aws-native/pull/287)

## 0.9.0 (December 15, 2021)

- Update to include the latest resource definitions
- Explain how to add missing region config
  [#220](https://github.com/pulumi/pulumi-aws-native/issues/220)
- Allow cancellation during ongoing operations
  [#43](https://github.com/pulumi/pulumi-aws-native/issues/43)

## 0.8.0 (November 30, 2021)

Update to include the latest resource definitions

## 0.7.1 (November 23 2021)

- Fix for "Custom providers can leak credentials to state file from environment variables"
  [#236](https://github.com/pulumi/pulumi-aws-native/issues/236)

  **PLEASE READ**

  If you set credentials through environment variables (e.g. `AWS_SECRET_ACCESS_KEY`) AND
  use the SDK to create a provider where these values are not explicitly set, (e.g. `new awsnative.Provider("...");`)
  prior versions of the `aws-native` provider may have included the credentials in the state in clear text.

  All users are recommended to upgrade their provider version to this or newer version and run a `pulumi up`.
  Please also rotate the affected credentials after all relevant stacks have been updated.

  You can check if your state file contains credentials by running `pulumi stack export | grep -A 3 "accessKey\|secretKey\|token"`
  and checking if any unencrypted values are produced. After the update these values will either not be present
  or be stored as encrypted secrets using your stack's preferred encryption provider.

  Note that the Pulumi state backend also encrypts the state as a whole and other state backends
  support a similar mechanism which should significantly limit exposure of the creds.
  Nonetheless, We sincerely regret the inconvenience this causes.

## 0.7.0 (November 22, 2021)

Update to include the latest resource definitions

## 0.6.0 (November 18, 2021)

- Add support for specifying max retries
  [#112](https://github.com/pulumi/pulumi-aws-native/issues/112)
- Fix modifying S3 bucket tags
  [#230](https://github.com/pulumi/pulumi-aws-native/issues/230)

## 0.5.0 (November 15, 2021)

- Improve retry logic to avoid excessive requests to APIs
  [#192](https://github.com/pulumi/pulumi-aws-native/issues/192)

## 0.4.0 (November 8, 2021)

- Update to pulumi v3.16.0 deps
  [#202](https://github.com/pulumi/pulumi-aws-native/pull/202)
- Add autonaming support
  [#156](https://github.com/pulumi/pulumi-aws-native/issues/156)
- Turn off replacement detection to prevent false replacements
  [#186](https://github.com/pulumi/pulumi-aws-native/issues/186)
- Support partial initialization errors
  [#208](https://github.com/pulumi/pulumi-aws-native/pull/208)

## 0.3.0 (November 2, 2021)

- Latest resource definitions

## 0.2.0 (October 8, 2021)

- Deduplicate type names [#160](https://github.com/pulumi/pulumi-aws-native/issues/160)
- [cf2pulumi] Fix supported check for config module [#165](https://github.com/pulumi/pulumi-aws-native/issues/165)
- Improve inline schema type representation [#167](https://github.com/pulumi/pulumi-aws-native/pull/167)

## 0.1.0 (September 30, 2021)

First public release
