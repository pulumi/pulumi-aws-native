// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisCustomActionSetParametersOperation
    {
        /// <summary>
        /// The parameter that determines the value configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisSetParameterValueConfiguration> ParameterValueConfigurations;

        [OutputConstructor]
        private AnalysisCustomActionSetParametersOperation(ImmutableArray<Outputs.AnalysisSetParameterValueConfiguration> parameterValueConfigurations)
        {
            ParameterValueConfigurations = parameterValueConfigurations;
        }
    }
}
