// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisDynamicDefaultValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The column that contains the default value of each user or group.
        /// </summary>
        [Input("defaultValueColumn", required: true)]
        public Input<Inputs.AnalysisColumnIdentifierArgs> DefaultValueColumn { get; set; } = null!;

        /// <summary>
        /// The column that contains the group name.
        /// </summary>
        [Input("groupNameColumn")]
        public Input<Inputs.AnalysisColumnIdentifierArgs>? GroupNameColumn { get; set; }

        /// <summary>
        /// The column that contains the username.
        /// </summary>
        [Input("userNameColumn")]
        public Input<Inputs.AnalysisColumnIdentifierArgs>? UserNameColumn { get; set; }

        public AnalysisDynamicDefaultValueArgs()
        {
        }
        public static new AnalysisDynamicDefaultValueArgs Empty => new AnalysisDynamicDefaultValueArgs();
    }
}
