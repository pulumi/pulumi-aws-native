// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisSheetImageSource
    {
        /// <summary>
        /// The source of the static file that contains the image.
        /// </summary>
        public readonly Outputs.AnalysisSheetImageStaticFileSource? SheetImageStaticFileSource;

        [OutputConstructor]
        private AnalysisSheetImageSource(Outputs.AnalysisSheetImageStaticFileSource? sheetImageStaticFileSource)
        {
            SheetImageStaticFileSource = sheetImageStaticFileSource;
        }
    }
}
