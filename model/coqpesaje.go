package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Pesaje
type CoqPesajeE struct {
	Uniqueid         int64       `json:"uniqueid,omitempty"`
	Owner            NullInt32   `json:"owner,omitempty"`
	Dispositivoid    NullInt32   `json:"dispositivoid,omitempty"`
	Id               int32       `json:"id,omitempty"`
	Sede             int32       `json:"sede"`
	Flag1            string      `json:"flag1,omitempty"`
	Flag2            string      `json:"flag2,omitempty"`
	CountryCode      NullString  `json:"countrycode,omitempty"`
	Campeonatoid     NullInt64   `json:"campeonato_id,omitempty"`
	FechaNumero      NullInt32   `json:"fecha_numero_id,omitempty"`
	FechaTorneoid    NullInt64   `json:"fecha_torneo_id,omitempty"`
	FechaGalponid    NullInt64   `json:"fecha_galpon_id,omitempty"`
	Peleaid          NullInt64   `json:"pelea_id,omitempty"`
	Galponid         NullInt64   `json:"galpon_id,omitempty"`
	Frenteid         NullInt64   `json:"frente_id,omitempty"`
	GalponText       NullString  `json:"galpon_text,omitempty"`
	Criadorid        NullInt64   `json:"criador_id,omitempty"`
	CriadorText      NullString  `json:"criador_text,omitempty"`
	Urbaniza         NullString  `json:"urbaniza,omitempty"`
	Galloid          NullInt64   `json:"gallo_id,omitempty"`
	PlacaPropietario NullString  `json:"placa_propietario,omitempty"`
	PlacaAsociacion  NullString  `json:"placa_asociacion,omitempty"`
	ColorGalloid     NullInt64   `json:"colo_gallo_id,omitempty"`
	ColorPicoid      NullInt64   `json:"color_pico_id,omitempty"`
	ColorPataid      NullInt64   `json:"color_pata_id,omitempty"`
	Peso             NullFloat64 `json:"peso,omitempty"`
	FechaRegistro    NullString  `json:"fecha_registro,omitempty"`
	Temp             NullString  `json:"temp,omitempty"`
	Invalidoind      NullInt32   `json:"invalido_ind,omitempty"`
	ExcluirRanking   NullInt32   `json:"excluir_ranking,omitempty"`
	Tardia           NullInt32   `json:"tardia,omitempty"`
	Libre            NullString  `json:"libre,omitempty"`
	Ruf1             NullString  `json:"ruf1,omitempty"`
	Ruf2             NullString  `json:"ruf2,omitempty"`
	Ruf3             NullString  `json:"ruf3,omitempty"`
	Iv               NullString  `json:"iv,omitempty"`
	Salt             NullString  `json:"salt,omitempty"`
	Checksum         NullString  `json:"checksum,omitempty"`
	FCreated         NullTime    `json:"fcreated,omitempty"`
	FUpdated         NullTime    `json:"fupdated,omitempty"`
	UCreated         NullString  `json:"ucreated,omitempty"`
	UUpdated         NullString  `json:"uupdated,omitempty"`
	Activo           int32       `json:"activo,omitempty"`
	Estadoreg        int32       `json:"estadoreg,omitempty"`
	TotalRecords     int64       `json:"total_records,omitempty"`
}

