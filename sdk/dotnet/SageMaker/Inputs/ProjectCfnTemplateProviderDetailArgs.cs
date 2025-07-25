// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// CloudFormation template provider details for a SageMaker project.
    /// </summary>
    public sealed class ProjectCfnTemplateProviderDetailArgs : global::Pulumi.ResourceArgs
    {
        [Input("parameters")]
        private InputList<Inputs.ProjectCfnStackParameterArgs>? _parameters;
        public InputList<Inputs.ProjectCfnStackParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ProjectCfnStackParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role used by the template provider.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// The name of the template used for the project.
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        /// <summary>
        /// The URL of the CloudFormation template.
        /// </summary>
        [Input("templateUrl", required: true)]
        public Input<string> TemplateUrl { get; set; } = null!;

        public ProjectCfnTemplateProviderDetailArgs()
        {
        }
        public static new ProjectCfnTemplateProviderDetailArgs Empty => new ProjectCfnTemplateProviderDetailArgs();
    }
}
