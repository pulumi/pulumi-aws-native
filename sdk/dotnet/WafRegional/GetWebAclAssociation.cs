// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WafRegional
{
    public static class GetWebAclAssociation
    {
        /// <summary>
        /// Resource Type definition for AWS::WAFRegional::WebACLAssociation
        /// </summary>
        public static Task<GetWebAclAssociationResult> InvokeAsync(GetWebAclAssociationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebAclAssociationResult>("aws-native:wafregional:getWebAclAssociation", args ?? new GetWebAclAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::WAFRegional::WebACLAssociation
        /// </summary>
        public static Output<GetWebAclAssociationResult> Invoke(GetWebAclAssociationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebAclAssociationResult>("aws-native:wafregional:getWebAclAssociation", args ?? new GetWebAclAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebAclAssociationArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetWebAclAssociationArgs()
        {
        }
        public static new GetWebAclAssociationArgs Empty => new GetWebAclAssociationArgs();
    }

    public sealed class GetWebAclAssociationInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetWebAclAssociationInvokeArgs()
        {
        }
        public static new GetWebAclAssociationInvokeArgs Empty => new GetWebAclAssociationInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebAclAssociationResult
    {
        public readonly string? Id;

        [OutputConstructor]
        private GetWebAclAssociationResult(string? id)
        {
            Id = id;
        }
    }
}