// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight
{
    public static class GetFolder
    {
        /// <summary>
        /// Definition of the AWS::QuickSight::Folder Resource Type.
        /// </summary>
        public static Task<GetFolderResult> InvokeAsync(GetFolderArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFolderResult>("aws-native:quicksight:getFolder", args ?? new GetFolderArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of the AWS::QuickSight::Folder Resource Type.
        /// </summary>
        public static Output<GetFolderResult> Invoke(GetFolderInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFolderResult>("aws-native:quicksight:getFolder", args ?? new GetFolderInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of the AWS::QuickSight::Folder Resource Type.
        /// </summary>
        public static Output<GetFolderResult> Invoke(GetFolderInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFolderResult>("aws-native:quicksight:getFolder", args ?? new GetFolderInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFolderArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID for the AWS account where you want to create the folder.
        /// </summary>
        [Input("awsAccountId", required: true)]
        public string AwsAccountId { get; set; } = null!;

        /// <summary>
        /// The ID of the folder.
        /// </summary>
        [Input("folderId", required: true)]
        public string FolderId { get; set; } = null!;

        public GetFolderArgs()
        {
        }
        public static new GetFolderArgs Empty => new GetFolderArgs();
    }

    public sealed class GetFolderInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID for the AWS account where you want to create the folder.
        /// </summary>
        [Input("awsAccountId", required: true)]
        public Input<string> AwsAccountId { get; set; } = null!;

        /// <summary>
        /// The ID of the folder.
        /// </summary>
        [Input("folderId", required: true)]
        public Input<string> FolderId { get; set; } = null!;

        public GetFolderInvokeArgs()
        {
        }
        public static new GetFolderInvokeArgs Empty => new GetFolderInvokeArgs();
    }


    [OutputType]
    public sealed class GetFolderResult
    {
        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) for the folder.&lt;/p&gt;
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// &lt;p&gt;The time that the folder was created.&lt;/p&gt;
        /// </summary>
        public readonly string? CreatedTime;
        /// <summary>
        /// &lt;p&gt;The time that the folder was last updated.&lt;/p&gt;
        /// </summary>
        public readonly string? LastUpdatedTime;
        /// <summary>
        /// A display name for the folder.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A structure that describes the principals and the resource-level permissions of a folder.
        /// 
        /// To specify no permissions, omit `Permissions` .
        /// </summary>
        public readonly ImmutableArray<Outputs.FolderResourcePermission> Permissions;
        /// <summary>
        /// A list of tags for the folders that you want to apply overrides to.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetFolderResult(
            string? arn,

            string? createdTime,

            string? lastUpdatedTime,

            string? name,

            ImmutableArray<Outputs.FolderResourcePermission> permissions,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Arn = arn;
            CreatedTime = createdTime;
            LastUpdatedTime = lastUpdatedTime;
            Name = name;
            Permissions = permissions;
            Tags = tags;
        }
    }
}
