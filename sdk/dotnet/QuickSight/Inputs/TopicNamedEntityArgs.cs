// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TopicNamedEntityArgs : global::Pulumi.ResourceArgs
    {
        [Input("definition")]
        private InputList<Inputs.TopicNamedEntityDefinitionArgs>? _definition;

        /// <summary>
        /// The definition of a named entity.
        /// </summary>
        public InputList<Inputs.TopicNamedEntityDefinitionArgs> Definition
        {
            get => _definition ?? (_definition = new InputList<Inputs.TopicNamedEntityDefinitionArgs>());
            set => _definition = value;
        }

        /// <summary>
        /// The description of the named entity.
        /// </summary>
        [Input("entityDescription")]
        public Input<string>? EntityDescription { get; set; }

        /// <summary>
        /// The name of the named entity.
        /// </summary>
        [Input("entityName", required: true)]
        public Input<string> EntityName { get; set; } = null!;

        [Input("entitySynonyms")]
        private InputList<string>? _entitySynonyms;

        /// <summary>
        /// The other names or aliases for the named entity.
        /// </summary>
        public InputList<string> EntitySynonyms
        {
            get => _entitySynonyms ?? (_entitySynonyms = new InputList<string>());
            set => _entitySynonyms = value;
        }

        /// <summary>
        /// The type of named entity that a topic represents.
        /// </summary>
        [Input("semanticEntityType")]
        public Input<Inputs.TopicSemanticEntityTypeArgs>? SemanticEntityType { get; set; }

        public TopicNamedEntityArgs()
        {
        }
        public static new TopicNamedEntityArgs Empty => new TopicNamedEntityArgs();
    }
}
