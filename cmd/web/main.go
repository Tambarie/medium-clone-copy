package main

import (
	"flag"
	"github.com/Tambarie/medium-clone/pkg/Database"
	"github.com/Tambarie/medium-clone/pkg/models"
	"net/http"
	"time"
)

type application struct {
	user      models.IUserCRUD
	post 	models.IPost
	comment models.IComments
}


func main()  {
	addr := flag.String("addr",":8081","pass the network to the address")
	flag.Parse()

	db := Database.ConnectDB()
	// initializing the app
	app := &application{
		user: &models.DbModel{
			Db: db,
		},
		post: &models.PostModel{
			Db: db,
		},

		comment: &models.CommentModel{
			Db: db,
		},
	}

	server := &http.Server{
		Handler: app.routes() ,
		Addr: *addr,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
		server.ListenAndServe()
}
