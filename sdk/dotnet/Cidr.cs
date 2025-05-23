// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative
{
    public static class Cidr
    {
        public static Task<CidrResult> InvokeAsync(CidrArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<CidrResult>("aws-native:index:cidr", args ?? new CidrArgs(), options.WithDefaults());

        public static Output<CidrResult> Invoke(CidrInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<CidrResult>("aws-native:index:cidr", args ?? new CidrInvokeArgs(), options.WithDefaults());

        public static Output<CidrResult> Invoke(CidrInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<CidrResult>("aws-native:index:cidr", args ?? new CidrInvokeArgs(), options.WithDefaults());
    }


    public sealed class CidrArgs : global::Pulumi.InvokeArgs
    {
        [Input("cidrBits", required: true)]
        public int CidrBits { get; set; }

        [Input("count", required: true)]
        public int Count { get; set; }

        [Input("ipBlock", required: true)]
        public string IpBlock { get; set; } = null!;

        public CidrArgs()
        {
        }
        public static new CidrArgs Empty => new CidrArgs();
    }

    public sealed class CidrInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("cidrBits", required: true)]
        public Input<int> CidrBits { get; set; } = null!;

        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        [Input("ipBlock", required: true)]
        public Input<string> IpBlock { get; set; } = null!;

        public CidrInvokeArgs()
        {
        }
        public static new CidrInvokeArgs Empty => new CidrInvokeArgs();
    }


    [OutputType]
    public sealed class CidrResult
    {
        public readonly ImmutableArray<string> Subnets;

        [OutputConstructor]
        private CidrResult(ImmutableArray<string> subnets)
        {
            Subnets = subnets;
        }
    }
}
