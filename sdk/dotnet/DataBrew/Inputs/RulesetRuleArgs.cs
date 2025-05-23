// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Inputs
{

    /// <summary>
    /// Data quality rule for a target resource (dataset)
    /// </summary>
    public sealed class RulesetRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The expression which includes column references, condition names followed by variable references, possibly grouped and combined with other conditions. For example, `(:col1 starts_with :prefix1 or :col1 starts_with :prefix2) and (:col1 ends_with :suffix1 or :col1 ends_with :suffix2)` . Column and value references are substitution variables that should start with the ':' symbol. Depending on the context, substitution variables' values can be either an actual value or a column name. These values are defined in the SubstitutionMap. If a CheckExpression starts with a column reference, then ColumnSelectors in the rule should be null. If ColumnSelectors has been defined, then there should be no columnn reference in the left side of a condition, for example, `is_between :val1 and :val2` .
        /// </summary>
        [Input("checkExpression", required: true)]
        public Input<string> CheckExpression { get; set; } = null!;

        [Input("columnSelectors")]
        private InputList<Inputs.RulesetColumnSelectorArgs>? _columnSelectors;

        /// <summary>
        /// List of column selectors. Selectors can be used to select columns using a name or regular expression from the dataset. Rule will be applied to selected columns.
        /// </summary>
        public InputList<Inputs.RulesetColumnSelectorArgs> ColumnSelectors
        {
            get => _columnSelectors ?? (_columnSelectors = new InputList<Inputs.RulesetColumnSelectorArgs>());
            set => _columnSelectors = value;
        }

        /// <summary>
        /// A value that specifies whether the rule is disabled. Once a rule is disabled, a profile job will not validate it during a job run. Default value is false.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Name of the rule
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("substitutionMap")]
        private InputList<Inputs.RulesetSubstitutionValueArgs>? _substitutionMap;

        /// <summary>
        /// The map of substitution variable names to their values used in a check expression. Variable names should start with a ':' (colon). Variable values can either be actual values or column names. To differentiate between the two, column names should be enclosed in backticks, for example, `":col1": "`Column A`".`
        /// </summary>
        public InputList<Inputs.RulesetSubstitutionValueArgs> SubstitutionMap
        {
            get => _substitutionMap ?? (_substitutionMap = new InputList<Inputs.RulesetSubstitutionValueArgs>());
            set => _substitutionMap = value;
        }

        /// <summary>
        /// The threshold used with a non-aggregate check expression. Non-aggregate check expressions will be applied to each row in a specific column, and the threshold will be used to determine whether the validation succeeds.
        /// </summary>
        [Input("threshold")]
        public Input<Inputs.RulesetThresholdArgs>? Threshold { get; set; }

        public RulesetRuleArgs()
        {
        }
        public static new RulesetRuleArgs Empty => new RulesetRuleArgs();
    }
}
