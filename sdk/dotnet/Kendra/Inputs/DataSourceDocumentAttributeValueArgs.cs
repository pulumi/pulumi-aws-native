// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    public sealed class DataSourceDocumentAttributeValueArgs : global::Pulumi.ResourceArgs
    {
        [Input("dateValue")]
        public Input<string>? DateValue { get; set; }

        [Input("longValue")]
        public Input<int>? LongValue { get; set; }

        [Input("stringListValue")]
        private InputList<string>? _stringListValue;
        public InputList<string> StringListValue
        {
            get => _stringListValue ?? (_stringListValue = new InputList<string>());
            set => _stringListValue = value;
        }

        [Input("stringValue")]
        public Input<string>? StringValue { get; set; }

        public DataSourceDocumentAttributeValueArgs()
        {
        }
        public static new DataSourceDocumentAttributeValueArgs Empty => new DataSourceDocumentAttributeValueArgs();
    }
}