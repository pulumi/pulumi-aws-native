// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex
{
    public static class GetBotVersion
    {
        /// <summary>
        /// A version is a numbered snapshot of your work that you can publish for use in different parts of your workflow, such as development, beta deployment, and production.
        /// </summary>
        public static Task<GetBotVersionResult> InvokeAsync(GetBotVersionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBotVersionResult>("aws-native:lex:getBotVersion", args ?? new GetBotVersionArgs(), options.WithDefaults());

        /// <summary>
        /// A version is a numbered snapshot of your work that you can publish for use in different parts of your workflow, such as development, beta deployment, and production.
        /// </summary>
        public static Output<GetBotVersionResult> Invoke(GetBotVersionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBotVersionResult>("aws-native:lex:getBotVersion", args ?? new GetBotVersionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// A version is a numbered snapshot of your work that you can publish for use in different parts of your workflow, such as development, beta deployment, and production.
        /// </summary>
        public static Output<GetBotVersionResult> Invoke(GetBotVersionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetBotVersionResult>("aws-native:lex:getBotVersion", args ?? new GetBotVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBotVersionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier of the bot.
        /// </summary>
        [Input("botId", required: true)]
        public string BotId { get; set; } = null!;

        /// <summary>
        /// The version of the bot.
        /// </summary>
        [Input("botVersion", required: true)]
        public string BotVersionValue { get; set; } = null!;

        public GetBotVersionArgs()
        {
        }
        public static new GetBotVersionArgs Empty => new GetBotVersionArgs();
    }

    public sealed class GetBotVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier of the bot.
        /// </summary>
        [Input("botId", required: true)]
        public Input<string> BotId { get; set; } = null!;

        /// <summary>
        /// The version of the bot.
        /// </summary>
        [Input("botVersion", required: true)]
        public Input<string> BotVersion { get; set; } = null!;

        public GetBotVersionInvokeArgs()
        {
        }
        public static new GetBotVersionInvokeArgs Empty => new GetBotVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetBotVersionResult
    {
        /// <summary>
        /// The version of the bot.
        /// </summary>
        public readonly string? BotVersionValue;
        /// <summary>
        /// The description of the version.
        /// </summary>
        public readonly string? Description;

        [OutputConstructor]
        private GetBotVersionResult(
            string? botVersion,

            string? description)
        {
            BotVersionValue = botVersion;
            Description = description;
        }
    }
}
