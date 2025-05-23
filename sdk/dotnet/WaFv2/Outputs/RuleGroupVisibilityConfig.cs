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
    /// Visibility Metric of the RuleGroup.
    /// </summary>
    [OutputType]
    public sealed class RuleGroupVisibilityConfig
    {
        /// <summary>
        /// Indicates whether the associated resource sends metrics to Amazon CloudWatch. For the list of available metrics, see [AWS WAF Metrics](https://docs.aws.amazon.com/waf/latest/developerguide/monitoring-cloudwatch.html#waf-metrics) in the *AWS WAF Developer Guide* .
        /// 
        /// For web ACLs, the metrics are for web requests that have the web ACL default action applied. AWS WAF applies the default action to web requests that pass the inspection of all rules in the web ACL without being either allowed or blocked. For more information,
        /// see [The web ACL default action](https://docs.aws.amazon.com/waf/latest/developerguide/web-acl-default-action.html) in the *AWS WAF Developer Guide* .
        /// </summary>
        public readonly bool CloudWatchMetricsEnabled;
        /// <summary>
        /// A name of the Amazon CloudWatch metric dimension. The name can contain only the characters: A-Z, a-z, 0-9, - (hyphen), and _ (underscore). The name can be from one to 128 characters long. It can't contain whitespace or metric names that are reserved for AWS WAF , for example `All` and `Default_Action` .
        /// </summary>
        public readonly string MetricName;
        /// <summary>
        /// Indicates whether AWS WAF should store a sampling of the web requests that match the rules. You can view the sampled requests through the AWS WAF console.
        /// 
        /// If you configure data protection for the web ACL, the protection applies to the web ACL's sampled web request data.
        /// 
        /// &gt; Request sampling doesn't provide a field redaction option, and any field redaction that you specify in your logging configuration doesn't affect sampling. You can only exclude fields from request sampling by disabling sampling in the web ACL visibility configuration or by configuring data protection for the web ACL.
        /// </summary>
        public readonly bool SampledRequestsEnabled;

        [OutputConstructor]
        private RuleGroupVisibilityConfig(
            bool cloudWatchMetricsEnabled,

            string metricName,

            bool sampledRequestsEnabled)
        {
            CloudWatchMetricsEnabled = cloudWatchMetricsEnabled;
            MetricName = metricName;
            SampledRequestsEnabled = sampledRequestsEnabled;
        }
    }
}
