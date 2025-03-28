// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of the AWS::QuickSight::Theme Resource Type.
 */
export function getTheme(args: GetThemeArgs, opts?: pulumi.InvokeOptions): Promise<GetThemeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:quicksight:getTheme", {
        "awsAccountId": args.awsAccountId,
        "themeId": args.themeId,
    }, opts);
}

export interface GetThemeArgs {
    /**
     * The ID of the AWS account where you want to store the new theme.
     */
    awsAccountId: string;
    /**
     * An ID for the theme that you want to create. The theme ID is unique per AWS Region in each AWS account.
     */
    themeId: string;
}

export interface GetThemeResult {
    /**
     * <p>The Amazon Resource Name (ARN) of the theme.</p>
     */
    readonly arn?: string;
    /**
     * The ID of the theme that a custom theme will inherit from. All themes inherit from one of the starting themes defined by Amazon QuickSight. For a list of the starting themes, use `ListThemes` or choose *Themes* from within an analysis.
     */
    readonly baseThemeId?: string;
    /**
     * The theme configuration, which contains the theme display properties.
     */
    readonly configuration?: outputs.quicksight.ThemeConfiguration;
    /**
     * <p>The date and time that the theme was created.</p>
     */
    readonly createdTime?: string;
    /**
     * <p>The date and time that the theme was last updated.</p>
     */
    readonly lastUpdatedTime?: string;
    /**
     * A display name for the theme.
     */
    readonly name?: string;
    /**
     * A valid grouping of resource permissions to apply to the new theme.
     */
    readonly permissions?: outputs.quicksight.ThemeResourcePermission[];
    /**
     * A map of the key-value pairs for the resource tag or tags that you want to add to the resource.
     */
    readonly tags?: outputs.Tag[];
    /**
     * Theme type.
     */
    readonly type?: enums.quicksight.ThemeType;
    readonly version?: outputs.quicksight.ThemeVersion;
    /**
     * A description of the first version of the theme that you're creating. Every time `UpdateTheme` is called, a new version is created. Each version of the theme has a description of the version in the `VersionDescription` field.
     */
    readonly versionDescription?: string;
}
/**
 * Definition of the AWS::QuickSight::Theme Resource Type.
 */
export function getThemeOutput(args: GetThemeOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetThemeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:quicksight:getTheme", {
        "awsAccountId": args.awsAccountId,
        "themeId": args.themeId,
    }, opts);
}

export interface GetThemeOutputArgs {
    /**
     * The ID of the AWS account where you want to store the new theme.
     */
    awsAccountId: pulumi.Input<string>;
    /**
     * An ID for the theme that you want to create. The theme ID is unique per AWS Region in each AWS account.
     */
    themeId: pulumi.Input<string>;
}
