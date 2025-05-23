// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Outputs
{

    [OutputType]
    public sealed class DomainConfigurationClientCertificateConfig
    {
        /// <summary>
        /// The ARN of the Lambda function that IoT invokes after mutual TLS authentication during the connection.
        /// </summary>
        public readonly string? ClientCertificateCallbackArn;

        [OutputConstructor]
        private DomainConfigurationClientCertificateConfig(string? clientCertificateCallbackArn)
        {
            ClientCertificateCallbackArn = clientCertificateCallbackArn;
        }
    }
}
