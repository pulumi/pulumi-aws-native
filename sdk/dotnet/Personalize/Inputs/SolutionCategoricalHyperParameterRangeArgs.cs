// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Personalize.Inputs
{

    /// <summary>
    /// Provides the name and values of a Categorical hyperparameter.
    /// </summary>
    public sealed class SolutionCategoricalHyperParameterRangeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the hyperparameter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// A list of the categories for the hyperparameter.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public SolutionCategoricalHyperParameterRangeArgs()
        {
        }
        public static new SolutionCategoricalHyperParameterRangeArgs Empty => new SolutionCategoricalHyperParameterRangeArgs();
    }
}
