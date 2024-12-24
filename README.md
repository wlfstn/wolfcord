# wolfcord
Framework built on top of toml, discordgo, &amp; pgx5.

## TOML Configs
- "config.toml"

## TOML config template
```toml
KeyLoc = '/var/.tokens/potatobot'
ServerID = '981337133713379898' #potato server

[Database]
host = 'potatodomain.dog'
port = 5432
user = 'hellopeski'
password = '/var/.tokens/potatoPostgres'
dbname = 'potato'

[[Whitelist]]
id = '110098133712340001'
name = 'Sawyer'

[[Whitelist]]
id = '110098133712340002'
name = 'Gavin'
```

## TOML Cmds examples
```toml
name = 'dj_logo'
description = 'Returns the logo of the DJ'

[[options]]
name = 'name'
description = 'The name of the DJ'
type = 3
required = true
autocomplete = true
```

```toml
name = 'event'
description = 'View iwait events'

[[options]]
name = 'list'
description = 'List the 10 most recent events'
type = 1

[[options]]
name = 'last'
description = 'List the lineup of the last event'
type = 1

[[options]]
name = 'date'
description = 'The date of the event'
type = 1
	[[options.options]]
	name = 'day'
	description = 'The date of the event'
	type = 3
	required = true
	autocomplete = true
```