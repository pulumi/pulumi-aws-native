// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RDS.Inputs
{

    /// <summary>
    /// The OptionConfiguration property type specifies an individual option, and its settings, within an AWS::RDS::OptionGroup resource.
    /// </summary>
    public sealed class OptionGroupOptionConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("dBSecurityGroupMemberships")]
        private InputList<string>? _dBSecurityGroupMemberships;

        /// <summary>
        /// A list of DBSecurityGroupMembership name strings used for this option.
        /// </summary>
        public InputList<string> DBSecurityGroupMemberships
        {
            get => _dBSecurityGroupMemberships ?? (_dBSecurityGroupMemberships = new InputList<string>());
            set => _dBSecurityGroupMemberships = value;
        }

        /// <summary>
        /// The configuration of options to include in a group.
        /// </summary>
        [Input("optionName", required: true)]
        public Input<string> OptionName { get; set; } = null!;

        [Input("optionSettings")]
        private InputList<Inputs.OptionGroupOptionSettingArgs>? _optionSettings;

        /// <summary>
        /// The option settings to include in an option group.
        /// </summary>
        public InputList<Inputs.OptionGroupOptionSettingArgs> OptionSettings
        {
            get => _optionSettings ?? (_optionSettings = new InputList<Inputs.OptionGroupOptionSettingArgs>());
            set => _optionSettings = value;
        }

        /// <summary>
        /// The version for the option.
        /// </summary>
        [Input("optionVersion")]
        public Input<string>? OptionVersion { get; set; }

        /// <summary>
        /// The optional port for the option.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("vpcSecurityGroupMemberships")]
        private InputList<string>? _vpcSecurityGroupMemberships;

        /// <summary>
        /// A list of VpcSecurityGroupMembership name strings used for this option.
        /// </summary>
        public InputList<string> VpcSecurityGroupMemberships
        {
            get => _vpcSecurityGroupMemberships ?? (_vpcSecurityGroupMemberships = new InputList<string>());
            set => _vpcSecurityGroupMemberships = value;
        }

        public OptionGroupOptionConfigurationArgs()
        {
        }
        public static new OptionGroupOptionConfigurationArgs Empty => new OptionGroupOptionConfigurationArgs();
    }
}