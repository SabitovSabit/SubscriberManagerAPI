package utils

import(
	"encoding/json"
	"net/http"
)

func JsonReponse(w http.ResponseWriter,data interface{}){
	json.NewEncoder(w).Encode(data)
}

func JsonDeserialize(data []byte, obj interface{}) {
	json.Unmarshal(data, &obj)
}