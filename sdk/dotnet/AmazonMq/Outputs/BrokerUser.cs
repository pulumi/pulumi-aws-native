// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmazonMq.Outputs
{

    [OutputType]
    public sealed class BrokerUser
    {
        public readonly bool? ConsoleAccess;
        public readonly ImmutableArray<string> Groups;
        public readonly string Password;
        public readonly bool? ReplicationUser;
        public readonly string Username;

        [OutputConstructor]
        private BrokerUser(
            bool? consoleAccess,

            ImmutableArray<string> groups,

            string password,

            bool? replicationUser,

            string username)
        {
            ConsoleAccess = consoleAccess;
            Groups = groups;
            Password = password;
            ReplicationUser = replicationUser;
            Username = username;
        }
    }
}