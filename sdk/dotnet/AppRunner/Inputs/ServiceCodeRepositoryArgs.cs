// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppRunner.Inputs
{

    /// <summary>
    /// Source Code Repository
    /// </summary>
    public sealed class ServiceCodeRepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration for building and running the service from a source code repository.
        /// 
        /// &gt; `CodeConfiguration` is required only for `CreateService` request.
        /// </summary>
        [Input("codeConfiguration")]
        public Input<Inputs.ServiceCodeConfigurationArgs>? CodeConfiguration { get; set; }

        /// <summary>
        /// Repository Url
        /// </summary>
        [Input("repositoryUrl", required: true)]
        public Input<string> RepositoryUrl { get; set; } = null!;

        /// <summary>
        /// The version that should be used within the source code repository.
        /// </summary>
        [Input("sourceCodeVersion", required: true)]
        public Input<Inputs.ServiceSourceCodeVersionArgs> SourceCodeVersion { get; set; } = null!;

        /// <summary>
        /// Source Directory
        /// </summary>
        [Input("sourceDirectory")]
        public Input<string>? SourceDirectory { get; set; }

        public ServiceCodeRepositoryArgs()
        {
        }
        public static new ServiceCodeRepositoryArgs Empty => new ServiceCodeRepositoryArgs();
    }
}
