// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SsmContacts
{
    public static class GetContactChannel
    {
        /// <summary>
        /// Resource Type definition for AWS::SSMContacts::ContactChannel
        /// </summary>
        public static Task<GetContactChannelResult> InvokeAsync(GetContactChannelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContactChannelResult>("aws-native:ssmcontacts:getContactChannel", args ?? new GetContactChannelArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SSMContacts::ContactChannel
        /// </summary>
        public static Output<GetContactChannelResult> Invoke(GetContactChannelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContactChannelResult>("aws-native:ssmcontacts:getContactChannel", args ?? new GetContactChannelInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SSMContacts::ContactChannel
        /// </summary>
        public static Output<GetContactChannelResult> Invoke(GetContactChannelInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetContactChannelResult>("aws-native:ssmcontacts:getContactChannel", args ?? new GetContactChannelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContactChannelArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the engagement to a contact channel.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetContactChannelArgs()
        {
        }
        public static new GetContactChannelArgs Empty => new GetContactChannelArgs();
    }

    public sealed class GetContactChannelInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the engagement to a contact channel.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetContactChannelInvokeArgs()
        {
        }
        public static new GetContactChannelInvokeArgs Empty => new GetContactChannelInvokeArgs();
    }


    [OutputType]
    public sealed class GetContactChannelResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the engagement to a contact channel.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The details that SSM Incident Manager uses when trying to engage the contact channel.
        /// </summary>
        public readonly string? ChannelAddress;
        /// <summary>
        /// The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters.
        /// </summary>
        public readonly string? ChannelName;

        [OutputConstructor]
        private GetContactChannelResult(
            string? arn,

            string? channelAddress,

            string? channelName)
        {
            Arn = arn;
            ChannelAddress = channelAddress;
            ChannelName = channelName;
        }
    }
}
