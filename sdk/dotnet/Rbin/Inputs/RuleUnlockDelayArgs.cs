// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Rbin.Inputs
{

    public sealed class RuleUnlockDelayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unit of time in which to measure the unlock delay. Currently, the unlock delay can be measure only in days.
        /// </summary>
        [Input("unlockDelayUnit")]
        public Input<Pulumi.AwsNative.Rbin.RuleUnlockDelayUnlockDelayUnit>? UnlockDelayUnit { get; set; }

        /// <summary>
        /// The unlock delay period, measured in the unit specified for UnlockDelayUnit.
        /// </summary>
        [Input("unlockDelayValue")]
        public Input<int>? UnlockDelayValue { get; set; }

        public RuleUnlockDelayArgs()
        {
        }
        public static new RuleUnlockDelayArgs Empty => new RuleUnlockDelayArgs();
    }
}
