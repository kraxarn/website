package teamspeak

type Client struct {
	ChannelId        string `json:"cid"`
	ClientId         string `json:"clid"`
	ClientDatabaseId string `json:"client_database_id"`
	ClientNickname   string `json:"client_nickname"`
	ClientType       string `json:"client_type"`
}

func (a Api) ClientList() (ApiResponse[[]Client], error) {
	var response ApiResponse[[]Client]
	err := a.get("/clientlist", &response)
	return response, err
}
