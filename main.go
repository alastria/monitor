package main

import (
	"fmt"
//	"math/big"
	"os"
	"time"
	"path/filepath"
	"encoding/hex"
	"math/rand"
	"io/ioutil"
	"strings"
	"io"
	"os/exec"

	"github.com/robfig/cron"
	"github.com/hashicorp/go-getter"
	//"github.com/alastria/monitor/services"
)

const STATIC_NODES = "/home/marcos/alastria/data/static-nodes.json"
const PERMISSIONED_NODES = "/home/marcos/alastria/data/permissioned-nodes.json"
var IDENTITY, NODE_TYPE, STATIC, PERMISSIONED string

func main() {
	IDENTITY = strings.Trim(getFile("/home/marcos/alastria/data/IDENTITY"), "\n")
	NODE_TYPE = strings.Trim(getFile("/home/marcos/alastria/data/NODE_TYPE"), "\n")
	STATIC = getFile(STATIC_NODES)
	PERMISSIONED = getFile(PERMISSIONED_NODES)
	update()
	c := cron.New()
	c.AddFunc("0 0 30 * * *", update)
	c.Start()

	for true {
		time.Sleep(10000 * time.Millisecond)
		fmt.Println("Esperando")
	}

	if true {
		fmt.Println("Terminé")
		return
	}

	/*var felicitador services.FelicitadorQuorumService

	argsWithoutProg := os.Args[1:]

	switch argsWithoutProg[0] {
	case "deploy":
		felicitador = services.NewFelicitadorQuorumService()
		address, err := felicitador.DeployFelicitadorContract()
		if err == nil {
			fmt.Print(fmt.Sprintf("%s", address))
		} else {
			fmt.Sprintln("Error: ", err)
		}
		break
	case "felicita":
		felicitador = services.NewFelicitadorQuorumServiceAddress(argsWithoutProg[1], "")

		err := felicitador.Felicita(argsWithoutProg[2], argsWithoutProg[3])
		if err != nil {
			fmt.Sprintln("Error: ", err)
		} else {
			fmt.Print("OK")
		}

		break
	case "cuantas":
		felicitador = services.NewFelicitadorQuorumServiceAddress(argsWithoutProg[1], "")

		cuantas, err := felicitador.LeerCuantasFelicitaciones()

		if err == nil {
			fmt.Print(fmt.Sprintf("%d", cuantas))
		} else {
			fmt.Sprintln("Error: ", err)
		}
		break
	case "felicitacion":
		felicitador = services.NewFelicitadorQuorumServiceAddress(argsWithoutProg[1], "")

		if len(argsWithoutProg) == 3 {
			var cual *big.Int
			i := new(big.Int)
			i, _ = i.SetString(argsWithoutProg[2], 10)
			cual = i
			felicitacion, err := felicitador.LeerFelicitacion(cual)
			if err == nil {
				fmt.Print(fmt.Sprintf("%s|&%s", felicitacion.NombreFelicitador, felicitacion.Mensaje))
			} else {
				fmt.Sprintln("Error: ", err)
			}

		}
		break
	default:
		fmt.Println("Forma de uso: \n felicitador deploy \n felicitador felicita <contract_address> <quien> <mensaje> \n felicitador cuantas <contract_address> \n felicitador felicitacion <contract_address> <cual>")
	}
*/
}

func update() {
	fmt.Printf("%s, %s, %s, %s", IDENTITY, NODE_TYPE, STATIC, PERMISSIONED)
	stfile, static := getGithub("https://raw.githubusercontent.com/alastria/alastria-node/feature/ibft/data/static-nodes.json")
	pmfile, permissioned := getGithub("https://raw.githubusercontent.com/alastria/alastria-node/feature/ibft/data/permissioned-nodes_" + NODE_TYPE + ".json")
	if strings.Compare(static, STATIC) != 0 || strings.Compare(permissioned, PERMISSIONED) != 0 {

		fmt.Println("Son distintos")
		if strings.Compare(static, STATIC) != 0 {
			fmt.Println("Actualizando static-nodes")
			copy(stfile, STATIC_NODES)
		}
		if strings.Compare(permissioned, PERMISSIONED) != 0 {
			fmt.Println("Actualizando permissioned-nodes")
			copy(pmfile, PERMISSIONED_NODES)
		}
		fmt.Println(strings.Trim(static, "]"), strings.Trim(STATIC, "]"))
		if !strings.Contains(strings.Trim(static, "]"), strings.Trim(STATIC, "]")) ||
				!strings.Contains(strings.Trim(permissioned, "]"), strings.Trim(PERMISSIONED, "]")) {
			fmt.Println("Hay que reiniciar el nodo...")
			out, err := exec.Command("$HOME/alastria-node/scripts/stop.sh").Output()
			fmt.Println(out, err)
			time.Sleep(15000 * time.Millisecond)
			out, err = exec.Command("$HOME/alastria-node/scripts/start.sh", "all").Output()
			fmt.Println(out, err)
		}
	}
}

func getGithub(url string) (filename, contenido string) {
	filename = tempFileName("monitor", ".json")
	err := getter.GetFile(filename, url)
	if err != nil {
		fmt.Println(err)
	}
	if err == nil {
		contenido = getFile(filename)
	}
	return
}

func getFile(fichero string) (contenido string) {
	data, _ := ioutil.ReadFile(fichero)
	contenido = string(data)
	return
}

func copy(stfrom, stto string) {
	fmt.Println(stfrom, stto)
	from, err := os.Open(stfrom)
	if err != nil {
		fmt.Println(err)
	}
	defer from.Close()
	to, err := os.Open(stto)
	if err != nil {
		fmt.Println(err)
	}
	defer to.Close()
	io.Copy(from, to)
}

func tempFileName(prefix, suffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return filepath.Join(os.TempDir(), prefix+hex.EncodeToString(randBytes)+suffix)
}
