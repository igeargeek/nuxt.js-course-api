module github.com/icecreamhotz/movie-ticket

go 1.12

require (
	github.com/google/wire v0.4.0
	github.com/icecreamhotz/movie-ticket/configs v0.0.0
	github.com/icecreamhotz/movie-ticket/controllers v0.0.0
	github.com/icecreamhotz/movie-ticket/database v0.0.0
	go.mongodb.org/mongo-driver v1.3.1
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073 // indirect
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a // indirect
	golang.org/x/sys v0.0.0-20200302150141-5c8b2ff67527 // indirect
	golang.org/x/tools v0.0.0-20200304143113-d6a4d55695f2 // indirect
)

replace (
	github.com/icecreamhotz/movie-ticket/configs v0.0.0 => ./configs
	github.com/icecreamhotz/movie-ticket/controllers v0.0.0 => ./controllers
	github.com/icecreamhotz/movie-ticket/database v0.0.0 => ./database
)
