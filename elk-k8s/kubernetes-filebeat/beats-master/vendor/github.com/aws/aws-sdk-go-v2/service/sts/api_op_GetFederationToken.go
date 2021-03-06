// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sts

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/GetFederationTokenRequest
type GetFederationTokenInput struct {
	_ struct{} `type:"structure"`

	// The duration, in seconds, that the session should last. Acceptable durations
	// for federation sessions range from 900 seconds (15 minutes) to 129,600 seconds
	// (36 hours), with 43,200 seconds (12 hours) as the default. Sessions obtained
	// using AWS account root user credentials are restricted to a maximum of 3,600
	// seconds (one hour). If the specified duration is longer than one hour, the
	// session obtained by using root user credentials defaults to one hour.
	DurationSeconds *int64 `min:"900" type:"integer"`

	// The name of the federated user. The name is used as an identifier for the
	// temporary security credentials (such as Bob). For example, you can reference
	// the federated user name in a resource-based policy, such as in an Amazon
	// S3 bucket policy.
	//
	// The regex used to validate this parameter is a string of characters consisting
	// of upper- and lower-case alphanumeric characters with no spaces. You can
	// also include underscores or any of the following characters: =,.@-
	//
	// Name is a required field
	Name *string `min:"2" type:"string" required:"true"`

	// An IAM policy in JSON format that you want to use as an inline session policy.
	//
	// You must pass an inline or managed session policy (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
	// to this operation. You can pass a single JSON policy document to use as an
	// inline session policy. You can also specify up to 10 managed policies to
	// use as managed session policies.
	//
	// This parameter is optional. However, if you do not pass any session policies,
	// then the resulting federated user session has no permissions. The only exception
	// is when the credentials are used to access a resource that has a resource-based
	// policy that specifically references the federated user session in the Principal
	// element of the policy.
	//
	// When you pass session policies, the session permissions are the intersection
	// of the IAM user policies and the session policies that you pass. This gives
	// you a way to further restrict the permissions for a federated user. You cannot
	// use session policies to grant more permissions than those that are defined
	// in the permissions policy of the IAM user. For more information, see Session
	// Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
	// in the IAM User Guide.
	//
	// The plain text that you use for both inline and managed session policies
	// shouldn't exceed 2048 characters. The JSON policy characters can be any ASCII
	// character from the space character to the end of the valid character list
	// (\u0020 through \u00FF). It can also include the tab (\u0009), linefeed (\u000A),
	// and carriage return (\u000D) characters.
	//
	// The characters in this parameter count towards the 2048 character session
	// policy guideline. However, an AWS conversion compresses the session policies
	// into a packed binary format that has a separate limit. This is the enforced
	// limit. The PackedPolicySize response element indicates by percentage how
	// close the policy is to the upper size limit.
	Policy *string `min:"1" type:"string"`

	// The Amazon Resource Names (ARNs) of the IAM managed policies that you want
	// to use as a managed session policy. The policies must exist in the same account
	// as the IAM user that is requesting federated access.
	//
	// You must pass an inline or managed session policy (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
	// to this operation. You can pass a single JSON policy document to use as an
	// inline session policy. You can also specify up to 10 managed policies to
	// use as managed session policies. The plain text that you use for both inline
	// and managed session policies shouldn't exceed 2048 characters. You can provide
	// up to 10 managed policy ARNs. For more information about ARNs, see Amazon
	// Resource Names (ARNs) and AWS Service Namespaces (general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// This parameter is optional. However, if you do not pass any session policies,
	// then the resulting federated user session has no permissions. The only exception
	// is when the credentials are used to access a resource that has a resource-based
	// policy that specifically references the federated user session in the Principal
	// element of the policy.
	//
	// When you pass session policies, the session permissions are the intersection
	// of the IAM user policies and the session policies that you pass. This gives
	// you a way to further restrict the permissions for a federated user. You cannot
	// use session policies to grant more permissions than those that are defined
	// in the permissions policy of the IAM user. For more information, see Session
	// Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
	// in the IAM User Guide.
	//
	// The characters in this parameter count towards the 2048 character session
	// policy guideline. However, an AWS conversion compresses the session policies
	// into a packed binary format that has a separate limit. This is the enforced
	// limit. The PackedPolicySize response element indicates by percentage how
	// close the policy is to the upper size limit.
	PolicyArns []PolicyDescriptorType `type:"list"`
}

