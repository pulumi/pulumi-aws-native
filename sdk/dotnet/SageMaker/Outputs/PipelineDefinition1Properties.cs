// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    [OutputType]
    public sealed class PipelineDefinition1Properties
    {
        public readonly Outputs.PipelineS3Location PipelineDefinitionS3Location;

        [OutputConstructor]
        private PipelineDefinition1Properties(Outputs.PipelineS3Location pipelineDefinitionS3Location)
        {
            PipelineDefinitionS3Location = pipelineDefinitionS3Location;
        }
    }
}