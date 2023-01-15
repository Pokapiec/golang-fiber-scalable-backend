# Golang app made to get more in touch with Go language and it's ecosystem

**It is supposed to teach me:**
- how to structure project for fairly large project
- how to deal with databases in golang language (especially gorm ORM and mongodb)
- how to setup developement environment for golang (Makefile and air for hot reloading)

And i hope the list will be expanding more and more !
:crossed_fingers:

# Using the project

For project to be working correctly 2 tools should be installed:

(Air for live server reloading on files change)
```
go install github.com/cosmtrek/air@latest 
```

(For generating swagger docs)
```
go install github.com/swaggo/swag/cmd/swag@latest 
```

In the project Makefile is used to simplify handling repetitive commands in developement process:
- make dev -> start server in reload mode
- make swagger -> format swagger comments and generate swagger docs
- make build -> generate production ready executable