// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles.Inputs
{

    /// <summary>
    /// The details of a single attribute item specified in the mathematical expression.
    /// </summary>
    public sealed class CalculatedAttributeDefinitionAttributeItemArgs : global::Pulumi.ResourceArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public CalculatedAttributeDefinitionAttributeItemArgs()
        {
        }
        public static new CalculatedAttributeDefinitionAttributeItemArgs Empty => new CalculatedAttributeDefinitionAttributeItemArgs();
    }
}