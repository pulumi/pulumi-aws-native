// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityLake.Inputs
{

    public sealed class SubscriberCustomLogSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name for a third-party custom source. This must be a Regionally unique value.
        /// </summary>
        [Input("sourceName")]
        public Input<string>? SourceName { get; set; }

        /// <summary>
        /// The version for a third-party custom source. This must be a Regionally unique value.
        /// </summary>
        [Input("sourceVersion")]
        public Input<string>? SourceVersion { get; set; }

        public SubscriberCustomLogSourceArgs()
        {
        }
        public static new SubscriberCustomLogSourceArgs Empty => new SubscriberCustomLogSourceArgs();
    }
}