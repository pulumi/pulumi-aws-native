// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Iam.Inputs
{

    /// <summary>
    /// Contains information about an attached policy.
    /// </summary>
    public sealed class UserPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy document.
        /// </summary>
        [Input("policyDocument", required: true)]
        public Input<object> PolicyDocument { get; set; } = null!;

        /// <summary>
        /// The friendly name (not ARN) identifying the policy.
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        public UserPolicyArgs()
        {
        }
        public static new UserPolicyArgs Empty => new UserPolicyArgs();
    }
}