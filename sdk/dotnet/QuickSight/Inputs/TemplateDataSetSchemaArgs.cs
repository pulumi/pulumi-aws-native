// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateDataSetSchemaArgs : global::Pulumi.ResourceArgs
    {
        [Input("columnSchemaList")]
        private InputList<Inputs.TemplateColumnSchemaArgs>? _columnSchemaList;
        public InputList<Inputs.TemplateColumnSchemaArgs> ColumnSchemaList
        {
            get => _columnSchemaList ?? (_columnSchemaList = new InputList<Inputs.TemplateColumnSchemaArgs>());
            set => _columnSchemaList = value;
        }

        public TemplateDataSetSchemaArgs()
        {
        }
        public static new TemplateDataSetSchemaArgs Empty => new TemplateDataSetSchemaArgs();
    }
}