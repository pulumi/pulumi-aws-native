// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;The column schema.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class TemplateColumnSchema
    {
        /// <summary>
        /// &lt;p&gt;The data type of the column schema.&lt;/p&gt;
        /// </summary>
        public readonly string? DataType;
        /// <summary>
        /// &lt;p&gt;The geographic role of the column schema.&lt;/p&gt;
        /// </summary>
        public readonly string? GeographicRole;
        /// <summary>
        /// &lt;p&gt;The name of the column schema.&lt;/p&gt;
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private TemplateColumnSchema(
            string? dataType,

            string? geographicRole,

            string? name)
        {
            DataType = dataType;
            GeographicRole = geographicRole;
            Name = name;
        }
    }
}
