# Accelone - Contacts API 

Basic example of CRUD API in golang

Requerimetns

- Go > 1.20


Folder structure
``` bash
├── README.md
├── controller
│   ├── contact.go
│   ├── controller.go
│   └── router.go
├── go.mod
├── go.sum
├── main.go
├── model
│   └── contact.go
├── postman-collection.json
├── repository
│   ├── contact.go
│   └── repository.go
└── usecase
    ├── contact.go
    └── usecase.go
```

To execute:

``` bash
go run main.go
```

See postman-collection.json for available endpoints