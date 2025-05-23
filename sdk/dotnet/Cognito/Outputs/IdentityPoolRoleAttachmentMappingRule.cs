// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Outputs
{

    [OutputType]
    public sealed class IdentityPoolRoleAttachmentMappingRule
    {
        public readonly string Claim;
        public readonly string MatchType;
        public readonly string RoleArn;
        public readonly string Value;

        [OutputConstructor]
        private IdentityPoolRoleAttachmentMappingRule(
            string claim,

            string matchType,

            string roleArn,

            string value)
        {
            Claim = claim;
            MatchType = matchType;
            RoleArn = roleArn;
            Value = value;
        }
    }
}
