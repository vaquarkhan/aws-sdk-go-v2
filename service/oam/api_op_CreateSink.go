// Code generated by smithy-go-codegen DO NOT EDIT.

package oam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this to create a sink in the current account, so that it can be used as a
// monitoring account in CloudWatch cross-account observability. A sink is a
// resource that represents an attachment point in a monitoring account. Source
// accounts can link to the sink to send observability data. After you create a
// sink, you must create a sink policy that allows source accounts to attach to it.
// For more information, see PutSinkPolicy
// (https://docs.aws.amazon.com/OAM/latest/APIReference/API_PutSinkPolicy.html).
// Each account can contain one sink. If you delete a sink, you can then create a
// new one in that account.
func (c *Client) CreateSink(ctx context.Context, params *CreateSinkInput, optFns ...func(*Options)) (*CreateSinkOutput, error) {
	if params == nil {
		params = &CreateSinkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSink", params, optFns, c.addOperationCreateSinkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSinkInput struct {

	// A name for the sink.
	//
	// This member is required.
	Name *string

	// Assigns one or more tags (key-value pairs) to the link. Tags can help you
	// organize and categorize your resources. You can also use them to scope user
	// permissions by granting a user permission to access or change only resources
	// with certain tag values. For more information about using tags to control
	// access, see Controlling access to Amazon Web Services resources using tags
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html).
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateSinkOutput struct {

	// The ARN of the sink that is newly created.
	Arn *string

	// The random ID string that Amazon Web Services generated as part of the sink ARN.
	Id *string

	// The name of the sink.
	Name *string

	// The tags assigned to the sink.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSinkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSink{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSink{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateSinkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSink(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateSink(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "oam",
		OperationName: "CreateSink",
	}
}