// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SsmContacts.Inputs
{

    /// <summary>
    /// The contacts or contact methods that the escalation plan or engagement plan is engaging.
    /// </summary>
    public sealed class ContactTargetsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Information about the contact channel that Incident Manager engages.
        /// </summary>
        [Input("channelTargetInfo")]
        public Input<Inputs.ContactChannelTargetInfoArgs>? ChannelTargetInfo { get; set; }

        /// <summary>
        /// The contact that Incident Manager is engaging during an incident.
        /// </summary>
        [Input("contactTargetInfo")]
        public Input<Inputs.ContactTargetInfoArgs>? ContactTargetInfo { get; set; }

        public ContactTargetsArgs()
        {
        }
        public static new ContactTargetsArgs Empty => new ContactTargetsArgs();
    }
}
