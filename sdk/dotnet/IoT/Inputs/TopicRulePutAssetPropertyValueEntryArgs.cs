// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    public sealed class TopicRulePutAssetPropertyValueEntryArgs : global::Pulumi.ResourceArgs
    {
        [Input("assetId")]
        public Input<string>? AssetId { get; set; }

        [Input("entryId")]
        public Input<string>? EntryId { get; set; }

        [Input("propertyAlias")]
        public Input<string>? PropertyAlias { get; set; }

        [Input("propertyId")]
        public Input<string>? PropertyId { get; set; }

        [Input("propertyValues", required: true)]
        private InputList<Inputs.TopicRuleAssetPropertyValueArgs>? _propertyValues;
        public InputList<Inputs.TopicRuleAssetPropertyValueArgs> PropertyValues
        {
            get => _propertyValues ?? (_propertyValues = new InputList<Inputs.TopicRuleAssetPropertyValueArgs>());
            set => _propertyValues = value;
        }

        public TopicRulePutAssetPropertyValueEntryArgs()
        {
        }
        public static new TopicRulePutAssetPropertyValueEntryArgs Empty => new TopicRulePutAssetPropertyValueEntryArgs();
    }
}