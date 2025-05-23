// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DeviceFarm.Inputs
{

    /// <summary>
    /// Represents a condition for a device pool.
    /// </summary>
    public sealed class DevicePoolRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The rule's stringified attribute.
        /// </summary>
        [Input("attribute")]
        public Input<Pulumi.AwsNative.DeviceFarm.DevicePoolRuleAttribute>? Attribute { get; set; }

        /// <summary>
        /// Specifies how Device Farm compares the rule's attribute to the value.
        /// </summary>
        [Input("operator")]
        public Input<Pulumi.AwsNative.DeviceFarm.DevicePoolRuleOperator>? Operator { get; set; }

        /// <summary>
        /// The rule's value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public DevicePoolRuleArgs()
        {
        }
        public static new DevicePoolRuleArgs Empty => new DevicePoolRuleArgs();
    }
}
