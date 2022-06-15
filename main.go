package main

import (
  "fmt"
  "log"
  "strconv"
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
)

type contact struct {
  ID string `json:"id"`
  FirstName string `json:"firstname"`
  LastName string `json:"lastname"`
  Phone string `json:"phone"`
  Email string `json:"email"`
}

var port string = ":3000"
var contacts []contact

// Controllers
func getContacts(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(contacts)
}

func getContact(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)

  for _, item := range contacts {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      break
    }
  }
}

func createContact(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var contact contact

  _ = json.NewDecoder(r.Body).Decode(&contact)

  contact.ID = strconv.Itoa(len(contacts) + 1)
  contacts = append(contacts, contact)

  json.NewEncoder(w).Encode(contacts)
}

func updateContact(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  var contact contact

  _ = json.NewDecoder(r.Body).Decode(&contact)

  for index, item := range contacts {
    if item.ID == params["id"] {
      contacts = append(contacts[:index], contacts[index + 1:]...)
      contact.ID = params["id"]
      contacts = append(contacts, contact)
      break
    }
  }
  json.NewEncoder(w).Encode(contact)
}

func deleteContact(w http.ResponseWriter, r *http.Request) {  
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)

  for index, item := range contacts {
    if item.ID == params["id"] {
      contacts = append(contacts[:index], contacts[index + 1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(contacts)
}

func main() {
  // Router
  r := mux.NewRouter()

  contacts = append(contacts, contact{ID: "1", FirstName: "John", LastName: "Doe", Phone: "12345678", Email: "luis@email.com"})
  contacts = append(contacts, contact{ID: "2", FirstName: "Rayn", LastName: "Ray", Phone: "234567532"})

  // Routes
  r.HandleFunc("/contacts", getContacts).Methods("GET")
  r.HandleFunc("/contacts/{id}", getContact).Methods("GET")
  r.HandleFunc("/contacts", createContact).Methods("POST")
  r.HandleFunc("/contacts/{id}", updateContact).Methods("PUT")
  r.HandleFunc("/contacts/{id}", deleteContact).Methods("DELETE")

  fmt.Printf("Server running on http://localhost%s", port)

  // Starting server
  log.Fatal(http.ListenAndServe(port, r))
}
