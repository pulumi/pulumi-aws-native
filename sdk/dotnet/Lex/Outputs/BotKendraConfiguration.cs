// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    [OutputType]
    public sealed class BotKendraConfiguration
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon Kendra index that you want the `AMAZON.KendraSearchIntent` intent to search. The index must be in the same account and Region as the Amazon Lex bot.
        /// </summary>
        public readonly string KendraIndex;
        /// <summary>
        /// A query filter that Amazon Lex sends to Amazon Kendra to filter the response from a query. The filter is in the format defined by Amazon Kendra. For more information, see [Filtering queries](https://docs.aws.amazon.com/kendra/latest/dg/filtering.html) .
        /// </summary>
        public readonly string? QueryFilterString;
        /// <summary>
        /// Determines whether the `AMAZON.KendraSearchIntent` intent uses a custom query string to query the Amazon Kendra index.
        /// </summary>
        public readonly bool? QueryFilterStringEnabled;

        [OutputConstructor]
        private BotKendraConfiguration(
            string kendraIndex,

            string? queryFilterString,

            bool? queryFilterStringEnabled)
        {
            KendraIndex = kendraIndex;
            QueryFilterString = queryFilterString;
            QueryFilterStringEnabled = queryFilterStringEnabled;
        }
    }
}
