// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowCustomConnectorSourcePropertiesDataTransferApiProperties
    {
        public readonly string Name;
        public readonly Pulumi.AwsNative.AppFlow.FlowCustomConnectorSourcePropertiesDataTransferApiPropertiesType Type;

        [OutputConstructor]
        private FlowCustomConnectorSourcePropertiesDataTransferApiProperties(
            string name,

            Pulumi.AwsNative.AppFlow.FlowCustomConnectorSourcePropertiesDataTransferApiPropertiesType type)
        {
            Name = name;
            Type = type;
        }
    }
}