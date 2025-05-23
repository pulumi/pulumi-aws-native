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
    /// &lt;p&gt;List of errors that occurred when the template version creation failed.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class TemplateError
    {
        /// <summary>
        /// &lt;p&gt;Description of the error type.&lt;/p&gt;
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// Type of error.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.TemplateErrorType? Type;
        /// <summary>
        /// &lt;p&gt;An error path that shows which entities caused the template error.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.TemplateEntity> ViolatedEntities;

        [OutputConstructor]
        private TemplateError(
            string? message,

            Pulumi.AwsNative.QuickSight.TemplateErrorType? type,

            ImmutableArray<Outputs.TemplateEntity> violatedEntities)
        {
            Message = message;
            Type = type;
            ViolatedEntities = violatedEntities;
        }
    }
}
