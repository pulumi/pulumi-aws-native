// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Inputs
{

    public sealed class ComponentDataConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("identifiers")]
        private InputList<string>? _identifiers;
        public InputList<string> Identifiers
        {
            get => _identifiers ?? (_identifiers = new InputList<string>());
            set => _identifiers = value;
        }

        [Input("model", required: true)]
        public Input<string> Model { get; set; } = null!;

        [Input("predicate")]
        public Input<Inputs.ComponentPredicateArgs>? Predicate { get; set; }

        [Input("sort")]
        private InputList<Inputs.ComponentSortPropertyArgs>? _sort;
        public InputList<Inputs.ComponentSortPropertyArgs> Sort
        {
            get => _sort ?? (_sort = new InputList<Inputs.ComponentSortPropertyArgs>());
            set => _sort = value;
        }

        public ComponentDataConfigurationArgs()
        {
        }
        public static new ComponentDataConfigurationArgs Empty => new ComponentDataConfigurationArgs();
    }
}