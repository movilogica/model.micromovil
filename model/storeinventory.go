package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Inventory
type StoreInventoryE struct {
	Uniqueid            int64       `json:"uniqueid,omitempty"`
	Owner               NullInt32   `json:"owner,omitempty"`
	Dispositivoid       NullInt32   `json:"dispositivoid,omitempty"`
	Id                  int32       `json:"id,omitempty"`
	Sede                int32       `json:"sede"`
	Flag1               string      `json:"flag1,omitempty"`
	Flag2               string      `json:"flag2,omitempty"`
	PersonaId           NullInt64   `json:"personaid,omitempty"`
	TokendataId         NullString  `json:"tokendataid,omitempty"`
	WarehouseId         NullInt64   `json:"warehouseid,omitempty"`
	WarehouseText       NullString  `json:"warehousetext,omitempty"`
	LocationId          NullInt64   `json:"locationid,omitempty"`
	LocationSeqId       NullString  `json:"locationseqid,omitempty"`
	InventoryTypeId     NullString  `json:"inventorytypeid,omitempty"`
	ProductId           NullInt64   `json:"productid,omitempty"`
	ProductCode         NullString  `json:"productcode,omitempty"`
	ProductText         NullString  `json:"producttext,omitempty"`
	DivisionCode        NullString  `json:"divisioncode,omitempty"`
	VirtualInfo         NullString  `json:"virtualinfo,omitempty"`
	Inventario          NullString  `json:"inventario,omitempty"`
	ProductTypeId       NullString  `json:"producttypeid,omitempty"`
	StatusItemId        NullString  `json:"statusitemid,omitempty"`
	ReceivedDate        NullTime    `json:"receiveddate,omitempty"`
	ManufacturedDate    NullTime    `json:"manufactureddate,omitempty"`
	ExpiredDate         NullTime    `json:"expireddate,omitempty"`
	ContainerId         NullString  `json:"containerid,omitempty"`
	LotId               NullString  `json:"lotid,omitempty"`
	LotExpired          NullTime    `json:"lotexpired,omitempty"`
	SkuNumber           NullString  `json:"skunumber,omitempty"`
	BinNumber           NullString  `json:"binnumber,omitempty"`
	SerialNumber        NullString  `json:"serialnumber,omitempty"`
	SoftIdentifier      NullString  `json:"softidentifier,omitempty"`
	BarcodeBox          NullString  `json:"barcodebox,omitempty"`
	BarcodeItem         NullString  `json:"barcodeitem,omitempty"`
	DocumentText        NullString  `json:"documenttext,omitempty"`
	PersonaText         NullString  `json:"personatext,omitempty"`
	SequenceId          NullInt64   `json:"sequenceid,omitempty"`
	ActivationNumber    NullString  `json:"activationnumber,omitempty"`
	ActivationValidThru NullTime    `json:"activationvalidthru,omitempty"`
	StockMinimo         NullInt64   `json:"stockminimo,omitempty"`
	UDisplay            NullString  `json:"udisplay,omitempty"`
	Uom                 NullString  `json:"uom,omitempty"`
	Quom                NullFloat64 `json:"quom,omitempty"`
	Quantity            NullFloat64 `json:"quantity,omitempty"`
	Qtotal              NullFloat64 `json:"qtotal,omitempty"`
	Qweight             NullFloat64 `json:"qweight,omitempty"`
	UomWeight           NullString  `json:"uomweight,omitempty"`
	Xs                  NullInt64   `json:"xs,omitempty"`
	S                   NullInt64   `json:"s,omitempty"`
	M                   NullInt64   `json:"m,omitempty"`
	L                   NullInt64   `json:"l,omitempty"`
	Xl                  NullInt64   `json:"xl,omitempty"`
	Xxl                 NullInt64   `json:"xxl,omitempty"`
	Xxxl                NullInt64   `json:"xxxl,omitempty"`
	Os                  NullInt64   `json:"os,omitempty"`
	StockTotal          NullFloat64 `json:"stocktotal,omitempty"`
	StockDisp           NullFloat64 `json:"stockdisp,omitempty"`
	StockTran           NullFloat64 `json:"stocktran,omitempty"`
	StockBloq           NullFloat64 `json:"stockbloq,omitempty"`
	Reactive            NullInt32   `json:"reactive,omitempty"`
	Pigment             NullInt32   `json:"pigment,omitempty"`
	Pfd                 NullInt32   `json:"pfd,omitempty"`
	DivisaId            NullInt64   `json:"divisaid,omitempty"`
	DivisaText          NullString  `json:"divisatext,omitempty"`
	DivisaSimbolo       NullString  `json:"divisasimbolo,omitempty"`
	DivisaDecimal       NullInt32   `json:"divisadecimal,omitempty"`
	CostUnit            NullFloat64 `json:"costunit,omitempty"`
	TasaVenta           NullFloat64 `json:"tasaventa,omitempty"`
	TasaCompra          NullFloat64 `json:"tasacompra,omitempty"`
	Notas               NullString  `json:"notas,omitempty"`
	Ruf1                NullString  `json:"ruf1,omitempty"`
	Ruf2                NullString  `json:"ruf2,omitempty"`
	Ruf3                NullString  `json:"ruf3,omitempty"`
	Iv                  NullString  `json:"iv,omitempty"`
	Salt                NullString  `json:"salt,omitempty"`
	Checksum            NullString  `json:"checksum,omitempty"`
	FCreated            NullTime    `json:"fcreated,omitempty"`
	FUpdated            NullTime    `json:"fupdated,omitempty"`
	UCreated            NullString  `json:"ucreated,omitempty"`
	UUpdated            NullString  `json:"uupdated,omitempty"`
	Activo              int32       `json:"activo,omitempty"`
	Estadoreg           int32       `json:"estadoreg,omitempty"`
	TotalRecords        int64       `json:"total_records,omitempty"`
}

