// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Deadline.Outputs
{

    [OutputType]
    public sealed class FleetConfiguration0Properties
    {
        public readonly Outputs.FleetCustomerManagedFleetConfiguration CustomerManaged;

        [OutputConstructor]
        private FleetConfiguration0Properties(Outputs.FleetCustomerManagedFleetConfiguration customerManaged)
        {
            CustomerManaged = customerManaged;
        }
    }
}
