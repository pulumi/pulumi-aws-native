// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGatewayV2.Outputs
{

    [OutputType]
    public sealed class ApiBodyS3Location
    {
        public readonly string? Bucket;
        public readonly string? Etag;
        public readonly string? Key;
        public readonly string? Version;

        [OutputConstructor]
        private ApiBodyS3Location(
            string? bucket,

            string? etag,

            string? key,

            string? version)
        {
            Bucket = bucket;
            Etag = etag;
            Key = key;
            Version = version;
        }
    }
}