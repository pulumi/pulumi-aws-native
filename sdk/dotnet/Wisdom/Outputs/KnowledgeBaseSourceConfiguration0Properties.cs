// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom.Outputs
{

    [OutputType]
    public sealed class KnowledgeBaseSourceConfiguration0Properties
    {
        public readonly Outputs.KnowledgeBaseAppIntegrationsConfiguration AppIntegrations;

        [OutputConstructor]
        private KnowledgeBaseSourceConfiguration0Properties(Outputs.KnowledgeBaseAppIntegrationsConfiguration appIntegrations)
        {
            AppIntegrations = appIntegrations;
        }
    }
}
