package controllers

import (
	"../utils"
	"../views"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ImageService struct {
	Db *sql.DB
	HomeView views.View
	ShowTableView *views.View
	ContactView views.View
}

type tableData struct {
	ID int // primary key
	ImageId int // the image id
	IsTagged int // sqlite boolean is an int (0 = false, 1 = true)
}

// tagImage reads the route variables for the current request
// identifies which image was tagged and sets the flag isTagged to true in the DB
func (s *ImageService) TagImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	utils.Must(err)

	// when an imageId is tagged, set isTagged to True in the DB
	statement, _ := s.Db.Prepare("INSERT INTO taggedImages (imageId, isTagged) VALUES (?, ?)")

	// Note: SQLite boolean is an int (0 = false, 1 = true)
	_, err = statement.Exec(id, 1)
	utils.Must(err)

	http.Redirect(w, r, "/", http.StatusFound)
	return
}

// When an image is tagged using the "Tag" button, a new entry is created in the DB
// showTable queries the DB for all DB entries of any tagged images
// and then renders a web page where these entries are being displayed to the user
func (s *ImageService) ShowTable(w http.ResponseWriter, r *http.Request) {
	// TODO: This is only a temporary feature which serves for development and will be removed in production
	//  This feature will not be available for the end-user and it was made only so that a member of Segmed can see it.

	// query the DB and create a JSON object from all the entries
	var id, imageId, isTagged int
	var data tableData
	var dataArr []tableData

	rows, _ := s.Db.Query("SELECT id, imageId, isTagged FROM taggedImages")
	for rows.Next() {
		rows.Scan(&id, &imageId, &isTagged)
		data = tableData {id, imageId, isTagged}
		dataArr = append(dataArr, data)
	}

	var jsonData []byte
	jsonData, err := json.Marshal(dataArr)
	utils.Must(err)

	w.Header().Set("Content-Type", "text/html")
	// render show table page with returned query as JSON
	utils.Must(s.ShowTableView.Render(w, string(jsonData)))
}