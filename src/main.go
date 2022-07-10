package main

import(
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main()  {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)

	if err != nil{
		fmt.Printf("Server Closed")
	}else {
		fmt.Printf("Server is running")
	}
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hi This is Yuvi !!")
}