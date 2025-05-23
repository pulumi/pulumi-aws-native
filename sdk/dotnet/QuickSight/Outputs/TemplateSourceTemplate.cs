// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;The source template of the template.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class TemplateSourceTemplate
    {
        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) of the resource.&lt;/p&gt;
        /// </summary>
        public readonly string Arn;

        [OutputConstructor]
        private TemplateSourceTemplate(string arn)
        {
            Arn = arn;
        }
    }
}
