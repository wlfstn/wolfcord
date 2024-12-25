package wc

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Config struct {
	BotName  string
	KeyLoc   string
	ServerID string
	Database DatabaseConfig
}

type Command struct {
	Name        string          `toml:"name"`
	Description string          `toml:"description"`
	Options     []CommandOption `toml:"options"`
}

type CommandOption struct {
	Name         string          `toml:"name"`
	Description  string          `toml:"description"`
	Type         int             `toml:"type"`
	Required     bool            `toml:"required"`
	Autocomplete bool            `toml:"autocomplete"`
	Options      []CommandOption `toml:"options"`
}

type AuthUserData struct {
	ID   string `toml:"id"`
	Name string `toml:"name"`
}

type AuthUsersList struct {
	Entities []AuthUserData `toml:"Whitelist"`
}
