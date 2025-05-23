// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53Resolver
{
    /// <summary>
    /// Resource Type definition for AWS::Route53Resolver::ResolverRule
    /// </summary>
    [AwsNativeResourceType("aws-native:route53resolver:ResolverRule")]
    public partial class ResolverRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the resolver rule.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps
        /// </summary>
        [Output("domainName")]
        public Output<string?> DomainName { get; private set; } = null!;

        /// <summary>
        /// The name for the Resolver rule
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the endpoint that the rule is associated with.
        /// </summary>
        [Output("resolverEndpointId")]
        public Output<string?> ResolverEndpointId { get; private set; } = null!;

        /// <summary>
        /// The ID of the endpoint that the rule is associated with.
        /// </summary>
        [Output("resolverRuleId")]
        public Output<string> ResolverRuleId { get; private set; } = null!;

        /// <summary>
        /// When you want to forward DNS queries for specified domain name to resolvers on your network, specify FORWARD. When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify SYSTEM.
        /// </summary>
        [Output("ruleType")]
        public Output<Pulumi.AwsNative.Route53Resolver.ResolverRuleRuleType> RuleType { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.
        /// </summary>
        [Output("targetIps")]
        public Output<ImmutableArray<Outputs.ResolverRuleTargetAddress>> TargetIps { get; private set; } = null!;


        /// <summary>
        /// Create a ResolverRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResolverRule(string name, ResolverRuleArgs args, CustomResourceOptions? options = null)
            : base("aws-native:route53resolver:ResolverRule", name, args ?? new ResolverRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResolverRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:route53resolver:ResolverRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "ruleType",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ResolverRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResolverRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ResolverRule(name, id, options);
        }
    }

    public sealed class ResolverRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// The name for the Resolver rule
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the endpoint that the rule is associated with.
        /// </summary>
        [Input("resolverEndpointId")]
        public Input<string>? ResolverEndpointId { get; set; }

        /// <summary>
        /// When you want to forward DNS queries for specified domain name to resolvers on your network, specify FORWARD. When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify SYSTEM.
        /// </summary>
        [Input("ruleType", required: true)]
        public Input<Pulumi.AwsNative.Route53Resolver.ResolverRuleRuleType> RuleType { get; set; } = null!;

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        [Input("targetIps")]
        private InputList<Inputs.ResolverRuleTargetAddressArgs>? _targetIps;

        /// <summary>
        /// An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.
        /// </summary>
        public InputList<Inputs.ResolverRuleTargetAddressArgs> TargetIps
        {
            get => _targetIps ?? (_targetIps = new InputList<Inputs.ResolverRuleTargetAddressArgs>());
            set => _targetIps = value;
        }

        public ResolverRuleArgs()
        {
        }
        public static new ResolverRuleArgs Empty => new ResolverRuleArgs();
    }
}
