// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class ModelCardMatrixMetricArgs : global::Pulumi.ResourceArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.SageMaker.ModelCardMatrixMetricType> Type { get; set; } = null!;

        [Input("value", required: true)]
        private InputList<ImmutableArray<double>>? _value;
        public InputList<ImmutableArray<double>> Value
        {
            get => _value ?? (_value = new InputList<ImmutableArray<double>>());
            set => _value = value;
        }

        [Input("xAxisName")]
        private InputList<string>? _xAxisName;
        public InputList<string> XAxisName
        {
            get => _xAxisName ?? (_xAxisName = new InputList<string>());
            set => _xAxisName = value;
        }

        [Input("yAxisName")]
        private InputList<string>? _yAxisName;
        public InputList<string> YAxisName
        {
            get => _yAxisName ?? (_yAxisName = new InputList<string>());
            set => _yAxisName = value;
        }

        public ModelCardMatrixMetricArgs()
        {
        }
        public static new ModelCardMatrixMetricArgs Empty => new ModelCardMatrixMetricArgs();
    }
}
