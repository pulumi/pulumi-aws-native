// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles.Inputs
{

    /// <summary>
    /// The threshold for the calculated attribute.
    /// </summary>
    public sealed class CalculatedAttributeDefinitionThresholdArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The operator of the threshold.
        /// </summary>
        [Input("operator", required: true)]
        public Input<Pulumi.AwsNative.CustomerProfiles.CalculatedAttributeDefinitionThresholdOperator> Operator { get; set; } = null!;

        /// <summary>
        /// The value of the threshold.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public CalculatedAttributeDefinitionThresholdArgs()
        {
        }
        public static new CalculatedAttributeDefinitionThresholdArgs Empty => new CalculatedAttributeDefinitionThresholdArgs();
    }
}
