// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Registers a package version.
 */
export class PackageVersion extends pulumi.CustomResource {
    /**
     * Get an existing PackageVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PackageVersion {
        return new PackageVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:panorama:PackageVersion';

    /**
     * Returns true if the given object is an instance of PackageVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PackageVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PackageVersion.__pulumiType;
    }

    /**
     * Whether the package version is the latest version.
     */
    public /*out*/ readonly isLatestPatch!: pulumi.Output<boolean>;
    /**
     * Whether to mark the new version as the latest version.
     */
    public readonly markLatest!: pulumi.Output<boolean | undefined>;
    /**
     * An owner account.
     */
    public readonly ownerAccount!: pulumi.Output<string | undefined>;
    /**
     * The package version's ARN.
     */
    public /*out*/ readonly packageArn!: pulumi.Output<string>;
    /**
     * A package ID.
     */
    public readonly packageId!: pulumi.Output<string>;
    /**
     * The package version's name.
     */
    public /*out*/ readonly packageName!: pulumi.Output<string>;
    /**
     * A package version.
     */
    public readonly packageVersion!: pulumi.Output<string>;
    /**
     * A patch version.
     */
    public readonly patchVersion!: pulumi.Output<string>;
    /**
     * The package version's registered time.
     */
    public /*out*/ readonly registeredTime!: pulumi.Output<number>;
    /**
     * The package version's status.
     */
    public /*out*/ readonly status!: pulumi.Output<enums.panorama.PackageVersionStatus>;
    /**
     * The package version's status description.
     */
    public /*out*/ readonly statusDescription!: pulumi.Output<string>;
    /**
     * If the version was marked latest, the new version to maker as latest.
     */
    public readonly updatedLatestPatchVersion!: pulumi.Output<string | undefined>;

    /**
     * Create a PackageVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PackageVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.packageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'packageId'");
            }
            if ((!args || args.packageVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'packageVersion'");
            }
            if ((!args || args.patchVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'patchVersion'");
            }
            resourceInputs["markLatest"] = args ? args.markLatest : undefined;
            resourceInputs["ownerAccount"] = args ? args.ownerAccount : undefined;
            resourceInputs["packageId"] = args ? args.packageId : undefined;
            resourceInputs["packageVersion"] = args ? args.packageVersion : undefined;
            resourceInputs["patchVersion"] = args ? args.patchVersion : undefined;
            resourceInputs["updatedLatestPatchVersion"] = args ? args.updatedLatestPatchVersion : undefined;
            resourceInputs["isLatestPatch"] = undefined /*out*/;
            resourceInputs["packageArn"] = undefined /*out*/;
            resourceInputs["packageName"] = undefined /*out*/;
            resourceInputs["registeredTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusDescription"] = undefined /*out*/;
        } else {
            resourceInputs["isLatestPatch"] = undefined /*out*/;
            resourceInputs["markLatest"] = undefined /*out*/;
            resourceInputs["ownerAccount"] = undefined /*out*/;
            resourceInputs["packageArn"] = undefined /*out*/;
            resourceInputs["packageId"] = undefined /*out*/;
            resourceInputs["packageName"] = undefined /*out*/;
            resourceInputs["packageVersion"] = undefined /*out*/;
            resourceInputs["patchVersion"] = undefined /*out*/;
            resourceInputs["registeredTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusDescription"] = undefined /*out*/;
            resourceInputs["updatedLatestPatchVersion"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["ownerAccount", "packageId", "packageVersion", "patchVersion"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(PackageVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PackageVersion resource.
 */
export interface PackageVersionArgs {
    /**
     * Whether to mark the new version as the latest version.
     */
    markLatest?: pulumi.Input<boolean>;
    /**
     * An owner account.
     */
    ownerAccount?: pulumi.Input<string>;
    /**
     * A package ID.
     */
    packageId: pulumi.Input<string>;
    /**
     * A package version.
     */
    packageVersion: pulumi.Input<string>;
    /**
     * A patch version.
     */
    patchVersion: pulumi.Input<string>;
    /**
     * If the version was marked latest, the new version to maker as latest.
     */
    updatedLatestPatchVersion?: pulumi.Input<string>;
}
