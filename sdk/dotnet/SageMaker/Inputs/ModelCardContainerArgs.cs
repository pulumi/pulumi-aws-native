// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class ModelCardContainerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Inference environment path. The Amazon EC2 Container Registry (Amazon ECR) path where inference code is stored.
        /// </summary>
        [Input("image", required: true)]
        public Input<string> Image { get; set; } = null!;

        /// <summary>
        /// The Amazon S3 path where the model artifacts, which result from model training, are stored.
        /// </summary>
        [Input("modelDataUrl")]
        public Input<string>? ModelDataUrl { get; set; }

        /// <summary>
        /// The name of a pre-trained machine learning benchmarked by Amazon SageMaker Inference Recommender model that matches your model.
        /// </summary>
        [Input("nearestModelName")]
        public Input<string>? NearestModelName { get; set; }

        public ModelCardContainerArgs()
        {
        }
        public static new ModelCardContainerArgs Empty => new ModelCardContainerArgs();
    }
}