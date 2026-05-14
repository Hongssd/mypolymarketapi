package mypolymarketapi

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"
)

// csv to json
func parseAccountingSnapshotZip(body []byte) (*DataDownloadAccountingSnapshotRes, error) {
	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return nil, err
	}

	res := &DataDownloadAccountingSnapshotRes{
		Positions: []DataAccountingSnapshotPositionRow{},
		Equities:  []DataAccountingSnapshotEquityRow{},
	}

	for _, f := range zipReader.File {
		if f.FileInfo().IsDir() {
			continue
		}
		if !strings.HasSuffix(strings.ToLower(f.Name), ".csv") {
			continue
		}

		reader, err := f.Open()
		if err != nil {
			return nil, err
		}

		var parseErr error
		switch strings.ToLower(f.Name) {
		case "positions.csv":
			res.Positions, parseErr = parsePositionCSVRows(reader)
		case "equity.csv":
			res.Equities, parseErr = parseEquityCSVRows(reader)
		default:
			// 未使用的 CSV 暂不输出，保持返回结构稳定
		}
		closeErr := reader.Close()
		if parseErr != nil {
			return nil, fmt.Errorf("parse csv file %s failed: %w", f.Name, parseErr)
		}
		if closeErr != nil {
			return nil, closeErr
		}
	}

	return res, nil
}

func parseEquityCSVRows(reader io.Reader) ([]DataAccountingSnapshotEquityRow, error) {
	csvReader := csv.NewReader(reader)
	csvReader.TrimLeadingSpace = true

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return []DataAccountingSnapshotEquityRow{}, nil
	}

	headers := records[0]
	rows := make([]DataAccountingSnapshotEquityRow, 0, len(records)-1)
	for i := 1; i < len(records); i++ {
		record := records[i]
		row := DataAccountingSnapshotEquityRow{}
		for j, header := range headers {
			val := ""
			if j < len(record) {
				val = record[j]
			}
			fillEquityRowByHeader(&row, header, val)
		}
		rows = append(rows, row)
	}

	return rows, nil
}

func parsePositionCSVRows(reader io.Reader) ([]DataAccountingSnapshotPositionRow, error) {
	csvReader := csv.NewReader(reader)
	csvReader.TrimLeadingSpace = true

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return []DataAccountingSnapshotPositionRow{}, nil
	}

	headers := records[0]
	rows := make([]DataAccountingSnapshotPositionRow, 0, len(records)-1)
	for i := 1; i < len(records); i++ {
		record := records[i]
		row := DataAccountingSnapshotPositionRow{}
		for j, header := range headers {
			val := ""
			if j < len(record) {
				val = record[j]
			}
			fillPositionRowByHeader(&row, header, val)
		}
		rows = append(rows, row)
	}

	return rows, nil
}

func fillEquityRowByHeader(row *DataAccountingSnapshotEquityRow, header, raw string) {
	switch strings.TrimSpace(header) {
	case "cashBalance":
		row.CashBalance = parseCSVFloat64(raw)
	case "positionsValue":
		row.PositionsValue = parseCSVFloat64(raw)
	case "equity":
		row.Equity = parseCSVFloat64(raw)
	case "valuationTime":
		val := strings.TrimSpace(raw)
		if val == "" {
			return
		}
		row.ValuationTime = val
		if t, err := time.Parse(time.RFC3339, val); err == nil {
			row.ValuationTimestamp = t.Unix()
		}
	default:
		return
	}
}

func fillPositionRowByHeader(row *DataAccountingSnapshotPositionRow, header, raw string) {
	switch strings.TrimSpace(header) {
	case "conditionId":
		row.ConditionID = parseCSVString(raw)
	case "asset":
		row.Asset = parseCSVString(raw)
	case "size":
		row.Size = parseCSVFloat64(raw)
	case "curPrice":
		row.CurPrice = parseCSVFloat64(raw)
	case "valuationTime":
		val := strings.TrimSpace(raw)
		if val == "" {
			return
		}
		row.ValuationTime = val
		if t, err := time.Parse(time.RFC3339, val); err == nil {
			row.ValuationTimestamp = t.Unix()
		}
	default:
		return
	}
}

