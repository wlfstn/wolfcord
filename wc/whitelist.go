package wc

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Instance struct {
	ID   string `toml:"id"`
	Name string `toml:"name"`
}

type Whitelist struct {
	Entities []Instance `toml:"Whitelist"`
}

func collectWhitelist(wlMap *map[string]string, file string) {
	var wl Whitelist

	_, err := toml.DecodeFile(file, &wl)
	if err != nil {
		fmt.Println("Whitelist Error:", err)
		os.Exit(1)
	}

	for _, entity := range wl.Entities {
		(*wlMap)[entity.ID] = entity.Name
	}
	fmt.Printf("Whitelist: %v\n", (*wlMap))
}
