// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    /// <summary>
    /// Used to enable or disable feedback forwarding for an identity.
    /// </summary>
    public sealed class EmailIdentityFeedbackAttributesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If the value is true, you receive email notifications when bounce or complaint events occur
        /// </summary>
        [Input("emailForwardingEnabled")]
        public Input<bool>? EmailForwardingEnabled { get; set; }

        public EmailIdentityFeedbackAttributesArgs()
        {
        }
        public static new EmailIdentityFeedbackAttributesArgs Empty => new EmailIdentityFeedbackAttributesArgs();
    }
}
