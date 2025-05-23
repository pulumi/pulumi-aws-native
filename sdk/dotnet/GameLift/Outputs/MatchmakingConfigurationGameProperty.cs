// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Outputs
{

    /// <summary>
    /// A key-value pair that contains information about a game session.
    /// </summary>
    [OutputType]
    public sealed class MatchmakingConfigurationGameProperty
    {
        /// <summary>
        /// The game property identifier.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The game property value.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private MatchmakingConfigurationGameProperty(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
