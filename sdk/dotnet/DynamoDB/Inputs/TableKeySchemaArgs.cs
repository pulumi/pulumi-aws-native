// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDB.Inputs
{

    public sealed class TableKeySchemaArgs : global::Pulumi.ResourceArgs
    {
        [Input("attributeName", required: true)]
        public Input<string> AttributeName { get; set; } = null!;

        [Input("keyType", required: true)]
        public Input<string> KeyType { get; set; } = null!;

        public TableKeySchemaArgs()
        {
        }
        public static new TableKeySchemaArgs Empty => new TableKeySchemaArgs();
    }
}