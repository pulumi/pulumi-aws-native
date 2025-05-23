// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone
{
    public static class GetDomainUnit
    {
        /// <summary>
        /// A domain unit enables you to easily organize your assets and other domain entities under specific business units and teams.
        /// </summary>
        public static Task<GetDomainUnitResult> InvokeAsync(GetDomainUnitArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDomainUnitResult>("aws-native:datazone:getDomainUnit", args ?? new GetDomainUnitArgs(), options.WithDefaults());

        /// <summary>
        /// A domain unit enables you to easily organize your assets and other domain entities under specific business units and teams.
        /// </summary>
        public static Output<GetDomainUnitResult> Invoke(GetDomainUnitInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainUnitResult>("aws-native:datazone:getDomainUnit", args ?? new GetDomainUnitInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// A domain unit enables you to easily organize your assets and other domain entities under specific business units and teams.
        /// </summary>
        public static Output<GetDomainUnitResult> Invoke(GetDomainUnitInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainUnitResult>("aws-native:datazone:getDomainUnit", args ?? new GetDomainUnitInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainUnitArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the domain where the domain unit was created.
        /// </summary>
        [Input("domainId", required: true)]
        public string DomainId { get; set; } = null!;

        /// <summary>
        /// The ID of the domain unit.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDomainUnitArgs()
        {
        }
        public static new GetDomainUnitArgs Empty => new GetDomainUnitArgs();
    }

    public sealed class GetDomainUnitInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the domain where the domain unit was created.
        /// </summary>
        [Input("domainId", required: true)]
        public Input<string> DomainId { get; set; } = null!;

        /// <summary>
        /// The ID of the domain unit.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDomainUnitInvokeArgs()
        {
        }
        public static new GetDomainUnitInvokeArgs Empty => new GetDomainUnitInvokeArgs();
    }


    [OutputType]
    public sealed class GetDomainUnitResult
    {
        /// <summary>
        /// The timestamp at which the domain unit was created.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// The description of the domain unit.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The ID of the domain where the domain unit was created.
        /// </summary>
        public readonly string? DomainId;
        /// <summary>
        /// The ID of the domain unit.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The identifier of the domain unit that you want to get.
        /// </summary>
        public readonly string? Identifier;
        /// <summary>
        /// The timestamp at which the domain unit was last updated.
        /// </summary>
        public readonly string? LastUpdatedAt;
        /// <summary>
        /// The name of the domain unit.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The ID of the parent domain unit.
        /// </summary>
        public readonly string? ParentDomainUnitId;

        [OutputConstructor]
        private GetDomainUnitResult(
            string? createdAt,

            string? description,

            string? domainId,

            string? id,

            string? identifier,

            string? lastUpdatedAt,

            string? name,

            string? parentDomainUnitId)
        {
            CreatedAt = createdAt;
            Description = description;
            DomainId = domainId;
            Id = id;
            Identifier = identifier;
            LastUpdatedAt = lastUpdatedAt;
            Name = name;
            ParentDomainUnitId = parentDomainUnitId;
        }
    }
}
