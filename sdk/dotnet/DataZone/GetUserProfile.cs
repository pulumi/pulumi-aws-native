// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone
{
    public static class GetUserProfile
    {
        /// <summary>
        /// A user profile represents Amazon DataZone users. Amazon DataZone supports both IAM roles and SSO identities to interact with the Amazon DataZone Management Console and the data portal for different purposes. Domain administrators use IAM roles to perform the initial administrative domain-related work in the Amazon DataZone Management Console, including creating new Amazon DataZone domains, configuring metadata form types, and implementing policies. Data workers use their SSO corporate identities via Identity Center to log into the Amazon DataZone Data Portal and access projects where they have memberships.
        /// </summary>
        public static Task<GetUserProfileResult> InvokeAsync(GetUserProfileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserProfileResult>("aws-native:datazone:getUserProfile", args ?? new GetUserProfileArgs(), options.WithDefaults());

        /// <summary>
        /// A user profile represents Amazon DataZone users. Amazon DataZone supports both IAM roles and SSO identities to interact with the Amazon DataZone Management Console and the data portal for different purposes. Domain administrators use IAM roles to perform the initial administrative domain-related work in the Amazon DataZone Management Console, including creating new Amazon DataZone domains, configuring metadata form types, and implementing policies. Data workers use their SSO corporate identities via Identity Center to log into the Amazon DataZone Data Portal and access projects where they have memberships.
        /// </summary>
        public static Output<GetUserProfileResult> Invoke(GetUserProfileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserProfileResult>("aws-native:datazone:getUserProfile", args ?? new GetUserProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserProfileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier of the Amazon DataZone domain in which the user profile is created.
        /// </summary>
        [Input("domainId", required: true)]
        public string DomainId { get; set; } = null!;

        /// <summary>
        /// The ID of the Amazon DataZone user profile.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetUserProfileArgs()
        {
        }
        public static new GetUserProfileArgs Empty => new GetUserProfileArgs();
    }

    public sealed class GetUserProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier of the Amazon DataZone domain in which the user profile is created.
        /// </summary>
        [Input("domainId", required: true)]
        public Input<string> DomainId { get; set; } = null!;

        /// <summary>
        /// The ID of the Amazon DataZone user profile.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetUserProfileInvokeArgs()
        {
        }
        public static new GetUserProfileInvokeArgs Empty => new GetUserProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserProfileResult
    {
        public readonly Union<Outputs.UserProfileDetails0Properties, Outputs.UserProfileDetails1Properties>? Details;
        /// <summary>
        /// The identifier of the Amazon DataZone domain in which the user profile is created.
        /// </summary>
        public readonly string? DomainId;
        /// <summary>
        /// The ID of the Amazon DataZone user profile.
        /// </summary>
        public readonly string? Id;
        public readonly Pulumi.AwsNative.DataZone.UserProfileStatus? Status;
        public readonly Pulumi.AwsNative.DataZone.UserProfileType? Type;

        [OutputConstructor]
        private GetUserProfileResult(
            Union<Outputs.UserProfileDetails0Properties, Outputs.UserProfileDetails1Properties>? details,

            string? domainId,

            string? id,

            Pulumi.AwsNative.DataZone.UserProfileStatus? status,

            Pulumi.AwsNative.DataZone.UserProfileType? type)
        {
            Details = details;
            DomainId = domainId;
            Id = id;
            Status = status;
            Type = type;
        }
    }
}