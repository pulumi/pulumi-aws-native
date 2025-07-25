// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Logs.Outputs
{

    /// <summary>
    /// Individual processor configuration
    /// </summary>
    [OutputType]
    public sealed class TransformerProcessor
    {
        /// <summary>
        /// Use this parameter to include the [addKeys](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-addKeys) processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerProcessorAddKeysProperties? AddKeys;
        /// <summary>
        /// Use this parameter to include the [copyValue](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-copyValue) processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerProcessorCopyValueProperties? CopyValue;
        /// <summary>
        /// Use this parameter to include the [CSV](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-CSV) processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerProcessorCsvProperties? Csv;
        /// <summary>
        /// Use this parameter to include the [datetimeConverter](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-datetimeConverter) processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerProcessorDateTimeConverterProperties? DateTimeConverter;
        /// <summary>
        /// Use this parameter to include the [deleteKeys](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-deleteKeys) processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerProcessorDeleteKeysProperties? DeleteKeys;
        /// <summary>
        /// Use this parameter to include the [grok](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-grok) processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerProcessorGrokProperties? Grok;
        /// <summary>
        /// Use this parameter to include the [listToMap](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-listToMap) processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerProcessorListToMapProperties? ListToMap;
        /// <summary>
        /// Use this parameter to include the [lowerCaseString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-lowerCaseString) processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerProcessorLowerCaseStringProperties? LowerCaseString;
        /// <summary>
        /// Use this parameter to include the [moveKeys](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-moveKeys) processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerProcessorMoveKeysProperties? MoveKeys;
        /// <summary>
        /// Use this parameter to include the [parseCloudfront](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parseCloudfront) processor in your transformer.
        /// 
        /// If you use this processor, it must be the first processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerParseCloudfront? ParseCloudfront;
        /// <summary>
        /// Use this parameter to include the [parseJSON](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parseJSON) processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerProcessorParseJsonProperties? ParseJson;
        /// <summary>
        /// Use this parameter to include the [parseKeyValue](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parseKeyValue) processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerProcessorParseKeyValueProperties? ParseKeyValue;
        /// <summary>
        /// Use this parameter to include the [parsePostGres](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-parsePostGres) processor in your transformer.
        /// 
        /// If you use this processor, it must be the first processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerParsePostgres? ParsePostgres;
        /// <summary>
        /// Use this parameter to include the [parseRoute53](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parseRoute53) processor in your transformer.
        /// 
        /// If you use this processor, it must be the first processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerParseRoute53? ParseRoute53;
        /// <summary>
        /// Use this parameter to convert logs into Open Cybersecurity Schema (OCSF) format.
        /// </summary>
        public readonly Outputs.TransformerParseToOcsf? ParseToOcsf;
        /// <summary>
        /// Use this parameter to include the [parseVPC](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parseVPC) processor in your transformer.
        /// 
        /// If you use this processor, it must be the first processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerParseVpc? ParseVpc;
        /// <summary>
        /// Use this parameter to include the [parseWAF](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-parseWAF) processor in your transformer.
        /// 
        /// If you use this processor, it must be the first processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerParseWaf? ParseWaf;
        /// <summary>
        /// Use this parameter to include the [renameKeys](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-renameKeys) processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerProcessorRenameKeysProperties? RenameKeys;
        /// <summary>
        /// Use this parameter to include the [splitString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-splitString) processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerProcessorSplitStringProperties? SplitString;
        /// <summary>
        /// Use this parameter to include the [substituteString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-substituteString) processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerProcessorSubstituteStringProperties? SubstituteString;
        /// <summary>
        /// Use this parameter to include the [trimString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-trimString) processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerProcessorTrimStringProperties? TrimString;
        /// <summary>
        /// Use this parameter to include the [typeConverter](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-typeConverter) processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerProcessorTypeConverterProperties? TypeConverter;
        /// <summary>
        /// Use this parameter to include the [upperCaseString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-upperCaseString) processor in your transformer.
        /// </summary>
        public readonly Outputs.TransformerProcessorUpperCaseStringProperties? UpperCaseString;

        [OutputConstructor]
        private TransformerProcessor(
            Outputs.TransformerProcessorAddKeysProperties? addKeys,

            Outputs.TransformerProcessorCopyValueProperties? copyValue,

            Outputs.TransformerProcessorCsvProperties? csv,

            Outputs.TransformerProcessorDateTimeConverterProperties? dateTimeConverter,

            Outputs.TransformerProcessorDeleteKeysProperties? deleteKeys,

            Outputs.TransformerProcessorGrokProperties? grok,

            Outputs.TransformerProcessorListToMapProperties? listToMap,

            Outputs.TransformerProcessorLowerCaseStringProperties? lowerCaseString,

            Outputs.TransformerProcessorMoveKeysProperties? moveKeys,

            Outputs.TransformerParseCloudfront? parseCloudfront,

            Outputs.TransformerProcessorParseJsonProperties? parseJson,

            Outputs.TransformerProcessorParseKeyValueProperties? parseKeyValue,

            Outputs.TransformerParsePostgres? parsePostgres,

            Outputs.TransformerParseRoute53? parseRoute53,

            Outputs.TransformerParseToOcsf? parseToOcsf,

            Outputs.TransformerParseVpc? parseVpc,

            Outputs.TransformerParseWaf? parseWaf,

            Outputs.TransformerProcessorRenameKeysProperties? renameKeys,

            Outputs.TransformerProcessorSplitStringProperties? splitString,

            Outputs.TransformerProcessorSubstituteStringProperties? substituteString,

            Outputs.TransformerProcessorTrimStringProperties? trimString,

            Outputs.TransformerProcessorTypeConverterProperties? typeConverter,

            Outputs.TransformerProcessorUpperCaseStringProperties? upperCaseString)
        {
            AddKeys = addKeys;
            CopyValue = copyValue;
            Csv = csv;
            DateTimeConverter = dateTimeConverter;
            DeleteKeys = deleteKeys;
            Grok = grok;
            ListToMap = listToMap;
            LowerCaseString = lowerCaseString;
            MoveKeys = moveKeys;
            ParseCloudfront = parseCloudfront;
            ParseJson = parseJson;
            ParseKeyValue = parseKeyValue;
            ParsePostgres = parsePostgres;
            ParseRoute53 = parseRoute53;
            ParseToOcsf = parseToOcsf;
            ParseVpc = parseVpc;
            ParseWaf = parseWaf;
            RenameKeys = renameKeys;
            SplitString = splitString;
            SubstituteString = substituteString;
            TrimString = trimString;
            TypeConverter = typeConverter;
            UpperCaseString = upperCaseString;
        }
    }
}
