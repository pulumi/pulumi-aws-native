// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetLaunchProfileArgs, GetLaunchProfileResult, GetLaunchProfileOutputArgs } from "./getLaunchProfile";
export const getLaunchProfile: typeof import("./getLaunchProfile").getLaunchProfile = null as any;
export const getLaunchProfileOutput: typeof import("./getLaunchProfile").getLaunchProfileOutput = null as any;

export { GetStreamingImageArgs, GetStreamingImageResult, GetStreamingImageOutputArgs } from "./getStreamingImage";
export const getStreamingImage: typeof import("./getStreamingImage").getStreamingImage = null as any;
export const getStreamingImageOutput: typeof import("./getStreamingImage").getStreamingImageOutput = null as any;

export { GetStudioArgs, GetStudioResult, GetStudioOutputArgs } from "./getStudio";
export const getStudio: typeof import("./getStudio").getStudio = null as any;
export const getStudioOutput: typeof import("./getStudio").getStudioOutput = null as any;

export { GetStudioComponentArgs, GetStudioComponentResult, GetStudioComponentOutputArgs } from "./getStudioComponent";
export const getStudioComponent: typeof import("./getStudioComponent").getStudioComponent = null as any;
export const getStudioComponentOutput: typeof import("./getStudioComponent").getStudioComponentOutput = null as any;

export { LaunchProfileArgs } from "./launchProfile";
export type LaunchProfile = import("./launchProfile").LaunchProfile;
export const LaunchProfile: typeof import("./launchProfile").LaunchProfile = null as any;

export { StreamingImageArgs } from "./streamingImage";
export type StreamingImage = import("./streamingImage").StreamingImage;
export const StreamingImage: typeof import("./streamingImage").StreamingImage = null as any;

export { StudioArgs } from "./studio";
export type Studio = import("./studio").Studio;
export const Studio: typeof import("./studio").Studio = null as any;

export { StudioComponentArgs } from "./studioComponent";
export type StudioComponent = import("./studioComponent").StudioComponent;
export const StudioComponent: typeof import("./studioComponent").StudioComponent = null as any;

utilities.lazyLoad(exports, ["getLaunchProfile","getLaunchProfileOutput"], () => require("./getLaunchProfile"));
utilities.lazyLoad(exports, ["getStreamingImage","getStreamingImageOutput"], () => require("./getStreamingImage"));
utilities.lazyLoad(exports, ["getStudio","getStudioOutput"], () => require("./getStudio"));
utilities.lazyLoad(exports, ["getStudioComponent","getStudioComponentOutput"], () => require("./getStudioComponent"));
utilities.lazyLoad(exports, ["LaunchProfile"], () => require("./launchProfile"));
utilities.lazyLoad(exports, ["StreamingImage"], () => require("./streamingImage"));
utilities.lazyLoad(exports, ["Studio"], () => require("./studio"));
utilities.lazyLoad(exports, ["StudioComponent"], () => require("./studioComponent"));

// Export enums:
export * from "../types/enums/nimblestudio";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:nimblestudio:LaunchProfile":
                return new LaunchProfile(name, <any>undefined, { urn })
            case "aws-native:nimblestudio:StreamingImage":
                return new StreamingImage(name, <any>undefined, { urn })
            case "aws-native:nimblestudio:Studio":
                return new Studio(name, <any>undefined, { urn })
            case "aws-native:nimblestudio:StudioComponent":
                return new StudioComponent(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "nimblestudio", _module)