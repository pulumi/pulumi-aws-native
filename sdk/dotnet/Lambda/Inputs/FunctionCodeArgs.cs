// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Inputs
{

    /// <summary>
    /// The [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) for a Lambda function. To deploy a function defined as a container image, you specify the location of a container image in the Amazon ECR registry. For a .zip file deployment package, you can specify the location of an object in Amazon S3. For Node.js and Python functions, you can specify the function code inline in the template.
    ///   When you specify source code inline for a Node.js function, the ``index`` file that CFN creates uses the extension ``.js``. This means that LAM treats the file as a CommonJS module. ES modules aren't supported for inline functions.
    ///   Changes to a deployment package in Amazon S3 or a container image in ECR are not detected automatically during stack updates. To update the function code, change the object key or version in the template.
    /// </summary>
    public sealed class FunctionCodeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// URI of a [container image](https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html) in the Amazon ECR registry.
        /// </summary>
        [Input("imageUri")]
        public Input<string>? ImageUri { get; set; }

        /// <summary>
        /// An Amazon S3 bucket in the same AWS-Region as your function. The bucket can be in a different AWS-account.
        /// </summary>
        [Input("s3Bucket")]
        public Input<string>? S3Bucket { get; set; }

        /// <summary>
        /// The Amazon S3 key of the deployment package.
        /// </summary>
        [Input("s3Key")]
        public Input<string>? S3Key { get; set; }

        /// <summary>
        /// For versioned objects, the version of the deployment package object to use.
        /// </summary>
        [Input("s3ObjectVersion")]
        public Input<string>? S3ObjectVersion { get; set; }

        /// <summary>
        /// The ARN of the KMSlong (KMS) customer managed key that's used to encrypt your function's .zip deployment package. If you don't provide a customer managed key, Lambda uses an [owned key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-cmk).
        /// </summary>
        [Input("sourceKmsKeyArn")]
        public Input<string>? SourceKmsKeyArn { get; set; }

        /// <summary>
        /// (Node.js and Python) The source code of your Lambda function. If you include your function source inline with this parameter, CFN places it in a file named ``index`` and zips it to create a [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html). This zip file cannot exceed 4MB. For the ``Handler`` property, the first part of the handler identifier must be ``index``. For example, ``index.handler``.
        ///   When you specify source code inline for a Node.js function, the ``index`` file that CFN creates uses the extension ``.js``. This means that LAM treats the file as a CommonJS module. ES modules aren't supported for inline functions.
        ///    For JSON, you must escape quotes and special characters such as newline (``\n``) with a backslash.
        ///  If you specify a function that interacts with an AWS CloudFormation custom resource, you don't have to write your own functions to send responses to the custom resource that invoked the function. AWS CloudFormation provides a response module ([cfn-response](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-lambda-function-code-cfnresponsemodule.html)) that simplifies sending responses. See [Using Lambda with CloudFormation](https://docs.aws.amazon.com/lambda/latest/dg/services-cloudformation.html) for details.
        /// </summary>
        [Input("zipFile")]
        public Input<string>? ZipFile { get; set; }

        public FunctionCodeArgs()
        {
        }
        public static new FunctionCodeArgs Empty => new FunctionCodeArgs();
    }
}
