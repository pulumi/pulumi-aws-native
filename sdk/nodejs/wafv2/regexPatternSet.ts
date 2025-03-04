// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Contains a list of Regular expressions based on the provided inputs. RegexPatternSet can be used with other WAF entities with RegexPatternSetReferenceStatement to perform other actions .
 */
export class RegexPatternSet extends pulumi.CustomResource {
    /**
     * Get an existing RegexPatternSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegexPatternSet {
        return new RegexPatternSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:wafv2:RegexPatternSet';

    /**
     * Returns true if the given object is an instance of RegexPatternSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegexPatternSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegexPatternSet.__pulumiType;
    }

    /**
     * ARN of the WAF entity.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Id of the RegexPatternSet
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * Description of the entity.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the RegexPatternSet.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The regular expression patterns in the set.
     */
    public readonly regularExpressionList!: pulumi.Output<string[]>;
    /**
     * Use CLOUDFRONT for CloudFront RegexPatternSet, use REGIONAL for Application Load Balancer and API Gateway.
     */
    public readonly scope!: pulumi.Output<enums.wafv2.RegexPatternSetScope>;
    /**
     * Key:value pairs associated with an AWS resource. The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
     *
     * > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a RegexPatternSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegexPatternSetArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.regularExpressionList === undefined) && !opts.urn) {
                throw new Error("Missing required property 'regularExpressionList'");
            }
            if ((!args || args.scope === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scope'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["regularExpressionList"] = args ? args.regularExpressionList : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["regularExpressionList"] = undefined /*out*/;
            resourceInputs["scope"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name", "scope"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(RegexPatternSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegexPatternSet resource.
 */
export interface RegexPatternSetArgs {
    /**
     * Description of the entity.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the RegexPatternSet.
     */
    name?: pulumi.Input<string>;
    /**
     * The regular expression patterns in the set.
     */
    regularExpressionList: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Use CLOUDFRONT for CloudFront RegexPatternSet, use REGIONAL for Application Load Balancer and API Gateway.
     */
    scope: pulumi.Input<enums.wafv2.RegexPatternSetScope>;
    /**
     * Key:value pairs associated with an AWS resource. The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
     *
     * > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
