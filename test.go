package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"

	fswap "github.com/fox-one/4swap-sdk-go"
	"github.com/fox-one/4swap-sdk-go/mtg"
	mixin "github.com/fox-one/mixin-sdk-go"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

var (
	// Specify the keystore file in the -config parameterp
	pin    = flag.String("pin", "178625", "")
	config = flag.String("config", "", "keystore file path")
)

type Empty struct {
	Data []Datum `json:"data"`
}

type Datum struct {
	Type           string      `json:"type"`
	AssetID        string      `json:"asset_id"`
	ChainID        string      `json:"chain_id"`
	Symbol         string      `json:"symbol"`
	Name           string      `json:"name"`
	IconURL        string      `json:"icon_url"`
	Balance        string      `json:"balance"`
	DepositEntries interface{} `json:"deposit_entries"`
	Destination    string      `json:"destination"`
	Topic          string      `json:"Topic"`
	PriceBtc       string      `json:"price_btc"`
	PriceUsd       string      `json:"price_usd"`
	ChangeBtc      string      `json:"change_btc"`
	ChangeUsd      string      `json:"change_usd"`
	AssetKey       string      `json:"asset_key"`
	MixinID        string      `json:"mixin_id"`
	Reserve        string      `json:"reserve"`
	Confirmations  int64       `json:"confirmations"`
	Capitalization int64       `json:"capitalization"`
	Liquidity      string      `json:"liquidity"`
}

func getAssetsID(name string) string {
	var empty Empty
	URL := "https://api.mixin.one/network/assets/search/" + name

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	json.Unmarshal([]byte(body), &empty)
	return empty.Data[0].AssetID
}

func makeSwap(ctx context.Context, InputAssetID, OutputAssetID string, client *mixin.Client, InputAmount float64) error {
	// 用 4swap's MTG api 端点
	fswap.UseEndpoint(fswap.MtgEndpoint)
	// read the mtg group
	// the group information would change frequently
	// it's recommended to save it for later use
	group, err := fswap.ReadGroup(ctx)
	if err != nil {
		return err
	}

	fmt.Println(group.Members)

	pairs, _ := fswap.ListPairs(ctx)
	sort.Slice(pairs, func(i, j int) bool {
		aLiquidity := pairs[i].BaseValue.Add(pairs[i].QuoteValue)
		bLiquidity := pairs[j].BaseValue.Add(pairs[j].QuoteValue)
		return aLiquidity.GreaterThan(bLiquidity)
	})

	preOrder, err := fswap.Route(pairs, InputAssetID, OutputAssetID, decimal.NewFromFloat(InputAmount))
	if err != nil {
		return err
	}
	log.Printf("Route: %v", preOrder.RouteAssets)

	followID, _ := uuid.NewV4()
	action := mtg.SwapAction(
		client.ClientID,
		followID.String(),
		OutputAssetID,
		preOrder.Routes,
		decimal.NewFromFloat(0.00000001),
	)

	// 生成 memo
	memo, err := action.Encode(group.PublicKey)
	if err != nil {
		return err
	}
	log.Println("memo", memo)
	tx, err := client.Transaction(ctx, &mixin.TransferInput{
		AssetID: InputAssetID,
		Amount:  decimal.RequireFromString("1000"),
		TraceID: mixin.RandomTraceID(),
		Memo:    memo,
		OpponentMultisig: struct {
			Receivers []string `json:"receivers,omitempty"`
			Threshold uint8    `json:"threshold,omitempty"`
		}{
			Receivers: group.Members,
			Threshold: uint8(group.Threshold),
		},
	}, *pin)
	if err != nil {
		println("交易出错")
		return err
	}
	log.Println("tx", tx)
	return nil
}

