// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QBusiness.Inputs
{

    public sealed class DataSourceDocumentAttributeValue2PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("longValue", required: true)]
        public Input<double> LongValue { get; set; } = null!;

        public DataSourceDocumentAttributeValue2PropertiesArgs()
        {
        }
        public static new DataSourceDocumentAttributeValue2PropertiesArgs Empty => new DataSourceDocumentAttributeValue2PropertiesArgs();
    }
}
