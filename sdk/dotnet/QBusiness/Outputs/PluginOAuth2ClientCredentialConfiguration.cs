// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QBusiness.Outputs
{

    [OutputType]
    public sealed class PluginOAuth2ClientCredentialConfiguration
    {
        public readonly string RoleArn;
        public readonly string SecretArn;

        [OutputConstructor]
        private PluginOAuth2ClientCredentialConfiguration(
            string roleArn,

            string secretArn)
        {
            RoleArn = roleArn;
            SecretArn = secretArn;
        }
    }
}