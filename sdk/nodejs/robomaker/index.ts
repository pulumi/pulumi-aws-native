// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { FleetArgs } from "./fleet";
export type Fleet = import("./fleet").Fleet;
export const Fleet: typeof import("./fleet").Fleet = null as any;
utilities.lazyLoad(exports, ["Fleet"], () => require("./fleet"));

export { GetFleetArgs, GetFleetResult, GetFleetOutputArgs } from "./getFleet";
export const getFleet: typeof import("./getFleet").getFleet = null as any;
export const getFleetOutput: typeof import("./getFleet").getFleetOutput = null as any;
utilities.lazyLoad(exports, ["getFleet","getFleetOutput"], () => require("./getFleet"));

export { GetRobotArgs, GetRobotResult, GetRobotOutputArgs } from "./getRobot";
export const getRobot: typeof import("./getRobot").getRobot = null as any;
export const getRobotOutput: typeof import("./getRobot").getRobotOutput = null as any;
utilities.lazyLoad(exports, ["getRobot","getRobotOutput"], () => require("./getRobot"));

export { GetRobotApplicationArgs, GetRobotApplicationResult, GetRobotApplicationOutputArgs } from "./getRobotApplication";
export const getRobotApplication: typeof import("./getRobotApplication").getRobotApplication = null as any;
export const getRobotApplicationOutput: typeof import("./getRobotApplication").getRobotApplicationOutput = null as any;
utilities.lazyLoad(exports, ["getRobotApplication","getRobotApplicationOutput"], () => require("./getRobotApplication"));

export { GetRobotApplicationVersionArgs, GetRobotApplicationVersionResult, GetRobotApplicationVersionOutputArgs } from "./getRobotApplicationVersion";
export const getRobotApplicationVersion: typeof import("./getRobotApplicationVersion").getRobotApplicationVersion = null as any;
export const getRobotApplicationVersionOutput: typeof import("./getRobotApplicationVersion").getRobotApplicationVersionOutput = null as any;
utilities.lazyLoad(exports, ["getRobotApplicationVersion","getRobotApplicationVersionOutput"], () => require("./getRobotApplicationVersion"));

export { GetSimulationApplicationArgs, GetSimulationApplicationResult, GetSimulationApplicationOutputArgs } from "./getSimulationApplication";
export const getSimulationApplication: typeof import("./getSimulationApplication").getSimulationApplication = null as any;
export const getSimulationApplicationOutput: typeof import("./getSimulationApplication").getSimulationApplicationOutput = null as any;
utilities.lazyLoad(exports, ["getSimulationApplication","getSimulationApplicationOutput"], () => require("./getSimulationApplication"));

export { GetSimulationApplicationVersionArgs, GetSimulationApplicationVersionResult, GetSimulationApplicationVersionOutputArgs } from "./getSimulationApplicationVersion";
export const getSimulationApplicationVersion: typeof import("./getSimulationApplicationVersion").getSimulationApplicationVersion = null as any;
export const getSimulationApplicationVersionOutput: typeof import("./getSimulationApplicationVersion").getSimulationApplicationVersionOutput = null as any;
utilities.lazyLoad(exports, ["getSimulationApplicationVersion","getSimulationApplicationVersionOutput"], () => require("./getSimulationApplicationVersion"));

export { RobotArgs } from "./robot";
export type Robot = import("./robot").Robot;
export const Robot: typeof import("./robot").Robot = null as any;
utilities.lazyLoad(exports, ["Robot"], () => require("./robot"));

export { RobotApplicationArgs } from "./robotApplication";
export type RobotApplication = import("./robotApplication").RobotApplication;
export const RobotApplication: typeof import("./robotApplication").RobotApplication = null as any;
utilities.lazyLoad(exports, ["RobotApplication"], () => require("./robotApplication"));

export { RobotApplicationVersionArgs } from "./robotApplicationVersion";
export type RobotApplicationVersion = import("./robotApplicationVersion").RobotApplicationVersion;
export const RobotApplicationVersion: typeof import("./robotApplicationVersion").RobotApplicationVersion = null as any;
utilities.lazyLoad(exports, ["RobotApplicationVersion"], () => require("./robotApplicationVersion"));

export { SimulationApplicationArgs } from "./simulationApplication";
export type SimulationApplication = import("./simulationApplication").SimulationApplication;
export const SimulationApplication: typeof import("./simulationApplication").SimulationApplication = null as any;
utilities.lazyLoad(exports, ["SimulationApplication"], () => require("./simulationApplication"));

export { SimulationApplicationVersionArgs } from "./simulationApplicationVersion";
export type SimulationApplicationVersion = import("./simulationApplicationVersion").SimulationApplicationVersion;
export const SimulationApplicationVersion: typeof import("./simulationApplicationVersion").SimulationApplicationVersion = null as any;
utilities.lazyLoad(exports, ["SimulationApplicationVersion"], () => require("./simulationApplicationVersion"));


// Export enums:
export * from "../types/enums/robomaker";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:robomaker:Fleet":
                return new Fleet(name, <any>undefined, { urn })
            case "aws-native:robomaker:Robot":
                return new Robot(name, <any>undefined, { urn })
            case "aws-native:robomaker:RobotApplication":
                return new RobotApplication(name, <any>undefined, { urn })
            case "aws-native:robomaker:RobotApplicationVersion":
                return new RobotApplicationVersion(name, <any>undefined, { urn })
            case "aws-native:robomaker:SimulationApplication":
                return new SimulationApplication(name, <any>undefined, { urn })
            case "aws-native:robomaker:SimulationApplicationVersion":
                return new SimulationApplicationVersion(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "robomaker", _module)
