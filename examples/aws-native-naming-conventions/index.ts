import * as aws from "@pulumi/aws-native";
import * as assert from "assert";

const fifoQueue = new aws.sqs.Queue("fifo-queue", { fifoQueue: true});
const standardQueue = new aws.sqs.Queue("standard-queue");

// Export the name of the fifo queue
export const fifoName = fifoQueue.queueName;
fifoName.apply(name => assert.match(name || "", new RegExp(".*\.fifo")));
export const standardName = standardQueue.queueName;
