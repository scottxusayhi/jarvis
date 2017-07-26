package helper

import (
	"fmt"
	"encoding/json"
	"git.oschina.net/k2ops/jarvis/server/api/model"
)

func scanStringValue(src interface{}) (string, error) {
	fmt.Printf("got %T type\n", src)
	var byteArrayValue []byte
	//var strValue string
	var ok bool
	byteArrayValue, ok = src.([]byte)
	if !ok {
		return "", fmt.Errorf("expect string but got %T instead", src)
	}
	return string(byteArrayValue), nil
}

func scanByteArrayValue(src interface{}) ([]byte, error) {
	fmt.Printf("got %T type\n", src)
	var byteArrayValue []byte
	var ok bool
	byteArrayValue, ok = src.([]byte)
	if !ok {
		return nil, fmt.Errorf("expect string but got %T instead", src)
	}
	return byteArrayValue, nil
}

// scan array of objects to struct Host,
// fields with error is kept unchanged
func scanValuesToHost(src []interface{}, dest *model.Host) {
	var strValue string
	var byteArrayValue []byte
	var err error
	// datacenter
	if strValue, err = scanStringValue(src[0]); err==nil {
		dest.DataCenter = strValue
	}
	// rack
	if strValue, err = scanStringValue(src[1]); err==nil {
		dest.Rack = strValue
	}
	// slot
	if strValue, err = scanStringValue(src[2]); err==nil {
		dest.Slot = strValue
	}
	// tags
	fmt.Println(dest.Tags)
	if byteArrayValue, err = scanByteArrayValue(src[4]); err==nil {
		json.Unmarshal(byteArrayValue, &dest.Tags)
	}
}

//func (m *JarvisMysqlBackend) GetOneHost(dc string, rack string, slot string, hostname string) (*model.Host, error) {
//	var err error
//	host := &model.Host{}
//	rows, _ := m.stmtGetOneHost.Query(dc, rack, slot, hostname)
//	columns, _ := rows.Columns()
//	count := len(columns)
//	values := make([]interface{}, count)
//	valuePtrs := make([]interface{}, count)
//
//	for rows.Next() {
//		// init valuePtrs with pointer of values
//		for i := range columns {
//			valuePtrs[i] = &values[i]
//		}
//		// scan row to values
//		err = rows.Scan(valuePtrs...)
//		if err != nil {
//			log.Error(err.Error())
//			return nil, err
//		}
//		fmt.Println(host)
//		scanValuesToHost(values, host)
//		fmt.Println(host)
//		return host, nil
//	}
//	return nil, sql.ErrNoRows
//}
