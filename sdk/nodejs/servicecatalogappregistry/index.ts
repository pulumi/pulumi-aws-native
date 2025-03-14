// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ApplicationArgs } from "./application";
export type Application = import("./application").Application;
export const Application: typeof import("./application").Application = null as any;
utilities.lazyLoad(exports, ["Application"], () => require("./application"));

export { AttributeGroupArgs } from "./attributeGroup";
export type AttributeGroup = import("./attributeGroup").AttributeGroup;
export const AttributeGroup: typeof import("./attributeGroup").AttributeGroup = null as any;
utilities.lazyLoad(exports, ["AttributeGroup"], () => require("./attributeGroup"));

export { AttributeGroupAssociationArgs } from "./attributeGroupAssociation";
export type AttributeGroupAssociation = import("./attributeGroupAssociation").AttributeGroupAssociation;
export const AttributeGroupAssociation: typeof import("./attributeGroupAssociation").AttributeGroupAssociation = null as any;
utilities.lazyLoad(exports, ["AttributeGroupAssociation"], () => require("./attributeGroupAssociation"));

export { GetApplicationArgs, GetApplicationResult, GetApplicationOutputArgs } from "./getApplication";
export const getApplication: typeof import("./getApplication").getApplication = null as any;
export const getApplicationOutput: typeof import("./getApplication").getApplicationOutput = null as any;
utilities.lazyLoad(exports, ["getApplication","getApplicationOutput"], () => require("./getApplication"));

export { GetAttributeGroupArgs, GetAttributeGroupResult, GetAttributeGroupOutputArgs } from "./getAttributeGroup";
export const getAttributeGroup: typeof import("./getAttributeGroup").getAttributeGroup = null as any;
export const getAttributeGroupOutput: typeof import("./getAttributeGroup").getAttributeGroupOutput = null as any;
utilities.lazyLoad(exports, ["getAttributeGroup","getAttributeGroupOutput"], () => require("./getAttributeGroup"));

export { GetAttributeGroupAssociationArgs, GetAttributeGroupAssociationResult, GetAttributeGroupAssociationOutputArgs } from "./getAttributeGroupAssociation";
export const getAttributeGroupAssociation: typeof import("./getAttributeGroupAssociation").getAttributeGroupAssociation = null as any;
export const getAttributeGroupAssociationOutput: typeof import("./getAttributeGroupAssociation").getAttributeGroupAssociationOutput = null as any;
utilities.lazyLoad(exports, ["getAttributeGroupAssociation","getAttributeGroupAssociationOutput"], () => require("./getAttributeGroupAssociation"));

export { GetResourceAssociationArgs, GetResourceAssociationResult, GetResourceAssociationOutputArgs } from "./getResourceAssociation";
export const getResourceAssociation: typeof import("./getResourceAssociation").getResourceAssociation = null as any;
export const getResourceAssociationOutput: typeof import("./getResourceAssociation").getResourceAssociationOutput = null as any;
utilities.lazyLoad(exports, ["getResourceAssociation","getResourceAssociationOutput"], () => require("./getResourceAssociation"));

export { ResourceAssociationArgs } from "./resourceAssociation";
export type ResourceAssociation = import("./resourceAssociation").ResourceAssociation;
export const ResourceAssociation: typeof import("./resourceAssociation").ResourceAssociation = null as any;
utilities.lazyLoad(exports, ["ResourceAssociation"], () => require("./resourceAssociation"));


// Export enums:
export * from "../types/enums/servicecatalogappregistry";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:servicecatalogappregistry:Application":
                return new Application(name, <any>undefined, { urn })
            case "aws-native:servicecatalogappregistry:AttributeGroup":
                return new AttributeGroup(name, <any>undefined, { urn })
            case "aws-native:servicecatalogappregistry:AttributeGroupAssociation":
                return new AttributeGroupAssociation(name, <any>undefined, { urn })
            case "aws-native:servicecatalogappregistry:ResourceAssociation":
                return new ResourceAssociation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "servicecatalogappregistry", _module)
