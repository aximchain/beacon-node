package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bnb-chain/node/app"
	"github.com/bnb-chain/node/app/config"
	axcInit "github.com/bnb-chain/node/cmd/axcchaind/init"
	"github.com/bnb-chain/node/common/utils"
)

var (
	chainID = "devnet-1000"
	nodeNum = 4
)

func main() {
	fmt.Println("start generate devnet configs")
	cwd, _ := os.Getwd()
	devnetHomeDir := path.Join(cwd, "build", "devnet")
	fmt.Println("devnet home dir:", devnetHomeDir)
	// clear devnetHomeDir
	err := os.RemoveAll(devnetHomeDir)
	if err != nil {
		panic(err)
	}
	// init nodes
	cdc := app.Codec
	ctx := app.ServerContext
	appInit := app.AximchainAppInit()
	ctxConfig := ctx.Config
	sdkConfig := sdk.GetConfig()
	sdkConfig.SetBech32PrefixForAccount(ctx.Bech32PrefixAccAddr, ctx.Bech32PrefixAccPub)
	sdkConfig.SetBech32PrefixForValidator(ctx.Bech32PrefixValAddr, ctx.Bech32PrefixValPub)
	sdkConfig.SetBech32PrefixForConsensusNode(ctx.Bech32PrefixConsAddr, ctx.Bech32PrefixConsPub)
	sdkConfig.Seal()
	var appState json.RawMessage
	var seeds string
	genesisTime := utils.Now()
	ServerContext := config.NewDefaultContext()
	nodesTemplateParams := make([]NodeTemplateParams, nodeNum)

	for i := 0; i < nodeNum; i++ {
		nodeName := fmt.Sprintf("node%d", i)
		nodeDir := path.Join(devnetHomeDir, nodeName, "testnoded")
		cliDir := path.Join(devnetHomeDir, nodeName, "testnodecli")
		ctxConfig.SetRoot(nodeDir)
		for _, subdir := range []string{"data", "config"} {
			err = os.MkdirAll(path.Join(nodeDir, subdir), os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
		// app.toml
		aximchainConfig := ServerContext.AximchainConfig
		aximchainConfig.UpgradeConfig.BEP3Height = 1
		aximchainConfig.UpgradeConfig.BEP8Height = 1
		aximchainConfig.UpgradeConfig.BEP12Height = 1
		aximchainConfig.UpgradeConfig.BEP67Height = 1
		aximchainConfig.UpgradeConfig.BEP70Height = 1
		aximchainConfig.UpgradeConfig.BEP82Height = 1
		aximchainConfig.UpgradeConfig.BEP84Height = 1
		aximchainConfig.UpgradeConfig.BEP87Height = 1
		aximchainConfig.UpgradeConfig.FixFailAckPackageHeight = 1
		aximchainConfig.UpgradeConfig.EnableAccountScriptsForCrossChainTransferHeight = 1
		aximchainConfig.UpgradeConfig.BEP128Height = 1
		aximchainConfig.UpgradeConfig.BEP151Height = 1
		aximchainConfig.UpgradeConfig.BEP153Height = 2
		aximchainConfig.UpgradeConfig.BEP159Height = 3
		aximchainConfig.UpgradeConfig.BEP159Phase2Height = 6
		aximchainConfig.BreatheBlockInterval = 5
		appConfigFilePath := filepath.Join(ctxConfig.RootDir, "config", "app.toml")
		config.WriteConfigFile(appConfigFilePath, aximchainConfig)
		// pk
		nodeID, pubKey := axcInit.InitializeNodeValidatorFiles(ctxConfig)
		ctxConfig.Moniker = nodeName
		valOperAddr, secret := axcInit.CreateValOperAccount(cliDir, ctxConfig.Moniker)
		fmt.Printf("%v secret: %v\n", nodeName, secret)
		if i == 0 {
			memo := fmt.Sprintf("%s@%s:26656", nodeID, "127.0.0.1")
			genTx := axcInit.PrepareCreateValidatorTx(cdc, chainID, ctxConfig.Moniker, memo, valOperAddr, pubKey)
			appState, err = appInit.AppGenState(cdc, []json.RawMessage{genTx})
			if err != nil {
				panic(err)
			}
			seeds = fmt.Sprintf("%s@172.20.0.100:26656", nodeID)
		} else {
			ctxConfig.P2P.Seeds = seeds
		}
		genFile := ctxConfig.GenesisFile()
		// genesis.json
		err = axcInit.ExportGenesisFileWithTime(genFile, chainID, nil, appState, genesisTime)
		if err != nil {
			panic(err)
		}
		// edit ctxConfig
		ctxConfig.LogLevel = "*:debug"
		// config.toml
		axcInit.WriteConfigFile(ctxConfig)
		// docker_compose.yml params
		node := NodeTemplateParams{Index: i, PortIP: i + 100, PortExpose1: 8000 + i, PortExpose2: 8100 + i}
		nodesTemplateParams[i] = node
	}
	dockerComposeTemplateParams := DockerComposeTemplateParams{
		Nodes: nodesTemplateParams,
	}
	WriteConfigFile(filepath.Join(devnetHomeDir, "docker-compose.yml"), &dockerComposeTemplateParams)
}
