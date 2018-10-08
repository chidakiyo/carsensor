package carsensor_api_go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// used car search

type UsedCarQuery struct {
	// apiキー
	Key string
	// 物件ID
	ID string
	// ブランドコード
	Brand string
	// 車種名
	Model string
	// 国コード
	Country string
	// 大エリアコード
	LargeArea string
	// 都道府県コード
	Pref string
	// ボディタイプコード
	Body string
	// 定員
	Person string
	// カラーコード
	Color string
	// 最低価格
	PriceMin uint64
	// 最高価格
	PriceMax uint64
	//キーワード
	Keyword string
	// 緯度
	Lat float64
	// 経度
	Lng float64
	// 検索範囲
	Range int64
	// 測地系
	Datum string
	// ミッション種類
	Mission int64
	// 禁煙車
	Nonsmoking int64
	// 本革シート
	Leather int64
	// 福祉車両
	Welfare int64
	// 登録年式(古い)
	YearOld int64
	// 登録年式(新しい)
	YearNew int64
	// 走行距離(最短)
	OddMin int64
	// 走行距離(最長)
	OddMax int64
	// ソート順
	Order int64
	// 検索の開始位置
	Start int64
	// 1ページあたりの取得数
	Count int64
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

func SearchUsedCar(param UsedCarQuery) UsedCarResponse {

	// TODO code利用する実装

	// TODO large_area

	url := URL_USEDCAR + "?key=" + param.Key + "&format=" + FORMAT + "&model=ジュリア" + "&count=1"

	fmt.Printf("[URL] %s\n", url)

	resp, err := http.Get(url) // TODO 全件
	defer resp.Body.Close()
	if err != nil {
		// TODO handle error
	}

	body, err := ioutil.ReadAll(resp.Body)

	response := &struct {
		Results UsedCarResponse `json:"results"`
	}{}
	if err := json.Unmarshal(body, response); err != nil {
		panic(err)
	}

	fmt.Printf("[[ RESPONSE BODY ]]\n%+v\n", string(body))
	fmt.Printf("[[ RESPONSE MAPPING ]]\n%+v\n", response)

	fmt.Printf("[[ URL ]]\n%+v\n", response.Results.UsedCars[0].URLS)

	fmt.Printf("[[ BODY ]]\n%+v\n", response.Results.UsedCars[0].Body)

	fmt.Printf("[[ SUB_IMG ]]\n%+v\n", response.Results.UsedCars[0])

	return response.Results
}
