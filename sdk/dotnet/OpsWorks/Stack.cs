// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpsWorks
{
    /// <summary>
    /// Resource Type definition for AWS::OpsWorks::Stack
    /// </summary>
    [Obsolete(@"Stack is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:opsworks:Stack")]
    public partial class Stack : global::Pulumi.CustomResource
    {
        [Output("agentVersion")]
        public Output<string?> AgentVersion { get; private set; } = null!;

        [Output("attributes")]
        public Output<object?> Attributes { get; private set; } = null!;

        [Output("chefConfiguration")]
        public Output<Outputs.StackChefConfiguration?> ChefConfiguration { get; private set; } = null!;

        [Output("cloneAppIds")]
        public Output<ImmutableArray<string>> CloneAppIds { get; private set; } = null!;

        [Output("clonePermissions")]
        public Output<bool?> ClonePermissions { get; private set; } = null!;

        [Output("configurationManager")]
        public Output<Outputs.StackConfigurationManager?> ConfigurationManager { get; private set; } = null!;

        [Output("customCookbooksSource")]
        public Output<Outputs.StackSource?> CustomCookbooksSource { get; private set; } = null!;

        [Output("customJson")]
        public Output<object?> CustomJson { get; private set; } = null!;

        [Output("defaultAvailabilityZone")]
        public Output<string?> DefaultAvailabilityZone { get; private set; } = null!;

        [Output("defaultInstanceProfileArn")]
        public Output<string> DefaultInstanceProfileArn { get; private set; } = null!;

        [Output("defaultOs")]
        public Output<string?> DefaultOs { get; private set; } = null!;

        [Output("defaultRootDeviceType")]
        public Output<string?> DefaultRootDeviceType { get; private set; } = null!;

        [Output("defaultSshKeyName")]
        public Output<string?> DefaultSshKeyName { get; private set; } = null!;

        [Output("defaultSubnetId")]
        public Output<string?> DefaultSubnetId { get; private set; } = null!;

        [Output("ecsClusterArn")]
        public Output<string?> EcsClusterArn { get; private set; } = null!;

        [Output("elasticIps")]
        public Output<ImmutableArray<Outputs.StackElasticIp>> ElasticIps { get; private set; } = null!;

        [Output("hostnameTheme")]
        public Output<string?> HostnameTheme { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("rdsDbInstances")]
        public Output<ImmutableArray<Outputs.StackRdsDbInstance>> RdsDbInstances { get; private set; } = null!;

        [Output("serviceRoleArn")]
        public Output<string> ServiceRoleArn { get; private set; } = null!;

        [Output("sourceStackId")]
        public Output<string?> SourceStackId { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.StackTag>> Tags { get; private set; } = null!;

        [Output("useCustomCookbooks")]
        public Output<bool?> UseCustomCookbooks { get; private set; } = null!;

        [Output("useOpsworksSecurityGroups")]
        public Output<bool?> UseOpsworksSecurityGroups { get; private set; } = null!;

        [Output("vpcId")]
        public Output<string?> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Stack resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Stack(string name, StackArgs args, CustomResourceOptions? options = null)
            : base("aws-native:opsworks:Stack", name, args ?? new StackArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Stack(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:opsworks:Stack", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Stack resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Stack Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Stack(name, id, options);
        }
    }

    public sealed class StackArgs : global::Pulumi.ResourceArgs
    {
        [Input("agentVersion")]
        public Input<string>? AgentVersion { get; set; }

        [Input("attributes")]
        public Input<object>? Attributes { get; set; }

        [Input("chefConfiguration")]
        public Input<Inputs.StackChefConfigurationArgs>? ChefConfiguration { get; set; }

        [Input("cloneAppIds")]
        private InputList<string>? _cloneAppIds;
        public InputList<string> CloneAppIds
        {
            get => _cloneAppIds ?? (_cloneAppIds = new InputList<string>());
            set => _cloneAppIds = value;
        }

        [Input("clonePermissions")]
        public Input<bool>? ClonePermissions { get; set; }

        [Input("configurationManager")]
        public Input<Inputs.StackConfigurationManagerArgs>? ConfigurationManager { get; set; }

        [Input("customCookbooksSource")]
        public Input<Inputs.StackSourceArgs>? CustomCookbooksSource { get; set; }

        [Input("customJson")]
        public Input<object>? CustomJson { get; set; }

        [Input("defaultAvailabilityZone")]
        public Input<string>? DefaultAvailabilityZone { get; set; }

        [Input("defaultInstanceProfileArn", required: true)]
        public Input<string> DefaultInstanceProfileArn { get; set; } = null!;

        [Input("defaultOs")]
        public Input<string>? DefaultOs { get; set; }

        [Input("defaultRootDeviceType")]
        public Input<string>? DefaultRootDeviceType { get; set; }

        [Input("defaultSshKeyName")]
        public Input<string>? DefaultSshKeyName { get; set; }

        [Input("defaultSubnetId")]
        public Input<string>? DefaultSubnetId { get; set; }

        [Input("ecsClusterArn")]
        public Input<string>? EcsClusterArn { get; set; }

        [Input("elasticIps")]
        private InputList<Inputs.StackElasticIpArgs>? _elasticIps;
        public InputList<Inputs.StackElasticIpArgs> ElasticIps
        {
            get => _elasticIps ?? (_elasticIps = new InputList<Inputs.StackElasticIpArgs>());
            set => _elasticIps = value;
        }

        [Input("hostnameTheme")]
        public Input<string>? HostnameTheme { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("rdsDbInstances")]
        private InputList<Inputs.StackRdsDbInstanceArgs>? _rdsDbInstances;
        public InputList<Inputs.StackRdsDbInstanceArgs> RdsDbInstances
        {
            get => _rdsDbInstances ?? (_rdsDbInstances = new InputList<Inputs.StackRdsDbInstanceArgs>());
            set => _rdsDbInstances = value;
        }

        [Input("serviceRoleArn", required: true)]
        public Input<string> ServiceRoleArn { get; set; } = null!;

        [Input("sourceStackId")]
        public Input<string>? SourceStackId { get; set; }

        [Input("tags")]
        private InputList<Inputs.StackTagArgs>? _tags;
        public InputList<Inputs.StackTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.StackTagArgs>());
            set => _tags = value;
        }

        [Input("useCustomCookbooks")]
        public Input<bool>? UseCustomCookbooks { get; set; }

        [Input("useOpsworksSecurityGroups")]
        public Input<bool>? UseOpsworksSecurityGroups { get; set; }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public StackArgs()
        {
        }
        public static new StackArgs Empty => new StackArgs();
    }
}