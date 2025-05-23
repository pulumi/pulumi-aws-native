// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFormation
{
    /// <summary>
    /// The Custom Resource Emulator allows you to use AWS CloudFormation Custom Resources directly in your Pulumi programs. It provides a way to invoke AWS Lambda functions that implement custom provisioning logic following the CloudFormation Custom Resource protocol.
    /// 
    /// &gt; **Note**: Currently, only Lambda-backed Custom Resources are supported. SNS-backed Custom Resources are not supported at this time.
    /// 
    /// ## Example Usage
    /// 
    /// A full example of creating a CloudFormation Custom Resource Lambda function and using it in Pulumi can be found [here](https://github.com/pulumi/pulumi-aws-native/tree/master/examples/cfn-custom-resource).
    /// 
    /// ## About CloudFormation Custom Resources
    /// 
    /// CloudFormation Custom Resources allow you to write custom provisioning logic for resources that aren't directly available as AWS CloudFormation resource types. Common use cases include:
    /// 
    /// - Implementing complex provisioning logic
    /// - Performing custom validations or transformations
    /// - Integrating with third-party services
    /// - Implementing organization-specific infrastructure patterns
    /// 
    /// For more information about CloudFormation Custom Resources, see [Custom Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources.html) in the AWS CloudFormation User Guide.
    /// 
    /// ## Permissions
    /// 
    /// The IAM principal used by your Pulumi program must have the following permissions:
    /// 
    /// 1. `lambda:InvokeFunction` on the Lambda function specified in `serviceToken`
    /// 2. S3 permissions on the bucket specified in `bucketName`:
    ///    - `s3:PutObject`
    ///    - `s3:GetObject`
    ///    - `s3:HeadObject`
    /// 
    /// ## Lambda Function Requirements
    /// 
    /// The Lambda function specified in `serviceToken` must implement the CloudFormation Custom Resource lifecycle.
    /// For detailed information about implementing Lambda-backed Custom Resources, see [AWS Lambda-backed Custom Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources-lambda.html) in the AWS CloudFormation User Guide.
    /// 
    /// ## Timeouts
    /// 
    /// Custom Resources have a default timeout of 60 minutes, matching the CloudFormation timeout for custom resource operations. You can customize it using the [`customTimeouts`](https://www.pulumi.com/docs/iac/concepts/options/customtimeouts/) resource option.
    /// </summary>
    [AwsNativeResourceType("aws-native:cloudformation:CustomResourceEmulator")]
    public partial class CustomResourceEmulator : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the S3 bucket to use for storing the response from the Custom Resource.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// The response data returned by invoking the Custom Resource.
        /// </summary>
        [Output("data")]
        public Output<ImmutableDictionary<string, object>> Data { get; private set; } = null!;

        /// <summary>
        /// Whether the response data contains sensitive information that should be marked as secret and not logged.
        /// </summary>
        [Output("noEcho")]
        public Output<bool> NoEcho { get; private set; } = null!;

        /// <summary>
        /// The name or unique identifier that corresponds to the `PhysicalResourceId` included in the Custom Resource response. If no `PhysicalResourceId` is provided in the response, a random ID will be generated.
        /// </summary>
        [Output("physicalResourceId")]
        public Output<string> PhysicalResourceId { get; private set; } = null!;

        /// <summary>
        /// The CloudFormation type of the Custom Resource provider. For example, `Custom::MyCustomResource`.
        /// </summary>
        [Output("resourceType")]
        public Output<string> ResourceType { get; private set; } = null!;

        /// <summary>
        /// The service token, such as a Lambda function ARN, that is invoked when the resource is created, updated, or deleted.
        /// </summary>
        [Output("serviceToken")]
        public Output<string> ServiceToken { get; private set; } = null!;

        /// <summary>
        /// A stand-in value for the CloudFormation stack ID.
        /// </summary>
        [Output("stackId")]
        public Output<string> StackId { get; private set; } = null!;


        /// <summary>
        /// Create a CustomResourceEmulator resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomResourceEmulator(string name, CustomResourceEmulatorArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cloudformation:CustomResourceEmulator", name, args ?? new CustomResourceEmulatorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomResourceEmulator(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cloudformation:CustomResourceEmulator", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CustomResourceEmulator resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomResourceEmulator Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CustomResourceEmulator(name, id, options);
        }
    }

    public sealed class CustomResourceEmulatorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The prefix to use for the bucket key when storing the response from the Custom Resource provider.
        /// </summary>
        [Input("bucketKeyPrefix", required: true)]
        public Input<string> BucketKeyPrefix { get; set; } = null!;

        /// <summary>
        /// The name of the S3 bucket to use for storing the response from the Custom Resource.
        /// 
        /// The IAM principal configured for the provider must have `s3:PutObject`, `s3:HeadObject` and `s3:GetObject` permissions on this bucket.
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        [Input("customResourceProperties", required: true)]
        private InputMap<object>? _customResourceProperties;

        /// <summary>
        /// The properties to pass as an input to the Custom Resource.
        /// The properties are passed as a map of key-value pairs whereas all primitive values (number, boolean) are converted to strings for CloudFormation interoperability.
        /// </summary>
        public InputMap<object> CustomResourceProperties
        {
            get => _customResourceProperties ?? (_customResourceProperties = new InputMap<object>());
            set => _customResourceProperties = value;
        }

        /// <summary>
        /// The CloudFormation type of the Custom Resource. For example, `Custom::MyCustomResource`.
        /// This is required for CloudFormation interoperability.
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        /// <summary>
        /// The service token to use for the Custom Resource. The service token is invoked when the resource is created, updated, or deleted.
        /// This can be a Lambda Function ARN with optional version or alias identifiers.
        /// 
        /// The IAM principal configured for the provider must have `lambda:InvokeFunction` permissions on this service token.
        /// </summary>
        [Input("serviceToken", required: true)]
        public Input<string> ServiceToken { get; set; } = null!;

        /// <summary>
        /// A stand-in value for the CloudFormation stack ID. This is required for CloudFormation interoperability.
        /// If not provided, the Pulumi Stack ID is used.
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        public CustomResourceEmulatorArgs()
        {
        }
        public static new CustomResourceEmulatorArgs Empty => new CustomResourceEmulatorArgs();
    }
}
