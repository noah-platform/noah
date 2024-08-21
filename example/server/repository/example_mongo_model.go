package repository

type ExampleDocument struct {
	ID    int    `bson:"exampleId"`
	Title string `bson:"title"`
}
