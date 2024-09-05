package controller

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/account"
	"github.com/fbsobreira/gotron-sdk/pkg/keys"
	"github.com/fbsobreira/gotron-sdk/pkg/mnemonic"
	"github.com/fbsobreira/gotron-sdk/pkg/store"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
)

func CreateWallet(ctx *gin.Context) {
	mnemonicStr := mnemonic.Generate()
	name := generateRandomName(10)
	creation := &account.Creation{
		Name:               name,
		Passphrase:         "",
		Mnemonic:           mnemonicStr,
		MnemonicPassphrase: "",
		HdAccountNumber:    new(uint32),
		HdIndexNumber:      new(uint32),
	}
	err := account.CreateNewLocalAccount(creation)
	if err != nil {
		fmt.Printf("Error creating account: %s\n", err)
	} else {
		fmt.Println("Account created successfully")
	}
	accountName, err := store.AddressFromAccountName(name)
	if err != nil {
		return
	}
	ppk, _ := keys.FromMnemonicSeedAndPassphrase(mnemonicStr, "", 0)
	//内存里清掉私钥防止泄露
	defer func() {
		ppk.Zero()
	}()
	privateKey := hex.EncodeToString(ppk.Serialize())
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "钱包创建成功", "data": gin.H{
		"private_key": privateKey,
		"public_key":  accountName,
	}})
}

func generateRandomName(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		result[i] = letters[num.Int64()]
	}
	return string(result)
}
