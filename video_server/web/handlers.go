package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"go_project/video_server/api/defs"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 真正的登录页面会先判断一下当前用户是否是注册用户
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")
	if err1 != nil || err2 != nil {
		p := &HomePage{Name: "ziyang"}
		t, err := template.ParseFiles("../template/home.html")
		if err != nil {
			log.Printf("Parsing template home.html error: %s", err)
			return
		}
		t.Execute(w, p)
		return
	}
	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		http.Redirect(w, r, "/username", http.StatusFound)
		return
	}
}

func userHomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cname, err1 := r.Cookie("username")
	_, err2 := r.Cookie("session")
	if err1 != nil || err2 != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	fname := r.FormValue("username")
	var p *UserPage
	if len(cname.Value) != 0 {
		p = &UserPage{Name: cname.Value}
	} else if len(fname) != 0 {
		p = &UserPage{Name: fname}
	}
	t, err := template.ParseFiles("../template/userhome.html")
	if err != nil {
		log.Printf("Parsing userhome.html error:%s", err)
		return
	}
	t.Execute(w, p)
}

func apiHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != http.MethodPost {
		re, _ := json.Marshal(ErrorRequestNotRecognized)
		io.WriteString(w, string(re))
		return
	}
	res, _ := ioutil.ReadAll(r.Body)
	apibody := &ApiBody{}
	if err := json.Unmarshal(res, apibody); err != nil {
		re, _ := json.Marshal(defs.ErrorRequestBodyParseFailed)
		io.WriteString(w, string(re))
		return
	}
	request(apibody, w, r)
	defer r.Body.Close()
}
