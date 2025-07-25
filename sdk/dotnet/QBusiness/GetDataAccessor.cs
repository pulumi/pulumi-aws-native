// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QBusiness
{
    public static class GetDataAccessor
    {
        /// <summary>
        /// Definition of AWS::QBusiness::DataAccessor Resource Type
        /// </summary>
        public static Task<GetDataAccessorResult> InvokeAsync(GetDataAccessorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDataAccessorResult>("aws-native:qbusiness:getDataAccessor", args ?? new GetDataAccessorArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::QBusiness::DataAccessor Resource Type
        /// </summary>
        public static Output<GetDataAccessorResult> Invoke(GetDataAccessorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDataAccessorResult>("aws-native:qbusiness:getDataAccessor", args ?? new GetDataAccessorInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::QBusiness::DataAccessor Resource Type
        /// </summary>
        public static Output<GetDataAccessorResult> Invoke(GetDataAccessorInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDataAccessorResult>("aws-native:qbusiness:getDataAccessor", args ?? new GetDataAccessorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataAccessorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier of the Amazon Q Business application.
        /// </summary>
        [Input("applicationId", required: true)]
        public string ApplicationId { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the data accessor.
        /// </summary>
        [Input("dataAccessorId", required: true)]
        public string DataAccessorId { get; set; } = null!;

        public GetDataAccessorArgs()
        {
        }
        public static new GetDataAccessorArgs Empty => new GetDataAccessorArgs();
    }

    public sealed class GetDataAccessorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier of the Amazon Q Business application.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the data accessor.
        /// </summary>
        [Input("dataAccessorId", required: true)]
        public Input<string> DataAccessorId { get; set; } = null!;

        public GetDataAccessorInvokeArgs()
        {
        }
        public static new GetDataAccessorInvokeArgs Empty => new GetDataAccessorInvokeArgs();
    }


    [OutputType]
    public sealed class GetDataAccessorResult
    {
        /// <summary>
        /// A list of action configurations specifying the allowed actions and any associated filters.
        /// </summary>
        public readonly ImmutableArray<Outputs.DataAccessorActionConfiguration> ActionConfigurations;
        /// <summary>
        /// The authentication configuration details for the data accessor. This specifies how the ISV authenticates when accessing data through this data accessor.
        /// </summary>
        public readonly Outputs.DataAccessorAuthenticationDetail? AuthenticationDetail;
        /// <summary>
        /// The timestamp when the data accessor was created.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the data accessor.
        /// </summary>
        public readonly string? DataAccessorArn;
        /// <summary>
        /// The unique identifier of the data accessor.
        /// </summary>
        public readonly string? DataAccessorId;
        /// <summary>
        /// The friendly name of the data accessor.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the associated IAM Identity Center application.
        /// </summary>
        public readonly string? IdcApplicationArn;
        /// <summary>
        /// The tags to associate with the data accessor.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The timestamp when the data accessor was last updated.
        /// </summary>
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private GetDataAccessorResult(
            ImmutableArray<Outputs.DataAccessorActionConfiguration> actionConfigurations,

            Outputs.DataAccessorAuthenticationDetail? authenticationDetail,

            string? createdAt,

            string? dataAccessorArn,

            string? dataAccessorId,

            string? displayName,

            string? idcApplicationArn,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            string? updatedAt)
        {
            ActionConfigurations = actionConfigurations;
            AuthenticationDetail = authenticationDetail;
            CreatedAt = createdAt;
            DataAccessorArn = dataAccessorArn;
            DataAccessorId = dataAccessorId;
            DisplayName = displayName;
            IdcApplicationArn = idcApplicationArn;
            Tags = tags;
            UpdatedAt = updatedAt;
        }
    }
}
