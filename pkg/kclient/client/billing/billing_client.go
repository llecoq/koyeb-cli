// Code generated by go-swagger; DO NOT EDIT.

package billing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new billing API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for billing API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	BillingCreateBillingInfo(params *BillingCreateBillingInfoParams, authInfo runtime.ClientAuthInfoWriter) (*BillingCreateBillingInfoOK, error)

	BillingGetBillingInfo(params *BillingGetBillingInfoParams, authInfo runtime.ClientAuthInfoWriter) (*BillingGetBillingInfoOK, error)

	BillingUpdateBilling(params *BillingUpdateBillingParams, authInfo runtime.ClientAuthInfoWriter) (*BillingUpdateBillingOK, error)

	BillingUpdateBilling2(params *BillingUpdateBilling2Params, authInfo runtime.ClientAuthInfoWriter) (*BillingUpdateBilling2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BillingCreateBillingInfo finishes registration
*/
func (a *Client) BillingCreateBillingInfo(params *BillingCreateBillingInfoParams, authInfo runtime.ClientAuthInfoWriter) (*BillingCreateBillingInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBillingCreateBillingInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "billing_CreateBillingInfo",
		Method:             "POST",
		PathPattern:        "/v1/account/billing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BillingCreateBillingInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BillingCreateBillingInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BillingCreateBillingInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BillingGetBillingInfo billings
*/
func (a *Client) BillingGetBillingInfo(params *BillingGetBillingInfoParams, authInfo runtime.ClientAuthInfoWriter) (*BillingGetBillingInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBillingGetBillingInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "billing_GetBillingInfo",
		Method:             "GET",
		PathPattern:        "/v1/account/billing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BillingGetBillingInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BillingGetBillingInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BillingGetBillingInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BillingUpdateBilling updates billing info
*/
func (a *Client) BillingUpdateBilling(params *BillingUpdateBillingParams, authInfo runtime.ClientAuthInfoWriter) (*BillingUpdateBillingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBillingUpdateBillingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "billing_UpdateBilling",
		Method:             "PUT",
		PathPattern:        "/v1/account/billing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BillingUpdateBillingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BillingUpdateBillingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BillingUpdateBillingDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BillingUpdateBilling2 updates billing info
*/
func (a *Client) BillingUpdateBilling2(params *BillingUpdateBilling2Params, authInfo runtime.ClientAuthInfoWriter) (*BillingUpdateBilling2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBillingUpdateBilling2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "billing_UpdateBilling2",
		Method:             "PATCH",
		PathPattern:        "/v1/account/billing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BillingUpdateBilling2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BillingUpdateBilling2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BillingUpdateBilling2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
