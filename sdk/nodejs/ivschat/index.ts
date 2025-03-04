// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetLoggingConfigurationArgs, GetLoggingConfigurationResult, GetLoggingConfigurationOutputArgs } from "./getLoggingConfiguration";
export const getLoggingConfiguration: typeof import("./getLoggingConfiguration").getLoggingConfiguration = null as any;
export const getLoggingConfigurationOutput: typeof import("./getLoggingConfiguration").getLoggingConfigurationOutput = null as any;
utilities.lazyLoad(exports, ["getLoggingConfiguration","getLoggingConfigurationOutput"], () => require("./getLoggingConfiguration"));

export { GetRoomArgs, GetRoomResult, GetRoomOutputArgs } from "./getRoom";
export const getRoom: typeof import("./getRoom").getRoom = null as any;
export const getRoomOutput: typeof import("./getRoom").getRoomOutput = null as any;
utilities.lazyLoad(exports, ["getRoom","getRoomOutput"], () => require("./getRoom"));

export { LoggingConfigurationArgs } from "./loggingConfiguration";
export type LoggingConfiguration = import("./loggingConfiguration").LoggingConfiguration;
export const LoggingConfiguration: typeof import("./loggingConfiguration").LoggingConfiguration = null as any;
utilities.lazyLoad(exports, ["LoggingConfiguration"], () => require("./loggingConfiguration"));

export { RoomArgs } from "./room";
export type Room = import("./room").Room;
export const Room: typeof import("./room").Room = null as any;
utilities.lazyLoad(exports, ["Room"], () => require("./room"));


// Export enums:
export * from "../types/enums/ivschat";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:ivschat:LoggingConfiguration":
                return new LoggingConfiguration(name, <any>undefined, { urn })
            case "aws-native:ivschat:Room":
                return new Room(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "ivschat", _module)
