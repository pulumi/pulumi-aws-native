// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QBusiness.Outputs
{

    [OutputType]
    public sealed class DataSourceDocumentAttributeCondition
    {
        public readonly string Key;
        public readonly Pulumi.AwsNative.QBusiness.DataSourceDocumentEnrichmentConditionOperator Operator;
        public readonly object? Value;

        [OutputConstructor]
        private DataSourceDocumentAttributeCondition(
            string key,

            Pulumi.AwsNative.QBusiness.DataSourceDocumentEnrichmentConditionOperator @operator,

            object? value)
        {
            Key = key;
            Operator = @operator;
            Value = value;
        }
    }
}