package task1

import (
  "context"
  "crypto/ecdsa"
  "go-eth-learn/internal/utils"
  "log"
  "math/big"

  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/crypto"
)

func Transfer(privateKeyHex string, targetAddress string, value *big.Int) string {
  client := utils.GetClient()
  privateKey, err := crypto.HexToECDSA(privateKeyHex)
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

  gasLimit := uint64(21000) // in units
  gasPrice, err := client.SuggestGasPrice(context.Background())
  if err != nil {
    log.Fatal(err)
  }
  toAddress := common.HexToAddress(targetAddress)
  tx := types.NewTx(&types.LegacyTx{
    Nonce:    nonce,
    To:       &toAddress,
    Value:    value,
    Gas:      gasLimit,
    GasPrice: gasPrice,
    Data:     nil,
  })
  chainID, err := client.NetworkID(context.Background())
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
  return signedTx.Hash().Hex()
}
