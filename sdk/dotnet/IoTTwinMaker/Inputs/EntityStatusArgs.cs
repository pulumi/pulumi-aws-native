// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTTwinMaker.Inputs
{

    public sealed class EntityStatusArgs : global::Pulumi.ResourceArgs
    {
        [Input("error")]
        public InputUnion<object, Inputs.EntityStatusErrorPropertiesArgs>? Error { get; set; }

        [Input("state")]
        public Input<Pulumi.AwsNative.IoTTwinMaker.EntityStatusState>? State { get; set; }

        public EntityStatusArgs()
        {
        }
        public static new EntityStatusArgs Empty => new EntityStatusArgs();
    }
}