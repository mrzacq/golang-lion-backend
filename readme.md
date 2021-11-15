### requirements

1. Postgres 14.1
2. go version go1.17.3

### running project

1. copy sql command in folder `db`
2. change db config in folder `app/database.go`
3. run command `go run main.go`
4. test using file `test.http`

### additional command

- init project
  `go mod init <projectname>`

- get module
  `go get <project url>`

### task

[x] register (fullname, username, password, gender)
[] login
[] hash user password
[x] get detail music / find music by id
[x] get all music
[] authentication on get music endpoint
[] favorit
