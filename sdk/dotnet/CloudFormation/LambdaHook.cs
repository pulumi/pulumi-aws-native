// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFormation
{
    /// <summary>
    /// This is a CloudFormation resource for the first-party AWS::Hooks::LambdaHook.
    /// </summary>
    [AwsNativeResourceType("aws-native:cloudformation:LambdaHook")]
    public partial class LambdaHook : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The typename alias for the hook.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// The execution role ARN assumed by Hooks to invoke Lambda.
        /// </summary>
        [Output("executionRole")]
        public Output<string> ExecutionRole { get; private set; } = null!;

        /// <summary>
        /// Attribute to specify CloudFormation behavior on hook failure.
        /// </summary>
        [Output("failureMode")]
        public Output<Pulumi.AwsNative.CloudFormation.LambdaHookFailureMode> FailureMode { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the activated hook
        /// </summary>
        [Output("hookArn")]
        public Output<string> HookArn { get; private set; } = null!;

        /// <summary>
        /// Attribute to specify which stacks this hook applies to or should get invoked for
        /// </summary>
        [Output("hookStatus")]
        public Output<Pulumi.AwsNative.CloudFormation.LambdaHookHookStatus> HookStatus { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN), Partial ARN, name, version, or alias of the Lambda function to invoke with this hook.
        /// </summary>
        [Output("lambdaFunction")]
        public Output<string> LambdaFunction { get; private set; } = null!;

        /// <summary>
        /// Filters to allow hooks to target specific stack attributes
        /// </summary>
        [Output("stackFilters")]
        public Output<Outputs.StackFiltersProperties?> StackFilters { get; private set; } = null!;

        /// <summary>
        /// Attribute to specify which targets should invoke the hook
        /// </summary>
        [Output("targetFilters")]
        public Output<Union<Outputs.TargetFilters0Properties, Outputs.TargetFilters1Properties>?> TargetFilters { get; private set; } = null!;

        /// <summary>
        /// Which operations should this Hook run against? Resource changes, stacks or change sets.
        /// </summary>
        [Output("targetOperations")]
        public Output<ImmutableArray<Pulumi.AwsNative.CloudFormation.LambdaHookTargetOperation>> TargetOperations { get; private set; } = null!;


        /// <summary>
        /// Create a LambdaHook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LambdaHook(string name, LambdaHookArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cloudformation:LambdaHook", name, args ?? new LambdaHookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LambdaHook(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cloudformation:LambdaHook", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "alias",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LambdaHook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LambdaHook Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LambdaHook(name, id, options);
        }
    }

    public sealed class LambdaHookArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The typename alias for the hook.
        /// </summary>
        [Input("alias", required: true)]
        public Input<string> Alias { get; set; } = null!;

        /// <summary>
        /// The execution role ARN assumed by Hooks to invoke Lambda.
        /// </summary>
        [Input("executionRole", required: true)]
        public Input<string> ExecutionRole { get; set; } = null!;

        /// <summary>
        /// Attribute to specify CloudFormation behavior on hook failure.
        /// </summary>
        [Input("failureMode", required: true)]
        public Input<Pulumi.AwsNative.CloudFormation.LambdaHookFailureMode> FailureMode { get; set; } = null!;

        /// <summary>
        /// Attribute to specify which stacks this hook applies to or should get invoked for
        /// </summary>
        [Input("hookStatus", required: true)]
        public Input<Pulumi.AwsNative.CloudFormation.LambdaHookHookStatus> HookStatus { get; set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN), Partial ARN, name, version, or alias of the Lambda function to invoke with this hook.
        /// </summary>
        [Input("lambdaFunction", required: true)]
        public Input<string> LambdaFunction { get; set; } = null!;

        /// <summary>
        /// Filters to allow hooks to target specific stack attributes
        /// </summary>
        [Input("stackFilters")]
        public Input<Inputs.StackFiltersPropertiesArgs>? StackFilters { get; set; }

        /// <summary>
        /// Attribute to specify which targets should invoke the hook
        /// </summary>
        [Input("targetFilters")]
        public InputUnion<Inputs.TargetFilters0PropertiesArgs, Inputs.TargetFilters1PropertiesArgs>? TargetFilters { get; set; }

        [Input("targetOperations", required: true)]
        private InputList<Pulumi.AwsNative.CloudFormation.LambdaHookTargetOperation>? _targetOperations;

        /// <summary>
        /// Which operations should this Hook run against? Resource changes, stacks or change sets.
        /// </summary>
        public InputList<Pulumi.AwsNative.CloudFormation.LambdaHookTargetOperation> TargetOperations
        {
            get => _targetOperations ?? (_targetOperations = new InputList<Pulumi.AwsNative.CloudFormation.LambdaHookTargetOperation>());
            set => _targetOperations = value;
        }

        public LambdaHookArgs()
        {
        }
        public static new LambdaHookArgs Empty => new LambdaHookArgs();
    }
}
