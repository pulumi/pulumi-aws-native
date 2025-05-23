// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchService.Outputs
{

    [OutputType]
    public sealed class DomainEndpointOptions
    {
        /// <summary>
        /// The fully qualified URL for your custom endpoint. Required if you enabled a custom endpoint for the domain.
        /// </summary>
        public readonly string? CustomEndpoint;
        /// <summary>
        /// The AWS Certificate Manager ARN for your domain's SSL/TLS certificate. Required if you enabled a custom endpoint for the domain.
        /// </summary>
        public readonly string? CustomEndpointCertificateArn;
        /// <summary>
        /// True to enable a custom endpoint for the domain. If enabled, you must also provide values for `CustomEndpoint` and `CustomEndpointCertificateArn` .
        /// </summary>
        public readonly bool? CustomEndpointEnabled;
        /// <summary>
        /// True to require that all traffic to the domain arrive over HTTPS. Required if you enable fine-grained access control in [AdvancedSecurityOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.html) .
        /// </summary>
        public readonly bool? EnforceHttps;
        /// <summary>
        /// The minimum TLS version required for traffic to the domain. The policy can be one of the following values:
        /// 
        /// - *Policy-Min-TLS-1-0-2019-07:* TLS security policy that supports TLS version 1.0 to TLS version 1.2
        /// - *Policy-Min-TLS-1-2-2019-07:* TLS security policy that supports only TLS version 1.2
        /// - *Policy-Min-TLS-1-2-PFS-2023-10:* TLS security policy that supports TLS version 1.2 to TLS version 1.3 with perfect forward secrecy cipher suites
        /// </summary>
        public readonly string? TlsSecurityPolicy;

        [OutputConstructor]
        private DomainEndpointOptions(
            string? customEndpoint,

            string? customEndpointCertificateArn,

            bool? customEndpointEnabled,

            bool? enforceHttps,

            string? tlsSecurityPolicy)
        {
            CustomEndpoint = customEndpoint;
            CustomEndpointCertificateArn = customEndpointCertificateArn;
            CustomEndpointEnabled = customEndpointEnabled;
            EnforceHttps = enforceHttps;
            TlsSecurityPolicy = tlsSecurityPolicy;
        }
    }
}
