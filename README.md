# fer
fer is not Ferdian

# Kumparan Microservices Generator 

you can make microservices start from proto. see the proto example in `pb/` folder

## Usage
you need to create proto file with services RPC with path like this
`pb/'$service/$protoname.proto`

example (see the proto example in folder `pb/`)
`pb/content/content_service.proto`
```
service ContentService{
     // Topic  << you must define the domain
     rpc CreateTopic(CreateTopicRequest) returns (Topic) {}
     rpc DeleteTopicByID(DeleteByIDRequest) returns (Empty) {}
 
     // Story  << you must define the domain
     rpc CreateStory(CreateTopicRequest) returns (Topic){}
     rpc DeleteStoryByID(DeleteByIDRequest) returns (Empty){}
 }
```

and create microservices
`fer project init --name content-service`
 
 - you will be asked to insert proto source path
 `pb/content/content_service.proto` 
 
 - new service will generated like this
 ```
-content-service/
    -client/    ->(Generated From Proto)
    -config/
    -console/
    -db/
    -event/
    -pb/
    -repository/
    -service/     ->(Generated From Proto)
    -worker/
    -config.yml
    -config.yml.dev
    -config.yml.example
    -config.yml.prod
    -config.yml.staging
    -go.mod       ->(mod is already with service name)
    -go.sum
    -LICENSE
    -main.go
    -Makefile
    -README.mod
 ```

## INSTALL
Clone the repository to your desired destination folder e.g :
```
cd ~ && git clone git@github.com:kumparan/fer.git
```
run build
```
go build
```
run  install
```
go install
```

It should be available as command now in terminal

## Feature
-   [x] Scaffold New Microservices
-   [x] Generate Service&test files and Client files From Proto
-   [ ] Generate Repository (include model)
-   [ ] Add worker with command
-   [ ] Add Nats Subscriber with command
-   [ ] db migration file generator