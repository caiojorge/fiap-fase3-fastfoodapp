package server

type Server interface {
	Initialization() Server
	Run(port string)
}
