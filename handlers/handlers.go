package handlers

import (
	"fmt"
	"jql-server/data"
	"log"
	"net/http"

	"github.com/sCuz12/celeritas"
	aliasparser "github.com/sCuz12/go-json-query-parser"
)

// Handlers is the type for handlers, and gives access to Celeritas and models
type Handlers struct {
	App    *celeritas.Celeritas
	Models data.Models
}

type QueryRequest struct {
	Query string `json:"query"`
}

func (h *Handlers) JsonSearch(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		h.App.ErrorLog.Println("Invalid request method",http.StatusMethodNotAllowed)
	} 


	query := r.FormValue("query")

	if query == "" {
		log.Printf("Query string is missing")
		http.Error(w, "Query string is missing", http.StatusBadRequest)
		return
	}
	
	jsonData := `[{"name": "John", "age": 30, "city": "New York"}, {"name": "Jane", "age": 25, "city": "Chicago"}]`

	err := r.ParseMultipartForm(10 << 20) //10MB

	if err != nil {
		h.App.ErrorLog.Printf("Failed to parse multi form",err)
		http.Error(w,"Invalid multiform data",http.StatusBadRequest)
	}

	// Retrieve the uploaded file
	file, fileHeader, err := r.FormFile("json-file")

	if err != nil {
		h.App.ErrorLog.Printf("Failed to retrieve file: %v", err)
		http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Log the filename
	filename := fileHeader.Filename
	log.Printf("Uploaded file: %s", filename)
	
	// h.App.WriteJSON(w,http.StatusAccepted,"Hello controller here")
	var jsonParser aliasparser.Query

	jsonParser.Parse(query)

	results,total,err := jsonParser.ProcessQuery(jsonData)

	if err != nil {
		h.App.ErrorLog.Println("Something went wrong")
		http.Error(w,"Something went wrong",http.StatusInternalServerError)
	}

	fmt.Println(total)

	h.App.WriteJSON(w,http.StatusAccepted,results)

}
