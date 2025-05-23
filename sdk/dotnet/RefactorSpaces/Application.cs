// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RefactorSpaces
{
    /// <summary>
    /// Definition of AWS::RefactorSpaces::Application Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:refactorspaces:Application")]
    public partial class Application : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The resource ID of the API Gateway for the proxy.
        /// </summary>
        [Output("apiGatewayId")]
        public Output<string> ApiGatewayId { get; private set; } = null!;

        /// <summary>
        /// The endpoint URL of the Amazon API Gateway proxy.
        /// </summary>
        [Output("apiGatewayProxy")]
        public Output<Outputs.ApplicationApiGatewayProxyInput?> ApiGatewayProxy { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the application.
        /// </summary>
        [Output("applicationIdentifier")]
        public Output<string> ApplicationIdentifier { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the application.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the environment.
        /// </summary>
        [Output("environmentIdentifier")]
        public Output<string> EnvironmentIdentifier { get; private set; } = null!;

        /// <summary>
        /// The name of the application.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Network Load Balancer .
        /// </summary>
        [Output("nlbArn")]
        public Output<string> NlbArn { get; private set; } = null!;

        /// <summary>
        /// The name of the Network Load Balancer configured by the API Gateway proxy.
        /// </summary>
        [Output("nlbName")]
        public Output<string> NlbName { get; private set; } = null!;

        /// <summary>
        /// The proxy type of the proxy created within the application.
        /// </summary>
        [Output("proxyType")]
        public Output<Pulumi.AwsNative.RefactorSpaces.ApplicationProxyType> ProxyType { get; private set; } = null!;

        /// <summary>
        /// The endpoint URL of the Amazon API Gateway proxy.
        /// </summary>
        [Output("proxyUrl")]
        public Output<string> ProxyUrl { get; private set; } = null!;

        /// <summary>
        /// The name of the API Gateway stage. The name defaults to `prod` .
        /// </summary>
        [Output("stageName")]
        public Output<string> StageName { get; private set; } = null!;

        /// <summary>
        /// Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the virtual private cloud (VPC).
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The `VpcLink` ID of the API Gateway proxy.
        /// </summary>
        [Output("vpcLinkId")]
        public Output<string> VpcLinkId { get; private set; } = null!;


        /// <summary>
        /// Create a Application resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Application(string name, ApplicationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:refactorspaces:Application", name, args ?? new ApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Application(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:refactorspaces:Application", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "apiGatewayProxy",
                    "environmentIdentifier",
                    "name",
                    "proxyType",
                    "vpcId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Application resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Application Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Application(name, id, options);
        }
    }

    public sealed class ApplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The endpoint URL of the Amazon API Gateway proxy.
        /// </summary>
        [Input("apiGatewayProxy")]
        public Input<Inputs.ApplicationApiGatewayProxyInputArgs>? ApiGatewayProxy { get; set; }

        /// <summary>
        /// The unique identifier of the environment.
        /// </summary>
        [Input("environmentIdentifier", required: true)]
        public Input<string> EnvironmentIdentifier { get; set; } = null!;

        /// <summary>
        /// The name of the application.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The proxy type of the proxy created within the application.
        /// </summary>
        [Input("proxyType", required: true)]
        public Input<Pulumi.AwsNative.RefactorSpaces.ApplicationProxyType> ProxyType { get; set; } = null!;

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the virtual private cloud (VPC).
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public ApplicationArgs()
        {
        }
        public static new ApplicationArgs Empty => new ApplicationArgs();
    }
}
