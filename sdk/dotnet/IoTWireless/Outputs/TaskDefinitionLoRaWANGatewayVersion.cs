// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTWireless.Outputs
{

    [OutputType]
    public sealed class TaskDefinitionLoRaWANGatewayVersion
    {
        public readonly string? Model;
        public readonly string? PackageVersion;
        public readonly string? Station;

        [OutputConstructor]
        private TaskDefinitionLoRaWANGatewayVersion(
            string? model,

            string? packageVersion,

            string? station)
        {
            Model = model;
            PackageVersion = packageVersion;
            Station = station;
        }
    }
}