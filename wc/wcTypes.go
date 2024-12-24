package wc

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

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Config struct {
	KeyLoc    string
	ServerID  string
	ChannelID string
	Database  DatabaseConfig
}

type AuthUserData struct {
	ID   string `toml:"id"`
	Name string `toml:"name"`
}

type AuthUsersList struct {
	Entities []AuthUserData `toml:"Whitelist"`
}
