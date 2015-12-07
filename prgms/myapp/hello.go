package kala

import (
"net/http"
"fmt"
)
func handler (res http.ResponseWriter, req *http.Request) {
fmt.Fprint(res, "ARE YOU PLANNING FOR YOUR HIGHER STUDIES ABROAD??\n")
fmt.Fprint(res, "IF YES ENTER /GRE or /GMAT in URL\n")

switch req.URL.Path{
	
case "/gre": {
	fmt.Fprint(res, "\nWrite GRE exam whose scale start from 160 to 340 \n ")
	fmt.Fprint(res, "Get minimum of 300 score")
}
	
case "/gmat": {
	fmt.Fprint(res, "\nwrite GMAT exam whose scale start from 200 to 800 \n ")
	fmt.Fprint(res, "get minimum of 550 score")
	}
default:
	fmt.Fprint(res, "\nnothing is entered in URL")
}

fmt.Fprint(res, "\nAlso write TOEFL or IELTS if you are international student \n ")
fmt.Fprint(res, "Get minimum of score 80 in TOEFL or score 6.5 in IELTS")
}

func init() {
http.HandleFunc("/", handler)
}