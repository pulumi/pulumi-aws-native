// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Backup.Inputs
{

    public sealed class FrameworkControlInputParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of a parameter, for example, `BackupPlanFrequency` .
        /// </summary>
        [Input("parameterName", required: true)]
        public Input<string> ParameterName { get; set; } = null!;

        /// <summary>
        /// The value of parameter, for example, `hourly` .
        /// </summary>
        [Input("parameterValue", required: true)]
        public Input<string> ParameterValue { get; set; } = null!;

        public FrameworkControlInputParameterArgs()
        {
        }
        public static new FrameworkControlInputParameterArgs Empty => new FrameworkControlInputParameterArgs();
    }
}
