// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio.Inputs
{

    public sealed class LaunchProfileStreamingSessionStorageRootArgs : global::Pulumi.ResourceArgs
    {
        [Input("linux")]
        public Input<string>? Linux { get; set; }

        [Input("windows")]
        public Input<string>? Windows { get; set; }

        public LaunchProfileStreamingSessionStorageRootArgs()
        {
        }
        public static new LaunchProfileStreamingSessionStorageRootArgs Empty => new LaunchProfileStreamingSessionStorageRootArgs();
    }
}
