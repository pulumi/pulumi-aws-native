// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Inputs
{

    public sealed class UserPoolRiskConfigurationAttachmentRiskExceptionConfigurationTypeArgs : global::Pulumi.ResourceArgs
    {
        [Input("blockedIpRangeList")]
        private InputList<string>? _blockedIpRangeList;

        /// <summary>
        /// An always-block IP address list. Overrides the risk decision and always blocks authentication requests. This parameter is displayed and set in CIDR notation.
        /// </summary>
        public InputList<string> BlockedIpRangeList
        {
            get => _blockedIpRangeList ?? (_blockedIpRangeList = new InputList<string>());
            set => _blockedIpRangeList = value;
        }

        [Input("skippedIpRangeList")]
        private InputList<string>? _skippedIpRangeList;

        /// <summary>
        /// An always-allow IP address list. Risk detection isn't performed on the IP addresses in this range list. This parameter is displayed and set in CIDR notation.
        /// </summary>
        public InputList<string> SkippedIpRangeList
        {
            get => _skippedIpRangeList ?? (_skippedIpRangeList = new InputList<string>());
            set => _skippedIpRangeList = value;
        }

        public UserPoolRiskConfigurationAttachmentRiskExceptionConfigurationTypeArgs()
        {
        }
        public static new UserPoolRiskConfigurationAttachmentRiskExceptionConfigurationTypeArgs Empty => new UserPoolRiskConfigurationAttachmentRiskExceptionConfigurationTypeArgs();
    }
}
