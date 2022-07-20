package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramento = 3
const delay = 5

func main() {
	exibirMenu()
	comando := leComando()

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do Programa...")
		os.Exit(-1)
	default:
		fmt.Println("Comando inválido")
	}
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
	return comando
}

func exibirMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()

	// Trantando erro de não encontrar arquivo
	if sites == nil {
		fmt.Println("Não foi possível ler o arquivo sites.txt")
		return
	}

	for teste := 0; teste < monitoramento; teste++ {
		escreverLogs("################################################################################\n")
		fmt.Println("----------------------------------------------------")
		fmt.Println("===================== TESTE", teste+1, "======================")
		fmt.Println("----------------------------------------------------")
		fmt.Println("\t Website\t|\tStatus")
		fmt.Println("----------------------------------------------------")
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}

}

func testaSite(site string) {

	now := time.Now()
	horaDeAgora := now.Format("15:04:05")

	resp, _ := http.Get(site)
	dominio := strings.Split(site, ".")[1]
	dominio = strings.Split(dominio, ".")[0]
	if resp.StatusCode == 200 {
		fmt.Println("\t", dominio, "\t|\tDisponível")
		escreverLogs("\t" + horaDeAgora + " => \t" + dominio + " \t\t\t|\tDisponível\n")
	} else {
		fmt.Println("\t", dominio, "\t|\tIndisponível")
		escreverLogs("\t" + horaDeAgora + " => \t" + dominio + " \t\t\t|\tIndisponível\n")
	}
}

func leSitesDoArquivo() []string {
	arquivo, err := os.Open("sites.txt")
	// arquivo, err := ioutil.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo: ", err)
		return nil
	}

	var sites []string
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func verificaSeExisteArquivo() {
	if _, err := os.Stat("logs.txt"); os.IsNotExist(err) {
		file, err := os.Create("logs.txt")
		if err != nil {
			fmt.Println("Erro ao criar arquivo: ", err)
			return
		}

		escritor := bufio.NewWriter(file)
		now := time.Now()
		diaDeHoje := now.Format("02/01/2006")
		escritor.WriteString("================================================================================\n")
		escritor.WriteString("=================================[ " + diaDeHoje + " ]=================================\n")
		escritor.WriteString("================================================================================\n")
		escritor.WriteString("\n")
		escritor.Flush()
		file.Close()

	}
}

func escreverLogs(log string) {
	verificaSeExisteArquivo()

	logs, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo: ", err)
		return
	}

	escritor := bufio.NewWriter(logs)
	info, err := os.Stat("logs.txt")
	modDate := info.ModTime()

	// Verificar se o dia é diferente do dia atual
	if modDate.Day() != time.Now().Day() {
		now := time.Now()
		diaDeHoje := now.Format("02/01/2006")
		escritor.WriteString("\n================================================================================\n")
		escritor.WriteString("=================================[ " + diaDeHoje + " ]=================================\n")
		escritor.WriteString("================================================================================\n")
		escritor.WriteString("\n")
	}
	escritor.WriteString(log)
	// Salvar o arquivo
	escritor.Flush()
	logs.Close()

}
