package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ilmsg/ideal-octo-sniffle/config"
	"github.com/ilmsg/ideal-octo-sniffle/database"
	"github.com/ilmsg/ideal-octo-sniffle/ideal"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	dbconfig := config.GetDBConfig()
	db, err := database.GetDatabase(dbconfig)
	if err != nil {
		return err
	}

	queries := ideal.New(db)

	// newUser, err := queries.CreateUser(ctx, ideal.CreateUserParams{
	// 	Email:    "scott@gmail.com",
	// 	Password: "scott",
	// })
	// if err != nil {
	// 	return err
	// }

	users, err := queries.ListUsers(ctx)
	if err != nil {
		return err
	}

	for _, user := range users {
		fmt.Printf("User{id=%d,email=%s,password=%s}\n", user.ID, user.Email, user.Password)
	}

	// userId, err := newUser.LastInsertId()
	// if err != nil {
	// 	return err
	// }
	// newStore, err := queries.CreateStore(ctx, ideal.CreateStoreParams{
	// 	Title:       "Store 1",
	// 	Description: "Store Description 1",
	// 	Userid:      userId,
	// })
	// if err != nil {
	// 	return err
	// }

	// storeId, err := newStore.LastInsertId()
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(storeId)

	stores, err := queries.ListStores(ctx)
	if err != nil {
		return err
	}

	for _, store := range stores {
		fmt.Printf("Store{id=%d,title=%s,description=%s,userId=%d}\n", store.ID, store.Title, store.Description, store.Userid)
	}

	return nil
}
