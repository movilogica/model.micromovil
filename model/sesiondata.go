package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
)

// User is the structure which holds one user from the database.
type SesionDataE struct {
	Uniqueid       int64      `json:"uniqueid"`
	Owner          NullInt32  `json:"owner,omitempty"`
	Dispositivoid  NullInt32  `json:"dispositivoid,omitempty"`
	Id             int32      `json:"id,omitempty"`
	Sede           int32      `json:"sede"`
	Flag1          string     `json:"flag1,omitempty"`
	Flag2          string     `json:"flag2,omitempty"`
	CredencialId   NullInt64  `json:"credencialid,omitempty"`
	PersonaId      NullInt64  `json:"personaid,omitempty"`
	Username       NullString `json:"username,omitempty"`
	TokenId        NullString `json:"tokenid,omitempty"`
	JsessionId     NullString `json:"jsessionid,omitempty"`
	Remotemachine  NullString `json:"remotemachine,omitempty"`
	Remotehost     NullString `json:"remotehost,omitempty"`
	Remoteport     NullInt32  `json:"remoteport,omitempty"`
	Headerdata     NullString `json:"headerdata,omitempty"`
	Headerchecksum NullString `json:"headerchecksum,omitempty"`
	LastActivityAt NullTime   `json:"lastactivity,omitempty"`
	Version        NullString `json:"version,omitempty"`
	Ruf1           NullString `json:"ruf1,omitempty"`
	Ruf2           NullString `json:"ruf2,omitempty"`
	Ruf3           NullString `json:"ruf3,omitempty"`
	Iv             NullString `json:"iv,omitempty"`
	Salt           NullString `json:"salt,omitempty"`
	Checksum       NullString `json:"checksum,omitempty"`
	FCreated       NullTime   `json:"fcreated,omitempty"`
	FUpdated       NullTime   `json:"fupdated,omitempty"`
	Estadoreg      NullInt64  `json:"estadoreg,omitempty"`
	Activo         NullInt64  `json:"activo,omitempty"`
	TotalRecords   NullInt64  `json:"total_records,omitempty"`
}

const querySelectSesion = `select * from security_sesiondata_list( $1, $2)`

func (e SesionDataE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

func (u *SesionDataE) ExistsSessionByToken(token string, jsonData string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "select 1 from security_sesiondata_list($1, $2)"

	// Se realiza la verificacion de la existencia de token
	var mapData map[string]any
	json.Unmarshal([]byte(jsonData), &mapData)
	if mapData == nil {
		mapData = make(map[string]any)
	}
	if _, ok := mapData["tokenid"]; !ok {
		if len(token) > 0 {
			mapData["tokenid"] = token
		} else {
			return 0, errors.New("You must be specified the token value")
		}
	}
	jsonBytes, _ := json.Marshal(mapData)

	///jsonText := fmt.Sprintf(`{"email":"%s"}`, email)
	row := db.QueryRowContext(ctx, query, token, string(jsonBytes))

	var exists NullInt32

	err := row.Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	if exists.Int32 > 0 {
		return 1, nil
	} else {
		return 0, nil
	}
}

// GetOne returns one user by id
func (u *SesionDataE) GetSessionByToken(token string, jsonData string) (SesionDataE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectSesion

	var rowdata SesionDataE

	// Se realiza la verificacion de la existencia de token
	var mapData map[string]any
	json.Unmarshal([]byte(jsonData), &mapData)
	if mapData == nil {
		mapData = make(map[string]any)
	}
	if _, ok := mapData["tokenid"]; !ok {
		if len(token) > 0 {
			mapData["tokenid"] = token
		} else {
			return rowdata, errors.New("You must be specified the token value")
		}
	}
	jsonBytes, _ := json.Marshal(mapData)

	///jsonText := fmt.Sprintf(`{"email":%s}`, email)
	row := db.QueryRowContext(ctx, query, "", string(jsonBytes))

	err := row.Scan(
		&rowdata.Uniqueid,
		&rowdata.Owner,
		&rowdata.Dispositivoid,
		&rowdata.Id,
		&rowdata.Sede,
		&rowdata.Flag1,
		&rowdata.Flag2,
		&rowdata.CredencialId,
		&rowdata.PersonaId,
		&rowdata.Username,
		&rowdata.TokenId,
		&rowdata.JsessionId,
		&rowdata.Remotemachine,
		&rowdata.Remotehost,
		&rowdata.Remoteport,
		&rowdata.Headerdata,
		&rowdata.Headerchecksum,
		&rowdata.LastActivityAt,
		&rowdata.Version,
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
		return rowdata, err
	}

	return rowdata, nil
}

// GetOne returns one user by id
func (u *SesionDataE) GetSessionByValue(token string, jsonData string) (SesionDataE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectSesion

	var rowdata SesionDataE

	// Se realiza la verificacion de la existencia de token
	var mapData map[string]any
	json.Unmarshal([]byte(jsonData), &mapData)
	if mapData == nil {
		return rowdata, errors.New("You must be specified any filter")
	}
	jsonBytes, _ := json.Marshal(mapData)
	log.Println(token, string(jsonBytes))
	row := db.QueryRowContext(ctx, query, token, jsonBytes)

	err := row.Scan(
		&rowdata.Uniqueid,
		&rowdata.Owner,
		&rowdata.Dispositivoid,
		&rowdata.Id,
		&rowdata.Sede,
		&rowdata.Flag1,
		&rowdata.Flag2,
		&rowdata.CredencialId,
		&rowdata.PersonaId,
		&rowdata.Username,
		&rowdata.TokenId,
		&rowdata.JsessionId,
		&rowdata.Remotemachine,
		&rowdata.Remotehost,
		&rowdata.Remoteport,
		&rowdata.Headerdata,
		&rowdata.Headerchecksum,
		&rowdata.LastActivityAt,
		&rowdata.Version,
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
		return rowdata, err
	}

	return rowdata, nil
}

func (u *SesionDataE) Save(token string, jsonData string, metricas string) (int64, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "select * from security_sesiondata_save($1, $2, $3)"

	///jsonText := fmt.Sprintf(`{"email":"%s"}`, email)
	row := db.QueryRowContext(ctx, query, token, jsonData, metricas)

	var uniqueid NullInt64
	var tokenid NullString

	err := row.Scan(&uniqueid, &tokenid)

	if err != nil && err != sql.ErrNoRows {
		return 0, "", err
	}
	return uniqueid.Int64, tokenid.String, nil
}
