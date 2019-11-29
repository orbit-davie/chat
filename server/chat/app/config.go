package main

type Config struct {
	Server    	ServerInfo  `toml:"server_info"`
	App 		App 		`toml:"app"`
	Mongo 		Mongo		`toml:"mongodb"`
	Address 	Address 	`toml:"address"`
}

type ServerInfo struct {
	Name 	string	`toml:"name"`
	Id   	string	`toml:"id"`
}

type App struct {
	Environment string `toml:"environment"`
}

type Mongo struct {
	Url      string `toml:"url"`
	Database string `toml:"database"`
}

type Address struct {
	Addr        string	`toml:"host"`
	Port		string	`toml:"port"`	// :50088
}
