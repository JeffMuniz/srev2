// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

// Contains the parameters for DeleteSpotDatafeedSubscription.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteSpotDatafeedSubscriptionRequest
type DeleteSpotDatafeedSubscriptionInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`
}

// String returns the string representation
func (s DeleteSpotDatafeedSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteSpotDatafeedSubscriptionOutput
type DeleteSpotDatafeedSubscriptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteSpotDatafeedSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteSpotDatafeedSubscription = "DeleteSpotDatafeedSubscription"

// DeleteSpotDatafeedSubscriptionRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the data feed for Spot Instances.
//
//    // Example sending a request using DeleteSpotDatafeedSubscriptionRequest.
//    req := client.DeleteSpotDatafeedSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteSpotDatafeedSubscription
func (c *Client) DeleteSpotDatafeedSubscriptionRequest(input *DeleteSpotDatafeedSubscriptionInput) DeleteSpotDatafeedSubscriptionRequest {
	op := &aws.Operation{
		Name:       opDeleteSpotDatafeedSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteSpotDatafeedSubscriptionInput{}
	}

	req := c.newRequest(op, input, &DeleteSpotDatafeedSubscriptionOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteSpotDatafeedSubscriptionRequest{Request: req, Input: input, Copy: c.DeleteSpotDatafeedSubscriptionRequest}
}

// DeleteSpotDatafeedSubscriptionRequest is the request type for the
// DeleteSpotDatafeedSubscription API operation.
type DeleteSpotDatafeedSubscriptionRequest struct {
	*aws.Request
	Input *DeleteSpotDatafeedSubscriptionInput
	Copy  func(*DeleteSpotDatafeedSubscriptionInput) DeleteSpotDatafeedSubscriptionRequest
}

// Send marshals and sends the DeleteSpotDatafeedSubscription API request.
func (r DeleteSpotDatafeedSubscriptionRequest) Send(ctx context.Context) (*DeleteSpotDatafeedSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSpotDatafeedSubscriptionResponse{
		DeleteSpotDatafeedSubscriptionOutput: r.Request.Data.(*DeleteSpotDatafeedSubscriptionOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSpotDatafeedSubscriptionResponse is the response type for the
// DeleteSpotDatafeedSubscription API operation.
type DeleteSpotDatafeedSubscriptionResponse struct {
	*DeleteSpotDatafeedSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSpotDatafeedSubscription request.
func (r *DeleteSpotDatafeedSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
