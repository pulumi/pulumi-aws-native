// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom.Inputs
{

    public sealed class KnowledgeBaseSeedUrlArgs : global::Pulumi.ResourceArgs
    {
        [Input("url")]
        public Input<string>? Url { get; set; }

        public KnowledgeBaseSeedUrlArgs()
        {
        }
        public static new KnowledgeBaseSeedUrlArgs Empty => new KnowledgeBaseSeedUrlArgs();
    }
}
