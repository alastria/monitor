package services

import (
	"math/big"
	"strings"

	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/alastria/felicitador/services/bindings"
)

type FelicitadorQuorumService struct {
	QuorumServices
	contractAddress string
}

func NewFelicitadorQuorumService() FelicitadorQuorumService {
	prefijo := beego.BConfig.RunMode + "::"
	return NewFelicitadorQuorumServiceAddress(beego.AppConfig.String(prefijo+"quorum.FelicitadorAddress"),
		prefijo)
}

func NewFelicitadorQuorumServiceAddress(address string, prefijo string) FelicitadorQuorumService {
	var err error

	newnetwork := FelicitadorQuorumService{
		NewQuorumServices(prefijo),
		address}
	if err != nil {
		newnetwork.log.Error("FelicitadorQuorumService no inicializada correctamente.", err)
	}
	return newnetwork
}

func (n FelicitadorQuorumService) Init() (wasOk bool, err error) {
	//n.log.Debug("FelicitadorQuorumService - Init")

	wasOk = true
	//n.log.Debug("FelicitadorQuorumService - End: %v", wasOk, err)
	return wasOk, err
}

func (n FelicitadorQuorumService) Felicita(felicitador string, mensaje string) (err error) {
	//n.log.Debug("Felicita")
	var transactor *bind.TransactOpts
	var contract *bindings.Felicitador

	transactor, err = n.createTransactor()
	if err == nil {
		contract, err = n.getContract()
		if err == nil {
			_, err = contract.Felicita(transactor, felicitador, mensaje)
		}
	}
	return err
}

func (n FelicitadorQuorumService) LeerCuantasFelicitaciones() (cuantas *big.Int, err error) {
	//n.log.Debug("LeerCuantasFelicitaciones")
	var contract *bindings.Felicitador
	contract, err = n.getContract()

	if err == nil {
		cuantas, err = contract.LeerCuantasFelicitaciones(&bind.CallOpts{})
	}

	return cuantas, err
}

func (n FelicitadorQuorumService) LeerFelicitacion(cual *big.Int) (felicitacion struct {
	NombreFelicitador string
	Mensaje           string
}, err error) {
	//n.log.Debug("LeerFelicitacion")

	var contract *bindings.Felicitador

	contract, err = n.getContract()
	if err == nil {
		felicitacion, err = contract.LeerFelicitacion(&bind.CallOpts{}, cual)
	}

	if err != nil {
		felicitacion.Mensaje = "Feliz Navidad"
		felicitacion.NombreFelicitador = "ERROR"
		err = nil
	}

	return felicitacion, err
}

func (n FelicitadorQuorumService) DeployFelicitadorContract() (string, error) {
	//n.log.Debug("FelicitadorQuorumService - DeployFelicitadorContract")
	client, err := n.connectToQuorumEthclient()
	if err == nil {
		transactor, err := n.createTransactor()

		if err == nil {
			// Deploy a new awesome contract for the binding demo
			address, _, Felicitador, err := bindings.DeployFelicitador(transactor, client)

			if err != nil || Felicitador == nil {
				n.log.Error("Failed to deploy new Felicitador contract: %v", err)
				return "", err
			}
			//n.log.Debug("Contract pending deploy: ", strings.ToLower(address.Hex()))
			//n.log.Debug("Transaction waiting to be mined: ", strings.ToLower(tx.Hash().Hex()))

			n.contractAddress = strings.ToLower(address.Hex())
		}
	}
	// Don't even wait, check its presence in the local pending state
	/*time.Sleep(250 * time.Millisecond)*/ // Allow it to be processed by the local node :P
	return n.contractAddress, err
}

func (n FelicitadorQuorumService) getContract() (contract *bindings.Felicitador, err error) {
	//n.log.Trace("FelicitadorQuorumService - setContract")
	var client *ethclient.Client

	client, err = n.connectToQuorumEthclient()
	if err == nil {
		if contract == nil {
			contract, err =
				bindings.NewFelicitador(
					common.HexToAddress(n.contractAddress), client)
		}
	}

	if err != nil {
		n.log.Debug("Object resume: %v", n.toString())
		n.log.Error("Failed to connect to the Ethereum client: %v", err)
	}
	return contract, err
}

func (n FelicitadorQuorumService) toString() (salida string) {
	salida = "{"
	salida += "'Host': \"" + n.host + "\", "
	salida += "'Port': \"" + strconv.Itoa(n.port) + "\", "
	salida += "'ContractAddress': \"" + n.contractAddress + "\", "
	salida += "'Prefijo': \"" + string(n.prefijo) + "\""
	salida += "}"
	return salida
}
