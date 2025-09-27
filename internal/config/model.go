package config

type Config struct {
	Server   Server
	Database Database
	Jwt      Jwt
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	MongoURI      string
	MongoDatabase string
}

type Jwt struct {
	Key string
	Exp int
}
