package main

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

//SendMessage Entry point
func SendMessage(targetQueueURL string) error {

	messageBody, errBody := json.Marshal("Bodystruct")
	if errBody != nil {
		return errBody
	}

	sqsMsgInput := convertToSQSMessageInput(messageBody, targetQueueURL)

	errSendTOSQS := sendMessageToQueue(sqsMsgInput)
	if errSendTOSQS != nil {
		return errSendTOSQS
	}
	return nil
}

func convertToSQSMessageInput(messageBody []byte, targetQueueURL string) sqs.SendMessageInput {

	sqsMsgInput := sqs.SendMessageInput{
		MessageBody: aws.String(string(messageBody)), // Required
		QueueUrl:    aws.String(targetQueueURL),      // Required
	}

	var attributes map[string]*sqs.MessageAttributeValue
	attributes = make(map[string]*sqs.MessageAttributeValue)

	attributes["one"] = &sqs.MessageAttributeValue{DataType: aws.String("String"), StringValue: aws.String("value1")}
	attributes["two"] = &sqs.MessageAttributeValue{DataType: aws.String("String"), StringValue: aws.String("value2")}

	sqsMsgInput.MessageAttributes = attributes

	return sqsMsgInput
}

//connector to aws
func sendMessageToQueue(sqsMsgInput sqs.SendMessageInput) error {

	region := "us-east-2"
	awsSession, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	svc := sqs.New(awsSession)

	_, errSendingSQSMsg := svc.SendMessage(&sqsMsgInput)
	if errSendingSQSMsg != nil {
		return errSendingSQSMsg
	}
	return nil
}
