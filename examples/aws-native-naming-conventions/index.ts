import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws-native";
import * as assert from "assert";

// Create an AWS resource (AWS SQS FIFO Queue)
const queue = new aws.sqs.Queue("my-queue",{
    fifoQueue: true,
});

// Export the name of the queue
export const name = queue.queueName;
name.apply(name => assert.match(name || "", new RegExp(".*\.fifo")));
