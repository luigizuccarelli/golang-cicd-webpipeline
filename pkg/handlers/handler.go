package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"lmzsoftware.com/lzuccarelli/golang-cicd-webconsole/pkg/connectors"
	"lmzsoftware.com/lzuccarelli/golang-cicd-webconsole/pkg/schema"
)

const (
	CONTENTTYPE     string = "Content-Type"
	APPLICATIONJSON string = "application/json"
)

func ResultsHandler(w http.ResponseWriter, r *http.Request, conn connectors.Clients) {
	var res *schema.ResultsSchema

	vars := mux.Vars(r)

	if vars["project"] == "" {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "ERROR missing param")
		return
	}

	// open the directory for the project
	files, err := ioutil.ReadDir(os.Getenv("CICD_CONSOLE_DIR") + "/" + vars["project"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "ERROR project folder")
		return
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().Before(files[j].ModTime())
	})

	//fmt.Println("DEBUG LMZ", files)

	var items []schema.ItemInfo
	var item = &schema.ItemInfo{}
	for _, f := range files {
		// ignore errors from pipelinecontent
		data, _ := ioutil.ReadFile(os.Getenv("CICD_CONSOLE_DIR") + "/" + vars["project"] + "/" + f.Name())
		item.Name = strings.ToUpper(f.Name()[:len(f.Name())-4])
		item.Time = f.ModTime().Unix()
		unixTimeUTC := time.Unix(f.ModTime().Unix(), 0)
		item.DisplayTime = fmt.Sprintf("%v", unixTimeUTC)
		replacer := strings.NewReplacer("\n", "<br>", "\x1b[1;34m [INFO] \x1b[0m", "<span style=\"color:#3498eb\">&nbsp;[INFO]</span>","\x1b[1;32m [DEBUG] \x1b[0m","<span style=\"color:#34eb40\">&nbsp;[DEBUG]</span>","\x1b[1;36m [TRACE] \x1b[0m","<span style=\"color:#d234eb\">&nbsp;[TRACE]</span>","\x1b[1;31m [ERROR] \x1b[0m","<span style=\"color:#eb4034\">&nbsp;[ERROR]</span>","INFO:","<span style=\"color:#3498eb\">&nbsp;INFO:</span>")
		item.Log =  replacer.Replace(string(data))
		if item.Log[0:3] == "PAS" {
			item.Pass = true
			item.Fail = false
		}
		if item.Log[0:3] == "ERR" {
			item.Pass = false
			item.Fail = true
		}
		items = append(items, *item)
		item = &schema.ItemInfo{}
	}
	addHeaders(w,r)
	res = &schema.ResultsSchema{StartTime: items[0].DisplayTime, StopTime: items[len(items)-1].DisplayTime, Items: items}
	html, err := ioutil.ReadFile(os.Getenv("TEMPLATE_FILE"))
	if err != nil {
		conn.Error("Rsultshandler %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "ERROR creating template")
		return
	}
	var tpl bytes.Buffer
	tmpl := template.New(vars["project"])
	tmp, er := tmpl.Parse(string(html))
	if er != nil {
		conn.Error("Rsultshandler %v", er)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "ERROR parsing template")
		return
	}
	tmp.Execute(&tpl, res)
	conn.Debug("ResultsHandler response : %s", tpl.String())
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", tpl.String())
}

func IsAlive(w http.ResponseWriter, r *http.Request) {
	addHeaders(w, r)
	fmt.Fprintf(w, "{ \"version\" : \""+os.Getenv("VERSION")+"\" , \"name\": \""+os.Getenv("NAME")+"\" }")
	return
}

// headers (with cors) utility
func addHeaders(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("API-KEY") != "" {
		w.Header().Set("API_KEY_PT", r.Header.Get("API_KEY"))
	}
	w.Header().Set(CONTENTTYPE, APPLICATIONJSON)
	// use this for cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept-Language")
}
