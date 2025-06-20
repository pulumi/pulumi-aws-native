// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::OpsWorksCM::Server
 */
export function getServer(args: GetServerArgs, opts?: pulumi.InvokeOptions): Promise<GetServerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:opsworkscm:getServer", {
        "id": args.id,
    }, opts);
}

export interface GetServerArgs {
    /**
     * The ID of the server.
     */
    id: string;
}

export interface GetServerResult {
    /**
     * The Amazon Resource Name (ARN) of the server, such as `arn:aws:OpsWorksCM:us-east-1:123456789012:server/server-a1bzhi` .
     */
    readonly arn?: string;
    /**
     * The number of automated backups that you want to keep. Whenever a new backup is created, AWS OpsWorks CM deletes the oldest backups if this number is exceeded. The default value is `1` .
     */
    readonly backupRetentionCount?: number;
    /**
     * Enable or disable scheduled backups. Valid values are `true` or `false` . The default value is `true` .
     */
    readonly disableAutomatedBackup?: boolean;
    /**
     * A DNS name that can be used to access the engine. Example: `myserver-asdfghjkl.us-east-1.opsworks.io` .
     */
    readonly endpoint?: string;
    /**
     * Optional engine attributes on a specified server.
     *
     * **Attributes accepted in a Chef createServer request:** - `CHEF_AUTOMATE_PIVOTAL_KEY` : A base64-encoded RSA public key. The corresponding private key is required to access the Chef API. When no CHEF_AUTOMATE_PIVOTAL_KEY is set, a private key is generated and returned in the response. When you are specifying the value of CHEF_AUTOMATE_PIVOTAL_KEY as a parameter in the AWS CloudFormation console, you must add newline ( `\n` ) characters at the end of each line of the pivotal key value.
     * - `CHEF_AUTOMATE_ADMIN_PASSWORD` : The password for the administrative user in the Chef Automate web-based dashboard. The password length is a minimum of eight characters, and a maximum of 32. The password can contain letters, numbers, and special characters (!/@#$%^&+=_). The password must contain at least one lower case letter, one upper case letter, one number, and one special character. When no CHEF_AUTOMATE_ADMIN_PASSWORD is set, one is generated and returned in the response.
     *
     * **Attributes accepted in a Puppet createServer request:** - `PUPPET_ADMIN_PASSWORD` : To work with the Puppet Enterprise console, a password must use ASCII characters.
     * - `PUPPET_R10K_REMOTE` : The r10k remote is the URL of your control repository (for example, ssh://git@your.git-repo.com:user/control-repo.git). Specifying an r10k remote opens TCP port 8170.
     * - `PUPPET_R10K_PRIVATE_KEY` : If you are using a private Git repository, add PUPPET_R10K_PRIVATE_KEY to specify a PEM-encoded private SSH key.
     */
    readonly engineAttributes?: outputs.opsworkscm.ServerEngineAttribute[];
    /**
     * The ID of the server.
     */
    readonly id?: string;
    /**
     * The start time for a one-hour period during which AWS OpsWorks CM backs up application-level data on your server if automated backups are enabled. Valid values must be specified in one of the following formats:
     *
     * - `HH:MM` for daily backups
     * - `DDD:HH:MM` for weekly backups
     *
     * `MM` must be specified as `00` . The specified time is in coordinated universal time (UTC). The default value is a random, daily start time.
     *
     * *Example:* `08:00` , which represents a daily start time of 08:00 UTC.
     *
     * *Example:* `Mon:08:00` , which represents a start time of every Monday at 08:00 UTC. (8:00 a.m.)
     */
    readonly preferredBackupWindow?: string;
    /**
     * The start time for a one-hour period each week during which AWS OpsWorks CM performs maintenance on the instance. Valid values must be specified in the following format: `DDD:HH:MM` . `MM` must be specified as `00` . The specified time is in coordinated universal time (UTC). The default value is a random one-hour period on Tuesday, Wednesday, or Friday. See `TimeWindowDefinition` for more information.
     *
     * *Example:* `Mon:08:00` , which represents a start time of every Monday at 08:00 UTC. (8:00 a.m.)
     */
    readonly preferredMaintenanceWindow?: string;
    /**
     * A map that contains tag keys and tag values to attach to an AWS OpsWorks for Chef Automate or OpsWorks for Puppet Enterprise server.
     *
     * - The key cannot be empty.
     * - The key can be a maximum of 127 characters, and can contain only Unicode letters, numbers, or separators, or the following special characters: `+ - = . _ : / @`
     * - The value can be a maximum 255 characters, and contain only Unicode letters, numbers, or separators, or the following special characters: `+ - = . _ : / @`
     * - Leading and trailing spaces are trimmed from both the key and value.
     * - A maximum of 50 user-applied tags is allowed for any AWS OpsWorks CM server.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource Type definition for AWS::OpsWorksCM::Server
 */
export function getServerOutput(args: GetServerOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:opsworkscm:getServer", {
        "id": args.id,
    }, opts);
}

export interface GetServerOutputArgs {
    /**
     * The ID of the server.
     */
    id: pulumi.Input<string>;
}
