package teamspeak

type HostInfo struct {
	HostTimestampUtc                  string `json:"host_timestamp_utc"`
	InstanceUptime                    string `json:"instance_uptime"`
	VirtualServersRunningTotal        string `json:"virtualservers_running_total"`
	VirtualServersTotalChannelsOnline string `json:"virtualservers_total_channels_online"`
	VirtualServersTotalClientsOnline  string `json:"virtualservers_total_clients_online"`
	VirtualServersTotalMaxClients     string `json:"virtualservers_total_maxclients"`
}

func (a Api) HostInfo() (ApiResponse[[]HostInfo], error) {
	var response ApiResponse[[]HostInfo]
	err := a.get("/hostinfo", &response)
	return response, err
}
