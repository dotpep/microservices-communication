# Microservices Communication in Dotnet Implementation

- Dotnet C# (will be switched/implemented on Golang)
- Two Services
- Event Driven Architecture
- SQL and InMemory (Cache) Databases
- API Gateway
- REST API
- gRPC (sync)
- Docker
- K8s
- RabbitMQ (async)

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

**DB Context:**

**Repository:**

Uses Interface concrete class pattern,
We will inject our repository through Dependency Injection (DI), in startup class.

Interface is Contract between apps/services and specifies what type of things,
method signatures effectively that our repository will support,
and then any concrete class can come along and implement those interfaces.

**DTO's:**

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

## Solution Architecture (System Design and Software/Service Architecuture)

Schema, Diagram, UML img's:

## Commands

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
