// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    [OutputType]
    public sealed class DataAutomationProjectSplitterConfiguration
    {
        /// <summary>
        /// Whether document splitter is enabled for a project.
        /// </summary>
        public readonly Pulumi.AwsNative.Bedrock.DataAutomationProjectState? State;

        [OutputConstructor]
        private DataAutomationProjectSplitterConfiguration(Pulumi.AwsNative.Bedrock.DataAutomationProjectState? state)
        {
            State = state;
        }
    }
}
