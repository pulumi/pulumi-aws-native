// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio.Inputs
{

    public sealed class LaunchProfileStreamConfigurationSessionStorageArgs : global::Pulumi.ResourceArgs
    {
        [Input("mode", required: true)]
        private InputList<string>? _mode;
        public InputList<string> Mode
        {
            get => _mode ?? (_mode = new InputList<string>());
            set => _mode = value;
        }

        [Input("root")]
        public Input<Inputs.LaunchProfileStreamingSessionStorageRootArgs>? Root { get; set; }

        public LaunchProfileStreamConfigurationSessionStorageArgs()
        {
        }
        public static new LaunchProfileStreamConfigurationSessionStorageArgs Empty => new LaunchProfileStreamConfigurationSessionStorageArgs();
    }
}
