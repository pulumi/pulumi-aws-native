// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const InAppTemplateAlignment = {
    Left: "LEFT",
    Center: "CENTER",
    Right: "RIGHT",
} as const;

export type InAppTemplateAlignment = (typeof InAppTemplateAlignment)[keyof typeof InAppTemplateAlignment];

export const InAppTemplateButtonAction = {
    Link: "LINK",
    DeepLink: "DEEP_LINK",
    Close: "CLOSE",
} as const;

export type InAppTemplateButtonAction = (typeof InAppTemplateButtonAction)[keyof typeof InAppTemplateButtonAction];

export const InAppTemplateLayout = {
    BottomBanner: "BOTTOM_BANNER",
    TopBanner: "TOP_BANNER",
    Overlays: "OVERLAYS",
    MobileFeed: "MOBILE_FEED",
    MiddleBanner: "MIDDLE_BANNER",
    Carousel: "CAROUSEL",
} as const;

export type InAppTemplateLayout = (typeof InAppTemplateLayout)[keyof typeof InAppTemplateLayout];