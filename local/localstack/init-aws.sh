#!/bin/sh

DEFAULT_SNS_ARN='arn:aws:sns:ap-northeast-2:000000000000:'
DEFAULT_SQS_ARN='arn:aws:sqs:ap-northeast-2:000000000000:'

# create SQS
awslocal sqs create-queue \
--queue-name 'ANY-QUEUE'

# create SNS
awslocal sns create-topic \
--name 'ANY-SNS'

# subscription SNS to SQS
awslocal sns subscribe \
--topic-arn $DEFAULT_SNS_ARN'ANY-SNS' \
--protocol sqs \
--notification-endpoint $DEFAULT_SQS_ARN'ANY-QUEUE' \
--attributes RawMessageDelivery=true