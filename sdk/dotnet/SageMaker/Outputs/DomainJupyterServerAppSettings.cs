// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// The JupyterServer app settings.
    /// </summary>
    [OutputType]
    public sealed class DomainJupyterServerAppSettings
    {
        /// <summary>
        /// The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the JupyterServer app.
        /// </summary>
        public readonly Outputs.DomainResourceSpec? DefaultResourceSpec;
        /// <summary>
        /// A list of LifecycleConfigArns available for use with JupyterServer apps.
        /// </summary>
        public readonly ImmutableArray<string> LifecycleConfigArns;

        [OutputConstructor]
        private DomainJupyterServerAppSettings(
            Outputs.DomainResourceSpec? defaultResourceSpec,

            ImmutableArray<string> lifecycleConfigArns)
        {
            DefaultResourceSpec = defaultResourceSpec;
            LifecycleConfigArns = lifecycleConfigArns;
        }
    }
}
