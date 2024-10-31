# Microservices Communication in Dotnet Implementation

**Tech Stack:**

- Dotnet C# (will be switched/implemented on Golang this two services)
- Two Services
- Event Driven Architecture
- SQL and InMemory (Cache) Databases
- REST API
- API Gateway
- gRPC (sync)
- RabbitMQ (async)
- Docker
- K8s

**SRC's:**

- [.NET Microservices YT](https://youtu.be/DgVjEo3OGBI?si=SRhBSwyhBf85bbRg)

## Demonstrations, Solution Architecture

`System Design and Software/Service Architecuture`

Schemas, Diagrams, UMLs, OpenAPI/Swagger endpoints docs img's:

## Architecture

### Event Driven Architecture

**Component Layer of EDA:**

- Models
- DTO's
- DB Context
- Repository

---

**Event Driven Architecture (Data):**

- (conn: DB Context) Models - mapped - (conn: Models) DTOs
- (conn: SQL DB) DB Context - read/write - (conn: DB Context) Repository (goes: Outside)
- Models is internal and DTOs is extertal.

**Main Components:**

- Models - internal data representation
- DB Context - mediates models down to persistent layer in Database (SQL Server)
- DTO's - extensive use of Data Transfer Objects - external representation of our model data and they're mapped to models
- Repository (Pattern) - just abstract away our DB Context implementation
- REST API (Controller - sync, in) we will say externally reach into a repository pull back any data and then send back HTTP response using DTO's, it is our external Contract going out to external Consumers
- Repository Read/Write with DB Context

---

**Models:**

Models with property, data annotations that will be auto validated.

**DB Context:**

DB Context that will connect context to InMemory DB with custom AppDbContext that inherit of DbContext with constructor and property of DbSet with Platform (models) class data type as plural Platforms.

**Repository:**

Uses Interface concrete class pattern,
We will inject our repository through Dependency Injection (DI), in startup class.

Interface is Contract between apps/services and specifies what type of things,
method signatures effectively that our repository will support,
and then any concrete class can come along and implement those interfaces.

**DB Preparation:**

It is for testing in our In-Memory DB, with some generated and populated data in DB.

Put some data to database.
We use In-Memory database as temporary database and it is used in tests.

**DTO's:**

DTO's (Data Transfer Object) - external represantataion of our internal models.

It is like contract for other services,
if you like will change your Models,
you will terminate this contract with other services like API's.
Abstract your external and internal, like using Interfaces but for DB and data.

- data privacy
- contractual coupling

Map DTO's:

- Model to PlatformReadDto
- PlatformCreatDto to Model

### Other Technologies and Tools

**REST API:**

---

**HTTP:**

---

**gRPC:**

---

**API Gateway:**

---

**RabbitMQ:**

---

**Docker and K8s:**

## Running Instructions

## Commands Log

### DotNET

`dotnet --list-sdks`
`dotnet --version`

`dotnet new webapi -n PlatformService -f net5.0` - create new webapi project
webapi is a template
-n flag is name of service/app
-f flag is stands for framework, you can specify version of dotnet like net7.0, net8.0

`code -r .\PlatformService\` - open this dir in vs code recursivly

`dotnet new globaljson` - set version of dotnet that will be used in this folder
change global.json (version to `dotnet --list-sdks` version):

```json
{
  "sdk": {
    "version": "5.0.408"
  }
}
```

`dotnet add package PackageName -v 5.0.8`
-v flag is --version

### Golang
