// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

import {Region} from "./index";

/**
 * The provider type for the AWS Cloud Control package. By default, resources use package-wide configuration settings, however an explicit `Provider` instance may be created and passed during resource construction to achieve fine-grained programmatic control over provider settings. See the [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'aws-native';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * The profile for API operations. If not set, the default profile created with `aws configure` will be used.
     */
    public readonly profile!: pulumi.Output<string | undefined>;
    /**
     * The region where AWS operations will take place. Examples are `us-east-1`, `us-west-2`, etc.
     */
    public readonly region!: pulumi.Output<Region | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role for Cloud Control API to use when performing this resource operation. Note, this is a unique feature for server side security enforcement, not to be confused with assumeRole, which is used to obtain temporary client credentials. If you do not specify a role, Cloud Control API uses a temporary session created using your AWS user credentials instead.
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;
    /**
     * The path to the shared credentials file. If not set this defaults to `~/.aws/credentials`.
     */
    public readonly sharedCredentialsFile!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["accessKey"] = args?.accessKey ? pulumi.secret(args.accessKey) : undefined;
            resourceInputs["allowedAccountIds"] = pulumi.output(args ? args.allowedAccountIds : undefined).apply(JSON.stringify);
            resourceInputs["assumeRole"] = pulumi.output(args ? args.assumeRole : undefined).apply(JSON.stringify);
            resourceInputs["autoNaming"] = pulumi.output(args ? (args.autoNaming ? pulumi.output(args.autoNaming).apply(inputs.providerAutoNamingArgsProvideDefaults) : undefined) : undefined).apply(JSON.stringify);
            resourceInputs["defaultTags"] = pulumi.output(args ? args.defaultTags : undefined).apply(JSON.stringify);
            resourceInputs["endpoints"] = pulumi.output(args ? args.endpoints : undefined).apply(JSON.stringify);
            resourceInputs["forbiddenAccountIds"] = pulumi.output(args ? args.forbiddenAccountIds : undefined).apply(JSON.stringify);
            resourceInputs["ignoreTags"] = pulumi.output(args ? args.ignoreTags : undefined).apply(JSON.stringify);
            resourceInputs["insecure"] = pulumi.output(args ? args.insecure : undefined).apply(JSON.stringify);
            resourceInputs["maxRetries"] = pulumi.output(args ? args.maxRetries : undefined).apply(JSON.stringify);
            resourceInputs["profile"] = (args ? args.profile : undefined) ?? utilities.getEnv("AWS_PROFILE");
            resourceInputs["region"] = (args ? args.region : undefined) ?? <any>utilities.getEnv("AWS_REGION", "AWS_DEFAULT_REGION");
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["s3ForcePathStyle"] = pulumi.output(args ? args.s3ForcePathStyle : undefined).apply(JSON.stringify);
            resourceInputs["secretKey"] = args?.secretKey ? pulumi.secret(args.secretKey) : undefined;
            resourceInputs["sharedCredentialsFile"] = (args ? args.sharedCredentialsFile : undefined) ?? utilities.getEnv("AWS_SHARED_CREDENTIALS_FILE");
            resourceInputs["skipCredentialsValidation"] = pulumi.output((args ? args.skipCredentialsValidation : undefined) ?? true).apply(JSON.stringify);
            resourceInputs["skipGetEc2Platforms"] = pulumi.output((args ? args.skipGetEc2Platforms : undefined) ?? true).apply(JSON.stringify);
            resourceInputs["skipMetadataApiCheck"] = pulumi.output((args ? args.skipMetadataApiCheck : undefined) ?? true).apply(JSON.stringify);
            resourceInputs["skipRegionValidation"] = pulumi.output((args ? args.skipRegionValidation : undefined) ?? true).apply(JSON.stringify);
            resourceInputs["skipRequestingAccountId"] = pulumi.output(args ? args.skipRequestingAccountId : undefined).apply(JSON.stringify);
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * The access key for API operations. You can retrieve this from the ‘Security & Credentials’ section of the AWS console.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * List of allowed AWS account IDs to prevent you from mistakenly using an incorrect one. Conflicts with `forbiddenAccountIds`.
     */
    allowedAccountIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configuration for retrieving temporary credentials from the STS service.
     */
    assumeRole?: pulumi.Input<inputs.ProviderAssumeRoleArgs>;
    /**
     * The configuration for automatically naming resources.
     */
    autoNaming?: pulumi.Input<inputs.ProviderAutoNamingArgs>;
    /**
     * Configuration block with resource tag settings to apply across all resources handled by this provider. This is designed to replace redundant per-resource `tags` configurations. Provider tags can be overridden with new values, but not excluded from specific resources. To override provider tag values, use the `tags` argument within a resource to configure new tag values for matching keys.
     */
    defaultTags?: pulumi.Input<inputs.ProviderDefaultTagsArgs>;
    /**
     * Configuration block for customizing service endpoints.
     */
    endpoints?: pulumi.Input<pulumi.Input<inputs.ProviderEndpointArgs>[]>;
    /**
     * List of forbidden AWS account IDs to prevent you from mistakenly using the wrong one (and potentially end up destroying a live environment). Conflicts with `allowedAccountIds`.
     */
    forbiddenAccountIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configuration block with resource tag settings to ignore across all resources handled by this provider (except any individual service tag resources such as `ec2.Tag`) for situations where external systems are managing certain resource tags.
     */
    ignoreTags?: pulumi.Input<inputs.ProviderIgnoreTagsArgs>;
    /**
     * Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`.
     */
    insecure?: pulumi.Input<boolean>;
    /**
     * The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.
     */
    maxRetries?: pulumi.Input<number>;
    /**
     * The profile for API operations. If not set, the default profile created with `aws configure` will be used.
     */
    profile?: pulumi.Input<string>;
    /**
     * The region where AWS operations will take place. Examples are `us-east-1`, `us-west-2`, etc.
     */
    region: pulumi.Input<Region>;
    /**
     * The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role for Cloud Control API to use when performing this resource operation. Note, this is a unique feature for server side security enforcement, not to be confused with assumeRole, which is used to obtain temporary client credentials. If you do not specify a role, Cloud Control API uses a temporary session created using your AWS user credentials instead.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * Set this to true to force the request to use path-style addressing, i.e., `http://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client will use virtual hosted bucket addressing when possible (`http://BUCKET.s3.amazonaws.com/KEY`). Specific to the Amazon S3 service.
     */
    s3ForcePathStyle?: pulumi.Input<boolean>;
    /**
     * The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * The path to the shared credentials file. If not set this defaults to `~/.aws/credentials`.
     */
    sharedCredentialsFile?: pulumi.Input<string>;
    /**
     * Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS available/implemented.
     */
    skipCredentialsValidation?: pulumi.Input<boolean>;
    /**
     * Skip getting the supported EC2 platforms. Used by users that don't have `ec2:DescribeAccountAttributes` permissions.
     */
    skipGetEc2Platforms?: pulumi.Input<boolean>;
    /**
     * Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to true prevents Pulumi from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.
     */
    skipMetadataApiCheck?: pulumi.Input<boolean>;
    /**
     * Skip static validation of region name. Used by users of alternative AWS-like APIs or users with access to regions that are not public.
     */
    skipRegionValidation?: pulumi.Input<boolean>;
    /**
     * Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.
     */
    skipRequestingAccountId?: pulumi.Input<boolean>;
    /**
     * Session token for validating temporary credentials. Typically provided after successful identity federation or Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit MFA code used to get temporary credentials.
     */
    token?: pulumi.Input<string>;
}
