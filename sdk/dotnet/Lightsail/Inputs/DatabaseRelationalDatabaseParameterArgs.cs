// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lightsail.Inputs
{

    /// <summary>
    /// Describes the parameters of the database.
    /// </summary>
    public sealed class DatabaseRelationalDatabaseParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the valid range of values for the parameter.
        /// </summary>
        [Input("allowedValues")]
        public Input<string>? AllowedValues { get; set; }

        /// <summary>
        /// Indicates when parameter updates are applied. Can be immediate or pending-reboot.
        /// </summary>
        [Input("applyMethod")]
        public Input<string>? ApplyMethod { get; set; }

        /// <summary>
        /// Specifies the engine-specific parameter type.
        /// </summary>
        [Input("applyType")]
        public Input<string>? ApplyType { get; set; }

        /// <summary>
        /// Specifies the valid data type for the parameter.
        /// </summary>
        [Input("dataType")]
        public Input<string>? DataType { get; set; }

        /// <summary>
        /// Provides a description of the parameter.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A Boolean value indicating whether the parameter can be modified.
        /// </summary>
        [Input("isModifiable")]
        public Input<bool>? IsModifiable { get; set; }

        /// <summary>
        /// Specifies the name of the parameter.
        /// </summary>
        [Input("parameterName")]
        public Input<string>? ParameterName { get; set; }

        /// <summary>
        /// Specifies the value of the parameter.
        /// </summary>
        [Input("parameterValue")]
        public Input<string>? ParameterValue { get; set; }

        public DatabaseRelationalDatabaseParameterArgs()
        {
        }
        public static new DatabaseRelationalDatabaseParameterArgs Empty => new DatabaseRelationalDatabaseParameterArgs();
    }
}