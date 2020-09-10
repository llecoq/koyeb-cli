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

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// NewOrganizationGithubInstallationCallbackParams creates a new OrganizationGithubInstallationCallbackParams object
// with the default values initialized.
func NewOrganizationGithubInstallationCallbackParams() *OrganizationGithubInstallationCallbackParams {
	var ()
	return &OrganizationGithubInstallationCallbackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationGithubInstallationCallbackParamsWithTimeout creates a new OrganizationGithubInstallationCallbackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrganizationGithubInstallationCallbackParamsWithTimeout(timeout time.Duration) *OrganizationGithubInstallationCallbackParams {
	var ()
	return &OrganizationGithubInstallationCallbackParams{

		timeout: timeout,
	}
}

// NewOrganizationGithubInstallationCallbackParamsWithContext creates a new OrganizationGithubInstallationCallbackParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrganizationGithubInstallationCallbackParamsWithContext(ctx context.Context) *OrganizationGithubInstallationCallbackParams {
	var ()
	return &OrganizationGithubInstallationCallbackParams{

		Context: ctx,
	}
}

// NewOrganizationGithubInstallationCallbackParamsWithHTTPClient creates a new OrganizationGithubInstallationCallbackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrganizationGithubInstallationCallbackParamsWithHTTPClient(client *http.Client) *OrganizationGithubInstallationCallbackParams {
	var ()
	return &OrganizationGithubInstallationCallbackParams{
		HTTPClient: client,
	}
}

/*OrganizationGithubInstallationCallbackParams contains all the parameters to send to the API endpoint
for the organization github installation callback operation typically these are written to a http.Request
*/
type OrganizationGithubInstallationCallbackParams struct {

	/*Body*/
	Body *models.AccountGithubInstallationCallbackRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the organization github installation callback params
func (o *OrganizationGithubInstallationCallbackParams) WithTimeout(timeout time.Duration) *OrganizationGithubInstallationCallbackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organization github installation callback params
func (o *OrganizationGithubInstallationCallbackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organization github installation callback params
func (o *OrganizationGithubInstallationCallbackParams) WithContext(ctx context.Context) *OrganizationGithubInstallationCallbackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organization github installation callback params
func (o *OrganizationGithubInstallationCallbackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organization github installation callback params
func (o *OrganizationGithubInstallationCallbackParams) WithHTTPClient(client *http.Client) *OrganizationGithubInstallationCallbackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organization github installation callback params
func (o *OrganizationGithubInstallationCallbackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the organization github installation callback params
func (o *OrganizationGithubInstallationCallbackParams) WithBody(body *models.AccountGithubInstallationCallbackRequest) *OrganizationGithubInstallationCallbackParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the organization github installation callback params
func (o *OrganizationGithubInstallationCallbackParams) SetBody(body *models.AccountGithubInstallationCallbackRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationGithubInstallationCallbackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
