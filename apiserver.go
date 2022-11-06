package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIServer struct {
	ListenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		ListenAddr: listenAddr,
	}
}

func (s *APIServer) Start() {
	e := echo.New()
	e.POST("/pod/:name", s.handleCreatePod)
	e.GET("/pods", s.handleGetPods)
	e.Start(s.ListenAddr)
}

func (s *APIServer) handleCreatePod(c echo.Context) error {
	name := c.Param("name")
	containers := []*Container{
		{
			name:  "test",
			image: "testimage:123",
			ports: map[string]int64{"containerPort": 80},
		},
	}

	if err := createPod(name, containers); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]any{"pod": "created successfully"})
}

// TODO: send api Pod response object instead of Pod object
func (s *APIServer) handleGetPods(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{"pods": &pods})
}

// change argument paramters, take in a podSpec later on
func createPod(name string, containers []*Container) error {

	if name == "" {
		return errors.New("pod name must be specified")
	}
	if containers == nil {
		//
		return errors.New("atleast one container must be specified")
	}
	pod := &Pod{
		name:       name,
		containers: containers,
	}

	pods = append(pods, pod)
	fmt.Println(pods)
	return nil
}
