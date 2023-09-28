package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	
	models "subscriptionApi/modelspck"
	sqlserver "subscriptionApi/sqlserverpck"
	utils "subscriptionApi/utilspck"

	"github.com/gorilla/mux"
)

func getSubscribersFromDb() []models.Subscriber {
	var subscribers []models.Subscriber

	result, err := sqlserver.Db.Query("select id,name,is_free,add_date from subscribers")

	fmt.Println(result)

	if err != nil {
		log.Println(err)
	}
	defer result.Close()

	for result.Next() {
		var subs models.Subscriber
		err = result.Scan(&subs.Id, &subs.Name, &subs.IsFree, &subs.AddDate)
		if err != nil {
			log.Println(err)
		}
		subscribers = append(subscribers, subs)
	}

	return subscribers
}

func allSubscibers(w http.ResponseWriter, r *http.Request) {
	utils.JsonReponse(w, getSubscribersFromDb())
}

func getSubsciberById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var sub models.Subscriber

	result, err := sqlserver.Db.Query("select id, name, is_free, add_date from subscribers where id = ?", id)

	for result.Next() {
		err = result.Scan(&sub.Id, &sub.Name, &sub.IsFree, &sub.AddDate)
	}

	if err != nil {
		log.Println(err)
		utils.JsonReponse(w, models.BaseResult{
			Result:  false,
			Message: err.Error(),
		})
	} else {
		if result.Next() {
			utils.JsonReponse(w, sub)
		} else {
			utils.JsonReponse(w, nil)
		}

	}

	defer result.Close()
}

func createSubscriber(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newSubs models.Subscriber

	utils.JsonDeserialize(reqBody, &newSubs)

	result, err := sqlserver.Db.Prepare("insert into subscribers (name,is_free,add_date) values (?,?,?)")

	_, err = result.Exec(newSubs.Name, newSubs.IsFree, newSubs.AddDate)

	if err != nil {
		log.Println(err)

		utils.JsonReponse(w, models.BaseResult{
			Result:  false,
			Message: err.Error(),
		})
	} else {
		utils.JsonReponse(w, models.BaseResult{
			Result:  true,
			Message: "subscriber has been created",
		})
	}

	defer result.Close()
}

func deleteSubscriber(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	result, err := sqlserver.Db.Prepare("delete from subscribers where id=?")

	_, err = result.Exec(id)

	if err != nil {
		log.Println(err)

		utils.JsonReponse(w, models.BaseResult{
			Result:  false,
			Message: err.Error(),
		})
	} else {
		utils.JsonReponse(w, models.BaseResult{
			Result:  true,
			Message: "subscriber has been deleted",
		})
	}

	defer result.Close()
}

func updateSubscriber(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")
	subs, err := ioutil.ReadAll(r.Body)
	var updateSub models.Subscriber
	utils.JsonDeserialize(subs, &updateSub)

	result, err := sqlserver.Db.Prepare("update subscribers set name=? ,is_free=?,add_date=?  where id=?")
	_, err = result.Exec(updateSub.Name, updateSub.IsFree, updateSub.AddDate, id)

	if err != nil {
		log.Println(err)

		utils.JsonReponse(w, models.BaseResult{
			Result:  false,
			Message: err.Error(),
		})
	} else {
		utils.JsonReponse(w, models.BaseResult{
			Result:  true,
			Message: "subscriber has been updated",
		})
	}

	defer result.Close()
}

func handleRequests() {
	myrouter := mux.NewRouter().StrictSlash(false)
	myrouter.HandleFunc("/subscribers", allSubscibers).Methods("GET")
	myrouter.HandleFunc("/subscriber", getSubsciberById).Methods("GET")
	myrouter.HandleFunc("/subscriber", deleteSubscriber).Methods("DELETE")
	myrouter.HandleFunc("/subscriber", createSubscriber).Methods("POST")
	myrouter.HandleFunc("/subscriber", updateSubscriber).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", myrouter))
}

func main() {
	
	sqlserver.Init()
	handleRequests()
	sqlserver.Db.Close()
}
