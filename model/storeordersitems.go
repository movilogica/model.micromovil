package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Ordenes de Warehouse - items
type StoreOrdersItemsE struct {
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
	ProductId           NullInt64   `json:"productid,omitempty"`
	ProductText         NullString  `json:"producttext,omitempty"`
	Secuencial          NullInt32   `json:"secuencial,omitempty"`
	OverrideglAccountId NullString  `json:"overrideglaccountid,omitempty"`
	SupplierProductId   NullString  `json:"supplierproductid,omitempty"`
	ProdCatalogId       NullInt32   `json:"prodcatalogid,omitempty"`
	ProducTypeId        NullInt32   `json:"productypeid,omitempty"`
	ProdCatalogText     NullString  `json:"prodcatalogtext,omitempty"`
	ProducTypeText      NullString  `json:"productypetext,omitempty"`
	UnitPrice           NullFloat64 `json:"unitprice,omitempty"`
	UDisplay            NullString  `json:"udisplay,omitempty"`
	Uom                 NullString  `json:"uom,omitempty"`
	Quom                NullFloat64 `json:"quom,omitempty"`
	Quantity            NullFloat64 `json:"quantity,omitempty"`
	Xs                  NullInt64   `json:"xs,omitempty"`
	S                   NullInt64   `json:"s,omitempty"`
	M                   NullInt64   `json:"m,omitempty"`
	L                   NullInt64   `json:"l,omitempty"`
	Xl                  NullInt64   `json:"xl,omitempty"`
	Xxl                 NullInt64   `json:"xxl,omitempty"`
	Xxxl                NullInt64   `json:"xxxl,omitempty"`
	Os                  NullInt64   `json:"os,omitempty"`
	QTotal              NullFloat64 `json:"qtotal,omitempty"`
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

func (e StoreOrdersItemsE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const queryListStoreOrdersItemsE = `select uniqueid, sede, flag1, flag2, secuencial, productid, producttext, unitprice, udisplay, uom, quom, quantity, qtotal, received, cancelled, subtotal, activo, estadoreg, total_records from store_orders_items_list( $1, $2)`
const queryLoadStoreOrdersItemsE = `select * from store_orders_items_list( $1, $2)`
const querySaveStoreOrdersItemsE = `SELECT store_orders_items_save($1, $2, $3)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *StoreOrdersItemsE) GetAll(token string, filter string) ([]*StoreOrdersItemsE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStoreOrdersItemsE

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

	var lista []*StoreOrdersItemsE

	/// `select uniqueid, sede, flag1, flag2, secuencial, productid, producttext, unitprice, udisplay, uom, quom,
	///         quantity, qtotal, received, cancelled, subtotal, activo, estadoreg, total_records
	for rows.Next() {
		var rowdata StoreOrdersItemsE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.Secuencial,
			&rowdata.ProductId,
			&rowdata.ProductText,
			&rowdata.UnitPrice,
			&rowdata.UDisplay,
			&rowdata.Uom,
			&rowdata.Quom,
			&rowdata.Quantity,
			&rowdata.QTotal,
			&rowdata.Received,
			&rowdata.Cancelled,
			&rowdata.Subtotal,
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
func (u *StoreOrdersItemsE) GetByUniqueid(token string, jsonText string) (*StoreOrdersItemsE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryLoadStoreOrdersItemsE

	var rowdata StoreOrdersItemsE
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
		&rowdata.ProductId,
		&rowdata.ProductText,
		&rowdata.OverrideglAccountId,
		&rowdata.SupplierProductId,
		&rowdata.ProdCatalogId,
		&rowdata.ProducTypeId,
		&rowdata.ProdCatalogText,
		&rowdata.ProducTypeText,
		&rowdata.UnitPrice,
		&rowdata.UDisplay,
		&rowdata.Uom,
		&rowdata.Quom,
		&rowdata.Quantity,
		&rowdata.Xs,
		&rowdata.S,
		&rowdata.M,
		&rowdata.L,
		&rowdata.Xl,
		&rowdata.Xxl,
		&rowdata.Xxxl,
		&rowdata.Os,
		&rowdata.QTotal,
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
func (u *StoreOrdersItemsE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreOrdersItemsE
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
func (u *StoreOrdersItemsE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreOrdersItemsE
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
func (u *StoreOrdersItemsE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := querySaveStoreOrdersItemsE
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
