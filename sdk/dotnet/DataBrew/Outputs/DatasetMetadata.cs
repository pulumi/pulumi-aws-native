// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    [OutputType]
    public sealed class DatasetMetadata
    {
        /// <summary>
        /// Arn of the source of the dataset. For e.g.: AppFlow Flow ARN.
        /// </summary>
        public readonly string? SourceArn;

        [OutputConstructor]
        private DatasetMetadata(string? sourceArn)
        {
            SourceArn = sourceArn;
        }
    }
}
