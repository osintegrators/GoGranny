package main

import(
	"fmt"
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"net/http"
	"encoding/json"
)


const requestPath = "/request/"
const requestPathLen = len(requestPath)

const Path = "/"
const PathLen = len(Path)

func main(){
    http.HandleFunc(requestPath, requestHandler)
    http.HandleFunc(Path, indexHandler)
    http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[PathLen:]
    filename := "index.html"
    if title == ""{
    	filename = "index.html"
    }else{
	    filename = title
	}
    http.ServeFile(w, r, filename)
}


func requestHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[requestPathLen:]
    switch(title){
    case "retrieveList":
    	out:= retrieveList()
    	fmt.Fprintf(w, out)
    	
    case "retrieveContact":
    	id := r.FormValue("id")
    	out:= retrieveContact(id)
    	fmt.Fprintf(w, out)
    
    case "saveContact":
    	id := r.FormValue("id")
    	name := r.FormValue("name")
    	address := r.FormValue("address")
    	phone := r.FormValue("phone")
    	email := r.FormValue("email")
    	out:= saveContact(id, name, address, phone, email)
    	fmt.Fprintf(w, out)
    
    case "deleteContact":
    	id := r.FormValue("id")
    	deleteContact(id)
    	
    }
}


type Contact struct {
    Id      int    `json:"pkId"`
    Name string	`json:"fldName"`
    Address   string	`json:"fldAddress"`
    Phone	string	`json:"fldPhone"`
    Email	string	`json:"fldEmail"`
}

func db_connect() *sql.DB{
	db, e := sql.Open("mysql", "root@/test?charset=utf8")
	if e != nil{
		panic(e)
	}
	return db
}

func retrieveList() string{
	sqlText := "SELECT pkId, fldName FROM tblContact ORDER BY fldName;"
	db := db_connect()
	defer db.Close()
	rows, _ := db.Query(sqlText)

	// Fetch rows
	var contact *Contact
	var ts string
	ret := "<option value=\"-1\" selected ></option>\n"
	for rows.Next() {
		contact = new(Contact)
		rows.Scan(&contact.Id, &contact.Name)
		ts = fmt.Sprintf("<option value='%d' >%s</option>\n",contact.Id,contact.Name)
		ret += ts
	}
	return ret
}

func retrieveContact(id string) string{
	if id == "-1"{
		return "-1"
	}
	sqlText := fmt.Sprintf("SELECT fldName, fldAddress, fldPhone, fldEmail FROM tblContact WHERE pkId = %s;",id)
	db := db_connect()
	defer db.Close()
	rows, _ := db.Query(sqlText)

	var contact *Contact
	var ret string
	var enc []byte
	
	contact = new(Contact)
	for rows.Next(){
	rows.Scan(&contact.Name, &contact.Address, &contact.Phone, &contact.Email)
		enc, _ = json.Marshal(contact)
		ret = string(enc)
	}
	
	return ret
}

func saveContact(id, name, address, phone, email string) string{
	db := db_connect()
	defer db.Close()
	var sqlText, pt1, pt2 string
	//var contact *Contact
	
	if id == "-1"{
		pt1 = "SELECT pkId FROM tblContact WHERE "
		pt2 = fmt.Sprintf("fldName = '%s' AND fldAddress = '%s' AND fldPhone = '%s' AND fldEmail = '%s';",name, address, phone, email)
		sqlText = pt1+pt2
		rows, _ := db.Query(sqlText)
		count := 0
		for rows.Next(){count++}
		if count == 0{
			pt1 = "INSERT INTO tblContact SET "
			pt2 = fmt.Sprintf("fldName = '%s', fldAddress = '%s', fldPhone = '%s', fldEmail = '%s';",name, address, phone, email)
			sqlText = pt1+pt2
		} else {
			return "same"
		}
	}else{
		pt1 = "UPDATE tblContact SET "
		pt2 = fmt.Sprintf("fldName = '%s', fldAddress = '%s', fldPhone = '%s', fldEmail = '%s' ",name, address, phone, email)
		pt3 := fmt.Sprintf("WHERE pkId = %s;",id)
		sqlText = pt1+pt2+pt3
	}
	db.Query(sqlText)
	
	return "ok"
}

func deleteContact(id string){
	sqlText := fmt.Sprintf("DELETE FROM tblContact WHERE pkId = %s;",id)
	db := db_connect()
	defer db.Close()
	db.Query(sqlText)
}