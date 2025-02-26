package koyeb

import (
	"fmt"

	"github.com/koyeb/koyeb-cli/pkg/koyeb/errors"
	"github.com/spf13/cobra"
)

func (h *ServiceHandler) Logs(ctx *CLIContext, cmd *cobra.Command, args []string) error {
	service, err := h.ResolveServiceArgs(ctx, args[0])
	if err != nil {
		return err
	}

	serviceDetail, resp, err := ctx.Client.ServicesApi.GetService(ctx.Context, service).Execute()
	if err != nil {
		return errors.NewCLIErrorFromAPIError(
			fmt.Sprintf("Error while retrieving the service `%s`", args[0]),
			err,
			resp,
		)
	}

	logsType := GetStringFlags(cmd, "type")
	serviceId := serviceDetail.Service.GetId()
	deploymentId := ""
	instanceId := GetStringFlags(cmd, "instance")

	if logsType == "build" {
		latestDeploy, resp, err := ctx.Client.DeploymentsApi.ListDeployments(ctx.Context).
			Limit("1").ServiceId(service).Execute()
		if err != nil {
			return errors.NewCLIErrorFromAPIError(
				fmt.Sprintf("Error while listing the deployments of the service `%s`", service),
				err,
				resp,
			)
		}
		if len(latestDeploy.GetDeployments()) == 0 {
			return &errors.CLIError{
				What: "Error while fetching the logs of your service",
				Why:  "we couldn't find the latest deployment of your service",
				Additional: []string{
					"Your service exists but has not been deployed yet",
				},
				Orig:     nil,
				Solution: "Try again in a few seconds. If the problem persists, please create an issue on https://github.com/koyeb/koyeb-cli/issues/new",
			}
		}
		deploymentId = *latestDeploy.GetDeployments()[0].Id
	}
	logsQuery, err := ctx.LogsClient.NewWatchLogsQuery(
		logsType,
		serviceId,
		deploymentId,
		instanceId,
	)
	if err != nil {
		return err
	}
	return logsQuery.PrintAll()
}
