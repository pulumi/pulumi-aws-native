// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppStream.Outputs
{

    [OutputType]
    public sealed class AppBlockScriptDetails
    {
        public readonly string? ExecutableParameters;
        public readonly string ExecutablePath;
        public readonly Outputs.AppBlockS3Location ScriptS3Location;
        public readonly int TimeoutInSeconds;

        [OutputConstructor]
        private AppBlockScriptDetails(
            string? executableParameters,

            string executablePath,

            Outputs.AppBlockS3Location scriptS3Location,

            int timeoutInSeconds)
        {
            ExecutableParameters = executableParameters;
            ExecutablePath = executablePath;
            ScriptS3Location = scriptS3Location;
            TimeoutInSeconds = timeoutInSeconds;
        }
    }
}