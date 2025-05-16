package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/zanjava/learn/store"
	"github.com/zanjava/learn/token"
	"log"
	"math"
	"math/big"
)

const (
	contractAddr = "0x32e74A928EA044993e40EC9cD8Dd6f0090c8E6CD"
)

func main() {
	//client := createClient()
	// 查询区块头信息，若传入 nil，它将返回最新的区块头
	//blockNumber := big.NewInt(5671744)

	// 查询区块头信息
	//getHeader(client, blockNumber)

	//chainID := getchainID(client)
	//fmt.Println("chainId:", chainID)

	//block := getBlockByNumber(client, blockNumber)

	//getReceiptByBlockHash(client, block.Hash())

	//getReceiptByBlockNumber(client, block.Number())

	//count := getBlockTransactionsCount(client, block.Hash())
	//fmt.Println("count:", count)
	//
	//for _, tx := range block.Transactions() {
	//	fmt.Println(tx.Hash().Hex())        // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
	//	fmt.Println(tx.Value().String())    // 100000000000000000
	//	fmt.Println(tx.Gas())               // 21000
	//	fmt.Println(tx.GasPrice().Uint64()) // 100000000000
	//	fmt.Println(tx.Nonce())             // 245132
	//	fmt.Println(tx.Data())              // []
	//	fmt.Println(tx.To().Hex())          // 0x8F9aFd209339088Ced7Bc0f57Fe08566ADda3587
	//
	//	if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
	//		fmt.Println("sender", sender.Hex()) // 0x2CdA41645F2dBffB852a605E92B185501801FC28
	//	} else {
	//		log.Fatal(err)
	//	}
	//
	//	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println(receipt.Status) // 1
	//	fmt.Println(receipt.Logs)   // []
	//	break
	//}

	//for idx := uint(0); idx < count; idx++ {
	//	tx, err := client.TransactionInBlock(context.Background(), block.Hash(), idx)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
	//	break
	//}
	//
	//txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	//tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(isPending)
	//fmt.Println(tx.Hash().Hex())

	//createNewWallet()

	//transferETH(client)

	//client := createClient()
	//queryAccountBalance(client, "0xE1B4Fd26B91eEc054a612E148a7C85753Af3d44e")
	//amount := new(big.Int)
	//amount.SetString("1000000000000000000000", 10)
	//transferToken(client, "0x28b149020d2152179873ec60bed6bf7cd705775d", "0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d", amount)

	//queryTokenBalance()
	//subscribeNewHeads()

	//使用 abigen 工具部署合约
	/*client := createClient()
	privateKey, err := crypto.HexToECDSA("f7dfbe2b391d314c391e26f2099c7b61dbde5069c810346d3c19b68c00ebd1e0")
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

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance*/

	// 加载合约
	client := createClient()
	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("f7dfbe2b391d314c391e26f2099c7b61dbde5069c810346d3c19b68c00ebd1e0")
	if err != nil {
		log.Fatal(err)
	}

	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("demo_save_key"))
	copy(value[:], []byte("demo_save_value11111"))

	chainID := getchainID(client)

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	tx, err := storeContract.SetItem(opt, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	valueInContract, err := storeContract.Items(callOpt, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("value in contract:", valueInContract)
	fmt.Println("is value saving in contract equals to origin value:", valueInContract == value)

}

func subscribeNewHeads() {
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/9cbc77586cd54219823a4c2a6d4d3cb6")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println("header.Number:", header.Number.Uint64())         // 5671744
			fmt.Println("header.Time:", header.Time)                      // 1712798400
			fmt.Println("header.Difficulty:", header.Difficulty.Uint64()) // 0
			fmt.Println("header.Hash:", header.Hash().Hex())
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}

/**
 * 查询代币余额
 */
func queryTokenBalance() {
	client := createClient()
	// Golem (GNT) Address
	tokenAddress := common.HexToAddress("0x7229939f9d747b6F934377FA023d9A0951846a57")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0xE1B4Fd26B91eEc054a612E148a7C85753Af3d44e")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
	fmt.Printf("wei: %s\n", bal)           // "wei: 74605500647408739782407023"
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}

/**
 * 查询账户余额
 */
func queryAccountBalance(client *ethclient.Client, account string) {
	account1 := common.HexToAddress(account)
	balance, err := client.BalanceAt(context.Background(), account1, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account1)
	fmt.Println(pendingBalance)
}

/**
 * 创建客户端
 */
func createClient() *ethclient.Client {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/9cbc77586cd54219823a4c2a6d4d3cb6")
	if err != nil {
		log.Fatal(err)
	}
	return client
}

/**
 * 查询区块头信息
 */
func getHeader(client *ethclient.Client, blockNumber *big.Int) *types.Header {
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("header.Number:", header.Number.Uint64())         // 5671744
	fmt.Println("header.Time:", header.Time)                      // 1712798400
	fmt.Println("header.Difficulty:", header.Difficulty.Uint64()) // 0
	fmt.Println("header.Hash:", header.Hash().Hex())              // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	return header
}

/**
 * 查询链ID
 */
func getchainID(client *ethclient.Client) *big.Int {
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chainID:", chainID) //11155111
	return chainID
}

/**
 * 查询区块信息
 */
func getBlockByNumber(client *ethclient.Client, blockNumber *big.Int) *types.Block {
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("block.Number:", block.Number().Uint64())               // 5671744
	fmt.Println("block.Time:", block.Time())                            // 1712798400
	fmt.Println("block.Difficulty:", block.Difficulty().Uint64())       // 0
	fmt.Println("block.Hash:", block.Hash().Hex())                      // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	fmt.Println("block.Transactions count:", len(block.Transactions())) // 70
	return block
}

/**
 * 查询区块交易数量
 */
func getBlockTransactionsCount(client *ethclient.Client, blockHash common.Hash) uint {
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transactions count:", count)
	return count
}

/**
 * 通过区块hash查询区块交易
 */
func getReceiptByBlockHash(client *ethclient.Client, blockHash common.Hash) []*types.Receipt {
	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err)
	}
	for _, receipt := range receiptByHash {
		fmt.Println(receipt.Status)           // 1
		fmt.Println(receipt.Logs)             // []
		fmt.Println(receipt.TxHash.Hex())     // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		fmt.Println(receipt.TransactionIndex) // 0
		break
	}
	return receiptByHash
}

