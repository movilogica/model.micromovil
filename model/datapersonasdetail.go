package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Personas
type DataPersonaDetailE struct {
	Uniqueid          int64      `json:"uniqueid,omitempty"`
	Owner             NullInt32  `json:"owner,omitempty"`
	Dispositivoid     NullInt32  `json:"dispositivoid,omitempty"`
	Id                int32      `json:"id,omitempty"`
	Sede              int32      `json:"sede"`
	Flag1             string     `json:"flag1,omitempty"`
	Flag2             string     `json:"flag2,omitempty"`
	PersonaId         NullInt64  `json:"personaid,omitempty"`
	Genero            NullString `json:"genero,omitempty"`
	FechaNacAt        NullTime   `json:"fechanac,omitempty"`
	LugarNac          NullString `json:"lugarnac,omitempty"`
	Nacionalidad      NullString `json:"nacionalidad,omitempty"`
	EstadoCivil       NullString `json:"estadocivil,omitempty"`
	DivisaId          NullInt64  `json:"divisaid,omitempty"`
	DivisaText        NullString `json:"divisatext,omitempty"`
	DivisaSimbolo     NullString `json:"divisasimbolo,omitempty"`
	DivisaDecimal     NullInt32  `json:"divisadecimal,omitempty"`
	GiroProfesionId   NullInt64  `json:"giroprofesionid,omitempty"`
	GiroProfesionText NullString `json:"giroprofesiontext,omitempty"`
	Ciiu              NullString `json:"ciiu,omitempty"`
	CiiuText          NullString `json:"ciiutext,omitempty"`
	GradoInstr        NullString `json:"gradoinstr,omitempty"`
	SitGradoInstr     NullString `json:"situaciongradoinstr,omitempty"`
	SitLaboral        NullString `json:"situacionlab,omitempty"`
	NombrePadre       NullString `json:"nombrepadre,omitempty"`
	NombreMadre       NullString `json:"nombremadre,omitempty"`
	Ruf1              NullString `json:"ruf1,omitempty"`
	Ruf2              NullString `json:"ruf2,omitempty"`
	Ruf3              NullString `json:"ruf3,omitempty"`
	Iv                NullString `json:"iv,omitempty"`
	Salt              NullString `json:"salt,omitempty"`
	Checksum          NullString `json:"checksum,omitempty"`
	FCreated          NullTime   `json:"fcreated,omitempty"`
	FUpdated          NullTime   `json:"fupdated,omitempty"`
	UCreated          NullString `json:"ucreated,omitempty"`
	UUpdated          NullString `json:"uupdated,omitempty"`
	Activo            int32      `json:"activo,omitempty"`
	Estadoreg         int32      `json:"estadoreg,omitempty"`
	TotalRecords      int64      `json:"total_records,omitempty"`
}

func (e DataPersonaDetailE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectDataPerDetail = `select * from data_personas_detail_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *DataPersonaDetailE) GetAll(token string, filter string) ([]*DataPersonaDetailE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectDataPerDetail

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

	var lista []*DataPersonaDetailE

	for rows.Next() {
		var rowdata DataPersonaDetailE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Owner,
			&rowdata.Dispositivoid,
			&rowdata.Id,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.PersonaId,
			&rowdata.Genero,
			&rowdata.FechaNacAt,
			&rowdata.LugarNac,
			&rowdata.Nacionalidad,
			&rowdata.EstadoCivil,
			&rowdata.DivisaId,
			&rowdata.DivisaText,
			&rowdata.DivisaSimbolo,
			&rowdata.DivisaDecimal,
			&rowdata.GiroProfesionId,
			&rowdata.GiroProfesionText,
			&rowdata.Ciiu,
			&rowdata.CiiuText,
			&rowdata.GradoInstr,
			&rowdata.SitGradoInstr,
			&rowdata.SitLaboral,
			&rowdata.NombrePadre,
			&rowdata.NombreMadre,
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
func (u *DataPersonaDetailE) GetByUniqueid(token string, uniqueid int) (*DataPersonaDetailE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectDataPerDetail

	var rowdata DataPersonaDetailE
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
		&rowdata.Genero,
		&rowdata.FechaNacAt,
		&rowdata.LugarNac,
		&rowdata.Nacionalidad,
		&rowdata.EstadoCivil,
		&rowdata.DivisaId,
		&rowdata.DivisaText,
		&rowdata.DivisaSimbolo,
		&rowdata.DivisaDecimal,
		&rowdata.GiroProfesionId,
		&rowdata.GiroProfesionText,
		&rowdata.Ciiu,
		&rowdata.CiiuText,
		&rowdata.GradoInstr,
		&rowdata.SitGradoInstr,
		&rowdata.SitLaboral,
		&rowdata.NombrePadre,
		&rowdata.NombreMadre,
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
func (u *DataPersonaDetailE) GetByPersonaid(token string, uniqueid int) (*DataPersonaDetailE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectDataPerDetail

	var rowdata DataPersonaDetailE
	jsonText := fmt.Sprintf(`{"personaid":%d}`, uniqueid)
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
		&rowdata.Genero,
		&rowdata.FechaNacAt,
		&rowdata.LugarNac,
		&rowdata.Nacionalidad,
		&rowdata.EstadoCivil,
		&rowdata.DivisaId,
		&rowdata.DivisaText,
		&rowdata.DivisaSimbolo,
		&rowdata.DivisaDecimal,
		&rowdata.GiroProfesionId,
		&rowdata.GiroProfesionText,
		&rowdata.Ciiu,
		&rowdata.CiiuText,
		&rowdata.GradoInstr,
		&rowdata.SitGradoInstr,
		&rowdata.SitLaboral,
		&rowdata.NombrePadre,
		&rowdata.NombreMadre,
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
func (u *DataPersonaDetailE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT data_personas_detail_save($1, $2, $3)`
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
func (u *DataPersonaDetailE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT data_personas_detail_save($1, $2, $3)`
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
func (u *DataPersonaDetailE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := `SELECT data_personas_detail_save($1, $2, $3)`
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
