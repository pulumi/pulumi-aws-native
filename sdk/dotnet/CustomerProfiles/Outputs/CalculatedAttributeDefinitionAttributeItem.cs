// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles.Outputs
{

    /// <summary>
    /// The details of a single attribute item specified in the mathematical expression.
    /// </summary>
    [OutputType]
    public sealed class CalculatedAttributeDefinitionAttributeItem
    {
        public readonly string Name;

        [OutputConstructor]
        private CalculatedAttributeDefinitionAttributeItem(string name)
        {
            Name = name;
        }
    }
}