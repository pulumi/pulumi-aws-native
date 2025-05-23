// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Redshift.Outputs
{

    [OutputType]
    public sealed class ClusterParameterGroupParameter
    {
        /// <summary>
        /// The name of the parameter.
        /// </summary>
        public readonly string ParameterName;
        /// <summary>
        /// The value of the parameter. If `ParameterName` is `wlm_json_configuration`, then the maximum size of `ParameterValue` is 8000 characters.
        /// </summary>
        public readonly string ParameterValue;

        [OutputConstructor]
        private ClusterParameterGroupParameter(
            string parameterName,

            string parameterValue)
        {
            ParameterName = parameterName;
            ParameterValue = parameterValue;
        }
    }
}
