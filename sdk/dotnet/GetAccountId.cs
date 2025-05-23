// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative
{
    public static class GetAccountId
    {
        public static Task<GetAccountIdResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccountIdResult>("aws-native:index:getAccountId", InvokeArgs.Empty, options.WithDefaults());

        public static Output<GetAccountIdResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountIdResult>("aws-native:index:getAccountId", InvokeArgs.Empty, options.WithDefaults());

        public static Output<GetAccountIdResult> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountIdResult>("aws-native:index:getAccountId", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetAccountIdResult
    {
        public readonly string AccountId;

        [OutputConstructor]
        private GetAccountIdResult(string accountId)
        {
            AccountId = accountId;
        }
    }
}
