package server

import (
	"fmt"
)

// Server is just a server, includes all basic information
type Server struct {
	Host 		string
	Name 		string
	Port	 	string
	UserName 	string
	Password	string
}

// HostName returns host name in detail
func (s *Server) HostName() string {
	if s.Name != "" {
		return fmt.Sprintf("%s (%s)", s.Name, s.Host)
	} 
	return s.Host
}

// Connect makes a SSH connection
func (s *Server) Connect() {
	fmt.Printf("Connect to %s\n", s.HostName())
}

// Detail returns the server information in detail
func (s *Server) Detail() string {
	if s.Name != "" {
		return fmt.Sprintf("%s (%s@%s:%s)", s.Name, s.UserName, s.Host, s.Port)
	}
	return fmt.Sprintf("(%s@%s:%s)", s.UserName, s.Host, s.Port)
}

// Print will print the server info to stdout
func (s *Server) Print() {
	fmt.Printf("%s", s.Detail())
}
