package model

import (
	"encoding/json"
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

type AuthPayload struct {
	Token    string `json:"token,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
	Otp      string `json:"otp,omitempty"`
	Teclado  string `json:"teclado,omitempty"`
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
