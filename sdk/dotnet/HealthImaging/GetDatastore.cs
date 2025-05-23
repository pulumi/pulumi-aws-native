// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.HealthImaging
{
    public static class GetDatastore
    {
        /// <summary>
        /// Definition of AWS::HealthImaging::Datastore Resource Type
        /// </summary>
        public static Task<GetDatastoreResult> InvokeAsync(GetDatastoreArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatastoreResult>("aws-native:healthimaging:getDatastore", args ?? new GetDatastoreArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::HealthImaging::Datastore Resource Type
        /// </summary>
        public static Output<GetDatastoreResult> Invoke(GetDatastoreInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatastoreResult>("aws-native:healthimaging:getDatastore", args ?? new GetDatastoreInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::HealthImaging::Datastore Resource Type
        /// </summary>
        public static Output<GetDatastoreResult> Invoke(GetDatastoreInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatastoreResult>("aws-native:healthimaging:getDatastore", args ?? new GetDatastoreInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatastoreArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The data store identifier.
        /// </summary>
        [Input("datastoreId", required: true)]
        public string DatastoreId { get; set; } = null!;

        public GetDatastoreArgs()
        {
        }
        public static new GetDatastoreArgs Empty => new GetDatastoreArgs();
    }

    public sealed class GetDatastoreInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The data store identifier.
        /// </summary>
        [Input("datastoreId", required: true)]
        public Input<string> DatastoreId { get; set; } = null!;

        public GetDatastoreInvokeArgs()
        {
        }
        public static new GetDatastoreInvokeArgs Empty => new GetDatastoreInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatastoreResult
    {
        /// <summary>
        /// The timestamp when the data store was created.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// The Amazon Resource Name (ARN) for the data store.
        /// </summary>
        public readonly string? DatastoreArn;
        /// <summary>
        /// The data store identifier.
        /// </summary>
        public readonly string? DatastoreId;
        /// <summary>
        /// The data store status.
        /// </summary>
        public readonly Pulumi.AwsNative.HealthImaging.DatastoreStatus? DatastoreStatus;
        /// <summary>
        /// The timestamp when the data store was last updated.
        /// </summary>
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private GetDatastoreResult(
            string? createdAt,

            string? datastoreArn,

            string? datastoreId,

            Pulumi.AwsNative.HealthImaging.DatastoreStatus? datastoreStatus,

            string? updatedAt)
        {
            CreatedAt = createdAt;
            DatastoreArn = datastoreArn;
            DatastoreId = datastoreId;
            DatastoreStatus = datastoreStatus;
            UpdatedAt = updatedAt;
        }
    }
}
