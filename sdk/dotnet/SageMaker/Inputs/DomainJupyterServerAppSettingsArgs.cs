// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// The JupyterServer app settings.
    /// </summary>
    public sealed class DomainJupyterServerAppSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the JupyterServer app.
        /// </summary>
        [Input("defaultResourceSpec")]
        public Input<Inputs.DomainResourceSpecArgs>? DefaultResourceSpec { get; set; }

        public DomainJupyterServerAppSettingsArgs()
        {
        }
        public static new DomainJupyterServerAppSettingsArgs Empty => new DomainJupyterServerAppSettingsArgs();
    }
}
