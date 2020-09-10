// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewOrganizationGithubInstallationRepoListParams creates a new OrganizationGithubInstallationRepoListParams object
// with the default values initialized.
func NewOrganizationGithubInstallationRepoListParams() *OrganizationGithubInstallationRepoListParams {

	return &OrganizationGithubInstallationRepoListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationGithubInstallationRepoListParamsWithTimeout creates a new OrganizationGithubInstallationRepoListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrganizationGithubInstallationRepoListParamsWithTimeout(timeout time.Duration) *OrganizationGithubInstallationRepoListParams {

	return &OrganizationGithubInstallationRepoListParams{

		timeout: timeout,
	}
}

// NewOrganizationGithubInstallationRepoListParamsWithContext creates a new OrganizationGithubInstallationRepoListParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrganizationGithubInstallationRepoListParamsWithContext(ctx context.Context) *OrganizationGithubInstallationRepoListParams {

	return &OrganizationGithubInstallationRepoListParams{

		Context: ctx,
	}
}

// NewOrganizationGithubInstallationRepoListParamsWithHTTPClient creates a new OrganizationGithubInstallationRepoListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrganizationGithubInstallationRepoListParamsWithHTTPClient(client *http.Client) *OrganizationGithubInstallationRepoListParams {

	return &OrganizationGithubInstallationRepoListParams{
		HTTPClient: client,
	}
}

/*OrganizationGithubInstallationRepoListParams contains all the parameters to send to the API endpoint
for the organization github installation repo list operation typically these are written to a http.Request
*/
type OrganizationGithubInstallationRepoListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the organization github installation repo list params
func (o *OrganizationGithubInstallationRepoListParams) WithTimeout(timeout time.Duration) *OrganizationGithubInstallationRepoListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organization github installation repo list params
func (o *OrganizationGithubInstallationRepoListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organization github installation repo list params
func (o *OrganizationGithubInstallationRepoListParams) WithContext(ctx context.Context) *OrganizationGithubInstallationRepoListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organization github installation repo list params
func (o *OrganizationGithubInstallationRepoListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organization github installation repo list params
func (o *OrganizationGithubInstallationRepoListParams) WithHTTPClient(client *http.Client) *OrganizationGithubInstallationRepoListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organization github installation repo list params
func (o *OrganizationGithubInstallationRepoListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationGithubInstallationRepoListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
