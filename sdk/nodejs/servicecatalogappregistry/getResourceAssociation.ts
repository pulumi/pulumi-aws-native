// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Schema for AWS::ServiceCatalogAppRegistry::ResourceAssociation
 */
export function getResourceAssociation(args: GetResourceAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceAssociationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:servicecatalogappregistry:getResourceAssociation", {
        "applicationArn": args.applicationArn,
        "resourceArn": args.resourceArn,
        "resourceType": args.resourceType,
    }, opts);
}

export interface GetResourceAssociationArgs {
    /**
     * The Amazon resource name (ARN) that specifies the application.
     */
    applicationArn: string;
    /**
     * The Amazon resource name (ARN) that specifies the resource.
     */
    resourceArn: string;
    /**
     * The type of the CFN Resource for now it's enum CFN_STACK.
     */
    resourceType: enums.servicecatalogappregistry.ResourceAssociationResourceType;
}

export interface GetResourceAssociationResult {
    /**
     * The Amazon resource name (ARN) that specifies the application.
     */
    readonly applicationArn?: string;
    /**
     * The Amazon resource name (ARN) that specifies the resource.
     */
    readonly resourceArn?: string;
}
/**
 * Resource Schema for AWS::ServiceCatalogAppRegistry::ResourceAssociation
 */
export function getResourceAssociationOutput(args: GetResourceAssociationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetResourceAssociationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:servicecatalogappregistry:getResourceAssociation", {
        "applicationArn": args.applicationArn,
        "resourceArn": args.resourceArn,
        "resourceType": args.resourceType,
    }, opts);
}

export interface GetResourceAssociationOutputArgs {
    /**
     * The Amazon resource name (ARN) that specifies the application.
     */
    applicationArn: pulumi.Input<string>;
    /**
     * The Amazon resource name (ARN) that specifies the resource.
     */
    resourceArn: pulumi.Input<string>;
    /**
     * The type of the CFN Resource for now it's enum CFN_STACK.
     */
    resourceType: pulumi.Input<enums.servicecatalogappregistry.ResourceAssociationResourceType>;
}
