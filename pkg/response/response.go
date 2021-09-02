package response

import (
	"encoding/json"
	"net/http"
)

//Message func
func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{"code": status, "msg": message}
}

//JSON func
func JSON(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}
