// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Transfer.Outputs
{

    [OutputType]
    public sealed class UserPosixProfile
    {
        public readonly double Gid;
        public readonly ImmutableArray<double> SecondaryGids;
        public readonly double Uid;

        [OutputConstructor]
        private UserPosixProfile(
            double gid,

            ImmutableArray<double> secondaryGids,

            double uid)
        {
            Gid = gid;
            SecondaryGids = secondaryGids;
            Uid = uid;
        }
    }
}