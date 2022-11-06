package main

import "net/http"

// other components in the ecosystem would use this client to
// interact with the API server

type APIServerClient struct {
	//host represents the hostname of apiserver
	host   string
	client *http.Client
}

func NewAPIServerClient(host string) *APIServerClient {
	return &APIServerClient{
		host:   host,
		client: http.DefaultClient,
	}
}
