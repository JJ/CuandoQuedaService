package main

// taken from https://github.com/tucnak/telebot
import (
	"os"
	"fmt"
	"strconv"
	"log"
	"encoding/json"	
	"net/http"
	"github.com/JJ/HitosIV"
)

func main() {
	http.HandleFunc("/", dame_hitos)
	port := os.Getenv("PORT")
	if port == "" {
		port = "31415"
	}
	bind := fmt.Sprintf(":%s", port)
	fmt.Printf("Escucha en %s...", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func dame_hitos(res http.ResponseWriter, req *http.Request) {
	este_hito, _ := strconv.ParseInt(req.URL.Query().Get("hito"), 10, 0)
	js,err := json.Marshal( HitosIV.Uno( uint(este_hito) ) )
	if ( err != nil ) {
		fmt.Fprintf( res, "Error %s", err )
	} else {
		res.Header().Set("Content-Type", "application/json")
		res.Write(js)
	}
}
