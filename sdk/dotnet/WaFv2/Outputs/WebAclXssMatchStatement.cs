// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
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
        public readonly Outputs.WebAclFieldToMatch FieldToMatch;
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