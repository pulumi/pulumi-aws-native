// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    public sealed class IndexValueImportanceItemArgs : global::Pulumi.ResourceArgs
    {
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("value")]
        public Input<int>? Value { get; set; }

        public IndexValueImportanceItemArgs()
        {
        }
        public static new IndexValueImportanceItemArgs Empty => new IndexValueImportanceItemArgs();
    }
}
