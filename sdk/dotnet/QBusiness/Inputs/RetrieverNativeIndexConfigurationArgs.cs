// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QBusiness.Inputs
{

    public sealed class RetrieverNativeIndexConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("indexId", required: true)]
        public Input<string> IndexId { get; set; } = null!;

        public RetrieverNativeIndexConfigurationArgs()
        {
        }
        public static new RetrieverNativeIndexConfigurationArgs Empty => new RetrieverNativeIndexConfigurationArgs();
    }
}
