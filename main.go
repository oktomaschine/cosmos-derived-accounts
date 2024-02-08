package main

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
)

const (
	MNEMONIC = "measure slogan connect luggage stereo federal stuff stomach stumble security end differ"
	COUNT    = 50
)

func main() {
	seed := bip39.NewSeed(MNEMONIC, "")

	master, ch := hd.ComputeMastersFromSeed(seed)

	fmt.Println("path,address,private key")

	for i := 0; i < COUNT; i++ {
		path := hd.CreateHDPath(types.CoinType, 0, uint32(i))
		priv, _ := hd.DerivePrivateKeyForPath(master, ch, path.String())

		cdc := getCodec()
		kr, _ := keyring.New("App name", "memory", "./", nil, cdc)

		kr.ImportPrivKeyHex(path.String(), hex.EncodeToString(priv), string(hd.Secp256k1.Name()))
		kr.ExportPrivKeyArmor(path.String(), "")

		acc, _ := kr.Key(path.String())
		addr, _ := acc.GetAddress()

		fmt.Printf("%s,%s,%s\n", path, addr.String(), hex.EncodeToString(priv))
	}
}

func getCodec() codec.Codec {
	registry := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(registry)
	return codec.NewProtoCodec(registry)
}
