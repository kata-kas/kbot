package setup

type Server struct {
	ID          string
	Environment string
}

var SERVERS = []Server{
	{
		ID:          "1234567890",
		Environment: ".env",
	},
}

func GetServer(id string) *Server {
	for _, server := range SERVERS {
		if server.ID == id {
			return &server
		}
	}

	return nil
}
