package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Galpon
type CoqGalponE struct {
	Uniqueid          int64      `json:"uniqueid,omitempty"`
	Owner             NullInt32  `json:"owner,omitempty"`
	Dispositivoid     NullInt32  `json:"dispositivoid,omitempty"`
	Id                int32      `json:"id,omitempty"`
	Sede              int32      `json:"sede"`
	Flag1             string     `json:"flag1,omitempty"`
	Flag2             string     `json:"flag2,omitempty"`
	CountryCode       NullString `json:"countrycode,omitempty"`
	Nombre            NullString `json:"nombre,omitempty"`
	Tokendataid       NullString `json:"tokendataid,omitempty"`
	PropietarioId     NullInt64  `json:"propietario_id,omitempty"`
	CriadorText       NullString `json:"criadortext,omitempty"`
	MovilCriador      NullString `json:"movilcriador,omitempty"`
	TorneoAnual       NullInt32  `json:"torneo_anual,omitempty"`
	GalponPadre       NullInt32  `json:"galpon_padre,omitempty"`
	Direccion         NullString `json:"direccion,omitempty"`
	Urbaniza          NullString `json:"urbaniza,omitempty"`
	Comodin           NullInt32  `json:"comodin,omitempty"`
	RegionId          NullInt64  `json:"regionid,omitempty"`
	RegionText        NullString `json:"regiontext,omitempty"`
	DepartamentoId    NullInt64  `json:"departamentoid,omitempty"`
	DepartamentoText  NullString `json:"departamentotext,omitempty"`
	ProvinciaId       NullInt64  `json:"provinciaid,omitempty"`
	ProvinciaText     NullString `json:"provinciatext,omitempty"`
	DistritoId        NullInt64  `json:"distritoid,omitempty"`
	DistritoText      NullString `json:"distritotext,omitempty"`
	CodigoPostal      NullString `json:"codigopostal,omitempty"`
	CodigoPostalText  NullString `json:"codigopostaltext,omitempty"`
	PaisId            NullInt64  `json:"paisid,omitempty"`
	PaisText          NullString `json:"paistext,omitempty"`
	Latitud           NullString `json:"latitud,omitempty"`
	Longitud          NullString `json:"longitud,omitempty"`
	NormalImageUrl    NullString `json:"normal_image_url,omitempty"`
	ThumbnailImageUrl NullString `json:"thumbnail_image_url,omitempty"`
	Ruf1              NullString `json:"ruf1,omitempty"`
	Ruf2              NullString `json:"ruf2,omitempty"`
	Ruf3              NullString `json:"ruf3,omitempty"`
	Iv                NullString `json:"iv,omitempty"`
	Salt              NullString `json:"salt,omitempty"`
	Checksum          NullString `json:"checksum,omitempty"`
	UCreated          NullString `json:"ucreated,omitempty"`
	UUpdated          NullString `json:"uupdated,omitempty"`
	FCreated          NullTime   `json:"fcreated,omitempty"`
	FUpdated          NullTime   `json:"fupdated,omitempty"`
	Activo            int32      `json:"activo,omitempty"`
	Estadoreg         int32      `json:"estadoreg,omitempty"`
	TotalRecords      int64      `json:"total_records"`
}

func (e CoqGalponE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectCoqGalpon = `select * from coq_galpon_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *CoqGalponE) GetAll(token string, filter string) ([]*CoqGalponE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectCoqGalpon

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

	var lista []*CoqGalponE

	for rows.Next() {
		var rowdata CoqGalponE
		err := rows.Scan(
			&rowdata.Uniqueid,
			//&rowdata.Owner,
			//&rowdata.Dispositivoid,
			&rowdata.Id,
			&rowdata.Sede,
			//&rowdata.Flag1,
			//&rowdata.Flag2,
			&rowdata.CountryCode,
			&rowdata.Nombre,
			&rowdata.Tokendataid,
			&rowdata.PropietarioId,
			&rowdata.CriadorText,
			&rowdata.MovilCriador,
			&rowdata.TorneoAnual,
			&rowdata.GalponPadre,
			&rowdata.Direccion,
			&rowdata.Urbaniza,
			&rowdata.Comodin,
			&rowdata.RegionId,
			&rowdata.RegionText,
			&rowdata.DepartamentoId,
			&rowdata.DepartamentoText,
			&rowdata.PropietarioId,
			&rowdata.ProvinciaText,
			&rowdata.DistritoId,
			&rowdata.DistritoText,
			&rowdata.CodigoPostal,
			&rowdata.CodigoPostalText,
			&rowdata.PaisId,
			&rowdata.PaisText,
			&rowdata.Latitud,
			&rowdata.Longitud,
			&rowdata.NormalImageUrl,
			&rowdata.ThumbnailImageUrl,
			&rowdata.Ruf1,
			&rowdata.Ruf2,
			&rowdata.Ruf3,
			&rowdata.Iv,
			&rowdata.Salt,
			&rowdata.Checksum,
			&rowdata.UCreated,
			&rowdata.UUpdated,
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
func (u *CoqGalponE) GetByUniqueid(token string, uniqueid int) (*CoqGalponE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectCoqGalpon

	var rowdata CoqGalponE
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
		&rowdata.Nombre,
		&rowdata.Tokendataid,
		&rowdata.PropietarioId,
		&rowdata.CriadorText,
		&rowdata.MovilCriador,
		&rowdata.TorneoAnual,
		&rowdata.GalponPadre,
		&rowdata.Direccion,
		&rowdata.Urbaniza,
		&rowdata.Comodin,
		&rowdata.RegionId,
		&rowdata.RegionText,
		&rowdata.DepartamentoId,
		&rowdata.DepartamentoText,
		&rowdata.PropietarioId,
		&rowdata.ProvinciaText,
		&rowdata.DistritoId,
		&rowdata.DistritoText,
		&rowdata.CodigoPostal,
		&rowdata.CodigoPostalText,
		&rowdata.PaisId,
		&rowdata.PaisText,
		&rowdata.Latitud,
		&rowdata.Longitud,
		&rowdata.NormalImageUrl,
		&rowdata.ThumbnailImageUrl,
		&rowdata.Ruf1,
		&rowdata.Ruf2,
		&rowdata.Ruf3,
		&rowdata.Iv,
		&rowdata.Salt,
		&rowdata.Checksum,
		&rowdata.UCreated,
		&rowdata.UUpdated,
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
func (u *CoqGalponE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT coq_galpon_save($1, $2, $3)`
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
func (u *CoqGalponE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT coq_galpon_save($1, $2, $3)`
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
func (u *CoqGalponE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"id":%d, 
							  "uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, id, 300)

	query := `SELECT coq_galpon_save($1, $2, $3)`
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
