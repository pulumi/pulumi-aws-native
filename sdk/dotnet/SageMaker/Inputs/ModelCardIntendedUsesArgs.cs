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
    /// Intended usage of model.
    /// </summary>
    public sealed class ModelCardIntendedUsesArgs : global::Pulumi.ResourceArgs
    {
        [Input("explanationsForRiskRating")]
        public Input<string>? ExplanationsForRiskRating { get; set; }

        [Input("factorsAffectingModelEfficiency")]
        public Input<string>? FactorsAffectingModelEfficiency { get; set; }

        /// <summary>
        /// intended use cases.
        /// </summary>
        [Input("intendedUses")]
        public Input<string>? IntendedUses { get; set; }

        /// <summary>
        /// Why the model was developed?
        /// </summary>
        [Input("purposeOfModel")]
        public Input<string>? PurposeOfModel { get; set; }

        [Input("riskRating")]
        public Input<Pulumi.AwsNative.SageMaker.ModelCardRiskRating>? RiskRating { get; set; }

        public ModelCardIntendedUsesArgs()
        {
        }
        public static new ModelCardIntendedUsesArgs Empty => new ModelCardIntendedUsesArgs();
    }
}