func parseCSVFloat64(raw string) float64 {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return 0
	}
	if num, err := strconv.ParseFloat(trimmed, 64); err == nil {
		return num
	}
	return 0
}

func parseCSVString(raw string) string {
	return strings.TrimSpace(raw)
}

type MySyncMap[K any, V any] struct {
	smap sync.Map
}

func NewMySyncMap[K any, V any]() MySyncMap[K, V] {
	return MySyncMap[K, V]{
		smap: sync.Map{},
	}
}
func (m *MySyncMap[K, V]) Load(k K) (V, bool) {
	v, ok := m.smap.Load(k)

	if ok {
		return v.(V), true
	}
	var resv V
	return resv, false
}
func (m *MySyncMap[K, V]) Store(k K, v V) {
	m.smap.Store(k, v)
}

func (m *MySyncMap[K, V]) Delete(k K) {
	m.smap.Delete(k)
}
func (m *MySyncMap[K, V]) Range(f func(k K, v V) bool) {
	m.smap.Range(func(k, v any) bool {
		return f(k.(K), v.(V))
	})
}

func (m *MySyncMap[K, V]) Length() int {
	length := 0
	m.Range(func(k K, v V) bool {
		length += 1
		return true
	})
	return length
}

func (m *MySyncMap[K, V]) MapValues(f func(k K, v V) V) *MySyncMap[K, V] {
	var res = NewMySyncMap[K, V]()
	m.Range(func(k K, v V) bool {
		res.Store(k, f(k, v))
		return true
	})
	return &res
}

// ParseClobTokenIDs 解析 Gamma Market 的 clobTokenIds 字段（模型里为 string）。
// 常见为 JSON 数组字符串，但部分市场为空、JSON null、或管道占位（如 "||"），直接 json.Unmarshal 到 []string 会失败。
func ParseClobTokenIDs(raw string) ([]string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" || raw == "null" {
		return nil, nil
	}
	var ids []string
	if err := json.Unmarshal([]byte(raw), &ids); err == nil {
		return ids, nil
	}
	var inner string
	if err := json.Unmarshal([]byte(raw), &inner); err == nil {
		inner = strings.TrimSpace(inner)
		if inner == "" || inner == "null" {
			return nil, nil
		}
		if err := json.Unmarshal([]byte(inner), &ids); err == nil {
			return ids, nil
		}
	}
	if strings.Contains(raw, "|") {
		ids = nil
		for _, p := range strings.Split(raw, "|") {
			p = strings.TrimSpace(p)
			if p != "" {
				ids = append(ids, p)
			}
		}
		if len(ids) > 0 {
			return ids, nil
		}
		return nil, nil
	}
	return nil, fmt.Errorf("parse clobTokenIds: invalid form %q", raw)
}

func GetAssetIDsFromEventSlug(slug string) ([]string, error) {
	p := &MyPolymarket{}
	gammaClient := p.NewGammaRestClient()
	res, err := gammaClient.NewGammaGetEventBySlug().Slug(slug).Do()
	if err != nil {
		return nil, err
	}
	var assetIDs []string
	for _, market := range res.Data.Markets {
		log.Infof("market: %+v", market.Slug)
		tokenIDs, err := ParseClobTokenIDs(market.ClobTokenIds)
		if err != nil {
			return nil, err
		}
		assetIDs = append(assetIDs, tokenIDs...)
	}

	return assetIDs, nil
}

func GetAssetIDsFromMarketSlug(slug string) ([]string, error) {
	p := &MyPolymarket{}
	gammaClient := p.NewGammaRestClient()
	res, err := gammaClient.NewGammaGetMarketBySlug().Slug(slug).Do()
	if err != nil {
		return nil, err
	}
	return ParseClobTokenIDs(res.Data.ClobTokenIds)
}
