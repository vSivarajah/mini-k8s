package main

var pods []*Pod

type Pod struct {
	name       string
	containers []*Container
}

type Container struct {
	name  string
	image string
	ports map[string]int64
}
