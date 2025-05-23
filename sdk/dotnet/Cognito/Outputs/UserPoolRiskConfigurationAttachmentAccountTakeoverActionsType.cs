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
    public sealed class UserPoolRiskConfigurationAttachmentAccountTakeoverActionsType
    {
        /// <summary>
        /// The action that you assign to a high-risk assessment by threat protection.
        /// </summary>
        public readonly Outputs.UserPoolRiskConfigurationAttachmentAccountTakeoverActionType? HighAction;
        /// <summary>
        /// The action that you assign to a low-risk assessment by threat protection.
        /// </summary>
        public readonly Outputs.UserPoolRiskConfigurationAttachmentAccountTakeoverActionType? LowAction;
        /// <summary>
        /// The action that you assign to a medium-risk assessment by threat protection.
        /// </summary>
        public readonly Outputs.UserPoolRiskConfigurationAttachmentAccountTakeoverActionType? MediumAction;

        [OutputConstructor]
        private UserPoolRiskConfigurationAttachmentAccountTakeoverActionsType(
            Outputs.UserPoolRiskConfigurationAttachmentAccountTakeoverActionType? highAction,

            Outputs.UserPoolRiskConfigurationAttachmentAccountTakeoverActionType? lowAction,

            Outputs.UserPoolRiskConfigurationAttachmentAccountTakeoverActionType? mediumAction)
        {
            HighAction = highAction;
            LowAction = lowAction;
            MediumAction = mediumAction;
        }
    }
}
