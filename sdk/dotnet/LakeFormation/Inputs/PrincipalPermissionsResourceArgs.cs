// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LakeFormation.Inputs
{

    public sealed class PrincipalPermissionsResourceArgs : global::Pulumi.ResourceArgs
    {
        [Input("catalog")]
        public Input<Inputs.PrincipalPermissionsCatalogResourceArgs>? Catalog { get; set; }

        [Input("dataCellsFilter")]
        public Input<Inputs.PrincipalPermissionsDataCellsFilterResourceArgs>? DataCellsFilter { get; set; }

        [Input("dataLocation")]
        public Input<Inputs.PrincipalPermissionsDataLocationResourceArgs>? DataLocation { get; set; }

        [Input("database")]
        public Input<Inputs.PrincipalPermissionsDatabaseResourceArgs>? Database { get; set; }

        [Input("lFTag")]
        public Input<Inputs.PrincipalPermissionsLFTagKeyResourceArgs>? LFTag { get; set; }

        [Input("lFTagPolicy")]
        public Input<Inputs.PrincipalPermissionsLFTagPolicyResourceArgs>? LFTagPolicy { get; set; }

        [Input("table")]
        public Input<Inputs.PrincipalPermissionsTableResourceArgs>? Table { get; set; }

        [Input("tableWithColumns")]
        public Input<Inputs.PrincipalPermissionsTableWithColumnsResourceArgs>? TableWithColumns { get; set; }

        public PrincipalPermissionsResourceArgs()
        {
        }
        public static new PrincipalPermissionsResourceArgs Empty => new PrincipalPermissionsResourceArgs();
    }
}