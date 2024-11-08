/**
* A sample Lambda function that looks up the latest AMI ID for a given region and architecture.
**/

// Map instance architectures to an AMI name pattern
var archToAMINamePattern = {
    "PV64": "amzn-ami-pv*x86_64-ebs",
    "HVM64": "al2023-ami-2023.*-kernel-*-x86_64",
    "HVMG2": "amzn-ami-graphics-hvm*x86_64-ebs*"
};
const { EC2Client, DescribeImagesCommand } = require("@aws-sdk/client-ec2");

exports.handler = async function(event, context) {
    const redactedEvent = { ...event, ResponseURL: "REDACTED" };
    console.log("REQUEST RECEIVED:\n" + JSON.stringify(redactedEvent));
    
    // For Delete requests, immediately send a SUCCESS response.
    if (event.RequestType == "Delete") {
        await sendResponse(event, context, "SUCCESS");
        return;
    }
 
    var responseStatus = "FAILED";
    var responseData = {};
 
    const ec2Client = new EC2Client({ region: event.ResourceProperties.Region });
    const describeImagesParams = {
        Filters: [{ Name: "name", Values: [archToAMINamePattern[event.ResourceProperties.Architecture]]}],
        Owners: [event.ResourceProperties.Architecture == "HVMG2" ? "679593333241" : "amazon"]
    };
 
    try {
        const describeImagesResult = await ec2Client.send(new DescribeImagesCommand(describeImagesParams));
        var images = describeImagesResult.Images;
        // Sort images by name in descending order. The names contain the AMI version, formatted as YYYY.MM.Ver.
        images.sort((x, y) => y.Name.localeCompare(x.Name));
        for (var j = 0; j < images.length; j++) {
            if (isBeta(images[j].Name)) continue;
            responseStatus = "SUCCESS";
            responseData["Id"] = images[j].ImageId;
            break;
        }
    } catch (err) {
        responseData = { Error: "DescribeImages call failed" };
        console.log(responseData.Error + ":\n", err);
    }
    
    await sendResponse(event, context, responseStatus, responseData);
};

// Check if the image is a beta or rc image. The Lambda function won't return any of those images.
function isBeta(imageName) {
    return imageName.toLowerCase().indexOf("beta") > -1 || imageName.toLowerCase().indexOf(".rc") > -1;
}

// Send response to the pre-signed S3 URL 
async function sendResponse(event, context, responseStatus, responseData) {
    var responseBody = JSON.stringify({
        Status: responseStatus,
        Reason: "See the details in CloudWatch Log Stream: " + context.logStreamName,
        PhysicalResourceId: context.logStreamName,
        StackId: event.StackId,
        RequestId: event.RequestId,
        LogicalResourceId: event.LogicalResourceId,
        Data: responseData
    });
 
    console.log("RESPONSE BODY:\n", responseBody);
 
    var https = require("https");
    var url = require("url");
 
    var parsedUrl = url.parse(event.ResponseURL);
    var options = {
        hostname: parsedUrl.hostname,
        port: 443,
        path: parsedUrl.path,
        method: "PUT",
        headers: {
            "content-type": "",
            "content-length": responseBody.length
        }
    };
 
    console.log("SENDING RESPONSE...\n");
 
    var request = https.request(options, function(response) {
        console.log("STATUS: " + response.statusCode);
        console.log("HEADERS: " + JSON.stringify(response.headers));
        // Tell AWS Lambda that the function execution is done  
        context.done();
    });
 
    request.on("error", function(error) {
        console.log("sendResponse Error:" + error);
        // Tell AWS Lambda that the function execution is done  
        context.done();
    });
  
    // write data to request body
    request.write(responseBody);
    request.end();
}
