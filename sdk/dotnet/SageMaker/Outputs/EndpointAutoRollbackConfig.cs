// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    [OutputType]
    public sealed class EndpointAutoRollbackConfig
    {
        /// <summary>
        /// List of CloudWatch alarms to monitor during the deployment. If any alarm goes off, the deployment is rolled back.
        /// </summary>
        public readonly ImmutableArray<Outputs.EndpointAlarm> Alarms;

        [OutputConstructor]
        private EndpointAutoRollbackConfig(ImmutableArray<Outputs.EndpointAlarm> alarms)
        {
            Alarms = alarms;
        }
    }
}
