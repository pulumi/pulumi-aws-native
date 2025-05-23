// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    /// <summary>
    /// Resource Schema of AWS::EC2::IPAMPool Type
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:IpamPool")]
    public partial class IpamPool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The address family of the address space in this pool. Either IPv4 or IPv6.
        /// </summary>
        [Output("addressFamily")]
        public Output<string> AddressFamily { get; private set; } = null!;

        /// <summary>
        /// The default netmask length for allocations made from this pool. This value is used when the netmask length of an allocation isn't specified.
        /// </summary>
        [Output("allocationDefaultNetmaskLength")]
        public Output<int?> AllocationDefaultNetmaskLength { get; private set; } = null!;

        /// <summary>
        /// The maximum allowed netmask length for allocations made from this pool.
        /// </summary>
        [Output("allocationMaxNetmaskLength")]
        public Output<int?> AllocationMaxNetmaskLength { get; private set; } = null!;

        /// <summary>
        /// The minimum allowed netmask length for allocations made from this pool.
        /// </summary>
        [Output("allocationMinNetmaskLength")]
        public Output<int?> AllocationMinNetmaskLength { get; private set; } = null!;

        /// <summary>
        /// When specified, an allocation will not be allowed unless a resource has a matching set of tags.
        /// </summary>
        [Output("allocationResourceTags")]
        public Output<ImmutableArray<Outputs.IpamPoolTag>> AllocationResourceTags { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IPAM Pool.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Determines what to do if IPAM discovers resources that haven't been assigned an allocation. If set to true, an allocation will be made automatically.
        /// </summary>
        [Output("autoImport")]
        public Output<bool?> AutoImport { get; private set; } = null!;

        /// <summary>
        /// Limits which service in Amazon Web Services that the pool can be used in.
        /// </summary>
        [Output("awsService")]
        public Output<Pulumi.AwsNative.Ec2.IpamPoolAwsService?> AwsService { get; private set; } = null!;

        /// <summary>
        /// The description of the IPAM pool.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IPAM this pool is a part of.
        /// </summary>
        [Output("ipamArn")]
        public Output<string> IpamArn { get; private set; } = null!;

        /// <summary>
        /// Id of the IPAM Pool.
        /// </summary>
        [Output("ipamPoolId")]
        public Output<string> IpamPoolId { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the scope this pool is a part of.
        /// </summary>
        [Output("ipamScopeArn")]
        public Output<string> IpamScopeArn { get; private set; } = null!;

        /// <summary>
        /// The Id of the scope this pool is a part of.
        /// </summary>
        [Output("ipamScopeId")]
        public Output<string> IpamScopeId { get; private set; } = null!;

        /// <summary>
        /// Determines whether this scope contains publicly routable space or space for a private network
        /// </summary>
        [Output("ipamScopeType")]
        public Output<Pulumi.AwsNative.Ec2.IpamPoolIpamScopeType> IpamScopeType { get; private set; } = null!;

        /// <summary>
        /// The region of this pool. If not set, this will default to "None" which will disable non-custom allocations. If the locale has been specified for the source pool, this value must match.
        /// </summary>
        [Output("locale")]
        public Output<string?> Locale { get; private set; } = null!;

        /// <summary>
        /// The depth of this pool in the source pool hierarchy.
        /// </summary>
        [Output("poolDepth")]
        public Output<int> PoolDepth { get; private set; } = null!;

        /// <summary>
        /// A list of cidrs representing the address space available for allocation in this pool.
        /// </summary>
        [Output("provisionedCidrs")]
        public Output<ImmutableArray<Outputs.IpamPoolProvisionedCidr>> ProvisionedCidrs { get; private set; } = null!;

        /// <summary>
        /// The IP address source for pools in the public scope. Only used for provisioning IP address CIDRs to pools in the public scope. Default is `byoip`.
        /// </summary>
        [Output("publicIpSource")]
        public Output<Pulumi.AwsNative.Ec2.IpamPoolPublicIpSource?> PublicIpSource { get; private set; } = null!;

        /// <summary>
        /// Determines whether or not address space from this pool is publicly advertised. Must be set if and only if the pool is IPv6.
        /// </summary>
        [Output("publiclyAdvertisable")]
        public Output<bool?> PubliclyAdvertisable { get; private set; } = null!;

        /// <summary>
        /// The Id of this pool's source. If set, all space provisioned in this pool must be free space provisioned in the parent pool.
        /// </summary>
        [Output("sourceIpamPoolId")]
        public Output<string?> SourceIpamPoolId { get; private set; } = null!;

        /// <summary>
        /// The resource used to provision CIDRs to a resource planning pool.
        /// </summary>
        [Output("sourceResource")]
        public Output<Outputs.IpamPoolSourceResource?> SourceResource { get; private set; } = null!;

        /// <summary>
        /// The state of this pool. This can be one of the following values: "create-in-progress", "create-complete", "modify-in-progress", "modify-complete", "delete-in-progress", or "delete-complete"
        /// </summary>
        [Output("state")]
        public Output<Pulumi.AwsNative.Ec2.IpamPoolState> State { get; private set; } = null!;

        /// <summary>
        /// An explanation of how the pool arrived at it current state.
        /// </summary>
        [Output("stateMessage")]
        public Output<string> StateMessage { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a IpamPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpamPool(string name, IpamPoolArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:IpamPool", name, args ?? new IpamPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpamPool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:IpamPool", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "addressFamily",
                    "awsService",
                    "ipamScopeId",
                    "locale",
                    "publicIpSource",
                    "publiclyAdvertisable",
                    "sourceIpamPoolId",
                    "sourceResource",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IpamPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpamPool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IpamPool(name, id, options);
        }
    }

    public sealed class IpamPoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The address family of the address space in this pool. Either IPv4 or IPv6.
        /// </summary>
        [Input("addressFamily", required: true)]
        public Input<string> AddressFamily { get; set; } = null!;

        /// <summary>
        /// The default netmask length for allocations made from this pool. This value is used when the netmask length of an allocation isn't specified.
        /// </summary>
        [Input("allocationDefaultNetmaskLength")]
        public Input<int>? AllocationDefaultNetmaskLength { get; set; }

        /// <summary>
        /// The maximum allowed netmask length for allocations made from this pool.
        /// </summary>
        [Input("allocationMaxNetmaskLength")]
        public Input<int>? AllocationMaxNetmaskLength { get; set; }

        /// <summary>
        /// The minimum allowed netmask length for allocations made from this pool.
        /// </summary>
        [Input("allocationMinNetmaskLength")]
        public Input<int>? AllocationMinNetmaskLength { get; set; }

        [Input("allocationResourceTags")]
        private InputList<Inputs.IpamPoolTagArgs>? _allocationResourceTags;

        /// <summary>
        /// When specified, an allocation will not be allowed unless a resource has a matching set of tags.
        /// </summary>
        public InputList<Inputs.IpamPoolTagArgs> AllocationResourceTags
        {
            get => _allocationResourceTags ?? (_allocationResourceTags = new InputList<Inputs.IpamPoolTagArgs>());
            set => _allocationResourceTags = value;
        }

        /// <summary>
        /// Determines what to do if IPAM discovers resources that haven't been assigned an allocation. If set to true, an allocation will be made automatically.
        /// </summary>
        [Input("autoImport")]
        public Input<bool>? AutoImport { get; set; }

        /// <summary>
        /// Limits which service in Amazon Web Services that the pool can be used in.
        /// </summary>
        [Input("awsService")]
        public Input<Pulumi.AwsNative.Ec2.IpamPoolAwsService>? AwsService { get; set; }

        /// <summary>
        /// The description of the IPAM pool.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Id of the scope this pool is a part of.
        /// </summary>
        [Input("ipamScopeId", required: true)]
        public Input<string> IpamScopeId { get; set; } = null!;

        /// <summary>
        /// The region of this pool. If not set, this will default to "None" which will disable non-custom allocations. If the locale has been specified for the source pool, this value must match.
        /// </summary>
        [Input("locale")]
        public Input<string>? Locale { get; set; }

        [Input("provisionedCidrs")]
        private InputList<Inputs.IpamPoolProvisionedCidrArgs>? _provisionedCidrs;

        /// <summary>
        /// A list of cidrs representing the address space available for allocation in this pool.
        /// </summary>
        public InputList<Inputs.IpamPoolProvisionedCidrArgs> ProvisionedCidrs
        {
            get => _provisionedCidrs ?? (_provisionedCidrs = new InputList<Inputs.IpamPoolProvisionedCidrArgs>());
            set => _provisionedCidrs = value;
        }

        /// <summary>
        /// The IP address source for pools in the public scope. Only used for provisioning IP address CIDRs to pools in the public scope. Default is `byoip`.
        /// </summary>
        [Input("publicIpSource")]
        public Input<Pulumi.AwsNative.Ec2.IpamPoolPublicIpSource>? PublicIpSource { get; set; }

        /// <summary>
        /// Determines whether or not address space from this pool is publicly advertised. Must be set if and only if the pool is IPv6.
        /// </summary>
        [Input("publiclyAdvertisable")]
        public Input<bool>? PubliclyAdvertisable { get; set; }

        /// <summary>
        /// The Id of this pool's source. If set, all space provisioned in this pool must be free space provisioned in the parent pool.
        /// </summary>
        [Input("sourceIpamPoolId")]
        public Input<string>? SourceIpamPoolId { get; set; }

        /// <summary>
        /// The resource used to provision CIDRs to a resource planning pool.
        /// </summary>
        [Input("sourceResource")]
        public Input<Inputs.IpamPoolSourceResourceArgs>? SourceResource { get; set; }

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

        public IpamPoolArgs()
        {
        }
        public static new IpamPoolArgs Empty => new IpamPoolArgs();
    }
}
