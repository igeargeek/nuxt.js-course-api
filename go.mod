module github.com/icecreamhotz/movie-ticket

go 1.12

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/go-playground/locales v0.12.1
	github.com/go-playground/universal-translator v0.16.0
	github.com/go-playground/validator v9.31.0+incompatible // indirect
	github.com/google/wire v0.4.0
	github.com/icecreamhotz/movie-ticket/configs v0.0.0
	github.com/icecreamhotz/movie-ticket/controllers v0.0.0
	github.com/icecreamhotz/movie-ticket/database v0.0.0
	github.com/icecreamhotz/movie-ticket/models v0.0.0
	github.com/icecreamhotz/movie-ticket/routes v0.0.0
	github.com/icecreamhotz/movie-ticket/utils v0.0.0
	go.mongodb.org/mongo-driver v1.3.1
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073 // indirect
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a // indirect
	golang.org/x/sys v0.0.0-20200302150141-5c8b2ff67527 // indirect
	golang.org/x/tools v0.0.0-20200304143113-d6a4d55695f2 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1
)

replace (
	github.com/icecreamhotz/movie-ticket/configs v0.0.0 => ./src/configs
	github.com/icecreamhotz/movie-ticket/controllers v0.0.0 => ./src/controllers
	github.com/icecreamhotz/movie-ticket/database v0.0.0 => ./src/database
	github.com/icecreamhotz/movie-ticket/models v0.0.0 => ./src/models
	github.com/icecreamhotz/movie-ticket/routes v0.0.0 => ./src/routes
	github.com/icecreamhotz/movie-ticket/utils v0.0.0 => ./src/utils
)
