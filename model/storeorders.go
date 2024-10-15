package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Ordenes de Warehouse
type StoreOrdersE struct {
	Uniqueid        int64       `json:"uniqueid,omitempty"`
	Owner           NullInt32   `json:"owner,omitempty"`
	Dispositivoid   NullInt32   `json:"dispositivoid,omitempty"`
	Id              int32       `json:"id,omitempty"`
	Sede            int32       `json:"sede"`
	Flag1           string      `json:"flag1,omitempty"`
	Flag2           string      `json:"flag2,omitempty"`
	PersonaId       NullInt64   `json:"personaid,omitempty"`
	TokendataId     NullString  `json:"tokendataid,omitempty"`
	BizPersonaId    NullInt64   `json:"bizpersonaid,omitempty"`
	PersonaText     NullString  `json:"personatext,omitempty"`
	Email           NullString  `json:"email,omitempty"`
	Movil           NullString  `json:"movil,omitempty"`
	OrderName       NullString  `json:"ordername,omitempty"`
	Priority        NullInt32   `json:"priority,omitempty"`
	ToWarehouseId   NullInt64   `json:"towarehouseid,omitempty"`
	ToStoreId       NullInt64   `json:"tostoreid,omitempty"`
	ToWarehouseText NullString  `json:"towarehousetext,omitempty"`
	ToStoreText     NullString  `json:"tostoretext,omitempty"`
	Fecha           NullTime    `json:"fecha,omitempty"`
	FEstimated      NullTime    `json:"festimated,omitempty"`
	FDelivery       NullTime    `json:"fdelivery,omitempty"`
	TipoOrdenId     NullInt64   `json:"tipordenid,omitempty"`
	TipoOrdenText   NullString  `json:"tipordentext,omitempty"`
	OrderTypeId     NullString  `json:"ordertypeid,omitempty"`
	StatusId        NullString  `json:"statusid,omitempty"`
	Numero          NullString  `json:"numero,omitempty"`
	NumeroExt       NullString  `json:"numeroext,omitempty"`
	Nroperacion     NullString  `json:"nroperacion,omitempty"`
	NroperacionMask NullString  `json:"nroperacionmask,omitempty"`
	DivisaId        NullInt64   `json:"divisaid,omitempty"`
	DivisaText      NullString  `json:"divisatext,omitempty"`
	DivisaSimbolo   NullString  `json:"divisasimbolo,omitempty"`
	DivisaDecimal   NullInt32   `json:"divisadecimal,omitempty"`
	TasaVenta       NullFloat64 `json:"tasaventa,omitempty"`
	TasaCompra      NullFloat64 `json:"tasacompra,omitempty"`
	TotaLineas      NullInt32   `json:"totalineas,omitempty"`
	TotalItems      NullFloat64 `json:"totalitems,omitempty"`
	TotalOrden      NullFloat64 `json:"totalorden,omitempty"`
	PrecioPromedio  NullFloat64 `json:"precioprom,omitempty"`
	Comentarios     NullString  `json:"comentarios,omitempty"`
	Latitud         NullFloat64 `json:"latitud,omitempty"`
	Longitud        NullFloat64 `json:"longitud,omitempty"`
	Checkin         NullInt32   `json:"checkin,omitempty"`
	Ruf1            NullString  `json:"ruf1,omitempty"`
	Ruf2            NullString  `json:"ruf2,omitempty"`
	Ruf3            NullString  `json:"ruf3,omitempty"`
	Iv              NullString  `json:"iv,omitempty"`
	Salt            NullString  `json:"salt,omitempty"`
	Checksum        NullString  `json:"checksum,omitempty"`
	FCreated        NullTime    `json:"fcreated,omitempty"`
	FUpdated        NullTime    `json:"fupdated,omitempty"`
	UCreated        NullString  `json:"ucreated,omitempty"`
	UUpdated        NullString  `json:"uupdated,omitempty"`
	Activo          int32       `json:"activo,omitempty"`
	Estadoreg       int32       `json:"estadoreg,omitempty"`
	TotalRecords    int64       `json:"total_records,omitempty"`
}

func (e StoreOrdersE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const queryListStoreOrdersE = `select uniqueid, sede, flag1, flag2, bizpersonaid, personatext, towarehousetext, tostoretext, fecha, fdelivery, ordertypeid, statusid, numero, nroperacionmask, divisasimbolo, totalorden activo, estadoreg, total_records from store_orders_list( $1, $2)`
const queryLoadStoreOrdersE = `select * from store_orders_list( $1, $2)`
const querySaveStoreOrdersE = `SELECT store_orders_save($1, $2, $3)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *StoreOrdersE) GetAll(token string, filter string) ([]*StoreOrdersE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStoreOrdersE

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

	var lista []*StoreOrdersE

	/// `select uniqueid, sede, flag1, flag2, bizpersonaid, personatext, towarehousetext, tostoretext, fecha, fdelivery,
	///         ordertypeid, statusid, numero, nroperacionmask, divisasimbolo, totalorden activo, estadoreg, total_records
	for rows.Next() {
		var rowdata StoreOrdersE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.BizPersonaId,
			&rowdata.PersonaText,
			&rowdata.ToWarehouseText,
			&rowdata.ToStoreText,
			&rowdata.Fecha,
			&rowdata.FDelivery,
			&rowdata.OrderTypeId,
			&rowdata.StatusId,
			&rowdata.Numero,
			&rowdata.NroperacionMask,
			&rowdata.DivisaSimbolo,
			&rowdata.TotalOrden,
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
func (u *StoreOrdersE) GetByUniqueid(token string, jsonText string) (*StoreOrdersE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryLoadStoreOrdersE

	var rowdata StoreOrdersE
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
		&rowdata.PersonaText,
		&rowdata.Email,
		&rowdata.Movil,
		&rowdata.OrderName,
		&rowdata.Priority,
		&rowdata.ToWarehouseId,
		&rowdata.ToStoreId,
		&rowdata.ToWarehouseText,
		&rowdata.ToStoreText,
		&rowdata.Fecha,
		&rowdata.FEstimated,
		&rowdata.FDelivery,
		&rowdata.TipoOrdenId,
		&rowdata.TipoOrdenText,
		&rowdata.OrderTypeId,
		&rowdata.StatusId,
		&rowdata.Numero,
		&rowdata.NumeroExt,
		&rowdata.Nroperacion,
		&rowdata.NroperacionMask,
		&rowdata.DivisaId,
		&rowdata.DivisaText,
		&rowdata.DivisaSimbolo,
		&rowdata.DivisaDecimal,
		&rowdata.TasaVenta,
		&rowdata.TasaCompra,
		&rowdata.TotaLineas,
		&rowdata.TotalItems,
		&rowdata.TotalOrden,
		&rowdata.PrecioPromedio,
		&rowdata.Comentarios,
		&rowdata.Latitud,
		&rowdata.Longitud,
		&rowdata.Checkin,
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
func (u *StoreOrdersE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreOrdersE
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
func (u *StoreOrdersE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreOrdersE
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
func (u *StoreOrdersE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := querySaveStoreOrdersE
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
