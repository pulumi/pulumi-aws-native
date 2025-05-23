// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RefactorSpaces.Inputs
{

    public sealed class RouteUriPathRouteInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to `ACTIVE` , traffic is forwarded to this route’s service after the route is created.
        /// </summary>
        [Input("activationState", required: true)]
        public Input<Pulumi.AwsNative.RefactorSpaces.RouteActivationState> ActivationState { get; set; } = null!;

        /// <summary>
        /// If set to `true` , this option appends the source path to the service URL endpoint.
        /// </summary>
        [Input("appendSourcePath")]
        public Input<bool>? AppendSourcePath { get; set; }

        /// <summary>
        /// Indicates whether to match all subpaths of the given source path. If this value is `false` , requests must match the source path exactly before they are forwarded to this route's service.
        /// </summary>
        [Input("includeChildPaths")]
        public Input<bool>? IncludeChildPaths { get; set; }

        [Input("methods")]
        private InputList<Pulumi.AwsNative.RefactorSpaces.RouteMethod>? _methods;

        /// <summary>
        /// A list of HTTP methods to match. An empty list matches all values. If a method is present, only HTTP requests using that method are forwarded to this route’s service.
        /// </summary>
        public InputList<Pulumi.AwsNative.RefactorSpaces.RouteMethod> Methods
        {
            get => _methods ?? (_methods = new InputList<Pulumi.AwsNative.RefactorSpaces.RouteMethod>());
            set => _methods = value;
        }

        /// <summary>
        /// This is the path that Refactor Spaces uses to match traffic. Paths must start with `/` and are relative to the base of the application. To use path parameters in the source path, add a variable in curly braces. For example, the resource path {user} represents a path parameter called 'user'.
        /// </summary>
        [Input("sourcePath")]
        public Input<string>? SourcePath { get; set; }

        public RouteUriPathRouteInputArgs()
        {
        }
        public static new RouteUriPathRouteInputArgs Empty => new RouteUriPathRouteInputArgs();
    }
}
