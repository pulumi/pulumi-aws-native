// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FinSpace.Inputs
{

    /// <summary>
    /// Additional parameters to identify Federation mode
    /// </summary>
    public sealed class EnvironmentFederationParametersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SAML metadata URL to link with the Environment
        /// </summary>
        [Input("applicationCallBackURL")]
        public Input<string>? ApplicationCallBackURL { get; set; }

        [Input("attributeMap")]
        private InputList<Inputs.EnvironmentFederationParametersAttributeMapItemPropertiesArgs>? _attributeMap;

        /// <summary>
        /// Attribute map for SAML configuration
        /// </summary>
        public InputList<Inputs.EnvironmentFederationParametersAttributeMapItemPropertiesArgs> AttributeMap
        {
            get => _attributeMap ?? (_attributeMap = new InputList<Inputs.EnvironmentFederationParametersAttributeMapItemPropertiesArgs>());
            set => _attributeMap = value;
        }

        /// <summary>
        /// Federation provider name to link with the Environment
        /// </summary>
        [Input("federationProviderName")]
        public Input<string>? FederationProviderName { get; set; }

        /// <summary>
        /// SAML metadata URL to link with the Environment
        /// </summary>
        [Input("federationURN")]
        public Input<string>? FederationURN { get; set; }

        /// <summary>
        /// SAML metadata document to link the federation provider to the Environment
        /// </summary>
        [Input("samlMetadataDocument")]
        public Input<string>? SamlMetadataDocument { get; set; }

        /// <summary>
        /// SAML metadata URL to link with the Environment
        /// </summary>
        [Input("samlMetadataURL")]
        public Input<string>? SamlMetadataURL { get; set; }

        public EnvironmentFederationParametersArgs()
        {
        }
        public static new EnvironmentFederationParametersArgs Empty => new EnvironmentFederationParametersArgs();
    }
}