// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class ModelPackageContainerDefinitionModelInputPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The input configuration object for the model.
        /// </summary>
        [Input("dataInputConfig", required: true)]
        public Input<string> DataInputConfig { get; set; } = null!;

        public ModelPackageContainerDefinitionModelInputPropertiesArgs()
        {
        }
        public static new ModelPackageContainerDefinitionModelInputPropertiesArgs Empty => new ModelPackageContainerDefinitionModelInputPropertiesArgs();
    }
}