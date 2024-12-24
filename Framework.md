## Welcome to the wolfcord framework
This will be a brief markdown file on how to get started, getting your project up and running.

## Project Setup
- 3 toml configs
-- config.toml
-- authorized.toml
-- directory of command.toml

## Resource Loading
- Load bot command files `ResourceLoadCommandFiles("folder" string)`
- Load SQL files `ResourceLoadSQL("file location" string)`
- Load files `ResourceLoadFile("file location" string)`


# Framework Purpose
- Structure bot embedded bot commands with TOML files.
- Have embedded commands be compared to remove, add, and update commands.
- Have bot and database configured from a TOML file.