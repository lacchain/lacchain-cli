package blockchain

import (
	"context"
//	"time"
	"fmt"
	"math/big"
	"strings"
	"crypto/ecdsa"
	//"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	//"github.com/lacchain/lacchain-cli/model"
	"github.com/lacchain/lacchain-cli/errors"
	log "github.com/lacchain/lacchain-cli/util"
	permissioning "github.com/lacchain/lacchain-cli/blockchain/contracts"
	//"github.com/ethereum/go-ethereum/common/hexutil"
)

const RelayABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_blocksFrequency\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_accountIngress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSigner\",\"type\":\"bool\"}],\"name\":\"BadSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIRelayHub.ErrorCode\",\"name\":\"errorCode\",\"type\":\"uint8\"}],\"name\":\"BadTransactionSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractDeployed\",\"type\":\"address\"}],\"name\":\"ContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"countExceeded\",\"type\":\"uint8\"}],\"name\":\"GasLimitExceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsedLastBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"averageLastBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGasLimit\",\"type\":\"uint256\"}],\"name\":\"GasLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsedLastBlocks\",\"type\":\"uint256\"}],\"name\":\"GasUsedByTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"NodeBlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"Recalculated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Relayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"TransactionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relay\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"charge\",\"type\":\"uint256\"}],\"name\":\"TransactionRelayed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNode\",\"type\":\"address\"}],\"name\":\"addNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"}],\"name\":\"deleteNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_byteCode\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"senderSignature\",\"type\":\"bytes\"}],\"name\":\"deployMetaTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"deployedAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasUsedLastBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMsgSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"encodedFunction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"senderSignature\",\"type\":\"bytes\"}],\"name\":\"relayMetaTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accountIngress\",\"type\":\"address\"}],\"name\":\"setAccounIngress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_blocksFrequency\",\"type\":\"uint8\"}],\"name\":\"setBlocksFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGasUsed\",\"type\":\"uint256\"}],\"name\":\"setGasUsedLastBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

//Client to manage Connection to Ethereum
type Client struct {
	client *ethclient.Client
}

//Connect to Ethereum
func (ec *Client) Connect(nodeURL string) error {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		msg := fmt.Sprintf("Can't connect to node %s", nodeURL)
		err = errors.FailedConnection.Wrapf(err,msg)
		return err
	}

	fmt.Sprintf("Connected to Ethereum Node:", nodeURL)
	ec.client = client
	return nil
}

//Close ethereum connection
func (ec *Client) Close() {
	ec.client.Close()
}

//GetEthclient ...
func (ec *Client) GetEthclient()(*ethclient.Client) {
	return ec.client
}

//ConfigCall for ethereum address
func (ec *Client) ConfigCall(fromAddress common.Address) (*bind.CallOpts){
	return &bind.CallOpts{Pending: false, From:fromAddress, BlockNumber:nil, Context:nil}
}

//ConfigTransaction from ethereum address contract
func (ec *Client) ConfigTransaction(key *ecdsa.PrivateKey, gasLimit uint64) (*bind.TransactOpts, error) {
/*	keystore, err := os.Open(keyStorePath)
	defer keystore.Close()
    if err != nil {
		msg := fmt.Sprintf("could not load keystore from location %s",keyStorePath)
		err = errors.FailedKeystore.Wrapf(err,msg)
		return nil,err
    }
*/	
	auth := bind.NewKeyedTransactor(key)

	nonce, err := ec.client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		msg := fmt.Sprintf("can't get pending nonce for:%s",auth.From)
		err = errors.FailedConfigTransaction.Wrapf(err,msg)
		return nil, err
	}

	gasPrice, err := ec.client.SuggestGasPrice(context.Background())
	if err != nil {
		msg := fmt.Sprintf("can't get gas price suggested")
		err = errors.FailedConfigTransaction.Wrapf(err,msg)
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = gasLimit + 300000 // in units
	auth.GasPrice = gasPrice

	log.GeneralLogger.Println("OptionsTransaction=[From:0x%x,nonce:%d,gasPrice:%s,gasLimit:%d", auth.From,nonce,gasPrice,auth.GasLimit)

	return auth, nil
}

//0x72756c6573000000000000000000000000000000000000000000000000000000
//GetNodeIngressContract ...
func (ec *Client) GetNodeIngressContract(contractAddress common.Address, name [32]byte)(string, error){
	contract, err := permissioning.NewPermissioning(contractAddress, ec.client)
	if err != nil {
		msg := fmt.Sprintf("can't instance NodeIngress contract %s", contractAddress)
		err = errors.FailedContract.Wrapf(err,msg)
		return "", err
	}

	nodeIngress, err := contract.GetContractAddress(&bind.CallOpts{},name)
	if err != nil {
		msg := fmt.Sprintf("can't get NodeRules contract address")
		err = errors.FailedContract.Wrapf(err,msg)
		return "", err
	}

	log.GeneralLogger.Println("NodeRules contract address:",nodeIngress.Hex())

	return nodeIngress.Hex(), nil
}

//GetNodeInformation ...
func (ec *Client) GetNodeInformation(contractAddress common.Address, enodeHigh, enodeLow [32]byte, enodeIp [16]byte, port uint16)(string, error){
	contract, err := permissioning.NewRules(contractAddress, ec.client)
	if err != nil {
		msg := fmt.Sprintf("can't instance NodeRules contract %s", contractAddress)
		err = errors.FailedContract.Wrapf(err,msg)
		return "", err
	}

	node, err := contract.GetByEnode(&bind.CallOpts{},enodeHigh,enodeLow,enodeIp,port)
	if err != nil {
		msg := fmt.Sprintf("can't get Node Information from NodeRules contract")
		err = errors.FailedContract.Wrapf(err,msg)
		return "", err
	}

	log.GeneralLogger.Println("Node:",node)

	return "", nil
}

//GenerateTransaction ...
func (ec *Client)GenerateTransaction(gasLimitTx uint64, relayAddress common.Address, from common.Address, to *common.Address, encodedFunction []byte, gasLimit *big.Int, nonce *big.Int, signature []byte, senderSignature []byte) (*types.Transaction){
	testabi, err := abi.JSON(strings.NewReader(RelayABI))
	if err != nil{
		log.GeneralLogger.Println("Error decoding ABI")
	}

	var bytesData []byte

	if (to!=nil){
		bytesData, _ = testabi.Pack("relayMetaTx",from,*to,encodedFunction,gasLimit,nonce,signature,senderSignature)
	}else{
		bytesData, _ = testabi.Pack("deployMetaTx",from,encodedFunction,gasLimit,nonce,signature,senderSignature)
	}
		tx := types.NewTransaction(0, relayAddress, big.NewInt(0), gasLimitTx, big.NewInt(0), bytesData)
	return tx
}

//GetTransactionReceipt ...
func (ec *Client) GetTransactionReceipt(transactionHash common.Hash)(*types.Receipt, error){
	receipt, err := ec.client.TransactionReceipt(context.Background(), transactionHash)
    if err != nil {
        return nil, err
	}
		
	log.GeneralLogger.Println("Status:",receipt.Status)
	log.GeneralLogger.Println("BlockNumber:",receipt.BlockNumber)

	return receipt,nil;
}