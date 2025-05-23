// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    /// <summary>
    /// Byte Match statement.
    /// </summary>
    [OutputType]
    public sealed class RuleGroupByteMatchStatement
    {
        /// <summary>
        /// The part of the web request that you want AWS WAF to inspect.
        /// </summary>
        public readonly Outputs.RuleGroupFieldToMatch FieldToMatch;
        /// <summary>
        /// The area within the portion of the web request that you want AWS WAF to search for `SearchString` . Valid values include the following:
        /// 
        /// *CONTAINS*
        /// 
        /// The specified part of the web request must include the value of `SearchString` , but the location doesn't matter.
        /// 
        /// *CONTAINS_WORD*
        /// 
        /// The specified part of the web request must include the value of `SearchString` , and `SearchString` must contain only alphanumeric characters or underscore (A-Z, a-z, 0-9, or _). In addition, `SearchString` must be a word, which means that both of the following are true:
        /// 
        /// - `SearchString` is at the beginning of the specified part of the web request or is preceded by a character other than an alphanumeric character or underscore (_). Examples include the value of a header and `;BadBot` .
        /// - `SearchString` is at the end of the specified part of the web request or is followed by a character other than an alphanumeric character or underscore (_), for example, `BadBot;` and `-BadBot;` .
        /// 
        /// *EXACTLY*
        /// 
        /// The value of the specified part of the web request must exactly match the value of `SearchString` .
        /// 
        /// *STARTS_WITH*
        /// 
        /// The value of `SearchString` must appear at the beginning of the specified part of the web request.
        /// 
        /// *ENDS_WITH*
        /// 
        /// The value of `SearchString` must appear at the end of the specified part of the web request.
        /// </summary>
        public readonly Pulumi.AwsNative.WaFv2.RuleGroupPositionalConstraint PositionalConstraint;
        /// <summary>
        /// A string value that you want AWS WAF to search for. AWS WAF searches only in the part of web requests that you designate for inspection in `FieldToMatch` . The maximum length of the value is 200 bytes. For alphabetic characters A-Z and a-z, the value is case sensitive.
        /// 
        /// Don't encode this string. Provide the value that you want AWS WAF to search for. AWS CloudFormation automatically base64 encodes the value for you.
        /// 
        /// For example, suppose the value of `Type` is `HEADER` and the value of `Data` is `User-Agent` . If you want to search the `User-Agent` header for the value `BadBot` , you provide the string `BadBot` in the value of `SearchString` .
        /// 
        /// You must specify either `SearchString` or `SearchStringBase64` in a `ByteMatchStatement` .
        /// </summary>
        public readonly string? SearchString;
        /// <summary>
        /// String to search for in a web request component, base64-encoded. If you don't want to encode the string, specify the unencoded value in `SearchString` instead.
        /// 
        /// You must specify either `SearchString` or `SearchStringBase64` in a `ByteMatchStatement` .
        /// </summary>
        public readonly string? SearchStringBase64;
        /// <summary>
        /// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection. If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleGroupTextTransformation> TextTransformations;

        [OutputConstructor]
        private RuleGroupByteMatchStatement(
            Outputs.RuleGroupFieldToMatch fieldToMatch,

            Pulumi.AwsNative.WaFv2.RuleGroupPositionalConstraint positionalConstraint,

            string? searchString,

            string? searchStringBase64,

            ImmutableArray<Outputs.RuleGroupTextTransformation> textTransformations)
        {
            FieldToMatch = fieldToMatch;
            PositionalConstraint = positionalConstraint;
            SearchString = searchString;
            SearchStringBase64 = searchStringBase64;
            TextTransformations = textTransformations;
        }
    }
}
