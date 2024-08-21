package main

//go:generate sqlc generate
//go:generate swag init --generalInfo ./handler/docs.go --parseDependency --output generated/docs
