// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodePipeline.Outputs
{

    [OutputType]
    public sealed class PipelineActionTypeId
    {
        public readonly string Category;
        public readonly string Owner;
        public readonly string Provider;
        public readonly string Version;

        [OutputConstructor]
        private PipelineActionTypeId(
            string category,

            string owner,

            string provider,

            string version)
        {
            Category = category;
            Owner = owner;
            Provider = provider;
            Version = version;
        }
    }
}