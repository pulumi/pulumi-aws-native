## HEAD (Unreleased)

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
