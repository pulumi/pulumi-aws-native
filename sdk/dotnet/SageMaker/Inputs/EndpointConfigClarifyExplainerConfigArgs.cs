// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class EndpointConfigClarifyExplainerConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("enableExplanations")]
        public Input<string>? EnableExplanations { get; set; }

        [Input("inferenceConfig")]
        public Input<Inputs.EndpointConfigClarifyInferenceConfigArgs>? InferenceConfig { get; set; }

        [Input("shapConfig", required: true)]
        public Input<Inputs.EndpointConfigClarifyShapConfigArgs> ShapConfig { get; set; } = null!;

        public EndpointConfigClarifyExplainerConfigArgs()
        {
        }
        public static new EndpointConfigClarifyExplainerConfigArgs Empty => new EndpointConfigClarifyExplainerConfigArgs();
    }
}