// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Inputs
{

    /// <summary>
    /// Specifies the configuration of the data source. It can be set to either glueRunConfiguration or redshiftRunConfiguration.
    /// </summary>
    public sealed class DataSourceConfigurationInput0PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("glueRunConfiguration")]
        public Input<Inputs.DataSourceGlueRunConfigurationInputArgs>? GlueRunConfiguration { get; set; }

        public DataSourceConfigurationInput0PropertiesArgs()
        {
        }
        public static new DataSourceConfigurationInput0PropertiesArgs Empty => new DataSourceConfigurationInput0PropertiesArgs();
    }
}