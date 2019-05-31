package commandparser

import (
	"log"
	"net/http"
)

type slackRequest struct {
	// Token string `json:"token"`
	// TeamID      string `json:"team_id"`
	// TeamDomain  string `json:"team_domain"`
	// ChannelID   string `json:"channel_id"`
	// ChannelName string `json:"channel_name"`
	// UserID      string `json:"user_id"`
	// UserName    string `json:"user_name"`
	// Command     string `json:"command"`
	Text string `json:"text"`
	// ResponseURL string `json:"response_url"`
	// TriggerID   string `json:"trigger_id"`
}

func ToilEliminatorCommandParser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	text := r.PostFormValue("user_name")
	log.Println(text)

	// Read the request body.
	// byteArray, err := ioutil.ReadAll(r.Body)
	// defer r.Body.Close()
	// if err != nil {
	// 	log.Printf("ioutil.ReadAll: %v", err)
	// 	http.Error(w, "Error reading request", http.StatusBadRequest)
	// 	return
	// }
	// log.Println(string(byteArray))
	// jsonBytes := ([]byte)(byteArray)

	// //parse json
	// var jsonBody map[string]interface{}
	// err = json.Unmarshal(data, &jsonBody)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// log.Printf("%v\n", jsonBody)

	// Parse the request body to get the topic name.
	// slackReq := slackRequest{}

	// slackReq := new(slackRequest)
	// if err := json.Unmarshal(byteArray, &slackReq); err != nil {
	// 	log.Printf("json.Unmarshal: %v", err)
	// 	http.Error(w, "Error parsing request", http.StatusBadRequest)
	// 	return
	// }
	// log.Println(slackReq)

	// //Read body data to parse json
	// body := make([]byte, length)
	// length, err = req.Body.Read(body)
	// if err != nil && err != io.EOF {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// create response message
	// msg := slackReq.Text
	msg := "accept"
	w.Write([]byte(msg))
}