func main() {
	flag.Parse()

	// Open the keystore file
	f, err := os.Open(*config)
	if err != nil {
		log.Panicln(err)
	}

	// Read the keystore file as json into mixin.Keystore, which is a go struct
	var store mixin.Keystore
	if err := json.NewDecoder(f).Decode(&store); err != nil {
		log.Panicln(err)
	}

	// Create a Mixin Client from the keystore, which is the instance to invoke Mixin APIs
	// client, err := mixin.NewFromKeystore(&store)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(store.Scope)
	// println(client.ClientID)
	ctx := context.Background()

	// snapshots, _ := client.ReadSnapshots(ctx, "", time.Now(), "", 1)
	// for _, snapshot := range snapshots {
	// 	fmt.Println("snapshot.Amount " + snapshot.Amount.String())
	// 	fmt.Println("snapshot.TraceID " + snapshot.TraceID)
	// 	fmt.Println(snapshot.CreatedAt)

	// }
	// fmt.Println(time.Now().UTC().Format("2022-09-30 08:23:26.58472 +0000 UTC"))
	assets, err := mixin.ReadTopNetworkAssets(ctx)
	if err != nil {
		log.Panicln(err)
		return
	}
	for _, asset := range assets {
		fmt.Println(asset.Symbol + " = " + asset.AssetID)
	}
}

// func swaptoMe(ctx context.Context, client *mixin.Client, InputAssetID string) {

// 	fswap.UseEndpoint(fswap.MtgEndpoint)
// 	group, err := fswap.ReadGroup(ctx)
// 	fmt.Println(group.Members)
// 	if err != nil {
// 		fmt.Println("read 4swap error")
// 		return
// 	}

// followID, _ := uuid.NewV4()
// action := mtg.SwapAction(
// 	"6a87e67f-02fb-47cf-b31f-32a13dd5b3d9",
// 	followID.String(),
// 	InputAssetID,
// 	"",
// 	decimal.NewFromFloat(0.00000001),
// )

// memo, err := action.Encode(group.PublicKey)
// if err != nil {
// 	return
// }
// tx, err := client.Transaction(ctx, &mixin.TransferInput{
// 	AssetID: InputAssetID,
// 	Amount:  decimal.RequireFromString("200"),
// 	TraceID: mixin.RandomTraceID(),
// 	Memo:    memo,
// 	OpponentMultisig: struct {
// 		Receivers []string `json:"receivers,omitempty"`
// 		Threshold uint8    `json:"threshold,omitempty"`
// 	}{
// 		Receivers: group.Members,
// 		Threshold: uint8(group.Threshold),
// 	},
// }, *pin)
// if err != nil {
// 	println("swap error")
// 	return
// }

// 	tx, err := client.Transfer(ctx, &mixin.TransferInput{
// 		AssetID:    InputAssetID,
// 		OpponentID: "6a87e67f-02fb-47cf-b31f-32a13dd5b3d9",
// 		Amount:     decimal.RequireFromString("200"),
// 		TraceID:    mixin.RandomTraceID(),
// 		Memo:       "refund",
// 	}, *pin)
// 	if err != nil {
// 		return
// 	}

// 	log.Println("tx", tx.SnapshotID)
// 	log.Println("tx SnapshotID", tx.SnapshotID)
// }

/*
* uid: 用户或机器人的 uuid
* sid: Session Id
* privateKey: 机器人私钥
* method: HTTP 请求方法 GET, POST
* url: HTTP 请求 URL 例如: /transfers
* body: HTTP 请求内容, 例如: {"pin": "encrypted pin token"}
 */
// func SignAuthenticationToken(uid, sid, privateKey, method, uri, body string) (string, error) {
// 	expire := time.Now().UTC().Add(time.Hour * 24 * 30 * 3)
// 	sum := sha256.Sum256([]byte(method + uri + body))
// 	jti, _ := uuid.NewV4()
// 	claims := jwt.MapClaims{
// 		"uid": uid,
// 		"sid": sid,
// 		"iat": time.Now().UTC().Unix(),
// 		"exp": expire.Unix(),
// 		"jti": jti,
// 		"sig": hex.EncodeToString(sum[:]),
// 		"scp": "FULL",
// 	}
// 	priv, err := base64.RawURLEncoding.DecodeString(privateKey)
// 	if err != nil {
// 		return "", err
// 	}
// 	// more validate the private key
// 	if len(priv) != 64 {
// 		return "", fmt.Errorf("Bad ed25519 private key %s", priv)
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
// 	return token.SignedString(ed25519.PrivateKey(priv))
// }
