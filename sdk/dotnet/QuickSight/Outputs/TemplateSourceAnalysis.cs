// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;The source analysis of the template.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class TemplateSourceAnalysis
    {
        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) of the resource.&lt;/p&gt;
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// &lt;p&gt;A structure containing information about the dataset references used as placeholders
        ///             in the template.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.TemplateDataSetReference> DataSetReferences;

        [OutputConstructor]
        private TemplateSourceAnalysis(
            string arn,

            ImmutableArray<Outputs.TemplateDataSetReference> dataSetReferences)
        {
            Arn = arn;
            DataSetReferences = dataSetReferences;
        }
    }
}