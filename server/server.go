package server

type Server interface {
	Init()
	Run()

	RegisterEndpoint(endpoint Endpoint)
}
