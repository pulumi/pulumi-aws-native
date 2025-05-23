// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    /// <summary>
    /// A list of files containing the environment variables to pass to a container. You can specify up to ten environment files. The file must have a ``.env`` file extension. Each line in an environment file should contain an environment variable in ``VARIABLE=VALUE`` format. Lines beginning with ``#`` are treated as comments and are ignored.
    ///  If there are environment variables specified using the ``environment`` parameter in a container definition, they take precedence over the variables contained within an environment file. If multiple environment files are specified that contain the same variable, they're processed from the top down. We recommend that you use unique variable names. For more information, see [Use a file to pass environment variables to a container](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/use-environment-file.html) in the *Amazon Elastic Container Service Developer Guide*.
    ///  Environment variable files are objects in Amazon S3 and all Amazon S3 security considerations apply. 
    ///  You must use the following platforms for the Fargate launch type:
    ///   +  Linux platform version ``1.4.0`` or later.
    ///   +  Windows platform version ``1.0.0`` or later.
    ///   
    ///  Consider the following when using the Fargate launch type:
    ///   +  The file is handled like a native Docker env-file.
    ///   +  There is no support for shell escape handling.
    ///   +  The container entry point interperts the ``VARIABLE`` values.
    /// </summary>
    public sealed class TaskDefinitionEnvironmentFileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The file type to use. Environment files are objects in Amazon S3. The only supported value is ``s3``.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon S3 object containing the environment variable file.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public TaskDefinitionEnvironmentFileArgs()
        {
        }
        public static new TaskDefinitionEnvironmentFileArgs Empty => new TaskDefinitionEnvironmentFileArgs();
    }
}