/**
 * 通过区块number查询区块交易
 */
func getReceiptByBlockNumber(client *ethclient.Client, blockNumber *big.Int) []*types.Receipt {
	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal(err)
	}
	for _, receipt := range receiptsByNum {
		fmt.Println(receipt.Status)           // 1
		fmt.Println(receipt.Logs)             // []
		fmt.Println(receipt.TxHash.Hex())     // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		fmt.Println(receipt.TransactionIndex) // 0
		break
	}
	return receiptsByNum
}

/**
 * 创建新钱包
 */
func createNewWallet() {
	//如果已经有了私钥的 Hex 字符串，也可以使用 HexToECDSA 方法恢复私钥：
	//privateKey, err := crypto.HexToECDSA("ccec5314acec3d18eae81b6bd988b844fc4f7f7d3c828b351de6d0fede02d3f2")
	//if err != nil {
	//	log.Fatal(err)
	//}
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 去掉'0x'
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04'
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)
	//hash := sha3.NewLegacyKeccak256()
	//hash.Write(publicKeyBytes[1:])
	//fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	//fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 原长32位，截去12位，保留后20位
}

/**
 * ETH转账
 */
func transferETH(client *ethclient.Client) {
	privateKey, err := crypto.HexToECDSA("f7dfbe2b391d314c391e26f2099c7b61dbde5069c810346d3c19b68c00ebd1e0")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("fromAddress:", fromAddress.Hex()) // 0xE1B4Fd26B91eEc054a612E148a7C85753Af3d44e
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(10000000000) // in wei (1 eth)
	gasLimit := uint64(21000)        // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	fmt.Println("chainID:", chainID) //11155111
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

/**
 * 代币转账
 */
func transferToken(client *ethclient.Client, tokenAddress string, toAddress string, amount *big.Int) {
	// Load your private key
	privateKey, err := crypto.HexToECDSA("f7dfbe2b391d314c391e26f2099c7b61dbde5069c810346d3c19b68c00ebd1e0")
	if err != nil {
		log.Fatal(err)
	}

	// Get the public address from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get the nonce for the sender's address
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Create the token contract address
	tokenContractAddress := common.HexToAddress(tokenAddress)

	// Create the destination address
	to := common.HexToAddress(toAddress)

	// Create the token transfer data
	// This is the ABI-encoded function call to the 'transfer' method
	// Function signature: transfer(address,uint256)
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := crypto.Keccak256(transferFnSignature)
	methodID := hash[:4]

	// Pad the address to 32 bytes
	paddedAddress := common.LeftPadBytes(to.Bytes(), 32)

	// Pad the amount to 32 bytes
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	// Combine the method ID and parameters
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// Estimate gas limit
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &tokenContractAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create the transaction
	tx := types.NewTransaction(nonce, tokenContractAddress, big.NewInt(0), gasLimit, gasPrice, data)

	// Get chain ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Sign the transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Token transfer sent: %s\n", signedTx.Hash().Hex())
}
