// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QBusiness.Inputs
{

    public sealed class ApplicationAutoSubscriptionConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes whether automatic subscriptions are enabled for an Amazon Q Business application using IAM identity federation for user management.
        /// </summary>
        [Input("autoSubscribe", required: true)]
        public Input<Pulumi.AwsNative.QBusiness.ApplicationAutoSubscriptionStatus> AutoSubscribe { get; set; } = null!;

        /// <summary>
        /// Describes the default subscription type assigned to an Amazon Q Business application using IAM identity federation for user management. If the value for `autoSubscribe` is set to `ENABLED` you must select a value for this field.
        /// </summary>
        [Input("defaultSubscriptionType")]
        public Input<Pulumi.AwsNative.QBusiness.ApplicationSubscriptionType>? DefaultSubscriptionType { get; set; }

        public ApplicationAutoSubscriptionConfigurationArgs()
        {
        }
        public static new ApplicationAutoSubscriptionConfigurationArgs Empty => new ApplicationAutoSubscriptionConfigurationArgs();
    }
}
