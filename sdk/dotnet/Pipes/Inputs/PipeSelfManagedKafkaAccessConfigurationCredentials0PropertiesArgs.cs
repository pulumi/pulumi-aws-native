// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Inputs
{

    public sealed class PipeSelfManagedKafkaAccessConfigurationCredentials0PropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional SecretManager ARN which stores the database credentials
        /// </summary>
        [Input("basicAuth", required: true)]
        public Input<string> BasicAuth { get; set; } = null!;

        public PipeSelfManagedKafkaAccessConfigurationCredentials0PropertiesArgs()
        {
        }
        public static new PipeSelfManagedKafkaAccessConfigurationCredentials0PropertiesArgs Empty => new PipeSelfManagedKafkaAccessConfigurationCredentials0PropertiesArgs();
    }
}
