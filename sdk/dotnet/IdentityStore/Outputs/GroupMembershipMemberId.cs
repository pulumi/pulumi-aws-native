// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IdentityStore.Outputs
{

    /// <summary>
    /// An object containing the identifier of a group member.
    /// </summary>
    [OutputType]
    public sealed class GroupMembershipMemberId
    {
        /// <summary>
        /// The identifier for a user in the identity store.
        /// </summary>
        public readonly string UserId;

        [OutputConstructor]
        private GroupMembershipMemberId(string userId)
        {
            UserId = userId;
        }
    }
}