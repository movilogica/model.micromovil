package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Tipos de autos
type UsuariosRolesE struct {
	Uniqueid      int64      `json:"uniqueid,omitempty"`
	Owner         NullInt32  `json:"owner,omitempty"`
	Dispositivoid int32      `json:"dispositivoid,omitempty"`
	Id            int32      `json:"id,omitempty"`
	Sede          int32      `json:"sede"`
	Flag1         string     `json:"flag1,omitempty"`
	Flag2         string     `json:"flag2,omitempty"`
	CredencialId  NullInt64  `json:"credencialid,omitempty"`
	RolId         NullInt64  `json:"rolid,omitempty"`
	RolCode       NullString `json:"rolcode,omitempty"`
	RolDescrip    NullString `json:"roldescrip,omitempty"`
	Profiles      NullString `json:"profiles,omitempty"`
	Ruf1          NullString `json:"ruf1,omitempty"`
	Ruf2          NullString `json:"ruf2,omitempty"`
	Ruf3          NullString `json:"ruf3,omitempty"`
	Iv            NullString `json:"iv,omitempty"`
	Salt          NullString `json:"salt,omitempty"`
	Checksum      NullString `json:"checksum,omitempty"`
	FCreated      NullTime   `json:"fcreated,omitempty"`
	FUpdated      NullTime   `json:"fupdated,omitempty"`
	Estadoreg     NullInt64  `json:"estadoreg,omitempty"`
	Activo        NullInt64  `json:"activo,omitempty"`
	TotalRecords  int64      `json:"total_records"`
}

func (e UsuariosRolesE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectUsuariosRoles = `select * from security_usuarios_roles_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *UsuariosRolesE) GetAll(token string, filter string) ([]*UsuariosRolesE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectUsuariosRoles

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

	var lista []*UsuariosRolesE

	for rows.Next() {
		var rowdata UsuariosRolesE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Owner,
			&rowdata.Dispositivoid,
			&rowdata.Id,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.CredencialId,
			&rowdata.RolId,
			&rowdata.RolCode,
			&rowdata.RolDescrip,
			&rowdata.Profiles,
			&rowdata.Ruf1,
			&rowdata.Ruf2,
			&rowdata.Ruf3,
			&rowdata.Iv,
			&rowdata.Salt,
			&rowdata.Checksum,
			&rowdata.FCreated,
			&rowdata.FUpdated,
			&rowdata.Estadoreg,
			&rowdata.Activo,
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
func (u *UsuariosRolesE) GetByUniqueid(token string, uniqueid int) (*UsuariosRolesE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectUsuariosRoles

	var rowdata UsuariosRolesE
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
		&rowdata.CredencialId,
		&rowdata.RolId,
		&rowdata.RolCode,
		&rowdata.RolDescrip,
		&rowdata.Profiles,
		&rowdata.Ruf1,
		&rowdata.Ruf2,
		&rowdata.Ruf3,
		&rowdata.Iv,
		&rowdata.Salt,
		&rowdata.Checksum,
		&rowdata.FCreated,
		&rowdata.FUpdated,
		&rowdata.Estadoreg,
		&rowdata.Activo,
		&rowdata.TotalRecords,
	)

	if err != nil {
		return nil, err
	}

	return &rowdata, nil
}
