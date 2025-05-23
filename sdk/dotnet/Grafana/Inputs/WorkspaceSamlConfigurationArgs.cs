// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Grafana.Inputs
{

    /// <summary>
    /// SAML configuration data associated with an AMG workspace.
    /// </summary>
    public sealed class WorkspaceSamlConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedOrganizations")]
        private InputList<string>? _allowedOrganizations;

        /// <summary>
        /// List of SAML organizations allowed to access Grafana.
        /// </summary>
        public InputList<string> AllowedOrganizations
        {
            get => _allowedOrganizations ?? (_allowedOrganizations = new InputList<string>());
            set => _allowedOrganizations = value;
        }

        /// <summary>
        /// A structure that defines which attributes in the SAML assertion are to be used to define information about the users authenticated by that IdP to use the workspace.
        /// </summary>
        [Input("assertionAttributes")]
        public Input<Inputs.WorkspaceAssertionAttributesArgs>? AssertionAttributes { get; set; }

        /// <summary>
        /// A structure containing the identity provider (IdP) metadata used to integrate the identity provider with this workspace.
        /// </summary>
        [Input("idpMetadata", required: true)]
        public Input<Inputs.WorkspaceIdpMetadataArgs> IdpMetadata { get; set; } = null!;

        /// <summary>
        /// The maximum lifetime an authenticated user can be logged in (in minutes) before being required to re-authenticate.
        /// </summary>
        [Input("loginValidityDuration")]
        public Input<double>? LoginValidityDuration { get; set; }

        /// <summary>
        /// A structure containing arrays that map group names in the SAML assertion to the Grafana `Admin` and `Editor` roles in the workspace.
        /// </summary>
        [Input("roleValues")]
        public Input<Inputs.WorkspaceRoleValuesArgs>? RoleValues { get; set; }

        public WorkspaceSamlConfigurationArgs()
        {
        }
        public static new WorkspaceSamlConfigurationArgs Empty => new WorkspaceSamlConfigurationArgs();
    }
}
