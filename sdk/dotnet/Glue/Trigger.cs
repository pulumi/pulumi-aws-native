// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue
{
    /// <summary>
    /// Resource Type definition for AWS::Glue::Trigger
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
    ///     var onDemandJobTrigger = new AwsNative.Glue.Trigger("onDemandJobTrigger", new()
    ///     {
    ///         Type = "ON_DEMAND",
    ///         Description = "DESCRIPTION_ON_DEMAND",
    ///         Actions = new[]
    ///         {
    ///             new AwsNative.Glue.Inputs.TriggerActionArgs
    ///             {
    ///                 JobName = "prod-job2",
    ///             },
    ///         },
    ///         Name = "prod-trigger1-ondemand",
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
    ///     var onDemandJobTrigger = new AwsNative.Glue.Trigger("onDemandJobTrigger", new()
    ///     {
    ///         Type = "ON_DEMAND",
    ///         Description = "DESCRIPTION_ON_DEMAND",
    ///         Actions = new[]
    ///         {
    ///             new AwsNative.Glue.Inputs.TriggerActionArgs
    ///             {
    ///                 JobName = "prod-job2",
    ///             },
    ///         },
    ///         Name = "prod-trigger1-ondemand",
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
    ///     var scheduledJobTrigger = new AwsNative.Glue.Trigger("scheduledJobTrigger", new()
    ///     {
    ///         Type = "SCHEDULED",
    ///         Description = "DESCRIPTION_SCHEDULED",
    ///         Schedule = "cron(0 */2 * * ? *)",
    ///         Actions = new[]
    ///         {
    ///             new AwsNative.Glue.Inputs.TriggerActionArgs
    ///             {
    ///                 JobName = "prod-job2",
    ///             },
    ///             new AwsNative.Glue.Inputs.TriggerActionArgs
    ///             {
    ///                 JobName = "prod-job3",
    ///                 Arguments = new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["--job-bookmark-option"] = "job-bookmark-enable",
    ///                 },
    ///             },
    ///         },
    ///         Name = "prod-trigger1-scheduled",
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
    ///     var scheduledJobTrigger = new AwsNative.Glue.Trigger("scheduledJobTrigger", new()
    ///     {
    ///         Type = "SCHEDULED",
    ///         Description = "DESCRIPTION_SCHEDULED",
    ///         Schedule = "cron(0 */2 * * ? *)",
    ///         Actions = new[]
    ///         {
    ///             new AwsNative.Glue.Inputs.TriggerActionArgs
    ///             {
    ///                 JobName = "prod-job2",
    ///             },
    ///             new AwsNative.Glue.Inputs.TriggerActionArgs
    ///             {
    ///                 JobName = "prod-job3",
    ///                 Arguments = new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["--job-bookmark-option"] = "job-bookmark-enable",
    ///                 },
    ///             },
    ///         },
    ///         Name = "prod-trigger1-scheduled",
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:glue:Trigger")]
    public partial class Trigger : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The actions initiated by this trigger.
        /// </summary>
        [Output("actions")]
        public Output<ImmutableArray<Outputs.TriggerAction>> Actions { get; private set; } = null!;

        /// <summary>
        /// A description of this trigger.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires.
        /// </summary>
        [Output("eventBatchingCondition")]
        public Output<Outputs.TriggerEventBatchingCondition?> EventBatchingCondition { get; private set; } = null!;

        /// <summary>
        /// The name of the trigger.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The predicate of this trigger, which defines when it will fire.
        /// </summary>
        [Output("predicate")]
        public Output<Outputs.TriggerPredicate?> Predicate { get; private set; } = null!;

        /// <summary>
        /// A cron expression used to specify the schedule.
        /// </summary>
        [Output("schedule")]
        public Output<string?> Schedule { get; private set; } = null!;

        /// <summary>
        /// Set to true to start SCHEDULED and CONDITIONAL triggers when created. True is not supported for ON_DEMAND triggers.
        /// </summary>
        [Output("startOnCreation")]
        public Output<bool?> StartOnCreation { get; private set; } = null!;

        /// <summary>
        /// The tags to use with this trigger.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Trigger` for more information about the expected schema for this property.
        /// </summary>
        [Output("tags")]
        public Output<object?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of trigger that this is.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The name of the workflow associated with the trigger.
        /// </summary>
        [Output("workflowName")]
        public Output<string?> WorkflowName { get; private set; } = null!;


        /// <summary>
        /// Create a Trigger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Trigger(string name, TriggerArgs args, CustomResourceOptions? options = null)
            : base("aws-native:glue:Trigger", name, args ?? new TriggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Trigger(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:glue:Trigger", name, null, MakeResourceOptions(options, id))
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
                    "type",
                    "workflowName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Trigger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Trigger Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Trigger(name, id, options);
        }
    }

    public sealed class TriggerArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions", required: true)]
        private InputList<Inputs.TriggerActionArgs>? _actions;

        /// <summary>
        /// The actions initiated by this trigger.
        /// </summary>
        public InputList<Inputs.TriggerActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.TriggerActionArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// A description of this trigger.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires.
        /// </summary>
        [Input("eventBatchingCondition")]
        public Input<Inputs.TriggerEventBatchingConditionArgs>? EventBatchingCondition { get; set; }

        /// <summary>
        /// The name of the trigger.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The predicate of this trigger, which defines when it will fire.
        /// </summary>
        [Input("predicate")]
        public Input<Inputs.TriggerPredicateArgs>? Predicate { get; set; }

        /// <summary>
        /// A cron expression used to specify the schedule.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// Set to true to start SCHEDULED and CONDITIONAL triggers when created. True is not supported for ON_DEMAND triggers.
        /// </summary>
        [Input("startOnCreation")]
        public Input<bool>? StartOnCreation { get; set; }

        /// <summary>
        /// The tags to use with this trigger.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Trigger` for more information about the expected schema for this property.
        /// </summary>
        [Input("tags")]
        public Input<object>? Tags { get; set; }

        /// <summary>
        /// The type of trigger that this is.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The name of the workflow associated with the trigger.
        /// </summary>
        [Input("workflowName")]
        public Input<string>? WorkflowName { get; set; }

        public TriggerArgs()
        {
        }
        public static new TriggerArgs Empty => new TriggerArgs();
    }
}
