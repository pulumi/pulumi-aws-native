// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LakeFormation.Outputs
{

    [OutputType]
    public sealed class TagAssociationResource
    {
        /// <summary>
        /// The identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your AWS Lake Formation environment.
        /// </summary>
        public readonly Outputs.TagAssociationCatalogResource? Catalog;
        /// <summary>
        /// The database for the resource. Unique to the Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database permissions to a principal.
        /// </summary>
        public readonly Outputs.TagAssociationDatabaseResource? Database;
        /// <summary>
        /// The table for the resource. A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
        /// </summary>
        public readonly Outputs.TagAssociationTableResource? Table;
        /// <summary>
        /// The table with columns for the resource. A principal with permissions to this resource can select metadata from the columns of a table in the Data Catalog and the underlying data in Amazon S3.
        /// </summary>
        public readonly Outputs.TagAssociationTableWithColumnsResource? TableWithColumns;

        [OutputConstructor]
        private TagAssociationResource(
            Outputs.TagAssociationCatalogResource? catalog,

            Outputs.TagAssociationDatabaseResource? database,

            Outputs.TagAssociationTableResource? table,

            Outputs.TagAssociationTableWithColumnsResource? tableWithColumns)
        {
            Catalog = catalog;
            Database = database;
            Table = table;
            TableWithColumns = tableWithColumns;
        }
    }
}
