// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardParameterDeclaration
    {
        public readonly Outputs.DashboardDateTimeParameterDeclaration? DateTimeParameterDeclaration;
        public readonly Outputs.DashboardDecimalParameterDeclaration? DecimalParameterDeclaration;
        public readonly Outputs.DashboardIntegerParameterDeclaration? IntegerParameterDeclaration;
        public readonly Outputs.DashboardStringParameterDeclaration? StringParameterDeclaration;

        [OutputConstructor]
        private DashboardParameterDeclaration(
            Outputs.DashboardDateTimeParameterDeclaration? dateTimeParameterDeclaration,

            Outputs.DashboardDecimalParameterDeclaration? decimalParameterDeclaration,

            Outputs.DashboardIntegerParameterDeclaration? integerParameterDeclaration,

            Outputs.DashboardStringParameterDeclaration? stringParameterDeclaration)
        {
            DateTimeParameterDeclaration = dateTimeParameterDeclaration;
            DecimalParameterDeclaration = decimalParameterDeclaration;
            IntegerParameterDeclaration = integerParameterDeclaration;
            StringParameterDeclaration = stringParameterDeclaration;
        }
    }
}