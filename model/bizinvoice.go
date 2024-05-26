package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
)

// Customer Medios
type BizInvoiceE struct {
	Uniqueid         int64                `json:"uniqueid,omitempty"`
	Owner            NullInt32            `json:"owner,omitempty"`
	Dispositivoid    NullInt32            `json:"dispositivoid,omitempty"`
	Id               int32                `json:"id,omitempty"`
	Sede             int32                `json:"sede"`
	Flag1            string               `json:"flag1,omitempty"`
	Flag2            string               `json:"flag2,omitempty"`
	PersonaId        NullInt64            `json:"personaid,omitempty"`
	TokendataId      NullString           `json:"tokendataid,omitempty"`
	BizPersonaId     NullInt64            `json:"bizpersonaid,omitempty"`
	BizPersonaText   NullString           `json:"bizpersonatext,omitempty"`
	BizPersonaRuc    NullString           `json:"bizpersonaruc,omitempty"`
	Email            NullString           `json:"email,omitempty"`
	Movil            NullString           `json:"movil,omitempty"`
	FormaPago        NullString           `json:"formapago,omitempty"`
	Fecha            NullTime             `json:"fecha,omitempty"`
	Fvenc            NullTime             `json:"fvenc,omitempty"`
	Fpago            NullTime             `json:"fpago,omitempty"`
	InvoiceTypeId    NullString           `json:"invoicetypeid,omitempty"`
	SubscriptionId   NullInt64            `json:"subscriptionid,omitempty"`
	StatusId         NullString           `json:"statusid,omitempty"`
	Numero           NullString           `json:"numero,omitempty"`
	NumeroText       NullString           `json:"numerotext,omitempty"`
	Nroperacion      NullString           `json:"nroperacion,omitempty"`
	Maskoperacion    NullString           `json:"maskperacion,omitempty"`
	DivisaId         NullInt64            `json:"divisaid,omitempty"`
	DivisaText       NullString           `json:"divisatext,omitempty"`
	DivisaSimbolo    NullString           `json:"divisasimbolo,omitempty"`
	DivisaDecimal    NullInt32            `json:"divisadecimal,omitempty"`
	TasaVenta        NullFloat64          `json:"tasaventa,omitempty"`
	TasaCompra       NullFloat64          `json:"tasacompra,omitempty"`
	TotalItems       NullFloat64          `json:"totalitems,omitempty"`
	TotalBruto       NullFloat64          `json:"totalbruto,omitempty"`
	TotalDctos       NullFloat64          `json:"totaldctos,omitempty"`
	TotalNeto        NullFloat64          `json:"totalneto,omitempty"`
	TotalImponible   NullFloat64          `json:"totalimponible,omitempty"`
	TotalImpuestos   NullFloat64          `json:"totalimpuestos,omitempty"`
	TotalRetenciones NullFloat64          `json:"totalretenciones,omitempty"`
	TotalFinal       NullFloat64          `json:"totalfinal,omitempty"`
	TotalPago        NullFloat64          `json:"totalpago,omitempty"`
	TotalDocAplic    NullFloat64          `json:"totaldocaplic,omitempty"`
	Comentarios      NullString           `json:"comentarios,omitempty"`
	Latitud          NullFloat64          `json:"latitud,omitempty"`
	Longitud         NullFloat64          `json:"longitud,omitempty"`
	Aplicar          NullString           `json:"aplicar,omitempty"`
	Ruf1             NullString           `json:"ruf1,omitempty"`
	Ruf2             NullString           `json:"ruf2,omitempty"`
	Ruf3             NullString           `json:"ruf3,omitempty"`
	Iv               NullString           `json:"iv,omitempty"`
	Salt             NullString           `json:"salt,omitempty"`
	Checksum         NullString           `json:"checksum,omitempty"`
	FCreated         NullTime             `json:"fcreated,omitempty"`
	FUpdated         NullTime             `json:"fupdated,omitempty"`
	UCreated         NullString           `json:"ucreated,omitempty"`
	UUpdated         NullString           `json:"uupdated,omitempty"`
	Activo           int32                `json:"activo,omitempty"`
	Estadoreg        int32                `json:"estadoreg,omitempty"`
	TotalRecords     int64                `json:"total_records"`
	Items            []*BizInvoiceDetE    `json:"items:omitempty"`
	Status           []*BizInvoiceStatusE `json:"status:omitempty"`
}

