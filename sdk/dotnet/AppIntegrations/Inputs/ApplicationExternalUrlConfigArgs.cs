// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppIntegrations.Inputs
{

    public sealed class ApplicationExternalUrlConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL to access the application.
        /// </summary>
        [Input("accessUrl", required: true)]
        public Input<string> AccessUrl { get; set; } = null!;

        [Input("approvedOrigins")]
        private InputList<string>? _approvedOrigins;

        /// <summary>
        /// Additional URLs to allow list if different than the access URL.
        /// </summary>
        public InputList<string> ApprovedOrigins
        {
            get => _approvedOrigins ?? (_approvedOrigins = new InputList<string>());
            set => _approvedOrigins = value;
        }

        public ApplicationExternalUrlConfigArgs()
        {
        }
        public static new ApplicationExternalUrlConfigArgs Empty => new ApplicationExternalUrlConfigArgs();
    }
}