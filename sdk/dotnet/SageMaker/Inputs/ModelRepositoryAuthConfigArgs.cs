// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class ModelRepositoryAuthConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("repositoryCredentialsProviderArn", required: true)]
        public Input<string> RepositoryCredentialsProviderArn { get; set; } = null!;

        public ModelRepositoryAuthConfigArgs()
        {
        }
        public static new ModelRepositoryAuthConfigArgs Empty => new ModelRepositoryAuthConfigArgs();
    }
}