// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// A collection of settings that update the current configuration for the RStudioServerPro Domain-level app.
    /// </summary>
    [OutputType]
    public sealed class DomainRStudioServerProDomainSettings
    {
        public readonly Outputs.DomainResourceSpec? DefaultResourceSpec;
        /// <summary>
        /// The ARN of the execution role for the RStudioServerPro Domain-level app.
        /// </summary>
        public readonly string DomainExecutionRoleArn;
        /// <summary>
        /// A URL pointing to an RStudio Connect server.
        /// </summary>
        public readonly string? RStudioConnectUrl;
        /// <summary>
        /// A URL pointing to an RStudio Package Manager server.
        /// </summary>
        public readonly string? RStudioPackageManagerUrl;

        [OutputConstructor]
        private DomainRStudioServerProDomainSettings(
            Outputs.DomainResourceSpec? defaultResourceSpec,

            string domainExecutionRoleArn,

            string? rStudioConnectUrl,

            string? rStudioPackageManagerUrl)
        {
            DefaultResourceSpec = defaultResourceSpec;
            DomainExecutionRoleArn = domainExecutionRoleArn;
            RStudioConnectUrl = rStudioConnectUrl;
            RStudioPackageManagerUrl = rStudioPackageManagerUrl;
        }
    }
}