// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// Business details.
    /// </summary>
    public sealed class ModelCardBusinessDetailsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// What business problem does the model solve?
        /// </summary>
        [Input("businessProblem")]
        public Input<string>? BusinessProblem { get; set; }

        /// <summary>
        /// Business stakeholders.
        /// </summary>
        [Input("businessStakeholders")]
        public Input<string>? BusinessStakeholders { get; set; }

        /// <summary>
        /// Line of business.
        /// </summary>
        [Input("lineOfBusiness")]
        public Input<string>? LineOfBusiness { get; set; }

        public ModelCardBusinessDetailsArgs()
        {
        }
        public static new ModelCardBusinessDetailsArgs Empty => new ModelCardBusinessDetailsArgs();
    }
}