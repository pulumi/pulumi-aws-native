// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodePipeline.Inputs
{

    /// <summary>
    /// Reserved for future use.
    /// </summary>
    public sealed class PipelineBlockerDeclarationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Reserved for future use.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Reserved for future use.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.CodePipeline.PipelineBlockerDeclarationType> Type { get; set; } = null!;

        public PipelineBlockerDeclarationArgs()
        {
        }
        public static new PipelineBlockerDeclarationArgs Empty => new PipelineBlockerDeclarationArgs();
    }
}
