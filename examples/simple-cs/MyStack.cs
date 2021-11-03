// Copyright 2021, Pulumi Corporation.  All rights reserved.

using Pulumi;
using Pulumi.AwsNative.Logs;

class MyStack : Stack
{
    public MyStack()
    {
        var logGroup = new LogGroup("log-test", new LogGroupArgs
        {
            RetentionInDays = 90
        });

        this.Arn = logGroup.Arn;
    }

    [Output]
    public Output<string> Arn { get; set; }
}
