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
    public sealed class MembershipProtectedQueryOutputConfiguration
    {
        /// <summary>
        /// Required configuration for a protected query with an `s3` output type.
        /// </summary>
        public readonly Outputs.MembershipProtectedQueryS3OutputConfiguration S3;

        [OutputConstructor]
        private MembershipProtectedQueryOutputConfiguration(Outputs.MembershipProtectedQueryS3OutputConfiguration s3)
        {
            S3 = s3;
        }
    }
}
