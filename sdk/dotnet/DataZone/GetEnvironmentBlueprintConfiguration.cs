// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone
{
    public static class GetEnvironmentBlueprintConfiguration
    {
        /// <summary>
        /// Definition of AWS::DataZone::EnvironmentBlueprintConfiguration Resource Type
        /// </summary>
        public static Task<GetEnvironmentBlueprintConfigurationResult> InvokeAsync(GetEnvironmentBlueprintConfigurationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEnvironmentBlueprintConfigurationResult>("aws-native:datazone:getEnvironmentBlueprintConfiguration", args ?? new GetEnvironmentBlueprintConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::DataZone::EnvironmentBlueprintConfiguration Resource Type
        /// </summary>
        public static Output<GetEnvironmentBlueprintConfigurationResult> Invoke(GetEnvironmentBlueprintConfigurationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnvironmentBlueprintConfigurationResult>("aws-native:datazone:getEnvironmentBlueprintConfiguration", args ?? new GetEnvironmentBlueprintConfigurationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::DataZone::EnvironmentBlueprintConfiguration Resource Type
        /// </summary>
        public static Output<GetEnvironmentBlueprintConfigurationResult> Invoke(GetEnvironmentBlueprintConfigurationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnvironmentBlueprintConfigurationResult>("aws-native:datazone:getEnvironmentBlueprintConfiguration", args ?? new GetEnvironmentBlueprintConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnvironmentBlueprintConfigurationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier of the Amazon DataZone domain in which an environment blueprint exists.
        /// </summary>
        [Input("domainId", required: true)]
        public string DomainId { get; set; } = null!;

        /// <summary>
        /// The identifier of the environment blueprint. This identifier should be used when creating environment profiles.
        /// </summary>
        [Input("environmentBlueprintId", required: true)]
        public string EnvironmentBlueprintId { get; set; } = null!;

        public GetEnvironmentBlueprintConfigurationArgs()
        {
        }
        public static new GetEnvironmentBlueprintConfigurationArgs Empty => new GetEnvironmentBlueprintConfigurationArgs();
    }

    public sealed class GetEnvironmentBlueprintConfigurationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier of the Amazon DataZone domain in which an environment blueprint exists.
        /// </summary>
        [Input("domainId", required: true)]
        public Input<string> DomainId { get; set; } = null!;

        /// <summary>
        /// The identifier of the environment blueprint. This identifier should be used when creating environment profiles.
        /// </summary>
        [Input("environmentBlueprintId", required: true)]
        public Input<string> EnvironmentBlueprintId { get; set; } = null!;

        public GetEnvironmentBlueprintConfigurationInvokeArgs()
        {
        }
        public static new GetEnvironmentBlueprintConfigurationInvokeArgs Empty => new GetEnvironmentBlueprintConfigurationInvokeArgs();
    }


    [OutputType]
    public sealed class GetEnvironmentBlueprintConfigurationResult
    {
        /// <summary>
        /// The timestamp of when an environment blueprint was created.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// The identifier of the Amazon DataZone domain in which an environment blueprint exists.
        /// </summary>
        public readonly string? DomainId;
        /// <summary>
        /// The enabled AWS Regions specified in a blueprint configuration.
        /// </summary>
        public readonly ImmutableArray<string> EnabledRegions;
        /// <summary>
        /// The identifier of the environment blueprint. This identifier should be used when creating environment profiles.
        /// </summary>
        public readonly string? EnvironmentBlueprintId;
        /// <summary>
        /// The ARN of the manage access role.
        /// </summary>
        public readonly string? ManageAccessRoleArn;
        /// <summary>
        /// The ARN of the provisioning role.
        /// </summary>
        public readonly string? ProvisioningRoleArn;
        /// <summary>
        /// The regional parameters of the environment blueprint.
        /// </summary>
        public readonly ImmutableArray<Outputs.EnvironmentBlueprintConfigurationRegionalParameter> RegionalParameters;
        /// <summary>
        /// The timestamp of when the environment blueprint was updated.
        /// </summary>
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private GetEnvironmentBlueprintConfigurationResult(
            string? createdAt,

            string? domainId,

            ImmutableArray<string> enabledRegions,

            string? environmentBlueprintId,

            string? manageAccessRoleArn,

            string? provisioningRoleArn,

            ImmutableArray<Outputs.EnvironmentBlueprintConfigurationRegionalParameter> regionalParameters,

            string? updatedAt)
        {
            CreatedAt = createdAt;
            DomainId = domainId;
            EnabledRegions = enabledRegions;
            EnvironmentBlueprintId = environmentBlueprintId;
            ManageAccessRoleArn = manageAccessRoleArn;
            ProvisioningRoleArn = provisioningRoleArn;
            RegionalParameters = regionalParameters;
            UpdatedAt = updatedAt;
        }
    }
}
