// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AuditManager.Outputs
{

    /// <summary>
    /// An AWS service such as Amazon S3, AWS CloudTrail, and so on.
    /// </summary>
    [OutputType]
    public sealed class AssessmentAWSService
    {
        public readonly string? ServiceName;

        [OutputConstructor]
        private AssessmentAWSService(string? serviceName)
        {
            ServiceName = serviceName;
        }
    }
}