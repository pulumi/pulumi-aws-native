// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses
{
    public static class GetContactList
    {
        /// <summary>
        /// Resource schema for AWS::SES::ContactList.
        /// </summary>
        public static Task<GetContactListResult> InvokeAsync(GetContactListArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContactListResult>("aws-native:ses:getContactList", args ?? new GetContactListArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::SES::ContactList.
        /// </summary>
        public static Output<GetContactListResult> Invoke(GetContactListInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContactListResult>("aws-native:ses:getContactList", args ?? new GetContactListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContactListArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the contact list.
        /// </summary>
        [Input("contactListName", required: true)]
        public string ContactListName { get; set; } = null!;

        public GetContactListArgs()
        {
        }
        public static new GetContactListArgs Empty => new GetContactListArgs();
    }

    public sealed class GetContactListInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the contact list.
        /// </summary>
        [Input("contactListName", required: true)]
        public Input<string> ContactListName { get; set; } = null!;

        public GetContactListInvokeArgs()
        {
        }
        public static new GetContactListInvokeArgs Empty => new GetContactListInvokeArgs();
    }


    [OutputType]
    public sealed class GetContactListResult
    {
        /// <summary>
        /// The description of the contact list.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The tags (keys and values) associated with the contact list.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContactListTag> Tags;
        /// <summary>
        /// The topics associated with the contact list.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContactListTopic> Topics;

        [OutputConstructor]
        private GetContactListResult(
            string? description,

            ImmutableArray<Outputs.ContactListTag> tags,

            ImmutableArray<Outputs.ContactListTopic> topics)
        {
            Description = description;
            Tags = tags;
            Topics = topics;
        }
    }
}