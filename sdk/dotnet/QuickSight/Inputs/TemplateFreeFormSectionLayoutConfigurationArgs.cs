// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateFreeFormSectionLayoutConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("elements", required: true)]
        private InputList<Inputs.TemplateFreeFormLayoutElementArgs>? _elements;
        public InputList<Inputs.TemplateFreeFormLayoutElementArgs> Elements
        {
            get => _elements ?? (_elements = new InputList<Inputs.TemplateFreeFormLayoutElementArgs>());
            set => _elements = value;
        }

        public TemplateFreeFormSectionLayoutConfigurationArgs()
        {
        }
        public static new TemplateFreeFormSectionLayoutConfigurationArgs Empty => new TemplateFreeFormSectionLayoutConfigurationArgs();
    }
}