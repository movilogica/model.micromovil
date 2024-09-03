package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"
)

type RequestPayload struct {
	Action   string         `json:"action"`
	Table    string         `json:"table,omitempty"`
	Data     any            `json:"data,omitempty"`
	Filter   any            `json:"filter,omitempty"`
	Metrica  MetricaPayload `json:"metrica,omitempty"`
	Auth     AuthPayload    `json:"auth,omitempty"`
	Log      LogPayLoad     `json:"log,omitempty"`
	Mail     MailPayload    `json:"mail,omitempty"`
	Uniqueid string         `json:"uniqueid,omitempty"`
	Uuid     string         `json:"uuid,omitempty"`
	Version  string         `json:"version,omitempty"`
	Redis    int            `json:"redis,omitempty"`
	Type     string         `json:"type,omitempty"`
}

type ResponsePayload struct {
	Error   bool   `json:"error"`
	Errorid int    `json:"errorid,omitempty"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type MetricaPayload struct {
	Username      string  `json:"username,omitempty"`
	Remotemachine string  `json:"remotemachine,omitempty"`
	Remotehost    string  `json:"remotehost,omitempty"`
	Remoteport    int32   `json:"remoteport,omitempty"`
	Headerdata    string  `json:"headerdata,omitempty"`
	Imei          string  `json:"imei,omitempty"`
	Cellinfo      string  `json:"cellinfo,omitempty"`
	Lat           float64 `json:"lat,omitempty"`
	Lon           float64 `json:"lon,omitempty"`
	CountryCode   string  `json:"countryCode,omitempty"`
	Region        string  `json:"region,omitempty"`
	RegionName    string  `json:"regionName,omitempty"`
	City          string  `json:"city,omitempty"`
	Zip           string  `json:"zip,omitempty"`
	TimeZone      string  `json:"timezone,omitempty"`
	Isp           string  `json:"isp,omitempty"`
	Org           string  `json:"org,omitempty"`
	As            string  `json:"as,omitempty"`
	Query         string  `json:"query,omitempty"`
}

func (m *MetricaPayload) IsAndroid() bool {
	return strings.Contains(strings.ToLower(m.Headerdata), "android")
}

func (m *MetricaPayload) IsIPhone() bool {
	return strings.Contains(strings.ToLower(m.Headerdata), "ios")
}

func (m *MetricaPayload) IsMobile() bool {
	return m.IsAndroid() || m.IsIPhone()
}

type AuthPayload struct {
	Token         string `json:"token,omitempty"`
	Imei          string `json:"imei,omitempty"`
	GoogleIdToken string `json:"googleidtoken,omitempty"`
	Nickname      string `json:"nickname,omitempty"`
	Username      string `json:"username,omitempty"`
	Email         string `json:"email,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Password      string `json:"password,omitempty"`
	Otp           string `json:"otp,omitempty"`
	Teclado       string `json:"teclado,omitempty"`
}

type LogPayLoad struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// ---- SessionData no debe de estar en el lado cliente, siempre debe estar en el Memcache ----
type SessionData struct {
	Id        int32  `json:"table,omitempty"` // Uniqueid de session data
	Name      string `json:"name,omitempty"`
	Uuid      string `json:"uuid,omitempty"`      // Uso interno y utilizado como campo clave especificado desde el cliente.
	TokenAuth string `json:"tokenauth,omitempty"` // Token de sesion data
	//	cookies   bool      `json:"cookies,omitempty"`
	ExpiryAt    time.Time            `json:"expiry,omitempty"`
	Language    string               `json:"language,omitempty"`
	Theme       string               `json:"theme,omitempty"` // Uso global
	SedeId      int32                `json:"sedeid,omitempty"`
	SedeText    string               `json:"sedetext,omitempty"`
	PersonaId   int64                `json:"personaid,omitempty"`
	Nroperacion string               `json:"nroperacion,omitempty"`
	TokendataId string               `json:"tokendataid,omitempty"`
	Location    DataComercioMinimalE `json:"location,omitempty"`
}

func (s *SessionData) IsExpired() bool {
	return s.ExpiryAt.Before(time.Now())
}

