
// Servidor service - servidor de serviço, que irá prestar algo ao usúario
//

package main

import (
	"fmt"
	"net/http"
)

func ola(w http.ResponseWriter, req *http.Request) {
	// as funções que servem como manipuladores recebem a http.ResponseWriter e a
	// http.Request como argumentos. O gravador de resposta é usada para preencher a
	// resposta HTTP. Aqui nossa resposta simples é anelas um "Olá\n"
	fmt.Fprintf(w, "Ola\n")
}

func cabecalhos(w http.ResponseWriter, req *http.Request) {
	// Esse manipulador faz algo um pouco mais sofisticado lendo todos os cabeçalhos de solicitação
	// HTTP e ecoando-os no corpo da resposta

	for nome, cabecalhos := range req.Header {
		for _, c := range cabecalhos {
			fmt.Fprintf(w, "%v: %v\n", nome, c)
		}
	}
}

func main() {
	//Um manipulador é um objeto que implementa http.handler
	// Uma maneira comum de escrever um manipulador(handler) é usar
	//o http.HandleFunc, rota de função que ele deve chamar "/Olá e /Cabecalho"
	//Ele configura o roteador padrão no pacote net/http e recebe uma função
	// como argumento. ("", Ola) e ("", cabecalhos)

	http.HandleFunc("/Ola", ola)
	http.HandleFunc("/cabecalhos", cabecalhos)

	//finalmente , chamamos o ListenAndServe com a porta ":8090"
	//e um manipulador nil para  que seja usado o roteador padrão que acabamos de configurar
	// Execute o servidor em segundo plano
	// acesse :http://localhost:8089/ola
	// http://localhost:8089/cabecalhos

	http.ListenAndServe(":8090", nil)

}
