// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    /// <summary>
    /// Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.
    /// </summary>
    [OutputType]
    public sealed class SentimentAnalysisSettingsProperties
    {
        /// <summary>
        /// Enable to call Amazon Comprehend for Sentiment natively within Lex
        /// </summary>
        public readonly bool DetectSentiment;

        [OutputConstructor]
        private SentimentAnalysisSettingsProperties(bool detectSentiment)
        {
            DetectSentiment = detectSentiment;
        }
    }
}
