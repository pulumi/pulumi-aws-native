// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTTwinMaker.Outputs
{

    [OutputType]
    public sealed class EntityStatusErrorProperties
    {
        public readonly Pulumi.AwsNative.IoTTwinMaker.EntityStatusErrorPropertiesCode? Code;
        public readonly string? Message;

        [OutputConstructor]
        private EntityStatusErrorProperties(
            Pulumi.AwsNative.IoTTwinMaker.EntityStatusErrorPropertiesCode? code,

            string? message)
        {
            Code = code;
            Message = message;
        }
    }
}