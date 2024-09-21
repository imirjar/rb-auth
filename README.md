![RB_Auth](docs/img/rb_auth.jpg)


# About app

.∧＿∧\
( ･ω･｡)つ━☆・*。\
⊂　 ノ 　　　・゜+.\
しーＪ　　　°。+ *´¨)\
　　　　　　　　　.· ´Authentification server.
 Powered by reDBeaver¸.·*´¨) 



### Running app
Install golang from page https://go.dev/doc/install

```bash 
make build_and_run
```

# App layers:
RB_AUTH app consists of 3 layers:
1) Gateway - layer of http/grpc/mpq connection
2) Service - provides authentication/authorization/reagistration logic
3) Storage - keeps the information about current Users/Groups/Roles

## Storage
1) memory - is a fake storage for developing mode only
2) psql - postgresql database

You can make your custom connection, which realize internal/service/service.go -> "storage" interface

## Service 
1) Authentificate - 
2) Authorizate
3) Registrate - register new User in system

## Gateway 
1) Sign in
2) Log in
3) Authenticate

# Test layers

## Mocks

#### 1 Gateway service mock
Implement service interface

```bash 
mockgen -destination=internal/gateway/http/mock_service.go -package=http github.com/imirjar/rb-auth/internal/gateway/http Service
```

#### 2 Service storage mock
Implement storage interface

```bash 
mockgen -destination=internal/gateway/http/mock_service.go -package=http github.com/imirjar/rb-auth/internal/gateway/http Service
```

#### 3 Storage db mock
Implement database functional

```bash 
mockgen -destination=internal/gateway/http/mock_service.go -package=http github.com/imirjar/rb-auth/internal/gateway/http Service
```

##  Run tests
```
make test
```

# Documentation

## Swagger

### Install swagger
```bash
go install github.com/swaggo/swag/cmd/swag@latest 
```

### Create swagger API doc

```bash
swag init -g internal/gateway/http/http.go
```