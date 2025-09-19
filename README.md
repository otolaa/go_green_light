# Application Greenlight

```
go mod init
go mod tidy
go run .
```

## Point API and REST 👾

| Method | URL Pattern | Action |
| :---:   | :---: | :---: |
| GET | /v1/healthcheck   | Show application health and version information |
| GET | /v1/movies  | Show the details of all movies |
| POST | /v1/movies   | Create a new movie |
| GET | /v1/movies/:id   | Show the details of a specific movie |
| PATCH | /v1/movies/:id   | Update the details of a specific movie |
| DELETE | /v1/movies/:id  | Delete a specific movie |
| POST | /v1/users  | Register a new user |
| PUT | /v1/users/activated  | Activate a specific user |
| PUT | /v1/users/password  | Update the password for a specific user |
| POST | /v1/tokens/authentication  | Generate a new authentication token |
| POST | /v1/tokens/password-reset  | Generate a new password-reset token |
| GET | /debug/vars  | Display application metrics |

## Generating the skeleton directory structure

```
.
|---bin
|---cmd
|   +---api
|       +---main.go
|---internal
|---migrations
|---remote
|---go.mod
|---Makefile
```

## Info for folder catalog ⚔️

```
bin — скомпилированные двоичные файлы, готовые к развёртыванию на рабочем сервере;

cmd/api —  основная функция приложения;

internal — различные вспомогательные пакеты;

migrations — файлы миграции для базы данных;

remote — файлы конфигурации и сценарии настроек для производственного сервера;

go.mod — информация о зависимостях проекта, версиях и путях к модулям;

Makefile — инструкции по автоматизации частых административных задач — проверка кода, создание двоичных файлов и выполнения миграций.
```

## Run 🚀

```
killall -9 go
$ go run ./cmd/api
```



## Use curl in terminal ⚙️

```
$ curl -i localhost:4000/v1/healthcheck
status: available
environment: development
version: 1.0.0

$ curl -i -X POST localhost:4000/v1/movies
create a new movies

$ curl -i localhost:4000/v1/movies/12
show the details of movies 12
```