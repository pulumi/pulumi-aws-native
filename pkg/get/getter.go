package get

import (
	"context"

	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/amplify"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/aws/aws-sdk-go/service/appstream"
	"github.com/aws/aws-sdk-go/service/appsync"
	"github.com/aws/aws-sdk-go/service/autoscalingplans"
	"github.com/aws/aws-sdk-go/service/backup"
	"github.com/aws/aws-sdk-go/service/cloud9"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/aws/aws-sdk-go/service/codecommit"
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go/service/dax"
	"github.com/aws/aws-sdk-go/service/directoryservice"
	"github.com/aws/aws-sdk-go/service/dlm"
	"github.com/aws/aws-sdk-go/service/docdb"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/emr"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/greengrass"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/inspector"
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iot1clickprojects"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/medialive"
	"github.com/aws/aws-sdk-go/service/mediastore"
	"github.com/aws/aws-sdk-go/service/mq"
	"github.com/aws/aws-sdk-go/service/neptune"
	"github.com/aws/aws-sdk-go/service/opsworks"
	"github.com/aws/aws-sdk-go/service/opsworkscm"
	"github.com/aws/aws-sdk-go/service/pinpoint"
	"github.com/aws/aws-sdk-go/service/pinpointemail"
	"github.com/aws/aws-sdk-go/service/ram"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/aws/aws-sdk-go/service/robomaker"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53resolver"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sagemaker"
	"github.com/aws/aws-sdk-go/service/servicecatalog"
	"github.com/aws/aws-sdk-go/service/servicediscovery"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/transfer"
)

type NotFoundError struct {
	cause error
}

func NewNotFoundError(cause error) error {
	return &NotFoundError{cause}
}

func (e *NotFoundError) Error() string {
	if e.cause == nil {
		return "Resource not found"
	}
	return e.cause.Error()
}

func IsNotFound(err error) bool {
	_, ok := err.(*NotFoundError)
	return ok
}

type Getter struct {
	amplify                  *amplify.Amplify
	apigateway               *apigateway.APIGateway
	apigatewayv2             *apigatewayv2.ApiGatewayV2
	appmesh                  *appmesh.AppMesh
	appstream                *appstream.AppStream
	appsync                  *appsync.AppSync
	autoscalingplans         *autoscalingplans.AutoScalingPlans
	backup                   *backup.Backup
	cloud9                   *cloud9.Cloud9
	cloudformation           *cloudformation.CloudFormation
	cloudfront               *cloudfront.CloudFront
	cloudtrail               *cloudtrail.CloudTrail
	cloudwatch               *cloudwatch.CloudWatch
	cloudwatchlogs           *cloudwatchlogs.CloudWatchLogs
	codebuild                *codebuild.CodeBuild
	codecommit               *codecommit.CodeCommit
	codepipeline             *codepipeline.CodePipeline
	cognitoidentity          *cognitoidentity.CognitoIdentity
	configservice            *configservice.ConfigService
	databasemigrationservice *databasemigrationservice.DatabaseMigrationService
	dax                      *dax.DAX
	directoryservice         *directoryservice.DirectoryService
	dlm                      *dlm.DLM
	docdb                    *docdb.DocDB
	dynamodb                 *dynamodb.DynamoDB
	ec2                      *ec2.EC2
	ecr                      *ecr.ECR
	ecs                      *ecs.ECS
	efs                      *efs.EFS
	eks                      *eks.EKS
	elasticache              *elasticache.ElastiCache
	elasticbeanstalk         *elasticbeanstalk.ElasticBeanstalk
	elasticsearchservice     *elasticsearchservice.ElasticsearchService
	elb                      *elb.ELB
	elbv2                    *elbv2.ELBV2
	emr                      *emr.EMR
	eventbridge              *eventbridge.EventBridge
	firehose                 *firehose.Firehose
	greengrass               *greengrass.Greengrass
	iam                      *iam.IAM
	inspector                *inspector.Inspector
	iot                      *iot.IoT
	iot1clickprojects        *iot1clickprojects.IoT1ClickProjects
	kinesis                  *kinesis.Kinesis
	kms                      *kms.KMS
	lambda                   *lambda.Lambda
	medialive                *medialive.MediaLive
	mediastore               *mediastore.MediaStore
	mq                       *mq.MQ
	neptune                  *neptune.Neptune
	opsworks                 *opsworks.OpsWorks
	opsworkscm               *opsworkscm.OpsWorksCM
	pinpoint                 *pinpoint.Pinpoint
	pinpointemail            *pinpointemail.PinpointEmail
	ram                      *ram.RAM
	rds                      *rds.RDS
	redshift                 *redshift.Redshift
	robomaker                *robomaker.RoboMaker
	route53                  *route53.Route53
	route53resolver          *route53resolver.Route53Resolver
	s3                       *s3.S3
	sagemaker                *sagemaker.SageMaker
	servicecatalog           *servicecatalog.ServiceCatalog
	servicediscovery         *servicediscovery.ServiceDiscovery
	sfn                      *sfn.SFN
	sns                      *sns.SNS
	sqs                      *sqs.SQS
	ssm                      *ssm.SSM
	transfer                 *transfer.Transfer
}

