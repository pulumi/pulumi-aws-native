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
    /// Xss Match Statement.
    /// </summary>
    [OutputType]
    public sealed class WebAclXssMatchStatement
    {
        /// <summary>
        /// The part of the web request that you want AWS WAF to inspect.
        /// </summary>
        public readonly Outputs.WebAclFieldToMatch FieldToMatch;
        /// <summary>
        /// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection. If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
        /// </summary>
        public readonly ImmutableArray<Outputs.WebAclTextTransformation> TextTransformations;

        [OutputConstructor]
        private WebAclXssMatchStatement(
            Outputs.WebAclFieldToMatch fieldToMatch,

            ImmutableArray<Outputs.WebAclTextTransformation> textTransformations)
        {
            FieldToMatch = fieldToMatch;
            TextTransformations = textTransformations;
        }
    }
}
