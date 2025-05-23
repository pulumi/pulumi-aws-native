// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SsmContacts.Outputs
{

    /// <summary>
    /// The contacts or contact methods that the escalation plan or engagement plan is engaging.
    /// </summary>
    [OutputType]
    public sealed class ContactTargets
    {
        /// <summary>
        /// Information about the contact channel that Incident Manager engages.
        /// </summary>
        public readonly Outputs.ContactChannelTargetInfo? ChannelTargetInfo;
        /// <summary>
        /// The contact that Incident Manager is engaging during an incident.
        /// </summary>
        public readonly Outputs.ContactTargetInfo? ContactTargetInfo;

        [OutputConstructor]
        private ContactTargets(
            Outputs.ContactChannelTargetInfo? channelTargetInfo,

            Outputs.ContactTargetInfo? contactTargetInfo)
        {
            ChannelTargetInfo = channelTargetInfo;
            ContactTargetInfo = contactTargetInfo;
        }
    }
}
