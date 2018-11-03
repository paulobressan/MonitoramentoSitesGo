//Pacote da aplicação atual
package main

//Importando pacotes externos
import (
	//pacote para trabalhar com texto
	"fmt"
	//pacote para manipular eventos do sistema operacional
	"os"
	//pacote para manipular requisições http
	"net/http"
	//pacote para analizar valores do códito
	"reflect"
	//Pacote time para manipular tempos de execução
	"time"
)

//constante de quantos monitoramento vai ser feito
const monitoramentos = 3

//Tempo de espera de cada monitoramento
const delay = 5

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
	fmt.Println("")
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
	//Array em go não pode iniciar sem um tamanho
	// var sites [4]string
	// sites[0] = "https://random-status-code.herokuapp.com/"
	// sites[1] = "https://google.com.br/"
	// sites[2] = "https://alura.com.br/"
	// sites[3] = "https://youtube.com.br/"

	//SLICES
	sites := []string{
		"https://random-status-code.herokuapp.com/",
		"https://google.com.br/",
		"https://alura.com.br/",
		"https://youtube.com.br/"}

	fmt.Println(sites)

	// For tradicional
	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println(sites[i])
	// }

	//For criado pelo Go com range
	// for index, site := range sites {
	// 	fmt.Println("Site passado", site, "Posição", index)
	// }
	for i := 0; i < monitoramentos; i++ {
		for index, site := range sites {
			fmt.Println("Testando site:", index, site)
			// realizando um get
			TestarSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")

	// Site para testar se esta online
	//site := "https://random-status-code.herokuapp.com/"

}

func TestarSite(site string) {
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

//trabalhando com coleções, slices
func exibirNomes() {
	nome := []string{"Paulo", "Bruna", "Rita"}
	//Adicionar novos itens no slice
	nome = append(nome, "Apararecida")
	fmt.Println(nome)
	//capturando o tipo da variavel
	fmt.Println(reflect.TypeOf(nome))
	//capturando o tamanho do array nome
	fmt.Println(len(nome))
	//capacidade do slice
	fmt.Println(cap(nome))
}
