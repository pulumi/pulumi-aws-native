// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Contains the Rules that identify the requests that you want to allow, block, or count. In a WebACL, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a WebACL, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the WebACL with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a WebACL, a request needs to match only one of the specifications to be allowed, blocked, or counted.
 */
export function getWebAcl(args: GetWebAclArgs, opts?: pulumi.InvokeOptions): Promise<GetWebAclResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:wafv2:getWebAcl", {
        "id": args.id,
        "name": args.name,
        "scope": args.scope,
    }, opts);
}

export interface GetWebAclArgs {
    /**
     * The ID of the web ACL.
     */
    id: string;
    /**
     * The name of the web ACL. You cannot change the name of a web ACL after you create it.
     */
    name: string;
    /**
     * Specifies whether this is for an Amazon CloudFront distribution or for a regional application. For an AWS Amplify application, use `CLOUDFRONT` . A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool, an AWS App Runner service, or an AWS Verified Access instance. Valid Values are `CLOUDFRONT` and `REGIONAL` .
     *
     * > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` . 
     *
     * For information about how to define the association of the web ACL with your resource, see `WebACLAssociation` .
     */
    scope: enums.wafv2.WebAclScope;
}

export interface GetWebAclResult {
    /**
     * The Amazon Resource Name (ARN) of the web ACL.
     */
    readonly arn?: string;
    /**
     * Specifies custom configurations for the associations between the web ACL and protected resources.
     *
     * Use this to customize the maximum size of the request body that your protected resources forward to AWS WAF for inspection. You can customize this setting for CloudFront, API Gateway, Amazon Cognito, App Runner, or Verified Access resources. The default setting is 16 KB (16,384 bytes).
     *
     * > You are charged additional fees when your protected resources forward body sizes that are larger than the default. For more information, see [AWS WAF Pricing](https://docs.aws.amazon.com/waf/pricing/) . 
     *
     * For Application Load Balancer and AWS AppSync , the limit is fixed at 8 KB (8,192 bytes).
     */
    readonly associationConfig?: outputs.wafv2.WebAclAssociationConfig;
    /**
     * The web ACL capacity units (WCUs) currently being used by this web ACL.
     *
     * AWS WAF uses WCUs to calculate and control the operating resources that are used to run your rules, rule groups, and web ACLs. AWS WAF calculates capacity differently for each rule type, to reflect the relative cost of each rule. Simple rules that cost little to run use fewer WCUs than more complex rules that use more processing power. Rule group capacity is fixed at creation, which helps users plan their web ACL WCU usage when they use a rule group. The WCU limit for web ACLs is 1,500.
     */
    readonly capacity?: number;
    /**
     * Specifies how AWS WAF should handle `CAPTCHA` evaluations for rules that don't have their own `CaptchaConfig` settings. If you don't specify this, AWS WAF uses its default settings for `CaptchaConfig` .
     */
    readonly captchaConfig?: outputs.wafv2.WebAclCaptchaConfig;
    /**
     * Specifies how AWS WAF should handle challenge evaluations for rules that don't have their own `ChallengeConfig` settings. If you don't specify this, AWS WAF uses its default settings for `ChallengeConfig` .
     */
    readonly challengeConfig?: outputs.wafv2.WebAclChallengeConfig;
    /**
     * A map of custom response keys and content bodies. When you create a rule with a block action, you can send a custom response to the web request. You define these for the web ACL, and then use them in the rules and default actions that you define in the web ACL.
     *
     * For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the *AWS WAF Developer Guide* .
     *
     * For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the *AWS WAF Developer Guide* .
     */
    readonly customResponseBodies?: {[key: string]: outputs.wafv2.WebAclCustomResponseBody};
    /**
     * Collection of dataProtects.
     */
    readonly dataProtectionConfig?: outputs.wafv2.WebAclDataProtectionConfig;
    /**
     * The action to perform if none of the `Rules` contained in the `WebACL` match.
     */
    readonly defaultAction?: outputs.wafv2.WebAclDefaultAction;
    /**
     * A description of the web ACL that helps with identification.
     */
    readonly description?: string;
    /**
     * The ID of the web ACL.
     */
    readonly id?: string;
    /**
     * The label namespace prefix for this web ACL. All labels added by rules in this web ACL have this prefix.
     *
     * The syntax for the label namespace prefix for a web ACL is the following: `awswaf:<account ID>:webacl:<web ACL name>:`
     *
     * When a rule with a label matches a web request, AWS WAF adds the fully qualified label to the request. A fully qualified label is made up of the label namespace from the rule group or web ACL where the rule is defined and the label from the rule, separated by a colon.
     */
    readonly labelNamespace?: string;
    /**
     * Configures the level of DDoS protection that applies to web ACLs associated with Application Load Balancers.
     */
    readonly onSourceDDoSProtectionConfig?: outputs.wafv2.WebAclOnSourceDDoSProtectionConfig;
    /**
     * Collection of Rules.
     */
    readonly rules?: outputs.wafv2.WebAclRule[];
    /**
     * Key:value pairs associated with an AWS resource. The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
     *
     * > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
     */
    readonly tags?: outputs.Tag[];
    /**
     * Specifies the domains that AWS WAF should accept in a web request token. This enables the use of tokens across multiple protected websites. When AWS WAF provides a token, it uses the domain of the AWS resource that the web ACL is protecting. If you don't specify a list of token domains, AWS WAF accepts tokens only for the domain of the protected resource. With a token domain list, AWS WAF accepts the resource's host domain plus all domains in the token domain list, including their prefixed subdomains.
     */
    readonly tokenDomains?: string[];
    /**
     * Defines and enables Amazon CloudWatch metrics and web request sample collection.
     */
    readonly visibilityConfig?: outputs.wafv2.WebAclVisibilityConfig;
}
/**
 * Contains the Rules that identify the requests that you want to allow, block, or count. In a WebACL, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a WebACL, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the WebACL with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a WebACL, a request needs to match only one of the specifications to be allowed, blocked, or counted.
 */
export function getWebAclOutput(args: GetWebAclOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetWebAclResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:wafv2:getWebAcl", {
        "id": args.id,
        "name": args.name,
        "scope": args.scope,
    }, opts);
}

export interface GetWebAclOutputArgs {
    /**
     * The ID of the web ACL.
     */
    id: pulumi.Input<string>;
    /**
     * The name of the web ACL. You cannot change the name of a web ACL after you create it.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies whether this is for an Amazon CloudFront distribution or for a regional application. For an AWS Amplify application, use `CLOUDFRONT` . A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool, an AWS App Runner service, or an AWS Verified Access instance. Valid Values are `CLOUDFRONT` and `REGIONAL` .
     *
     * > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` . 
     *
     * For information about how to define the association of the web ACL with your resource, see `WebACLAssociation` .
     */
    scope: pulumi.Input<enums.wafv2.WebAclScope>;
}
