// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect
{
    public static class GetView
    {
        /// <summary>
        /// Resource Type definition for AWS::Connect::View
        /// </summary>
        public static Task<GetViewResult> InvokeAsync(GetViewArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetViewResult>("aws-native:connect:getView", args ?? new GetViewArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Connect::View
        /// </summary>
        public static Output<GetViewResult> Invoke(GetViewInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetViewResult>("aws-native:connect:getView", args ?? new GetViewInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetViewArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the view.
        /// </summary>
        [Input("viewArn", required: true)]
        public string ViewArn { get; set; } = null!;

        public GetViewArgs()
        {
        }
        public static new GetViewArgs Empty => new GetViewArgs();
    }

    public sealed class GetViewInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the view.
        /// </summary>
        [Input("viewArn", required: true)]
        public Input<string> ViewArn { get; set; } = null!;

        public GetViewInvokeArgs()
        {
        }
        public static new GetViewInvokeArgs Empty => new GetViewInvokeArgs();
    }


    [OutputType]
    public sealed class GetViewResult
    {
        /// <summary>
        /// The actions of the view in an array.
        /// </summary>
        public readonly ImmutableArray<string> Actions;
        /// <summary>
        /// The description of the view.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the instance.
        /// </summary>
        public readonly string? InstanceArn;
        /// <summary>
        /// The name of the view.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// One or more tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.ViewTag> Tags;
        /// <summary>
        /// The template of the view as JSON.
        /// </summary>
        public readonly object? Template;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the view.
        /// </summary>
        public readonly string? ViewArn;
        /// <summary>
        /// The view content hash.
        /// </summary>
        public readonly string? ViewContentSha256;
        /// <summary>
        /// The view id of the view.
        /// </summary>
        public readonly string? ViewId;

        [OutputConstructor]
        private GetViewResult(
            ImmutableArray<string> actions,

            string? description,

            string? instanceArn,

            string? name,

            ImmutableArray<Outputs.ViewTag> tags,

            object? template,

            string? viewArn,

            string? viewContentSha256,

            string? viewId)
        {
            Actions = actions;
            Description = description;
            InstanceArn = instanceArn;
            Name = name;
            Tags = tags;
            Template = template;
            ViewArn = viewArn;
            ViewContentSha256 = viewContentSha256;
            ViewId = viewId;
        }
    }
}