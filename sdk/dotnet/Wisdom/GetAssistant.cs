// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom
{
    public static class GetAssistant
    {
        /// <summary>
        /// Definition of AWS::Wisdom::Assistant Resource Type
        /// </summary>
        public static Task<GetAssistantResult> InvokeAsync(GetAssistantArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAssistantResult>("aws-native:wisdom:getAssistant", args ?? new GetAssistantArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Wisdom::Assistant Resource Type
        /// </summary>
        public static Output<GetAssistantResult> Invoke(GetAssistantInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAssistantResult>("aws-native:wisdom:getAssistant", args ?? new GetAssistantInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Wisdom::Assistant Resource Type
        /// </summary>
        public static Output<GetAssistantResult> Invoke(GetAssistantInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAssistantResult>("aws-native:wisdom:getAssistant", args ?? new GetAssistantInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAssistantArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Wisdom assistant.
        /// </summary>
        [Input("assistantId", required: true)]
        public string AssistantId { get; set; } = null!;

        public GetAssistantArgs()
        {
        }
        public static new GetAssistantArgs Empty => new GetAssistantArgs();
    }

    public sealed class GetAssistantInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Wisdom assistant.
        /// </summary>
        [Input("assistantId", required: true)]
        public Input<string> AssistantId { get; set; } = null!;

        public GetAssistantInvokeArgs()
        {
        }
        public static new GetAssistantInvokeArgs Empty => new GetAssistantInvokeArgs();
    }


    [OutputType]
    public sealed class GetAssistantResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the assistant.
        /// </summary>
        public readonly string? AssistantArn;
        /// <summary>
        /// The ID of the Wisdom assistant.
        /// </summary>
        public readonly string? AssistantId;

        [OutputConstructor]
        private GetAssistantResult(
            string? assistantArn,

            string? assistantId)
        {
            AssistantArn = assistantArn;
            AssistantId = assistantId;
        }
    }
}
