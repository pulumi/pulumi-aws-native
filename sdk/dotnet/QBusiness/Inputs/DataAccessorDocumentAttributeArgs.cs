// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QBusiness.Inputs
{

    public sealed class DataAccessorDocumentAttributeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identifier for the attribute.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The value of the attribute.
        /// </summary>
        [Input("value", required: true)]
        public object Value { get; set; } = null!;

        public DataAccessorDocumentAttributeArgs()
        {
        }
        public static new DataAccessorDocumentAttributeArgs Empty => new DataAccessorDocumentAttributeArgs();
    }
}
