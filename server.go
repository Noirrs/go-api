package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	serverer()

}

func serverer() {
	var port = 3000
	var kazdata = "null"
	fmt.Println("Server is starting at http://localhost:" + strconv.Itoa(port))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		value := strings.Split(r.URL.Path, "/")
		if len(value[2]) < 1 {
			for _, v := range users {
				fmt.Println(len(users))
				if len(v.permissions) > 0 {
					if contains(v.permissions, "owner") {
						kazdata = "owner"
					} else if contains(v.permissions, "staff") {
						kazdata = "staff"
					} else {
						kazdata = "member"
					}
				}

				fmt.Fprintln(w, "ID: "+strconv.Itoa(v.id)+"  | Username: "+v.username+"  | Highest Permission "+kazdata)
			}
		} else if len(value[2]) > 0 {
			for _, user := range users {
				var kok, err = strconv.Atoi(value[2])
				if err != nil {
					fmt.Fprint(w, "Parameter doesn't have a type of number")
					return
				}
				if kok == user.id {
					if len(user.permissions) > 0 {
						if contains(user.permissions, "owner") {
							kazdata = "owner"
						} else if contains(user.permissions, "staff") {
							kazdata = "staff"
						} else if contains(user.permissions, "member") {
							kazdata = "member"
						}
					}
					fmt.Fprintln(w, "ID: "+strconv.Itoa(user.id)+"  | Username: "+user.username+"  | Highest Permission "+kazdata)
					return
					} 
			}
			  
				fmt.Fprintf(w, "Unknown User")
				return
			
		}
	})

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

var users = []User{
	{username: "Zürafa", id: 0, permissions: []string{"member", "staff", "owner"}},
	{username: "Tavuk", id: 1, permissions: []string{"member"}},
	{username: "Kaz", id: 2, permissions: []string{"staff"}},
}
