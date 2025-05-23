// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Outputs
{

    [OutputType]
    public sealed class PipeMultiMeasureAttributeMapping
    {
        /// <summary>
        /// Dynamic path to the measurement attribute in the source event.
        /// </summary>
        public readonly string MeasureValue;
        /// <summary>
        /// Data type of the measurement attribute in the source event.
        /// </summary>
        public readonly Pulumi.AwsNative.Pipes.PipeMeasureValueType MeasureValueType;
        /// <summary>
        /// Target measure name to be used.
        /// </summary>
        public readonly string MultiMeasureAttributeName;

        [OutputConstructor]
        private PipeMultiMeasureAttributeMapping(
            string measureValue,

            Pulumi.AwsNative.Pipes.PipeMeasureValueType measureValueType,

            string multiMeasureAttributeName)
        {
            MeasureValue = measureValue;
            MeasureValueType = measureValueType;
            MultiMeasureAttributeName = multiMeasureAttributeName;
        }
    }
}