type BizInvoiceDetE struct {
	Uniqueid            int64       `json:"uniqueid,omitempty"`
	Owner               NullInt32   `json:"owner,omitempty"`
	Dispositivoid       int32       `json:"dispositivoid,omitempty"`
	Id                  int32       `json:"id,omitempty"`
	Sede                int32       `json:"sede"`
	Flag1               string      `json:"flag1,omitempty"`
	Flag2               string      `json:"flag2,omitempty"`
	PersonaId           NullInt64   `json:"personaid,omitempty"`
	TokendataId         NullString  `json:"tokendataid,omitempty"`
	BizPersonaId        NullInt64   `json:"bizpersonaid,omitempty"`
	InvoiceId           NullInt64   `json:"invoiceid,omitempty"`
	Nroperacion         NullString  `json:"nroperacion,omitempty"`
	ProductId           NullInt64   `json:"productid,omitempty"`
	ProductText         NullString  `json:"producttext,omitempty"`
	Secuencial          NullInt32   `json:"secuencial,omitempty"`
	InvoiceItemTypeId   NullString  `json:"invoiceitemtypeid,omitempty"`
	OverrideGlAccountId NullString  `json:"overrideglaccountid,omitempty"`
	PrecioVenta         NullFloat64 `json:"precioventa,omitempty"`
	Cantidad            NullInt64   `json:"cantidad,omitempty"`
	UndVenta            NullString  `json:"undventa,omitempty"`
	Bruto               NullFloat64 `json:"bruto,omitempty"`
	DctoPorc            NullFloat64 `json:"dctoporc,omitempty"`
	DctoMonto           NullFloat64 `json:"dctomonto,omitempty"`
	Descuentos          NullFloat64 `json:"descuentos,omitempty"`
	Imponble            NullFloat64 `json:"imponible,omitempty"`
	Impuestos           NullFloat64 `json:"impuestos,omitempty"`
	Retenciones         NullFloat64 `json:"retenciones,omitempty"`
	Total               NullFloat64 `json:"total,omitempty"`
	Comentarios         NullString  `json:"comentarios,omitempty"`
	Ruf1                NullString  `json:"ruf1,omitempty"`
	Ruf2                NullString  `json:"ruf2,omitempty"`
	Ruf3                NullString  `json:"ruf3,omitempty"`
	Iv                  NullString  `json:"iv,omitempty"`
	Salt                NullString  `json:"salt,omitempty"`
	Checksum            NullString  `json:"checksum,omitempty"`
	FCreated            NullTime    `json:"fcreated,omitempty"`
	FUpdated            NullTime    `json:"fupdated,omitempty"`
	Activo              int32       `json:"activo,omitempty"`
	Estadoreg           int32       `json:"estadoreg,omitempty"`
	TotalRecords        int64       `json:"total_records"`
}

type BizInvoiceStatusE struct {
	Uniqueid      int64      `json:"uniqueid,omitempty"`
	Owner         NullInt32  `json:"owner,omitempty"`
	Dispositivoid int32      `json:"dispositivoid,omitempty"`
	Id            int32      `json:"id,omitempty"`
	Sede          int32      `json:"sede"`
	Flag1         string     `json:"flag1,omitempty"`
	Flag2         string     `json:"flag2,omitempty"`
	PersonaId     NullInt64  `json:"personaid,omitempty"`
	TokendataId   NullString `json:"tokendataid,omitempty"`
	BizPersonaId  NullInt64  `json:"bizpersonaid,omitempty"`
	InvoiceId     NullString `json:"invoiceid,omitempty"`
	StatusId      NullString `json:"statusid,omitempty"`
	StatusDetail  NullString `json:"statusdetail,omitempty"`
	Ruf1          NullString `json:"ruf1,omitempty"`
	Ruf2          NullString `json:"ruf2,omitempty"`
	Ruf3          NullString `json:"ruf3,omitempty"`
	Iv            NullString `json:"iv,omitempty"`
	Salt          NullString `json:"salt,omitempty"`
	Checksum      NullString `json:"checksum,omitempty"`
	FCreated      NullTime   `json:"fcreated,omitempty"`
	FUpdated      NullTime   `json:"fupdated,omitempty"`
	Activo        int32      `json:"activo,omitempty"`
	Estadoreg     int32      `json:"estadoreg,omitempty"`
	TotalRecords  int64      `json:"total_records"`
}

