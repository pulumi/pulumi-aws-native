// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2
{
    /// <summary>
    /// Contains the Rules that identify the requests that you want to allow, block, or count. In a WebACL, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a WebACL, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the WebACL with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a WebACL, a request needs to match only one of the specifications to be allowed, blocked, or counted.
    /// </summary>
    [AwsNativeResourceType("aws-native:wafv2:WebAcl")]
    public partial class WebAcl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the web ACL.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Specifies custom configurations for the associations between the web ACL and protected resources.
        /// 
        /// Use this to customize the maximum size of the request body that your protected resources forward to AWS WAF for inspection. You can customize this setting for CloudFront, API Gateway, Amazon Cognito, App Runner, or Verified Access resources. The default setting is 16 KB (16,384 bytes).
        /// 
        /// &gt; You are charged additional fees when your protected resources forward body sizes that are larger than the default. For more information, see [AWS WAF Pricing](https://docs.aws.amazon.com/waf/pricing/) . 
        /// 
        /// For Application Load Balancer and AWS AppSync , the limit is fixed at 8 KB (8,192 bytes).
        /// </summary>
        [Output("associationConfig")]
        public Output<Outputs.WebAclAssociationConfig?> AssociationConfig { get; private set; } = null!;

        /// <summary>
        /// The ID of the web ACL.
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// The web ACL capacity units (WCUs) currently being used by this web ACL.
        /// 
        /// AWS WAF uses WCUs to calculate and control the operating resources that are used to run your rules, rule groups, and web ACLs. AWS WAF calculates capacity differently for each rule type, to reflect the relative cost of each rule. Simple rules that cost little to run use fewer WCUs than more complex rules that use more processing power. Rule group capacity is fixed at creation, which helps users plan their web ACL WCU usage when they use a rule group. The WCU limit for web ACLs is 1,500.
        /// </summary>
        [Output("capacity")]
        public Output<int> Capacity { get; private set; } = null!;

        /// <summary>
        /// Specifies how AWS WAF should handle `CAPTCHA` evaluations for rules that don't have their own `CaptchaConfig` settings. If you don't specify this, AWS WAF uses its default settings for `CaptchaConfig` .
        /// </summary>
        [Output("captchaConfig")]
        public Output<Outputs.WebAclCaptchaConfig?> CaptchaConfig { get; private set; } = null!;

        /// <summary>
        /// Specifies how AWS WAF should handle challenge evaluations for rules that don't have their own `ChallengeConfig` settings. If you don't specify this, AWS WAF uses its default settings for `ChallengeConfig` .
        /// </summary>
        [Output("challengeConfig")]
        public Output<Outputs.WebAclChallengeConfig?> ChallengeConfig { get; private set; } = null!;

        /// <summary>
        /// A map of custom response keys and content bodies. When you create a rule with a block action, you can send a custom response to the web request. You define these for the web ACL, and then use them in the rules and default actions that you define in the web ACL.
        /// 
        /// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the *AWS WAF Developer Guide* .
        /// 
        /// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the *AWS WAF Developer Guide* .
        /// </summary>
        [Output("customResponseBodies")]
        public Output<ImmutableDictionary<string, Outputs.WebAclCustomResponseBody>?> CustomResponseBodies { get; private set; } = null!;

        /// <summary>
        /// Collection of dataProtects.
        /// </summary>
        [Output("dataProtectionConfig")]
        public Output<Outputs.WebAclDataProtectionConfig?> DataProtectionConfig { get; private set; } = null!;

        /// <summary>
        /// The action to perform if none of the `Rules` contained in the `WebACL` match.
        /// </summary>
        [Output("defaultAction")]
        public Output<Outputs.WebAclDefaultAction> DefaultAction { get; private set; } = null!;

        /// <summary>
        /// A description of the web ACL that helps with identification.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The label namespace prefix for this web ACL. All labels added by rules in this web ACL have this prefix.
        /// 
        /// The syntax for the label namespace prefix for a web ACL is the following: `awswaf:&lt;account ID&gt;:webacl:&lt;web ACL name&gt;:`
        /// 
        /// When a rule with a label matches a web request, AWS WAF adds the fully qualified label to the request. A fully qualified label is made up of the label namespace from the rule group or web ACL where the rule is defined and the label from the rule, separated by a colon.
        /// </summary>
        [Output("labelNamespace")]
        public Output<string> LabelNamespace { get; private set; } = null!;

        /// <summary>
        /// The name of the web ACL. You cannot change the name of a web ACL after you create it.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Configures the level of DDoS protection that applies to web ACLs associated with Application Load Balancers.
        /// </summary>
        [Output("onSourceDDoSProtectionConfig")]
        public Output<Outputs.WebAclOnSourceDDoSProtectionConfig?> OnSourceDDoSProtectionConfig { get; private set; } = null!;

        /// <summary>
        /// Collection of Rules.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.WebAclRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// Specifies whether this is for an Amazon CloudFront distribution or for a regional application. For an AWS Amplify application, use `CLOUDFRONT` . A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool, an AWS App Runner service, or an AWS Verified Access instance. Valid Values are `CLOUDFRONT` and `REGIONAL` .
        /// 
        /// &gt; For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` . 
        /// 
        /// For information about how to define the association of the web ACL with your resource, see `WebACLAssociation` .
        /// </summary>
        [Output("scope")]
        public Output<Pulumi.AwsNative.WaFv2.WebAclScope> Scope { get; private set; } = null!;

        /// <summary>
        /// Key:value pairs associated with an AWS resource. The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
        /// 
        /// &gt; To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the domains that AWS WAF should accept in a web request token. This enables the use of tokens across multiple protected websites. When AWS WAF provides a token, it uses the domain of the AWS resource that the web ACL is protecting. If you don't specify a list of token domains, AWS WAF accepts tokens only for the domain of the protected resource. With a token domain list, AWS WAF accepts the resource's host domain plus all domains in the token domain list, including their prefixed subdomains.
        /// </summary>
        [Output("tokenDomains")]
        public Output<ImmutableArray<string>> TokenDomains { get; private set; } = null!;

        /// <summary>
        /// Defines and enables Amazon CloudWatch metrics and web request sample collection.
        /// </summary>
        [Output("visibilityConfig")]
        public Output<Outputs.WebAclVisibilityConfig> VisibilityConfig { get; private set; } = null!;


        /// <summary>
        /// Create a WebAcl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebAcl(string name, WebAclArgs args, CustomResourceOptions? options = null)
            : base("aws-native:wafv2:WebAcl", name, args ?? new WebAclArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebAcl(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:wafv2:WebAcl", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "name",
                    "scope",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WebAcl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebAcl Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WebAcl(name, id, options);
        }
    }

    public sealed class WebAclArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies custom configurations for the associations between the web ACL and protected resources.
        /// 
        /// Use this to customize the maximum size of the request body that your protected resources forward to AWS WAF for inspection. You can customize this setting for CloudFront, API Gateway, Amazon Cognito, App Runner, or Verified Access resources. The default setting is 16 KB (16,384 bytes).
        /// 
        /// &gt; You are charged additional fees when your protected resources forward body sizes that are larger than the default. For more information, see [AWS WAF Pricing](https://docs.aws.amazon.com/waf/pricing/) . 
        /// 
        /// For Application Load Balancer and AWS AppSync , the limit is fixed at 8 KB (8,192 bytes).
        /// </summary>
        [Input("associationConfig")]
        public Input<Inputs.WebAclAssociationConfigArgs>? AssociationConfig { get; set; }

        /// <summary>
        /// Specifies how AWS WAF should handle `CAPTCHA` evaluations for rules that don't have their own `CaptchaConfig` settings. If you don't specify this, AWS WAF uses its default settings for `CaptchaConfig` .
        /// </summary>
        [Input("captchaConfig")]
        public Input<Inputs.WebAclCaptchaConfigArgs>? CaptchaConfig { get; set; }

        /// <summary>
        /// Specifies how AWS WAF should handle challenge evaluations for rules that don't have their own `ChallengeConfig` settings. If you don't specify this, AWS WAF uses its default settings for `ChallengeConfig` .
        /// </summary>
        [Input("challengeConfig")]
        public Input<Inputs.WebAclChallengeConfigArgs>? ChallengeConfig { get; set; }

        [Input("customResponseBodies")]
        private InputMap<Inputs.WebAclCustomResponseBodyArgs>? _customResponseBodies;

        /// <summary>
        /// A map of custom response keys and content bodies. When you create a rule with a block action, you can send a custom response to the web request. You define these for the web ACL, and then use them in the rules and default actions that you define in the web ACL.
        /// 
        /// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the *AWS WAF Developer Guide* .
        /// 
        /// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the *AWS WAF Developer Guide* .
        /// </summary>
        public InputMap<Inputs.WebAclCustomResponseBodyArgs> CustomResponseBodies
        {
            get => _customResponseBodies ?? (_customResponseBodies = new InputMap<Inputs.WebAclCustomResponseBodyArgs>());
            set => _customResponseBodies = value;
        }

        /// <summary>
        /// Collection of dataProtects.
        /// </summary>
        [Input("dataProtectionConfig")]
        public Input<Inputs.WebAclDataProtectionConfigArgs>? DataProtectionConfig { get; set; }

        /// <summary>
        /// The action to perform if none of the `Rules` contained in the `WebACL` match.
        /// </summary>
        [Input("defaultAction", required: true)]
        public Input<Inputs.WebAclDefaultActionArgs> DefaultAction { get; set; } = null!;

        /// <summary>
        /// A description of the web ACL that helps with identification.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the web ACL. You cannot change the name of a web ACL after you create it.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configures the level of DDoS protection that applies to web ACLs associated with Application Load Balancers.
        /// </summary>
        [Input("onSourceDDoSProtectionConfig")]
        public Input<Inputs.WebAclOnSourceDDoSProtectionConfigArgs>? OnSourceDDoSProtectionConfig { get; set; }

        [Input("rules")]
        private InputList<Inputs.WebAclRuleArgs>? _rules;

        /// <summary>
        /// Collection of Rules.
        /// </summary>
        public InputList<Inputs.WebAclRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.WebAclRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Specifies whether this is for an Amazon CloudFront distribution or for a regional application. For an AWS Amplify application, use `CLOUDFRONT` . A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool, an AWS App Runner service, or an AWS Verified Access instance. Valid Values are `CLOUDFRONT` and `REGIONAL` .
        /// 
        /// &gt; For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` . 
        /// 
        /// For information about how to define the association of the web ACL with your resource, see `WebACLAssociation` .
        /// </summary>
        [Input("scope", required: true)]
        public Input<Pulumi.AwsNative.WaFv2.WebAclScope> Scope { get; set; } = null!;

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// Key:value pairs associated with an AWS resource. The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
        /// 
        /// &gt; To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        [Input("tokenDomains")]
        private InputList<string>? _tokenDomains;

        /// <summary>
        /// Specifies the domains that AWS WAF should accept in a web request token. This enables the use of tokens across multiple protected websites. When AWS WAF provides a token, it uses the domain of the AWS resource that the web ACL is protecting. If you don't specify a list of token domains, AWS WAF accepts tokens only for the domain of the protected resource. With a token domain list, AWS WAF accepts the resource's host domain plus all domains in the token domain list, including their prefixed subdomains.
        /// </summary>
        public InputList<string> TokenDomains
        {
            get => _tokenDomains ?? (_tokenDomains = new InputList<string>());
            set => _tokenDomains = value;
        }

        /// <summary>
        /// Defines and enables Amazon CloudWatch metrics and web request sample collection.
        /// </summary>
        [Input("visibilityConfig", required: true)]
        public Input<Inputs.WebAclVisibilityConfigArgs> VisibilityConfig { get; set; } = null!;

        public WebAclArgs()
        {
        }
        public static new WebAclArgs Empty => new WebAclArgs();
    }
}
