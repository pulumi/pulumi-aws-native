// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Inputs
{

    public sealed class RuleGroupNotStatementArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The statement to negate. You can use any statement that can be nested.
        /// </summary>
        [Input("statement", required: true)]
        public Input<Inputs.RuleGroupStatementArgs> Statement { get; set; } = null!;

        public RuleGroupNotStatementArgs()
        {
        }
        public static new RuleGroupNotStatementArgs Empty => new RuleGroupNotStatementArgs();
    }
}
