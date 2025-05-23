// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
        /// <summary>
        /// The column that contains the default value of each user or group.
        /// </summary>
        public readonly Outputs.TemplateColumnIdentifier DefaultValueColumn;
        /// <summary>
        /// The column that contains the group name.
        /// </summary>
        public readonly Outputs.TemplateColumnIdentifier? GroupNameColumn;
        /// <summary>
        /// The column that contains the username.
        /// </summary>
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