func (e StoreInventoryE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

func (e StoreInventoryE) CreatedFormat() string {
	return e.FCreated.Time.Format("Jan 2006")
}

func (e StoreInventoryE) ColumnsToExport() []string {
	columnas := []string{"uniqueid", "warehousetext", "locationseqid", "productcode", "producttext", "stocktotal", "udisplay", "quom", "uom", "divisioncode", "lotid", "skunumber", "documenttext", "producttypeid", "statusitemid"}
	return columnas
}

const queryListStoreInventoryE = `select uniqueid, sede, flag1, flag2, locationid, locationseqid, warehousetext, inventorytypeid, productid, productcode, producttext, producttypeid, statusitemid, lotid, lotexpired, udisplay, uom, quom, quantity, qtotal, qweight, uomweight,stocktotal, stockdisp, lotid, skunumber, serialnumber, documenttext, sequenceid, receiveddate, manufactureddate, divisioncode, fcreated, activo, estadoreg, total_records from store_inventory_list( $1, $2)`
const queryListStoreInventoryExtendedE = `select uniqueid, sede, flag1, flag2, locationid, locationseqid, warehousetext, inventorytypeid, productid, productcode, producttext, producttypeid, statusitemid, lotid, lotexpired, udisplay, uom, quom, quantity, qtotal, qweight, uomweight,stocktotal, stockdisp, lotid, skunumber, serialnumber, documenttext, personatext, sequenceid, receiveddate, manufactureddate, divisioncode, fcreated, activo, estadoreg, total_records from store_inventory_extended_list( $1, $2)`
const queryListInventoryProductE = `select * from store_inventory_product_list( $1, $2)`
const queryListInventoryPresentaE = `select * from store_inventory_presenta_list( $1, $2)`
const queryListInventoryDivisionE = `select * from store_inventory_division_list( $1, $2)`
const queryListInventoryDocumentsE = `select * from store_inventory_documents_list( $1, $2)`
const queryLoadStoreInventoryE = `select * from store_inventory_list( $1, $2)`
const querySaveStoreInventoryE = `SELECT store_inventory_save($1, $2, $3)`
const queryListStockMinimoE = `select * from store_stock_minimo( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *StoreInventoryE) GetAll(token string, filter string) ([]*StoreInventoryE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStoreInventoryE

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

	var lista []*StoreInventoryE

	for rows.Next() {
		var rowdata StoreInventoryE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.LocationId,
			&rowdata.LocationSeqId,
			&rowdata.WarehouseText,
			&rowdata.InventoryTypeId,
			&rowdata.ProductId,
			&rowdata.ProductCode,
			&rowdata.ProductText,
			&rowdata.ProductTypeId,
			&rowdata.StatusItemId,
			&rowdata.LotId,
			&rowdata.LotExpired,
			&rowdata.UDisplay,
			&rowdata.Uom,
			&rowdata.Quom,
			&rowdata.Quantity,
			&rowdata.Qtotal,
			&rowdata.Qweight,
			&rowdata.UomWeight,
			&rowdata.StockTotal,
			&rowdata.StockDisp,
			&rowdata.LotId,
			&rowdata.SkuNumber,
			&rowdata.SerialNumber,
			&rowdata.DocumentText,
			&rowdata.SequenceId,
			&rowdata.ReceivedDate,
			&rowdata.ManufacturedDate,
			&rowdata.DivisionCode,
			&rowdata.FCreated,
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

func (u *StoreInventoryE) GetAllExtended(token string, filter string) ([]*StoreInventoryE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStoreInventoryExtendedE

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

	var lista []*StoreInventoryE

	for rows.Next() {
		var rowdata StoreInventoryE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.LocationId,
			&rowdata.LocationSeqId,
			&rowdata.WarehouseText,
			&rowdata.InventoryTypeId,
			&rowdata.ProductId,
			&rowdata.ProductCode,
			&rowdata.ProductText,
			&rowdata.ProductTypeId,
			&rowdata.StatusItemId,
			&rowdata.LotId,
			&rowdata.LotExpired,
			&rowdata.UDisplay,
			&rowdata.Uom,
			&rowdata.Quom,
			&rowdata.Quantity,
			&rowdata.Qtotal,
			&rowdata.Qweight,
			&rowdata.UomWeight,
			&rowdata.StockTotal,
			&rowdata.StockDisp,
			&rowdata.LotId,
			&rowdata.SkuNumber,
			&rowdata.SerialNumber,
			&rowdata.DocumentText,
			&rowdata.PersonaText,
			&rowdata.SequenceId,
			&rowdata.ReceivedDate,
			&rowdata.ManufacturedDate,
			&rowdata.DivisionCode,
			&rowdata.FCreated,
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
func (u *StoreInventoryE) GetByUniqueid(token string, jsonText string) (*StoreInventoryE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryLoadStoreInventoryE

	var rowdata StoreInventoryE
	log.Println("Where = " + string(jsonText))
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
		&rowdata.WarehouseId,
		&rowdata.WarehouseText,
		&rowdata.LocationId,
		&rowdata.LocationSeqId,
		&rowdata.InventoryTypeId,
		&rowdata.ProductId,
		&rowdata.ProductCode,
		&rowdata.ProductText,
		&rowdata.DivisionCode,
		&rowdata.VirtualInfo,
		&rowdata.Inventario,
		&rowdata.ProductTypeId,
		&rowdata.StatusItemId,
		&rowdata.ReceivedDate,
		&rowdata.ManufacturedDate,
		&rowdata.ExpiredDate,
		&rowdata.ContainerId,
		&rowdata.LotId,
		&rowdata.LotExpired,
		&rowdata.SkuNumber,
		&rowdata.BinNumber,
		&rowdata.SerialNumber,
		&rowdata.SoftIdentifier,
		&rowdata.BarcodeBox,
		&rowdata.BarcodeItem,
		&rowdata.DocumentText,
		&rowdata.SequenceId,
		&rowdata.ActivationNumber,
		&rowdata.ActivationValidThru,
		&rowdata.UDisplay,
		&rowdata.Uom,
		&rowdata.Quom,
		&rowdata.Quantity,
		&rowdata.Qtotal,
		&rowdata.Qweight,
		&rowdata.UomWeight,
		&rowdata.Xs,
		&rowdata.S,
		&rowdata.M,
		&rowdata.L,
		&rowdata.Xl,
		&rowdata.Xxl,
		&rowdata.Xxxl,
		&rowdata.Os,
		&rowdata.StockTotal,
		&rowdata.StockDisp,
		&rowdata.StockTran,
		&rowdata.StockBloq,
		&rowdata.Reactive,
		&rowdata.Pigment,
		&rowdata.Pfd,
		&rowdata.DivisaId,
		&rowdata.DivisaText,
		&rowdata.DivisaSimbolo,
		&rowdata.DivisaDecimal,
		&rowdata.CostUnit,
		&rowdata.TasaVenta,
		&rowdata.TasaCompra,
		&rowdata.Notas,
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

	return &rowdata, nil
}

func (u *StoreInventoryE) GetAllByPresentacion(token string, filter string) ([]*StoreInventoryE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListInventoryPresentaE

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

	var lista []*StoreInventoryE

	// `select uniqueid, sede, flag1, flag2, locationid, locationseqid, inventorytypeid, productid, producttext, statusitemid, lotid, udisplay, uom, quom, quantity, stocktotal, stockdisp, fcreated, activo, estadoreg, total_records from store_inventory_list( $1, $2)`
	for rows.Next() {
		var rowdata StoreInventoryE
		err := rows.Scan(
			&rowdata.ProductId,
			&rowdata.ProductText,
			&rowdata.UDisplay,
			&rowdata.Uom,
			&rowdata.Quom,
			&rowdata.Quantity,
			&rowdata.Qtotal,
			&rowdata.StockTotal,
			&rowdata.StockDisp,
			&rowdata.StockTran,
			&rowdata.StockBloq,
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

func (u *StoreInventoryE) GetAllByProducts(token string, filter string) ([]*StoreInventoryE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListInventoryProductE

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

	var lista []*StoreInventoryE

	// `select uniqueid, sede, flag1, flag2, locationid, locationseqid, inventorytypeid, productid, producttext, statusitemid, lotid, udisplay, uom, quom, quantity, stocktotal, stockdisp, fcreated, activo, estadoreg, total_records from store_inventory_list( $1, $2)`
	for rows.Next() {
		var rowdata StoreInventoryE
		err := rows.Scan(
			&rowdata.ProductId,
			&rowdata.ProductText,
			&rowdata.UDisplay,
			&rowdata.Uom,
			&rowdata.Quantity,
			&rowdata.Qtotal,
			&rowdata.StockTotal,
			&rowdata.StockDisp,
			&rowdata.StockTran,
			&rowdata.StockBloq,
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

func (u *StoreInventoryE) GetAllByDivision(token string, filter string) ([]*StoreInventoryE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListInventoryDivisionE

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

	var lista []*StoreInventoryE

	// `select uniqueid, sede, flag1, flag2, locationid, locationseqid, inventorytypeid, productid, producttext, statusitemid, lotid, udisplay, uom, quom, quantity, stocktotal, stockdisp, fcreated, activo, estadoreg, total_records from store_inventory_list( $1, $2)`
	for rows.Next() {
		var rowdata StoreInventoryE
		err := rows.Scan(
			&rowdata.DivisionCode,
			&rowdata.UDisplay,
			&rowdata.Uom,
			&rowdata.Quantity,
			&rowdata.Qtotal,
			&rowdata.StockTotal,
			&rowdata.StockDisp,
			&rowdata.StockTran,
			&rowdata.StockBloq,
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

func (u *StoreInventoryE) GetAllByDocument(token string, filter string) ([]*StoreInventoryE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListInventoryDocumentsE

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

	var lista []*StoreInventoryE

	// `select uniqueid, sede, flag1, flag2, locationid, locationseqid, inventorytypeid, productid, producttext, statusitemid, lotid, udisplay, uom, quom, quantity, stocktotal, stockdisp, fcreated, activo, estadoreg, total_records from store_inventory_list( $1, $2)`
	for rows.Next() {
		var rowdata StoreInventoryE
		err := rows.Scan(
			&rowdata.DocumentText,
			&rowdata.Quantity,
			&rowdata.Qtotal,
			&rowdata.StockTotal,
			&rowdata.StockDisp,
			&rowdata.StockTran,
			&rowdata.StockBloq,
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

func (u *StoreInventoryE) GetStockMinimo(token string, filter string) ([]*StoreInventoryE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStockMinimoE

	// Se deseenvuelve el JSON del Filter para adicionar filtros
	var mapFilter map[string]interface{}
	json.Unmarshal([]byte(filter), &mapFilter)
	if mapFilter == nil {
		mapFilter = make(map[string]interface{})
	}
	// --- Adicion de filtros
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

	var lista []*StoreInventoryE

	for rows.Next() {
		var rowdata StoreInventoryE
		err := rows.Scan(
			&rowdata.ProductId,
			&rowdata.ProductCode,
			&rowdata.ProductText,
			&rowdata.DivisionCode,
			&rowdata.UDisplay,
			&rowdata.StockMinimo,
			&rowdata.StockDisp,
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

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *StoreInventoryE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreInventoryE
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
func (u *StoreInventoryE) Delete(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Se deseenvuelve el JSON del Data para adicionar filtros
	var mapData map[string]interface{}
	json.Unmarshal([]byte(data), &mapData)
	if mapData == nil {
		mapData = make(map[string]interface{})
	}
	// --- Adicion de estado de eliminacion de record
	mapData["estadoreg"] = 300

	// Se empaqueta el JSON del Data
	jsonData, err := json.Marshal(mapData)
	if err != nil {
		log.Println("Error convirtiendo el Dato")
		return nil, err
	}
	log.Println("Data = " + string(jsonData))

	query := querySaveStoreInventoryE
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

// DeleteByID deletes one user from the database, by ID
func (u *StoreInventoryE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := querySaveStoreInventoryE
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

func (u *StoreInventoryE) GenerarSKU(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Printf("GenerarSKU [data]=%s [metricas]=%s\n", data, metricas)

	query := `CALL generate_skus($1, $2, $3, $4)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var v_uniqueid float64
	///_, err = stmt.ExecContext(ctx, token, data, metricas, &v_uniqueid)
	err = stmt.QueryRowContext(ctx, token, data, metricas, &v_uniqueid).Scan(&v_uniqueid)
	if err != nil {
		return nil, err
	}
	///defer result.Close()

	retorno := make(map[string]any)
	retorno["uniqueid"] = v_uniqueid

	log.Printf("GenerarSKU [ID]=%v\n", v_uniqueid)

	return retorno, nil
}

