// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Budgets.Outputs
{

    [OutputType]
    public sealed class BudgetAutoAdjustData
    {
        public readonly string AutoAdjustType;
        public readonly Outputs.BudgetHistoricalOptions? HistoricalOptions;

        [OutputConstructor]
        private BudgetAutoAdjustData(
            string autoAdjustType,

            Outputs.BudgetHistoricalOptions? historicalOptions)
        {
            AutoAdjustType = autoAdjustType;
            HistoricalOptions = historicalOptions;
        }
    }
}