package main

import (
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./estatico"))
	http.Handle("/", fs)
	log.Print("Listening on : 3000...")
	err := http.ListenAndServe("3000...", nil)
	if err != nil {
		log.Fatal(err)

	}

}

////////////////////////////////////------------------------------------////////////////////////////////////


<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8"/>
<title>O Servi do Winicius</title>
</head>
<body>
<h1> Bem Vindo a o Serve do Winicius</h1>
</body>
</html>
