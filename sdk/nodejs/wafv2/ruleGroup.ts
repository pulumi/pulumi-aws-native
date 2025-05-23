// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Contains the Rules that identify the requests that you want to allow, block, or count. In a RuleGroup, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a RuleGroup, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the RuleGroup with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a RuleGroup, a request needs to match only one of the specifications to be allowed, blocked, or counted.
 */
export class RuleGroup extends pulumi.CustomResource {
    /**
     * Get an existing RuleGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RuleGroup {
        return new RuleGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:wafv2:RuleGroup';

    /**
     * Returns true if the given object is an instance of RuleGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuleGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuleGroup.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the rule group.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Collection of Available Labels.
     */
    public readonly availableLabels!: pulumi.Output<outputs.wafv2.RuleGroupLabelSummary[] | undefined>;
    /**
     * The ID of the rule group.
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * The web ACL capacity units (WCUs) required for this rule group.
     *
     * When you create your own rule group, you define this, and you cannot change it after creation. When you add or modify the rules in a rule group, AWS WAF enforces this limit.
     *
     * AWS WAF uses WCUs to calculate and control the operating resources that are used to run your rules, rule groups, and web ACLs. AWS WAF calculates capacity differently for each rule type, to reflect the relative cost of each rule. Simple rules that cost little to run use fewer WCUs than more complex rules that use more processing power. Rule group capacity is fixed at creation, which helps users plan their web ACL WCU usage when they use a rule group. The WCU limit for web ACLs is 1,500.
     */
    public readonly capacity!: pulumi.Output<number>;
    /**
     * Collection of Consumed Labels.
     */
    public readonly consumedLabels!: pulumi.Output<outputs.wafv2.RuleGroupLabelSummary[] | undefined>;
    /**
     * A map of custom response keys and content bodies. When you create a rule with a block action, you can send a custom response to the web request. You define these for the rule group, and then use them in the rules that you define in the rule group.
     *
     * For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the *AWS WAF Developer Guide* .
     *
     * For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the *AWS WAF Developer Guide* .
     */
    public readonly customResponseBodies!: pulumi.Output<{[key: string]: outputs.wafv2.RuleGroupCustomResponseBody} | undefined>;
    /**
     * A description of the rule group that helps with identification.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The label namespace prefix for this rule group. All labels added by rules in this rule group have this prefix.
     *
     * The syntax for the label namespace prefix for a rule group is the following: `awswaf:<account ID>:rule group:<rule group name>:`
     *
     * When a rule with a label matches a web request, AWS WAF adds the fully qualified label to the request. A fully qualified label is made up of the label namespace from the rule group or web ACL where the rule is defined and the label from the rule, separated by a colon.
     */
    public /*out*/ readonly labelNamespace!: pulumi.Output<string>;
    /**
     * The name of the rule group. You cannot change the name of a rule group after you create it.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Collection of Rules.
     */
    public readonly rules!: pulumi.Output<outputs.wafv2.RuleGroupRule[] | undefined>;
    /**
     * Specifies whether this is for an Amazon CloudFront distribution or for a regional application. For an AWS Amplify application, use `CLOUDFRONT` . A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool, an AWS App Runner service, or an AWS Verified Access instance. Valid Values are `CLOUDFRONT` and `REGIONAL` .
     *
     * > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
     */
    public readonly scope!: pulumi.Output<enums.wafv2.RuleGroupScope>;
    /**
     * Key:value pairs associated with an AWS resource. The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
     *
     * > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * Defines and enables Amazon CloudWatch metrics and web request sample collection.
     */
    public readonly visibilityConfig!: pulumi.Output<outputs.wafv2.RuleGroupVisibilityConfig>;

    /**
     * Create a RuleGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.capacity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'capacity'");
            }
            if ((!args || args.scope === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scope'");
            }
            if ((!args || args.visibilityConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'visibilityConfig'");
            }
            resourceInputs["availableLabels"] = args ? args.availableLabels : undefined;
            resourceInputs["capacity"] = args ? args.capacity : undefined;
            resourceInputs["consumedLabels"] = args ? args.consumedLabels : undefined;
            resourceInputs["customResponseBodies"] = args ? args.customResponseBodies : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["visibilityConfig"] = args ? args.visibilityConfig : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["labelNamespace"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["availableLabels"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["capacity"] = undefined /*out*/;
            resourceInputs["consumedLabels"] = undefined /*out*/;
            resourceInputs["customResponseBodies"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["labelNamespace"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["rules"] = undefined /*out*/;
            resourceInputs["scope"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["visibilityConfig"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name", "scope"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(RuleGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RuleGroup resource.
 */
export interface RuleGroupArgs {
    /**
     * Collection of Available Labels.
     */
    availableLabels?: pulumi.Input<pulumi.Input<inputs.wafv2.RuleGroupLabelSummaryArgs>[]>;
    /**
     * The web ACL capacity units (WCUs) required for this rule group.
     *
     * When you create your own rule group, you define this, and you cannot change it after creation. When you add or modify the rules in a rule group, AWS WAF enforces this limit.
     *
     * AWS WAF uses WCUs to calculate and control the operating resources that are used to run your rules, rule groups, and web ACLs. AWS WAF calculates capacity differently for each rule type, to reflect the relative cost of each rule. Simple rules that cost little to run use fewer WCUs than more complex rules that use more processing power. Rule group capacity is fixed at creation, which helps users plan their web ACL WCU usage when they use a rule group. The WCU limit for web ACLs is 1,500.
     */
    capacity: pulumi.Input<number>;
    /**
     * Collection of Consumed Labels.
     */
    consumedLabels?: pulumi.Input<pulumi.Input<inputs.wafv2.RuleGroupLabelSummaryArgs>[]>;
    /**
     * A map of custom response keys and content bodies. When you create a rule with a block action, you can send a custom response to the web request. You define these for the rule group, and then use them in the rules that you define in the rule group.
     *
     * For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the *AWS WAF Developer Guide* .
     *
     * For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the *AWS WAF Developer Guide* .
     */
    customResponseBodies?: pulumi.Input<{[key: string]: pulumi.Input<inputs.wafv2.RuleGroupCustomResponseBodyArgs>}>;
    /**
     * A description of the rule group that helps with identification.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the rule group. You cannot change the name of a rule group after you create it.
     */
    name?: pulumi.Input<string>;
    /**
     * Collection of Rules.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.wafv2.RuleGroupRuleArgs>[]>;
    /**
     * Specifies whether this is for an Amazon CloudFront distribution or for a regional application. For an AWS Amplify application, use `CLOUDFRONT` . A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool, an AWS App Runner service, or an AWS Verified Access instance. Valid Values are `CLOUDFRONT` and `REGIONAL` .
     *
     * > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
     */
    scope: pulumi.Input<enums.wafv2.RuleGroupScope>;
    /**
     * Key:value pairs associated with an AWS resource. The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
     *
     * > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * Defines and enables Amazon CloudWatch metrics and web request sample collection.
     */
    visibilityConfig: pulumi.Input<inputs.wafv2.RuleGroupVisibilityConfigArgs>;
}
