// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { BackupPlanArgs } from "./backupPlan";
export type BackupPlan = import("./backupPlan").BackupPlan;
export const BackupPlan: typeof import("./backupPlan").BackupPlan = null as any;
utilities.lazyLoad(exports, ["BackupPlan"], () => require("./backupPlan"));

export { BackupSelectionArgs } from "./backupSelection";
export type BackupSelection = import("./backupSelection").BackupSelection;
export const BackupSelection: typeof import("./backupSelection").BackupSelection = null as any;
utilities.lazyLoad(exports, ["BackupSelection"], () => require("./backupSelection"));

export { BackupVaultArgs } from "./backupVault";
export type BackupVault = import("./backupVault").BackupVault;
export const BackupVault: typeof import("./backupVault").BackupVault = null as any;
utilities.lazyLoad(exports, ["BackupVault"], () => require("./backupVault"));

export { FrameworkArgs } from "./framework";
export type Framework = import("./framework").Framework;
export const Framework: typeof import("./framework").Framework = null as any;
utilities.lazyLoad(exports, ["Framework"], () => require("./framework"));

export { GetBackupPlanArgs, GetBackupPlanResult, GetBackupPlanOutputArgs } from "./getBackupPlan";
export const getBackupPlan: typeof import("./getBackupPlan").getBackupPlan = null as any;
export const getBackupPlanOutput: typeof import("./getBackupPlan").getBackupPlanOutput = null as any;
utilities.lazyLoad(exports, ["getBackupPlan","getBackupPlanOutput"], () => require("./getBackupPlan"));

export { GetBackupSelectionArgs, GetBackupSelectionResult, GetBackupSelectionOutputArgs } from "./getBackupSelection";
export const getBackupSelection: typeof import("./getBackupSelection").getBackupSelection = null as any;
export const getBackupSelectionOutput: typeof import("./getBackupSelection").getBackupSelectionOutput = null as any;
utilities.lazyLoad(exports, ["getBackupSelection","getBackupSelectionOutput"], () => require("./getBackupSelection"));

export { GetBackupVaultArgs, GetBackupVaultResult, GetBackupVaultOutputArgs } from "./getBackupVault";
export const getBackupVault: typeof import("./getBackupVault").getBackupVault = null as any;
export const getBackupVaultOutput: typeof import("./getBackupVault").getBackupVaultOutput = null as any;
utilities.lazyLoad(exports, ["getBackupVault","getBackupVaultOutput"], () => require("./getBackupVault"));

export { GetFrameworkArgs, GetFrameworkResult, GetFrameworkOutputArgs } from "./getFramework";
export const getFramework: typeof import("./getFramework").getFramework = null as any;
export const getFrameworkOutput: typeof import("./getFramework").getFrameworkOutput = null as any;
utilities.lazyLoad(exports, ["getFramework","getFrameworkOutput"], () => require("./getFramework"));

export { GetReportPlanArgs, GetReportPlanResult, GetReportPlanOutputArgs } from "./getReportPlan";
export const getReportPlan: typeof import("./getReportPlan").getReportPlan = null as any;
export const getReportPlanOutput: typeof import("./getReportPlan").getReportPlanOutput = null as any;
utilities.lazyLoad(exports, ["getReportPlan","getReportPlanOutput"], () => require("./getReportPlan"));

export { ReportPlanArgs } from "./reportPlan";
export type ReportPlan = import("./reportPlan").ReportPlan;
export const ReportPlan: typeof import("./reportPlan").ReportPlan = null as any;
utilities.lazyLoad(exports, ["ReportPlan"], () => require("./reportPlan"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:backup:BackupPlan":
                return new BackupPlan(name, <any>undefined, { urn })
            case "aws-native:backup:BackupSelection":
                return new BackupSelection(name, <any>undefined, { urn })
            case "aws-native:backup:BackupVault":
                return new BackupVault(name, <any>undefined, { urn })
            case "aws-native:backup:Framework":
                return new Framework(name, <any>undefined, { urn })
            case "aws-native:backup:ReportPlan":
                return new ReportPlan(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "backup", _module)