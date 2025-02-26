// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Schema for AWS::CodeStarConnections::RepositoryLink resource which is used to aggregate repository metadata relevant to synchronizing source provider content to AWS Resources.
 */
export function getRepositoryLink(args: GetRepositoryLinkArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryLinkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:codestarconnections:getRepositoryLink", {
        "repositoryLinkArn": args.repositoryLinkArn,
    }, opts);
}

export interface GetRepositoryLinkArgs {
    /**
     * A unique Amazon Resource Name (ARN) to designate the repository link.
     */
    repositoryLinkArn: string;
}

export interface GetRepositoryLinkResult {
    /**
     * The Amazon Resource Name (ARN) of the CodeStarConnection. The ARN is used as the connection reference when the connection is shared between AWS services.
     */
    readonly connectionArn?: string;
    /**
     * The ARN of the KMS key that the customer can optionally specify to use to encrypt RepositoryLink properties. If not specified, a default key will be used.
     */
    readonly encryptionKeyArn?: string;
    /**
     * The name of the external provider where your third-party code repository is configured.
     */
    readonly providerType?: enums.codestarconnections.RepositoryLinkProviderType;
    /**
     * A unique Amazon Resource Name (ARN) to designate the repository link.
     */
    readonly repositoryLinkArn?: string;
    /**
     * A UUID that uniquely identifies the RepositoryLink.
     */
    readonly repositoryLinkId?: string;
    /**
     * Specifies the tags applied to a RepositoryLink.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Schema for AWS::CodeStarConnections::RepositoryLink resource which is used to aggregate repository metadata relevant to synchronizing source provider content to AWS Resources.
 */
export function getRepositoryLinkOutput(args: GetRepositoryLinkOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRepositoryLinkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:codestarconnections:getRepositoryLink", {
        "repositoryLinkArn": args.repositoryLinkArn,
    }, opts);
}

export interface GetRepositoryLinkOutputArgs {
    /**
     * A unique Amazon Resource Name (ARN) to designate the repository link.
     */
    repositoryLinkArn: pulumi.Input<string>;
}
