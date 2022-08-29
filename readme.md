# Disclaimer:
> **_NOTE:_**  This is my one of the firs t encounter with go and gRPC so the quality of code may be very poor. It may contains traces of my C# and OOP based background.

# GRPC Mediator Service
This service is a gateway to provide gRPC interface to non-gRPC APIs e.g. HTTP+JSON. Multiple handlers of the request and response can be added to  the pipeline e.g. logging, checking permissions, calling backend API to process request and return response.

```mermaid
graph TB;
    client-- 1: protobuf request -->mediator_service;
    mediator_service-- 2: read proto descriptros -->schema_registry[(schema registry)]
    mediator_service-- 3: json request -->target_service;
    target_service-- 4: json response -->mediator_service
    mediator_service-- 5: protobuf response --> client;
```

The service uses a concept of handlers employs [Open-Closed principle](https://en.wikipedia.org/wiki/Open%E2%80%93closed_principle) to allow extend behavior without modifying existing code. Handlers are a series of operations that will be performed on the request and response in [chain of responsibility pattern](https://en.wikipedia.org/wiki/Open%E2%80%93closed_principle). More handlers can be added with a very small change (addition to the switch statement) to the pipeline code. As I learn more go, I would like to improve this part to remove any need to change the pipeline code.

## Handlers
Idea of handler is based on the ASP.NET Core middleware components that receives HTTPContext to process request and response parts. There is an interface defined that handlers must adhere to by implementing a "Process" method. This method receives context containing a value for EndPointContext.

```mermaid
flowchart TB;
    client-- 1: protobuf request -->mediator_service;
    mediator_service-- 2: read proto descriptros -->schema_registry[(schema registry)]
    mediator_service-- 3: endpoint context -->handlers{for each handlers};
    handlers -- 4: endpoint context --> mediator_service
    mediator_service-- 5: protobuf response --> client;
    handlers -- 3.1: endpoint context --> auth_handler[\handler 1 e.g.Authorisation handler\]
    auth_handler -- 3.1.1 endpoint context or error --> handlers
    handlers -- 3.2: endpoint context --> http_handler[\handler 2 e.g.HTTP+JSON handler\]
    http_handler -- 3.2.1: http json request --> target_service
    target_service-- 3.2.2: http json response --> http_handler
    http_handler -- 3.2.3: endpoint context --> handlers

```

```go
type Handler interface {
	Process(epCtx context.Context) error
}
```
Handler creation is performed by a method to abstract the creation from the service running the code. This can be enhanced to implement complete [Factory Method pattern] (https://en.wikipedia.org/wiki/Factory_method_pattern) when required

### Endpoint Context
Mediator service prepares the EndpointContext based on the mediator request fields and target proto options. Endpoint Context contains set of fields required by handlers to process the request and response:

```go
type EndpointContext struct {
	EndpointDescriptor protoreflect.MethodDescriptor //Method descriptor (reflection) for the target endpoint
	EndpointConfig     *gpb.EndpointConfig //Endpoint configuration options defined by the target endpoint
	Handlers           []Handler //Array of handlers appended in the required order of execution based on the EndpointConfig
	Request            *GRequest //Contains reference to the request message. Each handler can read, validate or enrich it
	Response           *GResponse //Contains reference to the response message. Each handler can read, validate or enrich it
}
```

This repository currently includes 3 handlers

1. HTTP Handler : Calls HTTP API by marshaling request to JSON and un-marshaling response from JSON
2. Entitlements Handler: Authorisation check
3. File Handler: For unit tests, it reads response from a local file with a static JSON content


### High Level Flow

1. Team defines custom protos (CP) for a backend service (BS)
    1. CP provides config for various handlers i.e. json on http
2. Client calls mediator service (GS) with CP request
3. GS parses CP and calls configured it to handlers
    1. http_backend converts CP to JSON and calls the backend service and returns converts JSON to CP and returns to GS
4. GS returns response in CP to Client

## Mediator Service

Mediator service accepts custom protos and returns custom protos based on the stream method. This service serves for all the service definition specified in its config and with protos available into its schema registry.
Mediator service expects the custom protos to decorate the endpoint with options configured in the gprotos.


At the centre of this service is the proto definition of the endpoints. A required set of handlers can be added to the Method definition for the endpoints in the proto file.

[US stats service endpoint proto message](/schemas/usstats/usstats.proto)

```protobuf
service StatsAPI {
  rpc GetStats(GetStatsRequest) returns (GetStatsResponse) {
    option (gservice.endpoint_config) = {
        entitlement_operations: ["appointments:read"],
        handlers: [
          {name:"entitlements",  options: [{key:"1", value:"appointments:read"}]},
          {name:"http-backend",  options: [
             {key:"auth_type",value:"JWT"},
               {key:"http_method",value:"GET"},
               {key:"url_pattern",value:"api/data?drilldowns={drilldowns}&measures={measures}"},
               {key:"host_config_key",value:"US-STATS"},
               {key:"body",value:"US-STATS"}
          ]}
        ],
    };
  }
}
```

Above proto definition provides an endpoint called "GetStats". It also specifies two handlers in the order of required execution i.e. check entitlements and target HTTP APi with type of authentication to use.

## Schema Registry

This project uses a very basic file-based schema registry. Use the following protoc command to generate the file descriptors for the target proto files including all its imports
```bash
protoc --include_imports --descriptor_set_out="<target path>/<file-name>.proto-registry.pb" -I<import files path> <source path>/<file-name>.proto
```

Store the output file to the local schemas\register folder. The service uses "protoregistry" to read the descriptors and parse into dynamic proto message using "dynamicpb"


```go
	tmpFile := "../schemas/register/" + filename + "-registry.pb"

	marshalledDescriptorSet, err := ioutil.ReadFile(tmpFile)
	if err != nil {
		return nil, err
	}
	descriptorSet := descriptorpb.FileDescriptorSet{}
	err = proto.Unmarshal(marshalledDescriptorSet, &descriptorSet)
```


## How to define new target
#### 1. Define protobuf for the target service
#### 2. Specify required handlers in the endpoint options

#### 3. Generate descriptor files and add to the schema registry

Generate descriptor files and store the output in the schema registry file. Following commands generate a descriptors for a sample file and add them to the registry folder

- go to med8r folder and run following command
- generate file descriptor set for stats.proto and usstats

```bash
cd med8r
protoc --include_imports --descriptor_set_out="schemas/register/stats.proto-registry.pb" -Ischemas schemas/statsservice/stats.proto

protoc --include_imports --descriptor_set_out="schemas/register/usstats.proto-registry.pb" -Ischemas schemas/usstats/usstats.proto

cd schemas
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative **/*.proto
```

## Test
go to gserver folder
```
go test
```
## Run
go to gserver folder
```
run: go run .
```
go to gclient folder
```
run: go run .
```

client will call server with its proto definition will get response in expected proto

