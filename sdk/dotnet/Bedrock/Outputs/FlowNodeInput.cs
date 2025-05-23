// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    /// <summary>
    /// Input to a node in a flow
    /// </summary>
    [OutputType]
    public sealed class FlowNodeInput
    {
        /// <summary>
        /// Expression for a node input in a flow
        /// </summary>
        public readonly string Expression;
        /// <summary>
        /// Name of a node input in a flow
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The data type of the input. If the input doesn't match this type at runtime, a validation error will be thrown.
        /// </summary>
        public readonly Pulumi.AwsNative.Bedrock.FlowNodeIoDataType Type;

        [OutputConstructor]
        private FlowNodeInput(
            string expression,

            string name,

            Pulumi.AwsNative.Bedrock.FlowNodeIoDataType type)
        {
            Expression = expression;
            Name = name;
            Type = type;
        }
    }
}
