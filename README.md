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
- Event Bus

**SRC's:**

- [.NET Microservices YT](https://youtu.be/DgVjEo3OGBI?si=SRhBSwyhBf85bbRg)

## TO DO

- Come Up with other microservices for this project
- [ ] Database ER-D Schemas
- [ ] API Documentation (Swagger/OpenAPI)
- [ ] Architectural Solution/Microservices Communication (images of Schemas...)

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

- PlatformService
- CommandsService

Services Response Example:

```json
```

## Demonstrations, Solution Architecture

`System Design and Software/Service Architecuture`

Schemas, Diagrams, UMLs, OpenAPI/Swagger endpoints docs img's:

## Documentation

### API Endpoints (Dotnet)

**PlatformService:**

- GET: `/api/platforms` Get Lists of All Platforms
- GET: `/api/platforms/1` Get Specific One Platform
- POST: `/api/platforms/` Create Platform

---

**CommandsService:**

- POST: `/api/c/platforms/` Test Inbound Connection (to PlatformService)

- GET: `/api/c/platforms` Get All Platforms
- GET: `/api/c/platforms/{platformId}/commands` Get All Commands for a Platform
- GET: `/api/c/platforms/{platformId}/commands/{commandId}` Get a Specific Command for a Platform
- POST: `/api/c/platforms/{platformId}/commands/` Create a Command for a Platform

### Database Entities/Schemas (Dotnet)

**PlatformDatabase:**

- Platform
    - Id
    - Name
    - Publisher
    - Cost

---

**CommandsDatabase:**

- Command
    - Id
    - HowTo
    - CommandLine
    - PlatformId
    - Platform (obj)

- Platform
    - Id
    - ExternalId
    - Name
    - Commands (obj)

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

Uses JSON in response, Have endpoints or URI's.
Works under HTTP (TCP) and have Header, Body, HTTP Status code and Use GET, POST, PUT, PATCH and DELETE HTTP methods.

---

**HTTP:**

Client requests to REST API.

---

**gRPC:**

---

**API Gateway:**

Ingress Nginx Controller and Load Balancer in K8s cluster.

---

**RabbitMQ (Message Bus) Architecture:**

`PlatformService` will `Publish` a **message** onto in **RabbitMQ Message Bus**,
and it doesn't care who is **listening** for those **messages** it's just putting
the information out there, it might be one service, or it might me 100 services, might be none,
doesn't matter it doesn't care,
and then on the `CommandService` side, we are going to `Subscribe` for those **events**.

**RabbitMQ Overview:**

RabbitMQ is Message Bus/Message Broker with Queue/Buffer, LIFO, Publisher/Subscriber or Producer/Consumer of Events, Tasks, Messages - Asyncronously.

- A Message Broker - it accepts and forwards messages.
- Messages are sent by Producers (or Publishers)
- Messages are received by Consumers (or Subscribers)
- Messages are stored on Queues (essentially a message buffer) (messages can be stored in the queue, so it has some degree of persistence, we're not going to be using that in our solution, so if our message bus crashes then we use or lose all our messages, but in production type environment you wouldn't do that you would allow messages to persist in the event of any failures)
- Exchanges can be used to add a degree of "routing" functionality (we're going to be using an exchage, but going to be routing)
- RabbitMQ Uses - Advanced Message Queuing Protocol (AMQP) and others

Idea: Messangers are published onto the queue if your services are overwhelmed and can't actually service those requests the message broker acts as a buffer for those messages and then as and when you bring more services online they kind of chew through the messages on the queue.

**Exchanges in RabbitMQ:**

4 Types of Exchange:

- Direct Exchange
- Fanout Exchange (we're going to use)
- Topic Exchange
- Header Exchange

**Direct Exchange:**

> Service 1 | Message Broker | Service 2

`Publusher (Publish RK="somekey") --> Exchange (Routes to: "somekey") -> Queue ||| --> (Consume) Consumer`

- Delivers Messages to queues based on a routing key
- Ideal for "direct" or unicast messaging

**Fanout Exchange:**

> Service 1 | Message Broker | Service 2

`Publusher (Publish) --> Exchange -> Queue 1 |||, Queue 2 ||| --> (Consume) Consumer 1, (Consume) Consumer 2, (Consume) Consumer 3`

- Delivers Messages to all Queues that are bound to the exchange
- It ignores the routing key
- Ideal for broadcas messages

Doesn't care who is listening just put messages to Queue.

**RabbitMQ in K8s:**

- K8s RabbitMQ .yml file configuration is not a Production Quality or Production Class deployment.

---

**Docker and K8s:**

**Docker Overview:**

Docker is ???.
Docker Container works above OS with Docker Engine (when VM like Virtual Box uses Hypervisor to make new OS top of your Host OS).

Docker Compose is kind of a middle ground between Docker and K8s,
there is also run up muliple containers, network them together,
it is good option in development type environment/stage
but in production-wise K8s is really the option.

**Kubernates Overview:**

Kubernates is Container Orchestrator, ...

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

- `kubectl apply -f .\commands-depl.yml`

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

Creating Database Storage Pod in K8s cluster:

- `kubectl get storageclass`

Three concepts of Storage K8s:

1. Persistent Volume Claim
2. Persistent Volume
3. Storage Class

`local-pvc.yml`:

- `kubectl apply -f .\local-pvc.yml`
- `kubectl get pvc`

---

Creating K8s secrets:

- `kubectl create secret generic mssql --from-literal=SA_PASSWORD="pas55w0rd!"`
- `kubectl create secret generic <mssql-name> --from-literal=<SA_PASSWORD-key>="<pas55w0rd!-value>"`

---

`mssql-plat-depl.yml`:

```yml
...
          env:
          - name: MSSQL_PID
            value: "Express"
          - name: ACCEPT_EULA
            value: "Y"
          - name: SA_PASSWORD
            valueFrom:
              secretKeyRef:
                name: <mssql-name>
                key: <SA_PASSWORD-key>
```

1. Deployment `msql`
2. Service `mssql-clusterip-srv` type:`ClusterIP`
3. Service `mssql-loadbalancer` type:`LoadBalancer`

- `kubectl apply -f .\mssql-plat-depl.yml`
- `kubectl get services`
- `kubectl get pods`

---

Connection to MsSQL (Microsoft SQL Server)

SQLTools and SQLTools Microsoft SQL Server driver (VS Code extentions to connect) (search: `@tag:sqltools-driver mssql` for driver) (SQLTools and MsSQL driver by Matheus Teixeira):

Connection string:

```txt
Server=localhost,1433;Database=Master;User Id=SA;Password=pas55w0rd!
```

Command:

- `docker ps`
- `docker exec -it <fa8cb70ec39f-CONTAINER-ID-or-Name> /opt/mssql-tools/bin/sqlcmd -S localhost -U sa`

MsSQL type commands:

1. select name from sys.databases;
2. go

- quit

---

`/PlatformService/appsettings.Production.json`:

- Do not use user as default `SA`/`sa` in `User Id=SA` in Production staging configuration, connection!

```json
{
    "CommandService": "http://commands-clusterip-srv:80/api/c/platforms",
    "ConnectionStrings":
    {
        "PlatformsConn": "Server=mssql-clusterip-srv,1433;Initial Catalog=platformsdb;User Id=SA;Password=pas55w0rd!;"
    }
}
```

---

Migrations of SQLServer in Production and InMemory Seeding data in Development

- `dotnet ef migrations add initialmigration`

If you get error with `file was not found.` e.g. for `dotnet ef`:

- `dotnet tool install --global dotnet-ef --version 5.*`

Link: [Command dotnet ef not found - Stackoverflow](https://stackoverflow.com/questions/57066856/command-dotnet-ef-not-found)

InMemory Database do not support migrations and we need to do some trick!
(We are using InMemory DB for Development and SQL Server as for Production)

in `PlatformService/Startup.cs` - (comment it when applying Migrations and uncomment when we done with this, and use it for Development):

```cs
            // Database
            //if (_env.IsProduction())
            //{
                Console.WriteLine("--> Using SqlServer Db");
                services.AddDbContext<AppDbContext>(opt =>
                    opt.UseSqlServer(Configuration.GetConnectionString("PlatformsConn"))
                );
            //}
            //else
            //{
            //    Console.WriteLine("--> Using InMem Db");
            //    services.AddDbContext<AppDbContext>(opt =>
            //        opt.UseInMemoryDatabase("InMem")
            //    );
            //}
```

and also:

```cs
            // Preperation DB
            //PrepDb.PrepPopulation(app, env.IsProduction());
```

it will look like after this command (`dotnet ef migrations add initialmigration`):

```bash
Build started...
Build succeeded.
--> Using SqlServer Db
--> CommandService Endpoint http://localhost:6000/api/c/platforms
Done. To undo this action, use 'ef migrations remove'
```

- make docker build for image and push it
- and kubectl restart of platforms-depl deployment

Check with:

- `docker ps` find mssql Container_ID
- `docker exec -it <e9dd6791525e-Container_ID> /opt/mssql-tools/bin/sqlcmd -S localhost -U sa`
- in container Mssql sqlcmd: `select name from sys.databases;`
- and write `go`
- this will show databases in sys and also our migrated databases called `platformsdb`

Additional check in Container MsSQL DB:

- `USE platformsdb;`
- `go`
- `SELECT * FROM Platforms;`
- `go`

---

Error Management and Kill Bad Deployment:

- `kubectl delete deployment platforms-depl`

---

RabbitMQ in K8s:

- `kubectl apply -f .\rabbitmq-depl.yml`
- `kubectl get deployments`, `kubectl get pods`
- `kubectl get services` show LoadBalancer and ClusterIP of rabbitmq
- to see logs: `kubectl get pods` or `docker ps` get Container_ID and `kubectl logs <pod-name> -f` or `docker logs <container-id> -f`
- Go to: `http://localhost:15672/` for Management Interface

RabbitMQ Management Interface Default Login Credentials:

- username: `guest`
- password: `guest`

For Test Publisher (in PlatformService):

- run `make run` or `dotnet run` in `PlatformService`
- request with POST to this endpoint: `http://localhost:5000/api/platforms`
- make a lot of requests (you can comment SyncClient instruction in `Controllers/PlatformsController.cs` and in `HttpPost`/`CreatePlatform()` method) and see difference between delay of sync and async.
- also see graph in `http://localhost:15672/` RabbitMQ Management Interface

```cs
// POST: /api/platforms
[HttpPost]
public ActionResult<PlatformReadDto> CreatePlatform(PlatformCreateDto platformCreateDto){
      // Message Broker Client - Send Async Message
    platformPublishedDto.Event = "Platform_Published";
}
```

- `platformPublishedDto` has Event property to store which event it is (definition of Event in string format)
- and for CreatePlatform() controller it will be setted as `Event="Platform_Published"` in PlatformPublishedDto.

For Test Listener (Consumer) (in CommandsService):

---

RabbitMQ theory:

- Message Bus Publisher (PlatformService)
- Event Processor/Processing
- Event Listener (CommandsService)

---

Dependency Injection Reasons - when you register your services in config services,
they have what's called a service lifetime (they have different lifetimes):

1. Singleton - created 1st time requested, subsequent requests use the same instance (registered services exist for lifetime of the application).
2. Scoped - same within a request but created for every new request (exist kind of for every kind of session).
3. Transient - new instance provided everytime, never the same/reused (exist ones every request).

When we come on to creating our Listening Service that is going to be created as a Singleton service,
ultimetly it's going to be ther for the lifetime of our application.
That service is going to call this service and in order for that to happen in order for us to inject this Event Processor service into our Listening service it too has to have a lifetime the same or greator than the service it's being injected to so this service is going to have to be a Singleton service.
We're going to have to inject it, well we're going to have to create it a reference to it in another way.

## Step by step

### Development (Localhost)

- PlatformService runs on `5000` http and `5001` https locally
- CommandsService runs on `6000` http and `6001` https locally

### Production (K8s Cluster)

1. `kubectl apply -f .\platforms-depl.yml`
2. `kubectl apply -f .\platforms-np-srv.yml`
3. `kubectl apply -f .\commands-depl.yml`
4. `kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.12.0-beta.0/deploy/static/provider/aws/deploy.yaml`
5. `kubectl apply -f .\ingress-srv.yml`
6. `kubectl apply -f .\local-pvc.yml`
7. `kubectl apply -f .\mssql-plat-depl.yml`
8. `kubectl apply -f .\rabbitmq-depl.yml`

---

Checking:

1. `kubectl get deployments`
2. `kubectl get pods`
3. `kubectl get services`
4. `kubectl get pvc`
5. `kubectl get namespace`
6. `kubectl get services --namespace=ingress-nginx`

---

Updating Code:

`platformservicedotnet` Or `commandservicedotnet`

- `docker build -t dotpep/platformservicedotnet .`
- `docker run -p 8080:80 dotpep/platformservicedotnet`
- `docker push dotpep/platformservicedotnet`

## Makefile
