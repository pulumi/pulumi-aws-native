// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    [OutputType]
    public sealed class DataSourceDocumentsMetadataConfiguration
    {
        public readonly string? S3Prefix;

        [OutputConstructor]
        private DataSourceDocumentsMetadataConfiguration(string? s3Prefix)
        {
            S3Prefix = s3Prefix;
        }
    }
}