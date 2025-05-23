// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Inputs
{

    /// <summary>
    /// A label for tagging Lex resources
    /// </summary>
    public sealed class BotAliasTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A string used to identify this tag
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// A string containing the value for the tag
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public BotAliasTagArgs()
        {
        }
        public static new BotAliasTagArgs Empty => new BotAliasTagArgs();
    }
}
