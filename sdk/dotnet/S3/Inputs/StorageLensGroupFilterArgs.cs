// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    /// <summary>
    /// Sets the Storage Lens Group filter.
    /// </summary>
    public sealed class StorageLensGroupFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("and")]
        public Input<Inputs.StorageLensGroupAndArgs>? And { get; set; }

        [Input("matchAnyPrefix")]
        private InputList<string>? _matchAnyPrefix;
        public InputList<string> MatchAnyPrefix
        {
            get => _matchAnyPrefix ?? (_matchAnyPrefix = new InputList<string>());
            set => _matchAnyPrefix = value;
        }

        [Input("matchAnySuffix")]
        private InputList<string>? _matchAnySuffix;
        public InputList<string> MatchAnySuffix
        {
            get => _matchAnySuffix ?? (_matchAnySuffix = new InputList<string>());
            set => _matchAnySuffix = value;
        }

        [Input("matchAnyTag")]
        private InputList<Inputs.StorageLensGroupTagArgs>? _matchAnyTag;
        public InputList<Inputs.StorageLensGroupTagArgs> MatchAnyTag
        {
            get => _matchAnyTag ?? (_matchAnyTag = new InputList<Inputs.StorageLensGroupTagArgs>());
            set => _matchAnyTag = value;
        }

        [Input("matchObjectAge")]
        public Input<Inputs.StorageLensGroupMatchObjectAgeArgs>? MatchObjectAge { get; set; }

        [Input("matchObjectSize")]
        public Input<Inputs.StorageLensGroupMatchObjectSizeArgs>? MatchObjectSize { get; set; }

        [Input("or")]
        public Input<Inputs.StorageLensGroupOrArgs>? Or { get; set; }

        public StorageLensGroupFilterArgs()
        {
        }
        public static new StorageLensGroupFilterArgs Empty => new StorageLensGroupFilterArgs();
    }
}