// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Omics.Inputs
{

    public sealed class AnnotationStoreReferenceItemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The reference's ARN.
        /// </summary>
        [Input("referenceArn", required: true)]
        public Input<string> ReferenceArn { get; set; } = null!;

        public AnnotationStoreReferenceItemArgs()
        {
        }
        public static new AnnotationStoreReferenceItemArgs Empty => new AnnotationStoreReferenceItemArgs();
    }
}