func (e BizInvoiceE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}
func (e BizInvoiceDetE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}
func (e BizInvoiceStatusE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

func (e BizInvoiceE) Pendiente() float64 {
	return e.TotalFinal.Float64 - e.TotalPago.Float64
}

func (c BizInvoiceE) NexOrderId() int {
	max := 1
	for _, v := range c.Items {
		if int(v.Secuencial.Int32) >= max {
			max = int(v.Secuencial.Int32) + 1
		}
	}
	return max
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func (i *BizInvoiceDetE) SubTotal() {
	fmt.Println("--- Calculando Item ---")
	i.Total.Float64 = float64(roundFloat(float64(float32(i.Cantidad.Int64)*float32(i.PrecioVenta.Float64)), 2))
	fmt.Printf("Precio:%f - Qty:%d - SubTotal:%f\n", i.PrecioVenta.Float64, i.Cantidad.Int64, i.Total.Float64)
}

func (c *BizInvoiceE) Totales() {
	fmt.Println("--- Calculando Totales ---")
	tot_items := 0.0
	tot_qty := 0
	var tot_sub float64
	var tot_final float64
	for _, v := range c.Items {
		tot_items += 1
		tot_qty += int(v.Cantidad.Int64)
		v.SubTotal()
		tot_sub += v.Total.Float64
		tot_final += v.Total.Float64
	}
	c.TotalItems.Float64 = tot_items
	///c.TotalQuantity = tot_qty
	c.TotalImponible.Float64 = tot_sub
	taxPerc := float64(c.TasaVenta.Float64) / float64(100)
	c.TotalImpuestos.Float64 = float64(roundFloat(float64(taxPerc*tot_sub), 2))
	c.TotalFinal.Float64 = float64(roundFloat(float64(c.TotalImponible.Float64+c.TotalImpuestos.Float64), 2))
	fmt.Printf("Items:%d Total:%d\n", c.TotalItems.Float64, c.TotalFinal.Float64)
}

const querySelectBizInvoiceCab = `select * from biz_invoice_cab_list( $1, $2)`
const querySelectBizInvoiceDet = `select * from biz_invoice_det_list( $1, $2)`
const querySelectBizInvoiceStatus = `select * from biz_invoice_status_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *BizInvoiceE) GetAll(token string, filter string) ([]*BizInvoiceE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBizInvoiceCab

	// Se deseenvuelve el JSON del Filter para adicionar filtros
	var mapFilter map[string]interface{}
	json.Unmarshal([]byte(filter), &mapFilter)
	if mapFilter == nil {
		mapFilter = make(map[string]interface{})
	}
	// --- Adicion de filtros
	// mapFilter["tipo"] = tabla
	// Se empaqueta el JSON del Filter
	jsonFilter, err := json.Marshal(mapFilter)
	if err != nil {
		log.Println("Error convirtiendo el Filter")
	}
	log.Println("Where = " + string(jsonFilter))

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx, token, string(jsonFilter))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []*BizInvoiceE

	for rows.Next() {
		var rowdata BizInvoiceE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Owner,
			&rowdata.Dispositivoid,
			&rowdata.Id,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.PersonaId,
			&rowdata.TokendataId,
			&rowdata.BizPersonaId,
			&rowdata.BizPersonaText,
			&rowdata.BizPersonaRuc,
			&rowdata.Email,
			&rowdata.Movil,
			&rowdata.FormaPago,
			&rowdata.Fecha,
			&rowdata.Fvenc,
			&rowdata.Fpago,
			&rowdata.InvoiceTypeId,
			&rowdata.SubscriptionId,
			&rowdata.StatusId,
			&rowdata.Numero,
			&rowdata.NumeroText,
			&rowdata.Nroperacion,
			&rowdata.Maskoperacion,
			&rowdata.DivisaId,
			&rowdata.DivisaText,
			&rowdata.DivisaSimbolo,
			&rowdata.DivisaDecimal,
			&rowdata.TasaVenta,
			&rowdata.TasaCompra,
			&rowdata.TotalItems,
			&rowdata.TotalBruto,
			&rowdata.TotalDctos,
			&rowdata.TotalNeto,
			&rowdata.TotalImponible,
			&rowdata.TotalImpuestos,
			&rowdata.TotalRetenciones,
			&rowdata.TotalFinal,
			&rowdata.TotalPago,
			&rowdata.TotalDocAplic,
			&rowdata.Comentarios,
			&rowdata.Latitud,
			&rowdata.Longitud,
			&rowdata.Aplicar,
			&rowdata.Ruf1,
			&rowdata.Ruf2,
			&rowdata.Ruf3,
			&rowdata.Iv,
			&rowdata.Salt,
			&rowdata.Checksum,
			&rowdata.FCreated,
			&rowdata.FUpdated,
			&rowdata.UCreated,
			&rowdata.UUpdated,
			&rowdata.Activo,
			&rowdata.Estadoreg,
			&rowdata.TotalRecords,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		lista = append(lista, &rowdata)
	}

	return lista, nil
}

// GetOne returns one user by id
func (u *BizInvoiceE) GetByUniqueid(token string, uniqueid int) (*BizInvoiceE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBizInvoiceCab

	var rowdata BizInvoiceE
	jsonText := fmt.Sprintf(`{"uniqueid":%d}`, uniqueid)
	row := db.QueryRowContext(ctx, query, token, jsonText)

	err := row.Scan(
		&rowdata.Uniqueid,
		&rowdata.Owner,
		&rowdata.Dispositivoid,
		&rowdata.Id,
		&rowdata.Sede,
		&rowdata.Flag1,
		&rowdata.Flag2,
		&rowdata.PersonaId,
		&rowdata.TokendataId,
		&rowdata.BizPersonaId,
		&rowdata.BizPersonaText,
		&rowdata.BizPersonaRuc,
		&rowdata.Email,
		&rowdata.Movil,
		&rowdata.FormaPago,
		&rowdata.Fecha,
		&rowdata.Fvenc,
		&rowdata.Fpago,
		&rowdata.InvoiceTypeId,
		&rowdata.SubscriptionId,
		&rowdata.StatusId,
		&rowdata.Numero,
		&rowdata.NumeroText,
		&rowdata.Nroperacion,
		&rowdata.Maskoperacion,
		&rowdata.DivisaId,
		&rowdata.DivisaText,
		&rowdata.DivisaSimbolo,
		&rowdata.DivisaDecimal,
		&rowdata.TasaVenta,
		&rowdata.TasaCompra,
		&rowdata.TotalItems,
		&rowdata.TotalBruto,
		&rowdata.TotalDctos,
		&rowdata.TotalNeto,
		&rowdata.TotalImponible,
		&rowdata.TotalImpuestos,
		&rowdata.TotalRetenciones,
		&rowdata.TotalFinal,
		&rowdata.TotalPago,
		&rowdata.TotalDocAplic,
		&rowdata.Comentarios,
		&rowdata.Latitud,
		&rowdata.Longitud,
		&rowdata.Aplicar,
		&rowdata.Ruf1,
		&rowdata.Ruf2,
		&rowdata.Ruf3,
		&rowdata.Iv,
		&rowdata.Salt,
		&rowdata.Checksum,
		&rowdata.FCreated,
		&rowdata.FUpdated,
		&rowdata.UCreated,
		&rowdata.UUpdated,
		&rowdata.Activo,
		&rowdata.Estadoreg,
		&rowdata.TotalRecords,
	)

	if err != nil {
		return nil, err
	}

	jsonText = fmt.Sprintf(`{"invoiceid":%d}`, uniqueid)

	/// Se adicionan los items del invoice
	itemE := BizInvoiceDetE{}
	detalle, err := itemE.GetAll(token, jsonText)
	if err == nil {
		rowdata.Items = detalle
	}
	/// Se adicionan los status del invoice
	statusE := BizInvoiceStatusE{}
	status, err := statusE.GetAll(token, jsonText)
	if err == nil {
		rowdata.Status = status
	}

	return &rowdata, nil
}

func (u *BizInvoiceDetE) GetAll(token string, filter string) ([]*BizInvoiceDetE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBizInvoiceDet

	// Se deseenvuelve el JSON del Filter para adicionar filtros
	var mapFilter map[string]interface{}
	json.Unmarshal([]byte(filter), &mapFilter)
	if mapFilter == nil {
		mapFilter = make(map[string]interface{})
	}
	// --- Adicion de filtros
	// mapFilter["tipo"] = tabla
	// Se empaqueta el JSON del Filter
	jsonFilter, err := json.Marshal(mapFilter)
	if err != nil {
		log.Println("Error convirtiendo el Filter")
	}
	log.Println("Where = " + string(jsonFilter))

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx, token, string(jsonFilter))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []*BizInvoiceDetE

	for rows.Next() {
		var rowdata BizInvoiceDetE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Owner,
			&rowdata.Dispositivoid,
			&rowdata.Id,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.PersonaId,
			&rowdata.TokendataId,
			&rowdata.BizPersonaId,
			&rowdata.InvoiceId,
			&rowdata.Nroperacion,
			&rowdata.ProductId,
			&rowdata.ProductText,
			&rowdata.Secuencial,
			&rowdata.InvoiceItemTypeId,
			&rowdata.OverrideGlAccountId,
			&rowdata.PrecioVenta,
			&rowdata.Cantidad,
			&rowdata.UndVenta,
			&rowdata.Bruto,
			&rowdata.DctoPorc,
			&rowdata.DctoMonto,
			&rowdata.Descuentos,
			&rowdata.Imponble,
			&rowdata.Impuestos,
			&rowdata.Retenciones,
			&rowdata.Total,
			&rowdata.Comentarios,
			&rowdata.Ruf1,
			&rowdata.Ruf2,
			&rowdata.Ruf3,
			&rowdata.Iv,
			&rowdata.Salt,
			&rowdata.Checksum,
			&rowdata.FCreated,
			&rowdata.FUpdated,
			&rowdata.Activo,
			&rowdata.Estadoreg,
			&rowdata.TotalRecords,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		lista = append(lista, &rowdata)
	}

	return lista, nil
}

func (u *BizInvoiceDetE) GetByUniqueid(token string, uniqueid int) (*BizInvoiceDetE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBizInvoiceDet

	var rowdata BizInvoiceDetE
	jsonText := fmt.Sprintf(`{"uniqueid":%d}`, uniqueid)
	row := db.QueryRowContext(ctx, query, token, jsonText)

	err := row.Scan(
		&rowdata.Uniqueid,
		&rowdata.Owner,
		&rowdata.Dispositivoid,
		&rowdata.Id,
		&rowdata.Sede,
		&rowdata.Flag1,
		&rowdata.Flag2,
		&rowdata.PersonaId,
		&rowdata.TokendataId,
		&rowdata.BizPersonaId,
		&rowdata.InvoiceId,
		&rowdata.Nroperacion,
		&rowdata.ProductId,
		&rowdata.ProductText,
		&rowdata.Secuencial,
		&rowdata.InvoiceItemTypeId,
		&rowdata.OverrideGlAccountId,
		&rowdata.PrecioVenta,
		&rowdata.Cantidad,
		&rowdata.UndVenta,
		&rowdata.Bruto,
		&rowdata.DctoPorc,
		&rowdata.DctoMonto,
		&rowdata.Descuentos,
		&rowdata.Imponble,
		&rowdata.Impuestos,
		&rowdata.Retenciones,
		&rowdata.Total,
		&rowdata.Comentarios,
		&rowdata.Ruf1,
		&rowdata.Ruf2,
		&rowdata.Ruf3,
		&rowdata.Iv,
		&rowdata.Salt,
		&rowdata.Checksum,
		&rowdata.FCreated,
		&rowdata.FUpdated,
		&rowdata.Activo,
		&rowdata.Estadoreg,
		&rowdata.TotalRecords,
	)

	if err != nil {
		return nil, err
	}

	/// Se adicionan los items del invoice

	return &rowdata, nil
}

func (u *BizInvoiceStatusE) GetAll(token string, filter string) ([]*BizInvoiceStatusE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBizInvoiceStatus

	// Se deseenvuelve el JSON del Filter para adicionar filtros
	var mapFilter map[string]interface{}
	json.Unmarshal([]byte(filter), &mapFilter)
	if mapFilter == nil {
		mapFilter = make(map[string]interface{})
	}
	// --- Adicion de filtros
	// mapFilter["tipo"] = tabla
	// Se empaqueta el JSON del Filter
	jsonFilter, err := json.Marshal(mapFilter)
	if err != nil {
		log.Println("Error convirtiendo el Filter")
	}
	log.Println("Where = " + string(jsonFilter))

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx, token, string(jsonFilter))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []*BizInvoiceStatusE

	for rows.Next() {
		var rowdata BizInvoiceStatusE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Owner,
			&rowdata.Dispositivoid,
			&rowdata.Id,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.PersonaId,
			&rowdata.TokendataId,
			&rowdata.BizPersonaId,
			&rowdata.InvoiceId,
			&rowdata.StatusId,
			&rowdata.StatusDetail,
			&rowdata.Ruf1,
			&rowdata.Ruf2,
			&rowdata.Ruf3,
			&rowdata.Iv,
			&rowdata.Salt,
			&rowdata.Checksum,
			&rowdata.FCreated,
			&rowdata.FUpdated,
			&rowdata.Activo,
			&rowdata.Estadoreg,
			&rowdata.TotalRecords,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		lista = append(lista, &rowdata)
	}

	return lista, nil
}

func (u *BizInvoiceStatusE) GetByUniqueid(token string, uniqueid int) (*BizInvoiceStatusE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBizInvoiceStatus

	var rowdata BizInvoiceStatusE
	jsonText := fmt.Sprintf(`{"uniqueid":%d}`, uniqueid)
	row := db.QueryRowContext(ctx, query, token, jsonText)

	err := row.Scan(
		&rowdata.Uniqueid,
		&rowdata.Owner,
		&rowdata.Dispositivoid,
		&rowdata.Id,
		&rowdata.Sede,
		&rowdata.Flag1,
		&rowdata.Flag2,
		&rowdata.PersonaId,
		&rowdata.TokendataId,
		&rowdata.BizPersonaId,
		&rowdata.InvoiceId,
		&rowdata.StatusId,
		&rowdata.StatusDetail,
		&rowdata.Ruf1,
		&rowdata.Ruf2,
		&rowdata.Ruf3,
		&rowdata.Iv,
		&rowdata.Salt,
		&rowdata.Checksum,
		&rowdata.FCreated,
		&rowdata.FUpdated,
		&rowdata.Activo,
		&rowdata.Estadoreg,
		&rowdata.TotalRecords,
	)

	if err != nil {
		return nil, err
	}

	/// Se adicionan los items del invoice

	return &rowdata, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *BizInvoiceE) Update(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Se deseenvuelve el JSON del Data para adicionar filtros
	var mapData map[string]interface{}
	json.Unmarshal([]byte(data), &mapData)
	if mapData == nil {
		mapData = make(map[string]interface{})
	}
	// --- Adicion de filtro de tipos de carros
	// mapData["tipo"] = tabla

	// Se empaqueta el JSON del Data
	jsonData, err := json.Marshal(mapData)
	if err != nil {
		log.Println("Error convirtiendo el Dato")
		return nil, err
	}
	log.Println("Data = " + string(jsonData))

	query := `SELECT biz_invoice_cab_save($1, $2, $3)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	result, err := stmt.QueryContext(ctx, token, string(jsonData), metricas)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var uniqueid int64

	if result.Next() {
		err := result.Scan(&uniqueid)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
	}

	retorno := make(map[string]any)
	retorno["uniqueid"] = uniqueid

	return retorno, nil
}

func (u *BizInvoiceDetE) Update(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Se deseenvuelve el JSON del Data para adicionar filtros
	var mapData map[string]interface{}
	json.Unmarshal([]byte(data), &mapData)
	if mapData == nil {
		mapData = make(map[string]interface{})
	}
	// --- Adicion de filtro de tipos de carros
	// mapData["tipo"] = tabla

	// Se empaqueta el JSON del Data
	jsonData, err := json.Marshal(mapData)
	if err != nil {
		log.Println("Error convirtiendo el Dato")
		return nil, err
	}
	log.Println("Data = " + string(jsonData))

	query := `SELECT biz_invoice_det_save($1, $2, $3)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	result, err := stmt.QueryContext(ctx, token, string(jsonData), metricas)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var uniqueid int64

	if result.Next() {
		err := result.Scan(&uniqueid)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
	}

	retorno := make(map[string]any)
	retorno["uniqueid"] = uniqueid

	return retorno, nil
}

func (u *BizInvoiceStatusE) Update(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Se deseenvuelve el JSON del Data para adicionar filtros
	var mapData map[string]interface{}
	json.Unmarshal([]byte(data), &mapData)
	if mapData == nil {
		mapData = make(map[string]interface{})
	}
	// --- Adicion de filtro de tipos de carros
	// mapData["tipo"] = tabla

	// Se empaqueta el JSON del Data
	jsonData, err := json.Marshal(mapData)
	if err != nil {
		log.Println("Error convirtiendo el Dato")
		return nil, err
	}
	log.Println("Data = " + string(jsonData))

	query := `SELECT biz_invoice_status_save($1, $2, $3)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	result, err := stmt.QueryContext(ctx, token, string(jsonData), metricas)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var uniqueid int64

	if result.Next() {
		err := result.Scan(&uniqueid)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
	}

	retorno := make(map[string]any)
	retorno["uniqueid"] = uniqueid

	return retorno, nil
}

// Delete deletes one user from the database, by User.ID
func (u *BizInvoiceE) Delete(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := `SELECT biz_invoice_cab_save($1, $2, $3)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	result, err := stmt.QueryContext(ctx, token, jsonText, metricas)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var uniqueid int64

	if result.Next() {
		err := result.Scan(&uniqueid)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
	}

	retorno := make(map[string]any)
	retorno["uniqueid"] = uniqueid

	return retorno, nil
}

// DeleteByID deletes one user from the database, by ID
func (u *BizInvoiceDetE) Delete(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := `SELECT biz_invoice_det_save($1, $2, $3)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	result, err := stmt.QueryContext(ctx, token, jsonText, metricas)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var uniqueid int64

	if result.Next() {
		err := result.Scan(&uniqueid)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
	}

	retorno := make(map[string]any)
	retorno["uniqueid"] = uniqueid

	return retorno, nil
}

// DeleteByID deletes one user from the database, by ID
func (u *BizInvoiceStatusE) Delete(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := `SELECT biz_invoice_status_save($1, $2, $3)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	result, err := stmt.QueryContext(ctx, token, jsonText, metricas)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var uniqueid int64

	if result.Next() {
		err := result.Scan(&uniqueid)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
	}

	retorno := make(map[string]any)
	retorno["uniqueid"] = uniqueid

	return retorno, nil
}

/*****
 * Este procedimiento registra datos en multiples tablas transaccionales.-
 *********/
func (u *BizInvoiceE) RegisterBill(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Printf("RegisterBill [data]=%s [metricas]=%s\n", data, metricas)

	query := `CALL generate_bill_ticket($1, $2, $3, $4)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var v_uniqueid float64
	_, err = stmt.ExecContext(ctx, token, data, metricas, &v_uniqueid)
	if err != nil {
		return nil, err
	}
	///defer result.Close()

	retorno := make(map[string]any)
	retorno["uniqueid"] = v_uniqueid

	log.Printf("RegisterBill [ID]=%v\n", v_uniqueid)

	return retorno, nil
}

/*****
 * Este procedimiento registra datos en multiples tablas transaccionales.-
 *********/
func (u *BizInvoiceE) AnularBill(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Printf("AnularBill [data]=%s [metricas]=%s\n", data, metricas)

	query := `CALL anular_bill_ticket($1, $2, $3, $4)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var v_uniqueid float64
	_, err = stmt.ExecContext(ctx, token, data, metricas, &v_uniqueid)
	if err != nil {
		return nil, err
	}
	///defer result.Close()

	retorno := make(map[string]any)
	retorno["uniqueid"] = v_uniqueid

	log.Printf("AnularBill [ID]=%v\n", v_uniqueid)

	return retorno, nil
}
