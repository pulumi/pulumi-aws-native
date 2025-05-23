// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms
{
    public static class GetCollaboration
    {
        /// <summary>
        /// Represents a collaboration between AWS accounts that allows for secure data collaboration
        /// </summary>
        public static Task<GetCollaborationResult> InvokeAsync(GetCollaborationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCollaborationResult>("aws-native:cleanrooms:getCollaboration", args ?? new GetCollaborationArgs(), options.WithDefaults());

        /// <summary>
        /// Represents a collaboration between AWS accounts that allows for secure data collaboration
        /// </summary>
        public static Output<GetCollaborationResult> Invoke(GetCollaborationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCollaborationResult>("aws-native:cleanrooms:getCollaboration", args ?? new GetCollaborationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Represents a collaboration between AWS accounts that allows for secure data collaboration
        /// </summary>
        public static Output<GetCollaborationResult> Invoke(GetCollaborationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCollaborationResult>("aws-native:cleanrooms:getCollaboration", args ?? new GetCollaborationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCollaborationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Returns the unique identifier of the specified collaboration.
        /// 
        /// Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`
        /// </summary>
        [Input("collaborationIdentifier", required: true)]
        public string CollaborationIdentifier { get; set; } = null!;

        public GetCollaborationArgs()
        {
        }
        public static new GetCollaborationArgs Empty => new GetCollaborationArgs();
    }

    public sealed class GetCollaborationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Returns the unique identifier of the specified collaboration.
        /// 
        /// Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`
        /// </summary>
        [Input("collaborationIdentifier", required: true)]
        public Input<string> CollaborationIdentifier { get; set; } = null!;

        public GetCollaborationInvokeArgs()
        {
        }
        public static new GetCollaborationInvokeArgs Empty => new GetCollaborationInvokeArgs();
    }


    [OutputType]
    public sealed class GetCollaborationResult
    {
        /// <summary>
        /// The analytics engine for the collaboration.
        /// </summary>
        public readonly Pulumi.AwsNative.CleanRooms.CollaborationAnalyticsEngine? AnalyticsEngine;
        /// <summary>
        /// Returns the Amazon Resource Name (ARN) of the specified collaboration.
        /// 
        /// Example: `arn:aws:cleanrooms:us-east-1:111122223333:collaboration/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Returns the unique identifier of the specified collaboration.
        /// 
        /// Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`
        /// </summary>
        public readonly string? CollaborationIdentifier;
        /// <summary>
        /// A description of the collaboration provided by the collaboration owner.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// A human-readable identifier provided by the collaboration owner. Display names are not unique.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetCollaborationResult(
            Pulumi.AwsNative.CleanRooms.CollaborationAnalyticsEngine? analyticsEngine,

            string? arn,

            string? collaborationIdentifier,

            string? description,

            string? name,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            AnalyticsEngine = analyticsEngine;
            Arn = arn;
            CollaborationIdentifier = collaborationIdentifier;
            Description = description;
            Name = name;
            Tags = tags;
        }
    }
}
