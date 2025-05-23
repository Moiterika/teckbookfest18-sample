# Project coding rules

## Standard

- Write Japanese comment
- If file name or identifier name are Japanese word, don't rename to English or "roma-ji".
- Don't add any package from github.com or go.dev.
  - If you need to add a package, ask me first.

## Architecture

- This project has the directory structure.
  - cmd
    - mhg
      - mhg is the ManHour Grouping tool.
    - tbf18
      - main.go is the application layer.
  - domain
    - entity files
    - interface files
      - i.e. reading/writing file interface
  - io
    - excel reader/writer implement
    - csv reader/writer implement

## Naming

- struct fields
  - must start with Fld
    - FldHoge
- Entity
  - file name is snake_case.
    - hoge_entity.go
  - struct name is PascalCase.
    - EntHoge
      - If Hoge is Japanese word, don't rename to English or "roma-ji".
- Repository
  - file name is snake_case.
    - hoge_repository.go
  - struct name is PascalCase.
    - RepoHoge
      - If Hoge is Japanese word, don't rename to English or "roma-ji".

## Answer

- Speak in Japanese