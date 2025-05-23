// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration.Inputs
{

    /// <summary>
    /// Input parameters in the form of key-value pairs for the conformance pack.
    /// </summary>
    public sealed class ConformancePackInputParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// One part of a key-value pair.
        /// </summary>
        [Input("parameterName", required: true)]
        public Input<string> ParameterName { get; set; } = null!;

        /// <summary>
        /// Another part of the key-value pair.
        /// </summary>
        [Input("parameterValue", required: true)]
        public Input<string> ParameterValue { get; set; } = null!;

        public ConformancePackInputParameterArgs()
        {
        }
        public static new ConformancePackInputParameterArgs Empty => new ConformancePackInputParameterArgs();
    }
}
