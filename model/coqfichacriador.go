package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Ficha Criador
type CoqFichaCriadorE struct {
	Uniqueid        int64      `json:"uniqueid,omitempty"`
	Owner           NullInt64  `json:"ownerid,omitempty"`
	Dispositivoid   NullInt64  `json:"dispositivoid,omitempty"`
	Id              int64      `json:"id,omitempty"`
	Sede            int64      `json:"sede"`
	Flag1           string     `json:"flag1,omitempty"`
	Flag2           string     `json:"flag2,omitempty"`
	CountryCode     NullString `json:"countrycode,omitempty"`
	PersonaId       NullInt64  `json:"personaid,omitempty"`
	Tokendataid     NullString `json:"tokendataid,omitempty"`
	Nombre          NullString `json:"nombre,omitempty"`
	Apellido        NullString `json:"apellido,omitempty"`
	ApellidoMaterno NullString `json:"apellido_materno,omitempty"`
	GalponText      NullString `json:"galpon_text,omitempty"`
	TipodocId       NullInt64  `json:"tipodoc_id,omitempty"`
	Nrodoc          NullString `json:"nrodoc,omitempty"`
	Dni             NullString `json:"dni,omitempty"`
	Direccion       NullString `json:"direccion,omitempty"`
	Urbaniza        NullString `json:"urbaniza,omitempty"`
	Email1          NullString `json:"email1,omitempty"`
	Email2          NullString `json:"email2,omitempty"`
	Celular1        NullString `json:"celular1,omitempty"`
	Celular2        NullString `json:"celular2,omitempty"`
	Celular3        NullString `json:"celular3,omitempty"`
	FInscripcion    NullTime   `json:"fecha_inscripcion,omitempty"`
	FBirthdate      NullTime   `json:"fecha_nacimiento,omitempty"`
	SocioCriador    NullInt32  `json:"socio_criador,omitempty"`
	FirmaReglPelea  NullInt32  `json:"firma_regl_pelea,omitempty"`
	FirmaCodEtica   NullInt32  `json:"firma_cod_etica,omitempty"`
	FirmaReglTorneo NullInt32  `json:"firma_regl_campeona,omitempty"`
	Ruf1            NullString `json:"ruf1,omitempty"`
	Ruf2            NullString `json:"ruf2,omitempty"`
	Ruf3            NullString `json:"ruf3,omitempty"`
	Iv              NullString `json:"iv,omitempty"`
	Salt            NullString `json:"salt,omitempty"`
	Checksum        NullString `json:"checksum,omitempty"`
	FCreated        NullTime   `json:"fcreated,omitempty"`
	FUpdated        NullTime   `json:"fupdated,omitempty"`
	Activo          int32      `json:"activo,omitempty"`
	Estadoreg       int32      `json:"estadoreg,omitempty"`
	TotalRecords    int64      `json:"total_records,omitempty"`
}

func (e CoqFichaCriadorE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectCoqFichaCriador = `select * from coq_ficha_criador_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *CoqFichaCriadorE) GetAll(token string, filter string) ([]*CoqFichaCriadorE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectCoqFichaCriador

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

	var lista []*CoqFichaCriadorE

	for rows.Next() {
		var rowdata CoqFichaCriadorE
		err := rows.Scan(
			&rowdata.Uniqueid,
			//&rowdata.Owner,
			//&rowdata.Dispositivoid,
			&rowdata.Id,
			&rowdata.Sede,
			//&rowdata.Flag1,
			//&rowdata.Flag2,
			&rowdata.CountryCode,
			&rowdata.PersonaId,
			&rowdata.Tokendataid,
			&rowdata.Nombre,
			&rowdata.Apellido,
			&rowdata.ApellidoMaterno,
			&rowdata.GalponText,
			&rowdata.TipodocId,
			&rowdata.Nrodoc,
			&rowdata.Dni,
			&rowdata.Direccion,
			&rowdata.Urbaniza,
			&rowdata.Email1,
			&rowdata.Email2,
			&rowdata.Celular1,
			&rowdata.Celular2,
			&rowdata.Celular3,
			&rowdata.FInscripcion,
			&rowdata.FBirthdate,
			&rowdata.SocioCriador,
			&rowdata.FirmaReglPelea,
			&rowdata.FirmaCodEtica,
			&rowdata.FirmaReglTorneo,
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
func (u *CoqFichaCriadorE) GetByUniqueid(token string, uniqueid int) (*CoqFichaCriadorE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectCoqFichaCriador

	var rowdata CoqFichaCriadorE
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
		&rowdata.PersonaId,
		&rowdata.Tokendataid,
		&rowdata.Nombre,
		&rowdata.Apellido,
		&rowdata.ApellidoMaterno,
		&rowdata.GalponText,
		&rowdata.TipodocId,
		&rowdata.Nrodoc,
		&rowdata.Dni,
		&rowdata.Direccion,
		&rowdata.Urbaniza,
		&rowdata.Email1,
		&rowdata.Email2,
		&rowdata.Celular1,
		&rowdata.Celular2,
		&rowdata.Celular3,
		&rowdata.FInscripcion,
		&rowdata.FBirthdate,
		&rowdata.SocioCriador,
		&rowdata.FirmaReglPelea,
		&rowdata.FirmaCodEtica,
		&rowdata.FirmaReglTorneo,
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

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *CoqFichaCriadorE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT coq_ficha_criador_save($1, $2, $3)`
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
func (u *CoqFichaCriadorE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT coq_ficha_criador_save($1, $2, $3)`
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
func (u *CoqFichaCriadorE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"id":%d, 
							  "uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, id, 300)

	query := `SELECT coq_ficha_criador_save($1, $2, $3)`
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
