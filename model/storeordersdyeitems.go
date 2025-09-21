package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Ordenes Dye House - items
type StoreOrdersDyeItemsE struct {
	Uniqueid            int64       `json:"uniqueid,omitempty"`
	Owner               NullInt32   `json:"owner,omitempty"`
	Dispositivoid       NullInt32   `json:"dispositivoid,omitempty"`
	Id                  int32       `json:"id,omitempty"`
	Sede                int32       `json:"sede"`
	Flag1               string      `json:"flag1,omitempty"`
	Flag2               string      `json:"flag2,omitempty"`
	PersonaId           NullInt64   `json:"personaid,omitempty"`
	TokendataId         NullString  `json:"tokendataid,omitempty"`
	BizPersonaId        NullInt64   `json:"bizpersonaid,omitempty"`
	OrderId             NullInt64   `json:"orderid,omitempty"`
	Nroperacion         NullString  `json:"nroperacion,omitempty"`
	Secuencial          NullInt32   `json:"secuencial,omitempty"`
	Fecha               NullTime    `json:"fecha,omitempty"`
	FDyeHouse           NullTime    `json:"fdyehouse,omitempty"`
	FEstimated          NullTime    `json:"festimated,omitempty"`
	FProcessed          NullTime    `json:"fprocessed,omitempty"`
	FCancelled          NullTime    `json:"fcancelled,omitempty"`
	StatusId            NullString  `json:"statusid,omitempty"`
	DyeOrder            NullString  `json:"dyeorder,omitempty"`
	CutTextil           NullString  `json:"cuttextil,omitempty"`
	ProductId           NullInt64   `json:"productid,omitempty"`
	ProductCode         NullString  `json:"productcode,omitempty"`
	ProductText         NullString  `json:"producttext,omitempty"`
	DivisionCode        NullString  `json:"divisioncode,omitempty"`
	ColorCode           NullString  `json:"colorcode,omitempty"`
	BackColor           NullString  `json:"backcolor,omitempty"`
	ForeColor           NullString  `json:"forecolor,omitempty"`
	Yield               NullString  `json:"yield,omitempty"`
	FinalWidth          NullFloat64 `json:"finalwidth,omitempty"`
	OverrideglAccountId NullString  `json:"overrideglaccountid,omitempty"`
	SupplierProductId   NullString  `json:"supplierproductid,omitempty"`
	ProductTypeId       NullString  `json:"producttypeid,omitempty"`
	TipoProductoId      NullString  `json:"tipoproductoid,omitempty"`
	TipoProductoText    NullString  `json:"tipoproductotext,omitempty"`
	Lotid               NullString  `json:"lotid,omitempty"`
	LotExpired          NullTime    `json:"lotexpired,omitempty"`
	Inventario          NullString  `json:"inventario,omitempty"`
	Features            NullString  `json:"features,omitempty"`
	UnitPrice           NullFloat64 `json:"unitprice,omitempty"`
	UDisplay            NullString  `json:"udisplay,omitempty"`
	Uom                 NullString  `json:"uom,omitempty"`
	Quom                NullFloat64 `json:"quom,omitempty"`
	Quantity            NullFloat64 `json:"quantity,omitempty"`
	QTotal              NullFloat64 `json:"qtotal,omitempty"`
	QGrossWeight        NullFloat64 `json:"qgrossweight,omitempty"`
	QNetWeight          NullFloat64 `json:"qnetweight,omitempty"`
	UomWeight           NullString  `json:"uomweight,omitempty"`
	Xs                  NullInt64   `json:"xs,omitempty"`
	S                   NullInt64   `json:"s,omitempty"`
	M                   NullInt64   `json:"m,omitempty"`
	L                   NullInt64   `json:"l,omitempty"`
	Xl                  NullInt64   `json:"xl,omitempty"`
	Xxl                 NullInt64   `json:"xxl,omitempty"`
	Xxxl                NullInt64   `json:"xxxl,omitempty"`
	Os                  NullInt64   `json:"os,omitempty"`
	DocumentText        NullString  `json:"documenttext,omitempty"`
	Roadmap             NullString  `json:"roadmap,omitempty"`
	Froadmap            NullTime    `json:"froadmap,omitempty"`
	Received            NullFloat64 `json:"received,omitempty"`
	Cancelled           NullFloat64 `json:"cancelled,omitempty"`
	Subtotal            NullFloat64 `json:"subtotal,omitempty"`
	Comentarios         NullString  `json:"comentarios,omitempty"`
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

func (e StoreOrdersDyeItemsE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

// Para la edicion se obtienen solo campos de la misma tabla.
const queryListEditOrdersDyeItemsE = `select uniqueid, sede, flag1, flag2, orderid, secuencial, productid, productcode, producttext, divisioncode, colorcode, backcolor, forecolor, yield, finalwidth, inventario, features, lotid, lotexpired, unitprice, udisplay, uom, quom, quantity, qtotal, qgrossweight, qnetweight, uomweight, documenttext, roadmap, froadmap, received, cancelled, subtotal, comentarios, activo, estadoreg, total_records from store_orders_dye_items_list( $1, $2)`

// La lista de items puede incluir columnas relacionadas de otras tablas
const queryListStoreOrdersDyeItemsE = `select uniqueid, sede, flag1, flag2, orderid, secuencial, fecha, fdyehouse, festimated, fprocessed, fcancelled, statusid, dyeorder, cuttextil, productid, productcode, producttext, divisioncode, colorcode, backcolor, forecolor, yield, finalwidth, inventario, features, lotid, lotexpired, unitprice, udisplay, uom, quom, quantity, qtotal, qgrossweight, qnetweight, uomweight, documenttext, roadmap, froadmap, received, cancelled, subtotal, comentarios, activo, estadoreg, total_records from store_orders_dye_items_list( $1, $2)`
const queryLoadStoreOrdersDyeItemsE = `select * from store_orders_dye_items_list( $1, $2)`
const querySaveStoreOrdersDyeItemsE = `SELECT store_orders_dye_items_save($1, $2, $3)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *StoreOrdersDyeItemsE) GetAll(token string, filter string) ([]*StoreOrdersDyeItemsE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStoreOrdersDyeItemsE

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

	var lista []*StoreOrdersDyeItemsE

	/// `select uniqueid, sede, flag1, flag2, orderid, secuencial, fecha, fdyehouse, festimated, fprocessed, fcancelled,
	//          statusid, dyeorder, cuttextil, productid, productcode, producttext, divisioncode, colortext, yield, finalwidth,
	//          features, lotid, lotexpired, unitprice, udisplay, uom, quom, quantity, qtotal, qgrossweight, qnetweight, uomweight, roadmap, froadmap, received, cancelled, subtotal, comentarios, activo, estadoreg, total_records
	for rows.Next() {
		var rowdata StoreOrdersDyeItemsE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.OrderId,
			&rowdata.Secuencial,
			&rowdata.Fecha,
			&rowdata.FDyeHouse,
			&rowdata.FEstimated,
			&rowdata.FProcessed,
			&rowdata.FCancelled,
			&rowdata.StatusId,
			&rowdata.DyeOrder,
			&rowdata.CutTextil,
			&rowdata.ProductId,
			&rowdata.ProductCode,
			&rowdata.ProductText,
			&rowdata.DivisionCode,
			&rowdata.ColorCode,
			&rowdata.BackColor,
			&rowdata.ForeColor,
			&rowdata.Yield,
			&rowdata.FinalWidth,
			&rowdata.Inventario,
			&rowdata.Features,
			&rowdata.Lotid,
			&rowdata.LotExpired,
			&rowdata.UnitPrice,
			&rowdata.UDisplay,
			&rowdata.Uom,
			&rowdata.Quom,
			&rowdata.Quantity,
			&rowdata.QTotal,
			&rowdata.QGrossWeight,
			&rowdata.QNetWeight,
			&rowdata.UomWeight,
			&rowdata.DocumentText,
			&rowdata.Roadmap,
			&rowdata.Froadmap,
			&rowdata.Received,
			&rowdata.Cancelled,
			&rowdata.Subtotal,
			&rowdata.Comentarios,
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
	log.Printf("Resultado items = %d records\r\n", len(lista))

	return lista, nil
}

func (u *StoreOrdersDyeItemsE) GetItems(token string, filter string) ([]*StoreOrdersDyeItemsE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListEditOrdersDyeItemsE

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

	var lista []*StoreOrdersDyeItemsE
	for rows.Next() {
		var rowdata StoreOrdersDyeItemsE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.OrderId,
			&rowdata.Secuencial,
			&rowdata.ProductId,
			&rowdata.ProductCode,
			&rowdata.ProductText,
			&rowdata.DivisionCode,
			&rowdata.ColorCode,
			&rowdata.BackColor,
			&rowdata.ForeColor,
			&rowdata.Yield,
			&rowdata.FinalWidth,
			&rowdata.Inventario,
			&rowdata.Features,
			&rowdata.Lotid,
			&rowdata.LotExpired,
			&rowdata.UnitPrice,
			&rowdata.UDisplay,
			&rowdata.Uom,
			&rowdata.Quom,
			&rowdata.Quantity,
			&rowdata.QTotal,
			&rowdata.QGrossWeight,
			&rowdata.QNetWeight,
			&rowdata.UomWeight,
			&rowdata.DocumentText,
			&rowdata.Roadmap,
			&rowdata.Froadmap,
			&rowdata.Received,
			&rowdata.Cancelled,
			&rowdata.Subtotal,
			&rowdata.Comentarios,
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
	log.Printf("Resultado items = %d records\r\n", len(lista))

	return lista, nil
}

// GetOne returns one user by id
func (u *StoreOrdersDyeItemsE) GetByUniqueid(token string, jsonText string) (*StoreOrdersDyeItemsE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryLoadStoreOrdersDyeItemsE

	var rowdata StoreOrdersDyeItemsE
	log.Printf("[%s] Where = %s\n", query, string(jsonText))
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
		&rowdata.OrderId,
		&rowdata.Nroperacion,
		&rowdata.Secuencial,
		&rowdata.Fecha,
		&rowdata.FDyeHouse,
		&rowdata.FEstimated,
		&rowdata.FProcessed,
		&rowdata.FCancelled,
		&rowdata.StatusId,
		&rowdata.DyeOrder,
		&rowdata.CutTextil,
		&rowdata.ProductId,
		&rowdata.ProductCode,
		&rowdata.ProductText,
		&rowdata.DivisionCode,
		&rowdata.ColorCode,
		&rowdata.BackColor,
		&rowdata.ForeColor,
		&rowdata.Yield,
		&rowdata.FinalWidth,
		&rowdata.OverrideglAccountId,
		&rowdata.SupplierProductId,
		&rowdata.ProductTypeId,
		&rowdata.TipoProductoId,
		&rowdata.TipoProductoText,
		&rowdata.Lotid,
		&rowdata.LotExpired,
		&rowdata.Inventario,
		&rowdata.Features,
		&rowdata.UnitPrice,
		&rowdata.UDisplay,
		&rowdata.Uom,
		&rowdata.Quom,
		&rowdata.Quantity,
		&rowdata.QTotal,
		&rowdata.QGrossWeight,
		&rowdata.QNetWeight,
		&rowdata.UomWeight,
		&rowdata.Xs,
		&rowdata.S,
		&rowdata.M,
		&rowdata.L,
		&rowdata.Xl,
		&rowdata.Xxl,
		&rowdata.Xxxl,
		&rowdata.Os,
		&rowdata.DocumentText,
		&rowdata.Roadmap,
		&rowdata.Froadmap,
		&rowdata.Received,
		&rowdata.Cancelled,
		&rowdata.Subtotal,
		&rowdata.Comentarios,
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

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *StoreOrdersDyeItemsE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreOrdersDyeItemsE
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
func (u *StoreOrdersDyeItemsE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreOrdersDyeItemsE
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
func (u *StoreOrdersDyeItemsE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := querySaveStoreOrdersDyeItemsE
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

func (u *StoreOrdersDyeItemsE) UMedidaText() string {
	umedidatext := ""
	/// 20Kgm/Rollo
	if u.Uom != u.UDisplay {
		if u.Quom.Float64 > 1 {
			umedidatext = fmt.Sprintf("%.2f%s//%s", u.Quom.Float64, u.Uom.String, u.UDisplay.String)
		} else {
			umedidatext = u.UDisplay.String
		}
	} else {
		if u.Quom.Float64 > 1 {
			umedidatext = fmt.Sprintf("%.2fUN//%s", u.Quom.Float64, u.UDisplay.String)
		} else {
			umedidatext = u.UDisplay.String
		}
	}
	if umedidatext == "" {
		return "-"
	} else {
		return umedidatext
	}
}
