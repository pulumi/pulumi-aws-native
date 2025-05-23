// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative
{
    public static class GetSsmParameterList
    {
        public static Task<GetSsmParameterListResult> InvokeAsync(GetSsmParameterListArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSsmParameterListResult>("aws-native:index:getSsmParameterList", args ?? new GetSsmParameterListArgs(), options.WithDefaults());

        public static Output<GetSsmParameterListResult> Invoke(GetSsmParameterListInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSsmParameterListResult>("aws-native:index:getSsmParameterList", args ?? new GetSsmParameterListInvokeArgs(), options.WithDefaults());

        public static Output<GetSsmParameterListResult> Invoke(GetSsmParameterListInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSsmParameterListResult>("aws-native:index:getSsmParameterList", args ?? new GetSsmParameterListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSsmParameterListArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetSsmParameterListArgs()
        {
        }
        public static new GetSsmParameterListArgs Empty => new GetSsmParameterListArgs();
    }

    public sealed class GetSsmParameterListInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetSsmParameterListInvokeArgs()
        {
        }
        public static new GetSsmParameterListInvokeArgs Empty => new GetSsmParameterListInvokeArgs();
    }


    [OutputType]
    public sealed class GetSsmParameterListResult
    {
        public readonly ImmutableArray<string> Value;

        [OutputConstructor]
        private GetSsmParameterListResult(ImmutableArray<string> value)
        {
            Value = value;
        }
    }
}
