// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisEmptyVisual
    {
        public readonly ImmutableArray<Outputs.AnalysisVisualCustomAction> Actions;
        public readonly string DataSetIdentifier;
        public readonly string VisualId;

        [OutputConstructor]
        private AnalysisEmptyVisual(
            ImmutableArray<Outputs.AnalysisVisualCustomAction> actions,

            string dataSetIdentifier,

            string visualId)
        {
            Actions = actions;
            DataSetIdentifier = dataSetIdentifier;
            VisualId = visualId;
        }
    }
}