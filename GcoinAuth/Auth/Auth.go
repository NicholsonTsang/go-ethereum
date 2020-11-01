package Auth

import (
	"github.com/ethereum/go-ethereum/GcoinAuth/config"
	"github.com/ethereum/go-ethereum/GcoinAuth/contracts"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

var(
	configuration = config.GetConfig()
	client, err1 = ethclient.Dial(configuration.CLIENT_ADDRESS)
	address = common.HexToAddress(configuration.AUTH_CONTRACT_ADDRESS)
	instance, err2 = contracts.NewNodeRegistrationImpl(address, client)
	pw string
	privatekey string
)

func CheckAdmin() bool{
	//configuration := config.GetConfig()
	//fmt.Println(configuration.AUTH_CONTRACT_ADDRESS)
	//client, err1 := ethclient.Dial(configuration.CLIENT_ADDRESS)
	//address := common.HexToAddress(configuration.AUTH_CONTRACT_ADDRESS)
	//instance, err2 := contracts.NewNodeRegistrationImpl(address, client)
	//configuration = configuration.GetConfig()
	client, err1 = ethclient.Dial(configuration.CLIENT_ADDRESS)
	if err1 != nil {

		fmt.Println(configuration.CLIENT_ADDRESS)
		log.Fatal(err1)
	}
	if err2 != nil {
		fmt.Println("error2")
		log.Fatal(err2)
	}
	//isAdmin, err := instance.SetItem(auth,[32]byte{},[32]byte{})
	hashedAdminPW, err := instance.HashedAdminPW(nil)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	fmt.Print("ENTER ADMIN PW: ")
	fmt.Scan(&pw)
	hashedPW := crypto.Keccak256([]byte(pw))
	slicehashedPW := hashedAdminPW[:]
	correctness := Equal(slicehashedPW,hashedPW)
	return correctness
}

func Signing() *bind.TransactOpts{
	fmt.Print("ENTER PRIVATE KEY FOR CONTRACT SIGNING: ")
	fmt.Scan(&privatekey)
	privateKey, err := crypto.HexToECDSA(privatekey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(4600000) // in units
	auth.GasPrice = gasPrice
	return auth
}

func AddPeerStateUpdate(hexApplicantAddress string) bool{
	if err1 != nil {
		log.Fatal(err1)
	}
	if err2 != nil {
		log.Fatal(err2)
	}
	auth := Signing()
	applicantAddress := common.HexToAddress(hexApplicantAddress)

	result, err := instance.AddingPeer(auth,applicantAddress)
	fmt.Println(result)
	if err != nil {
		fmt.Println("here")
		log.Fatal(err)
		return false
	}
	return true
}

func Equal(slice1, slice2 []byte) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
/***
func main()  {
	//fmt.Println(configuration.AUTH_CONTRACT_ADDRESS)
	correctness := checkAdmin()
	//correctness := addPeerStateUpdate("0xB0d66eD848f2AA34D7fC8B5a9F62B70721F44084")
	//correctness, _ := instance.HashedAdminPW(nil)
	//fmt.Println(configuration.AUTH_CONTRACT_ADDRESS)
	fmt.Print(correctness)
	//fmt.Println(configuration.AUTH_CONTRACT_ADDRESS)
}
***/
