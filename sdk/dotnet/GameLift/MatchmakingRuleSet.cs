// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift
{
    /// <summary>
    /// The AWS::GameLift::MatchmakingRuleSet resource creates an Amazon GameLift (GameLift) matchmaking rule set.
    /// 
    /// ## Example Usage
    /// ### Example
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AwsNative = Pulumi.AwsNative;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var matchmakingRuleSetResource = new AwsNative.GameLift.MatchmakingRuleSet("matchmakingRuleSetResource", new()
    ///     {
    ///         Name = "MyRuleSet",
    ///         RuleSetBody = "{\"name\": \"MyMatchmakingRuleSet\",\"ruleLanguageVersion\": \"1.0\", \"teams\": [{\"name\": \"MyTeam\",\"minPlayers\": 1,\"maxPlayers\": 20}]}",
    ///     });
    /// 
    ///     var matchmakingConfigurationResource = new AwsNative.GameLift.MatchmakingConfiguration("matchmakingConfigurationResource", new()
    ///     {
    ///         Name = "MyMatchmakingConfiguration",
    ///         AcceptanceRequired = true,
    ///         AcceptanceTimeoutSeconds = 60,
    ///         BackfillMode = AwsNative.GameLift.MatchmakingConfigurationBackfillMode.Manual,
    ///         CustomEventData = "MyCustomEventData",
    ///         Description = "A basic standalone matchmaking configuration",
    ///         FlexMatchMode = AwsNative.GameLift.MatchmakingConfigurationFlexMatchMode.Standalone,
    ///         RequestTimeoutSeconds = 100,
    ///         RuleSetName = matchmakingRuleSetResource.Id,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn =
    ///         {
    ///             matchmakingRuleSetResource,
    ///         },
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// ### Example
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AwsNative = Pulumi.AwsNative;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var matchmakingRuleSet = new AwsNative.GameLift.MatchmakingRuleSet("matchmakingRuleSet", new()
    ///     {
    ///         Name = "MyRuleSet",
    ///         RuleSetBody = "{\"name\": \"MyMatchmakingRuleSet\",\"ruleLanguageVersion\": \"1.0\", \"teams\": [{\"name\": \"MyTeam\",\"minPlayers\": 1,\"maxPlayers\": 20}]}",
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:gamelift:MatchmakingRuleSet")]
    public partial class MatchmakingRuleSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift matchmaking rule set resource and uniquely identifies it.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for the matchmaking rule set.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A collection of matchmaking rules, formatted as a JSON string.
        /// </summary>
        [Output("ruleSetBody")]
        public Output<string> RuleSetBody { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a MatchmakingRuleSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MatchmakingRuleSet(string name, MatchmakingRuleSetArgs args, CustomResourceOptions? options = null)
            : base("aws-native:gamelift:MatchmakingRuleSet", name, args ?? new MatchmakingRuleSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MatchmakingRuleSet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:gamelift:MatchmakingRuleSet", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "name",
                    "ruleSetBody",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MatchmakingRuleSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MatchmakingRuleSet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MatchmakingRuleSet(name, id, options);
        }
    }

    public sealed class MatchmakingRuleSetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique identifier for the matchmaking rule set.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A collection of matchmaking rules, formatted as a JSON string.
        /// </summary>
        [Input("ruleSetBody", required: true)]
        public Input<string> RuleSetBody { get; set; } = null!;

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public MatchmakingRuleSetArgs()
        {
        }
        public static new MatchmakingRuleSetArgs Empty => new MatchmakingRuleSetArgs();
    }
}
