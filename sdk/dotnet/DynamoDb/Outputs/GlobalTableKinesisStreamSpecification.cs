// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDb.Outputs
{

    [OutputType]
    public sealed class GlobalTableKinesisStreamSpecification
    {
        /// <summary>
        /// The precision for the time and date that the stream was created.
        /// </summary>
        public readonly Pulumi.AwsNative.DynamoDb.GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision? ApproximateCreationDateTimePrecision;
        /// <summary>
        /// The ARN for a specific Kinesis data stream.
        /// </summary>
        public readonly string StreamArn;

        [OutputConstructor]
        private GlobalTableKinesisStreamSpecification(
            Pulumi.AwsNative.DynamoDb.GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision? approximateCreationDateTimePrecision,

            string streamArn)
        {
            ApproximateCreationDateTimePrecision = approximateCreationDateTimePrecision;
            StreamArn = streamArn;
        }
    }
}
