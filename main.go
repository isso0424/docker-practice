package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func parseHeader(key string, values []string) string {
	switch len(values) {
	case 0:
		log.Printf("The key %s is empty\n", key)
		return ""
	case 1:
		return values[0]
	default:
		text := ""
		for index, value := range values {
			text += value
			if index + 1 != len(values) {
				text += ", "
			}
		}

		return text
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)

			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		header := w.Header()

		for key, value := range r.Header.Clone() {
			header.Add(key, parseHeader(key, value))
		}

		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
