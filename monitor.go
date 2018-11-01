//Pacote da aplicação atual
package main

//Importando pacotes externos
import (
	//pacote para trabalhar com texto
	"fmt"
	"net/http"

	//pacote para manipular eventos do sistema operacional
	"os"
	//pacote para manipular requisições http
)

//Função principal
func main() {
	for {
		exibeMenu()
		//Auto declarando variavel e recebendo o valor da função
		comando := lerComando()
		// IFS
		// if comando == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo logs...")
		// } else if comando == 0 {
		// 	fmt.Println("Saindo do programa")
		// } else {
		// 	fmt.Println("Não conheço este")
		// }

		//Switch case
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			//terminar execução com sucesso
			os.Exit(0)
		default:
			fmt.Println("Não conheço este")
			//terminar execução exibindo algum problema
			os.Exit(-1)
		}
	}
	//FUNÇÃO DE MULTIPLOS RETORNOS
	//recebendo multiplos valores de uma função, se não for do interece um dos valores, pordemos utilizar _ para ignorar
	// _, idade := devolveNomeEIdade()
	// fmt.Println("Idade :", idade)
	// nome, idade := devolveNomeEIdade()
	// fmt.Println(nome, "Idade :", idade)

	exibirIntroducao()
}

//Funções no go
func exibirIntroducao() {
	//Variavel auto declarativa
	//Exemplos,(var nome string = "Paulo", var nome = "Paulo")
	nome := "Paulo"
	versao := 1.0

	//Imprimir texto no console com auxilo do fmt
	fmt.Println("Olá, sr", nome)
	fmt.Println("A versão atual do sistema é", versao)
}

//Função com o retorno int
func lerComando() int {
	var comando int
	//Ler dados inserido no terminal com o fmt
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	//Retornando comando
	return comando
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

//Manipulando requisições http
func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// Site para testar se esta online
	site := "https://random-status-code.herokuapp.com/"
	// realizando um get
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "esta problema para carregar. Status Code:", resp.StatusCode)
	}
}

//função que retornar mais de um valor
func devolveNomeEIdade() (string, int) {
	nome := "Paulo"
	idade := 20
	return nome, idade
}
