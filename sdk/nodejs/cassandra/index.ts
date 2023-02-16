// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetKeyspaceArgs, GetKeyspaceResult, GetKeyspaceOutputArgs } from "./getKeyspace";
export const getKeyspace: typeof import("./getKeyspace").getKeyspace = null as any;
export const getKeyspaceOutput: typeof import("./getKeyspace").getKeyspaceOutput = null as any;
utilities.lazyLoad(exports, ["getKeyspace","getKeyspaceOutput"], () => require("./getKeyspace"));

export { GetTableArgs, GetTableResult, GetTableOutputArgs } from "./getTable";
export const getTable: typeof import("./getTable").getTable = null as any;
export const getTableOutput: typeof import("./getTable").getTableOutput = null as any;
utilities.lazyLoad(exports, ["getTable","getTableOutput"], () => require("./getTable"));

export { KeyspaceArgs } from "./keyspace";
export type Keyspace = import("./keyspace").Keyspace;
export const Keyspace: typeof import("./keyspace").Keyspace = null as any;
utilities.lazyLoad(exports, ["Keyspace"], () => require("./keyspace"));

export { TableArgs } from "./table";
export type Table = import("./table").Table;
export const Table: typeof import("./table").Table = null as any;
utilities.lazyLoad(exports, ["Table"], () => require("./table"));


// Export enums:
export * from "../types/enums/cassandra";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:cassandra:Keyspace":
                return new Keyspace(name, <any>undefined, { urn })
            case "aws-native:cassandra:Table":
                return new Table(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "cassandra", _module)