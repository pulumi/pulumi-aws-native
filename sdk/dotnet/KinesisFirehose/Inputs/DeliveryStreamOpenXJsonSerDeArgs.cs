// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Inputs
{

    public sealed class DeliveryStreamOpenXJsonSerDeArgs : global::Pulumi.ResourceArgs
    {
        [Input("caseInsensitive")]
        public Input<bool>? CaseInsensitive { get; set; }

        [Input("columnToJsonKeyMappings")]
        public Input<object>? ColumnToJsonKeyMappings { get; set; }

        [Input("convertDotsInJsonKeysToUnderscores")]
        public Input<bool>? ConvertDotsInJsonKeysToUnderscores { get; set; }

        public DeliveryStreamOpenXJsonSerDeArgs()
        {
        }
        public static new DeliveryStreamOpenXJsonSerDeArgs Empty => new DeliveryStreamOpenXJsonSerDeArgs();
    }
}