// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Eks
{
    /// <summary>
    /// An object representing an Amazon EKS AccessEntry.
    /// </summary>
    [AwsNativeResourceType("aws-native:eks:AccessEntry")]
    public partial class AccessEntry : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the access entry.
        /// </summary>
        [Output("accessEntryArn")]
        public Output<string> AccessEntryArn { get; private set; } = null!;

        /// <summary>
        /// An array of access policies that are associated with the access entry.
        /// </summary>
        [Output("accessPolicies")]
        public Output<ImmutableArray<Outputs.AccessEntryAccessPolicy>> AccessPolicies { get; private set; } = null!;

        /// <summary>
        /// The cluster that the access entry is created for.
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// The Kubernetes groups that the access entry is associated with.
        /// </summary>
        [Output("kubernetesGroups")]
        public Output<ImmutableArray<string>> KubernetesGroups { get; private set; } = null!;

        /// <summary>
        /// The principal ARN that the access entry is created for.
        /// </summary>
        [Output("principalArn")]
        public Output<string> PrincipalArn { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.AccessEntryTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The node type to associate with the access entry.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// The Kubernetes user that the access entry is associated with.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a AccessEntry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessEntry(string name, AccessEntryArgs args, CustomResourceOptions? options = null)
            : base("aws-native:eks:AccessEntry", name, args ?? new AccessEntryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessEntry(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:eks:AccessEntry", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "clusterName",
                    "principalArn",
                    "type",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AccessEntry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessEntry Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AccessEntry(name, id, options);
        }
    }

    public sealed class AccessEntryArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessPolicies")]
        private InputList<Inputs.AccessEntryAccessPolicyArgs>? _accessPolicies;

        /// <summary>
        /// An array of access policies that are associated with the access entry.
        /// </summary>
        public InputList<Inputs.AccessEntryAccessPolicyArgs> AccessPolicies
        {
            get => _accessPolicies ?? (_accessPolicies = new InputList<Inputs.AccessEntryAccessPolicyArgs>());
            set => _accessPolicies = value;
        }

        /// <summary>
        /// The cluster that the access entry is created for.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        [Input("kubernetesGroups")]
        private InputList<string>? _kubernetesGroups;

        /// <summary>
        /// The Kubernetes groups that the access entry is associated with.
        /// </summary>
        public InputList<string> KubernetesGroups
        {
            get => _kubernetesGroups ?? (_kubernetesGroups = new InputList<string>());
            set => _kubernetesGroups = value;
        }

        /// <summary>
        /// The principal ARN that the access entry is created for.
        /// </summary>
        [Input("principalArn", required: true)]
        public Input<string> PrincipalArn { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.AccessEntryTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.AccessEntryTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.AccessEntryTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The node type to associate with the access entry.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The Kubernetes user that the access entry is associated with.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public AccessEntryArgs()
        {
        }
        public static new AccessEntryArgs Empty => new AccessEntryArgs();
    }
}