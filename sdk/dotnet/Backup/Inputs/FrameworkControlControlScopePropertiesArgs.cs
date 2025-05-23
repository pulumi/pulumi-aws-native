// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Backup.Inputs
{

    /// <summary>
    /// The scope of a control. The control scope defines what the control will evaluate. Three examples of control scopes are: a specific backup plan, all backup plans with a specific tag, or all backup plans.
    /// </summary>
    public sealed class FrameworkControlControlScopePropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("complianceResourceIds")]
        private InputList<string>? _complianceResourceIds;

        /// <summary>
        /// The ID of the only AWS resource that you want your control scope to contain.
        /// </summary>
        public InputList<string> ComplianceResourceIds
        {
            get => _complianceResourceIds ?? (_complianceResourceIds = new InputList<string>());
            set => _complianceResourceIds = value;
        }

        [Input("complianceResourceTypes")]
        private InputList<string>? _complianceResourceTypes;

        /// <summary>
        /// Describes whether the control scope includes one or more types of resources, such as `EFS` or `RDS`.
        /// </summary>
        public InputList<string> ComplianceResourceTypes
        {
            get => _complianceResourceTypes ?? (_complianceResourceTypes = new InputList<string>());
            set => _complianceResourceTypes = value;
        }

        [Input("tags")]
        private InputList<Inputs.FrameworkTagArgs>? _tags;

        /// <summary>
        /// Describes whether the control scope includes resources with one or more tags. Each tag is a key-value pair.
        /// </summary>
        public InputList<Inputs.FrameworkTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.FrameworkTagArgs>());
            set => _tags = value;
        }

        public FrameworkControlControlScopePropertiesArgs()
        {
        }
        public static new FrameworkControlControlScopePropertiesArgs Empty => new FrameworkControlControlScopePropertiesArgs();
    }
}
