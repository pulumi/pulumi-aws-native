// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ClusterArgs } from "./cluster";
export type Cluster = import("./cluster").Cluster;
export const Cluster: typeof import("./cluster").Cluster = null as any;
utilities.lazyLoad(exports, ["Cluster"], () => require("./cluster"));

export { GetClusterArgs, GetClusterResult, GetClusterOutputArgs } from "./getCluster";
export const getCluster: typeof import("./getCluster").getCluster = null as any;
export const getClusterOutput: typeof import("./getCluster").getClusterOutput = null as any;
utilities.lazyLoad(exports, ["getCluster","getClusterOutput"], () => require("./getCluster"));

export { GetParameterGroupArgs, GetParameterGroupResult, GetParameterGroupOutputArgs } from "./getParameterGroup";
export const getParameterGroup: typeof import("./getParameterGroup").getParameterGroup = null as any;
export const getParameterGroupOutput: typeof import("./getParameterGroup").getParameterGroupOutput = null as any;
utilities.lazyLoad(exports, ["getParameterGroup","getParameterGroupOutput"], () => require("./getParameterGroup"));

export { GetSubnetGroupArgs, GetSubnetGroupResult, GetSubnetGroupOutputArgs } from "./getSubnetGroup";
export const getSubnetGroup: typeof import("./getSubnetGroup").getSubnetGroup = null as any;
export const getSubnetGroupOutput: typeof import("./getSubnetGroup").getSubnetGroupOutput = null as any;
utilities.lazyLoad(exports, ["getSubnetGroup","getSubnetGroupOutput"], () => require("./getSubnetGroup"));

export { ParameterGroupArgs } from "./parameterGroup";
export type ParameterGroup = import("./parameterGroup").ParameterGroup;
export const ParameterGroup: typeof import("./parameterGroup").ParameterGroup = null as any;
utilities.lazyLoad(exports, ["ParameterGroup"], () => require("./parameterGroup"));

export { SubnetGroupArgs } from "./subnetGroup";
export type SubnetGroup = import("./subnetGroup").SubnetGroup;
export const SubnetGroup: typeof import("./subnetGroup").SubnetGroup = null as any;
utilities.lazyLoad(exports, ["SubnetGroup"], () => require("./subnetGroup"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:dax:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "aws-native:dax:ParameterGroup":
                return new ParameterGroup(name, <any>undefined, { urn })
            case "aws-native:dax:SubnetGroup":
                return new SubnetGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "dax", _module)