// String returns the string representation
func (s GetFederationTokenInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFederationTokenInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFederationTokenInput"}
	if s.DurationSeconds != nil && *s.DurationSeconds < 900 {
		invalidParams.Add(aws.NewErrParamMinValue("DurationSeconds", 900))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 2))
	}
	if s.Policy != nil && len(*s.Policy) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Policy", 1))
	}
	if s.PolicyArns != nil {
		for i, v := range s.PolicyArns {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "PolicyArns", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful GetFederationToken request, including
// temporary AWS credentials that can be used to make AWS requests.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/GetFederationTokenResponse
type GetFederationTokenOutput struct {
	_ struct{} `type:"structure"`

	// The temporary security credentials, which include an access key ID, a secret
	// access key, and a security (or session) token.
	//
	// The size of the security token that STS API operations return is not fixed.
	// We strongly recommend that you make no assumptions about the maximum size.
	Credentials *Credentials `type:"structure"`

	// Identifiers for the federated user associated with the credentials (such
	// as arn:aws:sts::123456789012:federated-user/Bob or 123456789012:Bob). You
	// can use the federated user's ARN in your resource-based policies, such as
	// an Amazon S3 bucket policy.
	FederatedUser *FederatedUser `type:"structure"`

	// A percentage value indicating the size of the policy in packed form. The
	// service rejects policies for which the packed size is greater than 100 percent
	// of the allowed value.
	PackedPolicySize *int64 `type:"integer"`
}

// String returns the string representation
func (s GetFederationTokenOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetFederationToken = "GetFederationToken"

// GetFederationTokenRequest returns a request value for making API operation for
// AWS Security Token Service.
//
// Returns a set of temporary security credentials (consisting of an access
// key ID, a secret access key, and a security token) for a federated user.
// A typical use is in a proxy application that gets temporary security credentials
// on behalf of distributed applications inside a corporate network. You must
// call the GetFederationToken operation using the long-term security credentials
// of an IAM user. As a result, this call is appropriate in contexts where those
// credentials can be safely stored, usually in a server-based application.
// For a comparison of GetFederationToken with the other API operations that
// produce temporary credentials, see Requesting Temporary Security Credentials
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html)
// and Comparing the AWS STS API operations (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#stsapi_comparison)
// in the IAM User Guide.
//
// You can create a mobile-based or browser-based app that can authenticate
// users using a web identity provider like Login with Amazon, Facebook, Google,
// or an OpenID Connect-compatible identity provider. In this case, we recommend
// that you use Amazon Cognito (http://aws.amazon.com/cognito/) or AssumeRoleWithWebIdentity.
// For more information, see Federation Through a Web-based Identity Provider
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#api_assumerolewithwebidentity).
//
// You can also call GetFederationToken using the security credentials of an
// AWS account root user, but we do not recommend it. Instead, we recommend
// that you create an IAM user for the purpose of the proxy application. Then
// attach a policy to the IAM user that limits federated users to only the actions
// and resources that they need to access. For more information, see IAM Best
// Practices (https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html)
// in the IAM User Guide.
//
// The temporary credentials are valid for the specified duration, from 900
// seconds (15 minutes) up to a maximum of 129,600 seconds (36 hours). The default
// is 43,200 seconds (12 hours). Temporary credentials that are obtained by
// using AWS account root user credentials have a maximum duration of 3,600
// seconds (1 hour).
//
// The temporary security credentials created by GetFederationToken can be used
// to make API calls to any AWS service with the following exceptions:
//
//    * You cannot use these credentials to call any IAM API operations.
//
//    * You cannot call any STS API operations except GetCallerIdentity.
//
// Permissions
//
// You must pass an inline or managed session policy (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
// to this operation. You can pass a single JSON policy document to use as an
// inline session policy. You can also specify up to 10 managed policies to
// use as managed session policies. The plain text that you use for both inline
// and managed session policies shouldn't exceed 2048 characters.
//
// Though the session policy parameters are optional, if you do not pass a policy,
// then the resulting federated user session has no permissions. The only exception
// is when the credentials are used to access a resource that has a resource-based
// policy that specifically references the federated user session in the Principal
// element of the policy. When you pass session policies, the session permissions
// are the intersection of the IAM user policies and the session policies that
// you pass. This gives you a way to further restrict the permissions for a
// federated user. You cannot use session policies to grant more permissions
// than those that are defined in the permissions policy of the IAM user. For
// more information, see Session Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
// in the IAM User Guide. For information about using GetFederationToken to
// create temporary security credentials, see GetFederationToken???Federation
// Through a Custom Identity Broker (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#api_getfederationtoken).
//
//    // Example sending a request using GetFederationTokenRequest.
//    req := client.GetFederationTokenRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/GetFederationToken
func (c *Client) GetFederationTokenRequest(input *GetFederationTokenInput) GetFederationTokenRequest {
	op := &aws.Operation{
		Name:       opGetFederationToken,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetFederationTokenInput{}
	}

	req := c.newRequest(op, input, &GetFederationTokenOutput{})
	return GetFederationTokenRequest{Request: req, Input: input, Copy: c.GetFederationTokenRequest}
}

// GetFederationTokenRequest is the request type for the
// GetFederationToken API operation.
type GetFederationTokenRequest struct {
	*aws.Request
	Input *GetFederationTokenInput
	Copy  func(*GetFederationTokenInput) GetFederationTokenRequest
}

// Send marshals and sends the GetFederationToken API request.
func (r GetFederationTokenRequest) Send(ctx context.Context) (*GetFederationTokenResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFederationTokenResponse{
		GetFederationTokenOutput: r.Request.Data.(*GetFederationTokenOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFederationTokenResponse is the response type for the
// GetFederationToken API operation.
type GetFederationTokenResponse struct {
	*GetFederationTokenOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFederationToken request.
func (r *GetFederationTokenResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
