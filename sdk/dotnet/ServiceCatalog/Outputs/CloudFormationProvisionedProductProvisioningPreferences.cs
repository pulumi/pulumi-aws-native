// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceCatalog.Outputs
{

    [OutputType]
    public sealed class CloudFormationProvisionedProductProvisioningPreferences
    {
        /// <summary>
        /// One or more AWS accounts where the provisioned product will be available.
        /// 
        /// Applicable only to a `CFN_STACKSET` provisioned product type.
        /// 
        /// The specified accounts should be within the list of accounts from the `STACKSET` constraint. To get the list of accounts in the `STACKSET` constraint, use the `DescribeProvisioningParameters` operation.
        /// 
        /// If no values are specified, the default value is all acounts from the `STACKSET` constraint.
        /// </summary>
        public readonly ImmutableArray<string> StackSetAccounts;
        /// <summary>
        /// The number of accounts, per Region, for which this operation can fail before AWS Service Catalog stops the operation in that Region. If the operation is stopped in a Region, AWS Service Catalog doesn't attempt the operation in any subsequent Regions.
        /// 
        /// Applicable only to a `CFN_STACKSET` provisioned product type.
        /// 
        /// Conditional: You must specify either `StackSetFailureToleranceCount` or `StackSetFailureTolerancePercentage` , but not both.
        /// 
        /// The default value is `0` if no value is specified.
        /// </summary>
        public readonly int? StackSetFailureToleranceCount;
        /// <summary>
        /// The percentage of accounts, per Region, for which this stack operation can fail before AWS Service Catalog stops the operation in that Region. If the operation is stopped in a Region, AWS Service Catalog doesn't attempt the operation in any subsequent Regions.
        /// 
        /// When calculating the number of accounts based on the specified percentage, AWS Service Catalog rounds down to the next whole number.
        /// 
        /// Applicable only to a `CFN_STACKSET` provisioned product type.
        /// 
        /// Conditional: You must specify either `StackSetFailureToleranceCount` or `StackSetFailureTolerancePercentage` , but not both.
        /// </summary>
        public readonly int? StackSetFailureTolerancePercentage;
        /// <summary>
        /// The maximum number of accounts in which to perform this operation at one time. This is dependent on the value of `StackSetFailureToleranceCount` . `StackSetMaxConcurrentCount` is at most one more than the `StackSetFailureToleranceCount` .
        /// 
        /// Note that this setting lets you specify the maximum for operations. For large deployments, under certain circumstances the actual number of accounts acted upon concurrently may be lower due to service throttling.
        /// 
        /// Applicable only to a `CFN_STACKSET` provisioned product type.
        /// 
        /// Conditional: You must specify either `StackSetMaxConcurrentCount` or `StackSetMaxConcurrentPercentage` , but not both.
        /// </summary>
        public readonly int? StackSetMaxConcurrencyCount;
        /// <summary>
        /// The maximum percentage of accounts in which to perform this operation at one time.
        /// 
        /// When calculating the number of accounts based on the specified percentage, AWS Service Catalog rounds down to the next whole number. This is true except in cases where rounding down would result is zero. In this case, AWS Service Catalog sets the number as `1` instead.
        /// 
        /// Note that this setting lets you specify the maximum for operations. For large deployments, under certain circumstances the actual number of accounts acted upon concurrently may be lower due to service throttling.
        /// 
        /// Applicable only to a `CFN_STACKSET` provisioned product type.
        /// 
        /// Conditional: You must specify either `StackSetMaxConcurrentCount` or `StackSetMaxConcurrentPercentage` , but not both.
        /// </summary>
        public readonly int? StackSetMaxConcurrencyPercentage;
        /// <summary>
        /// Determines what action AWS Service Catalog performs to a stack set or a stack instance represented by the provisioned product. The default value is `UPDATE` if nothing is specified.
        /// 
        /// Applicable only to a `CFN_STACKSET` provisioned product type.
        /// 
        /// - **CREATE** - Creates a new stack instance in the stack set represented by the provisioned product. In this case, only new stack instances are created based on accounts and Regions; if new ProductId or ProvisioningArtifactID are passed, they will be ignored.
        /// - **UPDATE** - Updates the stack set represented by the provisioned product and also its stack instances.
        /// - **DELETE** - Deletes a stack instance in the stack set represented by the provisioned product.
        /// </summary>
        public readonly Pulumi.AwsNative.ServiceCatalog.CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType? StackSetOperationType;
        /// <summary>
        /// One or more AWS Regions where the provisioned product will be available.
        /// 
        /// Applicable only to a `CFN_STACKSET` provisioned product type.
        /// 
        /// The specified Regions should be within the list of Regions from the `STACKSET` constraint. To get the list of Regions in the `STACKSET` constraint, use the `DescribeProvisioningParameters` operation.
        /// 
        /// If no values are specified, the default value is all Regions from the `STACKSET` constraint.
        /// </summary>
        public readonly ImmutableArray<string> StackSetRegions;

        [OutputConstructor]
        private CloudFormationProvisionedProductProvisioningPreferences(
            ImmutableArray<string> stackSetAccounts,

            int? stackSetFailureToleranceCount,

            int? stackSetFailureTolerancePercentage,

            int? stackSetMaxConcurrencyCount,

            int? stackSetMaxConcurrencyPercentage,

            Pulumi.AwsNative.ServiceCatalog.CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType? stackSetOperationType,

            ImmutableArray<string> stackSetRegions)
        {
            StackSetAccounts = stackSetAccounts;
            StackSetFailureToleranceCount = stackSetFailureToleranceCount;
            StackSetFailureTolerancePercentage = stackSetFailureTolerancePercentage;
            StackSetMaxConcurrencyCount = stackSetMaxConcurrencyCount;
            StackSetMaxConcurrencyPercentage = stackSetMaxConcurrencyPercentage;
            StackSetOperationType = stackSetOperationType;
            StackSetRegions = stackSetRegions;
        }
    }
}
