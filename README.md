# fer
fer is not Ferdian

## INSTALL
```
go get github.com/kumparan/fer
```

It should be available as command now in terminal.

Check fer version
```
fer version
```

## Feature
-   [x] Scaffold New Microservices
-   [x] Generate Service&test files and Client files From Proto
-   [x] DB migration file generator
-   [x] Generate Repository (include model)
-   [ ] Add worker with command
-   [ ] Add Nats Subscriber with command

## Kumparan Microservices Generator 

you can make microservices start from proto. see the proto example in `pb/` folder

### Usage
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
`fer generate project --name content-service`
 
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
    -README.md
 ```

## DB MigrationFile Generator
You can create db migration file

`fer generate migration create_story`

and new migration will be created like this

`db/migration/20191007130809_create_story.sql`

## Repository & Model Generator
You can create repository and model file
`fer generate repository promoted_link`

and new repository and model will be created like this 
```
repository/model/promoted_link.go created
repository/promoted_link_repository.go created
```

