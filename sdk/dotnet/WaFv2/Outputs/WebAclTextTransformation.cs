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
    /// Text Transformation on the Search String before match.
    /// </summary>
    [OutputType]
    public sealed class WebAclTextTransformation
    {
        /// <summary>
        /// Sets the relative processing order for multiple transformations. AWS WAF processes all transformations, from lowest priority to highest, before inspecting the transformed content. The priorities don't need to be consecutive, but they must all be different.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// For detailed descriptions of each of the transformation types, see [Text transformations](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-statement-transformation.html) in the *AWS WAF Developer Guide* .
        /// </summary>
        public readonly Pulumi.AwsNative.WaFv2.WebAclTextTransformationType Type;

        [OutputConstructor]
        private WebAclTextTransformation(
            int priority,

            Pulumi.AwsNative.WaFv2.WebAclTextTransformationType type)
        {
            Priority = priority;
            Type = type;
        }
    }
}
