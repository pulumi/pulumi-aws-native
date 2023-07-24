// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Athena.Inputs
{

    /// <summary>
    /// Indicates that an Amazon S3 canned ACL should be set to control ownership of stored query results
    /// </summary>
    public sealed class WorkGroupAclConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("s3AclOption", required: true)]
        public Input<Pulumi.AwsNative.Athena.WorkGroupS3AclOption> S3AclOption { get; set; } = null!;

        public WorkGroupAclConfigurationArgs()
        {
        }
        public static new WorkGroupAclConfigurationArgs Empty => new WorkGroupAclConfigurationArgs();
    }
}