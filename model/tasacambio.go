package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Tipos de autos
type TasaCambioE struct {
	Uniqueid      int64       `json:"uniqueid,omitempty"`
	Owner         NullInt32   `json:"owner,omitempty"`
	Dispositivoid int32       `json:"dispositivoid,omitempty"`
	Id            int32       `json:"id,omitempty"`
	Sede          int32       `json:"sede"`
	Flag1         string      `json:"flag1,omitempty"`
	Flag2         string      `json:"flag2,omitempty"`
	Fecha         NullTime    `json:"fecha,omitempty"`
	DivisaId      NullInt64   `json:"divisaid,omitempty"`
	ForeignId     NullInt64   `json:"foreignid,omitempty"`
	ForeignText   NullString  `json:"foreigntext,omitempty"`
	Simbolo       NullString  `json:"simbolo,omitempty"`
	Decimales     NullInt32   `json:"decimales,omitempty"`
	Compra        NullFloat64 `json:"compra,omitempty"`
	Venta         NullFloat64 `json:"ventaday,omitempty"`
	Ruf1          NullString  `json:"ruf1,omitempty"`
	Ruf2          NullString  `json:"ruf2,omitempty"`
	Ruf3          NullString  `json:"ruf3,omitempty"`
	Iv            NullString  `json:"iv,omitempty"`
	Salt          NullString  `json:"salt,omitempty"`
	Checksum      NullString  `json:"checksum,omitempty"`
	FCreated      NullTime    `json:"fcreated,omitempty"`
	FUpdated      NullTime    `json:"fupdated,omitempty"`
	Activo        int32       `json:"activo,omitempty"`
	Estadoreg     int32       `json:"estadoreg,omitempty"`
	TotalRecords  int64       `json:"total_records"`
}

func (e TasaCambioE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectTasaCambio = `select * from param_tasacambio_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *TasaCambioE) GetAll(token string, filter string) ([]*TasaCambioE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectTasaCambio

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

	var lista []*TasaCambioE

	for rows.Next() {
		var rowdata TasaCambioE
		err := rows.Scan(
			&rowdata.DivisaId,
			&rowdata.Uniqueid,
			&rowdata.Owner,
			&rowdata.Dispositivoid,
			&rowdata.Id,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.Fecha,
			&rowdata.ForeignId,
			&rowdata.ForeignText,
			&rowdata.Simbolo,
			&rowdata.Decimales,
			&rowdata.Compra,
			&rowdata.Venta,
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
func (u *TasaCambioE) GetByUniqueid(token string, uniqueid int) (*TasaCambioE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectTasaCambio

	var rowdata TasaCambioE
	jsonText := fmt.Sprintf(`{"uniqueid":%d}`, uniqueid)
	row := db.QueryRowContext(ctx, query, token, jsonText)

	err := row.Scan(
		&rowdata.DivisaId,
		&rowdata.Uniqueid,
		&rowdata.Owner,
		&rowdata.Dispositivoid,
		&rowdata.Id,
		&rowdata.Sede,
		&rowdata.Flag1,
		&rowdata.Flag2,
		&rowdata.Fecha,
		&rowdata.ForeignId,
		&rowdata.ForeignText,
		&rowdata.Simbolo,
		&rowdata.Decimales,
		&rowdata.Compra,
		&rowdata.Venta,
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

// GetOne returns one user by id
func (u *TasaCambioE) GetTasaByDivisa(token string, divisaid int, foerignid int) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM param_tasacambio_load ($1, $2, $3)`

	row := db.QueryRowContext(ctx, query, token, divisaid, foerignid)

	var rowdata TasaCambioE
	err := row.Scan(
		&rowdata.Fecha,
		&rowdata.Compra,
		&rowdata.Venta,
		&rowdata.Simbolo,
		&rowdata.Decimales,
	)

	if err != nil {
		return nil, err
	}

	retorno := make(map[string]any)
	retorno["fecha"] = rowdata.Fecha.Time
	retorno["compra"] = rowdata.Compra.Float64
	retorno["venta"] = rowdata.Venta.Float64
	retorno["simbolo"] = rowdata.Simbolo.String
	retorno["decimales"] = rowdata.Decimales.Int32

	return retorno, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *TasaCambioE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT param_tasacambio_save($1, $2, $3)`
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
func (u *TasaCambioE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT param_tasacambio_save($1, $2, $3)`
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
func (u *TasaCambioE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := `SELECT param_tasacambio_save($1, $2, $3)`
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
