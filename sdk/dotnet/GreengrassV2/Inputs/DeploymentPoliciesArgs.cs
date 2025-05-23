// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GreengrassV2.Inputs
{

    public sealed class DeploymentPoliciesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The component update policy for the configuration deployment. This policy defines when it's safe to deploy the configuration to devices.
        /// </summary>
        [Input("componentUpdatePolicy")]
        public Input<Inputs.DeploymentComponentUpdatePolicyArgs>? ComponentUpdatePolicy { get; set; }

        /// <summary>
        /// The configuration validation policy for the configuration deployment. This policy defines how long each component has to validate its configure updates.
        /// </summary>
        [Input("configurationValidationPolicy")]
        public Input<Inputs.DeploymentConfigurationValidationPolicyArgs>? ConfigurationValidationPolicy { get; set; }

        /// <summary>
        /// The failure handling policy for the configuration deployment. This policy defines what to do if the deployment fails.
        /// 
        /// Default: `ROLLBACK`
        /// </summary>
        [Input("failureHandlingPolicy")]
        public Input<Pulumi.AwsNative.GreengrassV2.DeploymentPoliciesFailureHandlingPolicy>? FailureHandlingPolicy { get; set; }

        public DeploymentPoliciesArgs()
        {
        }
        public static new DeploymentPoliciesArgs Empty => new DeploymentPoliciesArgs();
    }
}
