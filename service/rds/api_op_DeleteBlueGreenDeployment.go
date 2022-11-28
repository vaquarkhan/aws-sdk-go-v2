// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a blue/green deployment. For more information, see Using Amazon RDS
// Blue/Green Deployments for database updates
// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/blue-green-deployments.html)
// in the Amazon RDS User Guide and  Using Amazon RDS Blue/Green Deployments for
// database updates
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/blue-green-deployments.html)
// in the Amazon Aurora User Guide.
func (c *Client) DeleteBlueGreenDeployment(ctx context.Context, params *DeleteBlueGreenDeploymentInput, optFns ...func(*Options)) (*DeleteBlueGreenDeploymentOutput, error) {
	if params == nil {
		params = &DeleteBlueGreenDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBlueGreenDeployment", params, optFns, c.addOperationDeleteBlueGreenDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBlueGreenDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBlueGreenDeploymentInput struct {

	// The blue/green deployment identifier of the deployment to be deleted. This
	// parameter isn't case-sensitive. Constraints:
	//
	// * Must match an existing
	// blue/green deployment identifier.
	//
	// This member is required.
	BlueGreenDeploymentIdentifier *string

	// A value that indicates whether to delete the resources in the green environment.
	DeleteTarget *bool

	noSmithyDocumentSerde
}

type DeleteBlueGreenDeploymentOutput struct {

	// Contains the details about a blue/green deployment. For more information, see
	// Using Amazon RDS Blue/Green Deployments for database updates
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/blue-green-deployments.html)
	// in the Amazon RDS User Guide and  Using Amazon RDS Blue/Green Deployments for
	// database updates
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/blue-green-deployments.html)
	// in the Amazon Aurora User Guide.
	BlueGreenDeployment *types.BlueGreenDeployment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteBlueGreenDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteBlueGreenDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteBlueGreenDeployment{}, middleware.After)
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
	if err = addOpDeleteBlueGreenDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBlueGreenDeployment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteBlueGreenDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteBlueGreenDeployment",
	}
}