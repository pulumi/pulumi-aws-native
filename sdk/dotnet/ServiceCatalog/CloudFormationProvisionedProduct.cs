// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceCatalog
{
    /// <summary>
    /// Resource Schema for AWS::ServiceCatalog::CloudFormationProvisionedProduct
    /// </summary>
    [AwsNativeResourceType("aws-native:servicecatalog:CloudFormationProvisionedProduct")]
    public partial class CloudFormationProvisionedProduct : global::Pulumi.CustomResource
    {
        [Output("acceptLanguage")]
        public Output<Pulumi.AwsNative.ServiceCatalog.CloudFormationProvisionedProductAcceptLanguage?> AcceptLanguage { get; private set; } = null!;

        [Output("cloudformationStackArn")]
        public Output<string> CloudformationStackArn { get; private set; } = null!;

        [Output("notificationArns")]
        public Output<ImmutableArray<string>> NotificationArns { get; private set; } = null!;

        /// <summary>
        /// List of key-value pair outputs.
        /// </summary>
        [Output("outputs")]
        public Output<object> Outputs { get; private set; } = null!;

        [Output("pathId")]
        public Output<string?> PathId { get; private set; } = null!;

        [Output("pathName")]
        public Output<string?> PathName { get; private set; } = null!;

        [Output("productId")]
        public Output<string?> ProductId { get; private set; } = null!;

        [Output("productName")]
        public Output<string?> ProductName { get; private set; } = null!;

        [Output("provisionedProductId")]
        public Output<string> ProvisionedProductId { get; private set; } = null!;

        [Output("provisionedProductName")]
        public Output<string?> ProvisionedProductName { get; private set; } = null!;

        [Output("provisioningArtifactId")]
        public Output<string?> ProvisioningArtifactId { get; private set; } = null!;

        [Output("provisioningArtifactName")]
        public Output<string?> ProvisioningArtifactName { get; private set; } = null!;

        [Output("provisioningParameters")]
        public Output<ImmutableArray<Outputs.CloudFormationProvisionedProductProvisioningParameter>> ProvisioningParameters { get; private set; } = null!;

        [Output("provisioningPreferences")]
        public Output<Outputs.CloudFormationProvisionedProductProvisioningPreferences?> ProvisioningPreferences { get; private set; } = null!;

        [Output("recordId")]
        public Output<string> RecordId { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.CloudFormationProvisionedProductTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a CloudFormationProvisionedProduct resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudFormationProvisionedProduct(string name, CloudFormationProvisionedProductArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:servicecatalog:CloudFormationProvisionedProduct", name, args ?? new CloudFormationProvisionedProductArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudFormationProvisionedProduct(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:servicecatalog:CloudFormationProvisionedProduct", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing CloudFormationProvisionedProduct resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudFormationProvisionedProduct Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CloudFormationProvisionedProduct(name, id, options);
        }
    }

    public sealed class CloudFormationProvisionedProductArgs : global::Pulumi.ResourceArgs
    {
        [Input("acceptLanguage")]
        public Input<Pulumi.AwsNative.ServiceCatalog.CloudFormationProvisionedProductAcceptLanguage>? AcceptLanguage { get; set; }

        [Input("notificationArns")]
        private InputList<string>? _notificationArns;
        public InputList<string> NotificationArns
        {
            get => _notificationArns ?? (_notificationArns = new InputList<string>());
            set => _notificationArns = value;
        }

        [Input("pathId")]
        public Input<string>? PathId { get; set; }

        [Input("pathName")]
        public Input<string>? PathName { get; set; }

        [Input("productId")]
        public Input<string>? ProductId { get; set; }

        [Input("productName")]
        public Input<string>? ProductName { get; set; }

        [Input("provisionedProductName")]
        public Input<string>? ProvisionedProductName { get; set; }

        [Input("provisioningArtifactId")]
        public Input<string>? ProvisioningArtifactId { get; set; }

        [Input("provisioningArtifactName")]
        public Input<string>? ProvisioningArtifactName { get; set; }

        [Input("provisioningParameters")]
        private InputList<Inputs.CloudFormationProvisionedProductProvisioningParameterArgs>? _provisioningParameters;
        public InputList<Inputs.CloudFormationProvisionedProductProvisioningParameterArgs> ProvisioningParameters
        {
            get => _provisioningParameters ?? (_provisioningParameters = new InputList<Inputs.CloudFormationProvisionedProductProvisioningParameterArgs>());
            set => _provisioningParameters = value;
        }

        [Input("provisioningPreferences")]
        public Input<Inputs.CloudFormationProvisionedProductProvisioningPreferencesArgs>? ProvisioningPreferences { get; set; }

        [Input("tags")]
        private InputList<Inputs.CloudFormationProvisionedProductTagArgs>? _tags;
        public InputList<Inputs.CloudFormationProvisionedProductTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.CloudFormationProvisionedProductTagArgs>());
            set => _tags = value;
        }

        public CloudFormationProvisionedProductArgs()
        {
        }
        public static new CloudFormationProvisionedProductArgs Empty => new CloudFormationProvisionedProductArgs();
    }
}