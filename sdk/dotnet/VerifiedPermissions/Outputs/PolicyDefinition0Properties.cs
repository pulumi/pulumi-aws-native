// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VerifiedPermissions.Outputs
{

    [OutputType]
    public sealed class PolicyDefinition0Properties
    {
        public readonly Outputs.PolicyStaticPolicyDefinition Static;

        [OutputConstructor]
        private PolicyDefinition0Properties(Outputs.PolicyStaticPolicyDefinition @static)
        {
            Static = @static;
        }
    }
}
