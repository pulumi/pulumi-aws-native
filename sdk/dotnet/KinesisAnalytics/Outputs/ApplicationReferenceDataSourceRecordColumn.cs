// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalytics.Outputs
{

    [OutputType]
    public sealed class ApplicationReferenceDataSourceRecordColumn
    {
        public readonly string? Mapping;
        public readonly string Name;
        public readonly string SqlType;

        [OutputConstructor]
        private ApplicationReferenceDataSourceRecordColumn(
            string? mapping,

            string name,

            string sqlType)
        {
            Mapping = mapping;
            Name = name;
            SqlType = sqlType;
        }
    }
}