func NewGetter(configProvider client.ConfigProvider) *Getter {
	return &Getter{
		amplify:                  amplify.New(configProvider),
		apigateway:               apigateway.New(configProvider),
		apigatewayv2:             apigatewayv2.New(configProvider),
		appmesh:                  appmesh.New(configProvider),
		appstream:                appstream.New(configProvider),
		appsync:                  appsync.New(configProvider),
		autoscalingplans:         autoscalingplans.New(configProvider),
		backup:                   backup.New(configProvider),
		cloud9:                   cloud9.New(configProvider),
		cloudformation:           cloudformation.New(configProvider),
		cloudfront:               cloudfront.New(configProvider),
		cloudtrail:               cloudtrail.New(configProvider),
		cloudwatch:               cloudwatch.New(configProvider),
		cloudwatchlogs:           cloudwatchlogs.New(configProvider),
		codebuild:                codebuild.New(configProvider),
		codecommit:               codecommit.New(configProvider),
		codepipeline:             codepipeline.New(configProvider),
		cognitoidentity:          cognitoidentity.New(configProvider),
		configservice:            configservice.New(configProvider),
		databasemigrationservice: databasemigrationservice.New(configProvider),
		dax:                      dax.New(configProvider),
		directoryservice:         directoryservice.New(configProvider),
		dlm:                      dlm.New(configProvider),
		docdb:                    docdb.New(configProvider),
		dynamodb:                 dynamodb.New(configProvider),
		ec2:                      ec2.New(configProvider),
		ecr:                      ecr.New(configProvider),
		ecs:                      ecs.New(configProvider),
		efs:                      efs.New(configProvider),
		eks:                      eks.New(configProvider),
		elasticache:              elasticache.New(configProvider),
		elasticbeanstalk:         elasticbeanstalk.New(configProvider),
		elasticsearchservice:     elasticsearchservice.New(configProvider),
		elb:                      elb.New(configProvider),
		elbv2:                    elbv2.New(configProvider),
		emr:                      emr.New(configProvider),
		eventbridge:              eventbridge.New(configProvider),
		firehose:                 firehose.New(configProvider),
		greengrass:               greengrass.New(configProvider),
		iam:                      iam.New(configProvider),
		inspector:                inspector.New(configProvider),
		iot:                      iot.New(configProvider),
		iot1clickprojects:        iot1clickprojects.New(configProvider),
		kinesis:                  kinesis.New(configProvider),
		kms:                      kms.New(configProvider),
		lambda:                   lambda.New(configProvider),
		medialive:                medialive.New(configProvider),
		mediastore:               mediastore.New(configProvider),
		mq:                       mq.New(configProvider),
		neptune:                  neptune.New(configProvider),
		opsworks:                 opsworks.New(configProvider),
		opsworkscm:               opsworkscm.New(configProvider),
		pinpoint:                 pinpoint.New(configProvider),
		pinpointemail:            pinpointemail.New(configProvider),
		ram:                      ram.New(configProvider),
		rds:                      rds.New(configProvider),
		redshift:                 redshift.New(configProvider),
		robomaker:                robomaker.New(configProvider),
		route53:                  route53.New(configProvider),
		route53resolver:          route53resolver.New(configProvider),
		s3:                       s3.New(configProvider),
		sagemaker:                sagemaker.New(configProvider),
		servicecatalog:           servicecatalog.New(configProvider),
		servicediscovery:         servicediscovery.New(configProvider),
		sfn:                      sfn.New(configProvider),
		sns:                      sns.New(configProvider),
		sqs:                      sqs.New(configProvider),
		ssm:                      ssm.New(configProvider),
		transfer:                 transfer.New(configProvider),
	}
}

