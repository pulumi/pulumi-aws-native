// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT
{
    public static class GetProvisioningTemplate
    {
        /// <summary>
        /// Creates a fleet provisioning template.
        /// </summary>
        public static Task<GetProvisioningTemplateResult> InvokeAsync(GetProvisioningTemplateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProvisioningTemplateResult>("aws-native:iot:getProvisioningTemplate", args ?? new GetProvisioningTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// Creates a fleet provisioning template.
        /// </summary>
        public static Output<GetProvisioningTemplateResult> Invoke(GetProvisioningTemplateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProvisioningTemplateResult>("aws-native:iot:getProvisioningTemplate", args ?? new GetProvisioningTemplateInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Creates a fleet provisioning template.
        /// </summary>
        public static Output<GetProvisioningTemplateResult> Invoke(GetProvisioningTemplateInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetProvisioningTemplateResult>("aws-native:iot:getProvisioningTemplate", args ?? new GetProvisioningTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProvisioningTemplateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the fleet provisioning template.
        /// </summary>
        [Input("templateName", required: true)]
        public string TemplateName { get; set; } = null!;

        public GetProvisioningTemplateArgs()
        {
        }
        public static new GetProvisioningTemplateArgs Empty => new GetProvisioningTemplateArgs();
    }

    public sealed class GetProvisioningTemplateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the fleet provisioning template.
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        public GetProvisioningTemplateInvokeArgs()
        {
        }
        public static new GetProvisioningTemplateInvokeArgs Empty => new GetProvisioningTemplateInvokeArgs();
    }


    [OutputType]
    public sealed class GetProvisioningTemplateResult
    {
        /// <summary>
        /// The description of the fleet provisioning template.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// True to enable the fleet provisioning template, otherwise false.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Creates a pre-provisioning hook template.
        /// </summary>
        public readonly Outputs.ProvisioningTemplateProvisioningHook? PreProvisioningHook;
        /// <summary>
        /// The role ARN for the role associated with the fleet provisioning template. This IoT role grants permission to provision a device.
        /// </summary>
        public readonly string? ProvisioningRoleArn;
        /// <summary>
        /// Metadata that can be used to manage the fleet provisioning template.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The ARN that identifies the provisioning template.
        /// </summary>
        public readonly string? TemplateArn;
        /// <summary>
        /// The JSON formatted contents of the fleet provisioning template version.
        /// </summary>
        public readonly string? TemplateBody;

        [OutputConstructor]
        private GetProvisioningTemplateResult(
            string? description,

            bool? enabled,

            Outputs.ProvisioningTemplateProvisioningHook? preProvisioningHook,

            string? provisioningRoleArn,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            string? templateArn,

            string? templateBody)
        {
            Description = description;
            Enabled = enabled;
            PreProvisioningHook = preProvisioningHook;
            ProvisioningRoleArn = provisioningRoleArn;
            Tags = tags;
            TemplateArn = templateArn;
            TemplateBody = templateBody;
        }
    }
}