func (u *StoreInventoryE) TransferStock(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Printf("TransferStock [data]=%s [metricas]=%s\n", data, metricas)

	query := `CALL transfer_stock($1, $2, $3, $4)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var v_uniqueid float64
	err = stmt.QueryRowContext(ctx, token, data, metricas, &v_uniqueid).Scan(&v_uniqueid)
	if err != nil {
		return nil, err
	}

	retorno := make(map[string]any)
	retorno["uniqueid"] = v_uniqueid

	log.Printf("TransferStock [ID]=%v\n", v_uniqueid)

	return retorno, nil
}

func (u *StoreInventoryE) AdjustStock(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Printf("AdjustStock [data]=%s [metricas]=%s\n", data, metricas)

	query := `CALL adjust_stock($1, $2, $3, $4)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var v_uniqueid float64
	err = stmt.QueryRowContext(ctx, token, data, metricas, &v_uniqueid).Scan(&v_uniqueid)
	if err != nil {
		return nil, err
	}

	retorno := make(map[string]any)
	retorno["uniqueid"] = v_uniqueid

	log.Printf("AdjustStock [ID]=%v\n", v_uniqueid)

	return retorno, nil
}

func (u *StoreInventoryE) UpdateInfoStock(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Printf("UpdateInfoStock [data]=%s [metricas]=%s\n", data, metricas)

	query := `CALL update_info_stock($1, $2, $3, $4)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var v_uniqueid float64
	err = stmt.QueryRowContext(ctx, token, data, metricas, &v_uniqueid).Scan(&v_uniqueid)
	if err != nil {
		return nil, err
	}
	///defer result.Close()

	retorno := make(map[string]any)
	retorno["uniqueid"] = v_uniqueid

	log.Printf("UpdateInfoStock [ID]=%v\n", v_uniqueid)

	return retorno, nil
}
