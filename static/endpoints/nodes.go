package endpoints

const (
	Nodes                               Endpoint = "/nodes"
	Nodes_Node_Qemu                     Endpoint = "/nodes/{val}/qemu"
	Nodes_Node_Qemu_VMID_               Endpoint = "/nodes/{val}/qemu/{val}"
	Nodes_node_Qemu_VMID_CloudinitDump  Endpoint = "/nodes/{val}/qemu/{val}/cloudinit/dump"
	Nodes_Node_Qemu_VMID_StatusCurrent  Endpoint = "/nodes/{val}/qemu/{val}/status/current"
	Nodes_Node_Qemu_VMID_StatusReboot   Endpoint = "/nodes/{val}/qemu/{val}/status/reboot"
	Nodes_Node_Qemu_VMID_StatusReset    Endpoint = "/nodes/{val}/qemu/{val}/status/reset"
	Nodes_Node_Qemu_VMID_StatusResume   Endpoint = "/nodes/{val}/qemu/{val}/status/resume"
	Nodes_Node_Qemu_VMID_StatusShutdown Endpoint = "/nodes/{val}/qemu/{val}/status/shutdown"
	Nodes_Node_Qemu_VMID_StatusStart    Endpoint = "/nodes/{val}/qemu/{val}/status/start"
	Nodes_Node_Qemu_VMID_StatusStop     Endpoint = "/nodes/{val}/qemu/{val}/status/stop"
	Nodes_Node_Qemu_VMID_StatusSuspend  Endpoint = "/nodes/{val}/qemu/{val}/status/suspend"
	Nodes_node_Qemu_VMID_Config         Endpoint = "/nodes/{val}/qemu/{val}/config"
	Nodes_Node_Status                   Endpoint = "/nodes/{val}/status"
	Nodes_Node_Time                     Endpoint = "/nodes/{val}/time"
	Nodes_Node_Version                  Endpoint = "/nodes/{val}/version"
	Nodes_Node_Config                   Endpoint = "/nodes/{val}/config"
	Nodes_Node_DNS                      Endpoint = "/nodes/{val}/dns"
	Nodes_Node_Migrateall               Endpoint = "/nodes/{val}/migrateall"
	Nodes_Node_Wakeonlan                Endpoint = "/nodes/{val}/wakeonlan"
)
