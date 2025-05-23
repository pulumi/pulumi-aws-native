// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    /// <summary>
    /// Various ways to encode a list of models in a CreateInferenceProfile request
    /// </summary>
    [OutputType]
    public sealed class ApplicationInferenceProfileInferenceProfileModelSourceProperties
    {
        /// <summary>
        /// Source arns for a custom inference profile to copy its regional load balancing config from. This
        /// can either be a foundation model or predefined inference profile ARN.
        /// </summary>
        public readonly string CopyFrom;

        [OutputConstructor]
        private ApplicationInferenceProfileInferenceProfileModelSourceProperties(string copyFrom)
        {
            CopyFrom = copyFrom;
        }
    }
}
