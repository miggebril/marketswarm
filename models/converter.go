package models

import (
	"database/sql"
	"marketswarm/lib/gorp"
	"encoding/json"
	"errors"
	"time"
)

type MarketswarmTypeConverter struct{}

func (me MarketswarmTypeConverter) ToDb(val interface{}) (interface{}, error) {

	//switch t := val.(type) {
		// case Broker:
		// 	b, err := json.Marshal(t)
		// 	if err != nil {
		// 		return "", err
		// 	}
		// 	return b, nil
		// case Counterparty:
		// 	b, err := json.Marshal(t)
		// 	if err != nil {
		// 		return "", err
		// 	}
		// 	return b, nil
		// case time.Time:
		// 	if t.IsZero() {
		// 		return nil, nil
		// 	}
		// 	return t, nil
		// case []string:
		// 	b, err := json.Marshal(t)
		// 	if err != nil {
		// 		return "", err
		// 	}
		// 	return b, nil
		// case []Fill:
		// 	b, err := json.Marshal(t)
		// 	if err != nil {
		// 		return "", err
		// 	}
		// 	return b, nil
		// case Portfolio:
		// 	b, err := json.Marshal(t)
		// 	if err != nil {
		// 		return "", err
		// 	}
		// 	return b, nil
		// case Fill:
		// 	b, err := json.Marshal(t)
		// 	if err != nil {
		// 		return "", err
		// 	}
		// 	return b, nil
	//}
	return val, nil
}


func (me MarketswarmTypeConverter) FromDb(target interface{}) (gorp.CustomScanner, bool) {
	switch target.(type) {
		case *int64:
			binder := func(holder, target interface{}) error {
				s, ok := holder.(*sql.NullInt64)
				if !ok {
					return errors.New("FromDb: Unable to convert int64 to *NullInt64")
				}
				sp, ok := target.(*int64)
				if !ok {
					return errors.New("FromDb: Unable to convert target to *int64")
				}
				if s.Valid {
					*sp = s.Int64
				} else {
					*sp = 0
				}
				return nil
			}
			return gorp.CustomScanner{new(sql.NullInt64), target, binder}, true
		case *int8:
			binder := func(holder, target interface{}) error {
				s, ok := holder.(*sql.NullInt64)
				if !ok {
					return errors.New("FromDb: Unable to convert int8 to *NullInt64")
				}
				sp, ok := target.(*int8)
				if !ok {
					return errors.New("FromDb: Unable to convert target to *int8")
				}
				if s.Valid {
					*sp = int8(s.Int64)
				} else {
					*sp = 0
				}
				return nil
			}
			return gorp.CustomScanner{new(sql.NullInt64), target, binder}, true
		case *int32:
			binder := func(holder, target interface{}) error {
				s, ok := holder.(*sql.NullInt64)
				if !ok {
					return errors.New("FromDb: Unable to convert int32 to *NullInt64")
				}
				sp, ok := target.(*int32)
				if !ok {
					return errors.New("FromDb: Unable to convert target to *int32")
				}
				if s.Valid {
					*sp = int32(s.Int64)
				} else {
					*sp = 0
				}
				return nil
			}
			return gorp.CustomScanner{new(sql.NullInt64), target, binder}, true
		case *int:
			binder := func(holder, target interface{}) error {
				s, ok := holder.(*sql.NullInt64)
				if !ok {
					return errors.New("FromDb: Unable to convert int to *NullInt64")
				}
				sp, ok := target.(*int)
				if !ok {
					return errors.New("FromDb: Unable to convert target to *int")
				}
				if s.Valid {
					*sp = int(s.Int64)
				} else {
					*sp = 0
				}
				return nil
			}
			return gorp.CustomScanner{new(sql.NullInt64), target, binder}, true
		case *bool:
			binder := func(holder, target interface{}) error {
				s, ok := holder.(*sql.NullBool)
				if !ok {
					return errors.New("FromDb: Unable to convert bool to *sql.NullBool")
				}
				bp, ok := target.(*bool)
				if !ok {
					return errors.New("FromDb: Unable to convert target to *bool")
				}
				if s.Valid && s.Bool {
					*bp = true
				}
				return nil
			}
			return gorp.CustomScanner{new(sql.NullBool), target, binder}, true
		case *string:
			binder := func(holder, target interface{}) error {
				s, ok := holder.(*sql.NullString)
				if !ok {
					return errors.New("FromDb: Unable to convert string to *NullString")
				}
				sp, ok := target.(*string)
				if !ok {
					return errors.New("FromDb: Unable to convert target to *string")
				}
				if s.Valid {
					*sp = s.String
				} else {
					*sp = ""
				}
				return nil
			}
			return gorp.CustomScanner{new(sql.NullString), target, binder}, true
		case *time.Time:
			binder := func(holder, target interface{}) error {
				s, ok := holder.(*gorp.NullTime)
				if !ok {
					return errors.New("FromDb: Unable to convert Time to *string")
				}
				sp, ok := target.(*time.Time)
				if !ok {
					return errors.New("FromDb: Unable to convert target to *time")
				}
				if (*s).Valid {
					*sp = (*s).Time.UTC()
				}
				return nil
			}
			return gorp.CustomScanner{new(gorp.NullTime), target, binder}, true
		case *[]string:
			binder := func(holder, target interface{}) error {
				s, ok := holder.(*[]byte)
				if !ok {
					return errors.New("FromDb: Unable to convert String to *string")
				}
				if string(*s) == "" || string(*s) == "null" {
					target = []string{}
					return nil
				}
				return json.Unmarshal(*s, target)
			}
			return gorp.CustomScanner{new([]byte), target, binder}, true
	}

	return gorp.CustomScanner{}, false
}