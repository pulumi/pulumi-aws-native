// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Inputs
{

    public sealed class FunctionRuntimeManagementConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique identifier for a runtime version arn
        /// </summary>
        [Input("runtimeVersionArn")]
        public Input<string>? RuntimeVersionArn { get; set; }

        /// <summary>
        /// Trigger for runtime update
        /// </summary>
        [Input("updateRuntimeOn", required: true)]
        public Input<Pulumi.AwsNative.Lambda.FunctionRuntimeManagementConfigUpdateRuntimeOn> UpdateRuntimeOn { get; set; } = null!;

        public FunctionRuntimeManagementConfigArgs()
        {
        }
        public static new FunctionRuntimeManagementConfigArgs Empty => new FunctionRuntimeManagementConfigArgs();
    }
}