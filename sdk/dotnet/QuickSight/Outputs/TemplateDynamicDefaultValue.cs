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
    public sealed class TemplateDynamicDefaultValue
    {
        public readonly Outputs.TemplateColumnIdentifier DefaultValueColumn;
        public readonly Outputs.TemplateColumnIdentifier? GroupNameColumn;
        public readonly Outputs.TemplateColumnIdentifier? UserNameColumn;

        [OutputConstructor]
        private TemplateDynamicDefaultValue(
            Outputs.TemplateColumnIdentifier defaultValueColumn,

            Outputs.TemplateColumnIdentifier? groupNameColumn,

            Outputs.TemplateColumnIdentifier? userNameColumn)
        {
            DefaultValueColumn = defaultValueColumn;
            GroupNameColumn = groupNameColumn;
            UserNameColumn = userNameColumn;
        }
    }
}