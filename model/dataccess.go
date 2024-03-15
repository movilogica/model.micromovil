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

const dbTimeout = time.Second * 3

var db *sql.DB

// New is the function used to create an instance of the data package. It returns the type
// Model, which embeds all the types we want to be available to our application.
func NewDataAccess(dbPool *sql.DB) DataAccess {
	db = dbPool

	return DataAccess{
		Parametros:              ParametersE{},
		Sede:                    SedeE{},
		Usuario:                 UsuarioE{},
		Credencial:              CredencialE{},
		Opciones:                OpcionesE{},
		Roles:                   RolesE{},
		RolesOpciones:           RolesOpcionesE{},
		RolesUsuarios:           RolesUsuariosE{},
		UsuariosRoles:           UsuariosRolesE{},
		UsuariosSedes:           UsuariosSedesE{},
		SesionData:              SesionDataE{},
		Ciiu:                    CiiuE{},
		Divisa:                  DivisaE{},
		Pais:                    PaisE{},
		Region:                  RegionE{},
		TipoDocumento:           TipoDocumentoE{},
		ZipCode:                 ZipCodeE{},
		Banco:                   BancosE{},
		CardBrand:               CardBrandE{},
		Dutie:                   DutieE{},
		DutieScope:              DutieScopeE{},
		Giros:                   GirosE{},
		Holiday:                 HolidayE{},
		TasaCambio:              TasaCambioE{},
		DataPersonas:            DataPersonasE{},
		DataPersonaDetail:       DataPersonaDetailE{},
		DataPersonaAddress:      DataPersonaAddressE{},
		DataPersonaTrusted:      DataPersonaTrustedE{},
		DataPersonasMedio:       DataPersonasMedioE{},
		DataPersonasRol:         DataPersonasRolE{},
		DataPersonasId:          DataPersonasIdE{},
		DataComercio:            DataComercioE{},
		DataComercioParam:       DataComercioParamE{},
		DataComercioPersonal:    DataComercioPersonalE{},
		DataComercioRole:        DataComercioRoleE{},
		DataTerminal:            DataTerminalE{},
		MovPersonas:             MovPersonasE{},
		MovTerminales:           MovTerminalesE{},
		BizPersonas:             BizPersonasE{},
		BizPersonasMedio:        BizPersonasMedioE{},
		BizPersonasVehicles:     BizPersonasVehiclesE{},
		BizNotesText:            BizNotesTextE{},
		BizParamInvoiceType:     BizParamInvoiceTypeE{},
		BizParamInvoiceItemType: BizParamInvoiceItemTypeE{},
		BizInvoice:              BizInvoiceE{},
		BizInvoiceDet:           BizInvoiceDetE{},
		BizInvoiceStatus:        BizInvoiceStatusE{},
		BizPayment:              BizPaymentE{},
		BizPaymentStatus:        BizPaymentStatusE{},
		BizKardex:               BizKardexE{},
		BizSubscriptions:        BizSubscriptionsE{},
		BizSubscriptionsPay:     BizSubscriptionsPayE{},
		StoreProduct:            StoreProductE{},
		StoreParamCategItem:     StoreParamCategItemE{},
		CustomerInfo:            CustomerInfoE{},
	}
}

// Models is the type for this package. Note that any model that is included as a member
// in this type is available to us throughout the application, anywhere that the
// app variable is used, provided that the model is also added in the New function.
type DataAccess struct {
	Parametros              ParametersE
	Sede                    SedeE
	Usuario                 UsuarioE
	Credencial              CredencialE
	Opciones                OpcionesE
	Roles                   RolesE
	RolesOpciones           RolesOpcionesE
	RolesUsuarios           RolesUsuariosE
	UsuariosRoles           UsuariosRolesE
	UsuariosSedes           UsuariosSedesE
	SesionData              SesionDataE
	Ciiu                    CiiuE
	Divisa                  DivisaE
	Pais                    PaisE
	Region                  RegionE
	TipoDocumento           TipoDocumentoE
	ZipCode                 ZipCodeE
	Banco                   BancosE
	CardBrand               CardBrandE
	Dutie                   DutieE
	DutieScope              DutieScopeE
	Giros                   GirosE
	Holiday                 HolidayE
	TasaCambio              TasaCambioE
	DataPersonas            DataPersonasE
	DataPersonaDetail       DataPersonaDetailE
	DataPersonaAddress      DataPersonaAddressE
	DataPersonaTrusted      DataPersonaTrustedE
	DataPersonasMedio       DataPersonasMedioE
	DataPersonasRol         DataPersonasRolE
	DataPersonasId          DataPersonasIdE
	DataComercio            DataComercioE
	DataComercioParam       DataComercioParamE
	DataComercioPersonal    DataComercioPersonalE
	DataComercioRole        DataComercioRoleE
	DataTerminal            DataTerminalE
	MovPersonas             MovPersonasE
	MovTerminales           MovTerminalesE
	BizPersonas             BizPersonasE
	BizPersonasMedio        BizPersonasMedioE
	BizPersonasVehicles     BizPersonasVehiclesE
	BizNotesText            BizNotesTextE
	BizParamInvoiceType     BizParamInvoiceTypeE
	BizParamInvoiceItemType BizParamInvoiceItemTypeE
	BizInvoice              BizInvoiceE
	BizInvoiceDet           BizInvoiceDetE
	BizInvoiceStatus        BizInvoiceStatusE
	BizPayment              BizPaymentE
	BizPaymentStatus        BizPaymentStatusE
	BizKardex               BizKardexE
	BizSubscriptions        BizSubscriptionsE
	BizSubscriptionsStatus  BizSubscriptionStatusE
	BizSubscriptionsPay     BizSubscriptionsPayE
	StoreProduct            StoreProductE
	StoreParamCategItem     StoreParamCategItemE
	CustomerInfo            CustomerInfoE
	ResumenGlobal           ResumenGlobalE
}

// CUSTOM NULL Handling structures
// NullInt64 is an alias for sql.NullInt64 data type
type NullInt32 sql.NullInt32

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
