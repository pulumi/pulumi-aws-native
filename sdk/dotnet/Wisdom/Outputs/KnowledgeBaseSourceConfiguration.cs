// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom.Outputs
{

    [OutputType]
    public sealed class KnowledgeBaseSourceConfiguration
    {
        /// <summary>
        /// Configuration information for Amazon AppIntegrations to automatically ingest content.
        /// </summary>
        public readonly Outputs.KnowledgeBaseAppIntegrationsConfiguration? AppIntegrations;

        [OutputConstructor]
        private KnowledgeBaseSourceConfiguration(Outputs.KnowledgeBaseAppIntegrationsConfiguration? appIntegrations)
        {
            AppIntegrations = appIntegrations;
        }
    }
}
