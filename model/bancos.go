package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Bancos
type BancosE struct {
	Uniqueid        int64      `json:"uniqueid,omitempty"`
	Owner           NullInt32  `json:"owner,omitempty"`
	Dispositivoid   int32      `json:"dispositivoid,omitempty"`
	Id              int32      `json:"id,omitempty"`
	Sede            int32      `json:"sede"`
	Flag1           string     `json:"flag1,omitempty"`
	Flag2           string     `json:"flag2,omitempty"`
	Code            NullString `json:"code,omitempty"`
	Descrip         NullString `json:"descrip,omitempty"`
	NombreCorto     NullString `json:"nombrecorto,omitempty"`
	TipoEntidad     NullString `json:"tipoentidad,omitempty"`
	AhorrosSoles    NullInt32  `json:"ahorros_soles,omitempty"`
	AhorrosDolares  NullInt32  `json:"ahorros_dolares,omitempty"`
	CtacteSoles     NullInt32  `json:"ctacte_soles,omitempty"`
	CtacteDolares   NullInt32  `json:"ctacte_dolares,omitempty"`
	RegexCtaSoles   NullString `json:"regex_ctasoles,omitempty"`
	RegexCtaDolares NullString `json:"regex_ctadolares,omitempty"`
	LenCtaSoles     NullInt32  `json:"len_ctasoles,omitempty"`
	LenCtaDolares   NullInt32  `json:"len_ctadolares,omitempty"`
	Url             NullString `json:"url,omitempty"`
	Phone           NullString `json:"phone,omitempty"`
	City            NullString `json:"city,omitempty"`
	Ruf1            NullString `json:"ruf1,omitempty"`
	Ruf2            NullString `json:"ruf2,omitempty"`
	Ruf3            NullString `json:"ruf3,omitempty"`
	Iv              NullString `json:"iv,omitempty"`
	Salt            NullString `json:"salt,omitempty"`
	Checksum        NullString `json:"checksum,omitempty"`
	FCreated        NullTime   `json:"fcreated,omitempty"`
	FUpdated        NullTime   `json:"fupdated,omitempty"`
	Estadoreg       NullInt64  `json:"estadoreg,omitempty"`
	Activo          NullInt64  `json:"activo,omitempty"`
	TotalRecords    int64      `json:"total_records"`
}

func (e BancosE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectBancos = `select * from param_bancos_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *BancosE) GetAll(token string, filter string) ([]*BancosE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBancos

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

	var lista []*BancosE

	for rows.Next() {
		var rowdata BancosE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Owner,
			&rowdata.Dispositivoid,
			&rowdata.Id,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.Code,
			&rowdata.Descrip,
			&rowdata.NombreCorto,
			&rowdata.TipoEntidad,
			&rowdata.AhorrosSoles,
			&rowdata.AhorrosDolares,
			&rowdata.CtacteSoles,
			&rowdata.CtacteDolares,
			&rowdata.RegexCtaSoles,
			&rowdata.RegexCtaDolares,
			&rowdata.LenCtaSoles,
			&rowdata.LenCtaDolares,
			&rowdata.Url,
			&rowdata.Phone,
			&rowdata.City,
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
func (u *BancosE) GetByUniqueid(token string, uniqueid int) (*BancosE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBancos

	var rowdata BancosE
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
		&rowdata.Code,
		&rowdata.Descrip,
		&rowdata.NombreCorto,
		&rowdata.TipoEntidad,
		&rowdata.AhorrosSoles,
		&rowdata.AhorrosDolares,
		&rowdata.CtacteSoles,
		&rowdata.CtacteDolares,
		&rowdata.RegexCtaSoles,
		&rowdata.RegexCtaDolares,
		&rowdata.LenCtaSoles,
		&rowdata.LenCtaDolares,
		&rowdata.Url,
		&rowdata.Phone,
		&rowdata.City,
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
func (u *BancosE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT param_bancos_save($1, $2, $3)`
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
func (u *BancosE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT param_bancos_save($1, $2, $3)`
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
func (u *BancosE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := `SELECT param_bancos_save($1, $2, $3)`
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
