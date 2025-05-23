// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    /// <summary>
    /// &lt;p&gt;Dashboard source template.&lt;/p&gt;
    /// </summary>
    public sealed class DashboardSourceTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) of the resource.&lt;/p&gt;
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        [Input("dataSetReferences", required: true)]
        private InputList<Inputs.DashboardDataSetReferenceArgs>? _dataSetReferences;

        /// <summary>
        /// &lt;p&gt;Dataset references.&lt;/p&gt;
        /// </summary>
        public InputList<Inputs.DashboardDataSetReferenceArgs> DataSetReferences
        {
            get => _dataSetReferences ?? (_dataSetReferences = new InputList<Inputs.DashboardDataSetReferenceArgs>());
            set => _dataSetReferences = value;
        }

        public DashboardSourceTemplateArgs()
        {
        }
        public static new DashboardSourceTemplateArgs Empty => new DashboardSourceTemplateArgs();
    }
}
