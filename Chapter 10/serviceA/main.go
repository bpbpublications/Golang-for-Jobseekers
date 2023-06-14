package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	gormMySQL "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"embed"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/spf13/cobra"
)

//go:embed migrations/*
var fs embed.FS

type Revenue struct {
	Datetime string `gorm:"primaryKey,autoIncrement"`
	Revenue  float64
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "This is a sample golang migrate application",
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
	Run: func(cmd *cobra.Command, args []string) {
		d, err := iofs.New(fs, "migrations")
		if err != nil {
			log.Fatal(err)
		}

		dbHost := os.Getenv("DB_HOST")

		m, err := migrate.NewWithSourceInstance(
			"iofs", d, fmt.Sprintf("mysql://user:password@(%v:3306)/application", dbHost))

		if err != nil {
			panic(fmt.Sprintf("unable to connect to database :: %v", err))
		}
		err = m.Up()
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("no change - no update done on database schema")
			os.Exit(0)
		}
		if err != nil {
			panic(fmt.Sprintf("unexpected error :: %v", err))
		}
	},
}

type RevenueGet struct {
	DB *gorm.DB
}

func (h RevenueGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rawDateTime := vars["datetime"]

	var u Revenue
	result := h.DB.First(&u, rawDateTime)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("bad connection"))
		return
	}

	rawResp, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(rawResp)
}

type RevenueCreate struct {
	DB *gorm.DB
}

func (h RevenueCreate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		return
	}

	type revenueCreate struct {
		Revenue float64 `json:"revenue"`
	}

	var uc revenueCreate
	json.Unmarshal(raw, &uc)

	u1 := Revenue{Datetime: time.Now().Format(time.RFC3339), Revenue: uc.Revenue}
	log.Printf("Revenue to be created on datetime: %v", u1.Datetime)
	result := h.DB.Create(&u1)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("bad connection :: %v", result.Error)))
		return
	}

	rawResp, _ := json.Marshal(u1)
	w.WriteHeader(http.StatusOK)
	w.Write(rawResp)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server start")
		dbHost := os.Getenv("DB_HOST")
		dsn := fmt.Sprintf("user:password@tcp(%v:3306)/application", dbHost)
		db, err := gorm.Open(gormMySQL.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("unable to connect to database :: %v", err))
		}

		r := mux.NewRouter()
		r.Handle("/revenue", RevenueCreate{DB: db}).Methods("POST")
		r.Handle("/revenue/{datetime}", RevenueGet{DB: db}).Methods("GET")

		srv := &http.Server{
			Handler: r,
			Addr:    ":8080",
		}

		log.Fatal(srv.ListenAndServe())
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(serverCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
