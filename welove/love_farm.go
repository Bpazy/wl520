package welove

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func FarmSign(love Love) (*http.Response, error) {
	u := "http://api.welove520.com/v1/game/farm/signin"
	sigEncoder := NewSig([]byte(KEY))
	// TODO access_token=xxxx&sig=5d74e30439a6656c12aa6dd2f2a5cd85(md5 possible)&ts=1546446867977(current timestamp)
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)

	d1 := Data{"access_token", love.AccessToken}
	d3 := Data{"app_key", love.AppKey}
	d2 := Data{"ts", timestamp}
	sig := sigEncoder.Encode("POST", u, d1, d3, d2)

	data := make(url.Values)
	data.Add("access_token", love.AccessToken)
	data.Add("sig", sig)
	data.Add("ts", timestamp)

	return NewWlHttpClient().Post(u, data)
}

type QueryItem struct {
	Result   int `json:"result"`
	Messages []struct {
		OpTime  int64 `json:"op_time"`
		MsgType int   `json:"msg_type"`
		AdItems []struct {
			ItemID        int    `json:"item_id"`
			Count         int    `json:"count"`
			OpTime        int64  `json:"op_time"`
			NeedHelp      int    `json:"need_help"`
			SellerFarmID  string `json:"seller_farm_id"`
			HeadURLFamale string `json:"head_url_famale"`
			HeadURLMale   string `json:"head_url_male"`
			ID            int    `json:"id"`
			FarmName      string `json:"farm_name"`
			Coin          int    `json:"coin"`
		} `json:"ad_items"`
	} `json:"messages"`
}

func QueryItems(accessToken string) QueryItem {
	u := "http://api.welove520.com/v1/game/farm/ad/query"
	d1 := Data{"access_token", accessToken}
	sigEncoder := NewSig([]byte(KEY))
	sig := sigEncoder.Encode("POST", u, d1)
	data := make(url.Values)
	data.Add("access_token", accessToken)
	data.Add("sig", sig)
	res, err := http.PostForm(u, data)
	if err != nil {
		panic(err)
	}
	bytes, _ := ioutil.ReadAll(res.Body)
	queryItem := QueryItem{}
	json.Unmarshal(bytes, &queryItem)
	return queryItem
}

type BuyItemStatus struct {
	Result   int `json:"result"`
	Messages []struct {
		StallItem struct {
			BuyerHeadURL  string `json:"buyer_head_url"`
			BuyerFarmName string `json:"buyer_farm_name"`
			ID            int    `json:"id"`
		} `json:"stall_item,omitempty"`
		OpTime     int64 `json:"op_time"`
		MsgType    int   `json:"msg_type"`
		Warehouses []struct {
			Category int `json:"category"`
			ItemsInc []struct {
				ItemID int `json:"item_id"`
				Count  int `json:"count"`
			} `json:"items_inc"`
		} `json:"warehouses,omitempty"`
		FarmID   string `json:"farm_id,omitempty"`
		GoldCost int    `json:"gold_cost,omitempty"`
	} `json:"messages"`
}

func BuyItem(accessToken, sellerFarmId string, stallSaleId int) BuyItemStatus {
	u := "http://api.welove520.com/v1/game/farm/stall/buy"
	d1 := Data{"access_token", accessToken}
	d2 := Data{"seller_farm_id", sellerFarmId}
	d3 := Data{"stall_sale_id", strconv.Itoa(stallSaleId)}
	sigEncoder := NewSig([]byte(KEY))
	sig := sigEncoder.Encode("POST", u, d1, d2, d3)

	data := make(url.Values)
	data.Add("access_token", accessToken)
	data.Add("seller_farm_id", sellerFarmId)
	data.Add("stall_sale_id", strconv.Itoa(stallSaleId))
	data.Add("sig", sig)
	res, err := http.PostForm(u, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	buyItemStatus := BuyItemStatus{}
	json.Unmarshal(bytes, &buyItemStatus)
	return buyItemStatus
}
