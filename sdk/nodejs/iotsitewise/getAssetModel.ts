// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::IoTSiteWise::AssetModel
 */
export function getAssetModel(args: GetAssetModelArgs, opts?: pulumi.InvokeOptions): Promise<GetAssetModelResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iotsitewise:getAssetModel", {
        "assetModelId": args.assetModelId,
    }, opts);
}

export interface GetAssetModelArgs {
    /**
     * The ID of the asset model.
     */
    assetModelId: string;
}

export interface GetAssetModelResult {
    /**
     * The ARN of the asset model, which has the following format.
     */
    readonly assetModelArn?: string;
    /**
     * The composite asset models that are part of this asset model. Composite asset models are asset models that contain specific properties.
     */
    readonly assetModelCompositeModels?: outputs.iotsitewise.AssetModelCompositeModel[];
    /**
     * A description for the asset model.
     */
    readonly assetModelDescription?: string;
    /**
     * The external ID of the asset model.
     */
    readonly assetModelExternalId?: string;
    /**
     * The hierarchy definitions of the asset model. Each hierarchy specifies an asset model whose assets can be children of any other assets created from this asset model. You can specify up to 10 hierarchies per asset model.
     */
    readonly assetModelHierarchies?: outputs.iotsitewise.AssetModelHierarchy[];
    /**
     * The ID of the asset model.
     */
    readonly assetModelId?: string;
    /**
     * A unique, friendly name for the asset model.
     */
    readonly assetModelName?: string;
    /**
     * The property definitions of the asset model. You can specify up to 200 properties per asset model.
     */
    readonly assetModelProperties?: outputs.iotsitewise.AssetModelProperty[];
    /**
     * A list of key-value pairs that contain metadata for the asset model.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource schema for AWS::IoTSiteWise::AssetModel
 */
export function getAssetModelOutput(args: GetAssetModelOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAssetModelResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:iotsitewise:getAssetModel", {
        "assetModelId": args.assetModelId,
    }, opts);
}

export interface GetAssetModelOutputArgs {
    /**
     * The ID of the asset model.
     */
    assetModelId: pulumi.Input<string>;
}
