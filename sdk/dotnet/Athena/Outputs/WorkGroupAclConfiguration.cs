// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Athena.Outputs
{

    /// <summary>
    /// Indicates that an Amazon S3 canned ACL should be set to control ownership of stored query results
    /// </summary>
    [OutputType]
    public sealed class WorkGroupAclConfiguration
    {
        public readonly Pulumi.AwsNative.Athena.WorkGroupS3AclOption S3AclOption;

        [OutputConstructor]
        private WorkGroupAclConfiguration(Pulumi.AwsNative.Athena.WorkGroupS3AclOption s3AclOption)
        {
            S3AclOption = s3AclOption;
        }
    }
}