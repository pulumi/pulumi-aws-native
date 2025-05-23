// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
        /// <summary>
        /// A date expressed as an ISO 8601 string.
        /// 
        /// It is important for the time zone to be included in the ISO 8601 date-time format. For example, 2012-03-25T12:30:10+01:00 is the ISO 8601 date-time format for March 25th 2012 at 12:30PM (plus 10 seconds) in Central European Time.
        /// </summary>
        [Input("dateValue")]
        public Input<string>? DateValue { get; set; }

        /// <summary>
        /// A long integer value.
        /// </summary>
        [Input("longValue")]
        public Input<int>? LongValue { get; set; }

        [Input("stringListValue")]
        private InputList<string>? _stringListValue;

        /// <summary>
        /// A list of strings. The default maximum length or number of strings is 10.
        /// </summary>
        public InputList<string> StringListValue
        {
            get => _stringListValue ?? (_stringListValue = new InputList<string>());
            set => _stringListValue = value;
        }

        /// <summary>
        /// A string, such as "department".
        /// </summary>
        [Input("stringValue")]
        public Input<string>? StringValue { get; set; }

        public DataSourceDocumentAttributeValueArgs()
        {
        }
        public static new DataSourceDocumentAttributeValueArgs Empty => new DataSourceDocumentAttributeValueArgs();
    }
}
