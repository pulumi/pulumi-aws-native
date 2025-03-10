// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift
{
    public static class GetAlias
    {
        /// <summary>
        /// The AWS::GameLift::Alias resource creates an alias for an Amazon GameLift (GameLift) fleet destination.
        /// </summary>
        public static Task<GetAliasResult> InvokeAsync(GetAliasArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAliasResult>("aws-native:gamelift:getAlias", args ?? new GetAliasArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::GameLift::Alias resource creates an alias for an Amazon GameLift (GameLift) fleet destination.
        /// </summary>
        public static Output<GetAliasResult> Invoke(GetAliasInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAliasResult>("aws-native:gamelift:getAlias", args ?? new GetAliasInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::GameLift::Alias resource creates an alias for an Amazon GameLift (GameLift) fleet destination.
        /// </summary>
        public static Output<GetAliasResult> Invoke(GetAliasInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAliasResult>("aws-native:gamelift:getAlias", args ?? new GetAliasInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAliasArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique alias ID
        /// </summary>
        [Input("aliasId", required: true)]
        public string AliasId { get; set; } = null!;

        public GetAliasArgs()
        {
        }
        public static new GetAliasArgs Empty => new GetAliasArgs();
    }

    public sealed class GetAliasInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique alias ID
        /// </summary>
        [Input("aliasId", required: true)]
        public Input<string> AliasId { get; set; } = null!;

        public GetAliasInvokeArgs()
        {
        }
        public static new GetAliasInvokeArgs Empty => new GetAliasInvokeArgs();
    }


    [OutputType]
    public sealed class GetAliasResult
    {
        /// <summary>
        /// Unique alias ID
        /// </summary>
        public readonly string? AliasId;
        /// <summary>
        /// A human-readable description of the alias.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// A descriptive label that is associated with an alias. Alias names do not need to be unique.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A routing configuration that specifies where traffic is directed for this alias, such as to a fleet or to a message.
        /// </summary>
        public readonly Outputs.AliasRoutingStrategy? RoutingStrategy;

        [OutputConstructor]
        private GetAliasResult(
            string? aliasId,

            string? description,

            string? name,

            Outputs.AliasRoutingStrategy? routingStrategy)
        {
            AliasId = aliasId;
            Description = description;
            Name = name;
            RoutingStrategy = routingStrategy;
        }
    }
}
