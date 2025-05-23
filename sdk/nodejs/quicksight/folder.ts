// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of the AWS::QuickSight::Folder Resource Type.
 */
export class Folder extends pulumi.CustomResource {
    /**
     * Get an existing Folder resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Folder {
        return new Folder(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:quicksight:Folder';

    /**
     * Returns true if the given object is an instance of Folder.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Folder {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Folder.__pulumiType;
    }

    /**
     * <p>The Amazon Resource Name (ARN) for the folder.</p>
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The ID for the AWS account where you want to create the folder.
     */
    public readonly awsAccountId!: pulumi.Output<string | undefined>;
    /**
     * <p>The time that the folder was created.</p>
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * The ID of the folder.
     */
    public readonly folderId!: pulumi.Output<string | undefined>;
    /**
     * The type of folder it is.
     */
    public readonly folderType!: pulumi.Output<enums.quicksight.FolderType | undefined>;
    /**
     * <p>The time that the folder was last updated.</p>
     */
    public /*out*/ readonly lastUpdatedTime!: pulumi.Output<string>;
    /**
     * A display name for the folder.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) for the folder.
     */
    public readonly parentFolderArn!: pulumi.Output<string | undefined>;
    /**
     * A structure that describes the principals and the resource-level permissions of a folder.
     *
     * To specify no permissions, omit `Permissions` .
     */
    public readonly permissions!: pulumi.Output<outputs.quicksight.FolderResourcePermission[] | undefined>;
    /**
     * The sharing scope of the folder.
     */
    public readonly sharingModel!: pulumi.Output<enums.quicksight.FolderSharingModel | undefined>;
    /**
     * A list of tags for the folders that you want to apply overrides to.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Folder resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FolderArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["awsAccountId"] = args ? args.awsAccountId : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["folderType"] = args ? args.folderType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parentFolderArn"] = args ? args.parentFolderArn : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["sharingModel"] = args ? args.sharingModel : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["lastUpdatedTime"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsAccountId"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["folderId"] = undefined /*out*/;
            resourceInputs["folderType"] = undefined /*out*/;
            resourceInputs["lastUpdatedTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parentFolderArn"] = undefined /*out*/;
            resourceInputs["permissions"] = undefined /*out*/;
            resourceInputs["sharingModel"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["awsAccountId", "folderId", "folderType", "parentFolderArn", "sharingModel"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Folder.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Folder resource.
 */
export interface FolderArgs {
    /**
     * The ID for the AWS account where you want to create the folder.
     */
    awsAccountId?: pulumi.Input<string>;
    /**
     * The ID of the folder.
     */
    folderId?: pulumi.Input<string>;
    /**
     * The type of folder it is.
     */
    folderType?: pulumi.Input<enums.quicksight.FolderType>;
    /**
     * A display name for the folder.
     */
    name?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) for the folder.
     */
    parentFolderArn?: pulumi.Input<string>;
    /**
     * A structure that describes the principals and the resource-level permissions of a folder.
     *
     * To specify no permissions, omit `Permissions` .
     */
    permissions?: pulumi.Input<pulumi.Input<inputs.quicksight.FolderResourcePermissionArgs>[]>;
    /**
     * The sharing scope of the folder.
     */
    sharingModel?: pulumi.Input<enums.quicksight.FolderSharingModel>;
    /**
     * A list of tags for the folders that you want to apply overrides to.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
