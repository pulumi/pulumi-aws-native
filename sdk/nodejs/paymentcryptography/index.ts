// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AliasArgs } from "./alias";
export type Alias = import("./alias").Alias;
export const Alias: typeof import("./alias").Alias = null as any;
utilities.lazyLoad(exports, ["Alias"], () => require("./alias"));

export { GetAliasArgs, GetAliasResult, GetAliasOutputArgs } from "./getAlias";
export const getAlias: typeof import("./getAlias").getAlias = null as any;
export const getAliasOutput: typeof import("./getAlias").getAliasOutput = null as any;
utilities.lazyLoad(exports, ["getAlias","getAliasOutput"], () => require("./getAlias"));

export { GetKeyArgs, GetKeyResult, GetKeyOutputArgs } from "./getKey";
export const getKey: typeof import("./getKey").getKey = null as any;
export const getKeyOutput: typeof import("./getKey").getKeyOutput = null as any;
utilities.lazyLoad(exports, ["getKey","getKeyOutput"], () => require("./getKey"));

export { KeyArgs } from "./key";
export type Key = import("./key").Key;
export const Key: typeof import("./key").Key = null as any;
utilities.lazyLoad(exports, ["Key"], () => require("./key"));


// Export enums:
export * from "../types/enums/paymentcryptography";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:paymentcryptography:Alias":
                return new Alias(name, <any>undefined, { urn })
            case "aws-native:paymentcryptography:Key":
                return new Key(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "paymentcryptography", _module)
