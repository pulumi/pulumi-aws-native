// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class ParallelismConfigurationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum parallel execution steps
        /// </summary>
        [Input("maxParallelExecutionSteps", required: true)]
        public Input<int> MaxParallelExecutionSteps { get; set; } = null!;

        public ParallelismConfigurationPropertiesArgs()
        {
        }
        public static new ParallelismConfigurationPropertiesArgs Empty => new ParallelismConfigurationPropertiesArgs();
    }
}