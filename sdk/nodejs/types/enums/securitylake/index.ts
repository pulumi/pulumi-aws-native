// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const SubscriberAccessTypesItem = {
    Lakeformation: "LAKEFORMATION",
    S3: "S3",
} as const;

export type SubscriberAccessTypesItem = (typeof SubscriberAccessTypesItem)[keyof typeof SubscriberAccessTypesItem];

export const SubscriberNotificationHttpsNotificationConfigurationHttpMethod = {
    Post: "POST",
    Put: "PUT",
} as const;

/**
 * The HTTPS method used for the notification subscription.
 */
export type SubscriberNotificationHttpsNotificationConfigurationHttpMethod = (typeof SubscriberNotificationHttpsNotificationConfigurationHttpMethod)[keyof typeof SubscriberNotificationHttpsNotificationConfigurationHttpMethod];