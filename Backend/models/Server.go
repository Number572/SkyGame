package models

type Server struct {
	Address Address
}

type Address struct {
	IPv4 string
	Port string
}
