// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Amplify.Inputs
{

    public sealed class BranchBackendArgs : global::Pulumi.ResourceArgs
    {
        [Input("stackArn")]
        public Input<string>? StackArn { get; set; }

        public BranchBackendArgs()
        {
        }
        public static new BranchBackendArgs Empty => new BranchBackendArgs();
    }
}