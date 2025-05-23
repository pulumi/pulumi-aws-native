// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Inputs
{

    public sealed class DeliveryStreamDatabaseColumnsArgs : global::Pulumi.ResourceArgs
    {
        [Input("exclude")]
        private InputList<string>? _exclude;
        public InputList<string> Exclude
        {
            get => _exclude ?? (_exclude = new InputList<string>());
            set => _exclude = value;
        }

        [Input("include")]
        private InputList<string>? _include;
        public InputList<string> Include
        {
            get => _include ?? (_include = new InputList<string>());
            set => _include = value;
        }

        public DeliveryStreamDatabaseColumnsArgs()
        {
        }
        public static new DeliveryStreamDatabaseColumnsArgs Empty => new DeliveryStreamDatabaseColumnsArgs();
    }
}
