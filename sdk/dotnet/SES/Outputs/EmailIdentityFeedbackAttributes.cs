// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SES.Outputs
{

    /// <summary>
    /// Used to enable or disable feedback forwarding for an identity.
    /// </summary>
    [OutputType]
    public sealed class EmailIdentityFeedbackAttributes
    {
        /// <summary>
        /// If the value is true, you receive email notifications when bounce or complaint events occur
        /// </summary>
        public readonly bool? EmailForwardingEnabled;

        [OutputConstructor]
        private EmailIdentityFeedbackAttributes(bool? emailForwardingEnabled)
        {
            EmailForwardingEnabled = emailForwardingEnabled;
        }
    }
}