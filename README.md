# wolfcord
Framework built on top of toml, discordgo, &amp; pgx5.

## TOML Configs
- Whitelist
-- "./resources/user_wl.toml"
-- "./resources/channels_wl.toml"
-- "config.toml"

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