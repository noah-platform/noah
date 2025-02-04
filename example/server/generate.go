package main

//go:generate sqlc generate
//go:generate swag init -v3.1 --generalInfo ./handler/docs.go --parseDependency --output generated/docs
