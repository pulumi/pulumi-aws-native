// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::QBusiness::Plugin Resource Type
 */
export function getPlugin(args: GetPluginArgs, opts?: pulumi.InvokeOptions): Promise<GetPluginResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:qbusiness:getPlugin", {
        "applicationId": args.applicationId,
        "pluginId": args.pluginId,
    }, opts);
}

export interface GetPluginArgs {
    /**
     * The identifier of the application that will contain the plugin.
     */
    applicationId: string;
    /**
     * The identifier of the plugin.
     */
    pluginId: string;
}

export interface GetPluginResult {
    /**
     * Authentication configuration information for an Amazon Q Business plugin.
     */
    readonly authConfiguration?: outputs.qbusiness.PluginAuthConfiguration0Properties | outputs.qbusiness.PluginAuthConfiguration1Properties | outputs.qbusiness.PluginAuthConfiguration2Properties;
    /**
     * The current status of a plugin. A plugin is modified asynchronously.
     */
    readonly buildStatus?: enums.qbusiness.PluginBuildStatus;
    /**
     * The timestamp for when the plugin was created.
     */
    readonly createdAt?: string;
    /**
     * Configuration information required to create a custom plugin.
     */
    readonly customPluginConfiguration?: outputs.qbusiness.PluginCustomPluginConfiguration;
    /**
     * The name of the plugin.
     */
    readonly displayName?: string;
    /**
     * The Amazon Resource Name (ARN) of a plugin.
     */
    readonly pluginArn?: string;
    /**
     * The identifier of the plugin.
     */
    readonly pluginId?: string;
    /**
     * The plugin server URL used for configuration.
     */
    readonly serverUrl?: string;
    /**
     * The current status of the plugin.
     */
    readonly state?: enums.qbusiness.PluginState;
    /**
     * A list of key-value pairs that identify or categorize the data source connector. You can also use tags to help control access to the data source connector. Tag keys and values can consist of Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The timestamp for when the plugin was last updated.
     */
    readonly updatedAt?: string;
}
/**
 * Definition of AWS::QBusiness::Plugin Resource Type
 */
export function getPluginOutput(args: GetPluginOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPluginResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:qbusiness:getPlugin", {
        "applicationId": args.applicationId,
        "pluginId": args.pluginId,
    }, opts);
}

export interface GetPluginOutputArgs {
    /**
     * The identifier of the application that will contain the plugin.
     */
    applicationId: pulumi.Input<string>;
    /**
     * The identifier of the plugin.
     */
    pluginId: pulumi.Input<string>;
}
