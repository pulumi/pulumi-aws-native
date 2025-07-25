// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const CloudAutonomousVmClusterComputeModel = {
    Ecpu: "ECPU",
    Ocpu: "OCPU",
} as const;

/**
 * The compute model of the Autonomous VM cluster: ECPU or OCPU.
 */
export type CloudAutonomousVmClusterComputeModel = (typeof CloudAutonomousVmClusterComputeModel)[keyof typeof CloudAutonomousVmClusterComputeModel];

export const CloudAutonomousVmClusterLicenseModel = {
    BringYourOwnLicense: "BRING_YOUR_OWN_LICENSE",
    LicenseIncluded: "LICENSE_INCLUDED",
} as const;

/**
 * The Oracle license model that applies to the Autonomous VM cluster. Valid values are LICENSE_INCLUDED or BRING_YOUR_OWN_LICENSE.
 */
export type CloudAutonomousVmClusterLicenseModel = (typeof CloudAutonomousVmClusterLicenseModel)[keyof typeof CloudAutonomousVmClusterLicenseModel];

export const CloudAutonomousVmClusterMaintenanceWindowDaysOfWeekItem = {
    Monday: "MONDAY",
    Tuesday: "TUESDAY",
    Wednesday: "WEDNESDAY",
    Thursday: "THURSDAY",
    Friday: "FRIDAY",
    Saturday: "SATURDAY",
    Sunday: "SUNDAY",
} as const;

export type CloudAutonomousVmClusterMaintenanceWindowDaysOfWeekItem = (typeof CloudAutonomousVmClusterMaintenanceWindowDaysOfWeekItem)[keyof typeof CloudAutonomousVmClusterMaintenanceWindowDaysOfWeekItem];

export const CloudAutonomousVmClusterMaintenanceWindowMonthsItem = {
    January: "JANUARY",
    February: "FEBRUARY",
    March: "MARCH",
    April: "APRIL",
    May: "MAY",
    June: "JUNE",
    July: "JULY",
    August: "AUGUST",
    September: "SEPTEMBER",
    October: "OCTOBER",
    November: "NOVEMBER",
    December: "DECEMBER",
} as const;

export type CloudAutonomousVmClusterMaintenanceWindowMonthsItem = (typeof CloudAutonomousVmClusterMaintenanceWindowMonthsItem)[keyof typeof CloudAutonomousVmClusterMaintenanceWindowMonthsItem];

export const CloudAutonomousVmClusterMaintenanceWindowPreference = {
    NoPreference: "NO_PREFERENCE",
    CustomPreference: "CUSTOM_PREFERENCE",
} as const;

/**
 * The preference for the maintenance window scheduling.
 */
export type CloudAutonomousVmClusterMaintenanceWindowPreference = (typeof CloudAutonomousVmClusterMaintenanceWindowPreference)[keyof typeof CloudAutonomousVmClusterMaintenanceWindowPreference];

export const CloudVmClusterLicenseModel = {
    BringYourOwnLicense: "BRING_YOUR_OWN_LICENSE",
    LicenseIncluded: "LICENSE_INCLUDED",
} as const;

/**
 * The Oracle license model applied to the VM cluster.
 */
export type CloudVmClusterLicenseModel = (typeof CloudVmClusterLicenseModel)[keyof typeof CloudVmClusterLicenseModel];
