// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QBusiness
{
    public static class GetApplication
    {
        /// <summary>
        /// Definition of AWS::QBusiness::Application Resource Type
        /// </summary>
        public static Task<GetApplicationResult> InvokeAsync(GetApplicationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationResult>("aws-native:qbusiness:getApplication", args ?? new GetApplicationArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::QBusiness::Application Resource Type
        /// </summary>
        public static Output<GetApplicationResult> Invoke(GetApplicationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationResult>("aws-native:qbusiness:getApplication", args ?? new GetApplicationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier for the Amazon Q Business application.
        /// </summary>
        [Input("applicationId", required: true)]
        public string ApplicationId { get; set; } = null!;

        public GetApplicationArgs()
        {
        }
        public static new GetApplicationArgs Empty => new GetApplicationArgs();
    }

    public sealed class GetApplicationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier for the Amazon Q Business application.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        public GetApplicationInvokeArgs()
        {
        }
        public static new GetApplicationInvokeArgs Empty => new GetApplicationInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon Q Business application.
        /// </summary>
        public readonly string? ApplicationArn;
        /// <summary>
        /// The identifier for the Amazon Q Business application.
        /// </summary>
        public readonly string? ApplicationId;
        /// <summary>
        /// Configuration information for the file upload during chat feature.
        /// </summary>
        public readonly Outputs.ApplicationAttachmentsConfiguration? AttachmentsConfiguration;
        /// <summary>
        /// The Unix timestamp when the Amazon Q Business application was created.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// A description for the Amazon Q Business application.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The name of the Amazon Q Business application.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the AWS IAM Identity Center instance attached to your Amazon Q Business application.
        /// </summary>
        public readonly string? IdentityCenterApplicationArn;
        /// <summary>
        /// Configuration information about Amazon Q Apps. (preview feature)
        /// </summary>
        public readonly Outputs.ApplicationQAppsConfiguration? QAppsConfiguration;
        /// <summary>
        /// The Amazon Resource Name (ARN) of an IAM role with permissions to access your Amazon CloudWatch logs and metrics.
        /// </summary>
        public readonly string? RoleArn;
        /// <summary>
        /// The status of the Amazon Q Business application. The application is ready to use when the status is `ACTIVE` .
        /// </summary>
        public readonly Pulumi.AwsNative.QBusiness.ApplicationStatus? Status;
        /// <summary>
        /// A list of key-value pairs that identify or categorize your Amazon Q Business application. You can also use tags to help control access to the application. Tag keys and values can consist of Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The Unix timestamp when the Amazon Q Business application was last updated.
        /// </summary>
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private GetApplicationResult(
            string? applicationArn,

            string? applicationId,

            Outputs.ApplicationAttachmentsConfiguration? attachmentsConfiguration,

            string? createdAt,

            string? description,

            string? displayName,

            string? identityCenterApplicationArn,

            Outputs.ApplicationQAppsConfiguration? qAppsConfiguration,

            string? roleArn,

            Pulumi.AwsNative.QBusiness.ApplicationStatus? status,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            string? updatedAt)
        {
            ApplicationArn = applicationArn;
            ApplicationId = applicationId;
            AttachmentsConfiguration = attachmentsConfiguration;
            CreatedAt = createdAt;
            Description = description;
            DisplayName = displayName;
            IdentityCenterApplicationArn = identityCenterApplicationArn;
            QAppsConfiguration = qAppsConfiguration;
            RoleArn = roleArn;
            Status = status;
            Tags = tags;
            UpdatedAt = updatedAt;
        }
    }
}