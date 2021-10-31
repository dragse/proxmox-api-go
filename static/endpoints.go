package static

import "regexp"

type Endpoint string

var /* const */ placeholderRegex = regexp.MustCompile("{val}")

const (
	EndpointAccessTicket Endpoint = "/access/ticket"

	EndpointCluster       Endpoint = "/cluster"
	EndpointClusterStatus Endpoint = "/cluster/status"
	EndpointClusterTasks  Endpoint = "/cluster/tasks"

	EndpointNodes                               Endpoint = "/nodes"
	EndpointNodes_Node_Qemu                     Endpoint = "/nodes/{val}/qemu"
	EndpointNodes_Node_Qemu_VMID_               Endpoint = "/nodes/{val}/qemu/{val}"
	EndpointNodes_node_Qemu_VMID_CloudinitDump  Endpoint = "/nodes/{val}/qemu/{val}/cloudinit/dump"
	EndpointNodes_Node_Qemu_VMID_StatusCurrent  Endpoint = "/nodes/{val}/qemu/{val}/status/current"
	EndpointNodes_Node_Qemu_VMID_StatusReboot   Endpoint = "/nodes/{val}/qemu/{val}/status/reboot"
	EndpointNodes_Node_Qemu_VMID_StatusReset    Endpoint = "/nodes/{val}/qemu/{val}/status/reset"
	EndpointNodes_Node_Qemu_VMID_StatusResume   Endpoint = "/nodes/{val}/qemu/{val}/status/resume"
	EndpointNodes_Node_Qemu_VMID_StatusShutdown Endpoint = "/nodes/{val}/qemu/{val}/status/shutdown"
	EndpointNodes_Node_Qemu_VMID_StatusStart    Endpoint = "/nodes/{val}/qemu/{val}/status/start"
	EndpointNodes_Node_Qemu_VMID_StatusStop     Endpoint = "/nodes/{val}/qemu/{val}/status/stop"
	EndpointNodes_Node_Qemu_VMID_StatusSuspend  Endpoint = "/nodes/{val}/qemu/{val}/status/suspend"
	EndpointNodes_node_Qemu_VMID_Config         Endpoint = "/nodes/{val}/qemu/{val}/config"
	EndpointNodes_Node_Status                   Endpoint = "/nodes/{val}/status"
	EndpointNodes_Node_Version                  Endpoint = "/nodes/{val}/version"
	EndpointNodes_Node_Time                     Endpoint = "/nodes/{val}/time"
	EndpointNodes_Node_Config                   Endpoint = "/nodes/{val}/config"
	EndpointNodes_Node_DNS                      Endpoint = "/nodes/{val}/dns"
	EndpointNodes_Node_Migrateall               Endpoint = "/nodes/{val}/migrateall"
	EndpointNodes_Node_Wakeonlan                Endpoint = "/nodes/{val}/wakeonlan"

	EndpointVersion Endpoint = "/version"
)

func (endpoint Endpoint) FormatValues(val ...string) Endpoint {
	i := 0
	format := placeholderRegex.ReplaceAllStringFunc(string(endpoint), func(s string) string {
		return val[i]
	})

	return Endpoint(format)
}
