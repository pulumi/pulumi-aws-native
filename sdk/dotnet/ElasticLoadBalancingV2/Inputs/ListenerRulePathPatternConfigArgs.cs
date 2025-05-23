// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Inputs
{

    /// <summary>
    /// Information about a path pattern condition.
    /// </summary>
    public sealed class ListenerRulePathPatternConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// The path patterns to compare against the request URL. The maximum size of each string is 128 characters. The comparison is case sensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character).
        ///  If you specify multiple strings, the condition is satisfied if one of them matches the request URL. The path pattern is compared only to the path of the URL, not to its query string.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public ListenerRulePathPatternConfigArgs()
        {
        }
        public static new ListenerRulePathPatternConfigArgs Empty => new ListenerRulePathPatternConfigArgs();
    }
}
