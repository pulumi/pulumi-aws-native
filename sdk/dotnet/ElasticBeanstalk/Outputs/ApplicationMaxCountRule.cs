// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticBeanstalk.Outputs
{

    [OutputType]
    public sealed class ApplicationMaxCountRule
    {
        /// <summary>
        /// Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.
        /// </summary>
        public readonly bool? DeleteSourceFromS3;
        /// <summary>
        /// Specify true to apply the rule, or false to disable it.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Specify the maximum number of application versions to retain.
        /// </summary>
        public readonly int? MaxCount;

        [OutputConstructor]
        private ApplicationMaxCountRule(
            bool? deleteSourceFromS3,

            bool? enabled,

            int? maxCount)
        {
            DeleteSourceFromS3 = deleteSourceFromS3;
            Enabled = enabled;
            MaxCount = maxCount;
        }
    }
}