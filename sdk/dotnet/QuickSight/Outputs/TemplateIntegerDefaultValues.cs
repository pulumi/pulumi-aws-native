// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TemplateIntegerDefaultValues
    {
        /// <summary>
        /// The dynamic value of the `IntegerDefaultValues` . Different defaults are displayed according to users, groups, and values mapping.
        /// </summary>
        public readonly Outputs.TemplateDynamicDefaultValue? DynamicValue;
        /// <summary>
        /// The static values of the `IntegerDefaultValues` .
        /// </summary>
        public readonly ImmutableArray<double> StaticValues;

        [OutputConstructor]
        private TemplateIntegerDefaultValues(
            Outputs.TemplateDynamicDefaultValue? dynamicValue,

            ImmutableArray<double> staticValues)
        {
            DynamicValue = dynamicValue;
            StaticValues = staticValues;
        }
    }
}
