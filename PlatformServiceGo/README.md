# Platform Service in Golang

## TODO

- [ ] restructure project
- [ ] rename package and project solution, go mod of `exampleFirst`
- [ ] add and use GORM
- [ ] add middleware logger and error levels, proper error handling
- [ ] be more clear storage (database) postgres (strcuts and interfaces) instead of `Service` make it concise/clear by `DbService` interface, `dbService` struct
- [ ] golang air (autoreload) air vs docker wathc (is it possible to use air in production in docker)
- [ ] also this `Service` interface and struct, move to storage package!
- [ ] add api versioning URI endpoint for all: api/v1/
- [ ] add api documentation with Swagger/OpenAPI

## Q/A

- gin vs go-chi
- diference between type Something struct { db database.Service } vs { DB database.Service } (i mean between lowercase and uppercase in struct members???)
- directory naming (plural or not)???
- gorm vs raw sql in golang projects
- golang struct vs database model json struct in project???
