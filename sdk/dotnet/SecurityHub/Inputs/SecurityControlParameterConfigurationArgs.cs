// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub.Inputs
{

    public sealed class SecurityControlParameterConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("value")]
        public Input<Inputs.SecurityControlParameterValueArgs>? Value { get; set; }

        [Input("valueType", required: true)]
        public Input<Pulumi.AwsNative.SecurityHub.SecurityControlParameterConfigurationValueType> ValueType { get; set; } = null!;

        public SecurityControlParameterConfigurationArgs()
        {
        }
        public static new SecurityControlParameterConfigurationArgs Empty => new SecurityControlParameterConfigurationArgs();
    }
}