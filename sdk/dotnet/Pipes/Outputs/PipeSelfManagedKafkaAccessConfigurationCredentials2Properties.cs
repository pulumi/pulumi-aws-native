// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Outputs
{

    [OutputType]
    public sealed class PipeSelfManagedKafkaAccessConfigurationCredentials2Properties
    {
        /// <summary>
        /// Optional SecretManager ARN which stores the database credentials
        /// </summary>
        public readonly string SaslScram256Auth;

        [OutputConstructor]
        private PipeSelfManagedKafkaAccessConfigurationCredentials2Properties(string saslScram256Auth)
        {
            SaslScram256Auth = saslScram256Auth;
        }
    }
}
