// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Organizations
{
    public static class GetAccount
    {
        /// <summary>
        /// You can use AWS::Organizations::Account to manage accounts in organization.
        /// </summary>
        public static Task<GetAccountResult> InvokeAsync(GetAccountArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccountResult>("aws-native:organizations:getAccount", args ?? new GetAccountArgs(), options.WithDefaults());

        /// <summary>
        /// You can use AWS::Organizations::Account to manage accounts in organization.
        /// </summary>
        public static Output<GetAccountResult> Invoke(GetAccountInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountResult>("aws-native:organizations:getAccount", args ?? new GetAccountInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// You can use AWS::Organizations::Account to manage accounts in organization.
        /// </summary>
        public static Output<GetAccountResult> Invoke(GetAccountInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountResult>("aws-native:organizations:getAccount", args ?? new GetAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccountArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// If the account was created successfully, the unique identifier (ID) of the new account.
        /// </summary>
        [Input("accountId", required: true)]
        public string AccountId { get; set; } = null!;

        public GetAccountArgs()
        {
        }
        public static new GetAccountArgs Empty => new GetAccountArgs();
    }

    public sealed class GetAccountInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// If the account was created successfully, the unique identifier (ID) of the new account.
        /// </summary>
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        public GetAccountInvokeArgs()
        {
        }
        public static new GetAccountInvokeArgs Empty => new GetAccountInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccountResult
    {
        /// <summary>
        /// If the account was created successfully, the unique identifier (ID) of the new account.
        /// </summary>
        public readonly string? AccountId;
        /// <summary>
        /// The friendly name of the member account.
        /// </summary>
        public readonly string? AccountName;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the account.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The email address of the owner to assign to the new member account.
        /// </summary>
        public readonly string? Email;
        /// <summary>
        /// The method by which the account joined the organization.
        /// </summary>
        public readonly Pulumi.AwsNative.Organizations.AccountJoinedMethod? JoinedMethod;
        /// <summary>
        /// The date the account became a part of the organization.
        /// </summary>
        public readonly string? JoinedTimestamp;
        /// <summary>
        /// List of parent nodes for the member account. Currently only one parent at a time is supported. Default is root.
        /// </summary>
        public readonly ImmutableArray<string> ParentIds;
        /// <summary>
        /// The status of the account in the organization.
        /// </summary>
        public readonly Pulumi.AwsNative.Organizations.AccountStatus? Status;
        /// <summary>
        /// A list of tags that you want to attach to the newly created account. For each tag in the list, you must specify both a tag key and a value.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetAccountResult(
            string? accountId,

            string? accountName,

            string? arn,

            string? email,

            Pulumi.AwsNative.Organizations.AccountJoinedMethod? joinedMethod,

            string? joinedTimestamp,

            ImmutableArray<string> parentIds,

            Pulumi.AwsNative.Organizations.AccountStatus? status,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            AccountId = accountId;
            AccountName = accountName;
            Arn = arn;
            Email = email;
            JoinedMethod = joinedMethod;
            JoinedTimestamp = joinedTimestamp;
            ParentIds = parentIds;
            Status = status;
            Tags = tags;
        }
    }
}
