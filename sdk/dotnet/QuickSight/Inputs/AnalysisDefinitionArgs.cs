// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisDefinitionArgs : global::Pulumi.ResourceArgs
    {
        [Input("analysisDefaults")]
        public Input<Inputs.AnalysisDefaultsArgs>? AnalysisDefaults { get; set; }

        [Input("calculatedFields")]
        private InputList<Inputs.AnalysisCalculatedFieldArgs>? _calculatedFields;

        /// <summary>
        /// An array of calculated field definitions for the analysis.
        /// </summary>
        public InputList<Inputs.AnalysisCalculatedFieldArgs> CalculatedFields
        {
            get => _calculatedFields ?? (_calculatedFields = new InputList<Inputs.AnalysisCalculatedFieldArgs>());
            set => _calculatedFields = value;
        }

        [Input("columnConfigurations")]
        private InputList<Inputs.AnalysisColumnConfigurationArgs>? _columnConfigurations;

        /// <summary>
        /// An array of analysis-level column configurations. Column configurations can be used to set default formatting for a column to be used throughout an analysis.
        /// </summary>
        public InputList<Inputs.AnalysisColumnConfigurationArgs> ColumnConfigurations
        {
            get => _columnConfigurations ?? (_columnConfigurations = new InputList<Inputs.AnalysisColumnConfigurationArgs>());
            set => _columnConfigurations = value;
        }

        [Input("dataSetIdentifierDeclarations", required: true)]
        private InputList<Inputs.AnalysisDataSetIdentifierDeclarationArgs>? _dataSetIdentifierDeclarations;

        /// <summary>
        /// An array of dataset identifier declarations. This mapping allows the usage of dataset identifiers instead of dataset ARNs throughout analysis sub-structures.
        /// </summary>
        public InputList<Inputs.AnalysisDataSetIdentifierDeclarationArgs> DataSetIdentifierDeclarations
        {
            get => _dataSetIdentifierDeclarations ?? (_dataSetIdentifierDeclarations = new InputList<Inputs.AnalysisDataSetIdentifierDeclarationArgs>());
            set => _dataSetIdentifierDeclarations = value;
        }

        [Input("filterGroups")]
        private InputList<Inputs.AnalysisFilterGroupArgs>? _filterGroups;

        /// <summary>
        /// Filter definitions for an analysis.
        /// 
        /// For more information, see [Filtering Data in Amazon QuickSight](https://docs.aws.amazon.com/quicksight/latest/user/adding-a-filter.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public InputList<Inputs.AnalysisFilterGroupArgs> FilterGroups
        {
            get => _filterGroups ?? (_filterGroups = new InputList<Inputs.AnalysisFilterGroupArgs>());
            set => _filterGroups = value;
        }

        /// <summary>
        /// An array of option definitions for an analysis.
        /// </summary>
        [Input("options")]
        public Input<Inputs.AnalysisAssetOptionsArgs>? Options { get; set; }

        [Input("parameterDeclarations")]
        private InputList<Inputs.AnalysisParameterDeclarationArgs>? _parameterDeclarations;

        /// <summary>
        /// An array of parameter declarations for an analysis.
        /// 
        /// Parameters are named variables that can transfer a value for use by an action or an object.
        /// 
        /// For more information, see [Parameters in Amazon QuickSight](https://docs.aws.amazon.com/quicksight/latest/user/parameters-in-quicksight.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public InputList<Inputs.AnalysisParameterDeclarationArgs> ParameterDeclarations
        {
            get => _parameterDeclarations ?? (_parameterDeclarations = new InputList<Inputs.AnalysisParameterDeclarationArgs>());
            set => _parameterDeclarations = value;
        }

        [Input("queryExecutionOptions")]
        public Input<Inputs.AnalysisQueryExecutionOptionsArgs>? QueryExecutionOptions { get; set; }

        [Input("sheets")]
        private InputList<Inputs.AnalysisSheetDefinitionArgs>? _sheets;

        /// <summary>
        /// An array of sheet definitions for an analysis. Each `SheetDefinition` provides detailed information about a sheet within this analysis.
        /// </summary>
        public InputList<Inputs.AnalysisSheetDefinitionArgs> Sheets
        {
            get => _sheets ?? (_sheets = new InputList<Inputs.AnalysisSheetDefinitionArgs>());
            set => _sheets = value;
        }

        [Input("staticFiles")]
        private InputList<Inputs.AnalysisStaticFileArgs>? _staticFiles;

        /// <summary>
        /// The static files for the definition.
        /// </summary>
        public InputList<Inputs.AnalysisStaticFileArgs> StaticFiles
        {
            get => _staticFiles ?? (_staticFiles = new InputList<Inputs.AnalysisStaticFileArgs>());
            set => _staticFiles = value;
        }

        public AnalysisDefinitionArgs()
        {
        }
        public static new AnalysisDefinitionArgs Empty => new AnalysisDefinitionArgs();
    }
}
