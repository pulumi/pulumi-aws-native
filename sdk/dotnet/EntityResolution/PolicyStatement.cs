// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EntityResolution
{
    /// <summary>
    /// Policy Statement defined in AWS Entity Resolution Service
    /// </summary>
    [AwsNativeResourceType("aws-native:entityresolution:PolicyStatement")]
    public partial class PolicyStatement : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The action that the principal can use on the resource.
        /// 
        /// For example, `entityresolution:GetIdMappingJob` , `entityresolution:GetMatchingJob` .
        /// </summary>
        [Output("action")]
        public Output<ImmutableArray<string>> Action { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the resource that will be accessed by the principal.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A set of condition keys that you can use in key policies.
        /// </summary>
        [Output("condition")]
        public Output<string?> Condition { get; private set; } = null!;

        /// <summary>
        /// Determines whether the permissions specified in the policy are to be allowed ( `Allow` ) or denied ( `Deny` ).
        /// 
        /// &gt; If you set the value of the `effect` parameter to `Deny` for the `AddPolicyStatement` operation, you must also set the value of the `effect` parameter in the `policy` to `Deny` for the `PutPolicy` operation.
        /// </summary>
        [Output("effect")]
        public Output<Pulumi.AwsNative.EntityResolution.PolicyStatementStatementEffect?> Effect { get; private set; } = null!;

        /// <summary>
        /// The AWS service or AWS account that can access the resource defined as ARN.
        /// </summary>
        [Output("principal")]
        public Output<ImmutableArray<string>> Principal { get; private set; } = null!;

        /// <summary>
        /// A statement identifier that differentiates the statement from others in the same policy.
        /// </summary>
        [Output("statementId")]
        public Output<string> StatementId { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyStatement resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyStatement(string name, PolicyStatementArgs args, CustomResourceOptions? options = null)
            : base("aws-native:entityresolution:PolicyStatement", name, args ?? new PolicyStatementArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyStatement(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:entityresolution:PolicyStatement", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "arn",
                    "statementId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PolicyStatement resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyStatement Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PolicyStatement(name, id, options);
        }
    }

    public sealed class PolicyStatementArgs : global::Pulumi.ResourceArgs
    {
        [Input("action")]
        private InputList<string>? _action;

        /// <summary>
        /// The action that the principal can use on the resource.
        /// 
        /// For example, `entityresolution:GetIdMappingJob` , `entityresolution:GetMatchingJob` .
        /// </summary>
        public InputList<string> Action
        {
            get => _action ?? (_action = new InputList<string>());
            set => _action = value;
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the resource that will be accessed by the principal.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        /// <summary>
        /// A set of condition keys that you can use in key policies.
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        /// <summary>
        /// Determines whether the permissions specified in the policy are to be allowed ( `Allow` ) or denied ( `Deny` ).
        /// 
        /// &gt; If you set the value of the `effect` parameter to `Deny` for the `AddPolicyStatement` operation, you must also set the value of the `effect` parameter in the `policy` to `Deny` for the `PutPolicy` operation.
        /// </summary>
        [Input("effect")]
        public Input<Pulumi.AwsNative.EntityResolution.PolicyStatementStatementEffect>? Effect { get; set; }

        [Input("principal")]
        private InputList<string>? _principal;

        /// <summary>
        /// The AWS service or AWS account that can access the resource defined as ARN.
        /// </summary>
        public InputList<string> Principal
        {
            get => _principal ?? (_principal = new InputList<string>());
            set => _principal = value;
        }

        /// <summary>
        /// A statement identifier that differentiates the statement from others in the same policy.
        /// </summary>
        [Input("statementId", required: true)]
        public Input<string> StatementId { get; set; } = null!;

        public PolicyStatementArgs()
        {
        }
        public static new PolicyStatementArgs Empty => new PolicyStatementArgs();
    }
}
