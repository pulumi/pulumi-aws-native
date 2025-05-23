// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs
{
    /// <summary>
    /// Resource Type definition for AWS::ECS::CapacityProvider.
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
    ///     var myCapacityProvider = new AwsNative.Ecs.CapacityProvider("myCapacityProvider", new()
    ///     {
    ///         AutoScalingGroupProvider = new AwsNative.Ecs.Inputs.CapacityProviderAutoScalingGroupProviderArgs
    ///         {
    ///             AutoScalingGroupArn = "arn:aws:autoscaling:us-west-2:123456789012:autoScalingGroup:a1b2c3d4-5678-90ab-cdef-EXAMPLE11111:autoScalingGroupName/MyAutoScalingGroup",
    ///             ManagedScaling = new AwsNative.Ecs.Inputs.CapacityProviderManagedScalingArgs
    ///             {
    ///                 MaximumScalingStepSize = 10,
    ///                 MinimumScalingStepSize = 1,
    ///                 Status = AwsNative.Ecs.CapacityProviderManagedScalingStatus.Enabled,
    ///                 TargetCapacity = 100,
    ///             },
    ///             ManagedTerminationProtection = AwsNative.Ecs.CapacityProviderAutoScalingGroupProviderManagedTerminationProtection.Enabled,
    ///         },
    ///         Tags = new[]
    ///         {
    ///             new AwsNative.Inputs.TagArgs
    ///             {
    ///                 Key = "environment",
    ///                 Value = "production",
    ///             },
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
    ///     var myCapacityProvider = new AwsNative.Ecs.CapacityProvider("myCapacityProvider", new()
    ///     {
    ///         AutoScalingGroupProvider = new AwsNative.Ecs.Inputs.CapacityProviderAutoScalingGroupProviderArgs
    ///         {
    ///             AutoScalingGroupArn = "arn:aws:autoscaling:us-west-2:123456789012:autoScalingGroup:a1b2c3d4-5678-90ab-cdef-EXAMPLE11111:autoScalingGroupName/MyAutoScalingGroup",
    ///             ManagedScaling = new AwsNative.Ecs.Inputs.CapacityProviderManagedScalingArgs
    ///             {
    ///                 MaximumScalingStepSize = 10,
    ///                 MinimumScalingStepSize = 1,
    ///                 Status = AwsNative.Ecs.CapacityProviderManagedScalingStatus.Enabled,
    ///                 TargetCapacity = 100,
    ///             },
    ///             ManagedTerminationProtection = AwsNative.Ecs.CapacityProviderAutoScalingGroupProviderManagedTerminationProtection.Enabled,
    ///         },
    ///         Tags = new[]
    ///         {
    ///             new AwsNative.Inputs.TagArgs
    ///             {
    ///                 Key = "environment",
    ///                 Value = "production",
    ///             },
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
    ///     var config = new Config();
    ///     var autoScalingGroupArn1 = config.Require("autoScalingGroupArn1");
    ///     var autoScalingGroupArn2 = config.Require("autoScalingGroupArn2");
    ///     var capacityProvider1 = new AwsNative.Ecs.CapacityProvider("capacityProvider1", new()
    ///     {
    ///         AutoScalingGroupProvider = new AwsNative.Ecs.Inputs.CapacityProviderAutoScalingGroupProviderArgs
    ///         {
    ///             AutoScalingGroupArn = autoScalingGroupArn1,
    ///             ManagedScaling = new AwsNative.Ecs.Inputs.CapacityProviderManagedScalingArgs
    ///             {
    ///                 Status = AwsNative.Ecs.CapacityProviderManagedScalingStatus.Enabled,
    ///             },
    ///             ManagedTerminationProtection = AwsNative.Ecs.CapacityProviderAutoScalingGroupProviderManagedTerminationProtection.Disabled,
    ///         },
    ///     });
    /// 
    ///     var capacityProvider2 = new AwsNative.Ecs.CapacityProvider("capacityProvider2", new()
    ///     {
    ///         AutoScalingGroupProvider = new AwsNative.Ecs.Inputs.CapacityProviderAutoScalingGroupProviderArgs
    ///         {
    ///             AutoScalingGroupArn = autoScalingGroupArn2,
    ///             ManagedScaling = new AwsNative.Ecs.Inputs.CapacityProviderManagedScalingArgs
    ///             {
    ///                 Status = AwsNative.Ecs.CapacityProviderManagedScalingStatus.Enabled,
    ///             },
    ///             ManagedTerminationProtection = AwsNative.Ecs.CapacityProviderAutoScalingGroupProviderManagedTerminationProtection.Disabled,
    ///         },
    ///     });
    /// 
    ///     var cluster = new AwsNative.Ecs.Cluster("cluster");
    /// 
    ///     var clusterCPAssociation = new AwsNative.Ecs.ClusterCapacityProviderAssociations("clusterCPAssociation", new()
    ///     {
    ///         Cluster = cluster.Id,
    ///         CapacityProviders = new[]
    ///         {
    ///             capacityProvider1.Id,
    ///             capacityProvider2.Id,
    ///         },
    ///         DefaultCapacityProviderStrategy = new[]
    ///         {
    ///             new AwsNative.Ecs.Inputs.ClusterCapacityProviderAssociationsCapacityProviderStrategyArgs
    ///             {
    ///                 Base = 2,
    ///                 Weight = 6,
    ///                 CapacityProvider = capacityProvider1.Id,
    ///             },
    ///             new AwsNative.Ecs.Inputs.ClusterCapacityProviderAssociationsCapacityProviderStrategyArgs
    ///             {
    ///                 Base = 0,
    ///                 Weight = 10,
    ///                 CapacityProvider = capacityProvider2.Id,
    ///             },
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
    ///     var config = new Config();
    ///     var autoScalingGroupArn1 = config.Require("autoScalingGroupArn1");
    ///     var autoScalingGroupArn2 = config.Require("autoScalingGroupArn2");
    ///     var capacityProvider1 = new AwsNative.Ecs.CapacityProvider("capacityProvider1", new()
    ///     {
    ///         AutoScalingGroupProvider = new AwsNative.Ecs.Inputs.CapacityProviderAutoScalingGroupProviderArgs
    ///         {
    ///             AutoScalingGroupArn = autoScalingGroupArn1,
    ///             ManagedScaling = new AwsNative.Ecs.Inputs.CapacityProviderManagedScalingArgs
    ///             {
    ///                 Status = AwsNative.Ecs.CapacityProviderManagedScalingStatus.Enabled,
    ///             },
    ///             ManagedTerminationProtection = AwsNative.Ecs.CapacityProviderAutoScalingGroupProviderManagedTerminationProtection.Disabled,
    ///         },
    ///     });
    /// 
    ///     var capacityProvider2 = new AwsNative.Ecs.CapacityProvider("capacityProvider2", new()
    ///     {
    ///         AutoScalingGroupProvider = new AwsNative.Ecs.Inputs.CapacityProviderAutoScalingGroupProviderArgs
    ///         {
    ///             AutoScalingGroupArn = autoScalingGroupArn2,
    ///             ManagedScaling = new AwsNative.Ecs.Inputs.CapacityProviderManagedScalingArgs
    ///             {
    ///                 Status = AwsNative.Ecs.CapacityProviderManagedScalingStatus.Enabled,
    ///             },
    ///             ManagedTerminationProtection = AwsNative.Ecs.CapacityProviderAutoScalingGroupProviderManagedTerminationProtection.Disabled,
    ///         },
    ///     });
    /// 
    ///     var cluster = new AwsNative.Ecs.Cluster("cluster");
    /// 
    ///     var clusterCPAssociation = new AwsNative.Ecs.ClusterCapacityProviderAssociations("clusterCPAssociation", new()
    ///     {
    ///         Cluster = cluster.Id,
    ///         CapacityProviders = new[]
    ///         {
    ///             capacityProvider1.Id,
    ///             capacityProvider2.Id,
    ///         },
    ///         DefaultCapacityProviderStrategy = new[]
    ///         {
    ///             new AwsNative.Ecs.Inputs.ClusterCapacityProviderAssociationsCapacityProviderStrategyArgs
    ///             {
    ///                 Base = 2,
    ///                 Weight = 6,
    ///                 CapacityProvider = capacityProvider1.Id,
    ///             },
    ///             new AwsNative.Ecs.Inputs.ClusterCapacityProviderAssociationsCapacityProviderStrategyArgs
    ///             {
    ///                 Base = 0,
    ///                 Weight = 10,
    ///                 CapacityProvider = capacityProvider2.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:ecs:CapacityProvider")]
    public partial class CapacityProvider : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Auto Scaling group settings for the capacity provider.
        /// </summary>
        [Output("autoScalingGroupProvider")]
        public Output<Outputs.CapacityProviderAutoScalingGroupProvider?> AutoScalingGroupProvider { get; private set; } = null!;

        /// <summary>
        /// The name of the capacity provider. If a name is specified, it cannot start with `aws` , `ecs` , or `fargate` . If no name is specified, a default name in the `CFNStackName-CFNResourceName-RandomString` format is used.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The metadata that you apply to the capacity provider to help you categorize and organize it. Each tag consists of a key and an optional value. You define both.
        /// 
        /// The following basic restrictions apply to tags:
        /// 
        /// - Maximum number of tags per resource - 50
        /// - For each resource, each tag key must be unique, and each tag key can have only one value.
        /// - Maximum key length - 128 Unicode characters in UTF-8
        /// - Maximum value length - 256 Unicode characters in UTF-8
        /// - If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
        /// - Tag keys and values are case-sensitive.
        /// - Do not use `aws:` , `AWS:` , or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a CapacityProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CapacityProvider(string name, CapacityProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ecs:CapacityProvider", name, args ?? new CapacityProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CapacityProvider(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ecs:CapacityProvider", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "autoScalingGroupProvider.autoScalingGroupArn",
                    "name",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CapacityProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CapacityProvider Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CapacityProvider(name, id, options);
        }
    }

    public sealed class CapacityProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Auto Scaling group settings for the capacity provider.
        /// </summary>
        [Input("autoScalingGroupProvider")]
        public Input<Inputs.CapacityProviderAutoScalingGroupProviderArgs>? AutoScalingGroupProvider { get; set; }

        /// <summary>
        /// The name of the capacity provider. If a name is specified, it cannot start with `aws` , `ecs` , or `fargate` . If no name is specified, a default name in the `CFNStackName-CFNResourceName-RandomString` format is used.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// The metadata that you apply to the capacity provider to help you categorize and organize it. Each tag consists of a key and an optional value. You define both.
        /// 
        /// The following basic restrictions apply to tags:
        /// 
        /// - Maximum number of tags per resource - 50
        /// - For each resource, each tag key must be unique, and each tag key can have only one value.
        /// - Maximum key length - 128 Unicode characters in UTF-8
        /// - Maximum value length - 256 Unicode characters in UTF-8
        /// - If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
        /// - Tag keys and values are case-sensitive.
        /// - Do not use `aws:` , `AWS:` , or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public CapacityProviderArgs()
        {
        }
        public static new CapacityProviderArgs Empty => new CapacityProviderArgs();
    }
}
