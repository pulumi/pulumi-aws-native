// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDb.Outputs
{

    [OutputType]
    public sealed class TableImportSourceSpecification
    {
        public readonly string? InputCompressionType;
        public readonly string InputFormat;
        public readonly Outputs.TableInputFormatOptions? InputFormatOptions;
        public readonly Outputs.TableS3BucketSource S3BucketSource;

        [OutputConstructor]
        private TableImportSourceSpecification(
            string? inputCompressionType,

            string inputFormat,

            Outputs.TableInputFormatOptions? inputFormatOptions,

            Outputs.TableS3BucketSource s3BucketSource)
        {
            InputCompressionType = inputCompressionType;
            InputFormat = inputFormat;
            InputFormatOptions = inputFormatOptions;
            S3BucketSource = s3BucketSource;
        }
    }
}