func (g *Getter) GetResourceAttributes(ctx context.Context, resourceType, physicalResourceID string) (map[string]interface{}, error) {
	switch resourceType {
	case "AWS::AmazonMQ::Broker":
		return g.getAmazonMQBrokerAttributes(ctx, physicalResourceID)
	case "AWS::AmazonMQ::Configuration":
		return g.getAmazonMQConfigurationAttributes(ctx, physicalResourceID)
	case "AWS::AmazonMQ::ConfigurationAssociation":
		return nil, nil
	case "AWS::Amplify::App":
		return g.getAmplifyAppAttributes(ctx, physicalResourceID)
	case "AWS::Amplify::Branch":
		return g.getAmplifyBranchAttributes(ctx, physicalResourceID)
	case "AWS::Amplify::Domain":
		return g.getAmplifyDomainAttributes(ctx, physicalResourceID)
	case "AWS::ApiGateway::Account":
		return nil, nil
	case "AWS::ApiGateway::ApiKey":
		return nil, nil
	case "AWS::ApiGateway::Authorizer":
		return nil, nil
	case "AWS::ApiGateway::BasePathMapping":
		return nil, nil
	case "AWS::ApiGateway::ClientCertificate":
		return nil, nil
	case "AWS::ApiGateway::Deployment":
		return nil, nil
	case "AWS::ApiGateway::DocumentationPart":
		return nil, nil
	case "AWS::ApiGateway::DocumentationVersion":
		return nil, nil
	case "AWS::ApiGateway::DomainName":
		return g.getApiGatewayDomainNameAttributes(ctx, physicalResourceID)
	case "AWS::ApiGateway::GatewayResponse":
		return nil, nil
	case "AWS::ApiGateway::Method":
		return nil, nil
	case "AWS::ApiGateway::Model":
		return nil, nil
	case "AWS::ApiGateway::RequestValidator":
		return nil, nil
	case "AWS::ApiGateway::Resource":
		return nil, nil
	case "AWS::ApiGateway::RestApi":
		return g.getApiGatewayRestApiAttributes(ctx, physicalResourceID)
	case "AWS::ApiGateway::Stage":
		return nil, nil
	case "AWS::ApiGateway::UsagePlan":
		return nil, nil
	case "AWS::ApiGateway::UsagePlanKey":
		return nil, nil
	case "AWS::ApiGateway::VpcLink":
		return nil, nil
	case "AWS::ApiGatewayV2::Api":
		return nil, nil
	case "AWS::ApiGatewayV2::ApiMapping":
		return nil, nil
	case "AWS::ApiGatewayV2::Authorizer":
		return nil, nil
	case "AWS::ApiGatewayV2::Deployment":
		return nil, nil
	case "AWS::ApiGatewayV2::DomainName":
		return g.getApiGatewayV2DomainNameAttributes(ctx, physicalResourceID)
	case "AWS::ApiGatewayV2::Integration":
		return nil, nil
	case "AWS::ApiGatewayV2::IntegrationResponse":
		return nil, nil
	case "AWS::ApiGatewayV2::Model":
		return nil, nil
	case "AWS::ApiGatewayV2::Route":
		return nil, nil
	case "AWS::ApiGatewayV2::RouteResponse":
		return nil, nil
	case "AWS::ApiGatewayV2::Stage":
		return nil, nil
	case "AWS::AppMesh::Mesh":
		return g.getAppMeshMeshAttributes(ctx, physicalResourceID)
	case "AWS::AppMesh::Route":
		return g.getAppMeshRouteAttributes(ctx, physicalResourceID)
	case "AWS::AppMesh::VirtualNode":
		return g.getAppMeshVirtualNodeAttributes(ctx, physicalResourceID)
	case "AWS::AppMesh::VirtualRouter":
		return g.getAppMeshVirtualRouterAttributes(ctx, physicalResourceID)
	case "AWS::AppMesh::VirtualService":
		return g.getAppMeshVirtualServiceAttributes(ctx, physicalResourceID)
	case "AWS::AppStream::DirectoryConfig":
		return nil, nil
	case "AWS::AppStream::Fleet":
		return nil, nil
	case "AWS::AppStream::ImageBuilder":
		return g.getAppStreamImageBuilderAttributes(ctx, physicalResourceID)
	case "AWS::AppStream::Stack":
		return nil, nil
	case "AWS::AppStream::StackFleetAssociation":
		return nil, nil
	case "AWS::AppStream::StackUserAssociation":
		return nil, nil
	case "AWS::AppStream::User":
		return nil, nil
	case "AWS::AppSync::ApiKey":
		return g.getAppSyncApiKeyAttributes(ctx, physicalResourceID)
	case "AWS::AppSync::DataSource":
		return g.getAppSyncDataSourceAttributes(ctx, physicalResourceID)
	case "AWS::AppSync::FunctionConfiguration":
		return g.getAppSyncFunctionConfigurationAttributes(ctx, physicalResourceID)
	case "AWS::AppSync::GraphQLApi":
		return g.getAppSyncGraphQLApiAttributes(ctx, physicalResourceID)
	case "AWS::AppSync::GraphQLSchema":
		return nil, nil
	case "AWS::AppSync::Resolver":
		return g.getAppSyncResolverAttributes(ctx, physicalResourceID)
	case "AWS::ApplicationAutoScaling::ScalableTarget":
		return nil, nil
	case "AWS::ApplicationAutoScaling::ScalingPolicy":
		return nil, nil
	case "AWS::Athena::NamedQuery":
		return nil, nil
	case "AWS::AutoScaling::AutoScalingGroup":
		return nil, nil
	case "AWS::AutoScaling::LaunchConfiguration":
		return nil, nil
	case "AWS::AutoScaling::LifecycleHook":
		return nil, nil
	case "AWS::AutoScaling::ScalingPolicy":
		return nil, nil
	case "AWS::AutoScaling::ScheduledAction":
		return nil, nil
	case "AWS::AutoScalingPlans::ScalingPlan":
		return g.getAutoScalingPlansScalingPlanAttributes(ctx, physicalResourceID)
	case "AWS::Backup::BackupPlan":
		return g.getBackupBackupPlanAttributes(ctx, physicalResourceID)
	case "AWS::Backup::BackupSelection":
		return g.getBackupBackupSelectionAttributes(ctx, physicalResourceID)
	case "AWS::Backup::BackupVault":
		return g.getBackupBackupVaultAttributes(ctx, physicalResourceID)
	case "AWS::Batch::ComputeEnvironment":
		return nil, nil
	case "AWS::Batch::JobDefinition":
		return nil, nil
	case "AWS::Batch::JobQueue":
		return nil, nil
	case "AWS::Budgets::Budget":
		return nil, nil
	case "AWS::CertificateManager::Certificate":
		return nil, nil
	case "AWS::Cloud9::EnvironmentEC2":
		return g.getCloud9EnvironmentEC2Attributes(ctx, physicalResourceID)
	case "AWS::CloudFormation::CustomResource":
		return nil, nil
	case "AWS::CloudFormation::Macro":
		return nil, nil
	case "AWS::CloudFormation::Stack":
		return nil, nil
	case "AWS::CloudFormation::WaitCondition":
		return g.getCloudFormationWaitConditionAttributes(ctx, physicalResourceID)
	case "AWS::CloudFormation::WaitConditionHandle":
		return nil, nil
	case "AWS::CloudFront::CloudFrontOriginAccessIdentity":
		return g.getCloudFrontCloudFrontOriginAccessIdentityAttributes(ctx, physicalResourceID)
	case "AWS::CloudFront::Distribution":
		return g.getCloudFrontDistributionAttributes(ctx, physicalResourceID)
	case "AWS::CloudFront::StreamingDistribution":
		return g.getCloudFrontStreamingDistributionAttributes(ctx, physicalResourceID)
	case "AWS::CloudTrail::Trail":
		return g.getCloudTrailTrailAttributes(ctx, physicalResourceID)
	case "AWS::CloudWatch::Alarm":
		return g.getCloudWatchAlarmAttributes(ctx, physicalResourceID)
	case "AWS::CloudWatch::AnomalyDetector":
		return nil, nil
	case "AWS::CloudWatch::Dashboard":
		return nil, nil
	case "AWS::CodeBuild::Project":
		return g.getCodeBuildProjectAttributes(ctx, physicalResourceID)
	case "AWS::CodeBuild::SourceCredential":
		return nil, nil
	case "AWS::CodeCommit::Repository":
		return g.getCodeCommitRepositoryAttributes(ctx, physicalResourceID)
	case "AWS::CodeDeploy::Application":
		return nil, nil
	case "AWS::CodeDeploy::DeploymentConfig":
		return nil, nil
	case "AWS::CodeDeploy::DeploymentGroup":
		return nil, nil
	case "AWS::CodePipeline::CustomActionType":
		return nil, nil
	case "AWS::CodePipeline::Pipeline":
		return g.getCodePipelinePipelineAttributes(ctx, physicalResourceID)
	case "AWS::CodePipeline::Webhook":
		return g.getCodePipelineWebhookAttributes(ctx, physicalResourceID)
	case "AWS::CodeStar::GitHubRepository":
		return nil, nil
	case "AWS::Cognito::IdentityPool":
		return g.getCognitoIdentityPoolAttributes(ctx, physicalResourceID)
	case "AWS::Cognito::IdentityPoolRoleAttachment":
		return nil, nil
	case "AWS::Cognito::UserPool":
		return g.getCognitoUserPoolAttributes(ctx, physicalResourceID)
	case "AWS::Cognito::UserPoolClient":
		return g.getCognitoUserPoolClientAttributes(ctx, physicalResourceID)
	case "AWS::Cognito::UserPoolDomain":
		return nil, nil
	case "AWS::Cognito::UserPoolGroup":
		return nil, nil
	case "AWS::Cognito::UserPoolIdentityProvider":
		return nil, nil
	case "AWS::Cognito::UserPoolResourceServer":
		return nil, nil
	case "AWS::Cognito::UserPoolRiskConfigurationAttachment":
		return nil, nil
	case "AWS::Cognito::UserPoolUICustomizationAttachment":
		return nil, nil
	case "AWS::Cognito::UserPoolUser":
		return nil, nil
	case "AWS::Cognito::UserPoolUserToGroupAttachment":
		return nil, nil
	case "AWS::Config::AggregationAuthorization":
		return nil, nil
	case "AWS::Config::ConfigRule":
		return g.getConfigConfigRuleAttributes(ctx, physicalResourceID)
	case "AWS::Config::ConfigurationAggregator":
		return nil, nil
	case "AWS::Config::ConfigurationRecorder":
		return nil, nil
	case "AWS::Config::DeliveryChannel":
		return nil, nil
	case "AWS::Config::OrganizationConfigRule":
		return nil, nil
	case "AWS::Config::RemediationConfiguration":
		return nil, nil
	case "AWS::DAX::Cluster":
		return g.getDAXClusterAttributes(ctx, physicalResourceID)
	case "AWS::DAX::ParameterGroup":
		return nil, nil
	case "AWS::DAX::SubnetGroup":
		return nil, nil
	case "AWS::DLM::LifecyclePolicy":
		return g.getDLMLifecyclePolicyAttributes(ctx, physicalResourceID)
	case "AWS::DMS::Certificate":
		return nil, nil
	case "AWS::DMS::Endpoint":
		return g.getDMSEndpointAttributes(ctx, physicalResourceID)
	case "AWS::DMS::EventSubscription":
		return nil, nil
	case "AWS::DMS::ReplicationInstance":
		return g.getDMSReplicationInstanceAttributes(ctx, physicalResourceID)
	case "AWS::DMS::ReplicationSubnetGroup":
		return nil, nil
	case "AWS::DMS::ReplicationTask":
		return nil, nil
	case "AWS::DataPipeline::Pipeline":
		return nil, nil
	case "AWS::DirectoryService::MicrosoftAD":
		return g.getDirectoryServiceMicrosoftADAttributes(ctx, physicalResourceID)
	case "AWS::DirectoryService::SimpleAD":
		return g.getDirectoryServiceSimpleADAttributes(ctx, physicalResourceID)
	case "AWS::DocDB::DBCluster":
		return g.getDocDBDBClusterAttributes(ctx, physicalResourceID)
	case "AWS::DocDB::DBClusterParameterGroup":
		return nil, nil
	case "AWS::DocDB::DBInstance":
		return g.getDocDBDBInstanceAttributes(ctx, physicalResourceID)
	case "AWS::DocDB::DBSubnetGroup":
		return nil, nil
	case "AWS::DynamoDB::Table":
		return g.getDynamoDBTableAttributes(ctx, physicalResourceID)
	case "AWS::EC2::CapacityReservation":
		return g.getEC2CapacityReservationAttributes(ctx, physicalResourceID)
	case "AWS::EC2::ClientVpnAuthorizationRule":
		return nil, nil
	case "AWS::EC2::ClientVpnEndpoint":
		return nil, nil
	case "AWS::EC2::ClientVpnRoute":
		return nil, nil
	case "AWS::EC2::ClientVpnTargetNetworkAssociation":
		return nil, nil
	case "AWS::EC2::CustomerGateway":
		return nil, nil
	case "AWS::EC2::DHCPOptions":
		return nil, nil
	case "AWS::EC2::EC2Fleet":
		return nil, nil
	case "AWS::EC2::EIP":
		return g.getEC2EIPAttributes(ctx, physicalResourceID)
	case "AWS::EC2::EIPAssociation":
		return nil, nil
	case "AWS::EC2::EgressOnlyInternetGateway":
		return nil, nil
	case "AWS::EC2::FlowLog":
		return nil, nil
	case "AWS::EC2::Host":
		return nil, nil
	case "AWS::EC2::Instance":
		return g.getEC2InstanceAttributes(ctx, physicalResourceID)
	case "AWS::EC2::InternetGateway":
		return nil, nil
	case "AWS::EC2::LaunchTemplate":
		return g.getEC2LaunchTemplateAttributes(ctx, physicalResourceID)
	case "AWS::EC2::NatGateway":
		return nil, nil
	case "AWS::EC2::NetworkAcl":
		return nil, nil
	case "AWS::EC2::NetworkAclEntry":
		return nil, nil
	case "AWS::EC2::NetworkInterface":
		return g.getEC2NetworkInterfaceAttributes(ctx, physicalResourceID)
	case "AWS::EC2::NetworkInterfaceAttachment":
		return nil, nil
	case "AWS::EC2::NetworkInterfacePermission":
		return nil, nil
	case "AWS::EC2::PlacementGroup":
		return nil, nil
	case "AWS::EC2::Route":
		return nil, nil
	case "AWS::EC2::RouteTable":
		return nil, nil
	case "AWS::EC2::SecurityGroup":
		return g.getEC2SecurityGroupAttributes(ctx, physicalResourceID)
	case "AWS::EC2::SecurityGroupEgress":
		return nil, nil
	case "AWS::EC2::SecurityGroupIngress":
		return nil, nil
	case "AWS::EC2::SpotFleet":
		return nil, nil
	case "AWS::EC2::Subnet":
		return g.getEC2SubnetAttributes(ctx, physicalResourceID)
	case "AWS::EC2::SubnetCidrBlock":
		return nil, nil
	case "AWS::EC2::SubnetNetworkAclAssociation":
		return g.getEC2SubnetNetworkAclAssociationAttributes(ctx, physicalResourceID)
	case "AWS::EC2::SubnetRouteTableAssociation":
		return nil, nil
	case "AWS::EC2::TrafficMirrorFilter":
		return nil, nil
	case "AWS::EC2::TrafficMirrorFilterRule":
		return nil, nil
	case "AWS::EC2::TrafficMirrorSession":
		return nil, nil
	case "AWS::EC2::TrafficMirrorTarget":
		return nil, nil
	case "AWS::EC2::TransitGateway":
		return nil, nil
	case "AWS::EC2::TransitGatewayAttachment":
		return nil, nil
	case "AWS::EC2::TransitGatewayRoute":
		return nil, nil
	case "AWS::EC2::TransitGatewayRouteTable":
		return nil, nil
	case "AWS::EC2::TransitGatewayRouteTableAssociation":
		return nil, nil
	case "AWS::EC2::TransitGatewayRouteTablePropagation":
		return nil, nil
	case "AWS::EC2::VPC":
		return g.getEC2VPCAttributes(ctx, physicalResourceID)
	case "AWS::EC2::VPCCidrBlock":
		return nil, nil
	case "AWS::EC2::VPCDHCPOptionsAssociation":
		return nil, nil
	case "AWS::EC2::VPCEndpoint":
		return g.getEC2VPCEndpointAttributes(ctx, physicalResourceID)
	case "AWS::EC2::VPCEndpointServicePermissions":
		return nil, nil
	case "AWS::EC2::VPCGatewayAttachment":
		return nil, nil
	case "AWS::EC2::VPCPeeringConnection":
		return nil, nil
	case "AWS::EC2::VPNConnection":
		return nil, nil
	case "AWS::EC2::VPNConnectionRoute":
		return nil, nil
	case "AWS::EC2::VPNGateway":
		return nil, nil
	case "AWS::EC2::VPNGatewayRoutePropagation":
		return nil, nil
	case "AWS::EC2::Volume":
		return nil, nil
	case "AWS::EC2::VolumeAttachment":
		return nil, nil
	case "AWS::ECR::Repository":
		return g.getECRRepositoryAttributes(ctx, physicalResourceID)
	case "AWS::ECS::Cluster":
		return g.getECSClusterAttributes(ctx, physicalResourceID)
	case "AWS::ECS::Service":
		return g.getECSServiceAttributes(ctx, physicalResourceID)
	case "AWS::ECS::TaskDefinition":
		return nil, nil
	case "AWS::EFS::FileSystem":
		return nil, nil
	case "AWS::EFS::MountTarget":
		return g.getEFSMountTargetAttributes(ctx, physicalResourceID)
	case "AWS::EKS::Cluster":
		return g.getEKSClusterAttributes(ctx, physicalResourceID)
	case "AWS::EMR::Cluster":
		return g.getEMRClusterAttributes(ctx, physicalResourceID)
	case "AWS::EMR::InstanceFleetConfig":
		return nil, nil
	case "AWS::EMR::InstanceGroupConfig":
		return nil, nil
	case "AWS::EMR::SecurityConfiguration":
		return nil, nil
	case "AWS::EMR::Step":
		return nil, nil
	case "AWS::ElastiCache::CacheCluster":
		return g.getElastiCacheCacheClusterAttributes(ctx, physicalResourceID)
	case "AWS::ElastiCache::ParameterGroup":
		return nil, nil
	case "AWS::ElastiCache::ReplicationGroup":
		return g.getElastiCacheReplicationGroupAttributes(ctx, physicalResourceID)
	case "AWS::ElastiCache::SecurityGroup":
		return nil, nil
	case "AWS::ElastiCache::SecurityGroupIngress":
		return nil, nil
	case "AWS::ElastiCache::SubnetGroup":
		return nil, nil
	case "AWS::ElasticBeanstalk::Application":
		return nil, nil
	case "AWS::ElasticBeanstalk::ApplicationVersion":
		return nil, nil
	case "AWS::ElasticBeanstalk::ConfigurationTemplate":
		return nil, nil
	case "AWS::ElasticBeanstalk::Environment":
		return g.getElasticBeanstalkEnvironmentAttributes(ctx, physicalResourceID)
	case "AWS::ElasticLoadBalancing::LoadBalancer":
		return g.getElasticLoadBalancingLoadBalancerAttributes(ctx, physicalResourceID)
	case "AWS::ElasticLoadBalancingV2::Listener":
		return nil, nil
	case "AWS::ElasticLoadBalancingV2::ListenerCertificate":
		return nil, nil
	case "AWS::ElasticLoadBalancingV2::ListenerRule":
		return nil, nil
	case "AWS::ElasticLoadBalancingV2::LoadBalancer":
		return g.getElasticLoadBalancingV2LoadBalancerAttributes(ctx, physicalResourceID)
	case "AWS::ElasticLoadBalancingV2::TargetGroup":
		return g.getElasticLoadBalancingV2TargetGroupAttributes(ctx, physicalResourceID)
	case "AWS::Elasticsearch::Domain":
		return g.getElasticsearchDomainAttributes(ctx, physicalResourceID)
	case "AWS::Events::EventBus":
		return g.getEventsEventBusAttributes(ctx, physicalResourceID)
	case "AWS::Events::EventBusPolicy":
		return nil, nil
	case "AWS::Events::Rule":
		return g.getEventsRuleAttributes(ctx, physicalResourceID)
	case "AWS::FSx::FileSystem":
		return nil, nil
	case "AWS::GameLift::Alias":
		return nil, nil
	case "AWS::GameLift::Build":
		return nil, nil
	case "AWS::GameLift::Fleet":
		return nil, nil
	case "AWS::Glue::Classifier":
		return nil, nil
	case "AWS::Glue::Connection":
		return nil, nil
	case "AWS::Glue::Crawler":
		return nil, nil
	case "AWS::Glue::DataCatalogEncryptionSettings":
		return nil, nil
	case "AWS::Glue::Database":
		return nil, nil
	case "AWS::Glue::DevEndpoint":
		return nil, nil
	case "AWS::Glue::Job":
		return nil, nil
	case "AWS::Glue::MLTransform":
		return nil, nil
	case "AWS::Glue::Partition":
		return nil, nil
	case "AWS::Glue::SecurityConfiguration":
		return nil, nil
	case "AWS::Glue::Table":
		return nil, nil
	case "AWS::Glue::Trigger":
		return nil, nil
	case "AWS::Glue::Workflow":
		return nil, nil
	case "AWS::Greengrass::ConnectorDefinition":
		return g.getGreengrassConnectorDefinitionAttributes(ctx, physicalResourceID)
	case "AWS::Greengrass::ConnectorDefinitionVersion":
		return nil, nil
	case "AWS::Greengrass::CoreDefinition":
		return g.getGreengrassCoreDefinitionAttributes(ctx, physicalResourceID)
	case "AWS::Greengrass::CoreDefinitionVersion":
		return nil, nil
	case "AWS::Greengrass::DeviceDefinition":
		return g.getGreengrassDeviceDefinitionAttributes(ctx, physicalResourceID)
	case "AWS::Greengrass::DeviceDefinitionVersion":
		return nil, nil
	case "AWS::Greengrass::FunctionDefinition":
		return g.getGreengrassFunctionDefinitionAttributes(ctx, physicalResourceID)
	case "AWS::Greengrass::FunctionDefinitionVersion":
		return nil, nil
	case "AWS::Greengrass::Group":
		return g.getGreengrassGroupAttributes(ctx, physicalResourceID)
	case "AWS::Greengrass::GroupVersion":
		return nil, nil
	case "AWS::Greengrass::LoggerDefinition":
		return g.getGreengrassLoggerDefinitionAttributes(ctx, physicalResourceID)
	case "AWS::Greengrass::LoggerDefinitionVersion":
		return nil, nil
	case "AWS::Greengrass::ResourceDefinition":
		return g.getGreengrassResourceDefinitionAttributes(ctx, physicalResourceID)
	case "AWS::Greengrass::ResourceDefinitionVersion":
		return nil, nil
	case "AWS::Greengrass::SubscriptionDefinition":
		return g.getGreengrassSubscriptionDefinitionAttributes(ctx, physicalResourceID)
	case "AWS::Greengrass::SubscriptionDefinitionVersion":
		return nil, nil
	case "AWS::GuardDuty::Detector":
		return nil, nil
	case "AWS::GuardDuty::Filter":
		return nil, nil
	case "AWS::GuardDuty::IPSet":
		return nil, nil
	case "AWS::GuardDuty::Master":
		return nil, nil
	case "AWS::GuardDuty::Member":
		return nil, nil
	case "AWS::GuardDuty::ThreatIntelSet":
		return nil, nil
	case "AWS::IAM::AccessKey":
		return g.getIAMAccessKeyAttributes(ctx, physicalResourceID)
	case "AWS::IAM::Group":
		return g.getIAMGroupAttributes(ctx, physicalResourceID)
	case "AWS::IAM::InstanceProfile":
		return g.getIAMInstanceProfileAttributes(ctx, physicalResourceID)
	case "AWS::IAM::ManagedPolicy":
		return nil, nil
	case "AWS::IAM::Policy":
		return nil, nil
	case "AWS::IAM::Role":
		return g.getIAMRoleAttributes(ctx, physicalResourceID)
	case "AWS::IAM::ServiceLinkedRole":
		return nil, nil
	case "AWS::IAM::User":
		return g.getIAMUserAttributes(ctx, physicalResourceID)
	case "AWS::IAM::UserToGroupAddition":
		return nil, nil
	case "AWS::Inspector::AssessmentTarget":
		return g.getInspectorAssessmentTargetAttributes(ctx, physicalResourceID)
	case "AWS::Inspector::AssessmentTemplate":
		return g.getInspectorAssessmentTemplateAttributes(ctx, physicalResourceID)
	case "AWS::Inspector::ResourceGroup":
		return g.getInspectorResourceGroupAttributes(ctx, physicalResourceID)
	case "AWS::IoT1Click::Device":
		return g.getIoT1ClickDeviceAttributes(ctx, physicalResourceID)
	case "AWS::IoT1Click::Placement":
		return g.getIoT1ClickPlacementAttributes(ctx, physicalResourceID)
	case "AWS::IoT1Click::Project":
		return g.getIoT1ClickProjectAttributes(ctx, physicalResourceID)
	case "AWS::IoT::Certificate":
		return g.getIoTCertificateAttributes(ctx, physicalResourceID)
	case "AWS::IoT::Policy":
		return g.getIoTPolicyAttributes(ctx, physicalResourceID)
	case "AWS::IoT::PolicyPrincipalAttachment":
		return nil, nil
	case "AWS::IoT::Thing":
		return nil, nil
	case "AWS::IoT::ThingPrincipalAttachment":
		return nil, nil
	case "AWS::IoT::TopicRule":
		return g.getIoTTopicRuleAttributes(ctx, physicalResourceID)
	case "AWS::IoTAnalytics::Channel":
		return nil, nil
	case "AWS::IoTAnalytics::Dataset":
		return nil, nil
	case "AWS::IoTAnalytics::Datastore":
		return nil, nil
	case "AWS::IoTAnalytics::Pipeline":
		return nil, nil
	case "AWS::IoTEvents::DetectorModel":
		return nil, nil
	case "AWS::IoTEvents::Input":
		return nil, nil
	case "AWS::IoTThingsGraph::FlowTemplate":
		return nil, nil
	case "AWS::KMS::Alias":
		return nil, nil
	case "AWS::KMS::Key":
		return g.getKMSKeyAttributes(ctx, physicalResourceID)
	case "AWS::Kinesis::Stream":
		return g.getKinesisStreamAttributes(ctx, physicalResourceID)
	case "AWS::Kinesis::StreamConsumer":
		return g.getKinesisStreamConsumerAttributes(ctx, physicalResourceID)
	case "AWS::KinesisAnalytics::Application":
		return nil, nil
	case "AWS::KinesisAnalytics::ApplicationOutput":
		return nil, nil
	case "AWS::KinesisAnalytics::ApplicationReferenceDataSource":
		return nil, nil
	case "AWS::KinesisAnalyticsV2::Application":
		return nil, nil
	case "AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption":
		return nil, nil
	case "AWS::KinesisAnalyticsV2::ApplicationOutput":
		return nil, nil
	case "AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource":
		return nil, nil
	case "AWS::KinesisFirehose::DeliveryStream":
		return g.getKinesisFirehoseDeliveryStreamAttributes(ctx, physicalResourceID)
	case "AWS::LakeFormation::DataLakeSettings":
		return nil, nil
	case "AWS::LakeFormation::Permissions":
		return nil, nil
	case "AWS::LakeFormation::Resource":
		return nil, nil
	case "AWS::Lambda::Alias":
		return nil, nil
	case "AWS::Lambda::EventSourceMapping":
		return nil, nil
	case "AWS::Lambda::Function":
		return g.getLambdaFunctionAttributes(ctx, physicalResourceID)
	case "AWS::Lambda::LayerVersion":
		return nil, nil
	case "AWS::Lambda::LayerVersionPermission":
		return nil, nil
	case "AWS::Lambda::Permission":
		return nil, nil
	case "AWS::Lambda::Version":
		return g.getLambdaVersionAttributes(ctx, physicalResourceID)
	case "AWS::Logs::Destination":
		return g.getLogsDestinationAttributes(ctx, physicalResourceID)
	case "AWS::Logs::LogGroup":
		return g.getLogsLogGroupAttributes(ctx, physicalResourceID)
	case "AWS::Logs::LogStream":
		return nil, nil
	case "AWS::Logs::MetricFilter":
		return nil, nil
	case "AWS::Logs::SubscriptionFilter":
		return nil, nil
	case "AWS::MSK::Cluster":
		return nil, nil
	case "AWS::MediaLive::Channel":
		return g.getMediaLiveChannelAttributes(ctx, physicalResourceID)
	case "AWS::MediaLive::Input":
		return g.getMediaLiveInputAttributes(ctx, physicalResourceID)
	case "AWS::MediaLive::InputSecurityGroup":
		return g.getMediaLiveInputSecurityGroupAttributes(ctx, physicalResourceID)
	case "AWS::MediaStore::Container":
		return g.getMediaStoreContainerAttributes(ctx, physicalResourceID)
	case "AWS::Neptune::DBCluster":
		return g.getNeptuneDBClusterAttributes(ctx, physicalResourceID)
	case "AWS::Neptune::DBClusterParameterGroup":
		return nil, nil
	case "AWS::Neptune::DBInstance":
		return g.getNeptuneDBInstanceAttributes(ctx, physicalResourceID)
	case "AWS::Neptune::DBParameterGroup":
		return nil, nil
	case "AWS::Neptune::DBSubnetGroup":
		return nil, nil
	case "AWS::OpsWorks::App":
		return nil, nil
	case "AWS::OpsWorks::ElasticLoadBalancerAttachment":
		return nil, nil
	case "AWS::OpsWorks::Instance":
		return g.getOpsWorksInstanceAttributes(ctx, physicalResourceID)
	case "AWS::OpsWorks::Layer":
		return nil, nil
	case "AWS::OpsWorks::Stack":
		return nil, nil
	case "AWS::OpsWorks::UserProfile":
		return g.getOpsWorksUserProfileAttributes(ctx, physicalResourceID)
	case "AWS::OpsWorks::Volume":
		return nil, nil
	case "AWS::OpsWorksCM::Server":
		return g.getOpsWorksCMServerAttributes(ctx, physicalResourceID)
	case "AWS::Pinpoint::ADMChannel":
		return nil, nil
	case "AWS::Pinpoint::APNSChannel":
		return nil, nil
	case "AWS::Pinpoint::APNSSandboxChannel":
		return nil, nil
	case "AWS::Pinpoint::APNSVoipChannel":
		return nil, nil
	case "AWS::Pinpoint::APNSVoipSandboxChannel":
		return nil, nil
	case "AWS::Pinpoint::App":
		return g.getPinpointAppAttributes(ctx, physicalResourceID)
	case "AWS::Pinpoint::ApplicationSettings":
		return nil, nil
	case "AWS::Pinpoint::BaiduChannel":
		return nil, nil
	case "AWS::Pinpoint::Campaign":
		return g.getPinpointCampaignAttributes(ctx, physicalResourceID)
	case "AWS::Pinpoint::EmailChannel":
		return nil, nil
	case "AWS::Pinpoint::EventStream":
		return nil, nil
	case "AWS::Pinpoint::GCMChannel":
		return nil, nil
	case "AWS::Pinpoint::SMSChannel":
		return nil, nil
	case "AWS::Pinpoint::Segment":
		return g.getPinpointSegmentAttributes(ctx, physicalResourceID)
	case "AWS::Pinpoint::VoiceChannel":
		return nil, nil
	case "AWS::PinpointEmail::ConfigurationSet":
		return nil, nil
	case "AWS::PinpointEmail::ConfigurationSetEventDestination":
		return nil, nil
	case "AWS::PinpointEmail::Identity":
		return g.getPinpointEmailIdentityAttributes(ctx, physicalResourceID)
	case "AWS::QLDB::Ledger":
		return nil, nil
	case "AWS::RAM::ResourceShare":
		return g.getRAMResourceShareAttributes(ctx, physicalResourceID)
	case "AWS::RDS::DBCluster":
		return g.getRDSDBClusterAttributes(ctx, physicalResourceID)
	case "AWS::RDS::DBClusterParameterGroup":
		return nil, nil
	case "AWS::RDS::DBInstance":
		return g.getRDSDBInstanceAttributes(ctx, physicalResourceID)
	case "AWS::RDS::DBParameterGroup":
		return nil, nil
	case "AWS::RDS::DBSecurityGroup":
		return nil, nil
	case "AWS::RDS::DBSecurityGroupIngress":
		return nil, nil
	case "AWS::RDS::DBSubnetGroup":
		return nil, nil
	case "AWS::RDS::EventSubscription":
		return nil, nil
	case "AWS::RDS::OptionGroup":
		return nil, nil
	case "AWS::Redshift::Cluster":
		return g.getRedshiftClusterAttributes(ctx, physicalResourceID)
	case "AWS::Redshift::ClusterParameterGroup":
		return nil, nil
	case "AWS::Redshift::ClusterSecurityGroup":
		return nil, nil
	case "AWS::Redshift::ClusterSecurityGroupIngress":
		return nil, nil
	case "AWS::Redshift::ClusterSubnetGroup":
		return nil, nil
	case "AWS::RoboMaker::Fleet":
		return g.getRoboMakerFleetAttributes(ctx, physicalResourceID)
	case "AWS::RoboMaker::Robot":
		return nil, nil
	case "AWS::RoboMaker::RobotApplication":
		return g.getRoboMakerRobotApplicationAttributes(ctx, physicalResourceID)
	case "AWS::RoboMaker::RobotApplicationVersion":
		return nil, nil
	case "AWS::RoboMaker::SimulationApplication":
		return g.getRoboMakerSimulationApplicationAttributes(ctx, physicalResourceID)
	case "AWS::RoboMaker::SimulationApplicationVersion":
		return nil, nil
	case "AWS::Route53::HealthCheck":
		return nil, nil
	case "AWS::Route53::HostedZone":
		return g.getRoute53HostedZoneAttributes(ctx, physicalResourceID)
	case "AWS::Route53::RecordSet":
		return nil, nil
	case "AWS::Route53::RecordSetGroup":
		return nil, nil
	case "AWS::Route53Resolver::ResolverEndpoint":
		return g.getRoute53ResolverResolverEndpointAttributes(ctx, physicalResourceID)
	case "AWS::Route53Resolver::ResolverRule":
		return g.getRoute53ResolverResolverRuleAttributes(ctx, physicalResourceID)
	case "AWS::Route53Resolver::ResolverRuleAssociation":
		return g.getRoute53ResolverResolverRuleAssociationAttributes(ctx, physicalResourceID)
	case "AWS::S3::Bucket":
		return g.getS3BucketAttributes(ctx, physicalResourceID)
	case "AWS::S3::BucketPolicy":
		return nil, nil
	case "AWS::SDB::Domain":
		return nil, nil
	case "AWS::SES::ConfigurationSet":
		return nil, nil
	case "AWS::SES::ConfigurationSetEventDestination":
		return nil, nil
	case "AWS::SES::ReceiptFilter":
		return nil, nil
	case "AWS::SES::ReceiptRule":
		return nil, nil
	case "AWS::SES::ReceiptRuleSet":
		return nil, nil
	case "AWS::SES::Template":
		return nil, nil
	case "AWS::SNS::Subscription":
		return nil, nil
	case "AWS::SNS::Topic":
		return g.getSNSTopicAttributes(ctx, physicalResourceID)
	case "AWS::SNS::TopicPolicy":
		return nil, nil
	case "AWS::SQS::Queue":
		return g.getSQSQueueAttributes(ctx, physicalResourceID)
	case "AWS::SQS::QueuePolicy":
		return nil, nil
	case "AWS::SSM::Association":
		return nil, nil
	case "AWS::SSM::Document":
		return nil, nil
	case "AWS::SSM::MaintenanceWindow":
		return nil, nil
	case "AWS::SSM::MaintenanceWindowTarget":
		return nil, nil
	case "AWS::SSM::MaintenanceWindowTask":
		return nil, nil
	case "AWS::SSM::Parameter":
		return g.getSSMParameterAttributes(ctx, physicalResourceID)
	case "AWS::SSM::PatchBaseline":
		return nil, nil
	case "AWS::SSM::ResourceDataSync":
		return nil, nil
	case "AWS::SageMaker::CodeRepository":
		return g.getSageMakerCodeRepositoryAttributes(ctx, physicalResourceID)
	case "AWS::SageMaker::Endpoint":
		return g.getSageMakerEndpointAttributes(ctx, physicalResourceID)
	case "AWS::SageMaker::EndpointConfig":
		return g.getSageMakerEndpointConfigAttributes(ctx, physicalResourceID)
	case "AWS::SageMaker::Model":
		return g.getSageMakerModelAttributes(ctx, physicalResourceID)
	case "AWS::SageMaker::NotebookInstance":
		return g.getSageMakerNotebookInstanceAttributes(ctx, physicalResourceID)
	case "AWS::SageMaker::NotebookInstanceLifecycleConfig":
		return g.getSageMakerNotebookInstanceLifecycleConfigAttributes(ctx, physicalResourceID)
	case "AWS::SageMaker::Workteam":
		return g.getSageMakerWorkteamAttributes(ctx, physicalResourceID)
	case "AWS::SecretsManager::ResourcePolicy":
		return nil, nil
	case "AWS::SecretsManager::RotationSchedule":
		return nil, nil
	case "AWS::SecretsManager::Secret":
		return nil, nil
	case "AWS::SecretsManager::SecretTargetAttachment":
		return nil, nil
	case "AWS::SecurityHub::Hub":
		return nil, nil
	case "AWS::ServiceCatalog::AcceptedPortfolioShare":
		return nil, nil
	case "AWS::ServiceCatalog::CloudFormationProduct":
		return g.getServiceCatalogCloudFormationProductAttributes(ctx, physicalResourceID)
	case "AWS::ServiceCatalog::CloudFormationProvisionedProduct":
		return g.getServiceCatalogCloudFormationProvisionedProductAttributes(ctx, physicalResourceID)
	case "AWS::ServiceCatalog::LaunchNotificationConstraint":
		return nil, nil
	case "AWS::ServiceCatalog::LaunchRoleConstraint":
		return nil, nil
	case "AWS::ServiceCatalog::LaunchTemplateConstraint":
		return nil, nil
	case "AWS::ServiceCatalog::Portfolio":
		return g.getServiceCatalogPortfolioAttributes(ctx, physicalResourceID)
	case "AWS::ServiceCatalog::PortfolioPrincipalAssociation":
		return nil, nil
	case "AWS::ServiceCatalog::PortfolioProductAssociation":
		return nil, nil
	case "AWS::ServiceCatalog::PortfolioShare":
		return nil, nil
	case "AWS::ServiceCatalog::ResourceUpdateConstraint":
		return nil, nil
	case "AWS::ServiceCatalog::StackSetConstraint":
		return nil, nil
	case "AWS::ServiceCatalog::TagOption":
		return nil, nil
	case "AWS::ServiceCatalog::TagOptionAssociation":
		return nil, nil
	case "AWS::ServiceDiscovery::HttpNamespace":
		return g.getServiceDiscoveryHttpNamespaceAttributes(ctx, physicalResourceID)
	case "AWS::ServiceDiscovery::Instance":
		return nil, nil
	case "AWS::ServiceDiscovery::PrivateDnsNamespace":
		return g.getServiceDiscoveryPrivateDnsNamespaceAttributes(ctx, physicalResourceID)
	case "AWS::ServiceDiscovery::PublicDnsNamespace":
		return g.getServiceDiscoveryPublicDnsNamespaceAttributes(ctx, physicalResourceID)
	case "AWS::ServiceDiscovery::Service":
		return g.getServiceDiscoveryServiceAttributes(ctx, physicalResourceID)
	case "AWS::StepFunctions::Activity":
		return g.getStepFunctionsActivityAttributes(ctx, physicalResourceID)
	case "AWS::StepFunctions::StateMachine":
		return g.getStepFunctionsStateMachineAttributes(ctx, physicalResourceID)
	case "AWS::Transfer::Server":
		return g.getTransferServerAttributes(ctx, physicalResourceID)
	case "AWS::Transfer::User":
		return g.getTransferUserAttributes(ctx, physicalResourceID)
	case "AWS::WAF::ByteMatchSet":
		return nil, nil
	case "AWS::WAF::IPSet":
		return nil, nil
	case "AWS::WAF::Rule":
		return nil, nil
	case "AWS::WAF::SizeConstraintSet":
		return nil, nil
	case "AWS::WAF::SqlInjectionMatchSet":
		return nil, nil
	case "AWS::WAF::WebACL":
		return nil, nil
	case "AWS::WAF::XssMatchSet":
		return nil, nil
	case "AWS::WAFRegional::ByteMatchSet":
		return nil, nil
	case "AWS::WAFRegional::GeoMatchSet":
		return nil, nil
	case "AWS::WAFRegional::IPSet":
		return nil, nil
	case "AWS::WAFRegional::RateBasedRule":
		return nil, nil
	case "AWS::WAFRegional::RegexPatternSet":
		return nil, nil
	case "AWS::WAFRegional::Rule":
		return nil, nil
	case "AWS::WAFRegional::SizeConstraintSet":
		return nil, nil
	case "AWS::WAFRegional::SqlInjectionMatchSet":
		return nil, nil
	case "AWS::WAFRegional::WebACL":
		return nil, nil
	case "AWS::WAFRegional::WebACLAssociation":
		return nil, nil
	case "AWS::WAFRegional::XssMatchSet":
		return nil, nil
	case "AWS::WorkSpaces::Workspace":
		return nil, nil
	case "Alexa::ASK::Skill":
		return nil, nil
	default:
		return nil, errors.Errorf("unknown resource type '%v'", resourceType)
	}
}
