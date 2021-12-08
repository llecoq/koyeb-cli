package koyeb

import (
	"github.com/koyeb/koyeb-api-client-go/api/v1/koyeb"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/idmapper"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/renderer"
	"github.com/spf13/cobra"
)

func (h *ServiceHandler) Describe(cmd *cobra.Command, args []string) error {
	ctx := h.ctxWithAuth
	res, _, err := h.client.ServicesApi.GetService(ctx, h.ResolveServiceShortID(args[0])).Execute()
	if err != nil {
		fatalApiError(err)
	}

	instancesRes, _, err := h.client.InstancesApi.ListInstances(ctx).Statuses([]string{
		string(koyeb.INSTANCESTATUS_ALLOCATING),
		string(koyeb.INSTANCESTATUS_STARTING),
		string(koyeb.INSTANCESTATUS_HEALTHY),
		string(koyeb.INSTANCESTATUS_UNHEALTHY),
		string(koyeb.INSTANCESTATUS_STOPPING),
	}).ServiceId(res.Service.GetId()).Execute()
	if err != nil {
		fatalApiError(err)
	}

	deploymentsRes, _, err := h.client.DeploymentsApi.ListDeployments(ctx).ServiceId(res.Service.GetId()).Execute()
	if err != nil {
		fatalApiError(err)
	}

	appMapper := idmapper.NewAppMapper(ctx, h.client)
	serviceMapper := idmapper.NewServiceMapper(ctx, h.client)

	full, _ := cmd.Flags().GetBool("full")
	getServiceReply := NewGetServiceReply(&res, full)
	listInstancesReply := NewListInstancesReply(instancesRes, appMapper, serviceMapper, full)
	listDeploymentsReply := NewListDeploymentsReply(&deploymentsRes, full)

	output, _ := cmd.Flags().GetString("output")
	return renderer.NewDescribeRenderer(
		getServiceReply,
		listDeploymentsReply,
		listInstancesReply,
	).Render(output)
}

type DescribeServiceReply struct {
	res  *koyeb.GetServiceReply
	full bool
}

func NewDescribeServiceReply(res *koyeb.GetServiceReply, full bool) *DescribeServiceReply {
	return &DescribeServiceReply{
		res:  res,
		full: full,
	}
}

func (a *DescribeServiceReply) MarshalBinary() ([]byte, error) {
	return a.res.GetService().MarshalJSON()
}

func (a *DescribeServiceReply) Title() string {
	return "Service"
}

func (a *DescribeServiceReply) Headers() []string {
	return []string{"id", "app", "name", "version", "status", "created_at", "updated_at"}
}

func (a *DescribeServiceReply) Fields() []map[string]string {
	res := []map[string]string{}
	item := a.res.GetService()
	fields := map[string]string{
		"id":         renderer.FormatID(item.GetId(), a.full),
		"app":        renderer.FormatID(item.GetAppId(), a.full),
		"name":       item.GetName(),
		"version":    item.GetVersion(),
		"status":     formatStatus(item.State.GetStatus()),
		"created_at": renderer.FormatTime(item.GetCreatedAt()),
		"updated_at": renderer.FormatTime(item.GetUpdatedAt()),
	}
	res = append(res, fields)
	return res
}
