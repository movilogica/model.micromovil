package model

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Almacenado cifrado en la maquina local - Info clave UUID (Autogenerado) - Sede actual y Apps disponibles
type CookieInfo struct {
	Uuid     string `json:"uuid,omitempty"`
	SedeId   int    `json:"sede"`
	SedeText string `json:"sede_text,omitempty"`
	SedeLogo string `json:"sede_logo,omitempty"`
	SedeApps string `json:"sede_apps,omitempty"`
}

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

func (m *MetricaPayload) GetHashMetricas() string {
	sum := sha256.Sum256([]byte(m.GetKeyMetricas()))
	return fmt.Sprintf("%x", sum)
}
func (m *MetricaPayload) GetKeyMetricas() string {
	return (m.Headerdata + m.Remotemachine)
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

type DataExchange struct {
	Uuid        string `json:"uuid,omitempty"`
	Topic       string `json:"topic,omitempty"`
	Action      string `json:"action,omitempty"`
	Data        string `json:"data,omitempty"`
	HashMetrica string `json:"hashmetrica,omitempty"`
	SedeId      int    `json:"sede,omitempty"`
}

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// ---- SessionData no debe de estar en el lado cliente, siempre debe estar en el Memcache ----
type SessionData struct {
	Id          int                   `json:"table,omitempty"` // Uniqueid de session data
	Name        string                `json:"name,omitempty"`
	Uuid        string                `json:"uuid,omitempty"`      // Uso interno y utilizado como campo clave especificado desde el cliente.
	TokenAuth   string                `json:"tokenauth,omitempty"` // Token de sesion data
	ExpiryAt    time.Time             `json:"expiry,omitempty"`
	Language    string                `json:"language,omitempty"`
	Theme       string                `json:"theme,omitempty"` // Uso global
	App         string                `json:app,omitempty`     // crm-store-balance-microfinanzas-etc...
	SubMenu     string                `json:submenu,omitempty` // current group menu izquierda seleccionado (empty==global)
	SideNav     int                   `json:sidenav,omitempty` /// 1-expand 0-hide
	SedeId      int                   `json:"sedeid,omitempty"`
	SedeText    string                `json:"sedetext,omitempty"`
	SedeLogo    string                `json:"sedelogo,omitempty"`
	SedeApps    string                `json:"sedeapps,omitempty"`
	OwnerId     int64                 `json:"ownerid,omitempty"`
	PersonaId   int64                 `json:"personaid,omitempty"`
	Nroperacion string                `json:"nroperacion,omitempty"`
	TokendataId string                `json:"tokendataid,omitempty"`
	ComercioId  int64                 `json:"comercioid,omitempty"`
	TradeName   string                `json:"tradename,omitempty"`
	ShortName   string                `json:"shortname,omitempty"`
	CountryCode string                `json:"countrycode,omitempty"`
	Location    *DataComercioMinimalE `json:"location,omitempty"`
	Sedes       []*SedeE              `json:"sedes"`
	Roles       []*UsuariosRolesE     `json:"roles,omitempty"`      /// Roles de nivel de sistemas
	RolNegocio  string                `json:"rolnegocio,omitempty"` /// Rol a nivel de negocio 'GUEST', 'OPERATOR', 'MANAGER'
	HashMetrica string                `json:"hashmetrica,omitempty"`
}

func (s *SessionData) BuildSessionData(sesionE SesionDataE, personaE *DataPersonasE, sedeE *SedeE, hashmetrica string) {
	// --- SESION EN MEMORIA CACHE
	s.Id = int(sesionE.Uniqueid)
	if sesionE.TokenId.Valid {
		s.TokenAuth = sesionE.TokenId.String
	}
	if sesionE.TokendataId.Valid {
		s.TokendataId = sesionE.TokendataId.String
	}
	if sesionE.JsessionId.Valid {
		s.Uuid = sesionE.JsessionId.String
	}
	if sesionE.ExpiryAt.Valid {
		s.ExpiryAt = sesionE.ExpiryAt.Time
	}
	s.PersonaId = personaE.Uniqueid
	s.OwnerId = personaE.Uniqueid
	s.Name = personaE.Nickname.String
	s.Nroperacion = personaE.Nroperacion.String
	s.CountryCode = personaE.CountryIso2.String
	if s.CountryCode == "" {
		s.CountryCode = "PE"
	}
	s.SedeId = int(sedeE.Uniqueid)
	s.SedeText = sedeE.Descrip.String
	s.SedeLogo = sedeE.UrlImage.String
	s.SedeApps = sedeE.Ruf1.String
	s.SubMenu = MENU_HOME
	s.SideNav = 1 /// Menu de la Izquierda expandido
	s.HashMetrica = hashmetrica
}

func (s *SessionData) IsExpired() bool {
	return s.ExpiryAt.Before(time.Now())
}

func (s *SessionData) IsAppCRM() bool {
	return strings.EqualFold(s.App, APP_CRM)
}

func (s *SessionData) IsAppStore() bool {
	return strings.EqualFold(s.App, APP_STORE)
}

func (s *SessionData) IsAppBalance() bool {
	return strings.EqualFold(s.App, APP_BALANCE)
}

func (s *SessionData) IsAppMicrofinanzas() bool {
	return strings.EqualFold(s.App, APP_MICROFINANZAS)
}

func (s *SessionData) IsAppInventory() bool {
	return strings.EqualFold(s.App, APP_INVENTORY)
}

func (s *SessionData) IsAppPedidos() bool {
	return strings.EqualFold(s.App, APP_PEDIDOS)
}

func (s *SessionData) IsAppTorneos() bool {
	return strings.EqualFold(s.App, APP_TORNEOS)
}

func (s *SessionData) IsRoleGuest() bool {
	role := strings.ToLower(ROLE_GUEST)
	return strings.Contains(strings.ToLower(s.RolNegocio), role) || s.IsRoleOperator()
}

func (s *SessionData) IsRoleOperator() bool {
	role := strings.ToLower(ROLE_OPERATOR)
	return strings.Contains(strings.ToLower(s.RolNegocio), role) || s.IsRoleManager()
}

func (s *SessionData) IsRoleManager() bool {
	role := strings.ToLower(ROLE_MANAGER)
	return strings.Contains(strings.ToLower(s.RolNegocio), role) || s.IsRoleOwner()
}

func (s *SessionData) IsRoleOwner() bool {
	return s.OwnerId == s.PersonaId
}

func (s *SessionData) IsSoporte() bool {
	hasrol := false
	for _, x := range s.Roles {
		hasrol = x.Soporte.Int32 == 1 || (x.Superadmin.Int32 == 1 && x.Activo == 1)
		if hasrol {
			break
		}
	}
	return hasrol
}
func (s *SessionData) IsDataCenter() bool {
	hasrol := false
	for _, x := range s.Roles {
		hasrol = x.Datacenter.Int32 == 1 || (x.Superadmin.Int32 == 1 && x.Activo == 1)
		if hasrol {
			break
		}
	}
	return hasrol
}
func (s *SessionData) IsManager() bool {
	hasrol := false
	for _, x := range s.Roles {
		hasrol = (x.Gerencial.Int32 == 1 && x.Activo == 1)
		if hasrol {
			break
		}
	}
	return hasrol
}
func (s *SessionData) IsSuperAdmin() bool {
	hasrol := false
	for _, x := range s.Roles {
		hasrol = (x.Superadmin.Int32 == 1 && x.Activo == 1)
		if hasrol {
			break
		}
	}
	return hasrol
}

// Actualizar en el futuro para considerar si la URL tiene el permiso
// Por ahora solo considera los tipos de ROL, tanto el rolemanager como el de Sistema
func (s *SessionData) HasPermission(url string) bool {
	/// Determinar las URL que no necesitan -> Location
	/*urlPublics := []string{"/datapers",
		"/datapers.do",
		"/sede",
		"/credencial",
		"/credencial.do",
		"/credencial/roles/",
		"/credencial.do/roles/",
		"/roles",
		"/roles.do",
		"/roles",
		"/paises",
		"/paises.do",
		"/regiones",
		"/regiones.do",
		"/departamentos",
		"/departamentos.do",
		"/provincias",
		"/provincias.do",
		"/distritos",
		"/distritos.do",
		"/zipcodes",
		"/zipcodes.do",
		"/param/variable",
		"/param.do/variable",
	}*/
	isPublic := false
	if (!isPublic && (s.Location == nil || s.Location.Uniqueid == 0)) ||
		!(s.IsSuperAdmin() || s.IsManager() ||
			s.IsRoleOperator() || s.IsRoleManager()) {
		log.Printf("global.SessionData:HasPersmission(%s) denied!\n", url)
		return false
	}
	log.Printf("global.SessionData:HasPersmission(%s) granted!\n", url)
	return true
}

func (s *SessionData) SetComercio(comercioE *DataComercioMinimalE) {
	s.TokendataId = comercioE.TokenDataId.String
	if comercioE.PersonaId.Valid {
		s.PersonaId = comercioE.PersonaId.Int64 /// nuevo valor asignado
	}
	s.Location = comercioE
	s.TradeName = comercioE.TradeName.String
	s.ShortName = comercioE.ShortName.String
	s.ComercioId = comercioE.Uniqueid
	s.RolNegocio = comercioE.RolNegocio.String
	/*if comercioE.RolNegocio.Valid {
		s.RolNegocio = comercioE.RolNegocio.String
	} else {
		s.RolNegocio = ""
	}*/
	/// Debemos seleccionar el primer App
	aplicacion := strings.Split(s.SedeApps, ",")
	if len(aplicacion) > 0 {
		s.App = aplicacion[0]
	}
}

func SetValue(obj any, field string, value any) {
	ref := reflect.ValueOf(obj)

	// if its a pointer, resolve its value
	if ref.Kind() == reflect.Ptr {
		ref = reflect.Indirect(ref)
	}

	if ref.Kind() == reflect.Interface {
		ref = ref.Elem()
	}

	// should double check we now have a struct (could still be anything)
	if ref.Kind() != reflect.Struct {
		log.Fatal("unexpected type")
	}

	prop := ref.FieldByName(field)
	prop.Set(reflect.ValueOf(value))
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
	if err != nil {
		var s string
		if err = json.Unmarshal(b, &s); err == nil {
			fmt.Printf("UnmarshalJSON=%s\n", s)
			x, err := strconv.ParseInt(s, 10, 32)
			if err != nil {
				ni.Int32 = int32(x)
			}
		}
	}
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
	if err != nil {
		var s string
		if err = json.Unmarshal(b, &s); err == nil {
			ni.Int64, err = strconv.ParseInt(s, 10, 64)
		}
	}
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
	//log.Println("NullTime - UnmarshalJSON =" + s)
	x, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		customLayout := "2006-01-02"
		x, err = time.Parse(customLayout, s)
		if err != nil {
			nt.Valid = false
			return err
		}
	}

	nt.Time = x
	nt.Valid = true
	return nil
}
