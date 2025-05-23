// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AuditManager.Inputs
{

    /// <summary>
    /// The wrapper that contains AWS Audit Manager role information, such as the role type and IAM ARN.
    /// </summary>
    public sealed class AssessmentRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// The type of customer persona.
        /// 
        /// &gt; In `CreateAssessment` , `roleType` can only be `PROCESS_OWNER` .
        /// &gt; 
        /// &gt; In `UpdateSettings` , `roleType` can only be `PROCESS_OWNER` .
        /// &gt; 
        /// &gt; In `BatchCreateDelegationByAssessment` , `roleType` can only be `RESOURCE_OWNER` .
        /// </summary>
        [Input("roleType")]
        public Input<Pulumi.AwsNative.AuditManager.AssessmentRoleType>? RoleType { get; set; }

        public AssessmentRoleArgs()
        {
        }
        public static new AssessmentRoleArgs Empty => new AssessmentRoleArgs();
    }
}