func (e CoqPesajeE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectCoqPesaje = `select * from coq_pesaje_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *CoqPesajeE) GetAll(token string, filter string) ([]*CoqPesajeE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectCoqPesaje

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

	var lista []*CoqPesajeE

	for rows.Next() {
		var rowdata CoqPesajeE
		err := rows.Scan(
			&rowdata.Uniqueid,
			//&rowdata.Owner,
			//&rowdata.Dispositivoid,
			&rowdata.Id,
			&rowdata.Sede,
			//&rowdata.Flag1,
			//&rowdata.Flag2,
			&rowdata.CountryCode,
			&rowdata.Campeonatoid,
			&rowdata.FechaNumero,
			&rowdata.FechaTorneoid,
			&rowdata.FechaGalponid,
			&rowdata.Peleaid,
			&rowdata.Galponid,
			&rowdata.Frenteid,
			&rowdata.GalponText,
			&rowdata.Criadorid,
			&rowdata.CriadorText,
			&rowdata.Galloid,
			&rowdata.PlacaPropietario,
			&rowdata.PlacaAsociacion,
			&rowdata.ColorGalloid,
			&rowdata.ColorPicoid,
			&rowdata.ColorPataid,
			&rowdata.Peso,
			&rowdata.FechaRegistro,
			&rowdata.Temp,
			&rowdata.Invalidoind,
			&rowdata.ExcluirRanking,
			&rowdata.Tardia,
			&rowdata.Libre,
			/*&rowdata.Ruf1,
			&rowdata.Ruf2,
			&rowdata.Ruf3,
			&rowdata.Iv,
			&rowdata.Salt,
			&rowdata.Checksum,*/
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
func (u *CoqPesajeE) GetByUniqueid(token string, uniqueid int) (*CoqPesajeE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectCoqPesaje

	var rowdata CoqPesajeE
	jsonText := fmt.Sprintf(`{"uniqueid":%d}`, uniqueid)
	row := db.QueryRowContext(ctx, query, token, jsonText)

	err := row.Scan(
		&rowdata.Uniqueid,
		//&rowdata.Owner,
		//&rowdata.Dispositivoid,
		&rowdata.Id,
		&rowdata.Sede,
		//&rowdata.Flag1,
		//&rowdata.Flag2,
		&rowdata.CountryCode,
		&rowdata.Campeonatoid,
		&rowdata.FechaNumero,
		&rowdata.FechaTorneoid,
		&rowdata.FechaGalponid,
		&rowdata.Peleaid,
		&rowdata.Galponid,
		&rowdata.Frenteid,
		&rowdata.GalponText,
		&rowdata.Criadorid,
		&rowdata.CriadorText,
		&rowdata.Galloid,
		&rowdata.PlacaPropietario,
		&rowdata.PlacaAsociacion,
		&rowdata.ColorGalloid,
		&rowdata.ColorPicoid,
		&rowdata.ColorPataid,
		&rowdata.Peso,
		&rowdata.FechaRegistro,
		&rowdata.Temp,
		&rowdata.Invalidoind,
		&rowdata.ExcluirRanking,
		&rowdata.Tardia,
		&rowdata.Libre,
		/*&rowdata.Ruf1,
		&rowdata.Ruf2,
		&rowdata.Ruf3,
		&rowdata.Iv,
		&rowdata.Salt,
		&rowdata.Checksum,*/
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

// GetInfo about one record
func (u *CoqPesajeE) GetInfo(token string, filter string) (*CoqPesajeE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from coq_pesaje_info( $1, $2)`

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

	var rowdata CoqPesajeE

	err = rows.Scan(
		&rowdata.Uniqueid,
		//&rowdata.Owner,
		//&rowdata.Dispositivoid,
		&rowdata.Id,
		&rowdata.Sede,
		//&rowdata.Flag1,
		//&rowdata.Flag2,
		&rowdata.CountryCode,
		&rowdata.Campeonatoid,
		&rowdata.FechaGalponid,
		&rowdata.FechaNumero,
		&rowdata.Peleaid,
		&rowdata.Galponid,
		&rowdata.Frenteid,
		&rowdata.GalponText,
		&rowdata.CriadorText,
		&rowdata.Urbaniza,
		&rowdata.Galloid,
		&rowdata.PlacaPropietario,
		&rowdata.PlacaAsociacion,
		&rowdata.ColorGalloid,
		&rowdata.ColorPicoid,
		&rowdata.ColorPataid,
		&rowdata.Peso,
		&rowdata.FechaRegistro,
		&rowdata.Temp,
		&rowdata.Invalidoind,
		&rowdata.ExcluirRanking,
		&rowdata.Tardia,
		&rowdata.Libre,
		/*&rowdata.Ruf1,
		&rowdata.Ruf2,
		&rowdata.Ruf3,
		&rowdata.Iv,
		&rowdata.Salt,
		&rowdata.Checksum,*/
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

// Validacion y registro de placa
func (u *CoqPesajeE) PlacaValidate(token string, filter string) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from coq_pesaje_placa_validate( $1, $2)`

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

	var jsonData []byte
	if err := rows.Scan(&jsonData); err != nil {
		log.Println("Error scanning", err)
		return nil, err
	}

	var infoResult map[string]interface{}
	json.Unmarshal(jsonData, &infoResult)

	return infoResult, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *CoqPesajeE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT coq_pesaje_save($1, $2, $3)`
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
func (u *CoqPesajeE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT coq_pesaje_delete($1, $2, $3)`
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
func (u *CoqPesajeE) DeleteByID(token string, id int, orden int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{  "id":%d, 
								"uniqueid":%d, 
								"orden":%d, 
								"estadoreg":%d
							  }`,
		id, id, orden, 300)

	query := `SELECT coq_pesaje_delete($1, $2, $3)`
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
