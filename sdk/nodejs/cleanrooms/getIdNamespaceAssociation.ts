// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Represents an association between an ID namespace and a collaboration
 */
export function getIdNamespaceAssociation(args: GetIdNamespaceAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetIdNamespaceAssociationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:cleanrooms:getIdNamespaceAssociation", {
        "idNamespaceAssociationIdentifier": args.idNamespaceAssociationIdentifier,
        "membershipIdentifier": args.membershipIdentifier,
    }, opts);
}

export interface GetIdNamespaceAssociationArgs {
    /**
     * The unique identifier of the ID namespace association that you want to retrieve.
     */
    idNamespaceAssociationIdentifier: string;
    /**
     * The unique identifier of the membership that contains the ID namespace association.
     */
    membershipIdentifier: string;
}

export interface GetIdNamespaceAssociationResult {
    /**
     * The Amazon Resource Name (ARN) of the ID namespace association.
     */
    readonly arn?: string;
    /**
     * The Amazon Resource Name (ARN) of the collaboration that contains this ID namespace association.
     */
    readonly collaborationArn?: string;
    /**
     * The unique identifier of the collaboration that contains this ID namespace association.
     */
    readonly collaborationIdentifier?: string;
    /**
     * The description of the ID namespace association.
     */
    readonly description?: string;
    /**
     * The configuration settings for the ID mapping table.
     */
    readonly idMappingConfig?: outputs.cleanrooms.IdNamespaceAssociationIdMappingConfig;
    /**
     * The unique identifier of the ID namespace association that you want to retrieve.
     */
    readonly idNamespaceAssociationIdentifier?: string;
    readonly inputReferenceProperties?: outputs.cleanrooms.IdNamespaceAssociationInputReferenceProperties;
    /**
     * The Amazon Resource Name (ARN) of the membership resource for this ID namespace association.
     */
    readonly membershipArn?: string;
    /**
     * The name of this ID namespace association.
     */
    readonly name?: string;
    /**
     * An optional label that you can assign to a resource when you create it. Each tag consists of a key and an optional value, both of which you define. When you use tagging, you can also use tag-based access control in IAM policies to control access to this resource.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Represents an association between an ID namespace and a collaboration
 */
export function getIdNamespaceAssociationOutput(args: GetIdNamespaceAssociationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetIdNamespaceAssociationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:cleanrooms:getIdNamespaceAssociation", {
        "idNamespaceAssociationIdentifier": args.idNamespaceAssociationIdentifier,
        "membershipIdentifier": args.membershipIdentifier,
    }, opts);
}

export interface GetIdNamespaceAssociationOutputArgs {
    /**
     * The unique identifier of the ID namespace association that you want to retrieve.
     */
    idNamespaceAssociationIdentifier: pulumi.Input<string>;
    /**
     * The unique identifier of the membership that contains the ID namespace association.
     */
    membershipIdentifier: pulumi.Input<string>;
}
