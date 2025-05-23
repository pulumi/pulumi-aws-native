// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub
{
    public static class GetDelegatedAdmin
    {
        /// <summary>
        /// The ``AWS::SecurityHub::DelegatedAdmin`` resource designates the delegated ASHlong administrator account for an organization. You must enable the integration between ASH and AOlong before you can designate a delegated ASH administrator. Only the management account for an organization can designate the delegated ASH administrator account. For more information, see [Designating the delegated administrator](https://docs.aws.amazon.com/securityhub/latest/userguide/designate-orgs-admin-account.html#designate-admin-instructions) in the *User Guide*.
        ///  To change the delegated administrator account, remove the current delegated administrator account, and then designate the new account.
        ///  To designate multiple delegated administrators in different organizations and AWS-Regions, we recommend using [mappings](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/mappings-section-structure.html).
        ///  Tags aren't supported for this resource.
        /// </summary>
        public static Task<GetDelegatedAdminResult> InvokeAsync(GetDelegatedAdminArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDelegatedAdminResult>("aws-native:securityhub:getDelegatedAdmin", args ?? new GetDelegatedAdminArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::SecurityHub::DelegatedAdmin`` resource designates the delegated ASHlong administrator account for an organization. You must enable the integration between ASH and AOlong before you can designate a delegated ASH administrator. Only the management account for an organization can designate the delegated ASH administrator account. For more information, see [Designating the delegated administrator](https://docs.aws.amazon.com/securityhub/latest/userguide/designate-orgs-admin-account.html#designate-admin-instructions) in the *User Guide*.
        ///  To change the delegated administrator account, remove the current delegated administrator account, and then designate the new account.
        ///  To designate multiple delegated administrators in different organizations and AWS-Regions, we recommend using [mappings](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/mappings-section-structure.html).
        ///  Tags aren't supported for this resource.
        /// </summary>
        public static Output<GetDelegatedAdminResult> Invoke(GetDelegatedAdminInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDelegatedAdminResult>("aws-native:securityhub:getDelegatedAdmin", args ?? new GetDelegatedAdminInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::SecurityHub::DelegatedAdmin`` resource designates the delegated ASHlong administrator account for an organization. You must enable the integration between ASH and AOlong before you can designate a delegated ASH administrator. Only the management account for an organization can designate the delegated ASH administrator account. For more information, see [Designating the delegated administrator](https://docs.aws.amazon.com/securityhub/latest/userguide/designate-orgs-admin-account.html#designate-admin-instructions) in the *User Guide*.
        ///  To change the delegated administrator account, remove the current delegated administrator account, and then designate the new account.
        ///  To designate multiple delegated administrators in different organizations and AWS-Regions, we recommend using [mappings](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/mappings-section-structure.html).
        ///  Tags aren't supported for this resource.
        /// </summary>
        public static Output<GetDelegatedAdminResult> Invoke(GetDelegatedAdminInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDelegatedAdminResult>("aws-native:securityhub:getDelegatedAdmin", args ?? new GetDelegatedAdminInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDelegatedAdminArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the delegated Security Hub administrator account, in the format of `accountID/Region` .
        /// </summary>
        [Input("delegatedAdminIdentifier", required: true)]
        public string DelegatedAdminIdentifier { get; set; } = null!;

        public GetDelegatedAdminArgs()
        {
        }
        public static new GetDelegatedAdminArgs Empty => new GetDelegatedAdminArgs();
    }

    public sealed class GetDelegatedAdminInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the delegated Security Hub administrator account, in the format of `accountID/Region` .
        /// </summary>
        [Input("delegatedAdminIdentifier", required: true)]
        public Input<string> DelegatedAdminIdentifier { get; set; } = null!;

        public GetDelegatedAdminInvokeArgs()
        {
        }
        public static new GetDelegatedAdminInvokeArgs Empty => new GetDelegatedAdminInvokeArgs();
    }


    [OutputType]
    public sealed class GetDelegatedAdminResult
    {
        /// <summary>
        /// The ID of the delegated Security Hub administrator account, in the format of `accountID/Region` .
        /// </summary>
        public readonly string? DelegatedAdminIdentifier;
        /// <summary>
        /// Whether the delegated Security Hub administrator is set for the organization.
        /// </summary>
        public readonly Pulumi.AwsNative.SecurityHub.DelegatedAdminStatus? Status;

        [OutputConstructor]
        private GetDelegatedAdminResult(
            string? delegatedAdminIdentifier,

            Pulumi.AwsNative.SecurityHub.DelegatedAdminStatus? status)
        {
            DelegatedAdminIdentifier = delegatedAdminIdentifier;
            Status = status;
        }
    }
}
