// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// Contains data, such as the inputs and targeted instance types that are used in the process of validating the model package.
    /// </summary>
    public sealed class ModelPackageValidationProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the profile for the model package.
        /// </summary>
        [Input("profileName", required: true)]
        public Input<string> ProfileName { get; set; } = null!;

        /// <summary>
        /// The `TransformJobDefinition` object that describes the transform job used for the validation of the model package.
        /// </summary>
        [Input("transformJobDefinition", required: true)]
        public Input<Inputs.ModelPackageTransformJobDefinitionArgs> TransformJobDefinition { get; set; } = null!;

        public ModelPackageValidationProfileArgs()
        {
        }
        public static new ModelPackageValidationProfileArgs Empty => new ModelPackageValidationProfileArgs();
    }
}
