// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Inputs
{

    /// <summary>
    /// Glue Connection Input
    /// </summary>
    public sealed class ConnectionGlueConnectionInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("athenaProperties")]
        private InputMap<string>? _athenaProperties;
        public InputMap<string> AthenaProperties
        {
            get => _athenaProperties ?? (_athenaProperties = new InputMap<string>());
            set => _athenaProperties = value;
        }

        [Input("authenticationConfiguration")]
        public Input<Inputs.ConnectionAuthenticationConfigurationInputArgs>? AuthenticationConfiguration { get; set; }

        [Input("connectionProperties")]
        private InputMap<string>? _connectionProperties;
        public InputMap<string> ConnectionProperties
        {
            get => _connectionProperties ?? (_connectionProperties = new InputMap<string>());
            set => _connectionProperties = value;
        }

        [Input("connectionType")]
        public Input<string>? ConnectionType { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("matchCriteria")]
        public Input<string>? MatchCriteria { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("physicalConnectionRequirements")]
        public Input<Inputs.ConnectionPhysicalConnectionRequirementsArgs>? PhysicalConnectionRequirements { get; set; }

        [Input("pythonProperties")]
        private InputMap<string>? _pythonProperties;
        public InputMap<string> PythonProperties
        {
            get => _pythonProperties ?? (_pythonProperties = new InputMap<string>());
            set => _pythonProperties = value;
        }

        [Input("sparkProperties")]
        private InputMap<string>? _sparkProperties;
        public InputMap<string> SparkProperties
        {
            get => _sparkProperties ?? (_sparkProperties = new InputMap<string>());
            set => _sparkProperties = value;
        }

        [Input("validateCredentials")]
        public Input<bool>? ValidateCredentials { get; set; }

        [Input("validateForComputeEnvironments")]
        private InputList<string>? _validateForComputeEnvironments;
        public InputList<string> ValidateForComputeEnvironments
        {
            get => _validateForComputeEnvironments ?? (_validateForComputeEnvironments = new InputList<string>());
            set => _validateForComputeEnvironments = value;
        }

        public ConnectionGlueConnectionInputArgs()
        {
        }
        public static new ConnectionGlueConnectionInputArgs Empty => new ConnectionGlueConnectionInputArgs();
    }
}
