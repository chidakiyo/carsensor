package carsensor_api_go

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

// used car search

type UsedCarQuery struct {
	// apiキー
	Key string `json:"key"`
	// 物件ID
	ID string `json:"id"`
	// ブランドコード
	Brand string `json:"brand"`
	// 車種名
	Model string `json:"model"`
	// 国コード
	Country string `json:"country"`
	// 大エリアコード
	LargeArea string `json:"large_area"`
	// 都道府県コード
	Pref string `json:"pref"`
	// ボディタイプコード
	Body string `json:"body"`
	// 定員
	Person string `json:"person"`
	// カラーコード
	Color string `json:"color"`
	// 最低価格
	PriceMin uint64 `json:"price_min"`
	// 最高価格
	PriceMax uint64 `json:"price_max"`
	//キーワード
	Keyword string `json:"keyword"`
	// 緯度
	Lat float64 `json:"lat"`
	// 経度
	Lng float64 `json:"lng"`
	// 検索範囲
	Range int64 `json:"range"`
	// 測地系
	Datum string `json:"datum"`
	// ミッション種類
	Mission int64 `json:"mission"`
	// 禁煙車
	Nonsmoking int64 `json:"nonsmoking"`
	// 本革シート
	Leather int64 `json:"leather"`
	// 福祉車両
	Welfare int64 `json:"welfare"`
	// 登録年式(古い)
	YearOld int64 `json:"year_old"`
	// 登録年式(新しい)
	YearNew int64 `json:"year_new"`
	// 走行距離(最短)
	OddMin int64 `json:"odd_min"`
	// 走行距離(最長)
	OddMax int64 `json:"odd_max"`
	// ソート順
	Order int64 `json:"order"`
	// 検索の開始位置
	Start int64 `json:"start"`
	// 1ページあたりの取得数
	Count int64 `json:"count"`
}

type UsedCarResponse struct {
	Base
	UsedCars []UsedCar `json:"usedcar"`
}

type UsedCar struct {
	// 物件ID
	ID string
	// ブランド
	Brand Brand
	// 車種名
	Model string
	// グレード名
	Grade string
	// 本体価格（応談があるのでstring型）
	Price int64
	// 車検
	Inspection string
	// 整備
	Maintenance string
	// 保証
	Warranty string
	// リサイクル
	Recycle string
	// エンジン型式
	Engine string
	// 説明文
	Desc string
	// ボディタイプ
	Body BodyType
	// 走行距離(不明があるのでstring)
	Odd string
	// 登録年
	Year string
	// 取扱店
	Shop Shop
	// カラー
	Color string
	// 整備に関するコメント
	MaintenanceComment string
	// 整備にかかる金額
	MaintenanceFee int64
	// 詳細ページURL
	URLS URLS
	// 保証に関するコメント
	WarrantyComment string
	// 保証距離
	WarrantyDistance string
	// 保証期間
	WarrantyLength string
	// 保証費用(円)
	WarrantyFee int64
	// 画像
	Photo Photo
}

type Shop struct {
	// 店名
	Name string
	// 都道府県名
	Perf Perf
	// 緯度
	Lat string
	// 軽度
	Lng string
	// 測地系
	Datum string
}

type Perf struct {
	// 都道府県コード
	Code int64
	// 都道府県名
	Name string
}

type Photo struct {
	// メイン
	Main PhotoMain `json:"main"`
	// サブ
	Sub []string `json:"sub"`
}

type PhotoMain struct {
	// メイン画像url(大)
	L string
	// メイン画像url(小)
	S string
	// キャプション
	Caption string
}

type PhotoAll struct {
	Front  PhotoMain `json:"front"`
	Rear   PhotoMain `json:"rear"`
	Inpane PhotoMain `json:"inpane"`
}

type URLS struct {
	// PC向けURL
	PC string
	// 携帯向けURL
	Mobile string
	// QRコード画像へのURL
	QR string
}

func SearchUsedCar(param UsedCarQuery) (UsedCarResponse, error) {

	u, err := CreateURL(URL_USEDCAR, param)
	if err != nil {
		err := errors.Wrap(err, "create url fail.")
		return UsedCarResponse{}, err // fail
	}

	fmt.Printf("URL : %s\n", u.String())

	resp, err := http.Get(u.String())
	defer resp.Body.Close()
	if err != nil {
		err := errors.Wrap(err, "http get fail.")
		return UsedCarResponse{}, err // fail
	}

	body, err := ioutil.ReadAll(resp.Body)

	response := &struct {
		Results UsedCarResponse `json:"results"`
	}{}
	if err := json.Unmarshal(body, response); err != nil {
		err := errors.Wrap(err, "json unmarshal fail.")
		return UsedCarResponse{}, err // fail
	}

	return response.Results, nil // success
}
