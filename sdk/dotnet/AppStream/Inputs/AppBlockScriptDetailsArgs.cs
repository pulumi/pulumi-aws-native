// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppStream.Inputs
{

    public sealed class AppBlockScriptDetailsArgs : global::Pulumi.ResourceArgs
    {
        [Input("executableParameters")]
        public Input<string>? ExecutableParameters { get; set; }

        [Input("executablePath", required: true)]
        public Input<string> ExecutablePath { get; set; } = null!;

        [Input("scriptS3Location", required: true)]
        public Input<Inputs.AppBlockS3LocationArgs> ScriptS3Location { get; set; } = null!;

        [Input("timeoutInSeconds", required: true)]
        public Input<int> TimeoutInSeconds { get; set; } = null!;

        public AppBlockScriptDetailsArgs()
        {
        }
        public static new AppBlockScriptDetailsArgs Empty => new AppBlockScriptDetailsArgs();
    }
}