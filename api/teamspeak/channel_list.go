package teamspeak

type Channel struct {
	ChannelName                 string `json:"channel_name"`
	ChannelNeededSubscribePower string `json:"channel_needed_subscribe_power"`
	ChannelOrder                string `json:"channel_order"`
	ChannelId                   string `json:"cid"`
	ParentId                    string `json:"pid"`
	TotalClients                string `json:"total_clients"`
}

func (a Api) ChannelList() (ApiResponse[[]Channel], error) {
	var response ApiResponse[[]Channel]
	err := a.get("/1/channellist", &response)
	return response, err
}
