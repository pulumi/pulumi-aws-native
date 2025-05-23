// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    /// <summary>
    /// Specifies the redirect behavior and when a redirect is applied. For more information about routing rules, see [Configuring advanced conditional redirects](https://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html#advanced-conditional-redirects) in the *Amazon S3 User Guide*.
    /// </summary>
    public sealed class BucketRoutingRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Container for redirect information. You can redirect requests to another host, to another page, or with another protocol. In the event of an error, you can specify a different error code to return.
        /// </summary>
        [Input("redirectRule", required: true)]
        public Input<Inputs.BucketRedirectRuleArgs> RedirectRule { get; set; } = null!;

        /// <summary>
        /// A container for describing a condition that must be met for the specified redirect to apply. For example, 1. If request is for pages in the ``/docs`` folder, redirect to the ``/documents`` folder. 2. If request results in HTTP error 4xx, redirect request to another host where you might process the error.
        /// </summary>
        [Input("routingRuleCondition")]
        public Input<Inputs.BucketRoutingRuleConditionArgs>? RoutingRuleCondition { get; set; }

        public BucketRoutingRuleArgs()
        {
        }
        public static new BucketRoutingRuleArgs Empty => new BucketRoutingRuleArgs();
    }
}
