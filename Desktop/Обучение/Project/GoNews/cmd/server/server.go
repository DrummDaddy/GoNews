package main

import (
	"GoNews/pkg/api"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/memdb"
	"GoNews/pkg/storage/mongo"
	"GoNews/pkg/storage/postgres"
	"log"
	"net/http"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	// Создаём объект сервера.
	var srv server

	// Создаём объекты баз данных.
	//
	// БД в памяти.
	db := memdb.New()

	// Реляционная БД PostgreSQL.
	db2, err := postgres.New("postgres://postgres:postgres@localhost/posts?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Документная БД MongoDB.
	db3, err := mongo.New("mongodb://localhost:27017", "go_news")
	if err != nil {
		log.Fatal(err)
	}

	// Закомментируйте или раскомментируйте нужный способ работы с базой данных:
	srv.db = db  // В памяти
	srv.db = db2 // PostgreSQL
	srv.db = db3 // MongoDB

	// Инициализируем хранилище сервера конкретной БД.
	srv.db = db2 // Для примера используем PostgreSQL.

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Запускаем веб-сервер на порту 8080 на всех интерфейсах.
	http.ListenAndServe(":8080", srv.api.Router())
}
