// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables or disables automated sensitive data discovery for an account.
func (c *Client) UpdateAutomatedDiscoveryConfiguration(ctx context.Context, params *UpdateAutomatedDiscoveryConfigurationInput, optFns ...func(*Options)) (*UpdateAutomatedDiscoveryConfigurationOutput, error) {
	if params == nil {
		params = &UpdateAutomatedDiscoveryConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAutomatedDiscoveryConfiguration", params, optFns, c.addOperationUpdateAutomatedDiscoveryConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAutomatedDiscoveryConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAutomatedDiscoveryConfigurationInput struct {

	// The new status of automated sensitive data discovery for the account. Valid
	// values are: ENABLED, start or resume automated sensitive data discovery
	// activities for the account; and, DISABLED, stop performing automated sensitive
	// data discovery activities for the account. When you enable automated sensitive
	// data discovery for the first time, Amazon Macie uses default configuration
	// settings to determine which data sources to analyze and which managed data
	// identifiers to use. To change these settings, use the UpdateClassificationScope
	// and UpdateSensitivityInspectionTemplate operations, respectively. If you change
	// the settings and subsequently disable the configuration, Amazon Macie retains
	// your changes.
	//
	// This member is required.
	Status types.AutomatedDiscoveryStatus

	noSmithyDocumentSerde
}

type UpdateAutomatedDiscoveryConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAutomatedDiscoveryConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAutomatedDiscoveryConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAutomatedDiscoveryConfiguration{}, middleware.After)
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
	if err = addOpUpdateAutomatedDiscoveryConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAutomatedDiscoveryConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAutomatedDiscoveryConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "UpdateAutomatedDiscoveryConfiguration",
	}
}