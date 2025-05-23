// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone
{
    public static class GetDomain
    {
        /// <summary>
        /// A domain is an organizing entity for connecting together assets, users, and their projects
        /// </summary>
        public static Task<GetDomainResult> InvokeAsync(GetDomainArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDomainResult>("aws-native:datazone:getDomain", args ?? new GetDomainArgs(), options.WithDefaults());

        /// <summary>
        /// A domain is an organizing entity for connecting together assets, users, and their projects
        /// </summary>
        public static Output<GetDomainResult> Invoke(GetDomainInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainResult>("aws-native:datazone:getDomain", args ?? new GetDomainInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// A domain is an organizing entity for connecting together assets, users, and their projects
        /// </summary>
        public static Output<GetDomainResult> Invoke(GetDomainInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainResult>("aws-native:datazone:getDomain", args ?? new GetDomainInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the Amazon DataZone domain.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDomainArgs()
        {
        }
        public static new GetDomainArgs Empty => new GetDomainArgs();
    }

    public sealed class GetDomainInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the Amazon DataZone domain.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDomainInvokeArgs()
        {
        }
        public static new GetDomainInvokeArgs Empty => new GetDomainInvokeArgs();
    }


    [OutputType]
    public sealed class GetDomainResult
    {
        /// <summary>
        /// The ARN of the Amazon DataZone domain.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The timestamp of when the Amazon DataZone domain was last updated.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// The description of the Amazon DataZone domain.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The domain execution role that is created when an Amazon DataZone domain is created. The domain execution role is created in the AWS account that houses the Amazon DataZone domain.
        /// </summary>
        public readonly string? DomainExecutionRole;
        /// <summary>
        /// The id of the Amazon DataZone domain.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The timestamp of when the Amazon DataZone domain was last updated.
        /// </summary>
        public readonly string? LastUpdatedAt;
        /// <summary>
        /// The identifier of the AWS account that manages the domain.
        /// </summary>
        public readonly string? ManagedAccountId;
        /// <summary>
        /// The name of the Amazon DataZone domain.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The URL of the data portal for this Amazon DataZone domain.
        /// </summary>
        public readonly string? PortalUrl;
        /// <summary>
        /// The ID of the root domain in Amazon Datazone.
        /// </summary>
        public readonly string? RootDomainUnitId;
        /// <summary>
        /// The service role of the domain that is created.
        /// </summary>
        public readonly string? ServiceRole;
        /// <summary>
        /// The single-sign on configuration of the Amazon DataZone domain.
        /// </summary>
        public readonly Outputs.DomainSingleSignOn? SingleSignOn;
        /// <summary>
        /// The status of the Amazon DataZone domain.
        /// </summary>
        public readonly Pulumi.AwsNative.DataZone.DomainStatus? Status;
        /// <summary>
        /// The tags specified for the Amazon DataZone domain.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetDomainResult(
            string? arn,

            string? createdAt,

            string? description,

            string? domainExecutionRole,

            string? id,

            string? lastUpdatedAt,

            string? managedAccountId,

            string? name,

            string? portalUrl,

            string? rootDomainUnitId,

            string? serviceRole,

            Outputs.DomainSingleSignOn? singleSignOn,

            Pulumi.AwsNative.DataZone.DomainStatus? status,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Arn = arn;
            CreatedAt = createdAt;
            Description = description;
            DomainExecutionRole = domainExecutionRole;
            Id = id;
            LastUpdatedAt = lastUpdatedAt;
            ManagedAccountId = managedAccountId;
            Name = name;
            PortalUrl = portalUrl;
            RootDomainUnitId = rootDomainUnitId;
            ServiceRole = serviceRole;
            SingleSignOn = singleSignOn;
            Status = status;
            Tags = tags;
        }
    }
}
