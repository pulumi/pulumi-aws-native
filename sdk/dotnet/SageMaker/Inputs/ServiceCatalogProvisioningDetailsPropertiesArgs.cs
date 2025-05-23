// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// Input ServiceCatalog Provisioning Details
    /// </summary>
    public sealed class ServiceCatalogProvisioningDetailsPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path identifier of the product. This value is optional if the product has a default path, and required if the product has more than one path.
        /// </summary>
        [Input("pathId")]
        public Input<string>? PathId { get; set; }

        /// <summary>
        /// The ID of the product to provision.
        /// </summary>
        [Input("productId", required: true)]
        public Input<string> ProductId { get; set; } = null!;

        /// <summary>
        /// The ID of the provisioning artifact.
        /// </summary>
        [Input("provisioningArtifactId")]
        public Input<string>? ProvisioningArtifactId { get; set; }

        [Input("provisioningParameters")]
        private InputList<Inputs.ProjectProvisioningParameterArgs>? _provisioningParameters;

        /// <summary>
        /// Parameters specified by the administrator that are required for provisioning the product.
        /// </summary>
        public InputList<Inputs.ProjectProvisioningParameterArgs> ProvisioningParameters
        {
            get => _provisioningParameters ?? (_provisioningParameters = new InputList<Inputs.ProjectProvisioningParameterArgs>());
            set => _provisioningParameters = value;
        }

        public ServiceCatalogProvisioningDetailsPropertiesArgs()
        {
        }
        public static new ServiceCatalogProvisioningDetailsPropertiesArgs Empty => new ServiceCatalogProvisioningDetailsPropertiesArgs();
    }
}
