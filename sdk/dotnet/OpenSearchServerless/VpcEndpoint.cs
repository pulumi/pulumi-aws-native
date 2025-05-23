// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchServerless
{
    /// <summary>
    /// Amazon OpenSearchServerless vpc endpoint resource
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
    ///     var testAOSSVpcEndpoint = new AwsNative.OpenSearchServerless.VpcEndpoint("testAOSSVpcEndpoint", new()
    ///     {
    ///         Name = "test-vpcendpoint",
    ///         VpcId = "vpc-0d728b8430292b3f4",
    ///         SubnetIds = new[]
    ///         {
    ///             "subnet-0e855f5722a9598ee",
    ///         },
    ///         SecurityGroupIds = new[]
    ///         {
    ///             "sg-03843b03f369eb245",
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
    ///     var testAOSSVpcEndpoint = new AwsNative.OpenSearchServerless.VpcEndpoint("testAOSSVpcEndpoint", new()
    ///     {
    ///         Name = "test-vpcendpoint",
    ///         VpcId = "vpc-0d728b8430292b3f4",
    ///         SubnetIds = new[]
    ///         {
    ///             "subnet-0e855f5722a9598ee",
    ///         },
    ///         SecurityGroupIds = new[]
    ///         {
    ///             "sg-03843b03f369eb245",
    ///         },
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:opensearchserverless:VpcEndpoint")]
    public partial class VpcEndpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The identifier of the VPC Endpoint
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// The name of the VPC Endpoint
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of one or more security groups to associate with the endpoint network interface
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// The ID of one or more subnets in which to create an endpoint network interface
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC in which the endpoint will be used.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcEndpoint(string name, VpcEndpointArgs args, CustomResourceOptions? options = null)
            : base("aws-native:opensearchserverless:VpcEndpoint", name, args ?? new VpcEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcEndpoint(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:opensearchserverless:VpcEndpoint", name, null, MakeResourceOptions(options, id))
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
                    "vpcId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VpcEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcEndpoint Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VpcEndpoint(name, id, options);
        }
    }

    public sealed class VpcEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the VPC Endpoint
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The ID of one or more security groups to associate with the endpoint network interface
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The ID of one or more subnets in which to create an endpoint network interface
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        /// <summary>
        /// The ID of the VPC in which the endpoint will be used.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public VpcEndpointArgs()
        {
        }
        public static new VpcEndpointArgs Empty => new VpcEndpointArgs();
    }
}
