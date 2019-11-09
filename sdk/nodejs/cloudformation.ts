import * as pulumi from "@pulumi/pulumi";

const __config = new pulumi.Config("cloudformation");
export const region = __config.get("region");
export const stack = __config.get("stack");

export interface StackArgs {
    stack?: pulumi.Input<string>;
    region: pulumi.Input<string>;
}

export class Stack extends pulumi.ProviderResource {
    constructor(name: string, args: StackArgs, opts?: pulumi.ResourceOptions) {
        const inputs: pulumi.Inputs = {
            stack: args.stack || name,
            region: args.region,
        };

        super("cloudformation", name, inputs, opts);
    }
}

export function getAccountId(opts?: pulumi.InvokeOptions): Promise<string> {
    return pulumi.runtime.invoke("cloudformation:index:getAccountId", {}, opts).then((res: any) => res.accountId);
}

export function getAzs(opts?: pulumi.InvokeOptions): Promise<string[]> {
    return pulumi.runtime.invoke("cloudformation:index:getAzs", {}, opts).then((res: any) => res.azs);
}

export function getStackId(opts?: pulumi.InvokeOptions): Promise<string> {
    return pulumi.runtime.invoke("cloudformation:index:getStackId", {}, opts).then((res: any) => res.stackId);
}

export interface ResourceOptions {
    logicalId?: pulumi.Input<string>;
    metadata?: pulumi.Inputs;
    creationPolicy?: pulumi.Inputs;
    deletionPolicy?: pulumi.Inputs;
    updatePolicy?: pulumi.Inputs;
    updateReplacePolicy?: pulumi.Inputs;
}

export interface TagAttributes {
    Key: string;
    Value: string;
}

export interface TagProperties {
    Key: pulumi.Input<string>;
    Value: pulumi.Input<string>;
}


export namespace ask {
    
    export interface SkillAttributes {
    }
    
    export interface SkillAuthenticationConfigurationAttributes {
        ClientId: string;
        ClientSecret: string;
        RefreshToken: string;
    }
    
    export interface SkillAuthenticationConfigurationProperties {
        ClientId: pulumi.Input<string>;
        ClientSecret: pulumi.Input<string>;
        RefreshToken: pulumi.Input<string>;
    }
    
    export interface SkillOverridesAttributes {
        Manifest?: string;
    }
    
    export interface SkillOverridesProperties {
        Manifest?: pulumi.Input<string>;
    }
    
    export interface SkillProperties {
        AuthenticationConfiguration: pulumi.Input<SkillAuthenticationConfigurationProperties>;
        SkillPackage: pulumi.Input<SkillSkillPackageProperties>;
        VendorId: pulumi.Input<string>;
    }
    
    export interface SkillSkillPackageAttributes {
        Overrides?: SkillOverridesAttributes;
        S3Bucket: string;
        S3BucketRole?: string;
        S3Key: string;
        S3ObjectVersion?: string;
    }
    
    export interface SkillSkillPackageProperties {
        Overrides?: pulumi.Input<SkillOverridesProperties>;
        S3Bucket: pulumi.Input<string>;
        S3BucketRole?: pulumi.Input<string>;
        S3Key: pulumi.Input<string>;
        S3ObjectVersion?: pulumi.Input<string>;
    }
    
    
    export class Skill extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SkillAttributes>;

        constructor(name: string, properties: SkillProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ASK:Skill", name, inputs, opts);
        }
    }
    
}

export namespace amazonmq {
    
    export interface BrokerAttributes {
        AmqpEndpoints: string[];
        Arn: string;
        ConfigurationId: string;
        ConfigurationRevision: number;
        IpAddresses: string[];
        MqttEndpoints: string[];
        OpenWireEndpoints: string[];
        StompEndpoints: string[];
        WssEndpoints: string[];
    }
    
    export interface BrokerConfigurationIdAttributes {
        Id: string;
        Revision: number;
    }
    
    export interface BrokerConfigurationIdProperties {
        Id: pulumi.Input<string>;
        Revision: pulumi.Input<number>;
    }
    
    export interface BrokerEncryptionOptionsAttributes {
        KmsKeyId?: string;
        UseAwsOwnedKey: boolean;
    }
    
    export interface BrokerEncryptionOptionsProperties {
        KmsKeyId?: pulumi.Input<string>;
        UseAwsOwnedKey: pulumi.Input<boolean>;
    }
    
    export interface BrokerLogListAttributes {
        Audit?: boolean;
        General?: boolean;
    }
    
    export interface BrokerLogListProperties {
        Audit?: pulumi.Input<boolean>;
        General?: pulumi.Input<boolean>;
    }
    
    export interface BrokerMaintenanceWindowAttributes {
        DayOfWeek: string;
        TimeOfDay: string;
        TimeZone: string;
    }
    
    export interface BrokerMaintenanceWindowProperties {
        DayOfWeek: pulumi.Input<string>;
        TimeOfDay: pulumi.Input<string>;
        TimeZone: pulumi.Input<string>;
    }
    
    export interface BrokerProperties {
        AutoMinorVersionUpgrade: pulumi.Input<boolean>;
        BrokerName: pulumi.Input<string>;
        Configuration?: pulumi.Input<BrokerConfigurationIdProperties>;
        DeploymentMode: pulumi.Input<string>;
        EncryptionOptions?: pulumi.Input<BrokerEncryptionOptionsProperties>;
        EngineType: pulumi.Input<string>;
        EngineVersion: pulumi.Input<string>;
        HostInstanceType: pulumi.Input<string>;
        Logs?: pulumi.Input<BrokerLogListProperties>;
        MaintenanceWindowStartTime?: pulumi.Input<BrokerMaintenanceWindowProperties>;
        PubliclyAccessible: pulumi.Input<boolean>;
        SecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
        SubnetIds?: pulumi.Input<pulumi.Input<string>[]>;
        Tags?: pulumi.Input<pulumi.Input<BrokerTagsEntryProperties>[]>;
        Users: pulumi.Input<pulumi.Input<BrokerUserProperties>[]>;
    }
    
    export interface BrokerTagsEntryAttributes {
        Key: string;
        Value: string;
    }
    
    export interface BrokerTagsEntryProperties {
        Key: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface BrokerUserAttributes {
        ConsoleAccess?: boolean;
        Groups?: string[];
        Password: string;
        Username: string;
    }
    
    export interface BrokerUserProperties {
        ConsoleAccess?: pulumi.Input<boolean>;
        Groups?: pulumi.Input<pulumi.Input<string>[]>;
        Password: pulumi.Input<string>;
        Username: pulumi.Input<string>;
    }
    
    export interface ConfigurationAssociationAttributes {
    }
    
    export interface ConfigurationAssociationConfigurationIdAttributes {
        Id: string;
        Revision: number;
    }
    
    export interface ConfigurationAssociationConfigurationIdProperties {
        Id: pulumi.Input<string>;
        Revision: pulumi.Input<number>;
    }
    
    export interface ConfigurationAssociationProperties {
        Broker: pulumi.Input<string>;
        Configuration: pulumi.Input<ConfigurationAssociationConfigurationIdProperties>;
    }
    
    export interface ConfigurationAttributes {
        Arn: string;
        Id: string;
        Revision: number;
    }
    
    export interface ConfigurationProperties {
        Data: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        EngineType: pulumi.Input<string>;
        EngineVersion: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<ConfigurationTagsEntryProperties>[]>;
    }
    
    export interface ConfigurationTagsEntryAttributes {
        Key: string;
        Value: string;
    }
    
    export interface ConfigurationTagsEntryProperties {
        Key: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    
    export class Broker extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<BrokerAttributes>;

        constructor(name: string, properties: BrokerProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AmazonMQ:Broker", name, inputs, opts);
        }
    }
    
    export class Configuration extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ConfigurationAttributes>;

        constructor(name: string, properties: ConfigurationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AmazonMQ:Configuration", name, inputs, opts);
        }
    }
    
    export class ConfigurationAssociation extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ConfigurationAssociationAttributes>;

        constructor(name: string, properties: ConfigurationAssociationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AmazonMQ:ConfigurationAssociation", name, inputs, opts);
        }
    }
    
}

export namespace amplify {
    
    export interface AppAttributes {
        AppId: string;
        AppName: string;
        Arn: string;
        DefaultDomain: string;
    }
    
    export interface AppAutoBranchCreationConfigAttributes {
        AutoBranchCreationPatterns?: string[];
        BasicAuthConfig?: AppBasicAuthConfigAttributes;
        BuildSpec?: string;
        EnableAutoBranchCreation?: boolean;
        EnableAutoBuild?: boolean;
        EnablePullRequestPreview?: boolean;
        EnvironmentVariables?: AppEnvironmentVariableAttributes[];
        PullRequestEnvironmentName?: string;
        Stage?: string;
    }
    
    export interface AppAutoBranchCreationConfigProperties {
        AutoBranchCreationPatterns?: pulumi.Input<pulumi.Input<string>[]>;
        BasicAuthConfig?: pulumi.Input<AppBasicAuthConfigProperties>;
        BuildSpec?: pulumi.Input<string>;
        EnableAutoBranchCreation?: pulumi.Input<boolean>;
        EnableAutoBuild?: pulumi.Input<boolean>;
        EnablePullRequestPreview?: pulumi.Input<boolean>;
        EnvironmentVariables?: pulumi.Input<pulumi.Input<AppEnvironmentVariableProperties>[]>;
        PullRequestEnvironmentName?: pulumi.Input<string>;
        Stage?: pulumi.Input<string>;
    }
    
    export interface AppBasicAuthConfigAttributes {
        EnableBasicAuth?: boolean;
        Password?: string;
        Username?: string;
    }
    
    export interface AppBasicAuthConfigProperties {
        EnableBasicAuth?: pulumi.Input<boolean>;
        Password?: pulumi.Input<string>;
        Username?: pulumi.Input<string>;
    }
    
    export interface AppCustomRuleAttributes {
        Condition?: string;
        Source: string;
        Status?: string;
        Target: string;
    }
    
    export interface AppCustomRuleProperties {
        Condition?: pulumi.Input<string>;
        Source: pulumi.Input<string>;
        Status?: pulumi.Input<string>;
        Target: pulumi.Input<string>;
    }
    
    export interface AppEnvironmentVariableAttributes {
        Name: string;
        Value: string;
    }
    
    export interface AppEnvironmentVariableProperties {
        Name: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface AppProperties {
        AccessToken?: pulumi.Input<string>;
        AutoBranchCreationConfig?: pulumi.Input<AppAutoBranchCreationConfigProperties>;
        BasicAuthConfig?: pulumi.Input<AppBasicAuthConfigProperties>;
        BuildSpec?: pulumi.Input<string>;
        CustomRules?: pulumi.Input<pulumi.Input<AppCustomRuleProperties>[]>;
        Description?: pulumi.Input<string>;
        EnvironmentVariables?: pulumi.Input<pulumi.Input<AppEnvironmentVariableProperties>[]>;
        IAMServiceRole?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        OauthToken?: pulumi.Input<string>;
        Repository?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface BranchAttributes {
        Arn: string;
        BranchName: string;
    }
    
    export interface BranchBasicAuthConfigAttributes {
        EnableBasicAuth?: boolean;
        Password: string;
        Username: string;
    }
    
    export interface BranchBasicAuthConfigProperties {
        EnableBasicAuth?: pulumi.Input<boolean>;
        Password: pulumi.Input<string>;
        Username: pulumi.Input<string>;
    }
    
    export interface BranchEnvironmentVariableAttributes {
        Name: string;
        Value: string;
    }
    
    export interface BranchEnvironmentVariableProperties {
        Name: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface BranchProperties {
        AppId: pulumi.Input<string>;
        BasicAuthConfig?: pulumi.Input<BranchBasicAuthConfigProperties>;
        BranchName: pulumi.Input<string>;
        BuildSpec?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        EnableAutoBuild?: pulumi.Input<boolean>;
        EnablePullRequestPreview?: pulumi.Input<boolean>;
        EnvironmentVariables?: pulumi.Input<pulumi.Input<BranchEnvironmentVariableProperties>[]>;
        PullRequestEnvironmentName?: pulumi.Input<string>;
        Stage?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface DomainAttributes {
        Arn: string;
        CertificateRecord: string;
        DomainName: string;
        DomainStatus: string;
        StatusReason: string;
    }
    
    export interface DomainProperties {
        AppId: pulumi.Input<string>;
        DomainName: pulumi.Input<string>;
        SubDomainSettings: pulumi.Input<pulumi.Input<DomainSubDomainSettingProperties>[]>;
    }
    
    export interface DomainSubDomainSettingAttributes {
        BranchName: string;
        Prefix: string;
    }
    
    export interface DomainSubDomainSettingProperties {
        BranchName: pulumi.Input<string>;
        Prefix: pulumi.Input<string>;
    }
    
    
    export class App extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AppAttributes>;

        constructor(name: string, properties: AppProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Amplify:App", name, inputs, opts);
        }
    }
    
    export class Branch extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<BranchAttributes>;

        constructor(name: string, properties: BranchProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Amplify:Branch", name, inputs, opts);
        }
    }
    
    export class Domain extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DomainAttributes>;

        constructor(name: string, properties: DomainProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Amplify:Domain", name, inputs, opts);
        }
    }
    
}

export namespace apigateway {
    
    export interface AccountAttributes {
    }
    
    export interface AccountProperties {
        CloudWatchRoleArn?: pulumi.Input<string>;
    }
    
    export interface ApiKeyAttributes {
    }
    
    export interface ApiKeyProperties {
        CustomerId?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        Enabled?: pulumi.Input<boolean>;
        GenerateDistinctId?: pulumi.Input<boolean>;
        Name?: pulumi.Input<string>;
        StageKeys?: pulumi.Input<pulumi.Input<ApiKeyStageKeyProperties>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        Value?: pulumi.Input<string>;
    }
    
    export interface ApiKeyStageKeyAttributes {
        RestApiId?: string;
        StageName?: string;
    }
    
    export interface ApiKeyStageKeyProperties {
        RestApiId?: pulumi.Input<string>;
        StageName?: pulumi.Input<string>;
    }
    
    export interface AuthorizerAttributes {
    }
    
    export interface AuthorizerProperties {
        AuthType?: pulumi.Input<string>;
        AuthorizerCredentials?: pulumi.Input<string>;
        AuthorizerResultTtlInSeconds?: pulumi.Input<number>;
        AuthorizerUri?: pulumi.Input<string>;
        IdentitySource?: pulumi.Input<string>;
        IdentityValidationExpression?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        ProviderARNs?: pulumi.Input<pulumi.Input<string>[]>;
        RestApiId: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface BasePathMappingAttributes {
    }
    
    export interface BasePathMappingProperties {
        BasePath?: pulumi.Input<string>;
        DomainName: pulumi.Input<string>;
        RestApiId?: pulumi.Input<string>;
        Stage?: pulumi.Input<string>;
    }
    
    export interface ClientCertificateAttributes {
    }
    
    export interface ClientCertificateProperties {
        Description?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface DeploymentAccessLogSettingAttributes {
        DestinationArn?: string;
        Format?: string;
    }
    
    export interface DeploymentAccessLogSettingProperties {
        DestinationArn?: pulumi.Input<string>;
        Format?: pulumi.Input<string>;
    }
    
    export interface DeploymentAttributes {
    }
    
    export interface DeploymentCanarySettingAttributes {
        PercentTraffic?: number;
        StageVariableOverrides?: {[key: string]: string};
        UseStageCache?: boolean;
    }
    
    export interface DeploymentCanarySettingProperties {
        PercentTraffic?: pulumi.Input<number>;
        StageVariableOverrides?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        UseStageCache?: pulumi.Input<boolean>;
    }
    
    export interface DeploymentDeploymentCanarySettingsAttributes {
        PercentTraffic?: number;
        StageVariableOverrides?: {[key: string]: string};
        UseStageCache?: boolean;
    }
    
    export interface DeploymentDeploymentCanarySettingsProperties {
        PercentTraffic?: pulumi.Input<number>;
        StageVariableOverrides?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        UseStageCache?: pulumi.Input<boolean>;
    }
    
    export interface DeploymentMethodSettingAttributes {
        CacheDataEncrypted?: boolean;
        CacheTtlInSeconds?: number;
        CachingEnabled?: boolean;
        DataTraceEnabled?: boolean;
        HttpMethod?: string;
        LoggingLevel?: string;
        MetricsEnabled?: boolean;
        ResourcePath?: string;
        ThrottlingBurstLimit?: number;
        ThrottlingRateLimit?: number;
    }
    
    export interface DeploymentMethodSettingProperties {
        CacheDataEncrypted?: pulumi.Input<boolean>;
        CacheTtlInSeconds?: pulumi.Input<number>;
        CachingEnabled?: pulumi.Input<boolean>;
        DataTraceEnabled?: pulumi.Input<boolean>;
        HttpMethod?: pulumi.Input<string>;
        LoggingLevel?: pulumi.Input<string>;
        MetricsEnabled?: pulumi.Input<boolean>;
        ResourcePath?: pulumi.Input<string>;
        ThrottlingBurstLimit?: pulumi.Input<number>;
        ThrottlingRateLimit?: pulumi.Input<number>;
    }
    
    export interface DeploymentProperties {
        DeploymentCanarySettings?: pulumi.Input<DeploymentDeploymentCanarySettingsProperties>;
        Description?: pulumi.Input<string>;
        RestApiId: pulumi.Input<string>;
        StageDescription?: pulumi.Input<DeploymentStageDescriptionProperties>;
        StageName?: pulumi.Input<string>;
    }
    
    export interface DeploymentStageDescriptionAttributes {
        AccessLogSetting?: DeploymentAccessLogSettingAttributes;
        CacheClusterEnabled?: boolean;
        CacheClusterSize?: string;
        CacheDataEncrypted?: boolean;
        CacheTtlInSeconds?: number;
        CachingEnabled?: boolean;
        CanarySetting?: DeploymentCanarySettingAttributes;
        ClientCertificateId?: string;
        DataTraceEnabled?: boolean;
        Description?: string;
        DocumentationVersion?: string;
        LoggingLevel?: string;
        MethodSettings?: DeploymentMethodSettingAttributes[];
        MetricsEnabled?: boolean;
        Tags?: TagAttributes[];
        ThrottlingBurstLimit?: number;
        ThrottlingRateLimit?: number;
        TracingEnabled?: boolean;
        Variables?: {[key: string]: string};
    }
    
    export interface DeploymentStageDescriptionProperties {
        AccessLogSetting?: pulumi.Input<DeploymentAccessLogSettingProperties>;
        CacheClusterEnabled?: pulumi.Input<boolean>;
        CacheClusterSize?: pulumi.Input<string>;
        CacheDataEncrypted?: pulumi.Input<boolean>;
        CacheTtlInSeconds?: pulumi.Input<number>;
        CachingEnabled?: pulumi.Input<boolean>;
        CanarySetting?: pulumi.Input<DeploymentCanarySettingProperties>;
        ClientCertificateId?: pulumi.Input<string>;
        DataTraceEnabled?: pulumi.Input<boolean>;
        Description?: pulumi.Input<string>;
        DocumentationVersion?: pulumi.Input<string>;
        LoggingLevel?: pulumi.Input<string>;
        MethodSettings?: pulumi.Input<pulumi.Input<DeploymentMethodSettingProperties>[]>;
        MetricsEnabled?: pulumi.Input<boolean>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        ThrottlingBurstLimit?: pulumi.Input<number>;
        ThrottlingRateLimit?: pulumi.Input<number>;
        TracingEnabled?: pulumi.Input<boolean>;
        Variables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    }
    
    export interface DocumentationPartAttributes {
    }
    
    export interface DocumentationPartLocationAttributes {
        Method?: string;
        Name?: string;
        Path?: string;
        StatusCode?: string;
        Type?: string;
    }
    
    export interface DocumentationPartLocationProperties {
        Method?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Path?: pulumi.Input<string>;
        StatusCode?: pulumi.Input<string>;
        Type?: pulumi.Input<string>;
    }
    
    export interface DocumentationPartProperties {
        Location: pulumi.Input<DocumentationPartLocationProperties>;
        Properties: pulumi.Input<string>;
        RestApiId: pulumi.Input<string>;
    }
    
    export interface DocumentationVersionAttributes {
    }
    
    export interface DocumentationVersionProperties {
        Description?: pulumi.Input<string>;
        DocumentationVersion: pulumi.Input<string>;
        RestApiId: pulumi.Input<string>;
    }
    
    export interface DomainNameAttributes {
        DistributionDomainName: string;
        DistributionHostedZoneId: string;
        RegionalDomainName: string;
        RegionalHostedZoneId: string;
    }
    
    export interface DomainNameEndpointConfigurationAttributes {
        Types?: string[];
    }
    
    export interface DomainNameEndpointConfigurationProperties {
        Types?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface DomainNameProperties {
        CertificateArn?: pulumi.Input<string>;
        DomainName: pulumi.Input<string>;
        EndpointConfiguration?: pulumi.Input<DomainNameEndpointConfigurationProperties>;
        RegionalCertificateArn?: pulumi.Input<string>;
        SecurityPolicy?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface GatewayResponseAttributes {
    }
    
    export interface GatewayResponseProperties {
        ResponseParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        ResponseTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        ResponseType: pulumi.Input<string>;
        RestApiId: pulumi.Input<string>;
        StatusCode?: pulumi.Input<string>;
    }
    
    export interface MethodAttributes {
    }
    
    export interface MethodIntegrationAttributes {
        CacheKeyParameters?: string[];
        CacheNamespace?: string;
        ConnectionId?: string;
        ConnectionType?: string;
        ContentHandling?: string;
        Credentials?: string;
        IntegrationHttpMethod?: string;
        IntegrationResponses?: MethodIntegrationResponseAttributes[];
        PassthroughBehavior?: string;
        RequestParameters?: {[key: string]: string};
        RequestTemplates?: {[key: string]: string};
        TimeoutInMillis?: number;
        Type?: string;
        Uri?: string;
    }
    
    export interface MethodIntegrationProperties {
        CacheKeyParameters?: pulumi.Input<pulumi.Input<string>[]>;
        CacheNamespace?: pulumi.Input<string>;
        ConnectionId?: pulumi.Input<string>;
        ConnectionType?: pulumi.Input<string>;
        ContentHandling?: pulumi.Input<string>;
        Credentials?: pulumi.Input<string>;
        IntegrationHttpMethod?: pulumi.Input<string>;
        IntegrationResponses?: pulumi.Input<pulumi.Input<MethodIntegrationResponseProperties>[]>;
        PassthroughBehavior?: pulumi.Input<string>;
        RequestParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        RequestTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        TimeoutInMillis?: pulumi.Input<number>;
        Type?: pulumi.Input<string>;
        Uri?: pulumi.Input<string>;
    }
    
    export interface MethodIntegrationResponseAttributes {
        ContentHandling?: string;
        ResponseParameters?: {[key: string]: string};
        ResponseTemplates?: {[key: string]: string};
        SelectionPattern?: string;
        StatusCode: string;
    }
    
    export interface MethodIntegrationResponseProperties {
        ContentHandling?: pulumi.Input<string>;
        ResponseParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        ResponseTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        SelectionPattern?: pulumi.Input<string>;
        StatusCode: pulumi.Input<string>;
    }
    
    export interface MethodMethodResponseAttributes {
        ResponseModels?: {[key: string]: string};
        ResponseParameters?: {[key: string]: boolean};
        StatusCode: string;
    }
    
    export interface MethodMethodResponseProperties {
        ResponseModels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        ResponseParameters?: pulumi.Input<{[key: string]: pulumi.Input<boolean>}>;
        StatusCode: pulumi.Input<string>;
    }
    
    export interface MethodProperties {
        ApiKeyRequired?: pulumi.Input<boolean>;
        AuthorizationScopes?: pulumi.Input<pulumi.Input<string>[]>;
        AuthorizationType?: pulumi.Input<string>;
        AuthorizerId?: pulumi.Input<string>;
        HttpMethod: pulumi.Input<string>;
        Integration?: pulumi.Input<MethodIntegrationProperties>;
        MethodResponses?: pulumi.Input<pulumi.Input<MethodMethodResponseProperties>[]>;
        OperationName?: pulumi.Input<string>;
        RequestModels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        RequestParameters?: pulumi.Input<{[key: string]: pulumi.Input<boolean>}>;
        RequestValidatorId?: pulumi.Input<string>;
        ResourceId: pulumi.Input<string>;
        RestApiId: pulumi.Input<string>;
    }
    
    export interface ModelAttributes {
    }
    
    export interface ModelProperties {
        ContentType?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        RestApiId: pulumi.Input<string>;
        Schema?: pulumi.Input<string>;
    }
    
    export interface RequestValidatorAttributes {
    }
    
    export interface RequestValidatorProperties {
        Name?: pulumi.Input<string>;
        RestApiId: pulumi.Input<string>;
        ValidateRequestBody?: pulumi.Input<boolean>;
        ValidateRequestParameters?: pulumi.Input<boolean>;
    }
    
    export interface ResourceAttributes {
    }
    
    export interface ResourceProperties {
        ParentId: pulumi.Input<string>;
        PathPart: pulumi.Input<string>;
        RestApiId: pulumi.Input<string>;
    }
    
    export interface RestApiAttributes {
        RootResourceId: string;
    }
    
    export interface RestApiEndpointConfigurationAttributes {
        Types?: string[];
    }
    
    export interface RestApiEndpointConfigurationProperties {
        Types?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface RestApiProperties {
        ApiKeySourceType?: pulumi.Input<string>;
        BinaryMediaTypes?: pulumi.Input<pulumi.Input<string>[]>;
        Body?: pulumi.Input<string>;
        BodyS3Location?: pulumi.Input<RestApiS3LocationProperties>;
        CloneFrom?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        EndpointConfiguration?: pulumi.Input<RestApiEndpointConfigurationProperties>;
        FailOnWarnings?: pulumi.Input<boolean>;
        MinimumCompressionSize?: pulumi.Input<number>;
        Name?: pulumi.Input<string>;
        Parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        Policy?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface RestApiS3LocationAttributes {
        Bucket?: string;
        ETag?: string;
        Key?: string;
        Version?: string;
    }
    
    export interface RestApiS3LocationProperties {
        Bucket?: pulumi.Input<string>;
        ETag?: pulumi.Input<string>;
        Key?: pulumi.Input<string>;
        Version?: pulumi.Input<string>;
    }
    
    export interface StageAccessLogSettingAttributes {
        DestinationArn?: string;
        Format?: string;
    }
    
    export interface StageAccessLogSettingProperties {
        DestinationArn?: pulumi.Input<string>;
        Format?: pulumi.Input<string>;
    }
    
    export interface StageAttributes {
    }
    
    export interface StageCanarySettingAttributes {
        DeploymentId?: string;
        PercentTraffic?: number;
        StageVariableOverrides?: {[key: string]: string};
        UseStageCache?: boolean;
    }
    
    export interface StageCanarySettingProperties {
        DeploymentId?: pulumi.Input<string>;
        PercentTraffic?: pulumi.Input<number>;
        StageVariableOverrides?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        UseStageCache?: pulumi.Input<boolean>;
    }
    
    export interface StageMethodSettingAttributes {
        CacheDataEncrypted?: boolean;
        CacheTtlInSeconds?: number;
        CachingEnabled?: boolean;
        DataTraceEnabled?: boolean;
        HttpMethod?: string;
        LoggingLevel?: string;
        MetricsEnabled?: boolean;
        ResourcePath?: string;
        ThrottlingBurstLimit?: number;
        ThrottlingRateLimit?: number;
    }
    
    export interface StageMethodSettingProperties {
        CacheDataEncrypted?: pulumi.Input<boolean>;
        CacheTtlInSeconds?: pulumi.Input<number>;
        CachingEnabled?: pulumi.Input<boolean>;
        DataTraceEnabled?: pulumi.Input<boolean>;
        HttpMethod?: pulumi.Input<string>;
        LoggingLevel?: pulumi.Input<string>;
        MetricsEnabled?: pulumi.Input<boolean>;
        ResourcePath?: pulumi.Input<string>;
        ThrottlingBurstLimit?: pulumi.Input<number>;
        ThrottlingRateLimit?: pulumi.Input<number>;
    }
    
    export interface StageProperties {
        AccessLogSetting?: pulumi.Input<StageAccessLogSettingProperties>;
        CacheClusterEnabled?: pulumi.Input<boolean>;
        CacheClusterSize?: pulumi.Input<string>;
        CanarySetting?: pulumi.Input<StageCanarySettingProperties>;
        ClientCertificateId?: pulumi.Input<string>;
        DeploymentId?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        DocumentationVersion?: pulumi.Input<string>;
        MethodSettings?: pulumi.Input<pulumi.Input<StageMethodSettingProperties>[]>;
        RestApiId: pulumi.Input<string>;
        StageName?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        TracingEnabled?: pulumi.Input<boolean>;
        Variables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    }
    
    export interface UsagePlanApiStageAttributes {
        ApiId?: string;
        Stage?: string;
        Throttle?: {[key: string]: UsagePlanThrottleSettingsAttributes};
    }
    
    export interface UsagePlanApiStageProperties {
        ApiId?: pulumi.Input<string>;
        Stage?: pulumi.Input<string>;
        Throttle?: pulumi.Input<{[key: string]: pulumi.Input<UsagePlanThrottleSettingsProperties>}>;
    }
    
    export interface UsagePlanAttributes {
    }
    
    export interface UsagePlanKeyAttributes {
    }
    
    export interface UsagePlanKeyProperties {
        KeyId: pulumi.Input<string>;
        KeyType: pulumi.Input<string>;
        UsagePlanId: pulumi.Input<string>;
    }
    
    export interface UsagePlanProperties {
        ApiStages?: pulumi.Input<pulumi.Input<UsagePlanApiStageProperties>[]>;
        Description?: pulumi.Input<string>;
        Quota?: pulumi.Input<UsagePlanQuotaSettingsProperties>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        Throttle?: pulumi.Input<UsagePlanThrottleSettingsProperties>;
        UsagePlanName?: pulumi.Input<string>;
    }
    
    export interface UsagePlanQuotaSettingsAttributes {
        Limit?: number;
        Offset?: number;
        Period?: string;
    }
    
    export interface UsagePlanQuotaSettingsProperties {
        Limit?: pulumi.Input<number>;
        Offset?: pulumi.Input<number>;
        Period?: pulumi.Input<string>;
    }
    
    export interface UsagePlanThrottleSettingsAttributes {
        BurstLimit?: number;
        RateLimit?: number;
    }
    
    export interface UsagePlanThrottleSettingsProperties {
        BurstLimit?: pulumi.Input<number>;
        RateLimit?: pulumi.Input<number>;
    }
    
    export interface VpcLinkAttributes {
    }
    
    export interface VpcLinkProperties {
        Description?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        TargetArns: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    
    export class Account extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AccountAttributes>;

        constructor(name: string, properties?: AccountProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:Account", name, inputs, opts);
        }
    }
    
    export class ApiKey extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ApiKeyAttributes>;

        constructor(name: string, properties?: ApiKeyProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:ApiKey", name, inputs, opts);
        }
    }
    
    export class Authorizer extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AuthorizerAttributes>;

        constructor(name: string, properties: AuthorizerProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:Authorizer", name, inputs, opts);
        }
    }
    
    export class BasePathMapping extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<BasePathMappingAttributes>;

        constructor(name: string, properties: BasePathMappingProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:BasePathMapping", name, inputs, opts);
        }
    }
    
    export class ClientCertificate extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ClientCertificateAttributes>;

        constructor(name: string, properties?: ClientCertificateProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:ClientCertificate", name, inputs, opts);
        }
    }
    
    export class Deployment extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DeploymentAttributes>;

        constructor(name: string, properties: DeploymentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:Deployment", name, inputs, opts);
        }
    }
    
    export class DocumentationPart extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DocumentationPartAttributes>;

        constructor(name: string, properties: DocumentationPartProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:DocumentationPart", name, inputs, opts);
        }
    }
    
    export class DocumentationVersion extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DocumentationVersionAttributes>;

        constructor(name: string, properties: DocumentationVersionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:DocumentationVersion", name, inputs, opts);
        }
    }
    
    export class DomainName extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DomainNameAttributes>;

        constructor(name: string, properties: DomainNameProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:DomainName", name, inputs, opts);
        }
    }
    
    export class GatewayResponse extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<GatewayResponseAttributes>;

        constructor(name: string, properties: GatewayResponseProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:GatewayResponse", name, inputs, opts);
        }
    }
    
    export class Method extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<MethodAttributes>;

        constructor(name: string, properties: MethodProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:Method", name, inputs, opts);
        }
    }
    
    export class Model extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ModelAttributes>;

        constructor(name: string, properties: ModelProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:Model", name, inputs, opts);
        }
    }
    
    export class RequestValidator extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RequestValidatorAttributes>;

        constructor(name: string, properties: RequestValidatorProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:RequestValidator", name, inputs, opts);
        }
    }
    
    export class Resource extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ResourceAttributes>;

        constructor(name: string, properties: ResourceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:Resource", name, inputs, opts);
        }
    }
    
    export class RestApi extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RestApiAttributes>;

        constructor(name: string, properties?: RestApiProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:RestApi", name, inputs, opts);
        }
    }
    
    export class Stage extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<StageAttributes>;

        constructor(name: string, properties: StageProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:Stage", name, inputs, opts);
        }
    }
    
    export class UsagePlan extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<UsagePlanAttributes>;

        constructor(name: string, properties?: UsagePlanProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:UsagePlan", name, inputs, opts);
        }
    }
    
    export class UsagePlanKey extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<UsagePlanKeyAttributes>;

        constructor(name: string, properties: UsagePlanKeyProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:UsagePlanKey", name, inputs, opts);
        }
    }
    
    export class VpcLink extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VpcLinkAttributes>;

        constructor(name: string, properties: VpcLinkProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGateway:VpcLink", name, inputs, opts);
        }
    }
    
}

export namespace apigatewayv2 {
    
    export interface ApiAttributes {
    }
    
    export interface ApiMappingAttributes {
    }
    
    export interface ApiMappingProperties {
        ApiId: pulumi.Input<string>;
        ApiMappingKey?: pulumi.Input<string>;
        DomainName: pulumi.Input<string>;
        Stage: pulumi.Input<string>;
    }
    
    export interface ApiProperties {
        ApiKeySelectionExpression?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        DisableSchemaValidation?: pulumi.Input<boolean>;
        Name: pulumi.Input<string>;
        ProtocolType: pulumi.Input<string>;
        RouteSelectionExpression: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
        Version?: pulumi.Input<string>;
    }
    
    export interface AuthorizerAttributes {
    }
    
    export interface AuthorizerProperties {
        ApiId: pulumi.Input<string>;
        AuthorizerCredentialsArn?: pulumi.Input<string>;
        AuthorizerResultTtlInSeconds?: pulumi.Input<number>;
        AuthorizerType: pulumi.Input<string>;
        AuthorizerUri: pulumi.Input<string>;
        IdentitySource: pulumi.Input<pulumi.Input<string>[]>;
        IdentityValidationExpression?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
    }
    
    export interface DeploymentAttributes {
    }
    
    export interface DeploymentProperties {
        ApiId: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        StageName?: pulumi.Input<string>;
    }
    
    export interface DomainNameAttributes {
        RegionalDomainName: string;
        RegionalHostedZoneId: string;
    }
    
    export interface DomainNameDomainNameConfigurationAttributes {
        CertificateArn?: string;
        CertificateName?: string;
        EndpointType?: string;
    }
    
    export interface DomainNameDomainNameConfigurationProperties {
        CertificateArn?: pulumi.Input<string>;
        CertificateName?: pulumi.Input<string>;
        EndpointType?: pulumi.Input<string>;
    }
    
    export interface DomainNameProperties {
        DomainName: pulumi.Input<string>;
        DomainNameConfigurations?: pulumi.Input<pulumi.Input<DomainNameDomainNameConfigurationProperties>[]>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface IntegrationAttributes {
    }
    
    export interface IntegrationProperties {
        ApiId: pulumi.Input<string>;
        ConnectionType?: pulumi.Input<string>;
        ContentHandlingStrategy?: pulumi.Input<string>;
        CredentialsArn?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        IntegrationMethod?: pulumi.Input<string>;
        IntegrationType: pulumi.Input<string>;
        IntegrationUri?: pulumi.Input<string>;
        PassthroughBehavior?: pulumi.Input<string>;
        RequestParameters?: pulumi.Input<string>;
        RequestTemplates?: pulumi.Input<string>;
        TemplateSelectionExpression?: pulumi.Input<string>;
        TimeoutInMillis?: pulumi.Input<number>;
    }
    
    export interface IntegrationResponseAttributes {
    }
    
    export interface IntegrationResponseProperties {
        ApiId: pulumi.Input<string>;
        ContentHandlingStrategy?: pulumi.Input<string>;
        IntegrationId: pulumi.Input<string>;
        IntegrationResponseKey: pulumi.Input<string>;
        ResponseParameters?: pulumi.Input<string>;
        ResponseTemplates?: pulumi.Input<string>;
        TemplateSelectionExpression?: pulumi.Input<string>;
    }
    
    export interface ModelAttributes {
    }
    
    export interface ModelProperties {
        ApiId: pulumi.Input<string>;
        ContentType?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        Schema: pulumi.Input<string>;
    }
    
    export interface RouteAttributes {
    }
    
    export interface RouteParameterConstraintsAttributes {
        Required: boolean;
    }
    
    export interface RouteParameterConstraintsProperties {
        Required: pulumi.Input<boolean>;
    }
    
    export interface RouteProperties {
        ApiId: pulumi.Input<string>;
        ApiKeyRequired?: pulumi.Input<boolean>;
        AuthorizationScopes?: pulumi.Input<pulumi.Input<string>[]>;
        AuthorizationType?: pulumi.Input<string>;
        AuthorizerId?: pulumi.Input<string>;
        ModelSelectionExpression?: pulumi.Input<string>;
        OperationName?: pulumi.Input<string>;
        RequestModels?: pulumi.Input<string>;
        RequestParameters?: pulumi.Input<string>;
        RouteKey: pulumi.Input<string>;
        RouteResponseSelectionExpression?: pulumi.Input<string>;
        Target?: pulumi.Input<string>;
    }
    
    export interface RouteResponseAttributes {
    }
    
    export interface RouteResponseParameterConstraintsAttributes {
        Required: boolean;
    }
    
    export interface RouteResponseParameterConstraintsProperties {
        Required: pulumi.Input<boolean>;
    }
    
    export interface RouteResponseProperties {
        ApiId: pulumi.Input<string>;
        ModelSelectionExpression?: pulumi.Input<string>;
        ResponseModels?: pulumi.Input<string>;
        ResponseParameters?: pulumi.Input<string>;
        RouteId: pulumi.Input<string>;
        RouteResponseKey: pulumi.Input<string>;
    }
    
    export interface StageAccessLogSettingsAttributes {
        DestinationArn?: string;
        Format?: string;
    }
    
    export interface StageAccessLogSettingsProperties {
        DestinationArn?: pulumi.Input<string>;
        Format?: pulumi.Input<string>;
    }
    
    export interface StageAttributes {
    }
    
    export interface StageProperties {
        AccessLogSettings?: pulumi.Input<StageAccessLogSettingsProperties>;
        ApiId: pulumi.Input<string>;
        ClientCertificateId?: pulumi.Input<string>;
        DefaultRouteSettings?: pulumi.Input<StageRouteSettingsProperties>;
        DeploymentId: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        RouteSettings?: pulumi.Input<string>;
        StageName: pulumi.Input<string>;
        StageVariables?: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface StageRouteSettingsAttributes {
        DataTraceEnabled?: boolean;
        DetailedMetricsEnabled?: boolean;
        LoggingLevel?: string;
        ThrottlingBurstLimit?: number;
        ThrottlingRateLimit?: number;
    }
    
    export interface StageRouteSettingsProperties {
        DataTraceEnabled?: pulumi.Input<boolean>;
        DetailedMetricsEnabled?: pulumi.Input<boolean>;
        LoggingLevel?: pulumi.Input<string>;
        ThrottlingBurstLimit?: pulumi.Input<number>;
        ThrottlingRateLimit?: pulumi.Input<number>;
    }
    
    
    export class Api extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ApiAttributes>;

        constructor(name: string, properties: ApiProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGatewayV2:Api", name, inputs, opts);
        }
    }
    
    export class ApiMapping extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ApiMappingAttributes>;

        constructor(name: string, properties: ApiMappingProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGatewayV2:ApiMapping", name, inputs, opts);
        }
    }
    
    export class Authorizer extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AuthorizerAttributes>;

        constructor(name: string, properties: AuthorizerProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGatewayV2:Authorizer", name, inputs, opts);
        }
    }
    
    export class Deployment extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DeploymentAttributes>;

        constructor(name: string, properties: DeploymentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGatewayV2:Deployment", name, inputs, opts);
        }
    }
    
    export class DomainName extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DomainNameAttributes>;

        constructor(name: string, properties: DomainNameProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGatewayV2:DomainName", name, inputs, opts);
        }
    }
    
    export class Integration extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<IntegrationAttributes>;

        constructor(name: string, properties: IntegrationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGatewayV2:Integration", name, inputs, opts);
        }
    }
    
    export class IntegrationResponse extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<IntegrationResponseAttributes>;

        constructor(name: string, properties: IntegrationResponseProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGatewayV2:IntegrationResponse", name, inputs, opts);
        }
    }
    
    export class Model extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ModelAttributes>;

        constructor(name: string, properties: ModelProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGatewayV2:Model", name, inputs, opts);
        }
    }
    
    export class Route extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RouteAttributes>;

        constructor(name: string, properties: RouteProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGatewayV2:Route", name, inputs, opts);
        }
    }
    
    export class RouteResponse extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RouteResponseAttributes>;

        constructor(name: string, properties: RouteResponseProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGatewayV2:RouteResponse", name, inputs, opts);
        }
    }
    
    export class Stage extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<StageAttributes>;

        constructor(name: string, properties: StageProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApiGatewayV2:Stage", name, inputs, opts);
        }
    }
    
}

export namespace appmesh {
    
    export interface MeshAttributes {
        Arn: string;
        MeshName: string;
        Uid: string;
    }
    
    export interface MeshEgressFilterAttributes {
        Type: string;
    }
    
    export interface MeshEgressFilterProperties {
        Type: pulumi.Input<string>;
    }
    
    export interface MeshMeshSpecAttributes {
        EgressFilter?: MeshEgressFilterAttributes;
    }
    
    export interface MeshMeshSpecProperties {
        EgressFilter?: pulumi.Input<MeshEgressFilterProperties>;
    }
    
    export interface MeshProperties {
        MeshName: pulumi.Input<string>;
        Spec?: pulumi.Input<MeshMeshSpecProperties>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface RouteAttributes {
        Arn: string;
        MeshName: string;
        RouteName: string;
        Uid: string;
        VirtualRouterName: string;
    }
    
    export interface RouteDurationAttributes {
        Unit: string;
        Value: number;
    }
    
    export interface RouteDurationProperties {
        Unit: pulumi.Input<string>;
        Value: pulumi.Input<number>;
    }
    
    export interface RouteGrpcRetryPolicyAttributes {
        GrpcRetryEvents?: string[];
        HttpRetryEvents?: string[];
        MaxRetries: number;
        PerRetryTimeout: RouteDurationAttributes;
        TcpRetryEvents?: string[];
    }
    
    export interface RouteGrpcRetryPolicyProperties {
        GrpcRetryEvents?: pulumi.Input<pulumi.Input<string>[]>;
        HttpRetryEvents?: pulumi.Input<pulumi.Input<string>[]>;
        MaxRetries: pulumi.Input<number>;
        PerRetryTimeout: pulumi.Input<RouteDurationProperties>;
        TcpRetryEvents?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface RouteGrpcRouteActionAttributes {
        WeightedTargets: RouteWeightedTargetAttributes[];
    }
    
    export interface RouteGrpcRouteActionProperties {
        WeightedTargets: pulumi.Input<pulumi.Input<RouteWeightedTargetProperties>[]>;
    }
    
    export interface RouteGrpcRouteAttributes {
        Action: RouteGrpcRouteActionAttributes;
        Match: RouteGrpcRouteMatchAttributes;
        RetryPolicy?: RouteGrpcRetryPolicyAttributes;
    }
    
    export interface RouteGrpcRouteMatchAttributes {
        Metadata?: RouteGrpcRouteMetadataAttributes[];
        MethodName?: string;
        ServiceName?: string;
    }
    
    export interface RouteGrpcRouteMatchProperties {
        Metadata?: pulumi.Input<pulumi.Input<RouteGrpcRouteMetadataProperties>[]>;
        MethodName?: pulumi.Input<string>;
        ServiceName?: pulumi.Input<string>;
    }
    
    export interface RouteGrpcRouteMetadataAttributes {
        Invert?: boolean;
        Match?: RouteGrpcRouteMetadataMatchMethodAttributes;
        Name: string;
    }
    
    export interface RouteGrpcRouteMetadataMatchMethodAttributes {
        Exact?: string;
        Prefix?: string;
        Range?: RouteMatchRangeAttributes;
        Regex?: string;
        Suffix?: string;
    }
    
    export interface RouteGrpcRouteMetadataMatchMethodProperties {
        Exact?: pulumi.Input<string>;
        Prefix?: pulumi.Input<string>;
        Range?: pulumi.Input<RouteMatchRangeProperties>;
        Regex?: pulumi.Input<string>;
        Suffix?: pulumi.Input<string>;
    }
    
    export interface RouteGrpcRouteMetadataProperties {
        Invert?: pulumi.Input<boolean>;
        Match?: pulumi.Input<RouteGrpcRouteMetadataMatchMethodProperties>;
        Name: pulumi.Input<string>;
    }
    
    export interface RouteGrpcRouteProperties {
        Action: pulumi.Input<RouteGrpcRouteActionProperties>;
        Match: pulumi.Input<RouteGrpcRouteMatchProperties>;
        RetryPolicy?: pulumi.Input<RouteGrpcRetryPolicyProperties>;
    }
    
    export interface RouteHeaderMatchMethodAttributes {
        Exact?: string;
        Prefix?: string;
        Range?: RouteMatchRangeAttributes;
        Regex?: string;
        Suffix?: string;
    }
    
    export interface RouteHeaderMatchMethodProperties {
        Exact?: pulumi.Input<string>;
        Prefix?: pulumi.Input<string>;
        Range?: pulumi.Input<RouteMatchRangeProperties>;
        Regex?: pulumi.Input<string>;
        Suffix?: pulumi.Input<string>;
    }
    
    export interface RouteHttpRetryPolicyAttributes {
        HttpRetryEvents?: string[];
        MaxRetries: number;
        PerRetryTimeout: RouteDurationAttributes;
        TcpRetryEvents?: string[];
    }
    
    export interface RouteHttpRetryPolicyProperties {
        HttpRetryEvents?: pulumi.Input<pulumi.Input<string>[]>;
        MaxRetries: pulumi.Input<number>;
        PerRetryTimeout: pulumi.Input<RouteDurationProperties>;
        TcpRetryEvents?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface RouteHttpRouteActionAttributes {
        WeightedTargets: RouteWeightedTargetAttributes[];
    }
    
    export interface RouteHttpRouteActionProperties {
        WeightedTargets: pulumi.Input<pulumi.Input<RouteWeightedTargetProperties>[]>;
    }
    
    export interface RouteHttpRouteAttributes {
        Action: RouteHttpRouteActionAttributes;
        Match: RouteHttpRouteMatchAttributes;
        RetryPolicy?: RouteHttpRetryPolicyAttributes;
    }
    
    export interface RouteHttpRouteHeaderAttributes {
        Invert?: boolean;
        Match?: RouteHeaderMatchMethodAttributes;
        Name: string;
    }
    
    export interface RouteHttpRouteHeaderProperties {
        Invert?: pulumi.Input<boolean>;
        Match?: pulumi.Input<RouteHeaderMatchMethodProperties>;
        Name: pulumi.Input<string>;
    }
    
    export interface RouteHttpRouteMatchAttributes {
        Headers?: RouteHttpRouteHeaderAttributes[];
        Method?: string;
        Prefix: string;
        Scheme?: string;
    }
    
    export interface RouteHttpRouteMatchProperties {
        Headers?: pulumi.Input<pulumi.Input<RouteHttpRouteHeaderProperties>[]>;
        Method?: pulumi.Input<string>;
        Prefix: pulumi.Input<string>;
        Scheme?: pulumi.Input<string>;
    }
    
    export interface RouteHttpRouteProperties {
        Action: pulumi.Input<RouteHttpRouteActionProperties>;
        Match: pulumi.Input<RouteHttpRouteMatchProperties>;
        RetryPolicy?: pulumi.Input<RouteHttpRetryPolicyProperties>;
    }
    
    export interface RouteMatchRangeAttributes {
        End: number;
        Start: number;
    }
    
    export interface RouteMatchRangeProperties {
        End: pulumi.Input<number>;
        Start: pulumi.Input<number>;
    }
    
    export interface RouteProperties {
        MeshName: pulumi.Input<string>;
        RouteName: pulumi.Input<string>;
        Spec: pulumi.Input<RouteRouteSpecProperties>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VirtualRouterName: pulumi.Input<string>;
    }
    
    export interface RouteRouteSpecAttributes {
        GrpcRoute?: RouteGrpcRouteAttributes;
        Http2Route?: RouteHttpRouteAttributes;
        HttpRoute?: RouteHttpRouteAttributes;
        Priority?: number;
        TcpRoute?: RouteTcpRouteAttributes;
    }
    
    export interface RouteRouteSpecProperties {
        GrpcRoute?: pulumi.Input<RouteGrpcRouteProperties>;
        Http2Route?: pulumi.Input<RouteHttpRouteProperties>;
        HttpRoute?: pulumi.Input<RouteHttpRouteProperties>;
        Priority?: pulumi.Input<number>;
        TcpRoute?: pulumi.Input<RouteTcpRouteProperties>;
    }
    
    export interface RouteTcpRouteActionAttributes {
        WeightedTargets: RouteWeightedTargetAttributes[];
    }
    
    export interface RouteTcpRouteActionProperties {
        WeightedTargets: pulumi.Input<pulumi.Input<RouteWeightedTargetProperties>[]>;
    }
    
    export interface RouteTcpRouteAttributes {
        Action: RouteTcpRouteActionAttributes;
    }
    
    export interface RouteTcpRouteProperties {
        Action: pulumi.Input<RouteTcpRouteActionProperties>;
    }
    
    export interface RouteWeightedTargetAttributes {
        VirtualNode: string;
        Weight: number;
    }
    
    export interface RouteWeightedTargetProperties {
        VirtualNode: pulumi.Input<string>;
        Weight: pulumi.Input<number>;
    }
    
    export interface VirtualNodeAccessLogAttributes {
        File?: VirtualNodeFileAccessLogAttributes;
    }
    
    export interface VirtualNodeAccessLogProperties {
        File?: pulumi.Input<VirtualNodeFileAccessLogProperties>;
    }
    
    export interface VirtualNodeAttributes {
        Arn: string;
        MeshName: string;
        Uid: string;
        VirtualNodeName: string;
    }
    
    export interface VirtualNodeAwsCloudMapInstanceAttributeAttributes {
        Key: string;
        Value: string;
    }
    
    export interface VirtualNodeAwsCloudMapInstanceAttributeProperties {
        Key: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface VirtualNodeAwsCloudMapServiceDiscoveryAttributes {
        Attributes?: VirtualNodeAwsCloudMapInstanceAttributeAttributes[];
        NamespaceName: string;
        ServiceName: string;
    }
    
    export interface VirtualNodeAwsCloudMapServiceDiscoveryProperties {
        Attributes?: pulumi.Input<pulumi.Input<VirtualNodeAwsCloudMapInstanceAttributeProperties>[]>;
        NamespaceName: pulumi.Input<string>;
        ServiceName: pulumi.Input<string>;
    }
    
    export interface VirtualNodeBackendAttributes {
        VirtualService?: VirtualNodeVirtualServiceBackendAttributes;
    }
    
    export interface VirtualNodeBackendProperties {
        VirtualService?: pulumi.Input<VirtualNodeVirtualServiceBackendProperties>;
    }
    
    export interface VirtualNodeDnsServiceDiscoveryAttributes {
        Hostname: string;
    }
    
    export interface VirtualNodeDnsServiceDiscoveryProperties {
        Hostname: pulumi.Input<string>;
    }
    
    export interface VirtualNodeFileAccessLogAttributes {
        Path: string;
    }
    
    export interface VirtualNodeFileAccessLogProperties {
        Path: pulumi.Input<string>;
    }
    
    export interface VirtualNodeHealthCheckAttributes {
        HealthyThreshold: number;
        IntervalMillis: number;
        Path?: string;
        Port?: number;
        Protocol: string;
        TimeoutMillis: number;
        UnhealthyThreshold: number;
    }
    
    export interface VirtualNodeHealthCheckProperties {
        HealthyThreshold: pulumi.Input<number>;
        IntervalMillis: pulumi.Input<number>;
        Path?: pulumi.Input<string>;
        Port?: pulumi.Input<number>;
        Protocol: pulumi.Input<string>;
        TimeoutMillis: pulumi.Input<number>;
        UnhealthyThreshold: pulumi.Input<number>;
    }
    
    export interface VirtualNodeListenerAttributes {
        HealthCheck?: VirtualNodeHealthCheckAttributes;
        PortMapping: VirtualNodePortMappingAttributes;
    }
    
    export interface VirtualNodeListenerProperties {
        HealthCheck?: pulumi.Input<VirtualNodeHealthCheckProperties>;
        PortMapping: pulumi.Input<VirtualNodePortMappingProperties>;
    }
    
    export interface VirtualNodeLoggingAttributes {
        AccessLog?: VirtualNodeAccessLogAttributes;
    }
    
    export interface VirtualNodeLoggingProperties {
        AccessLog?: pulumi.Input<VirtualNodeAccessLogProperties>;
    }
    
    export interface VirtualNodePortMappingAttributes {
        Port: number;
        Protocol: string;
    }
    
    export interface VirtualNodePortMappingProperties {
        Port: pulumi.Input<number>;
        Protocol: pulumi.Input<string>;
    }
    
    export interface VirtualNodeProperties {
        MeshName: pulumi.Input<string>;
        Spec: pulumi.Input<VirtualNodeVirtualNodeSpecProperties>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VirtualNodeName: pulumi.Input<string>;
    }
    
    export interface VirtualNodeServiceDiscoveryAttributes {
        AWSCloudMap?: VirtualNodeAwsCloudMapServiceDiscoveryAttributes;
        DNS?: VirtualNodeDnsServiceDiscoveryAttributes;
    }
    
    export interface VirtualNodeServiceDiscoveryProperties {
        AWSCloudMap?: pulumi.Input<VirtualNodeAwsCloudMapServiceDiscoveryProperties>;
        DNS?: pulumi.Input<VirtualNodeDnsServiceDiscoveryProperties>;
    }
    
    export interface VirtualNodeVirtualNodeSpecAttributes {
        Backends?: VirtualNodeBackendAttributes[];
        Listeners?: VirtualNodeListenerAttributes[];
        Logging?: VirtualNodeLoggingAttributes;
        ServiceDiscovery?: VirtualNodeServiceDiscoveryAttributes;
    }
    
    export interface VirtualNodeVirtualNodeSpecProperties {
        Backends?: pulumi.Input<pulumi.Input<VirtualNodeBackendProperties>[]>;
        Listeners?: pulumi.Input<pulumi.Input<VirtualNodeListenerProperties>[]>;
        Logging?: pulumi.Input<VirtualNodeLoggingProperties>;
        ServiceDiscovery?: pulumi.Input<VirtualNodeServiceDiscoveryProperties>;
    }
    
    export interface VirtualNodeVirtualServiceBackendAttributes {
        VirtualServiceName: string;
    }
    
    export interface VirtualNodeVirtualServiceBackendProperties {
        VirtualServiceName: pulumi.Input<string>;
    }
    
    export interface VirtualRouterAttributes {
        Arn: string;
        MeshName: string;
        Uid: string;
        VirtualRouterName: string;
    }
    
    export interface VirtualRouterPortMappingAttributes {
        Port: number;
        Protocol: string;
    }
    
    export interface VirtualRouterPortMappingProperties {
        Port: pulumi.Input<number>;
        Protocol: pulumi.Input<string>;
    }
    
    export interface VirtualRouterProperties {
        MeshName: pulumi.Input<string>;
        Spec: pulumi.Input<VirtualRouterVirtualRouterSpecProperties>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VirtualRouterName: pulumi.Input<string>;
    }
    
    export interface VirtualRouterVirtualRouterListenerAttributes {
        PortMapping: VirtualRouterPortMappingAttributes;
    }
    
    export interface VirtualRouterVirtualRouterListenerProperties {
        PortMapping: pulumi.Input<VirtualRouterPortMappingProperties>;
    }
    
    export interface VirtualRouterVirtualRouterSpecAttributes {
        Listeners: VirtualRouterVirtualRouterListenerAttributes[];
    }
    
    export interface VirtualRouterVirtualRouterSpecProperties {
        Listeners: pulumi.Input<pulumi.Input<VirtualRouterVirtualRouterListenerProperties>[]>;
    }
    
    export interface VirtualServiceAttributes {
        Arn: string;
        MeshName: string;
        Uid: string;
        VirtualServiceName: string;
    }
    
    export interface VirtualServiceProperties {
        MeshName: pulumi.Input<string>;
        Spec: pulumi.Input<VirtualServiceVirtualServiceSpecProperties>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VirtualServiceName: pulumi.Input<string>;
    }
    
    export interface VirtualServiceVirtualNodeServiceProviderAttributes {
        VirtualNodeName: string;
    }
    
    export interface VirtualServiceVirtualNodeServiceProviderProperties {
        VirtualNodeName: pulumi.Input<string>;
    }
    
    export interface VirtualServiceVirtualRouterServiceProviderAttributes {
        VirtualRouterName: string;
    }
    
    export interface VirtualServiceVirtualRouterServiceProviderProperties {
        VirtualRouterName: pulumi.Input<string>;
    }
    
    export interface VirtualServiceVirtualServiceProviderAttributes {
        VirtualNode?: VirtualServiceVirtualNodeServiceProviderAttributes;
        VirtualRouter?: VirtualServiceVirtualRouterServiceProviderAttributes;
    }
    
    export interface VirtualServiceVirtualServiceProviderProperties {
        VirtualNode?: pulumi.Input<VirtualServiceVirtualNodeServiceProviderProperties>;
        VirtualRouter?: pulumi.Input<VirtualServiceVirtualRouterServiceProviderProperties>;
    }
    
    export interface VirtualServiceVirtualServiceSpecAttributes {
        Provider?: VirtualServiceVirtualServiceProviderAttributes;
    }
    
    export interface VirtualServiceVirtualServiceSpecProperties {
        Provider?: pulumi.Input<VirtualServiceVirtualServiceProviderProperties>;
    }
    
    
    export class Mesh extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<MeshAttributes>;

        constructor(name: string, properties: MeshProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppMesh:Mesh", name, inputs, opts);
        }
    }
    
    export class Route extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RouteAttributes>;

        constructor(name: string, properties: RouteProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppMesh:Route", name, inputs, opts);
        }
    }
    
    export class VirtualNode extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VirtualNodeAttributes>;

        constructor(name: string, properties: VirtualNodeProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppMesh:VirtualNode", name, inputs, opts);
        }
    }
    
    export class VirtualRouter extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VirtualRouterAttributes>;

        constructor(name: string, properties: VirtualRouterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppMesh:VirtualRouter", name, inputs, opts);
        }
    }
    
    export class VirtualService extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VirtualServiceAttributes>;

        constructor(name: string, properties: VirtualServiceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppMesh:VirtualService", name, inputs, opts);
        }
    }
    
}

export namespace appstream {
    
    export interface DirectoryConfigAttributes {
    }
    
    export interface DirectoryConfigProperties {
        DirectoryName: pulumi.Input<string>;
        OrganizationalUnitDistinguishedNames: pulumi.Input<pulumi.Input<string>[]>;
        ServiceAccountCredentials: pulumi.Input<DirectoryConfigServiceAccountCredentialsProperties>;
    }
    
    export interface DirectoryConfigServiceAccountCredentialsAttributes {
        AccountName: string;
        AccountPassword: string;
    }
    
    export interface DirectoryConfigServiceAccountCredentialsProperties {
        AccountName: pulumi.Input<string>;
        AccountPassword: pulumi.Input<string>;
    }
    
    export interface FleetAttributes {
    }
    
    export interface FleetComputeCapacityAttributes {
        DesiredInstances: number;
    }
    
    export interface FleetComputeCapacityProperties {
        DesiredInstances: pulumi.Input<number>;
    }
    
    export interface FleetDomainJoinInfoAttributes {
        DirectoryName?: string;
        OrganizationalUnitDistinguishedName?: string;
    }
    
    export interface FleetDomainJoinInfoProperties {
        DirectoryName?: pulumi.Input<string>;
        OrganizationalUnitDistinguishedName?: pulumi.Input<string>;
    }
    
    export interface FleetProperties {
        ComputeCapacity: pulumi.Input<FleetComputeCapacityProperties>;
        Description?: pulumi.Input<string>;
        DisconnectTimeoutInSeconds?: pulumi.Input<number>;
        DisplayName?: pulumi.Input<string>;
        DomainJoinInfo?: pulumi.Input<FleetDomainJoinInfoProperties>;
        EnableDefaultInternetAccess?: pulumi.Input<boolean>;
        FleetType?: pulumi.Input<string>;
        IdleDisconnectTimeoutInSeconds?: pulumi.Input<number>;
        ImageArn?: pulumi.Input<string>;
        ImageName?: pulumi.Input<string>;
        InstanceType: pulumi.Input<string>;
        MaxUserDurationInSeconds?: pulumi.Input<number>;
        Name?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VpcConfig?: pulumi.Input<FleetVpcConfigProperties>;
    }
    
    export interface FleetVpcConfigAttributes {
        SecurityGroupIds?: string[];
        SubnetIds?: string[];
    }
    
    export interface FleetVpcConfigProperties {
        SecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
        SubnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface ImageBuilderAccessEndpointAttributes {
        EndpointType: string;
        VpceId: string;
    }
    
    export interface ImageBuilderAccessEndpointProperties {
        EndpointType: pulumi.Input<string>;
        VpceId: pulumi.Input<string>;
    }
    
    export interface ImageBuilderAttributes {
        StreamingUrl: string;
    }
    
    export interface ImageBuilderDomainJoinInfoAttributes {
        DirectoryName?: string;
        OrganizationalUnitDistinguishedName?: string;
    }
    
    export interface ImageBuilderDomainJoinInfoProperties {
        DirectoryName?: pulumi.Input<string>;
        OrganizationalUnitDistinguishedName?: pulumi.Input<string>;
    }
    
    export interface ImageBuilderProperties {
        AccessEndpoints?: pulumi.Input<pulumi.Input<ImageBuilderAccessEndpointProperties>[]>;
        AppstreamAgentVersion?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        DisplayName?: pulumi.Input<string>;
        DomainJoinInfo?: pulumi.Input<ImageBuilderDomainJoinInfoProperties>;
        EnableDefaultInternetAccess?: pulumi.Input<boolean>;
        ImageArn?: pulumi.Input<string>;
        ImageName?: pulumi.Input<string>;
        InstanceType: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VpcConfig?: pulumi.Input<ImageBuilderVpcConfigProperties>;
    }
    
    export interface ImageBuilderVpcConfigAttributes {
        SecurityGroupIds?: string[];
        SubnetIds?: string[];
    }
    
    export interface ImageBuilderVpcConfigProperties {
        SecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
        SubnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface StackAccessEndpointAttributes {
        EndpointType: string;
        VpceId: string;
    }
    
    export interface StackAccessEndpointProperties {
        EndpointType: pulumi.Input<string>;
        VpceId: pulumi.Input<string>;
    }
    
    export interface StackApplicationSettingsAttributes {
        Enabled: boolean;
        SettingsGroup?: string;
    }
    
    export interface StackApplicationSettingsProperties {
        Enabled: pulumi.Input<boolean>;
        SettingsGroup?: pulumi.Input<string>;
    }
    
    export interface StackAttributes {
    }
    
    export interface StackFleetAssociationAttributes {
    }
    
    export interface StackFleetAssociationProperties {
        FleetName: pulumi.Input<string>;
        StackName: pulumi.Input<string>;
    }
    
    export interface StackProperties {
        AccessEndpoints?: pulumi.Input<pulumi.Input<StackAccessEndpointProperties>[]>;
        ApplicationSettings?: pulumi.Input<StackApplicationSettingsProperties>;
        AttributesToDelete?: pulumi.Input<pulumi.Input<string>[]>;
        DeleteStorageConnectors?: pulumi.Input<boolean>;
        Description?: pulumi.Input<string>;
        DisplayName?: pulumi.Input<string>;
        EmbedHostDomains?: pulumi.Input<pulumi.Input<string>[]>;
        FeedbackURL?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        RedirectURL?: pulumi.Input<string>;
        StorageConnectors?: pulumi.Input<pulumi.Input<StackStorageConnectorProperties>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        UserSettings?: pulumi.Input<pulumi.Input<StackUserSettingProperties>[]>;
    }
    
    export interface StackStorageConnectorAttributes {
        ConnectorType: string;
        Domains?: string[];
        ResourceIdentifier?: string;
    }
    
    export interface StackStorageConnectorProperties {
        ConnectorType: pulumi.Input<string>;
        Domains?: pulumi.Input<pulumi.Input<string>[]>;
        ResourceIdentifier?: pulumi.Input<string>;
    }
    
    export interface StackUserAssociationAttributes {
    }
    
    export interface StackUserAssociationProperties {
        AuthenticationType: pulumi.Input<string>;
        SendEmailNotification?: pulumi.Input<boolean>;
        StackName: pulumi.Input<string>;
        UserName: pulumi.Input<string>;
    }
    
    export interface StackUserSettingAttributes {
        Action: string;
        Permission: string;
    }
    
    export interface StackUserSettingProperties {
        Action: pulumi.Input<string>;
        Permission: pulumi.Input<string>;
    }
    
    export interface UserAttributes {
    }
    
    export interface UserProperties {
        AuthenticationType: pulumi.Input<string>;
        FirstName?: pulumi.Input<string>;
        LastName?: pulumi.Input<string>;
        MessageAction?: pulumi.Input<string>;
        UserName: pulumi.Input<string>;
    }
    
    
    export class DirectoryConfig extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DirectoryConfigAttributes>;

        constructor(name: string, properties: DirectoryConfigProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppStream:DirectoryConfig", name, inputs, opts);
        }
    }
    
    export class Fleet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<FleetAttributes>;

        constructor(name: string, properties: FleetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppStream:Fleet", name, inputs, opts);
        }
    }
    
    export class ImageBuilder extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ImageBuilderAttributes>;

        constructor(name: string, properties: ImageBuilderProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppStream:ImageBuilder", name, inputs, opts);
        }
    }
    
    export class Stack extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<StackAttributes>;

        constructor(name: string, properties?: StackProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppStream:Stack", name, inputs, opts);
        }
    }
    
    export class StackFleetAssociation extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<StackFleetAssociationAttributes>;

        constructor(name: string, properties: StackFleetAssociationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppStream:StackFleetAssociation", name, inputs, opts);
        }
    }
    
    export class StackUserAssociation extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<StackUserAssociationAttributes>;

        constructor(name: string, properties: StackUserAssociationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppStream:StackUserAssociation", name, inputs, opts);
        }
    }
    
    export class User extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<UserAttributes>;

        constructor(name: string, properties: UserProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppStream:User", name, inputs, opts);
        }
    }
    
}

export namespace appsync {
    
    export interface ApiKeyAttributes {
        ApiKey: string;
        Arn: string;
    }
    
    export interface ApiKeyProperties {
        ApiId: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        Expires?: pulumi.Input<number>;
    }
    
    export interface DataSourceAttributes {
        DataSourceArn: string;
        Name: string;
    }
    
    export interface DataSourceAuthorizationConfigAttributes {
        AuthorizationType: string;
        AwsIamConfig?: DataSourceAwsIamConfigAttributes;
    }
    
    export interface DataSourceAuthorizationConfigProperties {
        AuthorizationType: pulumi.Input<string>;
        AwsIamConfig?: pulumi.Input<DataSourceAwsIamConfigProperties>;
    }
    
    export interface DataSourceAwsIamConfigAttributes {
        SigningRegion?: string;
        SigningServiceName?: string;
    }
    
    export interface DataSourceAwsIamConfigProperties {
        SigningRegion?: pulumi.Input<string>;
        SigningServiceName?: pulumi.Input<string>;
    }
    
    export interface DataSourceDynamoDBConfigAttributes {
        AwsRegion: string;
        TableName: string;
        UseCallerCredentials?: boolean;
    }
    
    export interface DataSourceDynamoDBConfigProperties {
        AwsRegion: pulumi.Input<string>;
        TableName: pulumi.Input<string>;
        UseCallerCredentials?: pulumi.Input<boolean>;
    }
    
    export interface DataSourceElasticsearchConfigAttributes {
        AwsRegion: string;
        Endpoint: string;
    }
    
    export interface DataSourceElasticsearchConfigProperties {
        AwsRegion: pulumi.Input<string>;
        Endpoint: pulumi.Input<string>;
    }
    
    export interface DataSourceHttpConfigAttributes {
        AuthorizationConfig?: DataSourceAuthorizationConfigAttributes;
        Endpoint: string;
    }
    
    export interface DataSourceHttpConfigProperties {
        AuthorizationConfig?: pulumi.Input<DataSourceAuthorizationConfigProperties>;
        Endpoint: pulumi.Input<string>;
    }
    
    export interface DataSourceLambdaConfigAttributes {
        LambdaFunctionArn: string;
    }
    
    export interface DataSourceLambdaConfigProperties {
        LambdaFunctionArn: pulumi.Input<string>;
    }
    
    export interface DataSourceProperties {
        ApiId: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        DynamoDBConfig?: pulumi.Input<DataSourceDynamoDBConfigProperties>;
        ElasticsearchConfig?: pulumi.Input<DataSourceElasticsearchConfigProperties>;
        HttpConfig?: pulumi.Input<DataSourceHttpConfigProperties>;
        LambdaConfig?: pulumi.Input<DataSourceLambdaConfigProperties>;
        Name: pulumi.Input<string>;
        RelationalDatabaseConfig?: pulumi.Input<DataSourceRelationalDatabaseConfigProperties>;
        ServiceRoleArn?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface DataSourceRdsHttpEndpointConfigAttributes {
        AwsRegion: string;
        AwsSecretStoreArn: string;
        DatabaseName?: string;
        DbClusterIdentifier: string;
        Schema?: string;
    }
    
    export interface DataSourceRdsHttpEndpointConfigProperties {
        AwsRegion: pulumi.Input<string>;
        AwsSecretStoreArn: pulumi.Input<string>;
        DatabaseName?: pulumi.Input<string>;
        DbClusterIdentifier: pulumi.Input<string>;
        Schema?: pulumi.Input<string>;
    }
    
    export interface DataSourceRelationalDatabaseConfigAttributes {
        RdsHttpEndpointConfig?: DataSourceRdsHttpEndpointConfigAttributes;
        RelationalDatabaseSourceType: string;
    }
    
    export interface DataSourceRelationalDatabaseConfigProperties {
        RdsHttpEndpointConfig?: pulumi.Input<DataSourceRdsHttpEndpointConfigProperties>;
        RelationalDatabaseSourceType: pulumi.Input<string>;
    }
    
    export interface FunctionConfigurationAttributes {
        DataSourceName: string;
        FunctionArn: string;
        FunctionId: string;
        Name: string;
    }
    
    export interface FunctionConfigurationProperties {
        ApiId: pulumi.Input<string>;
        DataSourceName: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        FunctionVersion: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        RequestMappingTemplate?: pulumi.Input<string>;
        RequestMappingTemplateS3Location?: pulumi.Input<string>;
        ResponseMappingTemplate?: pulumi.Input<string>;
        ResponseMappingTemplateS3Location?: pulumi.Input<string>;
    }
    
    export interface GraphQLApiAdditionalAuthenticationProviderAttributes {
        AuthenticationType: string;
        OpenIDConnectConfig?: GraphQLApiOpenIDConnectConfigAttributes;
        UserPoolConfig?: GraphQLApiCognitoUserPoolConfigAttributes;
    }
    
    export interface GraphQLApiAdditionalAuthenticationProviderProperties {
        AuthenticationType: pulumi.Input<string>;
        OpenIDConnectConfig?: pulumi.Input<GraphQLApiOpenIDConnectConfigProperties>;
        UserPoolConfig?: pulumi.Input<GraphQLApiCognitoUserPoolConfigProperties>;
    }
    
    export interface GraphQLApiAdditionalAuthenticationProvidersAttributes {
    }
    
    export interface GraphQLApiAdditionalAuthenticationProvidersProperties {
    }
    
    export interface GraphQLApiAttributes {
        ApiId: string;
        Arn: string;
        GraphQLUrl: string;
    }
    
    export interface GraphQLApiCognitoUserPoolConfigAttributes {
        AppIdClientRegex?: string;
        AwsRegion?: string;
        UserPoolId?: string;
    }
    
    export interface GraphQLApiCognitoUserPoolConfigProperties {
        AppIdClientRegex?: pulumi.Input<string>;
        AwsRegion?: pulumi.Input<string>;
        UserPoolId?: pulumi.Input<string>;
    }
    
    export interface GraphQLApiLogConfigAttributes {
        CloudWatchLogsRoleArn?: string;
        ExcludeVerboseContent?: boolean;
        FieldLogLevel?: string;
    }
    
    export interface GraphQLApiLogConfigProperties {
        CloudWatchLogsRoleArn?: pulumi.Input<string>;
        ExcludeVerboseContent?: pulumi.Input<boolean>;
        FieldLogLevel?: pulumi.Input<string>;
    }
    
    export interface GraphQLApiOpenIDConnectConfigAttributes {
        AuthTTL?: number;
        ClientId?: string;
        IatTTL?: number;
        Issuer?: string;
    }
    
    export interface GraphQLApiOpenIDConnectConfigProperties {
        AuthTTL?: pulumi.Input<number>;
        ClientId?: pulumi.Input<string>;
        IatTTL?: pulumi.Input<number>;
        Issuer?: pulumi.Input<string>;
    }
    
    export interface GraphQLApiProperties {
        AdditionalAuthenticationProviders?: pulumi.Input<GraphQLApiAdditionalAuthenticationProvidersProperties>;
        AuthenticationType: pulumi.Input<string>;
        LogConfig?: pulumi.Input<GraphQLApiLogConfigProperties>;
        Name: pulumi.Input<string>;
        OpenIDConnectConfig?: pulumi.Input<GraphQLApiOpenIDConnectConfigProperties>;
        Tags?: pulumi.Input<GraphQLApiTagsProperties>;
        UserPoolConfig?: pulumi.Input<GraphQLApiUserPoolConfigProperties>;
    }
    
    export interface GraphQLApiTagsAttributes {
    }
    
    export interface GraphQLApiTagsProperties {
    }
    
    export interface GraphQLApiUserPoolConfigAttributes {
        AppIdClientRegex?: string;
        AwsRegion?: string;
        DefaultAction?: string;
        UserPoolId?: string;
    }
    
    export interface GraphQLApiUserPoolConfigProperties {
        AppIdClientRegex?: pulumi.Input<string>;
        AwsRegion?: pulumi.Input<string>;
        DefaultAction?: pulumi.Input<string>;
        UserPoolId?: pulumi.Input<string>;
    }
    
    export interface GraphQLSchemaAttributes {
    }
    
    export interface GraphQLSchemaProperties {
        ApiId: pulumi.Input<string>;
        Definition?: pulumi.Input<string>;
        DefinitionS3Location?: pulumi.Input<string>;
    }
    
    export interface ResolverAttributes {
        FieldName: string;
        ResolverArn: string;
        TypeName: string;
    }
    
    export interface ResolverPipelineConfigAttributes {
        Functions?: string[];
    }
    
    export interface ResolverPipelineConfigProperties {
        Functions?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface ResolverProperties {
        ApiId: pulumi.Input<string>;
        DataSourceName?: pulumi.Input<string>;
        FieldName: pulumi.Input<string>;
        Kind?: pulumi.Input<string>;
        PipelineConfig?: pulumi.Input<ResolverPipelineConfigProperties>;
        RequestMappingTemplate?: pulumi.Input<string>;
        RequestMappingTemplateS3Location?: pulumi.Input<string>;
        ResponseMappingTemplate?: pulumi.Input<string>;
        ResponseMappingTemplateS3Location?: pulumi.Input<string>;
        TypeName: pulumi.Input<string>;
    }
    
    
    export class ApiKey extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ApiKeyAttributes>;

        constructor(name: string, properties: ApiKeyProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppSync:ApiKey", name, inputs, opts);
        }
    }
    
    export class DataSource extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DataSourceAttributes>;

        constructor(name: string, properties: DataSourceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppSync:DataSource", name, inputs, opts);
        }
    }
    
    export class FunctionConfiguration extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<FunctionConfigurationAttributes>;

        constructor(name: string, properties: FunctionConfigurationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppSync:FunctionConfiguration", name, inputs, opts);
        }
    }
    
    export class GraphQLApi extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<GraphQLApiAttributes>;

        constructor(name: string, properties: GraphQLApiProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppSync:GraphQLApi", name, inputs, opts);
        }
    }
    
    export class GraphQLSchema extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<GraphQLSchemaAttributes>;

        constructor(name: string, properties: GraphQLSchemaProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppSync:GraphQLSchema", name, inputs, opts);
        }
    }
    
    export class Resolver extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ResolverAttributes>;

        constructor(name: string, properties: ResolverProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AppSync:Resolver", name, inputs, opts);
        }
    }
    
}

export namespace applicationautoscaling {
    
    export interface ScalableTargetAttributes {
    }
    
    export interface ScalableTargetProperties {
        MaxCapacity: pulumi.Input<number>;
        MinCapacity: pulumi.Input<number>;
        ResourceId: pulumi.Input<string>;
        RoleARN: pulumi.Input<string>;
        ScalableDimension: pulumi.Input<string>;
        ScheduledActions?: pulumi.Input<pulumi.Input<ScalableTargetScheduledActionProperties>[]>;
        ServiceNamespace: pulumi.Input<string>;
        SuspendedState?: pulumi.Input<ScalableTargetSuspendedStateProperties>;
    }
    
    export interface ScalableTargetScalableTargetActionAttributes {
        MaxCapacity?: number;
        MinCapacity?: number;
    }
    
    export interface ScalableTargetScalableTargetActionProperties {
        MaxCapacity?: pulumi.Input<number>;
        MinCapacity?: pulumi.Input<number>;
    }
    
    export interface ScalableTargetScheduledActionAttributes {
        EndTime?: string;
        ScalableTargetAction?: ScalableTargetScalableTargetActionAttributes;
        Schedule: string;
        ScheduledActionName: string;
        StartTime?: string;
    }
    
    export interface ScalableTargetScheduledActionProperties {
        EndTime?: pulumi.Input<string>;
        ScalableTargetAction?: pulumi.Input<ScalableTargetScalableTargetActionProperties>;
        Schedule: pulumi.Input<string>;
        ScheduledActionName: pulumi.Input<string>;
        StartTime?: pulumi.Input<string>;
    }
    
    export interface ScalableTargetSuspendedStateAttributes {
        DynamicScalingInSuspended?: boolean;
        DynamicScalingOutSuspended?: boolean;
        ScheduledScalingSuspended?: boolean;
    }
    
    export interface ScalableTargetSuspendedStateProperties {
        DynamicScalingInSuspended?: pulumi.Input<boolean>;
        DynamicScalingOutSuspended?: pulumi.Input<boolean>;
        ScheduledScalingSuspended?: pulumi.Input<boolean>;
    }
    
    export interface ScalingPolicyAttributes {
    }
    
    export interface ScalingPolicyCustomizedMetricSpecificationAttributes {
        Dimensions?: ScalingPolicyMetricDimensionAttributes[];
        MetricName: string;
        Namespace: string;
        Statistic: string;
        Unit?: string;
    }
    
    export interface ScalingPolicyCustomizedMetricSpecificationProperties {
        Dimensions?: pulumi.Input<pulumi.Input<ScalingPolicyMetricDimensionProperties>[]>;
        MetricName: pulumi.Input<string>;
        Namespace: pulumi.Input<string>;
        Statistic: pulumi.Input<string>;
        Unit?: pulumi.Input<string>;
    }
    
    export interface ScalingPolicyMetricDimensionAttributes {
        Name: string;
        Value: string;
    }
    
    export interface ScalingPolicyMetricDimensionProperties {
        Name: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface ScalingPolicyPredefinedMetricSpecificationAttributes {
        PredefinedMetricType: string;
        ResourceLabel?: string;
    }
    
    export interface ScalingPolicyPredefinedMetricSpecificationProperties {
        PredefinedMetricType: pulumi.Input<string>;
        ResourceLabel?: pulumi.Input<string>;
    }
    
    export interface ScalingPolicyProperties {
        PolicyName: pulumi.Input<string>;
        PolicyType: pulumi.Input<string>;
        ResourceId?: pulumi.Input<string>;
        ScalableDimension?: pulumi.Input<string>;
        ScalingTargetId?: pulumi.Input<string>;
        ServiceNamespace?: pulumi.Input<string>;
        StepScalingPolicyConfiguration?: pulumi.Input<ScalingPolicyStepScalingPolicyConfigurationProperties>;
        TargetTrackingScalingPolicyConfiguration?: pulumi.Input<ScalingPolicyTargetTrackingScalingPolicyConfigurationProperties>;
    }
    
    export interface ScalingPolicyStepAdjustmentAttributes {
        MetricIntervalLowerBound?: number;
        MetricIntervalUpperBound?: number;
        ScalingAdjustment: number;
    }
    
    export interface ScalingPolicyStepAdjustmentProperties {
        MetricIntervalLowerBound?: pulumi.Input<number>;
        MetricIntervalUpperBound?: pulumi.Input<number>;
        ScalingAdjustment: pulumi.Input<number>;
    }
    
    export interface ScalingPolicyStepScalingPolicyConfigurationAttributes {
        AdjustmentType?: string;
        Cooldown?: number;
        MetricAggregationType?: string;
        MinAdjustmentMagnitude?: number;
        StepAdjustments?: ScalingPolicyStepAdjustmentAttributes[];
    }
    
    export interface ScalingPolicyStepScalingPolicyConfigurationProperties {
        AdjustmentType?: pulumi.Input<string>;
        Cooldown?: pulumi.Input<number>;
        MetricAggregationType?: pulumi.Input<string>;
        MinAdjustmentMagnitude?: pulumi.Input<number>;
        StepAdjustments?: pulumi.Input<pulumi.Input<ScalingPolicyStepAdjustmentProperties>[]>;
    }
    
    export interface ScalingPolicyTargetTrackingScalingPolicyConfigurationAttributes {
        CustomizedMetricSpecification?: ScalingPolicyCustomizedMetricSpecificationAttributes;
        DisableScaleIn?: boolean;
        PredefinedMetricSpecification?: ScalingPolicyPredefinedMetricSpecificationAttributes;
        ScaleInCooldown?: number;
        ScaleOutCooldown?: number;
        TargetValue: number;
    }
    
    export interface ScalingPolicyTargetTrackingScalingPolicyConfigurationProperties {
        CustomizedMetricSpecification?: pulumi.Input<ScalingPolicyCustomizedMetricSpecificationProperties>;
        DisableScaleIn?: pulumi.Input<boolean>;
        PredefinedMetricSpecification?: pulumi.Input<ScalingPolicyPredefinedMetricSpecificationProperties>;
        ScaleInCooldown?: pulumi.Input<number>;
        ScaleOutCooldown?: pulumi.Input<number>;
        TargetValue: pulumi.Input<number>;
    }
    
    
    export class ScalableTarget extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ScalableTargetAttributes>;

        constructor(name: string, properties: ScalableTargetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApplicationAutoScaling:ScalableTarget", name, inputs, opts);
        }
    }
    
    export class ScalingPolicy extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ScalingPolicyAttributes>;

        constructor(name: string, properties: ScalingPolicyProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ApplicationAutoScaling:ScalingPolicy", name, inputs, opts);
        }
    }
    
}

export namespace athena {
    
    export interface NamedQueryAttributes {
    }
    
    export interface NamedQueryProperties {
        Database: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        QueryString: pulumi.Input<string>;
    }
    
    
    export class NamedQuery extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<NamedQueryAttributes>;

        constructor(name: string, properties: NamedQueryProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Athena:NamedQuery", name, inputs, opts);
        }
    }
    
}

export namespace autoscaling {
    
    export interface AutoScalingGroupAttributes {
    }
    
    export interface AutoScalingGroupInstancesDistributionAttributes {
        OnDemandAllocationStrategy?: string;
        OnDemandBaseCapacity?: number;
        OnDemandPercentageAboveBaseCapacity?: number;
        SpotAllocationStrategy?: string;
        SpotInstancePools?: number;
        SpotMaxPrice?: string;
    }
    
    export interface AutoScalingGroupInstancesDistributionProperties {
        OnDemandAllocationStrategy?: pulumi.Input<string>;
        OnDemandBaseCapacity?: pulumi.Input<number>;
        OnDemandPercentageAboveBaseCapacity?: pulumi.Input<number>;
        SpotAllocationStrategy?: pulumi.Input<string>;
        SpotInstancePools?: pulumi.Input<number>;
        SpotMaxPrice?: pulumi.Input<string>;
    }
    
    export interface AutoScalingGroupLaunchTemplateAttributes {
        LaunchTemplateSpecification: AutoScalingGroupLaunchTemplateSpecificationAttributes;
        Overrides?: AutoScalingGroupLaunchTemplateOverridesAttributes[];
    }
    
    export interface AutoScalingGroupLaunchTemplateOverridesAttributes {
        InstanceType?: string;
    }
    
    export interface AutoScalingGroupLaunchTemplateOverridesProperties {
        InstanceType?: pulumi.Input<string>;
    }
    
    export interface AutoScalingGroupLaunchTemplateProperties {
        LaunchTemplateSpecification: pulumi.Input<AutoScalingGroupLaunchTemplateSpecificationProperties>;
        Overrides?: pulumi.Input<pulumi.Input<AutoScalingGroupLaunchTemplateOverridesProperties>[]>;
    }
    
    export interface AutoScalingGroupLaunchTemplateSpecificationAttributes {
        LaunchTemplateId?: string;
        LaunchTemplateName?: string;
        Version: string;
    }
    
    export interface AutoScalingGroupLaunchTemplateSpecificationProperties {
        LaunchTemplateId?: pulumi.Input<string>;
        LaunchTemplateName?: pulumi.Input<string>;
        Version: pulumi.Input<string>;
    }
    
    export interface AutoScalingGroupLifecycleHookSpecificationAttributes {
        DefaultResult?: string;
        HeartbeatTimeout?: number;
        LifecycleHookName: string;
        LifecycleTransition: string;
        NotificationMetadata?: string;
        NotificationTargetARN?: string;
        RoleARN?: string;
    }
    
    export interface AutoScalingGroupLifecycleHookSpecificationProperties {
        DefaultResult?: pulumi.Input<string>;
        HeartbeatTimeout?: pulumi.Input<number>;
        LifecycleHookName: pulumi.Input<string>;
        LifecycleTransition: pulumi.Input<string>;
        NotificationMetadata?: pulumi.Input<string>;
        NotificationTargetARN?: pulumi.Input<string>;
        RoleARN?: pulumi.Input<string>;
    }
    
    export interface AutoScalingGroupMetricsCollectionAttributes {
        Granularity: string;
        Metrics?: string[];
    }
    
    export interface AutoScalingGroupMetricsCollectionProperties {
        Granularity: pulumi.Input<string>;
        Metrics?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface AutoScalingGroupMixedInstancesPolicyAttributes {
        InstancesDistribution?: AutoScalingGroupInstancesDistributionAttributes;
        LaunchTemplate: AutoScalingGroupLaunchTemplateAttributes;
    }
    
    export interface AutoScalingGroupMixedInstancesPolicyProperties {
        InstancesDistribution?: pulumi.Input<AutoScalingGroupInstancesDistributionProperties>;
        LaunchTemplate: pulumi.Input<AutoScalingGroupLaunchTemplateProperties>;
    }
    
    export interface AutoScalingGroupNotificationConfigurationAttributes {
        NotificationTypes?: string[];
        TopicARN: string;
    }
    
    export interface AutoScalingGroupNotificationConfigurationProperties {
        NotificationTypes?: pulumi.Input<pulumi.Input<string>[]>;
        TopicARN: pulumi.Input<string>;
    }
    
    export interface AutoScalingGroupProperties {
        AutoScalingGroupName?: pulumi.Input<string>;
        AvailabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
        Cooldown?: pulumi.Input<string>;
        DesiredCapacity?: pulumi.Input<string>;
        HealthCheckGracePeriod?: pulumi.Input<number>;
        HealthCheckType?: pulumi.Input<string>;
        InstanceId?: pulumi.Input<string>;
        LaunchConfigurationName?: pulumi.Input<string>;
        LaunchTemplate?: pulumi.Input<AutoScalingGroupLaunchTemplateSpecificationProperties>;
        LifecycleHookSpecificationList?: pulumi.Input<pulumi.Input<AutoScalingGroupLifecycleHookSpecificationProperties>[]>;
        LoadBalancerNames?: pulumi.Input<pulumi.Input<string>[]>;
        MaxSize: pulumi.Input<string>;
        MetricsCollection?: pulumi.Input<pulumi.Input<AutoScalingGroupMetricsCollectionProperties>[]>;
        MinSize: pulumi.Input<string>;
        MixedInstancesPolicy?: pulumi.Input<AutoScalingGroupMixedInstancesPolicyProperties>;
        NotificationConfigurations?: pulumi.Input<pulumi.Input<AutoScalingGroupNotificationConfigurationProperties>[]>;
        PlacementGroup?: pulumi.Input<string>;
        ServiceLinkedRoleARN?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<AutoScalingGroupTagPropertyProperties>[]>;
        TargetGroupARNs?: pulumi.Input<pulumi.Input<string>[]>;
        TerminationPolicies?: pulumi.Input<pulumi.Input<string>[]>;
        VPCZoneIdentifier?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface AutoScalingGroupTagPropertyAttributes {
        Key: string;
        PropagateAtLaunch: boolean;
        Value: string;
    }
    
    export interface AutoScalingGroupTagPropertyProperties {
        Key: pulumi.Input<string>;
        PropagateAtLaunch: pulumi.Input<boolean>;
        Value: pulumi.Input<string>;
    }
    
    export interface LaunchConfigurationAttributes {
    }
    
    export interface LaunchConfigurationBlockDeviceAttributes {
        DeleteOnTermination?: boolean;
        Encrypted?: boolean;
        Iops?: number;
        SnapshotId?: string;
        VolumeSize?: number;
        VolumeType?: string;
    }
    
    export interface LaunchConfigurationBlockDeviceMappingAttributes {
        DeviceName: string;
        Ebs?: LaunchConfigurationBlockDeviceAttributes;
        NoDevice?: boolean;
        VirtualName?: string;
    }
    
    export interface LaunchConfigurationBlockDeviceMappingProperties {
        DeviceName: pulumi.Input<string>;
        Ebs?: pulumi.Input<LaunchConfigurationBlockDeviceProperties>;
        NoDevice?: pulumi.Input<boolean>;
        VirtualName?: pulumi.Input<string>;
    }
    
    export interface LaunchConfigurationBlockDeviceProperties {
        DeleteOnTermination?: pulumi.Input<boolean>;
        Encrypted?: pulumi.Input<boolean>;
        Iops?: pulumi.Input<number>;
        SnapshotId?: pulumi.Input<string>;
        VolumeSize?: pulumi.Input<number>;
        VolumeType?: pulumi.Input<string>;
    }
    
    export interface LaunchConfigurationProperties {
        AssociatePublicIpAddress?: pulumi.Input<boolean>;
        BlockDeviceMappings?: pulumi.Input<pulumi.Input<LaunchConfigurationBlockDeviceMappingProperties>[]>;
        ClassicLinkVPCId?: pulumi.Input<string>;
        ClassicLinkVPCSecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
        EbsOptimized?: pulumi.Input<boolean>;
        IamInstanceProfile?: pulumi.Input<string>;
        ImageId: pulumi.Input<string>;
        InstanceId?: pulumi.Input<string>;
        InstanceMonitoring?: pulumi.Input<boolean>;
        InstanceType: pulumi.Input<string>;
        KernelId?: pulumi.Input<string>;
        KeyName?: pulumi.Input<string>;
        LaunchConfigurationName?: pulumi.Input<string>;
        PlacementTenancy?: pulumi.Input<string>;
        RamDiskId?: pulumi.Input<string>;
        SecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
        SpotPrice?: pulumi.Input<string>;
        UserData?: pulumi.Input<string>;
    }
    
    export interface LifecycleHookAttributes {
    }
    
    export interface LifecycleHookProperties {
        AutoScalingGroupName: pulumi.Input<string>;
        DefaultResult?: pulumi.Input<string>;
        HeartbeatTimeout?: pulumi.Input<number>;
        LifecycleHookName?: pulumi.Input<string>;
        LifecycleTransition: pulumi.Input<string>;
        NotificationMetadata?: pulumi.Input<string>;
        NotificationTargetARN?: pulumi.Input<string>;
        RoleARN?: pulumi.Input<string>;
    }
    
    export interface ScalingPolicyAttributes {
    }
    
    export interface ScalingPolicyCustomizedMetricSpecificationAttributes {
        Dimensions?: ScalingPolicyMetricDimensionAttributes[];
        MetricName: string;
        Namespace: string;
        Statistic: string;
        Unit?: string;
    }
    
    export interface ScalingPolicyCustomizedMetricSpecificationProperties {
        Dimensions?: pulumi.Input<pulumi.Input<ScalingPolicyMetricDimensionProperties>[]>;
        MetricName: pulumi.Input<string>;
        Namespace: pulumi.Input<string>;
        Statistic: pulumi.Input<string>;
        Unit?: pulumi.Input<string>;
    }
    
    export interface ScalingPolicyMetricDimensionAttributes {
        Name: string;
        Value: string;
    }
    
    export interface ScalingPolicyMetricDimensionProperties {
        Name: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface ScalingPolicyPredefinedMetricSpecificationAttributes {
        PredefinedMetricType: string;
        ResourceLabel?: string;
    }
    
    export interface ScalingPolicyPredefinedMetricSpecificationProperties {
        PredefinedMetricType: pulumi.Input<string>;
        ResourceLabel?: pulumi.Input<string>;
    }
    
    export interface ScalingPolicyProperties {
        AdjustmentType?: pulumi.Input<string>;
        AutoScalingGroupName: pulumi.Input<string>;
        Cooldown?: pulumi.Input<string>;
        EstimatedInstanceWarmup?: pulumi.Input<number>;
        MetricAggregationType?: pulumi.Input<string>;
        MinAdjustmentMagnitude?: pulumi.Input<number>;
        PolicyType?: pulumi.Input<string>;
        ScalingAdjustment?: pulumi.Input<number>;
        StepAdjustments?: pulumi.Input<pulumi.Input<ScalingPolicyStepAdjustmentProperties>[]>;
        TargetTrackingConfiguration?: pulumi.Input<ScalingPolicyTargetTrackingConfigurationProperties>;
    }
    
    export interface ScalingPolicyStepAdjustmentAttributes {
        MetricIntervalLowerBound?: number;
        MetricIntervalUpperBound?: number;
        ScalingAdjustment: number;
    }
    
    export interface ScalingPolicyStepAdjustmentProperties {
        MetricIntervalLowerBound?: pulumi.Input<number>;
        MetricIntervalUpperBound?: pulumi.Input<number>;
        ScalingAdjustment: pulumi.Input<number>;
    }
    
    export interface ScalingPolicyTargetTrackingConfigurationAttributes {
        CustomizedMetricSpecification?: ScalingPolicyCustomizedMetricSpecificationAttributes;
        DisableScaleIn?: boolean;
        PredefinedMetricSpecification?: ScalingPolicyPredefinedMetricSpecificationAttributes;
        TargetValue: number;
    }
    
    export interface ScalingPolicyTargetTrackingConfigurationProperties {
        CustomizedMetricSpecification?: pulumi.Input<ScalingPolicyCustomizedMetricSpecificationProperties>;
        DisableScaleIn?: pulumi.Input<boolean>;
        PredefinedMetricSpecification?: pulumi.Input<ScalingPolicyPredefinedMetricSpecificationProperties>;
        TargetValue: pulumi.Input<number>;
    }
    
    export interface ScheduledActionAttributes {
    }
    
    export interface ScheduledActionProperties {
        AutoScalingGroupName: pulumi.Input<string>;
        DesiredCapacity?: pulumi.Input<number>;
        EndTime?: pulumi.Input<string>;
        MaxSize?: pulumi.Input<number>;
        MinSize?: pulumi.Input<number>;
        Recurrence?: pulumi.Input<string>;
        StartTime?: pulumi.Input<string>;
    }
    
    
    export class AutoScalingGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AutoScalingGroupAttributes>;

        constructor(name: string, properties: AutoScalingGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AutoScaling:AutoScalingGroup", name, inputs, opts);
        }
    }
    
    export class LaunchConfiguration extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<LaunchConfigurationAttributes>;

        constructor(name: string, properties: LaunchConfigurationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AutoScaling:LaunchConfiguration", name, inputs, opts);
        }
    }
    
    export class LifecycleHook extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<LifecycleHookAttributes>;

        constructor(name: string, properties: LifecycleHookProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AutoScaling:LifecycleHook", name, inputs, opts);
        }
    }
    
    export class ScalingPolicy extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ScalingPolicyAttributes>;

        constructor(name: string, properties: ScalingPolicyProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AutoScaling:ScalingPolicy", name, inputs, opts);
        }
    }
    
    export class ScheduledAction extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ScheduledActionAttributes>;

        constructor(name: string, properties: ScheduledActionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AutoScaling:ScheduledAction", name, inputs, opts);
        }
    }
    
}

export namespace autoscalingplans {
    
    export interface ScalingPlanApplicationSourceAttributes {
        CloudFormationStackARN?: string;
        TagFilters?: ScalingPlanTagFilterAttributes[];
    }
    
    export interface ScalingPlanApplicationSourceProperties {
        CloudFormationStackARN?: pulumi.Input<string>;
        TagFilters?: pulumi.Input<pulumi.Input<ScalingPlanTagFilterProperties>[]>;
    }
    
    export interface ScalingPlanAttributes {
        ScalingPlanName: string;
        ScalingPlanVersion: string;
    }
    
    export interface ScalingPlanCustomizedLoadMetricSpecificationAttributes {
        Dimensions?: ScalingPlanMetricDimensionAttributes[];
        MetricName: string;
        Namespace: string;
        Statistic: string;
        Unit?: string;
    }
    
    export interface ScalingPlanCustomizedLoadMetricSpecificationProperties {
        Dimensions?: pulumi.Input<pulumi.Input<ScalingPlanMetricDimensionProperties>[]>;
        MetricName: pulumi.Input<string>;
        Namespace: pulumi.Input<string>;
        Statistic: pulumi.Input<string>;
        Unit?: pulumi.Input<string>;
    }
    
    export interface ScalingPlanCustomizedScalingMetricSpecificationAttributes {
        Dimensions?: ScalingPlanMetricDimensionAttributes[];
        MetricName: string;
        Namespace: string;
        Statistic: string;
        Unit?: string;
    }
    
    export interface ScalingPlanCustomizedScalingMetricSpecificationProperties {
        Dimensions?: pulumi.Input<pulumi.Input<ScalingPlanMetricDimensionProperties>[]>;
        MetricName: pulumi.Input<string>;
        Namespace: pulumi.Input<string>;
        Statistic: pulumi.Input<string>;
        Unit?: pulumi.Input<string>;
    }
    
    export interface ScalingPlanMetricDimensionAttributes {
        Name: string;
        Value: string;
    }
    
    export interface ScalingPlanMetricDimensionProperties {
        Name: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface ScalingPlanPredefinedLoadMetricSpecificationAttributes {
        PredefinedLoadMetricType: string;
        ResourceLabel?: string;
    }
    
    export interface ScalingPlanPredefinedLoadMetricSpecificationProperties {
        PredefinedLoadMetricType: pulumi.Input<string>;
        ResourceLabel?: pulumi.Input<string>;
    }
    
    export interface ScalingPlanPredefinedScalingMetricSpecificationAttributes {
        PredefinedScalingMetricType: string;
        ResourceLabel?: string;
    }
    
    export interface ScalingPlanPredefinedScalingMetricSpecificationProperties {
        PredefinedScalingMetricType: pulumi.Input<string>;
        ResourceLabel?: pulumi.Input<string>;
    }
    
    export interface ScalingPlanProperties {
        ApplicationSource: pulumi.Input<ScalingPlanApplicationSourceProperties>;
        ScalingInstructions: pulumi.Input<pulumi.Input<ScalingPlanScalingInstructionProperties>[]>;
    }
    
    export interface ScalingPlanScalingInstructionAttributes {
        CustomizedLoadMetricSpecification?: ScalingPlanCustomizedLoadMetricSpecificationAttributes;
        DisableDynamicScaling?: boolean;
        MaxCapacity: number;
        MinCapacity: number;
        PredefinedLoadMetricSpecification?: ScalingPlanPredefinedLoadMetricSpecificationAttributes;
        PredictiveScalingMaxCapacityBehavior?: string;
        PredictiveScalingMaxCapacityBuffer?: number;
        PredictiveScalingMode?: string;
        ResourceId: string;
        ScalableDimension: string;
        ScalingPolicyUpdateBehavior?: string;
        ScheduledActionBufferTime?: number;
        ServiceNamespace: string;
        TargetTrackingConfigurations: ScalingPlanTargetTrackingConfigurationAttributes[];
    }
    
    export interface ScalingPlanScalingInstructionProperties {
        CustomizedLoadMetricSpecification?: pulumi.Input<ScalingPlanCustomizedLoadMetricSpecificationProperties>;
        DisableDynamicScaling?: pulumi.Input<boolean>;
        MaxCapacity: pulumi.Input<number>;
        MinCapacity: pulumi.Input<number>;
        PredefinedLoadMetricSpecification?: pulumi.Input<ScalingPlanPredefinedLoadMetricSpecificationProperties>;
        PredictiveScalingMaxCapacityBehavior?: pulumi.Input<string>;
        PredictiveScalingMaxCapacityBuffer?: pulumi.Input<number>;
        PredictiveScalingMode?: pulumi.Input<string>;
        ResourceId: pulumi.Input<string>;
        ScalableDimension: pulumi.Input<string>;
        ScalingPolicyUpdateBehavior?: pulumi.Input<string>;
        ScheduledActionBufferTime?: pulumi.Input<number>;
        ServiceNamespace: pulumi.Input<string>;
        TargetTrackingConfigurations: pulumi.Input<pulumi.Input<ScalingPlanTargetTrackingConfigurationProperties>[]>;
    }
    
    export interface ScalingPlanTagFilterAttributes {
        Key: string;
        Values?: string[];
    }
    
    export interface ScalingPlanTagFilterProperties {
        Key: pulumi.Input<string>;
        Values?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface ScalingPlanTargetTrackingConfigurationAttributes {
        CustomizedScalingMetricSpecification?: ScalingPlanCustomizedScalingMetricSpecificationAttributes;
        DisableScaleIn?: boolean;
        EstimatedInstanceWarmup?: number;
        PredefinedScalingMetricSpecification?: ScalingPlanPredefinedScalingMetricSpecificationAttributes;
        ScaleInCooldown?: number;
        ScaleOutCooldown?: number;
        TargetValue: number;
    }
    
    export interface ScalingPlanTargetTrackingConfigurationProperties {
        CustomizedScalingMetricSpecification?: pulumi.Input<ScalingPlanCustomizedScalingMetricSpecificationProperties>;
        DisableScaleIn?: pulumi.Input<boolean>;
        EstimatedInstanceWarmup?: pulumi.Input<number>;
        PredefinedScalingMetricSpecification?: pulumi.Input<ScalingPlanPredefinedScalingMetricSpecificationProperties>;
        ScaleInCooldown?: pulumi.Input<number>;
        ScaleOutCooldown?: pulumi.Input<number>;
        TargetValue: pulumi.Input<number>;
    }
    
    
    export class ScalingPlan extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ScalingPlanAttributes>;

        constructor(name: string, properties: ScalingPlanProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:AutoScalingPlans:ScalingPlan", name, inputs, opts);
        }
    }
    
}

export namespace backup {
    
    export interface BackupPlanAttributes {
        BackupPlanArn: string;
        BackupPlanId: string;
        VersionId: string;
    }
    
    export interface BackupPlanBackupPlanResourceTypeAttributes {
        BackupPlanName: string;
        BackupPlanRule: BackupPlanBackupRuleResourceTypeAttributes[];
    }
    
    export interface BackupPlanBackupPlanResourceTypeProperties {
        BackupPlanName: pulumi.Input<string>;
        BackupPlanRule: pulumi.Input<pulumi.Input<BackupPlanBackupRuleResourceTypeProperties>[]>;
    }
    
    export interface BackupPlanBackupRuleResourceTypeAttributes {
        CompletionWindowMinutes?: number;
        Lifecycle?: BackupPlanLifecycleResourceTypeAttributes;
        RecoveryPointTags?: string;
        RuleName: string;
        ScheduleExpression?: string;
        StartWindowMinutes?: number;
        TargetBackupVault: string;
    }
    
    export interface BackupPlanBackupRuleResourceTypeProperties {
        CompletionWindowMinutes?: pulumi.Input<number>;
        Lifecycle?: pulumi.Input<BackupPlanLifecycleResourceTypeProperties>;
        RecoveryPointTags?: pulumi.Input<string>;
        RuleName: pulumi.Input<string>;
        ScheduleExpression?: pulumi.Input<string>;
        StartWindowMinutes?: pulumi.Input<number>;
        TargetBackupVault: pulumi.Input<string>;
    }
    
    export interface BackupPlanLifecycleResourceTypeAttributes {
        DeleteAfterDays?: number;
        MoveToColdStorageAfterDays?: number;
    }
    
    export interface BackupPlanLifecycleResourceTypeProperties {
        DeleteAfterDays?: pulumi.Input<number>;
        MoveToColdStorageAfterDays?: pulumi.Input<number>;
    }
    
    export interface BackupPlanProperties {
        BackupPlan: pulumi.Input<BackupPlanBackupPlanResourceTypeProperties>;
        BackupPlanTags?: pulumi.Input<string>;
    }
    
    export interface BackupSelectionAttributes {
        BackupPlanId: string;
        SelectionId: string;
    }
    
    export interface BackupSelectionBackupSelectionResourceTypeAttributes {
        IamRoleArn: string;
        ListOfTags?: BackupSelectionConditionResourceTypeAttributes[];
        Resources?: string[];
        SelectionName: string;
    }
    
    export interface BackupSelectionBackupSelectionResourceTypeProperties {
        IamRoleArn: pulumi.Input<string>;
        ListOfTags?: pulumi.Input<pulumi.Input<BackupSelectionConditionResourceTypeProperties>[]>;
        Resources?: pulumi.Input<pulumi.Input<string>[]>;
        SelectionName: pulumi.Input<string>;
    }
    
    export interface BackupSelectionConditionResourceTypeAttributes {
        ConditionKey: string;
        ConditionType: string;
        ConditionValue: string;
    }
    
    export interface BackupSelectionConditionResourceTypeProperties {
        ConditionKey: pulumi.Input<string>;
        ConditionType: pulumi.Input<string>;
        ConditionValue: pulumi.Input<string>;
    }
    
    export interface BackupSelectionProperties {
        BackupPlanId: pulumi.Input<string>;
        BackupSelection: pulumi.Input<BackupSelectionBackupSelectionResourceTypeProperties>;
    }
    
    export interface BackupVaultAttributes {
        BackupVaultArn: string;
        BackupVaultName: string;
    }
    
    export interface BackupVaultNotificationObjectTypeAttributes {
        BackupVaultEvents: string[];
        SNSTopicArn: string;
    }
    
    export interface BackupVaultNotificationObjectTypeProperties {
        BackupVaultEvents: pulumi.Input<pulumi.Input<string>[]>;
        SNSTopicArn: pulumi.Input<string>;
    }
    
    export interface BackupVaultProperties {
        AccessPolicy?: pulumi.Input<string>;
        BackupVaultName: pulumi.Input<string>;
        BackupVaultTags?: pulumi.Input<string>;
        EncryptionKeyArn?: pulumi.Input<string>;
        Notifications?: pulumi.Input<BackupVaultNotificationObjectTypeProperties>;
    }
    
    
    export class BackupPlan extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<BackupPlanAttributes>;

        constructor(name: string, properties: BackupPlanProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Backup:BackupPlan", name, inputs, opts);
        }
    }
    
    export class BackupSelection extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<BackupSelectionAttributes>;

        constructor(name: string, properties: BackupSelectionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Backup:BackupSelection", name, inputs, opts);
        }
    }
    
    export class BackupVault extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<BackupVaultAttributes>;

        constructor(name: string, properties: BackupVaultProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Backup:BackupVault", name, inputs, opts);
        }
    }
    
}

export namespace batch {
    
    export interface ComputeEnvironmentAttributes {
    }
    
    export interface ComputeEnvironmentComputeResourcesAttributes {
        AllocationStrategy?: string;
        BidPercentage?: number;
        DesiredvCpus?: number;
        Ec2KeyPair?: string;
        ImageId?: string;
        InstanceRole: string;
        InstanceTypes: string[];
        LaunchTemplate?: ComputeEnvironmentLaunchTemplateSpecificationAttributes;
        MaxvCpus: number;
        MinvCpus: number;
        PlacementGroup?: string;
        SecurityGroupIds?: string[];
        SpotIamFleetRole?: string;
        Subnets: string[];
        Tags?: string;
        Type: string;
    }
    
    export interface ComputeEnvironmentComputeResourcesProperties {
        AllocationStrategy?: pulumi.Input<string>;
        BidPercentage?: pulumi.Input<number>;
        DesiredvCpus?: pulumi.Input<number>;
        Ec2KeyPair?: pulumi.Input<string>;
        ImageId?: pulumi.Input<string>;
        InstanceRole: pulumi.Input<string>;
        InstanceTypes: pulumi.Input<pulumi.Input<string>[]>;
        LaunchTemplate?: pulumi.Input<ComputeEnvironmentLaunchTemplateSpecificationProperties>;
        MaxvCpus: pulumi.Input<number>;
        MinvCpus: pulumi.Input<number>;
        PlacementGroup?: pulumi.Input<string>;
        SecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
        SpotIamFleetRole?: pulumi.Input<string>;
        Subnets: pulumi.Input<pulumi.Input<string>[]>;
        Tags?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface ComputeEnvironmentLaunchTemplateSpecificationAttributes {
        LaunchTemplateId?: string;
        LaunchTemplateName?: string;
        Version?: string;
    }
    
    export interface ComputeEnvironmentLaunchTemplateSpecificationProperties {
        LaunchTemplateId?: pulumi.Input<string>;
        LaunchTemplateName?: pulumi.Input<string>;
        Version?: pulumi.Input<string>;
    }
    
    export interface ComputeEnvironmentProperties {
        ComputeEnvironmentName?: pulumi.Input<string>;
        ComputeResources?: pulumi.Input<ComputeEnvironmentComputeResourcesProperties>;
        ServiceRole: pulumi.Input<string>;
        State?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface JobDefinitionAttributes {
    }
    
    export interface JobDefinitionContainerPropertiesAttributes {
        Command?: string[];
        Environment?: JobDefinitionEnvironmentAttributes[];
        Image: string;
        InstanceType?: string;
        JobRoleArn?: string;
        LinuxParameters?: JobDefinitionLinuxParametersAttributes;
        Memory: number;
        MountPoints?: JobDefinitionMountPointsAttributes[];
        Privileged?: boolean;
        ReadonlyRootFilesystem?: boolean;
        ResourceRequirements?: JobDefinitionResourceRequirementAttributes[];
        Ulimits?: JobDefinitionUlimitAttributes[];
        User?: string;
        Vcpus: number;
        Volumes?: JobDefinitionVolumesAttributes[];
    }
    
    export interface JobDefinitionContainerPropertiesProperties {
        Command?: pulumi.Input<pulumi.Input<string>[]>;
        Environment?: pulumi.Input<pulumi.Input<JobDefinitionEnvironmentProperties>[]>;
        Image: pulumi.Input<string>;
        InstanceType?: pulumi.Input<string>;
        JobRoleArn?: pulumi.Input<string>;
        LinuxParameters?: pulumi.Input<JobDefinitionLinuxParametersProperties>;
        Memory: pulumi.Input<number>;
        MountPoints?: pulumi.Input<pulumi.Input<JobDefinitionMountPointsProperties>[]>;
        Privileged?: pulumi.Input<boolean>;
        ReadonlyRootFilesystem?: pulumi.Input<boolean>;
        ResourceRequirements?: pulumi.Input<pulumi.Input<JobDefinitionResourceRequirementProperties>[]>;
        Ulimits?: pulumi.Input<pulumi.Input<JobDefinitionUlimitProperties>[]>;
        User?: pulumi.Input<string>;
        Vcpus: pulumi.Input<number>;
        Volumes?: pulumi.Input<pulumi.Input<JobDefinitionVolumesProperties>[]>;
    }
    
    export interface JobDefinitionDeviceAttributes {
        ContainerPath?: string;
        HostPath?: string;
        Permissions?: string[];
    }
    
    export interface JobDefinitionDeviceProperties {
        ContainerPath?: pulumi.Input<string>;
        HostPath?: pulumi.Input<string>;
        Permissions?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface JobDefinitionEnvironmentAttributes {
        Name?: string;
        Value?: string;
    }
    
    export interface JobDefinitionEnvironmentProperties {
        Name?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface JobDefinitionLinuxParametersAttributes {
        Devices?: JobDefinitionDeviceAttributes[];
    }
    
    export interface JobDefinitionLinuxParametersProperties {
        Devices?: pulumi.Input<pulumi.Input<JobDefinitionDeviceProperties>[]>;
    }
    
    export interface JobDefinitionMountPointsAttributes {
        ContainerPath?: string;
        ReadOnly?: boolean;
        SourceVolume?: string;
    }
    
    export interface JobDefinitionMountPointsProperties {
        ContainerPath?: pulumi.Input<string>;
        ReadOnly?: pulumi.Input<boolean>;
        SourceVolume?: pulumi.Input<string>;
    }
    
    export interface JobDefinitionNodePropertiesAttributes {
        MainNode: number;
        NodeRangeProperties: JobDefinitionNodeRangePropertyAttributes[];
        NumNodes: number;
    }
    
    export interface JobDefinitionNodePropertiesProperties {
        MainNode: pulumi.Input<number>;
        NodeRangeProperties: pulumi.Input<pulumi.Input<JobDefinitionNodeRangePropertyProperties>[]>;
        NumNodes: pulumi.Input<number>;
    }
    
    export interface JobDefinitionNodeRangePropertyAttributes {
        Container?: JobDefinitionContainerPropertiesAttributes;
        TargetNodes: string;
    }
    
    export interface JobDefinitionNodeRangePropertyProperties {
        Container?: pulumi.Input<JobDefinitionContainerPropertiesProperties>;
        TargetNodes: pulumi.Input<string>;
    }
    
    export interface JobDefinitionProperties {
        ContainerProperties?: pulumi.Input<JobDefinitionContainerPropertiesProperties>;
        JobDefinitionName?: pulumi.Input<string>;
        NodeProperties?: pulumi.Input<JobDefinitionNodePropertiesProperties>;
        Parameters?: pulumi.Input<string>;
        RetryStrategy?: pulumi.Input<JobDefinitionRetryStrategyProperties>;
        Timeout?: pulumi.Input<JobDefinitionTimeoutProperties>;
        Type: pulumi.Input<string>;
    }
    
    export interface JobDefinitionResourceRequirementAttributes {
        Type?: string;
        Value?: string;
    }
    
    export interface JobDefinitionResourceRequirementProperties {
        Type?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface JobDefinitionRetryStrategyAttributes {
        Attempts?: number;
    }
    
    export interface JobDefinitionRetryStrategyProperties {
        Attempts?: pulumi.Input<number>;
    }
    
    export interface JobDefinitionTimeoutAttributes {
        AttemptDurationSeconds?: number;
    }
    
    export interface JobDefinitionTimeoutProperties {
        AttemptDurationSeconds?: pulumi.Input<number>;
    }
    
    export interface JobDefinitionUlimitAttributes {
        HardLimit: number;
        Name: string;
        SoftLimit: number;
    }
    
    export interface JobDefinitionUlimitProperties {
        HardLimit: pulumi.Input<number>;
        Name: pulumi.Input<string>;
        SoftLimit: pulumi.Input<number>;
    }
    
    export interface JobDefinitionVolumesAttributes {
        Host?: JobDefinitionVolumesHostAttributes;
        Name?: string;
    }
    
    export interface JobDefinitionVolumesHostAttributes {
        SourcePath?: string;
    }
    
    export interface JobDefinitionVolumesHostProperties {
        SourcePath?: pulumi.Input<string>;
    }
    
    export interface JobDefinitionVolumesProperties {
        Host?: pulumi.Input<JobDefinitionVolumesHostProperties>;
        Name?: pulumi.Input<string>;
    }
    
    export interface JobQueueAttributes {
    }
    
    export interface JobQueueComputeEnvironmentOrderAttributes {
        ComputeEnvironment: string;
        Order: number;
    }
    
    export interface JobQueueComputeEnvironmentOrderProperties {
        ComputeEnvironment: pulumi.Input<string>;
        Order: pulumi.Input<number>;
    }
    
    export interface JobQueueProperties {
        ComputeEnvironmentOrder: pulumi.Input<pulumi.Input<JobQueueComputeEnvironmentOrderProperties>[]>;
        JobQueueName?: pulumi.Input<string>;
        Priority: pulumi.Input<number>;
        State?: pulumi.Input<string>;
    }
    
    
    export class ComputeEnvironment extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ComputeEnvironmentAttributes>;

        constructor(name: string, properties: ComputeEnvironmentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Batch:ComputeEnvironment", name, inputs, opts);
        }
    }
    
    export class JobDefinition extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<JobDefinitionAttributes>;

        constructor(name: string, properties: JobDefinitionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Batch:JobDefinition", name, inputs, opts);
        }
    }
    
    export class JobQueue extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<JobQueueAttributes>;

        constructor(name: string, properties: JobQueueProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Batch:JobQueue", name, inputs, opts);
        }
    }
    
}

export namespace budgets {
    
    export interface BudgetAttributes {
    }
    
    export interface BudgetBudgetDataAttributes {
        BudgetLimit?: BudgetSpendAttributes;
        BudgetName?: string;
        BudgetType: string;
        CostFilters?: string;
        CostTypes?: BudgetCostTypesAttributes;
        PlannedBudgetLimits?: string;
        TimePeriod?: BudgetTimePeriodAttributes;
        TimeUnit: string;
    }
    
    export interface BudgetBudgetDataProperties {
        BudgetLimit?: pulumi.Input<BudgetSpendProperties>;
        BudgetName?: pulumi.Input<string>;
        BudgetType: pulumi.Input<string>;
        CostFilters?: pulumi.Input<string>;
        CostTypes?: pulumi.Input<BudgetCostTypesProperties>;
        PlannedBudgetLimits?: pulumi.Input<string>;
        TimePeriod?: pulumi.Input<BudgetTimePeriodProperties>;
        TimeUnit: pulumi.Input<string>;
    }
    
    export interface BudgetCostTypesAttributes {
        IncludeCredit?: boolean;
        IncludeDiscount?: boolean;
        IncludeOtherSubscription?: boolean;
        IncludeRecurring?: boolean;
        IncludeRefund?: boolean;
        IncludeSubscription?: boolean;
        IncludeSupport?: boolean;
        IncludeTax?: boolean;
        IncludeUpfront?: boolean;
        UseAmortized?: boolean;
        UseBlended?: boolean;
    }
    
    export interface BudgetCostTypesProperties {
        IncludeCredit?: pulumi.Input<boolean>;
        IncludeDiscount?: pulumi.Input<boolean>;
        IncludeOtherSubscription?: pulumi.Input<boolean>;
        IncludeRecurring?: pulumi.Input<boolean>;
        IncludeRefund?: pulumi.Input<boolean>;
        IncludeSubscription?: pulumi.Input<boolean>;
        IncludeSupport?: pulumi.Input<boolean>;
        IncludeTax?: pulumi.Input<boolean>;
        IncludeUpfront?: pulumi.Input<boolean>;
        UseAmortized?: pulumi.Input<boolean>;
        UseBlended?: pulumi.Input<boolean>;
    }
    
    export interface BudgetNotificationAttributes {
        ComparisonOperator: string;
        NotificationType: string;
        Threshold: number;
        ThresholdType?: string;
    }
    
    export interface BudgetNotificationProperties {
        ComparisonOperator: pulumi.Input<string>;
        NotificationType: pulumi.Input<string>;
        Threshold: pulumi.Input<number>;
        ThresholdType?: pulumi.Input<string>;
    }
    
    export interface BudgetNotificationWithSubscribersAttributes {
        Notification: BudgetNotificationAttributes;
        Subscribers: BudgetSubscriberAttributes[];
    }
    
    export interface BudgetNotificationWithSubscribersProperties {
        Notification: pulumi.Input<BudgetNotificationProperties>;
        Subscribers: pulumi.Input<pulumi.Input<BudgetSubscriberProperties>[]>;
    }
    
    export interface BudgetProperties {
        Budget: pulumi.Input<BudgetBudgetDataProperties>;
        NotificationsWithSubscribers?: pulumi.Input<pulumi.Input<BudgetNotificationWithSubscribersProperties>[]>;
    }
    
    export interface BudgetSpendAttributes {
        Amount: number;
        Unit: string;
    }
    
    export interface BudgetSpendProperties {
        Amount: pulumi.Input<number>;
        Unit: pulumi.Input<string>;
    }
    
    export interface BudgetSubscriberAttributes {
        Address: string;
        SubscriptionType: string;
    }
    
    export interface BudgetSubscriberProperties {
        Address: pulumi.Input<string>;
        SubscriptionType: pulumi.Input<string>;
    }
    
    export interface BudgetTimePeriodAttributes {
        End?: string;
        Start?: string;
    }
    
    export interface BudgetTimePeriodProperties {
        End?: pulumi.Input<string>;
        Start?: pulumi.Input<string>;
    }
    
    
    export class Budget extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<BudgetAttributes>;

        constructor(name: string, properties: BudgetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Budgets:Budget", name, inputs, opts);
        }
    }
    
}

export namespace certificatemanager {
    
    export interface CertificateAttributes {
    }
    
    export interface CertificateDomainValidationOptionAttributes {
        DomainName: string;
        ValidationDomain: string;
    }
    
    export interface CertificateDomainValidationOptionProperties {
        DomainName: pulumi.Input<string>;
        ValidationDomain: pulumi.Input<string>;
    }
    
    export interface CertificateProperties {
        DomainName: pulumi.Input<string>;
        DomainValidationOptions?: pulumi.Input<pulumi.Input<CertificateDomainValidationOptionProperties>[]>;
        SubjectAlternativeNames?: pulumi.Input<pulumi.Input<string>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        ValidationMethod?: pulumi.Input<string>;
    }
    
    
    export class Certificate extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<CertificateAttributes>;

        constructor(name: string, properties: CertificateProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CertificateManager:Certificate", name, inputs, opts);
        }
    }
    
}

export namespace cloud9 {
    
    export interface EnvironmentEC2Attributes {
        Arn: string;
        Name: string;
    }
    
    export interface EnvironmentEC2Properties {
        AutomaticStopTimeMinutes?: pulumi.Input<number>;
        Description?: pulumi.Input<string>;
        InstanceType: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        OwnerArn?: pulumi.Input<string>;
        Repositories?: pulumi.Input<pulumi.Input<EnvironmentEC2RepositoryProperties>[]>;
        SubnetId?: pulumi.Input<string>;
    }
    
    export interface EnvironmentEC2RepositoryAttributes {
        PathComponent: string;
        RepositoryUrl: string;
    }
    
    export interface EnvironmentEC2RepositoryProperties {
        PathComponent: pulumi.Input<string>;
        RepositoryUrl: pulumi.Input<string>;
    }
    
    
    export class EnvironmentEC2 extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<EnvironmentEC2Attributes>;

        constructor(name: string, properties: EnvironmentEC2Properties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Cloud9:EnvironmentEC2", name, inputs, opts);
        }
    }
    
}

export namespace cloudformation {
    
    export interface CustomResourceAttributes {
    }
    
    export interface CustomResourceProperties {
        ServiceToken: pulumi.Input<string>;
    }
    
    export interface MacroAttributes {
    }
    
    export interface MacroProperties {
        Description?: pulumi.Input<string>;
        FunctionName: pulumi.Input<string>;
        LogGroupName?: pulumi.Input<string>;
        LogRoleARN?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
    }
    
    export interface StackAttributes {
    }
    
    export interface StackProperties {
        NotificationARNs?: pulumi.Input<pulumi.Input<string>[]>;
        Parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        TemplateURL: pulumi.Input<string>;
        TimeoutInMinutes?: pulumi.Input<number>;
    }
    
    export interface WaitConditionAttributes {
        Data: string;
    }
    
    export interface WaitConditionHandleAttributes {
    }
    
    export interface WaitConditionHandleProperties {
    }
    
    export interface WaitConditionProperties {
        Count?: pulumi.Input<number>;
        Handle?: pulumi.Input<string>;
        Timeout?: pulumi.Input<string>;
    }
    
    
    export class CustomResource extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<CustomResourceAttributes>;

        constructor(name: string, properties: CustomResourceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CloudFormation:CustomResource", name, inputs, opts);
        }
    }
    
    export class Macro extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<MacroAttributes>;

        constructor(name: string, properties: MacroProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CloudFormation:Macro", name, inputs, opts);
        }
    }
    
    export class Stack extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<StackAttributes>;

        constructor(name: string, properties: StackProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CloudFormation:Stack", name, inputs, opts);
        }
    }
    
    export class WaitCondition extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<WaitConditionAttributes>;

        constructor(name: string, properties?: WaitConditionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CloudFormation:WaitCondition", name, inputs, opts);
        }
    }
    
    export class WaitConditionHandle extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<WaitConditionHandleAttributes>;

        constructor(name: string, properties?: WaitConditionHandleProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CloudFormation:WaitConditionHandle", name, inputs, opts);
        }
    }
    
}

export namespace cloudfront {
    
    export interface CloudFrontOriginAccessIdentityAttributes {
        S3CanonicalUserId: string;
    }
    
    export interface CloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfigAttributes {
        Comment: string;
    }
    
    export interface CloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfigProperties {
        Comment: pulumi.Input<string>;
    }
    
    export interface CloudFrontOriginAccessIdentityProperties {
        CloudFrontOriginAccessIdentityConfig: pulumi.Input<CloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfigProperties>;
    }
    
    export interface DistributionAttributes {
        DomainName: string;
    }
    
    export interface DistributionCacheBehaviorAttributes {
        AllowedMethods?: string[];
        CachedMethods?: string[];
        Compress?: boolean;
        DefaultTTL?: number;
        FieldLevelEncryptionId?: string;
        ForwardedValues: DistributionForwardedValuesAttributes;
        LambdaFunctionAssociations?: DistributionLambdaFunctionAssociationAttributes[];
        MaxTTL?: number;
        MinTTL?: number;
        PathPattern: string;
        SmoothStreaming?: boolean;
        TargetOriginId: string;
        TrustedSigners?: string[];
        ViewerProtocolPolicy: string;
    }
    
    export interface DistributionCacheBehaviorProperties {
        AllowedMethods?: pulumi.Input<pulumi.Input<string>[]>;
        CachedMethods?: pulumi.Input<pulumi.Input<string>[]>;
        Compress?: pulumi.Input<boolean>;
        DefaultTTL?: pulumi.Input<number>;
        FieldLevelEncryptionId?: pulumi.Input<string>;
        ForwardedValues: pulumi.Input<DistributionForwardedValuesProperties>;
        LambdaFunctionAssociations?: pulumi.Input<pulumi.Input<DistributionLambdaFunctionAssociationProperties>[]>;
        MaxTTL?: pulumi.Input<number>;
        MinTTL?: pulumi.Input<number>;
        PathPattern: pulumi.Input<string>;
        SmoothStreaming?: pulumi.Input<boolean>;
        TargetOriginId: pulumi.Input<string>;
        TrustedSigners?: pulumi.Input<pulumi.Input<string>[]>;
        ViewerProtocolPolicy: pulumi.Input<string>;
    }
    
    export interface DistributionCookiesAttributes {
        Forward: string;
        WhitelistedNames?: string[];
    }
    
    export interface DistributionCookiesProperties {
        Forward: pulumi.Input<string>;
        WhitelistedNames?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface DistributionCustomErrorResponseAttributes {
        ErrorCachingMinTTL?: number;
        ErrorCode: number;
        ResponseCode?: number;
        ResponsePagePath?: string;
    }
    
    export interface DistributionCustomErrorResponseProperties {
        ErrorCachingMinTTL?: pulumi.Input<number>;
        ErrorCode: pulumi.Input<number>;
        ResponseCode?: pulumi.Input<number>;
        ResponsePagePath?: pulumi.Input<string>;
    }
    
    export interface DistributionCustomOriginConfigAttributes {
        HTTPPort?: number;
        HTTPSPort?: number;
        OriginKeepaliveTimeout?: number;
        OriginProtocolPolicy: string;
        OriginReadTimeout?: number;
        OriginSSLProtocols?: string[];
    }
    
    export interface DistributionCustomOriginConfigProperties {
        HTTPPort?: pulumi.Input<number>;
        HTTPSPort?: pulumi.Input<number>;
        OriginKeepaliveTimeout?: pulumi.Input<number>;
        OriginProtocolPolicy: pulumi.Input<string>;
        OriginReadTimeout?: pulumi.Input<number>;
        OriginSSLProtocols?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface DistributionDefaultCacheBehaviorAttributes {
        AllowedMethods?: string[];
        CachedMethods?: string[];
        Compress?: boolean;
        DefaultTTL?: number;
        FieldLevelEncryptionId?: string;
        ForwardedValues: DistributionForwardedValuesAttributes;
        LambdaFunctionAssociations?: DistributionLambdaFunctionAssociationAttributes[];
        MaxTTL?: number;
        MinTTL?: number;
        SmoothStreaming?: boolean;
        TargetOriginId: string;
        TrustedSigners?: string[];
        ViewerProtocolPolicy: string;
    }
    
    export interface DistributionDefaultCacheBehaviorProperties {
        AllowedMethods?: pulumi.Input<pulumi.Input<string>[]>;
        CachedMethods?: pulumi.Input<pulumi.Input<string>[]>;
        Compress?: pulumi.Input<boolean>;
        DefaultTTL?: pulumi.Input<number>;
        FieldLevelEncryptionId?: pulumi.Input<string>;
        ForwardedValues: pulumi.Input<DistributionForwardedValuesProperties>;
        LambdaFunctionAssociations?: pulumi.Input<pulumi.Input<DistributionLambdaFunctionAssociationProperties>[]>;
        MaxTTL?: pulumi.Input<number>;
        MinTTL?: pulumi.Input<number>;
        SmoothStreaming?: pulumi.Input<boolean>;
        TargetOriginId: pulumi.Input<string>;
        TrustedSigners?: pulumi.Input<pulumi.Input<string>[]>;
        ViewerProtocolPolicy: pulumi.Input<string>;
    }
    
    export interface DistributionDistributionConfigAttributes {
        Aliases?: string[];
        CacheBehaviors?: DistributionCacheBehaviorAttributes[];
        Comment?: string;
        CustomErrorResponses?: DistributionCustomErrorResponseAttributes[];
        DefaultCacheBehavior?: DistributionDefaultCacheBehaviorAttributes;
        DefaultRootObject?: string;
        Enabled: boolean;
        HttpVersion?: string;
        IPV6Enabled?: boolean;
        Logging?: DistributionLoggingAttributes;
        Origins?: DistributionOriginAttributes[];
        PriceClass?: string;
        Restrictions?: DistributionRestrictionsAttributes;
        ViewerCertificate?: DistributionViewerCertificateAttributes;
        WebACLId?: string;
    }
    
    export interface DistributionDistributionConfigProperties {
        Aliases?: pulumi.Input<pulumi.Input<string>[]>;
        CacheBehaviors?: pulumi.Input<pulumi.Input<DistributionCacheBehaviorProperties>[]>;
        Comment?: pulumi.Input<string>;
        CustomErrorResponses?: pulumi.Input<pulumi.Input<DistributionCustomErrorResponseProperties>[]>;
        DefaultCacheBehavior?: pulumi.Input<DistributionDefaultCacheBehaviorProperties>;
        DefaultRootObject?: pulumi.Input<string>;
        Enabled: pulumi.Input<boolean>;
        HttpVersion?: pulumi.Input<string>;
        IPV6Enabled?: pulumi.Input<boolean>;
        Logging?: pulumi.Input<DistributionLoggingProperties>;
        Origins?: pulumi.Input<pulumi.Input<DistributionOriginProperties>[]>;
        PriceClass?: pulumi.Input<string>;
        Restrictions?: pulumi.Input<DistributionRestrictionsProperties>;
        ViewerCertificate?: pulumi.Input<DistributionViewerCertificateProperties>;
        WebACLId?: pulumi.Input<string>;
    }
    
    export interface DistributionForwardedValuesAttributes {
        Cookies?: DistributionCookiesAttributes;
        Headers?: string[];
        QueryString: boolean;
        QueryStringCacheKeys?: string[];
    }
    
    export interface DistributionForwardedValuesProperties {
        Cookies?: pulumi.Input<DistributionCookiesProperties>;
        Headers?: pulumi.Input<pulumi.Input<string>[]>;
        QueryString: pulumi.Input<boolean>;
        QueryStringCacheKeys?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface DistributionGeoRestrictionAttributes {
        Locations?: string[];
        RestrictionType: string;
    }
    
    export interface DistributionGeoRestrictionProperties {
        Locations?: pulumi.Input<pulumi.Input<string>[]>;
        RestrictionType: pulumi.Input<string>;
    }
    
    export interface DistributionLambdaFunctionAssociationAttributes {
        EventType?: string;
        LambdaFunctionARN?: string;
    }
    
    export interface DistributionLambdaFunctionAssociationProperties {
        EventType?: pulumi.Input<string>;
        LambdaFunctionARN?: pulumi.Input<string>;
    }
    
    export interface DistributionLoggingAttributes {
        Bucket: string;
        IncludeCookies?: boolean;
        Prefix?: string;
    }
    
    export interface DistributionLoggingProperties {
        Bucket: pulumi.Input<string>;
        IncludeCookies?: pulumi.Input<boolean>;
        Prefix?: pulumi.Input<string>;
    }
    
    export interface DistributionOriginAttributes {
        CustomOriginConfig?: DistributionCustomOriginConfigAttributes;
        DomainName: string;
        Id: string;
        OriginCustomHeaders?: DistributionOriginCustomHeaderAttributes[];
        OriginPath?: string;
        S3OriginConfig?: DistributionS3OriginConfigAttributes;
    }
    
    export interface DistributionOriginCustomHeaderAttributes {
        HeaderName: string;
        HeaderValue: string;
    }
    
    export interface DistributionOriginCustomHeaderProperties {
        HeaderName: pulumi.Input<string>;
        HeaderValue: pulumi.Input<string>;
    }
    
    export interface DistributionOriginProperties {
        CustomOriginConfig?: pulumi.Input<DistributionCustomOriginConfigProperties>;
        DomainName: pulumi.Input<string>;
        Id: pulumi.Input<string>;
        OriginCustomHeaders?: pulumi.Input<pulumi.Input<DistributionOriginCustomHeaderProperties>[]>;
        OriginPath?: pulumi.Input<string>;
        S3OriginConfig?: pulumi.Input<DistributionS3OriginConfigProperties>;
    }
    
    export interface DistributionProperties {
        DistributionConfig: pulumi.Input<DistributionDistributionConfigProperties>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface DistributionRestrictionsAttributes {
        GeoRestriction: DistributionGeoRestrictionAttributes;
    }
    
    export interface DistributionRestrictionsProperties {
        GeoRestriction: pulumi.Input<DistributionGeoRestrictionProperties>;
    }
    
    export interface DistributionS3OriginConfigAttributes {
        OriginAccessIdentity?: string;
    }
    
    export interface DistributionS3OriginConfigProperties {
        OriginAccessIdentity?: pulumi.Input<string>;
    }
    
    export interface DistributionViewerCertificateAttributes {
        AcmCertificateArn?: string;
        CloudFrontDefaultCertificate?: boolean;
        IamCertificateId?: string;
        MinimumProtocolVersion?: string;
        SslSupportMethod?: string;
    }
    
    export interface DistributionViewerCertificateProperties {
        AcmCertificateArn?: pulumi.Input<string>;
        CloudFrontDefaultCertificate?: pulumi.Input<boolean>;
        IamCertificateId?: pulumi.Input<string>;
        MinimumProtocolVersion?: pulumi.Input<string>;
        SslSupportMethod?: pulumi.Input<string>;
    }
    
    export interface StreamingDistributionAttributes {
        DomainName: string;
    }
    
    export interface StreamingDistributionLoggingAttributes {
        Bucket: string;
        Enabled: boolean;
        Prefix: string;
    }
    
    export interface StreamingDistributionLoggingProperties {
        Bucket: pulumi.Input<string>;
        Enabled: pulumi.Input<boolean>;
        Prefix: pulumi.Input<string>;
    }
    
    export interface StreamingDistributionProperties {
        StreamingDistributionConfig: pulumi.Input<StreamingDistributionStreamingDistributionConfigProperties>;
        Tags: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface StreamingDistributionS3OriginAttributes {
        DomainName: string;
        OriginAccessIdentity: string;
    }
    
    export interface StreamingDistributionS3OriginProperties {
        DomainName: pulumi.Input<string>;
        OriginAccessIdentity: pulumi.Input<string>;
    }
    
    export interface StreamingDistributionStreamingDistributionConfigAttributes {
        Aliases?: string[];
        Comment: string;
        Enabled: boolean;
        Logging?: StreamingDistributionLoggingAttributes;
        PriceClass?: string;
        S3Origin: StreamingDistributionS3OriginAttributes;
        TrustedSigners: StreamingDistributionTrustedSignersAttributes;
    }
    
    export interface StreamingDistributionStreamingDistributionConfigProperties {
        Aliases?: pulumi.Input<pulumi.Input<string>[]>;
        Comment: pulumi.Input<string>;
        Enabled: pulumi.Input<boolean>;
        Logging?: pulumi.Input<StreamingDistributionLoggingProperties>;
        PriceClass?: pulumi.Input<string>;
        S3Origin: pulumi.Input<StreamingDistributionS3OriginProperties>;
        TrustedSigners: pulumi.Input<StreamingDistributionTrustedSignersProperties>;
    }
    
    export interface StreamingDistributionTrustedSignersAttributes {
        AwsAccountNumbers?: string[];
        Enabled: boolean;
    }
    
    export interface StreamingDistributionTrustedSignersProperties {
        AwsAccountNumbers?: pulumi.Input<pulumi.Input<string>[]>;
        Enabled: pulumi.Input<boolean>;
    }
    
    
    export class CloudFrontOriginAccessIdentity extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<CloudFrontOriginAccessIdentityAttributes>;

        constructor(name: string, properties: CloudFrontOriginAccessIdentityProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CloudFront:CloudFrontOriginAccessIdentity", name, inputs, opts);
        }
    }
    
    export class Distribution extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DistributionAttributes>;

        constructor(name: string, properties: DistributionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CloudFront:Distribution", name, inputs, opts);
        }
    }
    
    export class StreamingDistribution extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<StreamingDistributionAttributes>;

        constructor(name: string, properties: StreamingDistributionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CloudFront:StreamingDistribution", name, inputs, opts);
        }
    }
    
}

export namespace cloudtrail {
    
    export interface TrailAttributes {
        Arn: string;
        SnsTopicArn: string;
    }
    
    export interface TrailDataResourceAttributes {
        Type: string;
        Values?: string[];
    }
    
    export interface TrailDataResourceProperties {
        Type: pulumi.Input<string>;
        Values?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface TrailEventSelectorAttributes {
        DataResources?: TrailDataResourceAttributes[];
        IncludeManagementEvents?: boolean;
        ReadWriteType?: string;
    }
    
    export interface TrailEventSelectorProperties {
        DataResources?: pulumi.Input<pulumi.Input<TrailDataResourceProperties>[]>;
        IncludeManagementEvents?: pulumi.Input<boolean>;
        ReadWriteType?: pulumi.Input<string>;
    }
    
    export interface TrailProperties {
        CloudWatchLogsLogGroupArn?: pulumi.Input<string>;
        CloudWatchLogsRoleArn?: pulumi.Input<string>;
        EnableLogFileValidation?: pulumi.Input<boolean>;
        EventSelectors?: pulumi.Input<pulumi.Input<TrailEventSelectorProperties>[]>;
        IncludeGlobalServiceEvents?: pulumi.Input<boolean>;
        IsLogging: pulumi.Input<boolean>;
        IsMultiRegionTrail?: pulumi.Input<boolean>;
        KMSKeyId?: pulumi.Input<string>;
        S3BucketName: pulumi.Input<string>;
        S3KeyPrefix?: pulumi.Input<string>;
        SnsTopicName?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        TrailName?: pulumi.Input<string>;
    }
    
    
    export class Trail extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TrailAttributes>;

        constructor(name: string, properties: TrailProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CloudTrail:Trail", name, inputs, opts);
        }
    }
    
}

export namespace cloudwatch {
    
    export interface AlarmAttributes {
        Arn: string;
    }
    
    export interface AlarmDimensionAttributes {
        Name: string;
        Value: string;
    }
    
    export interface AlarmDimensionProperties {
        Name: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface AlarmMetricAttributes {
        Dimensions?: AlarmDimensionAttributes[];
        MetricName?: string;
        Namespace?: string;
    }
    
    export interface AlarmMetricDataQueryAttributes {
        Expression?: string;
        Id: string;
        Label?: string;
        MetricStat?: AlarmMetricStatAttributes;
        ReturnData?: boolean;
    }
    
    export interface AlarmMetricDataQueryProperties {
        Expression?: pulumi.Input<string>;
        Id: pulumi.Input<string>;
        Label?: pulumi.Input<string>;
        MetricStat?: pulumi.Input<AlarmMetricStatProperties>;
        ReturnData?: pulumi.Input<boolean>;
    }
    
    export interface AlarmMetricProperties {
        Dimensions?: pulumi.Input<pulumi.Input<AlarmDimensionProperties>[]>;
        MetricName?: pulumi.Input<string>;
        Namespace?: pulumi.Input<string>;
    }
    
    export interface AlarmMetricStatAttributes {
        Metric: AlarmMetricAttributes;
        Period: number;
        Stat: string;
        Unit?: string;
    }
    
    export interface AlarmMetricStatProperties {
        Metric: pulumi.Input<AlarmMetricProperties>;
        Period: pulumi.Input<number>;
        Stat: pulumi.Input<string>;
        Unit?: pulumi.Input<string>;
    }
    
    export interface AlarmProperties {
        ActionsEnabled?: pulumi.Input<boolean>;
        AlarmActions?: pulumi.Input<pulumi.Input<string>[]>;
        AlarmDescription?: pulumi.Input<string>;
        AlarmName?: pulumi.Input<string>;
        ComparisonOperator: pulumi.Input<string>;
        DatapointsToAlarm?: pulumi.Input<number>;
        Dimensions?: pulumi.Input<pulumi.Input<AlarmDimensionProperties>[]>;
        EvaluateLowSampleCountPercentile?: pulumi.Input<string>;
        EvaluationPeriods: pulumi.Input<number>;
        ExtendedStatistic?: pulumi.Input<string>;
        InsufficientDataActions?: pulumi.Input<pulumi.Input<string>[]>;
        MetricName?: pulumi.Input<string>;
        Metrics?: pulumi.Input<pulumi.Input<AlarmMetricDataQueryProperties>[]>;
        Namespace?: pulumi.Input<string>;
        OKActions?: pulumi.Input<pulumi.Input<string>[]>;
        Period?: pulumi.Input<number>;
        Statistic?: pulumi.Input<string>;
        Threshold?: pulumi.Input<number>;
        ThresholdMetricId?: pulumi.Input<string>;
        TreatMissingData?: pulumi.Input<string>;
        Unit?: pulumi.Input<string>;
    }
    
    export interface AnomalyDetectorAttributes {
    }
    
    export interface AnomalyDetectorConfigurationAttributes {
        ExcludedTimeRanges?: AnomalyDetectorRangeAttributes[];
        MetricTimeZone?: string;
    }
    
    export interface AnomalyDetectorConfigurationProperties {
        ExcludedTimeRanges?: pulumi.Input<pulumi.Input<AnomalyDetectorRangeProperties>[]>;
        MetricTimeZone?: pulumi.Input<string>;
    }
    
    export interface AnomalyDetectorDimensionAttributes {
        Name: string;
        Value: string;
    }
    
    export interface AnomalyDetectorDimensionProperties {
        Name: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface AnomalyDetectorProperties {
        Configuration?: pulumi.Input<AnomalyDetectorConfigurationProperties>;
        Dimensions?: pulumi.Input<pulumi.Input<AnomalyDetectorDimensionProperties>[]>;
        MetricName: pulumi.Input<string>;
        Namespace: pulumi.Input<string>;
        Stat: pulumi.Input<string>;
    }
    
    export interface AnomalyDetectorRangeAttributes {
        EndTime: string;
        StartTime: string;
    }
    
    export interface AnomalyDetectorRangeProperties {
        EndTime: pulumi.Input<string>;
        StartTime: pulumi.Input<string>;
    }
    
    export interface DashboardAttributes {
    }
    
    export interface DashboardProperties {
        DashboardBody: pulumi.Input<string>;
        DashboardName?: pulumi.Input<string>;
    }
    
    
    export class Alarm extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AlarmAttributes>;

        constructor(name: string, properties: AlarmProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CloudWatch:Alarm", name, inputs, opts);
        }
    }
    
    export class AnomalyDetector extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AnomalyDetectorAttributes>;

        constructor(name: string, properties: AnomalyDetectorProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CloudWatch:AnomalyDetector", name, inputs, opts);
        }
    }
    
    export class Dashboard extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DashboardAttributes>;

        constructor(name: string, properties: DashboardProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CloudWatch:Dashboard", name, inputs, opts);
        }
    }
    
}

export namespace codebuild {
    
    export interface ProjectArtifactsAttributes {
        ArtifactIdentifier?: string;
        EncryptionDisabled?: boolean;
        Location?: string;
        Name?: string;
        NamespaceType?: string;
        OverrideArtifactName?: boolean;
        Packaging?: string;
        Path?: string;
        Type: string;
    }
    
    export interface ProjectArtifactsProperties {
        ArtifactIdentifier?: pulumi.Input<string>;
        EncryptionDisabled?: pulumi.Input<boolean>;
        Location?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        NamespaceType?: pulumi.Input<string>;
        OverrideArtifactName?: pulumi.Input<boolean>;
        Packaging?: pulumi.Input<string>;
        Path?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface ProjectAttributes {
        Arn: string;
    }
    
    export interface ProjectCloudWatchLogsConfigAttributes {
        GroupName?: string;
        Status: string;
        StreamName?: string;
    }
    
    export interface ProjectCloudWatchLogsConfigProperties {
        GroupName?: pulumi.Input<string>;
        Status: pulumi.Input<string>;
        StreamName?: pulumi.Input<string>;
    }
    
    export interface ProjectEnvironmentAttributes {
        Certificate?: string;
        ComputeType: string;
        EnvironmentVariables?: ProjectEnvironmentVariableAttributes[];
        Image: string;
        ImagePullCredentialsType?: string;
        PrivilegedMode?: boolean;
        RegistryCredential?: ProjectRegistryCredentialAttributes;
        Type: string;
    }
    
    export interface ProjectEnvironmentProperties {
        Certificate?: pulumi.Input<string>;
        ComputeType: pulumi.Input<string>;
        EnvironmentVariables?: pulumi.Input<pulumi.Input<ProjectEnvironmentVariableProperties>[]>;
        Image: pulumi.Input<string>;
        ImagePullCredentialsType?: pulumi.Input<string>;
        PrivilegedMode?: pulumi.Input<boolean>;
        RegistryCredential?: pulumi.Input<ProjectRegistryCredentialProperties>;
        Type: pulumi.Input<string>;
    }
    
    export interface ProjectEnvironmentVariableAttributes {
        Name: string;
        Type?: string;
        Value: string;
    }
    
    export interface ProjectEnvironmentVariableProperties {
        Name: pulumi.Input<string>;
        Type?: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface ProjectFilterGroupAttributes {
    }
    
    export interface ProjectFilterGroupProperties {
    }
    
    export interface ProjectGitSubmodulesConfigAttributes {
        FetchSubmodules: boolean;
    }
    
    export interface ProjectGitSubmodulesConfigProperties {
        FetchSubmodules: pulumi.Input<boolean>;
    }
    
    export interface ProjectLogsConfigAttributes {
        CloudWatchLogs?: ProjectCloudWatchLogsConfigAttributes;
        S3Logs?: ProjectS3LogsConfigAttributes;
    }
    
    export interface ProjectLogsConfigProperties {
        CloudWatchLogs?: pulumi.Input<ProjectCloudWatchLogsConfigProperties>;
        S3Logs?: pulumi.Input<ProjectS3LogsConfigProperties>;
    }
    
    export interface ProjectProjectCacheAttributes {
        Location?: string;
        Modes?: string[];
        Type: string;
    }
    
    export interface ProjectProjectCacheProperties {
        Location?: pulumi.Input<string>;
        Modes?: pulumi.Input<pulumi.Input<string>[]>;
        Type: pulumi.Input<string>;
    }
    
    export interface ProjectProjectSourceVersionAttributes {
        SourceIdentifier: string;
        SourceVersion?: string;
    }
    
    export interface ProjectProjectSourceVersionProperties {
        SourceIdentifier: pulumi.Input<string>;
        SourceVersion?: pulumi.Input<string>;
    }
    
    export interface ProjectProjectTriggersAttributes {
        FilterGroups?: ProjectFilterGroupAttributes[];
        Webhook?: boolean;
    }
    
    export interface ProjectProjectTriggersProperties {
        FilterGroups?: pulumi.Input<pulumi.Input<ProjectFilterGroupProperties>[]>;
        Webhook?: pulumi.Input<boolean>;
    }
    
    export interface ProjectProperties {
        Artifacts: pulumi.Input<ProjectArtifactsProperties>;
        BadgeEnabled?: pulumi.Input<boolean>;
        Cache?: pulumi.Input<ProjectProjectCacheProperties>;
        Description?: pulumi.Input<string>;
        EncryptionKey?: pulumi.Input<string>;
        Environment: pulumi.Input<ProjectEnvironmentProperties>;
        LogsConfig?: pulumi.Input<ProjectLogsConfigProperties>;
        Name?: pulumi.Input<string>;
        QueuedTimeoutInMinutes?: pulumi.Input<number>;
        SecondaryArtifacts?: pulumi.Input<pulumi.Input<ProjectArtifactsProperties>[]>;
        SecondarySourceVersions?: pulumi.Input<pulumi.Input<ProjectProjectSourceVersionProperties>[]>;
        SecondarySources?: pulumi.Input<pulumi.Input<ProjectSourceProperties>[]>;
        ServiceRole: pulumi.Input<string>;
        Source: pulumi.Input<ProjectSourceProperties>;
        SourceVersion?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        TimeoutInMinutes?: pulumi.Input<number>;
        Triggers?: pulumi.Input<ProjectProjectTriggersProperties>;
        VpcConfig?: pulumi.Input<ProjectVpcConfigProperties>;
    }
    
    export interface ProjectRegistryCredentialAttributes {
        Credential: string;
        CredentialProvider: string;
    }
    
    export interface ProjectRegistryCredentialProperties {
        Credential: pulumi.Input<string>;
        CredentialProvider: pulumi.Input<string>;
    }
    
    export interface ProjectS3LogsConfigAttributes {
        EncryptionDisabled?: boolean;
        Location?: string;
        Status: string;
    }
    
    export interface ProjectS3LogsConfigProperties {
        EncryptionDisabled?: pulumi.Input<boolean>;
        Location?: pulumi.Input<string>;
        Status: pulumi.Input<string>;
    }
    
    export interface ProjectSourceAttributes {
        Auth?: ProjectSourceAuthAttributes;
        BuildSpec?: string;
        GitCloneDepth?: number;
        GitSubmodulesConfig?: ProjectGitSubmodulesConfigAttributes;
        InsecureSsl?: boolean;
        Location?: string;
        ReportBuildStatus?: boolean;
        SourceIdentifier?: string;
        Type: string;
    }
    
    export interface ProjectSourceAuthAttributes {
        Resource?: string;
        Type: string;
    }
    
    export interface ProjectSourceAuthProperties {
        Resource?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface ProjectSourceProperties {
        Auth?: pulumi.Input<ProjectSourceAuthProperties>;
        BuildSpec?: pulumi.Input<string>;
        GitCloneDepth?: pulumi.Input<number>;
        GitSubmodulesConfig?: pulumi.Input<ProjectGitSubmodulesConfigProperties>;
        InsecureSsl?: pulumi.Input<boolean>;
        Location?: pulumi.Input<string>;
        ReportBuildStatus?: pulumi.Input<boolean>;
        SourceIdentifier?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface ProjectVpcConfigAttributes {
        SecurityGroupIds?: string[];
        Subnets?: string[];
        VpcId?: string;
    }
    
    export interface ProjectVpcConfigProperties {
        SecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
        Subnets?: pulumi.Input<pulumi.Input<string>[]>;
        VpcId?: pulumi.Input<string>;
    }
    
    export interface ProjectWebhookFilterAttributes {
        ExcludeMatchedPattern?: boolean;
        Pattern: string;
        Type: string;
    }
    
    export interface ProjectWebhookFilterProperties {
        ExcludeMatchedPattern?: pulumi.Input<boolean>;
        Pattern: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface SourceCredentialAttributes {
    }
    
    export interface SourceCredentialProperties {
        AuthType: pulumi.Input<string>;
        ServerType: pulumi.Input<string>;
        Token: pulumi.Input<string>;
        Username?: pulumi.Input<string>;
    }
    
    
    export class Project extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ProjectAttributes>;

        constructor(name: string, properties: ProjectProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CodeBuild:Project", name, inputs, opts);
        }
    }
    
    export class SourceCredential extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SourceCredentialAttributes>;

        constructor(name: string, properties: SourceCredentialProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CodeBuild:SourceCredential", name, inputs, opts);
        }
    }
    
}

export namespace codecommit {
    
    export interface RepositoryAttributes {
        Arn: string;
        CloneUrlHttp: string;
        CloneUrlSsh: string;
        Name: string;
    }
    
    export interface RepositoryCodeAttributes {
        S3: RepositoryS3Attributes;
    }
    
    export interface RepositoryCodeProperties {
        S3: pulumi.Input<RepositoryS3Properties>;
    }
    
    export interface RepositoryProperties {
        Code?: pulumi.Input<RepositoryCodeProperties>;
        RepositoryDescription?: pulumi.Input<string>;
        RepositoryName: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        Triggers?: pulumi.Input<pulumi.Input<RepositoryRepositoryTriggerProperties>[]>;
    }
    
    export interface RepositoryRepositoryTriggerAttributes {
        Branches?: string[];
        CustomData?: string;
        DestinationArn: string;
        Events: string[];
        Name: string;
    }
    
    export interface RepositoryRepositoryTriggerProperties {
        Branches?: pulumi.Input<pulumi.Input<string>[]>;
        CustomData?: pulumi.Input<string>;
        DestinationArn: pulumi.Input<string>;
        Events: pulumi.Input<pulumi.Input<string>[]>;
        Name: pulumi.Input<string>;
    }
    
    export interface RepositoryS3Attributes {
        Bucket: string;
        Key: string;
        ObjectVersion?: string;
    }
    
    export interface RepositoryS3Properties {
        Bucket: pulumi.Input<string>;
        Key: pulumi.Input<string>;
        ObjectVersion?: pulumi.Input<string>;
    }
    
    
    export class Repository extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RepositoryAttributes>;

        constructor(name: string, properties: RepositoryProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CodeCommit:Repository", name, inputs, opts);
        }
    }
    
}

export namespace codedeploy {
    
    export interface ApplicationAttributes {
    }
    
    export interface ApplicationProperties {
        ApplicationName?: pulumi.Input<string>;
        ComputePlatform?: pulumi.Input<string>;
    }
    
    export interface DeploymentConfigAttributes {
    }
    
    export interface DeploymentConfigMinimumHealthyHostsAttributes {
        Type: string;
        Value: number;
    }
    
    export interface DeploymentConfigMinimumHealthyHostsProperties {
        Type: pulumi.Input<string>;
        Value: pulumi.Input<number>;
    }
    
    export interface DeploymentConfigProperties {
        DeploymentConfigName?: pulumi.Input<string>;
        MinimumHealthyHosts?: pulumi.Input<DeploymentConfigMinimumHealthyHostsProperties>;
    }
    
    export interface DeploymentGroupAlarmAttributes {
        Name?: string;
    }
    
    export interface DeploymentGroupAlarmConfigurationAttributes {
        Alarms?: DeploymentGroupAlarmAttributes[];
        Enabled?: boolean;
        IgnorePollAlarmFailure?: boolean;
    }
    
    export interface DeploymentGroupAlarmConfigurationProperties {
        Alarms?: pulumi.Input<pulumi.Input<DeploymentGroupAlarmProperties>[]>;
        Enabled?: pulumi.Input<boolean>;
        IgnorePollAlarmFailure?: pulumi.Input<boolean>;
    }
    
    export interface DeploymentGroupAlarmProperties {
        Name?: pulumi.Input<string>;
    }
    
    export interface DeploymentGroupAttributes {
    }
    
    export interface DeploymentGroupAutoRollbackConfigurationAttributes {
        Enabled?: boolean;
        Events?: string[];
    }
    
    export interface DeploymentGroupAutoRollbackConfigurationProperties {
        Enabled?: pulumi.Input<boolean>;
        Events?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface DeploymentGroupDeploymentAttributes {
        Description?: string;
        IgnoreApplicationStopFailures?: boolean;
        Revision: DeploymentGroupRevisionLocationAttributes;
    }
    
    export interface DeploymentGroupDeploymentProperties {
        Description?: pulumi.Input<string>;
        IgnoreApplicationStopFailures?: pulumi.Input<boolean>;
        Revision: pulumi.Input<DeploymentGroupRevisionLocationProperties>;
    }
    
    export interface DeploymentGroupDeploymentStyleAttributes {
        DeploymentOption?: string;
        DeploymentType?: string;
    }
    
    export interface DeploymentGroupDeploymentStyleProperties {
        DeploymentOption?: pulumi.Input<string>;
        DeploymentType?: pulumi.Input<string>;
    }
    
    export interface DeploymentGroupEC2TagFilterAttributes {
        Key?: string;
        Type?: string;
        Value?: string;
    }
    
    export interface DeploymentGroupEC2TagFilterProperties {
        Key?: pulumi.Input<string>;
        Type?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface DeploymentGroupEC2TagSetAttributes {
        Ec2TagSetList?: DeploymentGroupEC2TagSetListObjectAttributes[];
    }
    
    export interface DeploymentGroupEC2TagSetListObjectAttributes {
        Ec2TagGroup?: DeploymentGroupEC2TagFilterAttributes[];
    }
    
    export interface DeploymentGroupEC2TagSetListObjectProperties {
        Ec2TagGroup?: pulumi.Input<pulumi.Input<DeploymentGroupEC2TagFilterProperties>[]>;
    }
    
    export interface DeploymentGroupEC2TagSetProperties {
        Ec2TagSetList?: pulumi.Input<pulumi.Input<DeploymentGroupEC2TagSetListObjectProperties>[]>;
    }
    
    export interface DeploymentGroupELBInfoAttributes {
        Name?: string;
    }
    
    export interface DeploymentGroupELBInfoProperties {
        Name?: pulumi.Input<string>;
    }
    
    export interface DeploymentGroupGitHubLocationAttributes {
        CommitId: string;
        Repository: string;
    }
    
    export interface DeploymentGroupGitHubLocationProperties {
        CommitId: pulumi.Input<string>;
        Repository: pulumi.Input<string>;
    }
    
    export interface DeploymentGroupLoadBalancerInfoAttributes {
        ElbInfoList?: DeploymentGroupELBInfoAttributes[];
        TargetGroupInfoList?: DeploymentGroupTargetGroupInfoAttributes[];
    }
    
    export interface DeploymentGroupLoadBalancerInfoProperties {
        ElbInfoList?: pulumi.Input<pulumi.Input<DeploymentGroupELBInfoProperties>[]>;
        TargetGroupInfoList?: pulumi.Input<pulumi.Input<DeploymentGroupTargetGroupInfoProperties>[]>;
    }
    
    export interface DeploymentGroupOnPremisesTagSetAttributes {
        OnPremisesTagSetList?: DeploymentGroupOnPremisesTagSetListObjectAttributes[];
    }
    
    export interface DeploymentGroupOnPremisesTagSetListObjectAttributes {
        OnPremisesTagGroup?: DeploymentGroupTagFilterAttributes[];
    }
    
    export interface DeploymentGroupOnPremisesTagSetListObjectProperties {
        OnPremisesTagGroup?: pulumi.Input<pulumi.Input<DeploymentGroupTagFilterProperties>[]>;
    }
    
    export interface DeploymentGroupOnPremisesTagSetProperties {
        OnPremisesTagSetList?: pulumi.Input<pulumi.Input<DeploymentGroupOnPremisesTagSetListObjectProperties>[]>;
    }
    
    export interface DeploymentGroupProperties {
        AlarmConfiguration?: pulumi.Input<DeploymentGroupAlarmConfigurationProperties>;
        ApplicationName: pulumi.Input<string>;
        AutoRollbackConfiguration?: pulumi.Input<DeploymentGroupAutoRollbackConfigurationProperties>;
        AutoScalingGroups?: pulumi.Input<pulumi.Input<string>[]>;
        Deployment?: pulumi.Input<DeploymentGroupDeploymentProperties>;
        DeploymentConfigName?: pulumi.Input<string>;
        DeploymentGroupName?: pulumi.Input<string>;
        DeploymentStyle?: pulumi.Input<DeploymentGroupDeploymentStyleProperties>;
        Ec2TagFilters?: pulumi.Input<pulumi.Input<DeploymentGroupEC2TagFilterProperties>[]>;
        Ec2TagSet?: pulumi.Input<DeploymentGroupEC2TagSetProperties>;
        LoadBalancerInfo?: pulumi.Input<DeploymentGroupLoadBalancerInfoProperties>;
        OnPremisesInstanceTagFilters?: pulumi.Input<pulumi.Input<DeploymentGroupTagFilterProperties>[]>;
        OnPremisesTagSet?: pulumi.Input<DeploymentGroupOnPremisesTagSetProperties>;
        ServiceRoleArn: pulumi.Input<string>;
        TriggerConfigurations?: pulumi.Input<pulumi.Input<DeploymentGroupTriggerConfigProperties>[]>;
    }
    
    export interface DeploymentGroupRevisionLocationAttributes {
        GitHubLocation?: DeploymentGroupGitHubLocationAttributes;
        RevisionType?: string;
        S3Location?: DeploymentGroupS3LocationAttributes;
    }
    
    export interface DeploymentGroupRevisionLocationProperties {
        GitHubLocation?: pulumi.Input<DeploymentGroupGitHubLocationProperties>;
        RevisionType?: pulumi.Input<string>;
        S3Location?: pulumi.Input<DeploymentGroupS3LocationProperties>;
    }
    
    export interface DeploymentGroupS3LocationAttributes {
        Bucket: string;
        BundleType?: string;
        ETag?: string;
        Key: string;
        Version?: string;
    }
    
    export interface DeploymentGroupS3LocationProperties {
        Bucket: pulumi.Input<string>;
        BundleType?: pulumi.Input<string>;
        ETag?: pulumi.Input<string>;
        Key: pulumi.Input<string>;
        Version?: pulumi.Input<string>;
    }
    
    export interface DeploymentGroupTagFilterAttributes {
        Key?: string;
        Type?: string;
        Value?: string;
    }
    
    export interface DeploymentGroupTagFilterProperties {
        Key?: pulumi.Input<string>;
        Type?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface DeploymentGroupTargetGroupInfoAttributes {
        Name?: string;
    }
    
    export interface DeploymentGroupTargetGroupInfoProperties {
        Name?: pulumi.Input<string>;
    }
    
    export interface DeploymentGroupTriggerConfigAttributes {
        TriggerEvents?: string[];
        TriggerName?: string;
        TriggerTargetArn?: string;
    }
    
    export interface DeploymentGroupTriggerConfigProperties {
        TriggerEvents?: pulumi.Input<pulumi.Input<string>[]>;
        TriggerName?: pulumi.Input<string>;
        TriggerTargetArn?: pulumi.Input<string>;
    }
    
    
    export class Application extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ApplicationAttributes>;

        constructor(name: string, properties?: ApplicationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CodeDeploy:Application", name, inputs, opts);
        }
    }
    
    export class DeploymentConfig extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DeploymentConfigAttributes>;

        constructor(name: string, properties?: DeploymentConfigProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CodeDeploy:DeploymentConfig", name, inputs, opts);
        }
    }
    
    export class DeploymentGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DeploymentGroupAttributes>;

        constructor(name: string, properties: DeploymentGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CodeDeploy:DeploymentGroup", name, inputs, opts);
        }
    }
    
}

export namespace codepipeline {
    
    export interface CustomActionTypeArtifactDetailsAttributes {
        MaximumCount: number;
        MinimumCount: number;
    }
    
    export interface CustomActionTypeArtifactDetailsProperties {
        MaximumCount: pulumi.Input<number>;
        MinimumCount: pulumi.Input<number>;
    }
    
    export interface CustomActionTypeAttributes {
    }
    
    export interface CustomActionTypeConfigurationPropertiesAttributes {
        Description?: string;
        Key: boolean;
        Name: string;
        Queryable?: boolean;
        Required: boolean;
        Secret: boolean;
        Type?: string;
    }
    
    export interface CustomActionTypeConfigurationPropertiesProperties {
        Description?: pulumi.Input<string>;
        Key: pulumi.Input<boolean>;
        Name: pulumi.Input<string>;
        Queryable?: pulumi.Input<boolean>;
        Required: pulumi.Input<boolean>;
        Secret: pulumi.Input<boolean>;
        Type?: pulumi.Input<string>;
    }
    
    export interface CustomActionTypeProperties {
        Category: pulumi.Input<string>;
        ConfigurationProperties?: pulumi.Input<pulumi.Input<CustomActionTypeConfigurationPropertiesProperties>[]>;
        InputArtifactDetails: pulumi.Input<CustomActionTypeArtifactDetailsProperties>;
        OutputArtifactDetails: pulumi.Input<CustomActionTypeArtifactDetailsProperties>;
        Provider: pulumi.Input<string>;
        Settings?: pulumi.Input<CustomActionTypeSettingsProperties>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        Version: pulumi.Input<string>;
    }
    
    export interface CustomActionTypeSettingsAttributes {
        EntityUrlTemplate?: string;
        ExecutionUrlTemplate?: string;
        RevisionUrlTemplate?: string;
        ThirdPartyConfigurationUrl?: string;
    }
    
    export interface CustomActionTypeSettingsProperties {
        EntityUrlTemplate?: pulumi.Input<string>;
        ExecutionUrlTemplate?: pulumi.Input<string>;
        RevisionUrlTemplate?: pulumi.Input<string>;
        ThirdPartyConfigurationUrl?: pulumi.Input<string>;
    }
    
    export interface PipelineActionDeclarationAttributes {
        ActionTypeId: PipelineActionTypeIdAttributes;
        Configuration?: string;
        InputArtifacts?: PipelineInputArtifactAttributes[];
        Name: string;
        OutputArtifacts?: PipelineOutputArtifactAttributes[];
        Region?: string;
        RoleArn?: string;
        RunOrder?: number;
    }
    
    export interface PipelineActionDeclarationProperties {
        ActionTypeId: pulumi.Input<PipelineActionTypeIdProperties>;
        Configuration?: pulumi.Input<string>;
        InputArtifacts?: pulumi.Input<pulumi.Input<PipelineInputArtifactProperties>[]>;
        Name: pulumi.Input<string>;
        OutputArtifacts?: pulumi.Input<pulumi.Input<PipelineOutputArtifactProperties>[]>;
        Region?: pulumi.Input<string>;
        RoleArn?: pulumi.Input<string>;
        RunOrder?: pulumi.Input<number>;
    }
    
    export interface PipelineActionTypeIdAttributes {
        Category: string;
        Owner: string;
        Provider: string;
        Version: string;
    }
    
    export interface PipelineActionTypeIdProperties {
        Category: pulumi.Input<string>;
        Owner: pulumi.Input<string>;
        Provider: pulumi.Input<string>;
        Version: pulumi.Input<string>;
    }
    
    export interface PipelineArtifactStoreAttributes {
        EncryptionKey?: PipelineEncryptionKeyAttributes;
        Location: string;
        Type: string;
    }
    
    export interface PipelineArtifactStoreMapAttributes {
        ArtifactStore: PipelineArtifactStoreAttributes;
        Region: string;
    }
    
    export interface PipelineArtifactStoreMapProperties {
        ArtifactStore: pulumi.Input<PipelineArtifactStoreProperties>;
        Region: pulumi.Input<string>;
    }
    
    export interface PipelineArtifactStoreProperties {
        EncryptionKey?: pulumi.Input<PipelineEncryptionKeyProperties>;
        Location: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface PipelineAttributes {
        Version: string;
    }
    
    export interface PipelineBlockerDeclarationAttributes {
        Name: string;
        Type: string;
    }
    
    export interface PipelineBlockerDeclarationProperties {
        Name: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface PipelineEncryptionKeyAttributes {
        Id: string;
        Type: string;
    }
    
    export interface PipelineEncryptionKeyProperties {
        Id: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface PipelineInputArtifactAttributes {
        Name: string;
    }
    
    export interface PipelineInputArtifactProperties {
        Name: pulumi.Input<string>;
    }
    
    export interface PipelineOutputArtifactAttributes {
        Name: string;
    }
    
    export interface PipelineOutputArtifactProperties {
        Name: pulumi.Input<string>;
    }
    
    export interface PipelineProperties {
        ArtifactStore?: pulumi.Input<PipelineArtifactStoreProperties>;
        ArtifactStores?: pulumi.Input<pulumi.Input<PipelineArtifactStoreMapProperties>[]>;
        DisableInboundStageTransitions?: pulumi.Input<pulumi.Input<PipelineStageTransitionProperties>[]>;
        Name?: pulumi.Input<string>;
        RestartExecutionOnUpdate?: pulumi.Input<boolean>;
        RoleArn: pulumi.Input<string>;
        Stages: pulumi.Input<pulumi.Input<PipelineStageDeclarationProperties>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface PipelineStageDeclarationAttributes {
        Actions: PipelineActionDeclarationAttributes[];
        Blockers?: PipelineBlockerDeclarationAttributes[];
        Name: string;
    }
    
    export interface PipelineStageDeclarationProperties {
        Actions: pulumi.Input<pulumi.Input<PipelineActionDeclarationProperties>[]>;
        Blockers?: pulumi.Input<pulumi.Input<PipelineBlockerDeclarationProperties>[]>;
        Name: pulumi.Input<string>;
    }
    
    export interface PipelineStageTransitionAttributes {
        Reason: string;
        StageName: string;
    }
    
    export interface PipelineStageTransitionProperties {
        Reason: pulumi.Input<string>;
        StageName: pulumi.Input<string>;
    }
    
    export interface WebhookAttributes {
        Url: string;
    }
    
    export interface WebhookProperties {
        Authentication: pulumi.Input<string>;
        AuthenticationConfiguration: pulumi.Input<WebhookWebhookAuthConfigurationProperties>;
        Filters: pulumi.Input<pulumi.Input<WebhookWebhookFilterRuleProperties>[]>;
        Name?: pulumi.Input<string>;
        RegisterWithThirdParty?: pulumi.Input<boolean>;
        TargetAction: pulumi.Input<string>;
        TargetPipeline: pulumi.Input<string>;
        TargetPipelineVersion: pulumi.Input<number>;
    }
    
    export interface WebhookWebhookAuthConfigurationAttributes {
        AllowedIPRange?: string;
        SecretToken?: string;
    }
    
    export interface WebhookWebhookAuthConfigurationProperties {
        AllowedIPRange?: pulumi.Input<string>;
        SecretToken?: pulumi.Input<string>;
    }
    
    export interface WebhookWebhookFilterRuleAttributes {
        JsonPath: string;
        MatchEquals?: string;
    }
    
    export interface WebhookWebhookFilterRuleProperties {
        JsonPath: pulumi.Input<string>;
        MatchEquals?: pulumi.Input<string>;
    }
    
    
    export class CustomActionType extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<CustomActionTypeAttributes>;

        constructor(name: string, properties: CustomActionTypeProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CodePipeline:CustomActionType", name, inputs, opts);
        }
    }
    
    export class Pipeline extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PipelineAttributes>;

        constructor(name: string, properties: PipelineProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CodePipeline:Pipeline", name, inputs, opts);
        }
    }
    
    export class Webhook extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<WebhookAttributes>;

        constructor(name: string, properties: WebhookProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CodePipeline:Webhook", name, inputs, opts);
        }
    }
    
}

export namespace codestar {
    
    export interface GitHubRepositoryAttributes {
    }
    
    export interface GitHubRepositoryCodeAttributes {
        S3: GitHubRepositoryS3Attributes;
    }
    
    export interface GitHubRepositoryCodeProperties {
        S3: pulumi.Input<GitHubRepositoryS3Properties>;
    }
    
    export interface GitHubRepositoryProperties {
        Code?: pulumi.Input<GitHubRepositoryCodeProperties>;
        EnableIssues?: pulumi.Input<boolean>;
        IsPrivate?: pulumi.Input<boolean>;
        RepositoryAccessToken: pulumi.Input<string>;
        RepositoryDescription?: pulumi.Input<string>;
        RepositoryName: pulumi.Input<string>;
        RepositoryOwner: pulumi.Input<string>;
    }
    
    export interface GitHubRepositoryS3Attributes {
        Bucket: string;
        Key: string;
        ObjectVersion?: string;
    }
    
    export interface GitHubRepositoryS3Properties {
        Bucket: pulumi.Input<string>;
        Key: pulumi.Input<string>;
        ObjectVersion?: pulumi.Input<string>;
    }
    
    
    export class GitHubRepository extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<GitHubRepositoryAttributes>;

        constructor(name: string, properties: GitHubRepositoryProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:CodeStar:GitHubRepository", name, inputs, opts);
        }
    }
    
}

export namespace cognito {
    
    export interface IdentityPoolAttributes {
        Name: string;
    }
    
    export interface IdentityPoolCognitoIdentityProviderAttributes {
        ClientId?: string;
        ProviderName?: string;
        ServerSideTokenCheck?: boolean;
    }
    
    export interface IdentityPoolCognitoIdentityProviderProperties {
        ClientId?: pulumi.Input<string>;
        ProviderName?: pulumi.Input<string>;
        ServerSideTokenCheck?: pulumi.Input<boolean>;
    }
    
    export interface IdentityPoolCognitoStreamsAttributes {
        RoleArn?: string;
        StreamName?: string;
        StreamingStatus?: string;
    }
    
    export interface IdentityPoolCognitoStreamsProperties {
        RoleArn?: pulumi.Input<string>;
        StreamName?: pulumi.Input<string>;
        StreamingStatus?: pulumi.Input<string>;
    }
    
    export interface IdentityPoolProperties {
        AllowUnauthenticatedIdentities: pulumi.Input<boolean>;
        CognitoEvents?: pulumi.Input<string>;
        CognitoIdentityProviders?: pulumi.Input<pulumi.Input<IdentityPoolCognitoIdentityProviderProperties>[]>;
        CognitoStreams?: pulumi.Input<IdentityPoolCognitoStreamsProperties>;
        DeveloperProviderName?: pulumi.Input<string>;
        IdentityPoolName?: pulumi.Input<string>;
        OpenIdConnectProviderARNs?: pulumi.Input<pulumi.Input<string>[]>;
        PushSync?: pulumi.Input<IdentityPoolPushSyncProperties>;
        SamlProviderARNs?: pulumi.Input<pulumi.Input<string>[]>;
        SupportedLoginProviders?: pulumi.Input<string>;
    }
    
    export interface IdentityPoolPushSyncAttributes {
        ApplicationArns?: string[];
        RoleArn?: string;
    }
    
    export interface IdentityPoolPushSyncProperties {
        ApplicationArns?: pulumi.Input<pulumi.Input<string>[]>;
        RoleArn?: pulumi.Input<string>;
    }
    
    export interface IdentityPoolRoleAttachmentAttributes {
    }
    
    export interface IdentityPoolRoleAttachmentMappingRuleAttributes {
        Claim: string;
        MatchType: string;
        RoleARN: string;
        Value: string;
    }
    
    export interface IdentityPoolRoleAttachmentMappingRuleProperties {
        Claim: pulumi.Input<string>;
        MatchType: pulumi.Input<string>;
        RoleARN: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface IdentityPoolRoleAttachmentProperties {
        IdentityPoolId: pulumi.Input<string>;
        RoleMappings?: pulumi.Input<string>;
        Roles?: pulumi.Input<string>;
    }
    
    export interface IdentityPoolRoleAttachmentRoleMappingAttributes {
        AmbiguousRoleResolution?: string;
        IdentityProvider?: string;
        RulesConfiguration?: IdentityPoolRoleAttachmentRulesConfigurationTypeAttributes;
        Type: string;
    }
    
    export interface IdentityPoolRoleAttachmentRoleMappingProperties {
        AmbiguousRoleResolution?: pulumi.Input<string>;
        IdentityProvider?: pulumi.Input<string>;
        RulesConfiguration?: pulumi.Input<IdentityPoolRoleAttachmentRulesConfigurationTypeProperties>;
        Type: pulumi.Input<string>;
    }
    
    export interface IdentityPoolRoleAttachmentRulesConfigurationTypeAttributes {
        Rules: IdentityPoolRoleAttachmentMappingRuleAttributes[];
    }
    
    export interface IdentityPoolRoleAttachmentRulesConfigurationTypeProperties {
        Rules: pulumi.Input<pulumi.Input<IdentityPoolRoleAttachmentMappingRuleProperties>[]>;
    }
    
    export interface UserPoolAdminCreateUserConfigAttributes {
        AllowAdminCreateUserOnly?: boolean;
        InviteMessageTemplate?: UserPoolInviteMessageTemplateAttributes;
        UnusedAccountValidityDays?: number;
    }
    
    export interface UserPoolAdminCreateUserConfigProperties {
        AllowAdminCreateUserOnly?: pulumi.Input<boolean>;
        InviteMessageTemplate?: pulumi.Input<UserPoolInviteMessageTemplateProperties>;
        UnusedAccountValidityDays?: pulumi.Input<number>;
    }
    
    export interface UserPoolAttributes {
        Arn: string;
        ProviderName: string;
        ProviderURL: string;
    }
    
    export interface UserPoolClientAnalyticsConfigurationAttributes {
        ApplicationId?: string;
        ExternalId?: string;
        RoleArn?: string;
        UserDataShared?: boolean;
    }
    
    export interface UserPoolClientAnalyticsConfigurationProperties {
        ApplicationId?: pulumi.Input<string>;
        ExternalId?: pulumi.Input<string>;
        RoleArn?: pulumi.Input<string>;
        UserDataShared?: pulumi.Input<boolean>;
    }
    
    export interface UserPoolClientAttributes {
        ClientSecret: string;
        Name: string;
    }
    
    export interface UserPoolClientProperties {
        AllowedOAuthFlows?: pulumi.Input<pulumi.Input<string>[]>;
        AllowedOAuthFlowsUserPoolClient?: pulumi.Input<boolean>;
        AllowedOAuthScopes?: pulumi.Input<pulumi.Input<string>[]>;
        AnalyticsConfiguration?: pulumi.Input<UserPoolClientAnalyticsConfigurationProperties>;
        CallbackURLs?: pulumi.Input<pulumi.Input<string>[]>;
        ClientName?: pulumi.Input<string>;
        DefaultRedirectURI?: pulumi.Input<string>;
        ExplicitAuthFlows?: pulumi.Input<pulumi.Input<string>[]>;
        GenerateSecret?: pulumi.Input<boolean>;
        LogoutURLs?: pulumi.Input<pulumi.Input<string>[]>;
        ReadAttributes?: pulumi.Input<pulumi.Input<string>[]>;
        RefreshTokenValidity?: pulumi.Input<number>;
        SupportedIdentityProviders?: pulumi.Input<pulumi.Input<string>[]>;
        UserPoolId: pulumi.Input<string>;
        WriteAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface UserPoolDeviceConfigurationAttributes {
        ChallengeRequiredOnNewDevice?: boolean;
        DeviceOnlyRememberedOnUserPrompt?: boolean;
    }
    
    export interface UserPoolDeviceConfigurationProperties {
        ChallengeRequiredOnNewDevice?: pulumi.Input<boolean>;
        DeviceOnlyRememberedOnUserPrompt?: pulumi.Input<boolean>;
    }
    
    export interface UserPoolDomainAttributes {
    }
    
    export interface UserPoolDomainCustomDomainConfigTypeAttributes {
        CertificateArn?: string;
    }
    
    export interface UserPoolDomainCustomDomainConfigTypeProperties {
        CertificateArn?: pulumi.Input<string>;
    }
    
    export interface UserPoolDomainProperties {
        CustomDomainConfig?: pulumi.Input<UserPoolDomainCustomDomainConfigTypeProperties>;
        Domain: pulumi.Input<string>;
        UserPoolId: pulumi.Input<string>;
    }
    
    export interface UserPoolEmailConfigurationAttributes {
        EmailSendingAccount?: string;
        ReplyToEmailAddress?: string;
        SourceArn?: string;
    }
    
    export interface UserPoolEmailConfigurationProperties {
        EmailSendingAccount?: pulumi.Input<string>;
        ReplyToEmailAddress?: pulumi.Input<string>;
        SourceArn?: pulumi.Input<string>;
    }
    
    export interface UserPoolGroupAttributes {
    }
    
    export interface UserPoolGroupProperties {
        Description?: pulumi.Input<string>;
        GroupName?: pulumi.Input<string>;
        Precedence?: pulumi.Input<number>;
        RoleArn?: pulumi.Input<string>;
        UserPoolId: pulumi.Input<string>;
    }
    
    export interface UserPoolIdentityProviderAttributes {
    }
    
    export interface UserPoolIdentityProviderProperties {
        AttributeMapping?: pulumi.Input<string>;
        IdpIdentifiers?: pulumi.Input<pulumi.Input<string>[]>;
        ProviderDetails?: pulumi.Input<string>;
        ProviderName: pulumi.Input<string>;
        ProviderType: pulumi.Input<string>;
        UserPoolId: pulumi.Input<string>;
    }
    
    export interface UserPoolInviteMessageTemplateAttributes {
        EmailMessage?: string;
        EmailSubject?: string;
        SMSMessage?: string;
    }
    
    export interface UserPoolInviteMessageTemplateProperties {
        EmailMessage?: pulumi.Input<string>;
        EmailSubject?: pulumi.Input<string>;
        SMSMessage?: pulumi.Input<string>;
    }
    
    export interface UserPoolLambdaConfigAttributes {
        CreateAuthChallenge?: string;
        CustomMessage?: string;
        DefineAuthChallenge?: string;
        PostAuthentication?: string;
        PostConfirmation?: string;
        PreAuthentication?: string;
        PreSignUp?: string;
        PreTokenGeneration?: string;
        UserMigration?: string;
        VerifyAuthChallengeResponse?: string;
    }
    
    export interface UserPoolLambdaConfigProperties {
        CreateAuthChallenge?: pulumi.Input<string>;
        CustomMessage?: pulumi.Input<string>;
        DefineAuthChallenge?: pulumi.Input<string>;
        PostAuthentication?: pulumi.Input<string>;
        PostConfirmation?: pulumi.Input<string>;
        PreAuthentication?: pulumi.Input<string>;
        PreSignUp?: pulumi.Input<string>;
        PreTokenGeneration?: pulumi.Input<string>;
        UserMigration?: pulumi.Input<string>;
        VerifyAuthChallengeResponse?: pulumi.Input<string>;
    }
    
    export interface UserPoolNumberAttributeConstraintsAttributes {
        MaxValue?: string;
        MinValue?: string;
    }
    
    export interface UserPoolNumberAttributeConstraintsProperties {
        MaxValue?: pulumi.Input<string>;
        MinValue?: pulumi.Input<string>;
    }
    
    export interface UserPoolPasswordPolicyAttributes {
        MinimumLength?: number;
        RequireLowercase?: boolean;
        RequireNumbers?: boolean;
        RequireSymbols?: boolean;
        RequireUppercase?: boolean;
        TemporaryPasswordValidityDays?: number;
    }
    
    export interface UserPoolPasswordPolicyProperties {
        MinimumLength?: pulumi.Input<number>;
        RequireLowercase?: pulumi.Input<boolean>;
        RequireNumbers?: pulumi.Input<boolean>;
        RequireSymbols?: pulumi.Input<boolean>;
        RequireUppercase?: pulumi.Input<boolean>;
        TemporaryPasswordValidityDays?: pulumi.Input<number>;
    }
    
    export interface UserPoolPoliciesAttributes {
        PasswordPolicy?: UserPoolPasswordPolicyAttributes;
    }
    
    export interface UserPoolPoliciesProperties {
        PasswordPolicy?: pulumi.Input<UserPoolPasswordPolicyProperties>;
    }
    
    export interface UserPoolProperties {
        AdminCreateUserConfig?: pulumi.Input<UserPoolAdminCreateUserConfigProperties>;
        AliasAttributes?: pulumi.Input<pulumi.Input<string>[]>;
        AutoVerifiedAttributes?: pulumi.Input<pulumi.Input<string>[]>;
        DeviceConfiguration?: pulumi.Input<UserPoolDeviceConfigurationProperties>;
        EmailConfiguration?: pulumi.Input<UserPoolEmailConfigurationProperties>;
        EmailVerificationMessage?: pulumi.Input<string>;
        EmailVerificationSubject?: pulumi.Input<string>;
        EnabledMfas?: pulumi.Input<pulumi.Input<string>[]>;
        LambdaConfig?: pulumi.Input<UserPoolLambdaConfigProperties>;
        MfaConfiguration?: pulumi.Input<string>;
        Policies?: pulumi.Input<UserPoolPoliciesProperties>;
        Schema?: pulumi.Input<pulumi.Input<UserPoolSchemaAttributeProperties>[]>;
        SmsAuthenticationMessage?: pulumi.Input<string>;
        SmsConfiguration?: pulumi.Input<UserPoolSmsConfigurationProperties>;
        SmsVerificationMessage?: pulumi.Input<string>;
        UserPoolAddOns?: pulumi.Input<UserPoolUserPoolAddOnsProperties>;
        UserPoolName?: pulumi.Input<string>;
        UserPoolTags?: pulumi.Input<string>;
        UsernameAttributes?: pulumi.Input<pulumi.Input<string>[]>;
        VerificationMessageTemplate?: pulumi.Input<UserPoolVerificationMessageTemplateProperties>;
    }
    
    export interface UserPoolResourceServerAttributes {
    }
    
    export interface UserPoolResourceServerProperties {
        Identifier: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        Scopes?: pulumi.Input<pulumi.Input<UserPoolResourceServerResourceServerScopeTypeProperties>[]>;
        UserPoolId: pulumi.Input<string>;
    }
    
    export interface UserPoolResourceServerResourceServerScopeTypeAttributes {
        ScopeDescription: string;
        ScopeName: string;
    }
    
    export interface UserPoolResourceServerResourceServerScopeTypeProperties {
        ScopeDescription: pulumi.Input<string>;
        ScopeName: pulumi.Input<string>;
    }
    
    export interface UserPoolRiskConfigurationAttachmentAccountTakeoverActionTypeAttributes {
        EventAction: string;
        Notify: boolean;
    }
    
    export interface UserPoolRiskConfigurationAttachmentAccountTakeoverActionTypeProperties {
        EventAction: pulumi.Input<string>;
        Notify: pulumi.Input<boolean>;
    }
    
    export interface UserPoolRiskConfigurationAttachmentAccountTakeoverActionsTypeAttributes {
        HighAction?: UserPoolRiskConfigurationAttachmentAccountTakeoverActionTypeAttributes;
        LowAction?: UserPoolRiskConfigurationAttachmentAccountTakeoverActionTypeAttributes;
        MediumAction?: UserPoolRiskConfigurationAttachmentAccountTakeoverActionTypeAttributes;
    }
    
    export interface UserPoolRiskConfigurationAttachmentAccountTakeoverActionsTypeProperties {
        HighAction?: pulumi.Input<UserPoolRiskConfigurationAttachmentAccountTakeoverActionTypeProperties>;
        LowAction?: pulumi.Input<UserPoolRiskConfigurationAttachmentAccountTakeoverActionTypeProperties>;
        MediumAction?: pulumi.Input<UserPoolRiskConfigurationAttachmentAccountTakeoverActionTypeProperties>;
    }
    
    export interface UserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationTypeAttributes {
        Actions: UserPoolRiskConfigurationAttachmentAccountTakeoverActionsTypeAttributes;
        NotifyConfiguration?: UserPoolRiskConfigurationAttachmentNotifyConfigurationTypeAttributes;
    }
    
    export interface UserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationTypeProperties {
        Actions: pulumi.Input<UserPoolRiskConfigurationAttachmentAccountTakeoverActionsTypeProperties>;
        NotifyConfiguration?: pulumi.Input<UserPoolRiskConfigurationAttachmentNotifyConfigurationTypeProperties>;
    }
    
    export interface UserPoolRiskConfigurationAttachmentAttributes {
    }
    
    export interface UserPoolRiskConfigurationAttachmentCompromisedCredentialsActionsTypeAttributes {
        EventAction: string;
    }
    
    export interface UserPoolRiskConfigurationAttachmentCompromisedCredentialsActionsTypeProperties {
        EventAction: pulumi.Input<string>;
    }
    
    export interface UserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationTypeAttributes {
        Actions: UserPoolRiskConfigurationAttachmentCompromisedCredentialsActionsTypeAttributes;
        EventFilter?: string[];
    }
    
    export interface UserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationTypeProperties {
        Actions: pulumi.Input<UserPoolRiskConfigurationAttachmentCompromisedCredentialsActionsTypeProperties>;
        EventFilter?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface UserPoolRiskConfigurationAttachmentNotifyConfigurationTypeAttributes {
        BlockEmail?: UserPoolRiskConfigurationAttachmentNotifyEmailTypeAttributes;
        From?: string;
        MfaEmail?: UserPoolRiskConfigurationAttachmentNotifyEmailTypeAttributes;
        NoActionEmail?: UserPoolRiskConfigurationAttachmentNotifyEmailTypeAttributes;
        ReplyTo?: string;
        SourceArn: string;
    }
    
    export interface UserPoolRiskConfigurationAttachmentNotifyConfigurationTypeProperties {
        BlockEmail?: pulumi.Input<UserPoolRiskConfigurationAttachmentNotifyEmailTypeProperties>;
        From?: pulumi.Input<string>;
        MfaEmail?: pulumi.Input<UserPoolRiskConfigurationAttachmentNotifyEmailTypeProperties>;
        NoActionEmail?: pulumi.Input<UserPoolRiskConfigurationAttachmentNotifyEmailTypeProperties>;
        ReplyTo?: pulumi.Input<string>;
        SourceArn: pulumi.Input<string>;
    }
    
    export interface UserPoolRiskConfigurationAttachmentNotifyEmailTypeAttributes {
        HtmlBody?: string;
        Subject: string;
        TextBody?: string;
    }
    
    export interface UserPoolRiskConfigurationAttachmentNotifyEmailTypeProperties {
        HtmlBody?: pulumi.Input<string>;
        Subject: pulumi.Input<string>;
        TextBody?: pulumi.Input<string>;
    }
    
    export interface UserPoolRiskConfigurationAttachmentProperties {
        AccountTakeoverRiskConfiguration?: pulumi.Input<UserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationTypeProperties>;
        ClientId: pulumi.Input<string>;
        CompromisedCredentialsRiskConfiguration?: pulumi.Input<UserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationTypeProperties>;
        RiskExceptionConfiguration?: pulumi.Input<UserPoolRiskConfigurationAttachmentRiskExceptionConfigurationTypeProperties>;
        UserPoolId: pulumi.Input<string>;
    }
    
    export interface UserPoolRiskConfigurationAttachmentRiskExceptionConfigurationTypeAttributes {
        BlockedIPRangeList?: string[];
        SkippedIPRangeList?: string[];
    }
    
    export interface UserPoolRiskConfigurationAttachmentRiskExceptionConfigurationTypeProperties {
        BlockedIPRangeList?: pulumi.Input<pulumi.Input<string>[]>;
        SkippedIPRangeList?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface UserPoolSchemaAttributeAttributes {
        AttributeDataType?: string;
        DeveloperOnlyAttribute?: boolean;
        Mutable?: boolean;
        Name?: string;
        NumberAttributeConstraints?: UserPoolNumberAttributeConstraintsAttributes;
        Required?: boolean;
        StringAttributeConstraints?: UserPoolStringAttributeConstraintsAttributes;
    }
    
    export interface UserPoolSchemaAttributeProperties {
        AttributeDataType?: pulumi.Input<string>;
        DeveloperOnlyAttribute?: pulumi.Input<boolean>;
        Mutable?: pulumi.Input<boolean>;
        Name?: pulumi.Input<string>;
        NumberAttributeConstraints?: pulumi.Input<UserPoolNumberAttributeConstraintsProperties>;
        Required?: pulumi.Input<boolean>;
        StringAttributeConstraints?: pulumi.Input<UserPoolStringAttributeConstraintsProperties>;
    }
    
    export interface UserPoolSmsConfigurationAttributes {
        ExternalId?: string;
        SnsCallerArn?: string;
    }
    
    export interface UserPoolSmsConfigurationProperties {
        ExternalId?: pulumi.Input<string>;
        SnsCallerArn?: pulumi.Input<string>;
    }
    
    export interface UserPoolStringAttributeConstraintsAttributes {
        MaxLength?: string;
        MinLength?: string;
    }
    
    export interface UserPoolStringAttributeConstraintsProperties {
        MaxLength?: pulumi.Input<string>;
        MinLength?: pulumi.Input<string>;
    }
    
    export interface UserPoolUICustomizationAttachmentAttributes {
    }
    
    export interface UserPoolUICustomizationAttachmentProperties {
        CSS?: pulumi.Input<string>;
        ClientId: pulumi.Input<string>;
        UserPoolId: pulumi.Input<string>;
    }
    
    export interface UserPoolUserAttributeTypeAttributes {
        Name?: string;
        Value?: string;
    }
    
    export interface UserPoolUserAttributeTypeProperties {
        Name?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface UserPoolUserAttributes {
    }
    
    export interface UserPoolUserPoolAddOnsAttributes {
        AdvancedSecurityMode?: string;
    }
    
    export interface UserPoolUserPoolAddOnsProperties {
        AdvancedSecurityMode?: pulumi.Input<string>;
    }
    
    export interface UserPoolUserProperties {
        DesiredDeliveryMediums?: pulumi.Input<pulumi.Input<string>[]>;
        ForceAliasCreation?: pulumi.Input<boolean>;
        MessageAction?: pulumi.Input<string>;
        UserAttributes?: pulumi.Input<pulumi.Input<UserPoolUserAttributeTypeProperties>[]>;
        UserPoolId: pulumi.Input<string>;
        Username?: pulumi.Input<string>;
        ValidationData?: pulumi.Input<pulumi.Input<UserPoolUserAttributeTypeProperties>[]>;
    }
    
    export interface UserPoolUserToGroupAttachmentAttributes {
    }
    
    export interface UserPoolUserToGroupAttachmentProperties {
        GroupName: pulumi.Input<string>;
        UserPoolId: pulumi.Input<string>;
        Username: pulumi.Input<string>;
    }
    
    export interface UserPoolVerificationMessageTemplateAttributes {
        DefaultEmailOption?: string;
        EmailMessage?: string;
        EmailMessageByLink?: string;
        EmailSubject?: string;
        EmailSubjectByLink?: string;
        SmsMessage?: string;
    }
    
    export interface UserPoolVerificationMessageTemplateProperties {
        DefaultEmailOption?: pulumi.Input<string>;
        EmailMessage?: pulumi.Input<string>;
        EmailMessageByLink?: pulumi.Input<string>;
        EmailSubject?: pulumi.Input<string>;
        EmailSubjectByLink?: pulumi.Input<string>;
        SmsMessage?: pulumi.Input<string>;
    }
    
    
    export class IdentityPool extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<IdentityPoolAttributes>;

        constructor(name: string, properties: IdentityPoolProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Cognito:IdentityPool", name, inputs, opts);
        }
    }
    
    export class IdentityPoolRoleAttachment extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<IdentityPoolRoleAttachmentAttributes>;

        constructor(name: string, properties: IdentityPoolRoleAttachmentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Cognito:IdentityPoolRoleAttachment", name, inputs, opts);
        }
    }
    
    export class UserPool extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<UserPoolAttributes>;

        constructor(name: string, properties?: UserPoolProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Cognito:UserPool", name, inputs, opts);
        }
    }
    
    export class UserPoolClient extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<UserPoolClientAttributes>;

        constructor(name: string, properties: UserPoolClientProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Cognito:UserPoolClient", name, inputs, opts);
        }
    }
    
    export class UserPoolDomain extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<UserPoolDomainAttributes>;

        constructor(name: string, properties: UserPoolDomainProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Cognito:UserPoolDomain", name, inputs, opts);
        }
    }
    
    export class UserPoolGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<UserPoolGroupAttributes>;

        constructor(name: string, properties: UserPoolGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Cognito:UserPoolGroup", name, inputs, opts);
        }
    }
    
    export class UserPoolIdentityProvider extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<UserPoolIdentityProviderAttributes>;

        constructor(name: string, properties: UserPoolIdentityProviderProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Cognito:UserPoolIdentityProvider", name, inputs, opts);
        }
    }
    
    export class UserPoolResourceServer extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<UserPoolResourceServerAttributes>;

        constructor(name: string, properties: UserPoolResourceServerProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Cognito:UserPoolResourceServer", name, inputs, opts);
        }
    }
    
    export class UserPoolRiskConfigurationAttachment extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<UserPoolRiskConfigurationAttachmentAttributes>;

        constructor(name: string, properties: UserPoolRiskConfigurationAttachmentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Cognito:UserPoolRiskConfigurationAttachment", name, inputs, opts);
        }
    }
    
    export class UserPoolUICustomizationAttachment extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<UserPoolUICustomizationAttachmentAttributes>;

        constructor(name: string, properties: UserPoolUICustomizationAttachmentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Cognito:UserPoolUICustomizationAttachment", name, inputs, opts);
        }
    }
    
    export class UserPoolUser extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<UserPoolUserAttributes>;

        constructor(name: string, properties: UserPoolUserProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Cognito:UserPoolUser", name, inputs, opts);
        }
    }
    
    export class UserPoolUserToGroupAttachment extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<UserPoolUserToGroupAttachmentAttributes>;

        constructor(name: string, properties: UserPoolUserToGroupAttachmentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Cognito:UserPoolUserToGroupAttachment", name, inputs, opts);
        }
    }
    
}

export namespace config {
    
    export interface AggregationAuthorizationAttributes {
    }
    
    export interface AggregationAuthorizationProperties {
        AuthorizedAccountId: pulumi.Input<string>;
        AuthorizedAwsRegion: pulumi.Input<string>;
    }
    
    export interface ConfigRuleAttributes {
        Arn: string;
        "Compliance.Type": string;
        ConfigRuleId: string;
    }
    
    export interface ConfigRuleProperties {
        ConfigRuleName?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        InputParameters?: pulumi.Input<string>;
        MaximumExecutionFrequency?: pulumi.Input<string>;
        Scope?: pulumi.Input<ConfigRuleScopeProperties>;
        Source: pulumi.Input<ConfigRuleSourceProperties>;
    }
    
    export interface ConfigRuleScopeAttributes {
        ComplianceResourceId?: string;
        ComplianceResourceTypes?: string[];
        TagKey?: string;
        TagValue?: string;
    }
    
    export interface ConfigRuleScopeProperties {
        ComplianceResourceId?: pulumi.Input<string>;
        ComplianceResourceTypes?: pulumi.Input<pulumi.Input<string>[]>;
        TagKey?: pulumi.Input<string>;
        TagValue?: pulumi.Input<string>;
    }
    
    export interface ConfigRuleSourceAttributes {
        Owner: string;
        SourceDetails?: ConfigRuleSourceDetailAttributes[];
        SourceIdentifier: string;
    }
    
    export interface ConfigRuleSourceDetailAttributes {
        EventSource: string;
        MaximumExecutionFrequency?: string;
        MessageType: string;
    }
    
    export interface ConfigRuleSourceDetailProperties {
        EventSource: pulumi.Input<string>;
        MaximumExecutionFrequency?: pulumi.Input<string>;
        MessageType: pulumi.Input<string>;
    }
    
    export interface ConfigRuleSourceProperties {
        Owner: pulumi.Input<string>;
        SourceDetails?: pulumi.Input<pulumi.Input<ConfigRuleSourceDetailProperties>[]>;
        SourceIdentifier: pulumi.Input<string>;
    }
    
    export interface ConfigurationAggregatorAccountAggregationSourceAttributes {
        AccountIds: string[];
        AllAwsRegions?: boolean;
        AwsRegions?: string[];
    }
    
    export interface ConfigurationAggregatorAccountAggregationSourceProperties {
        AccountIds: pulumi.Input<pulumi.Input<string>[]>;
        AllAwsRegions?: pulumi.Input<boolean>;
        AwsRegions?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface ConfigurationAggregatorAttributes {
    }
    
    export interface ConfigurationAggregatorOrganizationAggregationSourceAttributes {
        AllAwsRegions?: boolean;
        AwsRegions?: string[];
        RoleArn: string;
    }
    
    export interface ConfigurationAggregatorOrganizationAggregationSourceProperties {
        AllAwsRegions?: pulumi.Input<boolean>;
        AwsRegions?: pulumi.Input<pulumi.Input<string>[]>;
        RoleArn: pulumi.Input<string>;
    }
    
    export interface ConfigurationAggregatorProperties {
        AccountAggregationSources?: pulumi.Input<pulumi.Input<ConfigurationAggregatorAccountAggregationSourceProperties>[]>;
        ConfigurationAggregatorName: pulumi.Input<string>;
        OrganizationAggregationSource?: pulumi.Input<ConfigurationAggregatorOrganizationAggregationSourceProperties>;
    }
    
    export interface ConfigurationRecorderAttributes {
    }
    
    export interface ConfigurationRecorderProperties {
        Name?: pulumi.Input<string>;
        RecordingGroup?: pulumi.Input<ConfigurationRecorderRecordingGroupProperties>;
        RoleARN: pulumi.Input<string>;
    }
    
    export interface ConfigurationRecorderRecordingGroupAttributes {
        AllSupported?: boolean;
        IncludeGlobalResourceTypes?: boolean;
        ResourceTypes?: string[];
    }
    
    export interface ConfigurationRecorderRecordingGroupProperties {
        AllSupported?: pulumi.Input<boolean>;
        IncludeGlobalResourceTypes?: pulumi.Input<boolean>;
        ResourceTypes?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface DeliveryChannelAttributes {
    }
    
    export interface DeliveryChannelConfigSnapshotDeliveryPropertiesAttributes {
        DeliveryFrequency?: string;
    }
    
    export interface DeliveryChannelConfigSnapshotDeliveryPropertiesProperties {
        DeliveryFrequency?: pulumi.Input<string>;
    }
    
    export interface DeliveryChannelProperties {
        ConfigSnapshotDeliveryProperties?: pulumi.Input<DeliveryChannelConfigSnapshotDeliveryPropertiesProperties>;
        Name?: pulumi.Input<string>;
        S3BucketName: pulumi.Input<string>;
        S3KeyPrefix?: pulumi.Input<string>;
        SnsTopicARN?: pulumi.Input<string>;
    }
    
    export interface OrganizationConfigRuleAttributes {
    }
    
    export interface OrganizationConfigRuleOrganizationCustomRuleMetadataAttributes {
        Description?: string;
        InputParameters?: string;
        LambdaFunctionArn: string;
        MaximumExecutionFrequency?: string;
        OrganizationConfigRuleTriggerTypes: string[];
        ResourceIdScope?: string;
        ResourceTypesScope?: string[];
        TagKeyScope?: string;
        TagValueScope?: string;
    }
    
    export interface OrganizationConfigRuleOrganizationCustomRuleMetadataProperties {
        Description?: pulumi.Input<string>;
        InputParameters?: pulumi.Input<string>;
        LambdaFunctionArn: pulumi.Input<string>;
        MaximumExecutionFrequency?: pulumi.Input<string>;
        OrganizationConfigRuleTriggerTypes: pulumi.Input<pulumi.Input<string>[]>;
        ResourceIdScope?: pulumi.Input<string>;
        ResourceTypesScope?: pulumi.Input<pulumi.Input<string>[]>;
        TagKeyScope?: pulumi.Input<string>;
        TagValueScope?: pulumi.Input<string>;
    }
    
    export interface OrganizationConfigRuleOrganizationManagedRuleMetadataAttributes {
        Description?: string;
        InputParameters?: string;
        MaximumExecutionFrequency?: string;
        ResourceIdScope?: string;
        ResourceTypesScope?: string[];
        RuleIdentifier: string;
        TagKeyScope?: string;
        TagValueScope?: string;
    }
    
    export interface OrganizationConfigRuleOrganizationManagedRuleMetadataProperties {
        Description?: pulumi.Input<string>;
        InputParameters?: pulumi.Input<string>;
        MaximumExecutionFrequency?: pulumi.Input<string>;
        ResourceIdScope?: pulumi.Input<string>;
        ResourceTypesScope?: pulumi.Input<pulumi.Input<string>[]>;
        RuleIdentifier: pulumi.Input<string>;
        TagKeyScope?: pulumi.Input<string>;
        TagValueScope?: pulumi.Input<string>;
    }
    
    export interface OrganizationConfigRuleProperties {
        ExcludedAccounts?: pulumi.Input<pulumi.Input<string>[]>;
        OrganizationConfigRuleName: pulumi.Input<string>;
        OrganizationCustomRuleMetadata?: pulumi.Input<OrganizationConfigRuleOrganizationCustomRuleMetadataProperties>;
        OrganizationManagedRuleMetadata?: pulumi.Input<OrganizationConfigRuleOrganizationManagedRuleMetadataProperties>;
    }
    
    export interface RemediationConfigurationAttributes {
    }
    
    export interface RemediationConfigurationExecutionControlsAttributes {
        SsmControls?: RemediationConfigurationSsmControlsAttributes;
    }
    
    export interface RemediationConfigurationExecutionControlsProperties {
        SsmControls?: pulumi.Input<RemediationConfigurationSsmControlsProperties>;
    }
    
    export interface RemediationConfigurationProperties {
        Automatic?: pulumi.Input<boolean>;
        ConfigRuleName: pulumi.Input<string>;
        ExecutionControls?: pulumi.Input<RemediationConfigurationExecutionControlsProperties>;
        MaximumAutomaticAttempts?: pulumi.Input<number>;
        Parameters?: pulumi.Input<string>;
        ResourceType?: pulumi.Input<string>;
        RetryAttemptSeconds?: pulumi.Input<number>;
        TargetId: pulumi.Input<string>;
        TargetType: pulumi.Input<string>;
        TargetVersion?: pulumi.Input<string>;
    }
    
    export interface RemediationConfigurationRemediationParameterValueAttributes {
        ResourceValue?: RemediationConfigurationResourceValueAttributes;
        StaticValue?: RemediationConfigurationStaticValueAttributes;
    }
    
    export interface RemediationConfigurationRemediationParameterValueProperties {
        ResourceValue?: pulumi.Input<RemediationConfigurationResourceValueProperties>;
        StaticValue?: pulumi.Input<RemediationConfigurationStaticValueProperties>;
    }
    
    export interface RemediationConfigurationResourceValueAttributes {
        Value?: string;
    }
    
    export interface RemediationConfigurationResourceValueProperties {
        Value?: pulumi.Input<string>;
    }
    
    export interface RemediationConfigurationSsmControlsAttributes {
        ConcurrentExecutionRatePercentage?: number;
        ErrorPercentage?: number;
    }
    
    export interface RemediationConfigurationSsmControlsProperties {
        ConcurrentExecutionRatePercentage?: pulumi.Input<number>;
        ErrorPercentage?: pulumi.Input<number>;
    }
    
    export interface RemediationConfigurationStaticValueAttributes {
        Values?: string[];
    }
    
    export interface RemediationConfigurationStaticValueProperties {
        Values?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    
    export class AggregationAuthorization extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AggregationAuthorizationAttributes>;

        constructor(name: string, properties: AggregationAuthorizationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Config:AggregationAuthorization", name, inputs, opts);
        }
    }
    
    export class ConfigRule extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ConfigRuleAttributes>;

        constructor(name: string, properties: ConfigRuleProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Config:ConfigRule", name, inputs, opts);
        }
    }
    
    export class ConfigurationAggregator extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ConfigurationAggregatorAttributes>;

        constructor(name: string, properties: ConfigurationAggregatorProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Config:ConfigurationAggregator", name, inputs, opts);
        }
    }
    
    export class ConfigurationRecorder extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ConfigurationRecorderAttributes>;

        constructor(name: string, properties: ConfigurationRecorderProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Config:ConfigurationRecorder", name, inputs, opts);
        }
    }
    
    export class DeliveryChannel extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DeliveryChannelAttributes>;

        constructor(name: string, properties: DeliveryChannelProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Config:DeliveryChannel", name, inputs, opts);
        }
    }
    
    export class OrganizationConfigRule extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<OrganizationConfigRuleAttributes>;

        constructor(name: string, properties: OrganizationConfigRuleProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Config:OrganizationConfigRule", name, inputs, opts);
        }
    }
    
    export class RemediationConfiguration extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RemediationConfigurationAttributes>;

        constructor(name: string, properties: RemediationConfigurationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Config:RemediationConfiguration", name, inputs, opts);
        }
    }
    
}

export namespace dax {
    
    export interface ClusterAttributes {
        Arn: string;
        ClusterDiscoveryEndpoint: string;
    }
    
    export interface ClusterProperties {
        AvailabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
        ClusterName?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        IAMRoleARN: pulumi.Input<string>;
        NodeType: pulumi.Input<string>;
        NotificationTopicARN?: pulumi.Input<string>;
        ParameterGroupName?: pulumi.Input<string>;
        PreferredMaintenanceWindow?: pulumi.Input<string>;
        ReplicationFactor: pulumi.Input<number>;
        SSESpecification?: pulumi.Input<ClusterSSESpecificationProperties>;
        SecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
        SubnetGroupName?: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface ClusterSSESpecificationAttributes {
        SSEEnabled?: boolean;
    }
    
    export interface ClusterSSESpecificationProperties {
        SSEEnabled?: pulumi.Input<boolean>;
    }
    
    export interface ParameterGroupAttributes {
    }
    
    export interface ParameterGroupProperties {
        Description?: pulumi.Input<string>;
        ParameterGroupName?: pulumi.Input<string>;
        ParameterNameValues?: pulumi.Input<string>;
    }
    
    export interface SubnetGroupAttributes {
    }
    
    export interface SubnetGroupProperties {
        Description?: pulumi.Input<string>;
        SubnetGroupName?: pulumi.Input<string>;
        SubnetIds: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    
    export class Cluster extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ClusterAttributes>;

        constructor(name: string, properties: ClusterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DAX:Cluster", name, inputs, opts);
        }
    }
    
    export class ParameterGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ParameterGroupAttributes>;

        constructor(name: string, properties?: ParameterGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DAX:ParameterGroup", name, inputs, opts);
        }
    }
    
    export class SubnetGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SubnetGroupAttributes>;

        constructor(name: string, properties: SubnetGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DAX:SubnetGroup", name, inputs, opts);
        }
    }
    
}

export namespace dlm {
    
    export interface LifecyclePolicyAttributes {
        Arn: string;
    }
    
    export interface LifecyclePolicyCreateRuleAttributes {
        Interval: number;
        IntervalUnit: string;
        Times?: string[];
    }
    
    export interface LifecyclePolicyCreateRuleProperties {
        Interval: pulumi.Input<number>;
        IntervalUnit: pulumi.Input<string>;
        Times?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface LifecyclePolicyParametersAttributes {
        ExcludeBootVolume?: boolean;
    }
    
    export interface LifecyclePolicyParametersProperties {
        ExcludeBootVolume?: pulumi.Input<boolean>;
    }
    
    export interface LifecyclePolicyPolicyDetailsAttributes {
        Parameters?: LifecyclePolicyParametersAttributes;
        PolicyType?: string;
        ResourceTypes?: string[];
        Schedules?: LifecyclePolicyScheduleAttributes[];
        TargetTags?: TagAttributes[];
    }
    
    export interface LifecyclePolicyPolicyDetailsProperties {
        Parameters?: pulumi.Input<LifecyclePolicyParametersProperties>;
        PolicyType?: pulumi.Input<string>;
        ResourceTypes?: pulumi.Input<pulumi.Input<string>[]>;
        Schedules?: pulumi.Input<pulumi.Input<LifecyclePolicyScheduleProperties>[]>;
        TargetTags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface LifecyclePolicyProperties {
        Description?: pulumi.Input<string>;
        ExecutionRoleArn?: pulumi.Input<string>;
        PolicyDetails?: pulumi.Input<LifecyclePolicyPolicyDetailsProperties>;
        State?: pulumi.Input<string>;
    }
    
    export interface LifecyclePolicyRetainRuleAttributes {
        Count: number;
    }
    
    export interface LifecyclePolicyRetainRuleProperties {
        Count: pulumi.Input<number>;
    }
    
    export interface LifecyclePolicyScheduleAttributes {
        CopyTags?: boolean;
        CreateRule?: LifecyclePolicyCreateRuleAttributes;
        Name?: string;
        RetainRule?: LifecyclePolicyRetainRuleAttributes;
        TagsToAdd?: TagAttributes[];
        VariableTags?: TagAttributes[];
    }
    
    export interface LifecyclePolicyScheduleProperties {
        CopyTags?: pulumi.Input<boolean>;
        CreateRule?: pulumi.Input<LifecyclePolicyCreateRuleProperties>;
        Name?: pulumi.Input<string>;
        RetainRule?: pulumi.Input<LifecyclePolicyRetainRuleProperties>;
        TagsToAdd?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VariableTags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    
    export class LifecyclePolicy extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<LifecyclePolicyAttributes>;

        constructor(name: string, properties?: LifecyclePolicyProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DLM:LifecyclePolicy", name, inputs, opts);
        }
    }
    
}

export namespace dms {
    
    export interface CertificateAttributes {
    }
    
    export interface CertificateProperties {
        CertificateIdentifier?: pulumi.Input<string>;
        CertificatePem?: pulumi.Input<string>;
        CertificateWallet?: pulumi.Input<string>;
    }
    
    export interface EndpointAttributes {
        ExternalId: string;
    }
    
    export interface EndpointDynamoDbSettingsAttributes {
        ServiceAccessRoleArn?: string;
    }
    
    export interface EndpointDynamoDbSettingsProperties {
        ServiceAccessRoleArn?: pulumi.Input<string>;
    }
    
    export interface EndpointElasticsearchSettingsAttributes {
        EndpointUri?: string;
        ErrorRetryDuration?: number;
        FullLoadErrorPercentage?: number;
        ServiceAccessRoleArn?: string;
    }
    
    export interface EndpointElasticsearchSettingsProperties {
        EndpointUri?: pulumi.Input<string>;
        ErrorRetryDuration?: pulumi.Input<number>;
        FullLoadErrorPercentage?: pulumi.Input<number>;
        ServiceAccessRoleArn?: pulumi.Input<string>;
    }
    
    export interface EndpointKinesisSettingsAttributes {
        MessageFormat?: string;
        ServiceAccessRoleArn?: string;
        StreamArn?: string;
    }
    
    export interface EndpointKinesisSettingsProperties {
        MessageFormat?: pulumi.Input<string>;
        ServiceAccessRoleArn?: pulumi.Input<string>;
        StreamArn?: pulumi.Input<string>;
    }
    
    export interface EndpointMongoDbSettingsAttributes {
        AuthMechanism?: string;
        AuthSource?: string;
        AuthType?: string;
        DatabaseName?: string;
        DocsToInvestigate?: string;
        ExtractDocId?: string;
        NestingLevel?: string;
        Password?: string;
        Port?: number;
        ServerName?: string;
        Username?: string;
    }
    
    export interface EndpointMongoDbSettingsProperties {
        AuthMechanism?: pulumi.Input<string>;
        AuthSource?: pulumi.Input<string>;
        AuthType?: pulumi.Input<string>;
        DatabaseName?: pulumi.Input<string>;
        DocsToInvestigate?: pulumi.Input<string>;
        ExtractDocId?: pulumi.Input<string>;
        NestingLevel?: pulumi.Input<string>;
        Password?: pulumi.Input<string>;
        Port?: pulumi.Input<number>;
        ServerName?: pulumi.Input<string>;
        Username?: pulumi.Input<string>;
    }
    
    export interface EndpointProperties {
        CertificateArn?: pulumi.Input<string>;
        DatabaseName?: pulumi.Input<string>;
        DynamoDbSettings?: pulumi.Input<EndpointDynamoDbSettingsProperties>;
        ElasticsearchSettings?: pulumi.Input<EndpointElasticsearchSettingsProperties>;
        EndpointIdentifier?: pulumi.Input<string>;
        EndpointType: pulumi.Input<string>;
        EngineName: pulumi.Input<string>;
        ExtraConnectionAttributes?: pulumi.Input<string>;
        KinesisSettings?: pulumi.Input<EndpointKinesisSettingsProperties>;
        KmsKeyId?: pulumi.Input<string>;
        MongoDbSettings?: pulumi.Input<EndpointMongoDbSettingsProperties>;
        Password?: pulumi.Input<string>;
        Port?: pulumi.Input<number>;
        S3Settings?: pulumi.Input<EndpointS3SettingsProperties>;
        ServerName?: pulumi.Input<string>;
        SslMode?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        Username?: pulumi.Input<string>;
    }
    
    export interface EndpointS3SettingsAttributes {
        BucketFolder?: string;
        BucketName?: string;
        CompressionType?: string;
        CsvDelimiter?: string;
        CsvRowDelimiter?: string;
        ExternalTableDefinition?: string;
        ServiceAccessRoleArn?: string;
    }
    
    export interface EndpointS3SettingsProperties {
        BucketFolder?: pulumi.Input<string>;
        BucketName?: pulumi.Input<string>;
        CompressionType?: pulumi.Input<string>;
        CsvDelimiter?: pulumi.Input<string>;
        CsvRowDelimiter?: pulumi.Input<string>;
        ExternalTableDefinition?: pulumi.Input<string>;
        ServiceAccessRoleArn?: pulumi.Input<string>;
    }
    
    export interface EventSubscriptionAttributes {
    }
    
    export interface EventSubscriptionProperties {
        Enabled?: pulumi.Input<boolean>;
        EventCategories?: pulumi.Input<pulumi.Input<string>[]>;
        SnsTopicArn: pulumi.Input<string>;
        SourceIds?: pulumi.Input<pulumi.Input<string>[]>;
        SourceType?: pulumi.Input<string>;
        SubscriptionName?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface ReplicationInstanceAttributes {
        ReplicationInstancePrivateIpAddresses: string[];
        ReplicationInstancePublicIpAddresses: string[];
    }
    
    export interface ReplicationInstanceProperties {
        AllocatedStorage?: pulumi.Input<number>;
        AllowMajorVersionUpgrade?: pulumi.Input<boolean>;
        AutoMinorVersionUpgrade?: pulumi.Input<boolean>;
        AvailabilityZone?: pulumi.Input<string>;
        EngineVersion?: pulumi.Input<string>;
        KmsKeyId?: pulumi.Input<string>;
        MultiAZ?: pulumi.Input<boolean>;
        PreferredMaintenanceWindow?: pulumi.Input<string>;
        PubliclyAccessible?: pulumi.Input<boolean>;
        ReplicationInstanceClass: pulumi.Input<string>;
        ReplicationInstanceIdentifier?: pulumi.Input<string>;
        ReplicationSubnetGroupIdentifier?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface ReplicationSubnetGroupAttributes {
    }
    
    export interface ReplicationSubnetGroupProperties {
        ReplicationSubnetGroupDescription: pulumi.Input<string>;
        ReplicationSubnetGroupIdentifier?: pulumi.Input<string>;
        SubnetIds: pulumi.Input<pulumi.Input<string>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface ReplicationTaskAttributes {
    }
    
    export interface ReplicationTaskProperties {
        CdcStartPosition?: pulumi.Input<string>;
        CdcStartTime?: pulumi.Input<number>;
        CdcStopPosition?: pulumi.Input<string>;
        MigrationType: pulumi.Input<string>;
        ReplicationInstanceArn: pulumi.Input<string>;
        ReplicationTaskIdentifier?: pulumi.Input<string>;
        ReplicationTaskSettings?: pulumi.Input<string>;
        SourceEndpointArn: pulumi.Input<string>;
        TableMappings: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        TargetEndpointArn: pulumi.Input<string>;
    }
    
    
    export class Certificate extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<CertificateAttributes>;

        constructor(name: string, properties?: CertificateProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DMS:Certificate", name, inputs, opts);
        }
    }
    
    export class Endpoint extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<EndpointAttributes>;

        constructor(name: string, properties: EndpointProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DMS:Endpoint", name, inputs, opts);
        }
    }
    
    export class EventSubscription extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<EventSubscriptionAttributes>;

        constructor(name: string, properties: EventSubscriptionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DMS:EventSubscription", name, inputs, opts);
        }
    }
    
    export class ReplicationInstance extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ReplicationInstanceAttributes>;

        constructor(name: string, properties: ReplicationInstanceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DMS:ReplicationInstance", name, inputs, opts);
        }
    }
    
    export class ReplicationSubnetGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ReplicationSubnetGroupAttributes>;

        constructor(name: string, properties: ReplicationSubnetGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DMS:ReplicationSubnetGroup", name, inputs, opts);
        }
    }
    
    export class ReplicationTask extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ReplicationTaskAttributes>;

        constructor(name: string, properties: ReplicationTaskProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DMS:ReplicationTask", name, inputs, opts);
        }
    }
    
}

export namespace datapipeline {
    
    export interface PipelineAttributes {
    }
    
    export interface PipelineFieldAttributes {
        Key: string;
        RefValue?: string;
        StringValue?: string;
    }
    
    export interface PipelineFieldProperties {
        Key: pulumi.Input<string>;
        RefValue?: pulumi.Input<string>;
        StringValue?: pulumi.Input<string>;
    }
    
    export interface PipelineParameterAttributeAttributes {
        Key: string;
        StringValue: string;
    }
    
    export interface PipelineParameterAttributeProperties {
        Key: pulumi.Input<string>;
        StringValue: pulumi.Input<string>;
    }
    
    export interface PipelineParameterObjectAttributes {
        Attributes: PipelineParameterAttributeAttributes[];
        Id: string;
    }
    
    export interface PipelineParameterObjectProperties {
        Attributes: pulumi.Input<pulumi.Input<PipelineParameterAttributeProperties>[]>;
        Id: pulumi.Input<string>;
    }
    
    export interface PipelineParameterValueAttributes {
        Id: string;
        StringValue: string;
    }
    
    export interface PipelineParameterValueProperties {
        Id: pulumi.Input<string>;
        StringValue: pulumi.Input<string>;
    }
    
    export interface PipelinePipelineObjectAttributes {
        Fields: PipelineFieldAttributes[];
        Id: string;
        Name: string;
    }
    
    export interface PipelinePipelineObjectProperties {
        Fields: pulumi.Input<pulumi.Input<PipelineFieldProperties>[]>;
        Id: pulumi.Input<string>;
        Name: pulumi.Input<string>;
    }
    
    export interface PipelinePipelineTagAttributes {
        Key: string;
        Value: string;
    }
    
    export interface PipelinePipelineTagProperties {
        Key: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface PipelineProperties {
        Activate?: pulumi.Input<boolean>;
        Description?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        ParameterObjects: pulumi.Input<pulumi.Input<PipelineParameterObjectProperties>[]>;
        ParameterValues?: pulumi.Input<pulumi.Input<PipelineParameterValueProperties>[]>;
        PipelineObjects?: pulumi.Input<pulumi.Input<PipelinePipelineObjectProperties>[]>;
        PipelineTags?: pulumi.Input<pulumi.Input<PipelinePipelineTagProperties>[]>;
    }
    
    
    export class Pipeline extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PipelineAttributes>;

        constructor(name: string, properties: PipelineProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DataPipeline:Pipeline", name, inputs, opts);
        }
    }
    
}

export namespace directoryservice {
    
    export interface MicrosoftADAttributes {
        Alias: string;
        DnsIpAddresses: string[];
    }
    
    export interface MicrosoftADProperties {
        CreateAlias?: pulumi.Input<boolean>;
        Edition?: pulumi.Input<string>;
        EnableSso?: pulumi.Input<boolean>;
        Name: pulumi.Input<string>;
        Password: pulumi.Input<string>;
        ShortName?: pulumi.Input<string>;
        VpcSettings: pulumi.Input<MicrosoftADVpcSettingsProperties>;
    }
    
    export interface MicrosoftADVpcSettingsAttributes {
        SubnetIds: string[];
        VpcId: string;
    }
    
    export interface MicrosoftADVpcSettingsProperties {
        SubnetIds: pulumi.Input<pulumi.Input<string>[]>;
        VpcId: pulumi.Input<string>;
    }
    
    export interface SimpleADAttributes {
        Alias: string;
        DnsIpAddresses: string[];
    }
    
    export interface SimpleADProperties {
        CreateAlias?: pulumi.Input<boolean>;
        Description?: pulumi.Input<string>;
        EnableSso?: pulumi.Input<boolean>;
        Name: pulumi.Input<string>;
        Password: pulumi.Input<string>;
        ShortName?: pulumi.Input<string>;
        Size: pulumi.Input<string>;
        VpcSettings: pulumi.Input<SimpleADVpcSettingsProperties>;
    }
    
    export interface SimpleADVpcSettingsAttributes {
        SubnetIds: string[];
        VpcId: string;
    }
    
    export interface SimpleADVpcSettingsProperties {
        SubnetIds: pulumi.Input<pulumi.Input<string>[]>;
        VpcId: pulumi.Input<string>;
    }
    
    
    export class MicrosoftAD extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<MicrosoftADAttributes>;

        constructor(name: string, properties: MicrosoftADProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DirectoryService:MicrosoftAD", name, inputs, opts);
        }
    }
    
    export class SimpleAD extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SimpleADAttributes>;

        constructor(name: string, properties: SimpleADProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DirectoryService:SimpleAD", name, inputs, opts);
        }
    }
    
}

export namespace docdb {
    
    export interface DBClusterAttributes {
        ClusterResourceId: string;
        Endpoint: string;
        Port: string;
        ReadEndpoint: string;
    }
    
    export interface DBClusterParameterGroupAttributes {
    }
    
    export interface DBClusterParameterGroupProperties {
        Description: pulumi.Input<string>;
        Family: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Parameters: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface DBClusterProperties {
        AvailabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
        BackupRetentionPeriod?: pulumi.Input<number>;
        DBClusterIdentifier?: pulumi.Input<string>;
        DBClusterParameterGroupName?: pulumi.Input<string>;
        DBSubnetGroupName?: pulumi.Input<string>;
        EnableCloudwatchLogsExports?: pulumi.Input<pulumi.Input<string>[]>;
        EngineVersion?: pulumi.Input<string>;
        KmsKeyId?: pulumi.Input<string>;
        MasterUserPassword?: pulumi.Input<string>;
        MasterUsername?: pulumi.Input<string>;
        Port?: pulumi.Input<number>;
        PreferredBackupWindow?: pulumi.Input<string>;
        PreferredMaintenanceWindow?: pulumi.Input<string>;
        SnapshotIdentifier?: pulumi.Input<string>;
        StorageEncrypted?: pulumi.Input<boolean>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface DBInstanceAttributes {
        Endpoint: string;
        Port: string;
    }
    
    export interface DBInstanceProperties {
        AutoMinorVersionUpgrade?: pulumi.Input<boolean>;
        AvailabilityZone?: pulumi.Input<string>;
        DBClusterIdentifier: pulumi.Input<string>;
        DBInstanceClass: pulumi.Input<string>;
        DBInstanceIdentifier?: pulumi.Input<string>;
        PreferredMaintenanceWindow?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface DBSubnetGroupAttributes {
    }
    
    export interface DBSubnetGroupProperties {
        DBSubnetGroupDescription: pulumi.Input<string>;
        DBSubnetGroupName?: pulumi.Input<string>;
        SubnetIds: pulumi.Input<pulumi.Input<string>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    
    export class DBCluster extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DBClusterAttributes>;

        constructor(name: string, properties?: DBClusterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DocDB:DBCluster", name, inputs, opts);
        }
    }
    
    export class DBClusterParameterGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DBClusterParameterGroupAttributes>;

        constructor(name: string, properties: DBClusterParameterGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DocDB:DBClusterParameterGroup", name, inputs, opts);
        }
    }
    
    export class DBInstance extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DBInstanceAttributes>;

        constructor(name: string, properties: DBInstanceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DocDB:DBInstance", name, inputs, opts);
        }
    }
    
    export class DBSubnetGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DBSubnetGroupAttributes>;

        constructor(name: string, properties: DBSubnetGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DocDB:DBSubnetGroup", name, inputs, opts);
        }
    }
    
}

export namespace dynamodb {
    
    export interface TableAttributeDefinitionAttributes {
        AttributeName: string;
        AttributeType: string;
    }
    
    export interface TableAttributeDefinitionProperties {
        AttributeName: pulumi.Input<string>;
        AttributeType: pulumi.Input<string>;
    }
    
    export interface TableAttributes {
        Arn: string;
        StreamArn: string;
    }
    
    export interface TableGlobalSecondaryIndexAttributes {
        IndexName: string;
        KeySchema: TableKeySchemaAttributes[];
        Projection: TableProjectionAttributes;
        ProvisionedThroughput?: TableProvisionedThroughputAttributes;
    }
    
    export interface TableGlobalSecondaryIndexProperties {
        IndexName: pulumi.Input<string>;
        KeySchema: pulumi.Input<pulumi.Input<TableKeySchemaProperties>[]>;
        Projection: pulumi.Input<TableProjectionProperties>;
        ProvisionedThroughput?: pulumi.Input<TableProvisionedThroughputProperties>;
    }
    
    export interface TableKeySchemaAttributes {
        AttributeName: string;
        KeyType: string;
    }
    
    export interface TableKeySchemaProperties {
        AttributeName: pulumi.Input<string>;
        KeyType: pulumi.Input<string>;
    }
    
    export interface TableLocalSecondaryIndexAttributes {
        IndexName: string;
        KeySchema: TableKeySchemaAttributes[];
        Projection: TableProjectionAttributes;
    }
    
    export interface TableLocalSecondaryIndexProperties {
        IndexName: pulumi.Input<string>;
        KeySchema: pulumi.Input<pulumi.Input<TableKeySchemaProperties>[]>;
        Projection: pulumi.Input<TableProjectionProperties>;
    }
    
    export interface TablePointInTimeRecoverySpecificationAttributes {
        PointInTimeRecoveryEnabled?: boolean;
    }
    
    export interface TablePointInTimeRecoverySpecificationProperties {
        PointInTimeRecoveryEnabled?: pulumi.Input<boolean>;
    }
    
    export interface TableProjectionAttributes {
        NonKeyAttributes?: string[];
        ProjectionType?: string;
    }
    
    export interface TableProjectionProperties {
        NonKeyAttributes?: pulumi.Input<pulumi.Input<string>[]>;
        ProjectionType?: pulumi.Input<string>;
    }
    
    export interface TableProperties {
        AttributeDefinitions?: pulumi.Input<pulumi.Input<TableAttributeDefinitionProperties>[]>;
        BillingMode?: pulumi.Input<string>;
        GlobalSecondaryIndexes?: pulumi.Input<pulumi.Input<TableGlobalSecondaryIndexProperties>[]>;
        KeySchema: pulumi.Input<pulumi.Input<TableKeySchemaProperties>[]>;
        LocalSecondaryIndexes?: pulumi.Input<pulumi.Input<TableLocalSecondaryIndexProperties>[]>;
        PointInTimeRecoverySpecification?: pulumi.Input<TablePointInTimeRecoverySpecificationProperties>;
        ProvisionedThroughput?: pulumi.Input<TableProvisionedThroughputProperties>;
        SSESpecification?: pulumi.Input<TableSSESpecificationProperties>;
        StreamSpecification?: pulumi.Input<TableStreamSpecificationProperties>;
        TableName?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        TimeToLiveSpecification?: pulumi.Input<TableTimeToLiveSpecificationProperties>;
    }
    
    export interface TableProvisionedThroughputAttributes {
        ReadCapacityUnits: number;
        WriteCapacityUnits: number;
    }
    
    export interface TableProvisionedThroughputProperties {
        ReadCapacityUnits: pulumi.Input<number>;
        WriteCapacityUnits: pulumi.Input<number>;
    }
    
    export interface TableSSESpecificationAttributes {
        KMSMasterKeyId?: string;
        SSEEnabled: boolean;
        SSEType?: string;
    }
    
    export interface TableSSESpecificationProperties {
        KMSMasterKeyId?: pulumi.Input<string>;
        SSEEnabled: pulumi.Input<boolean>;
        SSEType?: pulumi.Input<string>;
    }
    
    export interface TableStreamSpecificationAttributes {
        StreamViewType: string;
    }
    
    export interface TableStreamSpecificationProperties {
        StreamViewType: pulumi.Input<string>;
    }
    
    export interface TableTimeToLiveSpecificationAttributes {
        AttributeName: string;
        Enabled: boolean;
    }
    
    export interface TableTimeToLiveSpecificationProperties {
        AttributeName: pulumi.Input<string>;
        Enabled: pulumi.Input<boolean>;
    }
    
    
    export class Table extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TableAttributes>;

        constructor(name: string, properties: TableProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:DynamoDB:Table", name, inputs, opts);
        }
    }
    
}

export namespace ec2 {
    
    export interface CapacityReservationAttributes {
        AvailabilityZone: string;
        AvailableInstanceCount: number;
        InstanceType: string;
        Tenancy: string;
        TotalInstanceCount: number;
    }
    
    export interface CapacityReservationProperties {
        AvailabilityZone: pulumi.Input<string>;
        EbsOptimized?: pulumi.Input<boolean>;
        EndDate?: pulumi.Input<string>;
        EndDateType?: pulumi.Input<string>;
        EphemeralStorage?: pulumi.Input<boolean>;
        InstanceCount: pulumi.Input<number>;
        InstanceMatchCriteria?: pulumi.Input<string>;
        InstancePlatform: pulumi.Input<string>;
        InstanceType: pulumi.Input<string>;
        TagSpecifications?: pulumi.Input<pulumi.Input<CapacityReservationTagSpecificationProperties>[]>;
        Tenancy?: pulumi.Input<string>;
    }
    
    export interface CapacityReservationTagSpecificationAttributes {
        ResourceType?: string;
        Tags?: TagAttributes[];
    }
    
    export interface CapacityReservationTagSpecificationProperties {
        ResourceType?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface ClientVpnAuthorizationRuleAttributes {
    }
    
    export interface ClientVpnAuthorizationRuleProperties {
        AccessGroupId?: pulumi.Input<string>;
        AuthorizeAllGroups?: pulumi.Input<boolean>;
        ClientVpnEndpointId: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        TargetNetworkCidr: pulumi.Input<string>;
    }
    
    export interface ClientVpnEndpointAttributes {
    }
    
    export interface ClientVpnEndpointCertificateAuthenticationRequestAttributes {
        ClientRootCertificateChainArn: string;
    }
    
    export interface ClientVpnEndpointCertificateAuthenticationRequestProperties {
        ClientRootCertificateChainArn: pulumi.Input<string>;
    }
    
    export interface ClientVpnEndpointClientAuthenticationRequestAttributes {
        ActiveDirectory?: ClientVpnEndpointDirectoryServiceAuthenticationRequestAttributes;
        MutualAuthentication?: ClientVpnEndpointCertificateAuthenticationRequestAttributes;
        Type: string;
    }
    
    export interface ClientVpnEndpointClientAuthenticationRequestProperties {
        ActiveDirectory?: pulumi.Input<ClientVpnEndpointDirectoryServiceAuthenticationRequestProperties>;
        MutualAuthentication?: pulumi.Input<ClientVpnEndpointCertificateAuthenticationRequestProperties>;
        Type: pulumi.Input<string>;
    }
    
    export interface ClientVpnEndpointConnectionLogOptionsAttributes {
        CloudwatchLogGroup?: string;
        CloudwatchLogStream?: string;
        Enabled: boolean;
    }
    
    export interface ClientVpnEndpointConnectionLogOptionsProperties {
        CloudwatchLogGroup?: pulumi.Input<string>;
        CloudwatchLogStream?: pulumi.Input<string>;
        Enabled: pulumi.Input<boolean>;
    }
    
    export interface ClientVpnEndpointDirectoryServiceAuthenticationRequestAttributes {
        DirectoryId: string;
    }
    
    export interface ClientVpnEndpointDirectoryServiceAuthenticationRequestProperties {
        DirectoryId: pulumi.Input<string>;
    }
    
    export interface ClientVpnEndpointProperties {
        AuthenticationOptions: pulumi.Input<pulumi.Input<ClientVpnEndpointClientAuthenticationRequestProperties>[]>;
        ClientCidrBlock: pulumi.Input<string>;
        ConnectionLogOptions: pulumi.Input<ClientVpnEndpointConnectionLogOptionsProperties>;
        Description?: pulumi.Input<string>;
        DnsServers?: pulumi.Input<pulumi.Input<string>[]>;
        ServerCertificateArn: pulumi.Input<string>;
        SplitTunnel?: pulumi.Input<boolean>;
        TagSpecifications?: pulumi.Input<pulumi.Input<ClientVpnEndpointTagSpecificationProperties>[]>;
        TransportProtocol?: pulumi.Input<string>;
    }
    
    export interface ClientVpnEndpointTagSpecificationAttributes {
        ResourceType: string;
        Tags: TagAttributes[];
    }
    
    export interface ClientVpnEndpointTagSpecificationProperties {
        ResourceType: pulumi.Input<string>;
        Tags: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface ClientVpnRouteAttributes {
    }
    
    export interface ClientVpnRouteProperties {
        ClientVpnEndpointId: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        DestinationCidrBlock: pulumi.Input<string>;
        TargetVpcSubnetId: pulumi.Input<string>;
    }
    
    export interface ClientVpnTargetNetworkAssociationAttributes {
    }
    
    export interface ClientVpnTargetNetworkAssociationProperties {
        ClientVpnEndpointId: pulumi.Input<string>;
        SubnetId: pulumi.Input<string>;
    }
    
    export interface CustomerGatewayAttributes {
    }
    
    export interface CustomerGatewayProperties {
        BgpAsn: pulumi.Input<number>;
        IpAddress: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        Type: pulumi.Input<string>;
    }
    
    export interface DHCPOptionsAttributes {
    }
    
    export interface DHCPOptionsProperties {
        DomainName?: pulumi.Input<string>;
        DomainNameServers?: pulumi.Input<pulumi.Input<string>[]>;
        NetbiosNameServers?: pulumi.Input<pulumi.Input<string>[]>;
        NetbiosNodeType?: pulumi.Input<number>;
        NtpServers?: pulumi.Input<pulumi.Input<string>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface EC2FleetAttributes {
    }
    
    export interface EC2FleetFleetLaunchTemplateConfigRequestAttributes {
        LaunchTemplateSpecification?: EC2FleetFleetLaunchTemplateSpecificationRequestAttributes;
        Overrides?: EC2FleetFleetLaunchTemplateOverridesRequestAttributes[];
    }
    
    export interface EC2FleetFleetLaunchTemplateConfigRequestProperties {
        LaunchTemplateSpecification?: pulumi.Input<EC2FleetFleetLaunchTemplateSpecificationRequestProperties>;
        Overrides?: pulumi.Input<pulumi.Input<EC2FleetFleetLaunchTemplateOverridesRequestProperties>[]>;
    }
    
    export interface EC2FleetFleetLaunchTemplateOverridesRequestAttributes {
        AvailabilityZone?: string;
        InstanceType?: string;
        MaxPrice?: string;
        Priority?: number;
        SubnetId?: string;
        WeightedCapacity?: number;
    }
    
    export interface EC2FleetFleetLaunchTemplateOverridesRequestProperties {
        AvailabilityZone?: pulumi.Input<string>;
        InstanceType?: pulumi.Input<string>;
        MaxPrice?: pulumi.Input<string>;
        Priority?: pulumi.Input<number>;
        SubnetId?: pulumi.Input<string>;
        WeightedCapacity?: pulumi.Input<number>;
    }
    
    export interface EC2FleetFleetLaunchTemplateSpecificationRequestAttributes {
        LaunchTemplateId?: string;
        LaunchTemplateName?: string;
        Version?: string;
    }
    
    export interface EC2FleetFleetLaunchTemplateSpecificationRequestProperties {
        LaunchTemplateId?: pulumi.Input<string>;
        LaunchTemplateName?: pulumi.Input<string>;
        Version?: pulumi.Input<string>;
    }
    
    export interface EC2FleetOnDemandOptionsRequestAttributes {
        AllocationStrategy?: string;
    }
    
    export interface EC2FleetOnDemandOptionsRequestProperties {
        AllocationStrategy?: pulumi.Input<string>;
    }
    
    export interface EC2FleetProperties {
        ExcessCapacityTerminationPolicy?: pulumi.Input<string>;
        LaunchTemplateConfigs: pulumi.Input<pulumi.Input<EC2FleetFleetLaunchTemplateConfigRequestProperties>[]>;
        OnDemandOptions?: pulumi.Input<EC2FleetOnDemandOptionsRequestProperties>;
        ReplaceUnhealthyInstances?: pulumi.Input<boolean>;
        SpotOptions?: pulumi.Input<EC2FleetSpotOptionsRequestProperties>;
        TagSpecifications?: pulumi.Input<pulumi.Input<EC2FleetTagSpecificationProperties>[]>;
        TargetCapacitySpecification: pulumi.Input<EC2FleetTargetCapacitySpecificationRequestProperties>;
        TerminateInstancesWithExpiration?: pulumi.Input<boolean>;
        Type?: pulumi.Input<string>;
        ValidFrom?: pulumi.Input<string>;
        ValidUntil?: pulumi.Input<string>;
    }
    
    export interface EC2FleetSpotOptionsRequestAttributes {
        AllocationStrategy?: string;
        InstanceInterruptionBehavior?: string;
        InstancePoolsToUseCount?: number;
    }
    
    export interface EC2FleetSpotOptionsRequestProperties {
        AllocationStrategy?: pulumi.Input<string>;
        InstanceInterruptionBehavior?: pulumi.Input<string>;
        InstancePoolsToUseCount?: pulumi.Input<number>;
    }
    
    export interface EC2FleetTagRequestAttributes {
        Key?: string;
        Value?: string;
    }
    
    export interface EC2FleetTagRequestProperties {
        Key?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface EC2FleetTagSpecificationAttributes {
        ResourceType?: string;
        Tags?: EC2FleetTagRequestAttributes[];
    }
    
    export interface EC2FleetTagSpecificationProperties {
        ResourceType?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<EC2FleetTagRequestProperties>[]>;
    }
    
    export interface EC2FleetTargetCapacitySpecificationRequestAttributes {
        DefaultTargetCapacityType?: string;
        OnDemandTargetCapacity?: number;
        SpotTargetCapacity?: number;
        TotalTargetCapacity: number;
    }
    
    export interface EC2FleetTargetCapacitySpecificationRequestProperties {
        DefaultTargetCapacityType?: pulumi.Input<string>;
        OnDemandTargetCapacity?: pulumi.Input<number>;
        SpotTargetCapacity?: pulumi.Input<number>;
        TotalTargetCapacity: pulumi.Input<number>;
    }
    
    export interface EIPAssociationAttributes {
    }
    
    export interface EIPAssociationProperties {
        AllocationId?: pulumi.Input<string>;
        EIP?: pulumi.Input<string>;
        InstanceId?: pulumi.Input<string>;
        NetworkInterfaceId?: pulumi.Input<string>;
        PrivateIpAddress?: pulumi.Input<string>;
    }
    
    export interface EIPAttributes {
        AllocationId: string;
    }
    
    export interface EIPProperties {
        Domain?: pulumi.Input<string>;
        InstanceId?: pulumi.Input<string>;
        PublicIpv4Pool?: pulumi.Input<string>;
    }
    
    export interface EgressOnlyInternetGatewayAttributes {
    }
    
    export interface EgressOnlyInternetGatewayProperties {
        VpcId: pulumi.Input<string>;
    }
    
    export interface FlowLogAttributes {
    }
    
    export interface FlowLogProperties {
        DeliverLogsPermissionArn?: pulumi.Input<string>;
        LogDestination?: pulumi.Input<string>;
        LogDestinationType?: pulumi.Input<string>;
        LogGroupName?: pulumi.Input<string>;
        ResourceId: pulumi.Input<string>;
        ResourceType: pulumi.Input<string>;
        TrafficType: pulumi.Input<string>;
    }
    
    export interface HostAttributes {
    }
    
    export interface HostProperties {
        AutoPlacement?: pulumi.Input<string>;
        AvailabilityZone: pulumi.Input<string>;
        HostRecovery?: pulumi.Input<string>;
        InstanceType: pulumi.Input<string>;
    }
    
    export interface InstanceAssociationParameterAttributes {
        Key: string;
        Value: string[];
    }
    
    export interface InstanceAssociationParameterProperties {
        Key: pulumi.Input<string>;
        Value: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface InstanceAttributes {
        AvailabilityZone: string;
        PrivateDnsName: string;
        PrivateIp: string;
        PublicDnsName: string;
        PublicIp: string;
    }
    
    export interface InstanceBlockDeviceMappingAttributes {
        DeviceName: string;
        Ebs?: InstanceEbsAttributes;
        NoDevice?: InstanceNoDeviceAttributes;
        VirtualName?: string;
    }
    
    export interface InstanceBlockDeviceMappingProperties {
        DeviceName: pulumi.Input<string>;
        Ebs?: pulumi.Input<InstanceEbsProperties>;
        NoDevice?: pulumi.Input<InstanceNoDeviceProperties>;
        VirtualName?: pulumi.Input<string>;
    }
    
    export interface InstanceCpuOptionsAttributes {
        CoreCount?: number;
        ThreadsPerCore?: number;
    }
    
    export interface InstanceCpuOptionsProperties {
        CoreCount?: pulumi.Input<number>;
        ThreadsPerCore?: pulumi.Input<number>;
    }
    
    export interface InstanceCreditSpecificationAttributes {
        CPUCredits?: string;
    }
    
    export interface InstanceCreditSpecificationProperties {
        CPUCredits?: pulumi.Input<string>;
    }
    
    export interface InstanceEbsAttributes {
        DeleteOnTermination?: boolean;
        Encrypted?: boolean;
        Iops?: number;
        KmsKeyId?: string;
        SnapshotId?: string;
        VolumeSize?: number;
        VolumeType?: string;
    }
    
    export interface InstanceEbsProperties {
        DeleteOnTermination?: pulumi.Input<boolean>;
        Encrypted?: pulumi.Input<boolean>;
        Iops?: pulumi.Input<number>;
        KmsKeyId?: pulumi.Input<string>;
        SnapshotId?: pulumi.Input<string>;
        VolumeSize?: pulumi.Input<number>;
        VolumeType?: pulumi.Input<string>;
    }
    
    export interface InstanceElasticGpuSpecificationAttributes {
        Type: string;
    }
    
    export interface InstanceElasticGpuSpecificationProperties {
        Type: pulumi.Input<string>;
    }
    
    export interface InstanceElasticInferenceAcceleratorAttributes {
        Type: string;
    }
    
    export interface InstanceElasticInferenceAcceleratorProperties {
        Type: pulumi.Input<string>;
    }
    
    export interface InstanceInstanceIpv6AddressAttributes {
        Ipv6Address: string;
    }
    
    export interface InstanceInstanceIpv6AddressProperties {
        Ipv6Address: pulumi.Input<string>;
    }
    
    export interface InstanceLaunchTemplateSpecificationAttributes {
        LaunchTemplateId?: string;
        LaunchTemplateName?: string;
        Version: string;
    }
    
    export interface InstanceLaunchTemplateSpecificationProperties {
        LaunchTemplateId?: pulumi.Input<string>;
        LaunchTemplateName?: pulumi.Input<string>;
        Version: pulumi.Input<string>;
    }
    
    export interface InstanceLicenseSpecificationAttributes {
        LicenseConfigurationArn: string;
    }
    
    export interface InstanceLicenseSpecificationProperties {
        LicenseConfigurationArn: pulumi.Input<string>;
    }
    
    export interface InstanceNetworkInterfaceAttributes {
        AssociatePublicIpAddress?: boolean;
        DeleteOnTermination?: boolean;
        Description?: string;
        DeviceIndex: string;
        GroupSet?: string[];
        Ipv6AddressCount?: number;
        Ipv6Addresses?: InstanceInstanceIpv6AddressAttributes[];
        NetworkInterfaceId?: string;
        PrivateIpAddress?: string;
        PrivateIpAddresses?: InstancePrivateIpAddressSpecificationAttributes[];
        SecondaryPrivateIpAddressCount?: number;
        SubnetId?: string;
    }
    
    export interface InstanceNetworkInterfaceProperties {
        AssociatePublicIpAddress?: pulumi.Input<boolean>;
        DeleteOnTermination?: pulumi.Input<boolean>;
        Description?: pulumi.Input<string>;
        DeviceIndex: pulumi.Input<string>;
        GroupSet?: pulumi.Input<pulumi.Input<string>[]>;
        Ipv6AddressCount?: pulumi.Input<number>;
        Ipv6Addresses?: pulumi.Input<pulumi.Input<InstanceInstanceIpv6AddressProperties>[]>;
        NetworkInterfaceId?: pulumi.Input<string>;
        PrivateIpAddress?: pulumi.Input<string>;
        PrivateIpAddresses?: pulumi.Input<pulumi.Input<InstancePrivateIpAddressSpecificationProperties>[]>;
        SecondaryPrivateIpAddressCount?: pulumi.Input<number>;
        SubnetId?: pulumi.Input<string>;
    }
    
    export interface InstanceNoDeviceAttributes {
    }
    
    export interface InstanceNoDeviceProperties {
    }
    
    export interface InstancePrivateIpAddressSpecificationAttributes {
        Primary: boolean;
        PrivateIpAddress: string;
    }
    
    export interface InstancePrivateIpAddressSpecificationProperties {
        Primary: pulumi.Input<boolean>;
        PrivateIpAddress: pulumi.Input<string>;
    }
    
    export interface InstanceProperties {
        AdditionalInfo?: pulumi.Input<string>;
        Affinity?: pulumi.Input<string>;
        AvailabilityZone?: pulumi.Input<string>;
        BlockDeviceMappings?: pulumi.Input<pulumi.Input<InstanceBlockDeviceMappingProperties>[]>;
        CpuOptions?: pulumi.Input<InstanceCpuOptionsProperties>;
        CreditSpecification?: pulumi.Input<InstanceCreditSpecificationProperties>;
        DisableApiTermination?: pulumi.Input<boolean>;
        EbsOptimized?: pulumi.Input<boolean>;
        ElasticGpuSpecifications?: pulumi.Input<pulumi.Input<InstanceElasticGpuSpecificationProperties>[]>;
        ElasticInferenceAccelerators?: pulumi.Input<pulumi.Input<InstanceElasticInferenceAcceleratorProperties>[]>;
        HostId?: pulumi.Input<string>;
        IamInstanceProfile?: pulumi.Input<string>;
        ImageId?: pulumi.Input<string>;
        InstanceInitiatedShutdownBehavior?: pulumi.Input<string>;
        InstanceType?: pulumi.Input<string>;
        Ipv6AddressCount?: pulumi.Input<number>;
        Ipv6Addresses?: pulumi.Input<pulumi.Input<InstanceInstanceIpv6AddressProperties>[]>;
        KernelId?: pulumi.Input<string>;
        KeyName?: pulumi.Input<string>;
        LaunchTemplate?: pulumi.Input<InstanceLaunchTemplateSpecificationProperties>;
        LicenseSpecifications?: pulumi.Input<pulumi.Input<InstanceLicenseSpecificationProperties>[]>;
        Monitoring?: pulumi.Input<boolean>;
        NetworkInterfaces?: pulumi.Input<pulumi.Input<InstanceNetworkInterfaceProperties>[]>;
        PlacementGroupName?: pulumi.Input<string>;
        PrivateIpAddress?: pulumi.Input<string>;
        RamdiskId?: pulumi.Input<string>;
        SecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
        SecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
        SourceDestCheck?: pulumi.Input<boolean>;
        SsmAssociations?: pulumi.Input<pulumi.Input<InstanceSsmAssociationProperties>[]>;
        SubnetId?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        Tenancy?: pulumi.Input<string>;
        UserData?: pulumi.Input<string>;
        Volumes?: pulumi.Input<pulumi.Input<InstanceVolumeProperties>[]>;
    }
    
    export interface InstanceSsmAssociationAttributes {
        AssociationParameters?: InstanceAssociationParameterAttributes[];
        DocumentName: string;
    }
    
    export interface InstanceSsmAssociationProperties {
        AssociationParameters?: pulumi.Input<pulumi.Input<InstanceAssociationParameterProperties>[]>;
        DocumentName: pulumi.Input<string>;
    }
    
    export interface InstanceVolumeAttributes {
        Device: string;
        VolumeId: string;
    }
    
    export interface InstanceVolumeProperties {
        Device: pulumi.Input<string>;
        VolumeId: pulumi.Input<string>;
    }
    
    export interface InternetGatewayAttributes {
    }
    
    export interface InternetGatewayProperties {
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface LaunchTemplateAttributes {
        DefaultVersionNumber: string;
        LatestVersionNumber: string;
    }
    
    export interface LaunchTemplateBlockDeviceMappingAttributes {
        DeviceName?: string;
        Ebs?: LaunchTemplateEbsAttributes;
        NoDevice?: string;
        VirtualName?: string;
    }
    
    export interface LaunchTemplateBlockDeviceMappingProperties {
        DeviceName?: pulumi.Input<string>;
        Ebs?: pulumi.Input<LaunchTemplateEbsProperties>;
        NoDevice?: pulumi.Input<string>;
        VirtualName?: pulumi.Input<string>;
    }
    
    export interface LaunchTemplateCapacityReservationPreferenceAttributes {
    }
    
    export interface LaunchTemplateCapacityReservationPreferenceProperties {
    }
    
    export interface LaunchTemplateCapacityReservationSpecificationAttributes {
        CapacityReservationPreference?: LaunchTemplateCapacityReservationPreferenceAttributes;
        CapacityReservationTarget?: LaunchTemplateCapacityReservationTargetAttributes;
    }
    
    export interface LaunchTemplateCapacityReservationSpecificationProperties {
        CapacityReservationPreference?: pulumi.Input<LaunchTemplateCapacityReservationPreferenceProperties>;
        CapacityReservationTarget?: pulumi.Input<LaunchTemplateCapacityReservationTargetProperties>;
    }
    
    export interface LaunchTemplateCapacityReservationTargetAttributes {
        CapacityReservationId?: string;
    }
    
    export interface LaunchTemplateCapacityReservationTargetProperties {
        CapacityReservationId?: pulumi.Input<string>;
    }
    
    export interface LaunchTemplateCpuOptionsAttributes {
        CoreCount?: number;
        ThreadsPerCore?: number;
    }
    
    export interface LaunchTemplateCpuOptionsProperties {
        CoreCount?: pulumi.Input<number>;
        ThreadsPerCore?: pulumi.Input<number>;
    }
    
    export interface LaunchTemplateCreditSpecificationAttributes {
        CpuCredits?: string;
    }
    
    export interface LaunchTemplateCreditSpecificationProperties {
        CpuCredits?: pulumi.Input<string>;
    }
    
    export interface LaunchTemplateEbsAttributes {
        DeleteOnTermination?: boolean;
        Encrypted?: boolean;
        Iops?: number;
        KmsKeyId?: string;
        SnapshotId?: string;
        VolumeSize?: number;
        VolumeType?: string;
    }
    
    export interface LaunchTemplateEbsProperties {
        DeleteOnTermination?: pulumi.Input<boolean>;
        Encrypted?: pulumi.Input<boolean>;
        Iops?: pulumi.Input<number>;
        KmsKeyId?: pulumi.Input<string>;
        SnapshotId?: pulumi.Input<string>;
        VolumeSize?: pulumi.Input<number>;
        VolumeType?: pulumi.Input<string>;
    }
    
    export interface LaunchTemplateElasticGpuSpecificationAttributes {
        Type?: string;
    }
    
    export interface LaunchTemplateElasticGpuSpecificationProperties {
        Type?: pulumi.Input<string>;
    }
    
    export interface LaunchTemplateHibernationOptionsAttributes {
        Configured?: boolean;
    }
    
    export interface LaunchTemplateHibernationOptionsProperties {
        Configured?: pulumi.Input<boolean>;
    }
    
    export interface LaunchTemplateIamInstanceProfileAttributes {
        Arn?: string;
        Name?: string;
    }
    
    export interface LaunchTemplateIamInstanceProfileProperties {
        Arn?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
    }
    
    export interface LaunchTemplateInstanceMarketOptionsAttributes {
        MarketType?: string;
        SpotOptions?: LaunchTemplateSpotOptionsAttributes;
    }
    
    export interface LaunchTemplateInstanceMarketOptionsProperties {
        MarketType?: pulumi.Input<string>;
        SpotOptions?: pulumi.Input<LaunchTemplateSpotOptionsProperties>;
    }
    
    export interface LaunchTemplateIpv6AddAttributes {
        Ipv6Address?: string;
    }
    
    export interface LaunchTemplateIpv6AddProperties {
        Ipv6Address?: pulumi.Input<string>;
    }
    
    export interface LaunchTemplateLaunchTemplateDataAttributes {
        BlockDeviceMappings?: LaunchTemplateBlockDeviceMappingAttributes[];
        CapacityReservationSpecification?: LaunchTemplateCapacityReservationSpecificationAttributes;
        CpuOptions?: LaunchTemplateCpuOptionsAttributes;
        CreditSpecification?: LaunchTemplateCreditSpecificationAttributes;
        DisableApiTermination?: boolean;
        EbsOptimized?: boolean;
        ElasticGpuSpecifications?: LaunchTemplateElasticGpuSpecificationAttributes[];
        ElasticInferenceAccelerators?: LaunchTemplateLaunchTemplateElasticInferenceAcceleratorAttributes[];
        HibernationOptions?: LaunchTemplateHibernationOptionsAttributes;
        IamInstanceProfile?: LaunchTemplateIamInstanceProfileAttributes;
        ImageId?: string;
        InstanceInitiatedShutdownBehavior?: string;
        InstanceMarketOptions?: LaunchTemplateInstanceMarketOptionsAttributes;
        InstanceType?: string;
        KernelId?: string;
        KeyName?: string;
        LicenseSpecifications?: LaunchTemplateLicenseSpecificationAttributes[];
        Monitoring?: LaunchTemplateMonitoringAttributes;
        NetworkInterfaces?: LaunchTemplateNetworkInterfaceAttributes[];
        Placement?: LaunchTemplatePlacementAttributes;
        RamDiskId?: string;
        SecurityGroupIds?: string[];
        SecurityGroups?: string[];
        TagSpecifications?: LaunchTemplateTagSpecificationAttributes[];
        UserData?: string;
    }
    
    export interface LaunchTemplateLaunchTemplateDataProperties {
        BlockDeviceMappings?: pulumi.Input<pulumi.Input<LaunchTemplateBlockDeviceMappingProperties>[]>;
        CapacityReservationSpecification?: pulumi.Input<LaunchTemplateCapacityReservationSpecificationProperties>;
        CpuOptions?: pulumi.Input<LaunchTemplateCpuOptionsProperties>;
        CreditSpecification?: pulumi.Input<LaunchTemplateCreditSpecificationProperties>;
        DisableApiTermination?: pulumi.Input<boolean>;
        EbsOptimized?: pulumi.Input<boolean>;
        ElasticGpuSpecifications?: pulumi.Input<pulumi.Input<LaunchTemplateElasticGpuSpecificationProperties>[]>;
        ElasticInferenceAccelerators?: pulumi.Input<pulumi.Input<LaunchTemplateLaunchTemplateElasticInferenceAcceleratorProperties>[]>;
        HibernationOptions?: pulumi.Input<LaunchTemplateHibernationOptionsProperties>;
        IamInstanceProfile?: pulumi.Input<LaunchTemplateIamInstanceProfileProperties>;
        ImageId?: pulumi.Input<string>;
        InstanceInitiatedShutdownBehavior?: pulumi.Input<string>;
        InstanceMarketOptions?: pulumi.Input<LaunchTemplateInstanceMarketOptionsProperties>;
        InstanceType?: pulumi.Input<string>;
        KernelId?: pulumi.Input<string>;
        KeyName?: pulumi.Input<string>;
        LicenseSpecifications?: pulumi.Input<pulumi.Input<LaunchTemplateLicenseSpecificationProperties>[]>;
        Monitoring?: pulumi.Input<LaunchTemplateMonitoringProperties>;
        NetworkInterfaces?: pulumi.Input<pulumi.Input<LaunchTemplateNetworkInterfaceProperties>[]>;
        Placement?: pulumi.Input<LaunchTemplatePlacementProperties>;
        RamDiskId?: pulumi.Input<string>;
        SecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
        SecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
        TagSpecifications?: pulumi.Input<pulumi.Input<LaunchTemplateTagSpecificationProperties>[]>;
        UserData?: pulumi.Input<string>;
    }
    
    export interface LaunchTemplateLaunchTemplateElasticInferenceAcceleratorAttributes {
        Type?: string;
    }
    
    export interface LaunchTemplateLaunchTemplateElasticInferenceAcceleratorProperties {
        Type?: pulumi.Input<string>;
    }
    
    export interface LaunchTemplateLicenseSpecificationAttributes {
        LicenseConfigurationArn?: string;
    }
    
    export interface LaunchTemplateLicenseSpecificationProperties {
        LicenseConfigurationArn?: pulumi.Input<string>;
    }
    
    export interface LaunchTemplateMonitoringAttributes {
        Enabled?: boolean;
    }
    
    export interface LaunchTemplateMonitoringProperties {
        Enabled?: pulumi.Input<boolean>;
    }
    
    export interface LaunchTemplateNetworkInterfaceAttributes {
        AssociatePublicIpAddress?: boolean;
        DeleteOnTermination?: boolean;
        Description?: string;
        DeviceIndex?: number;
        Groups?: string[];
        InterfaceType?: string;
        Ipv6AddressCount?: number;
        Ipv6Addresses?: LaunchTemplateIpv6AddAttributes[];
        NetworkInterfaceId?: string;
        PrivateIpAddress?: string;
        PrivateIpAddresses?: LaunchTemplatePrivateIpAddAttributes[];
        SecondaryPrivateIpAddressCount?: number;
        SubnetId?: string;
    }
    
    export interface LaunchTemplateNetworkInterfaceProperties {
        AssociatePublicIpAddress?: pulumi.Input<boolean>;
        DeleteOnTermination?: pulumi.Input<boolean>;
        Description?: pulumi.Input<string>;
        DeviceIndex?: pulumi.Input<number>;
        Groups?: pulumi.Input<pulumi.Input<string>[]>;
        InterfaceType?: pulumi.Input<string>;
        Ipv6AddressCount?: pulumi.Input<number>;
        Ipv6Addresses?: pulumi.Input<pulumi.Input<LaunchTemplateIpv6AddProperties>[]>;
        NetworkInterfaceId?: pulumi.Input<string>;
        PrivateIpAddress?: pulumi.Input<string>;
        PrivateIpAddresses?: pulumi.Input<pulumi.Input<LaunchTemplatePrivateIpAddProperties>[]>;
        SecondaryPrivateIpAddressCount?: pulumi.Input<number>;
        SubnetId?: pulumi.Input<string>;
    }
    
    export interface LaunchTemplatePlacementAttributes {
        Affinity?: string;
        AvailabilityZone?: string;
        GroupName?: string;
        HostId?: string;
        Tenancy?: string;
    }
    
    export interface LaunchTemplatePlacementProperties {
        Affinity?: pulumi.Input<string>;
        AvailabilityZone?: pulumi.Input<string>;
        GroupName?: pulumi.Input<string>;
        HostId?: pulumi.Input<string>;
        Tenancy?: pulumi.Input<string>;
    }
    
    export interface LaunchTemplatePrivateIpAddAttributes {
        Primary?: boolean;
        PrivateIpAddress?: string;
    }
    
    export interface LaunchTemplatePrivateIpAddProperties {
        Primary?: pulumi.Input<boolean>;
        PrivateIpAddress?: pulumi.Input<string>;
    }
    
    export interface LaunchTemplateProperties {
        LaunchTemplateData?: pulumi.Input<LaunchTemplateLaunchTemplateDataProperties>;
        LaunchTemplateName?: pulumi.Input<string>;
    }
    
    export interface LaunchTemplateSpotOptionsAttributes {
        BlockDurationMinutes?: number;
        InstanceInterruptionBehavior?: string;
        MaxPrice?: string;
        SpotInstanceType?: string;
        ValidUntil?: string;
    }
    
    export interface LaunchTemplateSpotOptionsProperties {
        BlockDurationMinutes?: pulumi.Input<number>;
        InstanceInterruptionBehavior?: pulumi.Input<string>;
        MaxPrice?: pulumi.Input<string>;
        SpotInstanceType?: pulumi.Input<string>;
        ValidUntil?: pulumi.Input<string>;
    }
    
    export interface LaunchTemplateTagSpecificationAttributes {
        ResourceType?: string;
        Tags?: TagAttributes[];
    }
    
    export interface LaunchTemplateTagSpecificationProperties {
        ResourceType?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface NatGatewayAttributes {
    }
    
    export interface NatGatewayProperties {
        AllocationId: pulumi.Input<string>;
        SubnetId: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface NetworkAclAttributes {
    }
    
    export interface NetworkAclEntryAttributes {
    }
    
    export interface NetworkAclEntryIcmpAttributes {
        Code?: number;
        Type?: number;
    }
    
    export interface NetworkAclEntryIcmpProperties {
        Code?: pulumi.Input<number>;
        Type?: pulumi.Input<number>;
    }
    
    export interface NetworkAclEntryPortRangeAttributes {
        From?: number;
        To?: number;
    }
    
    export interface NetworkAclEntryPortRangeProperties {
        From?: pulumi.Input<number>;
        To?: pulumi.Input<number>;
    }
    
    export interface NetworkAclEntryProperties {
        CidrBlock: pulumi.Input<string>;
        Egress?: pulumi.Input<boolean>;
        Icmp?: pulumi.Input<NetworkAclEntryIcmpProperties>;
        Ipv6CidrBlock?: pulumi.Input<string>;
        NetworkAclId: pulumi.Input<string>;
        PortRange?: pulumi.Input<NetworkAclEntryPortRangeProperties>;
        Protocol: pulumi.Input<number>;
        RuleAction: pulumi.Input<string>;
        RuleNumber: pulumi.Input<number>;
    }
    
    export interface NetworkAclProperties {
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VpcId: pulumi.Input<string>;
    }
    
    export interface NetworkInterfaceAttachmentAttributes {
    }
    
    export interface NetworkInterfaceAttachmentProperties {
        DeleteOnTermination?: pulumi.Input<boolean>;
        DeviceIndex: pulumi.Input<string>;
        InstanceId: pulumi.Input<string>;
        NetworkInterfaceId: pulumi.Input<string>;
    }
    
    export interface NetworkInterfaceAttributes {
        PrimaryPrivateIpAddress: string;
        SecondaryPrivateIpAddresses: string[];
    }
    
    export interface NetworkInterfaceInstanceIpv6AddressAttributes {
        Ipv6Address: string;
    }
    
    export interface NetworkInterfaceInstanceIpv6AddressProperties {
        Ipv6Address: pulumi.Input<string>;
    }
    
    export interface NetworkInterfacePermissionAttributes {
    }
    
    export interface NetworkInterfacePermissionProperties {
        AwsAccountId: pulumi.Input<string>;
        NetworkInterfaceId: pulumi.Input<string>;
        Permission: pulumi.Input<string>;
    }
    
    export interface NetworkInterfacePrivateIpAddressSpecificationAttributes {
        Primary: boolean;
        PrivateIpAddress: string;
    }
    
    export interface NetworkInterfacePrivateIpAddressSpecificationProperties {
        Primary: pulumi.Input<boolean>;
        PrivateIpAddress: pulumi.Input<string>;
    }
    
    export interface NetworkInterfaceProperties {
        Description?: pulumi.Input<string>;
        GroupSet?: pulumi.Input<pulumi.Input<string>[]>;
        InterfaceType?: pulumi.Input<string>;
        Ipv6AddressCount?: pulumi.Input<number>;
        Ipv6Addresses?: pulumi.Input<NetworkInterfaceInstanceIpv6AddressProperties>;
        PrivateIpAddress?: pulumi.Input<string>;
        PrivateIpAddresses?: pulumi.Input<pulumi.Input<NetworkInterfacePrivateIpAddressSpecificationProperties>[]>;
        SecondaryPrivateIpAddressCount?: pulumi.Input<number>;
        SourceDestCheck?: pulumi.Input<boolean>;
        SubnetId: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface PlacementGroupAttributes {
    }
    
    export interface PlacementGroupProperties {
        Strategy?: pulumi.Input<string>;
    }
    
    export interface RouteAttributes {
    }
    
    export interface RouteProperties {
        DestinationCidrBlock?: pulumi.Input<string>;
        DestinationIpv6CidrBlock?: pulumi.Input<string>;
        EgressOnlyInternetGatewayId?: pulumi.Input<string>;
        GatewayId?: pulumi.Input<string>;
        InstanceId?: pulumi.Input<string>;
        NatGatewayId?: pulumi.Input<string>;
        NetworkInterfaceId?: pulumi.Input<string>;
        RouteTableId: pulumi.Input<string>;
        TransitGatewayId?: pulumi.Input<string>;
        VpcPeeringConnectionId?: pulumi.Input<string>;
    }
    
    export interface RouteTableAttributes {
    }
    
    export interface RouteTableProperties {
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VpcId: pulumi.Input<string>;
    }
    
    export interface SecurityGroupAttributes {
        GroupId: string;
        VpcId: string;
    }
    
    export interface SecurityGroupEgressAttributes {
    }
    
    export interface SecurityGroupEgressProperties {
        CidrIp?: pulumi.Input<string>;
        CidrIpv6?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        DestinationPrefixListId?: pulumi.Input<string>;
        DestinationSecurityGroupId?: pulumi.Input<string>;
        FromPort?: pulumi.Input<number>;
        GroupId: pulumi.Input<string>;
        IpProtocol: pulumi.Input<string>;
        ToPort?: pulumi.Input<number>;
    }
    
    export interface SecurityGroupIngressAttributes {
    }
    
    export interface SecurityGroupIngressProperties {
        CidrIp?: pulumi.Input<string>;
        CidrIpv6?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        FromPort?: pulumi.Input<number>;
        GroupId?: pulumi.Input<string>;
        GroupName?: pulumi.Input<string>;
        IpProtocol: pulumi.Input<string>;
        SourcePrefixListId?: pulumi.Input<string>;
        SourceSecurityGroupId?: pulumi.Input<string>;
        SourceSecurityGroupName?: pulumi.Input<string>;
        SourceSecurityGroupOwnerId?: pulumi.Input<string>;
        ToPort?: pulumi.Input<number>;
    }
    
    export interface SecurityGroupProperties {
        GroupDescription: pulumi.Input<string>;
        GroupName?: pulumi.Input<string>;
        SecurityGroupEgress?: pulumi.Input<pulumi.Input<SecurityGroupEgressProperties>[]>;
        SecurityGroupIngress?: pulumi.Input<pulumi.Input<SecurityGroupIngressProperties>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VpcId?: pulumi.Input<string>;
    }
    
    export interface SpotFleetAttributes {
    }
    
    export interface SpotFleetBlockDeviceMappingAttributes {
        DeviceName: string;
        Ebs?: SpotFleetEbsBlockDeviceAttributes;
        NoDevice?: string;
        VirtualName?: string;
    }
    
    export interface SpotFleetBlockDeviceMappingProperties {
        DeviceName: pulumi.Input<string>;
        Ebs?: pulumi.Input<SpotFleetEbsBlockDeviceProperties>;
        NoDevice?: pulumi.Input<string>;
        VirtualName?: pulumi.Input<string>;
    }
    
    export interface SpotFleetClassicLoadBalancerAttributes {
        Name: string;
    }
    
    export interface SpotFleetClassicLoadBalancerProperties {
        Name: pulumi.Input<string>;
    }
    
    export interface SpotFleetClassicLoadBalancersConfigAttributes {
        ClassicLoadBalancers: SpotFleetClassicLoadBalancerAttributes[];
    }
    
    export interface SpotFleetClassicLoadBalancersConfigProperties {
        ClassicLoadBalancers: pulumi.Input<pulumi.Input<SpotFleetClassicLoadBalancerProperties>[]>;
    }
    
    export interface SpotFleetEbsBlockDeviceAttributes {
        DeleteOnTermination?: boolean;
        Encrypted?: boolean;
        Iops?: number;
        SnapshotId?: string;
        VolumeSize?: number;
        VolumeType?: string;
    }
    
    export interface SpotFleetEbsBlockDeviceProperties {
        DeleteOnTermination?: pulumi.Input<boolean>;
        Encrypted?: pulumi.Input<boolean>;
        Iops?: pulumi.Input<number>;
        SnapshotId?: pulumi.Input<string>;
        VolumeSize?: pulumi.Input<number>;
        VolumeType?: pulumi.Input<string>;
    }
    
    export interface SpotFleetFleetLaunchTemplateSpecificationAttributes {
        LaunchTemplateId?: string;
        LaunchTemplateName?: string;
        Version: string;
    }
    
    export interface SpotFleetFleetLaunchTemplateSpecificationProperties {
        LaunchTemplateId?: pulumi.Input<string>;
        LaunchTemplateName?: pulumi.Input<string>;
        Version: pulumi.Input<string>;
    }
    
    export interface SpotFleetGroupIdentifierAttributes {
        GroupId: string;
    }
    
    export interface SpotFleetGroupIdentifierProperties {
        GroupId: pulumi.Input<string>;
    }
    
    export interface SpotFleetIamInstanceProfileSpecificationAttributes {
        Arn?: string;
    }
    
    export interface SpotFleetIamInstanceProfileSpecificationProperties {
        Arn?: pulumi.Input<string>;
    }
    
    export interface SpotFleetInstanceIpv6AddressAttributes {
        Ipv6Address: string;
    }
    
    export interface SpotFleetInstanceIpv6AddressProperties {
        Ipv6Address: pulumi.Input<string>;
    }
    
    export interface SpotFleetInstanceNetworkInterfaceSpecificationAttributes {
        AssociatePublicIpAddress?: boolean;
        DeleteOnTermination?: boolean;
        Description?: string;
        DeviceIndex?: number;
        Groups?: string[];
        Ipv6AddressCount?: number;
        Ipv6Addresses?: SpotFleetInstanceIpv6AddressAttributes[];
        NetworkInterfaceId?: string;
        PrivateIpAddresses?: SpotFleetPrivateIpAddressSpecificationAttributes[];
        SecondaryPrivateIpAddressCount?: number;
        SubnetId?: string;
    }
    
    export interface SpotFleetInstanceNetworkInterfaceSpecificationProperties {
        AssociatePublicIpAddress?: pulumi.Input<boolean>;
        DeleteOnTermination?: pulumi.Input<boolean>;
        Description?: pulumi.Input<string>;
        DeviceIndex?: pulumi.Input<number>;
        Groups?: pulumi.Input<pulumi.Input<string>[]>;
        Ipv6AddressCount?: pulumi.Input<number>;
        Ipv6Addresses?: pulumi.Input<pulumi.Input<SpotFleetInstanceIpv6AddressProperties>[]>;
        NetworkInterfaceId?: pulumi.Input<string>;
        PrivateIpAddresses?: pulumi.Input<pulumi.Input<SpotFleetPrivateIpAddressSpecificationProperties>[]>;
        SecondaryPrivateIpAddressCount?: pulumi.Input<number>;
        SubnetId?: pulumi.Input<string>;
    }
    
    export interface SpotFleetLaunchTemplateConfigAttributes {
        LaunchTemplateSpecification?: SpotFleetFleetLaunchTemplateSpecificationAttributes;
        Overrides?: SpotFleetLaunchTemplateOverridesAttributes[];
    }
    
    export interface SpotFleetLaunchTemplateConfigProperties {
        LaunchTemplateSpecification?: pulumi.Input<SpotFleetFleetLaunchTemplateSpecificationProperties>;
        Overrides?: pulumi.Input<pulumi.Input<SpotFleetLaunchTemplateOverridesProperties>[]>;
    }
    
    export interface SpotFleetLaunchTemplateOverridesAttributes {
        AvailabilityZone?: string;
        InstanceType?: string;
        SpotPrice?: string;
        SubnetId?: string;
        WeightedCapacity?: number;
    }
    
    export interface SpotFleetLaunchTemplateOverridesProperties {
        AvailabilityZone?: pulumi.Input<string>;
        InstanceType?: pulumi.Input<string>;
        SpotPrice?: pulumi.Input<string>;
        SubnetId?: pulumi.Input<string>;
        WeightedCapacity?: pulumi.Input<number>;
    }
    
    export interface SpotFleetLoadBalancersConfigAttributes {
        ClassicLoadBalancersConfig?: SpotFleetClassicLoadBalancersConfigAttributes;
        TargetGroupsConfig?: SpotFleetTargetGroupsConfigAttributes;
    }
    
    export interface SpotFleetLoadBalancersConfigProperties {
        ClassicLoadBalancersConfig?: pulumi.Input<SpotFleetClassicLoadBalancersConfigProperties>;
        TargetGroupsConfig?: pulumi.Input<SpotFleetTargetGroupsConfigProperties>;
    }
    
    export interface SpotFleetPrivateIpAddressSpecificationAttributes {
        Primary?: boolean;
        PrivateIpAddress: string;
    }
    
    export interface SpotFleetPrivateIpAddressSpecificationProperties {
        Primary?: pulumi.Input<boolean>;
        PrivateIpAddress: pulumi.Input<string>;
    }
    
    export interface SpotFleetProperties {
        SpotFleetRequestConfigData: pulumi.Input<SpotFleetSpotFleetRequestConfigDataProperties>;
    }
    
    export interface SpotFleetSpotFleetLaunchSpecificationAttributes {
        BlockDeviceMappings?: SpotFleetBlockDeviceMappingAttributes[];
        EbsOptimized?: boolean;
        IamInstanceProfile?: SpotFleetIamInstanceProfileSpecificationAttributes;
        ImageId: string;
        InstanceType: string;
        KernelId?: string;
        KeyName?: string;
        Monitoring?: SpotFleetSpotFleetMonitoringAttributes;
        NetworkInterfaces?: SpotFleetInstanceNetworkInterfaceSpecificationAttributes[];
        Placement?: SpotFleetSpotPlacementAttributes;
        RamdiskId?: string;
        SecurityGroups?: SpotFleetGroupIdentifierAttributes[];
        SpotPrice?: string;
        SubnetId?: string;
        TagSpecifications?: SpotFleetSpotFleetTagSpecificationAttributes[];
        UserData?: string;
        WeightedCapacity?: number;
    }
    
    export interface SpotFleetSpotFleetLaunchSpecificationProperties {
        BlockDeviceMappings?: pulumi.Input<pulumi.Input<SpotFleetBlockDeviceMappingProperties>[]>;
        EbsOptimized?: pulumi.Input<boolean>;
        IamInstanceProfile?: pulumi.Input<SpotFleetIamInstanceProfileSpecificationProperties>;
        ImageId: pulumi.Input<string>;
        InstanceType: pulumi.Input<string>;
        KernelId?: pulumi.Input<string>;
        KeyName?: pulumi.Input<string>;
        Monitoring?: pulumi.Input<SpotFleetSpotFleetMonitoringProperties>;
        NetworkInterfaces?: pulumi.Input<pulumi.Input<SpotFleetInstanceNetworkInterfaceSpecificationProperties>[]>;
        Placement?: pulumi.Input<SpotFleetSpotPlacementProperties>;
        RamdiskId?: pulumi.Input<string>;
        SecurityGroups?: pulumi.Input<pulumi.Input<SpotFleetGroupIdentifierProperties>[]>;
        SpotPrice?: pulumi.Input<string>;
        SubnetId?: pulumi.Input<string>;
        TagSpecifications?: pulumi.Input<pulumi.Input<SpotFleetSpotFleetTagSpecificationProperties>[]>;
        UserData?: pulumi.Input<string>;
        WeightedCapacity?: pulumi.Input<number>;
    }
    
    export interface SpotFleetSpotFleetMonitoringAttributes {
        Enabled?: boolean;
    }
    
    export interface SpotFleetSpotFleetMonitoringProperties {
        Enabled?: pulumi.Input<boolean>;
    }
    
    export interface SpotFleetSpotFleetRequestConfigDataAttributes {
        AllocationStrategy?: string;
        ExcessCapacityTerminationPolicy?: string;
        IamFleetRole: string;
        InstanceInterruptionBehavior?: string;
        LaunchSpecifications?: SpotFleetSpotFleetLaunchSpecificationAttributes[];
        LaunchTemplateConfigs?: SpotFleetLaunchTemplateConfigAttributes[];
        LoadBalancersConfig?: SpotFleetLoadBalancersConfigAttributes;
        ReplaceUnhealthyInstances?: boolean;
        SpotPrice?: string;
        TargetCapacity: number;
        TerminateInstancesWithExpiration?: boolean;
        Type?: string;
        ValidFrom?: string;
        ValidUntil?: string;
    }
    
    export interface SpotFleetSpotFleetRequestConfigDataProperties {
        AllocationStrategy?: pulumi.Input<string>;
        ExcessCapacityTerminationPolicy?: pulumi.Input<string>;
        IamFleetRole: pulumi.Input<string>;
        InstanceInterruptionBehavior?: pulumi.Input<string>;
        LaunchSpecifications?: pulumi.Input<pulumi.Input<SpotFleetSpotFleetLaunchSpecificationProperties>[]>;
        LaunchTemplateConfigs?: pulumi.Input<pulumi.Input<SpotFleetLaunchTemplateConfigProperties>[]>;
        LoadBalancersConfig?: pulumi.Input<SpotFleetLoadBalancersConfigProperties>;
        ReplaceUnhealthyInstances?: pulumi.Input<boolean>;
        SpotPrice?: pulumi.Input<string>;
        TargetCapacity: pulumi.Input<number>;
        TerminateInstancesWithExpiration?: pulumi.Input<boolean>;
        Type?: pulumi.Input<string>;
        ValidFrom?: pulumi.Input<string>;
        ValidUntil?: pulumi.Input<string>;
    }
    
    export interface SpotFleetSpotFleetTagSpecificationAttributes {
        ResourceType?: string;
        Tags?: TagAttributes[];
    }
    
    export interface SpotFleetSpotFleetTagSpecificationProperties {
        ResourceType?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface SpotFleetSpotPlacementAttributes {
        AvailabilityZone?: string;
        GroupName?: string;
        Tenancy?: string;
    }
    
    export interface SpotFleetSpotPlacementProperties {
        AvailabilityZone?: pulumi.Input<string>;
        GroupName?: pulumi.Input<string>;
        Tenancy?: pulumi.Input<string>;
    }
    
    export interface SpotFleetTargetGroupAttributes {
        Arn: string;
    }
    
    export interface SpotFleetTargetGroupProperties {
        Arn: pulumi.Input<string>;
    }
    
    export interface SpotFleetTargetGroupsConfigAttributes {
        TargetGroups: SpotFleetTargetGroupAttributes[];
    }
    
    export interface SpotFleetTargetGroupsConfigProperties {
        TargetGroups: pulumi.Input<pulumi.Input<SpotFleetTargetGroupProperties>[]>;
    }
    
    export interface SubnetAttributes {
        AvailabilityZone: string;
        Ipv6CidrBlocks: string[];
        NetworkAclAssociationId: string;
        VpcId: string;
    }
    
    export interface SubnetCidrBlockAttributes {
    }
    
    export interface SubnetCidrBlockProperties {
        Ipv6CidrBlock: pulumi.Input<string>;
        SubnetId: pulumi.Input<string>;
    }
    
    export interface SubnetNetworkAclAssociationAttributes {
        AssociationId: string;
    }
    
    export interface SubnetNetworkAclAssociationProperties {
        NetworkAclId: pulumi.Input<string>;
        SubnetId: pulumi.Input<string>;
    }
    
    export interface SubnetProperties {
        AssignIpv6AddressOnCreation?: pulumi.Input<boolean>;
        AvailabilityZone?: pulumi.Input<string>;
        CidrBlock: pulumi.Input<string>;
        Ipv6CidrBlock?: pulumi.Input<string>;
        MapPublicIpOnLaunch?: pulumi.Input<boolean>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VpcId: pulumi.Input<string>;
    }
    
    export interface SubnetRouteTableAssociationAttributes {
    }
    
    export interface SubnetRouteTableAssociationProperties {
        RouteTableId: pulumi.Input<string>;
        SubnetId: pulumi.Input<string>;
    }
    
    export interface TrafficMirrorFilterAttributes {
    }
    
    export interface TrafficMirrorFilterProperties {
        Description?: pulumi.Input<string>;
        NetworkServices?: pulumi.Input<pulumi.Input<string>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface TrafficMirrorFilterRuleAttributes {
    }
    
    export interface TrafficMirrorFilterRuleProperties {
        Description?: pulumi.Input<string>;
        DestinationCidrBlock: pulumi.Input<string>;
        DestinationPortRange?: pulumi.Input<TrafficMirrorFilterRuleTrafficMirrorPortRangeProperties>;
        Protocol?: pulumi.Input<number>;
        RuleAction: pulumi.Input<string>;
        RuleNumber: pulumi.Input<number>;
        SourceCidrBlock: pulumi.Input<string>;
        SourcePortRange?: pulumi.Input<TrafficMirrorFilterRuleTrafficMirrorPortRangeProperties>;
        TrafficDirection: pulumi.Input<string>;
        TrafficMirrorFilterId: pulumi.Input<string>;
    }
    
    export interface TrafficMirrorFilterRuleTrafficMirrorPortRangeAttributes {
        FromPort: number;
        ToPort: number;
    }
    
    export interface TrafficMirrorFilterRuleTrafficMirrorPortRangeProperties {
        FromPort: pulumi.Input<number>;
        ToPort: pulumi.Input<number>;
    }
    
    export interface TrafficMirrorSessionAttributes {
    }
    
    export interface TrafficMirrorSessionProperties {
        Description?: pulumi.Input<string>;
        NetworkInterfaceId: pulumi.Input<string>;
        PacketLength?: pulumi.Input<number>;
        SessionNumber: pulumi.Input<number>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        TrafficMirrorFilterId: pulumi.Input<string>;
        TrafficMirrorTargetId: pulumi.Input<string>;
        VirtualNetworkId?: pulumi.Input<number>;
    }
    
    export interface TrafficMirrorTargetAttributes {
    }
    
    export interface TrafficMirrorTargetProperties {
        Description?: pulumi.Input<string>;
        NetworkInterfaceId?: pulumi.Input<string>;
        NetworkLoadBalancerArn?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface TransitGatewayAttachmentAttributes {
    }
    
    export interface TransitGatewayAttachmentProperties {
        SubnetIds: pulumi.Input<pulumi.Input<string>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        TransitGatewayId: pulumi.Input<string>;
        VpcId: pulumi.Input<string>;
    }
    
    export interface TransitGatewayAttributes {
    }
    
    export interface TransitGatewayProperties {
        AmazonSideAsn?: pulumi.Input<number>;
        AutoAcceptSharedAttachments?: pulumi.Input<string>;
        DefaultRouteTableAssociation?: pulumi.Input<string>;
        DefaultRouteTablePropagation?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        DnsSupport?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VpnEcmpSupport?: pulumi.Input<string>;
    }
    
    export interface TransitGatewayRouteAttributes {
    }
    
    export interface TransitGatewayRouteProperties {
        Blackhole?: pulumi.Input<boolean>;
        DestinationCidrBlock?: pulumi.Input<string>;
        TransitGatewayAttachmentId?: pulumi.Input<string>;
        TransitGatewayRouteTableId: pulumi.Input<string>;
    }
    
    export interface TransitGatewayRouteTableAssociationAttributes {
    }
    
    export interface TransitGatewayRouteTableAssociationProperties {
        TransitGatewayAttachmentId: pulumi.Input<string>;
        TransitGatewayRouteTableId: pulumi.Input<string>;
    }
    
    export interface TransitGatewayRouteTableAttributes {
    }
    
    export interface TransitGatewayRouteTablePropagationAttributes {
    }
    
    export interface TransitGatewayRouteTablePropagationProperties {
        TransitGatewayAttachmentId: pulumi.Input<string>;
        TransitGatewayRouteTableId: pulumi.Input<string>;
    }
    
    export interface TransitGatewayRouteTableProperties {
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        TransitGatewayId: pulumi.Input<string>;
    }
    
    export interface VPCAttributes {
        CidrBlock: string;
        CidrBlockAssociations: string[];
        DefaultNetworkAcl: string;
        DefaultSecurityGroup: string;
        Ipv6CidrBlocks: string[];
    }
    
    export interface VPCCidrBlockAttributes {
    }
    
    export interface VPCCidrBlockProperties {
        AmazonProvidedIpv6CidrBlock?: pulumi.Input<boolean>;
        CidrBlock?: pulumi.Input<string>;
        VpcId: pulumi.Input<string>;
    }
    
    export interface VPCDHCPOptionsAssociationAttributes {
    }
    
    export interface VPCDHCPOptionsAssociationProperties {
        DhcpOptionsId: pulumi.Input<string>;
        VpcId: pulumi.Input<string>;
    }
    
    export interface VPCEndpointAttributes {
        CreationTimestamp: string;
        DnsEntries: string[];
        NetworkInterfaceIds: string[];
    }
    
    export interface VPCEndpointProperties {
        PolicyDocument?: pulumi.Input<string>;
        PrivateDnsEnabled?: pulumi.Input<boolean>;
        RouteTableIds?: pulumi.Input<pulumi.Input<string>[]>;
        SecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
        ServiceName: pulumi.Input<string>;
        SubnetIds?: pulumi.Input<pulumi.Input<string>[]>;
        VpcEndpointType?: pulumi.Input<string>;
        VpcId: pulumi.Input<string>;
    }
    
    export interface VPCEndpointServicePermissionsAttributes {
    }
    
    export interface VPCEndpointServicePermissionsProperties {
        AllowedPrincipals?: pulumi.Input<pulumi.Input<string>[]>;
        ServiceId: pulumi.Input<string>;
    }
    
    export interface VPCGatewayAttachmentAttributes {
    }
    
    export interface VPCGatewayAttachmentProperties {
        InternetGatewayId?: pulumi.Input<string>;
        VpcId: pulumi.Input<string>;
        VpnGatewayId?: pulumi.Input<string>;
    }
    
    export interface VPCPeeringConnectionAttributes {
    }
    
    export interface VPCPeeringConnectionProperties {
        PeerOwnerId?: pulumi.Input<string>;
        PeerRegion?: pulumi.Input<string>;
        PeerRoleArn?: pulumi.Input<string>;
        PeerVpcId: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VpcId: pulumi.Input<string>;
    }
    
    export interface VPCProperties {
        CidrBlock: pulumi.Input<string>;
        EnableDnsHostnames?: pulumi.Input<boolean>;
        EnableDnsSupport?: pulumi.Input<boolean>;
        InstanceTenancy?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface VPNConnectionAttributes {
    }
    
    export interface VPNConnectionProperties {
        CustomerGatewayId: pulumi.Input<string>;
        StaticRoutesOnly?: pulumi.Input<boolean>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        TransitGatewayId?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
        VpnGatewayId?: pulumi.Input<string>;
        VpnTunnelOptionsSpecifications?: pulumi.Input<pulumi.Input<VPNConnectionVpnTunnelOptionsSpecificationProperties>[]>;
    }
    
    export interface VPNConnectionRouteAttributes {
    }
    
    export interface VPNConnectionRouteProperties {
        DestinationCidrBlock: pulumi.Input<string>;
        VpnConnectionId: pulumi.Input<string>;
    }
    
    export interface VPNConnectionVpnTunnelOptionsSpecificationAttributes {
        PreSharedKey?: string;
        TunnelInsideCidr?: string;
    }
    
    export interface VPNConnectionVpnTunnelOptionsSpecificationProperties {
        PreSharedKey?: pulumi.Input<string>;
        TunnelInsideCidr?: pulumi.Input<string>;
    }
    
    export interface VPNGatewayAttributes {
    }
    
    export interface VPNGatewayProperties {
        AmazonSideAsn?: pulumi.Input<number>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        Type: pulumi.Input<string>;
    }
    
    export interface VPNGatewayRoutePropagationAttributes {
    }
    
    export interface VPNGatewayRoutePropagationProperties {
        RouteTableIds: pulumi.Input<pulumi.Input<string>[]>;
        VpnGatewayId: pulumi.Input<string>;
    }
    
    export interface VolumeAttachmentAttributes {
    }
    
    export interface VolumeAttachmentProperties {
        Device: pulumi.Input<string>;
        InstanceId: pulumi.Input<string>;
        VolumeId: pulumi.Input<string>;
    }
    
    export interface VolumeAttributes {
    }
    
    export interface VolumeProperties {
        AutoEnableIO?: pulumi.Input<boolean>;
        AvailabilityZone: pulumi.Input<string>;
        Encrypted?: pulumi.Input<boolean>;
        Iops?: pulumi.Input<number>;
        KmsKeyId?: pulumi.Input<string>;
        Size?: pulumi.Input<number>;
        SnapshotId?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VolumeType?: pulumi.Input<string>;
    }
    
    
    export class CapacityReservation extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<CapacityReservationAttributes>;

        constructor(name: string, properties: CapacityReservationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:CapacityReservation", name, inputs, opts);
        }
    }
    
    export class ClientVpnAuthorizationRule extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ClientVpnAuthorizationRuleAttributes>;

        constructor(name: string, properties: ClientVpnAuthorizationRuleProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:ClientVpnAuthorizationRule", name, inputs, opts);
        }
    }
    
    export class ClientVpnEndpoint extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ClientVpnEndpointAttributes>;

        constructor(name: string, properties: ClientVpnEndpointProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:ClientVpnEndpoint", name, inputs, opts);
        }
    }
    
    export class ClientVpnRoute extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ClientVpnRouteAttributes>;

        constructor(name: string, properties: ClientVpnRouteProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:ClientVpnRoute", name, inputs, opts);
        }
    }
    
    export class ClientVpnTargetNetworkAssociation extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ClientVpnTargetNetworkAssociationAttributes>;

        constructor(name: string, properties: ClientVpnTargetNetworkAssociationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:ClientVpnTargetNetworkAssociation", name, inputs, opts);
        }
    }
    
    export class CustomerGateway extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<CustomerGatewayAttributes>;

        constructor(name: string, properties: CustomerGatewayProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:CustomerGateway", name, inputs, opts);
        }
    }
    
    export class DHCPOptions extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DHCPOptionsAttributes>;

        constructor(name: string, properties?: DHCPOptionsProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:DHCPOptions", name, inputs, opts);
        }
    }
    
    export class EC2Fleet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<EC2FleetAttributes>;

        constructor(name: string, properties: EC2FleetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:EC2Fleet", name, inputs, opts);
        }
    }
    
    export class EIP extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<EIPAttributes>;

        constructor(name: string, properties?: EIPProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:EIP", name, inputs, opts);
        }
    }
    
    export class EIPAssociation extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<EIPAssociationAttributes>;

        constructor(name: string, properties?: EIPAssociationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:EIPAssociation", name, inputs, opts);
        }
    }
    
    export class EgressOnlyInternetGateway extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<EgressOnlyInternetGatewayAttributes>;

        constructor(name: string, properties: EgressOnlyInternetGatewayProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:EgressOnlyInternetGateway", name, inputs, opts);
        }
    }
    
    export class FlowLog extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<FlowLogAttributes>;

        constructor(name: string, properties: FlowLogProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:FlowLog", name, inputs, opts);
        }
    }
    
    export class Host extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<HostAttributes>;

        constructor(name: string, properties: HostProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:Host", name, inputs, opts);
        }
    }
    
    export class Instance extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<InstanceAttributes>;

        constructor(name: string, properties?: InstanceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:Instance", name, inputs, opts);
        }
    }
    
    export class InternetGateway extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<InternetGatewayAttributes>;

        constructor(name: string, properties?: InternetGatewayProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:InternetGateway", name, inputs, opts);
        }
    }
    
    export class LaunchTemplate extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<LaunchTemplateAttributes>;

        constructor(name: string, properties?: LaunchTemplateProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:LaunchTemplate", name, inputs, opts);
        }
    }
    
    export class NatGateway extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<NatGatewayAttributes>;

        constructor(name: string, properties: NatGatewayProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:NatGateway", name, inputs, opts);
        }
    }
    
    export class NetworkAcl extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<NetworkAclAttributes>;

        constructor(name: string, properties: NetworkAclProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:NetworkAcl", name, inputs, opts);
        }
    }
    
    export class NetworkAclEntry extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<NetworkAclEntryAttributes>;

        constructor(name: string, properties: NetworkAclEntryProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:NetworkAclEntry", name, inputs, opts);
        }
    }
    
    export class NetworkInterface extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<NetworkInterfaceAttributes>;

        constructor(name: string, properties: NetworkInterfaceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:NetworkInterface", name, inputs, opts);
        }
    }
    
    export class NetworkInterfaceAttachment extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<NetworkInterfaceAttachmentAttributes>;

        constructor(name: string, properties: NetworkInterfaceAttachmentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:NetworkInterfaceAttachment", name, inputs, opts);
        }
    }
    
    export class NetworkInterfacePermission extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<NetworkInterfacePermissionAttributes>;

        constructor(name: string, properties: NetworkInterfacePermissionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:NetworkInterfacePermission", name, inputs, opts);
        }
    }
    
    export class PlacementGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PlacementGroupAttributes>;

        constructor(name: string, properties?: PlacementGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:PlacementGroup", name, inputs, opts);
        }
    }
    
    export class Route extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RouteAttributes>;

        constructor(name: string, properties: RouteProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:Route", name, inputs, opts);
        }
    }
    
    export class RouteTable extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RouteTableAttributes>;

        constructor(name: string, properties: RouteTableProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:RouteTable", name, inputs, opts);
        }
    }
    
    export class SecurityGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SecurityGroupAttributes>;

        constructor(name: string, properties: SecurityGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:SecurityGroup", name, inputs, opts);
        }
    }
    
    export class SecurityGroupEgress extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SecurityGroupEgressAttributes>;

        constructor(name: string, properties: SecurityGroupEgressProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:SecurityGroupEgress", name, inputs, opts);
        }
    }
    
    export class SecurityGroupIngress extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SecurityGroupIngressAttributes>;

        constructor(name: string, properties: SecurityGroupIngressProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:SecurityGroupIngress", name, inputs, opts);
        }
    }
    
    export class SpotFleet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SpotFleetAttributes>;

        constructor(name: string, properties: SpotFleetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:SpotFleet", name, inputs, opts);
        }
    }
    
    export class Subnet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SubnetAttributes>;

        constructor(name: string, properties: SubnetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:Subnet", name, inputs, opts);
        }
    }
    
    export class SubnetCidrBlock extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SubnetCidrBlockAttributes>;

        constructor(name: string, properties: SubnetCidrBlockProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:SubnetCidrBlock", name, inputs, opts);
        }
    }
    
    export class SubnetNetworkAclAssociation extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SubnetNetworkAclAssociationAttributes>;

        constructor(name: string, properties: SubnetNetworkAclAssociationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:SubnetNetworkAclAssociation", name, inputs, opts);
        }
    }
    
    export class SubnetRouteTableAssociation extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SubnetRouteTableAssociationAttributes>;

        constructor(name: string, properties: SubnetRouteTableAssociationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:SubnetRouteTableAssociation", name, inputs, opts);
        }
    }
    
    export class TrafficMirrorFilter extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TrafficMirrorFilterAttributes>;

        constructor(name: string, properties?: TrafficMirrorFilterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:TrafficMirrorFilter", name, inputs, opts);
        }
    }
    
    export class TrafficMirrorFilterRule extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TrafficMirrorFilterRuleAttributes>;

        constructor(name: string, properties: TrafficMirrorFilterRuleProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:TrafficMirrorFilterRule", name, inputs, opts);
        }
    }
    
    export class TrafficMirrorSession extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TrafficMirrorSessionAttributes>;

        constructor(name: string, properties: TrafficMirrorSessionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:TrafficMirrorSession", name, inputs, opts);
        }
    }
    
    export class TrafficMirrorTarget extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TrafficMirrorTargetAttributes>;

        constructor(name: string, properties?: TrafficMirrorTargetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:TrafficMirrorTarget", name, inputs, opts);
        }
    }
    
    export class TransitGateway extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TransitGatewayAttributes>;

        constructor(name: string, properties?: TransitGatewayProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:TransitGateway", name, inputs, opts);
        }
    }
    
    export class TransitGatewayAttachment extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TransitGatewayAttachmentAttributes>;

        constructor(name: string, properties: TransitGatewayAttachmentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:TransitGatewayAttachment", name, inputs, opts);
        }
    }
    
    export class TransitGatewayRoute extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TransitGatewayRouteAttributes>;

        constructor(name: string, properties: TransitGatewayRouteProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:TransitGatewayRoute", name, inputs, opts);
        }
    }
    
    export class TransitGatewayRouteTable extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TransitGatewayRouteTableAttributes>;

        constructor(name: string, properties: TransitGatewayRouteTableProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:TransitGatewayRouteTable", name, inputs, opts);
        }
    }
    
    export class TransitGatewayRouteTableAssociation extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TransitGatewayRouteTableAssociationAttributes>;

        constructor(name: string, properties: TransitGatewayRouteTableAssociationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:TransitGatewayRouteTableAssociation", name, inputs, opts);
        }
    }
    
    export class TransitGatewayRouteTablePropagation extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TransitGatewayRouteTablePropagationAttributes>;

        constructor(name: string, properties: TransitGatewayRouteTablePropagationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:TransitGatewayRouteTablePropagation", name, inputs, opts);
        }
    }
    
    export class VPC extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VPCAttributes>;

        constructor(name: string, properties: VPCProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:VPC", name, inputs, opts);
        }
    }
    
    export class VPCCidrBlock extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VPCCidrBlockAttributes>;

        constructor(name: string, properties: VPCCidrBlockProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:VPCCidrBlock", name, inputs, opts);
        }
    }
    
    export class VPCDHCPOptionsAssociation extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VPCDHCPOptionsAssociationAttributes>;

        constructor(name: string, properties: VPCDHCPOptionsAssociationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:VPCDHCPOptionsAssociation", name, inputs, opts);
        }
    }
    
    export class VPCEndpoint extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VPCEndpointAttributes>;

        constructor(name: string, properties: VPCEndpointProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:VPCEndpoint", name, inputs, opts);
        }
    }
    
    export class VPCEndpointServicePermissions extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VPCEndpointServicePermissionsAttributes>;

        constructor(name: string, properties: VPCEndpointServicePermissionsProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:VPCEndpointServicePermissions", name, inputs, opts);
        }
    }
    
    export class VPCGatewayAttachment extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VPCGatewayAttachmentAttributes>;

        constructor(name: string, properties: VPCGatewayAttachmentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:VPCGatewayAttachment", name, inputs, opts);
        }
    }
    
    export class VPCPeeringConnection extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VPCPeeringConnectionAttributes>;

        constructor(name: string, properties: VPCPeeringConnectionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:VPCPeeringConnection", name, inputs, opts);
        }
    }
    
    export class VPNConnection extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VPNConnectionAttributes>;

        constructor(name: string, properties: VPNConnectionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:VPNConnection", name, inputs, opts);
        }
    }
    
    export class VPNConnectionRoute extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VPNConnectionRouteAttributes>;

        constructor(name: string, properties: VPNConnectionRouteProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:VPNConnectionRoute", name, inputs, opts);
        }
    }
    
    export class VPNGateway extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VPNGatewayAttributes>;

        constructor(name: string, properties: VPNGatewayProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:VPNGateway", name, inputs, opts);
        }
    }
    
    export class VPNGatewayRoutePropagation extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VPNGatewayRoutePropagationAttributes>;

        constructor(name: string, properties: VPNGatewayRoutePropagationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:VPNGatewayRoutePropagation", name, inputs, opts);
        }
    }
    
    export class Volume extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VolumeAttributes>;

        constructor(name: string, properties: VolumeProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:Volume", name, inputs, opts);
        }
    }
    
    export class VolumeAttachment extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VolumeAttachmentAttributes>;

        constructor(name: string, properties: VolumeAttachmentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EC2:VolumeAttachment", name, inputs, opts);
        }
    }
    
}

export namespace ecr {
    
    export interface RepositoryAttributes {
        Arn: string;
    }
    
    export interface RepositoryLifecyclePolicyAttributes {
        LifecyclePolicyText?: string;
        RegistryId?: string;
    }
    
    export interface RepositoryLifecyclePolicyProperties {
        LifecyclePolicyText?: pulumi.Input<string>;
        RegistryId?: pulumi.Input<string>;
    }
    
    export interface RepositoryProperties {
        LifecyclePolicy?: pulumi.Input<RepositoryLifecyclePolicyProperties>;
        RepositoryName?: pulumi.Input<string>;
        RepositoryPolicyText?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    
    export class Repository extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RepositoryAttributes>;

        constructor(name: string, properties?: RepositoryProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ECR:Repository", name, inputs, opts);
        }
    }
    
}

export namespace ecs {
    
    export interface ClusterAttributes {
        Arn: string;
    }
    
    export interface ClusterProperties {
        ClusterName?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface ServiceAttributes {
        Name: string;
    }
    
    export interface ServiceAwsVpcConfigurationAttributes {
        AssignPublicIp?: string;
        SecurityGroups?: string[];
        Subnets: string[];
    }
    
    export interface ServiceAwsVpcConfigurationProperties {
        AssignPublicIp?: pulumi.Input<string>;
        SecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
        Subnets: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface ServiceDeploymentConfigurationAttributes {
        MaximumPercent?: number;
        MinimumHealthyPercent?: number;
    }
    
    export interface ServiceDeploymentConfigurationProperties {
        MaximumPercent?: pulumi.Input<number>;
        MinimumHealthyPercent?: pulumi.Input<number>;
    }
    
    export interface ServiceLoadBalancerAttributes {
        ContainerName?: string;
        ContainerPort: number;
        LoadBalancerName?: string;
        TargetGroupArn?: string;
    }
    
    export interface ServiceLoadBalancerProperties {
        ContainerName?: pulumi.Input<string>;
        ContainerPort: pulumi.Input<number>;
        LoadBalancerName?: pulumi.Input<string>;
        TargetGroupArn?: pulumi.Input<string>;
    }
    
    export interface ServiceNetworkConfigurationAttributes {
        AwsvpcConfiguration?: ServiceAwsVpcConfigurationAttributes;
    }
    
    export interface ServiceNetworkConfigurationProperties {
        AwsvpcConfiguration?: pulumi.Input<ServiceAwsVpcConfigurationProperties>;
    }
    
    export interface ServicePlacementConstraintAttributes {
        Expression?: string;
        Type: string;
    }
    
    export interface ServicePlacementConstraintProperties {
        Expression?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface ServicePlacementStrategyAttributes {
        Field?: string;
        Type: string;
    }
    
    export interface ServicePlacementStrategyProperties {
        Field?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface ServiceProperties {
        Cluster?: pulumi.Input<string>;
        DeploymentConfiguration?: pulumi.Input<ServiceDeploymentConfigurationProperties>;
        DesiredCount?: pulumi.Input<number>;
        EnableECSManagedTags?: pulumi.Input<boolean>;
        HealthCheckGracePeriodSeconds?: pulumi.Input<number>;
        LaunchType?: pulumi.Input<string>;
        LoadBalancers?: pulumi.Input<pulumi.Input<ServiceLoadBalancerProperties>[]>;
        NetworkConfiguration?: pulumi.Input<ServiceNetworkConfigurationProperties>;
        PlacementConstraints?: pulumi.Input<pulumi.Input<ServicePlacementConstraintProperties>[]>;
        PlacementStrategies?: pulumi.Input<pulumi.Input<ServicePlacementStrategyProperties>[]>;
        PlatformVersion?: pulumi.Input<string>;
        PropagateTags?: pulumi.Input<string>;
        Role?: pulumi.Input<string>;
        SchedulingStrategy?: pulumi.Input<string>;
        ServiceName?: pulumi.Input<string>;
        ServiceRegistries?: pulumi.Input<pulumi.Input<ServiceServiceRegistryProperties>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        TaskDefinition: pulumi.Input<string>;
    }
    
    export interface ServiceServiceRegistryAttributes {
        ContainerName?: string;
        ContainerPort?: number;
        Port?: number;
        RegistryArn?: string;
    }
    
    export interface ServiceServiceRegistryProperties {
        ContainerName?: pulumi.Input<string>;
        ContainerPort?: pulumi.Input<number>;
        Port?: pulumi.Input<number>;
        RegistryArn?: pulumi.Input<string>;
    }
    
    export interface TaskDefinitionAttributes {
    }
    
    export interface TaskDefinitionContainerDefinitionAttributes {
        Command?: string[];
        Cpu?: number;
        DependsOn?: TaskDefinitionContainerDependencyAttributes[];
        DisableNetworking?: boolean;
        DnsSearchDomains?: string[];
        DnsServers?: string[];
        DockerLabels?: {[key: string]: string};
        DockerSecurityOptions?: string[];
        EntryPoint?: string[];
        Environment?: TaskDefinitionKeyValuePairAttributes[];
        Essential?: boolean;
        ExtraHosts?: TaskDefinitionHostEntryAttributes[];
        HealthCheck?: TaskDefinitionHealthCheckAttributes;
        Hostname?: string;
        Image?: string;
        Interactive?: boolean;
        Links?: string[];
        LinuxParameters?: TaskDefinitionLinuxParametersAttributes;
        LogConfiguration?: TaskDefinitionLogConfigurationAttributes;
        Memory?: number;
        MemoryReservation?: number;
        MountPoints?: TaskDefinitionMountPointAttributes[];
        Name?: string;
        PortMappings?: TaskDefinitionPortMappingAttributes[];
        Privileged?: boolean;
        PseudoTerminal?: boolean;
        ReadonlyRootFilesystem?: boolean;
        RepositoryCredentials?: TaskDefinitionRepositoryCredentialsAttributes;
        ResourceRequirements?: TaskDefinitionResourceRequirementAttributes[];
        Secrets?: TaskDefinitionSecretAttributes[];
        StartTimeout?: number;
        StopTimeout?: number;
        SystemControls?: TaskDefinitionSystemControlAttributes[];
        Ulimits?: TaskDefinitionUlimitAttributes[];
        User?: string;
        VolumesFrom?: TaskDefinitionVolumeFromAttributes[];
        WorkingDirectory?: string;
    }
    
    export interface TaskDefinitionContainerDefinitionProperties {
        Command?: pulumi.Input<pulumi.Input<string>[]>;
        Cpu?: pulumi.Input<number>;
        DependsOn?: pulumi.Input<pulumi.Input<TaskDefinitionContainerDependencyProperties>[]>;
        DisableNetworking?: pulumi.Input<boolean>;
        DnsSearchDomains?: pulumi.Input<pulumi.Input<string>[]>;
        DnsServers?: pulumi.Input<pulumi.Input<string>[]>;
        DockerLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        DockerSecurityOptions?: pulumi.Input<pulumi.Input<string>[]>;
        EntryPoint?: pulumi.Input<pulumi.Input<string>[]>;
        Environment?: pulumi.Input<pulumi.Input<TaskDefinitionKeyValuePairProperties>[]>;
        Essential?: pulumi.Input<boolean>;
        ExtraHosts?: pulumi.Input<pulumi.Input<TaskDefinitionHostEntryProperties>[]>;
        HealthCheck?: pulumi.Input<TaskDefinitionHealthCheckProperties>;
        Hostname?: pulumi.Input<string>;
        Image?: pulumi.Input<string>;
        Interactive?: pulumi.Input<boolean>;
        Links?: pulumi.Input<pulumi.Input<string>[]>;
        LinuxParameters?: pulumi.Input<TaskDefinitionLinuxParametersProperties>;
        LogConfiguration?: pulumi.Input<TaskDefinitionLogConfigurationProperties>;
        Memory?: pulumi.Input<number>;
        MemoryReservation?: pulumi.Input<number>;
        MountPoints?: pulumi.Input<pulumi.Input<TaskDefinitionMountPointProperties>[]>;
        Name?: pulumi.Input<string>;
        PortMappings?: pulumi.Input<pulumi.Input<TaskDefinitionPortMappingProperties>[]>;
        Privileged?: pulumi.Input<boolean>;
        PseudoTerminal?: pulumi.Input<boolean>;
        ReadonlyRootFilesystem?: pulumi.Input<boolean>;
        RepositoryCredentials?: pulumi.Input<TaskDefinitionRepositoryCredentialsProperties>;
        ResourceRequirements?: pulumi.Input<pulumi.Input<TaskDefinitionResourceRequirementProperties>[]>;
        Secrets?: pulumi.Input<pulumi.Input<TaskDefinitionSecretProperties>[]>;
        StartTimeout?: pulumi.Input<number>;
        StopTimeout?: pulumi.Input<number>;
        SystemControls?: pulumi.Input<pulumi.Input<TaskDefinitionSystemControlProperties>[]>;
        Ulimits?: pulumi.Input<pulumi.Input<TaskDefinitionUlimitProperties>[]>;
        User?: pulumi.Input<string>;
        VolumesFrom?: pulumi.Input<pulumi.Input<TaskDefinitionVolumeFromProperties>[]>;
        WorkingDirectory?: pulumi.Input<string>;
    }
    
    export interface TaskDefinitionContainerDependencyAttributes {
        Condition: string;
        ContainerName: string;
    }
    
    export interface TaskDefinitionContainerDependencyProperties {
        Condition: pulumi.Input<string>;
        ContainerName: pulumi.Input<string>;
    }
    
    export interface TaskDefinitionDeviceAttributes {
        ContainerPath?: string;
        HostPath: string;
        Permissions?: string[];
    }
    
    export interface TaskDefinitionDeviceProperties {
        ContainerPath?: pulumi.Input<string>;
        HostPath: pulumi.Input<string>;
        Permissions?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface TaskDefinitionDockerVolumeConfigurationAttributes {
        Autoprovision?: boolean;
        Driver?: string;
        DriverOpts?: {[key: string]: string};
        Labels?: {[key: string]: string};
        Scope?: string;
    }
    
    export interface TaskDefinitionDockerVolumeConfigurationProperties {
        Autoprovision?: pulumi.Input<boolean>;
        Driver?: pulumi.Input<string>;
        DriverOpts?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        Labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        Scope?: pulumi.Input<string>;
    }
    
    export interface TaskDefinitionHealthCheckAttributes {
        Command: string[];
        Interval?: number;
        Retries?: number;
        StartPeriod?: number;
        Timeout?: number;
    }
    
    export interface TaskDefinitionHealthCheckProperties {
        Command: pulumi.Input<pulumi.Input<string>[]>;
        Interval?: pulumi.Input<number>;
        Retries?: pulumi.Input<number>;
        StartPeriod?: pulumi.Input<number>;
        Timeout?: pulumi.Input<number>;
    }
    
    export interface TaskDefinitionHostEntryAttributes {
        Hostname: string;
        IpAddress: string;
    }
    
    export interface TaskDefinitionHostEntryProperties {
        Hostname: pulumi.Input<string>;
        IpAddress: pulumi.Input<string>;
    }
    
    export interface TaskDefinitionHostVolumePropertiesAttributes {
        SourcePath?: string;
    }
    
    export interface TaskDefinitionHostVolumePropertiesProperties {
        SourcePath?: pulumi.Input<string>;
    }
    
    export interface TaskDefinitionInferenceAcceleratorAttributes {
        DeviceName?: string;
        DevicePolicy?: string;
        DeviceType?: string;
    }
    
    export interface TaskDefinitionInferenceAcceleratorProperties {
        DeviceName?: pulumi.Input<string>;
        DevicePolicy?: pulumi.Input<string>;
        DeviceType?: pulumi.Input<string>;
    }
    
    export interface TaskDefinitionKernelCapabilitiesAttributes {
        Add?: string[];
        Drop?: string[];
    }
    
    export interface TaskDefinitionKernelCapabilitiesProperties {
        Add?: pulumi.Input<pulumi.Input<string>[]>;
        Drop?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface TaskDefinitionKeyValuePairAttributes {
        Name?: string;
        Value?: string;
    }
    
    export interface TaskDefinitionKeyValuePairProperties {
        Name?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface TaskDefinitionLinuxParametersAttributes {
        Capabilities?: TaskDefinitionKernelCapabilitiesAttributes;
        Devices?: TaskDefinitionDeviceAttributes[];
        InitProcessEnabled?: boolean;
        SharedMemorySize?: number;
        Tmpfs?: TaskDefinitionTmpfsAttributes[];
    }
    
    export interface TaskDefinitionLinuxParametersProperties {
        Capabilities?: pulumi.Input<TaskDefinitionKernelCapabilitiesProperties>;
        Devices?: pulumi.Input<pulumi.Input<TaskDefinitionDeviceProperties>[]>;
        InitProcessEnabled?: pulumi.Input<boolean>;
        SharedMemorySize?: pulumi.Input<number>;
        Tmpfs?: pulumi.Input<pulumi.Input<TaskDefinitionTmpfsProperties>[]>;
    }
    
    export interface TaskDefinitionLogConfigurationAttributes {
        LogDriver: string;
        Options?: {[key: string]: string};
        SecretOptions?: TaskDefinitionSecretAttributes[];
    }
    
    export interface TaskDefinitionLogConfigurationProperties {
        LogDriver: pulumi.Input<string>;
        Options?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        SecretOptions?: pulumi.Input<pulumi.Input<TaskDefinitionSecretProperties>[]>;
    }
    
    export interface TaskDefinitionMountPointAttributes {
        ContainerPath?: string;
        ReadOnly?: boolean;
        SourceVolume?: string;
    }
    
    export interface TaskDefinitionMountPointProperties {
        ContainerPath?: pulumi.Input<string>;
        ReadOnly?: pulumi.Input<boolean>;
        SourceVolume?: pulumi.Input<string>;
    }
    
    export interface TaskDefinitionPortMappingAttributes {
        ContainerPort?: number;
        HostPort?: number;
        Protocol?: string;
    }
    
    export interface TaskDefinitionPortMappingProperties {
        ContainerPort?: pulumi.Input<number>;
        HostPort?: pulumi.Input<number>;
        Protocol?: pulumi.Input<string>;
    }
    
    export interface TaskDefinitionProperties {
        ContainerDefinitions?: pulumi.Input<pulumi.Input<TaskDefinitionContainerDefinitionProperties>[]>;
        Cpu?: pulumi.Input<string>;
        ExecutionRoleArn?: pulumi.Input<string>;
        Family?: pulumi.Input<string>;
        InferenceAccelerators?: pulumi.Input<pulumi.Input<TaskDefinitionInferenceAcceleratorProperties>[]>;
        IpcMode?: pulumi.Input<string>;
        Memory?: pulumi.Input<string>;
        NetworkMode?: pulumi.Input<string>;
        PidMode?: pulumi.Input<string>;
        PlacementConstraints?: pulumi.Input<pulumi.Input<TaskDefinitionTaskDefinitionPlacementConstraintProperties>[]>;
        ProxyConfiguration?: pulumi.Input<TaskDefinitionProxyConfigurationProperties>;
        RequiresCompatibilities?: pulumi.Input<pulumi.Input<string>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        TaskRoleArn?: pulumi.Input<string>;
        Volumes?: pulumi.Input<pulumi.Input<TaskDefinitionVolumeProperties>[]>;
    }
    
    export interface TaskDefinitionProxyConfigurationAttributes {
        ContainerName: string;
        ProxyConfigurationProperties?: TaskDefinitionKeyValuePairAttributes[];
        Type?: string;
    }
    
    export interface TaskDefinitionProxyConfigurationProperties {
        ContainerName: pulumi.Input<string>;
        ProxyConfigurationProperties?: pulumi.Input<pulumi.Input<TaskDefinitionKeyValuePairProperties>[]>;
        Type?: pulumi.Input<string>;
    }
    
    export interface TaskDefinitionRepositoryCredentialsAttributes {
        CredentialsParameter?: string;
    }
    
    export interface TaskDefinitionRepositoryCredentialsProperties {
        CredentialsParameter?: pulumi.Input<string>;
    }
    
    export interface TaskDefinitionResourceRequirementAttributes {
        Type: string;
        Value: string;
    }
    
    export interface TaskDefinitionResourceRequirementProperties {
        Type: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface TaskDefinitionSecretAttributes {
        Name: string;
        ValueFrom: string;
    }
    
    export interface TaskDefinitionSecretProperties {
        Name: pulumi.Input<string>;
        ValueFrom: pulumi.Input<string>;
    }
    
    export interface TaskDefinitionSystemControlAttributes {
        Namespace: string;
        Value: string;
    }
    
    export interface TaskDefinitionSystemControlProperties {
        Namespace: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface TaskDefinitionTaskDefinitionPlacementConstraintAttributes {
        Expression?: string;
        Type: string;
    }
    
    export interface TaskDefinitionTaskDefinitionPlacementConstraintProperties {
        Expression?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface TaskDefinitionTmpfsAttributes {
        ContainerPath?: string;
        MountOptions?: string[];
        Size: number;
    }
    
    export interface TaskDefinitionTmpfsProperties {
        ContainerPath?: pulumi.Input<string>;
        MountOptions?: pulumi.Input<pulumi.Input<string>[]>;
        Size: pulumi.Input<number>;
    }
    
    export interface TaskDefinitionUlimitAttributes {
        HardLimit: number;
        Name: string;
        SoftLimit: number;
    }
    
    export interface TaskDefinitionUlimitProperties {
        HardLimit: pulumi.Input<number>;
        Name: pulumi.Input<string>;
        SoftLimit: pulumi.Input<number>;
    }
    
    export interface TaskDefinitionVolumeAttributes {
        DockerVolumeConfiguration?: TaskDefinitionDockerVolumeConfigurationAttributes;
        Host?: TaskDefinitionHostVolumePropertiesAttributes;
        Name?: string;
    }
    
    export interface TaskDefinitionVolumeFromAttributes {
        ReadOnly?: boolean;
        SourceContainer?: string;
    }
    
    export interface TaskDefinitionVolumeFromProperties {
        ReadOnly?: pulumi.Input<boolean>;
        SourceContainer?: pulumi.Input<string>;
    }
    
    export interface TaskDefinitionVolumeProperties {
        DockerVolumeConfiguration?: pulumi.Input<TaskDefinitionDockerVolumeConfigurationProperties>;
        Host?: pulumi.Input<TaskDefinitionHostVolumePropertiesProperties>;
        Name?: pulumi.Input<string>;
    }
    
    
    export class Cluster extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ClusterAttributes>;

        constructor(name: string, properties?: ClusterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ECS:Cluster", name, inputs, opts);
        }
    }
    
    export class Service extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ServiceAttributes>;

        constructor(name: string, properties: ServiceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ECS:Service", name, inputs, opts);
        }
    }
    
    export class TaskDefinition extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TaskDefinitionAttributes>;

        constructor(name: string, properties?: TaskDefinitionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ECS:TaskDefinition", name, inputs, opts);
        }
    }
    
}

export namespace efs {
    
    export interface FileSystemAttributes {
    }
    
    export interface FileSystemElasticFileSystemTagAttributes {
        Key: string;
        Value: string;
    }
    
    export interface FileSystemElasticFileSystemTagProperties {
        Key: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface FileSystemLifecyclePolicyAttributes {
        TransitionToIA: string;
    }
    
    export interface FileSystemLifecyclePolicyProperties {
        TransitionToIA: pulumi.Input<string>;
    }
    
    export interface FileSystemProperties {
        Encrypted?: pulumi.Input<boolean>;
        FileSystemTags?: pulumi.Input<pulumi.Input<FileSystemElasticFileSystemTagProperties>[]>;
        KmsKeyId?: pulumi.Input<string>;
        LifecyclePolicies?: pulumi.Input<pulumi.Input<FileSystemLifecyclePolicyProperties>[]>;
        PerformanceMode?: pulumi.Input<string>;
        ProvisionedThroughputInMibps?: pulumi.Input<number>;
        ThroughputMode?: pulumi.Input<string>;
    }
    
    export interface MountTargetAttributes {
        IpAddress: string;
    }
    
    export interface MountTargetProperties {
        FileSystemId: pulumi.Input<string>;
        IpAddress?: pulumi.Input<string>;
        SecurityGroups: pulumi.Input<pulumi.Input<string>[]>;
        SubnetId: pulumi.Input<string>;
    }
    
    
    export class FileSystem extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<FileSystemAttributes>;

        constructor(name: string, properties?: FileSystemProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EFS:FileSystem", name, inputs, opts);
        }
    }
    
    export class MountTarget extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<MountTargetAttributes>;

        constructor(name: string, properties: MountTargetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EFS:MountTarget", name, inputs, opts);
        }
    }
    
}

export namespace eks {
    
    export interface ClusterAttributes {
        Arn: string;
        CertificateAuthorityData: string;
        Endpoint: string;
    }
    
    export interface ClusterProperties {
        Name?: pulumi.Input<string>;
        ResourcesVpcConfig: pulumi.Input<ClusterResourcesVpcConfigProperties>;
        RoleArn: pulumi.Input<string>;
        Version?: pulumi.Input<string>;
    }
    
    export interface ClusterResourcesVpcConfigAttributes {
        SecurityGroupIds?: string[];
        SubnetIds: string[];
    }
    
    export interface ClusterResourcesVpcConfigProperties {
        SecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
        SubnetIds: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    
    export class Cluster extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ClusterAttributes>;

        constructor(name: string, properties: ClusterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EKS:Cluster", name, inputs, opts);
        }
    }
    
}

export namespace emr {
    
    export interface ClusterApplicationAttributes {
        AdditionalInfo?: {[key: string]: string};
        Args?: string[];
        Name?: string;
        Version?: string;
    }
    
    export interface ClusterApplicationProperties {
        AdditionalInfo?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        Args?: pulumi.Input<pulumi.Input<string>[]>;
        Name?: pulumi.Input<string>;
        Version?: pulumi.Input<string>;
    }
    
    export interface ClusterAttributes {
        MasterPublicDNS: string;
    }
    
    export interface ClusterAutoScalingPolicyAttributes {
        Constraints: ClusterScalingConstraintsAttributes;
        Rules: ClusterScalingRuleAttributes[];
    }
    
    export interface ClusterAutoScalingPolicyProperties {
        Constraints: pulumi.Input<ClusterScalingConstraintsProperties>;
        Rules: pulumi.Input<pulumi.Input<ClusterScalingRuleProperties>[]>;
    }
    
    export interface ClusterBootstrapActionConfigAttributes {
        Name: string;
        ScriptBootstrapAction: ClusterScriptBootstrapActionConfigAttributes;
    }
    
    export interface ClusterBootstrapActionConfigProperties {
        Name: pulumi.Input<string>;
        ScriptBootstrapAction: pulumi.Input<ClusterScriptBootstrapActionConfigProperties>;
    }
    
    export interface ClusterCloudWatchAlarmDefinitionAttributes {
        ComparisonOperator: string;
        Dimensions?: ClusterMetricDimensionAttributes[];
        EvaluationPeriods?: number;
        MetricName: string;
        Namespace?: string;
        Period: number;
        Statistic?: string;
        Threshold: number;
        Unit?: string;
    }
    
    export interface ClusterCloudWatchAlarmDefinitionProperties {
        ComparisonOperator: pulumi.Input<string>;
        Dimensions?: pulumi.Input<pulumi.Input<ClusterMetricDimensionProperties>[]>;
        EvaluationPeriods?: pulumi.Input<number>;
        MetricName: pulumi.Input<string>;
        Namespace?: pulumi.Input<string>;
        Period: pulumi.Input<number>;
        Statistic?: pulumi.Input<string>;
        Threshold: pulumi.Input<number>;
        Unit?: pulumi.Input<string>;
    }
    
    export interface ClusterConfigurationAttributes {
        Classification?: string;
        ConfigurationProperties?: {[key: string]: string};
        Configurations?: ClusterConfigurationAttributes[];
    }
    
    export interface ClusterConfigurationProperties {
        Classification?: pulumi.Input<string>;
        ConfigurationProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        Configurations?: pulumi.Input<pulumi.Input<ClusterConfigurationProperties>[]>;
    }
    
    export interface ClusterEbsBlockDeviceConfigAttributes {
        VolumeSpecification: ClusterVolumeSpecificationAttributes;
        VolumesPerInstance?: number;
    }
    
    export interface ClusterEbsBlockDeviceConfigProperties {
        VolumeSpecification: pulumi.Input<ClusterVolumeSpecificationProperties>;
        VolumesPerInstance?: pulumi.Input<number>;
    }
    
    export interface ClusterEbsConfigurationAttributes {
        EbsBlockDeviceConfigs?: ClusterEbsBlockDeviceConfigAttributes[];
        EbsOptimized?: boolean;
    }
    
    export interface ClusterEbsConfigurationProperties {
        EbsBlockDeviceConfigs?: pulumi.Input<pulumi.Input<ClusterEbsBlockDeviceConfigProperties>[]>;
        EbsOptimized?: pulumi.Input<boolean>;
    }
    
    export interface ClusterHadoopJarStepConfigAttributes {
        Args?: string[];
        Jar: string;
        MainClass?: string;
        StepProperties?: ClusterKeyValueAttributes[];
    }
    
    export interface ClusterHadoopJarStepConfigProperties {
        Args?: pulumi.Input<pulumi.Input<string>[]>;
        Jar: pulumi.Input<string>;
        MainClass?: pulumi.Input<string>;
        StepProperties?: pulumi.Input<pulumi.Input<ClusterKeyValueProperties>[]>;
    }
    
    export interface ClusterInstanceFleetConfigAttributes {
        InstanceTypeConfigs?: ClusterInstanceTypeConfigAttributes[];
        LaunchSpecifications?: ClusterInstanceFleetProvisioningSpecificationsAttributes;
        Name?: string;
        TargetOnDemandCapacity?: number;
        TargetSpotCapacity?: number;
    }
    
    export interface ClusterInstanceFleetConfigProperties {
        InstanceTypeConfigs?: pulumi.Input<pulumi.Input<ClusterInstanceTypeConfigProperties>[]>;
        LaunchSpecifications?: pulumi.Input<ClusterInstanceFleetProvisioningSpecificationsProperties>;
        Name?: pulumi.Input<string>;
        TargetOnDemandCapacity?: pulumi.Input<number>;
        TargetSpotCapacity?: pulumi.Input<number>;
    }
    
    export interface ClusterInstanceFleetProvisioningSpecificationsAttributes {
        SpotSpecification: ClusterSpotProvisioningSpecificationAttributes;
    }
    
    export interface ClusterInstanceFleetProvisioningSpecificationsProperties {
        SpotSpecification: pulumi.Input<ClusterSpotProvisioningSpecificationProperties>;
    }
    
    export interface ClusterInstanceGroupConfigAttributes {
        AutoScalingPolicy?: ClusterAutoScalingPolicyAttributes;
        BidPrice?: string;
        Configurations?: ClusterConfigurationAttributes[];
        EbsConfiguration?: ClusterEbsConfigurationAttributes;
        InstanceCount: number;
        InstanceType: string;
        Market?: string;
        Name?: string;
    }
    
    export interface ClusterInstanceGroupConfigProperties {
        AutoScalingPolicy?: pulumi.Input<ClusterAutoScalingPolicyProperties>;
        BidPrice?: pulumi.Input<string>;
        Configurations?: pulumi.Input<pulumi.Input<ClusterConfigurationProperties>[]>;
        EbsConfiguration?: pulumi.Input<ClusterEbsConfigurationProperties>;
        InstanceCount: pulumi.Input<number>;
        InstanceType: pulumi.Input<string>;
        Market?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
    }
    
    export interface ClusterInstanceTypeConfigAttributes {
        BidPrice?: string;
        BidPriceAsPercentageOfOnDemandPrice?: number;
        Configurations?: ClusterConfigurationAttributes[];
        EbsConfiguration?: ClusterEbsConfigurationAttributes;
        InstanceType: string;
        WeightedCapacity?: number;
    }
    
    export interface ClusterInstanceTypeConfigProperties {
        BidPrice?: pulumi.Input<string>;
        BidPriceAsPercentageOfOnDemandPrice?: pulumi.Input<number>;
        Configurations?: pulumi.Input<pulumi.Input<ClusterConfigurationProperties>[]>;
        EbsConfiguration?: pulumi.Input<ClusterEbsConfigurationProperties>;
        InstanceType: pulumi.Input<string>;
        WeightedCapacity?: pulumi.Input<number>;
    }
    
    export interface ClusterJobFlowInstancesConfigAttributes {
        AdditionalMasterSecurityGroups?: string[];
        AdditionalSlaveSecurityGroups?: string[];
        CoreInstanceFleet?: ClusterInstanceFleetConfigAttributes;
        CoreInstanceGroup?: ClusterInstanceGroupConfigAttributes;
        Ec2KeyName?: string;
        Ec2SubnetId?: string;
        Ec2SubnetIds?: string[];
        EmrManagedMasterSecurityGroup?: string;
        EmrManagedSlaveSecurityGroup?: string;
        HadoopVersion?: string;
        KeepJobFlowAliveWhenNoSteps?: boolean;
        MasterInstanceFleet?: ClusterInstanceFleetConfigAttributes;
        MasterInstanceGroup?: ClusterInstanceGroupConfigAttributes;
        Placement?: ClusterPlacementTypeAttributes;
        ServiceAccessSecurityGroup?: string;
        TerminationProtected?: boolean;
    }
    
    export interface ClusterJobFlowInstancesConfigProperties {
        AdditionalMasterSecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
        AdditionalSlaveSecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
        CoreInstanceFleet?: pulumi.Input<ClusterInstanceFleetConfigProperties>;
        CoreInstanceGroup?: pulumi.Input<ClusterInstanceGroupConfigProperties>;
        Ec2KeyName?: pulumi.Input<string>;
        Ec2SubnetId?: pulumi.Input<string>;
        Ec2SubnetIds?: pulumi.Input<pulumi.Input<string>[]>;
        EmrManagedMasterSecurityGroup?: pulumi.Input<string>;
        EmrManagedSlaveSecurityGroup?: pulumi.Input<string>;
        HadoopVersion?: pulumi.Input<string>;
        KeepJobFlowAliveWhenNoSteps?: pulumi.Input<boolean>;
        MasterInstanceFleet?: pulumi.Input<ClusterInstanceFleetConfigProperties>;
        MasterInstanceGroup?: pulumi.Input<ClusterInstanceGroupConfigProperties>;
        Placement?: pulumi.Input<ClusterPlacementTypeProperties>;
        ServiceAccessSecurityGroup?: pulumi.Input<string>;
        TerminationProtected?: pulumi.Input<boolean>;
    }
    
    export interface ClusterKerberosAttributesAttributes {
        ADDomainJoinPassword?: string;
        ADDomainJoinUser?: string;
        CrossRealmTrustPrincipalPassword?: string;
        KdcAdminPassword: string;
        Realm: string;
    }
    
    export interface ClusterKerberosAttributesProperties {
        ADDomainJoinPassword?: pulumi.Input<string>;
        ADDomainJoinUser?: pulumi.Input<string>;
        CrossRealmTrustPrincipalPassword?: pulumi.Input<string>;
        KdcAdminPassword: pulumi.Input<string>;
        Realm: pulumi.Input<string>;
    }
    
    export interface ClusterKeyValueAttributes {
        Key?: string;
        Value?: string;
    }
    
    export interface ClusterKeyValueProperties {
        Key?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface ClusterMetricDimensionAttributes {
        Key: string;
        Value: string;
    }
    
    export interface ClusterMetricDimensionProperties {
        Key: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface ClusterPlacementTypeAttributes {
        AvailabilityZone: string;
    }
    
    export interface ClusterPlacementTypeProperties {
        AvailabilityZone: pulumi.Input<string>;
    }
    
    export interface ClusterProperties {
        AdditionalInfo?: pulumi.Input<string>;
        Applications?: pulumi.Input<pulumi.Input<ClusterApplicationProperties>[]>;
        AutoScalingRole?: pulumi.Input<string>;
        BootstrapActions?: pulumi.Input<pulumi.Input<ClusterBootstrapActionConfigProperties>[]>;
        Configurations?: pulumi.Input<pulumi.Input<ClusterConfigurationProperties>[]>;
        CustomAmiId?: pulumi.Input<string>;
        EbsRootVolumeSize?: pulumi.Input<number>;
        Instances: pulumi.Input<ClusterJobFlowInstancesConfigProperties>;
        JobFlowRole: pulumi.Input<string>;
        KerberosAttributes?: pulumi.Input<ClusterKerberosAttributesProperties>;
        LogUri?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        ReleaseLabel?: pulumi.Input<string>;
        ScaleDownBehavior?: pulumi.Input<string>;
        SecurityConfiguration?: pulumi.Input<string>;
        ServiceRole: pulumi.Input<string>;
        Steps?: pulumi.Input<pulumi.Input<ClusterStepConfigProperties>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VisibleToAllUsers?: pulumi.Input<boolean>;
    }
    
    export interface ClusterScalingActionAttributes {
        Market?: string;
        SimpleScalingPolicyConfiguration: ClusterSimpleScalingPolicyConfigurationAttributes;
    }
    
    export interface ClusterScalingActionProperties {
        Market?: pulumi.Input<string>;
        SimpleScalingPolicyConfiguration: pulumi.Input<ClusterSimpleScalingPolicyConfigurationProperties>;
    }
    
    export interface ClusterScalingConstraintsAttributes {
        MaxCapacity: number;
        MinCapacity: number;
    }
    
    export interface ClusterScalingConstraintsProperties {
        MaxCapacity: pulumi.Input<number>;
        MinCapacity: pulumi.Input<number>;
    }
    
    export interface ClusterScalingRuleAttributes {
        Action: ClusterScalingActionAttributes;
        Description?: string;
        Name: string;
        Trigger: ClusterScalingTriggerAttributes;
    }
    
    export interface ClusterScalingRuleProperties {
        Action: pulumi.Input<ClusterScalingActionProperties>;
        Description?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        Trigger: pulumi.Input<ClusterScalingTriggerProperties>;
    }
    
    export interface ClusterScalingTriggerAttributes {
        CloudWatchAlarmDefinition: ClusterCloudWatchAlarmDefinitionAttributes;
    }
    
    export interface ClusterScalingTriggerProperties {
        CloudWatchAlarmDefinition: pulumi.Input<ClusterCloudWatchAlarmDefinitionProperties>;
    }
    
    export interface ClusterScriptBootstrapActionConfigAttributes {
        Args?: string[];
        Path: string;
    }
    
    export interface ClusterScriptBootstrapActionConfigProperties {
        Args?: pulumi.Input<pulumi.Input<string>[]>;
        Path: pulumi.Input<string>;
    }
    
    export interface ClusterSimpleScalingPolicyConfigurationAttributes {
        AdjustmentType?: string;
        CoolDown?: number;
        ScalingAdjustment: number;
    }
    
    export interface ClusterSimpleScalingPolicyConfigurationProperties {
        AdjustmentType?: pulumi.Input<string>;
        CoolDown?: pulumi.Input<number>;
        ScalingAdjustment: pulumi.Input<number>;
    }
    
    export interface ClusterSpotProvisioningSpecificationAttributes {
        BlockDurationMinutes?: number;
        TimeoutAction: string;
        TimeoutDurationMinutes: number;
    }
    
    export interface ClusterSpotProvisioningSpecificationProperties {
        BlockDurationMinutes?: pulumi.Input<number>;
        TimeoutAction: pulumi.Input<string>;
        TimeoutDurationMinutes: pulumi.Input<number>;
    }
    
    export interface ClusterStepConfigAttributes {
        ActionOnFailure?: string;
        HadoopJarStep: ClusterHadoopJarStepConfigAttributes;
        Name: string;
    }
    
    export interface ClusterStepConfigProperties {
        ActionOnFailure?: pulumi.Input<string>;
        HadoopJarStep: pulumi.Input<ClusterHadoopJarStepConfigProperties>;
        Name: pulumi.Input<string>;
    }
    
    export interface ClusterVolumeSpecificationAttributes {
        Iops?: number;
        SizeInGB: number;
        VolumeType: string;
    }
    
    export interface ClusterVolumeSpecificationProperties {
        Iops?: pulumi.Input<number>;
        SizeInGB: pulumi.Input<number>;
        VolumeType: pulumi.Input<string>;
    }
    
    export interface InstanceFleetConfigAttributes {
    }
    
    export interface InstanceFleetConfigConfigurationAttributes {
        Classification?: string;
        ConfigurationProperties?: {[key: string]: string};
        Configurations?: InstanceFleetConfigConfigurationAttributes[];
    }
    
    export interface InstanceFleetConfigConfigurationProperties {
        Classification?: pulumi.Input<string>;
        ConfigurationProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        Configurations?: pulumi.Input<pulumi.Input<InstanceFleetConfigConfigurationProperties>[]>;
    }
    
    export interface InstanceFleetConfigEbsBlockDeviceConfigAttributes {
        VolumeSpecification: InstanceFleetConfigVolumeSpecificationAttributes;
        VolumesPerInstance?: number;
    }
    
    export interface InstanceFleetConfigEbsBlockDeviceConfigProperties {
        VolumeSpecification: pulumi.Input<InstanceFleetConfigVolumeSpecificationProperties>;
        VolumesPerInstance?: pulumi.Input<number>;
    }
    
    export interface InstanceFleetConfigEbsConfigurationAttributes {
        EbsBlockDeviceConfigs?: InstanceFleetConfigEbsBlockDeviceConfigAttributes[];
        EbsOptimized?: boolean;
    }
    
    export interface InstanceFleetConfigEbsConfigurationProperties {
        EbsBlockDeviceConfigs?: pulumi.Input<pulumi.Input<InstanceFleetConfigEbsBlockDeviceConfigProperties>[]>;
        EbsOptimized?: pulumi.Input<boolean>;
    }
    
    export interface InstanceFleetConfigInstanceFleetProvisioningSpecificationsAttributes {
        SpotSpecification: InstanceFleetConfigSpotProvisioningSpecificationAttributes;
    }
    
    export interface InstanceFleetConfigInstanceFleetProvisioningSpecificationsProperties {
        SpotSpecification: pulumi.Input<InstanceFleetConfigSpotProvisioningSpecificationProperties>;
    }
    
    export interface InstanceFleetConfigInstanceTypeConfigAttributes {
        BidPrice?: string;
        BidPriceAsPercentageOfOnDemandPrice?: number;
        Configurations?: InstanceFleetConfigConfigurationAttributes[];
        EbsConfiguration?: InstanceFleetConfigEbsConfigurationAttributes;
        InstanceType: string;
        WeightedCapacity?: number;
    }
    
    export interface InstanceFleetConfigInstanceTypeConfigProperties {
        BidPrice?: pulumi.Input<string>;
        BidPriceAsPercentageOfOnDemandPrice?: pulumi.Input<number>;
        Configurations?: pulumi.Input<pulumi.Input<InstanceFleetConfigConfigurationProperties>[]>;
        EbsConfiguration?: pulumi.Input<InstanceFleetConfigEbsConfigurationProperties>;
        InstanceType: pulumi.Input<string>;
        WeightedCapacity?: pulumi.Input<number>;
    }
    
    export interface InstanceFleetConfigProperties {
        ClusterId: pulumi.Input<string>;
        InstanceFleetType: pulumi.Input<string>;
        InstanceTypeConfigs?: pulumi.Input<pulumi.Input<InstanceFleetConfigInstanceTypeConfigProperties>[]>;
        LaunchSpecifications?: pulumi.Input<InstanceFleetConfigInstanceFleetProvisioningSpecificationsProperties>;
        Name?: pulumi.Input<string>;
        TargetOnDemandCapacity?: pulumi.Input<number>;
        TargetSpotCapacity?: pulumi.Input<number>;
    }
    
    export interface InstanceFleetConfigSpotProvisioningSpecificationAttributes {
        BlockDurationMinutes?: number;
        TimeoutAction: string;
        TimeoutDurationMinutes: number;
    }
    
    export interface InstanceFleetConfigSpotProvisioningSpecificationProperties {
        BlockDurationMinutes?: pulumi.Input<number>;
        TimeoutAction: pulumi.Input<string>;
        TimeoutDurationMinutes: pulumi.Input<number>;
    }
    
    export interface InstanceFleetConfigVolumeSpecificationAttributes {
        Iops?: number;
        SizeInGB: number;
        VolumeType: string;
    }
    
    export interface InstanceFleetConfigVolumeSpecificationProperties {
        Iops?: pulumi.Input<number>;
        SizeInGB: pulumi.Input<number>;
        VolumeType: pulumi.Input<string>;
    }
    
    export interface InstanceGroupConfigAttributes {
    }
    
    export interface InstanceGroupConfigAutoScalingPolicyAttributes {
        Constraints: InstanceGroupConfigScalingConstraintsAttributes;
        Rules: InstanceGroupConfigScalingRuleAttributes[];
    }
    
    export interface InstanceGroupConfigAutoScalingPolicyProperties {
        Constraints: pulumi.Input<InstanceGroupConfigScalingConstraintsProperties>;
        Rules: pulumi.Input<pulumi.Input<InstanceGroupConfigScalingRuleProperties>[]>;
    }
    
    export interface InstanceGroupConfigCloudWatchAlarmDefinitionAttributes {
        ComparisonOperator: string;
        Dimensions?: InstanceGroupConfigMetricDimensionAttributes[];
        EvaluationPeriods?: number;
        MetricName: string;
        Namespace?: string;
        Period: number;
        Statistic?: string;
        Threshold: number;
        Unit?: string;
    }
    
    export interface InstanceGroupConfigCloudWatchAlarmDefinitionProperties {
        ComparisonOperator: pulumi.Input<string>;
        Dimensions?: pulumi.Input<pulumi.Input<InstanceGroupConfigMetricDimensionProperties>[]>;
        EvaluationPeriods?: pulumi.Input<number>;
        MetricName: pulumi.Input<string>;
        Namespace?: pulumi.Input<string>;
        Period: pulumi.Input<number>;
        Statistic?: pulumi.Input<string>;
        Threshold: pulumi.Input<number>;
        Unit?: pulumi.Input<string>;
    }
    
    export interface InstanceGroupConfigConfigurationAttributes {
        Classification?: string;
        ConfigurationProperties?: {[key: string]: string};
        Configurations?: InstanceGroupConfigConfigurationAttributes[];
    }
    
    export interface InstanceGroupConfigConfigurationProperties {
        Classification?: pulumi.Input<string>;
        ConfigurationProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        Configurations?: pulumi.Input<pulumi.Input<InstanceGroupConfigConfigurationProperties>[]>;
    }
    
    export interface InstanceGroupConfigEbsBlockDeviceConfigAttributes {
        VolumeSpecification: InstanceGroupConfigVolumeSpecificationAttributes;
        VolumesPerInstance?: number;
    }
    
    export interface InstanceGroupConfigEbsBlockDeviceConfigProperties {
        VolumeSpecification: pulumi.Input<InstanceGroupConfigVolumeSpecificationProperties>;
        VolumesPerInstance?: pulumi.Input<number>;
    }
    
    export interface InstanceGroupConfigEbsConfigurationAttributes {
        EbsBlockDeviceConfigs?: InstanceGroupConfigEbsBlockDeviceConfigAttributes[];
        EbsOptimized?: boolean;
    }
    
    export interface InstanceGroupConfigEbsConfigurationProperties {
        EbsBlockDeviceConfigs?: pulumi.Input<pulumi.Input<InstanceGroupConfigEbsBlockDeviceConfigProperties>[]>;
        EbsOptimized?: pulumi.Input<boolean>;
    }
    
    export interface InstanceGroupConfigMetricDimensionAttributes {
        Key: string;
        Value: string;
    }
    
    export interface InstanceGroupConfigMetricDimensionProperties {
        Key: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface InstanceGroupConfigProperties {
        AutoScalingPolicy?: pulumi.Input<InstanceGroupConfigAutoScalingPolicyProperties>;
        BidPrice?: pulumi.Input<string>;
        Configurations?: pulumi.Input<pulumi.Input<InstanceGroupConfigConfigurationProperties>[]>;
        EbsConfiguration?: pulumi.Input<InstanceGroupConfigEbsConfigurationProperties>;
        InstanceCount: pulumi.Input<number>;
        InstanceRole: pulumi.Input<string>;
        InstanceType: pulumi.Input<string>;
        JobFlowId: pulumi.Input<string>;
        Market?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
    }
    
    export interface InstanceGroupConfigScalingActionAttributes {
        Market?: string;
        SimpleScalingPolicyConfiguration: InstanceGroupConfigSimpleScalingPolicyConfigurationAttributes;
    }
    
    export interface InstanceGroupConfigScalingActionProperties {
        Market?: pulumi.Input<string>;
        SimpleScalingPolicyConfiguration: pulumi.Input<InstanceGroupConfigSimpleScalingPolicyConfigurationProperties>;
    }
    
    export interface InstanceGroupConfigScalingConstraintsAttributes {
        MaxCapacity: number;
        MinCapacity: number;
    }
    
    export interface InstanceGroupConfigScalingConstraintsProperties {
        MaxCapacity: pulumi.Input<number>;
        MinCapacity: pulumi.Input<number>;
    }
    
    export interface InstanceGroupConfigScalingRuleAttributes {
        Action: InstanceGroupConfigScalingActionAttributes;
        Description?: string;
        Name: string;
        Trigger: InstanceGroupConfigScalingTriggerAttributes;
    }
    
    export interface InstanceGroupConfigScalingRuleProperties {
        Action: pulumi.Input<InstanceGroupConfigScalingActionProperties>;
        Description?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        Trigger: pulumi.Input<InstanceGroupConfigScalingTriggerProperties>;
    }
    
    export interface InstanceGroupConfigScalingTriggerAttributes {
        CloudWatchAlarmDefinition: InstanceGroupConfigCloudWatchAlarmDefinitionAttributes;
    }
    
    export interface InstanceGroupConfigScalingTriggerProperties {
        CloudWatchAlarmDefinition: pulumi.Input<InstanceGroupConfigCloudWatchAlarmDefinitionProperties>;
    }
    
    export interface InstanceGroupConfigSimpleScalingPolicyConfigurationAttributes {
        AdjustmentType?: string;
        CoolDown?: number;
        ScalingAdjustment: number;
    }
    
    export interface InstanceGroupConfigSimpleScalingPolicyConfigurationProperties {
        AdjustmentType?: pulumi.Input<string>;
        CoolDown?: pulumi.Input<number>;
        ScalingAdjustment: pulumi.Input<number>;
    }
    
    export interface InstanceGroupConfigVolumeSpecificationAttributes {
        Iops?: number;
        SizeInGB: number;
        VolumeType: string;
    }
    
    export interface InstanceGroupConfigVolumeSpecificationProperties {
        Iops?: pulumi.Input<number>;
        SizeInGB: pulumi.Input<number>;
        VolumeType: pulumi.Input<string>;
    }
    
    export interface SecurityConfigurationAttributes {
    }
    
    export interface SecurityConfigurationProperties {
        Name?: pulumi.Input<string>;
        SecurityConfiguration: pulumi.Input<string>;
    }
    
    export interface StepAttributes {
    }
    
    export interface StepHadoopJarStepConfigAttributes {
        Args?: string[];
        Jar: string;
        MainClass?: string;
        StepProperties?: StepKeyValueAttributes[];
    }
    
    export interface StepHadoopJarStepConfigProperties {
        Args?: pulumi.Input<pulumi.Input<string>[]>;
        Jar: pulumi.Input<string>;
        MainClass?: pulumi.Input<string>;
        StepProperties?: pulumi.Input<pulumi.Input<StepKeyValueProperties>[]>;
    }
    
    export interface StepKeyValueAttributes {
        Key?: string;
        Value?: string;
    }
    
    export interface StepKeyValueProperties {
        Key?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface StepProperties {
        ActionOnFailure: pulumi.Input<string>;
        HadoopJarStep: pulumi.Input<StepHadoopJarStepConfigProperties>;
        JobFlowId: pulumi.Input<string>;
        Name: pulumi.Input<string>;
    }
    
    
    export class Cluster extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ClusterAttributes>;

        constructor(name: string, properties: ClusterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EMR:Cluster", name, inputs, opts);
        }
    }
    
    export class InstanceFleetConfig extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<InstanceFleetConfigAttributes>;

        constructor(name: string, properties: InstanceFleetConfigProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EMR:InstanceFleetConfig", name, inputs, opts);
        }
    }
    
    export class InstanceGroupConfig extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<InstanceGroupConfigAttributes>;

        constructor(name: string, properties: InstanceGroupConfigProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EMR:InstanceGroupConfig", name, inputs, opts);
        }
    }
    
    export class SecurityConfiguration extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SecurityConfigurationAttributes>;

        constructor(name: string, properties: SecurityConfigurationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EMR:SecurityConfiguration", name, inputs, opts);
        }
    }
    
    export class Step extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<StepAttributes>;

        constructor(name: string, properties: StepProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:EMR:Step", name, inputs, opts);
        }
    }
    
}

export namespace elasticache {
    
    export interface CacheClusterAttributes {
        "ConfigurationEndpoint.Address": string;
        "ConfigurationEndpoint.Port": string;
        "RedisEndpoint.Address": string;
        "RedisEndpoint.Port": string;
    }
    
    export interface CacheClusterProperties {
        AZMode?: pulumi.Input<string>;
        AutoMinorVersionUpgrade?: pulumi.Input<boolean>;
        CacheNodeType: pulumi.Input<string>;
        CacheParameterGroupName?: pulumi.Input<string>;
        CacheSecurityGroupNames?: pulumi.Input<pulumi.Input<string>[]>;
        CacheSubnetGroupName?: pulumi.Input<string>;
        ClusterName?: pulumi.Input<string>;
        Engine: pulumi.Input<string>;
        EngineVersion?: pulumi.Input<string>;
        NotificationTopicArn?: pulumi.Input<string>;
        NumCacheNodes: pulumi.Input<number>;
        Port?: pulumi.Input<number>;
        PreferredAvailabilityZone?: pulumi.Input<string>;
        PreferredAvailabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
        PreferredMaintenanceWindow?: pulumi.Input<string>;
        SnapshotArns?: pulumi.Input<pulumi.Input<string>[]>;
        SnapshotName?: pulumi.Input<string>;
        SnapshotRetentionLimit?: pulumi.Input<number>;
        SnapshotWindow?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface ParameterGroupAttributes {
    }
    
    export interface ParameterGroupProperties {
        CacheParameterGroupFamily: pulumi.Input<string>;
        Description: pulumi.Input<string>;
        Properties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    }
    
    export interface ReplicationGroupAttributes {
        "ConfigurationEndPoint.Address": string;
        "ConfigurationEndPoint.Port": string;
        "PrimaryEndPoint.Address": string;
        "PrimaryEndPoint.Port": string;
        "ReadEndPoint.Addresses": string;
        "ReadEndPoint.Addresses.List": string[];
        "ReadEndPoint.Ports": string;
        "ReadEndPoint.Ports.List": string[];
    }
    
    export interface ReplicationGroupNodeGroupConfigurationAttributes {
        NodeGroupId?: string;
        PrimaryAvailabilityZone?: string;
        ReplicaAvailabilityZones?: string[];
        ReplicaCount?: number;
        Slots?: string;
    }
    
    export interface ReplicationGroupNodeGroupConfigurationProperties {
        NodeGroupId?: pulumi.Input<string>;
        PrimaryAvailabilityZone?: pulumi.Input<string>;
        ReplicaAvailabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
        ReplicaCount?: pulumi.Input<number>;
        Slots?: pulumi.Input<string>;
    }
    
    export interface ReplicationGroupProperties {
        AtRestEncryptionEnabled?: pulumi.Input<boolean>;
        AuthToken?: pulumi.Input<string>;
        AutoMinorVersionUpgrade?: pulumi.Input<boolean>;
        AutomaticFailoverEnabled?: pulumi.Input<boolean>;
        CacheNodeType?: pulumi.Input<string>;
        CacheParameterGroupName?: pulumi.Input<string>;
        CacheSecurityGroupNames?: pulumi.Input<pulumi.Input<string>[]>;
        CacheSubnetGroupName?: pulumi.Input<string>;
        Engine?: pulumi.Input<string>;
        EngineVersion?: pulumi.Input<string>;
        KmsKeyId?: pulumi.Input<string>;
        NodeGroupConfiguration?: pulumi.Input<pulumi.Input<ReplicationGroupNodeGroupConfigurationProperties>[]>;
        NotificationTopicArn?: pulumi.Input<string>;
        NumCacheClusters?: pulumi.Input<number>;
        NumNodeGroups?: pulumi.Input<number>;
        Port?: pulumi.Input<number>;
        PreferredCacheClusterAZs?: pulumi.Input<pulumi.Input<string>[]>;
        PreferredMaintenanceWindow?: pulumi.Input<string>;
        PrimaryClusterId?: pulumi.Input<string>;
        ReplicasPerNodeGroup?: pulumi.Input<number>;
        ReplicationGroupDescription: pulumi.Input<string>;
        ReplicationGroupId?: pulumi.Input<string>;
        SecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
        SnapshotArns?: pulumi.Input<pulumi.Input<string>[]>;
        SnapshotName?: pulumi.Input<string>;
        SnapshotRetentionLimit?: pulumi.Input<number>;
        SnapshotWindow?: pulumi.Input<string>;
        SnapshottingClusterId?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        TransitEncryptionEnabled?: pulumi.Input<boolean>;
    }
    
    export interface SecurityGroupAttributes {
    }
    
    export interface SecurityGroupIngressAttributes {
    }
    
    export interface SecurityGroupIngressProperties {
        CacheSecurityGroupName: pulumi.Input<string>;
        EC2SecurityGroupName: pulumi.Input<string>;
        EC2SecurityGroupOwnerId?: pulumi.Input<string>;
    }
    
    export interface SecurityGroupProperties {
        Description: pulumi.Input<string>;
    }
    
    export interface SubnetGroupAttributes {
    }
    
    export interface SubnetGroupProperties {
        CacheSubnetGroupName?: pulumi.Input<string>;
        Description: pulumi.Input<string>;
        SubnetIds: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    
    export class CacheCluster extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<CacheClusterAttributes>;

        constructor(name: string, properties: CacheClusterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ElastiCache:CacheCluster", name, inputs, opts);
        }
    }
    
    export class ParameterGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ParameterGroupAttributes>;

        constructor(name: string, properties: ParameterGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ElastiCache:ParameterGroup", name, inputs, opts);
        }
    }
    
    export class ReplicationGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ReplicationGroupAttributes>;

        constructor(name: string, properties: ReplicationGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ElastiCache:ReplicationGroup", name, inputs, opts);
        }
    }
    
    export class SecurityGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SecurityGroupAttributes>;

        constructor(name: string, properties: SecurityGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ElastiCache:SecurityGroup", name, inputs, opts);
        }
    }
    
    export class SecurityGroupIngress extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SecurityGroupIngressAttributes>;

        constructor(name: string, properties: SecurityGroupIngressProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ElastiCache:SecurityGroupIngress", name, inputs, opts);
        }
    }
    
    export class SubnetGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SubnetGroupAttributes>;

        constructor(name: string, properties: SubnetGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ElastiCache:SubnetGroup", name, inputs, opts);
        }
    }
    
}

export namespace elasticbeanstalk {
    
    export interface ApplicationApplicationResourceLifecycleConfigAttributes {
        ServiceRole?: string;
        VersionLifecycleConfig?: ApplicationApplicationVersionLifecycleConfigAttributes;
    }
    
    export interface ApplicationApplicationResourceLifecycleConfigProperties {
        ServiceRole?: pulumi.Input<string>;
        VersionLifecycleConfig?: pulumi.Input<ApplicationApplicationVersionLifecycleConfigProperties>;
    }
    
    export interface ApplicationApplicationVersionLifecycleConfigAttributes {
        MaxAgeRule?: ApplicationMaxAgeRuleAttributes;
        MaxCountRule?: ApplicationMaxCountRuleAttributes;
    }
    
    export interface ApplicationApplicationVersionLifecycleConfigProperties {
        MaxAgeRule?: pulumi.Input<ApplicationMaxAgeRuleProperties>;
        MaxCountRule?: pulumi.Input<ApplicationMaxCountRuleProperties>;
    }
    
    export interface ApplicationAttributes {
    }
    
    export interface ApplicationMaxAgeRuleAttributes {
        DeleteSourceFromS3?: boolean;
        Enabled?: boolean;
        MaxAgeInDays?: number;
    }
    
    export interface ApplicationMaxAgeRuleProperties {
        DeleteSourceFromS3?: pulumi.Input<boolean>;
        Enabled?: pulumi.Input<boolean>;
        MaxAgeInDays?: pulumi.Input<number>;
    }
    
    export interface ApplicationMaxCountRuleAttributes {
        DeleteSourceFromS3?: boolean;
        Enabled?: boolean;
        MaxCount?: number;
    }
    
    export interface ApplicationMaxCountRuleProperties {
        DeleteSourceFromS3?: pulumi.Input<boolean>;
        Enabled?: pulumi.Input<boolean>;
        MaxCount?: pulumi.Input<number>;
    }
    
    export interface ApplicationProperties {
        ApplicationName?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        ResourceLifecycleConfig?: pulumi.Input<ApplicationApplicationResourceLifecycleConfigProperties>;
    }
    
    export interface ApplicationVersionAttributes {
    }
    
    export interface ApplicationVersionProperties {
        ApplicationName: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        SourceBundle: pulumi.Input<ApplicationVersionSourceBundleProperties>;
    }
    
    export interface ApplicationVersionSourceBundleAttributes {
        S3Bucket: string;
        S3Key: string;
    }
    
    export interface ApplicationVersionSourceBundleProperties {
        S3Bucket: pulumi.Input<string>;
        S3Key: pulumi.Input<string>;
    }
    
    export interface ConfigurationTemplateAttributes {
    }
    
    export interface ConfigurationTemplateConfigurationOptionSettingAttributes {
        Namespace: string;
        OptionName: string;
        ResourceName?: string;
        Value?: string;
    }
    
    export interface ConfigurationTemplateConfigurationOptionSettingProperties {
        Namespace: pulumi.Input<string>;
        OptionName: pulumi.Input<string>;
        ResourceName?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface ConfigurationTemplateProperties {
        ApplicationName: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        EnvironmentId?: pulumi.Input<string>;
        OptionSettings?: pulumi.Input<pulumi.Input<ConfigurationTemplateConfigurationOptionSettingProperties>[]>;
        PlatformArn?: pulumi.Input<string>;
        SolutionStackName?: pulumi.Input<string>;
        SourceConfiguration?: pulumi.Input<ConfigurationTemplateSourceConfigurationProperties>;
    }
    
    export interface ConfigurationTemplateSourceConfigurationAttributes {
        ApplicationName: string;
        TemplateName: string;
    }
    
    export interface ConfigurationTemplateSourceConfigurationProperties {
        ApplicationName: pulumi.Input<string>;
        TemplateName: pulumi.Input<string>;
    }
    
    export interface EnvironmentAttributes {
        EndpointURL: string;
    }
    
    export interface EnvironmentOptionSettingAttributes {
        Namespace: string;
        OptionName: string;
        ResourceName?: string;
        Value?: string;
    }
    
    export interface EnvironmentOptionSettingProperties {
        Namespace: pulumi.Input<string>;
        OptionName: pulumi.Input<string>;
        ResourceName?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface EnvironmentProperties {
        ApplicationName: pulumi.Input<string>;
        CNAMEPrefix?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        EnvironmentName?: pulumi.Input<string>;
        OptionSettings?: pulumi.Input<pulumi.Input<EnvironmentOptionSettingProperties>[]>;
        PlatformArn?: pulumi.Input<string>;
        SolutionStackName?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        TemplateName?: pulumi.Input<string>;
        Tier?: pulumi.Input<EnvironmentTierProperties>;
        VersionLabel?: pulumi.Input<string>;
    }
    
    export interface EnvironmentTierAttributes {
        Name?: string;
        Type?: string;
        Version?: string;
    }
    
    export interface EnvironmentTierProperties {
        Name?: pulumi.Input<string>;
        Type?: pulumi.Input<string>;
        Version?: pulumi.Input<string>;
    }
    
    
    export class Application extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ApplicationAttributes>;

        constructor(name: string, properties?: ApplicationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ElasticBeanstalk:Application", name, inputs, opts);
        }
    }
    
    export class ApplicationVersion extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ApplicationVersionAttributes>;

        constructor(name: string, properties: ApplicationVersionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ElasticBeanstalk:ApplicationVersion", name, inputs, opts);
        }
    }
    
    export class ConfigurationTemplate extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ConfigurationTemplateAttributes>;

        constructor(name: string, properties: ConfigurationTemplateProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ElasticBeanstalk:ConfigurationTemplate", name, inputs, opts);
        }
    }
    
    export class Environment extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<EnvironmentAttributes>;

        constructor(name: string, properties: EnvironmentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ElasticBeanstalk:Environment", name, inputs, opts);
        }
    }
    
}

export namespace elasticloadbalancing {
    
    export interface LoadBalancerAccessLoggingPolicyAttributes {
        EmitInterval?: number;
        Enabled: boolean;
        S3BucketName: string;
        S3BucketPrefix?: string;
    }
    
    export interface LoadBalancerAccessLoggingPolicyProperties {
        EmitInterval?: pulumi.Input<number>;
        Enabled: pulumi.Input<boolean>;
        S3BucketName: pulumi.Input<string>;
        S3BucketPrefix?: pulumi.Input<string>;
    }
    
    export interface LoadBalancerAppCookieStickinessPolicyAttributes {
        CookieName: string;
        PolicyName: string;
    }
    
    export interface LoadBalancerAppCookieStickinessPolicyProperties {
        CookieName: pulumi.Input<string>;
        PolicyName: pulumi.Input<string>;
    }
    
    export interface LoadBalancerAttributes {
        CanonicalHostedZoneName: string;
        CanonicalHostedZoneNameID: string;
        DNSName: string;
        "SourceSecurityGroup.GroupName": string;
        "SourceSecurityGroup.OwnerAlias": string;
    }
    
    export interface LoadBalancerConnectionDrainingPolicyAttributes {
        Enabled: boolean;
        Timeout?: number;
    }
    
    export interface LoadBalancerConnectionDrainingPolicyProperties {
        Enabled: pulumi.Input<boolean>;
        Timeout?: pulumi.Input<number>;
    }
    
    export interface LoadBalancerConnectionSettingsAttributes {
        IdleTimeout: number;
    }
    
    export interface LoadBalancerConnectionSettingsProperties {
        IdleTimeout: pulumi.Input<number>;
    }
    
    export interface LoadBalancerHealthCheckAttributes {
        HealthyThreshold: string;
        Interval: string;
        Target: string;
        Timeout: string;
        UnhealthyThreshold: string;
    }
    
    export interface LoadBalancerHealthCheckProperties {
        HealthyThreshold: pulumi.Input<string>;
        Interval: pulumi.Input<string>;
        Target: pulumi.Input<string>;
        Timeout: pulumi.Input<string>;
        UnhealthyThreshold: pulumi.Input<string>;
    }
    
    export interface LoadBalancerLBCookieStickinessPolicyAttributes {
        CookieExpirationPeriod?: string;
        PolicyName?: string;
    }
    
    export interface LoadBalancerLBCookieStickinessPolicyProperties {
        CookieExpirationPeriod?: pulumi.Input<string>;
        PolicyName?: pulumi.Input<string>;
    }
    
    export interface LoadBalancerListenersAttributes {
        InstancePort: string;
        InstanceProtocol?: string;
        LoadBalancerPort: string;
        PolicyNames?: string[];
        Protocol: string;
        SSLCertificateId?: string;
    }
    
    export interface LoadBalancerListenersProperties {
        InstancePort: pulumi.Input<string>;
        InstanceProtocol?: pulumi.Input<string>;
        LoadBalancerPort: pulumi.Input<string>;
        PolicyNames?: pulumi.Input<pulumi.Input<string>[]>;
        Protocol: pulumi.Input<string>;
        SSLCertificateId?: pulumi.Input<string>;
    }
    
    export interface LoadBalancerPoliciesAttributes {
        Attributes: string[];
        InstancePorts?: string[];
        LoadBalancerPorts?: string[];
        PolicyName: string;
        PolicyType: string;
    }
    
    export interface LoadBalancerPoliciesProperties {
        Attributes: pulumi.Input<pulumi.Input<string>[]>;
        InstancePorts?: pulumi.Input<pulumi.Input<string>[]>;
        LoadBalancerPorts?: pulumi.Input<pulumi.Input<string>[]>;
        PolicyName: pulumi.Input<string>;
        PolicyType: pulumi.Input<string>;
    }
    
    export interface LoadBalancerProperties {
        AccessLoggingPolicy?: pulumi.Input<LoadBalancerAccessLoggingPolicyProperties>;
        AppCookieStickinessPolicy?: pulumi.Input<pulumi.Input<LoadBalancerAppCookieStickinessPolicyProperties>[]>;
        AvailabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
        ConnectionDrainingPolicy?: pulumi.Input<LoadBalancerConnectionDrainingPolicyProperties>;
        ConnectionSettings?: pulumi.Input<LoadBalancerConnectionSettingsProperties>;
        CrossZone?: pulumi.Input<boolean>;
        HealthCheck?: pulumi.Input<LoadBalancerHealthCheckProperties>;
        Instances?: pulumi.Input<pulumi.Input<string>[]>;
        LBCookieStickinessPolicy?: pulumi.Input<pulumi.Input<LoadBalancerLBCookieStickinessPolicyProperties>[]>;
        Listeners: pulumi.Input<pulumi.Input<LoadBalancerListenersProperties>[]>;
        LoadBalancerName?: pulumi.Input<string>;
        Policies?: pulumi.Input<pulumi.Input<LoadBalancerPoliciesProperties>[]>;
        Scheme?: pulumi.Input<string>;
        SecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
        Subnets?: pulumi.Input<pulumi.Input<string>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    
    export class LoadBalancer extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<LoadBalancerAttributes>;

        constructor(name: string, properties: LoadBalancerProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ElasticLoadBalancing:LoadBalancer", name, inputs, opts);
        }
    }
    
}

export namespace elasticloadbalancingv2 {
    
    export interface ListenerActionAttributes {
        AuthenticateCognitoConfig?: ListenerAuthenticateCognitoConfigAttributes;
        AuthenticateOidcConfig?: ListenerAuthenticateOidcConfigAttributes;
        FixedResponseConfig?: ListenerFixedResponseConfigAttributes;
        Order?: number;
        RedirectConfig?: ListenerRedirectConfigAttributes;
        TargetGroupArn?: string;
        Type: string;
    }
    
    export interface ListenerActionProperties {
        AuthenticateCognitoConfig?: pulumi.Input<ListenerAuthenticateCognitoConfigProperties>;
        AuthenticateOidcConfig?: pulumi.Input<ListenerAuthenticateOidcConfigProperties>;
        FixedResponseConfig?: pulumi.Input<ListenerFixedResponseConfigProperties>;
        Order?: pulumi.Input<number>;
        RedirectConfig?: pulumi.Input<ListenerRedirectConfigProperties>;
        TargetGroupArn?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface ListenerAttributes {
    }
    
    export interface ListenerAuthenticateCognitoConfigAttributes {
        AuthenticationRequestExtraParams?: {[key: string]: string};
        OnUnauthenticatedRequest?: string;
        Scope?: string;
        SessionCookieName?: string;
        SessionTimeout?: number;
        UserPoolArn: string;
        UserPoolClientId: string;
        UserPoolDomain: string;
    }
    
    export interface ListenerAuthenticateCognitoConfigProperties {
        AuthenticationRequestExtraParams?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        OnUnauthenticatedRequest?: pulumi.Input<string>;
        Scope?: pulumi.Input<string>;
        SessionCookieName?: pulumi.Input<string>;
        SessionTimeout?: pulumi.Input<number>;
        UserPoolArn: pulumi.Input<string>;
        UserPoolClientId: pulumi.Input<string>;
        UserPoolDomain: pulumi.Input<string>;
    }
    
    export interface ListenerAuthenticateOidcConfigAttributes {
        AuthenticationRequestExtraParams?: {[key: string]: string};
        AuthorizationEndpoint: string;
        ClientId: string;
        ClientSecret: string;
        Issuer: string;
        OnUnauthenticatedRequest?: string;
        Scope?: string;
        SessionCookieName?: string;
        SessionTimeout?: number;
        TokenEndpoint: string;
        UserInfoEndpoint: string;
    }
    
    export interface ListenerAuthenticateOidcConfigProperties {
        AuthenticationRequestExtraParams?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        AuthorizationEndpoint: pulumi.Input<string>;
        ClientId: pulumi.Input<string>;
        ClientSecret: pulumi.Input<string>;
        Issuer: pulumi.Input<string>;
        OnUnauthenticatedRequest?: pulumi.Input<string>;
        Scope?: pulumi.Input<string>;
        SessionCookieName?: pulumi.Input<string>;
        SessionTimeout?: pulumi.Input<number>;
        TokenEndpoint: pulumi.Input<string>;
        UserInfoEndpoint: pulumi.Input<string>;
    }
    
    export interface ListenerCertificateAttributes {
    }
    
    export interface ListenerCertificateCertificateAttributes {
        CertificateArn?: string;
    }
    
    export interface ListenerCertificateCertificateProperties {
        CertificateArn?: pulumi.Input<string>;
    }
    
    export interface ListenerCertificateProperties {
        Certificates: pulumi.Input<pulumi.Input<ListenerCertificateCertificateProperties>[]>;
        ListenerArn: pulumi.Input<string>;
    }
    
    export interface ListenerFixedResponseConfigAttributes {
        ContentType?: string;
        MessageBody?: string;
        StatusCode: string;
    }
    
    export interface ListenerFixedResponseConfigProperties {
        ContentType?: pulumi.Input<string>;
        MessageBody?: pulumi.Input<string>;
        StatusCode: pulumi.Input<string>;
    }
    
    export interface ListenerProperties {
        Certificates?: pulumi.Input<pulumi.Input<ListenerCertificateProperties>[]>;
        DefaultActions: pulumi.Input<pulumi.Input<ListenerActionProperties>[]>;
        LoadBalancerArn: pulumi.Input<string>;
        Port: pulumi.Input<number>;
        Protocol: pulumi.Input<string>;
        SslPolicy?: pulumi.Input<string>;
    }
    
    export interface ListenerRedirectConfigAttributes {
        Host?: string;
        Path?: string;
        Port?: string;
        Protocol?: string;
        Query?: string;
        StatusCode: string;
    }
    
    export interface ListenerRedirectConfigProperties {
        Host?: pulumi.Input<string>;
        Path?: pulumi.Input<string>;
        Port?: pulumi.Input<string>;
        Protocol?: pulumi.Input<string>;
        Query?: pulumi.Input<string>;
        StatusCode: pulumi.Input<string>;
    }
    
    export interface ListenerRuleActionAttributes {
        AuthenticateCognitoConfig?: ListenerRuleAuthenticateCognitoConfigAttributes;
        AuthenticateOidcConfig?: ListenerRuleAuthenticateOidcConfigAttributes;
        FixedResponseConfig?: ListenerRuleFixedResponseConfigAttributes;
        Order?: number;
        RedirectConfig?: ListenerRuleRedirectConfigAttributes;
        TargetGroupArn?: string;
        Type: string;
    }
    
    export interface ListenerRuleActionProperties {
        AuthenticateCognitoConfig?: pulumi.Input<ListenerRuleAuthenticateCognitoConfigProperties>;
        AuthenticateOidcConfig?: pulumi.Input<ListenerRuleAuthenticateOidcConfigProperties>;
        FixedResponseConfig?: pulumi.Input<ListenerRuleFixedResponseConfigProperties>;
        Order?: pulumi.Input<number>;
        RedirectConfig?: pulumi.Input<ListenerRuleRedirectConfigProperties>;
        TargetGroupArn?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface ListenerRuleAttributes {
    }
    
    export interface ListenerRuleAuthenticateCognitoConfigAttributes {
        AuthenticationRequestExtraParams?: {[key: string]: string};
        OnUnauthenticatedRequest?: string;
        Scope?: string;
        SessionCookieName?: string;
        SessionTimeout?: number;
        UserPoolArn: string;
        UserPoolClientId: string;
        UserPoolDomain: string;
    }
    
    export interface ListenerRuleAuthenticateCognitoConfigProperties {
        AuthenticationRequestExtraParams?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        OnUnauthenticatedRequest?: pulumi.Input<string>;
        Scope?: pulumi.Input<string>;
        SessionCookieName?: pulumi.Input<string>;
        SessionTimeout?: pulumi.Input<number>;
        UserPoolArn: pulumi.Input<string>;
        UserPoolClientId: pulumi.Input<string>;
        UserPoolDomain: pulumi.Input<string>;
    }
    
    export interface ListenerRuleAuthenticateOidcConfigAttributes {
        AuthenticationRequestExtraParams?: {[key: string]: string};
        AuthorizationEndpoint: string;
        ClientId: string;
        ClientSecret: string;
        Issuer: string;
        OnUnauthenticatedRequest?: string;
        Scope?: string;
        SessionCookieName?: string;
        SessionTimeout?: number;
        TokenEndpoint: string;
        UserInfoEndpoint: string;
    }
    
    export interface ListenerRuleAuthenticateOidcConfigProperties {
        AuthenticationRequestExtraParams?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        AuthorizationEndpoint: pulumi.Input<string>;
        ClientId: pulumi.Input<string>;
        ClientSecret: pulumi.Input<string>;
        Issuer: pulumi.Input<string>;
        OnUnauthenticatedRequest?: pulumi.Input<string>;
        Scope?: pulumi.Input<string>;
        SessionCookieName?: pulumi.Input<string>;
        SessionTimeout?: pulumi.Input<number>;
        TokenEndpoint: pulumi.Input<string>;
        UserInfoEndpoint: pulumi.Input<string>;
    }
    
    export interface ListenerRuleFixedResponseConfigAttributes {
        ContentType?: string;
        MessageBody?: string;
        StatusCode: string;
    }
    
    export interface ListenerRuleFixedResponseConfigProperties {
        ContentType?: pulumi.Input<string>;
        MessageBody?: pulumi.Input<string>;
        StatusCode: pulumi.Input<string>;
    }
    
    export interface ListenerRuleHostHeaderConfigAttributes {
        Values?: string[];
    }
    
    export interface ListenerRuleHostHeaderConfigProperties {
        Values?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface ListenerRuleHttpHeaderConfigAttributes {
        HttpHeaderName?: string;
        Values?: string[];
    }
    
    export interface ListenerRuleHttpHeaderConfigProperties {
        HttpHeaderName?: pulumi.Input<string>;
        Values?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface ListenerRuleHttpRequestMethodConfigAttributes {
        Values?: string[];
    }
    
    export interface ListenerRuleHttpRequestMethodConfigProperties {
        Values?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface ListenerRulePathPatternConfigAttributes {
        Values?: string[];
    }
    
    export interface ListenerRulePathPatternConfigProperties {
        Values?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface ListenerRuleProperties {
        Actions: pulumi.Input<pulumi.Input<ListenerRuleActionProperties>[]>;
        Conditions: pulumi.Input<pulumi.Input<ListenerRuleRuleConditionProperties>[]>;
        ListenerArn: pulumi.Input<string>;
        Priority: pulumi.Input<number>;
    }
    
    export interface ListenerRuleQueryStringConfigAttributes {
        Values?: ListenerRuleQueryStringKeyValueAttributes[];
    }
    
    export interface ListenerRuleQueryStringConfigProperties {
        Values?: pulumi.Input<pulumi.Input<ListenerRuleQueryStringKeyValueProperties>[]>;
    }
    
    export interface ListenerRuleQueryStringKeyValueAttributes {
        Key?: string;
        Value?: string;
    }
    
    export interface ListenerRuleQueryStringKeyValueProperties {
        Key?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface ListenerRuleRedirectConfigAttributes {
        Host?: string;
        Path?: string;
        Port?: string;
        Protocol?: string;
        Query?: string;
        StatusCode: string;
    }
    
    export interface ListenerRuleRedirectConfigProperties {
        Host?: pulumi.Input<string>;
        Path?: pulumi.Input<string>;
        Port?: pulumi.Input<string>;
        Protocol?: pulumi.Input<string>;
        Query?: pulumi.Input<string>;
        StatusCode: pulumi.Input<string>;
    }
    
    export interface ListenerRuleRuleConditionAttributes {
        Field?: string;
        HostHeaderConfig?: ListenerRuleHostHeaderConfigAttributes;
        HttpHeaderConfig?: ListenerRuleHttpHeaderConfigAttributes;
        HttpRequestMethodConfig?: ListenerRuleHttpRequestMethodConfigAttributes;
        PathPatternConfig?: ListenerRulePathPatternConfigAttributes;
        QueryStringConfig?: ListenerRuleQueryStringConfigAttributes;
        SourceIpConfig?: ListenerRuleSourceIpConfigAttributes;
        Values?: string[];
    }
    
    export interface ListenerRuleRuleConditionProperties {
        Field?: pulumi.Input<string>;
        HostHeaderConfig?: pulumi.Input<ListenerRuleHostHeaderConfigProperties>;
        HttpHeaderConfig?: pulumi.Input<ListenerRuleHttpHeaderConfigProperties>;
        HttpRequestMethodConfig?: pulumi.Input<ListenerRuleHttpRequestMethodConfigProperties>;
        PathPatternConfig?: pulumi.Input<ListenerRulePathPatternConfigProperties>;
        QueryStringConfig?: pulumi.Input<ListenerRuleQueryStringConfigProperties>;
        SourceIpConfig?: pulumi.Input<ListenerRuleSourceIpConfigProperties>;
        Values?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface ListenerRuleSourceIpConfigAttributes {
        Values?: string[];
    }
    
    export interface ListenerRuleSourceIpConfigProperties {
        Values?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface LoadBalancerAttributes {
        CanonicalHostedZoneID: string;
        DNSName: string;
        LoadBalancerFullName: string;
        LoadBalancerName: string;
        SecurityGroups: string[];
    }
    
    export interface LoadBalancerLoadBalancerAttributeAttributes {
        Key?: string;
        Value?: string;
    }
    
    export interface LoadBalancerLoadBalancerAttributeProperties {
        Key?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface LoadBalancerProperties {
        IpAddressType?: pulumi.Input<string>;
        LoadBalancerAttributes?: pulumi.Input<pulumi.Input<LoadBalancerLoadBalancerAttributeProperties>[]>;
        Name?: pulumi.Input<string>;
        Scheme?: pulumi.Input<string>;
        SecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
        SubnetMappings?: pulumi.Input<pulumi.Input<LoadBalancerSubnetMappingProperties>[]>;
        Subnets?: pulumi.Input<pulumi.Input<string>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        Type?: pulumi.Input<string>;
    }
    
    export interface LoadBalancerSubnetMappingAttributes {
        AllocationId: string;
        SubnetId: string;
    }
    
    export interface LoadBalancerSubnetMappingProperties {
        AllocationId: pulumi.Input<string>;
        SubnetId: pulumi.Input<string>;
    }
    
    export interface TargetGroupAttributes {
        LoadBalancerArns: string[];
        TargetGroupFullName: string;
        TargetGroupName: string;
    }
    
    export interface TargetGroupMatcherAttributes {
        HttpCode: string;
    }
    
    export interface TargetGroupMatcherProperties {
        HttpCode: pulumi.Input<string>;
    }
    
    export interface TargetGroupProperties {
        HealthCheckEnabled?: pulumi.Input<boolean>;
        HealthCheckIntervalSeconds?: pulumi.Input<number>;
        HealthCheckPath?: pulumi.Input<string>;
        HealthCheckPort?: pulumi.Input<string>;
        HealthCheckProtocol?: pulumi.Input<string>;
        HealthCheckTimeoutSeconds?: pulumi.Input<number>;
        HealthyThresholdCount?: pulumi.Input<number>;
        Matcher?: pulumi.Input<TargetGroupMatcherProperties>;
        Name?: pulumi.Input<string>;
        Port?: pulumi.Input<number>;
        Protocol?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        TargetGroupAttributes?: pulumi.Input<pulumi.Input<TargetGroupTargetGroupAttributeProperties>[]>;
        TargetType?: pulumi.Input<string>;
        Targets?: pulumi.Input<pulumi.Input<TargetGroupTargetDescriptionProperties>[]>;
        UnhealthyThresholdCount?: pulumi.Input<number>;
        VpcId?: pulumi.Input<string>;
    }
    
    export interface TargetGroupTargetDescriptionAttributes {
        AvailabilityZone?: string;
        Id: string;
        Port?: number;
    }
    
    export interface TargetGroupTargetDescriptionProperties {
        AvailabilityZone?: pulumi.Input<string>;
        Id: pulumi.Input<string>;
        Port?: pulumi.Input<number>;
    }
    
    export interface TargetGroupTargetGroupAttributeAttributes {
        Key?: string;
        Value?: string;
    }
    
    export interface TargetGroupTargetGroupAttributeProperties {
        Key?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    
    export class Listener extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ListenerAttributes>;

        constructor(name: string, properties: ListenerProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ElasticLoadBalancingV2:Listener", name, inputs, opts);
        }
    }
    
    export class ListenerCertificate extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ListenerCertificateAttributes>;

        constructor(name: string, properties: ListenerCertificateProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ElasticLoadBalancingV2:ListenerCertificate", name, inputs, opts);
        }
    }
    
    export class ListenerRule extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ListenerRuleAttributes>;

        constructor(name: string, properties: ListenerRuleProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ElasticLoadBalancingV2:ListenerRule", name, inputs, opts);
        }
    }
    
    export class LoadBalancer extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<LoadBalancerAttributes>;

        constructor(name: string, properties?: LoadBalancerProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ElasticLoadBalancingV2:LoadBalancer", name, inputs, opts);
        }
    }
    
    export class TargetGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TargetGroupAttributes>;

        constructor(name: string, properties?: TargetGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ElasticLoadBalancingV2:TargetGroup", name, inputs, opts);
        }
    }
    
}

export namespace elasticsearch {
    
    export interface DomainAttributes {
        Arn: string;
        DomainArn: string;
        DomainEndpoint: string;
    }
    
    export interface DomainEBSOptionsAttributes {
        EBSEnabled?: boolean;
        Iops?: number;
        VolumeSize?: number;
        VolumeType?: string;
    }
    
    export interface DomainEBSOptionsProperties {
        EBSEnabled?: pulumi.Input<boolean>;
        Iops?: pulumi.Input<number>;
        VolumeSize?: pulumi.Input<number>;
        VolumeType?: pulumi.Input<string>;
    }
    
    export interface DomainElasticsearchClusterConfigAttributes {
        DedicatedMasterCount?: number;
        DedicatedMasterEnabled?: boolean;
        DedicatedMasterType?: string;
        InstanceCount?: number;
        InstanceType?: string;
        ZoneAwarenessConfig?: DomainZoneAwarenessConfigAttributes;
        ZoneAwarenessEnabled?: boolean;
    }
    
    export interface DomainElasticsearchClusterConfigProperties {
        DedicatedMasterCount?: pulumi.Input<number>;
        DedicatedMasterEnabled?: pulumi.Input<boolean>;
        DedicatedMasterType?: pulumi.Input<string>;
        InstanceCount?: pulumi.Input<number>;
        InstanceType?: pulumi.Input<string>;
        ZoneAwarenessConfig?: pulumi.Input<DomainZoneAwarenessConfigProperties>;
        ZoneAwarenessEnabled?: pulumi.Input<boolean>;
    }
    
    export interface DomainEncryptionAtRestOptionsAttributes {
        Enabled?: boolean;
        KmsKeyId?: string;
    }
    
    export interface DomainEncryptionAtRestOptionsProperties {
        Enabled?: pulumi.Input<boolean>;
        KmsKeyId?: pulumi.Input<string>;
    }
    
    export interface DomainLogPublishingOptionAttributes {
        CloudWatchLogsLogGroupArn?: string;
        Enabled?: boolean;
    }
    
    export interface DomainLogPublishingOptionProperties {
        CloudWatchLogsLogGroupArn?: pulumi.Input<string>;
        Enabled?: pulumi.Input<boolean>;
    }
    
    export interface DomainNodeToNodeEncryptionOptionsAttributes {
        Enabled?: boolean;
    }
    
    export interface DomainNodeToNodeEncryptionOptionsProperties {
        Enabled?: pulumi.Input<boolean>;
    }
    
    export interface DomainProperties {
        AccessPolicies?: pulumi.Input<string>;
        AdvancedOptions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        DomainName?: pulumi.Input<string>;
        EBSOptions?: pulumi.Input<DomainEBSOptionsProperties>;
        ElasticsearchClusterConfig?: pulumi.Input<DomainElasticsearchClusterConfigProperties>;
        ElasticsearchVersion?: pulumi.Input<string>;
        EncryptionAtRestOptions?: pulumi.Input<DomainEncryptionAtRestOptionsProperties>;
        LogPublishingOptions?: pulumi.Input<{[key: string]: pulumi.Input<DomainLogPublishingOptionProperties>}>;
        NodeToNodeEncryptionOptions?: pulumi.Input<DomainNodeToNodeEncryptionOptionsProperties>;
        SnapshotOptions?: pulumi.Input<DomainSnapshotOptionsProperties>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VPCOptions?: pulumi.Input<DomainVPCOptionsProperties>;
    }
    
    export interface DomainSnapshotOptionsAttributes {
        AutomatedSnapshotStartHour?: number;
    }
    
    export interface DomainSnapshotOptionsProperties {
        AutomatedSnapshotStartHour?: pulumi.Input<number>;
    }
    
    export interface DomainVPCOptionsAttributes {
        SecurityGroupIds?: string[];
        SubnetIds?: string[];
    }
    
    export interface DomainVPCOptionsProperties {
        SecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
        SubnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface DomainZoneAwarenessConfigAttributes {
        AvailabilityZoneCount?: number;
    }
    
    export interface DomainZoneAwarenessConfigProperties {
        AvailabilityZoneCount?: pulumi.Input<number>;
    }
    
    
    export class Domain extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DomainAttributes>;

        constructor(name: string, properties?: DomainProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Elasticsearch:Domain", name, inputs, opts);
        }
    }
    
}

export namespace events {
    
    export interface EventBusAttributes {
        Arn: string;
        Name: string;
        Policy: string;
    }
    
    export interface EventBusPolicyAttributes {
    }
    
    export interface EventBusPolicyConditionAttributes {
        Key?: string;
        Type?: string;
        Value?: string;
    }
    
    export interface EventBusPolicyConditionProperties {
        Key?: pulumi.Input<string>;
        Type?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface EventBusPolicyProperties {
        Action: pulumi.Input<string>;
        Condition?: pulumi.Input<EventBusPolicyConditionProperties>;
        EventBusName?: pulumi.Input<string>;
        Principal: pulumi.Input<string>;
        StatementId: pulumi.Input<string>;
    }
    
    export interface EventBusProperties {
        EventSourceName?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
    }
    
    export interface RuleAttributes {
        Arn: string;
    }
    
    export interface RuleAwsVpcConfigurationAttributes {
        AssignPublicIp?: string;
        SecurityGroups?: string[];
        Subnets: string[];
    }
    
    export interface RuleAwsVpcConfigurationProperties {
        AssignPublicIp?: pulumi.Input<string>;
        SecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
        Subnets: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface RuleBatchArrayPropertiesAttributes {
        Size?: number;
    }
    
    export interface RuleBatchArrayPropertiesProperties {
        Size?: pulumi.Input<number>;
    }
    
    export interface RuleBatchParametersAttributes {
        ArrayProperties?: RuleBatchArrayPropertiesAttributes;
        JobDefinition: string;
        JobName: string;
        RetryStrategy?: RuleBatchRetryStrategyAttributes;
    }
    
    export interface RuleBatchParametersProperties {
        ArrayProperties?: pulumi.Input<RuleBatchArrayPropertiesProperties>;
        JobDefinition: pulumi.Input<string>;
        JobName: pulumi.Input<string>;
        RetryStrategy?: pulumi.Input<RuleBatchRetryStrategyProperties>;
    }
    
    export interface RuleBatchRetryStrategyAttributes {
        Attempts?: number;
    }
    
    export interface RuleBatchRetryStrategyProperties {
        Attempts?: pulumi.Input<number>;
    }
    
    export interface RuleEcsParametersAttributes {
        Group?: string;
        LaunchType?: string;
        NetworkConfiguration?: RuleNetworkConfigurationAttributes;
        PlatformVersion?: string;
        TaskCount?: number;
        TaskDefinitionArn: string;
    }
    
    export interface RuleEcsParametersProperties {
        Group?: pulumi.Input<string>;
        LaunchType?: pulumi.Input<string>;
        NetworkConfiguration?: pulumi.Input<RuleNetworkConfigurationProperties>;
        PlatformVersion?: pulumi.Input<string>;
        TaskCount?: pulumi.Input<number>;
        TaskDefinitionArn: pulumi.Input<string>;
    }
    
    export interface RuleInputTransformerAttributes {
        InputPathsMap?: {[key: string]: string};
        InputTemplate: string;
    }
    
    export interface RuleInputTransformerProperties {
        InputPathsMap?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        InputTemplate: pulumi.Input<string>;
    }
    
    export interface RuleKinesisParametersAttributes {
        PartitionKeyPath: string;
    }
    
    export interface RuleKinesisParametersProperties {
        PartitionKeyPath: pulumi.Input<string>;
    }
    
    export interface RuleNetworkConfigurationAttributes {
        AwsVpcConfiguration?: RuleAwsVpcConfigurationAttributes;
    }
    
    export interface RuleNetworkConfigurationProperties {
        AwsVpcConfiguration?: pulumi.Input<RuleAwsVpcConfigurationProperties>;
    }
    
    export interface RuleProperties {
        Description?: pulumi.Input<string>;
        EventBusName?: pulumi.Input<string>;
        EventPattern?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        RoleArn?: pulumi.Input<string>;
        ScheduleExpression?: pulumi.Input<string>;
        State?: pulumi.Input<string>;
        Targets?: pulumi.Input<pulumi.Input<RuleTargetProperties>[]>;
    }
    
    export interface RuleRunCommandParametersAttributes {
        RunCommandTargets: RuleRunCommandTargetAttributes[];
    }
    
    export interface RuleRunCommandParametersProperties {
        RunCommandTargets: pulumi.Input<pulumi.Input<RuleRunCommandTargetProperties>[]>;
    }
    
    export interface RuleRunCommandTargetAttributes {
        Key: string;
        Values: string[];
    }
    
    export interface RuleRunCommandTargetProperties {
        Key: pulumi.Input<string>;
        Values: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface RuleSqsParametersAttributes {
        MessageGroupId: string;
    }
    
    export interface RuleSqsParametersProperties {
        MessageGroupId: pulumi.Input<string>;
    }
    
    export interface RuleTargetAttributes {
        Arn: string;
        BatchParameters?: RuleBatchParametersAttributes;
        EcsParameters?: RuleEcsParametersAttributes;
        Id: string;
        Input?: string;
        InputPath?: string;
        InputTransformer?: RuleInputTransformerAttributes;
        KinesisParameters?: RuleKinesisParametersAttributes;
        RoleArn?: string;
        RunCommandParameters?: RuleRunCommandParametersAttributes;
        SqsParameters?: RuleSqsParametersAttributes;
    }
    
    export interface RuleTargetProperties {
        Arn: pulumi.Input<string>;
        BatchParameters?: pulumi.Input<RuleBatchParametersProperties>;
        EcsParameters?: pulumi.Input<RuleEcsParametersProperties>;
        Id: pulumi.Input<string>;
        Input?: pulumi.Input<string>;
        InputPath?: pulumi.Input<string>;
        InputTransformer?: pulumi.Input<RuleInputTransformerProperties>;
        KinesisParameters?: pulumi.Input<RuleKinesisParametersProperties>;
        RoleArn?: pulumi.Input<string>;
        RunCommandParameters?: pulumi.Input<RuleRunCommandParametersProperties>;
        SqsParameters?: pulumi.Input<RuleSqsParametersProperties>;
    }
    
    
    export class EventBus extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<EventBusAttributes>;

        constructor(name: string, properties: EventBusProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Events:EventBus", name, inputs, opts);
        }
    }
    
    export class EventBusPolicy extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<EventBusPolicyAttributes>;

        constructor(name: string, properties: EventBusPolicyProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Events:EventBusPolicy", name, inputs, opts);
        }
    }
    
    export class Rule extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RuleAttributes>;

        constructor(name: string, properties?: RuleProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Events:Rule", name, inputs, opts);
        }
    }
    
}

export namespace fsx {
    
    export interface FileSystemAttributes {
    }
    
    export interface FileSystemLustreConfigurationAttributes {
        ExportPath?: string;
        ImportPath?: string;
        ImportedFileChunkSize?: number;
        WeeklyMaintenanceStartTime?: string;
    }
    
    export interface FileSystemLustreConfigurationProperties {
        ExportPath?: pulumi.Input<string>;
        ImportPath?: pulumi.Input<string>;
        ImportedFileChunkSize?: pulumi.Input<number>;
        WeeklyMaintenanceStartTime?: pulumi.Input<string>;
    }
    
    export interface FileSystemProperties {
        BackupId?: pulumi.Input<string>;
        FileSystemType: pulumi.Input<string>;
        KmsKeyId?: pulumi.Input<string>;
        LustreConfiguration?: pulumi.Input<FileSystemLustreConfigurationProperties>;
        SecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
        StorageCapacity?: pulumi.Input<number>;
        SubnetIds: pulumi.Input<pulumi.Input<string>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        WindowsConfiguration?: pulumi.Input<FileSystemWindowsConfigurationProperties>;
    }
    
    export interface FileSystemSelfManagedActiveDirectoryConfigurationAttributes {
        DnsIps?: string[];
        DomainName?: string;
        FileSystemAdministratorsGroup?: string;
        OrganizationalUnitDistinguishedName?: string;
        Password?: string;
        UserName?: string;
    }
    
    export interface FileSystemSelfManagedActiveDirectoryConfigurationProperties {
        DnsIps?: pulumi.Input<pulumi.Input<string>[]>;
        DomainName?: pulumi.Input<string>;
        FileSystemAdministratorsGroup?: pulumi.Input<string>;
        OrganizationalUnitDistinguishedName?: pulumi.Input<string>;
        Password?: pulumi.Input<string>;
        UserName?: pulumi.Input<string>;
    }
    
    export interface FileSystemWindowsConfigurationAttributes {
        ActiveDirectoryId?: string;
        AutomaticBackupRetentionDays?: number;
        CopyTagsToBackups?: boolean;
        DailyAutomaticBackupStartTime?: string;
        SelfManagedActiveDirectoryConfiguration?: FileSystemSelfManagedActiveDirectoryConfigurationAttributes;
        ThroughputCapacity?: number;
        WeeklyMaintenanceStartTime?: string;
    }
    
    export interface FileSystemWindowsConfigurationProperties {
        ActiveDirectoryId?: pulumi.Input<string>;
        AutomaticBackupRetentionDays?: pulumi.Input<number>;
        CopyTagsToBackups?: pulumi.Input<boolean>;
        DailyAutomaticBackupStartTime?: pulumi.Input<string>;
        SelfManagedActiveDirectoryConfiguration?: pulumi.Input<FileSystemSelfManagedActiveDirectoryConfigurationProperties>;
        ThroughputCapacity?: pulumi.Input<number>;
        WeeklyMaintenanceStartTime?: pulumi.Input<string>;
    }
    
    
    export class FileSystem extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<FileSystemAttributes>;

        constructor(name: string, properties: FileSystemProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:FSx:FileSystem", name, inputs, opts);
        }
    }
    
}

export namespace gamelift {
    
    export interface AliasAttributes {
    }
    
    export interface AliasProperties {
        Description?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        RoutingStrategy: pulumi.Input<AliasRoutingStrategyProperties>;
    }
    
    export interface AliasRoutingStrategyAttributes {
        FleetId?: string;
        Message?: string;
        Type: string;
    }
    
    export interface AliasRoutingStrategyProperties {
        FleetId?: pulumi.Input<string>;
        Message?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface BuildAttributes {
    }
    
    export interface BuildProperties {
        Name?: pulumi.Input<string>;
        StorageLocation?: pulumi.Input<BuildS3LocationProperties>;
        Version?: pulumi.Input<string>;
    }
    
    export interface BuildS3LocationAttributes {
        Bucket: string;
        Key: string;
        RoleArn: string;
    }
    
    export interface BuildS3LocationProperties {
        Bucket: pulumi.Input<string>;
        Key: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
    }
    
    export interface FleetAttributes {
    }
    
    export interface FleetIpPermissionAttributes {
        FromPort: number;
        IpRange: string;
        Protocol: string;
        ToPort: number;
    }
    
    export interface FleetIpPermissionProperties {
        FromPort: pulumi.Input<number>;
        IpRange: pulumi.Input<string>;
        Protocol: pulumi.Input<string>;
        ToPort: pulumi.Input<number>;
    }
    
    export interface FleetProperties {
        BuildId: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        DesiredEC2Instances: pulumi.Input<number>;
        EC2InboundPermissions?: pulumi.Input<pulumi.Input<FleetIpPermissionProperties>[]>;
        EC2InstanceType: pulumi.Input<string>;
        LogPaths?: pulumi.Input<pulumi.Input<string>[]>;
        MaxSize?: pulumi.Input<number>;
        MinSize?: pulumi.Input<number>;
        Name: pulumi.Input<string>;
        ServerLaunchParameters?: pulumi.Input<string>;
        ServerLaunchPath: pulumi.Input<string>;
    }
    
    
    export class Alias extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AliasAttributes>;

        constructor(name: string, properties: AliasProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:GameLift:Alias", name, inputs, opts);
        }
    }
    
    export class Build extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<BuildAttributes>;

        constructor(name: string, properties?: BuildProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:GameLift:Build", name, inputs, opts);
        }
    }
    
    export class Fleet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<FleetAttributes>;

        constructor(name: string, properties: FleetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:GameLift:Fleet", name, inputs, opts);
        }
    }
    
}

export namespace glue {
    
    export interface ClassifierAttributes {
    }
    
    export interface ClassifierCsvClassifierAttributes {
        AllowSingleColumn?: boolean;
        ContainsHeader?: string;
        Delimiter?: string;
        DisableValueTrimming?: boolean;
        Header?: string[];
        Name?: string;
        QuoteSymbol?: string;
    }
    
    export interface ClassifierCsvClassifierProperties {
        AllowSingleColumn?: pulumi.Input<boolean>;
        ContainsHeader?: pulumi.Input<string>;
        Delimiter?: pulumi.Input<string>;
        DisableValueTrimming?: pulumi.Input<boolean>;
        Header?: pulumi.Input<pulumi.Input<string>[]>;
        Name?: pulumi.Input<string>;
        QuoteSymbol?: pulumi.Input<string>;
    }
    
    export interface ClassifierGrokClassifierAttributes {
        Classification: string;
        CustomPatterns?: string;
        GrokPattern: string;
        Name?: string;
    }
    
    export interface ClassifierGrokClassifierProperties {
        Classification: pulumi.Input<string>;
        CustomPatterns?: pulumi.Input<string>;
        GrokPattern: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
    }
    
    export interface ClassifierJsonClassifierAttributes {
        JsonPath: string;
        Name?: string;
    }
    
    export interface ClassifierJsonClassifierProperties {
        JsonPath: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
    }
    
    export interface ClassifierProperties {
        CsvClassifier?: pulumi.Input<ClassifierCsvClassifierProperties>;
        GrokClassifier?: pulumi.Input<ClassifierGrokClassifierProperties>;
        JsonClassifier?: pulumi.Input<ClassifierJsonClassifierProperties>;
        XMLClassifier?: pulumi.Input<ClassifierXMLClassifierProperties>;
    }
    
    export interface ClassifierXMLClassifierAttributes {
        Classification: string;
        Name?: string;
        RowTag: string;
    }
    
    export interface ClassifierXMLClassifierProperties {
        Classification: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        RowTag: pulumi.Input<string>;
    }
    
    export interface ConnectionAttributes {
    }
    
    export interface ConnectionConnectionInputAttributes {
        ConnectionProperties: string;
        ConnectionType: string;
        Description?: string;
        MatchCriteria?: string[];
        Name?: string;
        PhysicalConnectionRequirements?: ConnectionPhysicalConnectionRequirementsAttributes;
    }
    
    export interface ConnectionConnectionInputProperties {
        ConnectionProperties: pulumi.Input<string>;
        ConnectionType: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        MatchCriteria?: pulumi.Input<pulumi.Input<string>[]>;
        Name?: pulumi.Input<string>;
        PhysicalConnectionRequirements?: pulumi.Input<ConnectionPhysicalConnectionRequirementsProperties>;
    }
    
    export interface ConnectionPhysicalConnectionRequirementsAttributes {
        AvailabilityZone?: string;
        SecurityGroupIdList?: string[];
        SubnetId?: string;
    }
    
    export interface ConnectionPhysicalConnectionRequirementsProperties {
        AvailabilityZone?: pulumi.Input<string>;
        SecurityGroupIdList?: pulumi.Input<pulumi.Input<string>[]>;
        SubnetId?: pulumi.Input<string>;
    }
    
    export interface ConnectionProperties {
        CatalogId: pulumi.Input<string>;
        ConnectionInput: pulumi.Input<ConnectionConnectionInputProperties>;
    }
    
    export interface CrawlerAttributes {
    }
    
    export interface CrawlerCatalogTargetAttributes {
        DatabaseName?: string;
        Tables?: string[];
    }
    
    export interface CrawlerCatalogTargetProperties {
        DatabaseName?: pulumi.Input<string>;
        Tables?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface CrawlerDynamoDBTargetAttributes {
        Path?: string;
    }
    
    export interface CrawlerDynamoDBTargetProperties {
        Path?: pulumi.Input<string>;
    }
    
    export interface CrawlerJdbcTargetAttributes {
        ConnectionName?: string;
        Exclusions?: string[];
        Path?: string;
    }
    
    export interface CrawlerJdbcTargetProperties {
        ConnectionName?: pulumi.Input<string>;
        Exclusions?: pulumi.Input<pulumi.Input<string>[]>;
        Path?: pulumi.Input<string>;
    }
    
    export interface CrawlerProperties {
        Classifiers?: pulumi.Input<pulumi.Input<string>[]>;
        Configuration?: pulumi.Input<string>;
        CrawlerSecurityConfiguration?: pulumi.Input<string>;
        DatabaseName?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Role: pulumi.Input<string>;
        Schedule?: pulumi.Input<CrawlerScheduleProperties>;
        SchemaChangePolicy?: pulumi.Input<CrawlerSchemaChangePolicyProperties>;
        TablePrefix?: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
        Targets: pulumi.Input<CrawlerTargetsProperties>;
    }
    
    export interface CrawlerS3TargetAttributes {
        Exclusions?: string[];
        Path?: string;
    }
    
    export interface CrawlerS3TargetProperties {
        Exclusions?: pulumi.Input<pulumi.Input<string>[]>;
        Path?: pulumi.Input<string>;
    }
    
    export interface CrawlerScheduleAttributes {
        ScheduleExpression?: string;
    }
    
    export interface CrawlerScheduleProperties {
        ScheduleExpression?: pulumi.Input<string>;
    }
    
    export interface CrawlerSchemaChangePolicyAttributes {
        DeleteBehavior?: string;
        UpdateBehavior?: string;
    }
    
    export interface CrawlerSchemaChangePolicyProperties {
        DeleteBehavior?: pulumi.Input<string>;
        UpdateBehavior?: pulumi.Input<string>;
    }
    
    export interface CrawlerTargetsAttributes {
        CatalogTargets?: CrawlerCatalogTargetAttributes[];
        DynamoDBTargets?: CrawlerDynamoDBTargetAttributes[];
        JdbcTargets?: CrawlerJdbcTargetAttributes[];
        S3Targets?: CrawlerS3TargetAttributes[];
    }
    
    export interface CrawlerTargetsProperties {
        CatalogTargets?: pulumi.Input<pulumi.Input<CrawlerCatalogTargetProperties>[]>;
        DynamoDBTargets?: pulumi.Input<pulumi.Input<CrawlerDynamoDBTargetProperties>[]>;
        JdbcTargets?: pulumi.Input<pulumi.Input<CrawlerJdbcTargetProperties>[]>;
        S3Targets?: pulumi.Input<pulumi.Input<CrawlerS3TargetProperties>[]>;
    }
    
    export interface DataCatalogEncryptionSettingsAttributes {
    }
    
    export interface DataCatalogEncryptionSettingsConnectionPasswordEncryptionAttributes {
        KmsKeyId?: string;
        ReturnConnectionPasswordEncrypted?: boolean;
    }
    
    export interface DataCatalogEncryptionSettingsConnectionPasswordEncryptionProperties {
        KmsKeyId?: pulumi.Input<string>;
        ReturnConnectionPasswordEncrypted?: pulumi.Input<boolean>;
    }
    
    export interface DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsAttributes {
        ConnectionPasswordEncryption?: DataCatalogEncryptionSettingsConnectionPasswordEncryptionAttributes;
        EncryptionAtRest?: DataCatalogEncryptionSettingsEncryptionAtRestAttributes;
    }
    
    export interface DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsProperties {
        ConnectionPasswordEncryption?: pulumi.Input<DataCatalogEncryptionSettingsConnectionPasswordEncryptionProperties>;
        EncryptionAtRest?: pulumi.Input<DataCatalogEncryptionSettingsEncryptionAtRestProperties>;
    }
    
    export interface DataCatalogEncryptionSettingsEncryptionAtRestAttributes {
        CatalogEncryptionMode?: string;
        SseAwsKmsKeyId?: string;
    }
    
    export interface DataCatalogEncryptionSettingsEncryptionAtRestProperties {
        CatalogEncryptionMode?: pulumi.Input<string>;
        SseAwsKmsKeyId?: pulumi.Input<string>;
    }
    
    export interface DataCatalogEncryptionSettingsProperties {
        CatalogId: pulumi.Input<string>;
        DataCatalogEncryptionSettings: pulumi.Input<DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsProperties>;
    }
    
    export interface DatabaseAttributes {
    }
    
    export interface DatabaseDatabaseInputAttributes {
        Description?: string;
        LocationUri?: string;
        Name?: string;
        Parameters?: string;
    }
    
    export interface DatabaseDatabaseInputProperties {
        Description?: pulumi.Input<string>;
        LocationUri?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Parameters?: pulumi.Input<string>;
    }
    
    export interface DatabaseProperties {
        CatalogId: pulumi.Input<string>;
        DatabaseInput: pulumi.Input<DatabaseDatabaseInputProperties>;
    }
    
    export interface DevEndpointAttributes {
    }
    
    export interface DevEndpointProperties {
        Arguments?: pulumi.Input<string>;
        EndpointName?: pulumi.Input<string>;
        ExtraJarsS3Path?: pulumi.Input<string>;
        ExtraPythonLibsS3Path?: pulumi.Input<string>;
        GlueVersion?: pulumi.Input<string>;
        NumberOfNodes?: pulumi.Input<number>;
        NumberOfWorkers?: pulumi.Input<number>;
        PublicKey?: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
        SecurityConfiguration?: pulumi.Input<string>;
        SecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
        SubnetId?: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
        WorkerType?: pulumi.Input<string>;
    }
    
    export interface JobAttributes {
    }
    
    export interface JobConnectionsListAttributes {
        Connections?: string[];
    }
    
    export interface JobConnectionsListProperties {
        Connections?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface JobExecutionPropertyAttributes {
        MaxConcurrentRuns?: number;
    }
    
    export interface JobExecutionPropertyProperties {
        MaxConcurrentRuns?: pulumi.Input<number>;
    }
    
    export interface JobJobCommandAttributes {
        Name?: string;
        PythonVersion?: string;
        ScriptLocation?: string;
    }
    
    export interface JobJobCommandProperties {
        Name?: pulumi.Input<string>;
        PythonVersion?: pulumi.Input<string>;
        ScriptLocation?: pulumi.Input<string>;
    }
    
    export interface JobNotificationPropertyAttributes {
        NotifyDelayAfter?: number;
    }
    
    export interface JobNotificationPropertyProperties {
        NotifyDelayAfter?: pulumi.Input<number>;
    }
    
    export interface JobProperties {
        AllocatedCapacity?: pulumi.Input<number>;
        Command: pulumi.Input<JobJobCommandProperties>;
        Connections?: pulumi.Input<JobConnectionsListProperties>;
        DefaultArguments?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        ExecutionProperty?: pulumi.Input<JobExecutionPropertyProperties>;
        GlueVersion?: pulumi.Input<string>;
        LogUri?: pulumi.Input<string>;
        MaxCapacity?: pulumi.Input<number>;
        MaxRetries?: pulumi.Input<number>;
        Name?: pulumi.Input<string>;
        NotificationProperty?: pulumi.Input<JobNotificationPropertyProperties>;
        NumberOfWorkers?: pulumi.Input<number>;
        Role: pulumi.Input<string>;
        SecurityConfiguration?: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
        Timeout?: pulumi.Input<number>;
        WorkerType?: pulumi.Input<string>;
    }
    
    export interface MLTransformAttributes {
    }
    
    export interface MLTransformFindMatchesParametersAttributes {
        AccuracyCostTradeoff?: number;
        EnforceProvidedLabels?: boolean;
        PrecisionRecallTradeoff?: number;
        PrimaryKeyColumnName: string;
    }
    
    export interface MLTransformFindMatchesParametersProperties {
        AccuracyCostTradeoff?: pulumi.Input<number>;
        EnforceProvidedLabels?: pulumi.Input<boolean>;
        PrecisionRecallTradeoff?: pulumi.Input<number>;
        PrimaryKeyColumnName: pulumi.Input<string>;
    }
    
    export interface MLTransformGlueTablesAttributes {
        CatalogId?: string;
        ConnectionName?: string;
        DatabaseName: string;
        TableName: string;
    }
    
    export interface MLTransformGlueTablesProperties {
        CatalogId?: pulumi.Input<string>;
        ConnectionName?: pulumi.Input<string>;
        DatabaseName: pulumi.Input<string>;
        TableName: pulumi.Input<string>;
    }
    
    export interface MLTransformInputRecordTablesAttributes {
        GlueTables?: MLTransformGlueTablesAttributes[];
    }
    
    export interface MLTransformInputRecordTablesProperties {
        GlueTables?: pulumi.Input<pulumi.Input<MLTransformGlueTablesProperties>[]>;
    }
    
    export interface MLTransformProperties {
        Description?: pulumi.Input<string>;
        InputRecordTables: pulumi.Input<MLTransformInputRecordTablesProperties>;
        MaxCapacity?: pulumi.Input<number>;
        MaxRetries?: pulumi.Input<number>;
        Name?: pulumi.Input<string>;
        NumberOfWorkers?: pulumi.Input<number>;
        Role: pulumi.Input<string>;
        Timeout?: pulumi.Input<number>;
        TransformParameters: pulumi.Input<MLTransformTransformParametersProperties>;
        WorkerType?: pulumi.Input<string>;
    }
    
    export interface MLTransformTransformParametersAttributes {
        FindMatchesParameters?: MLTransformFindMatchesParametersAttributes;
        TransformType: string;
    }
    
    export interface MLTransformTransformParametersProperties {
        FindMatchesParameters?: pulumi.Input<MLTransformFindMatchesParametersProperties>;
        TransformType: pulumi.Input<string>;
    }
    
    export interface PartitionAttributes {
    }
    
    export interface PartitionColumnAttributes {
        Comment?: string;
        Name: string;
        Type?: string;
    }
    
    export interface PartitionColumnProperties {
        Comment?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        Type?: pulumi.Input<string>;
    }
    
    export interface PartitionOrderAttributes {
        Column: string;
        SortOrder?: number;
    }
    
    export interface PartitionOrderProperties {
        Column: pulumi.Input<string>;
        SortOrder?: pulumi.Input<number>;
    }
    
    export interface PartitionPartitionInputAttributes {
        Parameters?: string;
        StorageDescriptor?: PartitionStorageDescriptorAttributes;
        Values: string[];
    }
    
    export interface PartitionPartitionInputProperties {
        Parameters?: pulumi.Input<string>;
        StorageDescriptor?: pulumi.Input<PartitionStorageDescriptorProperties>;
        Values: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface PartitionProperties {
        CatalogId: pulumi.Input<string>;
        DatabaseName: pulumi.Input<string>;
        PartitionInput: pulumi.Input<PartitionPartitionInputProperties>;
        TableName: pulumi.Input<string>;
    }
    
    export interface PartitionSerdeInfoAttributes {
        Name?: string;
        Parameters?: string;
        SerializationLibrary?: string;
    }
    
    export interface PartitionSerdeInfoProperties {
        Name?: pulumi.Input<string>;
        Parameters?: pulumi.Input<string>;
        SerializationLibrary?: pulumi.Input<string>;
    }
    
    export interface PartitionSkewedInfoAttributes {
        SkewedColumnNames?: string[];
        SkewedColumnValueLocationMaps?: string;
        SkewedColumnValues?: string[];
    }
    
    export interface PartitionSkewedInfoProperties {
        SkewedColumnNames?: pulumi.Input<pulumi.Input<string>[]>;
        SkewedColumnValueLocationMaps?: pulumi.Input<string>;
        SkewedColumnValues?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface PartitionStorageDescriptorAttributes {
        BucketColumns?: string[];
        Columns?: PartitionColumnAttributes[];
        Compressed?: boolean;
        InputFormat?: string;
        Location?: string;
        NumberOfBuckets?: number;
        OutputFormat?: string;
        Parameters?: string;
        SerdeInfo?: PartitionSerdeInfoAttributes;
        SkewedInfo?: PartitionSkewedInfoAttributes;
        SortColumns?: PartitionOrderAttributes[];
        StoredAsSubDirectories?: boolean;
    }
    
    export interface PartitionStorageDescriptorProperties {
        BucketColumns?: pulumi.Input<pulumi.Input<string>[]>;
        Columns?: pulumi.Input<pulumi.Input<PartitionColumnProperties>[]>;
        Compressed?: pulumi.Input<boolean>;
        InputFormat?: pulumi.Input<string>;
        Location?: pulumi.Input<string>;
        NumberOfBuckets?: pulumi.Input<number>;
        OutputFormat?: pulumi.Input<string>;
        Parameters?: pulumi.Input<string>;
        SerdeInfo?: pulumi.Input<PartitionSerdeInfoProperties>;
        SkewedInfo?: pulumi.Input<PartitionSkewedInfoProperties>;
        SortColumns?: pulumi.Input<pulumi.Input<PartitionOrderProperties>[]>;
        StoredAsSubDirectories?: pulumi.Input<boolean>;
    }
    
    export interface SecurityConfigurationAttributes {
    }
    
    export interface SecurityConfigurationCloudWatchEncryptionAttributes {
        CloudWatchEncryptionMode?: string;
        KmsKeyArn?: string;
    }
    
    export interface SecurityConfigurationCloudWatchEncryptionProperties {
        CloudWatchEncryptionMode?: pulumi.Input<string>;
        KmsKeyArn?: pulumi.Input<string>;
    }
    
    export interface SecurityConfigurationEncryptionConfigurationAttributes {
        CloudWatchEncryption?: SecurityConfigurationCloudWatchEncryptionAttributes;
        JobBookmarksEncryption?: SecurityConfigurationJobBookmarksEncryptionAttributes;
        S3Encryptions?: SecurityConfigurationS3EncryptionsAttributes;
    }
    
    export interface SecurityConfigurationEncryptionConfigurationProperties {
        CloudWatchEncryption?: pulumi.Input<SecurityConfigurationCloudWatchEncryptionProperties>;
        JobBookmarksEncryption?: pulumi.Input<SecurityConfigurationJobBookmarksEncryptionProperties>;
        S3Encryptions?: pulumi.Input<SecurityConfigurationS3EncryptionsProperties>;
    }
    
    export interface SecurityConfigurationJobBookmarksEncryptionAttributes {
        JobBookmarksEncryptionMode?: string;
        KmsKeyArn?: string;
    }
    
    export interface SecurityConfigurationJobBookmarksEncryptionProperties {
        JobBookmarksEncryptionMode?: pulumi.Input<string>;
        KmsKeyArn?: pulumi.Input<string>;
    }
    
    export interface SecurityConfigurationProperties {
        EncryptionConfiguration: pulumi.Input<SecurityConfigurationEncryptionConfigurationProperties>;
        Name: pulumi.Input<string>;
    }
    
    export interface SecurityConfigurationS3EncryptionAttributes {
        KmsKeyArn?: string;
        S3EncryptionMode?: string;
    }
    
    export interface SecurityConfigurationS3EncryptionProperties {
        KmsKeyArn?: pulumi.Input<string>;
        S3EncryptionMode?: pulumi.Input<string>;
    }
    
    export interface SecurityConfigurationS3EncryptionsAttributes {
    }
    
    export interface SecurityConfigurationS3EncryptionsProperties {
    }
    
    export interface TableAttributes {
    }
    
    export interface TableColumnAttributes {
        Comment?: string;
        Name: string;
        Type?: string;
    }
    
    export interface TableColumnProperties {
        Comment?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        Type?: pulumi.Input<string>;
    }
    
    export interface TableOrderAttributes {
        Column: string;
        SortOrder: number;
    }
    
    export interface TableOrderProperties {
        Column: pulumi.Input<string>;
        SortOrder: pulumi.Input<number>;
    }
    
    export interface TableProperties {
        CatalogId: pulumi.Input<string>;
        DatabaseName: pulumi.Input<string>;
        TableInput: pulumi.Input<TableTableInputProperties>;
    }
    
    export interface TableSerdeInfoAttributes {
        Name?: string;
        Parameters?: string;
        SerializationLibrary?: string;
    }
    
    export interface TableSerdeInfoProperties {
        Name?: pulumi.Input<string>;
        Parameters?: pulumi.Input<string>;
        SerializationLibrary?: pulumi.Input<string>;
    }
    
    export interface TableSkewedInfoAttributes {
        SkewedColumnNames?: string[];
        SkewedColumnValueLocationMaps?: string;
        SkewedColumnValues?: string[];
    }
    
    export interface TableSkewedInfoProperties {
        SkewedColumnNames?: pulumi.Input<pulumi.Input<string>[]>;
        SkewedColumnValueLocationMaps?: pulumi.Input<string>;
        SkewedColumnValues?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface TableStorageDescriptorAttributes {
        BucketColumns?: string[];
        Columns?: TableColumnAttributes[];
        Compressed?: boolean;
        InputFormat?: string;
        Location?: string;
        NumberOfBuckets?: number;
        OutputFormat?: string;
        Parameters?: string;
        SerdeInfo?: TableSerdeInfoAttributes;
        SkewedInfo?: TableSkewedInfoAttributes;
        SortColumns?: TableOrderAttributes[];
        StoredAsSubDirectories?: boolean;
    }
    
    export interface TableStorageDescriptorProperties {
        BucketColumns?: pulumi.Input<pulumi.Input<string>[]>;
        Columns?: pulumi.Input<pulumi.Input<TableColumnProperties>[]>;
        Compressed?: pulumi.Input<boolean>;
        InputFormat?: pulumi.Input<string>;
        Location?: pulumi.Input<string>;
        NumberOfBuckets?: pulumi.Input<number>;
        OutputFormat?: pulumi.Input<string>;
        Parameters?: pulumi.Input<string>;
        SerdeInfo?: pulumi.Input<TableSerdeInfoProperties>;
        SkewedInfo?: pulumi.Input<TableSkewedInfoProperties>;
        SortColumns?: pulumi.Input<pulumi.Input<TableOrderProperties>[]>;
        StoredAsSubDirectories?: pulumi.Input<boolean>;
    }
    
    export interface TableTableInputAttributes {
        Description?: string;
        Name?: string;
        Owner?: string;
        Parameters?: string;
        PartitionKeys?: TableColumnAttributes[];
        Retention?: number;
        StorageDescriptor?: TableStorageDescriptorAttributes;
        TableType?: string;
        ViewExpandedText?: string;
        ViewOriginalText?: string;
    }
    
    export interface TableTableInputProperties {
        Description?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Owner?: pulumi.Input<string>;
        Parameters?: pulumi.Input<string>;
        PartitionKeys?: pulumi.Input<pulumi.Input<TableColumnProperties>[]>;
        Retention?: pulumi.Input<number>;
        StorageDescriptor?: pulumi.Input<TableStorageDescriptorProperties>;
        TableType?: pulumi.Input<string>;
        ViewExpandedText?: pulumi.Input<string>;
        ViewOriginalText?: pulumi.Input<string>;
    }
    
    export interface TriggerActionAttributes {
        Arguments?: string;
        CrawlerName?: string;
        JobName?: string;
        NotificationProperty?: TriggerNotificationPropertyAttributes;
        SecurityConfiguration?: string;
        Timeout?: number;
    }
    
    export interface TriggerActionProperties {
        Arguments?: pulumi.Input<string>;
        CrawlerName?: pulumi.Input<string>;
        JobName?: pulumi.Input<string>;
        NotificationProperty?: pulumi.Input<TriggerNotificationPropertyProperties>;
        SecurityConfiguration?: pulumi.Input<string>;
        Timeout?: pulumi.Input<number>;
    }
    
    export interface TriggerAttributes {
    }
    
    export interface TriggerConditionAttributes {
        CrawlState?: string;
        CrawlerName?: string;
        JobName?: string;
        LogicalOperator?: string;
        State?: string;
    }
    
    export interface TriggerConditionProperties {
        CrawlState?: pulumi.Input<string>;
        CrawlerName?: pulumi.Input<string>;
        JobName?: pulumi.Input<string>;
        LogicalOperator?: pulumi.Input<string>;
        State?: pulumi.Input<string>;
    }
    
    export interface TriggerNotificationPropertyAttributes {
        NotifyDelayAfter?: number;
    }
    
    export interface TriggerNotificationPropertyProperties {
        NotifyDelayAfter?: pulumi.Input<number>;
    }
    
    export interface TriggerPredicateAttributes {
        Conditions?: TriggerConditionAttributes[];
        Logical?: string;
    }
    
    export interface TriggerPredicateProperties {
        Conditions?: pulumi.Input<pulumi.Input<TriggerConditionProperties>[]>;
        Logical?: pulumi.Input<string>;
    }
    
    export interface TriggerProperties {
        Actions: pulumi.Input<pulumi.Input<TriggerActionProperties>[]>;
        Description?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Predicate?: pulumi.Input<TriggerPredicateProperties>;
        Schedule?: pulumi.Input<string>;
        StartOnCreation?: pulumi.Input<boolean>;
        Tags?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
        WorkflowName?: pulumi.Input<string>;
    }
    
    export interface WorkflowAttributes {
    }
    
    export interface WorkflowProperties {
        DefaultRunProperties?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
    }
    
    
    export class Classifier extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ClassifierAttributes>;

        constructor(name: string, properties?: ClassifierProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Glue:Classifier", name, inputs, opts);
        }
    }
    
    export class Connection extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ConnectionAttributes>;

        constructor(name: string, properties: ConnectionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Glue:Connection", name, inputs, opts);
        }
    }
    
    export class Crawler extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<CrawlerAttributes>;

        constructor(name: string, properties: CrawlerProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Glue:Crawler", name, inputs, opts);
        }
    }
    
    export class DataCatalogEncryptionSettings extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DataCatalogEncryptionSettingsAttributes>;

        constructor(name: string, properties: DataCatalogEncryptionSettingsProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Glue:DataCatalogEncryptionSettings", name, inputs, opts);
        }
    }
    
    export class Database extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DatabaseAttributes>;

        constructor(name: string, properties: DatabaseProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Glue:Database", name, inputs, opts);
        }
    }
    
    export class DevEndpoint extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DevEndpointAttributes>;

        constructor(name: string, properties: DevEndpointProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Glue:DevEndpoint", name, inputs, opts);
        }
    }
    
    export class Job extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<JobAttributes>;

        constructor(name: string, properties: JobProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Glue:Job", name, inputs, opts);
        }
    }
    
    export class MLTransform extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<MLTransformAttributes>;

        constructor(name: string, properties: MLTransformProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Glue:MLTransform", name, inputs, opts);
        }
    }
    
    export class Partition extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PartitionAttributes>;

        constructor(name: string, properties: PartitionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Glue:Partition", name, inputs, opts);
        }
    }
    
    export class SecurityConfiguration extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SecurityConfigurationAttributes>;

        constructor(name: string, properties: SecurityConfigurationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Glue:SecurityConfiguration", name, inputs, opts);
        }
    }
    
    export class Table extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TableAttributes>;

        constructor(name: string, properties: TableProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Glue:Table", name, inputs, opts);
        }
    }
    
    export class Trigger extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TriggerAttributes>;

        constructor(name: string, properties: TriggerProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Glue:Trigger", name, inputs, opts);
        }
    }
    
    export class Workflow extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<WorkflowAttributes>;

        constructor(name: string, properties?: WorkflowProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Glue:Workflow", name, inputs, opts);
        }
    }
    
}

export namespace greengrass {
    
    export interface ConnectorDefinitionAttributes {
        Arn: string;
        Id: string;
        LatestVersionArn: string;
        Name: string;
    }
    
    export interface ConnectorDefinitionConnectorAttributes {
        ConnectorArn: string;
        Id: string;
        Parameters?: string;
    }
    
    export interface ConnectorDefinitionConnectorDefinitionVersionAttributes {
        Connectors: ConnectorDefinitionConnectorAttributes[];
    }
    
    export interface ConnectorDefinitionConnectorDefinitionVersionProperties {
        Connectors: pulumi.Input<pulumi.Input<ConnectorDefinitionConnectorProperties>[]>;
    }
    
    export interface ConnectorDefinitionConnectorProperties {
        ConnectorArn: pulumi.Input<string>;
        Id: pulumi.Input<string>;
        Parameters?: pulumi.Input<string>;
    }
    
    export interface ConnectorDefinitionProperties {
        InitialVersion?: pulumi.Input<ConnectorDefinitionConnectorDefinitionVersionProperties>;
        Name: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface ConnectorDefinitionVersionAttributes {
    }
    
    export interface ConnectorDefinitionVersionConnectorAttributes {
        ConnectorArn: string;
        Id: string;
        Parameters?: string;
    }
    
    export interface ConnectorDefinitionVersionConnectorProperties {
        ConnectorArn: pulumi.Input<string>;
        Id: pulumi.Input<string>;
        Parameters?: pulumi.Input<string>;
    }
    
    export interface ConnectorDefinitionVersionProperties {
        ConnectorDefinitionId: pulumi.Input<string>;
        Connectors: pulumi.Input<pulumi.Input<ConnectorDefinitionVersionConnectorProperties>[]>;
    }
    
    export interface CoreDefinitionAttributes {
        Arn: string;
        Id: string;
        LatestVersionArn: string;
        Name: string;
    }
    
    export interface CoreDefinitionCoreAttributes {
        CertificateArn: string;
        Id: string;
        SyncShadow?: boolean;
        ThingArn: string;
    }
    
    export interface CoreDefinitionCoreDefinitionVersionAttributes {
        Cores: CoreDefinitionCoreAttributes[];
    }
    
    export interface CoreDefinitionCoreDefinitionVersionProperties {
        Cores: pulumi.Input<pulumi.Input<CoreDefinitionCoreProperties>[]>;
    }
    
    export interface CoreDefinitionCoreProperties {
        CertificateArn: pulumi.Input<string>;
        Id: pulumi.Input<string>;
        SyncShadow?: pulumi.Input<boolean>;
        ThingArn: pulumi.Input<string>;
    }
    
    export interface CoreDefinitionProperties {
        InitialVersion?: pulumi.Input<CoreDefinitionCoreDefinitionVersionProperties>;
        Name: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface CoreDefinitionVersionAttributes {
    }
    
    export interface CoreDefinitionVersionCoreAttributes {
        CertificateArn: string;
        Id: string;
        SyncShadow?: boolean;
        ThingArn: string;
    }
    
    export interface CoreDefinitionVersionCoreProperties {
        CertificateArn: pulumi.Input<string>;
        Id: pulumi.Input<string>;
        SyncShadow?: pulumi.Input<boolean>;
        ThingArn: pulumi.Input<string>;
    }
    
    export interface CoreDefinitionVersionProperties {
        CoreDefinitionId: pulumi.Input<string>;
        Cores: pulumi.Input<pulumi.Input<CoreDefinitionVersionCoreProperties>[]>;
    }
    
    export interface DeviceDefinitionAttributes {
        Arn: string;
        Id: string;
        LatestVersionArn: string;
        Name: string;
    }
    
    export interface DeviceDefinitionDeviceAttributes {
        CertificateArn: string;
        Id: string;
        SyncShadow?: boolean;
        ThingArn: string;
    }
    
    export interface DeviceDefinitionDeviceDefinitionVersionAttributes {
        Devices: DeviceDefinitionDeviceAttributes[];
    }
    
    export interface DeviceDefinitionDeviceDefinitionVersionProperties {
        Devices: pulumi.Input<pulumi.Input<DeviceDefinitionDeviceProperties>[]>;
    }
    
    export interface DeviceDefinitionDeviceProperties {
        CertificateArn: pulumi.Input<string>;
        Id: pulumi.Input<string>;
        SyncShadow?: pulumi.Input<boolean>;
        ThingArn: pulumi.Input<string>;
    }
    
    export interface DeviceDefinitionProperties {
        InitialVersion?: pulumi.Input<DeviceDefinitionDeviceDefinitionVersionProperties>;
        Name: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface DeviceDefinitionVersionAttributes {
    }
    
    export interface DeviceDefinitionVersionDeviceAttributes {
        CertificateArn: string;
        Id: string;
        SyncShadow?: boolean;
        ThingArn: string;
    }
    
    export interface DeviceDefinitionVersionDeviceProperties {
        CertificateArn: pulumi.Input<string>;
        Id: pulumi.Input<string>;
        SyncShadow?: pulumi.Input<boolean>;
        ThingArn: pulumi.Input<string>;
    }
    
    export interface DeviceDefinitionVersionProperties {
        DeviceDefinitionId: pulumi.Input<string>;
        Devices: pulumi.Input<pulumi.Input<DeviceDefinitionVersionDeviceProperties>[]>;
    }
    
    export interface FunctionDefinitionAttributes {
        Arn: string;
        Id: string;
        LatestVersionArn: string;
        Name: string;
    }
    
    export interface FunctionDefinitionDefaultConfigAttributes {
        Execution: FunctionDefinitionExecutionAttributes;
    }
    
    export interface FunctionDefinitionDefaultConfigProperties {
        Execution: pulumi.Input<FunctionDefinitionExecutionProperties>;
    }
    
    export interface FunctionDefinitionEnvironmentAttributes {
        AccessSysfs?: boolean;
        Execution?: FunctionDefinitionExecutionAttributes;
        ResourceAccessPolicies?: FunctionDefinitionResourceAccessPolicyAttributes[];
        Variables?: string;
    }
    
    export interface FunctionDefinitionEnvironmentProperties {
        AccessSysfs?: pulumi.Input<boolean>;
        Execution?: pulumi.Input<FunctionDefinitionExecutionProperties>;
        ResourceAccessPolicies?: pulumi.Input<pulumi.Input<FunctionDefinitionResourceAccessPolicyProperties>[]>;
        Variables?: pulumi.Input<string>;
    }
    
    export interface FunctionDefinitionExecutionAttributes {
        IsolationMode?: string;
        RunAs?: FunctionDefinitionRunAsAttributes;
    }
    
    export interface FunctionDefinitionExecutionProperties {
        IsolationMode?: pulumi.Input<string>;
        RunAs?: pulumi.Input<FunctionDefinitionRunAsProperties>;
    }
    
    export interface FunctionDefinitionFunctionAttributes {
        FunctionArn: string;
        FunctionConfiguration: FunctionDefinitionFunctionConfigurationAttributes;
        Id: string;
    }
    
    export interface FunctionDefinitionFunctionConfigurationAttributes {
        EncodingType?: string;
        Environment?: FunctionDefinitionEnvironmentAttributes;
        ExecArgs?: string;
        Executable?: string;
        MemorySize?: number;
        Pinned?: boolean;
        Timeout?: number;
    }
    
    export interface FunctionDefinitionFunctionConfigurationProperties {
        EncodingType?: pulumi.Input<string>;
        Environment?: pulumi.Input<FunctionDefinitionEnvironmentProperties>;
        ExecArgs?: pulumi.Input<string>;
        Executable?: pulumi.Input<string>;
        MemorySize?: pulumi.Input<number>;
        Pinned?: pulumi.Input<boolean>;
        Timeout?: pulumi.Input<number>;
    }
    
    export interface FunctionDefinitionFunctionDefinitionVersionAttributes {
        DefaultConfig?: FunctionDefinitionDefaultConfigAttributes;
        Functions: FunctionDefinitionFunctionAttributes[];
    }
    
    export interface FunctionDefinitionFunctionDefinitionVersionProperties {
        DefaultConfig?: pulumi.Input<FunctionDefinitionDefaultConfigProperties>;
        Functions: pulumi.Input<pulumi.Input<FunctionDefinitionFunctionProperties>[]>;
    }
    
    export interface FunctionDefinitionFunctionProperties {
        FunctionArn: pulumi.Input<string>;
        FunctionConfiguration: pulumi.Input<FunctionDefinitionFunctionConfigurationProperties>;
        Id: pulumi.Input<string>;
    }
    
    export interface FunctionDefinitionProperties {
        InitialVersion?: pulumi.Input<FunctionDefinitionFunctionDefinitionVersionProperties>;
        Name: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface FunctionDefinitionResourceAccessPolicyAttributes {
        Permission?: string;
        ResourceId: string;
    }
    
    export interface FunctionDefinitionResourceAccessPolicyProperties {
        Permission?: pulumi.Input<string>;
        ResourceId: pulumi.Input<string>;
    }
    
    export interface FunctionDefinitionRunAsAttributes {
        Gid?: number;
        Uid?: number;
    }
    
    export interface FunctionDefinitionRunAsProperties {
        Gid?: pulumi.Input<number>;
        Uid?: pulumi.Input<number>;
    }
    
    export interface FunctionDefinitionVersionAttributes {
    }
    
    export interface FunctionDefinitionVersionDefaultConfigAttributes {
        Execution: FunctionDefinitionVersionExecutionAttributes;
    }
    
    export interface FunctionDefinitionVersionDefaultConfigProperties {
        Execution: pulumi.Input<FunctionDefinitionVersionExecutionProperties>;
    }
    
    export interface FunctionDefinitionVersionEnvironmentAttributes {
        AccessSysfs?: boolean;
        Execution?: FunctionDefinitionVersionExecutionAttributes;
        ResourceAccessPolicies?: FunctionDefinitionVersionResourceAccessPolicyAttributes[];
        Variables?: string;
    }
    
    export interface FunctionDefinitionVersionEnvironmentProperties {
        AccessSysfs?: pulumi.Input<boolean>;
        Execution?: pulumi.Input<FunctionDefinitionVersionExecutionProperties>;
        ResourceAccessPolicies?: pulumi.Input<pulumi.Input<FunctionDefinitionVersionResourceAccessPolicyProperties>[]>;
        Variables?: pulumi.Input<string>;
    }
    
    export interface FunctionDefinitionVersionExecutionAttributes {
        IsolationMode?: string;
        RunAs?: FunctionDefinitionVersionRunAsAttributes;
    }
    
    export interface FunctionDefinitionVersionExecutionProperties {
        IsolationMode?: pulumi.Input<string>;
        RunAs?: pulumi.Input<FunctionDefinitionVersionRunAsProperties>;
    }
    
    export interface FunctionDefinitionVersionFunctionAttributes {
        FunctionArn: string;
        FunctionConfiguration: FunctionDefinitionVersionFunctionConfigurationAttributes;
        Id: string;
    }
    
    export interface FunctionDefinitionVersionFunctionConfigurationAttributes {
        EncodingType?: string;
        Environment?: FunctionDefinitionVersionEnvironmentAttributes;
        ExecArgs?: string;
        Executable?: string;
        MemorySize?: number;
        Pinned?: boolean;
        Timeout?: number;
    }
    
    export interface FunctionDefinitionVersionFunctionConfigurationProperties {
        EncodingType?: pulumi.Input<string>;
        Environment?: pulumi.Input<FunctionDefinitionVersionEnvironmentProperties>;
        ExecArgs?: pulumi.Input<string>;
        Executable?: pulumi.Input<string>;
        MemorySize?: pulumi.Input<number>;
        Pinned?: pulumi.Input<boolean>;
        Timeout?: pulumi.Input<number>;
    }
    
    export interface FunctionDefinitionVersionFunctionProperties {
        FunctionArn: pulumi.Input<string>;
        FunctionConfiguration: pulumi.Input<FunctionDefinitionVersionFunctionConfigurationProperties>;
        Id: pulumi.Input<string>;
    }
    
    export interface FunctionDefinitionVersionProperties {
        DefaultConfig?: pulumi.Input<FunctionDefinitionVersionDefaultConfigProperties>;
        FunctionDefinitionId: pulumi.Input<string>;
        Functions: pulumi.Input<pulumi.Input<FunctionDefinitionVersionFunctionProperties>[]>;
    }
    
    export interface FunctionDefinitionVersionResourceAccessPolicyAttributes {
        Permission?: string;
        ResourceId: string;
    }
    
    export interface FunctionDefinitionVersionResourceAccessPolicyProperties {
        Permission?: pulumi.Input<string>;
        ResourceId: pulumi.Input<string>;
    }
    
    export interface FunctionDefinitionVersionRunAsAttributes {
        Gid?: number;
        Uid?: number;
    }
    
    export interface FunctionDefinitionVersionRunAsProperties {
        Gid?: pulumi.Input<number>;
        Uid?: pulumi.Input<number>;
    }
    
    export interface GroupAttributes {
        Arn: string;
        Id: string;
        LatestVersionArn: string;
        Name: string;
        RoleArn: string;
        RoleAttachedAt: string;
    }
    
    export interface GroupGroupVersionAttributes {
        ConnectorDefinitionVersionArn?: string;
        CoreDefinitionVersionArn?: string;
        DeviceDefinitionVersionArn?: string;
        FunctionDefinitionVersionArn?: string;
        LoggerDefinitionVersionArn?: string;
        ResourceDefinitionVersionArn?: string;
        SubscriptionDefinitionVersionArn?: string;
    }
    
    export interface GroupGroupVersionProperties {
        ConnectorDefinitionVersionArn?: pulumi.Input<string>;
        CoreDefinitionVersionArn?: pulumi.Input<string>;
        DeviceDefinitionVersionArn?: pulumi.Input<string>;
        FunctionDefinitionVersionArn?: pulumi.Input<string>;
        LoggerDefinitionVersionArn?: pulumi.Input<string>;
        ResourceDefinitionVersionArn?: pulumi.Input<string>;
        SubscriptionDefinitionVersionArn?: pulumi.Input<string>;
    }
    
    export interface GroupProperties {
        InitialVersion?: pulumi.Input<GroupGroupVersionProperties>;
        Name: pulumi.Input<string>;
        RoleArn?: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface GroupVersionAttributes {
    }
    
    export interface GroupVersionProperties {
        ConnectorDefinitionVersionArn?: pulumi.Input<string>;
        CoreDefinitionVersionArn?: pulumi.Input<string>;
        DeviceDefinitionVersionArn?: pulumi.Input<string>;
        FunctionDefinitionVersionArn?: pulumi.Input<string>;
        GroupId: pulumi.Input<string>;
        LoggerDefinitionVersionArn?: pulumi.Input<string>;
        ResourceDefinitionVersionArn?: pulumi.Input<string>;
        SubscriptionDefinitionVersionArn?: pulumi.Input<string>;
    }
    
    export interface LoggerDefinitionAttributes {
        Arn: string;
        Id: string;
        LatestVersionArn: string;
        Name: string;
    }
    
    export interface LoggerDefinitionLoggerAttributes {
        Component: string;
        Id: string;
        Level: string;
        Space?: number;
        Type: string;
    }
    
    export interface LoggerDefinitionLoggerDefinitionVersionAttributes {
        Loggers: LoggerDefinitionLoggerAttributes[];
    }
    
    export interface LoggerDefinitionLoggerDefinitionVersionProperties {
        Loggers: pulumi.Input<pulumi.Input<LoggerDefinitionLoggerProperties>[]>;
    }
    
    export interface LoggerDefinitionLoggerProperties {
        Component: pulumi.Input<string>;
        Id: pulumi.Input<string>;
        Level: pulumi.Input<string>;
        Space?: pulumi.Input<number>;
        Type: pulumi.Input<string>;
    }
    
    export interface LoggerDefinitionProperties {
        InitialVersion?: pulumi.Input<LoggerDefinitionLoggerDefinitionVersionProperties>;
        Name: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface LoggerDefinitionVersionAttributes {
    }
    
    export interface LoggerDefinitionVersionLoggerAttributes {
        Component: string;
        Id: string;
        Level: string;
        Space?: number;
        Type: string;
    }
    
    export interface LoggerDefinitionVersionLoggerProperties {
        Component: pulumi.Input<string>;
        Id: pulumi.Input<string>;
        Level: pulumi.Input<string>;
        Space?: pulumi.Input<number>;
        Type: pulumi.Input<string>;
    }
    
    export interface LoggerDefinitionVersionProperties {
        LoggerDefinitionId: pulumi.Input<string>;
        Loggers: pulumi.Input<pulumi.Input<LoggerDefinitionVersionLoggerProperties>[]>;
    }
    
    export interface ResourceDefinitionAttributes {
        Arn: string;
        Id: string;
        LatestVersionArn: string;
        Name: string;
    }
    
    export interface ResourceDefinitionGroupOwnerSettingAttributes {
        AutoAddGroupOwner: boolean;
        GroupOwner?: string;
    }
    
    export interface ResourceDefinitionGroupOwnerSettingProperties {
        AutoAddGroupOwner: pulumi.Input<boolean>;
        GroupOwner?: pulumi.Input<string>;
    }
    
    export interface ResourceDefinitionLocalDeviceResourceDataAttributes {
        GroupOwnerSetting?: ResourceDefinitionGroupOwnerSettingAttributes;
        SourcePath: string;
    }
    
    export interface ResourceDefinitionLocalDeviceResourceDataProperties {
        GroupOwnerSetting?: pulumi.Input<ResourceDefinitionGroupOwnerSettingProperties>;
        SourcePath: pulumi.Input<string>;
    }
    
    export interface ResourceDefinitionLocalVolumeResourceDataAttributes {
        DestinationPath: string;
        GroupOwnerSetting?: ResourceDefinitionGroupOwnerSettingAttributes;
        SourcePath: string;
    }
    
    export interface ResourceDefinitionLocalVolumeResourceDataProperties {
        DestinationPath: pulumi.Input<string>;
        GroupOwnerSetting?: pulumi.Input<ResourceDefinitionGroupOwnerSettingProperties>;
        SourcePath: pulumi.Input<string>;
    }
    
    export interface ResourceDefinitionProperties {
        InitialVersion?: pulumi.Input<ResourceDefinitionResourceDefinitionVersionProperties>;
        Name: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface ResourceDefinitionResourceDataContainerAttributes {
        LocalDeviceResourceData?: ResourceDefinitionLocalDeviceResourceDataAttributes;
        LocalVolumeResourceData?: ResourceDefinitionLocalVolumeResourceDataAttributes;
        S3MachineLearningModelResourceData?: ResourceDefinitionS3MachineLearningModelResourceDataAttributes;
        SageMakerMachineLearningModelResourceData?: ResourceDefinitionSageMakerMachineLearningModelResourceDataAttributes;
        SecretsManagerSecretResourceData?: ResourceDefinitionSecretsManagerSecretResourceDataAttributes;
    }
    
    export interface ResourceDefinitionResourceDataContainerProperties {
        LocalDeviceResourceData?: pulumi.Input<ResourceDefinitionLocalDeviceResourceDataProperties>;
        LocalVolumeResourceData?: pulumi.Input<ResourceDefinitionLocalVolumeResourceDataProperties>;
        S3MachineLearningModelResourceData?: pulumi.Input<ResourceDefinitionS3MachineLearningModelResourceDataProperties>;
        SageMakerMachineLearningModelResourceData?: pulumi.Input<ResourceDefinitionSageMakerMachineLearningModelResourceDataProperties>;
        SecretsManagerSecretResourceData?: pulumi.Input<ResourceDefinitionSecretsManagerSecretResourceDataProperties>;
    }
    
    export interface ResourceDefinitionResourceDefinitionVersionAttributes {
        Resources: ResourceDefinitionResourceInstanceAttributes[];
    }
    
    export interface ResourceDefinitionResourceDefinitionVersionProperties {
        Resources: pulumi.Input<pulumi.Input<ResourceDefinitionResourceInstanceProperties>[]>;
    }
    
    export interface ResourceDefinitionResourceInstanceAttributes {
        Id: string;
        Name: string;
        ResourceDataContainer: ResourceDefinitionResourceDataContainerAttributes;
    }
    
    export interface ResourceDefinitionResourceInstanceProperties {
        Id: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        ResourceDataContainer: pulumi.Input<ResourceDefinitionResourceDataContainerProperties>;
    }
    
    export interface ResourceDefinitionS3MachineLearningModelResourceDataAttributes {
        DestinationPath: string;
        S3Uri: string;
    }
    
    export interface ResourceDefinitionS3MachineLearningModelResourceDataProperties {
        DestinationPath: pulumi.Input<string>;
        S3Uri: pulumi.Input<string>;
    }
    
    export interface ResourceDefinitionSageMakerMachineLearningModelResourceDataAttributes {
        DestinationPath: string;
        SageMakerJobArn: string;
    }
    
    export interface ResourceDefinitionSageMakerMachineLearningModelResourceDataProperties {
        DestinationPath: pulumi.Input<string>;
        SageMakerJobArn: pulumi.Input<string>;
    }
    
    export interface ResourceDefinitionSecretsManagerSecretResourceDataAttributes {
        ARN: string;
        AdditionalStagingLabelsToDownload?: string[];
    }
    
    export interface ResourceDefinitionSecretsManagerSecretResourceDataProperties {
        ARN: pulumi.Input<string>;
        AdditionalStagingLabelsToDownload?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface ResourceDefinitionVersionAttributes {
    }
    
    export interface ResourceDefinitionVersionGroupOwnerSettingAttributes {
        AutoAddGroupOwner: boolean;
        GroupOwner?: string;
    }
    
    export interface ResourceDefinitionVersionGroupOwnerSettingProperties {
        AutoAddGroupOwner: pulumi.Input<boolean>;
        GroupOwner?: pulumi.Input<string>;
    }
    
    export interface ResourceDefinitionVersionLocalDeviceResourceDataAttributes {
        GroupOwnerSetting?: ResourceDefinitionVersionGroupOwnerSettingAttributes;
        SourcePath: string;
    }
    
    export interface ResourceDefinitionVersionLocalDeviceResourceDataProperties {
        GroupOwnerSetting?: pulumi.Input<ResourceDefinitionVersionGroupOwnerSettingProperties>;
        SourcePath: pulumi.Input<string>;
    }
    
    export interface ResourceDefinitionVersionLocalVolumeResourceDataAttributes {
        DestinationPath: string;
        GroupOwnerSetting?: ResourceDefinitionVersionGroupOwnerSettingAttributes;
        SourcePath: string;
    }
    
    export interface ResourceDefinitionVersionLocalVolumeResourceDataProperties {
        DestinationPath: pulumi.Input<string>;
        GroupOwnerSetting?: pulumi.Input<ResourceDefinitionVersionGroupOwnerSettingProperties>;
        SourcePath: pulumi.Input<string>;
    }
    
    export interface ResourceDefinitionVersionProperties {
        ResourceDefinitionId: pulumi.Input<string>;
        Resources: pulumi.Input<pulumi.Input<ResourceDefinitionVersionResourceInstanceProperties>[]>;
    }
    
    export interface ResourceDefinitionVersionResourceDataContainerAttributes {
        LocalDeviceResourceData?: ResourceDefinitionVersionLocalDeviceResourceDataAttributes;
        LocalVolumeResourceData?: ResourceDefinitionVersionLocalVolumeResourceDataAttributes;
        S3MachineLearningModelResourceData?: ResourceDefinitionVersionS3MachineLearningModelResourceDataAttributes;
        SageMakerMachineLearningModelResourceData?: ResourceDefinitionVersionSageMakerMachineLearningModelResourceDataAttributes;
        SecretsManagerSecretResourceData?: ResourceDefinitionVersionSecretsManagerSecretResourceDataAttributes;
    }
    
    export interface ResourceDefinitionVersionResourceDataContainerProperties {
        LocalDeviceResourceData?: pulumi.Input<ResourceDefinitionVersionLocalDeviceResourceDataProperties>;
        LocalVolumeResourceData?: pulumi.Input<ResourceDefinitionVersionLocalVolumeResourceDataProperties>;
        S3MachineLearningModelResourceData?: pulumi.Input<ResourceDefinitionVersionS3MachineLearningModelResourceDataProperties>;
        SageMakerMachineLearningModelResourceData?: pulumi.Input<ResourceDefinitionVersionSageMakerMachineLearningModelResourceDataProperties>;
        SecretsManagerSecretResourceData?: pulumi.Input<ResourceDefinitionVersionSecretsManagerSecretResourceDataProperties>;
    }
    
    export interface ResourceDefinitionVersionResourceInstanceAttributes {
        Id: string;
        Name: string;
        ResourceDataContainer: ResourceDefinitionVersionResourceDataContainerAttributes;
    }
    
    export interface ResourceDefinitionVersionResourceInstanceProperties {
        Id: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        ResourceDataContainer: pulumi.Input<ResourceDefinitionVersionResourceDataContainerProperties>;
    }
    
    export interface ResourceDefinitionVersionS3MachineLearningModelResourceDataAttributes {
        DestinationPath: string;
        S3Uri: string;
    }
    
    export interface ResourceDefinitionVersionS3MachineLearningModelResourceDataProperties {
        DestinationPath: pulumi.Input<string>;
        S3Uri: pulumi.Input<string>;
    }
    
    export interface ResourceDefinitionVersionSageMakerMachineLearningModelResourceDataAttributes {
        DestinationPath: string;
        SageMakerJobArn: string;
    }
    
    export interface ResourceDefinitionVersionSageMakerMachineLearningModelResourceDataProperties {
        DestinationPath: pulumi.Input<string>;
        SageMakerJobArn: pulumi.Input<string>;
    }
    
    export interface ResourceDefinitionVersionSecretsManagerSecretResourceDataAttributes {
        ARN: string;
        AdditionalStagingLabelsToDownload?: string[];
    }
    
    export interface ResourceDefinitionVersionSecretsManagerSecretResourceDataProperties {
        ARN: pulumi.Input<string>;
        AdditionalStagingLabelsToDownload?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface SubscriptionDefinitionAttributes {
        Arn: string;
        Id: string;
        LatestVersionArn: string;
        Name: string;
    }
    
    export interface SubscriptionDefinitionProperties {
        InitialVersion?: pulumi.Input<SubscriptionDefinitionSubscriptionDefinitionVersionProperties>;
        Name: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface SubscriptionDefinitionSubscriptionAttributes {
        Id: string;
        Source: string;
        Subject: string;
        Target: string;
    }
    
    export interface SubscriptionDefinitionSubscriptionDefinitionVersionAttributes {
        Subscriptions: SubscriptionDefinitionSubscriptionAttributes[];
    }
    
    export interface SubscriptionDefinitionSubscriptionDefinitionVersionProperties {
        Subscriptions: pulumi.Input<pulumi.Input<SubscriptionDefinitionSubscriptionProperties>[]>;
    }
    
    export interface SubscriptionDefinitionSubscriptionProperties {
        Id: pulumi.Input<string>;
        Source: pulumi.Input<string>;
        Subject: pulumi.Input<string>;
        Target: pulumi.Input<string>;
    }
    
    export interface SubscriptionDefinitionVersionAttributes {
    }
    
    export interface SubscriptionDefinitionVersionProperties {
        SubscriptionDefinitionId: pulumi.Input<string>;
        Subscriptions: pulumi.Input<pulumi.Input<SubscriptionDefinitionVersionSubscriptionProperties>[]>;
    }
    
    export interface SubscriptionDefinitionVersionSubscriptionAttributes {
        Id: string;
        Source: string;
        Subject: string;
        Target: string;
    }
    
    export interface SubscriptionDefinitionVersionSubscriptionProperties {
        Id: pulumi.Input<string>;
        Source: pulumi.Input<string>;
        Subject: pulumi.Input<string>;
        Target: pulumi.Input<string>;
    }
    
    
    export class ConnectorDefinition extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ConnectorDefinitionAttributes>;

        constructor(name: string, properties: ConnectorDefinitionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Greengrass:ConnectorDefinition", name, inputs, opts);
        }
    }
    
    export class ConnectorDefinitionVersion extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ConnectorDefinitionVersionAttributes>;

        constructor(name: string, properties: ConnectorDefinitionVersionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Greengrass:ConnectorDefinitionVersion", name, inputs, opts);
        }
    }
    
    export class CoreDefinition extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<CoreDefinitionAttributes>;

        constructor(name: string, properties: CoreDefinitionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Greengrass:CoreDefinition", name, inputs, opts);
        }
    }
    
    export class CoreDefinitionVersion extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<CoreDefinitionVersionAttributes>;

        constructor(name: string, properties: CoreDefinitionVersionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Greengrass:CoreDefinitionVersion", name, inputs, opts);
        }
    }
    
    export class DeviceDefinition extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DeviceDefinitionAttributes>;

        constructor(name: string, properties: DeviceDefinitionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Greengrass:DeviceDefinition", name, inputs, opts);
        }
    }
    
    export class DeviceDefinitionVersion extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DeviceDefinitionVersionAttributes>;

        constructor(name: string, properties: DeviceDefinitionVersionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Greengrass:DeviceDefinitionVersion", name, inputs, opts);
        }
    }
    
    export class FunctionDefinition extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<FunctionDefinitionAttributes>;

        constructor(name: string, properties: FunctionDefinitionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Greengrass:FunctionDefinition", name, inputs, opts);
        }
    }
    
    export class FunctionDefinitionVersion extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<FunctionDefinitionVersionAttributes>;

        constructor(name: string, properties: FunctionDefinitionVersionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Greengrass:FunctionDefinitionVersion", name, inputs, opts);
        }
    }
    
    export class Group extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<GroupAttributes>;

        constructor(name: string, properties: GroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Greengrass:Group", name, inputs, opts);
        }
    }
    
    export class GroupVersion extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<GroupVersionAttributes>;

        constructor(name: string, properties: GroupVersionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Greengrass:GroupVersion", name, inputs, opts);
        }
    }
    
    export class LoggerDefinition extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<LoggerDefinitionAttributes>;

        constructor(name: string, properties: LoggerDefinitionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Greengrass:LoggerDefinition", name, inputs, opts);
        }
    }
    
    export class LoggerDefinitionVersion extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<LoggerDefinitionVersionAttributes>;

        constructor(name: string, properties: LoggerDefinitionVersionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Greengrass:LoggerDefinitionVersion", name, inputs, opts);
        }
    }
    
    export class ResourceDefinition extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ResourceDefinitionAttributes>;

        constructor(name: string, properties: ResourceDefinitionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Greengrass:ResourceDefinition", name, inputs, opts);
        }
    }
    
    export class ResourceDefinitionVersion extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ResourceDefinitionVersionAttributes>;

        constructor(name: string, properties: ResourceDefinitionVersionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Greengrass:ResourceDefinitionVersion", name, inputs, opts);
        }
    }
    
    export class SubscriptionDefinition extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SubscriptionDefinitionAttributes>;

        constructor(name: string, properties: SubscriptionDefinitionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Greengrass:SubscriptionDefinition", name, inputs, opts);
        }
    }
    
    export class SubscriptionDefinitionVersion extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SubscriptionDefinitionVersionAttributes>;

        constructor(name: string, properties: SubscriptionDefinitionVersionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Greengrass:SubscriptionDefinitionVersion", name, inputs, opts);
        }
    }
    
}

export namespace guardduty {
    
    export interface DetectorAttributes {
    }
    
    export interface DetectorProperties {
        Enable: pulumi.Input<boolean>;
        FindingPublishingFrequency?: pulumi.Input<string>;
    }
    
    export interface FilterAttributes {
    }
    
    export interface FilterConditionAttributes {
        Eq?: string[];
        Gte?: number;
        Lt?: number;
        Lte?: number;
        Neq?: string[];
    }
    
    export interface FilterConditionProperties {
        Eq?: pulumi.Input<pulumi.Input<string>[]>;
        Gte?: pulumi.Input<number>;
        Lt?: pulumi.Input<number>;
        Lte?: pulumi.Input<number>;
        Neq?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface FilterFindingCriteriaAttributes {
        Criterion?: string;
        ItemType?: FilterConditionAttributes;
    }
    
    export interface FilterFindingCriteriaProperties {
        Criterion?: pulumi.Input<string>;
        ItemType?: pulumi.Input<FilterConditionProperties>;
    }
    
    export interface FilterProperties {
        Action: pulumi.Input<string>;
        Description: pulumi.Input<string>;
        DetectorId: pulumi.Input<string>;
        FindingCriteria: pulumi.Input<FilterFindingCriteriaProperties>;
        Name?: pulumi.Input<string>;
        Rank: pulumi.Input<number>;
    }
    
    export interface IPSetAttributes {
    }
    
    export interface IPSetProperties {
        Activate: pulumi.Input<boolean>;
        DetectorId: pulumi.Input<string>;
        Format: pulumi.Input<string>;
        Location: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
    }
    
    export interface MasterAttributes {
    }
    
    export interface MasterProperties {
        DetectorId: pulumi.Input<string>;
        InvitationId?: pulumi.Input<string>;
        MasterId: pulumi.Input<string>;
    }
    
    export interface MemberAttributes {
    }
    
    export interface MemberProperties {
        DetectorId: pulumi.Input<string>;
        DisableEmailNotification?: pulumi.Input<boolean>;
        Email: pulumi.Input<string>;
        MemberId: pulumi.Input<string>;
        Message?: pulumi.Input<string>;
        Status?: pulumi.Input<string>;
    }
    
    export interface ThreatIntelSetAttributes {
    }
    
    export interface ThreatIntelSetProperties {
        Activate: pulumi.Input<boolean>;
        DetectorId: pulumi.Input<string>;
        Format: pulumi.Input<string>;
        Location: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
    }
    
    
    export class Detector extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DetectorAttributes>;

        constructor(name: string, properties: DetectorProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:GuardDuty:Detector", name, inputs, opts);
        }
    }
    
    export class Filter extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<FilterAttributes>;

        constructor(name: string, properties: FilterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:GuardDuty:Filter", name, inputs, opts);
        }
    }
    
    export class IPSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<IPSetAttributes>;

        constructor(name: string, properties: IPSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:GuardDuty:IPSet", name, inputs, opts);
        }
    }
    
    export class Master extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<MasterAttributes>;

        constructor(name: string, properties: MasterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:GuardDuty:Master", name, inputs, opts);
        }
    }
    
    export class Member extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<MemberAttributes>;

        constructor(name: string, properties: MemberProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:GuardDuty:Member", name, inputs, opts);
        }
    }
    
    export class ThreatIntelSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ThreatIntelSetAttributes>;

        constructor(name: string, properties: ThreatIntelSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:GuardDuty:ThreatIntelSet", name, inputs, opts);
        }
    }
    
}

export namespace iam {
    
    export interface AccessKeyAttributes {
        SecretAccessKey: string;
    }
    
    export interface AccessKeyProperties {
        Serial?: pulumi.Input<number>;
        Status?: pulumi.Input<string>;
        UserName: pulumi.Input<string>;
    }
    
    export interface GroupAttributes {
        Arn: string;
    }
    
    export interface GroupPolicyAttributes {
        PolicyDocument: string;
        PolicyName: string;
    }
    
    export interface GroupPolicyProperties {
        PolicyDocument: pulumi.Input<string>;
        PolicyName: pulumi.Input<string>;
    }
    
    export interface GroupProperties {
        GroupName?: pulumi.Input<string>;
        ManagedPolicyArns?: pulumi.Input<pulumi.Input<string>[]>;
        Path?: pulumi.Input<string>;
        Policies?: pulumi.Input<pulumi.Input<GroupPolicyProperties>[]>;
    }
    
    export interface InstanceProfileAttributes {
        Arn: string;
    }
    
    export interface InstanceProfileProperties {
        InstanceProfileName?: pulumi.Input<string>;
        Path?: pulumi.Input<string>;
        Roles: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface ManagedPolicyAttributes {
    }
    
    export interface ManagedPolicyProperties {
        Description?: pulumi.Input<string>;
        Groups?: pulumi.Input<pulumi.Input<string>[]>;
        ManagedPolicyName?: pulumi.Input<string>;
        Path?: pulumi.Input<string>;
        PolicyDocument: pulumi.Input<string>;
        Roles?: pulumi.Input<pulumi.Input<string>[]>;
        Users?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface PolicyAttributes {
    }
    
    export interface PolicyProperties {
        Groups?: pulumi.Input<pulumi.Input<string>[]>;
        PolicyDocument: pulumi.Input<string>;
        PolicyName: pulumi.Input<string>;
        Roles?: pulumi.Input<pulumi.Input<string>[]>;
        Users?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface RoleAttributes {
        Arn: string;
        RoleId: string;
    }
    
    export interface RolePolicyAttributes {
        PolicyDocument: string;
        PolicyName: string;
    }
    
    export interface RolePolicyProperties {
        PolicyDocument: pulumi.Input<string>;
        PolicyName: pulumi.Input<string>;
    }
    
    export interface RoleProperties {
        AssumeRolePolicyDocument: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        ManagedPolicyArns?: pulumi.Input<pulumi.Input<string>[]>;
        MaxSessionDuration?: pulumi.Input<number>;
        Path?: pulumi.Input<string>;
        PermissionsBoundary?: pulumi.Input<string>;
        Policies?: pulumi.Input<pulumi.Input<RolePolicyProperties>[]>;
        RoleName?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface ServiceLinkedRoleAttributes {
    }
    
    export interface ServiceLinkedRoleProperties {
        AWSServiceName: pulumi.Input<string>;
        CustomSuffix?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
    }
    
    export interface UserAttributes {
        Arn: string;
    }
    
    export interface UserLoginProfileAttributes {
        Password: string;
        PasswordResetRequired?: boolean;
    }
    
    export interface UserLoginProfileProperties {
        Password: pulumi.Input<string>;
        PasswordResetRequired?: pulumi.Input<boolean>;
    }
    
    export interface UserPolicyAttributes {
        PolicyDocument: string;
        PolicyName: string;
    }
    
    export interface UserPolicyProperties {
        PolicyDocument: pulumi.Input<string>;
        PolicyName: pulumi.Input<string>;
    }
    
    export interface UserProperties {
        Groups?: pulumi.Input<pulumi.Input<string>[]>;
        LoginProfile?: pulumi.Input<UserLoginProfileProperties>;
        ManagedPolicyArns?: pulumi.Input<pulumi.Input<string>[]>;
        Path?: pulumi.Input<string>;
        PermissionsBoundary?: pulumi.Input<string>;
        Policies?: pulumi.Input<pulumi.Input<UserPolicyProperties>[]>;
        UserName?: pulumi.Input<string>;
    }
    
    export interface UserToGroupAdditionAttributes {
    }
    
    export interface UserToGroupAdditionProperties {
        GroupName: pulumi.Input<string>;
        Users: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    
    export class AccessKey extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AccessKeyAttributes>;

        constructor(name: string, properties: AccessKeyProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IAM:AccessKey", name, inputs, opts);
        }
    }
    
    export class Group extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<GroupAttributes>;

        constructor(name: string, properties?: GroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IAM:Group", name, inputs, opts);
        }
    }
    
    export class InstanceProfile extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<InstanceProfileAttributes>;

        constructor(name: string, properties: InstanceProfileProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IAM:InstanceProfile", name, inputs, opts);
        }
    }
    
    export class ManagedPolicy extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ManagedPolicyAttributes>;

        constructor(name: string, properties: ManagedPolicyProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IAM:ManagedPolicy", name, inputs, opts);
        }
    }
    
    export class Policy extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PolicyAttributes>;

        constructor(name: string, properties: PolicyProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IAM:Policy", name, inputs, opts);
        }
    }
    
    export class Role extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RoleAttributes>;

        constructor(name: string, properties: RoleProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IAM:Role", name, inputs, opts);
        }
    }
    
    export class ServiceLinkedRole extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ServiceLinkedRoleAttributes>;

        constructor(name: string, properties: ServiceLinkedRoleProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IAM:ServiceLinkedRole", name, inputs, opts);
        }
    }
    
    export class User extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<UserAttributes>;

        constructor(name: string, properties?: UserProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IAM:User", name, inputs, opts);
        }
    }
    
    export class UserToGroupAddition extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<UserToGroupAdditionAttributes>;

        constructor(name: string, properties: UserToGroupAdditionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IAM:UserToGroupAddition", name, inputs, opts);
        }
    }
    
}

export namespace inspector {
    
    export interface AssessmentTargetAttributes {
        Arn: string;
    }
    
    export interface AssessmentTargetProperties {
        AssessmentTargetName?: pulumi.Input<string>;
        ResourceGroupArn?: pulumi.Input<string>;
    }
    
    export interface AssessmentTemplateAttributes {
        Arn: string;
    }
    
    export interface AssessmentTemplateProperties {
        AssessmentTargetArn: pulumi.Input<string>;
        AssessmentTemplateName?: pulumi.Input<string>;
        DurationInSeconds: pulumi.Input<number>;
        RulesPackageArns: pulumi.Input<pulumi.Input<string>[]>;
        UserAttributesForFindings?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface ResourceGroupAttributes {
        Arn: string;
    }
    
    export interface ResourceGroupProperties {
        ResourceGroupTags: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    
    export class AssessmentTarget extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AssessmentTargetAttributes>;

        constructor(name: string, properties?: AssessmentTargetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Inspector:AssessmentTarget", name, inputs, opts);
        }
    }
    
    export class AssessmentTemplate extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AssessmentTemplateAttributes>;

        constructor(name: string, properties: AssessmentTemplateProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Inspector:AssessmentTemplate", name, inputs, opts);
        }
    }
    
    export class ResourceGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ResourceGroupAttributes>;

        constructor(name: string, properties: ResourceGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Inspector:ResourceGroup", name, inputs, opts);
        }
    }
    
}

export namespace iot {
    
    export interface CertificateAttributes {
        Arn: string;
    }
    
    export interface CertificateProperties {
        CertificateSigningRequest: pulumi.Input<string>;
        Status: pulumi.Input<string>;
    }
    
    export interface PolicyAttributes {
        Arn: string;
    }
    
    export interface PolicyPrincipalAttachmentAttributes {
    }
    
    export interface PolicyPrincipalAttachmentProperties {
        PolicyName: pulumi.Input<string>;
        Principal: pulumi.Input<string>;
    }
    
    export interface PolicyProperties {
        PolicyDocument: pulumi.Input<string>;
        PolicyName?: pulumi.Input<string>;
    }
    
    export interface ThingAttributePayloadAttributes {
        Attributes?: {[key: string]: string};
    }
    
    export interface ThingAttributePayloadProperties {
        Attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    }
    
    export interface ThingAttributes {
    }
    
    export interface ThingPrincipalAttachmentAttributes {
    }
    
    export interface ThingPrincipalAttachmentProperties {
        Principal: pulumi.Input<string>;
        ThingName: pulumi.Input<string>;
    }
    
    export interface ThingProperties {
        AttributePayload?: pulumi.Input<ThingAttributePayloadProperties>;
        ThingName?: pulumi.Input<string>;
    }
    
    export interface TopicRuleActionAttributes {
        CloudwatchAlarm?: TopicRuleCloudwatchAlarmActionAttributes;
        CloudwatchMetric?: TopicRuleCloudwatchMetricActionAttributes;
        DynamoDB?: TopicRuleDynamoDBActionAttributes;
        DynamoDBv2?: TopicRuleDynamoDBv2ActionAttributes;
        Elasticsearch?: TopicRuleElasticsearchActionAttributes;
        Firehose?: TopicRuleFirehoseActionAttributes;
        IotAnalytics?: TopicRuleIotAnalyticsActionAttributes;
        Kinesis?: TopicRuleKinesisActionAttributes;
        Lambda?: TopicRuleLambdaActionAttributes;
        Republish?: TopicRuleRepublishActionAttributes;
        S3?: TopicRuleS3ActionAttributes;
        Sns?: TopicRuleSnsActionAttributes;
        Sqs?: TopicRuleSqsActionAttributes;
        StepFunctions?: TopicRuleStepFunctionsActionAttributes;
    }
    
    export interface TopicRuleActionProperties {
        CloudwatchAlarm?: pulumi.Input<TopicRuleCloudwatchAlarmActionProperties>;
        CloudwatchMetric?: pulumi.Input<TopicRuleCloudwatchMetricActionProperties>;
        DynamoDB?: pulumi.Input<TopicRuleDynamoDBActionProperties>;
        DynamoDBv2?: pulumi.Input<TopicRuleDynamoDBv2ActionProperties>;
        Elasticsearch?: pulumi.Input<TopicRuleElasticsearchActionProperties>;
        Firehose?: pulumi.Input<TopicRuleFirehoseActionProperties>;
        IotAnalytics?: pulumi.Input<TopicRuleIotAnalyticsActionProperties>;
        Kinesis?: pulumi.Input<TopicRuleKinesisActionProperties>;
        Lambda?: pulumi.Input<TopicRuleLambdaActionProperties>;
        Republish?: pulumi.Input<TopicRuleRepublishActionProperties>;
        S3?: pulumi.Input<TopicRuleS3ActionProperties>;
        Sns?: pulumi.Input<TopicRuleSnsActionProperties>;
        Sqs?: pulumi.Input<TopicRuleSqsActionProperties>;
        StepFunctions?: pulumi.Input<TopicRuleStepFunctionsActionProperties>;
    }
    
    export interface TopicRuleAttributes {
        Arn: string;
    }
    
    export interface TopicRuleCloudwatchAlarmActionAttributes {
        AlarmName: string;
        RoleArn: string;
        StateReason: string;
        StateValue: string;
    }
    
    export interface TopicRuleCloudwatchAlarmActionProperties {
        AlarmName: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
        StateReason: pulumi.Input<string>;
        StateValue: pulumi.Input<string>;
    }
    
    export interface TopicRuleCloudwatchMetricActionAttributes {
        MetricName: string;
        MetricNamespace: string;
        MetricTimestamp?: string;
        MetricUnit: string;
        MetricValue: string;
        RoleArn: string;
    }
    
    export interface TopicRuleCloudwatchMetricActionProperties {
        MetricName: pulumi.Input<string>;
        MetricNamespace: pulumi.Input<string>;
        MetricTimestamp?: pulumi.Input<string>;
        MetricUnit: pulumi.Input<string>;
        MetricValue: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
    }
    
    export interface TopicRuleDynamoDBActionAttributes {
        HashKeyField: string;
        HashKeyType?: string;
        HashKeyValue: string;
        PayloadField?: string;
        RangeKeyField?: string;
        RangeKeyType?: string;
        RangeKeyValue?: string;
        RoleArn: string;
        TableName: string;
    }
    
    export interface TopicRuleDynamoDBActionProperties {
        HashKeyField: pulumi.Input<string>;
        HashKeyType?: pulumi.Input<string>;
        HashKeyValue: pulumi.Input<string>;
        PayloadField?: pulumi.Input<string>;
        RangeKeyField?: pulumi.Input<string>;
        RangeKeyType?: pulumi.Input<string>;
        RangeKeyValue?: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
        TableName: pulumi.Input<string>;
    }
    
    export interface TopicRuleDynamoDBv2ActionAttributes {
        PutItem?: TopicRulePutItemInputAttributes;
        RoleArn?: string;
    }
    
    export interface TopicRuleDynamoDBv2ActionProperties {
        PutItem?: pulumi.Input<TopicRulePutItemInputProperties>;
        RoleArn?: pulumi.Input<string>;
    }
    
    export interface TopicRuleElasticsearchActionAttributes {
        Endpoint: string;
        Id: string;
        Index: string;
        RoleArn: string;
        Type: string;
    }
    
    export interface TopicRuleElasticsearchActionProperties {
        Endpoint: pulumi.Input<string>;
        Id: pulumi.Input<string>;
        Index: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface TopicRuleFirehoseActionAttributes {
        DeliveryStreamName: string;
        RoleArn: string;
        Separator?: string;
    }
    
    export interface TopicRuleFirehoseActionProperties {
        DeliveryStreamName: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
        Separator?: pulumi.Input<string>;
    }
    
    export interface TopicRuleIotAnalyticsActionAttributes {
        ChannelName: string;
        RoleArn: string;
    }
    
    export interface TopicRuleIotAnalyticsActionProperties {
        ChannelName: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
    }
    
    export interface TopicRuleKinesisActionAttributes {
        PartitionKey?: string;
        RoleArn: string;
        StreamName: string;
    }
    
    export interface TopicRuleKinesisActionProperties {
        PartitionKey?: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
        StreamName: pulumi.Input<string>;
    }
    
    export interface TopicRuleLambdaActionAttributes {
        FunctionArn?: string;
    }
    
    export interface TopicRuleLambdaActionProperties {
        FunctionArn?: pulumi.Input<string>;
    }
    
    export interface TopicRuleProperties {
        RuleName?: pulumi.Input<string>;
        TopicRulePayload: pulumi.Input<TopicRuleTopicRulePayloadProperties>;
    }
    
    export interface TopicRulePutItemInputAttributes {
        TableName: string;
    }
    
    export interface TopicRulePutItemInputProperties {
        TableName: pulumi.Input<string>;
    }
    
    export interface TopicRuleRepublishActionAttributes {
        RoleArn: string;
        Topic: string;
    }
    
    export interface TopicRuleRepublishActionProperties {
        RoleArn: pulumi.Input<string>;
        Topic: pulumi.Input<string>;
    }
    
    export interface TopicRuleS3ActionAttributes {
        BucketName: string;
        Key: string;
        RoleArn: string;
    }
    
    export interface TopicRuleS3ActionProperties {
        BucketName: pulumi.Input<string>;
        Key: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
    }
    
    export interface TopicRuleSnsActionAttributes {
        MessageFormat?: string;
        RoleArn: string;
        TargetArn: string;
    }
    
    export interface TopicRuleSnsActionProperties {
        MessageFormat?: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
        TargetArn: pulumi.Input<string>;
    }
    
    export interface TopicRuleSqsActionAttributes {
        QueueUrl: string;
        RoleArn: string;
        UseBase64?: boolean;
    }
    
    export interface TopicRuleSqsActionProperties {
        QueueUrl: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
        UseBase64?: pulumi.Input<boolean>;
    }
    
    export interface TopicRuleStepFunctionsActionAttributes {
        ExecutionNamePrefix?: string;
        RoleArn: string;
        StateMachineName: string;
    }
    
    export interface TopicRuleStepFunctionsActionProperties {
        ExecutionNamePrefix?: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
        StateMachineName: pulumi.Input<string>;
    }
    
    export interface TopicRuleTopicRulePayloadAttributes {
        Actions: TopicRuleActionAttributes[];
        AwsIotSqlVersion?: string;
        Description?: string;
        ErrorAction?: TopicRuleActionAttributes;
        RuleDisabled: boolean;
        Sql: string;
    }
    
    export interface TopicRuleTopicRulePayloadProperties {
        Actions: pulumi.Input<pulumi.Input<TopicRuleActionProperties>[]>;
        AwsIotSqlVersion?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        ErrorAction?: pulumi.Input<TopicRuleActionProperties>;
        RuleDisabled: pulumi.Input<boolean>;
        Sql: pulumi.Input<string>;
    }
    
    
    export class Certificate extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<CertificateAttributes>;

        constructor(name: string, properties: CertificateProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IoT:Certificate", name, inputs, opts);
        }
    }
    
    export class Policy extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PolicyAttributes>;

        constructor(name: string, properties: PolicyProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IoT:Policy", name, inputs, opts);
        }
    }
    
    export class PolicyPrincipalAttachment extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PolicyPrincipalAttachmentAttributes>;

        constructor(name: string, properties: PolicyPrincipalAttachmentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IoT:PolicyPrincipalAttachment", name, inputs, opts);
        }
    }
    
    export class Thing extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ThingAttributes>;

        constructor(name: string, properties?: ThingProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IoT:Thing", name, inputs, opts);
        }
    }
    
    export class ThingPrincipalAttachment extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ThingPrincipalAttachmentAttributes>;

        constructor(name: string, properties: ThingPrincipalAttachmentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IoT:ThingPrincipalAttachment", name, inputs, opts);
        }
    }
    
    export class TopicRule extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TopicRuleAttributes>;

        constructor(name: string, properties: TopicRuleProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IoT:TopicRule", name, inputs, opts);
        }
    }
    
}

export namespace iot1click {
    
    export interface DeviceAttributes {
        Arn: string;
        DeviceId: string;
        Enabled: boolean;
    }
    
    export interface DeviceProperties {
        DeviceId: pulumi.Input<string>;
        Enabled: pulumi.Input<boolean>;
    }
    
    export interface PlacementAttributes {
        PlacementName: string;
        ProjectName: string;
    }
    
    export interface PlacementProperties {
        AssociatedDevices?: pulumi.Input<string>;
        Attributes?: pulumi.Input<string>;
        PlacementName?: pulumi.Input<string>;
        ProjectName: pulumi.Input<string>;
    }
    
    export interface ProjectAttributes {
        Arn: string;
        ProjectName: string;
    }
    
    export interface ProjectDeviceTemplateAttributes {
        CallbackOverrides?: string;
        DeviceType?: string;
    }
    
    export interface ProjectDeviceTemplateProperties {
        CallbackOverrides?: pulumi.Input<string>;
        DeviceType?: pulumi.Input<string>;
    }
    
    export interface ProjectPlacementTemplateAttributes {
        DefaultAttributes?: string;
        DeviceTemplates?: string;
    }
    
    export interface ProjectPlacementTemplateProperties {
        DefaultAttributes?: pulumi.Input<string>;
        DeviceTemplates?: pulumi.Input<string>;
    }
    
    export interface ProjectProperties {
        Description?: pulumi.Input<string>;
        PlacementTemplate: pulumi.Input<ProjectPlacementTemplateProperties>;
        ProjectName?: pulumi.Input<string>;
    }
    
    
    export class Device extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DeviceAttributes>;

        constructor(name: string, properties: DeviceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IoT1Click:Device", name, inputs, opts);
        }
    }
    
    export class Placement extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PlacementAttributes>;

        constructor(name: string, properties: PlacementProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IoT1Click:Placement", name, inputs, opts);
        }
    }
    
    export class Project extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ProjectAttributes>;

        constructor(name: string, properties: ProjectProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IoT1Click:Project", name, inputs, opts);
        }
    }
    
}

export namespace iotanalytics {
    
    export interface ChannelAttributes {
    }
    
    export interface ChannelChannelStorageAttributes {
        CustomerManagedS3?: ChannelCustomerManagedS3Attributes;
        ServiceManagedS3?: ChannelServiceManagedS3Attributes;
    }
    
    export interface ChannelChannelStorageProperties {
        CustomerManagedS3?: pulumi.Input<ChannelCustomerManagedS3Properties>;
        ServiceManagedS3?: pulumi.Input<ChannelServiceManagedS3Properties>;
    }
    
    export interface ChannelCustomerManagedS3Attributes {
        Bucket: string;
        KeyPrefix?: string;
        RoleArn: string;
    }
    
    export interface ChannelCustomerManagedS3Properties {
        Bucket: pulumi.Input<string>;
        KeyPrefix?: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
    }
    
    export interface ChannelProperties {
        ChannelName?: pulumi.Input<string>;
        ChannelStorage?: pulumi.Input<ChannelChannelStorageProperties>;
        RetentionPeriod?: pulumi.Input<ChannelRetentionPeriodProperties>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface ChannelRetentionPeriodAttributes {
        NumberOfDays?: number;
        Unlimited?: boolean;
    }
    
    export interface ChannelRetentionPeriodProperties {
        NumberOfDays?: pulumi.Input<number>;
        Unlimited?: pulumi.Input<boolean>;
    }
    
    export interface ChannelServiceManagedS3Attributes {
    }
    
    export interface ChannelServiceManagedS3Properties {
    }
    
    export interface DatasetActionAttributes {
        ActionName: string;
        ContainerAction?: DatasetContainerActionAttributes;
        QueryAction?: DatasetQueryActionAttributes;
    }
    
    export interface DatasetActionProperties {
        ActionName: pulumi.Input<string>;
        ContainerAction?: pulumi.Input<DatasetContainerActionProperties>;
        QueryAction?: pulumi.Input<DatasetQueryActionProperties>;
    }
    
    export interface DatasetAttributes {
    }
    
    export interface DatasetContainerActionAttributes {
        ExecutionRoleArn: string;
        Image: string;
        ResourceConfiguration: DatasetResourceConfigurationAttributes;
        Variables?: DatasetVariableAttributes[];
    }
    
    export interface DatasetContainerActionProperties {
        ExecutionRoleArn: pulumi.Input<string>;
        Image: pulumi.Input<string>;
        ResourceConfiguration: pulumi.Input<DatasetResourceConfigurationProperties>;
        Variables?: pulumi.Input<pulumi.Input<DatasetVariableProperties>[]>;
    }
    
    export interface DatasetDatasetContentDeliveryRuleAttributes {
        Destination: DatasetDatasetContentDeliveryRuleDestinationAttributes;
        EntryName?: string;
    }
    
    export interface DatasetDatasetContentDeliveryRuleDestinationAttributes {
        IotEventsDestinationConfiguration?: DatasetIotEventsDestinationConfigurationAttributes;
        S3DestinationConfiguration?: DatasetS3DestinationConfigurationAttributes;
    }
    
    export interface DatasetDatasetContentDeliveryRuleDestinationProperties {
        IotEventsDestinationConfiguration?: pulumi.Input<DatasetIotEventsDestinationConfigurationProperties>;
        S3DestinationConfiguration?: pulumi.Input<DatasetS3DestinationConfigurationProperties>;
    }
    
    export interface DatasetDatasetContentDeliveryRuleProperties {
        Destination: pulumi.Input<DatasetDatasetContentDeliveryRuleDestinationProperties>;
        EntryName?: pulumi.Input<string>;
    }
    
    export interface DatasetDatasetContentVersionValueAttributes {
        DatasetName?: string;
    }
    
    export interface DatasetDatasetContentVersionValueProperties {
        DatasetName?: pulumi.Input<string>;
    }
    
    export interface DatasetDeltaTimeAttributes {
        OffsetSeconds: number;
        TimeExpression: string;
    }
    
    export interface DatasetDeltaTimeProperties {
        OffsetSeconds: pulumi.Input<number>;
        TimeExpression: pulumi.Input<string>;
    }
    
    export interface DatasetFilterAttributes {
        DeltaTime?: DatasetDeltaTimeAttributes;
    }
    
    export interface DatasetFilterProperties {
        DeltaTime?: pulumi.Input<DatasetDeltaTimeProperties>;
    }
    
    export interface DatasetGlueConfigurationAttributes {
        DatabaseName: string;
        TableName: string;
    }
    
    export interface DatasetGlueConfigurationProperties {
        DatabaseName: pulumi.Input<string>;
        TableName: pulumi.Input<string>;
    }
    
    export interface DatasetIotEventsDestinationConfigurationAttributes {
        InputName: string;
        RoleArn: string;
    }
    
    export interface DatasetIotEventsDestinationConfigurationProperties {
        InputName: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
    }
    
    export interface DatasetOutputFileUriValueAttributes {
        FileName?: string;
    }
    
    export interface DatasetOutputFileUriValueProperties {
        FileName?: pulumi.Input<string>;
    }
    
    export interface DatasetProperties {
        Actions: pulumi.Input<pulumi.Input<DatasetActionProperties>[]>;
        ContentDeliveryRules?: pulumi.Input<pulumi.Input<DatasetDatasetContentDeliveryRuleProperties>[]>;
        DatasetName?: pulumi.Input<string>;
        RetentionPeriod?: pulumi.Input<DatasetRetentionPeriodProperties>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        Triggers?: pulumi.Input<pulumi.Input<DatasetTriggerProperties>[]>;
        VersioningConfiguration?: pulumi.Input<DatasetVersioningConfigurationProperties>;
    }
    
    export interface DatasetQueryActionAttributes {
        Filters?: DatasetFilterAttributes[];
        SqlQuery: string;
    }
    
    export interface DatasetQueryActionProperties {
        Filters?: pulumi.Input<pulumi.Input<DatasetFilterProperties>[]>;
        SqlQuery: pulumi.Input<string>;
    }
    
    export interface DatasetResourceConfigurationAttributes {
        ComputeType: string;
        VolumeSizeInGB: number;
    }
    
    export interface DatasetResourceConfigurationProperties {
        ComputeType: pulumi.Input<string>;
        VolumeSizeInGB: pulumi.Input<number>;
    }
    
    export interface DatasetRetentionPeriodAttributes {
        NumberOfDays: number;
        Unlimited: boolean;
    }
    
    export interface DatasetRetentionPeriodProperties {
        NumberOfDays: pulumi.Input<number>;
        Unlimited: pulumi.Input<boolean>;
    }
    
    export interface DatasetS3DestinationConfigurationAttributes {
        Bucket: string;
        GlueConfiguration?: DatasetGlueConfigurationAttributes;
        Key: string;
        RoleArn: string;
    }
    
    export interface DatasetS3DestinationConfigurationProperties {
        Bucket: pulumi.Input<string>;
        GlueConfiguration?: pulumi.Input<DatasetGlueConfigurationProperties>;
        Key: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
    }
    
    export interface DatasetScheduleAttributes {
        ScheduleExpression: string;
    }
    
    export interface DatasetScheduleProperties {
        ScheduleExpression: pulumi.Input<string>;
    }
    
    export interface DatasetTriggerAttributes {
        Schedule?: DatasetScheduleAttributes;
        TriggeringDataset?: DatasetTriggeringDatasetAttributes;
    }
    
    export interface DatasetTriggerProperties {
        Schedule?: pulumi.Input<DatasetScheduleProperties>;
        TriggeringDataset?: pulumi.Input<DatasetTriggeringDatasetProperties>;
    }
    
    export interface DatasetTriggeringDatasetAttributes {
        DatasetName: string;
    }
    
    export interface DatasetTriggeringDatasetProperties {
        DatasetName: pulumi.Input<string>;
    }
    
    export interface DatasetVariableAttributes {
        DatasetContentVersionValue?: DatasetDatasetContentVersionValueAttributes;
        DoubleValue?: number;
        OutputFileUriValue?: DatasetOutputFileUriValueAttributes;
        StringValue?: string;
        VariableName: string;
    }
    
    export interface DatasetVariableProperties {
        DatasetContentVersionValue?: pulumi.Input<DatasetDatasetContentVersionValueProperties>;
        DoubleValue?: pulumi.Input<number>;
        OutputFileUriValue?: pulumi.Input<DatasetOutputFileUriValueProperties>;
        StringValue?: pulumi.Input<string>;
        VariableName: pulumi.Input<string>;
    }
    
    export interface DatasetVersioningConfigurationAttributes {
        MaxVersions?: number;
        Unlimited?: boolean;
    }
    
    export interface DatasetVersioningConfigurationProperties {
        MaxVersions?: pulumi.Input<number>;
        Unlimited?: pulumi.Input<boolean>;
    }
    
    export interface DatastoreAttributes {
    }
    
    export interface DatastoreCustomerManagedS3Attributes {
        Bucket: string;
        KeyPrefix?: string;
        RoleArn: string;
    }
    
    export interface DatastoreCustomerManagedS3Properties {
        Bucket: pulumi.Input<string>;
        KeyPrefix?: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
    }
    
    export interface DatastoreDatastoreStorageAttributes {
        CustomerManagedS3?: DatastoreCustomerManagedS3Attributes;
        ServiceManagedS3?: DatastoreServiceManagedS3Attributes;
    }
    
    export interface DatastoreDatastoreStorageProperties {
        CustomerManagedS3?: pulumi.Input<DatastoreCustomerManagedS3Properties>;
        ServiceManagedS3?: pulumi.Input<DatastoreServiceManagedS3Properties>;
    }
    
    export interface DatastoreProperties {
        DatastoreName?: pulumi.Input<string>;
        DatastoreStorage?: pulumi.Input<DatastoreDatastoreStorageProperties>;
        RetentionPeriod?: pulumi.Input<DatastoreRetentionPeriodProperties>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface DatastoreRetentionPeriodAttributes {
        NumberOfDays?: number;
        Unlimited?: boolean;
    }
    
    export interface DatastoreRetentionPeriodProperties {
        NumberOfDays?: pulumi.Input<number>;
        Unlimited?: pulumi.Input<boolean>;
    }
    
    export interface DatastoreServiceManagedS3Attributes {
    }
    
    export interface DatastoreServiceManagedS3Properties {
    }
    
    export interface PipelineActivityAttributes {
        AddAttributes?: PipelineAddAttributesAttributes;
        Channel?: PipelineChannelAttributes;
        Datastore?: PipelineDatastoreAttributes;
        DeviceRegistryEnrich?: PipelineDeviceRegistryEnrichAttributes;
        DeviceShadowEnrich?: PipelineDeviceShadowEnrichAttributes;
        Filter?: PipelineFilterAttributes;
        Lambda?: PipelineLambdaAttributes;
        Math?: PipelineMathAttributes;
        RemoveAttributes?: PipelineRemoveAttributesAttributes;
        SelectAttributes?: PipelineSelectAttributesAttributes;
    }
    
    export interface PipelineActivityProperties {
        AddAttributes?: pulumi.Input<PipelineAddAttributesProperties>;
        Channel?: pulumi.Input<PipelineChannelProperties>;
        Datastore?: pulumi.Input<PipelineDatastoreProperties>;
        DeviceRegistryEnrich?: pulumi.Input<PipelineDeviceRegistryEnrichProperties>;
        DeviceShadowEnrich?: pulumi.Input<PipelineDeviceShadowEnrichProperties>;
        Filter?: pulumi.Input<PipelineFilterProperties>;
        Lambda?: pulumi.Input<PipelineLambdaProperties>;
        Math?: pulumi.Input<PipelineMathProperties>;
        RemoveAttributes?: pulumi.Input<PipelineRemoveAttributesProperties>;
        SelectAttributes?: pulumi.Input<PipelineSelectAttributesProperties>;
    }
    
    export interface PipelineAddAttributesAttributes {
        Attributes?: string;
        Name?: string;
        Next?: string;
    }
    
    export interface PipelineAddAttributesProperties {
        Attributes?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Next?: pulumi.Input<string>;
    }
    
    export interface PipelineAttributes {
    }
    
    export interface PipelineChannelAttributes {
        ChannelName?: string;
        Name?: string;
        Next?: string;
    }
    
    export interface PipelineChannelProperties {
        ChannelName?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Next?: pulumi.Input<string>;
    }
    
    export interface PipelineDatastoreAttributes {
        DatastoreName?: string;
        Name?: string;
    }
    
    export interface PipelineDatastoreProperties {
        DatastoreName?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
    }
    
    export interface PipelineDeviceRegistryEnrichAttributes {
        Attribute?: string;
        Name?: string;
        Next?: string;
        RoleArn?: string;
        ThingName?: string;
    }
    
    export interface PipelineDeviceRegistryEnrichProperties {
        Attribute?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Next?: pulumi.Input<string>;
        RoleArn?: pulumi.Input<string>;
        ThingName?: pulumi.Input<string>;
    }
    
    export interface PipelineDeviceShadowEnrichAttributes {
        Attribute?: string;
        Name?: string;
        Next?: string;
        RoleArn?: string;
        ThingName?: string;
    }
    
    export interface PipelineDeviceShadowEnrichProperties {
        Attribute?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Next?: pulumi.Input<string>;
        RoleArn?: pulumi.Input<string>;
        ThingName?: pulumi.Input<string>;
    }
    
    export interface PipelineFilterAttributes {
        Filter?: string;
        Name?: string;
        Next?: string;
    }
    
    export interface PipelineFilterProperties {
        Filter?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Next?: pulumi.Input<string>;
    }
    
    export interface PipelineLambdaAttributes {
        BatchSize?: number;
        LambdaName?: string;
        Name?: string;
        Next?: string;
    }
    
    export interface PipelineLambdaProperties {
        BatchSize?: pulumi.Input<number>;
        LambdaName?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Next?: pulumi.Input<string>;
    }
    
    export interface PipelineMathAttributes {
        Attribute?: string;
        Math?: string;
        Name?: string;
        Next?: string;
    }
    
    export interface PipelineMathProperties {
        Attribute?: pulumi.Input<string>;
        Math?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Next?: pulumi.Input<string>;
    }
    
    export interface PipelineProperties {
        PipelineActivities: pulumi.Input<pulumi.Input<PipelineActivityProperties>[]>;
        PipelineName?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface PipelineRemoveAttributesAttributes {
        Attributes?: string[];
        Name?: string;
        Next?: string;
    }
    
    export interface PipelineRemoveAttributesProperties {
        Attributes?: pulumi.Input<pulumi.Input<string>[]>;
        Name?: pulumi.Input<string>;
        Next?: pulumi.Input<string>;
    }
    
    export interface PipelineSelectAttributesAttributes {
        Attributes?: string[];
        Name?: string;
        Next?: string;
    }
    
    export interface PipelineSelectAttributesProperties {
        Attributes?: pulumi.Input<pulumi.Input<string>[]>;
        Name?: pulumi.Input<string>;
        Next?: pulumi.Input<string>;
    }
    
    
    export class Channel extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ChannelAttributes>;

        constructor(name: string, properties?: ChannelProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IoTAnalytics:Channel", name, inputs, opts);
        }
    }
    
    export class Dataset extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DatasetAttributes>;

        constructor(name: string, properties: DatasetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IoTAnalytics:Dataset", name, inputs, opts);
        }
    }
    
    export class Datastore extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DatastoreAttributes>;

        constructor(name: string, properties?: DatastoreProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IoTAnalytics:Datastore", name, inputs, opts);
        }
    }
    
    export class Pipeline extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PipelineAttributes>;

        constructor(name: string, properties: PipelineProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IoTAnalytics:Pipeline", name, inputs, opts);
        }
    }
    
}

export namespace iotevents {
    
    export interface DetectorModelActionAttributes {
        ClearTimer?: DetectorModelClearTimerAttributes;
        Firehose?: DetectorModelFirehoseAttributes;
        IotEvents?: DetectorModelIotEventsAttributes;
        IotTopicPublish?: DetectorModelIotTopicPublishAttributes;
        Lambda?: DetectorModelLambdaAttributes;
        ResetTimer?: DetectorModelResetTimerAttributes;
        SetTimer?: DetectorModelSetTimerAttributes;
        SetVariable?: DetectorModelSetVariableAttributes;
        Sns?: DetectorModelSnsAttributes;
        Sqs?: DetectorModelSqsAttributes;
    }
    
    export interface DetectorModelActionProperties {
        ClearTimer?: pulumi.Input<DetectorModelClearTimerProperties>;
        Firehose?: pulumi.Input<DetectorModelFirehoseProperties>;
        IotEvents?: pulumi.Input<DetectorModelIotEventsProperties>;
        IotTopicPublish?: pulumi.Input<DetectorModelIotTopicPublishProperties>;
        Lambda?: pulumi.Input<DetectorModelLambdaProperties>;
        ResetTimer?: pulumi.Input<DetectorModelResetTimerProperties>;
        SetTimer?: pulumi.Input<DetectorModelSetTimerProperties>;
        SetVariable?: pulumi.Input<DetectorModelSetVariableProperties>;
        Sns?: pulumi.Input<DetectorModelSnsProperties>;
        Sqs?: pulumi.Input<DetectorModelSqsProperties>;
    }
    
    export interface DetectorModelAttributes {
    }
    
    export interface DetectorModelClearTimerAttributes {
        TimerName?: string;
    }
    
    export interface DetectorModelClearTimerProperties {
        TimerName?: pulumi.Input<string>;
    }
    
    export interface DetectorModelDetectorModelDefinitionAttributes {
        InitialStateName?: string;
        States?: DetectorModelStateAttributes[];
    }
    
    export interface DetectorModelDetectorModelDefinitionProperties {
        InitialStateName?: pulumi.Input<string>;
        States?: pulumi.Input<pulumi.Input<DetectorModelStateProperties>[]>;
    }
    
    export interface DetectorModelEventAttributes {
        Actions?: DetectorModelActionAttributes[];
        Condition?: string;
        EventName?: string;
    }
    
    export interface DetectorModelEventProperties {
        Actions?: pulumi.Input<pulumi.Input<DetectorModelActionProperties>[]>;
        Condition?: pulumi.Input<string>;
        EventName?: pulumi.Input<string>;
    }
    
    export interface DetectorModelFirehoseAttributes {
        DeliveryStreamName?: string;
        Separator?: string;
    }
    
    export interface DetectorModelFirehoseProperties {
        DeliveryStreamName?: pulumi.Input<string>;
        Separator?: pulumi.Input<string>;
    }
    
    export interface DetectorModelIotEventsAttributes {
        InputName?: string;
    }
    
    export interface DetectorModelIotEventsProperties {
        InputName?: pulumi.Input<string>;
    }
    
    export interface DetectorModelIotTopicPublishAttributes {
        MqttTopic?: string;
    }
    
    export interface DetectorModelIotTopicPublishProperties {
        MqttTopic?: pulumi.Input<string>;
    }
    
    export interface DetectorModelLambdaAttributes {
        FunctionArn?: string;
    }
    
    export interface DetectorModelLambdaProperties {
        FunctionArn?: pulumi.Input<string>;
    }
    
    export interface DetectorModelOnEnterAttributes {
        Events?: DetectorModelEventAttributes[];
    }
    
    export interface DetectorModelOnEnterProperties {
        Events?: pulumi.Input<pulumi.Input<DetectorModelEventProperties>[]>;
    }
    
    export interface DetectorModelOnExitAttributes {
        Events?: DetectorModelEventAttributes[];
    }
    
    export interface DetectorModelOnExitProperties {
        Events?: pulumi.Input<pulumi.Input<DetectorModelEventProperties>[]>;
    }
    
    export interface DetectorModelOnInputAttributes {
        Events?: DetectorModelEventAttributes[];
        TransitionEvents?: DetectorModelTransitionEventAttributes[];
    }
    
    export interface DetectorModelOnInputProperties {
        Events?: pulumi.Input<pulumi.Input<DetectorModelEventProperties>[]>;
        TransitionEvents?: pulumi.Input<pulumi.Input<DetectorModelTransitionEventProperties>[]>;
    }
    
    export interface DetectorModelProperties {
        DetectorModelDefinition?: pulumi.Input<DetectorModelDetectorModelDefinitionProperties>;
        DetectorModelDescription?: pulumi.Input<string>;
        DetectorModelName?: pulumi.Input<string>;
        Key?: pulumi.Input<string>;
        RoleArn?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface DetectorModelResetTimerAttributes {
        TimerName?: string;
    }
    
    export interface DetectorModelResetTimerProperties {
        TimerName?: pulumi.Input<string>;
    }
    
    export interface DetectorModelSetTimerAttributes {
        Seconds?: number;
        TimerName?: string;
    }
    
    export interface DetectorModelSetTimerProperties {
        Seconds?: pulumi.Input<number>;
        TimerName?: pulumi.Input<string>;
    }
    
    export interface DetectorModelSetVariableAttributes {
        Value?: string;
        VariableName?: string;
    }
    
    export interface DetectorModelSetVariableProperties {
        Value?: pulumi.Input<string>;
        VariableName?: pulumi.Input<string>;
    }
    
    export interface DetectorModelSnsAttributes {
        TargetArn?: string;
    }
    
    export interface DetectorModelSnsProperties {
        TargetArn?: pulumi.Input<string>;
    }
    
    export interface DetectorModelSqsAttributes {
        QueueUrl?: string;
        UseBase64?: boolean;
    }
    
    export interface DetectorModelSqsProperties {
        QueueUrl?: pulumi.Input<string>;
        UseBase64?: pulumi.Input<boolean>;
    }
    
    export interface DetectorModelStateAttributes {
        OnEnter?: DetectorModelOnEnterAttributes;
        OnExit?: DetectorModelOnExitAttributes;
        OnInput?: DetectorModelOnInputAttributes;
        StateName?: string;
    }
    
    export interface DetectorModelStateProperties {
        OnEnter?: pulumi.Input<DetectorModelOnEnterProperties>;
        OnExit?: pulumi.Input<DetectorModelOnExitProperties>;
        OnInput?: pulumi.Input<DetectorModelOnInputProperties>;
        StateName?: pulumi.Input<string>;
    }
    
    export interface DetectorModelTransitionEventAttributes {
        Actions?: DetectorModelActionAttributes[];
        Condition?: string;
        EventName?: string;
        NextState?: string;
    }
    
    export interface DetectorModelTransitionEventProperties {
        Actions?: pulumi.Input<pulumi.Input<DetectorModelActionProperties>[]>;
        Condition?: pulumi.Input<string>;
        EventName?: pulumi.Input<string>;
        NextState?: pulumi.Input<string>;
    }
    
    export interface InputAttributeAttributes {
        JsonPath?: string;
    }
    
    export interface InputAttributeProperties {
        JsonPath?: pulumi.Input<string>;
    }
    
    export interface InputAttributes {
    }
    
    export interface InputInputDefinitionAttributes {
        Attributes?: InputAttributeAttributes[];
    }
    
    export interface InputInputDefinitionProperties {
        Attributes?: pulumi.Input<pulumi.Input<InputAttributeProperties>[]>;
    }
    
    export interface InputProperties {
        InputDefinition?: pulumi.Input<InputInputDefinitionProperties>;
        InputDescription?: pulumi.Input<string>;
        InputName?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    
    export class DetectorModel extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DetectorModelAttributes>;

        constructor(name: string, properties?: DetectorModelProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IoTEvents:DetectorModel", name, inputs, opts);
        }
    }
    
    export class Input extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<InputAttributes>;

        constructor(name: string, properties?: InputProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IoTEvents:Input", name, inputs, opts);
        }
    }
    
}

export namespace iotthingsgraph {
    
    export interface FlowTemplateAttributes {
    }
    
    export interface FlowTemplateDefinitionDocumentAttributes {
        Language: string;
        Text: string;
    }
    
    export interface FlowTemplateDefinitionDocumentProperties {
        Language: pulumi.Input<string>;
        Text: pulumi.Input<string>;
    }
    
    export interface FlowTemplateProperties {
        CompatibleNamespaceVersion?: pulumi.Input<number>;
        Definition: pulumi.Input<FlowTemplateDefinitionDocumentProperties>;
    }
    
    
    export class FlowTemplate extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<FlowTemplateAttributes>;

        constructor(name: string, properties: FlowTemplateProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:IoTThingsGraph:FlowTemplate", name, inputs, opts);
        }
    }
    
}

export namespace kms {
    
    export interface AliasAttributes {
    }
    
    export interface AliasProperties {
        AliasName: pulumi.Input<string>;
        TargetKeyId: pulumi.Input<string>;
    }
    
    export interface KeyAttributes {
        Arn: string;
    }
    
    export interface KeyProperties {
        Description?: pulumi.Input<string>;
        EnableKeyRotation?: pulumi.Input<boolean>;
        Enabled?: pulumi.Input<boolean>;
        KeyPolicy: pulumi.Input<string>;
        KeyUsage?: pulumi.Input<string>;
        PendingWindowInDays?: pulumi.Input<number>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    
    export class Alias extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AliasAttributes>;

        constructor(name: string, properties: AliasProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:KMS:Alias", name, inputs, opts);
        }
    }
    
    export class Key extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<KeyAttributes>;

        constructor(name: string, properties: KeyProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:KMS:Key", name, inputs, opts);
        }
    }
    
}

export namespace kinesis {
    
    export interface StreamAttributes {
        Arn: string;
    }
    
    export interface StreamConsumerAttributes {
        ConsumerARN: string;
        ConsumerCreationTimestamp: string;
        ConsumerName: string;
        ConsumerStatus: string;
        StreamARN: string;
    }
    
    export interface StreamConsumerProperties {
        ConsumerName: pulumi.Input<string>;
        StreamARN: pulumi.Input<string>;
    }
    
    export interface StreamProperties {
        Name?: pulumi.Input<string>;
        RetentionPeriodHours?: pulumi.Input<number>;
        ShardCount: pulumi.Input<number>;
        StreamEncryption?: pulumi.Input<StreamStreamEncryptionProperties>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface StreamStreamEncryptionAttributes {
        EncryptionType: string;
        KeyId: string;
    }
    
    export interface StreamStreamEncryptionProperties {
        EncryptionType: pulumi.Input<string>;
        KeyId: pulumi.Input<string>;
    }
    
    
    export class Stream extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<StreamAttributes>;

        constructor(name: string, properties: StreamProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Kinesis:Stream", name, inputs, opts);
        }
    }
    
    export class StreamConsumer extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<StreamConsumerAttributes>;

        constructor(name: string, properties: StreamConsumerProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Kinesis:StreamConsumer", name, inputs, opts);
        }
    }
    
}

export namespace kinesisanalytics {
    
    export interface ApplicationAttributes {
    }
    
    export interface ApplicationCSVMappingParametersAttributes {
        RecordColumnDelimiter: string;
        RecordRowDelimiter: string;
    }
    
    export interface ApplicationCSVMappingParametersProperties {
        RecordColumnDelimiter: pulumi.Input<string>;
        RecordRowDelimiter: pulumi.Input<string>;
    }
    
    export interface ApplicationInputAttributes {
        InputParallelism?: ApplicationInputParallelismAttributes;
        InputProcessingConfiguration?: ApplicationInputProcessingConfigurationAttributes;
        InputSchema: ApplicationInputSchemaAttributes;
        KinesisFirehoseInput?: ApplicationKinesisFirehoseInputAttributes;
        KinesisStreamsInput?: ApplicationKinesisStreamsInputAttributes;
        NamePrefix: string;
    }
    
    export interface ApplicationInputLambdaProcessorAttributes {
        ResourceARN: string;
        RoleARN: string;
    }
    
    export interface ApplicationInputLambdaProcessorProperties {
        ResourceARN: pulumi.Input<string>;
        RoleARN: pulumi.Input<string>;
    }
    
    export interface ApplicationInputParallelismAttributes {
        Count?: number;
    }
    
    export interface ApplicationInputParallelismProperties {
        Count?: pulumi.Input<number>;
    }
    
    export interface ApplicationInputProcessingConfigurationAttributes {
        InputLambdaProcessor?: ApplicationInputLambdaProcessorAttributes;
    }
    
    export interface ApplicationInputProcessingConfigurationProperties {
        InputLambdaProcessor?: pulumi.Input<ApplicationInputLambdaProcessorProperties>;
    }
    
    export interface ApplicationInputProperties {
        InputParallelism?: pulumi.Input<ApplicationInputParallelismProperties>;
        InputProcessingConfiguration?: pulumi.Input<ApplicationInputProcessingConfigurationProperties>;
        InputSchema: pulumi.Input<ApplicationInputSchemaProperties>;
        KinesisFirehoseInput?: pulumi.Input<ApplicationKinesisFirehoseInputProperties>;
        KinesisStreamsInput?: pulumi.Input<ApplicationKinesisStreamsInputProperties>;
        NamePrefix: pulumi.Input<string>;
    }
    
    export interface ApplicationInputSchemaAttributes {
        RecordColumns: ApplicationRecordColumnAttributes[];
        RecordEncoding?: string;
        RecordFormat: ApplicationRecordFormatAttributes;
    }
    
    export interface ApplicationInputSchemaProperties {
        RecordColumns: pulumi.Input<pulumi.Input<ApplicationRecordColumnProperties>[]>;
        RecordEncoding?: pulumi.Input<string>;
        RecordFormat: pulumi.Input<ApplicationRecordFormatProperties>;
    }
    
    export interface ApplicationJSONMappingParametersAttributes {
        RecordRowPath: string;
    }
    
    export interface ApplicationJSONMappingParametersProperties {
        RecordRowPath: pulumi.Input<string>;
    }
    
    export interface ApplicationKinesisFirehoseInputAttributes {
        ResourceARN: string;
        RoleARN: string;
    }
    
    export interface ApplicationKinesisFirehoseInputProperties {
        ResourceARN: pulumi.Input<string>;
        RoleARN: pulumi.Input<string>;
    }
    
    export interface ApplicationKinesisStreamsInputAttributes {
        ResourceARN: string;
        RoleARN: string;
    }
    
    export interface ApplicationKinesisStreamsInputProperties {
        ResourceARN: pulumi.Input<string>;
        RoleARN: pulumi.Input<string>;
    }
    
    export interface ApplicationMappingParametersAttributes {
        CSVMappingParameters?: ApplicationCSVMappingParametersAttributes;
        JSONMappingParameters?: ApplicationJSONMappingParametersAttributes;
    }
    
    export interface ApplicationMappingParametersProperties {
        CSVMappingParameters?: pulumi.Input<ApplicationCSVMappingParametersProperties>;
        JSONMappingParameters?: pulumi.Input<ApplicationJSONMappingParametersProperties>;
    }
    
    export interface ApplicationOutputAttributes {
    }
    
    export interface ApplicationOutputDestinationSchemaAttributes {
        RecordFormatType?: string;
    }
    
    export interface ApplicationOutputDestinationSchemaProperties {
        RecordFormatType?: pulumi.Input<string>;
    }
    
    export interface ApplicationOutputKinesisFirehoseOutputAttributes {
        ResourceARN: string;
        RoleARN: string;
    }
    
    export interface ApplicationOutputKinesisFirehoseOutputProperties {
        ResourceARN: pulumi.Input<string>;
        RoleARN: pulumi.Input<string>;
    }
    
    export interface ApplicationOutputKinesisStreamsOutputAttributes {
        ResourceARN: string;
        RoleARN: string;
    }
    
    export interface ApplicationOutputKinesisStreamsOutputProperties {
        ResourceARN: pulumi.Input<string>;
        RoleARN: pulumi.Input<string>;
    }
    
    export interface ApplicationOutputLambdaOutputAttributes {
        ResourceARN: string;
        RoleARN: string;
    }
    
    export interface ApplicationOutputLambdaOutputProperties {
        ResourceARN: pulumi.Input<string>;
        RoleARN: pulumi.Input<string>;
    }
    
    export interface ApplicationOutputOutputAttributes {
        DestinationSchema: ApplicationOutputDestinationSchemaAttributes;
        KinesisFirehoseOutput?: ApplicationOutputKinesisFirehoseOutputAttributes;
        KinesisStreamsOutput?: ApplicationOutputKinesisStreamsOutputAttributes;
        LambdaOutput?: ApplicationOutputLambdaOutputAttributes;
        Name?: string;
    }
    
    export interface ApplicationOutputOutputProperties {
        DestinationSchema: pulumi.Input<ApplicationOutputDestinationSchemaProperties>;
        KinesisFirehoseOutput?: pulumi.Input<ApplicationOutputKinesisFirehoseOutputProperties>;
        KinesisStreamsOutput?: pulumi.Input<ApplicationOutputKinesisStreamsOutputProperties>;
        LambdaOutput?: pulumi.Input<ApplicationOutputLambdaOutputProperties>;
        Name?: pulumi.Input<string>;
    }
    
    export interface ApplicationOutputProperties {
        ApplicationName: pulumi.Input<string>;
        Output: pulumi.Input<ApplicationOutputOutputProperties>;
    }
    
    export interface ApplicationProperties {
        ApplicationCode?: pulumi.Input<string>;
        ApplicationDescription?: pulumi.Input<string>;
        ApplicationName?: pulumi.Input<string>;
        Inputs: pulumi.Input<pulumi.Input<ApplicationInputProperties>[]>;
    }
    
    export interface ApplicationRecordColumnAttributes {
        Mapping?: string;
        Name: string;
        SqlType: string;
    }
    
    export interface ApplicationRecordColumnProperties {
        Mapping?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        SqlType: pulumi.Input<string>;
    }
    
    export interface ApplicationRecordFormatAttributes {
        MappingParameters?: ApplicationMappingParametersAttributes;
        RecordFormatType: string;
    }
    
    export interface ApplicationRecordFormatProperties {
        MappingParameters?: pulumi.Input<ApplicationMappingParametersProperties>;
        RecordFormatType: pulumi.Input<string>;
    }
    
    export interface ApplicationReferenceDataSourceAttributes {
    }
    
    export interface ApplicationReferenceDataSourceCSVMappingParametersAttributes {
        RecordColumnDelimiter: string;
        RecordRowDelimiter: string;
    }
    
    export interface ApplicationReferenceDataSourceCSVMappingParametersProperties {
        RecordColumnDelimiter: pulumi.Input<string>;
        RecordRowDelimiter: pulumi.Input<string>;
    }
    
    export interface ApplicationReferenceDataSourceJSONMappingParametersAttributes {
        RecordRowPath: string;
    }
    
    export interface ApplicationReferenceDataSourceJSONMappingParametersProperties {
        RecordRowPath: pulumi.Input<string>;
    }
    
    export interface ApplicationReferenceDataSourceMappingParametersAttributes {
        CSVMappingParameters?: ApplicationReferenceDataSourceCSVMappingParametersAttributes;
        JSONMappingParameters?: ApplicationReferenceDataSourceJSONMappingParametersAttributes;
    }
    
    export interface ApplicationReferenceDataSourceMappingParametersProperties {
        CSVMappingParameters?: pulumi.Input<ApplicationReferenceDataSourceCSVMappingParametersProperties>;
        JSONMappingParameters?: pulumi.Input<ApplicationReferenceDataSourceJSONMappingParametersProperties>;
    }
    
    export interface ApplicationReferenceDataSourceProperties {
        ApplicationName: pulumi.Input<string>;
        ReferenceDataSource: pulumi.Input<ApplicationReferenceDataSourceReferenceDataSourceProperties>;
    }
    
    export interface ApplicationReferenceDataSourceRecordColumnAttributes {
        Mapping?: string;
        Name: string;
        SqlType: string;
    }
    
    export interface ApplicationReferenceDataSourceRecordColumnProperties {
        Mapping?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        SqlType: pulumi.Input<string>;
    }
    
    export interface ApplicationReferenceDataSourceRecordFormatAttributes {
        MappingParameters?: ApplicationReferenceDataSourceMappingParametersAttributes;
        RecordFormatType: string;
    }
    
    export interface ApplicationReferenceDataSourceRecordFormatProperties {
        MappingParameters?: pulumi.Input<ApplicationReferenceDataSourceMappingParametersProperties>;
        RecordFormatType: pulumi.Input<string>;
    }
    
    export interface ApplicationReferenceDataSourceReferenceDataSourceAttributes {
        ReferenceSchema: ApplicationReferenceDataSourceReferenceSchemaAttributes;
        S3ReferenceDataSource?: ApplicationReferenceDataSourceS3ReferenceDataSourceAttributes;
        TableName?: string;
    }
    
    export interface ApplicationReferenceDataSourceReferenceDataSourceProperties {
        ReferenceSchema: pulumi.Input<ApplicationReferenceDataSourceReferenceSchemaProperties>;
        S3ReferenceDataSource?: pulumi.Input<ApplicationReferenceDataSourceS3ReferenceDataSourceProperties>;
        TableName?: pulumi.Input<string>;
    }
    
    export interface ApplicationReferenceDataSourceReferenceSchemaAttributes {
        RecordColumns: ApplicationReferenceDataSourceRecordColumnAttributes[];
        RecordEncoding?: string;
        RecordFormat: ApplicationReferenceDataSourceRecordFormatAttributes;
    }
    
    export interface ApplicationReferenceDataSourceReferenceSchemaProperties {
        RecordColumns: pulumi.Input<pulumi.Input<ApplicationReferenceDataSourceRecordColumnProperties>[]>;
        RecordEncoding?: pulumi.Input<string>;
        RecordFormat: pulumi.Input<ApplicationReferenceDataSourceRecordFormatProperties>;
    }
    
    export interface ApplicationReferenceDataSourceS3ReferenceDataSourceAttributes {
        BucketARN: string;
        FileKey: string;
        ReferenceRoleARN: string;
    }
    
    export interface ApplicationReferenceDataSourceS3ReferenceDataSourceProperties {
        BucketARN: pulumi.Input<string>;
        FileKey: pulumi.Input<string>;
        ReferenceRoleARN: pulumi.Input<string>;
    }
    
    
    export class Application extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ApplicationAttributes>;

        constructor(name: string, properties: ApplicationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:KinesisAnalytics:Application", name, inputs, opts);
        }
    }
    
    export class ApplicationOutput extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ApplicationOutputAttributes>;

        constructor(name: string, properties: ApplicationOutputProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:KinesisAnalytics:ApplicationOutput", name, inputs, opts);
        }
    }
    
    export class ApplicationReferenceDataSource extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ApplicationReferenceDataSourceAttributes>;

        constructor(name: string, properties: ApplicationReferenceDataSourceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:KinesisAnalytics:ApplicationReferenceDataSource", name, inputs, opts);
        }
    }
    
}

export namespace kinesisanalyticsv2 {
    
    export interface ApplicationApplicationCodeConfigurationAttributes {
        CodeContent: ApplicationCodeContentAttributes;
        CodeContentType: string;
    }
    
    export interface ApplicationApplicationCodeConfigurationProperties {
        CodeContent: pulumi.Input<ApplicationCodeContentProperties>;
        CodeContentType: pulumi.Input<string>;
    }
    
    export interface ApplicationApplicationConfigurationAttributes {
        ApplicationCodeConfiguration?: ApplicationApplicationCodeConfigurationAttributes;
        ApplicationSnapshotConfiguration?: ApplicationApplicationSnapshotConfigurationAttributes;
        EnvironmentProperties?: ApplicationEnvironmentPropertiesAttributes;
        FlinkApplicationConfiguration?: ApplicationFlinkApplicationConfigurationAttributes;
        SqlApplicationConfiguration?: ApplicationSqlApplicationConfigurationAttributes;
    }
    
    export interface ApplicationApplicationConfigurationProperties {
        ApplicationCodeConfiguration?: pulumi.Input<ApplicationApplicationCodeConfigurationProperties>;
        ApplicationSnapshotConfiguration?: pulumi.Input<ApplicationApplicationSnapshotConfigurationProperties>;
        EnvironmentProperties?: pulumi.Input<ApplicationEnvironmentPropertiesProperties>;
        FlinkApplicationConfiguration?: pulumi.Input<ApplicationFlinkApplicationConfigurationProperties>;
        SqlApplicationConfiguration?: pulumi.Input<ApplicationSqlApplicationConfigurationProperties>;
    }
    
    export interface ApplicationApplicationSnapshotConfigurationAttributes {
        SnapshotsEnabled: boolean;
    }
    
    export interface ApplicationApplicationSnapshotConfigurationProperties {
        SnapshotsEnabled: pulumi.Input<boolean>;
    }
    
    export interface ApplicationAttributes {
    }
    
    export interface ApplicationCSVMappingParametersAttributes {
        RecordColumnDelimiter: string;
        RecordRowDelimiter: string;
    }
    
    export interface ApplicationCSVMappingParametersProperties {
        RecordColumnDelimiter: pulumi.Input<string>;
        RecordRowDelimiter: pulumi.Input<string>;
    }
    
    export interface ApplicationCheckpointConfigurationAttributes {
        CheckpointInterval?: number;
        CheckpointingEnabled?: boolean;
        ConfigurationType: string;
        MinPauseBetweenCheckpoints?: number;
    }
    
    export interface ApplicationCheckpointConfigurationProperties {
        CheckpointInterval?: pulumi.Input<number>;
        CheckpointingEnabled?: pulumi.Input<boolean>;
        ConfigurationType: pulumi.Input<string>;
        MinPauseBetweenCheckpoints?: pulumi.Input<number>;
    }
    
    export interface ApplicationCloudWatchLoggingOptionAttributes {
    }
    
    export interface ApplicationCloudWatchLoggingOptionCloudWatchLoggingOptionAttributes {
        LogStreamARN: string;
    }
    
    export interface ApplicationCloudWatchLoggingOptionCloudWatchLoggingOptionProperties {
        LogStreamARN: pulumi.Input<string>;
    }
    
    export interface ApplicationCloudWatchLoggingOptionProperties {
        ApplicationName: pulumi.Input<string>;
        CloudWatchLoggingOption: pulumi.Input<ApplicationCloudWatchLoggingOptionCloudWatchLoggingOptionProperties>;
    }
    
    export interface ApplicationCodeContentAttributes {
        S3ContentLocation?: ApplicationS3ContentLocationAttributes;
        TextContent?: string;
        ZipFileContent?: string;
    }
    
    export interface ApplicationCodeContentProperties {
        S3ContentLocation?: pulumi.Input<ApplicationS3ContentLocationProperties>;
        TextContent?: pulumi.Input<string>;
        ZipFileContent?: pulumi.Input<string>;
    }
    
    export interface ApplicationEnvironmentPropertiesAttributes {
        PropertyGroups?: ApplicationPropertyGroupAttributes[];
    }
    
    export interface ApplicationEnvironmentPropertiesProperties {
        PropertyGroups?: pulumi.Input<pulumi.Input<ApplicationPropertyGroupProperties>[]>;
    }
    
    export interface ApplicationFlinkApplicationConfigurationAttributes {
        CheckpointConfiguration?: ApplicationCheckpointConfigurationAttributes;
        MonitoringConfiguration?: ApplicationMonitoringConfigurationAttributes;
        ParallelismConfiguration?: ApplicationParallelismConfigurationAttributes;
    }
    
    export interface ApplicationFlinkApplicationConfigurationProperties {
        CheckpointConfiguration?: pulumi.Input<ApplicationCheckpointConfigurationProperties>;
        MonitoringConfiguration?: pulumi.Input<ApplicationMonitoringConfigurationProperties>;
        ParallelismConfiguration?: pulumi.Input<ApplicationParallelismConfigurationProperties>;
    }
    
    export interface ApplicationInputAttributes {
        InputParallelism?: ApplicationInputParallelismAttributes;
        InputProcessingConfiguration?: ApplicationInputProcessingConfigurationAttributes;
        InputSchema: ApplicationInputSchemaAttributes;
        KinesisFirehoseInput?: ApplicationKinesisFirehoseInputAttributes;
        KinesisStreamsInput?: ApplicationKinesisStreamsInputAttributes;
        NamePrefix: string;
    }
    
    export interface ApplicationInputLambdaProcessorAttributes {
        ResourceARN: string;
    }
    
    export interface ApplicationInputLambdaProcessorProperties {
        ResourceARN: pulumi.Input<string>;
    }
    
    export interface ApplicationInputParallelismAttributes {
        Count?: number;
    }
    
    export interface ApplicationInputParallelismProperties {
        Count?: pulumi.Input<number>;
    }
    
    export interface ApplicationInputProcessingConfigurationAttributes {
        InputLambdaProcessor?: ApplicationInputLambdaProcessorAttributes;
    }
    
    export interface ApplicationInputProcessingConfigurationProperties {
        InputLambdaProcessor?: pulumi.Input<ApplicationInputLambdaProcessorProperties>;
    }
    
    export interface ApplicationInputProperties {
        InputParallelism?: pulumi.Input<ApplicationInputParallelismProperties>;
        InputProcessingConfiguration?: pulumi.Input<ApplicationInputProcessingConfigurationProperties>;
        InputSchema: pulumi.Input<ApplicationInputSchemaProperties>;
        KinesisFirehoseInput?: pulumi.Input<ApplicationKinesisFirehoseInputProperties>;
        KinesisStreamsInput?: pulumi.Input<ApplicationKinesisStreamsInputProperties>;
        NamePrefix: pulumi.Input<string>;
    }
    
    export interface ApplicationInputSchemaAttributes {
        RecordColumns: ApplicationRecordColumnAttributes[];
        RecordEncoding?: string;
        RecordFormat: ApplicationRecordFormatAttributes;
    }
    
    export interface ApplicationInputSchemaProperties {
        RecordColumns: pulumi.Input<pulumi.Input<ApplicationRecordColumnProperties>[]>;
        RecordEncoding?: pulumi.Input<string>;
        RecordFormat: pulumi.Input<ApplicationRecordFormatProperties>;
    }
    
    export interface ApplicationJSONMappingParametersAttributes {
        RecordRowPath: string;
    }
    
    export interface ApplicationJSONMappingParametersProperties {
        RecordRowPath: pulumi.Input<string>;
    }
    
    export interface ApplicationKinesisFirehoseInputAttributes {
        ResourceARN: string;
    }
    
    export interface ApplicationKinesisFirehoseInputProperties {
        ResourceARN: pulumi.Input<string>;
    }
    
    export interface ApplicationKinesisStreamsInputAttributes {
        ResourceARN: string;
    }
    
    export interface ApplicationKinesisStreamsInputProperties {
        ResourceARN: pulumi.Input<string>;
    }
    
    export interface ApplicationMappingParametersAttributes {
        CSVMappingParameters?: ApplicationCSVMappingParametersAttributes;
        JSONMappingParameters?: ApplicationJSONMappingParametersAttributes;
    }
    
    export interface ApplicationMappingParametersProperties {
        CSVMappingParameters?: pulumi.Input<ApplicationCSVMappingParametersProperties>;
        JSONMappingParameters?: pulumi.Input<ApplicationJSONMappingParametersProperties>;
    }
    
    export interface ApplicationMonitoringConfigurationAttributes {
        ConfigurationType: string;
        LogLevel?: string;
        MetricsLevel?: string;
    }
    
    export interface ApplicationMonitoringConfigurationProperties {
        ConfigurationType: pulumi.Input<string>;
        LogLevel?: pulumi.Input<string>;
        MetricsLevel?: pulumi.Input<string>;
    }
    
    export interface ApplicationOutputAttributes {
    }
    
    export interface ApplicationOutputDestinationSchemaAttributes {
        RecordFormatType?: string;
    }
    
    export interface ApplicationOutputDestinationSchemaProperties {
        RecordFormatType?: pulumi.Input<string>;
    }
    
    export interface ApplicationOutputKinesisFirehoseOutputAttributes {
        ResourceARN: string;
    }
    
    export interface ApplicationOutputKinesisFirehoseOutputProperties {
        ResourceARN: pulumi.Input<string>;
    }
    
    export interface ApplicationOutputKinesisStreamsOutputAttributes {
        ResourceARN: string;
    }
    
    export interface ApplicationOutputKinesisStreamsOutputProperties {
        ResourceARN: pulumi.Input<string>;
    }
    
    export interface ApplicationOutputLambdaOutputAttributes {
        ResourceARN: string;
    }
    
    export interface ApplicationOutputLambdaOutputProperties {
        ResourceARN: pulumi.Input<string>;
    }
    
    export interface ApplicationOutputOutputAttributes {
        DestinationSchema: ApplicationOutputDestinationSchemaAttributes;
        KinesisFirehoseOutput?: ApplicationOutputKinesisFirehoseOutputAttributes;
        KinesisStreamsOutput?: ApplicationOutputKinesisStreamsOutputAttributes;
        LambdaOutput?: ApplicationOutputLambdaOutputAttributes;
        Name?: string;
    }
    
    export interface ApplicationOutputOutputProperties {
        DestinationSchema: pulumi.Input<ApplicationOutputDestinationSchemaProperties>;
        KinesisFirehoseOutput?: pulumi.Input<ApplicationOutputKinesisFirehoseOutputProperties>;
        KinesisStreamsOutput?: pulumi.Input<ApplicationOutputKinesisStreamsOutputProperties>;
        LambdaOutput?: pulumi.Input<ApplicationOutputLambdaOutputProperties>;
        Name?: pulumi.Input<string>;
    }
    
    export interface ApplicationOutputProperties {
        ApplicationName: pulumi.Input<string>;
        Output: pulumi.Input<ApplicationOutputOutputProperties>;
    }
    
    export interface ApplicationParallelismConfigurationAttributes {
        AutoScalingEnabled?: boolean;
        ConfigurationType: string;
        Parallelism?: number;
        ParallelismPerKPU?: number;
    }
    
    export interface ApplicationParallelismConfigurationProperties {
        AutoScalingEnabled?: pulumi.Input<boolean>;
        ConfigurationType: pulumi.Input<string>;
        Parallelism?: pulumi.Input<number>;
        ParallelismPerKPU?: pulumi.Input<number>;
    }
    
    export interface ApplicationProperties {
        ApplicationConfiguration?: pulumi.Input<ApplicationApplicationConfigurationProperties>;
        ApplicationDescription?: pulumi.Input<string>;
        ApplicationName?: pulumi.Input<string>;
        RuntimeEnvironment: pulumi.Input<string>;
        ServiceExecutionRole: pulumi.Input<string>;
    }
    
    export interface ApplicationPropertyGroupAttributes {
        PropertyGroupId?: string;
        PropertyMap?: string;
    }
    
    export interface ApplicationPropertyGroupProperties {
        PropertyGroupId?: pulumi.Input<string>;
        PropertyMap?: pulumi.Input<string>;
    }
    
    export interface ApplicationRecordColumnAttributes {
        Mapping?: string;
        Name: string;
        SqlType: string;
    }
    
    export interface ApplicationRecordColumnProperties {
        Mapping?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        SqlType: pulumi.Input<string>;
    }
    
    export interface ApplicationRecordFormatAttributes {
        MappingParameters?: ApplicationMappingParametersAttributes;
        RecordFormatType: string;
    }
    
    export interface ApplicationRecordFormatProperties {
        MappingParameters?: pulumi.Input<ApplicationMappingParametersProperties>;
        RecordFormatType: pulumi.Input<string>;
    }
    
    export interface ApplicationReferenceDataSourceAttributes {
    }
    
    export interface ApplicationReferenceDataSourceCSVMappingParametersAttributes {
        RecordColumnDelimiter: string;
        RecordRowDelimiter: string;
    }
    
    export interface ApplicationReferenceDataSourceCSVMappingParametersProperties {
        RecordColumnDelimiter: pulumi.Input<string>;
        RecordRowDelimiter: pulumi.Input<string>;
    }
    
    export interface ApplicationReferenceDataSourceJSONMappingParametersAttributes {
        RecordRowPath: string;
    }
    
    export interface ApplicationReferenceDataSourceJSONMappingParametersProperties {
        RecordRowPath: pulumi.Input<string>;
    }
    
    export interface ApplicationReferenceDataSourceMappingParametersAttributes {
        CSVMappingParameters?: ApplicationReferenceDataSourceCSVMappingParametersAttributes;
        JSONMappingParameters?: ApplicationReferenceDataSourceJSONMappingParametersAttributes;
    }
    
    export interface ApplicationReferenceDataSourceMappingParametersProperties {
        CSVMappingParameters?: pulumi.Input<ApplicationReferenceDataSourceCSVMappingParametersProperties>;
        JSONMappingParameters?: pulumi.Input<ApplicationReferenceDataSourceJSONMappingParametersProperties>;
    }
    
    export interface ApplicationReferenceDataSourceProperties {
        ApplicationName: pulumi.Input<string>;
        ReferenceDataSource: pulumi.Input<ApplicationReferenceDataSourceReferenceDataSourceProperties>;
    }
    
    export interface ApplicationReferenceDataSourceRecordColumnAttributes {
        Mapping?: string;
        Name: string;
        SqlType: string;
    }
    
    export interface ApplicationReferenceDataSourceRecordColumnProperties {
        Mapping?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        SqlType: pulumi.Input<string>;
    }
    
    export interface ApplicationReferenceDataSourceRecordFormatAttributes {
        MappingParameters?: ApplicationReferenceDataSourceMappingParametersAttributes;
        RecordFormatType: string;
    }
    
    export interface ApplicationReferenceDataSourceRecordFormatProperties {
        MappingParameters?: pulumi.Input<ApplicationReferenceDataSourceMappingParametersProperties>;
        RecordFormatType: pulumi.Input<string>;
    }
    
    export interface ApplicationReferenceDataSourceReferenceDataSourceAttributes {
        ReferenceSchema: ApplicationReferenceDataSourceReferenceSchemaAttributes;
        S3ReferenceDataSource?: ApplicationReferenceDataSourceS3ReferenceDataSourceAttributes;
        TableName?: string;
    }
    
    export interface ApplicationReferenceDataSourceReferenceDataSourceProperties {
        ReferenceSchema: pulumi.Input<ApplicationReferenceDataSourceReferenceSchemaProperties>;
        S3ReferenceDataSource?: pulumi.Input<ApplicationReferenceDataSourceS3ReferenceDataSourceProperties>;
        TableName?: pulumi.Input<string>;
    }
    
    export interface ApplicationReferenceDataSourceReferenceSchemaAttributes {
        RecordColumns: ApplicationReferenceDataSourceRecordColumnAttributes[];
        RecordEncoding?: string;
        RecordFormat: ApplicationReferenceDataSourceRecordFormatAttributes;
    }
    
    export interface ApplicationReferenceDataSourceReferenceSchemaProperties {
        RecordColumns: pulumi.Input<pulumi.Input<ApplicationReferenceDataSourceRecordColumnProperties>[]>;
        RecordEncoding?: pulumi.Input<string>;
        RecordFormat: pulumi.Input<ApplicationReferenceDataSourceRecordFormatProperties>;
    }
    
    export interface ApplicationReferenceDataSourceS3ReferenceDataSourceAttributes {
        BucketARN: string;
        FileKey: string;
    }
    
    export interface ApplicationReferenceDataSourceS3ReferenceDataSourceProperties {
        BucketARN: pulumi.Input<string>;
        FileKey: pulumi.Input<string>;
    }
    
    export interface ApplicationS3ContentLocationAttributes {
        BucketARN?: string;
        FileKey?: string;
        ObjectVersion?: string;
    }
    
    export interface ApplicationS3ContentLocationProperties {
        BucketARN?: pulumi.Input<string>;
        FileKey?: pulumi.Input<string>;
        ObjectVersion?: pulumi.Input<string>;
    }
    
    export interface ApplicationSqlApplicationConfigurationAttributes {
        Inputs?: ApplicationInputAttributes[];
    }
    
    export interface ApplicationSqlApplicationConfigurationProperties {
        Inputs?: pulumi.Input<pulumi.Input<ApplicationInputProperties>[]>;
    }
    
    
    export class Application extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ApplicationAttributes>;

        constructor(name: string, properties: ApplicationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:KinesisAnalyticsV2:Application", name, inputs, opts);
        }
    }
    
    export class ApplicationCloudWatchLoggingOption extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ApplicationCloudWatchLoggingOptionAttributes>;

        constructor(name: string, properties: ApplicationCloudWatchLoggingOptionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:KinesisAnalyticsV2:ApplicationCloudWatchLoggingOption", name, inputs, opts);
        }
    }
    
    export class ApplicationOutput extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ApplicationOutputAttributes>;

        constructor(name: string, properties: ApplicationOutputProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:KinesisAnalyticsV2:ApplicationOutput", name, inputs, opts);
        }
    }
    
    export class ApplicationReferenceDataSource extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ApplicationReferenceDataSourceAttributes>;

        constructor(name: string, properties: ApplicationReferenceDataSourceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:KinesisAnalyticsV2:ApplicationReferenceDataSource", name, inputs, opts);
        }
    }
    
}

export namespace kinesisfirehose {
    
    export interface DeliveryStreamAttributes {
        Arn: string;
    }
    
    export interface DeliveryStreamBufferingHintsAttributes {
        IntervalInSeconds: number;
        SizeInMBs: number;
    }
    
    export interface DeliveryStreamBufferingHintsProperties {
        IntervalInSeconds: pulumi.Input<number>;
        SizeInMBs: pulumi.Input<number>;
    }
    
    export interface DeliveryStreamCloudWatchLoggingOptionsAttributes {
        Enabled?: boolean;
        LogGroupName?: string;
        LogStreamName?: string;
    }
    
    export interface DeliveryStreamCloudWatchLoggingOptionsProperties {
        Enabled?: pulumi.Input<boolean>;
        LogGroupName?: pulumi.Input<string>;
        LogStreamName?: pulumi.Input<string>;
    }
    
    export interface DeliveryStreamCopyCommandAttributes {
        CopyOptions?: string;
        DataTableColumns?: string;
        DataTableName: string;
    }
    
    export interface DeliveryStreamCopyCommandProperties {
        CopyOptions?: pulumi.Input<string>;
        DataTableColumns?: pulumi.Input<string>;
        DataTableName: pulumi.Input<string>;
    }
    
    export interface DeliveryStreamDataFormatConversionConfigurationAttributes {
        Enabled: boolean;
        InputFormatConfiguration: DeliveryStreamInputFormatConfigurationAttributes;
        OutputFormatConfiguration: DeliveryStreamOutputFormatConfigurationAttributes;
        SchemaConfiguration: DeliveryStreamSchemaConfigurationAttributes;
    }
    
    export interface DeliveryStreamDataFormatConversionConfigurationProperties {
        Enabled: pulumi.Input<boolean>;
        InputFormatConfiguration: pulumi.Input<DeliveryStreamInputFormatConfigurationProperties>;
        OutputFormatConfiguration: pulumi.Input<DeliveryStreamOutputFormatConfigurationProperties>;
        SchemaConfiguration: pulumi.Input<DeliveryStreamSchemaConfigurationProperties>;
    }
    
    export interface DeliveryStreamDeserializerAttributes {
        HiveJsonSerDe?: DeliveryStreamHiveJsonSerDeAttributes;
        OpenXJsonSerDe?: DeliveryStreamOpenXJsonSerDeAttributes;
    }
    
    export interface DeliveryStreamDeserializerProperties {
        HiveJsonSerDe?: pulumi.Input<DeliveryStreamHiveJsonSerDeProperties>;
        OpenXJsonSerDe?: pulumi.Input<DeliveryStreamOpenXJsonSerDeProperties>;
    }
    
    export interface DeliveryStreamElasticsearchBufferingHintsAttributes {
        IntervalInSeconds: number;
        SizeInMBs: number;
    }
    
    export interface DeliveryStreamElasticsearchBufferingHintsProperties {
        IntervalInSeconds: pulumi.Input<number>;
        SizeInMBs: pulumi.Input<number>;
    }
    
    export interface DeliveryStreamElasticsearchDestinationConfigurationAttributes {
        BufferingHints: DeliveryStreamElasticsearchBufferingHintsAttributes;
        CloudWatchLoggingOptions?: DeliveryStreamCloudWatchLoggingOptionsAttributes;
        DomainARN: string;
        IndexName: string;
        IndexRotationPeriod: string;
        ProcessingConfiguration?: DeliveryStreamProcessingConfigurationAttributes;
        RetryOptions: DeliveryStreamElasticsearchRetryOptionsAttributes;
        RoleARN: string;
        S3BackupMode: string;
        S3Configuration: DeliveryStreamS3DestinationConfigurationAttributes;
        TypeName: string;
    }
    
    export interface DeliveryStreamElasticsearchDestinationConfigurationProperties {
        BufferingHints: pulumi.Input<DeliveryStreamElasticsearchBufferingHintsProperties>;
        CloudWatchLoggingOptions?: pulumi.Input<DeliveryStreamCloudWatchLoggingOptionsProperties>;
        DomainARN: pulumi.Input<string>;
        IndexName: pulumi.Input<string>;
        IndexRotationPeriod: pulumi.Input<string>;
        ProcessingConfiguration?: pulumi.Input<DeliveryStreamProcessingConfigurationProperties>;
        RetryOptions: pulumi.Input<DeliveryStreamElasticsearchRetryOptionsProperties>;
        RoleARN: pulumi.Input<string>;
        S3BackupMode: pulumi.Input<string>;
        S3Configuration: pulumi.Input<DeliveryStreamS3DestinationConfigurationProperties>;
        TypeName: pulumi.Input<string>;
    }
    
    export interface DeliveryStreamElasticsearchRetryOptionsAttributes {
        DurationInSeconds: number;
    }
    
    export interface DeliveryStreamElasticsearchRetryOptionsProperties {
        DurationInSeconds: pulumi.Input<number>;
    }
    
    export interface DeliveryStreamEncryptionConfigurationAttributes {
        KMSEncryptionConfig?: DeliveryStreamKMSEncryptionConfigAttributes;
        NoEncryptionConfig?: string;
    }
    
    export interface DeliveryStreamEncryptionConfigurationProperties {
        KMSEncryptionConfig?: pulumi.Input<DeliveryStreamKMSEncryptionConfigProperties>;
        NoEncryptionConfig?: pulumi.Input<string>;
    }
    
    export interface DeliveryStreamExtendedS3DestinationConfigurationAttributes {
        BucketARN: string;
        BufferingHints: DeliveryStreamBufferingHintsAttributes;
        CloudWatchLoggingOptions?: DeliveryStreamCloudWatchLoggingOptionsAttributes;
        CompressionFormat: string;
        DataFormatConversionConfiguration?: DeliveryStreamDataFormatConversionConfigurationAttributes;
        EncryptionConfiguration?: DeliveryStreamEncryptionConfigurationAttributes;
        ErrorOutputPrefix?: string;
        Prefix?: string;
        ProcessingConfiguration?: DeliveryStreamProcessingConfigurationAttributes;
        RoleARN: string;
        S3BackupConfiguration?: DeliveryStreamS3DestinationConfigurationAttributes;
        S3BackupMode?: string;
    }
    
    export interface DeliveryStreamExtendedS3DestinationConfigurationProperties {
        BucketARN: pulumi.Input<string>;
        BufferingHints: pulumi.Input<DeliveryStreamBufferingHintsProperties>;
        CloudWatchLoggingOptions?: pulumi.Input<DeliveryStreamCloudWatchLoggingOptionsProperties>;
        CompressionFormat: pulumi.Input<string>;
        DataFormatConversionConfiguration?: pulumi.Input<DeliveryStreamDataFormatConversionConfigurationProperties>;
        EncryptionConfiguration?: pulumi.Input<DeliveryStreamEncryptionConfigurationProperties>;
        ErrorOutputPrefix?: pulumi.Input<string>;
        Prefix?: pulumi.Input<string>;
        ProcessingConfiguration?: pulumi.Input<DeliveryStreamProcessingConfigurationProperties>;
        RoleARN: pulumi.Input<string>;
        S3BackupConfiguration?: pulumi.Input<DeliveryStreamS3DestinationConfigurationProperties>;
        S3BackupMode?: pulumi.Input<string>;
    }
    
    export interface DeliveryStreamHiveJsonSerDeAttributes {
        TimestampFormats?: string[];
    }
    
    export interface DeliveryStreamHiveJsonSerDeProperties {
        TimestampFormats?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface DeliveryStreamInputFormatConfigurationAttributes {
        Deserializer: DeliveryStreamDeserializerAttributes;
    }
    
    export interface DeliveryStreamInputFormatConfigurationProperties {
        Deserializer: pulumi.Input<DeliveryStreamDeserializerProperties>;
    }
    
    export interface DeliveryStreamKMSEncryptionConfigAttributes {
        AWSKMSKeyARN: string;
    }
    
    export interface DeliveryStreamKMSEncryptionConfigProperties {
        AWSKMSKeyARN: pulumi.Input<string>;
    }
    
    export interface DeliveryStreamKinesisStreamSourceConfigurationAttributes {
        KinesisStreamARN: string;
        RoleARN: string;
    }
    
    export interface DeliveryStreamKinesisStreamSourceConfigurationProperties {
        KinesisStreamARN: pulumi.Input<string>;
        RoleARN: pulumi.Input<string>;
    }
    
    export interface DeliveryStreamOpenXJsonSerDeAttributes {
        CaseInsensitive?: boolean;
        ColumnToJsonKeyMappings?: {[key: string]: string};
        ConvertDotsInJsonKeysToUnderscores?: boolean;
    }
    
    export interface DeliveryStreamOpenXJsonSerDeProperties {
        CaseInsensitive?: pulumi.Input<boolean>;
        ColumnToJsonKeyMappings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        ConvertDotsInJsonKeysToUnderscores?: pulumi.Input<boolean>;
    }
    
    export interface DeliveryStreamOrcSerDeAttributes {
        BlockSizeBytes?: number;
        BloomFilterColumns?: string[];
        BloomFilterFalsePositiveProbability?: number;
        Compression?: string;
        DictionaryKeyThreshold?: number;
        EnablePadding?: boolean;
        FormatVersion?: string;
        PaddingTolerance?: number;
        RowIndexStride?: number;
        StripeSizeBytes?: number;
    }
    
    export interface DeliveryStreamOrcSerDeProperties {
        BlockSizeBytes?: pulumi.Input<number>;
        BloomFilterColumns?: pulumi.Input<pulumi.Input<string>[]>;
        BloomFilterFalsePositiveProbability?: pulumi.Input<number>;
        Compression?: pulumi.Input<string>;
        DictionaryKeyThreshold?: pulumi.Input<number>;
        EnablePadding?: pulumi.Input<boolean>;
        FormatVersion?: pulumi.Input<string>;
        PaddingTolerance?: pulumi.Input<number>;
        RowIndexStride?: pulumi.Input<number>;
        StripeSizeBytes?: pulumi.Input<number>;
    }
    
    export interface DeliveryStreamOutputFormatConfigurationAttributes {
        Serializer: DeliveryStreamSerializerAttributes;
    }
    
    export interface DeliveryStreamOutputFormatConfigurationProperties {
        Serializer: pulumi.Input<DeliveryStreamSerializerProperties>;
    }
    
    export interface DeliveryStreamParquetSerDeAttributes {
        BlockSizeBytes?: number;
        Compression?: string;
        EnableDictionaryCompression?: boolean;
        MaxPaddingBytes?: number;
        PageSizeBytes?: number;
        WriterVersion?: string;
    }
    
    export interface DeliveryStreamParquetSerDeProperties {
        BlockSizeBytes?: pulumi.Input<number>;
        Compression?: pulumi.Input<string>;
        EnableDictionaryCompression?: pulumi.Input<boolean>;
        MaxPaddingBytes?: pulumi.Input<number>;
        PageSizeBytes?: pulumi.Input<number>;
        WriterVersion?: pulumi.Input<string>;
    }
    
    export interface DeliveryStreamProcessingConfigurationAttributes {
        Enabled?: boolean;
        Processors?: DeliveryStreamProcessorAttributes[];
    }
    
    export interface DeliveryStreamProcessingConfigurationProperties {
        Enabled?: pulumi.Input<boolean>;
        Processors?: pulumi.Input<pulumi.Input<DeliveryStreamProcessorProperties>[]>;
    }
    
    export interface DeliveryStreamProcessorAttributes {
        Parameters: DeliveryStreamProcessorParameterAttributes[];
        Type: string;
    }
    
    export interface DeliveryStreamProcessorParameterAttributes {
        ParameterName: string;
        ParameterValue: string;
    }
    
    export interface DeliveryStreamProcessorParameterProperties {
        ParameterName: pulumi.Input<string>;
        ParameterValue: pulumi.Input<string>;
    }
    
    export interface DeliveryStreamProcessorProperties {
        Parameters: pulumi.Input<pulumi.Input<DeliveryStreamProcessorParameterProperties>[]>;
        Type: pulumi.Input<string>;
    }
    
    export interface DeliveryStreamProperties {
        DeliveryStreamName?: pulumi.Input<string>;
        DeliveryStreamType?: pulumi.Input<string>;
        ElasticsearchDestinationConfiguration?: pulumi.Input<DeliveryStreamElasticsearchDestinationConfigurationProperties>;
        ExtendedS3DestinationConfiguration?: pulumi.Input<DeliveryStreamExtendedS3DestinationConfigurationProperties>;
        KinesisStreamSourceConfiguration?: pulumi.Input<DeliveryStreamKinesisStreamSourceConfigurationProperties>;
        RedshiftDestinationConfiguration?: pulumi.Input<DeliveryStreamRedshiftDestinationConfigurationProperties>;
        S3DestinationConfiguration?: pulumi.Input<DeliveryStreamS3DestinationConfigurationProperties>;
        SplunkDestinationConfiguration?: pulumi.Input<DeliveryStreamSplunkDestinationConfigurationProperties>;
    }
    
    export interface DeliveryStreamRedshiftDestinationConfigurationAttributes {
        CloudWatchLoggingOptions?: DeliveryStreamCloudWatchLoggingOptionsAttributes;
        ClusterJDBCURL: string;
        CopyCommand: DeliveryStreamCopyCommandAttributes;
        Password: string;
        ProcessingConfiguration?: DeliveryStreamProcessingConfigurationAttributes;
        RoleARN: string;
        S3Configuration: DeliveryStreamS3DestinationConfigurationAttributes;
        Username: string;
    }
    
    export interface DeliveryStreamRedshiftDestinationConfigurationProperties {
        CloudWatchLoggingOptions?: pulumi.Input<DeliveryStreamCloudWatchLoggingOptionsProperties>;
        ClusterJDBCURL: pulumi.Input<string>;
        CopyCommand: pulumi.Input<DeliveryStreamCopyCommandProperties>;
        Password: pulumi.Input<string>;
        ProcessingConfiguration?: pulumi.Input<DeliveryStreamProcessingConfigurationProperties>;
        RoleARN: pulumi.Input<string>;
        S3Configuration: pulumi.Input<DeliveryStreamS3DestinationConfigurationProperties>;
        Username: pulumi.Input<string>;
    }
    
    export interface DeliveryStreamS3DestinationConfigurationAttributes {
        BucketARN: string;
        BufferingHints: DeliveryStreamBufferingHintsAttributes;
        CloudWatchLoggingOptions?: DeliveryStreamCloudWatchLoggingOptionsAttributes;
        CompressionFormat: string;
        EncryptionConfiguration?: DeliveryStreamEncryptionConfigurationAttributes;
        ErrorOutputPrefix?: string;
        Prefix?: string;
        RoleARN: string;
    }
    
    export interface DeliveryStreamS3DestinationConfigurationProperties {
        BucketARN: pulumi.Input<string>;
        BufferingHints: pulumi.Input<DeliveryStreamBufferingHintsProperties>;
        CloudWatchLoggingOptions?: pulumi.Input<DeliveryStreamCloudWatchLoggingOptionsProperties>;
        CompressionFormat: pulumi.Input<string>;
        EncryptionConfiguration?: pulumi.Input<DeliveryStreamEncryptionConfigurationProperties>;
        ErrorOutputPrefix?: pulumi.Input<string>;
        Prefix?: pulumi.Input<string>;
        RoleARN: pulumi.Input<string>;
    }
    
    export interface DeliveryStreamSchemaConfigurationAttributes {
        CatalogId: string;
        DatabaseName: string;
        Region: string;
        RoleARN: string;
        TableName: string;
        VersionId: string;
    }
    
    export interface DeliveryStreamSchemaConfigurationProperties {
        CatalogId: pulumi.Input<string>;
        DatabaseName: pulumi.Input<string>;
        Region: pulumi.Input<string>;
        RoleARN: pulumi.Input<string>;
        TableName: pulumi.Input<string>;
        VersionId: pulumi.Input<string>;
    }
    
    export interface DeliveryStreamSerializerAttributes {
        OrcSerDe?: DeliveryStreamOrcSerDeAttributes;
        ParquetSerDe?: DeliveryStreamParquetSerDeAttributes;
    }
    
    export interface DeliveryStreamSerializerProperties {
        OrcSerDe?: pulumi.Input<DeliveryStreamOrcSerDeProperties>;
        ParquetSerDe?: pulumi.Input<DeliveryStreamParquetSerDeProperties>;
    }
    
    export interface DeliveryStreamSplunkDestinationConfigurationAttributes {
        CloudWatchLoggingOptions?: DeliveryStreamCloudWatchLoggingOptionsAttributes;
        HECAcknowledgmentTimeoutInSeconds?: number;
        HECEndpoint: string;
        HECEndpointType: string;
        HECToken: string;
        ProcessingConfiguration?: DeliveryStreamProcessingConfigurationAttributes;
        RetryOptions?: DeliveryStreamSplunkRetryOptionsAttributes;
        S3BackupMode?: string;
        S3Configuration: DeliveryStreamS3DestinationConfigurationAttributes;
    }
    
    export interface DeliveryStreamSplunkDestinationConfigurationProperties {
        CloudWatchLoggingOptions?: pulumi.Input<DeliveryStreamCloudWatchLoggingOptionsProperties>;
        HECAcknowledgmentTimeoutInSeconds?: pulumi.Input<number>;
        HECEndpoint: pulumi.Input<string>;
        HECEndpointType: pulumi.Input<string>;
        HECToken: pulumi.Input<string>;
        ProcessingConfiguration?: pulumi.Input<DeliveryStreamProcessingConfigurationProperties>;
        RetryOptions?: pulumi.Input<DeliveryStreamSplunkRetryOptionsProperties>;
        S3BackupMode?: pulumi.Input<string>;
        S3Configuration: pulumi.Input<DeliveryStreamS3DestinationConfigurationProperties>;
    }
    
    export interface DeliveryStreamSplunkRetryOptionsAttributes {
        DurationInSeconds: number;
    }
    
    export interface DeliveryStreamSplunkRetryOptionsProperties {
        DurationInSeconds: pulumi.Input<number>;
    }
    
    
    export class DeliveryStream extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DeliveryStreamAttributes>;

        constructor(name: string, properties?: DeliveryStreamProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:KinesisFirehose:DeliveryStream", name, inputs, opts);
        }
    }
    
}

export namespace lakeformation {
    
    export interface DataLakeSettingsAdminsAttributes {
    }
    
    export interface DataLakeSettingsAdminsProperties {
    }
    
    export interface DataLakeSettingsAttributes {
    }
    
    export interface DataLakeSettingsDataLakePrincipalAttributes {
        DataLakePrincipalIdentifier?: string;
    }
    
    export interface DataLakeSettingsDataLakePrincipalProperties {
        DataLakePrincipalIdentifier?: pulumi.Input<string>;
    }
    
    export interface DataLakeSettingsProperties {
        Admins?: pulumi.Input<DataLakeSettingsAdminsProperties>;
    }
    
    export interface PermissionsAttributes {
    }
    
    export interface PermissionsDataLakePrincipalAttributes {
        DataLakePrincipalIdentifier?: string;
    }
    
    export interface PermissionsDataLakePrincipalProperties {
        DataLakePrincipalIdentifier?: pulumi.Input<string>;
    }
    
    export interface PermissionsDatabaseResourceAttributes {
        Name?: string;
    }
    
    export interface PermissionsDatabaseResourceProperties {
        Name?: pulumi.Input<string>;
    }
    
    export interface PermissionsProperties {
        DataLakePrincipal: pulumi.Input<PermissionsDataLakePrincipalProperties>;
        Permissions?: pulumi.Input<pulumi.Input<string>[]>;
        PermissionsWithGrantOption?: pulumi.Input<pulumi.Input<string>[]>;
        Resource: pulumi.Input<PermissionsResourceProperties>;
    }
    
    export interface PermissionsResourceAttributes {
        DatabaseResource?: PermissionsDatabaseResourceAttributes;
        TableResource?: PermissionsTableResourceAttributes;
    }
    
    export interface PermissionsResourceProperties {
        DatabaseResource?: pulumi.Input<PermissionsDatabaseResourceProperties>;
        TableResource?: pulumi.Input<PermissionsTableResourceProperties>;
    }
    
    export interface PermissionsTableResourceAttributes {
        DatabaseName?: string;
        Name?: string;
    }
    
    export interface PermissionsTableResourceProperties {
        DatabaseName?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
    }
    
    export interface ResourceAttributes {
    }
    
    export interface ResourceProperties {
        ResourceArn: pulumi.Input<string>;
        RoleArn?: pulumi.Input<string>;
        UseServiceLinkedRole: pulumi.Input<boolean>;
    }
    
    
    export class DataLakeSettings extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DataLakeSettingsAttributes>;

        constructor(name: string, properties?: DataLakeSettingsProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:LakeFormation:DataLakeSettings", name, inputs, opts);
        }
    }
    
    export class Permissions extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PermissionsAttributes>;

        constructor(name: string, properties: PermissionsProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:LakeFormation:Permissions", name, inputs, opts);
        }
    }
    
    export class Resource extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ResourceAttributes>;

        constructor(name: string, properties: ResourceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:LakeFormation:Resource", name, inputs, opts);
        }
    }
    
}

export namespace lambda {
    
    export interface AliasAliasRoutingConfigurationAttributes {
        AdditionalVersionWeights: AliasVersionWeightAttributes[];
    }
    
    export interface AliasAliasRoutingConfigurationProperties {
        AdditionalVersionWeights: pulumi.Input<pulumi.Input<AliasVersionWeightProperties>[]>;
    }
    
    export interface AliasAttributes {
    }
    
    export interface AliasProperties {
        Description?: pulumi.Input<string>;
        FunctionName: pulumi.Input<string>;
        FunctionVersion: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        RoutingConfig?: pulumi.Input<AliasAliasRoutingConfigurationProperties>;
    }
    
    export interface AliasVersionWeightAttributes {
        FunctionVersion: string;
        FunctionWeight: number;
    }
    
    export interface AliasVersionWeightProperties {
        FunctionVersion: pulumi.Input<string>;
        FunctionWeight: pulumi.Input<number>;
    }
    
    export interface EventSourceMappingAttributes {
    }
    
    export interface EventSourceMappingProperties {
        BatchSize?: pulumi.Input<number>;
        Enabled?: pulumi.Input<boolean>;
        EventSourceArn: pulumi.Input<string>;
        FunctionName: pulumi.Input<string>;
        MaximumBatchingWindowInSeconds?: pulumi.Input<number>;
        StartingPosition?: pulumi.Input<string>;
    }
    
    export interface FunctionAttributes {
        Arn: string;
    }
    
    export interface FunctionCodeAttributes {
        S3Bucket?: string;
        S3Key?: string;
        S3ObjectVersion?: string;
        ZipFile?: string;
    }
    
    export interface FunctionCodeProperties {
        S3Bucket?: pulumi.Input<string>;
        S3Key?: pulumi.Input<string>;
        S3ObjectVersion?: pulumi.Input<string>;
        ZipFile?: pulumi.Input<string>;
    }
    
    export interface FunctionDeadLetterConfigAttributes {
        TargetArn?: string;
    }
    
    export interface FunctionDeadLetterConfigProperties {
        TargetArn?: pulumi.Input<string>;
    }
    
    export interface FunctionEnvironmentAttributes {
        Variables?: {[key: string]: string};
    }
    
    export interface FunctionEnvironmentProperties {
        Variables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    }
    
    export interface FunctionProperties {
        Code: pulumi.Input<FunctionCodeProperties>;
        DeadLetterConfig?: pulumi.Input<FunctionDeadLetterConfigProperties>;
        Description?: pulumi.Input<string>;
        Environment?: pulumi.Input<FunctionEnvironmentProperties>;
        FunctionName?: pulumi.Input<string>;
        Handler: pulumi.Input<string>;
        KmsKeyArn?: pulumi.Input<string>;
        Layers?: pulumi.Input<pulumi.Input<string>[]>;
        MemorySize?: pulumi.Input<number>;
        ReservedConcurrentExecutions?: pulumi.Input<number>;
        Role: pulumi.Input<string>;
        Runtime: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        Timeout?: pulumi.Input<number>;
        TracingConfig?: pulumi.Input<FunctionTracingConfigProperties>;
        VpcConfig?: pulumi.Input<FunctionVpcConfigProperties>;
    }
    
    export interface FunctionTracingConfigAttributes {
        Mode?: string;
    }
    
    export interface FunctionTracingConfigProperties {
        Mode?: pulumi.Input<string>;
    }
    
    export interface FunctionVpcConfigAttributes {
        SecurityGroupIds: string[];
        SubnetIds: string[];
    }
    
    export interface FunctionVpcConfigProperties {
        SecurityGroupIds: pulumi.Input<pulumi.Input<string>[]>;
        SubnetIds: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface LayerVersionAttributes {
    }
    
    export interface LayerVersionContentAttributes {
        S3Bucket: string;
        S3Key: string;
        S3ObjectVersion?: string;
    }
    
    export interface LayerVersionContentProperties {
        S3Bucket: pulumi.Input<string>;
        S3Key: pulumi.Input<string>;
        S3ObjectVersion?: pulumi.Input<string>;
    }
    
    export interface LayerVersionPermissionAttributes {
    }
    
    export interface LayerVersionPermissionProperties {
        Action: pulumi.Input<string>;
        LayerVersionArn: pulumi.Input<string>;
        OrganizationId?: pulumi.Input<string>;
        Principal: pulumi.Input<string>;
    }
    
    export interface LayerVersionProperties {
        CompatibleRuntimes?: pulumi.Input<pulumi.Input<string>[]>;
        Content: pulumi.Input<LayerVersionContentProperties>;
        Description?: pulumi.Input<string>;
        LayerName?: pulumi.Input<string>;
        LicenseInfo?: pulumi.Input<string>;
    }
    
    export interface PermissionAttributes {
    }
    
    export interface PermissionProperties {
        Action: pulumi.Input<string>;
        EventSourceToken?: pulumi.Input<string>;
        FunctionName: pulumi.Input<string>;
        Principal: pulumi.Input<string>;
        SourceAccount?: pulumi.Input<string>;
        SourceArn?: pulumi.Input<string>;
    }
    
    export interface VersionAttributes {
        Version: string;
    }
    
    export interface VersionProperties {
        CodeSha256?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        FunctionName: pulumi.Input<string>;
    }
    
    
    export class Alias extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AliasAttributes>;

        constructor(name: string, properties: AliasProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Lambda:Alias", name, inputs, opts);
        }
    }
    
    export class EventSourceMapping extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<EventSourceMappingAttributes>;

        constructor(name: string, properties: EventSourceMappingProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Lambda:EventSourceMapping", name, inputs, opts);
        }
    }
    
    export class Function extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<FunctionAttributes>;

        constructor(name: string, properties: FunctionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Lambda:Function", name, inputs, opts);
        }
    }
    
    export class LayerVersion extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<LayerVersionAttributes>;

        constructor(name: string, properties: LayerVersionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Lambda:LayerVersion", name, inputs, opts);
        }
    }
    
    export class LayerVersionPermission extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<LayerVersionPermissionAttributes>;

        constructor(name: string, properties: LayerVersionPermissionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Lambda:LayerVersionPermission", name, inputs, opts);
        }
    }
    
    export class Permission extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PermissionAttributes>;

        constructor(name: string, properties: PermissionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Lambda:Permission", name, inputs, opts);
        }
    }
    
    export class Version extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VersionAttributes>;

        constructor(name: string, properties: VersionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Lambda:Version", name, inputs, opts);
        }
    }
    
}

export namespace logs {
    
    export interface DestinationAttributes {
        Arn: string;
    }
    
    export interface DestinationProperties {
        DestinationName: pulumi.Input<string>;
        DestinationPolicy: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
        TargetArn: pulumi.Input<string>;
    }
    
    export interface LogGroupAttributes {
        Arn: string;
    }
    
    export interface LogGroupProperties {
        LogGroupName?: pulumi.Input<string>;
        RetentionInDays?: pulumi.Input<number>;
    }
    
    export interface LogStreamAttributes {
    }
    
    export interface LogStreamProperties {
        LogGroupName: pulumi.Input<string>;
        LogStreamName?: pulumi.Input<string>;
    }
    
    export interface MetricFilterAttributes {
    }
    
    export interface MetricFilterMetricTransformationAttributes {
        DefaultValue?: number;
        MetricName: string;
        MetricNamespace: string;
        MetricValue: string;
    }
    
    export interface MetricFilterMetricTransformationProperties {
        DefaultValue?: pulumi.Input<number>;
        MetricName: pulumi.Input<string>;
        MetricNamespace: pulumi.Input<string>;
        MetricValue: pulumi.Input<string>;
    }
    
    export interface MetricFilterProperties {
        FilterPattern: pulumi.Input<string>;
        LogGroupName: pulumi.Input<string>;
        MetricTransformations: pulumi.Input<pulumi.Input<MetricFilterMetricTransformationProperties>[]>;
    }
    
    export interface SubscriptionFilterAttributes {
    }
    
    export interface SubscriptionFilterProperties {
        DestinationArn: pulumi.Input<string>;
        FilterPattern: pulumi.Input<string>;
        LogGroupName: pulumi.Input<string>;
        RoleArn?: pulumi.Input<string>;
    }
    
    
    export class Destination extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DestinationAttributes>;

        constructor(name: string, properties: DestinationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Logs:Destination", name, inputs, opts);
        }
    }
    
    export class LogGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<LogGroupAttributes>;

        constructor(name: string, properties?: LogGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Logs:LogGroup", name, inputs, opts);
        }
    }
    
    export class LogStream extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<LogStreamAttributes>;

        constructor(name: string, properties: LogStreamProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Logs:LogStream", name, inputs, opts);
        }
    }
    
    export class MetricFilter extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<MetricFilterAttributes>;

        constructor(name: string, properties: MetricFilterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Logs:MetricFilter", name, inputs, opts);
        }
    }
    
    export class SubscriptionFilter extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SubscriptionFilterAttributes>;

        constructor(name: string, properties: SubscriptionFilterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Logs:SubscriptionFilter", name, inputs, opts);
        }
    }
    
}

export namespace msk {
    
    export interface ClusterAttributes {
    }
    
    export interface ClusterBrokerNodeGroupInfoAttributes {
        BrokerAZDistribution?: string;
        ClientSubnets: string[];
        InstanceType: string;
        SecurityGroups?: string[];
        StorageInfo?: ClusterStorageInfoAttributes;
    }
    
    export interface ClusterBrokerNodeGroupInfoProperties {
        BrokerAZDistribution?: pulumi.Input<string>;
        ClientSubnets: pulumi.Input<pulumi.Input<string>[]>;
        InstanceType: pulumi.Input<string>;
        SecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
        StorageInfo?: pulumi.Input<ClusterStorageInfoProperties>;
    }
    
    export interface ClusterClientAuthenticationAttributes {
        Tls?: ClusterTlsAttributes;
    }
    
    export interface ClusterClientAuthenticationProperties {
        Tls?: pulumi.Input<ClusterTlsProperties>;
    }
    
    export interface ClusterConfigurationInfoAttributes {
        Arn: string;
        Revision: number;
    }
    
    export interface ClusterConfigurationInfoProperties {
        Arn: pulumi.Input<string>;
        Revision: pulumi.Input<number>;
    }
    
    export interface ClusterEBSStorageInfoAttributes {
        VolumeSize?: number;
    }
    
    export interface ClusterEBSStorageInfoProperties {
        VolumeSize?: pulumi.Input<number>;
    }
    
    export interface ClusterEncryptionAtRestAttributes {
        DataVolumeKMSKeyId: string;
    }
    
    export interface ClusterEncryptionAtRestProperties {
        DataVolumeKMSKeyId: pulumi.Input<string>;
    }
    
    export interface ClusterEncryptionInTransitAttributes {
        ClientBroker?: string;
        InCluster?: boolean;
    }
    
    export interface ClusterEncryptionInTransitProperties {
        ClientBroker?: pulumi.Input<string>;
        InCluster?: pulumi.Input<boolean>;
    }
    
    export interface ClusterEncryptionInfoAttributes {
        EncryptionAtRest?: ClusterEncryptionAtRestAttributes;
        EncryptionInTransit?: ClusterEncryptionInTransitAttributes;
    }
    
    export interface ClusterEncryptionInfoProperties {
        EncryptionAtRest?: pulumi.Input<ClusterEncryptionAtRestProperties>;
        EncryptionInTransit?: pulumi.Input<ClusterEncryptionInTransitProperties>;
    }
    
    export interface ClusterProperties {
        BrokerNodeGroupInfo: pulumi.Input<ClusterBrokerNodeGroupInfoProperties>;
        ClientAuthentication?: pulumi.Input<ClusterClientAuthenticationProperties>;
        ClusterName: pulumi.Input<string>;
        ConfigurationInfo?: pulumi.Input<ClusterConfigurationInfoProperties>;
        EncryptionInfo?: pulumi.Input<ClusterEncryptionInfoProperties>;
        EnhancedMonitoring?: pulumi.Input<string>;
        KafkaVersion: pulumi.Input<string>;
        NumberOfBrokerNodes: pulumi.Input<number>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface ClusterStorageInfoAttributes {
        EBSStorageInfo?: ClusterEBSStorageInfoAttributes;
    }
    
    export interface ClusterStorageInfoProperties {
        EBSStorageInfo?: pulumi.Input<ClusterEBSStorageInfoProperties>;
    }
    
    export interface ClusterTlsAttributes {
        CertificateAuthorityArnList?: string[];
    }
    
    export interface ClusterTlsProperties {
        CertificateAuthorityArnList?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    
    export class Cluster extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ClusterAttributes>;

        constructor(name: string, properties: ClusterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:MSK:Cluster", name, inputs, opts);
        }
    }
    
}

export namespace medialive {
    
    export interface ChannelAribSourceSettingsAttributes {
    }
    
    export interface ChannelAribSourceSettingsProperties {
    }
    
    export interface ChannelAttributes {
        Arn: string;
        Inputs: string[];
    }
    
    export interface ChannelAudioLanguageSelectionAttributes {
        LanguageCode?: string;
        LanguageSelectionPolicy?: string;
    }
    
    export interface ChannelAudioLanguageSelectionProperties {
        LanguageCode?: pulumi.Input<string>;
        LanguageSelectionPolicy?: pulumi.Input<string>;
    }
    
    export interface ChannelAudioPidSelectionAttributes {
        Pid?: number;
    }
    
    export interface ChannelAudioPidSelectionProperties {
        Pid?: pulumi.Input<number>;
    }
    
    export interface ChannelAudioSelectorAttributes {
        Name?: string;
        SelectorSettings?: ChannelAudioSelectorSettingsAttributes;
    }
    
    export interface ChannelAudioSelectorProperties {
        Name?: pulumi.Input<string>;
        SelectorSettings?: pulumi.Input<ChannelAudioSelectorSettingsProperties>;
    }
    
    export interface ChannelAudioSelectorSettingsAttributes {
        AudioLanguageSelection?: ChannelAudioLanguageSelectionAttributes;
        AudioPidSelection?: ChannelAudioPidSelectionAttributes;
    }
    
    export interface ChannelAudioSelectorSettingsProperties {
        AudioLanguageSelection?: pulumi.Input<ChannelAudioLanguageSelectionProperties>;
        AudioPidSelection?: pulumi.Input<ChannelAudioPidSelectionProperties>;
    }
    
    export interface ChannelCaptionSelectorAttributes {
        LanguageCode?: string;
        Name?: string;
        SelectorSettings?: ChannelCaptionSelectorSettingsAttributes;
    }
    
    export interface ChannelCaptionSelectorProperties {
        LanguageCode?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        SelectorSettings?: pulumi.Input<ChannelCaptionSelectorSettingsProperties>;
    }
    
    export interface ChannelCaptionSelectorSettingsAttributes {
        AribSourceSettings?: ChannelAribSourceSettingsAttributes;
        DvbSubSourceSettings?: ChannelDvbSubSourceSettingsAttributes;
        EmbeddedSourceSettings?: ChannelEmbeddedSourceSettingsAttributes;
        Scte20SourceSettings?: ChannelScte20SourceSettingsAttributes;
        Scte27SourceSettings?: ChannelScte27SourceSettingsAttributes;
        TeletextSourceSettings?: ChannelTeletextSourceSettingsAttributes;
    }
    
    export interface ChannelCaptionSelectorSettingsProperties {
        AribSourceSettings?: pulumi.Input<ChannelAribSourceSettingsProperties>;
        DvbSubSourceSettings?: pulumi.Input<ChannelDvbSubSourceSettingsProperties>;
        EmbeddedSourceSettings?: pulumi.Input<ChannelEmbeddedSourceSettingsProperties>;
        Scte20SourceSettings?: pulumi.Input<ChannelScte20SourceSettingsProperties>;
        Scte27SourceSettings?: pulumi.Input<ChannelScte27SourceSettingsProperties>;
        TeletextSourceSettings?: pulumi.Input<ChannelTeletextSourceSettingsProperties>;
    }
    
    export interface ChannelDvbSubSourceSettingsAttributes {
        Pid?: number;
    }
    
    export interface ChannelDvbSubSourceSettingsProperties {
        Pid?: pulumi.Input<number>;
    }
    
    export interface ChannelEmbeddedSourceSettingsAttributes {
        Convert608To708?: string;
        Scte20Detection?: string;
        Source608ChannelNumber?: number;
        Source608TrackNumber?: number;
    }
    
    export interface ChannelEmbeddedSourceSettingsProperties {
        Convert608To708?: pulumi.Input<string>;
        Scte20Detection?: pulumi.Input<string>;
        Source608ChannelNumber?: pulumi.Input<number>;
        Source608TrackNumber?: pulumi.Input<number>;
    }
    
    export interface ChannelHlsInputSettingsAttributes {
        Bandwidth?: number;
        BufferSegments?: number;
        Retries?: number;
        RetryInterval?: number;
    }
    
    export interface ChannelHlsInputSettingsProperties {
        Bandwidth?: pulumi.Input<number>;
        BufferSegments?: pulumi.Input<number>;
        Retries?: pulumi.Input<number>;
        RetryInterval?: pulumi.Input<number>;
    }
    
    export interface ChannelInputAttachmentAttributes {
        InputAttachmentName?: string;
        InputId?: string;
        InputSettings?: ChannelInputSettingsAttributes;
    }
    
    export interface ChannelInputAttachmentProperties {
        InputAttachmentName?: pulumi.Input<string>;
        InputId?: pulumi.Input<string>;
        InputSettings?: pulumi.Input<ChannelInputSettingsProperties>;
    }
    
    export interface ChannelInputSettingsAttributes {
        AudioSelectors?: ChannelAudioSelectorAttributes[];
        CaptionSelectors?: ChannelCaptionSelectorAttributes[];
        DeblockFilter?: string;
        DenoiseFilter?: string;
        FilterStrength?: number;
        InputFilter?: string;
        NetworkInputSettings?: ChannelNetworkInputSettingsAttributes;
        SourceEndBehavior?: string;
        VideoSelector?: ChannelVideoSelectorAttributes;
    }
    
    export interface ChannelInputSettingsProperties {
        AudioSelectors?: pulumi.Input<pulumi.Input<ChannelAudioSelectorProperties>[]>;
        CaptionSelectors?: pulumi.Input<pulumi.Input<ChannelCaptionSelectorProperties>[]>;
        DeblockFilter?: pulumi.Input<string>;
        DenoiseFilter?: pulumi.Input<string>;
        FilterStrength?: pulumi.Input<number>;
        InputFilter?: pulumi.Input<string>;
        NetworkInputSettings?: pulumi.Input<ChannelNetworkInputSettingsProperties>;
        SourceEndBehavior?: pulumi.Input<string>;
        VideoSelector?: pulumi.Input<ChannelVideoSelectorProperties>;
    }
    
    export interface ChannelInputSpecificationAttributes {
        Codec?: string;
        MaximumBitrate?: string;
        Resolution?: string;
    }
    
    export interface ChannelInputSpecificationProperties {
        Codec?: pulumi.Input<string>;
        MaximumBitrate?: pulumi.Input<string>;
        Resolution?: pulumi.Input<string>;
    }
    
    export interface ChannelMediaPackageOutputDestinationSettingsAttributes {
        ChannelId?: string;
    }
    
    export interface ChannelMediaPackageOutputDestinationSettingsProperties {
        ChannelId?: pulumi.Input<string>;
    }
    
    export interface ChannelNetworkInputSettingsAttributes {
        HlsInputSettings?: ChannelHlsInputSettingsAttributes;
        ServerValidation?: string;
    }
    
    export interface ChannelNetworkInputSettingsProperties {
        HlsInputSettings?: pulumi.Input<ChannelHlsInputSettingsProperties>;
        ServerValidation?: pulumi.Input<string>;
    }
    
    export interface ChannelOutputDestinationAttributes {
        Id?: string;
        MediaPackageSettings?: ChannelMediaPackageOutputDestinationSettingsAttributes[];
        Settings?: ChannelOutputDestinationSettingsAttributes[];
    }
    
    export interface ChannelOutputDestinationProperties {
        Id?: pulumi.Input<string>;
        MediaPackageSettings?: pulumi.Input<pulumi.Input<ChannelMediaPackageOutputDestinationSettingsProperties>[]>;
        Settings?: pulumi.Input<pulumi.Input<ChannelOutputDestinationSettingsProperties>[]>;
    }
    
    export interface ChannelOutputDestinationSettingsAttributes {
        PasswordParam?: string;
        StreamName?: string;
        Url?: string;
        Username?: string;
    }
    
    export interface ChannelOutputDestinationSettingsProperties {
        PasswordParam?: pulumi.Input<string>;
        StreamName?: pulumi.Input<string>;
        Url?: pulumi.Input<string>;
        Username?: pulumi.Input<string>;
    }
    
    export interface ChannelProperties {
        ChannelClass?: pulumi.Input<string>;
        Destinations?: pulumi.Input<pulumi.Input<ChannelOutputDestinationProperties>[]>;
        EncoderSettings?: pulumi.Input<string>;
        InputAttachments?: pulumi.Input<pulumi.Input<ChannelInputAttachmentProperties>[]>;
        InputSpecification?: pulumi.Input<ChannelInputSpecificationProperties>;
        LogLevel?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        RoleArn?: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface ChannelScte20SourceSettingsAttributes {
        Convert608To708?: string;
        Source608ChannelNumber?: number;
    }
    
    export interface ChannelScte20SourceSettingsProperties {
        Convert608To708?: pulumi.Input<string>;
        Source608ChannelNumber?: pulumi.Input<number>;
    }
    
    export interface ChannelScte27SourceSettingsAttributes {
        Pid?: number;
    }
    
    export interface ChannelScte27SourceSettingsProperties {
        Pid?: pulumi.Input<number>;
    }
    
    export interface ChannelTeletextSourceSettingsAttributes {
        PageNumber?: string;
    }
    
    export interface ChannelTeletextSourceSettingsProperties {
        PageNumber?: pulumi.Input<string>;
    }
    
    export interface ChannelVideoSelectorAttributes {
        ColorSpace?: string;
        ColorSpaceUsage?: string;
        SelectorSettings?: ChannelVideoSelectorSettingsAttributes;
    }
    
    export interface ChannelVideoSelectorPidAttributes {
        Pid?: number;
    }
    
    export interface ChannelVideoSelectorPidProperties {
        Pid?: pulumi.Input<number>;
    }
    
    export interface ChannelVideoSelectorProgramIdAttributes {
        ProgramId?: number;
    }
    
    export interface ChannelVideoSelectorProgramIdProperties {
        ProgramId?: pulumi.Input<number>;
    }
    
    export interface ChannelVideoSelectorProperties {
        ColorSpace?: pulumi.Input<string>;
        ColorSpaceUsage?: pulumi.Input<string>;
        SelectorSettings?: pulumi.Input<ChannelVideoSelectorSettingsProperties>;
    }
    
    export interface ChannelVideoSelectorSettingsAttributes {
        VideoSelectorPid?: ChannelVideoSelectorPidAttributes;
        VideoSelectorProgramId?: ChannelVideoSelectorProgramIdAttributes;
    }
    
    export interface ChannelVideoSelectorSettingsProperties {
        VideoSelectorPid?: pulumi.Input<ChannelVideoSelectorPidProperties>;
        VideoSelectorProgramId?: pulumi.Input<ChannelVideoSelectorProgramIdProperties>;
    }
    
    export interface InputAttributes {
        Arn: string;
        Destinations: string[];
        Sources: string[];
    }
    
    export interface InputInputDestinationRequestAttributes {
        StreamName?: string;
    }
    
    export interface InputInputDestinationRequestProperties {
        StreamName?: pulumi.Input<string>;
    }
    
    export interface InputInputSourceRequestAttributes {
        PasswordParam?: string;
        Url?: string;
        Username?: string;
    }
    
    export interface InputInputSourceRequestProperties {
        PasswordParam?: pulumi.Input<string>;
        Url?: pulumi.Input<string>;
        Username?: pulumi.Input<string>;
    }
    
    export interface InputInputVpcRequestAttributes {
        SecurityGroupIds?: string[];
        SubnetIds?: string[];
    }
    
    export interface InputInputVpcRequestProperties {
        SecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
        SubnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface InputMediaConnectFlowRequestAttributes {
        FlowArn?: string;
    }
    
    export interface InputMediaConnectFlowRequestProperties {
        FlowArn?: pulumi.Input<string>;
    }
    
    export interface InputProperties {
        Destinations?: pulumi.Input<pulumi.Input<InputInputDestinationRequestProperties>[]>;
        InputSecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
        MediaConnectFlows?: pulumi.Input<pulumi.Input<InputMediaConnectFlowRequestProperties>[]>;
        Name?: pulumi.Input<string>;
        RoleArn?: pulumi.Input<string>;
        Sources?: pulumi.Input<pulumi.Input<InputInputSourceRequestProperties>[]>;
        Tags?: pulumi.Input<string>;
        Type?: pulumi.Input<string>;
        Vpc?: pulumi.Input<InputInputVpcRequestProperties>;
    }
    
    export interface InputSecurityGroupAttributes {
        Arn: string;
    }
    
    export interface InputSecurityGroupInputWhitelistRuleCidrAttributes {
        Cidr?: string;
    }
    
    export interface InputSecurityGroupInputWhitelistRuleCidrProperties {
        Cidr?: pulumi.Input<string>;
    }
    
    export interface InputSecurityGroupProperties {
        Tags?: pulumi.Input<string>;
        WhitelistRules?: pulumi.Input<pulumi.Input<InputSecurityGroupInputWhitelistRuleCidrProperties>[]>;
    }
    
    
    export class Channel extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ChannelAttributes>;

        constructor(name: string, properties?: ChannelProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:MediaLive:Channel", name, inputs, opts);
        }
    }
    
    export class Input extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<InputAttributes>;

        constructor(name: string, properties?: InputProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:MediaLive:Input", name, inputs, opts);
        }
    }
    
    export class InputSecurityGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<InputSecurityGroupAttributes>;

        constructor(name: string, properties?: InputSecurityGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:MediaLive:InputSecurityGroup", name, inputs, opts);
        }
    }
    
}

export namespace mediastore {
    
    export interface ContainerAttributes {
        Endpoint: string;
    }
    
    export interface ContainerCorsRuleAttributes {
        AllowedHeaders?: string[];
        AllowedMethods?: string[];
        AllowedOrigins?: string[];
        ExposeHeaders?: string[];
        MaxAgeSeconds?: number;
    }
    
    export interface ContainerCorsRuleProperties {
        AllowedHeaders?: pulumi.Input<pulumi.Input<string>[]>;
        AllowedMethods?: pulumi.Input<pulumi.Input<string>[]>;
        AllowedOrigins?: pulumi.Input<pulumi.Input<string>[]>;
        ExposeHeaders?: pulumi.Input<pulumi.Input<string>[]>;
        MaxAgeSeconds?: pulumi.Input<number>;
    }
    
    export interface ContainerProperties {
        AccessLoggingEnabled?: pulumi.Input<boolean>;
        ContainerName: pulumi.Input<string>;
        CorsPolicy?: pulumi.Input<pulumi.Input<ContainerCorsRuleProperties>[]>;
        LifecyclePolicy?: pulumi.Input<string>;
        Policy?: pulumi.Input<string>;
    }
    
    
    export class Container extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ContainerAttributes>;

        constructor(name: string, properties: ContainerProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:MediaStore:Container", name, inputs, opts);
        }
    }
    
}

export namespace neptune {
    
    export interface DBClusterAttributes {
        ClusterResourceId: string;
        Endpoint: string;
        Port: string;
        ReadEndpoint: string;
    }
    
    export interface DBClusterParameterGroupAttributes {
    }
    
    export interface DBClusterParameterGroupProperties {
        Description: pulumi.Input<string>;
        Family: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Parameters: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface DBClusterProperties {
        AvailabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
        BackupRetentionPeriod?: pulumi.Input<number>;
        DBClusterIdentifier?: pulumi.Input<string>;
        DBClusterParameterGroupName?: pulumi.Input<string>;
        DBSubnetGroupName?: pulumi.Input<string>;
        EnableCloudwatchLogsExports?: pulumi.Input<pulumi.Input<string>[]>;
        IamAuthEnabled?: pulumi.Input<boolean>;
        KmsKeyId?: pulumi.Input<string>;
        Port?: pulumi.Input<number>;
        PreferredBackupWindow?: pulumi.Input<string>;
        PreferredMaintenanceWindow?: pulumi.Input<string>;
        SnapshotIdentifier?: pulumi.Input<string>;
        StorageEncrypted?: pulumi.Input<boolean>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface DBInstanceAttributes {
        Endpoint: string;
        Port: string;
    }
    
    export interface DBInstanceProperties {
        AllowMajorVersionUpgrade?: pulumi.Input<boolean>;
        AutoMinorVersionUpgrade?: pulumi.Input<boolean>;
        AvailabilityZone?: pulumi.Input<string>;
        DBClusterIdentifier?: pulumi.Input<string>;
        DBInstanceClass: pulumi.Input<string>;
        DBInstanceIdentifier?: pulumi.Input<string>;
        DBParameterGroupName?: pulumi.Input<string>;
        DBSnapshotIdentifier?: pulumi.Input<string>;
        DBSubnetGroupName?: pulumi.Input<string>;
        PreferredMaintenanceWindow?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface DBParameterGroupAttributes {
    }
    
    export interface DBParameterGroupProperties {
        Description: pulumi.Input<string>;
        Family: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Parameters: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface DBSubnetGroupAttributes {
    }
    
    export interface DBSubnetGroupProperties {
        DBSubnetGroupDescription: pulumi.Input<string>;
        DBSubnetGroupName?: pulumi.Input<string>;
        SubnetIds: pulumi.Input<pulumi.Input<string>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    
    export class DBCluster extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DBClusterAttributes>;

        constructor(name: string, properties?: DBClusterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Neptune:DBCluster", name, inputs, opts);
        }
    }
    
    export class DBClusterParameterGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DBClusterParameterGroupAttributes>;

        constructor(name: string, properties: DBClusterParameterGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Neptune:DBClusterParameterGroup", name, inputs, opts);
        }
    }
    
    export class DBInstance extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DBInstanceAttributes>;

        constructor(name: string, properties: DBInstanceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Neptune:DBInstance", name, inputs, opts);
        }
    }
    
    export class DBParameterGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DBParameterGroupAttributes>;

        constructor(name: string, properties: DBParameterGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Neptune:DBParameterGroup", name, inputs, opts);
        }
    }
    
    export class DBSubnetGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DBSubnetGroupAttributes>;

        constructor(name: string, properties: DBSubnetGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Neptune:DBSubnetGroup", name, inputs, opts);
        }
    }
    
}

export namespace opsworks {
    
    export interface AppAttributes {
    }
    
    export interface AppDataSourceAttributes {
        Arn?: string;
        DatabaseName?: string;
        Type?: string;
    }
    
    export interface AppDataSourceProperties {
        Arn?: pulumi.Input<string>;
        DatabaseName?: pulumi.Input<string>;
        Type?: pulumi.Input<string>;
    }
    
    export interface AppEnvironmentVariableAttributes {
        Key: string;
        Secure?: boolean;
        Value: string;
    }
    
    export interface AppEnvironmentVariableProperties {
        Key: pulumi.Input<string>;
        Secure?: pulumi.Input<boolean>;
        Value: pulumi.Input<string>;
    }
    
    export interface AppProperties {
        AppSource?: pulumi.Input<AppSourceProperties>;
        Attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        DataSources?: pulumi.Input<pulumi.Input<AppDataSourceProperties>[]>;
        Description?: pulumi.Input<string>;
        Domains?: pulumi.Input<pulumi.Input<string>[]>;
        EnableSsl?: pulumi.Input<boolean>;
        Environment?: pulumi.Input<pulumi.Input<AppEnvironmentVariableProperties>[]>;
        Name: pulumi.Input<string>;
        Shortname?: pulumi.Input<string>;
        SslConfiguration?: pulumi.Input<AppSslConfigurationProperties>;
        StackId: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface AppSourceAttributes {
        Password?: string;
        Revision?: string;
        SshKey?: string;
        Type?: string;
        Url?: string;
        Username?: string;
    }
    
    export interface AppSourceProperties {
        Password?: pulumi.Input<string>;
        Revision?: pulumi.Input<string>;
        SshKey?: pulumi.Input<string>;
        Type?: pulumi.Input<string>;
        Url?: pulumi.Input<string>;
        Username?: pulumi.Input<string>;
    }
    
    export interface AppSslConfigurationAttributes {
        Certificate?: string;
        Chain?: string;
        PrivateKey?: string;
    }
    
    export interface AppSslConfigurationProperties {
        Certificate?: pulumi.Input<string>;
        Chain?: pulumi.Input<string>;
        PrivateKey?: pulumi.Input<string>;
    }
    
    export interface ElasticLoadBalancerAttachmentAttributes {
    }
    
    export interface ElasticLoadBalancerAttachmentProperties {
        ElasticLoadBalancerName: pulumi.Input<string>;
        LayerId: pulumi.Input<string>;
    }
    
    export interface InstanceAttributes {
        AvailabilityZone: string;
        PrivateDnsName: string;
        PrivateIp: string;
        PublicDnsName: string;
        PublicIp: string;
    }
    
    export interface InstanceBlockDeviceMappingAttributes {
        DeviceName?: string;
        Ebs?: InstanceEbsBlockDeviceAttributes;
        NoDevice?: string;
        VirtualName?: string;
    }
    
    export interface InstanceBlockDeviceMappingProperties {
        DeviceName?: pulumi.Input<string>;
        Ebs?: pulumi.Input<InstanceEbsBlockDeviceProperties>;
        NoDevice?: pulumi.Input<string>;
        VirtualName?: pulumi.Input<string>;
    }
    
    export interface InstanceEbsBlockDeviceAttributes {
        DeleteOnTermination?: boolean;
        Iops?: number;
        SnapshotId?: string;
        VolumeSize?: number;
        VolumeType?: string;
    }
    
    export interface InstanceEbsBlockDeviceProperties {
        DeleteOnTermination?: pulumi.Input<boolean>;
        Iops?: pulumi.Input<number>;
        SnapshotId?: pulumi.Input<string>;
        VolumeSize?: pulumi.Input<number>;
        VolumeType?: pulumi.Input<string>;
    }
    
    export interface InstanceProperties {
        AgentVersion?: pulumi.Input<string>;
        AmiId?: pulumi.Input<string>;
        Architecture?: pulumi.Input<string>;
        AutoScalingType?: pulumi.Input<string>;
        AvailabilityZone?: pulumi.Input<string>;
        BlockDeviceMappings?: pulumi.Input<pulumi.Input<InstanceBlockDeviceMappingProperties>[]>;
        EbsOptimized?: pulumi.Input<boolean>;
        ElasticIps?: pulumi.Input<pulumi.Input<string>[]>;
        Hostname?: pulumi.Input<string>;
        InstallUpdatesOnBoot?: pulumi.Input<boolean>;
        InstanceType: pulumi.Input<string>;
        LayerIds: pulumi.Input<pulumi.Input<string>[]>;
        Os?: pulumi.Input<string>;
        RootDeviceType?: pulumi.Input<string>;
        SshKeyName?: pulumi.Input<string>;
        StackId: pulumi.Input<string>;
        SubnetId?: pulumi.Input<string>;
        Tenancy?: pulumi.Input<string>;
        TimeBasedAutoScaling?: pulumi.Input<InstanceTimeBasedAutoScalingProperties>;
        VirtualizationType?: pulumi.Input<string>;
        Volumes?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface InstanceTimeBasedAutoScalingAttributes {
        Friday?: {[key: string]: string};
        Monday?: {[key: string]: string};
        Saturday?: {[key: string]: string};
        Sunday?: {[key: string]: string};
        Thursday?: {[key: string]: string};
        Tuesday?: {[key: string]: string};
        Wednesday?: {[key: string]: string};
    }
    
    export interface InstanceTimeBasedAutoScalingProperties {
        Friday?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        Monday?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        Saturday?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        Sunday?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        Thursday?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        Tuesday?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        Wednesday?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    }
    
    export interface LayerAttributes {
    }
    
    export interface LayerAutoScalingThresholdsAttributes {
        CpuThreshold?: number;
        IgnoreMetricsTime?: number;
        InstanceCount?: number;
        LoadThreshold?: number;
        MemoryThreshold?: number;
        ThresholdsWaitTime?: number;
    }
    
    export interface LayerAutoScalingThresholdsProperties {
        CpuThreshold?: pulumi.Input<number>;
        IgnoreMetricsTime?: pulumi.Input<number>;
        InstanceCount?: pulumi.Input<number>;
        LoadThreshold?: pulumi.Input<number>;
        MemoryThreshold?: pulumi.Input<number>;
        ThresholdsWaitTime?: pulumi.Input<number>;
    }
    
    export interface LayerLifecycleEventConfigurationAttributes {
        ShutdownEventConfiguration?: LayerShutdownEventConfigurationAttributes;
    }
    
    export interface LayerLifecycleEventConfigurationProperties {
        ShutdownEventConfiguration?: pulumi.Input<LayerShutdownEventConfigurationProperties>;
    }
    
    export interface LayerLoadBasedAutoScalingAttributes {
        DownScaling?: LayerAutoScalingThresholdsAttributes;
        Enable?: boolean;
        UpScaling?: LayerAutoScalingThresholdsAttributes;
    }
    
    export interface LayerLoadBasedAutoScalingProperties {
        DownScaling?: pulumi.Input<LayerAutoScalingThresholdsProperties>;
        Enable?: pulumi.Input<boolean>;
        UpScaling?: pulumi.Input<LayerAutoScalingThresholdsProperties>;
    }
    
    export interface LayerProperties {
        Attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        AutoAssignElasticIps: pulumi.Input<boolean>;
        AutoAssignPublicIps: pulumi.Input<boolean>;
        CustomInstanceProfileArn?: pulumi.Input<string>;
        CustomJson?: pulumi.Input<string>;
        CustomRecipes?: pulumi.Input<LayerRecipesProperties>;
        CustomSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
        EnableAutoHealing: pulumi.Input<boolean>;
        InstallUpdatesOnBoot?: pulumi.Input<boolean>;
        LifecycleEventConfiguration?: pulumi.Input<LayerLifecycleEventConfigurationProperties>;
        LoadBasedAutoScaling?: pulumi.Input<LayerLoadBasedAutoScalingProperties>;
        Name: pulumi.Input<string>;
        Packages?: pulumi.Input<pulumi.Input<string>[]>;
        Shortname: pulumi.Input<string>;
        StackId: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        Type: pulumi.Input<string>;
        UseEbsOptimizedInstances?: pulumi.Input<boolean>;
        VolumeConfigurations?: pulumi.Input<pulumi.Input<LayerVolumeConfigurationProperties>[]>;
    }
    
    export interface LayerRecipesAttributes {
        Configure?: string[];
        Deploy?: string[];
        Setup?: string[];
        Shutdown?: string[];
        Undeploy?: string[];
    }
    
    export interface LayerRecipesProperties {
        Configure?: pulumi.Input<pulumi.Input<string>[]>;
        Deploy?: pulumi.Input<pulumi.Input<string>[]>;
        Setup?: pulumi.Input<pulumi.Input<string>[]>;
        Shutdown?: pulumi.Input<pulumi.Input<string>[]>;
        Undeploy?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface LayerShutdownEventConfigurationAttributes {
        DelayUntilElbConnectionsDrained?: boolean;
        ExecutionTimeout?: number;
    }
    
    export interface LayerShutdownEventConfigurationProperties {
        DelayUntilElbConnectionsDrained?: pulumi.Input<boolean>;
        ExecutionTimeout?: pulumi.Input<number>;
    }
    
    export interface LayerVolumeConfigurationAttributes {
        Encrypted?: boolean;
        Iops?: number;
        MountPoint?: string;
        NumberOfDisks?: number;
        RaidLevel?: number;
        Size?: number;
        VolumeType?: string;
    }
    
    export interface LayerVolumeConfigurationProperties {
        Encrypted?: pulumi.Input<boolean>;
        Iops?: pulumi.Input<number>;
        MountPoint?: pulumi.Input<string>;
        NumberOfDisks?: pulumi.Input<number>;
        RaidLevel?: pulumi.Input<number>;
        Size?: pulumi.Input<number>;
        VolumeType?: pulumi.Input<string>;
    }
    
    export interface StackAttributes {
    }
    
    export interface StackChefConfigurationAttributes {
        BerkshelfVersion?: string;
        ManageBerkshelf?: boolean;
    }
    
    export interface StackChefConfigurationProperties {
        BerkshelfVersion?: pulumi.Input<string>;
        ManageBerkshelf?: pulumi.Input<boolean>;
    }
    
    export interface StackElasticIpAttributes {
        Ip: string;
        Name?: string;
    }
    
    export interface StackElasticIpProperties {
        Ip: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
    }
    
    export interface StackProperties {
        AgentVersion?: pulumi.Input<string>;
        Attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        ChefConfiguration?: pulumi.Input<StackChefConfigurationProperties>;
        CloneAppIds?: pulumi.Input<pulumi.Input<string>[]>;
        ClonePermissions?: pulumi.Input<boolean>;
        ConfigurationManager?: pulumi.Input<StackStackConfigurationManagerProperties>;
        CustomCookbooksSource?: pulumi.Input<StackSourceProperties>;
        CustomJson?: pulumi.Input<string>;
        DefaultAvailabilityZone?: pulumi.Input<string>;
        DefaultInstanceProfileArn: pulumi.Input<string>;
        DefaultOs?: pulumi.Input<string>;
        DefaultRootDeviceType?: pulumi.Input<string>;
        DefaultSshKeyName?: pulumi.Input<string>;
        DefaultSubnetId?: pulumi.Input<string>;
        EcsClusterArn?: pulumi.Input<string>;
        ElasticIps?: pulumi.Input<pulumi.Input<StackElasticIpProperties>[]>;
        HostnameTheme?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        RdsDbInstances?: pulumi.Input<pulumi.Input<StackRdsDbInstanceProperties>[]>;
        ServiceRoleArn: pulumi.Input<string>;
        SourceStackId?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        UseCustomCookbooks?: pulumi.Input<boolean>;
        UseOpsworksSecurityGroups?: pulumi.Input<boolean>;
        VpcId?: pulumi.Input<string>;
    }
    
    export interface StackRdsDbInstanceAttributes {
        DbPassword: string;
        DbUser: string;
        RdsDbInstanceArn: string;
    }
    
    export interface StackRdsDbInstanceProperties {
        DbPassword: pulumi.Input<string>;
        DbUser: pulumi.Input<string>;
        RdsDbInstanceArn: pulumi.Input<string>;
    }
    
    export interface StackSourceAttributes {
        Password?: string;
        Revision?: string;
        SshKey?: string;
        Type?: string;
        Url?: string;
        Username?: string;
    }
    
    export interface StackSourceProperties {
        Password?: pulumi.Input<string>;
        Revision?: pulumi.Input<string>;
        SshKey?: pulumi.Input<string>;
        Type?: pulumi.Input<string>;
        Url?: pulumi.Input<string>;
        Username?: pulumi.Input<string>;
    }
    
    export interface StackStackConfigurationManagerAttributes {
        Name?: string;
        Version?: string;
    }
    
    export interface StackStackConfigurationManagerProperties {
        Name?: pulumi.Input<string>;
        Version?: pulumi.Input<string>;
    }
    
    export interface UserProfileAttributes {
        SshUsername: string;
    }
    
    export interface UserProfileProperties {
        AllowSelfManagement?: pulumi.Input<boolean>;
        IamUserArn: pulumi.Input<string>;
        SshPublicKey?: pulumi.Input<string>;
        SshUsername?: pulumi.Input<string>;
    }
    
    export interface VolumeAttributes {
    }
    
    export interface VolumeProperties {
        Ec2VolumeId: pulumi.Input<string>;
        MountPoint?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        StackId: pulumi.Input<string>;
    }
    
    
    export class App extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AppAttributes>;

        constructor(name: string, properties: AppProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:OpsWorks:App", name, inputs, opts);
        }
    }
    
    export class ElasticLoadBalancerAttachment extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ElasticLoadBalancerAttachmentAttributes>;

        constructor(name: string, properties: ElasticLoadBalancerAttachmentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:OpsWorks:ElasticLoadBalancerAttachment", name, inputs, opts);
        }
    }
    
    export class Instance extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<InstanceAttributes>;

        constructor(name: string, properties: InstanceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:OpsWorks:Instance", name, inputs, opts);
        }
    }
    
    export class Layer extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<LayerAttributes>;

        constructor(name: string, properties: LayerProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:OpsWorks:Layer", name, inputs, opts);
        }
    }
    
    export class Stack extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<StackAttributes>;

        constructor(name: string, properties: StackProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:OpsWorks:Stack", name, inputs, opts);
        }
    }
    
    export class UserProfile extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<UserProfileAttributes>;

        constructor(name: string, properties: UserProfileProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:OpsWorks:UserProfile", name, inputs, opts);
        }
    }
    
    export class Volume extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VolumeAttributes>;

        constructor(name: string, properties: VolumeProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:OpsWorks:Volume", name, inputs, opts);
        }
    }
    
}

export namespace opsworkscm {
    
    export interface ServerAttributes {
        Arn: string;
        Endpoint: string;
    }
    
    export interface ServerEngineAttributeAttributes {
        Name?: string;
        Value?: string;
    }
    
    export interface ServerEngineAttributeProperties {
        Name?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface ServerProperties {
        AssociatePublicIpAddress?: pulumi.Input<boolean>;
        BackupId?: pulumi.Input<string>;
        BackupRetentionCount?: pulumi.Input<number>;
        DisableAutomatedBackup?: pulumi.Input<boolean>;
        Engine?: pulumi.Input<string>;
        EngineAttributes?: pulumi.Input<pulumi.Input<ServerEngineAttributeProperties>[]>;
        EngineModel?: pulumi.Input<string>;
        EngineVersion?: pulumi.Input<string>;
        InstanceProfileArn: pulumi.Input<string>;
        InstanceType: pulumi.Input<string>;
        KeyPair?: pulumi.Input<string>;
        PreferredBackupWindow?: pulumi.Input<string>;
        PreferredMaintenanceWindow?: pulumi.Input<string>;
        SecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
        ServerName?: pulumi.Input<string>;
        ServiceRoleArn: pulumi.Input<string>;
        SubnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    
    export class Server extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ServerAttributes>;

        constructor(name: string, properties: ServerProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:OpsWorksCM:Server", name, inputs, opts);
        }
    }
    
}

export namespace pinpoint {
    
    export interface ADMChannelAttributes {
    }
    
    export interface ADMChannelProperties {
        ApplicationId: pulumi.Input<string>;
        ClientId: pulumi.Input<string>;
        ClientSecret: pulumi.Input<string>;
        Enabled?: pulumi.Input<boolean>;
    }
    
    export interface APNSChannelAttributes {
    }
    
    export interface APNSChannelProperties {
        ApplicationId: pulumi.Input<string>;
        BundleId?: pulumi.Input<string>;
        Certificate?: pulumi.Input<string>;
        DefaultAuthenticationMethod?: pulumi.Input<string>;
        Enabled?: pulumi.Input<boolean>;
        PrivateKey?: pulumi.Input<string>;
        TeamId?: pulumi.Input<string>;
        TokenKey?: pulumi.Input<string>;
        TokenKeyId?: pulumi.Input<string>;
    }
    
    export interface APNSSandboxChannelAttributes {
    }
    
    export interface APNSSandboxChannelProperties {
        ApplicationId: pulumi.Input<string>;
        BundleId?: pulumi.Input<string>;
        Certificate?: pulumi.Input<string>;
        DefaultAuthenticationMethod?: pulumi.Input<string>;
        Enabled?: pulumi.Input<boolean>;
        PrivateKey?: pulumi.Input<string>;
        TeamId?: pulumi.Input<string>;
        TokenKey?: pulumi.Input<string>;
        TokenKeyId?: pulumi.Input<string>;
    }
    
    export interface APNSVoipChannelAttributes {
    }
    
    export interface APNSVoipChannelProperties {
        ApplicationId: pulumi.Input<string>;
        BundleId?: pulumi.Input<string>;
        Certificate?: pulumi.Input<string>;
        DefaultAuthenticationMethod?: pulumi.Input<string>;
        Enabled?: pulumi.Input<boolean>;
        PrivateKey?: pulumi.Input<string>;
        TeamId?: pulumi.Input<string>;
        TokenKey?: pulumi.Input<string>;
        TokenKeyId?: pulumi.Input<string>;
    }
    
    export interface APNSVoipSandboxChannelAttributes {
    }
    
    export interface APNSVoipSandboxChannelProperties {
        ApplicationId: pulumi.Input<string>;
        BundleId?: pulumi.Input<string>;
        Certificate?: pulumi.Input<string>;
        DefaultAuthenticationMethod?: pulumi.Input<string>;
        Enabled?: pulumi.Input<boolean>;
        PrivateKey?: pulumi.Input<string>;
        TeamId?: pulumi.Input<string>;
        TokenKey?: pulumi.Input<string>;
        TokenKeyId?: pulumi.Input<string>;
    }
    
    export interface AppAttributes {
        Arn: string;
    }
    
    export interface AppProperties {
        Name: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface ApplicationSettingsAttributes {
    }
    
    export interface ApplicationSettingsCampaignHookAttributes {
        LambdaFunctionName?: string;
        Mode?: string;
        WebUrl?: string;
    }
    
    export interface ApplicationSettingsCampaignHookProperties {
        LambdaFunctionName?: pulumi.Input<string>;
        Mode?: pulumi.Input<string>;
        WebUrl?: pulumi.Input<string>;
    }
    
    export interface ApplicationSettingsLimitsAttributes {
        Daily?: number;
        MaximumDuration?: number;
        MessagesPerSecond?: number;
        Total?: number;
    }
    
    export interface ApplicationSettingsLimitsProperties {
        Daily?: pulumi.Input<number>;
        MaximumDuration?: pulumi.Input<number>;
        MessagesPerSecond?: pulumi.Input<number>;
        Total?: pulumi.Input<number>;
    }
    
    export interface ApplicationSettingsProperties {
        ApplicationId: pulumi.Input<string>;
        CampaignHook?: pulumi.Input<ApplicationSettingsCampaignHookProperties>;
        CloudWatchMetricsEnabled?: pulumi.Input<boolean>;
        Limits?: pulumi.Input<ApplicationSettingsLimitsProperties>;
        QuietTime?: pulumi.Input<ApplicationSettingsQuietTimeProperties>;
    }
    
    export interface ApplicationSettingsQuietTimeAttributes {
        End: string;
        Start: string;
    }
    
    export interface ApplicationSettingsQuietTimeProperties {
        End: pulumi.Input<string>;
        Start: pulumi.Input<string>;
    }
    
    export interface BaiduChannelAttributes {
    }
    
    export interface BaiduChannelProperties {
        ApiKey: pulumi.Input<string>;
        ApplicationId: pulumi.Input<string>;
        Enabled?: pulumi.Input<boolean>;
        SecretKey: pulumi.Input<string>;
    }
    
    export interface CampaignAttributeDimensionAttributes {
        AttributeType?: string;
        Values?: string[];
    }
    
    export interface CampaignAttributeDimensionProperties {
        AttributeType?: pulumi.Input<string>;
        Values?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface CampaignAttributes {
        Arn: string;
        CampaignId: string;
    }
    
    export interface CampaignCampaignEmailMessageAttributes {
        Body?: string;
        FromAddress?: string;
        HtmlBody?: string;
        Title?: string;
    }
    
    export interface CampaignCampaignEmailMessageProperties {
        Body?: pulumi.Input<string>;
        FromAddress?: pulumi.Input<string>;
        HtmlBody?: pulumi.Input<string>;
        Title?: pulumi.Input<string>;
    }
    
    export interface CampaignCampaignEventFilterAttributes {
        Dimensions?: CampaignEventDimensionsAttributes;
        FilterType?: string;
    }
    
    export interface CampaignCampaignEventFilterProperties {
        Dimensions?: pulumi.Input<CampaignEventDimensionsProperties>;
        FilterType?: pulumi.Input<string>;
    }
    
    export interface CampaignCampaignHookAttributes {
        LambdaFunctionName?: string;
        Mode?: string;
        WebUrl?: string;
    }
    
    export interface CampaignCampaignHookProperties {
        LambdaFunctionName?: pulumi.Input<string>;
        Mode?: pulumi.Input<string>;
        WebUrl?: pulumi.Input<string>;
    }
    
    export interface CampaignCampaignSmsMessageAttributes {
        Body?: string;
        MessageType?: string;
        SenderId?: string;
    }
    
    export interface CampaignCampaignSmsMessageProperties {
        Body?: pulumi.Input<string>;
        MessageType?: pulumi.Input<string>;
        SenderId?: pulumi.Input<string>;
    }
    
    export interface CampaignEventDimensionsAttributes {
        Attributes?: string;
        EventType?: CampaignSetDimensionAttributes;
        Metrics?: string;
    }
    
    export interface CampaignEventDimensionsProperties {
        Attributes?: pulumi.Input<string>;
        EventType?: pulumi.Input<CampaignSetDimensionProperties>;
        Metrics?: pulumi.Input<string>;
    }
    
    export interface CampaignLimitsAttributes {
        Daily?: number;
        MaximumDuration?: number;
        MessagesPerSecond?: number;
        Total?: number;
    }
    
    export interface CampaignLimitsProperties {
        Daily?: pulumi.Input<number>;
        MaximumDuration?: pulumi.Input<number>;
        MessagesPerSecond?: pulumi.Input<number>;
        Total?: pulumi.Input<number>;
    }
    
    export interface CampaignMessageAttributes {
        Action?: string;
        Body?: string;
        ImageIconUrl?: string;
        ImageSmallIconUrl?: string;
        ImageUrl?: string;
        JsonBody?: string;
        MediaUrl?: string;
        RawContent?: string;
        SilentPush?: boolean;
        TimeToLive?: number;
        Title?: string;
        Url?: string;
    }
    
    export interface CampaignMessageConfigurationAttributes {
        ADMMessage?: CampaignMessageAttributes;
        APNSMessage?: CampaignMessageAttributes;
        BaiduMessage?: CampaignMessageAttributes;
        DefaultMessage?: CampaignMessageAttributes;
        EmailMessage?: CampaignCampaignEmailMessageAttributes;
        GCMMessage?: CampaignMessageAttributes;
        SMSMessage?: CampaignCampaignSmsMessageAttributes;
    }
    
    export interface CampaignMessageConfigurationProperties {
        ADMMessage?: pulumi.Input<CampaignMessageProperties>;
        APNSMessage?: pulumi.Input<CampaignMessageProperties>;
        BaiduMessage?: pulumi.Input<CampaignMessageProperties>;
        DefaultMessage?: pulumi.Input<CampaignMessageProperties>;
        EmailMessage?: pulumi.Input<CampaignCampaignEmailMessageProperties>;
        GCMMessage?: pulumi.Input<CampaignMessageProperties>;
        SMSMessage?: pulumi.Input<CampaignCampaignSmsMessageProperties>;
    }
    
    export interface CampaignMessageProperties {
        Action?: pulumi.Input<string>;
        Body?: pulumi.Input<string>;
        ImageIconUrl?: pulumi.Input<string>;
        ImageSmallIconUrl?: pulumi.Input<string>;
        ImageUrl?: pulumi.Input<string>;
        JsonBody?: pulumi.Input<string>;
        MediaUrl?: pulumi.Input<string>;
        RawContent?: pulumi.Input<string>;
        SilentPush?: pulumi.Input<boolean>;
        TimeToLive?: pulumi.Input<number>;
        Title?: pulumi.Input<string>;
        Url?: pulumi.Input<string>;
    }
    
    export interface CampaignMetricDimensionAttributes {
        ComparisonOperator?: string;
        Value?: number;
    }
    
    export interface CampaignMetricDimensionProperties {
        ComparisonOperator?: pulumi.Input<string>;
        Value?: pulumi.Input<number>;
    }
    
    export interface CampaignProperties {
        AdditionalTreatments?: pulumi.Input<pulumi.Input<CampaignWriteTreatmentResourceProperties>[]>;
        ApplicationId: pulumi.Input<string>;
        CampaignHook?: pulumi.Input<CampaignCampaignHookProperties>;
        Description?: pulumi.Input<string>;
        HoldoutPercent?: pulumi.Input<number>;
        IsPaused?: pulumi.Input<boolean>;
        Limits?: pulumi.Input<CampaignLimitsProperties>;
        MessageConfiguration: pulumi.Input<CampaignMessageConfigurationProperties>;
        Name: pulumi.Input<string>;
        Schedule: pulumi.Input<CampaignScheduleProperties>;
        SegmentId: pulumi.Input<string>;
        SegmentVersion?: pulumi.Input<number>;
        Tags?: pulumi.Input<string>;
        TreatmentDescription?: pulumi.Input<string>;
        TreatmentName?: pulumi.Input<string>;
    }
    
    export interface CampaignQuietTimeAttributes {
        End: string;
        Start: string;
    }
    
    export interface CampaignQuietTimeProperties {
        End: pulumi.Input<string>;
        Start: pulumi.Input<string>;
    }
    
    export interface CampaignScheduleAttributes {
        EndTime?: string;
        EventFilter?: CampaignCampaignEventFilterAttributes;
        Frequency?: string;
        IsLocalTime?: boolean;
        QuietTime?: CampaignQuietTimeAttributes;
        StartTime?: string;
        TimeZone?: string;
    }
    
    export interface CampaignScheduleProperties {
        EndTime?: pulumi.Input<string>;
        EventFilter?: pulumi.Input<CampaignCampaignEventFilterProperties>;
        Frequency?: pulumi.Input<string>;
        IsLocalTime?: pulumi.Input<boolean>;
        QuietTime?: pulumi.Input<CampaignQuietTimeProperties>;
        StartTime?: pulumi.Input<string>;
        TimeZone?: pulumi.Input<string>;
    }
    
    export interface CampaignSetDimensionAttributes {
        DimensionType?: string;
        Values?: string[];
    }
    
    export interface CampaignSetDimensionProperties {
        DimensionType?: pulumi.Input<string>;
        Values?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface CampaignWriteTreatmentResourceAttributes {
        MessageConfiguration?: CampaignMessageConfigurationAttributes;
        Schedule?: CampaignScheduleAttributes;
        SizePercent?: number;
        TreatmentDescription?: string;
        TreatmentName?: string;
    }
    
    export interface CampaignWriteTreatmentResourceProperties {
        MessageConfiguration?: pulumi.Input<CampaignMessageConfigurationProperties>;
        Schedule?: pulumi.Input<CampaignScheduleProperties>;
        SizePercent?: pulumi.Input<number>;
        TreatmentDescription?: pulumi.Input<string>;
        TreatmentName?: pulumi.Input<string>;
    }
    
    export interface EmailChannelAttributes {
    }
    
    export interface EmailChannelProperties {
        ApplicationId: pulumi.Input<string>;
        ConfigurationSet?: pulumi.Input<string>;
        Enabled?: pulumi.Input<boolean>;
        FromAddress: pulumi.Input<string>;
        Identity: pulumi.Input<string>;
        RoleArn?: pulumi.Input<string>;
    }
    
    export interface EmailTemplateAttributes {
        Arn: string;
    }
    
    export interface EmailTemplateProperties {
        HtmlPart?: pulumi.Input<string>;
        Subject: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
        TemplateName: pulumi.Input<string>;
        TextPart?: pulumi.Input<string>;
    }
    
    export interface EventStreamAttributes {
    }
    
    export interface EventStreamProperties {
        ApplicationId: pulumi.Input<string>;
        DestinationStreamArn: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
    }
    
    export interface GCMChannelAttributes {
    }
    
    export interface GCMChannelProperties {
        ApiKey: pulumi.Input<string>;
        ApplicationId: pulumi.Input<string>;
        Enabled?: pulumi.Input<boolean>;
    }
    
    export interface PushTemplateAPNSPushNotificationTemplateAttributes {
        Action?: string;
        Body?: string;
        MediaUrl?: string;
        Sound?: string;
        Title?: string;
        Url?: string;
    }
    
    export interface PushTemplateAPNSPushNotificationTemplateProperties {
        Action?: pulumi.Input<string>;
        Body?: pulumi.Input<string>;
        MediaUrl?: pulumi.Input<string>;
        Sound?: pulumi.Input<string>;
        Title?: pulumi.Input<string>;
        Url?: pulumi.Input<string>;
    }
    
    export interface PushTemplateAndroidPushNotificationTemplateAttributes {
        Action?: string;
        Body?: string;
        ImageIconUrl?: string;
        ImageUrl?: string;
        SmallImageIconUrl?: string;
        Sound?: string;
        Title?: string;
        Url?: string;
    }
    
    export interface PushTemplateAndroidPushNotificationTemplateProperties {
        Action?: pulumi.Input<string>;
        Body?: pulumi.Input<string>;
        ImageIconUrl?: pulumi.Input<string>;
        ImageUrl?: pulumi.Input<string>;
        SmallImageIconUrl?: pulumi.Input<string>;
        Sound?: pulumi.Input<string>;
        Title?: pulumi.Input<string>;
        Url?: pulumi.Input<string>;
    }
    
    export interface PushTemplateAttributes {
        Arn: string;
    }
    
    export interface PushTemplateDefaultPushNotificationTemplateAttributes {
        Action?: string;
        Body?: string;
        Sound?: string;
        Title?: string;
        Url?: string;
    }
    
    export interface PushTemplateDefaultPushNotificationTemplateProperties {
        Action?: pulumi.Input<string>;
        Body?: pulumi.Input<string>;
        Sound?: pulumi.Input<string>;
        Title?: pulumi.Input<string>;
        Url?: pulumi.Input<string>;
    }
    
    export interface PushTemplateProperties {
        ADM?: pulumi.Input<PushTemplateAndroidPushNotificationTemplateProperties>;
        APNS?: pulumi.Input<PushTemplateAPNSPushNotificationTemplateProperties>;
        Baidu?: pulumi.Input<PushTemplateAndroidPushNotificationTemplateProperties>;
        Default?: pulumi.Input<PushTemplateDefaultPushNotificationTemplateProperties>;
        GCM?: pulumi.Input<PushTemplateAndroidPushNotificationTemplateProperties>;
        Tags?: pulumi.Input<string>;
        TemplateName: pulumi.Input<string>;
    }
    
    export interface SMSChannelAttributes {
    }
    
    export interface SMSChannelProperties {
        ApplicationId: pulumi.Input<string>;
        Enabled?: pulumi.Input<boolean>;
        SenderId?: pulumi.Input<string>;
        ShortCode?: pulumi.Input<string>;
    }
    
    export interface SegmentAttributeDimensionAttributes {
        AttributeType?: string;
        Values?: string[];
    }
    
    export interface SegmentAttributeDimensionProperties {
        AttributeType?: pulumi.Input<string>;
        Values?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface SegmentAttributes {
        Arn: string;
        SegmentId: string;
    }
    
    export interface SegmentBehaviorAttributes {
        Recency?: SegmentRecencyAttributes;
    }
    
    export interface SegmentBehaviorProperties {
        Recency?: pulumi.Input<SegmentRecencyProperties>;
    }
    
    export interface SegmentCoordinatesAttributes {
        Latitude: number;
        Longitude: number;
    }
    
    export interface SegmentCoordinatesProperties {
        Latitude: pulumi.Input<number>;
        Longitude: pulumi.Input<number>;
    }
    
    export interface SegmentDemographicAttributes {
        AppVersion?: SegmentSetDimensionAttributes;
        Channel?: SegmentSetDimensionAttributes;
        DeviceType?: SegmentSetDimensionAttributes;
        Make?: SegmentSetDimensionAttributes;
        Model?: SegmentSetDimensionAttributes;
        Platform?: SegmentSetDimensionAttributes;
    }
    
    export interface SegmentDemographicProperties {
        AppVersion?: pulumi.Input<SegmentSetDimensionProperties>;
        Channel?: pulumi.Input<SegmentSetDimensionProperties>;
        DeviceType?: pulumi.Input<SegmentSetDimensionProperties>;
        Make?: pulumi.Input<SegmentSetDimensionProperties>;
        Model?: pulumi.Input<SegmentSetDimensionProperties>;
        Platform?: pulumi.Input<SegmentSetDimensionProperties>;
    }
    
    export interface SegmentGPSPointAttributes {
        Coordinates: SegmentCoordinatesAttributes;
        RangeInKilometers: number;
    }
    
    export interface SegmentGPSPointProperties {
        Coordinates: pulumi.Input<SegmentCoordinatesProperties>;
        RangeInKilometers: pulumi.Input<number>;
    }
    
    export interface SegmentGroupsAttributes {
        Dimensions?: SegmentSegmentDimensionsAttributes[];
        SourceSegments?: SegmentSourceSegmentsAttributes[];
        SourceType?: string;
        Type?: string;
    }
    
    export interface SegmentGroupsProperties {
        Dimensions?: pulumi.Input<pulumi.Input<SegmentSegmentDimensionsProperties>[]>;
        SourceSegments?: pulumi.Input<pulumi.Input<SegmentSourceSegmentsProperties>[]>;
        SourceType?: pulumi.Input<string>;
        Type?: pulumi.Input<string>;
    }
    
    export interface SegmentLocationAttributes {
        Country?: SegmentSetDimensionAttributes;
        GPSPoint?: SegmentGPSPointAttributes;
    }
    
    export interface SegmentLocationProperties {
        Country?: pulumi.Input<SegmentSetDimensionProperties>;
        GPSPoint?: pulumi.Input<SegmentGPSPointProperties>;
    }
    
    export interface SegmentProperties {
        ApplicationId: pulumi.Input<string>;
        Dimensions?: pulumi.Input<SegmentSegmentDimensionsProperties>;
        Name: pulumi.Input<string>;
        SegmentGroups?: pulumi.Input<SegmentSegmentGroupsProperties>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface SegmentRecencyAttributes {
        Duration: string;
        RecencyType: string;
    }
    
    export interface SegmentRecencyProperties {
        Duration: pulumi.Input<string>;
        RecencyType: pulumi.Input<string>;
    }
    
    export interface SegmentSegmentDimensionsAttributes {
        Attributes?: string;
        Behavior?: SegmentBehaviorAttributes;
        Demographic?: SegmentDemographicAttributes;
        Location?: SegmentLocationAttributes;
        Metrics?: string;
        UserAttributes?: string;
    }
    
    export interface SegmentSegmentDimensionsProperties {
        Attributes?: pulumi.Input<string>;
        Behavior?: pulumi.Input<SegmentBehaviorProperties>;
        Demographic?: pulumi.Input<SegmentDemographicProperties>;
        Location?: pulumi.Input<SegmentLocationProperties>;
        Metrics?: pulumi.Input<string>;
        UserAttributes?: pulumi.Input<string>;
    }
    
    export interface SegmentSegmentGroupsAttributes {
        Groups?: SegmentGroupsAttributes[];
        Include?: string;
    }
    
    export interface SegmentSegmentGroupsProperties {
        Groups?: pulumi.Input<pulumi.Input<SegmentGroupsProperties>[]>;
        Include?: pulumi.Input<string>;
    }
    
    export interface SegmentSetDimensionAttributes {
        DimensionType?: string;
        Values?: string[];
    }
    
    export interface SegmentSetDimensionProperties {
        DimensionType?: pulumi.Input<string>;
        Values?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface SegmentSourceSegmentsAttributes {
        Id: string;
        Version?: number;
    }
    
    export interface SegmentSourceSegmentsProperties {
        Id: pulumi.Input<string>;
        Version?: pulumi.Input<number>;
    }
    
    export interface SmsTemplateAttributes {
        Arn: string;
    }
    
    export interface SmsTemplateProperties {
        Body: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
        TemplateName: pulumi.Input<string>;
    }
    
    export interface VoiceChannelAttributes {
    }
    
    export interface VoiceChannelProperties {
        ApplicationId: pulumi.Input<string>;
        Enabled?: pulumi.Input<boolean>;
    }
    
    
    export class ADMChannel extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ADMChannelAttributes>;

        constructor(name: string, properties: ADMChannelProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:ADMChannel", name, inputs, opts);
        }
    }
    
    export class APNSChannel extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<APNSChannelAttributes>;

        constructor(name: string, properties: APNSChannelProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:APNSChannel", name, inputs, opts);
        }
    }
    
    export class APNSSandboxChannel extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<APNSSandboxChannelAttributes>;

        constructor(name: string, properties: APNSSandboxChannelProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:APNSSandboxChannel", name, inputs, opts);
        }
    }
    
    export class APNSVoipChannel extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<APNSVoipChannelAttributes>;

        constructor(name: string, properties: APNSVoipChannelProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:APNSVoipChannel", name, inputs, opts);
        }
    }
    
    export class APNSVoipSandboxChannel extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<APNSVoipSandboxChannelAttributes>;

        constructor(name: string, properties: APNSVoipSandboxChannelProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:APNSVoipSandboxChannel", name, inputs, opts);
        }
    }
    
    export class App extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AppAttributes>;

        constructor(name: string, properties: AppProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:App", name, inputs, opts);
        }
    }
    
    export class ApplicationSettings extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ApplicationSettingsAttributes>;

        constructor(name: string, properties: ApplicationSettingsProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:ApplicationSettings", name, inputs, opts);
        }
    }
    
    export class BaiduChannel extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<BaiduChannelAttributes>;

        constructor(name: string, properties: BaiduChannelProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:BaiduChannel", name, inputs, opts);
        }
    }
    
    export class Campaign extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<CampaignAttributes>;

        constructor(name: string, properties: CampaignProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:Campaign", name, inputs, opts);
        }
    }
    
    export class EmailChannel extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<EmailChannelAttributes>;

        constructor(name: string, properties: EmailChannelProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:EmailChannel", name, inputs, opts);
        }
    }
    
    export class EmailTemplate extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<EmailTemplateAttributes>;

        constructor(name: string, properties: EmailTemplateProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:EmailTemplate", name, inputs, opts);
        }
    }
    
    export class EventStream extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<EventStreamAttributes>;

        constructor(name: string, properties: EventStreamProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:EventStream", name, inputs, opts);
        }
    }
    
    export class GCMChannel extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<GCMChannelAttributes>;

        constructor(name: string, properties: GCMChannelProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:GCMChannel", name, inputs, opts);
        }
    }
    
    export class PushTemplate extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PushTemplateAttributes>;

        constructor(name: string, properties: PushTemplateProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:PushTemplate", name, inputs, opts);
        }
    }
    
    export class SMSChannel extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SMSChannelAttributes>;

        constructor(name: string, properties: SMSChannelProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:SMSChannel", name, inputs, opts);
        }
    }
    
    export class Segment extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SegmentAttributes>;

        constructor(name: string, properties: SegmentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:Segment", name, inputs, opts);
        }
    }
    
    export class SmsTemplate extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SmsTemplateAttributes>;

        constructor(name: string, properties: SmsTemplateProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:SmsTemplate", name, inputs, opts);
        }
    }
    
    export class VoiceChannel extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<VoiceChannelAttributes>;

        constructor(name: string, properties: VoiceChannelProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Pinpoint:VoiceChannel", name, inputs, opts);
        }
    }
    
}

export namespace pinpointemail {
    
    export interface ConfigurationSetAttributes {
    }
    
    export interface ConfigurationSetDeliveryOptionsAttributes {
        SendingPoolName?: string;
    }
    
    export interface ConfigurationSetDeliveryOptionsProperties {
        SendingPoolName?: pulumi.Input<string>;
    }
    
    export interface ConfigurationSetEventDestinationAttributes {
    }
    
    export interface ConfigurationSetEventDestinationCloudWatchDestinationAttributes {
        DimensionConfigurations?: ConfigurationSetEventDestinationDimensionConfigurationAttributes[];
    }
    
    export interface ConfigurationSetEventDestinationCloudWatchDestinationProperties {
        DimensionConfigurations?: pulumi.Input<pulumi.Input<ConfigurationSetEventDestinationDimensionConfigurationProperties>[]>;
    }
    
    export interface ConfigurationSetEventDestinationDimensionConfigurationAttributes {
        DefaultDimensionValue: string;
        DimensionName: string;
        DimensionValueSource: string;
    }
    
    export interface ConfigurationSetEventDestinationDimensionConfigurationProperties {
        DefaultDimensionValue: pulumi.Input<string>;
        DimensionName: pulumi.Input<string>;
        DimensionValueSource: pulumi.Input<string>;
    }
    
    export interface ConfigurationSetEventDestinationEventDestinationAttributes {
        CloudWatchDestination?: ConfigurationSetEventDestinationCloudWatchDestinationAttributes;
        Enabled?: boolean;
        KinesisFirehoseDestination?: ConfigurationSetEventDestinationKinesisFirehoseDestinationAttributes;
        MatchingEventTypes: string[];
        PinpointDestination?: ConfigurationSetEventDestinationPinpointDestinationAttributes;
        SnsDestination?: ConfigurationSetEventDestinationSnsDestinationAttributes;
    }
    
    export interface ConfigurationSetEventDestinationEventDestinationProperties {
        CloudWatchDestination?: pulumi.Input<ConfigurationSetEventDestinationCloudWatchDestinationProperties>;
        Enabled?: pulumi.Input<boolean>;
        KinesisFirehoseDestination?: pulumi.Input<ConfigurationSetEventDestinationKinesisFirehoseDestinationProperties>;
        MatchingEventTypes: pulumi.Input<pulumi.Input<string>[]>;
        PinpointDestination?: pulumi.Input<ConfigurationSetEventDestinationPinpointDestinationProperties>;
        SnsDestination?: pulumi.Input<ConfigurationSetEventDestinationSnsDestinationProperties>;
    }
    
    export interface ConfigurationSetEventDestinationKinesisFirehoseDestinationAttributes {
        DeliveryStreamArn: string;
        IamRoleArn: string;
    }
    
    export interface ConfigurationSetEventDestinationKinesisFirehoseDestinationProperties {
        DeliveryStreamArn: pulumi.Input<string>;
        IamRoleArn: pulumi.Input<string>;
    }
    
    export interface ConfigurationSetEventDestinationPinpointDestinationAttributes {
        ApplicationArn?: string;
    }
    
    export interface ConfigurationSetEventDestinationPinpointDestinationProperties {
        ApplicationArn?: pulumi.Input<string>;
    }
    
    export interface ConfigurationSetEventDestinationProperties {
        ConfigurationSetName: pulumi.Input<string>;
        EventDestination?: pulumi.Input<ConfigurationSetEventDestinationEventDestinationProperties>;
        EventDestinationName: pulumi.Input<string>;
    }
    
    export interface ConfigurationSetEventDestinationSnsDestinationAttributes {
        TopicArn: string;
    }
    
    export interface ConfigurationSetEventDestinationSnsDestinationProperties {
        TopicArn: pulumi.Input<string>;
    }
    
    export interface ConfigurationSetProperties {
        DeliveryOptions?: pulumi.Input<ConfigurationSetDeliveryOptionsProperties>;
        Name: pulumi.Input<string>;
        ReputationOptions?: pulumi.Input<ConfigurationSetReputationOptionsProperties>;
        SendingOptions?: pulumi.Input<ConfigurationSetSendingOptionsProperties>;
        Tags?: pulumi.Input<pulumi.Input<ConfigurationSetTagsProperties>[]>;
        TrackingOptions?: pulumi.Input<ConfigurationSetTrackingOptionsProperties>;
    }
    
    export interface ConfigurationSetReputationOptionsAttributes {
        ReputationMetricsEnabled?: boolean;
    }
    
    export interface ConfigurationSetReputationOptionsProperties {
        ReputationMetricsEnabled?: pulumi.Input<boolean>;
    }
    
    export interface ConfigurationSetSendingOptionsAttributes {
        SendingEnabled?: boolean;
    }
    
    export interface ConfigurationSetSendingOptionsProperties {
        SendingEnabled?: pulumi.Input<boolean>;
    }
    
    export interface ConfigurationSetTagsAttributes {
        Key?: string;
        Value?: string;
    }
    
    export interface ConfigurationSetTagsProperties {
        Key?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface ConfigurationSetTrackingOptionsAttributes {
        CustomRedirectDomain?: string;
    }
    
    export interface ConfigurationSetTrackingOptionsProperties {
        CustomRedirectDomain?: pulumi.Input<string>;
    }
    
    export interface IdentityAttributes {
        IdentityDNSRecordName1: string;
        IdentityDNSRecordName2: string;
        IdentityDNSRecordName3: string;
        IdentityDNSRecordValue1: string;
        IdentityDNSRecordValue2: string;
        IdentityDNSRecordValue3: string;
    }
    
    export interface IdentityMailFromAttributesAttributes {
        BehaviorOnMxFailure?: string;
        MailFromDomain?: string;
    }
    
    export interface IdentityMailFromAttributesProperties {
        BehaviorOnMxFailure?: pulumi.Input<string>;
        MailFromDomain?: pulumi.Input<string>;
    }
    
    export interface IdentityProperties {
        DkimSigningEnabled?: pulumi.Input<boolean>;
        FeedbackForwardingEnabled?: pulumi.Input<boolean>;
        MailFromAttributes?: pulumi.Input<IdentityMailFromAttributesProperties>;
        Name: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<IdentityTagsProperties>[]>;
    }
    
    export interface IdentityTagsAttributes {
        Key?: string;
        Value?: string;
    }
    
    export interface IdentityTagsProperties {
        Key?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    
    export class ConfigurationSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ConfigurationSetAttributes>;

        constructor(name: string, properties: ConfigurationSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:PinpointEmail:ConfigurationSet", name, inputs, opts);
        }
    }
    
    export class ConfigurationSetEventDestination extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ConfigurationSetEventDestinationAttributes>;

        constructor(name: string, properties: ConfigurationSetEventDestinationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:PinpointEmail:ConfigurationSetEventDestination", name, inputs, opts);
        }
    }
    
    export class Identity extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<IdentityAttributes>;

        constructor(name: string, properties: IdentityProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:PinpointEmail:Identity", name, inputs, opts);
        }
    }
    
}

export namespace qldb {
    
    export interface LedgerAttributes {
    }
    
    export interface LedgerProperties {
        DeletionProtection?: pulumi.Input<boolean>;
        Name?: pulumi.Input<string>;
        PermissionsMode: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    
    export class Ledger extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<LedgerAttributes>;

        constructor(name: string, properties: LedgerProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:QLDB:Ledger", name, inputs, opts);
        }
    }
    
}

export namespace ram {
    
    export interface ResourceShareAttributes {
        Arn: string;
    }
    
    export interface ResourceShareProperties {
        AllowExternalPrincipals?: pulumi.Input<boolean>;
        Name: pulumi.Input<string>;
        Principals?: pulumi.Input<pulumi.Input<string>[]>;
        ResourceArns?: pulumi.Input<pulumi.Input<string>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    
    export class ResourceShare extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ResourceShareAttributes>;

        constructor(name: string, properties: ResourceShareProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:RAM:ResourceShare", name, inputs, opts);
        }
    }
    
}

export namespace rds {
    
    export interface DBClusterAttributes {
        "Endpoint.Address": string;
        "Endpoint.Port": string;
        "ReadEndpoint.Address": string;
    }
    
    export interface DBClusterDBClusterRoleAttributes {
        FeatureName?: string;
        RoleArn: string;
    }
    
    export interface DBClusterDBClusterRoleProperties {
        FeatureName?: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
    }
    
    export interface DBClusterParameterGroupAttributes {
    }
    
    export interface DBClusterParameterGroupProperties {
        Description: pulumi.Input<string>;
        Family: pulumi.Input<string>;
        Parameters: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface DBClusterProperties {
        AssociatedRoles?: pulumi.Input<pulumi.Input<DBClusterDBClusterRoleProperties>[]>;
        AvailabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
        BacktrackWindow?: pulumi.Input<number>;
        BackupRetentionPeriod?: pulumi.Input<number>;
        DBClusterIdentifier?: pulumi.Input<string>;
        DBClusterParameterGroupName?: pulumi.Input<string>;
        DBSubnetGroupName?: pulumi.Input<string>;
        DatabaseName?: pulumi.Input<string>;
        DeletionProtection?: pulumi.Input<boolean>;
        EnableCloudwatchLogsExports?: pulumi.Input<pulumi.Input<string>[]>;
        EnableIAMDatabaseAuthentication?: pulumi.Input<boolean>;
        Engine: pulumi.Input<string>;
        EngineMode?: pulumi.Input<string>;
        EngineVersion?: pulumi.Input<string>;
        KmsKeyId?: pulumi.Input<string>;
        MasterUserPassword?: pulumi.Input<string>;
        MasterUsername?: pulumi.Input<string>;
        Port?: pulumi.Input<number>;
        PreferredBackupWindow?: pulumi.Input<string>;
        PreferredMaintenanceWindow?: pulumi.Input<string>;
        ReplicationSourceIdentifier?: pulumi.Input<string>;
        RestoreType?: pulumi.Input<string>;
        ScalingConfiguration?: pulumi.Input<DBClusterScalingConfigurationProperties>;
        SnapshotIdentifier?: pulumi.Input<string>;
        SourceDBClusterIdentifier?: pulumi.Input<string>;
        SourceRegion?: pulumi.Input<string>;
        StorageEncrypted?: pulumi.Input<boolean>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        UseLatestRestorableTime?: pulumi.Input<boolean>;
        VpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface DBClusterScalingConfigurationAttributes {
        AutoPause?: boolean;
        MaxCapacity?: number;
        MinCapacity?: number;
        SecondsUntilAutoPause?: number;
    }
    
    export interface DBClusterScalingConfigurationProperties {
        AutoPause?: pulumi.Input<boolean>;
        MaxCapacity?: pulumi.Input<number>;
        MinCapacity?: pulumi.Input<number>;
        SecondsUntilAutoPause?: pulumi.Input<number>;
    }
    
    export interface DBInstanceAttributes {
        "Endpoint.Address": string;
        "Endpoint.Port": string;
    }
    
    export interface DBInstanceDBInstanceRoleAttributes {
        FeatureName: string;
        RoleArn: string;
    }
    
    export interface DBInstanceDBInstanceRoleProperties {
        FeatureName: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
    }
    
    export interface DBInstanceProcessorFeatureAttributes {
        Name?: string;
        Value?: string;
    }
    
    export interface DBInstanceProcessorFeatureProperties {
        Name?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface DBInstanceProperties {
        AllocatedStorage?: pulumi.Input<string>;
        AllowMajorVersionUpgrade?: pulumi.Input<boolean>;
        AssociatedRoles?: pulumi.Input<pulumi.Input<DBInstanceDBInstanceRoleProperties>[]>;
        AutoMinorVersionUpgrade?: pulumi.Input<boolean>;
        AvailabilityZone?: pulumi.Input<string>;
        BackupRetentionPeriod?: pulumi.Input<number>;
        CharacterSetName?: pulumi.Input<string>;
        CopyTagsToSnapshot?: pulumi.Input<boolean>;
        DBClusterIdentifier?: pulumi.Input<string>;
        DBInstanceClass: pulumi.Input<string>;
        DBInstanceIdentifier?: pulumi.Input<string>;
        DBName?: pulumi.Input<string>;
        DBParameterGroupName?: pulumi.Input<string>;
        DBSecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
        DBSnapshotIdentifier?: pulumi.Input<string>;
        DBSubnetGroupName?: pulumi.Input<string>;
        DeleteAutomatedBackups?: pulumi.Input<boolean>;
        DeletionProtection?: pulumi.Input<boolean>;
        Domain?: pulumi.Input<string>;
        DomainIAMRoleName?: pulumi.Input<string>;
        EnableCloudwatchLogsExports?: pulumi.Input<pulumi.Input<string>[]>;
        EnableIAMDatabaseAuthentication?: pulumi.Input<boolean>;
        EnablePerformanceInsights?: pulumi.Input<boolean>;
        Engine?: pulumi.Input<string>;
        EngineVersion?: pulumi.Input<string>;
        Iops?: pulumi.Input<number>;
        KmsKeyId?: pulumi.Input<string>;
        LicenseModel?: pulumi.Input<string>;
        MasterUserPassword?: pulumi.Input<string>;
        MasterUsername?: pulumi.Input<string>;
        MonitoringInterval?: pulumi.Input<number>;
        MonitoringRoleArn?: pulumi.Input<string>;
        MultiAZ?: pulumi.Input<boolean>;
        OptionGroupName?: pulumi.Input<string>;
        PerformanceInsightsKMSKeyId?: pulumi.Input<string>;
        PerformanceInsightsRetentionPeriod?: pulumi.Input<number>;
        Port?: pulumi.Input<string>;
        PreferredBackupWindow?: pulumi.Input<string>;
        PreferredMaintenanceWindow?: pulumi.Input<string>;
        ProcessorFeatures?: pulumi.Input<pulumi.Input<DBInstanceProcessorFeatureProperties>[]>;
        PromotionTier?: pulumi.Input<number>;
        PubliclyAccessible?: pulumi.Input<boolean>;
        SourceDBInstanceIdentifier?: pulumi.Input<string>;
        SourceRegion?: pulumi.Input<string>;
        StorageEncrypted?: pulumi.Input<boolean>;
        StorageType?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        Timezone?: pulumi.Input<string>;
        UseDefaultProcessorFeatures?: pulumi.Input<boolean>;
        VPCSecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface DBParameterGroupAttributes {
    }
    
    export interface DBParameterGroupProperties {
        Description: pulumi.Input<string>;
        Family: pulumi.Input<string>;
        Parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface DBSecurityGroupAttributes {
    }
    
    export interface DBSecurityGroupIngressAttributes {
    }
    
    export interface DBSecurityGroupIngressProperties {
        CIDRIP?: pulumi.Input<string>;
        DBSecurityGroupName: pulumi.Input<string>;
        EC2SecurityGroupId?: pulumi.Input<string>;
        EC2SecurityGroupName?: pulumi.Input<string>;
        EC2SecurityGroupOwnerId?: pulumi.Input<string>;
    }
    
    export interface DBSecurityGroupProperties {
        DBSecurityGroupIngress: pulumi.Input<pulumi.Input<DBSecurityGroupIngressProperties>[]>;
        EC2VpcId?: pulumi.Input<string>;
        GroupDescription: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface DBSubnetGroupAttributes {
    }
    
    export interface DBSubnetGroupProperties {
        DBSubnetGroupDescription: pulumi.Input<string>;
        DBSubnetGroupName?: pulumi.Input<string>;
        SubnetIds: pulumi.Input<pulumi.Input<string>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface EventSubscriptionAttributes {
    }
    
    export interface EventSubscriptionProperties {
        Enabled?: pulumi.Input<boolean>;
        EventCategories?: pulumi.Input<pulumi.Input<string>[]>;
        SnsTopicArn: pulumi.Input<string>;
        SourceIds?: pulumi.Input<pulumi.Input<string>[]>;
        SourceType?: pulumi.Input<string>;
    }
    
    export interface OptionGroupAttributes {
    }
    
    export interface OptionGroupOptionConfigurationAttributes {
        DBSecurityGroupMemberships?: string[];
        OptionName: string;
        OptionSettings?: OptionGroupOptionSettingAttributes[];
        OptionVersion?: string;
        Port?: number;
        VpcSecurityGroupMemberships?: string[];
    }
    
    export interface OptionGroupOptionConfigurationProperties {
        DBSecurityGroupMemberships?: pulumi.Input<pulumi.Input<string>[]>;
        OptionName: pulumi.Input<string>;
        OptionSettings?: pulumi.Input<pulumi.Input<OptionGroupOptionSettingProperties>[]>;
        OptionVersion?: pulumi.Input<string>;
        Port?: pulumi.Input<number>;
        VpcSecurityGroupMemberships?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface OptionGroupOptionSettingAttributes {
        Name?: string;
        Value?: string;
    }
    
    export interface OptionGroupOptionSettingProperties {
        Name?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface OptionGroupProperties {
        EngineName: pulumi.Input<string>;
        MajorEngineVersion: pulumi.Input<string>;
        OptionConfigurations: pulumi.Input<pulumi.Input<OptionGroupOptionConfigurationProperties>[]>;
        OptionGroupDescription: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    
    export class DBCluster extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DBClusterAttributes>;

        constructor(name: string, properties: DBClusterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:RDS:DBCluster", name, inputs, opts);
        }
    }
    
    export class DBClusterParameterGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DBClusterParameterGroupAttributes>;

        constructor(name: string, properties: DBClusterParameterGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:RDS:DBClusterParameterGroup", name, inputs, opts);
        }
    }
    
    export class DBInstance extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DBInstanceAttributes>;

        constructor(name: string, properties: DBInstanceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:RDS:DBInstance", name, inputs, opts);
        }
    }
    
    export class DBParameterGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DBParameterGroupAttributes>;

        constructor(name: string, properties: DBParameterGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:RDS:DBParameterGroup", name, inputs, opts);
        }
    }
    
    export class DBSecurityGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DBSecurityGroupAttributes>;

        constructor(name: string, properties: DBSecurityGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:RDS:DBSecurityGroup", name, inputs, opts);
        }
    }
    
    export class DBSecurityGroupIngress extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DBSecurityGroupIngressAttributes>;

        constructor(name: string, properties: DBSecurityGroupIngressProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:RDS:DBSecurityGroupIngress", name, inputs, opts);
        }
    }
    
    export class DBSubnetGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DBSubnetGroupAttributes>;

        constructor(name: string, properties: DBSubnetGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:RDS:DBSubnetGroup", name, inputs, opts);
        }
    }
    
    export class EventSubscription extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<EventSubscriptionAttributes>;

        constructor(name: string, properties: EventSubscriptionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:RDS:EventSubscription", name, inputs, opts);
        }
    }
    
    export class OptionGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<OptionGroupAttributes>;

        constructor(name: string, properties: OptionGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:RDS:OptionGroup", name, inputs, opts);
        }
    }
    
}

export namespace redshift {
    
    export interface ClusterAttributes {
        "Endpoint.Address": string;
        "Endpoint.Port": string;
    }
    
    export interface ClusterLoggingPropertiesAttributes {
        BucketName: string;
        S3KeyPrefix?: string;
    }
    
    export interface ClusterLoggingPropertiesProperties {
        BucketName: pulumi.Input<string>;
        S3KeyPrefix?: pulumi.Input<string>;
    }
    
    export interface ClusterParameterGroupAttributes {
    }
    
    export interface ClusterParameterGroupParameterAttributes {
        ParameterName: string;
        ParameterValue: string;
    }
    
    export interface ClusterParameterGroupParameterProperties {
        ParameterName: pulumi.Input<string>;
        ParameterValue: pulumi.Input<string>;
    }
    
    export interface ClusterParameterGroupProperties {
        Description: pulumi.Input<string>;
        ParameterGroupFamily: pulumi.Input<string>;
        Parameters?: pulumi.Input<pulumi.Input<ClusterParameterGroupParameterProperties>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface ClusterProperties {
        AllowVersionUpgrade?: pulumi.Input<boolean>;
        AutomatedSnapshotRetentionPeriod?: pulumi.Input<number>;
        AvailabilityZone?: pulumi.Input<string>;
        ClusterIdentifier?: pulumi.Input<string>;
        ClusterParameterGroupName?: pulumi.Input<string>;
        ClusterSecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
        ClusterSubnetGroupName?: pulumi.Input<string>;
        ClusterType: pulumi.Input<string>;
        ClusterVersion?: pulumi.Input<string>;
        DBName: pulumi.Input<string>;
        ElasticIp?: pulumi.Input<string>;
        Encrypted?: pulumi.Input<boolean>;
        HsmClientCertificateIdentifier?: pulumi.Input<string>;
        HsmConfigurationIdentifier?: pulumi.Input<string>;
        IamRoles?: pulumi.Input<pulumi.Input<string>[]>;
        KmsKeyId?: pulumi.Input<string>;
        LoggingProperties?: pulumi.Input<ClusterLoggingPropertiesProperties>;
        MasterUserPassword: pulumi.Input<string>;
        MasterUsername: pulumi.Input<string>;
        NodeType: pulumi.Input<string>;
        NumberOfNodes?: pulumi.Input<number>;
        OwnerAccount?: pulumi.Input<string>;
        Port?: pulumi.Input<number>;
        PreferredMaintenanceWindow?: pulumi.Input<string>;
        PubliclyAccessible?: pulumi.Input<boolean>;
        SnapshotClusterIdentifier?: pulumi.Input<string>;
        SnapshotIdentifier?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface ClusterSecurityGroupAttributes {
    }
    
    export interface ClusterSecurityGroupIngressAttributes {
    }
    
    export interface ClusterSecurityGroupIngressProperties {
        CIDRIP?: pulumi.Input<string>;
        ClusterSecurityGroupName: pulumi.Input<string>;
        EC2SecurityGroupName?: pulumi.Input<string>;
        EC2SecurityGroupOwnerId?: pulumi.Input<string>;
    }
    
    export interface ClusterSecurityGroupProperties {
        Description: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface ClusterSubnetGroupAttributes {
    }
    
    export interface ClusterSubnetGroupProperties {
        Description: pulumi.Input<string>;
        SubnetIds: pulumi.Input<pulumi.Input<string>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    
    export class Cluster extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ClusterAttributes>;

        constructor(name: string, properties: ClusterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Redshift:Cluster", name, inputs, opts);
        }
    }
    
    export class ClusterParameterGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ClusterParameterGroupAttributes>;

        constructor(name: string, properties: ClusterParameterGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Redshift:ClusterParameterGroup", name, inputs, opts);
        }
    }
    
    export class ClusterSecurityGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ClusterSecurityGroupAttributes>;

        constructor(name: string, properties: ClusterSecurityGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Redshift:ClusterSecurityGroup", name, inputs, opts);
        }
    }
    
    export class ClusterSecurityGroupIngress extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ClusterSecurityGroupIngressAttributes>;

        constructor(name: string, properties: ClusterSecurityGroupIngressProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Redshift:ClusterSecurityGroupIngress", name, inputs, opts);
        }
    }
    
    export class ClusterSubnetGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ClusterSubnetGroupAttributes>;

        constructor(name: string, properties: ClusterSubnetGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Redshift:ClusterSubnetGroup", name, inputs, opts);
        }
    }
    
}

export namespace robomaker {
    
    export interface FleetAttributes {
        Arn: string;
    }
    
    export interface FleetProperties {
        Name?: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface RobotApplicationAttributes {
        Arn: string;
        CurrentRevisionId: string;
    }
    
    export interface RobotApplicationProperties {
        CurrentRevisionId?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        RobotSoftwareSuite: pulumi.Input<RobotApplicationRobotSoftwareSuiteProperties>;
        Sources: pulumi.Input<pulumi.Input<RobotApplicationSourceConfigProperties>[]>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface RobotApplicationRobotSoftwareSuiteAttributes {
        Name: string;
        Version: string;
    }
    
    export interface RobotApplicationRobotSoftwareSuiteProperties {
        Name: pulumi.Input<string>;
        Version: pulumi.Input<string>;
    }
    
    export interface RobotApplicationSourceConfigAttributes {
        Architecture: string;
        S3Bucket: string;
        S3Key: string;
    }
    
    export interface RobotApplicationSourceConfigProperties {
        Architecture: pulumi.Input<string>;
        S3Bucket: pulumi.Input<string>;
        S3Key: pulumi.Input<string>;
    }
    
    export interface RobotApplicationVersionAttributes {
    }
    
    export interface RobotApplicationVersionProperties {
        Application: pulumi.Input<string>;
        CurrentRevisionId?: pulumi.Input<string>;
    }
    
    export interface RobotAttributes {
    }
    
    export interface RobotProperties {
        Architecture: pulumi.Input<string>;
        Fleet?: pulumi.Input<string>;
        GreengrassGroupId: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface SimulationApplicationAttributes {
        Arn: string;
        CurrentRevisionId: string;
    }
    
    export interface SimulationApplicationProperties {
        CurrentRevisionId?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        RenderingEngine: pulumi.Input<SimulationApplicationRenderingEngineProperties>;
        RobotSoftwareSuite: pulumi.Input<SimulationApplicationRobotSoftwareSuiteProperties>;
        SimulationSoftwareSuite: pulumi.Input<SimulationApplicationSimulationSoftwareSuiteProperties>;
        Sources: pulumi.Input<pulumi.Input<SimulationApplicationSourceConfigProperties>[]>;
        Tags?: pulumi.Input<string>;
    }
    
    export interface SimulationApplicationRenderingEngineAttributes {
        Name: string;
        Version: string;
    }
    
    export interface SimulationApplicationRenderingEngineProperties {
        Name: pulumi.Input<string>;
        Version: pulumi.Input<string>;
    }
    
    export interface SimulationApplicationRobotSoftwareSuiteAttributes {
        Name: string;
        Version: string;
    }
    
    export interface SimulationApplicationRobotSoftwareSuiteProperties {
        Name: pulumi.Input<string>;
        Version: pulumi.Input<string>;
    }
    
    export interface SimulationApplicationSimulationSoftwareSuiteAttributes {
        Name: string;
        Version: string;
    }
    
    export interface SimulationApplicationSimulationSoftwareSuiteProperties {
        Name: pulumi.Input<string>;
        Version: pulumi.Input<string>;
    }
    
    export interface SimulationApplicationSourceConfigAttributes {
        Architecture: string;
        S3Bucket: string;
        S3Key: string;
    }
    
    export interface SimulationApplicationSourceConfigProperties {
        Architecture: pulumi.Input<string>;
        S3Bucket: pulumi.Input<string>;
        S3Key: pulumi.Input<string>;
    }
    
    export interface SimulationApplicationVersionAttributes {
    }
    
    export interface SimulationApplicationVersionProperties {
        Application: pulumi.Input<string>;
        CurrentRevisionId?: pulumi.Input<string>;
    }
    
    
    export class Fleet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<FleetAttributes>;

        constructor(name: string, properties?: FleetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:RoboMaker:Fleet", name, inputs, opts);
        }
    }
    
    export class Robot extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RobotAttributes>;

        constructor(name: string, properties: RobotProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:RoboMaker:Robot", name, inputs, opts);
        }
    }
    
    export class RobotApplication extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RobotApplicationAttributes>;

        constructor(name: string, properties: RobotApplicationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:RoboMaker:RobotApplication", name, inputs, opts);
        }
    }
    
    export class RobotApplicationVersion extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RobotApplicationVersionAttributes>;

        constructor(name: string, properties: RobotApplicationVersionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:RoboMaker:RobotApplicationVersion", name, inputs, opts);
        }
    }
    
    export class SimulationApplication extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SimulationApplicationAttributes>;

        constructor(name: string, properties: SimulationApplicationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:RoboMaker:SimulationApplication", name, inputs, opts);
        }
    }
    
    export class SimulationApplicationVersion extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SimulationApplicationVersionAttributes>;

        constructor(name: string, properties: SimulationApplicationVersionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:RoboMaker:SimulationApplicationVersion", name, inputs, opts);
        }
    }
    
}

export namespace route53 {
    
    export interface HealthCheckAlarmIdentifierAttributes {
        Name: string;
        Region: string;
    }
    
    export interface HealthCheckAlarmIdentifierProperties {
        Name: pulumi.Input<string>;
        Region: pulumi.Input<string>;
    }
    
    export interface HealthCheckAttributes {
    }
    
    export interface HealthCheckHealthCheckConfigAttributes {
        AlarmIdentifier?: HealthCheckAlarmIdentifierAttributes;
        ChildHealthChecks?: string[];
        EnableSNI?: boolean;
        FailureThreshold?: number;
        FullyQualifiedDomainName?: string;
        HealthThreshold?: number;
        IPAddress?: string;
        InsufficientDataHealthStatus?: string;
        Inverted?: boolean;
        MeasureLatency?: boolean;
        Port?: number;
        Regions?: string[];
        RequestInterval?: number;
        ResourcePath?: string;
        SearchString?: string;
        Type: string;
    }
    
    export interface HealthCheckHealthCheckConfigProperties {
        AlarmIdentifier?: pulumi.Input<HealthCheckAlarmIdentifierProperties>;
        ChildHealthChecks?: pulumi.Input<pulumi.Input<string>[]>;
        EnableSNI?: pulumi.Input<boolean>;
        FailureThreshold?: pulumi.Input<number>;
        FullyQualifiedDomainName?: pulumi.Input<string>;
        HealthThreshold?: pulumi.Input<number>;
        IPAddress?: pulumi.Input<string>;
        InsufficientDataHealthStatus?: pulumi.Input<string>;
        Inverted?: pulumi.Input<boolean>;
        MeasureLatency?: pulumi.Input<boolean>;
        Port?: pulumi.Input<number>;
        Regions?: pulumi.Input<pulumi.Input<string>[]>;
        RequestInterval?: pulumi.Input<number>;
        ResourcePath?: pulumi.Input<string>;
        SearchString?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface HealthCheckHealthCheckTagAttributes {
        Key: string;
        Value: string;
    }
    
    export interface HealthCheckHealthCheckTagProperties {
        Key: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface HealthCheckProperties {
        HealthCheckConfig: pulumi.Input<HealthCheckHealthCheckConfigProperties>;
        HealthCheckTags?: pulumi.Input<pulumi.Input<HealthCheckHealthCheckTagProperties>[]>;
    }
    
    export interface HostedZoneAttributes {
        NameServers: string[];
    }
    
    export interface HostedZoneHostedZoneConfigAttributes {
        Comment?: string;
    }
    
    export interface HostedZoneHostedZoneConfigProperties {
        Comment?: pulumi.Input<string>;
    }
    
    export interface HostedZoneHostedZoneTagAttributes {
        Key: string;
        Value: string;
    }
    
    export interface HostedZoneHostedZoneTagProperties {
        Key: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface HostedZoneProperties {
        HostedZoneConfig?: pulumi.Input<HostedZoneHostedZoneConfigProperties>;
        HostedZoneTags?: pulumi.Input<pulumi.Input<HostedZoneHostedZoneTagProperties>[]>;
        Name: pulumi.Input<string>;
        QueryLoggingConfig?: pulumi.Input<HostedZoneQueryLoggingConfigProperties>;
        VPCs?: pulumi.Input<pulumi.Input<HostedZoneVPCProperties>[]>;
    }
    
    export interface HostedZoneQueryLoggingConfigAttributes {
        CloudWatchLogsLogGroupArn: string;
    }
    
    export interface HostedZoneQueryLoggingConfigProperties {
        CloudWatchLogsLogGroupArn: pulumi.Input<string>;
    }
    
    export interface HostedZoneVPCAttributes {
        VPCId: string;
        VPCRegion: string;
    }
    
    export interface HostedZoneVPCProperties {
        VPCId: pulumi.Input<string>;
        VPCRegion: pulumi.Input<string>;
    }
    
    export interface RecordSetAliasTargetAttributes {
        DNSName: string;
        EvaluateTargetHealth?: boolean;
        HostedZoneId: string;
    }
    
    export interface RecordSetAliasTargetProperties {
        DNSName: pulumi.Input<string>;
        EvaluateTargetHealth?: pulumi.Input<boolean>;
        HostedZoneId: pulumi.Input<string>;
    }
    
    export interface RecordSetAttributes {
    }
    
    export interface RecordSetGeoLocationAttributes {
        ContinentCode?: string;
        CountryCode?: string;
        SubdivisionCode?: string;
    }
    
    export interface RecordSetGeoLocationProperties {
        ContinentCode?: pulumi.Input<string>;
        CountryCode?: pulumi.Input<string>;
        SubdivisionCode?: pulumi.Input<string>;
    }
    
    export interface RecordSetGroupAliasTargetAttributes {
        DNSName: string;
        EvaluateTargetHealth?: boolean;
        HostedZoneId: string;
    }
    
    export interface RecordSetGroupAliasTargetProperties {
        DNSName: pulumi.Input<string>;
        EvaluateTargetHealth?: pulumi.Input<boolean>;
        HostedZoneId: pulumi.Input<string>;
    }
    
    export interface RecordSetGroupAttributes {
    }
    
    export interface RecordSetGroupGeoLocationAttributes {
        ContinentCode?: string;
        CountryCode?: string;
        SubdivisionCode?: string;
    }
    
    export interface RecordSetGroupGeoLocationProperties {
        ContinentCode?: pulumi.Input<string>;
        CountryCode?: pulumi.Input<string>;
        SubdivisionCode?: pulumi.Input<string>;
    }
    
    export interface RecordSetGroupProperties {
        Comment?: pulumi.Input<string>;
        HostedZoneId?: pulumi.Input<string>;
        HostedZoneName?: pulumi.Input<string>;
        RecordSets?: pulumi.Input<pulumi.Input<RecordSetGroupRecordSetProperties>[]>;
    }
    
    export interface RecordSetGroupRecordSetAttributes {
        AliasTarget?: RecordSetGroupAliasTargetAttributes;
        Comment?: string;
        Failover?: string;
        GeoLocation?: RecordSetGroupGeoLocationAttributes;
        HealthCheckId?: string;
        HostedZoneId?: string;
        HostedZoneName?: string;
        MultiValueAnswer?: boolean;
        Name: string;
        Region?: string;
        ResourceRecords?: string[];
        SetIdentifier?: string;
        TTL?: string;
        Type: string;
        Weight?: number;
    }
    
    export interface RecordSetGroupRecordSetProperties {
        AliasTarget?: pulumi.Input<RecordSetGroupAliasTargetProperties>;
        Comment?: pulumi.Input<string>;
        Failover?: pulumi.Input<string>;
        GeoLocation?: pulumi.Input<RecordSetGroupGeoLocationProperties>;
        HealthCheckId?: pulumi.Input<string>;
        HostedZoneId?: pulumi.Input<string>;
        HostedZoneName?: pulumi.Input<string>;
        MultiValueAnswer?: pulumi.Input<boolean>;
        Name: pulumi.Input<string>;
        Region?: pulumi.Input<string>;
        ResourceRecords?: pulumi.Input<pulumi.Input<string>[]>;
        SetIdentifier?: pulumi.Input<string>;
        TTL?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
        Weight?: pulumi.Input<number>;
    }
    
    export interface RecordSetProperties {
        AliasTarget?: pulumi.Input<RecordSetAliasTargetProperties>;
        Comment?: pulumi.Input<string>;
        Failover?: pulumi.Input<string>;
        GeoLocation?: pulumi.Input<RecordSetGeoLocationProperties>;
        HealthCheckId?: pulumi.Input<string>;
        HostedZoneId?: pulumi.Input<string>;
        HostedZoneName?: pulumi.Input<string>;
        MultiValueAnswer?: pulumi.Input<boolean>;
        Name: pulumi.Input<string>;
        Region?: pulumi.Input<string>;
        ResourceRecords?: pulumi.Input<pulumi.Input<string>[]>;
        SetIdentifier?: pulumi.Input<string>;
        TTL?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
        Weight?: pulumi.Input<number>;
    }
    
    
    export class HealthCheck extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<HealthCheckAttributes>;

        constructor(name: string, properties: HealthCheckProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Route53:HealthCheck", name, inputs, opts);
        }
    }
    
    export class HostedZone extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<HostedZoneAttributes>;

        constructor(name: string, properties: HostedZoneProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Route53:HostedZone", name, inputs, opts);
        }
    }
    
    export class RecordSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RecordSetAttributes>;

        constructor(name: string, properties: RecordSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Route53:RecordSet", name, inputs, opts);
        }
    }
    
    export class RecordSetGroup extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RecordSetGroupAttributes>;

        constructor(name: string, properties?: RecordSetGroupProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Route53:RecordSetGroup", name, inputs, opts);
        }
    }
    
}

export namespace route53resolver {
    
    export interface ResolverEndpointAttributes {
        Arn: string;
        Direction: string;
        HostVPCId: string;
        IpAddressCount: string;
        Name: string;
        ResolverEndpointId: string;
    }
    
    export interface ResolverEndpointIpAddressRequestAttributes {
        Ip?: string;
        SubnetId: string;
    }
    
    export interface ResolverEndpointIpAddressRequestProperties {
        Ip?: pulumi.Input<string>;
        SubnetId: pulumi.Input<string>;
    }
    
    export interface ResolverEndpointProperties {
        Direction: pulumi.Input<string>;
        IpAddresses: pulumi.Input<pulumi.Input<ResolverEndpointIpAddressRequestProperties>[]>;
        Name?: pulumi.Input<string>;
        SecurityGroupIds: pulumi.Input<pulumi.Input<string>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface ResolverRuleAssociationAttributes {
        Name: string;
        ResolverRuleAssociationId: string;
        ResolverRuleId: string;
        VPCId: string;
    }
    
    export interface ResolverRuleAssociationProperties {
        Name?: pulumi.Input<string>;
        ResolverRuleId: pulumi.Input<string>;
        VPCId: pulumi.Input<string>;
    }
    
    export interface ResolverRuleAttributes {
        Arn: string;
        DomainName: string;
        Name: string;
        ResolverEndpointId: string;
        ResolverRuleId: string;
        TargetIps: string;
    }
    
    export interface ResolverRuleProperties {
        DomainName: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        ResolverEndpointId?: pulumi.Input<string>;
        RuleType: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        TargetIps?: pulumi.Input<pulumi.Input<ResolverRuleTargetAddressProperties>[]>;
    }
    
    export interface ResolverRuleTargetAddressAttributes {
        Ip: string;
        Port?: string;
    }
    
    export interface ResolverRuleTargetAddressProperties {
        Ip: pulumi.Input<string>;
        Port?: pulumi.Input<string>;
    }
    
    
    export class ResolverEndpoint extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ResolverEndpointAttributes>;

        constructor(name: string, properties: ResolverEndpointProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Route53Resolver:ResolverEndpoint", name, inputs, opts);
        }
    }
    
    export class ResolverRule extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ResolverRuleAttributes>;

        constructor(name: string, properties: ResolverRuleProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Route53Resolver:ResolverRule", name, inputs, opts);
        }
    }
    
    export class ResolverRuleAssociation extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ResolverRuleAssociationAttributes>;

        constructor(name: string, properties: ResolverRuleAssociationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Route53Resolver:ResolverRuleAssociation", name, inputs, opts);
        }
    }
    
}

export namespace s3 {
    
    export interface BucketAbortIncompleteMultipartUploadAttributes {
        DaysAfterInitiation: number;
    }
    
    export interface BucketAbortIncompleteMultipartUploadProperties {
        DaysAfterInitiation: pulumi.Input<number>;
    }
    
    export interface BucketAccelerateConfigurationAttributes {
        AccelerationStatus: string;
    }
    
    export interface BucketAccelerateConfigurationProperties {
        AccelerationStatus: pulumi.Input<string>;
    }
    
    export interface BucketAccessControlTranslationAttributes {
        Owner: string;
    }
    
    export interface BucketAccessControlTranslationProperties {
        Owner: pulumi.Input<string>;
    }
    
    export interface BucketAnalyticsConfigurationAttributes {
        Id: string;
        Prefix?: string;
        StorageClassAnalysis: BucketStorageClassAnalysisAttributes;
        TagFilters?: BucketTagFilterAttributes[];
    }
    
    export interface BucketAnalyticsConfigurationProperties {
        Id: pulumi.Input<string>;
        Prefix?: pulumi.Input<string>;
        StorageClassAnalysis: pulumi.Input<BucketStorageClassAnalysisProperties>;
        TagFilters?: pulumi.Input<pulumi.Input<BucketTagFilterProperties>[]>;
    }
    
    export interface BucketAttributes {
        Arn: string;
        DomainName: string;
        DualStackDomainName: string;
        RegionalDomainName: string;
        WebsiteURL: string;
    }
    
    export interface BucketBucketEncryptionAttributes {
        ServerSideEncryptionConfiguration: BucketServerSideEncryptionRuleAttributes[];
    }
    
    export interface BucketBucketEncryptionProperties {
        ServerSideEncryptionConfiguration: pulumi.Input<pulumi.Input<BucketServerSideEncryptionRuleProperties>[]>;
    }
    
    export interface BucketCorsConfigurationAttributes {
        CorsRules: BucketCorsRuleAttributes[];
    }
    
    export interface BucketCorsConfigurationProperties {
        CorsRules: pulumi.Input<pulumi.Input<BucketCorsRuleProperties>[]>;
    }
    
    export interface BucketCorsRuleAttributes {
        AllowedHeaders?: string[];
        AllowedMethods: string[];
        AllowedOrigins: string[];
        ExposedHeaders?: string[];
        Id?: string;
        MaxAge?: number;
    }
    
    export interface BucketCorsRuleProperties {
        AllowedHeaders?: pulumi.Input<pulumi.Input<string>[]>;
        AllowedMethods: pulumi.Input<pulumi.Input<string>[]>;
        AllowedOrigins: pulumi.Input<pulumi.Input<string>[]>;
        ExposedHeaders?: pulumi.Input<pulumi.Input<string>[]>;
        Id?: pulumi.Input<string>;
        MaxAge?: pulumi.Input<number>;
    }
    
    export interface BucketDataExportAttributes {
        Destination: BucketDestinationAttributes;
        OutputSchemaVersion: string;
    }
    
    export interface BucketDataExportProperties {
        Destination: pulumi.Input<BucketDestinationProperties>;
        OutputSchemaVersion: pulumi.Input<string>;
    }
    
    export interface BucketDefaultRetentionAttributes {
        Days?: number;
        Mode?: string;
        Years?: number;
    }
    
    export interface BucketDefaultRetentionProperties {
        Days?: pulumi.Input<number>;
        Mode?: pulumi.Input<string>;
        Years?: pulumi.Input<number>;
    }
    
    export interface BucketDestinationAttributes {
        BucketAccountId?: string;
        BucketArn: string;
        Format: string;
        Prefix?: string;
    }
    
    export interface BucketDestinationProperties {
        BucketAccountId?: pulumi.Input<string>;
        BucketArn: pulumi.Input<string>;
        Format: pulumi.Input<string>;
        Prefix?: pulumi.Input<string>;
    }
    
    export interface BucketEncryptionConfigurationAttributes {
        ReplicaKmsKeyID: string;
    }
    
    export interface BucketEncryptionConfigurationProperties {
        ReplicaKmsKeyID: pulumi.Input<string>;
    }
    
    export interface BucketFilterRuleAttributes {
        Name: string;
        Value: string;
    }
    
    export interface BucketFilterRuleProperties {
        Name: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface BucketInventoryConfigurationAttributes {
        Destination: BucketDestinationAttributes;
        Enabled: boolean;
        Id: string;
        IncludedObjectVersions: string;
        OptionalFields?: string[];
        Prefix?: string;
        ScheduleFrequency: string;
    }
    
    export interface BucketInventoryConfigurationProperties {
        Destination: pulumi.Input<BucketDestinationProperties>;
        Enabled: pulumi.Input<boolean>;
        Id: pulumi.Input<string>;
        IncludedObjectVersions: pulumi.Input<string>;
        OptionalFields?: pulumi.Input<pulumi.Input<string>[]>;
        Prefix?: pulumi.Input<string>;
        ScheduleFrequency: pulumi.Input<string>;
    }
    
    export interface BucketLambdaConfigurationAttributes {
        Event: string;
        Filter?: BucketNotificationFilterAttributes;
        Function: string;
    }
    
    export interface BucketLambdaConfigurationProperties {
        Event: pulumi.Input<string>;
        Filter?: pulumi.Input<BucketNotificationFilterProperties>;
        Function: pulumi.Input<string>;
    }
    
    export interface BucketLifecycleConfigurationAttributes {
        Rules: BucketRuleAttributes[];
    }
    
    export interface BucketLifecycleConfigurationProperties {
        Rules: pulumi.Input<pulumi.Input<BucketRuleProperties>[]>;
    }
    
    export interface BucketLoggingConfigurationAttributes {
        DestinationBucketName?: string;
        LogFilePrefix?: string;
    }
    
    export interface BucketLoggingConfigurationProperties {
        DestinationBucketName?: pulumi.Input<string>;
        LogFilePrefix?: pulumi.Input<string>;
    }
    
    export interface BucketMetricsConfigurationAttributes {
        Id: string;
        Prefix?: string;
        TagFilters?: BucketTagFilterAttributes[];
    }
    
    export interface BucketMetricsConfigurationProperties {
        Id: pulumi.Input<string>;
        Prefix?: pulumi.Input<string>;
        TagFilters?: pulumi.Input<pulumi.Input<BucketTagFilterProperties>[]>;
    }
    
    export interface BucketNoncurrentVersionTransitionAttributes {
        StorageClass: string;
        TransitionInDays: number;
    }
    
    export interface BucketNoncurrentVersionTransitionProperties {
        StorageClass: pulumi.Input<string>;
        TransitionInDays: pulumi.Input<number>;
    }
    
    export interface BucketNotificationConfigurationAttributes {
        LambdaConfigurations?: BucketLambdaConfigurationAttributes[];
        QueueConfigurations?: BucketQueueConfigurationAttributes[];
        TopicConfigurations?: BucketTopicConfigurationAttributes[];
    }
    
    export interface BucketNotificationConfigurationProperties {
        LambdaConfigurations?: pulumi.Input<pulumi.Input<BucketLambdaConfigurationProperties>[]>;
        QueueConfigurations?: pulumi.Input<pulumi.Input<BucketQueueConfigurationProperties>[]>;
        TopicConfigurations?: pulumi.Input<pulumi.Input<BucketTopicConfigurationProperties>[]>;
    }
    
    export interface BucketNotificationFilterAttributes {
        S3Key: BucketS3KeyFilterAttributes;
    }
    
    export interface BucketNotificationFilterProperties {
        S3Key: pulumi.Input<BucketS3KeyFilterProperties>;
    }
    
    export interface BucketObjectLockConfigurationAttributes {
        ObjectLockEnabled?: string;
        Rule?: BucketObjectLockRuleAttributes;
    }
    
    export interface BucketObjectLockConfigurationProperties {
        ObjectLockEnabled?: pulumi.Input<string>;
        Rule?: pulumi.Input<BucketObjectLockRuleProperties>;
    }
    
    export interface BucketObjectLockRuleAttributes {
        DefaultRetention?: BucketDefaultRetentionAttributes;
    }
    
    export interface BucketObjectLockRuleProperties {
        DefaultRetention?: pulumi.Input<BucketDefaultRetentionProperties>;
    }
    
    export interface BucketPolicyAttributes {
    }
    
    export interface BucketPolicyProperties {
        Bucket: pulumi.Input<string>;
        PolicyDocument: pulumi.Input<string>;
    }
    
    export interface BucketProperties {
        AccelerateConfiguration?: pulumi.Input<BucketAccelerateConfigurationProperties>;
        AccessControl?: pulumi.Input<string>;
        AnalyticsConfigurations?: pulumi.Input<pulumi.Input<BucketAnalyticsConfigurationProperties>[]>;
        BucketEncryption?: pulumi.Input<BucketBucketEncryptionProperties>;
        BucketName?: pulumi.Input<string>;
        CorsConfiguration?: pulumi.Input<BucketCorsConfigurationProperties>;
        InventoryConfigurations?: pulumi.Input<pulumi.Input<BucketInventoryConfigurationProperties>[]>;
        LifecycleConfiguration?: pulumi.Input<BucketLifecycleConfigurationProperties>;
        LoggingConfiguration?: pulumi.Input<BucketLoggingConfigurationProperties>;
        MetricsConfigurations?: pulumi.Input<pulumi.Input<BucketMetricsConfigurationProperties>[]>;
        NotificationConfiguration?: pulumi.Input<BucketNotificationConfigurationProperties>;
        ObjectLockConfiguration?: pulumi.Input<BucketObjectLockConfigurationProperties>;
        ObjectLockEnabled?: pulumi.Input<boolean>;
        PublicAccessBlockConfiguration?: pulumi.Input<BucketPublicAccessBlockConfigurationProperties>;
        ReplicationConfiguration?: pulumi.Input<BucketReplicationConfigurationProperties>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VersioningConfiguration?: pulumi.Input<BucketVersioningConfigurationProperties>;
        WebsiteConfiguration?: pulumi.Input<BucketWebsiteConfigurationProperties>;
    }
    
    export interface BucketPublicAccessBlockConfigurationAttributes {
        BlockPublicAcls?: boolean;
        BlockPublicPolicy?: boolean;
        IgnorePublicAcls?: boolean;
        RestrictPublicBuckets?: boolean;
    }
    
    export interface BucketPublicAccessBlockConfigurationProperties {
        BlockPublicAcls?: pulumi.Input<boolean>;
        BlockPublicPolicy?: pulumi.Input<boolean>;
        IgnorePublicAcls?: pulumi.Input<boolean>;
        RestrictPublicBuckets?: pulumi.Input<boolean>;
    }
    
    export interface BucketQueueConfigurationAttributes {
        Event: string;
        Filter?: BucketNotificationFilterAttributes;
        Queue: string;
    }
    
    export interface BucketQueueConfigurationProperties {
        Event: pulumi.Input<string>;
        Filter?: pulumi.Input<BucketNotificationFilterProperties>;
        Queue: pulumi.Input<string>;
    }
    
    export interface BucketRedirectAllRequestsToAttributes {
        HostName: string;
        Protocol?: string;
    }
    
    export interface BucketRedirectAllRequestsToProperties {
        HostName: pulumi.Input<string>;
        Protocol?: pulumi.Input<string>;
    }
    
    export interface BucketRedirectRuleAttributes {
        HostName?: string;
        HttpRedirectCode?: string;
        Protocol?: string;
        ReplaceKeyPrefixWith?: string;
        ReplaceKeyWith?: string;
    }
    
    export interface BucketRedirectRuleProperties {
        HostName?: pulumi.Input<string>;
        HttpRedirectCode?: pulumi.Input<string>;
        Protocol?: pulumi.Input<string>;
        ReplaceKeyPrefixWith?: pulumi.Input<string>;
        ReplaceKeyWith?: pulumi.Input<string>;
    }
    
    export interface BucketReplicationConfigurationAttributes {
        Role: string;
        Rules: BucketReplicationRuleAttributes[];
    }
    
    export interface BucketReplicationConfigurationProperties {
        Role: pulumi.Input<string>;
        Rules: pulumi.Input<pulumi.Input<BucketReplicationRuleProperties>[]>;
    }
    
    export interface BucketReplicationDestinationAttributes {
        AccessControlTranslation?: BucketAccessControlTranslationAttributes;
        Account?: string;
        Bucket: string;
        EncryptionConfiguration?: BucketEncryptionConfigurationAttributes;
        StorageClass?: string;
    }
    
    export interface BucketReplicationDestinationProperties {
        AccessControlTranslation?: pulumi.Input<BucketAccessControlTranslationProperties>;
        Account?: pulumi.Input<string>;
        Bucket: pulumi.Input<string>;
        EncryptionConfiguration?: pulumi.Input<BucketEncryptionConfigurationProperties>;
        StorageClass?: pulumi.Input<string>;
    }
    
    export interface BucketReplicationRuleAttributes {
        Destination: BucketReplicationDestinationAttributes;
        Id?: string;
        Prefix: string;
        SourceSelectionCriteria?: BucketSourceSelectionCriteriaAttributes;
        Status: string;
    }
    
    export interface BucketReplicationRuleProperties {
        Destination: pulumi.Input<BucketReplicationDestinationProperties>;
        Id?: pulumi.Input<string>;
        Prefix: pulumi.Input<string>;
        SourceSelectionCriteria?: pulumi.Input<BucketSourceSelectionCriteriaProperties>;
        Status: pulumi.Input<string>;
    }
    
    export interface BucketRoutingRuleAttributes {
        RedirectRule: BucketRedirectRuleAttributes;
        RoutingRuleCondition?: BucketRoutingRuleConditionAttributes;
    }
    
    export interface BucketRoutingRuleConditionAttributes {
        HttpErrorCodeReturnedEquals?: string;
        KeyPrefixEquals?: string;
    }
    
    export interface BucketRoutingRuleConditionProperties {
        HttpErrorCodeReturnedEquals?: pulumi.Input<string>;
        KeyPrefixEquals?: pulumi.Input<string>;
    }
    
    export interface BucketRoutingRuleProperties {
        RedirectRule: pulumi.Input<BucketRedirectRuleProperties>;
        RoutingRuleCondition?: pulumi.Input<BucketRoutingRuleConditionProperties>;
    }
    
    export interface BucketRuleAttributes {
        AbortIncompleteMultipartUpload?: BucketAbortIncompleteMultipartUploadAttributes;
        ExpirationDate?: string;
        ExpirationInDays?: number;
        Id?: string;
        NoncurrentVersionExpirationInDays?: number;
        NoncurrentVersionTransition?: BucketNoncurrentVersionTransitionAttributes;
        NoncurrentVersionTransitions?: BucketNoncurrentVersionTransitionAttributes[];
        Prefix?: string;
        Status: string;
        TagFilters?: BucketTagFilterAttributes[];
        Transition?: BucketTransitionAttributes;
        Transitions?: BucketTransitionAttributes[];
    }
    
    export interface BucketRuleProperties {
        AbortIncompleteMultipartUpload?: pulumi.Input<BucketAbortIncompleteMultipartUploadProperties>;
        ExpirationDate?: pulumi.Input<string>;
        ExpirationInDays?: pulumi.Input<number>;
        Id?: pulumi.Input<string>;
        NoncurrentVersionExpirationInDays?: pulumi.Input<number>;
        NoncurrentVersionTransition?: pulumi.Input<BucketNoncurrentVersionTransitionProperties>;
        NoncurrentVersionTransitions?: pulumi.Input<pulumi.Input<BucketNoncurrentVersionTransitionProperties>[]>;
        Prefix?: pulumi.Input<string>;
        Status: pulumi.Input<string>;
        TagFilters?: pulumi.Input<pulumi.Input<BucketTagFilterProperties>[]>;
        Transition?: pulumi.Input<BucketTransitionProperties>;
        Transitions?: pulumi.Input<pulumi.Input<BucketTransitionProperties>[]>;
    }
    
    export interface BucketS3KeyFilterAttributes {
        Rules: BucketFilterRuleAttributes[];
    }
    
    export interface BucketS3KeyFilterProperties {
        Rules: pulumi.Input<pulumi.Input<BucketFilterRuleProperties>[]>;
    }
    
    export interface BucketServerSideEncryptionByDefaultAttributes {
        KMSMasterKeyID?: string;
        SSEAlgorithm: string;
    }
    
    export interface BucketServerSideEncryptionByDefaultProperties {
        KMSMasterKeyID?: pulumi.Input<string>;
        SSEAlgorithm: pulumi.Input<string>;
    }
    
    export interface BucketServerSideEncryptionRuleAttributes {
        ServerSideEncryptionByDefault?: BucketServerSideEncryptionByDefaultAttributes;
    }
    
    export interface BucketServerSideEncryptionRuleProperties {
        ServerSideEncryptionByDefault?: pulumi.Input<BucketServerSideEncryptionByDefaultProperties>;
    }
    
    export interface BucketSourceSelectionCriteriaAttributes {
        SseKmsEncryptedObjects: BucketSseKmsEncryptedObjectsAttributes;
    }
    
    export interface BucketSourceSelectionCriteriaProperties {
        SseKmsEncryptedObjects: pulumi.Input<BucketSseKmsEncryptedObjectsProperties>;
    }
    
    export interface BucketSseKmsEncryptedObjectsAttributes {
        Status: string;
    }
    
    export interface BucketSseKmsEncryptedObjectsProperties {
        Status: pulumi.Input<string>;
    }
    
    export interface BucketStorageClassAnalysisAttributes {
        DataExport?: BucketDataExportAttributes;
    }
    
    export interface BucketStorageClassAnalysisProperties {
        DataExport?: pulumi.Input<BucketDataExportProperties>;
    }
    
    export interface BucketTagFilterAttributes {
        Key: string;
        Value: string;
    }
    
    export interface BucketTagFilterProperties {
        Key: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface BucketTopicConfigurationAttributes {
        Event: string;
        Filter?: BucketNotificationFilterAttributes;
        Topic: string;
    }
    
    export interface BucketTopicConfigurationProperties {
        Event: pulumi.Input<string>;
        Filter?: pulumi.Input<BucketNotificationFilterProperties>;
        Topic: pulumi.Input<string>;
    }
    
    export interface BucketTransitionAttributes {
        StorageClass: string;
        TransitionDate?: string;
        TransitionInDays?: number;
    }
    
    export interface BucketTransitionProperties {
        StorageClass: pulumi.Input<string>;
        TransitionDate?: pulumi.Input<string>;
        TransitionInDays?: pulumi.Input<number>;
    }
    
    export interface BucketVersioningConfigurationAttributes {
        Status: string;
    }
    
    export interface BucketVersioningConfigurationProperties {
        Status: pulumi.Input<string>;
    }
    
    export interface BucketWebsiteConfigurationAttributes {
        ErrorDocument?: string;
        IndexDocument?: string;
        RedirectAllRequestsTo?: BucketRedirectAllRequestsToAttributes;
        RoutingRules?: BucketRoutingRuleAttributes[];
    }
    
    export interface BucketWebsiteConfigurationProperties {
        ErrorDocument?: pulumi.Input<string>;
        IndexDocument?: pulumi.Input<string>;
        RedirectAllRequestsTo?: pulumi.Input<BucketRedirectAllRequestsToProperties>;
        RoutingRules?: pulumi.Input<pulumi.Input<BucketRoutingRuleProperties>[]>;
    }
    
    
    export class Bucket extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<BucketAttributes>;

        constructor(name: string, properties?: BucketProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:S3:Bucket", name, inputs, opts);
        }
    }
    
    export class BucketPolicy extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<BucketPolicyAttributes>;

        constructor(name: string, properties: BucketPolicyProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:S3:BucketPolicy", name, inputs, opts);
        }
    }
    
}

export namespace sdb {
    
    export interface DomainAttributes {
    }
    
    export interface DomainProperties {
        Description?: pulumi.Input<string>;
    }
    
    
    export class Domain extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DomainAttributes>;

        constructor(name: string, properties?: DomainProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SDB:Domain", name, inputs, opts);
        }
    }
    
}

export namespace ses {
    
    export interface ConfigurationSetAttributes {
    }
    
    export interface ConfigurationSetEventDestinationAttributes {
    }
    
    export interface ConfigurationSetEventDestinationCloudWatchDestinationAttributes {
        DimensionConfigurations?: ConfigurationSetEventDestinationDimensionConfigurationAttributes[];
    }
    
    export interface ConfigurationSetEventDestinationCloudWatchDestinationProperties {
        DimensionConfigurations?: pulumi.Input<pulumi.Input<ConfigurationSetEventDestinationDimensionConfigurationProperties>[]>;
    }
    
    export interface ConfigurationSetEventDestinationDimensionConfigurationAttributes {
        DefaultDimensionValue: string;
        DimensionName: string;
        DimensionValueSource: string;
    }
    
    export interface ConfigurationSetEventDestinationDimensionConfigurationProperties {
        DefaultDimensionValue: pulumi.Input<string>;
        DimensionName: pulumi.Input<string>;
        DimensionValueSource: pulumi.Input<string>;
    }
    
    export interface ConfigurationSetEventDestinationEventDestinationAttributes {
        CloudWatchDestination?: ConfigurationSetEventDestinationCloudWatchDestinationAttributes;
        Enabled?: boolean;
        KinesisFirehoseDestination?: ConfigurationSetEventDestinationKinesisFirehoseDestinationAttributes;
        MatchingEventTypes: string[];
        Name?: string;
    }
    
    export interface ConfigurationSetEventDestinationEventDestinationProperties {
        CloudWatchDestination?: pulumi.Input<ConfigurationSetEventDestinationCloudWatchDestinationProperties>;
        Enabled?: pulumi.Input<boolean>;
        KinesisFirehoseDestination?: pulumi.Input<ConfigurationSetEventDestinationKinesisFirehoseDestinationProperties>;
        MatchingEventTypes: pulumi.Input<pulumi.Input<string>[]>;
        Name?: pulumi.Input<string>;
    }
    
    export interface ConfigurationSetEventDestinationKinesisFirehoseDestinationAttributes {
        DeliveryStreamARN: string;
        IAMRoleARN: string;
    }
    
    export interface ConfigurationSetEventDestinationKinesisFirehoseDestinationProperties {
        DeliveryStreamARN: pulumi.Input<string>;
        IAMRoleARN: pulumi.Input<string>;
    }
    
    export interface ConfigurationSetEventDestinationProperties {
        ConfigurationSetName: pulumi.Input<string>;
        EventDestination: pulumi.Input<ConfigurationSetEventDestinationEventDestinationProperties>;
    }
    
    export interface ConfigurationSetProperties {
        Name?: pulumi.Input<string>;
    }
    
    export interface ReceiptFilterAttributes {
    }
    
    export interface ReceiptFilterFilterAttributes {
        IpFilter: ReceiptFilterIpFilterAttributes;
        Name?: string;
    }
    
    export interface ReceiptFilterFilterProperties {
        IpFilter: pulumi.Input<ReceiptFilterIpFilterProperties>;
        Name?: pulumi.Input<string>;
    }
    
    export interface ReceiptFilterIpFilterAttributes {
        Cidr: string;
        Policy: string;
    }
    
    export interface ReceiptFilterIpFilterProperties {
        Cidr: pulumi.Input<string>;
        Policy: pulumi.Input<string>;
    }
    
    export interface ReceiptFilterProperties {
        Filter: pulumi.Input<ReceiptFilterFilterProperties>;
    }
    
    export interface ReceiptRuleActionAttributes {
        AddHeaderAction?: ReceiptRuleAddHeaderActionAttributes;
        BounceAction?: ReceiptRuleBounceActionAttributes;
        LambdaAction?: ReceiptRuleLambdaActionAttributes;
        S3Action?: ReceiptRuleS3ActionAttributes;
        SNSAction?: ReceiptRuleSNSActionAttributes;
        StopAction?: ReceiptRuleStopActionAttributes;
        WorkmailAction?: ReceiptRuleWorkmailActionAttributes;
    }
    
    export interface ReceiptRuleActionProperties {
        AddHeaderAction?: pulumi.Input<ReceiptRuleAddHeaderActionProperties>;
        BounceAction?: pulumi.Input<ReceiptRuleBounceActionProperties>;
        LambdaAction?: pulumi.Input<ReceiptRuleLambdaActionProperties>;
        S3Action?: pulumi.Input<ReceiptRuleS3ActionProperties>;
        SNSAction?: pulumi.Input<ReceiptRuleSNSActionProperties>;
        StopAction?: pulumi.Input<ReceiptRuleStopActionProperties>;
        WorkmailAction?: pulumi.Input<ReceiptRuleWorkmailActionProperties>;
    }
    
    export interface ReceiptRuleAddHeaderActionAttributes {
        HeaderName: string;
        HeaderValue: string;
    }
    
    export interface ReceiptRuleAddHeaderActionProperties {
        HeaderName: pulumi.Input<string>;
        HeaderValue: pulumi.Input<string>;
    }
    
    export interface ReceiptRuleAttributes {
    }
    
    export interface ReceiptRuleBounceActionAttributes {
        Message: string;
        Sender: string;
        SmtpReplyCode: string;
        StatusCode?: string;
        TopicArn?: string;
    }
    
    export interface ReceiptRuleBounceActionProperties {
        Message: pulumi.Input<string>;
        Sender: pulumi.Input<string>;
        SmtpReplyCode: pulumi.Input<string>;
        StatusCode?: pulumi.Input<string>;
        TopicArn?: pulumi.Input<string>;
    }
    
    export interface ReceiptRuleLambdaActionAttributes {
        FunctionArn: string;
        InvocationType?: string;
        TopicArn?: string;
    }
    
    export interface ReceiptRuleLambdaActionProperties {
        FunctionArn: pulumi.Input<string>;
        InvocationType?: pulumi.Input<string>;
        TopicArn?: pulumi.Input<string>;
    }
    
    export interface ReceiptRuleProperties {
        After?: pulumi.Input<string>;
        Rule: pulumi.Input<ReceiptRuleRuleProperties>;
        RuleSetName: pulumi.Input<string>;
    }
    
    export interface ReceiptRuleRuleAttributes {
        Actions?: ReceiptRuleActionAttributes[];
        Enabled?: boolean;
        Name?: string;
        Recipients?: string[];
        ScanEnabled?: boolean;
        TlsPolicy?: string;
    }
    
    export interface ReceiptRuleRuleProperties {
        Actions?: pulumi.Input<pulumi.Input<ReceiptRuleActionProperties>[]>;
        Enabled?: pulumi.Input<boolean>;
        Name?: pulumi.Input<string>;
        Recipients?: pulumi.Input<pulumi.Input<string>[]>;
        ScanEnabled?: pulumi.Input<boolean>;
        TlsPolicy?: pulumi.Input<string>;
    }
    
    export interface ReceiptRuleS3ActionAttributes {
        BucketName: string;
        KmsKeyArn?: string;
        ObjectKeyPrefix?: string;
        TopicArn?: string;
    }
    
    export interface ReceiptRuleS3ActionProperties {
        BucketName: pulumi.Input<string>;
        KmsKeyArn?: pulumi.Input<string>;
        ObjectKeyPrefix?: pulumi.Input<string>;
        TopicArn?: pulumi.Input<string>;
    }
    
    export interface ReceiptRuleSNSActionAttributes {
        Encoding?: string;
        TopicArn?: string;
    }
    
    export interface ReceiptRuleSNSActionProperties {
        Encoding?: pulumi.Input<string>;
        TopicArn?: pulumi.Input<string>;
    }
    
    export interface ReceiptRuleSetAttributes {
    }
    
    export interface ReceiptRuleSetProperties {
        RuleSetName?: pulumi.Input<string>;
    }
    
    export interface ReceiptRuleStopActionAttributes {
        Scope: string;
        TopicArn?: string;
    }
    
    export interface ReceiptRuleStopActionProperties {
        Scope: pulumi.Input<string>;
        TopicArn?: pulumi.Input<string>;
    }
    
    export interface ReceiptRuleWorkmailActionAttributes {
        OrganizationArn: string;
        TopicArn?: string;
    }
    
    export interface ReceiptRuleWorkmailActionProperties {
        OrganizationArn: pulumi.Input<string>;
        TopicArn?: pulumi.Input<string>;
    }
    
    export interface TemplateAttributes {
    }
    
    export interface TemplateProperties {
        Template?: pulumi.Input<TemplateTemplateProperties>;
    }
    
    export interface TemplateTemplateAttributes {
        HtmlPart?: string;
        SubjectPart?: string;
        TemplateName?: string;
        TextPart?: string;
    }
    
    export interface TemplateTemplateProperties {
        HtmlPart?: pulumi.Input<string>;
        SubjectPart?: pulumi.Input<string>;
        TemplateName?: pulumi.Input<string>;
        TextPart?: pulumi.Input<string>;
    }
    
    
    export class ConfigurationSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ConfigurationSetAttributes>;

        constructor(name: string, properties?: ConfigurationSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SES:ConfigurationSet", name, inputs, opts);
        }
    }
    
    export class ConfigurationSetEventDestination extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ConfigurationSetEventDestinationAttributes>;

        constructor(name: string, properties: ConfigurationSetEventDestinationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SES:ConfigurationSetEventDestination", name, inputs, opts);
        }
    }
    
    export class ReceiptFilter extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ReceiptFilterAttributes>;

        constructor(name: string, properties: ReceiptFilterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SES:ReceiptFilter", name, inputs, opts);
        }
    }
    
    export class ReceiptRule extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ReceiptRuleAttributes>;

        constructor(name: string, properties: ReceiptRuleProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SES:ReceiptRule", name, inputs, opts);
        }
    }
    
    export class ReceiptRuleSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ReceiptRuleSetAttributes>;

        constructor(name: string, properties?: ReceiptRuleSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SES:ReceiptRuleSet", name, inputs, opts);
        }
    }
    
    export class Template extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TemplateAttributes>;

        constructor(name: string, properties?: TemplateProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SES:Template", name, inputs, opts);
        }
    }
    
}

export namespace sns {
    
    export interface SubscriptionAttributes {
    }
    
    export interface SubscriptionProperties {
        DeliveryPolicy?: pulumi.Input<string>;
        Endpoint?: pulumi.Input<string>;
        FilterPolicy?: pulumi.Input<string>;
        Protocol: pulumi.Input<string>;
        RawMessageDelivery?: pulumi.Input<boolean>;
        Region?: pulumi.Input<string>;
        TopicArn: pulumi.Input<string>;
    }
    
    export interface TopicAttributes {
        TopicName: string;
    }
    
    export interface TopicPolicyAttributes {
    }
    
    export interface TopicPolicyProperties {
        PolicyDocument: pulumi.Input<string>;
        Topics: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface TopicProperties {
        DisplayName?: pulumi.Input<string>;
        KmsMasterKeyId?: pulumi.Input<string>;
        Subscription?: pulumi.Input<pulumi.Input<TopicSubscriptionProperties>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        TopicName?: pulumi.Input<string>;
    }
    
    export interface TopicSubscriptionAttributes {
        Endpoint: string;
        Protocol: string;
    }
    
    export interface TopicSubscriptionProperties {
        Endpoint: pulumi.Input<string>;
        Protocol: pulumi.Input<string>;
    }
    
    
    export class Subscription extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SubscriptionAttributes>;

        constructor(name: string, properties: SubscriptionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SNS:Subscription", name, inputs, opts);
        }
    }
    
    export class Topic extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TopicAttributes>;

        constructor(name: string, properties?: TopicProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SNS:Topic", name, inputs, opts);
        }
    }
    
    export class TopicPolicy extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TopicPolicyAttributes>;

        constructor(name: string, properties: TopicPolicyProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SNS:TopicPolicy", name, inputs, opts);
        }
    }
    
}

export namespace sqs {
    
    export interface QueueAttributes {
        Arn: string;
        QueueName: string;
    }
    
    export interface QueuePolicyAttributes {
    }
    
    export interface QueuePolicyProperties {
        PolicyDocument: pulumi.Input<string>;
        Queues: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface QueueProperties {
        ContentBasedDeduplication?: pulumi.Input<boolean>;
        DelaySeconds?: pulumi.Input<number>;
        FifoQueue?: pulumi.Input<boolean>;
        KmsDataKeyReusePeriodSeconds?: pulumi.Input<number>;
        KmsMasterKeyId?: pulumi.Input<string>;
        MaximumMessageSize?: pulumi.Input<number>;
        MessageRetentionPeriod?: pulumi.Input<number>;
        QueueName?: pulumi.Input<string>;
        ReceiveMessageWaitTimeSeconds?: pulumi.Input<number>;
        RedrivePolicy?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VisibilityTimeout?: pulumi.Input<number>;
    }
    
    
    export class Queue extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<QueueAttributes>;

        constructor(name: string, properties?: QueueProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SQS:Queue", name, inputs, opts);
        }
    }
    
    export class QueuePolicy extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<QueuePolicyAttributes>;

        constructor(name: string, properties: QueuePolicyProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SQS:QueuePolicy", name, inputs, opts);
        }
    }
    
}

export namespace ssm {
    
    export interface AssociationAttributes {
    }
    
    export interface AssociationInstanceAssociationOutputLocationAttributes {
        S3Location?: AssociationS3OutputLocationAttributes;
    }
    
    export interface AssociationInstanceAssociationOutputLocationProperties {
        S3Location?: pulumi.Input<AssociationS3OutputLocationProperties>;
    }
    
    export interface AssociationParameterValuesAttributes {
        ParameterValues: string[];
    }
    
    export interface AssociationParameterValuesProperties {
        ParameterValues: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface AssociationProperties {
        AssociationName?: pulumi.Input<string>;
        DocumentVersion?: pulumi.Input<string>;
        InstanceId?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        OutputLocation?: pulumi.Input<AssociationInstanceAssociationOutputLocationProperties>;
        Parameters?: pulumi.Input<{[key: string]: pulumi.Input<AssociationParameterValuesProperties>}>;
        ScheduleExpression?: pulumi.Input<string>;
        Targets?: pulumi.Input<pulumi.Input<AssociationTargetProperties>[]>;
    }
    
    export interface AssociationS3OutputLocationAttributes {
        OutputS3BucketName?: string;
        OutputS3KeyPrefix?: string;
    }
    
    export interface AssociationS3OutputLocationProperties {
        OutputS3BucketName?: pulumi.Input<string>;
        OutputS3KeyPrefix?: pulumi.Input<string>;
    }
    
    export interface AssociationTargetAttributes {
        Key: string;
        Values: string[];
    }
    
    export interface AssociationTargetProperties {
        Key: pulumi.Input<string>;
        Values: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface DocumentAttributes {
    }
    
    export interface DocumentProperties {
        Content: pulumi.Input<string>;
        DocumentType?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface MaintenanceWindowAttributes {
    }
    
    export interface MaintenanceWindowProperties {
        AllowUnassociatedTargets: pulumi.Input<boolean>;
        Cutoff: pulumi.Input<number>;
        Description?: pulumi.Input<string>;
        Duration: pulumi.Input<number>;
        EndDate?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        Schedule: pulumi.Input<string>;
        ScheduleTimezone?: pulumi.Input<string>;
        StartDate?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface MaintenanceWindowTargetAttributes {
    }
    
    export interface MaintenanceWindowTargetProperties {
        Description?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        OwnerInformation?: pulumi.Input<string>;
        ResourceType: pulumi.Input<string>;
        Targets: pulumi.Input<pulumi.Input<MaintenanceWindowTargetTargetsProperties>[]>;
        WindowId: pulumi.Input<string>;
    }
    
    export interface MaintenanceWindowTargetTargetsAttributes {
        Key: string;
        Values?: string[];
    }
    
    export interface MaintenanceWindowTargetTargetsProperties {
        Key: pulumi.Input<string>;
        Values?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface MaintenanceWindowTaskAttributes {
    }
    
    export interface MaintenanceWindowTaskLoggingInfoAttributes {
        Region: string;
        S3Bucket: string;
        S3Prefix?: string;
    }
    
    export interface MaintenanceWindowTaskLoggingInfoProperties {
        Region: pulumi.Input<string>;
        S3Bucket: pulumi.Input<string>;
        S3Prefix?: pulumi.Input<string>;
    }
    
    export interface MaintenanceWindowTaskMaintenanceWindowAutomationParametersAttributes {
        DocumentVersion?: string;
        Parameters?: string;
    }
    
    export interface MaintenanceWindowTaskMaintenanceWindowAutomationParametersProperties {
        DocumentVersion?: pulumi.Input<string>;
        Parameters?: pulumi.Input<string>;
    }
    
    export interface MaintenanceWindowTaskMaintenanceWindowLambdaParametersAttributes {
        ClientContext?: string;
        Payload?: string;
        Qualifier?: string;
    }
    
    export interface MaintenanceWindowTaskMaintenanceWindowLambdaParametersProperties {
        ClientContext?: pulumi.Input<string>;
        Payload?: pulumi.Input<string>;
        Qualifier?: pulumi.Input<string>;
    }
    
    export interface MaintenanceWindowTaskMaintenanceWindowRunCommandParametersAttributes {
        Comment?: string;
        DocumentHash?: string;
        DocumentHashType?: string;
        NotificationConfig?: MaintenanceWindowTaskNotificationConfigAttributes;
        OutputS3BucketName?: string;
        OutputS3KeyPrefix?: string;
        Parameters?: string;
        ServiceRoleArn?: string;
        TimeoutSeconds?: number;
    }
    
    export interface MaintenanceWindowTaskMaintenanceWindowRunCommandParametersProperties {
        Comment?: pulumi.Input<string>;
        DocumentHash?: pulumi.Input<string>;
        DocumentHashType?: pulumi.Input<string>;
        NotificationConfig?: pulumi.Input<MaintenanceWindowTaskNotificationConfigProperties>;
        OutputS3BucketName?: pulumi.Input<string>;
        OutputS3KeyPrefix?: pulumi.Input<string>;
        Parameters?: pulumi.Input<string>;
        ServiceRoleArn?: pulumi.Input<string>;
        TimeoutSeconds?: pulumi.Input<number>;
    }
    
    export interface MaintenanceWindowTaskMaintenanceWindowStepFunctionsParametersAttributes {
        Input?: string;
        Name?: string;
    }
    
    export interface MaintenanceWindowTaskMaintenanceWindowStepFunctionsParametersProperties {
        Input?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
    }
    
    export interface MaintenanceWindowTaskNotificationConfigAttributes {
        NotificationArn: string;
        NotificationEvents?: string[];
        NotificationType?: string;
    }
    
    export interface MaintenanceWindowTaskNotificationConfigProperties {
        NotificationArn: pulumi.Input<string>;
        NotificationEvents?: pulumi.Input<pulumi.Input<string>[]>;
        NotificationType?: pulumi.Input<string>;
    }
    
    export interface MaintenanceWindowTaskProperties {
        Description?: pulumi.Input<string>;
        LoggingInfo?: pulumi.Input<MaintenanceWindowTaskLoggingInfoProperties>;
        MaxConcurrency: pulumi.Input<string>;
        MaxErrors: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Priority: pulumi.Input<number>;
        ServiceRoleArn?: pulumi.Input<string>;
        Targets: pulumi.Input<pulumi.Input<MaintenanceWindowTaskTargetProperties>[]>;
        TaskArn: pulumi.Input<string>;
        TaskInvocationParameters?: pulumi.Input<MaintenanceWindowTaskTaskInvocationParametersProperties>;
        TaskParameters?: pulumi.Input<string>;
        TaskType: pulumi.Input<string>;
        WindowId: pulumi.Input<string>;
    }
    
    export interface MaintenanceWindowTaskTargetAttributes {
        Key: string;
        Values?: string[];
    }
    
    export interface MaintenanceWindowTaskTargetProperties {
        Key: pulumi.Input<string>;
        Values?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface MaintenanceWindowTaskTaskInvocationParametersAttributes {
        MaintenanceWindowAutomationParameters?: MaintenanceWindowTaskMaintenanceWindowAutomationParametersAttributes;
        MaintenanceWindowLambdaParameters?: MaintenanceWindowTaskMaintenanceWindowLambdaParametersAttributes;
        MaintenanceWindowRunCommandParameters?: MaintenanceWindowTaskMaintenanceWindowRunCommandParametersAttributes;
        MaintenanceWindowStepFunctionsParameters?: MaintenanceWindowTaskMaintenanceWindowStepFunctionsParametersAttributes;
    }
    
    export interface MaintenanceWindowTaskTaskInvocationParametersProperties {
        MaintenanceWindowAutomationParameters?: pulumi.Input<MaintenanceWindowTaskMaintenanceWindowAutomationParametersProperties>;
        MaintenanceWindowLambdaParameters?: pulumi.Input<MaintenanceWindowTaskMaintenanceWindowLambdaParametersProperties>;
        MaintenanceWindowRunCommandParameters?: pulumi.Input<MaintenanceWindowTaskMaintenanceWindowRunCommandParametersProperties>;
        MaintenanceWindowStepFunctionsParameters?: pulumi.Input<MaintenanceWindowTaskMaintenanceWindowStepFunctionsParametersProperties>;
    }
    
    export interface ParameterAttributes {
        Type: string;
        Value: string;
    }
    
    export interface ParameterProperties {
        AllowedPattern?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Policies?: pulumi.Input<string>;
        Tags?: pulumi.Input<string>;
        Tier?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface PatchBaselineAttributes {
    }
    
    export interface PatchBaselinePatchFilterAttributes {
        Key?: string;
        Values?: string[];
    }
    
    export interface PatchBaselinePatchFilterGroupAttributes {
        PatchFilters?: PatchBaselinePatchFilterAttributes[];
    }
    
    export interface PatchBaselinePatchFilterGroupProperties {
        PatchFilters?: pulumi.Input<pulumi.Input<PatchBaselinePatchFilterProperties>[]>;
    }
    
    export interface PatchBaselinePatchFilterProperties {
        Key?: pulumi.Input<string>;
        Values?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface PatchBaselinePatchSourceAttributes {
        Configuration?: string;
        Name?: string;
        Products?: string[];
    }
    
    export interface PatchBaselinePatchSourceProperties {
        Configuration?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        Products?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface PatchBaselineProperties {
        ApprovalRules?: pulumi.Input<PatchBaselineRuleGroupProperties>;
        ApprovedPatches?: pulumi.Input<pulumi.Input<string>[]>;
        ApprovedPatchesComplianceLevel?: pulumi.Input<string>;
        ApprovedPatchesEnableNonSecurity?: pulumi.Input<boolean>;
        Description?: pulumi.Input<string>;
        GlobalFilters?: pulumi.Input<PatchBaselinePatchFilterGroupProperties>;
        Name: pulumi.Input<string>;
        OperatingSystem?: pulumi.Input<string>;
        PatchGroups?: pulumi.Input<pulumi.Input<string>[]>;
        RejectedPatches?: pulumi.Input<pulumi.Input<string>[]>;
        RejectedPatchesAction?: pulumi.Input<string>;
        Sources?: pulumi.Input<pulumi.Input<PatchBaselinePatchSourceProperties>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface PatchBaselineRuleAttributes {
        ApproveAfterDays?: number;
        ComplianceLevel?: string;
        EnableNonSecurity?: boolean;
        PatchFilterGroup?: PatchBaselinePatchFilterGroupAttributes;
    }
    
    export interface PatchBaselineRuleGroupAttributes {
        PatchRules?: PatchBaselineRuleAttributes[];
    }
    
    export interface PatchBaselineRuleGroupProperties {
        PatchRules?: pulumi.Input<pulumi.Input<PatchBaselineRuleProperties>[]>;
    }
    
    export interface PatchBaselineRuleProperties {
        ApproveAfterDays?: pulumi.Input<number>;
        ComplianceLevel?: pulumi.Input<string>;
        EnableNonSecurity?: pulumi.Input<boolean>;
        PatchFilterGroup?: pulumi.Input<PatchBaselinePatchFilterGroupProperties>;
    }
    
    export interface ResourceDataSyncAttributes {
    }
    
    export interface ResourceDataSyncProperties {
        BucketName: pulumi.Input<string>;
        BucketPrefix?: pulumi.Input<string>;
        BucketRegion: pulumi.Input<string>;
        KMSKeyArn?: pulumi.Input<string>;
        SyncFormat: pulumi.Input<string>;
        SyncName: pulumi.Input<string>;
    }
    
    
    export class Association extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AssociationAttributes>;

        constructor(name: string, properties: AssociationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SSM:Association", name, inputs, opts);
        }
    }
    
    export class Document extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<DocumentAttributes>;

        constructor(name: string, properties: DocumentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SSM:Document", name, inputs, opts);
        }
    }
    
    export class MaintenanceWindow extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<MaintenanceWindowAttributes>;

        constructor(name: string, properties: MaintenanceWindowProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SSM:MaintenanceWindow", name, inputs, opts);
        }
    }
    
    export class MaintenanceWindowTarget extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<MaintenanceWindowTargetAttributes>;

        constructor(name: string, properties: MaintenanceWindowTargetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SSM:MaintenanceWindowTarget", name, inputs, opts);
        }
    }
    
    export class MaintenanceWindowTask extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<MaintenanceWindowTaskAttributes>;

        constructor(name: string, properties: MaintenanceWindowTaskProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SSM:MaintenanceWindowTask", name, inputs, opts);
        }
    }
    
    export class Parameter extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ParameterAttributes>;

        constructor(name: string, properties: ParameterProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SSM:Parameter", name, inputs, opts);
        }
    }
    
    export class PatchBaseline extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PatchBaselineAttributes>;

        constructor(name: string, properties: PatchBaselineProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SSM:PatchBaseline", name, inputs, opts);
        }
    }
    
    export class ResourceDataSync extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ResourceDataSyncAttributes>;

        constructor(name: string, properties: ResourceDataSyncProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SSM:ResourceDataSync", name, inputs, opts);
        }
    }
    
}

export namespace sagemaker {
    
    export interface CodeRepositoryAttributes {
        CodeRepositoryName: string;
    }
    
    export interface CodeRepositoryGitConfigAttributes {
        Branch?: string;
        RepositoryUrl: string;
        SecretArn?: string;
    }
    
    export interface CodeRepositoryGitConfigProperties {
        Branch?: pulumi.Input<string>;
        RepositoryUrl: pulumi.Input<string>;
        SecretArn?: pulumi.Input<string>;
    }
    
    export interface CodeRepositoryProperties {
        CodeRepositoryName?: pulumi.Input<string>;
        GitConfig: pulumi.Input<CodeRepositoryGitConfigProperties>;
    }
    
    export interface EndpointAttributes {
        EndpointName: string;
    }
    
    export interface EndpointConfigAttributes {
        EndpointConfigName: string;
    }
    
    export interface EndpointConfigProductionVariantAttributes {
        AcceleratorType?: string;
        InitialInstanceCount: number;
        InitialVariantWeight: number;
        InstanceType: string;
        ModelName: string;
        VariantName: string;
    }
    
    export interface EndpointConfigProductionVariantProperties {
        AcceleratorType?: pulumi.Input<string>;
        InitialInstanceCount: pulumi.Input<number>;
        InitialVariantWeight: pulumi.Input<number>;
        InstanceType: pulumi.Input<string>;
        ModelName: pulumi.Input<string>;
        VariantName: pulumi.Input<string>;
    }
    
    export interface EndpointConfigProperties {
        EndpointConfigName?: pulumi.Input<string>;
        KmsKeyId?: pulumi.Input<string>;
        ProductionVariants: pulumi.Input<pulumi.Input<EndpointConfigProductionVariantProperties>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface EndpointProperties {
        EndpointConfigName: pulumi.Input<string>;
        EndpointName?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface ModelAttributes {
        ModelName: string;
    }
    
    export interface ModelContainerDefinitionAttributes {
        ContainerHostname?: string;
        Environment?: string;
        Image: string;
        ModelDataUrl?: string;
    }
    
    export interface ModelContainerDefinitionProperties {
        ContainerHostname?: pulumi.Input<string>;
        Environment?: pulumi.Input<string>;
        Image: pulumi.Input<string>;
        ModelDataUrl?: pulumi.Input<string>;
    }
    
    export interface ModelProperties {
        Containers?: pulumi.Input<pulumi.Input<ModelContainerDefinitionProperties>[]>;
        ExecutionRoleArn: pulumi.Input<string>;
        ModelName?: pulumi.Input<string>;
        PrimaryContainer?: pulumi.Input<ModelContainerDefinitionProperties>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VpcConfig?: pulumi.Input<ModelVpcConfigProperties>;
    }
    
    export interface ModelVpcConfigAttributes {
        SecurityGroupIds: string[];
        Subnets: string[];
    }
    
    export interface ModelVpcConfigProperties {
        SecurityGroupIds: pulumi.Input<pulumi.Input<string>[]>;
        Subnets: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface NotebookInstanceAttributes {
        NotebookInstanceName: string;
    }
    
    export interface NotebookInstanceLifecycleConfigAttributes {
        NotebookInstanceLifecycleConfigName: string;
    }
    
    export interface NotebookInstanceLifecycleConfigNotebookInstanceLifecycleHookAttributes {
        Content?: string;
    }
    
    export interface NotebookInstanceLifecycleConfigNotebookInstanceLifecycleHookProperties {
        Content?: pulumi.Input<string>;
    }
    
    export interface NotebookInstanceLifecycleConfigProperties {
        NotebookInstanceLifecycleConfigName?: pulumi.Input<string>;
        OnCreate?: pulumi.Input<pulumi.Input<NotebookInstanceLifecycleConfigNotebookInstanceLifecycleHookProperties>[]>;
        OnStart?: pulumi.Input<pulumi.Input<NotebookInstanceLifecycleConfigNotebookInstanceLifecycleHookProperties>[]>;
    }
    
    export interface NotebookInstanceProperties {
        AcceleratorTypes?: pulumi.Input<pulumi.Input<string>[]>;
        AdditionalCodeRepositories?: pulumi.Input<pulumi.Input<string>[]>;
        DefaultCodeRepository?: pulumi.Input<string>;
        DirectInternetAccess?: pulumi.Input<string>;
        InstanceType: pulumi.Input<string>;
        KmsKeyId?: pulumi.Input<string>;
        LifecycleConfigName?: pulumi.Input<string>;
        NotebookInstanceName?: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
        RootAccess?: pulumi.Input<string>;
        SecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
        SubnetId?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        VolumeSizeInGB?: pulumi.Input<number>;
    }
    
    export interface WorkteamAttributes {
        WorkteamName: string;
    }
    
    export interface WorkteamCognitoMemberDefinitionAttributes {
        CognitoClientId: string;
        CognitoUserGroup: string;
        CognitoUserPool: string;
    }
    
    export interface WorkteamCognitoMemberDefinitionProperties {
        CognitoClientId: pulumi.Input<string>;
        CognitoUserGroup: pulumi.Input<string>;
        CognitoUserPool: pulumi.Input<string>;
    }
    
    export interface WorkteamMemberDefinitionAttributes {
        CognitoMemberDefinition: WorkteamCognitoMemberDefinitionAttributes;
    }
    
    export interface WorkteamMemberDefinitionProperties {
        CognitoMemberDefinition: pulumi.Input<WorkteamCognitoMemberDefinitionProperties>;
    }
    
    export interface WorkteamNotificationConfigurationAttributes {
        NotificationTopicArn: string;
    }
    
    export interface WorkteamNotificationConfigurationProperties {
        NotificationTopicArn: pulumi.Input<string>;
    }
    
    export interface WorkteamProperties {
        Description?: pulumi.Input<string>;
        MemberDefinitions?: pulumi.Input<pulumi.Input<WorkteamMemberDefinitionProperties>[]>;
        NotificationConfiguration?: pulumi.Input<WorkteamNotificationConfigurationProperties>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        WorkteamName?: pulumi.Input<string>;
    }
    
    
    export class CodeRepository extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<CodeRepositoryAttributes>;

        constructor(name: string, properties: CodeRepositoryProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SageMaker:CodeRepository", name, inputs, opts);
        }
    }
    
    export class Endpoint extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<EndpointAttributes>;

        constructor(name: string, properties: EndpointProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SageMaker:Endpoint", name, inputs, opts);
        }
    }
    
    export class EndpointConfig extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<EndpointConfigAttributes>;

        constructor(name: string, properties: EndpointConfigProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SageMaker:EndpointConfig", name, inputs, opts);
        }
    }
    
    export class Model extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ModelAttributes>;

        constructor(name: string, properties: ModelProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SageMaker:Model", name, inputs, opts);
        }
    }
    
    export class NotebookInstance extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<NotebookInstanceAttributes>;

        constructor(name: string, properties: NotebookInstanceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SageMaker:NotebookInstance", name, inputs, opts);
        }
    }
    
    export class NotebookInstanceLifecycleConfig extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<NotebookInstanceLifecycleConfigAttributes>;

        constructor(name: string, properties?: NotebookInstanceLifecycleConfigProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SageMaker:NotebookInstanceLifecycleConfig", name, inputs, opts);
        }
    }
    
    export class Workteam extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<WorkteamAttributes>;

        constructor(name: string, properties?: WorkteamProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SageMaker:Workteam", name, inputs, opts);
        }
    }
    
}

export namespace secretsmanager {
    
    export interface ResourcePolicyAttributes {
    }
    
    export interface ResourcePolicyProperties {
        ResourcePolicy: pulumi.Input<string>;
        SecretId: pulumi.Input<string>;
    }
    
    export interface RotationScheduleAttributes {
    }
    
    export interface RotationScheduleProperties {
        RotationLambdaARN?: pulumi.Input<string>;
        RotationRules?: pulumi.Input<RotationScheduleRotationRulesProperties>;
        SecretId: pulumi.Input<string>;
    }
    
    export interface RotationScheduleRotationRulesAttributes {
        AutomaticallyAfterDays?: number;
    }
    
    export interface RotationScheduleRotationRulesProperties {
        AutomaticallyAfterDays?: pulumi.Input<number>;
    }
    
    export interface SecretAttributes {
    }
    
    export interface SecretGenerateSecretStringAttributes {
        ExcludeCharacters?: string;
        ExcludeLowercase?: boolean;
        ExcludeNumbers?: boolean;
        ExcludePunctuation?: boolean;
        ExcludeUppercase?: boolean;
        GenerateStringKey?: string;
        IncludeSpace?: boolean;
        PasswordLength?: number;
        RequireEachIncludedType?: boolean;
        SecretStringTemplate?: string;
    }
    
    export interface SecretGenerateSecretStringProperties {
        ExcludeCharacters?: pulumi.Input<string>;
        ExcludeLowercase?: pulumi.Input<boolean>;
        ExcludeNumbers?: pulumi.Input<boolean>;
        ExcludePunctuation?: pulumi.Input<boolean>;
        ExcludeUppercase?: pulumi.Input<boolean>;
        GenerateStringKey?: pulumi.Input<string>;
        IncludeSpace?: pulumi.Input<boolean>;
        PasswordLength?: pulumi.Input<number>;
        RequireEachIncludedType?: pulumi.Input<boolean>;
        SecretStringTemplate?: pulumi.Input<string>;
    }
    
    export interface SecretProperties {
        Description?: pulumi.Input<string>;
        GenerateSecretString?: pulumi.Input<SecretGenerateSecretStringProperties>;
        KmsKeyId?: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
        SecretString?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface SecretTargetAttachmentAttributes {
    }
    
    export interface SecretTargetAttachmentProperties {
        SecretId: pulumi.Input<string>;
        TargetId: pulumi.Input<string>;
        TargetType: pulumi.Input<string>;
    }
    
    
    export class ResourcePolicy extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ResourcePolicyAttributes>;

        constructor(name: string, properties: ResourcePolicyProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SecretsManager:ResourcePolicy", name, inputs, opts);
        }
    }
    
    export class RotationSchedule extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RotationScheduleAttributes>;

        constructor(name: string, properties: RotationScheduleProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SecretsManager:RotationSchedule", name, inputs, opts);
        }
    }
    
    export class Secret extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SecretAttributes>;

        constructor(name: string, properties?: SecretProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SecretsManager:Secret", name, inputs, opts);
        }
    }
    
    export class SecretTargetAttachment extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SecretTargetAttachmentAttributes>;

        constructor(name: string, properties: SecretTargetAttachmentProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SecretsManager:SecretTargetAttachment", name, inputs, opts);
        }
    }
    
}

export namespace securityhub {
    
    export interface HubAttributes {
    }
    
    export interface HubProperties {
        Tags?: pulumi.Input<string>;
    }
    
    
    export class Hub extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<HubAttributes>;

        constructor(name: string, properties?: HubProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:SecurityHub:Hub", name, inputs, opts);
        }
    }
    
}

export namespace servicecatalog {
    
    export interface AcceptedPortfolioShareAttributes {
    }
    
    export interface AcceptedPortfolioShareProperties {
        AcceptLanguage?: pulumi.Input<string>;
        PortfolioId: pulumi.Input<string>;
    }
    
    export interface CloudFormationProductAttributes {
        ProductName: string;
        ProvisioningArtifactIds: string;
        ProvisioningArtifactNames: string;
    }
    
    export interface CloudFormationProductProperties {
        AcceptLanguage?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        Distributor?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        Owner: pulumi.Input<string>;
        ProvisioningArtifactParameters: pulumi.Input<pulumi.Input<CloudFormationProductProvisioningArtifactPropertiesProperties>[]>;
        SupportDescription?: pulumi.Input<string>;
        SupportEmail?: pulumi.Input<string>;
        SupportUrl?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface CloudFormationProductProvisioningArtifactPropertiesAttributes {
        Description?: string;
        DisableTemplateValidation?: boolean;
        Info: string;
        Name?: string;
    }
    
    export interface CloudFormationProductProvisioningArtifactPropertiesProperties {
        Description?: pulumi.Input<string>;
        DisableTemplateValidation?: pulumi.Input<boolean>;
        Info: pulumi.Input<string>;
        Name?: pulumi.Input<string>;
    }
    
    export interface CloudFormationProvisionedProductAttributes {
        CloudformationStackArn: string;
        RecordId: string;
    }
    
    export interface CloudFormationProvisionedProductProperties {
        AcceptLanguage?: pulumi.Input<string>;
        NotificationArns?: pulumi.Input<pulumi.Input<string>[]>;
        PathId?: pulumi.Input<string>;
        ProductId?: pulumi.Input<string>;
        ProductName?: pulumi.Input<string>;
        ProvisionedProductName?: pulumi.Input<string>;
        ProvisioningArtifactId?: pulumi.Input<string>;
        ProvisioningArtifactName?: pulumi.Input<string>;
        ProvisioningParameters?: pulumi.Input<pulumi.Input<CloudFormationProvisionedProductProvisioningParameterProperties>[]>;
        ProvisioningPreferences?: pulumi.Input<CloudFormationProvisionedProductProvisioningPreferencesProperties>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface CloudFormationProvisionedProductProvisioningParameterAttributes {
        Key?: string;
        Value?: string;
    }
    
    export interface CloudFormationProvisionedProductProvisioningParameterProperties {
        Key?: pulumi.Input<string>;
        Value?: pulumi.Input<string>;
    }
    
    export interface CloudFormationProvisionedProductProvisioningPreferencesAttributes {
        StackSetAccounts?: string[];
        StackSetFailureToleranceCount?: number;
        StackSetFailureTolerancePercentage?: number;
        StackSetMaxConcurrencyCount?: number;
        StackSetMaxConcurrencyPercentage?: number;
        StackSetOperationType?: string;
        StackSetRegions?: string[];
    }
    
    export interface CloudFormationProvisionedProductProvisioningPreferencesProperties {
        StackSetAccounts?: pulumi.Input<pulumi.Input<string>[]>;
        StackSetFailureToleranceCount?: pulumi.Input<number>;
        StackSetFailureTolerancePercentage?: pulumi.Input<number>;
        StackSetMaxConcurrencyCount?: pulumi.Input<number>;
        StackSetMaxConcurrencyPercentage?: pulumi.Input<number>;
        StackSetOperationType?: pulumi.Input<string>;
        StackSetRegions?: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface LaunchNotificationConstraintAttributes {
    }
    
    export interface LaunchNotificationConstraintProperties {
        AcceptLanguage?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        NotificationArns: pulumi.Input<pulumi.Input<string>[]>;
        PortfolioId: pulumi.Input<string>;
        ProductId: pulumi.Input<string>;
    }
    
    export interface LaunchRoleConstraintAttributes {
    }
    
    export interface LaunchRoleConstraintProperties {
        AcceptLanguage?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        PortfolioId: pulumi.Input<string>;
        ProductId: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
    }
    
    export interface LaunchTemplateConstraintAttributes {
    }
    
    export interface LaunchTemplateConstraintProperties {
        AcceptLanguage?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        PortfolioId: pulumi.Input<string>;
        ProductId: pulumi.Input<string>;
        Rules: pulumi.Input<string>;
    }
    
    export interface PortfolioAttributes {
        PortfolioName: string;
    }
    
    export interface PortfolioPrincipalAssociationAttributes {
    }
    
    export interface PortfolioPrincipalAssociationProperties {
        AcceptLanguage?: pulumi.Input<string>;
        PortfolioId: pulumi.Input<string>;
        PrincipalARN: pulumi.Input<string>;
        PrincipalType: pulumi.Input<string>;
    }
    
    export interface PortfolioProductAssociationAttributes {
    }
    
    export interface PortfolioProductAssociationProperties {
        AcceptLanguage?: pulumi.Input<string>;
        PortfolioId: pulumi.Input<string>;
        ProductId: pulumi.Input<string>;
        SourcePortfolioId?: pulumi.Input<string>;
    }
    
    export interface PortfolioProperties {
        AcceptLanguage?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        DisplayName: pulumi.Input<string>;
        ProviderName: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface PortfolioShareAttributes {
    }
    
    export interface PortfolioShareProperties {
        AcceptLanguage?: pulumi.Input<string>;
        AccountId: pulumi.Input<string>;
        PortfolioId: pulumi.Input<string>;
    }
    
    export interface ResourceUpdateConstraintAttributes {
    }
    
    export interface ResourceUpdateConstraintProperties {
        AcceptLanguage?: pulumi.Input<string>;
        Description?: pulumi.Input<string>;
        PortfolioId: pulumi.Input<string>;
        ProductId: pulumi.Input<string>;
        TagUpdateOnProvisionedProduct: pulumi.Input<string>;
    }
    
    export interface StackSetConstraintAttributes {
    }
    
    export interface StackSetConstraintProperties {
        AcceptLanguage?: pulumi.Input<string>;
        AccountList: pulumi.Input<pulumi.Input<string>[]>;
        AdminRole: pulumi.Input<string>;
        Description: pulumi.Input<string>;
        ExecutionRole: pulumi.Input<string>;
        PortfolioId: pulumi.Input<string>;
        ProductId: pulumi.Input<string>;
        RegionList: pulumi.Input<pulumi.Input<string>[]>;
        StackInstanceControl: pulumi.Input<string>;
    }
    
    export interface TagOptionAssociationAttributes {
    }
    
    export interface TagOptionAssociationProperties {
        ResourceId: pulumi.Input<string>;
        TagOptionId: pulumi.Input<string>;
    }
    
    export interface TagOptionAttributes {
    }
    
    export interface TagOptionProperties {
        Active?: pulumi.Input<boolean>;
        Key: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    
    export class AcceptedPortfolioShare extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<AcceptedPortfolioShareAttributes>;

        constructor(name: string, properties: AcceptedPortfolioShareProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceCatalog:AcceptedPortfolioShare", name, inputs, opts);
        }
    }
    
    export class CloudFormationProduct extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<CloudFormationProductAttributes>;

        constructor(name: string, properties: CloudFormationProductProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceCatalog:CloudFormationProduct", name, inputs, opts);
        }
    }
    
    export class CloudFormationProvisionedProduct extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<CloudFormationProvisionedProductAttributes>;

        constructor(name: string, properties?: CloudFormationProvisionedProductProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceCatalog:CloudFormationProvisionedProduct", name, inputs, opts);
        }
    }
    
    export class LaunchNotificationConstraint extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<LaunchNotificationConstraintAttributes>;

        constructor(name: string, properties: LaunchNotificationConstraintProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceCatalog:LaunchNotificationConstraint", name, inputs, opts);
        }
    }
    
    export class LaunchRoleConstraint extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<LaunchRoleConstraintAttributes>;

        constructor(name: string, properties: LaunchRoleConstraintProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceCatalog:LaunchRoleConstraint", name, inputs, opts);
        }
    }
    
    export class LaunchTemplateConstraint extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<LaunchTemplateConstraintAttributes>;

        constructor(name: string, properties: LaunchTemplateConstraintProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceCatalog:LaunchTemplateConstraint", name, inputs, opts);
        }
    }
    
    export class Portfolio extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PortfolioAttributes>;

        constructor(name: string, properties: PortfolioProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceCatalog:Portfolio", name, inputs, opts);
        }
    }
    
    export class PortfolioPrincipalAssociation extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PortfolioPrincipalAssociationAttributes>;

        constructor(name: string, properties: PortfolioPrincipalAssociationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceCatalog:PortfolioPrincipalAssociation", name, inputs, opts);
        }
    }
    
    export class PortfolioProductAssociation extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PortfolioProductAssociationAttributes>;

        constructor(name: string, properties: PortfolioProductAssociationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceCatalog:PortfolioProductAssociation", name, inputs, opts);
        }
    }
    
    export class PortfolioShare extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PortfolioShareAttributes>;

        constructor(name: string, properties: PortfolioShareProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceCatalog:PortfolioShare", name, inputs, opts);
        }
    }
    
    export class ResourceUpdateConstraint extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ResourceUpdateConstraintAttributes>;

        constructor(name: string, properties: ResourceUpdateConstraintProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceCatalog:ResourceUpdateConstraint", name, inputs, opts);
        }
    }
    
    export class StackSetConstraint extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<StackSetConstraintAttributes>;

        constructor(name: string, properties: StackSetConstraintProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceCatalog:StackSetConstraint", name, inputs, opts);
        }
    }
    
    export class TagOption extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TagOptionAttributes>;

        constructor(name: string, properties: TagOptionProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceCatalog:TagOption", name, inputs, opts);
        }
    }
    
    export class TagOptionAssociation extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<TagOptionAssociationAttributes>;

        constructor(name: string, properties: TagOptionAssociationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceCatalog:TagOptionAssociation", name, inputs, opts);
        }
    }
    
}

export namespace servicediscovery {
    
    export interface HttpNamespaceAttributes {
        Arn: string;
        Id: string;
    }
    
    export interface HttpNamespaceProperties {
        Description?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
    }
    
    export interface InstanceAttributes {
    }
    
    export interface InstanceProperties {
        InstanceAttributes: pulumi.Input<string>;
        InstanceId?: pulumi.Input<string>;
        ServiceId: pulumi.Input<string>;
    }
    
    export interface PrivateDnsNamespaceAttributes {
        Arn: string;
        Id: string;
    }
    
    export interface PrivateDnsNamespaceProperties {
        Description?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        Vpc: pulumi.Input<string>;
    }
    
    export interface PublicDnsNamespaceAttributes {
        Arn: string;
        Id: string;
    }
    
    export interface PublicDnsNamespaceProperties {
        Description?: pulumi.Input<string>;
        Name: pulumi.Input<string>;
    }
    
    export interface ServiceAttributes {
        Arn: string;
        Id: string;
        Name: string;
    }
    
    export interface ServiceDnsConfigAttributes {
        DnsRecords: ServiceDnsRecordAttributes[];
        NamespaceId?: string;
        RoutingPolicy?: string;
    }
    
    export interface ServiceDnsConfigProperties {
        DnsRecords: pulumi.Input<pulumi.Input<ServiceDnsRecordProperties>[]>;
        NamespaceId?: pulumi.Input<string>;
        RoutingPolicy?: pulumi.Input<string>;
    }
    
    export interface ServiceDnsRecordAttributes {
        TTL: number;
        Type: string;
    }
    
    export interface ServiceDnsRecordProperties {
        TTL: pulumi.Input<number>;
        Type: pulumi.Input<string>;
    }
    
    export interface ServiceHealthCheckConfigAttributes {
        FailureThreshold?: number;
        ResourcePath?: string;
        Type: string;
    }
    
    export interface ServiceHealthCheckConfigProperties {
        FailureThreshold?: pulumi.Input<number>;
        ResourcePath?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface ServiceHealthCheckCustomConfigAttributes {
        FailureThreshold?: number;
    }
    
    export interface ServiceHealthCheckCustomConfigProperties {
        FailureThreshold?: pulumi.Input<number>;
    }
    
    export interface ServiceProperties {
        Description?: pulumi.Input<string>;
        DnsConfig?: pulumi.Input<ServiceDnsConfigProperties>;
        HealthCheckConfig?: pulumi.Input<ServiceHealthCheckConfigProperties>;
        HealthCheckCustomConfig?: pulumi.Input<ServiceHealthCheckCustomConfigProperties>;
        Name?: pulumi.Input<string>;
        NamespaceId?: pulumi.Input<string>;
    }
    
    
    export class HttpNamespace extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<HttpNamespaceAttributes>;

        constructor(name: string, properties: HttpNamespaceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceDiscovery:HttpNamespace", name, inputs, opts);
        }
    }
    
    export class Instance extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<InstanceAttributes>;

        constructor(name: string, properties: InstanceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceDiscovery:Instance", name, inputs, opts);
        }
    }
    
    export class PrivateDnsNamespace extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PrivateDnsNamespaceAttributes>;

        constructor(name: string, properties: PrivateDnsNamespaceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceDiscovery:PrivateDnsNamespace", name, inputs, opts);
        }
    }
    
    export class PublicDnsNamespace extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<PublicDnsNamespaceAttributes>;

        constructor(name: string, properties: PublicDnsNamespaceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceDiscovery:PublicDnsNamespace", name, inputs, opts);
        }
    }
    
    export class Service extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ServiceAttributes>;

        constructor(name: string, properties?: ServiceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:ServiceDiscovery:Service", name, inputs, opts);
        }
    }
    
}

export namespace stepfunctions {
    
    export interface ActivityAttributes {
        Name: string;
    }
    
    export interface ActivityProperties {
        Name: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<ActivityTagsEntryProperties>[]>;
    }
    
    export interface ActivityTagsEntryAttributes {
        Key: string;
        Value: string;
    }
    
    export interface ActivityTagsEntryProperties {
        Key: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface StateMachineAttributes {
        Name: string;
    }
    
    export interface StateMachineProperties {
        DefinitionString: pulumi.Input<string>;
        RoleArn: pulumi.Input<string>;
        StateMachineName?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<StateMachineTagsEntryProperties>[]>;
    }
    
    export interface StateMachineTagsEntryAttributes {
        Key: string;
        Value: string;
    }
    
    export interface StateMachineTagsEntryProperties {
        Key: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    
    export class Activity extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ActivityAttributes>;

        constructor(name: string, properties: ActivityProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:StepFunctions:Activity", name, inputs, opts);
        }
    }
    
    export class StateMachine extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<StateMachineAttributes>;

        constructor(name: string, properties: StateMachineProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:StepFunctions:StateMachine", name, inputs, opts);
        }
    }
    
}

export namespace transfer {
    
    export interface ServerAttributes {
        Arn: string;
        ServerId: string;
    }
    
    export interface ServerEndpointDetailsAttributes {
        VpcEndpointId: string;
    }
    
    export interface ServerEndpointDetailsProperties {
        VpcEndpointId: pulumi.Input<string>;
    }
    
    export interface ServerIdentityProviderDetailsAttributes {
        InvocationRole: string;
        Url: string;
    }
    
    export interface ServerIdentityProviderDetailsProperties {
        InvocationRole: pulumi.Input<string>;
        Url: pulumi.Input<string>;
    }
    
    export interface ServerProperties {
        EndpointDetails?: pulumi.Input<ServerEndpointDetailsProperties>;
        EndpointType?: pulumi.Input<string>;
        IdentityProviderDetails?: pulumi.Input<ServerIdentityProviderDetailsProperties>;
        IdentityProviderType?: pulumi.Input<string>;
        LoggingRole?: pulumi.Input<string>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
    }
    
    export interface UserAttributes {
        Arn: string;
        ServerId: string;
        UserName: string;
    }
    
    export interface UserProperties {
        HomeDirectory?: pulumi.Input<string>;
        Policy?: pulumi.Input<string>;
        Role: pulumi.Input<string>;
        ServerId: pulumi.Input<string>;
        SshPublicKeys?: pulumi.Input<pulumi.Input<UserSshPublicKeyProperties>[]>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        UserName: pulumi.Input<string>;
    }
    
    export interface UserSshPublicKeyAttributes {
    }
    
    export interface UserSshPublicKeyProperties {
    }
    
    
    export class Server extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ServerAttributes>;

        constructor(name: string, properties?: ServerProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Transfer:Server", name, inputs, opts);
        }
    }
    
    export class User extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<UserAttributes>;

        constructor(name: string, properties: UserProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:Transfer:User", name, inputs, opts);
        }
    }
    
}

export namespace waf {
    
    export interface ByteMatchSetAttributes {
    }
    
    export interface ByteMatchSetByteMatchTupleAttributes {
        FieldToMatch: ByteMatchSetFieldToMatchAttributes;
        PositionalConstraint: string;
        TargetString?: string;
        TargetStringBase64?: string;
        TextTransformation: string;
    }
    
    export interface ByteMatchSetByteMatchTupleProperties {
        FieldToMatch: pulumi.Input<ByteMatchSetFieldToMatchProperties>;
        PositionalConstraint: pulumi.Input<string>;
        TargetString?: pulumi.Input<string>;
        TargetStringBase64?: pulumi.Input<string>;
        TextTransformation: pulumi.Input<string>;
    }
    
    export interface ByteMatchSetFieldToMatchAttributes {
        Data?: string;
        Type: string;
    }
    
    export interface ByteMatchSetFieldToMatchProperties {
        Data?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface ByteMatchSetProperties {
        ByteMatchTuples?: pulumi.Input<pulumi.Input<ByteMatchSetByteMatchTupleProperties>[]>;
        Name: pulumi.Input<string>;
    }
    
    export interface IPSetAttributes {
    }
    
    export interface IPSetIPSetDescriptorAttributes {
        Type: string;
        Value: string;
    }
    
    export interface IPSetIPSetDescriptorProperties {
        Type: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface IPSetProperties {
        IPSetDescriptors?: pulumi.Input<pulumi.Input<IPSetIPSetDescriptorProperties>[]>;
        Name: pulumi.Input<string>;
    }
    
    export interface RuleAttributes {
    }
    
    export interface RulePredicateAttributes {
        DataId: string;
        Negated: boolean;
        Type: string;
    }
    
    export interface RulePredicateProperties {
        DataId: pulumi.Input<string>;
        Negated: pulumi.Input<boolean>;
        Type: pulumi.Input<string>;
    }
    
    export interface RuleProperties {
        MetricName: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        Predicates?: pulumi.Input<pulumi.Input<RulePredicateProperties>[]>;
    }
    
    export interface SizeConstraintSetAttributes {
    }
    
    export interface SizeConstraintSetFieldToMatchAttributes {
        Data?: string;
        Type: string;
    }
    
    export interface SizeConstraintSetFieldToMatchProperties {
        Data?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface SizeConstraintSetProperties {
        Name: pulumi.Input<string>;
        SizeConstraints: pulumi.Input<pulumi.Input<SizeConstraintSetSizeConstraintProperties>[]>;
    }
    
    export interface SizeConstraintSetSizeConstraintAttributes {
        ComparisonOperator: string;
        FieldToMatch: SizeConstraintSetFieldToMatchAttributes;
        Size: number;
        TextTransformation: string;
    }
    
    export interface SizeConstraintSetSizeConstraintProperties {
        ComparisonOperator: pulumi.Input<string>;
        FieldToMatch: pulumi.Input<SizeConstraintSetFieldToMatchProperties>;
        Size: pulumi.Input<number>;
        TextTransformation: pulumi.Input<string>;
    }
    
    export interface SqlInjectionMatchSetAttributes {
    }
    
    export interface SqlInjectionMatchSetFieldToMatchAttributes {
        Data?: string;
        Type: string;
    }
    
    export interface SqlInjectionMatchSetFieldToMatchProperties {
        Data?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface SqlInjectionMatchSetProperties {
        Name: pulumi.Input<string>;
        SqlInjectionMatchTuples?: pulumi.Input<pulumi.Input<SqlInjectionMatchSetSqlInjectionMatchTupleProperties>[]>;
    }
    
    export interface SqlInjectionMatchSetSqlInjectionMatchTupleAttributes {
        FieldToMatch: SqlInjectionMatchSetFieldToMatchAttributes;
        TextTransformation: string;
    }
    
    export interface SqlInjectionMatchSetSqlInjectionMatchTupleProperties {
        FieldToMatch: pulumi.Input<SqlInjectionMatchSetFieldToMatchProperties>;
        TextTransformation: pulumi.Input<string>;
    }
    
    export interface WebACLActivatedRuleAttributes {
        Action?: WebACLWafActionAttributes;
        Priority: number;
        RuleId: string;
    }
    
    export interface WebACLActivatedRuleProperties {
        Action?: pulumi.Input<WebACLWafActionProperties>;
        Priority: pulumi.Input<number>;
        RuleId: pulumi.Input<string>;
    }
    
    export interface WebACLAttributes {
    }
    
    export interface WebACLProperties {
        DefaultAction: pulumi.Input<WebACLWafActionProperties>;
        MetricName: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        Rules?: pulumi.Input<pulumi.Input<WebACLActivatedRuleProperties>[]>;
    }
    
    export interface WebACLWafActionAttributes {
        Type: string;
    }
    
    export interface WebACLWafActionProperties {
        Type: pulumi.Input<string>;
    }
    
    export interface XssMatchSetAttributes {
    }
    
    export interface XssMatchSetFieldToMatchAttributes {
        Data?: string;
        Type: string;
    }
    
    export interface XssMatchSetFieldToMatchProperties {
        Data?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface XssMatchSetProperties {
        Name: pulumi.Input<string>;
        XssMatchTuples: pulumi.Input<pulumi.Input<XssMatchSetXssMatchTupleProperties>[]>;
    }
    
    export interface XssMatchSetXssMatchTupleAttributes {
        FieldToMatch: XssMatchSetFieldToMatchAttributes;
        TextTransformation: string;
    }
    
    export interface XssMatchSetXssMatchTupleProperties {
        FieldToMatch: pulumi.Input<XssMatchSetFieldToMatchProperties>;
        TextTransformation: pulumi.Input<string>;
    }
    
    
    export class ByteMatchSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ByteMatchSetAttributes>;

        constructor(name: string, properties: ByteMatchSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAF:ByteMatchSet", name, inputs, opts);
        }
    }
    
    export class IPSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<IPSetAttributes>;

        constructor(name: string, properties: IPSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAF:IPSet", name, inputs, opts);
        }
    }
    
    export class Rule extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RuleAttributes>;

        constructor(name: string, properties: RuleProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAF:Rule", name, inputs, opts);
        }
    }
    
    export class SizeConstraintSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SizeConstraintSetAttributes>;

        constructor(name: string, properties: SizeConstraintSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAF:SizeConstraintSet", name, inputs, opts);
        }
    }
    
    export class SqlInjectionMatchSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SqlInjectionMatchSetAttributes>;

        constructor(name: string, properties: SqlInjectionMatchSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAF:SqlInjectionMatchSet", name, inputs, opts);
        }
    }
    
    export class WebACL extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<WebACLAttributes>;

        constructor(name: string, properties: WebACLProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAF:WebACL", name, inputs, opts);
        }
    }
    
    export class XssMatchSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<XssMatchSetAttributes>;

        constructor(name: string, properties: XssMatchSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAF:XssMatchSet", name, inputs, opts);
        }
    }
    
}

export namespace wafregional {
    
    export interface ByteMatchSetAttributes {
    }
    
    export interface ByteMatchSetByteMatchTupleAttributes {
        FieldToMatch: ByteMatchSetFieldToMatchAttributes;
        PositionalConstraint: string;
        TargetString?: string;
        TargetStringBase64?: string;
        TextTransformation: string;
    }
    
    export interface ByteMatchSetByteMatchTupleProperties {
        FieldToMatch: pulumi.Input<ByteMatchSetFieldToMatchProperties>;
        PositionalConstraint: pulumi.Input<string>;
        TargetString?: pulumi.Input<string>;
        TargetStringBase64?: pulumi.Input<string>;
        TextTransformation: pulumi.Input<string>;
    }
    
    export interface ByteMatchSetFieldToMatchAttributes {
        Data?: string;
        Type: string;
    }
    
    export interface ByteMatchSetFieldToMatchProperties {
        Data?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface ByteMatchSetProperties {
        ByteMatchTuples?: pulumi.Input<pulumi.Input<ByteMatchSetByteMatchTupleProperties>[]>;
        Name: pulumi.Input<string>;
    }
    
    export interface GeoMatchSetAttributes {
    }
    
    export interface GeoMatchSetGeoMatchConstraintAttributes {
        Type: string;
        Value: string;
    }
    
    export interface GeoMatchSetGeoMatchConstraintProperties {
        Type: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface GeoMatchSetProperties {
        GeoMatchConstraints?: pulumi.Input<pulumi.Input<GeoMatchSetGeoMatchConstraintProperties>[]>;
        Name: pulumi.Input<string>;
    }
    
    export interface IPSetAttributes {
    }
    
    export interface IPSetIPSetDescriptorAttributes {
        Type: string;
        Value: string;
    }
    
    export interface IPSetIPSetDescriptorProperties {
        Type: pulumi.Input<string>;
        Value: pulumi.Input<string>;
    }
    
    export interface IPSetProperties {
        IPSetDescriptors?: pulumi.Input<pulumi.Input<IPSetIPSetDescriptorProperties>[]>;
        Name: pulumi.Input<string>;
    }
    
    export interface RateBasedRuleAttributes {
    }
    
    export interface RateBasedRulePredicateAttributes {
        DataId: string;
        Negated: boolean;
        Type: string;
    }
    
    export interface RateBasedRulePredicateProperties {
        DataId: pulumi.Input<string>;
        Negated: pulumi.Input<boolean>;
        Type: pulumi.Input<string>;
    }
    
    export interface RateBasedRuleProperties {
        MatchPredicates?: pulumi.Input<pulumi.Input<RateBasedRulePredicateProperties>[]>;
        MetricName: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        RateKey: pulumi.Input<string>;
        RateLimit: pulumi.Input<number>;
    }
    
    export interface RegexPatternSetAttributes {
    }
    
    export interface RegexPatternSetProperties {
        Name: pulumi.Input<string>;
        RegexPatternStrings: pulumi.Input<pulumi.Input<string>[]>;
    }
    
    export interface RuleAttributes {
    }
    
    export interface RulePredicateAttributes {
        DataId: string;
        Negated: boolean;
        Type: string;
    }
    
    export interface RulePredicateProperties {
        DataId: pulumi.Input<string>;
        Negated: pulumi.Input<boolean>;
        Type: pulumi.Input<string>;
    }
    
    export interface RuleProperties {
        MetricName: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        Predicates?: pulumi.Input<pulumi.Input<RulePredicateProperties>[]>;
    }
    
    export interface SizeConstraintSetAttributes {
    }
    
    export interface SizeConstraintSetFieldToMatchAttributes {
        Data?: string;
        Type: string;
    }
    
    export interface SizeConstraintSetFieldToMatchProperties {
        Data?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface SizeConstraintSetProperties {
        Name: pulumi.Input<string>;
        SizeConstraints?: pulumi.Input<pulumi.Input<SizeConstraintSetSizeConstraintProperties>[]>;
    }
    
    export interface SizeConstraintSetSizeConstraintAttributes {
        ComparisonOperator: string;
        FieldToMatch: SizeConstraintSetFieldToMatchAttributes;
        Size: number;
        TextTransformation: string;
    }
    
    export interface SizeConstraintSetSizeConstraintProperties {
        ComparisonOperator: pulumi.Input<string>;
        FieldToMatch: pulumi.Input<SizeConstraintSetFieldToMatchProperties>;
        Size: pulumi.Input<number>;
        TextTransformation: pulumi.Input<string>;
    }
    
    export interface SqlInjectionMatchSetAttributes {
    }
    
    export interface SqlInjectionMatchSetFieldToMatchAttributes {
        Data?: string;
        Type: string;
    }
    
    export interface SqlInjectionMatchSetFieldToMatchProperties {
        Data?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface SqlInjectionMatchSetProperties {
        Name: pulumi.Input<string>;
        SqlInjectionMatchTuples?: pulumi.Input<pulumi.Input<SqlInjectionMatchSetSqlInjectionMatchTupleProperties>[]>;
    }
    
    export interface SqlInjectionMatchSetSqlInjectionMatchTupleAttributes {
        FieldToMatch: SqlInjectionMatchSetFieldToMatchAttributes;
        TextTransformation: string;
    }
    
    export interface SqlInjectionMatchSetSqlInjectionMatchTupleProperties {
        FieldToMatch: pulumi.Input<SqlInjectionMatchSetFieldToMatchProperties>;
        TextTransformation: pulumi.Input<string>;
    }
    
    export interface WebACLActionAttributes {
        Type: string;
    }
    
    export interface WebACLActionProperties {
        Type: pulumi.Input<string>;
    }
    
    export interface WebACLAssociationAttributes {
    }
    
    export interface WebACLAssociationProperties {
        ResourceArn: pulumi.Input<string>;
        WebACLId: pulumi.Input<string>;
    }
    
    export interface WebACLAttributes {
    }
    
    export interface WebACLProperties {
        DefaultAction: pulumi.Input<WebACLActionProperties>;
        MetricName: pulumi.Input<string>;
        Name: pulumi.Input<string>;
        Rules?: pulumi.Input<pulumi.Input<WebACLRuleProperties>[]>;
    }
    
    export interface WebACLRuleAttributes {
        Action: WebACLActionAttributes;
        Priority: number;
        RuleId: string;
    }
    
    export interface WebACLRuleProperties {
        Action: pulumi.Input<WebACLActionProperties>;
        Priority: pulumi.Input<number>;
        RuleId: pulumi.Input<string>;
    }
    
    export interface XssMatchSetAttributes {
    }
    
    export interface XssMatchSetFieldToMatchAttributes {
        Data?: string;
        Type: string;
    }
    
    export interface XssMatchSetFieldToMatchProperties {
        Data?: pulumi.Input<string>;
        Type: pulumi.Input<string>;
    }
    
    export interface XssMatchSetProperties {
        Name: pulumi.Input<string>;
        XssMatchTuples?: pulumi.Input<pulumi.Input<XssMatchSetXssMatchTupleProperties>[]>;
    }
    
    export interface XssMatchSetXssMatchTupleAttributes {
        FieldToMatch: XssMatchSetFieldToMatchAttributes;
        TextTransformation: string;
    }
    
    export interface XssMatchSetXssMatchTupleProperties {
        FieldToMatch: pulumi.Input<XssMatchSetFieldToMatchProperties>;
        TextTransformation: pulumi.Input<string>;
    }
    
    
    export class ByteMatchSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<ByteMatchSetAttributes>;

        constructor(name: string, properties: ByteMatchSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAFRegional:ByteMatchSet", name, inputs, opts);
        }
    }
    
    export class GeoMatchSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<GeoMatchSetAttributes>;

        constructor(name: string, properties: GeoMatchSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAFRegional:GeoMatchSet", name, inputs, opts);
        }
    }
    
    export class IPSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<IPSetAttributes>;

        constructor(name: string, properties: IPSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAFRegional:IPSet", name, inputs, opts);
        }
    }
    
    export class RateBasedRule extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RateBasedRuleAttributes>;

        constructor(name: string, properties: RateBasedRuleProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAFRegional:RateBasedRule", name, inputs, opts);
        }
    }
    
    export class RegexPatternSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RegexPatternSetAttributes>;

        constructor(name: string, properties: RegexPatternSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAFRegional:RegexPatternSet", name, inputs, opts);
        }
    }
    
    export class Rule extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<RuleAttributes>;

        constructor(name: string, properties: RuleProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAFRegional:Rule", name, inputs, opts);
        }
    }
    
    export class SizeConstraintSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SizeConstraintSetAttributes>;

        constructor(name: string, properties: SizeConstraintSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAFRegional:SizeConstraintSet", name, inputs, opts);
        }
    }
    
    export class SqlInjectionMatchSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<SqlInjectionMatchSetAttributes>;

        constructor(name: string, properties: SqlInjectionMatchSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAFRegional:SqlInjectionMatchSet", name, inputs, opts);
        }
    }
    
    export class WebACL extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<WebACLAttributes>;

        constructor(name: string, properties: WebACLProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAFRegional:WebACL", name, inputs, opts);
        }
    }
    
    export class WebACLAssociation extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<WebACLAssociationAttributes>;

        constructor(name: string, properties: WebACLAssociationProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAFRegional:WebACLAssociation", name, inputs, opts);
        }
    }
    
    export class XssMatchSet extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<XssMatchSetAttributes>;

        constructor(name: string, properties: XssMatchSetProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WAFRegional:XssMatchSet", name, inputs, opts);
        }
    }
    
}

export namespace workspaces {
    
    export interface WorkspaceAttributes {
    }
    
    export interface WorkspaceProperties {
        BundleId: pulumi.Input<string>;
        DirectoryId: pulumi.Input<string>;
        RootVolumeEncryptionEnabled?: pulumi.Input<boolean>;
        Tags?: pulumi.Input<pulumi.Input<TagProperties>[]>;
        UserName: pulumi.Input<string>;
        UserVolumeEncryptionEnabled?: pulumi.Input<boolean>;
        VolumeEncryptionKey?: pulumi.Input<string>;
        WorkspaceProperties?: pulumi.Input<WorkspaceWorkspacePropertiesProperties>;
    }
    
    export interface WorkspaceWorkspacePropertiesAttributes {
        ComputeTypeName?: string;
        RootVolumeSizeGib?: number;
        RunningMode?: string;
        RunningModeAutoStopTimeoutInMinutes?: number;
        UserVolumeSizeGib?: number;
    }
    
    export interface WorkspaceWorkspacePropertiesProperties {
        ComputeTypeName?: pulumi.Input<string>;
        RootVolumeSizeGib?: pulumi.Input<number>;
        RunningMode?: pulumi.Input<string>;
        RunningModeAutoStopTimeoutInMinutes?: pulumi.Input<number>;
        UserVolumeSizeGib?: pulumi.Input<number>;
    }
    
    
    export class Workspace extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<WorkspaceAttributes>;

        constructor(name: string, properties: WorkspaceProperties, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("cloudformation:WorkSpaces:Workspace", name, inputs, opts);
        }
    }
    
}

