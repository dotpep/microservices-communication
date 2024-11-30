# Microservices Communication in Dotnet Implementation

**Tech Stack:**

- Dotnet C# (will be switched/implemented on Golang this two services)
- Two Services
- Event Driven Architecture
- SQL and InMemory (Cache) Databases
- REST API
- API Gateway (Ingress Nginx Controller, Ingress Nginx Load Balancer)
- gRPC (sync)
- RabbitMQ (async)
- Docker
- K8s

**SRC's:**

- [.NET Microservices YT](https://youtu.be/DgVjEo3OGBI?si=SRhBSwyhBf85bbRg)

## TO DO

-
- Come Up with other microservices for this project

---

**Golang:**

- Transition to Golang Microservices (From DotNET)
    - How to use Golang PlatformService (instead of Dotnet)
    - is it needed? need I delete them?
    - where to store Dotnet service? what is fallback?
- migrate/convert Dotnet PlatformService to Golang PlatformService (PlatformServiceGo)
- how to structure Golang project
- Golang Event-Driven Architecture

## Services

- PlatformService runs on `5000` http and `5001` https locally
- CommandsService runs on `6000` http and `6001` https locally

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

**Controller and Actions:**

Controller - API Endpoints.
For testing API, I'm using Insomnia API client.

Endpoint: `http://localhost:5000/api/platforms`

- Get Lists of All Platforms
- Get Specific One Platform by ID
- Create Platform

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

Docker is ???.
Docker Container works above OS with Docker Engine (when VM like Virtual Box uses Hypervisor to make new OS top of your Host OS).

Docker Compose is kind of a middle ground between Docker and K8s,
there is also run up muliple containers, network them together,
it is good option in development type environment/stage
but in production-wise K8s is really the option.

Kubernates is Container Orchestrator,

**Kubernates Architecture:**

`/K8s` is production deployment k8s files

K8s terminology:

- Cluster is
- Node is
    - Node port (3xxxx:80) 3xxxx is internal, 80 is external (3xxxx will be generated)
- Pod is Container Service (like our Platform Service Container) with (80:666)
- Port mapping is
    - Cluster IP
- communcation between K8s

API Gateway in K8s:

Pod with Ingress Nginx Controller...

**Questions:**

- API Gateway vs Ingress Controller vs Load Balancer vs Reverse Proxy
- Kubernates Imperative (Command Line) and Declarative (Config Files)
- K8s microservices directory structure solution

## Running Instructions

## Commands Log

### DotNET

To start:

`dotnet build`
`dotnet run`

---

To configure version:

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

---

to Fix bug:

```bash
dotnet run
Building...
--> Seeding Data...
crit: Microsoft.AspNetCore.Server.Kestrel[0]
      Unable to start Kestrel.
```

[Stackoverflow - Unable to start Kestrel when 'dotnet run'](https://stackoverflow.com/questions/57736568/unable-to-start-kestrel-when-dotnet-run)

1. `dotnet dev-certs https --clean`
2. `dotnet dev-certs https --trust`

### Golang

### Docker

Docker:

- Dockerfile --> Docker Image

push image to docker hub - `docker build -t dotpep/platformservicedotnet .`

- `docker run -p 8080:80 -d dotpep/platformservicedotnet` (-p flag port, mapping) (-d flag detached) (docker hub image)

Commands:

- `docker run` (start new id container)
- `docker ps` (what containers are in running)
- `docker stop <container-id>` (stop container by id)
- `docker start <container-id>` (restart existed container with id)
- `docker logs -f <container-id>` (check docker logs, with -f flag attached)

A lot usage:

- `docker push dotpep/platformservicedotnet`
- `docker build`

### K8s

- `kubectl version`
- `kubectl apply -f .\platforms-depl.yml` (`cd .\K8s\`)
- `kubectl get deployments`
- `kubectl get pods`
- `kubectl delete deployment <service-deployments-name>`
- `kubectl delete pod <name>`

---

If we try to delete `platformservice` with logs of web-api, k8s will automatically up it, even if we delete it.

For deletion and stoping you need to delete deployments (when running k8s there will be 2 service container in `docker ps`).

- `kubectl delete deployments platforms-depl`

---

We have service up and running on K8s cluster
but we have no way of accessing it yet.

We need to create a **Node Port** that will actually give us access to service running in k8s cluster.

- `kubectl apply -f .\platforms-np-srv.yml`
- `kubectl get services`

Node Port - 3xxxx will be generated with random!

---

Naming K8s directory .yml files:

- `platforms-depl.tml` platform service deployment
- `platforms-np-srv.tml` platform node port service

---

Second Service (CommandService):

- `docker build -t dotpep/commandservicedotnet .`
- `docker run -p 8080:80 dotpep/commandservicedotnet`

Make `appsettings.Production.json` in PlatformService that sends HttpCommandDataClient to CommandsService:

```json
{
    "CommandService": "http://commands-clusterip-srv:80/api/c/platforms"
}
```

- `commands-clusterip-srv:80` is service host in /K8s/commands-depl.yml conf file `ClusterIP` second section

Because of changes and creating new file, you need to build and push Docker Image to Hub: `docker build -t dotpep/platformservicedotnet .` and `docker push dotpep/platformservicedotnet`

---

- `kubectl get deployments`
- `kubectl rollout restart deployment platforms-depl`
- `kubectl logs <pod> -f`

---

Setting Up - Ingress Nginx Controller (API Gateway)

links:

- [github.com/kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx)
- [kubernetes.github.io/ingress-nginx/deploy/#docker-desktop](https://kubernetes.github.io/ingress-nginx/deploy/#docker-desktop)
- [raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.12.0-beta.0/deploy/static/provider/aws/deploy.yaml](https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.12.0-beta.0/deploy/static/provider/aws/deploy.yaml)

Command:

- `kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.12.0-beta.0/deploy/static/provider/aws/deploy.yaml` (in `kubernetes.github.io/ingress-nginx/deploy/#docker-desktop` link)

Additional:

- `kubectl get namespace`
- `kubectl get deployments --namespace=ingress-nginx` shows name:`ingress-nginx-controller`
- `kubectl get pods --namespace=ingress-nginx`
- `kubectl logs <pod> --namespace=ingress-nginx -f`
- `kubectl get services --namespace=ingress-nginx` shows name:`ingress-nginx-controller` and type:`LoadBalancer`

---

`ingress-srv.yml`:

- rules: - host: `microcomm.com`

Steps:

- `cd C:\Windows\System32\drivers\etc`
- `ls` shows files like `hosts`, `networks`, `protocol`, `services` ...
- `vim .\hosts` and add `127.0.0.1 microcomm.com`
- `cat .\hosts` ensure that you added it correctly in the end of file

Applying ingress-srv.yml:

- `kubectl apply -f .\ingress-srv.yml`

```bash
Warning: annotation "kubernetes.io/ingress.class" is deprecated, please use 'spec.ingressClassName' instead
```

Solution for warning:

- [kubernetes ingress service annotations, kubernetes.io/ingress.class annotation is officially deprecated:](https://stackoverflow.com/questions/64262770/kubernetes-ingress-service-annotations)

Try to Ping this endpoint:

- `http://microcomm.com/api/platforms`

---
