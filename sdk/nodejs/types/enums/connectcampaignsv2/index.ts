// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const CampaignCommunicationLimitTimeUnit = {
    Day: "DAY",
} as const;

/**
 * The communication limit time unit
 */
export type CampaignCommunicationLimitTimeUnit = (typeof CampaignCommunicationLimitTimeUnit)[keyof typeof CampaignCommunicationLimitTimeUnit];

export const CampaignDayOfWeek = {
    Monday: "MONDAY",
    Tuesday: "TUESDAY",
    Wednesday: "WEDNESDAY",
    Thursday: "THURSDAY",
    Friday: "FRIDAY",
    Saturday: "SATURDAY",
    Sunday: "SUNDAY",
} as const;

/**
 * Day of week
 */
export type CampaignDayOfWeek = (typeof CampaignDayOfWeek)[keyof typeof CampaignDayOfWeek];

export const CampaignInstanceLimitsHandling = {
    OptIn: "OPT_IN",
    OptOut: "OPT_OUT",
} as const;

/**
 * Enumeration of Instance Limits handling in a Campaign
 */
export type CampaignInstanceLimitsHandling = (typeof CampaignInstanceLimitsHandling)[keyof typeof CampaignInstanceLimitsHandling];

export const CampaignLocalTimeZoneDetectionType = {
    ZipCode: "ZIP_CODE",
    AreaCode: "AREA_CODE",
} as const;

/**
 * Local TimeZone Detection method
 */
export type CampaignLocalTimeZoneDetectionType = (typeof CampaignLocalTimeZoneDetectionType)[keyof typeof CampaignLocalTimeZoneDetectionType];
