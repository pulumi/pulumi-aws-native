// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    [OutputType]
    public sealed class ModelCardContainer
    {
        /// <summary>
        /// Inference environment path. The Amazon EC2 Container Registry (Amazon ECR) path where inference code is stored.
        /// </summary>
        public readonly string Image;
        /// <summary>
        /// The Amazon S3 path where the model artifacts, which result from model training, are stored.
        /// </summary>
        public readonly string? ModelDataUrl;
        /// <summary>
        /// The name of a pre-trained machine learning benchmarked by Amazon SageMaker Inference Recommender model that matches your model.
        /// </summary>
        public readonly string? NearestModelName;

        [OutputConstructor]
        private ModelCardContainer(
            string image,

            string? modelDataUrl,

            string? nearestModelName)
        {
            Image = image;
            ModelDataUrl = modelDataUrl;
            NearestModelName = nearestModelName;
        }
    }
}