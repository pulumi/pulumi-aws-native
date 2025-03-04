// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The request to create a new origin access identity (OAI). An origin access identity is a special CloudFront user that you can associate with Amazon S3 origins, so that you can secure all or just some of your Amazon S3 content. For more information, see [Restricting Access to Amazon S3 Content by Using an Origin Access Identity](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html) in the *Amazon CloudFront Developer Guide*.
 *
 * ## Example Usage
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const cloudfrontoriginaccessidentity = new aws_native.cloudfront.CloudFrontOriginAccessIdentity("cloudfrontoriginaccessidentity", {cloudFrontOriginAccessIdentityConfig: {
 *     comment: "string-value",
 * }});
 *
 * ```
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const cloudfrontoriginaccessidentity = new aws_native.cloudfront.CloudFrontOriginAccessIdentity("cloudfrontoriginaccessidentity", {cloudFrontOriginAccessIdentityConfig: {
 *     comment: "string-value",
 * }});
 *
 * ```
 */
export class CloudFrontOriginAccessIdentity extends pulumi.CustomResource {
    /**
     * Get an existing CloudFrontOriginAccessIdentity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CloudFrontOriginAccessIdentity {
        return new CloudFrontOriginAccessIdentity(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cloudfront:CloudFrontOriginAccessIdentity';

    /**
     * Returns true if the given object is an instance of CloudFrontOriginAccessIdentity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudFrontOriginAccessIdentity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudFrontOriginAccessIdentity.__pulumiType;
    }

    /**
     * The ID for the origin access identity, for example, `E74FTE3AJFJ256A` .
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * The current configuration information for the identity.
     */
    public readonly cloudFrontOriginAccessIdentityConfig!: pulumi.Output<outputs.cloudfront.CloudFrontOriginAccessIdentityConfig>;
    /**
     * The Amazon S3 canonical user ID for the origin access identity, used when giving the origin access identity read permission to an object in Amazon S3. For example: `b970b42360b81c8ddbd79d2f5df0069ba9033c8a79655752abe380cd6d63ba8bcf23384d568fcf89fc49700b5e11a0fd` .
     */
    public /*out*/ readonly s3CanonicalUserId!: pulumi.Output<string>;

    /**
     * Create a CloudFrontOriginAccessIdentity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudFrontOriginAccessIdentityArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.cloudFrontOriginAccessIdentityConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudFrontOriginAccessIdentityConfig'");
            }
            resourceInputs["cloudFrontOriginAccessIdentityConfig"] = args ? args.cloudFrontOriginAccessIdentityConfig : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["s3CanonicalUserId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["cloudFrontOriginAccessIdentityConfig"] = undefined /*out*/;
            resourceInputs["s3CanonicalUserId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CloudFrontOriginAccessIdentity.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CloudFrontOriginAccessIdentity resource.
 */
export interface CloudFrontOriginAccessIdentityArgs {
    /**
     * The current configuration information for the identity.
     */
    cloudFrontOriginAccessIdentityConfig: pulumi.Input<inputs.cloudfront.CloudFrontOriginAccessIdentityConfigArgs>;
}
