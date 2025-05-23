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
    /// Allocates an internet gateway for use with a VPC. After creating the Internet gateway, you then attach it to a VPC.
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
    ///     var myInternetGateway = new AwsNative.Ec2.InternetGateway("myInternetGateway", new()
    ///     {
    ///         Tags = new[]
    ///         {
    ///             new AwsNative.Inputs.TagArgs
    ///             {
    ///                 Key = "stack",
    ///                 Value = "production",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:InternetGateway")]
    public partial class InternetGateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the internet gateway.
        /// </summary>
        [Output("internetGatewayId")]
        public Output<string> InternetGatewayId { get; private set; } = null!;

        /// <summary>
        /// Any tags to assign to the internet gateway.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a InternetGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InternetGateway(string name, InternetGatewayArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ec2:InternetGateway", name, args ?? new InternetGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InternetGateway(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:InternetGateway", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing InternetGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InternetGateway Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new InternetGateway(name, id, options);
        }
    }

    public sealed class InternetGatewayArgs : global::Pulumi.ResourceArgs
    {
        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// Any tags to assign to the internet gateway.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public InternetGatewayArgs()
        {
        }
        public static new InternetGatewayArgs Empty => new InternetGatewayArgs();
    }
}
