// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeleteDBParameterGroupMessage
type DeleteDBParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB parameter group.
	//
	// Constraints:
	//
	//    * Must be the name of an existing DB parameter group
	//
	//    * You can't delete a default DB parameter group
	//
	//    * Can't be associated with any DB instances
	//
	// DBParameterGroupName is a required field
	DBParameterGroupName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDBParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDBParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDBParameterGroupInput"}

	if s.DBParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBParameterGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeleteDBParameterGroupOutput
type DeleteDBParameterGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDBParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDBParameterGroup = "DeleteDBParameterGroup"

// DeleteDBParameterGroupRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Deletes a specified DB parameter group. The DB parameter group to be deleted
// can't be associated with any DB instances.
//
//    // Example sending a request using DeleteDBParameterGroupRequest.
//    req := client.DeleteDBParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeleteDBParameterGroup
func (c *Client) DeleteDBParameterGroupRequest(input *DeleteDBParameterGroupInput) DeleteDBParameterGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteDBParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDBParameterGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteDBParameterGroupOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteDBParameterGroupRequest{Request: req, Input: input, Copy: c.DeleteDBParameterGroupRequest}
}

// DeleteDBParameterGroupRequest is the request type for the
// DeleteDBParameterGroup API operation.
type DeleteDBParameterGroupRequest struct {
	*aws.Request
	Input *DeleteDBParameterGroupInput
	Copy  func(*DeleteDBParameterGroupInput) DeleteDBParameterGroupRequest
}

// Send marshals and sends the DeleteDBParameterGroup API request.
func (r DeleteDBParameterGroupRequest) Send(ctx context.Context) (*DeleteDBParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDBParameterGroupResponse{
		DeleteDBParameterGroupOutput: r.Request.Data.(*DeleteDBParameterGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDBParameterGroupResponse is the response type for the
// DeleteDBParameterGroup API operation.
type DeleteDBParameterGroupResponse struct {
	*DeleteDBParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDBParameterGroup request.
func (r *DeleteDBParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
