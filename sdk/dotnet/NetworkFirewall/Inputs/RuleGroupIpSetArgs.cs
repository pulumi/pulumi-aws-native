// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Inputs
{

    public sealed class RuleGroupIpSetArgs : global::Pulumi.ResourceArgs
    {
        [Input("definition")]
        private InputList<string>? _definition;
        public InputList<string> Definition
        {
            get => _definition ?? (_definition = new InputList<string>());
            set => _definition = value;
        }

        public RuleGroupIpSetArgs()
        {
        }
        public static new RuleGroupIpSetArgs Empty => new RuleGroupIpSetArgs();
    }
}
