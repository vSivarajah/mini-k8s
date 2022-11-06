package main

type Manager struct {
	//Client enables you to communicate with mini-k8s api-server
	apiClient APIServerClient

	//probably needs to some hashmaps for workersqueue
	//also need some other maps to store other important information
	//maybe a caching system?
}

func NewManager(host string) *Manager {
	client := NewAPIServerClient(host)

	return &Manager{
		apiClient: *client,
	}
}
