// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Rds
{
    public static class GetOptionGroup
    {
        /// <summary>
        /// The ``AWS::RDS::OptionGroup`` resource creates or updates an option group, to enable and configure features that are specific to a particular DB engine.
        /// </summary>
        public static Task<GetOptionGroupResult> InvokeAsync(GetOptionGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOptionGroupResult>("aws-native:rds:getOptionGroup", args ?? new GetOptionGroupArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::RDS::OptionGroup`` resource creates or updates an option group, to enable and configure features that are specific to a particular DB engine.
        /// </summary>
        public static Output<GetOptionGroupResult> Invoke(GetOptionGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOptionGroupResult>("aws-native:rds:getOptionGroup", args ?? new GetOptionGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::RDS::OptionGroup`` resource creates or updates an option group, to enable and configure features that are specific to a particular DB engine.
        /// </summary>
        public static Output<GetOptionGroupResult> Invoke(GetOptionGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOptionGroupResult>("aws-native:rds:getOptionGroup", args ?? new GetOptionGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOptionGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the option group to be created.
        ///  Constraints:
        ///   +  Must be 1 to 255 letters, numbers, or hyphens
        ///   +  First character must be a letter
        ///   +  Can't end with a hyphen or contain two consecutive hyphens
        ///   
        ///  Example: ``myoptiongroup``
        ///  If you don't specify a value for ``OptionGroupName`` property, a name is automatically created for the option group.
        ///   This value is stored as a lowercase string.
        /// </summary>
        [Input("optionGroupName", required: true)]
        public string OptionGroupName { get; set; } = null!;

        public GetOptionGroupArgs()
        {
        }
        public static new GetOptionGroupArgs Empty => new GetOptionGroupArgs();
    }

    public sealed class GetOptionGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the option group to be created.
        ///  Constraints:
        ///   +  Must be 1 to 255 letters, numbers, or hyphens
        ///   +  First character must be a letter
        ///   +  Can't end with a hyphen or contain two consecutive hyphens
        ///   
        ///  Example: ``myoptiongroup``
        ///  If you don't specify a value for ``OptionGroupName`` property, a name is automatically created for the option group.
        ///   This value is stored as a lowercase string.
        /// </summary>
        [Input("optionGroupName", required: true)]
        public Input<string> OptionGroupName { get; set; } = null!;

        public GetOptionGroupInvokeArgs()
        {
        }
        public static new GetOptionGroupInvokeArgs Empty => new GetOptionGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetOptionGroupResult
    {
        /// <summary>
        /// A list of all available options for an option group.
        /// </summary>
        public readonly ImmutableArray<Outputs.OptionGroupOptionConfiguration> OptionConfigurations;
        /// <summary>
        /// Tags to assign to the option group.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetOptionGroupResult(
            ImmutableArray<Outputs.OptionGroupOptionConfiguration> optionConfigurations,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            OptionConfigurations = optionConfigurations;
            Tags = tags;
        }
    }
}
