// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway.Outputs
{

    [OutputType]
    public sealed class DocumentationPartLocation
    {
        /// <summary>
        /// The HTTP verb of a method.
        /// </summary>
        public readonly string? Method;
        /// <summary>
        /// The name of the targeted API entity.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The URL path of the target.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// The HTTP status code of a response.
        /// </summary>
        public readonly string? StatusCode;
        /// <summary>
        /// The type of API entity that the documentation content applies to.
        /// </summary>
        public readonly Pulumi.AwsNative.ApiGateway.DocumentationPartLocationType? Type;

        [OutputConstructor]
        private DocumentationPartLocation(
            string? method,

            string? name,

            string? path,

            string? statusCode,

            Pulumi.AwsNative.ApiGateway.DocumentationPartLocationType? type)
        {
            Method = method;
            Name = name;
            Path = path;
            StatusCode = statusCode;
            Type = type;
        }
    }
}