// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    /// <summary>
    /// A key-value pair to associate expression's substitution variable names with their values
    /// </summary>
    [OutputType]
    public sealed class RulesetSubstitutionValue
    {
        /// <summary>
        /// Value or column name
        /// </summary>
        public readonly string Value;
        /// <summary>
        /// Variable name
        /// </summary>
        public readonly string ValueReference;

        [OutputConstructor]
        private RulesetSubstitutionValue(
            string value,

            string valueReference)
        {
            Value = value;
            ValueReference = valueReference;
        }
    }
}