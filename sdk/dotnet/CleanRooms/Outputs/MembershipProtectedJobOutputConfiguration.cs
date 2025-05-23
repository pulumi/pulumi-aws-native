// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Outputs
{

    [OutputType]
    public sealed class MembershipProtectedJobOutputConfiguration
    {
        /// <summary>
        /// Contains the configuration to write the job results to S3.
        /// </summary>
        public readonly Outputs.MembershipProtectedJobS3OutputConfigurationInput S3;

        [OutputConstructor]
        private MembershipProtectedJobOutputConfiguration(Outputs.MembershipProtectedJobS3OutputConfigurationInput s3)
        {
            S3 = s3;
        }
    }
}
