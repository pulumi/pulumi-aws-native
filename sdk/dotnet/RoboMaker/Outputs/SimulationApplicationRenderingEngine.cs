// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RoboMaker.Outputs
{

    /// <summary>
    /// Information about a rendering engine.
    /// </summary>
    [OutputType]
    public sealed class SimulationApplicationRenderingEngine
    {
        /// <summary>
        /// The name of the rendering engine.
        /// </summary>
        public readonly Pulumi.AwsNative.RoboMaker.SimulationApplicationRenderingEngineName Name;
        /// <summary>
        /// The version of the rendering engine.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private SimulationApplicationRenderingEngine(
            Pulumi.AwsNative.RoboMaker.SimulationApplicationRenderingEngineName name,

            string version)
        {
            Name = name;
            Version = version;
        }
    }
}