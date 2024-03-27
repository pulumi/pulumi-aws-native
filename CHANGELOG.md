# Change Log

## 0.100.0 (2024-03-27)

- [Fix replaceOnChanges for Python programs with multi-word properties](https://github.com/pulumi/pulumi-aws-native/pull/1424)
- [Fix handling of AlreadyExists error while creating a resource](https://github.com/pulumi/pulumi-aws-native/pull/1439)
- [Improve cf2pulumi error messages](https://github.com/pulumi/pulumi-aws-native/pull/1438)
- [Skip aws:cdk:path metadata which fails conversion](https://github.com/pulumi/pulumi-aws-native/pull/1444)
- [Fix extension resource](https://github.com/pulumi/pulumi-aws-native/pull/1406)
- [Don't send writeOnly properties if they're also createOnly](https://github.com/pulumi/pulumi-aws-native/pull/1448)
- [Improve autonaming coverage](https://github.com/pulumi/pulumi-aws-native/pull/1443)

## 0.99.0 (2024-03-14)

- [Fix multi-type resolutions](https://github.com/pulumi/pulumi-aws-native/pull/1383) includes fixes for inline role policies
- [Include "id" resource properties as "awsId"](https://github.com/pulumi/pulumi-aws-native/pull/1387)
- [Add more descriptive explanation for missing aws native region config](https://github.com/pulumi/pulumi-aws-native/pull/1355)
- [Send unchanged write-only properties as adds on update](https://github.com/pulumi/pulumi-aws-native/pull/1395)
- [Fix parsing provider config defaultTags](https://github.com/pulumi/pulumi-aws-native/pull/1393)
- [Remove resources not supported by CloudControl from the SDK](https://github.com/pulumi/pulumi-aws-native/pull/1402)
- [Fix updates with secret values](https://github.com/pulumi/pulumi-aws-native/pull/1397)

### Breaking Changes

EC2 Instance properties changed:

- `Id` removed.
- `Affinity` changed from `string` to `InstanceAffinity` enum.
- `CpuOptions` type name changed from `InstanceCpuOptions` to `CpuOptionsProperties`.
- `CreditSpecification` type name changed from `InstanceCreditSpecification` to `CreditSpecificationPropertiesArgs`.
- `EnclaveOptions` type name changed from `InstanceHibernationOptions` to `HibernationOptionsProperties`.
- `HibernationOptions` type name changed from `InstanceHibernationOptions` to `HibernationOptionsProperties`

## 0.98.0 (2024-02-29)

- **BREAKING CHANGE**: [De-duplicate types for simple arrays of tags](https://github.com/pulumi/pulumi-aws-native/pull/1348)
  - See PR for a complete list of resources affected (512)
- [Add types for 9 untyped properties](https://github.com/pulumi/pulumi-aws-native/pull/1365)
- [Fix generating types for refs to map types](https://github.com/pulumi/pulumi-aws-native/pull/1363)
- [Implement defaultTags configuration](https://github.com/pulumi/pulumi-aws-native/issues/1369)
- [Fix precedence of credential sources](https://github.com/pulumi/pulumi-aws-native/pull/1378)
- [Fix handling of write-only properties](https://github.com/pulumi/pulumi-aws-native/pull/1377)

#### Resources
- `🟢` "aws-native:customerprofiles:Domain": required inputs: "defaultExpirationDays" input has changed to Required
- `🟢` "aws-native:customerprofiles:ObjectType": required inputs: "description" input has changed to Required
- "aws-native:dynamodb:Table":
    - `🟡` inputs: "resourcePolicy" missing
    - `🟡` properties: "resourcePolicy" missing output "resourcePolicy"
- `🟢` "aws-native:mediapackagev2:Channel": required inputs: "channelGroupName" input has changed to Required
- "aws-native:mediapackagev2:ChannelPolicy": required inputs:
    - `🟢` "channelGroupName" input has changed to Required
    - `🟢` "channelName" input has changed to Required
- "aws-native:mediapackagev2:OriginEndpoint":
    - `🟢` required: "containerType" property is no longer Required
    - required inputs:
        - `🟢` "channelGroupName" input has changed to Required
        - `🟢` "channelName" input has changed to Required
- "aws-native:mediapackagev2:OriginEndpointPolicy": required inputs:
    - `🟢` "channelGroupName" input has changed to Required
    - `🟢` "channelName" input has changed to Required
    - `🟢` "originEndpointName" input has changed to Required
#### Types
- `🔴` "aws-native:dynamodb:TableResourcePolicy" missing
- `🟡` "aws-native:dynamodb:TableStreamSpecification": properties: "resourcePolicy" missing
- `🟡` "aws-native:iottwinmaker:EntityProperty": properties: "definition" type changed from "#/types/aws-native:iottwinmaker:EntityPropertyDefinitionProperties" to "#/types/aws-native:iottwinmaker:EntityDefinition"
- `🔴` "aws-native:iottwinmaker:EntityPropertyDefinitionProperties" missing

#### New resources:

- `controltower.EnabledBaseline`
- `guardduty.Master`
- `guardduty.Member`

#### New functions:

- `controltower.getEnabledBaseline`
- `guardduty.getMember`
<!-- thollander/actions-comment-pull-request "schemaCheck" -->

## 0.97.0 (2024-02-21)

- [Fix types which are maps](https://github.com/pulumi/pulumi-aws-native/pull/1342)
  - Resources parameters affected:
    - `aws-native:apigateway:Method`: `requestModels`, `requestParameters`
    - `aws-native:apigateway:RestApi`: `parameters`
    - `aws-native:apigateway:Stage`: `variables`
    - `aws-native:apigatewayv2:Api`: `tags`
    - `aws-native:apigatewayv2:DomainName`: `tags`
    - `aws-native:apigatewayv2:VpcLink`: `tags`
    - `aws-native:appconfig:Extension`: `actions`, `parameters`
    - `aws-native:appconfig:ExtensionAssociation`: `parameters`
    - `aws-native:athena:DataCatalog`: `parameters`
    - `aws-native:backup:BackupPlan`: `backupPlanTags`
    - `aws-native:backup:BackupVault`: `backupVaultTags`
    - `aws-native:batch:ComputeEnvironment`: `tags`
    - `aws-native:batch:JobQueue`: `tags`
    - `aws-native:batch:SchedulingPolicy`: `tags`
    - `aws-native:cloudformation:Stack`: `parameters`
    - `aws-native:codestarnotifications:NotificationRule`: `tags`
    - `aws-native:cognito:IdentityPoolRoleAttachment`: `roleMappings`, `roles`
    - `aws-native:cognito:UserPool`: `userPoolTags`
    - `aws-native:cognito:UserPoolUser`: `clientMetadata`
    - `aws-native:eks:Nodegroup`: `labels`, `tags`
    - `aws-native:fis:ExperimentTemplate`: `tags`
    - `aws-native:greengrassv2:Deployment`: `components`, `tags`
    - `aws-native:imagebuilder:Component`: `tags`
    - `aws-native:imagebuilder:ContainerRecipe`: `tags`
    - `aws-native:imagebuilder:DistributionConfiguration`: `tags`
    - `aws-native:imagebuilder:Image`: `tags`
    - `aws-native:imagebuilder:ImagePipeline`: `tags`
    - `aws-native:imagebuilder:ImageRecipe`: `tags`
    - `aws-native:imagebuilder:InfrastructureConfiguration`: `resourceTags`, `tags`
    - `aws-native:imagebuilder:LifecyclePolicy`: `tags`
    - `aws-native:imagebuilder:Workflow`: `tags`
    - `aws-native:iot:Authorizer`: `tokenSigningPublicKeys`
    - `aws-native:iot:SecurityProfile`: `alertTargets`
    - `aws-native:iottwinmaker:ComponentType`: `compositeComponentTypes`, `functions`, `propertyDefinitions`, `propertyGroups`, `tags`
    - `aws-native:iottwinmaker:Entity`: `components`, `compositeComponents`, `tags`
    - `aws-native:iottwinmaker:Scene`: `sceneMetadata`, `tags`
    - `aws-native:iottwinmaker:SyncJob`: `tags`
    - `aws-native:iottwinmaker:Workspace`: `tags`
    - `aws-native:kafkaconnect:Connector`: `connectorConfiguration`
    - `aws-native:msk:Cluster`: `tags`
    - `aws-native:msk:ServerlessCluster`: `tags`
    - `aws-native:opensearchservice:Domain`: `advancedOptions`, `logPublishingOptions`
    - `aws-native:servicecatalog:CloudFormationProvisionedProduct`, `outputs`
    - `aws-native:ssm:Parameter`: `tags`

  - Type properties affected:
    - `aws-native:apigateway:DeploymentCanarySetting`: `stageVariableOverrides`
    - `aws-native:apigateway:DeploymentCanarySettings`: `stageVariableOverrides`
    - `aws-native:apigateway:DeploymentStageDescription`: `variables`
    - `aws-native:apigateway:MethodIntegration`: `requestParameters`, `requestTemplates`
    - `aws-native:apigateway:MethodIntegrationResponse`: `responseParameters`, `responseTemplates`
    - `aws-native:apigateway:MethodResponse`: `responseModels`, `responseParameters`
    - `aws-native:apigateway:StageCanarySetting`: `stageVariableOverrides`
    - `aws-native:apigateway:UsagePlanApiStage`: `throttle`
    - `aws-native:appintegrations:DataIntegrationFileConfiguration`: `filters`
    - `aws-native:backup:BackupPlanBackupRuleResourceType`: `recoveryPointTags`
    - `aws-native:batch:ComputeEnvironmentComputeResources`: `tags`
    - `aws-native:ecs:ServiceLogConfiguration`: `options`
    - `aws-native:ecs:TaskDefinitionContainerDefinition`: `dockerLabels`
    - `aws-native:ecs:TaskDefinitionDockerVolumeConfiguration`: `driverOpts`, `labels`
    - `aws-native:ecs:TaskDefinitionFirelensConfiguration`: `options`
    - `aws-native:ecs:TaskDefinitionLogConfiguration`: `options`
    - `aws-native:elasticloadbalancingv2:ListenerRuleAuthenticateCognitoConfig`: `authenticationRequestExtraParams`
    - `aws-native:elasticloadbalancingv2:ListenerRuleAuthenticateOidcConfig`: `authenticationRequestExtraParams`
    - `aws-native:emrserverless:ApplicationConfigurationObject`: `properties`
    - `aws-native:entityresolution:IdMappingWorkflowProviderProperties`: `providerConfiguration`
    - `aws-native:entityresolution:MatchingWorkflowProviderProperties`: `providerConfiguration`
    - `aws-native:events:RuleHttpParameters`: `headerParameters`, `queryStringParameters`
    - `aws-native:events:RuleInputTransformer`: `inputPathsMap`
    - `aws-native:greengrassv2:ComponentVersionComponentPlatform`: `attributes`
    - `aws-native:greengrassv2:ComponentVersionLambdaExecutionParameters`: `environmentVariables`
    - `aws-native:greengrassv2:ComponentVersionLambdaFunctionRecipeSource`: `componentDependencies`
    - `aws-native:guardduty:FilterFindingCriteria`: `criterion`
    - `aws-native:imagebuilder:DistributionConfigurationAmiDistributionConfiguration`: `amiTags`
    - `aws-native:imagebuilder:LifecyclePolicyAmiExclusionRules`: `tagMap`
    - `aws-native:imagebuilder:LifecyclePolicyExclusionRules`: `tagMap`
    - `aws-native:imagebuilder:LifecyclePolicyResourceSelection`: `tagMap`

### Breaking Changes

#### Resources

- "aws-native:ec2:PrefixList": required: `maxEntries` property is no longer Required

#### New resources:

- `cognito.UserPoolRiskConfigurationAttachment`

#### New functions:

- `cognito.getUserPoolRiskConfigurationAttachment`
- `ec2.getTransitGatewayRouteTableAssociation`


## 0.96.0 (2024-02-09)

- [Added configuration option `aws-native:config:skipCredentialsValidation` (or via environment variable `AWS_SKIP_CREDENTIALS_VALIDATION`)](https://github.com/pulumi/pulumi-aws-native/issues/1326)
- [Upgraded internal dependencies](https://github.com/pulumi/pulumi-aws-native/issues/1328)
- [Implemented custom endpoints configuration](https://github.com/pulumi/pulumi-aws-native/issues/1332)
- [Automated upgrade to latest specifications](https://github.com/pulumi/pulumi-aws-native/pull/1333)

### Breaking Changes

#### Resources

- "aws-native:cognito:UserPoolIdentityProvider": required: "providerDetails" property is no longer Required
- "aws-native:ssmguiconnect:Preferences" removed
- "aws-native:verifiedpermissions:IdentitySource": "policyStoreId" input has changed to Required
- "aws-native:verifiedpermissions:PolicyTemplate": "policyStoreId" input has changed to Required

#### Functions

- "aws-native:ssmguiconnect:getPreferences" removed

#### Types

- "aws-native:wafv2:LoggingConfigurationFieldToMatch": property "jsonBody" removed

### New resources

- `appconfig.Environment`
- `appconfig.HostedConfigurationVersion`
- `ec2.SecurityGroupIngress`
- `ec2.TransitGatewayRouteTableAssociation`

### New functions

- `appconfig.getEnvironment`
- `appconfig.getHostedConfigurationVersion`
- `ec2.getSecurityGroupIngress`

## 0.95.0 (2024-01-25)

### Breaking Changes

#### Resources
- "aws-native:batch:JobDefinition":
    - `🟡` inputs: "timeout" type changed from "#/types/aws-native:batch:JobDefinitionJobTimeout" to "#/types/aws-native:batch:JobDefinitionTimeout"
    - properties:
        - `🟡` "containerOrchestrationType" missing output "containerOrchestrationType"
        - `🟡` "jobDefinitionArn" missing output "jobDefinitionArn"
        - `🟡` "revision" missing output "revision"
        - `🟡` "status" missing output "status"
        - `🟡` "timeout" type changed from "#/types/aws-native:batch:JobDefinitionJobTimeout" to "#/types/aws-native:batch:JobDefinitionTimeout"
- "aws-native:iot:FleetMetric": properties:
    - `🟡` "creationDate" type changed from "number" to "string"
    - `🟡` "lastModifiedDate" type changed from "number" to "string"
#### Functions
- "aws-native:batch:getJobDefinition": inputs:
    - `🟡` "jobDefinitionArn" missing input "jobDefinitionArn"
    - `🟢` required: "id" input has changed to Required
#### Types
- "aws-native:batch:JobDefinitionContainerProperties": properties:
    - `🟡` "ephemeralStorage" type changed from "#/types/aws-native:batch:JobDefinitionContainerPropertiesEphemeralStorageProperties" to "#/types/aws-native:batch:JobDefinitionEphemeralStorage"
    - `🟡` "fargatePlatformConfiguration" type changed from "#/types/aws-native:batch:JobDefinitionContainerPropertiesFargatePlatformConfigurationProperties" to "#/types/aws-native:batch:JobDefinitionFargatePlatformConfiguration"
    - `🟡` "linuxParameters" type changed from "#/types/aws-native:batch:JobDefinitionContainerPropertiesLinuxParametersProperties" to "#/types/aws-native:batch:JobDefinitionLinuxParameters"
    - `🟡` "logConfiguration" type changed from "#/types/aws-native:batch:JobDefinitionContainerPropertiesLogConfigurationProperties" to "#/types/aws-native:batch:JobDefinitionLogConfiguration"
    - `🟡` "mountPoints": items type changed from "#/types/aws-native:batch:JobDefinitionMountPoint" to "#/types/aws-native:batch:JobDefinitionMountPoints"
    - `🟡` "networkConfiguration" type changed from "#/types/aws-native:batch:JobDefinitionContainerPropertiesNetworkConfigurationProperties" to "#/types/aws-native:batch:JobDefinitionNetworkConfiguration"
    - `🟡` "runtimePlatform" type changed from "#/types/aws-native:batch:JobDefinitionContainerPropertiesRuntimePlatformProperties" to "#/types/aws-native:batch:JobDefinitionRuntimePlatform"
    - `🟡` "volumes": items type changed from "#/types/aws-native:batch:JobDefinitionVolume" to "#/types/aws-native:batch:JobDefinitionVolumes"
- `🔴` "aws-native:batch:JobDefinitionContainerPropertiesEphemeralStorageProperties" missing
- `🔴` "aws-native:batch:JobDefinitionContainerPropertiesFargatePlatformConfigurationProperties" missing
- `🔴` "aws-native:batch:JobDefinitionContainerPropertiesLinuxParametersProperties" missing
- `🔴` "aws-native:batch:JobDefinitionContainerPropertiesLogConfigurationProperties" missing
- `🔴` "aws-native:batch:JobDefinitionContainerPropertiesNetworkConfigurationProperties" missing
- `🔴` "aws-native:batch:JobDefinitionContainerPropertiesRuntimePlatformProperties" missing
- `🔴` "aws-native:batch:JobDefinitionEfsAuthorizationConfig" missing
- `🟡` "aws-native:batch:JobDefinitionEfsVolumeConfiguration": properties: "authorizationConfig" type changed from "#/types/aws-native:batch:JobDefinitionEfsAuthorizationConfig" to "#/types/aws-native:batch:JobDefinitionAuthorizationConfig"
- `🔴` "aws-native:batch:JobDefinitionEksMetadata" missing
- `🔴` "aws-native:batch:JobDefinitionEksPodProperties" missing
- `🟡` "aws-native:batch:JobDefinitionEksProperties": properties: "podProperties" type changed from "#/types/aws-native:batch:JobDefinitionEksPodProperties" to "#/types/aws-native:batch:JobDefinitionPodProperties"
- `🔴` "aws-native:batch:JobDefinitionHost" missing
- `🔴` "aws-native:batch:JobDefinitionJobTimeout" missing
- `🔴` "aws-native:batch:JobDefinitionMountPoint" missing
- `🔴` "aws-native:batch:JobDefinitionVolume" missing
- `🟡` "aws-native:elasticache:ServerlessCacheEndpoint": properties: "port" type changed from "integer" to "string"

#### New resources:

- `codebuild.Fleet`
- `cognito.IdentityPoolRoleAttachment`
- `cognito.UserPoolDomain`
- `cognito.UserPoolIdentityProvider`
- `cognito.UserPoolResourceServer`
- `cognito.UserPoolUiCustomizationAttachment`
- `datazone.DataSource`
- `datazone.Domain`
- `datazone.Environment`
- `datazone.EnvironmentBlueprintConfiguration`
- `datazone.EnvironmentProfile`
- `datazone.Project`
- `datazone.SubscriptionTarget`
- `guardduty.Filter`
- `ivs.Stage`
- `ssmguiconnect.Preferences`

#### New functions:

- `codebuild.getFleet`
- `cognito.getIdentityPoolRoleAttachment`
- `cognito.getUserPoolDomain`
- `cognito.getUserPoolIdentityProvider`
- `cognito.getUserPoolResourceServer`
- `cognito.getUserPoolUiCustomizationAttachment`
- `datazone.getDataSource`
- `datazone.getDomain`
- `datazone.getEnvironment`
- `datazone.getEnvironmentBlueprintConfiguration`
- `datazone.getEnvironmentProfile`
- `datazone.getProject`
- `datazone.getSubscriptionTarget`
- `guardduty.getFilter`
- `ivs.getStage`
- `ssmguiconnect.getPreferences`

## 0.94.0 (2024-01-17)

Breaking changes to types:

- "aws-native:cognito:IdentityPoolCognitoIdentityProvider": required:
    - `🟢` "clientId" property has changed to Required
    - `🟢` "providerName" property has changed to Required
- `🟡` "aws-native:connect:RuleActions": properties: "endAssociatedTaskActions" removed

## 0.93.0 (2024-01-12)

### Breaking Changes

#### Resources
- `🟢` "aws-native:ec2:SubnetCidrBlock": required: "ipv6CidrBlock" property is no longer Required
- `🟢` "aws-native:guardduty:IpSet": required: "name" property is no longer Required
- `🟢` "aws-native:guardduty:ThreatIntelSet": required: "name" property is no longer Required
- "aws-native:quicksight:Analysis": required:
    - `🟢` "errors" property is no longer Required
    - `🟢` "sheets" property is no longer Required
- `🟢` "aws-native:verifiedpermissions:Policy": required inputs: "policyStoreId" input has changed to Required
#### Types
- `🟢` "aws-native:imagebuilder:LifecyclePolicyRecipeSelection": required: "semanticVersion" property has changed to Required
- `🟡` "aws-native:quicksight:AnalysisResourcePermission": properties: "resource" missing
- `🟡` "aws-native:quicksight:DashboardResourcePermission": properties: "resource" missing
- `🟡` "aws-native:quicksight:TemplateResourcePermission": properties: "resource" missing
- `🟡` "aws-native:quicksight:ThemeResourcePermission": properties: "resource" missing

#### New resources:

- `ec2.Instance`
- `ec2.SecurityGroup`
- `ssm.PatchBaseline`

#### New functions:

- `ec2.getInstance`
- `ec2.getSecurityGroup`
- `ssm.getPatchBaseline`
  
## 0.92.0 (2023-12-29)

### Breaking Changes

#### Resources
- "aws-native:connect:PhoneNumber": required:
    - 🟢 "countryCode" property is no longer Required
    - 🟢 "type" property is no longer Required
- 🟡 "aws-native:location:Map": properties: "dataSource" missing output "dataSource"
- 🟢 "aws-native:redshift:Cluster": required: "masterUserPassword" property is no longer Required

#### Types
- 🟢 "aws-native:redshift:ClusterLoggingProperties": required: "bucketName" property is no longer Required

### New resources
- `cloudfront.KeyValueStore`
- `connect.PredefinedAttribute`
- `location.ApiKey`
- `neptunegraph.Graph`
- `neptunegraph.PrivateGraphEndpoint`
- `networkfirewall.TlsInspectionConfiguration`

### New functions
- `cloudfront.getKeyValueStore`
- `connect.getPredefinedAttribute`
- `location.getApiKey`
- `neptunegraph.getGraph`
- `neptunegraph.getPrivateGraphEndpoint`
- `networkfirewall.getTlsInspectionConfiguration`

## 0.91.0 (23-12-22)

Fix: Handle PENDING status (<https://github.com/pulumi/pulumi-aws-native/issues/1219>)

### Breaking Changes

#### Resources

- "aws-native:ec2:Subnet":
    - inputs:
        - `🟡` "ipv4NetmaskLength" missing
        - `🟡` "ipv6NetmaskLength" missing
    - properties:
        - `🟡` "ipv4NetmaskLength" missing output "ipv4NetmaskLength"
        - `🟡` "ipv6NetmaskLength" missing output "ipv6NetmaskLength"
    - `🟢` required: "ipv6CidrBlocks" property is no longer Required
- "aws-native:elasticache:ServerlessCache": required:
    - `🟢` "endpoint" property is no longer Required
    - `🟢` "readerEndpoint" property is no longer Required
- "aws-native:lambda:Function":
    - `🟡` inputs: "policy" missing
    - `🟡` properties: "policy" missing output "policy"

#### Functions

- "aws-native:lambda:getEventInvokeConfig": inputs:
    - `🟡` "id" missing input "id"
    - required:
        - `🟢` "functionName" input has changed to Required
        - `🟢` "qualifier" input has changed to Required
- "aws-native:route53resolver:getResolverConfig": inputs:
    - `🟡` "id" missing input "id"
    - `🟢` required: "resourceId" input has changed to Required

#### Types
- "aws-native:autoscaling:AutoScalingGroupInstanceRequirements": required:
    - `🟢` "memoryMiB" property has changed to Required
    - `🟢` "vCpuCount" property has changed to Required
- `🟢` "aws-native:connect:InstanceStorageConfigKinesisVideoStreamConfig": required: "encryptionConfig" property has changed to Required
- `🟢` "aws-native:osis:PipelineLogPublishingOptionsCloudWatchLogDestinationProperties": required: "logGroup" property has changed to Required
- `🟢` "aws-native:osis:PipelineVpcOptions": required: "subnetIds" property has changed to Required

### New resources

- `b2bi.Capability`
- `b2bi.Partnership`
- `b2bi.Profile`
- `b2bi.Transformer`
- `batch.JobDefinition`
- `ce.AnomalyMonitor`
- `ce.AnomalySubscription`
- `dms.DataProvider`
- `dms.InstanceProfile`
- `dms.MigrationProject`
- `ec2.SecurityGroupEgress`
- `ec2.SnapshotBlockPublicAccess`
- `eks.AccessEntry`
- `eventschemas.Discoverer`
- `eventschemas.Registry`
- `eventschemas.Schema`
- `fis.TargetAccountConfiguration`
- `imagebuilder.Workflow`
- `iot.CertificateProvider`
- `securityhub.Hub`

### New functions

- `b2bi.getCapability`
- `b2bi.getPartnership`
- `b2bi.getProfile`
- `b2bi.getTransformer`
- `batch.getJobDefinition`
- `ce.getAnomalyMonitor`
- `ce.getAnomalySubscription`
- `dms.getDataProvider`
- `dms.getInstanceProfile`
- `dms.getMigrationProject`
- `ec2.getSecurityGroupEgress`
- `ec2.getSnapshotBlockPublicAccess`
- `eks.getAccessEntry`
- `eventschemas.getDiscoverer`
- `eventschemas.getRegistry`
- `eventschemas.getSchema`
- `fis.getTargetAccountConfiguration`
- `imagebuilder.getWorkflow`
- `iot.getCertificateProvider`
- `securityhub.getHub`

## 0.90.0 (2023-12-04)

### New resources:

- arczonalshift.ZonalAutoshiftConfiguration
- workspacesthinclient.Environment

### New functions:

- arczonalshift.getZonalAutoshiftConfiguration
- workspacesthinclient.getEnvironment

## 0.89.0 (2023-11-30)

### New resources:

- sagemaker.InferenceComponent

### New functions:

- sagemaker.getInferenceComponent

## 0.88.0 (2023-11-29)

### New resources:

- elasticache.ServerlessCache
- s3express.BucketPolicy
- s3express.DirectoryBucket

### New functions:

- elasticache.getServerlessCache
- s3express.getBucketPolicy
- s3express.getDirectoryBucket

## 0.87.0 (2023-11-28)

### Upstream breaking changes:

- Function "aws-native:ec2:getTransitGatewayVpcAttachment" missing output "addSubnetIds"
- Function "aws-native:ec2:getTransitGatewayVpcAttachment" missing output "removeSubnetIds"

### New resources:

- codestarconnections.RepositoryLink
- codestarconnections.SyncConfiguration
- controltower.LandingZone
- eks.PodIdentityAssociation
- elasticloadbalancingv2.TrustStore
- elasticloadbalancingv2.TrustStoreRevocation
- logs.Delivery
- logs.DeliveryDestination
- logs.DeliverySource
- logs.LogAnomalyDetector
- s3.AccessGrant
- s3.AccessGrantsInstance
- s3.AccessGrantsLocation

### New functions:

- codestarconnections.getRepositoryLink
- codestarconnections.getSyncConfiguration
- controltower.getEnabledControl
- controltower.getLandingZone
- eks.getPodIdentityAssociation
- elasticloadbalancingv2.getTrustStore
- elasticloadbalancingv2.getTrustStoreRevocation
- logs.getDelivery
- logs.getDeliveryDestination
- logs.getDeliverySource
- logs.getLogAnomalyDetector
- s3.getAccessGrant
- s3.getAccessGrantsInstance
- s3.getAccessGrantsLocation

## 0.86.0 (2023-11-21)

### Upstream breaking changes

- Resource "aws-native:eks:Cluster" removed property "accessConfig"

## 0.85.0 (2023-11-18)

### Upstream breaking changes:

- Resource "aws-native:resourceexplorer2:View" input "filters" type changed from "#/types/aws-native:resourceexplorer2:ViewFilters" to "#/types/aws-native:resourceexplorer2:ViewSearchFilter"
- Resource "aws-native:resourceexplorer2:View" output "filters" type changed from "#/types/aws-native:resourceexplorer2:ViewFilters" to "#/types/aws-native:resourceexplorer2:ViewSearchFilter"
- Function "aws-native:resourceexplorer2:getView" output "filters" type changed from "#/types/aws-native:resourceexplorer2:ViewFilters" to "#/types/aws-native:resourceexplorer2:ViewSearchFilter"
- Function "aws-native:appsync:getResolver" missing input "id"
- Function "aws-native:appsync:getResolver" missing output "responseMappingTemplateS3Location"
- Function "aws-native:appsync:getResolver" missing output "codeS3Location"
- Function "aws-native:appsync:getResolver" missing output "requestMappingTemplateS3Location"
- Function "aws-native:appsync:getResolver" missing output "id"
- Type "aws-native:resourceexplorer2:ViewFilters" missing

### New resources:

- `applicationautoscaling.ScalingPolicy`
- `backup.RestoreTestingPlan`
- `backup.RestoreTestingSelection`
- `gamelift.GameSessionQueue`
- `gamelift.MatchmakingConfiguration`
- `gamelift.MatchmakingRuleSet`
- `gamelift.Script`
- `guardduty.IpSet`
- `guardduty.ThreatIntelSet`
- `imagebuilder.LifecyclePolicy`
- `medialive.Multiplex`
- `medialive.Multiplexprogram`
- `opensearchserverless.LifecyclePolicy`
- `s3.StorageLensGroup`

### New functions:

- `applicationautoscaling.getScalingPolicy`
- `backup.getRestoreTestingPlan`
- `backup.getRestoreTestingSelection`
- `gamelift.getGameSessionQueue`
- `gamelift.getMatchmakingConfiguration`
- `gamelift.getMatchmakingRuleSet`
- `gamelift.getScript`
- `guardduty.getIpSet`
- `guardduty.getThreatIntelSet`
- `imagebuilder.getLifecyclePolicy`
- `medialive.getMultiplex`
- `medialive.getMultiplexprogram`
- `opensearchserverless.getLifecyclePolicy`
- `s3.getStorageLensGroup`

## 0.84.0 (2023-11-06)

#### Breaking Changes:

- Function "aws-native:autoscaling:getAutoScalingGroup" missing input "id"
- Function "aws-native:autoscaling:getAutoScalingGroup" missing output "id"
- Function "aws-native:appsync:getResolver" missing input "resolverArn"
- Type "aws-native:autoscaling:AutoScalingGroupNotificationConfiguration" input "topicArn" type changed from "array" to "string"
had &{Type:string Ref: AdditionalProperties: Items: OneOf:[]} but now has no type

#### New resources:

- `appconfig.ConfigurationProfile`

#### New functions:

- `appconfig.getConfigurationProfile`

## 0.83.0 (2023-10-31)

#### Breaking Changes:

Resource "aws-native:autoscaling:AutoScalingGroup" missing input "notificationConfiguration"
Resource "aws-native:autoscaling:AutoScalingGroup" missing output "notificationConfiguration"
Resource "aws-native:events:Rule" input "state" type changed from "string" to "#/types/aws-native:events:RuleState"
Resource "aws-native:events:Rule" input "eventPattern" type changed from "pulumi.json#/Any" to "string"
Resource "aws-native:events:Rule" output "state" type changed from "string" to "#/types/aws-native:events:RuleState"
Resource "aws-native:events:Rule" output "eventPattern" type changed from "pulumi.json#/Any" to "string"
Resource "aws-native:s3:Bucket" input "accessControl" type changed from "string" to "#/types/aws-native:s3:BucketAccessControl"
Resource "aws-native:s3:Bucket" output "accessControl" type changed from "string" to "#/types/aws-native:s3:BucketAccessControl"
Function "aws-native:s3:getBucket" missing input "id"
Function "aws-native:s3:getBucket" missing output "id"
Function "aws-native:s3:getBucket" missing output "accessControl"
Function "aws-native:cognito:getUserPoolClient" missing input "id"
Function "aws-native:cognito:getUserPoolClient" missing output "id"
Function "aws-native:autoscaling:getAutoScalingGroup" missing input "autoScalingGroupName"
Function "aws-native:autoscaling:getAutoScalingGroup" missing output "notificationConfiguration"
Function "aws-native:events:getRule" missing input "id"
Function "aws-native:events:getRule" missing output "id"
Function "aws-native:events:getRule" output "state" type changed from "string" to "#/types/aws-native:events:RuleState"
Function "aws-native:events:getRule" output "eventPattern" type changed from "pulumi.json#/Any" to "string"
Type "aws-native:s3:BucketCorsRule" input "allowedMethods" items type changed from "string" to "#/types/aws-native:s3:BucketCorsRuleAllowedMethodsItem"
Type "aws-native:s3:BucketMetrics" input "status" type changed from "string" to "#/types/aws-native:s3:BucketMetricsStatus"
Type "aws-native:s3:BucketInventoryConfiguration" input "includedObjectVersions" type changed from "string" to "#/types/aws-native:s3:BucketInventoryConfigurationIncludedObjectVersions"
Type "aws-native:s3:BucketInventoryConfiguration" input "optionalFields" items type changed from "string" to "#/types/aws-native:s3:BucketInventoryConfigurationOptionalFieldsItem"
Type "aws-native:s3:BucketInventoryConfiguration" input "scheduleFrequency" type changed from "string" to "#/types/aws-native:s3:BucketInventoryConfigurationScheduleFrequency"
Type "aws-native:s3:BucketIntelligentTieringConfiguration" input "status" type changed from "string" to "#/types/aws-native:s3:BucketIntelligentTieringConfigurationStatus"
Type "aws-native:autoscaling:AutoScalingGroupNotificationConfiguration" input "topicArn" type changed from "string" to "array"
had no type but now has &{Type:string Ref: AdditionalProperties: Items: OneOf:[]}
Type "aws-native:s3:BucketRule" input "objectSizeGreaterThan" type changed from "integer" to "string"
Type "aws-native:s3:BucketRule" input "objectSizeLessThan" type changed from "integer" to "string"
Type "aws-native:s3:BucketRule" input "status" type changed from "string" to "#/types/aws-native:s3:BucketRuleStatus"
Type "aws-native:s3:BucketReplicationTime" input "status" type changed from "string" to "#/types/aws-native:s3:BucketReplicationTimeStatus"
Type "aws-native:s3:BucketVersioningConfiguration" input "status" type changed from "string" to "#/types/aws-native:s3:BucketVersioningConfigurationStatus"
Type "aws-native:s3:BucketReplicaModifications" input "status" type changed from "string" to "#/types/aws-native:s3:BucketReplicaModificationsStatus"
Type "aws-native:s3:BucketDestination" input "format" type changed from "string" to "#/types/aws-native:s3:BucketDestinationFormat"
Type "aws-native:s3:BucketOwnershipControlsRule" input "objectOwnership" type changed from "string" to "#/types/aws-native:s3:BucketOwnershipControlsRuleObjectOwnership"
Type "aws-native:s3:BucketDefaultRetention" input "mode" type changed from "string" to "#/types/aws-native:s3:BucketDefaultRetentionMode"
Type "aws-native:s3:BucketNoncurrentVersionTransition" input "storageClass" type changed from "string" to "#/types/aws-native:s3:BucketNoncurrentVersionTransitionStorageClass"
Type "aws-native:s3:BucketReplicationDestination" input "storageClass" type changed from "string" to "#/types/aws-native:s3:BucketReplicationDestinationStorageClass"
Type "aws-native:s3:BucketReplicationRule" input "status" type changed from "string" to "#/types/aws-native:s3:BucketReplicationRuleStatus"
Type "aws-native:s3:BucketServerSideEncryptionByDefault" input "sseAlgorithm" type changed from "string" to "#/types/aws-native:s3:BucketServerSideEncryptionByDefaultSseAlgorithm"
Type "aws-native:s3:BucketTiering" input "accessTier" type changed from "string" to "#/types/aws-native:s3:BucketTieringAccessTier"
Type "aws-native:s3:BucketSseKmsEncryptedObjects" input "status" type changed from "string" to "#/types/aws-native:s3:BucketSseKmsEncryptedObjectsStatus"
Type "aws-native:s3:BucketDeleteMarkerReplication" input "status" type changed from "string" to "#/types/aws-native:s3:BucketDeleteMarkerReplicationStatus"
Type "aws-native:s3:BucketRedirectAllRequestsTo" input "protocol" type changed from "string" to "#/types/aws-native:s3:BucketRedirectAllRequestsToProtocol"
Type "aws-native:s3:BucketRedirectRule" input "protocol" type changed from "string" to "#/types/aws-native:s3:BucketRedirectRuleProtocol"
Type "aws-native:s3:BucketAccelerateConfiguration" input "accelerationStatus" type changed from "string" to "#/types/aws-native:s3:BucketAccelerateConfigurationAccelerationStatus"
Type "aws-native:s3:BucketTransition" input "storageClass" type changed from "string" to "#/types/aws-native:s3:BucketTransitionStorageClass"

#### New resources:

    appsync.Resolver
    cognito.UserPoolUserToGroupAttachment
    ec2.VpcGatewayAttachment
    iam.User

#### New functions:

    appsync.getResolver
    ec2.getVpcGatewayAttachment
    iam.getUser

## 0.82.1 (2023-10-30)

Update pulumi version to 3.91.1

## 0.82.0 (2023-10-25)

Resource "aws-native:s3:Bucket" input "accessControl" type changed from "#/types/aws-native:s3:BucketAccessControl" to "string"
Resource "aws-native:s3:Bucket" output "accessControl" type changed from "#/types/aws-native:s3:BucketAccessControl" to "string"
Resource "aws-native:events:Rule" input "state" type changed from "#/types/aws-native:events:RuleState" to "string"
Resource "aws-native:events:Rule" input "eventPattern" type changed from "string" to "pulumi.json#/Any"
Resource "aws-native:events:Rule" output "eventPattern" type changed from "string" to "pulumi.json#/Any"
Resource "aws-native:events:Rule" output "state" type changed from "#/types/aws-native:events:RuleState" to "string"
Resource "aws-native:connect:ContactFlowModule" input "state" type changed from "#/types/aws-native:connect:ContactFlowModuleState" to "string"
Resource "aws-native:connect:ContactFlowModule" output "state" type changed from "#/types/aws-native:connect:ContactFlowModuleState" to "string"
Resource "aws-native:connect:ContactFlowModule" output "status" type changed from "#/types/aws-native:connect:ContactFlowModuleStatus" to "string"
Resource "aws-native:cognito:IdentityPool" missing input "identityPoolTags"
Resource "aws-native:cognito:IdentityPool" missing output "identityPoolTags"
Function "aws-native:servicecatalogappregistry:getResourceAssociation" missing input "id"
Function "aws-native:servicecatalogappregistry:getResourceAssociation" missing output "id"
Function "aws-native:cognito:getIdentityPool" missing output "identityPoolTags"
Function "aws-native:iot:getJobTemplate" missing output "destinationPackageVersions"
Function "aws-native:iot:getJobTemplate" missing output "jobExecutionsRetryConfig"
Function "aws-native:iot:getJobTemplate" missing output "maintenanceWindows"
Function "aws-native:servicecatalogappregistry:getAttributeGroupAssociation" missing input "id"
Function "aws-native:servicecatalogappregistry:getAttributeGroupAssociation" missing output "id"
Function "aws-native:s3:getBucket" missing input "bucketName"
Function "aws-native:connect:getContactFlowModule" output "state" type changed from "#/types/aws-native:connect:ContactFlowModuleState" to "string"
Function "aws-native:connect:getContactFlowModule" output "status" type changed from "#/types/aws-native:connect:ContactFlowModuleStatus" to "string"
Function "aws-native:events:getRule" missing input "arn"
Function "aws-native:events:getRule" missing output "eventBusName"
Function "aws-native:events:getRule" output "eventPattern" type changed from "string" to "pulumi.json#/Any"
Function "aws-native:events:getRule" output "state" type changed from "#/types/aws-native:events:RuleState" to "string"
Type "aws-native:s3:BucketReplicationRule" input "status" type changed from "#/types/aws-native:s3:BucketReplicationRuleStatus" to "string"
Type "aws-native:s3:BucketTransitionStorageClass" missing
Type "aws-native:s3:BucketRedirectRuleProtocol" missing
Type "aws-native:s3:BucketReplicaModifications" input "status" type changed from "#/types/aws-native:s3:BucketReplicaModificationsStatus" to "string"
Type "aws-native:s3:BucketIntelligentTieringConfigurationStatus" missing
Type "aws-native:s3:BucketReplicationDestinationStorageClass" missing
Type "aws-native:s3:BucketDestinationFormat" missing
Type "aws-native:s3:BucketCorsRuleAllowedMethodsItem" missing
Type "aws-native:s3:BucketTieringAccessTier" missing
Type "aws-native:s3:BucketDefaultRetentionMode" missing
Type "aws-native:connect:ContactFlowModuleStatus" missing
Type "aws-native:s3:BucketMetrics" input "status" type changed from "#/types/aws-native:s3:BucketMetricsStatus" to "string"
Type "aws-native:s3:BucketMetricsStatus" missing
Type "aws-native:s3:BucketAccelerateConfigurationAccelerationStatus" missing
Type "aws-native:s3:BucketRule" input "status" type changed from "#/types/aws-native:s3:BucketRuleStatus" to "string"
Type "aws-native:s3:BucketRule" input "objectSizeLessThan" type changed from "string" to "integer"
Type "aws-native:s3:BucketRule" input "objectSizeGreaterThan" type changed from "string" to "integer"
Type "aws-native:connect:ContactFlowModuleState" missing
Type "aws-native:s3:BucketReplicationTime" input "status" type changed from "#/types/aws-native:s3:BucketReplicationTimeStatus" to "string"
Type "aws-native:s3:BucketCorsRule" input "allowedMethods" items type changed from "#/types/aws-native:s3:BucketCorsRuleAllowedMethodsItem" to "string"
Type "aws-native:s3:BucketReplicaModificationsStatus" missing
Type "aws-native:s3:BucketReplicationDestination" input "storageClass" type changed from "#/types/aws-native:s3:BucketReplicationDestinationStorageClass" to "string"
Type "aws-native:s3:BucketInventoryConfiguration" input "includedObjectVersions" type changed from "#/types/aws-native:s3:BucketInventoryConfigurationIncludedObjectVersions" to "string"
Type "aws-native:s3:BucketInventoryConfiguration" input "optionalFields" items type changed from "#/types/aws-native:s3:BucketInventoryConfigurationOptionalFieldsItem" to "string"
Type "aws-native:s3:BucketInventoryConfiguration" input "scheduleFrequency" type changed from "#/types/aws-native:s3:BucketInventoryConfigurationScheduleFrequency" to "string"
Type "aws-native:s3:BucketOwnershipControlsRule" input "objectOwnership" type changed from "#/types/aws-native:s3:BucketOwnershipControlsRuleObjectOwnership" to "string"
Type "aws-native:s3:BucketNoncurrentVersionTransitionStorageClass" missing
Type "aws-native:s3:BucketReplicationRuleStatus" missing
Type "aws-native:s3:BucketAccelerateConfiguration" input "accelerationStatus" type changed from "#/types/aws-native:s3:BucketAccelerateConfigurationAccelerationStatus" to "string"
Type "aws-native:s3:BucketSseKmsEncryptedObjects" input "status" type changed from "#/types/aws-native:s3:BucketSseKmsEncryptedObjectsStatus" to "string"
Type "aws-native:s3:BucketTransition" input "storageClass" type changed from "#/types/aws-native:s3:BucketTransitionStorageClass" to "string"
Type "aws-native:s3:BucketDeleteMarkerReplicationStatus" missing
Type "aws-native:s3:BucketDestination" input "format" type changed from "#/types/aws-native:s3:BucketDestinationFormat" to "string"
Type "aws-native:s3:BucketRuleStatus" missing
Type "aws-native:s3:BucketSseKmsEncryptedObjectsStatus" missing
Type "aws-native:s3:BucketDefaultRetention" input "mode" type changed from "#/types/aws-native:s3:BucketDefaultRetentionMode" to "string"
Type "aws-native:s3:BucketRedirectAllRequestsToProtocol" missing
Type "aws-native:events:RuleRedshiftDataParameters" missing property "sqls"
Type "aws-native:s3:BucketTiering" input "accessTier" type changed from "#/types/aws-native:s3:BucketTieringAccessTier" to "string"
Type "aws-native:s3:BucketDeleteMarkerReplication" input "status" type changed from "#/types/aws-native:s3:BucketDeleteMarkerReplicationStatus" to "string"
Type "aws-native:s3:BucketServerSideEncryptionByDefault" input "sseAlgorithm" type changed from "#/types/aws-native:s3:BucketServerSideEncryptionByDefaultSseAlgorithm" to "string"
Type "aws-native:s3:BucketServerSideEncryptionByDefaultSseAlgorithm" missing
Type "aws-native:s3:BucketAccessControl" missing
Type "aws-native:s3:BucketRedirectAllRequestsTo" input "protocol" type changed from "#/types/aws-native:s3:BucketRedirectAllRequestsToProtocol" to "string"
Type "aws-native:s3:BucketInventoryConfigurationOptionalFieldsItem" missing
Type "aws-native:s3:BucketNoncurrentVersionTransition" input "storageClass" type changed from "#/types/aws-native:s3:BucketNoncurrentVersionTransitionStorageClass" to "string"
Type "aws-native:s3:BucketVersioningConfigurationStatus" missing
Type "aws-native:s3:BucketRedirectRule" input "protocol" type changed from "#/types/aws-native:s3:BucketRedirectRuleProtocol" to "string"
Type "aws-native:s3:BucketInventoryConfigurationScheduleFrequency" missing
Type "aws-native:s3:BucketIntelligentTieringConfiguration" input "status" type changed from "#/types/aws-native:s3:BucketIntelligentTieringConfigurationStatus" to "string"
Type "aws-native:s3:BucketInventoryConfigurationIncludedObjectVersions" missing
Type "aws-native:events:RuleState" missing
Type "aws-native:s3:BucketReplicationTimeStatus" missing
Type "aws-native:s3:BucketOwnershipControlsRuleObjectOwnership" missing
Type "aws-native:s3:BucketVersioningConfiguration" input "status" type changed from "#/types/aws-native:s3:BucketVersioningConfigurationStatus" to "string"
Type "aws-native:cognito:IdentityPoolTag" missing

### New resources

- `appconfig.Application`
- `appsync.FunctionConfiguration`
- `cognito.UserPool`
- `cognito.UserPoolClient`
- `cognito.UserPoolGroup`
- `entityresolution.IdMappingWorkflow`
- `iam.Group`
- `msk.Replicator`

### New functions

- `appconfig.getApplication`
- `appsync.getFunctionConfiguration`
- `cognito.getUserPool`
- `cognito.getUserPoolClient`
- `cognito.getUserPoolGroup`
- `entityresolution.getIdMappingWorkflow`
- `iam.getGroup`
- `msk.getReplicator`

## 0.81.0 (2023-10-16)

Resource "aws-native:autoscaling:AutoScalingGroup" property "instanceMaintenancePolicy" removed
Type "aws-native:kendra:DataSourceConfiguration" property "templateConfiguration" removed

### New resources

- `cognito.IdentityPool`
- `events.Rule`
- `lambda.Version`

### New functions

- `cognito.getIdentityPool`
- `events.getRule`
- `lambda.getVersion`

## 0.80.0 (2023-10-10)

Function "aws-native:connect:getPhoneNumber" output "description" removed
Function "aws-native:s3:getBucketPolicy" property "id" removed

### New resources

- `apigatewayv2.DomainName`
- `cognito.LogDeliveryConfiguration`
- `iot.SoftwarePackage`
- `iot.SoftwarePackageVersion`

### New functions

- `apigatewayv2.getDomainName`
- `cognito.getLogDeliveryConfiguration`
- `iot.getSoftwarePackage`
- `iot.getSoftwarePackageVersion`

## 0.79.0 (2023-10-02)

### New resources:
- `cognito.UserPoolUser`
- `events.EventBus`
- `events.Rule`
- `ssm.Parameter`

### New functions:
- `events.getEventBus`
- `events.getRule`
- `ssm.getParameter`


## 0.75.0 (2023-08-28)

### Breaking Changes

Function "aws-native:ec2:getPlacementGroup" missing output "tags"

### New resources:

entityresolution.MatchingWorkflow
workspacesweb.BrowserSettings
workspacesweb.IdentityProvider
workspacesweb.IpAccessSettings
workspacesweb.NetworkSettings
workspacesweb.Portal
workspacesweb.TrustStore
workspacesweb.UserAccessLoggingSettings
workspacesweb.UserSettings

### New functions:

entityresolution.getMatchingWorkflow
workspacesweb.getBrowserSettings
workspacesweb.getIdentityProvider
workspacesweb.getIpAccessSettings
workspacesweb.getNetworkSettings
workspacesweb.getPortal
workspacesweb.getTrustStore
workspacesweb.getUserAccessLoggingSettings
workspacesweb.getUserSettings

## Misc

Began marking createOnlyProperties with the replaceOnChanges flag

## 0.74.0 (2023-08-22)

### Breaking Changes

Function "aws-native:lambda:getFunction" missing output "snapStart"
Function "aws-native:quicksight:getAnalysis" missing output "status"
Function "aws-native:ec2:getEipAssociation" missing output "networkInterfaceId"
Function "aws-native:ec2:getEipAssociation" missing output "privateIpAddress"
Function "aws-native:ec2:getEipAssociation" missing output "allocationId"
Function "aws-native:ec2:getEipAssociation" missing output "eip"
Function "aws-native:ec2:getEipAssociation" missing output "instanceId"
Function "aws-native:iam:getManagedPolicy" missing input "id"
Function "aws-native:iam:getManagedPolicy" missing output "id"

### New resources

cloudformation.Stack
datasync.LocationAzureBlob
ec2.InstanceConnectEndpoint
emr.WalWorkspace
entityresolution.SchemaMapping
route53resolver.OutpostResolver

### New functions

cloudformation.getStack
datasync.getLocationAzureBlob
ec2.getInstanceConnectEndpoint
emr.getWalWorkspace
entityresolution.getSchemaMapping
route53resolver.getOutpostResolver

## 0.72.0 (2023-08-04)

This release includes a large number of breaking changes as we now UPPERCASE initialisms
to TitleCase in all property/type/module/resource names. 


### Breaking Changes

Resource "aws-native:wafv2:WebACLAssociation" missing
Resource "aws-native:ec2:IPAMResourceDiscovery" missing
Resource "aws-native:ec2:IPAMScope" missing
Resource "aws-native:ec2:VPNConnection" missing
Resource "aws-native:directoryservice:SimpleAD" missing
Resource "aws-native:dynamodb:GlobalTable" input "sseSpecification" type changed from "#/types/aws-native:dynamodb:GlobalTableSSESpecification" to "#/types/aws-native:dynamodb:GlobalTableSseSpecification"
Resource "aws-native:dynamodb:GlobalTable" output "sseSpecification" type changed from "#/types/aws-native:dynamodb:GlobalTableSSESpecification" to "#/types/aws-native:dynamodb:GlobalTableSseSpecification"
Resource "aws-native:iam:OIDCProvider" missing
Resource "aws-native:datasync:LocationNFS" missing
Resource "aws-native🛡️DRTAccess" missing
Resource "aws-native:datasync:LocationS3" input "s3Config" type changed from "#/types/aws-native:datasync:LocationS3S3Config" to "#/types/aws-native:datasync:LocationS3s3Config"
Resource "aws-native:datasync:LocationS3" output "s3Config" type changed from "#/types/aws-native:datasync:LocationS3S3Config" to "#/types/aws-native:datasync:LocationS3s3Config"
Resource "aws-native:synthetics:Canary" input "vpcConfig" type changed from "#/types/aws-native:synthetics:CanaryVPCConfig" to "#/types/aws-native:synthetics:CanaryVpcConfig"
Resource "aws-native:synthetics:Canary" output "vpcConfig" type changed from "#/types/aws-native:synthetics:CanaryVPCConfig" to "#/types/aws-native:synthetics:CanaryVpcConfig"
Resource "aws-native:datasync:LocationEFS" missing
Resource "aws-native:ec2:VPNGateway" missing
Resource "aws-native:ec2:EC2Fleet" missing
Resource "aws-native:rds:DBSubnetGroup" missing
Resource "aws-native:route53resolver:ResolverDNSSECConfig" missing
Resource "aws-native:ec2:VPC" missing
Resource "aws-native:rds:DBProxyEndpoint" missing
Resource "aws-native:ec2:VPCPeeringConnection" missing
Resource "aws-native:iotwireless:WirelessDevice" input "loRaWan" type changed from "#/types/aws-native:iotwireless:WirelessDeviceLoRaWANDevice" to "#/types/aws-native:iotwireless:WirelessDeviceLoRaWanDevice"
Resource "aws-native:iotwireless:WirelessDevice" output "loRaWan" type changed from "#/types/aws-native:iotwireless:WirelessDeviceLoRaWANDevice" to "#/types/aws-native:iotwireless:WirelessDeviceLoRaWanDevice"
Resource "aws-native:configuration:ConformancePack" input "templateSsmDocumentDetails" type changed from "#/types/aws-native:configuration:TemplateSSMDocumentDetailsProperties" to "#/types/aws-native:configuration:TemplateSsmDocumentDetailsProperties"
Resource "aws-native:configuration:ConformancePack" output "templateSsmDocumentDetails" type changed from "#/types/aws-native:configuration:TemplateSSMDocumentDetailsProperties" to "#/types/aws-native:configuration:TemplateSsmDocumentDetailsProperties"
Resource "aws-native:ec2:EIPAssociation" missing
Resource "aws-native:apigatewayv2:Authorizer" input "jwtConfiguration" type changed from "#/types/aws-native:apigatewayv2:AuthorizerJWTConfiguration" to "#/types/aws-native:apigatewayv2:AuthorizerJwtConfiguration"
Resource "aws-native:apigatewayv2:Authorizer" output "jwtConfiguration" type changed from "#/types/aws-native:apigatewayv2:AuthorizerJWTConfiguration" to "#/types/aws-native:apigatewayv2:AuthorizerJwtConfiguration"
Resource "aws-native:datasync:LocationFSxOpenZFS" missing
Resource "aws-native:rolesanywhere:CRL" missing
Resource "aws-native:iotwireless:MulticastGroup" input "loRaWan" type changed from "#/types/aws-native:iotwireless:MulticastGroupLoRaWAN" to "#/types/aws-native:iotwireless:MulticastGroupLoRaWan"
Resource "aws-native:iotwireless:MulticastGroup" output "loRaWan" type changed from "#/types/aws-native:iotwireless:MulticastGroupLoRaWAN" to "#/types/aws-native:iotwireless:MulticastGroupLoRaWan"
Resource "aws-native:rds:DBProxy" missing
Resource "aws-native:datasync:LocationFSxONTAP" missing
Resource "aws-native:omics:AnnotationStore" input "storeOptions" type changed from "#/types/aws-native:omics:StoreOptionsProperties" to "#/types/aws-native:omics:AnnotationStoreStoreOptionsProperties"
Resource "aws-native:omics:AnnotationStore" output "storeOptions" type changed from "#/types/aws-native:omics:StoreOptionsProperties" to "#/types/aws-native:omics:AnnotationStoreStoreOptionsProperties"
Resource "aws-native:lambda:EventSourceMapping" input "documentDbEventSourceConfig" type changed from "#/types/aws-native:lambda:EventSourceMappingDocumentDBEventSourceConfig" to "#/types/aws-native:lambda:EventSourceMappingDocumentDbEventSourceConfig"
Resource "aws-native:lambda:EventSourceMapping" output "documentDbEventSourceConfig" type changed from "#/types/aws-native:lambda:EventSourceMappingDocumentDBEventSourceConfig" to "#/types/aws-native:lambda:EventSourceMappingDocumentDbEventSourceConfig"
Resource "aws-native:datasync:LocationHDFS" missing
Resource "aws-native:ec2:EIP" missing
Resource "aws-native:datasync:LocationSMB" missing
Resource "aws-native:rds:DBCluster" missing
Resource "aws-native:route53:HostedZone" input "vpcs" items type changed from "#/types/aws-native:route53:HostedZoneVPC" to "#/types/aws-native:route53:HostedZoneVpc"
Resource "aws-native:route53:HostedZone" output "vpcs" items type changed from "#/types/aws-native:route53:HostedZoneVPC" to "#/types/aws-native:route53:HostedZoneVpc"
Resource "aws-native:rds:DBProxyTargetGroup" missing
Resource "aws-native:route53:DNSSEC" missing
Resource "aws-native:ec2:DHCPOptions" missing
Resource "aws-native:ec2:IPAM" missing
Resource "aws-native:quicksight:VPCConnection" missing
Resource "aws-native:networkmanager:Device" input "awsLocation" type changed from "#/types/aws-native:networkmanager:DeviceAWSLocation" to "#/types/aws-native:networkmanager:DeviceAwsLocation"
Resource "aws-native:networkmanager:Device" output "awsLocation" type changed from "#/types/aws-native:networkmanager:DeviceAWSLocation" to "#/types/aws-native:networkmanager:DeviceAwsLocation"
Resource "aws-native:neptune:DBCluster" missing
Resource "aws-native:ec2:LocalGatewayRouteTableVPCAssociation" missing
Resource "aws-native:lakeformation:TagAssociation" input "lfTags" items type changed from "#/types/aws-native:lakeformation:TagAssociationLFTagPair" to "#/types/aws-native:lakeformation:TagAssociationLfTagPair"
Resource "aws-native:lakeformation:TagAssociation" output "lfTags" items type changed from "#/types/aws-native:lakeformation:TagAssociationLFTagPair" to "#/types/aws-native:lakeformation:TagAssociationLfTagPair"
Resource "aws-native:auditmanager:Assessment" input "awsAccount" type changed from "#/types/aws-native:auditmanager:AssessmentAWSAccount" to "#/types/aws-native:auditmanager:AssessmentAwsAccount"
Resource "aws-native:auditmanager:Assessment" output "awsAccount" type changed from "#/types/aws-native:auditmanager:AssessmentAWSAccount" to "#/types/aws-native:auditmanager:AssessmentAwsAccount"
Resource "aws-native:fms:Policy" input "includeMap" type changed from "#/types/aws-native:fms:PolicyIEMap" to "#/types/aws-native:fms:PolicyIeMap"
Resource "aws-native:fms:Policy" input "excludeMap" type changed from "#/types/aws-native:fms:PolicyIEMap" to "#/types/aws-native:fms:PolicyIeMap"
Resource "aws-native:fms:Policy" output "includeMap" type changed from "#/types/aws-native:fms:PolicyIEMap" to "#/types/aws-native:fms:PolicyIeMap"
Resource "aws-native:fms:Policy" output "excludeMap" type changed from "#/types/aws-native:fms:PolicyIEMap" to "#/types/aws-native:fms:PolicyIeMap"
Resource "aws-native:dynamodb:Table" input "sseSpecification" type changed from "#/types/aws-native:dynamodb:TableSSESpecification" to "#/types/aws-native:dynamodb:TableSseSpecification"
Resource "aws-native:dynamodb:Table" output "sseSpecification" type changed from "#/types/aws-native:dynamodb:TableSSESpecification" to "#/types/aws-native:dynamodb:TableSseSpecification"
Resource "aws-native:ec2:VPNConnectionRoute" missing
Resource "aws-native:iotwireless:WirelessGateway" input "loRaWan" type changed from "#/types/aws-native:iotwireless:WirelessGatewayLoRaWANGateway" to "#/types/aws-native:iotwireless:WirelessGatewayLoRaWanGateway"
Resource "aws-native:iotwireless:WirelessGateway" output "loRaWan" type changed from "#/types/aws-native:iotwireless:WirelessGatewayLoRaWANGateway" to "#/types/aws-native:iotwireless:WirelessGatewayLoRaWanGateway"
Resource "aws-native:devicefarm:VPCEConfiguration" missing
Resource "aws-native:iam:VirtualMFADevice" missing
Resource "aws-native:iotwireless:FuotaTask" input "loRaWan" type changed from "#/types/aws-native:iotwireless:FuotaTaskLoRaWAN" to "#/types/aws-native:iotwireless:FuotaTaskLoRaWan"
Resource "aws-native:iotwireless:FuotaTask" output "loRaWan" type changed from "#/types/aws-native:iotwireless:FuotaTaskLoRaWAN" to "#/types/aws-native:iotwireless:FuotaTaskLoRaWan"
Resource "aws-native:ec2:VPCEndpointServicePermissions" missing
Resource "aws-native:rds:CustomDBEngineVersion" missing
Resource "aws-native:ec2:VPCDHCPOptionsAssociation" missing
Resource "aws-native:rds:DBClusterParameterGroup" missing
Resource "aws-native:memorydb:ACL" missing
Resource "aws-native:rds:DBInstance" missing
Resource "aws-native:healthlake:FHIRDatastore" missing
Resource "aws-native:iam:SAMLProvider" missing
Resource "aws-native:opensearchservice:Domain" input "ebsOptions" type changed from "#/types/aws-native:opensearchservice:DomainEBSOptions" to "#/types/aws-native:opensearchservice:DomainEbsOptions"
Resource "aws-native:opensearchservice:Domain" input "vpcOptions" type changed from "#/types/aws-native:opensearchservice:DomainVPCOptions" to "#/types/aws-native:opensearchservice:DomainVpcOptions"
Resource "aws-native:opensearchservice:Domain" output "ebsOptions" type changed from "#/types/aws-native:opensearchservice:DomainEBSOptions" to "#/types/aws-native:opensearchservice:DomainEbsOptions"
Resource "aws-native:opensearchservice:Domain" output "vpcOptions" type changed from "#/types/aws-native:opensearchservice:DomainVPCOptions" to "#/types/aws-native:opensearchservice:DomainVpcOptions"
Resource "aws-native:iot:CACertificate" missing
Resource "aws-native:ec2:IPAMAllocation" missing
Resource "aws-native:ec2:VPCEndpoint" missing
Resource "aws-native:iotwireless:TaskDefinition" input "loRaWanUpdateGatewayTaskEntry" type changed from "#/types/aws-native:iotwireless:TaskDefinitionLoRaWANUpdateGatewayTaskEntry" to "#/types/aws-native:iotwireless:TaskDefinitionLoRaWanUpdateGatewayTaskEntry"
Resource "aws-native:iotwireless:TaskDefinition" output "loRaWanUpdateGatewayTaskEntry" type changed from "#/types/aws-native:iotwireless:TaskDefinitionLoRaWANUpdateGatewayTaskEntry" to "#/types/aws-native:iotwireless:TaskDefinitionLoRaWanUpdateGatewayTaskEntry"
Resource "aws-native:rds:DBParameterGroup" missing
Resource "aws-native:ec2:IPAMResourceDiscoveryAssociation" missing
Resource "aws-native:iotwireless:ServiceProfile" input "loRaWan" type changed from "#/types/aws-native:iotwireless:ServiceProfileLoRaWANServiceProfile" to "#/types/aws-native:iotwireless:ServiceProfileLoRaWanServiceProfile"
Resource "aws-native:iotwireless:ServiceProfile" output "loRaWan" type changed from "#/types/aws-native:iotwireless:ServiceProfileLoRaWANServiceProfile" to "#/types/aws-native:iotwireless:ServiceProfileLoRaWanServiceProfile"
Resource "aws-native:amplifyuibuilder:Form" input "cta" type changed from "#/types/aws-native:amplifyuibuilder:FormCTA" to "#/types/aws-native:amplifyuibuilder:FormCta"
Resource "aws-native:amplifyuibuilder:Form" output "cta" type changed from "#/types/aws-native:amplifyuibuilder:FormCTA" to "#/types/aws-native:amplifyuibuilder:FormCta"
Resource "aws-native:ec2:IPAMPool" missing
Resource "aws-native:ec2:IPAMPoolCidr" missing
Resource "aws-native:iotwireless:DeviceProfile" input "loRaWan" type changed from "#/types/aws-native:iotwireless:DeviceProfileLoRaWANDeviceProfile" to "#/types/aws-native:iotwireless:DeviceProfileLoRaWanDeviceProfile"
Resource "aws-native:iotwireless:DeviceProfile" output "loRaWan" type changed from "#/types/aws-native:iotwireless:DeviceProfileLoRaWANDeviceProfile" to "#/types/aws-native:iotwireless:DeviceProfileLoRaWanDeviceProfile"
Resource "aws-native:wafv2:WebACL" missing
Resource "aws-native:wafv2:IPSet" missing
Resource "aws-native:ec2:VPCEndpointService" missing
Function "aws-native:rds:getDBProxyEndpoint" missing
Function "aws-native:ec2:getIPAMResourceDiscovery" missing
Function "aws-native:iam:getOIDCProvider" missing
Function "aws-native:rds:getCustomDBEngineVersion" missing
Function "aws-native:networkmanager:getDevice" output "awsLocation" type changed from "#/types/aws-native:networkmanager:DeviceAWSLocation" to "#/types/aws-native:networkmanager:DeviceAwsLocation"
Function "aws-native:amplifyuibuilder:getForm" output "cta" type changed from "#/types/aws-native:amplifyuibuilder:FormCTA" to "#/types/aws-native:amplifyuibuilder:FormCta"
Function "aws-native:ec2:getDHCPOptions" missing
Function "aws-native:ec2:getEC2Fleet" missing
Function "aws-native:rds:getDBClusterParameterGroup" missing
Function "aws-native:iotwireless:getServiceProfile" output "loRaWan" type changed from "#/types/aws-native:iotwireless:ServiceProfileLoRaWANServiceProfile" to "#/types/aws-native:iotwireless:ServiceProfileLoRaWanServiceProfile"
Function "aws-native:rds:getDBParameterGroup" missing
Function "aws-native:ec2:getIPAMPool" missing
Function "aws-native:rolesanywhere:getCRL" missing
Function "aws-native:route53:getHostedZone" output "vpcs" items type changed from "#/types/aws-native:route53:HostedZoneVPC" to "#/types/aws-native:route53:HostedZoneVpc"
Function "aws-native:iotwireless:getTaskDefinition" output "loRaWanUpdateGatewayTaskEntry" type changed from "#/types/aws-native:iotwireless:TaskDefinitionLoRaWANUpdateGatewayTaskEntry" to "#/types/aws-native:iotwireless:TaskDefinitionLoRaWanUpdateGatewayTaskEntry"
Function "aws-native:dynamodb:getTable" output "sseSpecification" type changed from "#/types/aws-native:dynamodb:TableSSESpecification" to "#/types/aws-native:dynamodb:TableSseSpecification"
Function "aws-native:iam:getVirtualMFADevice" missing
Function "aws-native:ec2:getVPCEndpoint" missing
Function "aws-native:ec2:getIPAMResourceDiscoveryAssociation" missing
Function "aws-native:ec2:getVPCEndpointServicePermissions" missing
Function "aws-native:opensearchservice:getDomain" output "ebsOptions" type changed from "#/types/aws-native:opensearchservice:DomainEBSOptions" to "#/types/aws-native:opensearchservice:DomainEbsOptions"
Function "aws-native:opensearchservice:getDomain" output "vpcOptions" type changed from "#/types/aws-native:opensearchservice:DomainVPCOptions" to "#/types/aws-native:opensearchservice:DomainVpcOptions"
Function "aws-native:healthlake:getFHIRDatastore" missing
Function "aws-native:rds:getDBProxy" missing
Function "aws-native:ec2:getEIP" missing
Function "aws-native:datasync:getLocationFSxONTAP" missing
Function "aws-native:iotwireless:getWirelessGateway" output "loRaWan" type changed from "#/types/aws-native:iotwireless:WirelessGatewayLoRaWANGateway" to "#/types/aws-native:iotwireless:WirelessGatewayLoRaWanGateway"
Function "aws-native:ec2:getIPAMAllocation" missing
Function "aws-native:ec2:getVPCPeeringConnection" missing
Function "aws-native:route53resolver:getResolverDNSSECConfig" missing
Function "aws-native:apigatewayv2:getAuthorizer" output "jwtConfiguration" type changed from "#/types/aws-native:apigatewayv2:AuthorizerJWTConfiguration" to "#/types/aws-native:apigatewayv2:AuthorizerJwtConfiguration"
Function "aws-native:ec2:getIPAMPoolCidr" missing
Function "aws-native:datasync:getLocationSMB" missing
Function "aws-native🛡️getDRTAccess" missing
Function "aws-native:iotwireless:getMulticastGroup" output "loRaWan" type changed from "#/types/aws-native:iotwireless:MulticastGroupLoRaWAN" to "#/types/aws-native:iotwireless:MulticastGroupLoRaWan"
Function "aws-native:ec2:getVPCEndpointService" missing
Function "aws-native:iam:getSAMLProvider" missing
Function "aws-native:rds:getDBInstance" missing
Function "aws-native:synthetics:getCanary" output "vpcConfig" type changed from "#/types/aws-native:synthetics:CanaryVPCConfig" to "#/types/aws-native:synthetics:CanaryVpcConfig"
Function "aws-native:rds:getDBCluster" missing
Function "aws-native:iotwireless:getFuotaTask" output "loRaWan" type changed from "#/types/aws-native:iotwireless:FuotaTaskLoRaWAN" to "#/types/aws-native:iotwireless:FuotaTaskLoRaWan"
Function "aws-native:ec2:getVPC" missing
Function "aws-native:neptune:getDBCluster" missing
Function "aws-native:devicefarm:getVPCEConfiguration" missing
Function "aws-native:dynamodb:getGlobalTable" output "sseSpecification" type changed from "#/types/aws-native:dynamodb:GlobalTableSSESpecification" to "#/types/aws-native:dynamodb:GlobalTableSseSpecification"
Function "aws-native:iotwireless:getWirelessDevice" output "loRaWan" type changed from "#/types/aws-native:iotwireless:WirelessDeviceLoRaWANDevice" to "#/types/aws-native:iotwireless:WirelessDeviceLoRaWanDevice"
Function "aws-native:ec2:getEIPAssociation" missing
Function "aws-native:iotwireless:getDeviceProfile" output "loRaWan" type changed from "#/types/aws-native:iotwireless:DeviceProfileLoRaWANDeviceProfile" to "#/types/aws-native:iotwireless:DeviceProfileLoRaWanDeviceProfile"
Function "aws-native:iot:getCACertificate" missing
Function "aws-native:wafv2:getWebACL" missing
Function "aws-native:datasync:getLocationHDFS" missing
Function "aws-native:ec2:getIPAM" missing
Function "aws-native:rds:getDBProxyTargetGroup" missing
Function "aws-native:datasync:getLocationFSxOpenZFS" missing
Function "aws-native:ec2:getVPNConnection" missing
Function "aws-native:memorydb:getACL" missing
Function "aws-native:datasync:getLocationNFS" missing
Function "aws-native:directoryservice:getSimpleAD" missing
Function "aws-native:wafv2:getIPSet" missing
Function "aws-native:rds:getDBSubnetGroup" missing
Function "aws-native:fms:getPolicy" output "excludeMap" type changed from "#/types/aws-native:fms:PolicyIEMap" to "#/types/aws-native:fms:PolicyIeMap"
Function "aws-native:fms:getPolicy" output "includeMap" type changed from "#/types/aws-native:fms:PolicyIEMap" to "#/types/aws-native:fms:PolicyIeMap"
Function "aws-native:ec2:getLocalGatewayRouteTableVPCAssociation" missing
Function "aws-native:ec2:getIPAMScope" missing
Function "aws-native:ec2:getVPNGateway" missing
Function "aws-native:lambda:getEventSourceMapping" output "documentDbEventSourceConfig" type changed from "#/types/aws-native:lambda:EventSourceMappingDocumentDBEventSourceConfig" to "#/types/aws-native:lambda:EventSourceMappingDocumentDbEventSourceConfig"
Function "aws-native:quicksight:getVPCConnection" missing
Function "aws-native:datasync:getLocationEFS" missing
Type "aws-native:quicksight:ThemeUIColorPalette" missing
Type "aws-native:ec2:EC2FleetAcceleratorCountRequest" missing
Type "aws-native:rds:DBClusterEndpoint" missing
Type "aws-native:wafv2:WebACLCustomRequestHandling" missing
Type "aws-native:wafv2:WebACLRuleGroupReferenceStatement" missing
Type "aws-native:appflow:FlowDestinationConnectorProperties" input "sapoData" type changed from "#/types/aws-native:appflow:FlowSAPODataDestinationProperties" to "#/types/aws-native:appflow:FlowSapoDataDestinationProperties"
Type "aws-native:pipes:SelfManagedKafkaAccessConfigurationCredentials2Properties" missing
Type "aws-native:lex:BotSSMLMessage" missing
Type "aws-native:wafv2:WebACLRateLimitLabelNamespace" missing
Type "aws-native:m2:Definition0Properties" missing
Type "aws-native:ec2:IPAMPoolTag" missing
Type "aws-native:lakeformation:PrincipalPermissionsResource" input "lfTag" type changed from "#/types/aws-native:lakeformation:PrincipalPermissionsLFTagKeyResource" to "#/types/aws-native:lakeformation:PrincipalPermissionsLfTagKeyResource"
Type "aws-native:lakeformation:PrincipalPermissionsResource" input "lfTagPolicy" type changed from "#/types/aws-native:lakeformation:PrincipalPermissionsLFTagPolicyResource" to "#/types/aws-native:lakeformation:PrincipalPermissionsLfTagPolicyResource"
Type "aws-native:datasync:LocationHDFSTag" missing
Type "aws-native:datasync:LocationEFSInTransitEncryption" missing
Type "aws-native:wafv2:WebACLHeaderMatchPattern" missing
Type "aws-native:quicksight:DashboardKPIPrimaryValueConditionalFormatting" missing
Type "aws-native:quicksight:DashboardTableFieldURLConfiguration" missing
Type "aws-native:wafv2:WebACLTag" missing
Type "aws-native:rds:DBProxyTargetGroupTargetGroupName" missing
Type "aws-native:wafv2:WebACLResponseInspectionStatusCode" missing
Type "aws-native:wafv2:WebACLRegexPatternSetReferenceStatement" missing
Type "aws-native:rds:DBProxyTagFormat" missing
Type "aws-native:wafv2:WebACLImmunityTimeProperty" missing
Type "aws-native:ec2:IPAMScopeTag" missing
Type "aws-native:iotwireless:FuotaTaskLoRaWAN" missing
Type "aws-native:wafv2:WebACLRateLimitIP" missing
Type "aws-native:wafv2:WebACLRateBasedStatementAggregateKeyType" missing
Type "aws-native:amplifyuibuilder:FieldPosition0Properties" missing
Type "aws-native:datasync:LocationFSxOpenZFSMountOptionsVersion" missing
Type "aws-native:wafv2:WebACLAWSManagedRulesBotControlRuleSet" missing
Type "aws-native:wafv2:WebACLAndStatement" missing
Type "aws-native:pipes:MSKAccessCredentials1Properties" missing
Type "aws-native:kinesisfirehose:DeliveryStreamSplunkDestinationConfiguration" input "hecEndpointType" type changed from "#/types/aws-native:kinesisfirehose:DeliveryStreamSplunkDestinationConfigurationHECEndpointType" to "#/types/aws-native:kinesisfirehose:DeliveryStreamSplunkDestinationConfigurationHecEndpointType"
Type "aws-native:wafv2:WebACLDefaultAction" missing
Type "aws-native:wafv2:WebACLExcludedRule" missing
Type "aws-native:cloudfront:ResponseHeadersPolicyXSSProtection" missing
Type "aws-native:lookoutmetrics:AlertSNSConfiguration" missing
Type "aws-native:rds:DBInstanceRole" missing
Type "aws-native:ecs:TaskDefinitionAuthorizationConfigIAM" missing
Type "aws-native:quicksight:TemplateURLTargetConfiguration" missing
Type "aws-native:ecs:TaskDefinitionEFSVolumeConfiguration" missing
Type "aws-native:opensearchservice:DomainAdvancedSecurityOptionsInput" input "samlOptions" type changed from "#/types/aws-native:opensearchservice:DomainSAMLOptions" to "#/types/aws-native:opensearchservice:DomainSamlOptions"
Type "aws-native:wafv2:IPSetScope" missing
Type "aws-native:healthlake:FHIRDatastoreIdentityProviderConfiguration" missing
Type "aws-native:pipes:SelfManagedKafkaAccessConfigurationCredentials3Properties" missing
Type "aws-native:datasync:LocationNFSOnPremConfig" missing
Type "aws-native:msk:ClusterEBSStorageInfo" missing
Type "aws-native:appflow:ConnectorProfileSAPODataConnectorProfileCredentialsOAuthCredentialsProperties" missing
Type "aws-native:s3:BucketServerSideEncryptionByDefault" input "sseAlgorithm" type changed from "#/types/aws-native:s3:BucketServerSideEncryptionByDefaultSSEAlgorithm" to "#/types/aws-native:s3:BucketServerSideEncryptionByDefaultSseAlgorithm"
Type "aws-native:quicksight:AnalysisURLTargetConfiguration" missing
Type "aws-native:quicksight:TemplateKPIConditionalFormatting" missing
Type "aws-native:appflow:ConnectorProfileSAPODataConnectorProfileCredentials" missing
Type "aws-native:dynamodb:GlobalTableSSESpecification" missing
Type "aws-native:ec2:EC2FleetTagSpecificationResourceType" missing
Type "aws-native:appflow:FlowSourceConnectorProperties" input "sapoData" type changed from "#/types/aws-native:appflow:FlowSAPODataSourceProperties" to "#/types/aws-native:appflow:FlowSapoDataSourceProperties"
Type "aws-native:wafv2:WebACLMapMatchScope" missing
Type "aws-native:quicksight:TemplateKPIOptions" missing
Type "aws-native:rds:DBInstanceProcessorFeature" missing
Type "aws-native:iotwireless:TaskDefinitionLoRaWANGatewayVersion" missing
Type "aws-native:ec2:EC2FleetInstanceRequirementsRequestAcceleratorManufacturersItem" missing
Type "aws-native:ec2:EC2FleetInstanceRequirementsRequest" missing
Type "aws-native:wafv2:WebACLFieldIdentifier" missing
Type "aws-native:quicksight:TemplateKPIConfiguration" missing
Type "aws-native:quicksight:AnalysisKPIVisual" missing
Type "aws-native:iotwireless:TaskDefinitionLoRaWANUpdateGatewayTaskCreate" missing
Type "aws-native:iotwireless:WirelessGatewayLoRaWANGateway" missing
Type "aws-native:ec2:EC2FleetFleetLaunchTemplateOverridesRequest" missing
Type "aws-native:ec2:VPNGatewayTag" missing
Type "aws-native:rds:DBProxyEngineFamily" missing
Type "aws-native:iot:CACertificateAutoRegistrationStatus" missing
Type "aws-native:ecs:TaskDefinitionVolume" input "efsVolumeConfiguration" type changed from "#/types/aws-native:ecs:TaskDefinitionEFSVolumeConfiguration" to "#/types/aws-native:ecs:TaskDefinitionEfsVolumeConfiguration"
Type "aws-native:wafv2:WebACLRequestInspectionACFPPayloadType" missing
Type "aws-native:wafv2:WebACLHeaders" missing
Type "aws-native:quicksight:AnalysisKPISortConfiguration" missing
Type "aws-native:wafv2:WebACLAllowAction" missing
Type "aws-native:wafv2:WebACLNotStatement" missing
Type "aws-native:devicefarm:VPCEConfigurationTag" missing
Type "aws-native:datasync:LocationHDFSQopConfiguration" missing
Type "aws-native:wafv2:RuleGroupRateLimitIP" missing
Type "aws-native:rds:DBClusterServerlessV2ScalingConfiguration" missing
Type "aws-native:datasync:LocationFSxONTAPNfsMountOptionsVersion" missing
Type "aws-native:wafv2:WebACLRegexMatchStatement" missing
Type "aws-native:quicksight:AnalysisKPIConditionalFormatting" missing
Type "aws-native:wafv2:WebACLChallengeConfig" missing
Type "aws-native:wafv2:WebACLJsonMatchPattern" missing
Type "aws-native:ec2:LaunchTemplateTotalLocalStorageGB" missing
Type "aws-native:ec2:VPCEndpointVpcEndpointType" missing
Type "aws-native:rds:DBProxyTargetGroupConnectionPoolConfigurationInfoFormat" missing
Type "aws-native:wafv2:WebACLCaptchaConfig" missing
Type "aws-native:wafv2:RuleGroupRateLimitForwardedIP" missing
Type "aws-native:lambda:EventSourceMappingDocumentDBEventSourceConfigFullDocument" missing
Type "aws-native:quicksight:AnalysisTableFieldOption" input "urlStyling" type changed from "#/types/aws-native:quicksight:AnalysisTableFieldURLConfiguration" to "#/types/aws-native:quicksight:AnalysisTableFieldUrlConfiguration"
Type "aws-native:rds:DBProxyAuthFormatClientPasswordAuthType" missing
Type "aws-native:ec2:LaunchTemplateInstanceRequirements" input "totalLocalStorageGb" type changed from "#/types/aws-native:ec2:LaunchTemplateTotalLocalStorageGB" to "#/types/aws-native:ec2:LaunchTemplateTotalLocalStorageGb"
Type "aws-native:lakeformation:PrincipalPermissionsLFTagPolicyResource" missing
Type "aws-native:iotevents:DetectorModelAction" input "dynamoDb" type changed from "#/types/aws-native:iotevents:DetectorModelDynamoDB" to "#/types/aws-native:iotevents:DetectorModelDynamoDb"
Type "aws-native:appstream:Tag0Properties" missing
Type "aws-native:memorydb:ACLTag" missing
Type "aws-native:iot:MitigationActionUpdateCACertificateParams" missing
Type "aws-native:pipes:SelfManagedKafkaAccessConfigurationCredentials0Properties" missing
Type "aws-native:datasync:LocationSMBTag" missing
Type "aws-native:datasync:LocationFSxOpenZFSNFS" missing
Type "aws-native:ec2:IPAMScopeIpamScopeType" missing
Type "aws-native:wafv2:WebACLTextTransformation" missing
Type "aws-native:ec2:EC2FleetMemoryMiBRequest" missing
Type "aws-native:rds:DBProxyEndpointTargetRole" missing
Type "aws-native:wafv2:WebACLResponseInspection" missing
Type "aws-native:quicksight:DashboardTableFieldLinkConfiguration" input "target" type changed from "#/types/aws-native:quicksight:DashboardURLTargetConfiguration" to "#/types/aws-native:quicksight:DashboardUrlTargetConfiguration"
Type "aws-native:wafv2:WebACLOrStatement" missing
Type "aws-native:route53recoveryreadiness:ResourceSetTargetResource" input "nlbResource" type changed from "#/types/aws-native:route53recoveryreadiness:ResourceSetNLBResource" to "#/types/aws-native:route53recoveryreadiness:ResourceSetNlbResource"
Type "aws-native:applicationinsights:ApplicationConfigurationDetails" input "haClusterPrometheusExporter" type changed from "#/types/aws-native:applicationinsights:ApplicationHAClusterPrometheusExporter" to "#/types/aws-native:applicationinsights:ApplicationHaClusterPrometheusExporter"
Type "aws-native:applicationinsights:ApplicationConfigurationDetails" input "hanaPrometheusExporter" type changed from "#/types/aws-native:applicationinsights:ApplicationHANAPrometheusExporter" to "#/types/aws-native:applicationinsights:ApplicationHanaPrometheusExporter"
Type "aws-native:applicationinsights:ApplicationConfigurationDetails" input "jmxPrometheusExporter" type changed from "#/types/aws-native:applicationinsights:ApplicationJMXPrometheusExporter" to "#/types/aws-native:applicationinsights:ApplicationJmxPrometheusExporter"
Type "aws-native:datasync:LocationFSxOpenZFSMountOptions" missing
Type "aws-native:networkfirewall:RuleGroupMatchAttributes" input "tcpFlags" items type changed from "#/types/aws-native:networkfirewall:RuleGroupTCPFlagField" to "#/types/aws-native:networkfirewall:RuleGroupTcpFlagField"
Type "aws-native:ec2:EC2FleetSpotOptionsRequestAllocationStrategy" missing
Type "aws-native:iotevents:DetectorModelDynamoDB" missing
Type "aws-native:pipes:PipeSourceRabbitMQBrokerParameters" missing
Type "aws-native:quicksight:TemplateKPISortConfiguration" missing
Type "aws-native:wafv2:WebACLByteMatchStatement" missing
Type "aws-native:wafv2:RuleGroupRateBasedStatementCustomKey" input "forwardedIp" type changed from "#/types/aws-native:wafv2:RuleGroupRateLimitForwardedIP" to "#/types/aws-native:wafv2:RuleGroupRateLimitForwardedIp"
Type "aws-native:wafv2:RuleGroupRateBasedStatementCustomKey" input "httpMethod" type changed from "#/types/aws-native:wafv2:RuleGroupRateLimitHTTPMethod" to "#/types/aws-native:wafv2:RuleGroupRateLimitHttpMethod"
Type "aws-native:wafv2:RuleGroupRateBasedStatementCustomKey" input "ip" type changed from "#/types/aws-native:wafv2:RuleGroupRateLimitIP" to "#/types/aws-native:wafv2:RuleGroupRateLimitIp"
Type "aws-native:wafv2:WebACLSizeConstraintStatement" missing
Type "aws-native:ec2:EC2FleetInstanceRequirementsRequestInstanceGenerationsItem" missing
Type "aws-native:healthlake:FHIRDatastoreKmsEncryptionConfig" missing
Type "aws-native:quicksight:VPCConnectionTag" missing
Type "aws-native:ec2:EC2FleetTotalLocalStorageGBRequest" missing
Type "aws-native:wafv2:WebACLLabelMatchStatement" missing
Type "aws-native:wafv2:WebACLIPSetForwardedIPConfigurationFallbackBehavior" missing
Type "aws-native:neptune:DBClusterServerlessScalingConfiguration" missing
Type "aws-native:pipes:MQBrokerAccessCredentialsProperties" missing
Type "aws-native:lookoutmetrics:AnomalyDetectorRDSSourceConfig" missing
Type "aws-native:quicksight:AnalysisTableFieldLinkConfiguration" input "target" type changed from "#/types/aws-native:quicksight:AnalysisURLTargetConfiguration" to "#/types/aws-native:quicksight:AnalysisUrlTargetConfiguration"
Type "aws-native:rds:DBProxyAuthFormat" missing
Type "aws-native:wafv2:WebACLLabelMatchScope" missing
Type "aws-native:wafv2:WebACLRequestInspectionACFP" missing
Type "aws-native:wafv2:WebACLRuleAction" missing
Type "aws-native:configuration:TemplateSSMDocumentDetailsProperties" missing
Type "aws-native:wafv2:WebACLIPSetForwardedIPConfigurationPosition" missing
Type "aws-native:quicksight:AnalysisKPIOptions" missing
Type "aws-native:wafv2:RuleGroupIPSetForwardedIPConfigurationFallbackBehavior" missing
Type "aws-native:iotwireless:TaskDefinitionLoRaWANUpdateGatewayTaskEntry" missing
Type "aws-native:datasync:LocationSMBMountOptionsVersion" missing
Type "aws-native:quicksight:DashboardKPIFieldWells" missing
Type "aws-native:pipes:PipeSourceActiveMQBrokerParameters" missing
Type "aws-native:ec2:EC2FleetNetworkInterfaceCountRequest" missing
Type "aws-native:wafv2:WebACLCustomResponse" missing
Type "aws-native:ec2:EC2FleetTargetCapacitySpecificationRequestDefaultTargetCapacityType" missing
Type "aws-native:ec2:EC2FleetSpotOptionsRequest" missing
Type "aws-native:quicksight:AnalysisKPIPrimaryValueConditionalFormatting" missing
Type "aws-native:wafv2:WebACLManagedRuleGroupStatement" missing
Type "aws-native:s3outposts:FilterAndOperator0Properties" missing
Type "aws-native:ec2:EC2FleetInstanceRequirementsRequestBurstablePerformance" missing
Type "aws-native:wafv2:WebACLSqliMatchStatement" missing
Type "aws-native:wafv2:WebACLFieldToMatchSingleHeaderProperties" missing
Type "aws-native:ec2:IPAMTag" missing
Type "aws-native:datasync:LocationFSxOpenZFSProtocol" missing
Type "aws-native:wafv2:RuleGroupCustomRequestHandling" input "insertHeaders" items type changed from "#/types/aws-native:wafv2:RuleGroupCustomHTTPHeader" to "#/types/aws-native:wafv2:RuleGroupCustomHttpHeader"
Type "aws-native:wafv2:WebACLOverrideAction" missing
Type "aws-native:amplifyuibuilder:FormCTA" missing
Type "aws-native:wafv2:WebACLIPSetForwardedIPConfiguration" missing
Type "aws-native:wafv2:WebACLCustomHTTPHeader" missing
Type "aws-native:datasync:LocationFSxOpenZFSTag" missing
Type "aws-native:datasync:LocationEFSEc2Config" missing
Type "aws-native:pipes:PipeSourceDynamoDBStreamParameters" missing
Type "aws-native:wafv2:WebACLAWSManagedRulesACFPRuleSet" missing
Type "aws-native:iotevents:AlarmModelDynamoDB" missing
Type "aws-native:dynamodb:GlobalTableReplicaSSESpecification" missing
Type "aws-native:quicksight:AnalysisTableFieldURLConfiguration" missing
Type "aws-native:datasync:LocationFSxONTAPSmbMountOptions" missing
Type "aws-native:quicksight:TemplateKPIVisual" missing
Type "aws-native:wafv2:WebACLAssociationConfig" missing
Type "aws-native:datasync:LocationHDFSQopConfigurationRpcProtection" missing
Type "aws-native:quicksight:TemplateKPIFieldWells" missing
Type "aws-native:lambda:EventSourceMappingDocumentDBEventSourceConfig" missing
Type "aws-native:quicksight:DashboardUIState" missing
Type "aws-native:kinesisfirehose:DeliveryStreamKMSEncryptionConfig" missing
Type "aws-native:wafv2:WebACLAWSManagedRulesATPRuleSet" missing
Type "aws-native:lex:BotMessage" input "ssmlMessage" type changed from "#/types/aws-native:lex:BotSSMLMessage" to "#/types/aws-native:lex:BotSsmlMessage"
Type "aws-native:opensearchservice:DomainEBSOptions" missing
Type "aws-native:wafv2:WebACLAWSManagedRulesBotControlRuleSetInspectionLevel" missing
Type "aws-native:amplifyuibuilder:FieldPosition2Properties" missing
Type "aws-native:quicksight:DashboardCustomActionURLOperation" missing
Type "aws-native:quicksight:DashboardVisualCustomActionOperation" input "urlOperation" type changed from "#/types/aws-native:quicksight:DashboardCustomActionURLOperation" to "#/types/aws-native:quicksight:DashboardCustomActionUrlOperation"
Type "aws-native:iot:TopicRuleDynamoDBAction" missing
Type "aws-native:rolesanywhere:SourceData1Properties" missing
Type "aws-native:ec2:EC2FleetOnDemandOptionsRequest" missing
Type "aws-native:wafv2:WebACLBodyParsingFallbackBehavior" missing
Type "aws-native:wafv2:WebACLRuleActionOverride" missing
Type "aws-native:quicksight:DashboardKPIConditionalFormatting" missing
Type "aws-native:iam:VirtualMFADeviceTag" missing
Type "aws-native:lakeformation:TagAssociationLFTagPair" missing
Type "aws-native:ec2:EC2FleetNetworkBandwidthGbpsRequest" missing
Type "aws-native:rds:DBInstanceTag" missing
Type "aws-native:ec2:EC2FleetCapacityReservationOptionsRequestUsageStrategy" missing
Type "aws-native:networkfirewall:RuleGroupTCPFlag" missing
Type "aws-native:quicksight:AnalysisKPIConditionalFormattingOption" missing
Type "aws-native:auditmanager:AssessmentScope" input "awsAccounts" items type changed from "#/types/aws-native:auditmanager:AssessmentAWSAccount" to "#/types/aws-native:auditmanager:AssessmentAwsAccount"
Type "aws-native:auditmanager:AssessmentScope" input "awsServices" items type changed from "#/types/aws-native:auditmanager:AssessmentAWSService" to "#/types/aws-native:auditmanager:AssessmentAwsService"
Type "aws-native:route53resolver:ResolverDNSSECConfigValidationStatus" missing
Type "aws-native:quicksight:TemplateTableFieldOption" input "urlStyling" type changed from "#/types/aws-native:quicksight:TemplateTableFieldURLConfiguration" to "#/types/aws-native:quicksight:TemplateTableFieldUrlConfiguration"
Type "aws-native:appflow:FlowSAPODataSourceProperties" missing
Type "aws-native:rds:DBInstanceProcessorFeatureName" missing
Type "aws-native:wafv2:WebACLJsonMatchScope" missing
Type "aws-native:forecast:AttributesItemProperties" missing
Type "aws-native:appflow:FlowSAPODataDestinationProperties" missing
Type "aws-native:appflow:FlowConnectorOperator" input "sapoData" type changed from "#/types/aws-native:appflow:FlowSAPODataConnectorOperator" to "#/types/aws-native:appflow:FlowSapoDataConnectorOperator"
Type "aws-native:ec2:EC2FleetInstanceRequirementsRequestLocalStorageTypesItem" missing
Type "aws-native:appflow:FlowSAPODataConnectorOperator" missing
Type "aws-native:wafv2:WebACLCookieMatchPattern" missing
Type "aws-native:rds:DBSubnetGroupTag" missing
Type "aws-native:wafv2:WebACLForwardedIPConfigurationFallbackBehavior" missing
Type "aws-native:quicksight:DashboardKPIVisual" missing
Type "aws-native:ec2:EC2FleetExcessCapacityTerminationPolicy" missing
Type "aws-native:rds:DBClusterReadEndpoint" missing
Type "aws-native:ec2:EC2FleetFleetLaunchTemplateSpecificationRequest" missing
Type "aws-native:ec2:IPAMResourceDiscoveryAssociationTag" missing
Type "aws-native:wafv2:RuleGroupRateLimitHTTPMethod" missing
Type "aws-native:quicksight:VPCConnectionResourceStatus" missing
Type "aws-native:wafv2:WebACLLabel" missing
Type "aws-native:iot:CACertificateTag" missing
Type "aws-native:iam:OIDCProviderTag" missing
Type "aws-native:route53recoveryreadiness:ResourceSetResource" input "dnsTargetResource" type changed from "#/types/aws-native:route53recoveryreadiness:ResourceSetDNSTargetResource" to "#/types/aws-native:route53recoveryreadiness:ResourceSetDnsTargetResource"
Type "aws-native:ec2:IPAMPoolProvisionedCidr" missing
Type "aws-native:appflow:ConnectorProfileSAPODataConnectorProfileProperties" missing
Type "aws-native:quicksight:TemplateTableFieldURLConfiguration" missing
Type "aws-native:ec2:EC2FleetInstanceRequirementsRequestBareMetal" missing
Type "aws-native:rds:DBClusterScalingConfiguration" missing
Type "aws-native:wafv2:RuleGroupStatement" input "ipSetReferenceStatement" type changed from "#/types/aws-native:wafv2:RuleGroupIPSetReferenceStatement" to "#/types/aws-native:wafv2:RuleGroupIpSetReferenceStatement"
Type "aws-native:iotevents:AlarmModelAlarmAction" input "dynamoDb" type changed from "#/types/aws-native:iotevents:AlarmModelDynamoDB" to "#/types/aws-native:iotevents:AlarmModelDynamoDb"
Type "aws-native:wafv2:RuleGroupCustomResponse" input "responseHeaders" items type changed from "#/types/aws-native:wafv2:RuleGroupCustomHTTPHeader" to "#/types/aws-native:wafv2:RuleGroupCustomHttpHeader"
Type "aws-native:wafv2:WebACLJsonBody" missing
Type "aws-native:healthlake:FHIRDatastoreKmsEncryptionConfigCmkType" missing
Type "aws-native:quicksight:TemplateTableFieldLinkConfiguration" input "target" type changed from "#/types/aws-native:quicksight:TemplateURLTargetConfiguration" to "#/types/aws-native:quicksight:TemplateUrlTargetConfiguration"
Type "aws-native:wafv2:WebACLSensitivityLevel" missing
Type "aws-native:ec2:EIPTag" missing
Type "aws-native:quicksight:TemplateVisual" input "kpiVisual" type changed from "#/types/aws-native:quicksight:TemplateKPIVisual" to "#/types/aws-native:quicksight:TemplateKpiVisual"
Type "aws-native:ecs:TaskDefinitionEFSVolumeConfigurationTransitEncryption" missing
Type "aws-native:quicksight:TemplateVisualCustomActionOperation" input "urlOperation" type changed from "#/types/aws-native:quicksight:TemplateCustomActionURLOperation" to "#/types/aws-native:quicksight:TemplateCustomActionUrlOperation"
Type "aws-native:wafv2:RuleGroupForwardedIPConfigurationFallbackBehavior" missing
Type "aws-native:wafv2:RuleGroupForwardedIPConfiguration" missing
Type "aws-native:quicksight:ThemeConfiguration" input "uiColorPalette" type changed from "#/types/aws-native:quicksight:ThemeUIColorPalette" to "#/types/aws-native:quicksight:ThemeUiColorPalette"
Type "aws-native:apigatewayv2:AuthorizerJWTConfiguration" missing
Type "aws-native:route53recoveryreadiness:ResourceSetDNSTargetResource" missing
Type "aws-native:ec2:VPCTag" missing
Type "aws-native:wafv2:WebACLRateLimitQueryArgument" missing
Type "aws-native:forecast:SchemaProperties" input "attributes" items type changed from "#/types/aws-native:forecast:AttributesItemProperties" to "#/types/aws-native:forecast:DatasetAttributesItemProperties"
Type "aws-native:wafv2:WebACLPositionalConstraint" missing
Type "aws-native:iot:MitigationActionActionParams" input "updateCaCertificateParams" type changed from "#/types/aws-native:iot:MitigationActionUpdateCACertificateParams" to "#/types/aws-native:iot:MitigationActionUpdateCaCertificateParams"
Type "aws-native:ec2:EC2FleetInstanceRequirementsRequestCpuManufacturersItem" missing
Type "aws-native:applicationinsights:ApplicationJMXPrometheusExporter" missing
Type "aws-native:wafv2:WebACLChallengeAction" missing
Type "aws-native:quicksight:AnalysisKPIProgressBarConditionalFormatting" missing
Type "aws-native:opensearchservice:DomainSAMLOptions" missing
Type "aws-native:wafv2:WebACLRateLimitHTTPMethod" missing
Type "aws-native:ec2:LocalGatewayRouteTableVPCAssociationTag" missing
Type "aws-native:networkmanager:DeviceAWSLocation" missing
Type "aws-native:datasync:LocationFSxONTAPTag" missing
Type "aws-native:wafv2:IPSetTag" missing
Type "aws-native:quicksight:TemplateKPIProgressBarConditionalFormatting" missing
Type "aws-native:omics:StoreOptions0Properties" missing
Type "aws-native:rds:CustomDBEngineVersionStatus" missing
Type "aws-native:wafv2:WebACLIPSetReferenceStatement" missing
Type "aws-native:wafv2:WebACLForwardedIPConfiguration" missing
Type "aws-native:healthlake:FHIRDatastoreDatastoreTypeVersion" missing
Type "aws-native:ec2:SpotFleetInstanceRequirementsRequest" input "totalLocalStorageGb" type changed from "#/types/aws-native:ec2:SpotFleetTotalLocalStorageGBRequest" to "#/types/aws-native:ec2:SpotFleetTotalLocalStorageGbRequest"
Type "aws-native:kinesisfirehose:DeliveryStreamEncryptionConfiguration" input "kmsEncryptionConfig" type changed from "#/types/aws-native:kinesisfirehose:DeliveryStreamKMSEncryptionConfig" to "#/types/aws-native:kinesisfirehose:DeliveryStreamKmsEncryptionConfig"
Type "aws-native:ec2:EC2FleetMemoryGiBPerVCpuRequest" missing
Type "aws-native:wafv2:WebACLGeoMatchStatement" missing
Type "aws-native:ec2:EC2FleetType" missing
Type "aws-native:wafv2:WebACLRequestInspection" missing
Type "aws-native:rds:DBProxyAuthFormatIAMAuth" missing
Type "aws-native:msk:ClusterStorageInfo" input "ebsStorageInfo" type changed from "#/types/aws-native:msk:ClusterEBSStorageInfo" to "#/types/aws-native:msk:ClusterEbsStorageInfo"
Type "aws-native:wafv2:RuleGroupCustomHTTPHeader" missing
Type "aws-native:quicksight:VPCConnectionAvailabilityStatus" missing
Type "aws-native:wafv2:WebACLRateLimitHeader" missing
Type "aws-native:wafv2:RuleGroupGeoMatchStatement" input "forwardedIpConfig" type changed from "#/types/aws-native:wafv2:RuleGroupForwardedIPConfiguration" to "#/types/aws-native:wafv2:RuleGroupForwardedIpConfiguration"
Type "aws-native:quicksight:VPCConnectionNetworkInterfaceStatus" missing
Type "aws-native:ec2:EC2FleetVCpuCountRangeRequest" missing
Type "aws-native:iotwireless:DeviceProfileLoRaWANDeviceProfile" missing
Type "aws-native:kinesisanalyticsv2:ApplicationJSONMappingParameters" missing
Type "aws-native:datasync:LocationEFSTag" missing
Type "aws-native:ec2:IPAMResourceDiscoveryIpamOperatingRegion" missing
Type "aws-native:ec2:EC2FleetSpotOptionsRequestInstanceInterruptionBehavior" missing
Type "aws-native:wafv2:WebACLSizeConstraintStatementComparisonOperator" missing
Type "aws-native:quicksight:DashboardKPIProgressBarConditionalFormatting" missing
Type "aws-native:lookoutmetrics:AnomalyDetectorMetricSource" input "rdsSourceConfig" type changed from "#/types/aws-native:lookoutmetrics:AnomalyDetectorRDSSourceConfig" to "#/types/aws-native:lookoutmetrics:AnomalyDetectorRdsSourceConfig"
Type "aws-native:synthetics:CanaryVPCConfig" missing
Type "aws-native:dynamodb:TableSSESpecification" missing
Type "aws-native:rolesanywhere:SourceData0Properties" missing
Type "aws-native:wafv2:WebACLBody" missing
Type "aws-native:rds:DBProxyEndpointTagFormat" missing
Type "aws-native:s3:BucketServerSideEncryptionByDefaultSSEAlgorithm" missing
Type "aws-native:route53:HostedZoneVPC" missing
Type "aws-native:wafv2:WebACLRequestBody" missing
Type "aws-native:quicksight:TemplateKPIPrimaryValueConditionalFormatting" missing
Type "aws-native:ec2:DHCPOptionsTag" missing
Type "aws-native:wafv2:WebACLScope" missing
Type "aws-native:datasync:LocationNFSMountOptions" missing
Type "aws-native:wafv2:WebACLRateLimitQueryString" missing
Type "aws-native:lookoutmetrics:AlertAction" input "snsConfiguration" type changed from "#/types/aws-native:lookoutmetrics:AlertSNSConfiguration" to "#/types/aws-native:lookoutmetrics:AlertSnsConfiguration"
Type "aws-native:datasync:LocationNFSTag" missing
Type "aws-native:fms:PolicyIEMap" missing
Type "aws-native:dynamodb:GlobalTableReplicaSpecification" input "sseSpecification" type changed from "#/types/aws-native:dynamodb:GlobalTableReplicaSSESpecification" to "#/types/aws-native:dynamodb:GlobalTableReplicaSseSpecification"
Type "aws-native:datasync:LocationNFSMountOptionsVersion" missing
Type "aws-native:quicksight:DashboardPublishOptions" input "exportToCsvOption" type changed from "#/types/aws-native:quicksight:DashboardExportToCSVOption" to "#/types/aws-native:quicksight:DashboardExportToCsvOption"
Type "aws-native:quicksight:DashboardVisual" input "kpiVisual" type changed from "#/types/aws-native:quicksight:DashboardKPIVisual" to "#/types/aws-native:quicksight:DashboardKpiVisual"
Type "aws-native:auditmanager:AssessmentAWSService" missing
Type "aws-native:lakeformation:PrincipalPermissionsLFTagKeyResource" missing
Type "aws-native:quicksight:AnalysisKPIConfiguration" missing
Type "aws-native:quicksight:TemplateCustomActionURLOperation" missing
Type "aws-native:ec2:IPAMPoolIpamScopeType" missing
Type "aws-native:wafv2:WebACLStatement" missing
Type "aws-native:ec2:EC2FleetTargetCapacitySpecificationRequest" missing
Type "aws-native:iam:SAMLProviderTag" missing
Type "aws-native:pipes:PipeDynamoDBStreamStartPosition" missing
Type "aws-native:personalize:SolutionConfigAutoMLConfigProperties" missing
Type "aws-native:wafv2:WebACLOversizeHandling" missing
Type "aws-native:appflow:ConnectorProfileCredentials" input "sapoData" type changed from "#/types/aws-native:appflow:ConnectorProfileSAPODataConnectorProfileCredentials" to "#/types/aws-native:appflow:ConnectorProfileSapoDataConnectorProfileCredentials"
Type "aws-native:ec2:IPAMPoolState" missing
Type "aws-native:wafv2:WebACLRateBasedStatement" missing
Type "aws-native:applicationinsights:ApplicationHANAPrometheusExporter" missing
Type "aws-native:networkfirewall:RuleGroupTCPFlagField" missing
Type "aws-native:wafv2:WebACLFieldToMatch" missing
Type "aws-native:ec2:EC2FleetCapacityReservationOptionsRequest" missing
Type "aws-native:ec2:EC2FleetInstanceRequirementsRequestLocalStorage" missing
Type "aws-native:ec2:EC2FleetMaintenanceStrategies" missing
Type "aws-native:wafv2:WebACLRateLimitCookie" missing

New resources:
datasync.LocationEfs
datasync.LocationFSxOntap
datasync.LocationFSxOpenZfs
datasync.LocationHdfs
datasync.LocationNfs
datasync.LocationSmb
devicefarm.VpceConfiguration
directoryservice.SimpleAd
ec2.DhcpOptions
ec2.Ec2Fleet
ec2.Eip
ec2.EipAssociation
ec2.Ipam
ec2.IpamAllocation
ec2.IpamPool
ec2.IpamPoolCidr
ec2.IpamResourceDiscovery
ec2.IpamResourceDiscoveryAssociation
ec2.IpamScope
ec2.LocalGatewayRouteTableVpcAssociation
ec2.Vpc
ec2.VpcEndpoint
ec2.VpcEndpointService
ec2.VpcEndpointServicePermissions
ec2.VpcPeeringConnection
ec2.VpcdhcpOptionsAssociation
ec2.VpnConnection
ec2.VpnConnectionRoute
ec2.VpnGateway
healthlake.FhirDatastore
iam.OidcProvider
iam.SamlProvider
iam.VirtualMfaDevice
iot.CaCertificate
memorydb.Acl
neptune.DbCluster
quicksight.VpcConnection
rds.CustomDbEngineVersion
rds.DbCluster
rds.DbClusterParameterGroup
rds.DbInstance
rds.DbParameterGroup
rds.DbProxy
rds.DbProxyEndpoint
rds.DbProxyTargetGroup
rds.DbSubnetGroup
rolesanywhere.Crl
route53.Dnssec
route53resolver.ResolverDnssecConfig
shield.DrtAccess
wafv2.IpSet
wafv2.WebAcl
wafv2.WebAclAssociation
New functions:
datasync.getLocationEfs
datasync.getLocationFSxOntap
datasync.getLocationFSxOpenZfs
datasync.getLocationHdfs
datasync.getLocationNfs
datasync.getLocationSmb
devicefarm.getVpceConfiguration
directoryservice.getSimpleAd
ec2.getDhcpOptions
ec2.getEc2Fleet
ec2.getEip
ec2.getEipAssociation
ec2.getIpam
ec2.getIpamAllocation
ec2.getIpamPool
ec2.getIpamPoolCidr
ec2.getIpamResourceDiscovery
ec2.getIpamResourceDiscoveryAssociation
ec2.getIpamScope
ec2.getLocalGatewayRouteTableVpcAssociation
ec2.getVpc
ec2.getVpcEndpoint
ec2.getVpcEndpointService
ec2.getVpcEndpointServicePermissions
ec2.getVpcPeeringConnection
ec2.getVpnConnection
ec2.getVpnGateway
healthlake.getFhirDatastore
iam.getOidcProvider
iam.getSamlProvider
iam.getVirtualMfaDevice
iot.getCaCertificate
memorydb.getAcl
neptune.getDbCluster
quicksight.getVpcConnection
rds.getCustomDbEngineVersion
rds.getDbCluster
rds.getDbClusterParameterGroup
rds.getDbInstance
rds.getDbParameterGroup
rds.getDbProxy
rds.getDbProxyEndpoint
rds.getDbProxyTargetGroup
rds.getDbSubnetGroup
rolesanywhere.getCrl
route53resolver.getResolverDnssecConfig
shield.getDrtAccess
wafv2.getIpSet
wafv2.getWebAcl

#### Resources
- "aws-native:amplify:App":
    - `🟡` inputs: "iAMServiceRole" missing
    - `🟡` properties: "iAMServiceRole" missing output "iAMServiceRole"
- "aws-native:amplify:Domain":
    - `🟡` inputs: "autoSubDomainIAMRole" missing
    - `🟡` properties: "autoSubDomainIAMRole" missing output "autoSubDomainIAMRole"
- `🟡` "aws-native:apigateway:ApiKey": properties: "aPIKeyId" missing output "aPIKeyId"
- "aws-native:apigateway:Authorizer":
    - `🟡` inputs: "providerARNs" missing
    - `🟡` properties: "providerARNs" missing output "providerARNs"
- "aws-native:appflow:ConnectorProfile":
    - `🟡` inputs: "kMSArn" missing
    - `🟡` properties: "kMSArn" missing output "kMSArn"
- "aws-native:appflow:Flow":
    - `🟡` inputs: "kMSArn" missing
    - `🟡` properties: "kMSArn" missing output "kMSArn"
- "aws-native:appintegrations:DataIntegration":
    - `🟡` inputs: "sourceURI" missing
    - `🟡` properties: "sourceURI" missing output "sourceURI"
    - `🟢` required inputs: "sourceUri" input has changed to Required
- "aws-native:applicationautoscaling:ScalableTarget":
    - `🟡` inputs: "roleARN" missing
    - `🟡` properties: "roleARN" missing output "roleARN"
- "aws-native:applicationinsights:Application":
    - inputs:
        - `🟡` "cWEMonitorEnabled" missing
        - `🟡` "opsItemSNSTopicArn" missing
    - properties:
        - `🟡` "applicationARN" missing output "applicationARN"
        - `🟡` "cWEMonitorEnabled" missing output "cWEMonitorEnabled"
        - `🟡` "opsItemSNSTopicArn" missing output "opsItemSNSTopicArn"
- "aws-native:autoscaling:LaunchConfiguration":
    - inputs:
        - `🟡` "classicLinkVPCId" missing
        - `🟡` "classicLinkVPCSecurityGroups" missing
    - properties:
        - `🟡` "classicLinkVPCId" missing output "classicLinkVPCId"
        - `🟡` "classicLinkVPCSecurityGroups" missing output "classicLinkVPCSecurityGroups"
- "aws-native:autoscaling:LifecycleHook":
    - inputs:
        - `🟡` "notificationTargetARN" missing
        - `🟡` "roleARN" missing
    - properties:
        - `🟡` "notificationTargetARN" missing output "notificationTargetARN"
        - `🟡` "roleARN" missing output "roleARN"
- "aws-native:cloudformation:StackSet":
    - inputs:
        - `🟡` "administrationRoleARN" missing
        - `🟡` "templateURL" missing
    - properties:
        - `🟡` "administrationRoleARN" missing output "administrationRoleARN"
        - `🟡` "templateURL" missing output "templateURL"
- `🟡` "aws-native:cloudfront:Function": properties: "functionARN" missing output "functionARN"
- "aws-native:cloudtrail:Trail":
    - `🟡` inputs: "kMSKeyId" missing
    - `🟡` properties: "kMSKeyId" missing output "kMSKeyId"
- "aws-native:cloudwatch:CompositeAlarm":
    - `🟡` inputs: "oKActions" missing
    - `🟡` properties: "oKActions" missing output "oKActions"
- "aws-native:configuration:ConformancePack":
    - `🟡` inputs: "templateSSMDocumentDetails" missing
    - `🟡` properties: "templateSSMDocumentDetails" missing output "templateSSMDocumentDetails"
- "aws-native:dynamodb:GlobalTable":
    - `🟡` inputs: "sSESpecification" missing
    - `🟡` properties: "sSESpecification" missing output "sSESpecification"
- "aws-native:dynamodb:Table":
    - `🟡` inputs: "sSESpecification" missing
    - `🟡` properties: "sSESpecification" missing output "sSESpecification"
- "aws-native:ec2:EIPAssociation":
    - `🟡` inputs: "eIP" missing
    - `🟡` properties: "eIP" missing output "eIP"
- `🟡` "aws-native:ec2:VPNGateway": properties: "vPNGatewayId" missing output "vPNGatewayId"
- "aws-native:ec2:Volume":
    - `🟡` inputs: "autoEnableIO" missing
    - `🟡` properties: "autoEnableIO" missing output "autoEnableIO"
- "aws-native:ecs:Service":
    - `🟡` inputs: "enableECSManagedTags" missing
    - `🟡` properties: "enableECSManagedTags" missing output "enableECSManagedTags"
- "aws-native:elasticbeanstalk:Environment":
    - `🟡` inputs: "cNAMEPrefix" missing
    - properties:
        - `🟡` "cNAMEPrefix" missing output "cNAMEPrefix"
        - `🟡` "endpointURL" missing output "endpointURL"
- `🟡` "aws-native:fsx:DataRepositoryAssociation": properties: "resourceARN" missing output "resourceARN"
- "aws-native:gamelift:Fleet":
    - inputs:
        - `🟡` "desiredEC2Instances" missing
        - `🟡` "eC2InboundPermissions" missing
        - `🟡` "eC2InstanceType" missing
        - `🟡` "instanceRoleARN" missing
    - properties:
        - `🟡` "desiredEC2Instances" missing output "desiredEC2Instances"
        - `🟡` "eC2InboundPermissions" missing output "eC2InboundPermissions"
        - `🟡` "eC2InstanceType" missing output "eC2InstanceType"
        - `🟡` "instanceRoleARN" missing output "instanceRoleARN"
- "aws-native:iam:ServiceLinkedRole":
    - `🟡` inputs: "aWSServiceName" missing
    - `🟡` properties: "aWSServiceName" missing output "aWSServiceName"
- "aws-native:iot:CACertificate":
    - `🟡` inputs: "cACertificatePem" missing
    - `🟡` properties: "cACertificatePem" missing output "cACertificatePem"
    - `🟢` required inputs: "caCertificatePem" input has changed to Required
- "aws-native:iot:Certificate":
    - `🟡` inputs: "cACertificatePem" missing
    - `🟡` properties: "cACertificatePem" missing output "cACertificatePem"
- "aws-native:iotwireless:DeviceProfile":
    - `🟡` inputs: "loRaWAN" missing
    - `🟡` properties: "loRaWAN" missing output "loRaWAN"
- "aws-native:iotwireless:FuotaTask":
    - `🟡` inputs: "loRaWAN" missing
    - `🟡` properties: "loRaWAN" missing output "loRaWAN"
    - `🟢` required inputs: "loRaWan" input has changed to Required
- "aws-native:iotwireless:MulticastGroup":
    - `🟡` inputs: "loRaWAN" missing
    - `🟡` properties: "loRaWAN" missing output "loRaWAN"
    - `🟢` required inputs: "loRaWan" input has changed to Required
- "aws-native:iotwireless:ServiceProfile":
    - `🟡` inputs: "loRaWAN" missing
    - `🟡` properties: "loRaWAN" missing output "loRaWAN"
- "aws-native:iotwireless:TaskDefinition":
    - `🟡` inputs: "loRaWANUpdateGatewayTaskEntry" missing
    - `🟡` properties: "loRaWANUpdateGatewayTaskEntry" missing output "loRaWANUpdateGatewayTaskEntry"
- "aws-native:iotwireless:WirelessDevice":
    - `🟡` inputs: "loRaWAN" missing
    - `🟡` properties: "loRaWAN" missing output "loRaWAN"
- "aws-native:iotwireless:WirelessGateway":
    - `🟡` inputs: "loRaWAN" missing
    - `🟡` properties: "loRaWAN" missing output "loRaWAN"
    - `🟢` required inputs: "loRaWan" input has changed to Required
- "aws-native:lakeformation:TagAssociation":
    - `🟡` inputs: "lFTags" missing
    - `🟡` properties: "lFTags" missing output "lFTags"
    - `🟢` required inputs: "lfTags" input has changed to Required
- "aws-native:lambda:EventSourceMapping":
    - `🟡` inputs: "documentDBEventSourceConfig" missing
    - `🟡` properties: "documentDBEventSourceConfig" missing output "documentDBEventSourceConfig"
- "aws-native:lambda:Permission":
    - `🟡` inputs: "principalOrgID" missing
    - `🟡` properties: "principalOrgID" missing output "principalOrgID"
- "aws-native:lex:Bot":
    - `🟡` inputs: "idleSessionTTLInSeconds" missing
    - `🟡` properties: "idleSessionTTLInSeconds" missing output "idleSessionTTLInSeconds"
    - `🟢` required inputs: "idleSessionTtlInSeconds" input has changed to Required
- "aws-native:licensemanager:License":
    - `🟡` inputs: "productSKU" missing
    - `🟡` properties: "productSKU" missing output "productSKU"
- "aws-native:lightsail:LoadBalancer":
    - `🟡` inputs: "sessionStickinessLBCookieDurationSeconds" missing
    - `🟡` properties: "sessionStickinessLBCookieDurationSeconds" missing output "sessionStickinessLBCookieDurationSeconds"
- "aws-native:memorydb:ACL":
    - `🟡` inputs: "aCLName" missing
    - `🟡` properties: "aCLName" missing output "aCLName"
- "aws-native:memorydb:Cluster":
    - inputs:
        - `🟡` "aCLName" missing
        - `🟡` "tLSEnabled" missing
    - properties:
        - `🟡` "aCLName" missing output "aCLName"
        - `🟡` "aRN" missing output "aRN"
        - `🟡` "tLSEnabled" missing output "tLSEnabled"
    - `🟢` required inputs: "aclName" input has changed to Required
- `🟡` "aws-native:memorydb:ParameterGroup": properties: "aRN" missing output "aRN"
- `🟡` "aws-native:memorydb:SubnetGroup": properties: "aRN" missing output "aRN"
- "aws-native:neptune:DBCluster":
    - inputs:
        - `🟡` "dBClusterIdentifier" missing
        - `🟡` "dBClusterParameterGroupName" missing
        - `🟡` "dBInstanceParameterGroupName" missing
        - `🟡` "dBSubnetGroupName" missing
        - `🟡` "sourceDBClusterIdentifier" missing
    - properties:
        - `🟡` "dBClusterIdentifier" missing output "dBClusterIdentifier"
        - `🟡` "dBClusterParameterGroupName" missing output "dBClusterParameterGroupName"
        - `🟡` "dBInstanceParameterGroupName" missing output "dBInstanceParameterGroupName"
        - `🟡` "dBSubnetGroupName" missing output "dBSubnetGroupName"
        - `🟡` "sourceDBClusterIdentifier" missing output "sourceDBClusterIdentifier"
- "aws-native:networkmanager:Device":
    - `🟡` inputs: "aWSLocation" missing
    - `🟡` properties: "aWSLocation" missing output "aWSLocation"
- "aws-native:opensearchservice:Domain":
    - inputs:
        - `🟡` "eBSOptions" missing
        - `🟡` "vPCOptions" missing
    - properties:
        - `🟡` "eBSOptions" missing output "eBSOptions"
        - `🟡` "vPCOptions" missing output "vPCOptions"
- "aws-native:personalize:Solution":
    - inputs:
        - `🟡` "performAutoML" missing
        - `🟡` "performHPO" missing
    - properties:
        - `🟡` "performAutoML" missing output "performAutoML"
        - `🟡` "performHPO" missing output "performHPO"
- "aws-native:quicksight:VPCConnection":
    - `🟡` inputs: "vPCConnectionId" missing
    - properties:
        - `🟡` "vPCConnectionId" missing output "vPCConnectionId"
        - `🟡` "vPCId" missing output "vPCId"
- "aws-native:rds:CustomDBEngineVersion":
    - `🟡` inputs: "kMSKeyId" missing
    - properties:
        - `🟡` "dBEngineVersionArn" missing output "dBEngineVersionArn"
        - `🟡` "kMSKeyId" missing output "kMSKeyId"
- "aws-native:rds:DBCluster":
    - inputs:
        - `🟡` "dBClusterIdentifier" missing
        - `🟡` "dBClusterInstanceClass" missing
        - `🟡` "dBClusterParameterGroupName" missing
        - `🟡` "dBInstanceParameterGroupName" missing
        - `🟡` "dBSubnetGroupName" missing
        - `🟡` "dBSystemId" missing
        - `🟡` "domainIAMRoleName" missing
        - `🟡` "enableIAMDatabaseAuthentication" missing
        - `🟡` "sourceDBClusterIdentifier" missing
    - properties:
        - `🟡` "dBClusterArn" missing output "dBClusterArn"
        - `🟡` "dBClusterIdentifier" missing output "dBClusterIdentifier"
        - `🟡` "dBClusterInstanceClass" missing output "dBClusterInstanceClass"
        - `🟡` "dBClusterParameterGroupName" missing output "dBClusterParameterGroupName"
        - `🟡` "dBClusterResourceId" missing output "dBClusterResourceId"
        - `🟡` "dBInstanceParameterGroupName" missing output "dBInstanceParameterGroupName"
        - `🟡` "dBSubnetGroupName" missing output "dBSubnetGroupName"
        - `🟡` "dBSystemId" missing output "dBSystemId"
        - `🟡` "domainIAMRoleName" missing output "domainIAMRoleName"
        - `🟡` "enableIAMDatabaseAuthentication" missing output "enableIAMDatabaseAuthentication"
        - `🟡` "sourceDBClusterIdentifier" missing output "sourceDBClusterIdentifier"
- "aws-native:rds:DBClusterParameterGroup":
    - `🟡` inputs: "dBClusterParameterGroupName" missing
    - `🟡` properties: "dBClusterParameterGroupName" missing output "dBClusterParameterGroupName"
- "aws-native:rds:DBInstance":
    - inputs:
        - `🟡` "cACertificateIdentifier" missing
        - `🟡` "customIAMInstanceProfile" missing
        - `🟡` "dBClusterIdentifier" missing
        - `🟡` "dBClusterSnapshotIdentifier" missing
        - `🟡` "dBInstanceClass" missing
        - `🟡` "dBInstanceIdentifier" missing
        - `🟡` "dBName" missing
        - `🟡` "dBParameterGroupName" missing
        - `🟡` "dBSecurityGroups" missing
        - `🟡` "dBSnapshotIdentifier" missing
        - `🟡` "dBSubnetGroupName" missing
        - `🟡` "domainIAMRoleName" missing
        - `🟡` "enableIAMDatabaseAuthentication" missing
        - `🟡` "multiAZ" missing
        - `🟡` "performanceInsightsKMSKeyId" missing
        - `🟡` "sourceDBClusterIdentifier" missing
        - `🟡` "sourceDBInstanceAutomatedBackupsArn" missing
        - `🟡` "sourceDBInstanceIdentifier" missing
        - `🟡` "vPCSecurityGroups" missing
    - properties:
        - `🟡` "cACertificateIdentifier" missing output "cACertificateIdentifier"
        - `🟡` "customIAMInstanceProfile" missing output "customIAMInstanceProfile"
        - `🟡` "dBClusterIdentifier" missing output "dBClusterIdentifier"
        - `🟡` "dBClusterSnapshotIdentifier" missing output "dBClusterSnapshotIdentifier"
        - `🟡` "dBInstanceArn" missing output "dBInstanceArn"
        - `🟡` "dBInstanceClass" missing output "dBInstanceClass"
        - `🟡` "dBInstanceIdentifier" missing output "dBInstanceIdentifier"
        - `🟡` "dBName" missing output "dBName"
        - `🟡` "dBParameterGroupName" missing output "dBParameterGroupName"
        - `🟡` "dBSecurityGroups" missing output "dBSecurityGroups"
        - `🟡` "dBSnapshotIdentifier" missing output "dBSnapshotIdentifier"
        - `🟡` "dBSubnetGroupName" missing output "dBSubnetGroupName"
        - `🟡` "dBSystemId" missing output "dBSystemId"
        - `🟡` "domainIAMRoleName" missing output "domainIAMRoleName"
        - `🟡` "enableIAMDatabaseAuthentication" missing output "enableIAMDatabaseAuthentication"
        - `🟡` "multiAZ" missing output "multiAZ"
        - `🟡` "performanceInsightsKMSKeyId" missing output "performanceInsightsKMSKeyId"
        - `🟡` "sourceDBClusterIdentifier" missing output "sourceDBClusterIdentifier"
        - `🟡` "sourceDBInstanceAutomatedBackupsArn" missing output "sourceDBInstanceAutomatedBackupsArn"
        - `🟡` "sourceDBInstanceIdentifier" missing output "sourceDBInstanceIdentifier"
        - `🟡` "vPCSecurityGroups" missing output "vPCSecurityGroups"
- "aws-native:rds:DBParameterGroup":
    - `🟡` inputs: "dBParameterGroupName" missing
    - `🟡` properties: "dBParameterGroupName" missing output "dBParameterGroupName"
- "aws-native:rds:DBProxy":
    - inputs:
        - `🟡` "dBProxyName" missing
        - `🟡` "requireTLS" missing
    - properties:
        - `🟡` "dBProxyArn" missing output "dBProxyArn"
        - `🟡` "dBProxyName" missing output "dBProxyName"
        - `🟡` "requireTLS" missing output "requireTLS"
- "aws-native:rds:DBProxyEndpoint":
    - inputs:
        - `🟡` "dBProxyEndpointName" missing
        - `🟡` "dBProxyName" missing
    - properties:
        - `🟡` "dBProxyEndpointArn" missing output "dBProxyEndpointArn"
        - `🟡` "dBProxyEndpointName" missing output "dBProxyEndpointName"
        - `🟡` "dBProxyName" missing output "dBProxyName"
    - `🟢` required inputs: "dbProxyName" input has changed to Required
- "aws-native:rds:DBProxyTargetGroup":
    - inputs:
        - `🟡` "dBClusterIdentifiers" missing
        - `🟡` "dBInstanceIdentifiers" missing
        - `🟡` "dBProxyName" missing
    - properties:
        - `🟡` "dBClusterIdentifiers" missing output "dBClusterIdentifiers"
        - `🟡` "dBInstanceIdentifiers" missing output "dBInstanceIdentifiers"
        - `🟡` "dBProxyName" missing output "dBProxyName"
    - `🟢` required inputs: "dbProxyName" input has changed to Required
- "aws-native:rds:DBSubnetGroup":
    - inputs:
        - `🟡` "dBSubnetGroupDescription" missing
        - `🟡` "dBSubnetGroupName" missing
    - properties:
        - `🟡` "dBSubnetGroupDescription" missing output "dBSubnetGroupDescription"
        - `🟡` "dBSubnetGroupName" missing output "dBSubnetGroupName"
    - `🟢` required inputs: "dbSubnetGroupDescription" input has changed to Required
- "aws-native:rds:GlobalCluster":
    - `🟡` inputs: "sourceDBClusterIdentifier" missing
    - `🟡` properties: "sourceDBClusterIdentifier" missing output "sourceDBClusterIdentifier"
- "aws-native:redshift:Cluster":
    - `🟡` inputs: "dBName" missing
    - `🟡` properties: "dBName" missing output "dBName"
    - `🟢` required inputs: "dbName" input has changed to Required
- "aws-native:redshift:EndpointAuthorization": properties:
    - `🟡` "allowedAllVPCs" missing output "allowedAllVPCs"
    - `🟡` "allowedVPCs" missing output "allowedVPCs"
- "aws-native:route53:HostedZone":
    - `🟡` inputs: "vPCs" missing
    - `🟡` properties: "vPCs" missing output "vPCs"
- "aws-native:route53resolver:ResolverRuleAssociation":
    - `🟡` inputs: "vPCId" missing
    - `🟡` properties: "vPCId" missing output "vPCId"
    - `🟢` required inputs: "vpcId" input has changed to Required
- `🟡` "aws-native:s3:Bucket": properties: "websiteURL" missing output "websiteURL"
- "aws-native:sagemaker:ImageVersion":
    - `🟡` inputs: "mLFramework" missing
    - `🟡` properties: "mLFramework" missing output "mLFramework"
- "aws-native:ses:EmailIdentity": properties:
    - `🟡` "dkimDNSTokenName1" missing output "dkimDNSTokenName1"
    - `🟡` "dkimDNSTokenName2" missing output "dkimDNSTokenName2"
    - `🟡` "dkimDNSTokenName3" missing output "dkimDNSTokenName3"
    - `🟡` "dkimDNSTokenValue1" missing output "dkimDNSTokenValue1"
    - `🟡` "dkimDNSTokenValue2" missing output "dkimDNSTokenValue2"
    - `🟡` "dkimDNSTokenValue3" missing output "dkimDNSTokenValue3"
- "aws-native:ssm:ResourceDataSync":
    - `🟡` inputs: "kMSKeyArn" missing
    - `🟡` properties: "kMSKeyArn" missing output "kMSKeyArn"
- "aws-native:synthetics:Canary":
    - `🟡` inputs: "vPCConfig" missing
    - `🟡` properties: "vPCConfig" missing output "vPCConfig"
- "aws-native:timestream:ScheduledQuery": properties:
    - `🟡` "sQErrorReportConfiguration" missing output "sQErrorReportConfiguration"
    - `🟡` "sQKmsKeyId" missing output "sQKmsKeyId"
    - `🟡` "sQName" missing output "sQName"
    - `🟡` "sQNotificationConfiguration" missing output "sQNotificationConfiguration"
    - `🟡` "sQQueryString" missing output "sQQueryString"
    - `🟡` "sQScheduleConfiguration" missing output "sQScheduleConfiguration"
    - `🟡` "sQScheduledQueryExecutionRoleArn" missing output "sQScheduledQueryExecutionRoleArn"
    - `🟡` "sQTargetConfiguration" missing output "sQTargetConfiguration"
- "aws-native:wafv2:IPSet":
    - `🟡` inputs: "iPAddressVersion" missing
    - `🟡` properties: "iPAddressVersion" missing output "iPAddressVersion"
    - `🟢` required inputs: "ipAddressVersion" input has changed to Required
- "aws-native:wafv2:WebACLAssociation":
    - `🟡` inputs: "webACLArn" missing
    - `🟡` properties: "webACLArn" missing output "webACLArn"
    - `🟢` required inputs: "webAclArn" input has changed to Required
- `🟡` "aws-native:xray:Group": properties: "groupARN" missing output "groupARN"
- `🟡` "aws-native:xray:SamplingRule": properties: "ruleARN" missing output "ruleARN"
#### Functions
- "aws-native:apigateway:getApiKey": inputs:
    - `🟡` "aPIKeyId" missing input "aPIKeyId"
    - `🟢` required: "apiKeyId" input has changed to Required
- "aws-native:applicationinsights:getApplication": inputs:
    - `🟡` "applicationARN" missing input "applicationARN"No new resources/functions.

## 0.71.0

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
