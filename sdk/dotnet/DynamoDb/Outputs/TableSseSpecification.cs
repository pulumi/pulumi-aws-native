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
    public sealed class TableSseSpecification
    {
        public readonly string? KmsMasterKeyId;
        public readonly bool SseEnabled;
        public readonly string? SseType;

        [OutputConstructor]
        private TableSseSpecification(
            string? kmsMasterKeyId,

            bool sseEnabled,

            string? sseType)
        {
            KmsMasterKeyId = kmsMasterKeyId;
            SseEnabled = sseEnabled;
            SseType = sseType;
        }
    }
}