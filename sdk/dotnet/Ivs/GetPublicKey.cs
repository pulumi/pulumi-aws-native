// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ivs
{
    public static class GetPublicKey
    {
        /// <summary>
        /// Resource Type definition for AWS::IVS::PublicKey
        /// </summary>
        public static Task<GetPublicKeyResult> InvokeAsync(GetPublicKeyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPublicKeyResult>("aws-native:ivs:getPublicKey", args ?? new GetPublicKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IVS::PublicKey
        /// </summary>
        public static Output<GetPublicKeyResult> Invoke(GetPublicKeyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPublicKeyResult>("aws-native:ivs:getPublicKey", args ?? new GetPublicKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPublicKeyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Key-pair identifier.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetPublicKeyArgs()
        {
        }
        public static new GetPublicKeyArgs Empty => new GetPublicKeyArgs();
    }

    public sealed class GetPublicKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Key-pair identifier.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetPublicKeyInvokeArgs()
        {
        }
        public static new GetPublicKeyInvokeArgs Empty => new GetPublicKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetPublicKeyResult
    {
        /// <summary>
        /// Key-pair identifier.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Key-pair identifier.
        /// </summary>
        public readonly string? Fingerprint;
        /// <summary>
        /// A list of key-value pairs that contain metadata for the asset model.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetPublicKeyResult(
            string? arn,

            string? fingerprint,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Arn = arn;
            Fingerprint = fingerprint;
            Tags = tags;
        }
    }
}