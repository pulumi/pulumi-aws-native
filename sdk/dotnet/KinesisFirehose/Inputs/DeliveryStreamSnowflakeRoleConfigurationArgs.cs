// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Inputs
{

    public sealed class DeliveryStreamSnowflakeRoleConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable Snowflake role
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The Snowflake role you wish to configure
        /// </summary>
        [Input("snowflakeRole")]
        public Input<string>? SnowflakeRole { get; set; }

        public DeliveryStreamSnowflakeRoleConfigurationArgs()
        {
        }
        public static new DeliveryStreamSnowflakeRoleConfigurationArgs Empty => new DeliveryStreamSnowflakeRoleConfigurationArgs();
    }
}
