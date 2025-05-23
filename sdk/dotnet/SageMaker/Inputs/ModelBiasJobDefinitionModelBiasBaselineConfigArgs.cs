// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// Baseline configuration used to validate that the data conforms to the specified constraints and statistics.
    /// </summary>
    public sealed class ModelBiasJobDefinitionModelBiasBaselineConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the baseline model bias job.
        /// </summary>
        [Input("baseliningJobName")]
        public Input<string>? BaseliningJobName { get; set; }

        /// <summary>
        /// The constraints resource for a monitoring job.
        /// </summary>
        [Input("constraintsResource")]
        public Input<Inputs.ModelBiasJobDefinitionConstraintsResourceArgs>? ConstraintsResource { get; set; }

        public ModelBiasJobDefinitionModelBiasBaselineConfigArgs()
        {
        }
        public static new ModelBiasJobDefinitionModelBiasBaselineConfigArgs Empty => new ModelBiasJobDefinitionModelBiasBaselineConfigArgs();
    }
}
