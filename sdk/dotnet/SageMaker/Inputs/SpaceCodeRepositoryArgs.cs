// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class SpaceCodeRepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A CodeRepository (valid URL) to be used within Jupyter's Git extension.
        /// </summary>
        [Input("repositoryUrl", required: true)]
        public Input<string> RepositoryUrl { get; set; } = null!;

        public SpaceCodeRepositoryArgs()
        {
        }
        public static new SpaceCodeRepositoryArgs Empty => new SpaceCodeRepositoryArgs();
    }
}