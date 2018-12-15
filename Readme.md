#### Folders structure:

* `app/` - implementing business logic of app - 
Domain and Service

* `interface/` - implementing external interfaces - 
gRPC handlers, factory protobuf messages, map database operations into repository

* `infrastructure/` - interfaces which common for all apps -
 DB drivers and other tools

* `config/` - .toml configuration files

#### Inside `app/`:

> Domain - object of data

> Repository - methods for retrieving domain objects.

> Service - fetch Domain objects using Repositories and make operations

#### Inside `interface/`:

* `api/` - output folder for auto-generated gRPC code
* `factory/` - factories for convert:
    * TableRow <-> Domain
    * Message -> Domain
    * Message <- Domain
* `grpc-protofiles/` - git-submodule stores all protofiles across apps;
    It configured in `.gitmodules` files under root folder;
* `handler/` - implements gRPC handlers
* `repository/` - implements `app/` Repositories.
Using database commands from `infrastructure/`
* `tests/` - npm-package for testing gRPC service handlers

#### Relations:

1) `interface/` implements gRPC handlers which executes `app/` Services

2) `interface/` implements `app/` Repositories.
   Using database commands over tables from `infrastructure/`.

3) `app/` Services only who work with Domains using Repositories
