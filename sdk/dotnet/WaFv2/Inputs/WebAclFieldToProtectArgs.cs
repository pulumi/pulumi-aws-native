// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Inputs
{

    /// <summary>
    /// Field in log to protect.
    /// </summary>
    public sealed class WebAclFieldToProtectArgs : global::Pulumi.ResourceArgs
    {
        [Input("fieldKeys")]
        private InputList<string>? _fieldKeys;

        /// <summary>
        /// List of field keys to protect
        /// </summary>
        public InputList<string> FieldKeys
        {
            get => _fieldKeys ?? (_fieldKeys = new InputList<string>());
            set => _fieldKeys = value;
        }

        /// <summary>
        /// Field type to protect
        /// </summary>
        [Input("fieldType", required: true)]
        public Input<Pulumi.AwsNative.WaFv2.WebAclFieldToProtectFieldType> FieldType { get; set; } = null!;

        public WebAclFieldToProtectArgs()
        {
        }
        public static new WebAclFieldToProtectArgs Empty => new WebAclFieldToProtectArgs();
    }
}
