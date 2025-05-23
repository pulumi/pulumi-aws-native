// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses
{
    public static class GetMailManagerArchive
    {
        /// <summary>
        /// Definition of AWS::SES::MailManagerArchive Resource Type
        /// </summary>
        public static Task<GetMailManagerArchiveResult> InvokeAsync(GetMailManagerArchiveArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMailManagerArchiveResult>("aws-native:ses:getMailManagerArchive", args ?? new GetMailManagerArchiveArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::SES::MailManagerArchive Resource Type
        /// </summary>
        public static Output<GetMailManagerArchiveResult> Invoke(GetMailManagerArchiveInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMailManagerArchiveResult>("aws-native:ses:getMailManagerArchive", args ?? new GetMailManagerArchiveInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::SES::MailManagerArchive Resource Type
        /// </summary>
        public static Output<GetMailManagerArchiveResult> Invoke(GetMailManagerArchiveInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMailManagerArchiveResult>("aws-native:ses:getMailManagerArchive", args ?? new GetMailManagerArchiveInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMailManagerArchiveArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier of the archive.
        /// </summary>
        [Input("archiveId", required: true)]
        public string ArchiveId { get; set; } = null!;

        public GetMailManagerArchiveArgs()
        {
        }
        public static new GetMailManagerArchiveArgs Empty => new GetMailManagerArchiveArgs();
    }

    public sealed class GetMailManagerArchiveInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier of the archive.
        /// </summary>
        [Input("archiveId", required: true)]
        public Input<string> ArchiveId { get; set; } = null!;

        public GetMailManagerArchiveInvokeArgs()
        {
        }
        public static new GetMailManagerArchiveInvokeArgs Empty => new GetMailManagerArchiveInvokeArgs();
    }


    [OutputType]
    public sealed class GetMailManagerArchiveResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the archive.
        /// </summary>
        public readonly string? ArchiveArn;
        /// <summary>
        /// The unique identifier of the archive.
        /// </summary>
        public readonly string? ArchiveId;
        /// <summary>
        /// A unique name for the new archive.
        /// </summary>
        public readonly string? ArchiveName;
        /// <summary>
        /// The current state of the archive:
        /// 
        /// - `ACTIVE` – The archive is ready and available for use.
        /// - `PENDING_DELETION` – The archive has been marked for deletion and will be permanently deleted in 30 days. No further modifications can be made in this state.
        /// </summary>
        public readonly Pulumi.AwsNative.Ses.MailManagerArchiveArchiveState? ArchiveState;
        /// <summary>
        /// The period for retaining emails in the archive before automatic deletion.
        /// </summary>
        public readonly Outputs.MailManagerArchiveArchiveRetentionProperties? Retention;
        /// <summary>
        /// The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetMailManagerArchiveResult(
            string? archiveArn,

            string? archiveId,

            string? archiveName,

            Pulumi.AwsNative.Ses.MailManagerArchiveArchiveState? archiveState,

            Outputs.MailManagerArchiveArchiveRetentionProperties? retention,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            ArchiveArn = archiveArn;
            ArchiveId = archiveId;
            ArchiveName = archiveName;
            ArchiveState = archiveState;
            Retention = retention;
            Tags = tags;
        }
    }
}
