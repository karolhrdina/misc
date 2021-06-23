## HW (homework)

### How to ...

#### ... build the services 
* `make bin/port-domain` to build the `port-domain` service
* `make bin/client-api` to build the `client-api` service

This will produce binary in `bin/` directory

Running  `make clean/<service>` will delete the binary.

#### ... re-generate protocol buffers
`./scripts/gen-proto.sh` regenerates any changes of `proto/port-domain.proto` into `pb.go`

Please note, that this is poor man's protobuf generation. Although good enough
for a PoC like this, a better way would be to use something like
`https://github.com/bufbuild/buf` where using config yaml's you can
descriptively control what and how is generated without the need for the
nitty-gritty details of `protoc` and `protoc-gen-go-grpc` :) Not even
mentioning that the tools and their options keep changing with every major
version (not in backwards-compatible way) which makes it hard to write by hand.

#### ... re-generate sqlboiler 
First run `./scripts/run-postgres.sh`
then `sqlboiler psql` from `services/port-domain` directory, where the `sqlboiler.toml` is.

The `./scripts/run-postgres.sh` is also a poor man's way of doing the db
migration (even for development), but it's quick and ok for PoC. Better way of
doing this would be to ingergate migration tool (like sql-migrate) with the
make or compose tools or write an internal binary that would initialize an
migrate an empty postgres database.

#### ... build docker images
* `docker build -f docker/Dockerfile.port-domain --tag port-domain:latest .` to build the `port-domain` image
* `docker build -f docker/Dockerfile.client-api --tag client-api:latest .` to build the `client-api` image

#### ... re-generate Wire
Run `wire` in `services/port-domain` or `service/client-api`.

#### ... install dependencies
* google.golang.org/protobuf/cmd/protoc-gen-go
* google.golang.org/grpc/cmd/protoc-gen-go-grpc
* github.com/volatiletech/sqlboiler/v4
* github.com/google/wire
* github.com/golangci/golangci-lint

I am assuming that you have all the other tools (like docker, docker-compose, go compiler...) installed.

If I had time I would at least make a target in Makefile and make use of `go install ` a certain version of each tool.

### Running the app (i.e. demo)

`docker-compose --project-name=hw -f docker-compose-main.yml up`

Network is set to be reachable from outside. With the help of scripts from `testdata`, we can
* `./api_post.sh` which will make a POST request to /v1/ports/import 
* then we can list them back: `./api_list.sh`
* and even check the database: `./hw_psql.sh` (will ask for password, from `.env` file, it is: "password")
* using a bloomRPC or similar tool, we can even test gRPC of port-domain

### Tests 
Unfortunately, I ran out of time and wasn't able to write any (unit)tests.
Let me at least write here, what I would do if I had time.

* unit tests:
  `port-domain`:
    * the PortStorer interface can be mocked (using for e.g. stretchr testify/mock)
  and all cases, even errors, can be simulated using the On() methods and furthermore asserted
  at the end.
    * conversion functions.

  `client-api`:
    * gRPC client dependency is an interface too. Although I would probably use testify/mock for this
    (it's a small interface) but there  are mock generators fo gRPC clients as well.
    * unit tests for handlers, for the conversion functions

* integration tests:
  To test the whole thing, I'd make a docker image + compose to run the whole setup,
  and on the extra image I'd run the integration tests that would call the main interface (http rest api, i.e.
  the import as well as list) and could even check what has been stored in the database.

* gray-area:
  Sometimes it may be a good idea to run a test suite, but against a real dependency, like a DB.
  Here it would be the storage implementation. I would simply use stretchr testify Suite
  to create unit test suites, where before the the whole suite a db migration would be run 
  and before each test the db table would be truncated so that each test would have empty database.

I am more than happy to talk about all the actual test (corner)cases on the technical interview, I simply
ran out of time to write them out individually (similar time to actualy write the suites :shrug:)

### directory structure
* `services`:
  Client-api and port-domain applications (services).

* `docker`:
  Docker files for both services.

* `pkg`:
  Go packages. Basically code that is shared/used by more than one service.

* `proto`:
  gRPC protocol buffer definitions. They are generated into `pb.go` directory.

* `pb.go`: 
  Generated protoco buffer gRPC files and stubs.

* `scripts`:
  Helper scripts (for building services, generating protobufs, running postgres...) some of which are
  used by Makefile.

* `bin`:
  Built service binaries.
  
* `testdata`:
  Tools/scripts that I used to test the application.


