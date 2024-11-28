# Platform Service in Golang

## Tech Stack

- Golang
- Go-chi
- GORM
- Postgres
- Docker
- Makefile

## TODO

- [ ] restructure project
- [x] rename package and project solution, go mod of `exampleFirst`
- [x] add and use GORM
- [x] Clean and DI (using Dependency Injection with Factory Mehtod NewInstance)
- [ ] add api versioning URI endpoint for all: api/v1/
- [ ] use UUID instead of just ID in models
- [ ] golang air (autoreload) in docker, air vs docker watch (is it possible to use air in production in docker)
- [ ] add api documentation with Swagger/OpenAPI
- [ ] write Unit Tests, also Integrative Tests
- [ ] validate of GORM model when encoding `json` with `validate` like in Dotnet Data Annotations and Pydantic Schema in Python (platforms need `required`)
- [ ] add middleware logger and error levels, proper error handling
- [ ] be more clear storage (database) postgres (strcuts and interfaces) instead of `Service` make it concise/clear by `DbService` interface, `dbService` struct
- [ ] also this `Service` interface and struct, move to storage package!

## Q/A

- gin vs go-chi
- diference between type Something struct { db database.Service } vs { DB database.Service } (i mean between lowercase and uppercase in struct members???)
- directory naming (plural or not)???
- gorm vs raw sql in golang projects
- golang struct vs database model json struct in project???
- air vs docker compose watch
- air in production???