func MarshalJSON_Not_Nulls(e any) ([]byte, error) {
	v := reflect.ValueOf(e)

	result := make(map[string]interface{})

	for i := 0; i < v.NumField(); i++ {
		fieldName := strings.ToLower(v.Type().Field(i).Name)
		value := v.Field(i).Interface()
		typeValue := reflect.TypeOf(value)
		if typeValue == reflect.TypeOf(NullBool{}) {
			if !(value.(NullBool)).Valid {
				continue
			}
			value = (value.(NullBool)).Bool
		} else if typeValue == reflect.TypeOf(NullFloat64{}) {
			if !(value.(NullFloat64)).Valid {
				continue
			}
			value = (value.(NullFloat64)).Float64
		} else if typeValue == reflect.TypeOf(NullInt32{}) {
			if !(value.(NullInt32)).Valid {
				continue
			}
			value = (value.(NullInt32)).Int32
		} else if typeValue == reflect.TypeOf(NullInt64{}) {
			if !(value.(NullInt64)).Valid {
				continue
			}
			value = (value.(NullInt64)).Int64
		} else if typeValue == reflect.TypeOf(NullString{}) {
			if !(value.(NullString)).Valid {
				continue
			}
			value = (value.(NullString)).String
		} else if typeValue == reflect.TypeOf(NullTime{}) {
			if !(value.(NullTime)).Valid {
				continue
			}
			value = (value.(NullTime)).Time
		}
		result[fieldName] = value
	}

	return json.Marshal(result)
}

// CUSTOM NULL Handling structures
// NullInt64 is an alias for sql.NullInt64 data type
type NullInt32 sql.NullInt32

func (ns *NullInt32) Value(valor int32) {
	ns.Int32 = valor
	ns.Valid = true
}

// Scan implements the Scanner interface for NullInt64
func (ni *NullInt32) Scan(value interface{}) error {
	var i sql.NullInt32
	if err := i.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullInt32{i.Int32, false}
	} else {
		*ni = NullInt32{i.Int32, true}
	}
	return nil
}

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 sql.NullInt64

func (ns *NullInt64) Value(valor int64) {
	ns.Int64 = valor
	ns.Valid = true
}

// Scan implements the Scanner interface for NullInt64
func (ni *NullInt64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullInt64{i.Int64, false}
	} else {
		*ni = NullInt64{i.Int64, true}
	}
	return nil
}

// NullBool is an alias for sql.NullBool data type
type NullBool sql.NullBool

func (ns *NullBool) Value(valor bool) {
	ns.Bool = valor
	ns.Valid = true
}

// Scan implements the Scanner interface for NullBool
func (nb *NullBool) Scan(value interface{}) error {
	var b sql.NullBool
	if err := b.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nb = NullBool{b.Bool, false}
	} else {
		*nb = NullBool{b.Bool, true}
	}

	return nil
}

// NullFloat64 is an alias for sql.NullFloat64 data type
type NullFloat64 sql.NullFloat64

func (ns *NullFloat64) Value(valor float64) {
	ns.Float64 = valor
	ns.Valid = true
}

// Scan implements the Scanner interface for NullFloat64
func (nf *NullFloat64) Scan(value interface{}) error {
	var f sql.NullFloat64
	if err := f.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nf = NullFloat64{f.Float64, false}
	} else {
		*nf = NullFloat64{f.Float64, true}
	}

	return nil
}

// NullString is an alias for sql.NullString data type
type NullString sql.NullString

func (ns *NullString) Value(valor string) {
	ns.String = valor
	ns.Valid = true
}

// Scan implements the Scanner interface for NullString
func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = NullString{s.String, false}
	} else {
		*ns = NullString{s.String, true}
	}

	return nil
}

// NullTime is an alias for mysql.NullTime data type
type NullTime sql.NullTime

func (ns *NullTime) Value(valor time.Time) {
	ns.Time = valor
	ns.Valid = true
}

// Scan implements the Scanner interface for NullTime
func (nt *NullTime) Scan(value interface{}) error {
	var t sql.NullTime
	if err := t.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nt = NullTime{t.Time, false}
	} else {
		*nt = NullTime{t.Time, true}
	}

	return nil
}

// MarshalJSON for NullInt32
func (ni *NullInt32) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int32)
}

// UnmarshalJSON for NullInt32
func (ni *NullInt32) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int32)
	ni.Valid = (err == nil)
	return err
}

// MarshalJSON for NullInt64
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON for NullInt64
func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int64)
	ni.Valid = (err == nil)
	return err
}

// MarshalJSON for NullBool
func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nb.Bool)
}

// UnmarshalJSON for NullBool
func (nb *NullBool) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nb.Bool)
	nb.Valid = (err == nil)
	return err
}

// MarshalJSON for NullFloat64
func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}

// UnmarshalJSON for NullFloat64
func (nf *NullFloat64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nf.Float64)
	nf.Valid = (err == nil)
	return err
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}

// MarshalJSON for NullTime
func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339Nano))
	return []byte(val), nil
}

// UnmarshalJSON for NullTime
func (nt *NullTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	s := string(b)
	s = strings.Replace(s, "\"", "", 2)
	log.Println("NullTime - UnmarshalJSON =" + s)
	x, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		nt.Valid = false
		return err
	}

	nt.Time = x
	nt.Valid = true
	return nil
}
