// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppStream
{
    public static class GetDirectoryConfig
    {
        /// <summary>
        /// Resource Type definition for AWS::AppStream::DirectoryConfig
        /// </summary>
        public static Task<GetDirectoryConfigResult> InvokeAsync(GetDirectoryConfigArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDirectoryConfigResult>("aws-native:appstream:getDirectoryConfig", args ?? new GetDirectoryConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::AppStream::DirectoryConfig
        /// </summary>
        public static Output<GetDirectoryConfigResult> Invoke(GetDirectoryConfigInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDirectoryConfigResult>("aws-native:appstream:getDirectoryConfig", args ?? new GetDirectoryConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDirectoryConfigArgs : global::Pulumi.InvokeArgs
    {
        [Input("directoryName", required: true)]
        public string DirectoryName { get; set; } = null!;

        public GetDirectoryConfigArgs()
        {
        }
        public static new GetDirectoryConfigArgs Empty => new GetDirectoryConfigArgs();
    }

    public sealed class GetDirectoryConfigInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("directoryName", required: true)]
        public Input<string> DirectoryName { get; set; } = null!;

        public GetDirectoryConfigInvokeArgs()
        {
        }
        public static new GetDirectoryConfigInvokeArgs Empty => new GetDirectoryConfigInvokeArgs();
    }


    [OutputType]
    public sealed class GetDirectoryConfigResult
    {
        public readonly ImmutableArray<string> OrganizationalUnitDistinguishedNames;
        public readonly Outputs.DirectoryConfigServiceAccountCredentials? ServiceAccountCredentials;

        [OutputConstructor]
        private GetDirectoryConfigResult(
            ImmutableArray<string> organizationalUnitDistinguishedNames,

            Outputs.DirectoryConfigServiceAccountCredentials? serviceAccountCredentials)
        {
            OrganizationalUnitDistinguishedNames = organizationalUnitDistinguishedNames;
            ServiceAccountCredentials = serviceAccountCredentials;
        }
    }
}