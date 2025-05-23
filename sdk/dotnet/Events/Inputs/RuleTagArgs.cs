// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events.Inputs
{

    public sealed class RuleTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A string you can use to assign a value. The combination of tag keys and values can help you organize and categorize your resources.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The value for the specified tag key.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public RuleTagArgs()
        {
        }
        public static new RuleTagArgs Empty => new RuleTagArgs();
    }
}
