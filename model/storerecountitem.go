package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Inventory Recuento Items
type StoreRecountItemE struct {
	Uniqueid      int64       `json:"uniqueid,omitempty"`
	Owner         NullInt32   `json:"owner,omitempty"`
	Dispositivoid NullInt32   `json:"dispositivoid,omitempty"`
	Id            int32       `json:"id,omitempty"`
	Sede          int32       `json:"sede"`
	Flag1         string      `json:"flag1,omitempty"`
	Flag2         string      `json:"flag2,omitempty"`
	PersonaId     NullInt64   `json:"personaid,omitempty"`
	TokendataId   NullString  `json:"tokendataid,omitempty"`
	WarehouseId   NullInt64   `json:"warehouseid,omitempty"`
	RecountId     NullInt64   `json:"recountid,omitempty"`
	GroupId       NullInt64   `json:"groupid,omitempty"`
	GroupText     NullString  `json:"grouptext,omitempty"`
	PersonaText   NullString  `json:"personatext,omitempty"`
	Numero        NullInt64   `json:"numero,omitempty"`
	BoxId         NullInt64   `json:"boxid,omitempty"`
	BoxIdentifyId NullString  `json:"boxidentityid,omitempty"`
	BarcodeBox    NullString  `json:"barcodebox,omitempty"`
	BarcodeItem   NullString  `json:"barcodeitem,omitempty"`
	StyleCode     NullString  `json:"stylecode,omitempty"`
	StyleText     NullString  `json:"styletext,omitempty"`
	ColorCode     NullString  `json:"colorcode,omitempty"`
	DivisionCode  NullString  `json:"divisioncode,omitempty"`
	AreaId        NullString  `json:"areaid,omitempty"`
	AisleId       NullString  `json:"aisleid,omitempty"`
	SectionId     NullString  `json:"sectionid,omitempty"`
	LevelId       NullString  `json:"levelid,omitempty"`
	PositionId    NullString  `json:"positionid,omitempty"`
	UDisplay      NullString  `json:"udisplay,omitempty"`
	Uom           NullString  `json:"uom,omitempty"`
	Quom          NullInt64   `json:"quom,omitempty"`
	Quantity      NullFloat64 `json:"quantity,omitempty"`
	Xs            NullInt64   `json:"xs,omitempty"`
	S             NullInt64   `json:"s,omitempty"`
	M             NullInt64   `json:"m,omitempty"`
	L             NullInt64   `json:"l,omitempty"`
	Xl            NullInt64   `json:"xl,omitempty"`
	Xxl           NullInt64   `json:"xxl,omitempty"`
	Xxxl          NullInt64   `json:"xxxl,omitempty"`
	Os            NullInt64   `json:"os,omitempty"`
	Total         NullInt64   `json:"total,omitempty"`
	Reactive      NullInt32   `json:"reactive,omitempty"`
	Pigment       NullInt32   `json:"pigment,omitempty"`
	Pfd           NullInt32   `json:"pfd,omitempty"`
	UrlphotoImage NullString  `json:"urlphotoimage,omitempty"`
	Ruf1          NullString  `json:"ruf1,omitempty"`
	Ruf2          NullString  `json:"ruf2,omitempty"`
	Ruf3          NullString  `json:"ruf3,omitempty"`
	Iv            NullString  `json:"iv,omitempty"`
	Salt          NullString  `json:"salt,omitempty"`
	Checksum      NullString  `json:"checksum,omitempty"`
	FCreated      NullTime    `json:"fcreated,omitempty"`
	FUpdated      NullTime    `json:"fupdated,omitempty"`
	UCreated      NullString  `json:"ucreated,omitempty"`
	UUpdated      NullString  `json:"uupdated,omitempty"`
	Activo        int32       `json:"activo,omitempty"`
	Estadoreg     int32       `json:"estadoreg,omitempty"`
	TotalRecords  int64       `json:"total_records,omitempty"`
}

func (e StoreRecountItemE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

func (e StoreRecountItemE) CreatedFormat() string {
	return e.FCreated.Time.Format("Jan 2006")
}

const queryListStoreRecountItemE = `select * from store_recount_item_list( $1, $2)`
const querySaveStoreRecountItemE = `SELECT store_recount_item_save($1, $2, $3)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *StoreRecountItemE) GetAll(token string, filter string) ([]*StoreRecountItemE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStoreRecountItemE

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

	var lista []*StoreRecountItemE

	for rows.Next() {
		var rowdata StoreRecountItemE
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
			&rowdata.WarehouseId,
			&rowdata.RecountId,
			&rowdata.GroupId,
			&rowdata.GroupText,
			&rowdata.PersonaText,
			&rowdata.Numero,
			&rowdata.BoxId,
			&rowdata.BoxIdentifyId,
			&rowdata.BarcodeBox,
			&rowdata.BarcodeItem,
			&rowdata.StyleCode,
			&rowdata.StyleText,
			&rowdata.ColorCode,
			&rowdata.DivisionCode,
			&rowdata.AreaId,
			&rowdata.AisleId,
			&rowdata.SectionId,
			&rowdata.LevelId,
			&rowdata.PositionId,
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
			&rowdata.Total,
			&rowdata.Reactive,
			&rowdata.Pigment,
			&rowdata.Pfd,
			&rowdata.UrlphotoImage,
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

// GetOne returns one user by id
func (u *StoreRecountItemE) GetByUniqueid(token string, uniqueid int) (*StoreRecountItemE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStoreRecountItemE

	var rowdata StoreRecountItemE
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
		&rowdata.WarehouseId,
		&rowdata.RecountId,
		&rowdata.GroupId,
		&rowdata.GroupText,
		&rowdata.PersonaText,
		&rowdata.Numero,
		&rowdata.BoxId,
		&rowdata.BoxIdentifyId,
		&rowdata.BarcodeBox,
		&rowdata.BarcodeItem,
		&rowdata.StyleCode,
		&rowdata.StyleText,
		&rowdata.ColorCode,
		&rowdata.DivisionCode,
		&rowdata.AreaId,
		&rowdata.AisleId,
		&rowdata.SectionId,
		&rowdata.LevelId,
		&rowdata.PositionId,
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
		&rowdata.Total,
		&rowdata.Reactive,
		&rowdata.Pigment,
		&rowdata.Pfd,
		&rowdata.UrlphotoImage,
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

	return &rowdata, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *StoreRecountItemE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreRecountItemE
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
func (u *StoreRecountItemE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreRecountItemE
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
func (u *StoreRecountItemE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := querySaveStoreRecountItemE
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
