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
        [Input("accessUrl", required: true)]
        public Input<string> AccessUrl { get; set; } = null!;

        [Input("approvedOrigins", required: true)]
        private InputList<string>? _approvedOrigins;
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