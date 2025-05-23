// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Outputs
{

    [OutputType]
    public sealed class UserPoolClientAnalyticsConfiguration
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of an Amazon Pinpoint project that you want to connect to your user pool app client. Amazon Cognito publishes events to the Amazon Pinpoint project that `ApplicationArn` declares. You can also configure your application to pass an endpoint ID in the `AnalyticsMetadata` parameter of sign-in operations. The endpoint ID is information about the destination for push notifications
        /// </summary>
        public readonly string? ApplicationArn;
        /// <summary>
        /// Your Amazon Pinpoint project ID.
        /// </summary>
        public readonly string? ApplicationId;
        /// <summary>
        /// The [external ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html) of the role that Amazon Cognito assumes to send analytics data to Amazon Pinpoint.
        /// </summary>
        public readonly string? ExternalId;
        /// <summary>
        /// The ARN of an AWS Identity and Access Management role that has the permissions required for Amazon Cognito to publish events to Amazon Pinpoint analytics.
        /// </summary>
        public readonly string? RoleArn;
        /// <summary>
        /// If `UserDataShared` is `true` , Amazon Cognito includes user data in the events that it publishes to Amazon Pinpoint analytics.
        /// </summary>
        public readonly bool? UserDataShared;

        [OutputConstructor]
        private UserPoolClientAnalyticsConfiguration(
            string? applicationArn,

            string? applicationId,

            string? externalId,

            string? roleArn,

            bool? userDataShared)
        {
            ApplicationArn = applicationArn;
            ApplicationId = applicationId;
            ExternalId = externalId;
            RoleArn = roleArn;
            UserDataShared = userDataShared;
        }
    }
}
