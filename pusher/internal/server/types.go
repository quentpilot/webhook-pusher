package server

type HttpDriver interface {
	Load() // Loads routes handler
	Run()  // Runs the HTTP server
}

type HttpConfig struct {
	Host string
	Port uint
}
