// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Inputs
{

    public sealed class ResolverPipelineConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("functions")]
        private InputList<string>? _functions;
        public InputList<string> Functions
        {
            get => _functions ?? (_functions = new InputList<string>());
            set => _functions = value;
        }

        public ResolverPipelineConfigArgs()
        {
        }
        public static new ResolverPipelineConfigArgs Empty => new ResolverPipelineConfigArgs();
    }
}