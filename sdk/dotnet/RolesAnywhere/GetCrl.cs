// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RolesAnywhere
{
    public static class GetCrl
    {
        /// <summary>
        /// Definition of AWS::RolesAnywhere::CRL Resource Type
        /// </summary>
        public static Task<GetCrlResult> InvokeAsync(GetCrlArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCrlResult>("aws-native:rolesanywhere:getCrl", args ?? new GetCrlArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::RolesAnywhere::CRL Resource Type
        /// </summary>
        public static Output<GetCrlResult> Invoke(GetCrlInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCrlResult>("aws-native:rolesanywhere:getCrl", args ?? new GetCrlInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCrlArgs : global::Pulumi.InvokeArgs
    {
        [Input("crlId", required: true)]
        public string CrlId { get; set; } = null!;

        public GetCrlArgs()
        {
        }
        public static new GetCrlArgs Empty => new GetCrlArgs();
    }

    public sealed class GetCrlInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("crlId", required: true)]
        public Input<string> CrlId { get; set; } = null!;

        public GetCrlInvokeArgs()
        {
        }
        public static new GetCrlInvokeArgs Empty => new GetCrlInvokeArgs();
    }


    [OutputType]
    public sealed class GetCrlResult
    {
        public readonly string? CrlData;
        public readonly string? CrlId;
        public readonly bool? Enabled;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.CrlTag> Tags;
        public readonly string? TrustAnchorArn;

        [OutputConstructor]
        private GetCrlResult(
            string? crlData,

            string? crlId,

            bool? enabled,

            string? name,

            ImmutableArray<Outputs.CrlTag> tags,

            string? trustAnchorArn)
        {
            CrlData = crlData;
            CrlId = crlId;
            Enabled = enabled;
            Name = name;
            Tags = tags;
            TrustAnchorArn = trustAnchorArn;
        }
    }
}