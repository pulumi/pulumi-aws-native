// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Fms
{
    public static class GetResourceSet
    {
        /// <summary>
        /// Creates an AWS Firewall Manager resource set.
        /// </summary>
        public static Task<GetResourceSetResult> InvokeAsync(GetResourceSetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResourceSetResult>("aws-native:fms:getResourceSet", args ?? new GetResourceSetArgs(), options.WithDefaults());

        /// <summary>
        /// Creates an AWS Firewall Manager resource set.
        /// </summary>
        public static Output<GetResourceSetResult> Invoke(GetResourceSetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourceSetResult>("aws-native:fms:getResourceSet", args ?? new GetResourceSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourceSetArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetResourceSetArgs()
        {
        }
        public static new GetResourceSetArgs Empty => new GetResourceSetArgs();
    }

    public sealed class GetResourceSetInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetResourceSetInvokeArgs()
        {
        }
        public static new GetResourceSetInvokeArgs Empty => new GetResourceSetInvokeArgs();
    }


    [OutputType]
    public sealed class GetResourceSetResult
    {
        public readonly string? Description;
        public readonly string? Id;
        public readonly string? Name;
        public readonly ImmutableArray<string> ResourceTypeList;
        public readonly ImmutableArray<string> Resources;
        public readonly ImmutableArray<Outputs.ResourceSetTag> Tags;

        [OutputConstructor]
        private GetResourceSetResult(
            string? description,

            string? id,

            string? name,

            ImmutableArray<string> resourceTypeList,

            ImmutableArray<string> resources,

            ImmutableArray<Outputs.ResourceSetTag> tags)
        {
            Description = description;
            Id = id;
            Name = name;
            ResourceTypeList = resourceTypeList;
            Resources = resources;
            Tags = tags;
        }
    }
}