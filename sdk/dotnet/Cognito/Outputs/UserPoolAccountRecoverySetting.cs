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
    public sealed class UserPoolAccountRecoverySetting
    {
        /// <summary>
        /// The list of options and priorities for user message delivery in forgot-password operations. Sets or displays user pool preferences for email or SMS message priority, whether users should fall back to a second delivery method, and whether passwords should only be reset by administrators.
        /// </summary>
        public readonly ImmutableArray<Outputs.UserPoolRecoveryOption> RecoveryMechanisms;

        [OutputConstructor]
        private UserPoolAccountRecoverySetting(ImmutableArray<Outputs.UserPoolRecoveryOption> recoveryMechanisms)
        {
            RecoveryMechanisms = recoveryMechanisms;
        }
    }
}
