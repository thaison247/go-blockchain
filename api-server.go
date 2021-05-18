package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/thaison247/go-blKockchain/app"
	"github.com/thaison247/go-blockchain/configs"
	"github.com/thaison247/go-blockchain/structs"
	"net/http"
	"strconv"
	"time"
)

var nodeID = configs.GetConfig().NODE_ID

func main() {
	e := echo.New()

	app.WebController(e)

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "explorer.html", nil)
	})
	e.GET("/block/all", getBlocks)
	e.GET("/mempool", getMempool)
	//port := configs.GetConfig().WEB_PORT
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s",nodeID)))
}

func getBlocks(c echo.Context) error {
	bc := structs.NewBlockchain(nodeID)
	defer bc.DB.Close()

	res := make([]map[string]interface{},0)

	bci := bc.Iterator()

	for {
		block := bci.Next()

		var blockMap =  map[string]interface{}{}

		fmt.Printf("============ Block %x ============\n", block.Hash)
		blockMap["hash"] = fmt.Sprintf("%x", block.Hash)
		fmt.Printf("Height: %d\n", block.Height)
		blockMap["height"] = block.Height
		fmt.Printf("Prev. block: %x\n", block.PrevBlockHash)
		blockMap["prev_block"] = fmt.Sprintf("%x", block.PrevBlockHash)
		pow := structs.NewProofOfWork(block)
		fmt.Printf("Timestamp: %v\n", time.Unix(block.Timestamp, 0).Format("02-01-2006 03:04:05"))
		blockMap["timestamp"] = time.Unix(block.Timestamp, 0).Format("02-01-2006 03:04:05")
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
		blockMap["pow"] = strconv.FormatBool(pow.Validate())

		//blockMap["transactions"] = block.Transactions
		transactions := make([]map[string]interface{}, 0)
		for _, tx := range block.Transactions {
			txMap := map[string]interface{}{}
			fmt.Println(tx)
			txMap["id"] = fmt.Sprintf("%x", tx.ID)

			vin := make([]map[string]interface{}, 0)
			for _, in := range tx.Vin {
				txInput := make(map[string]interface{})
				txInput["txID"] = fmt.Sprintf("%x", in.Txid)
				txInput["vout"] = in.Vout
				txInput["signature"] = fmt.Sprintf("%x", in.Signature)
				txInput["pubKey"] = fmt.Sprintf("%x", structs.Base58Decode(in.PubKey))

				vin = append(vin, txInput)
			}
			txMap["vin"] = vin

			vout := make([]map[string]interface{}, 0)
			for _, out := range tx.Vout {
				txOutput := make(map[string]interface{})
				txOutput["value"] = out.Value
				txOutput["pubKeyHash"] = fmt.Sprintf("%x", out.PubKeyHash)

				vout = append(vout, txOutput)
			}
			txMap["vout"] = vout

			transactions = append(transactions, txMap)
		}

		blockMap["transactions"] = transactions
		fmt.Printf("\n\n")

		res = append(res, blockMap)

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return c.JSON(http.StatusOK, res)
}

func getMempool(c echo.Context) error {
	pendingTx := GetMemPool()
	return c.JSON(http.StatusOK, pendingTx)
}


