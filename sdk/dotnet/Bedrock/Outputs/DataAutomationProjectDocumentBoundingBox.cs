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
    public sealed class DataAutomationProjectDocumentBoundingBox
    {
        /// <summary>
        /// Whether bounding boxes are enabled for documents.
        /// </summary>
        public readonly Pulumi.AwsNative.Bedrock.DataAutomationProjectState State;

        [OutputConstructor]
        private DataAutomationProjectDocumentBoundingBox(Pulumi.AwsNative.Bedrock.DataAutomationProjectState state)
        {
            State = state;
        }
    }
}
