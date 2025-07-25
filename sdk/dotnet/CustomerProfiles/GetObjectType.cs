// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles
{
    public static class GetObjectType
    {
        /// <summary>
        /// An ObjectType resource of Amazon Connect Customer Profiles
        /// </summary>
        public static Task<GetObjectTypeResult> InvokeAsync(GetObjectTypeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetObjectTypeResult>("aws-native:customerprofiles:getObjectType", args ?? new GetObjectTypeArgs(), options.WithDefaults());

        /// <summary>
        /// An ObjectType resource of Amazon Connect Customer Profiles
        /// </summary>
        public static Output<GetObjectTypeResult> Invoke(GetObjectTypeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetObjectTypeResult>("aws-native:customerprofiles:getObjectType", args ?? new GetObjectTypeInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// An ObjectType resource of Amazon Connect Customer Profiles
        /// </summary>
        public static Output<GetObjectTypeResult> Invoke(GetObjectTypeInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetObjectTypeResult>("aws-native:customerprofiles:getObjectType", args ?? new GetObjectTypeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetObjectTypeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name of the domain.
        /// </summary>
        [Input("domainName", required: true)]
        public string DomainName { get; set; } = null!;

        /// <summary>
        /// The name of the profile object type.
        /// </summary>
        [Input("objectTypeName", required: true)]
        public string ObjectTypeName { get; set; } = null!;

        public GetObjectTypeArgs()
        {
        }
        public static new GetObjectTypeArgs Empty => new GetObjectTypeArgs();
    }

    public sealed class GetObjectTypeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name of the domain.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// The name of the profile object type.
        /// </summary>
        [Input("objectTypeName", required: true)]
        public Input<string> ObjectTypeName { get; set; } = null!;

        public GetObjectTypeInvokeArgs()
        {
        }
        public static new GetObjectTypeInvokeArgs Empty => new GetObjectTypeInvokeArgs();
    }


    [OutputType]
    public sealed class GetObjectTypeResult
    {
        /// <summary>
        /// Indicates whether a profile should be created when data is received.
        /// </summary>
        public readonly bool? AllowProfileCreation;
        /// <summary>
        /// The time of this integration got created.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// Description of the profile object type.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The default encryption key
        /// </summary>
        public readonly string? EncryptionKey;
        /// <summary>
        /// The default number of days until the data within the domain expires.
        /// </summary>
        public readonly int? ExpirationDays;
        /// <summary>
        /// A list of the name and ObjectType field.
        /// </summary>
        public readonly ImmutableArray<Outputs.ObjectTypeFieldMap> Fields;
        /// <summary>
        /// A list of unique keys that can be used to map data to the profile.
        /// </summary>
        public readonly ImmutableArray<Outputs.ObjectTypeKeyMap> Keys;
        /// <summary>
        /// The time of this integration got last updated at.
        /// </summary>
        public readonly string? LastUpdatedAt;
        /// <summary>
        /// The maximum available number of profile objects
        /// </summary>
        public readonly int? MaxAvailableProfileObjectCount;
        /// <summary>
        /// The maximum number of profile objects for this object type
        /// </summary>
        public readonly int? MaxProfileObjectCount;
        /// <summary>
        /// The format of your sourceLastUpdatedTimestamp that was previously set up.
        /// </summary>
        public readonly string? SourceLastUpdatedTimestampFormat;
        /// <summary>
        /// The tags (keys and values) associated with the integration.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// A unique identifier for the object template.
        /// </summary>
        public readonly string? TemplateId;

        [OutputConstructor]
        private GetObjectTypeResult(
            bool? allowProfileCreation,

            string? createdAt,

            string? description,

            string? encryptionKey,

            int? expirationDays,

            ImmutableArray<Outputs.ObjectTypeFieldMap> fields,

            ImmutableArray<Outputs.ObjectTypeKeyMap> keys,

            string? lastUpdatedAt,

            int? maxAvailableProfileObjectCount,

            int? maxProfileObjectCount,

            string? sourceLastUpdatedTimestampFormat,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            string? templateId)
        {
            AllowProfileCreation = allowProfileCreation;
            CreatedAt = createdAt;
            Description = description;
            EncryptionKey = encryptionKey;
            ExpirationDays = expirationDays;
            Fields = fields;
            Keys = keys;
            LastUpdatedAt = lastUpdatedAt;
            MaxAvailableProfileObjectCount = maxAvailableProfileObjectCount;
            MaxProfileObjectCount = maxProfileObjectCount;
            SourceLastUpdatedTimestampFormat = sourceLastUpdatedTimestampFormat;
            Tags = tags;
            TemplateId = templateId;
        }
    }
}
