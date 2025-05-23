// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    /// <summary>
    /// Parameters to define a mitigation action that adds a blank policy to restrict permissions.
    /// </summary>
    public sealed class MitigationActionReplaceDefaultPolicyVersionParamsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the template to be applied. The only supported value is `BLANK_POLICY` .
        /// </summary>
        [Input("templateName", required: true)]
        public Input<Pulumi.AwsNative.IoT.MitigationActionReplaceDefaultPolicyVersionParamsTemplateName> TemplateName { get; set; } = null!;

        public MitigationActionReplaceDefaultPolicyVersionParamsArgs()
        {
        }
        public static new MitigationActionReplaceDefaultPolicyVersionParamsArgs Empty => new MitigationActionReplaceDefaultPolicyVersionParamsArgs();
    }
}
