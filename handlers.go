package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"html/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var b Bookings

	res, err := http.Get("http://127.0.0.1:8080/bookings")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code errorL %d %s",
			res.StatusCode,
			res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)

	json.Unmarshal([]byte(body), &b)

	t, err := template.ParseFiles("tmpl/index.html")
	if err != nil {
		log.Print("template parsing errorL ", err)
	}
	err = t.Execute(w, b)
	if err != nil {
		log.Print("template executing error: ", err)
	}
	//fmt.Println(b[0].Court)
}
