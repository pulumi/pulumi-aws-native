// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Outputs
{

    [OutputType]
    public sealed class AnalysisTemplateAnalysisParameter
    {
        /// <summary>
        /// Optional. The default value that is applied in the analysis template. The member who can query can override this value in the query editor.
        /// </summary>
        public readonly string? DefaultValue;
        /// <summary>
        /// The name of the parameter. The name must use only alphanumeric, underscore (_), or hyphen (-) characters but cannot start or end with a hyphen.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The type of parameter.
        /// </summary>
        public readonly Pulumi.AwsNative.CleanRooms.AnalysisTemplateAnalysisParameterType Type;

        [OutputConstructor]
        private AnalysisTemplateAnalysisParameter(
            string? defaultValue,

            string name,

            Pulumi.AwsNative.CleanRooms.AnalysisTemplateAnalysisParameterType type)
        {
            DefaultValue = defaultValue;
            Name = name;
            Type = type;
        }
    }
}
