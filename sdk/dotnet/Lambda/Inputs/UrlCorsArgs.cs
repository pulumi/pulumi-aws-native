// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Inputs
{

    public sealed class UrlCorsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether credentials are included in the CORS request.
        /// </summary>
        [Input("allowCredentials")]
        public Input<bool>? AllowCredentials { get; set; }

        [Input("allowHeaders")]
        private InputList<string>? _allowHeaders;

        /// <summary>
        /// Represents a collection of allowed headers.
        /// </summary>
        public InputList<string> AllowHeaders
        {
            get => _allowHeaders ?? (_allowHeaders = new InputList<string>());
            set => _allowHeaders = value;
        }

        [Input("allowMethods")]
        private InputList<Pulumi.AwsNative.Lambda.UrlAllowMethodsItem>? _allowMethods;

        /// <summary>
        /// Represents a collection of allowed HTTP methods.
        /// </summary>
        public InputList<Pulumi.AwsNative.Lambda.UrlAllowMethodsItem> AllowMethods
        {
            get => _allowMethods ?? (_allowMethods = new InputList<Pulumi.AwsNative.Lambda.UrlAllowMethodsItem>());
            set => _allowMethods = value;
        }

        [Input("allowOrigins")]
        private InputList<string>? _allowOrigins;

        /// <summary>
        /// Represents a collection of allowed origins.
        /// </summary>
        public InputList<string> AllowOrigins
        {
            get => _allowOrigins ?? (_allowOrigins = new InputList<string>());
            set => _allowOrigins = value;
        }

        [Input("exposeHeaders")]
        private InputList<string>? _exposeHeaders;

        /// <summary>
        /// Represents a collection of exposed headers.
        /// </summary>
        public InputList<string> ExposeHeaders
        {
            get => _exposeHeaders ?? (_exposeHeaders = new InputList<string>());
            set => _exposeHeaders = value;
        }

        [Input("maxAge")]
        public Input<int>? MaxAge { get; set; }

        public UrlCorsArgs()
        {
        }
        public static new UrlCorsArgs Empty => new UrlCorsArgs();
    }
}