// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    public sealed class DataAutomationProjectDocumentOverrideConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether document splitter is enabled for a project.
        /// </summary>
        [Input("splitter")]
        public Input<Inputs.DataAutomationProjectSplitterConfigurationArgs>? Splitter { get; set; }

        public DataAutomationProjectDocumentOverrideConfigurationArgs()
        {
        }
        public static new DataAutomationProjectDocumentOverrideConfigurationArgs Empty => new DataAutomationProjectDocumentOverrideConfigurationArgs();
    }
}
