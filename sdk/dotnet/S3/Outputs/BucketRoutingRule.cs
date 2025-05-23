// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    /// <summary>
    /// Specifies the redirect behavior and when a redirect is applied. For more information about routing rules, see [Configuring advanced conditional redirects](https://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html#advanced-conditional-redirects) in the *Amazon S3 User Guide*.
    /// </summary>
    [OutputType]
    public sealed class BucketRoutingRule
    {
        /// <summary>
        /// Container for redirect information. You can redirect requests to another host, to another page, or with another protocol. In the event of an error, you can specify a different error code to return.
        /// </summary>
        public readonly Outputs.BucketRedirectRule RedirectRule;
        /// <summary>
        /// A container for describing a condition that must be met for the specified redirect to apply. For example, 1. If request is for pages in the ``/docs`` folder, redirect to the ``/documents`` folder. 2. If request results in HTTP error 4xx, redirect request to another host where you might process the error.
        /// </summary>
        public readonly Outputs.BucketRoutingRuleCondition? RoutingRuleCondition;

        [OutputConstructor]
        private BucketRoutingRule(
            Outputs.BucketRedirectRule redirectRule,

            Outputs.BucketRoutingRuleCondition? routingRuleCondition)
        {
            RedirectRule = redirectRule;
            RoutingRuleCondition = routingRuleCondition;
        }
    }
}
