// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    /// <summary>
    /// Configurations for Redshift query engine provisioned auth setup
    /// </summary>
    public sealed class KnowledgeBaseRedshiftProvisionedAuthConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Redshift database user
        /// </summary>
        [Input("databaseUser")]
        public Input<string>? DatabaseUser { get; set; }

        /// <summary>
        /// The type of authentication to use.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.Bedrock.KnowledgeBaseRedshiftProvisionedAuthType> Type { get; set; } = null!;

        /// <summary>
        /// The ARN of an Secrets Manager secret for authentication.
        /// </summary>
        [Input("usernamePasswordSecretArn")]
        public Input<string>? UsernamePasswordSecretArn { get; set; }

        public KnowledgeBaseRedshiftProvisionedAuthConfigurationArgs()
        {
        }
        public static new KnowledgeBaseRedshiftProvisionedAuthConfigurationArgs Empty => new KnowledgeBaseRedshiftProvisionedAuthConfigurationArgs();
    }
}
