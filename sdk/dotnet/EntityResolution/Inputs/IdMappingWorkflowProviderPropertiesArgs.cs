// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EntityResolution.Inputs
{

    public sealed class IdMappingWorkflowProviderPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("intermediateSourceConfiguration")]
        public Input<Inputs.IdMappingWorkflowIntermediateSourceConfigurationArgs>? IntermediateSourceConfiguration { get; set; }

        /// <summary>
        /// Additional Provider configuration that would be required for the provider service. The Configuration must be in JSON string format
        /// </summary>
        [Input("providerConfiguration")]
        public Input<object>? ProviderConfiguration { get; set; }

        /// <summary>
        /// Arn of the Provider Service being used.
        /// </summary>
        [Input("providerServiceArn", required: true)]
        public Input<string> ProviderServiceArn { get; set; } = null!;

        public IdMappingWorkflowProviderPropertiesArgs()
        {
        }
        public static new IdMappingWorkflowProviderPropertiesArgs Empty => new IdMappingWorkflowProviderPropertiesArgs();
    }
}