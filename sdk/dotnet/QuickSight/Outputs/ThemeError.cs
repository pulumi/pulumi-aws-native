// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;Theme error.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class ThemeError
    {
        /// <summary>
        /// &lt;p&gt;The error message.&lt;/p&gt;
        /// </summary>
        public readonly string? Message;
        public readonly Pulumi.AwsNative.QuickSight.ThemeErrorType? Type;

        [OutputConstructor]
        private ThemeError(
            string? message,

            Pulumi.AwsNative.QuickSight.ThemeErrorType? type)
        {
            Message = message;
            Type = type;
        }
    }
}