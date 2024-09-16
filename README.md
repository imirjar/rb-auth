![RB_Auth](docs/img/rb_auth.jpg)

Authentification server. Powered by reDBeaver


# About app

RB_AUTH app consists of 3 layers:
1) Gateway - layer of http/grpc/mpq connection
2) Service - provides authentication/authorization/reagistration logic
3) Storage - keeps the information about current Users/Groups/Roles

## Storage
1) memory - is a fake storage for developing mode only
2) psql - postgresql database

You can make your custom connection, witch realize internal/service/service.go -> "storage" interface

## Service 
1) Authentificate - 
2) Authorizate
3) Registrate - register new User in system

## Gateway 
1) Sign in
2) Log in
3) Authenticate

### 1. Gaterway tests

#### 1.1 Gateway service mock
For good tests we are using mockgen util

```bash 
!!

mockgen -destination=internal/gateway/http/mock_service.go -package=http github.com/imirjar/rb-auth/internal/gateway/http Service
```

#### 1.2
#### 1.3