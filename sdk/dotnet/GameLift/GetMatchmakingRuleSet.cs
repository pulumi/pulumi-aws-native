// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift
{
    public static class GetMatchmakingRuleSet
    {
        /// <summary>
        /// Resource Type definition for AWS::GameLift::MatchmakingRuleSet
        /// </summary>
        public static Task<GetMatchmakingRuleSetResult> InvokeAsync(GetMatchmakingRuleSetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMatchmakingRuleSetResult>("aws-native:gamelift:getMatchmakingRuleSet", args ?? new GetMatchmakingRuleSetArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::GameLift::MatchmakingRuleSet
        /// </summary>
        public static Output<GetMatchmakingRuleSetResult> Invoke(GetMatchmakingRuleSetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMatchmakingRuleSetResult>("aws-native:gamelift:getMatchmakingRuleSet", args ?? new GetMatchmakingRuleSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMatchmakingRuleSetArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetMatchmakingRuleSetArgs()
        {
        }
        public static new GetMatchmakingRuleSetArgs Empty => new GetMatchmakingRuleSetArgs();
    }

    public sealed class GetMatchmakingRuleSetInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetMatchmakingRuleSetInvokeArgs()
        {
        }
        public static new GetMatchmakingRuleSetInvokeArgs Empty => new GetMatchmakingRuleSetInvokeArgs();
    }


    [OutputType]
    public sealed class GetMatchmakingRuleSetResult
    {
        public readonly string? Arn;
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.MatchmakingRuleSetTag> Tags;

        [OutputConstructor]
        private GetMatchmakingRuleSetResult(
            string? arn,

            string? id,

            ImmutableArray<Outputs.MatchmakingRuleSetTag> tags)
        {
            Arn = arn;
            Id = id;
            Tags = tags;
        }
    }
}