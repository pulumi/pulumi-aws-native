// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CertificateManager
{
    public static class GetAccount
    {
        /// <summary>
        /// Resource schema for AWS::CertificateManager::Account.
        /// </summary>
        public static Task<GetAccountResult> InvokeAsync(GetAccountArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccountResult>("aws-native:certificatemanager:getAccount", args ?? new GetAccountArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::CertificateManager::Account.
        /// </summary>
        public static Output<GetAccountResult> Invoke(GetAccountInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountResult>("aws-native:certificatemanager:getAccount", args ?? new GetAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccountArgs : global::Pulumi.InvokeArgs
    {
        [Input("accountId", required: true)]
        public string AccountId { get; set; } = null!;

        public GetAccountArgs()
        {
        }
        public static new GetAccountArgs Empty => new GetAccountArgs();
    }

    public sealed class GetAccountInvokeArgs : global::Pulumi.InvokeArgs
    {
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
        public readonly string? AccountId;
        public readonly Outputs.AccountExpiryEventsConfiguration? ExpiryEventsConfiguration;

        [OutputConstructor]
        private GetAccountResult(
            string? accountId,

            Outputs.AccountExpiryEventsConfiguration? expiryEventsConfiguration)
        {
            AccountId = accountId;
            ExpiryEventsConfiguration = expiryEventsConfiguration;
        }
    }
}