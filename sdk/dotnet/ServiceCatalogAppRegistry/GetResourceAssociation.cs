// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceCatalogAppRegistry
{
    public static class GetResourceAssociation
    {
        /// <summary>
        /// Resource Schema for AWS::ServiceCatalogAppRegistry::ResourceAssociation
        /// </summary>
        public static Task<GetResourceAssociationResult> InvokeAsync(GetResourceAssociationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResourceAssociationResult>("aws-native:servicecatalogappregistry:getResourceAssociation", args ?? new GetResourceAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Schema for AWS::ServiceCatalogAppRegistry::ResourceAssociation
        /// </summary>
        public static Output<GetResourceAssociationResult> Invoke(GetResourceAssociationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourceAssociationResult>("aws-native:servicecatalogappregistry:getResourceAssociation", args ?? new GetResourceAssociationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Schema for AWS::ServiceCatalogAppRegistry::ResourceAssociation
        /// </summary>
        public static Output<GetResourceAssociationResult> Invoke(GetResourceAssociationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourceAssociationResult>("aws-native:servicecatalogappregistry:getResourceAssociation", args ?? new GetResourceAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourceAssociationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon resource name (ARN) that specifies the application.
        /// </summary>
        [Input("applicationArn", required: true)]
        public string ApplicationArn { get; set; } = null!;

        /// <summary>
        /// The Amazon resource name (ARN) that specifies the resource.
        /// </summary>
        [Input("resourceArn", required: true)]
        public string ResourceArn { get; set; } = null!;

        /// <summary>
        /// The type of the CFN Resource for now it's enum CFN_STACK.
        /// </summary>
        [Input("resourceType", required: true)]
        public Pulumi.AwsNative.ServiceCatalogAppRegistry.ResourceAssociationResourceType ResourceType { get; set; }

        public GetResourceAssociationArgs()
        {
        }
        public static new GetResourceAssociationArgs Empty => new GetResourceAssociationArgs();
    }

    public sealed class GetResourceAssociationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon resource name (ARN) that specifies the application.
        /// </summary>
        [Input("applicationArn", required: true)]
        public Input<string> ApplicationArn { get; set; } = null!;

        /// <summary>
        /// The Amazon resource name (ARN) that specifies the resource.
        /// </summary>
        [Input("resourceArn", required: true)]
        public Input<string> ResourceArn { get; set; } = null!;

        /// <summary>
        /// The type of the CFN Resource for now it's enum CFN_STACK.
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<Pulumi.AwsNative.ServiceCatalogAppRegistry.ResourceAssociationResourceType> ResourceType { get; set; } = null!;

        public GetResourceAssociationInvokeArgs()
        {
        }
        public static new GetResourceAssociationInvokeArgs Empty => new GetResourceAssociationInvokeArgs();
    }


    [OutputType]
    public sealed class GetResourceAssociationResult
    {
        /// <summary>
        /// The Amazon resource name (ARN) that specifies the application.
        /// </summary>
        public readonly string? ApplicationArn;
        /// <summary>
        /// The Amazon resource name (ARN) that specifies the resource.
        /// </summary>
        public readonly string? ResourceArn;

        [OutputConstructor]
        private GetResourceAssociationResult(
            string? applicationArn,

            string? resourceArn)
        {
            ApplicationArn = applicationArn;
            ResourceArn = resourceArn;
        }
    }
}
