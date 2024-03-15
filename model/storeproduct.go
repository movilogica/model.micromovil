package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Items
type StoreProductE struct {
	Uniqueid        int64       `json:"uniqueid,omitempty"`
	Owner           NullInt32   `json:"owner,omitempty"`
	Dispositivoid   int32       `json:"dispositivoid,omitempty"`
	Id              int32       `json:"id,omitempty"`
	Sede            int32       `json:"sede"`
	Flag1           string      `json:"flag1,omitempty"`
	Flag2           string      `json:"flag2,omitempty"`
	PersonaId       NullInt64   `json:"personaid,omitempty"`
	TokendataId     NullString  `json:"tokendataid,omitempty"`
	ParentId        NullInt64   `json:"parentid,omitempty"`
	Code            NullString  `json:"code,omitempty"`
	ProductName     NullString  `json:"productname,omitempty"`
	InternalName    NullString  `json:"internalname,omitempty"`
	PriceDetailText NullString  `json:"pricedetailtext,omitempty"`
	DetailScreen    NullString  `json:"detailscreen,omitempty"`
	ProductTypeId   NullString  `json:"producttypeid,omitempty"`
	UPresent        NullString  `json:"upresent,omitempty"`
	UMedPack        NullString  `json:"umedpack,omitempty"`
	QByPack         NullFloat64 `json:"qbypack,omitempty"`
	CategItemId     NullInt64   `json:"categitemid,omitempty"`
	CategItemText   NullString  `json:"categitemtext,omitempty"`
	Color           NullString  `json:"color,omitempty"`
	Marca           NullString  `json:"marca,omitempty"`
	Price           NullFloat64 `json:"price,omitempty"`
	Purchase        NullFloat64 `json:"purchase,omitempty"`
	StockMinimo     NullInt32   `json:"stockminimo,omitempty"`
	SmallImageUrl   NullString  `json:"smallimageurl,omitempty"`
	MediumImageUrl  NullString  `json:"mediumimageurl,omitempty"`
	LargeImageUrl   NullString  `json:"largeimageurl,omitempty"`
	Ruf1            NullString  `json:"ruf1,omitempty"`
	Ruf2            NullString  `json:"ruf2,omitempty"`
	Ruf3            NullString  `json:"ruf3,omitempty"`
	Iv              NullString  `json:"iv,omitempty"`
	Salt            NullString  `json:"salt,omitempty"`
	Checksum        NullString  `json:"checksum,omitempty"`
	FCreated        NullTime    `json:"fcreated,omitempty"`
	FUpdated        NullTime    `json:"fupdated,omitempty"`
	Activo          NullInt64   `json:"activo,omitempty"`
	Estadoreg       NullInt64   `json:"estadoreg,omitempty"`
	TotalRecords    int64       `json:"total_records"`
}

func (e StoreProductE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectStoreProducts = `select * from store_products_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *StoreProductE) GetAll(token string, filter string) ([]*StoreProductE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectStoreProducts

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

	var lista []*StoreProductE

	for rows.Next() {
		var rowdata StoreProductE
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
			&rowdata.ParentId,
			&rowdata.Code,
			&rowdata.ProductName,
			&rowdata.InternalName,
			&rowdata.PriceDetailText,
			&rowdata.DetailScreen,
			&rowdata.ProductTypeId,
			&rowdata.UPresent,
			&rowdata.UMedPack,
			&rowdata.QByPack,
			&rowdata.CategItemId,
			&rowdata.CategItemText,
			&rowdata.Color,
			&rowdata.Marca,
			&rowdata.Price,
			&rowdata.Purchase,
			&rowdata.StockMinimo,
			&rowdata.SmallImageUrl,
			&rowdata.MediumImageUrl,
			&rowdata.LargeImageUrl,
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
func (u *StoreProductE) GetByUniqueid(token string, uniqueid int) (*StoreProductE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectStoreProducts

	var rowdata StoreProductE
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
		&rowdata.ParentId,
		&rowdata.Code,
		&rowdata.ProductName,
		&rowdata.InternalName,
		&rowdata.PriceDetailText,
		&rowdata.DetailScreen,
		&rowdata.ProductTypeId,
		&rowdata.UPresent,
		&rowdata.UMedPack,
		&rowdata.QByPack,
		&rowdata.CategItemId,
		&rowdata.CategItemText,
		&rowdata.Color,
		&rowdata.Marca,
		&rowdata.Price,
		&rowdata.Purchase,
		&rowdata.StockMinimo,
		&rowdata.SmallImageUrl,
		&rowdata.MediumImageUrl,
		&rowdata.LargeImageUrl,
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
func (u *StoreProductE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT store_products_save($1, $2, $3)`
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
func (u *StoreProductE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT store_products_save($1, $2, $3)`
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
func (u *StoreProductE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := `SELECT store_products_save($1, $2, $3)`
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
