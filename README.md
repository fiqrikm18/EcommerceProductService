# Ecommerce Product Service
This is ad simple product service api using echo and postgres as database.

## Table Of Content
- [Pre-requisite](#pre-requisite)
- [How To Install](#how-to-install)
- [Scripts](#scripts)
- [Project Structure](#project-structure)

## Pre-requisite
- [Go 1.24](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [go-migrate](https://github.com/golang-migrate/migrate)
- [swaggo](https://github.com/swaggo/swag)

## How to install
To run the application make sure pre-requisite above are installed on your machine, here are step to run the application.
1. Clone application from this repository using blow command
```
git clone git@github.com:fiqrikm18/EcommerceProductService.git
```
2. Copy or rename `.env.example` file to `.env`
3. Change confgirations inside file `.env`, adjust to the settings on your local machine, for example
```
APPLICATION_PORT=8080
APPLICATION_HOST=localhost
APPLICATION_ROOT_CMD=cmd/app/main.go

POSTGRES_DSN="host=127.0.0.1 user=postgres password=postgres dbname=ecommerce port=5432 sslmode=disable TimeZone=Asia/Jakarta"
POSTGRES_MIGRATION_DSN="postgres://postgres:postgres@127.0.0.1:5432/ecommerce?sslmode=disable"
```

if you want running othe application using docker you can replace host on database section to container name where database used, in this configuration you can use

```
APPLICATION_PORT=8080
APPLICATION_HOST=localhost
APPLICATION_ROOT_CMD=cmd/app/main.go

POSTGRES_DSN="host=postgres-db user=postgres password=postgres dbname=ecommerce port=5432 sslmode=disable TimeZone=Asia/Jakarta"
POSTGRES_MIGRATION_DSN="postgres://postgres:postgres@postgres-db:5432/ecommerce?sslmode=disable"
```

`postgres-db` are service name on `docker-compose.yml` file

4. Run `make verify` to install dependency application used for
5. Run `make run-migration` to run migration to the database
6. Run `mnake generate-docs` to generate swagger documentation
7. If you want to run application ond development mode you can run `make dev` then open `http://localhost:8080/docs/index.html` to show documentation of API
7. If you want to run application ond production mode you can run `make docker-run` then open `http://localhost:8080/docs/index.html` to show documentation of API, if you want using docker, if you want to running it manually
```
go build -o product-service ./cmd/app
./product-service
```
then open `http://localhost:8080/docs/index.html` to show documentation of API

## Scripts
This application contains some script to make easier for development
1. Install dependencies
```
make deps
```

2. Run golang lint
```
make golangci-lint
```

3. Run gosec
```
make gosec
```

4. Run govulncheck
```
make govulncheck
```

5. Install dependencies the run golangcilint, gosec, and govulncheck
```
make verify
```

6. Run on development mode
```
make dev
```

7. Generate swagger documentation
```
make generate-docs
```

8. Generate mock interface
```
make generate-mock
```

9. Create Migration
```
make create-migration seq=initial_migration
```

10. Run Migration
```
make run-migration
```

11. Rollback Migration
```
make rollback-migration
```

12. Run using docker
```
make docker-run
```

13. Shutdown docker
```
make docker-down
```

## Project Structure
```
├── cmd/
│   ├── app/
│   │   └── main.go
│   └── root.go
├── common/
│   └── router.go
├── config/
│   ├── application.go
│   └── database.go
├── constants/
│   └── database.go
├── db/
│   ├── migrations/
│       ├── 000001_intial_migration.down.sql
│       └── 000001_intial_migration.up.sql
├── internal/
│   ├── domain/
│       ├── brand/
│       │   ├── dto/
│       │   │   └── brand_dto.go
│       │   ├── entity/
│       │   │   └── Brand.go
│       │   ├── presenter/
│       │   │   └── brand_presenter.go
│       │   ├── repository/
│       │   │   └── brand_repository.go
│       │   ├── usecase/
│       │   │   └── brand_usecase.go
│       │   └── dependency.go
│       ├── product/
│           ├── dto/
│           │   └── product_dto.go
│           ├── entity/
│           │   └── Product.go
│           ├── presenter/
│           │   └── product_presenter.go
│           ├── repository/
│           │   └── product_repository.go
│           ├── usecase/
│           │   └── product_usecase.go
│           └── dependency.go
├── pkg/
│   ├── response/
│   │   └── http.go
│   ├── validator/
│       └── request.go
├── scripts/
│   └── install-go-migrate.sh
├── Dockerfile
├── Makefile
├── README.md
├── docker-compose.yml
├── go.mod
└── go.sum
```
