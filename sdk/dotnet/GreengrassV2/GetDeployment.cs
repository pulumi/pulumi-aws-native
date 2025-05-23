// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GreengrassV2
{
    public static class GetDeployment
    {
        /// <summary>
        /// Resource for Greengrass V2 deployment.
        /// </summary>
        public static Task<GetDeploymentResult> InvokeAsync(GetDeploymentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDeploymentResult>("aws-native:greengrassv2:getDeployment", args ?? new GetDeploymentArgs(), options.WithDefaults());

        /// <summary>
        /// Resource for Greengrass V2 deployment.
        /// </summary>
        public static Output<GetDeploymentResult> Invoke(GetDeploymentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeploymentResult>("aws-native:greengrassv2:getDeployment", args ?? new GetDeploymentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource for Greengrass V2 deployment.
        /// </summary>
        public static Output<GetDeploymentResult> Invoke(GetDeploymentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeploymentResult>("aws-native:greengrassv2:getDeployment", args ?? new GetDeploymentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeploymentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the deployment.
        /// </summary>
        [Input("deploymentId", required: true)]
        public string DeploymentId { get; set; } = null!;

        public GetDeploymentArgs()
        {
        }
        public static new GetDeploymentArgs Empty => new GetDeploymentArgs();
    }

    public sealed class GetDeploymentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the deployment.
        /// </summary>
        [Input("deploymentId", required: true)]
        public Input<string> DeploymentId { get; set; } = null!;

        public GetDeploymentInvokeArgs()
        {
        }
        public static new GetDeploymentInvokeArgs Empty => new GetDeploymentInvokeArgs();
    }


    [OutputType]
    public sealed class GetDeploymentResult
    {
        /// <summary>
        /// The ID of the deployment.
        /// </summary>
        public readonly string? DeploymentId;
        /// <summary>
        /// Application-specific metadata to attach to the deployment. You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tag your AWS IoT Greengrass Version 2 resources](https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html) in the *AWS IoT Greengrass V2 Developer Guide* .
        /// 
        /// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
        /// 
        /// ```json
        /// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
        /// }
        /// ```
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetDeploymentResult(
            string? deploymentId,

            ImmutableDictionary<string, string>? tags)
        {
            DeploymentId = deploymentId;
            Tags = tags;
        }
    }
}
