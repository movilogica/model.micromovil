package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Tipos de autos
type CredencialE struct {
	Uniqueid          int64      `json:"uniqueid,omitempty"`
	Owner             NullInt32  `json:"owner,omitempty"`
	Dispositivoid     NullInt32  `json:"dispositivoid,omitempty"`
	Id                int32      `json:"id,omitempty"`
	Sede              int32      `json:"sede"`
	Flag1             string     `json:"flag1,omitempty"`
	Flag2             string     `json:"flag2,omitempty"`
	Nivel             NullInt32  `json:"nivel,omitempty"`
	Tipo              NullString `json:"tipo,omitempty"`
	Username          NullString `json:"username,omitempty"`
	Email             NullString `json:"email,omitempty"`
	Movil             NullString `json:"movil,omitempty"`
	Imei              NullString `json:"imei,omitempty"`
	Social            NullString `json:"social,omitempty"`
	Googleidtoken     NullString `json:"googleidtoken,omitempty"`
	Password          NullString `json:"password,omitempty"`
	ModoAutenticacion NullInt32  `json:"modoautenticacion,omitempty"`
	ModoAcceso        NullInt32  `json:"modoacceso,omitempty"`
	Changepwd         NullInt32  `json:"changepwd,omitempty"`
	NextChangepwdAt   NullTime   `json:"fnextchangepwd,omitempty"`
	LastAccessAt      NullTime   `json:"flastaccess,omitempty"`
	SedeActual        NullInt64  `json:"sedeactual,omitempty"`
	Welcome           NullInt32  `json:"welcome,omitempty"`
	WelcomeAt         NullTime   `json:"fwelcome,omitempty"`
	Notifier          NullInt32  `json:"notifier,omitempty"`
	NotifierAt        NullTime   `json:"fnotifier,omitempty"`
	Mailing           NullInt32  `json:"mailing,omitempty"`
	MailingAt         NullTime   `json:"fmailing,omitempty"`
	Pdf               NullInt32  `json:"pdf,omitempty"`
	PdfAt             NullTime   `json:"fpdf,omitempty"`
	Idioma            NullString `json:"idioma,omitempty"`
	Avatar            NullString `json:"avatar,omitempty"`
	TokenTerminal     NullString `json:"tokenterminal,omitempty"`
	Manual            NullInt32  `json:"manual,omitempty"`
	HasRole           NullInt32  `json:"hasrole,omitempty"`
	Ruf1              NullString `json:"ruf1,omitempty"`
	Ruf2              NullString `json:"ruf2,omitempty"`
	Ruf3              NullString `json:"ruf3,omitempty"`
	Iv                NullString `json:"iv,omitempty"`
	Salt              NullString `json:"salt,omitempty"`
	Checksum          NullString `json:"checksum,omitempty"`
	FCreated          NullTime   `json:"fcreated,omitempty"`
	FUpdated          NullTime   `json:"fupdated,omitempty"`
	Activo            int32      `json:"activo,omitempty"`
	Estadoreg         int32      `json:"estadoreg,omitempty"`
	TotalRecords      int64      `json:"total_records,omitempty"`
}

func (e CredencialE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectCredencial = `select * from security_credenciales_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *CredencialE) GetAll(token string, filter string) ([]*CredencialE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectCredencial

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

	var lista []*CredencialE

	for rows.Next() {
		var rowdata CredencialE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Owner,
			&rowdata.Dispositivoid,
			&rowdata.Id,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.Nivel,
			&rowdata.Tipo,
			&rowdata.Username,
			&rowdata.Email,
			&rowdata.Movil,
			&rowdata.Imei,
			&rowdata.Social,
			&rowdata.Googleidtoken,
			&rowdata.Password,
			&rowdata.ModoAutenticacion,
			&rowdata.ModoAcceso,
			&rowdata.Changepwd,
			&rowdata.NextChangepwdAt,
			&rowdata.LastAccessAt,
			&rowdata.SedeActual,
			&rowdata.Welcome,
			&rowdata.WelcomeAt,
			&rowdata.Notifier,
			&rowdata.NotifierAt,
			&rowdata.Mailing,
			&rowdata.MailingAt,
			&rowdata.Pdf,
			&rowdata.PdfAt,
			&rowdata.Idioma,
			&rowdata.Avatar,
			&rowdata.TokenTerminal,
			&rowdata.Manual,
			&rowdata.HasRole,
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

// GetByField returns one record by filter
/*func (u *SedeE) GetByField(fieldname string, value string, token string) (*SedeE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select uniqueid, owner, dispositivoid, id, sede, flag1, flag2,
			         code, descrip,
					 ruf1, ruf2, ruf3, fcreated, fupdated, estadoreg, activo, total_records
			    from sys_sedes_list( $1, $2)`

	rows, err := db.QueryContext(ctx, query, token, `{"tipo":"TYPE_OF_CARS"}`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result *Parameters

	if rows.Next() {
		var rowdata Parameters
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.Tipo,
			&rowdata.Secuencial,
			&rowdata.Code,
			&rowdata.Descrip,
			&rowdata.CreatedAt,
			&rowdata.UpdatedAt,
			&rowdata.Activo,
			&rowdata.Estadoreg,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		result = &rowdata
	}

	return result, nil
}*/

// GetOne returns one user by id
func (u *CredencialE) GetByUniqueid(token string, uniqueid int) (*CredencialE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectCredencial

	var rowdata CredencialE
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
		&rowdata.Nivel,
		&rowdata.Tipo,
		&rowdata.Username,
		&rowdata.Email,
		&rowdata.Movil,
		&rowdata.Imei,
		&rowdata.Social,
		&rowdata.Googleidtoken,
		&rowdata.Password,
		&rowdata.ModoAutenticacion,
		&rowdata.ModoAcceso,
		&rowdata.Changepwd,
		&rowdata.NextChangepwdAt,
		&rowdata.LastAccessAt,
		&rowdata.SedeActual,
		&rowdata.Welcome,
		&rowdata.WelcomeAt,
		&rowdata.Notifier,
		&rowdata.NotifierAt,
		&rowdata.Mailing,
		&rowdata.MailingAt,
		&rowdata.Pdf,
		&rowdata.PdfAt,
		&rowdata.Idioma,
		&rowdata.Avatar,
		&rowdata.TokenTerminal,
		&rowdata.Manual,
		&rowdata.HasRole,
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
func (u *CredencialE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT security_credenciales_save($1, $2, $3)`
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
func (u *CredencialE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT security_credenciales_save($1, $2, $3)`
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
func (u *CredencialE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := `SELECT security_credenciales_save($1, $2, $3)`
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

func (u *CredencialE) ExistsUserByEmail(token string, jsonData string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "select 1 from security_credenciales_list($1, $2)"

	///jsonText := fmt.Sprintf(`{"email":"%s"}`, email)
	row := db.QueryRowContext(ctx, query, token, jsonData)

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

func (u *CredencialE) ExistsUserByPhone(token string, jsonData string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "select 1 from security_credenciales_list($1, $2)"

	///jsonText := fmt.Sprintf(`{"movil":"%s"}`, phone)
	row := db.QueryRowContext(ctx, query, token, jsonData)

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
func (u *CredencialE) GetUserByEmailOrPhone(token string, jsonData string) (*CredencialE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectCredencial

	var rowdata CredencialE
	///jsonText := fmt.Sprintf(`{"email":%s}`, email)
	row := db.QueryRowContext(ctx, query, token, jsonData)

	err := row.Scan(
		&rowdata.Uniqueid,
		&rowdata.Owner,
		&rowdata.Dispositivoid,
		&rowdata.Id,
		&rowdata.Sede,
		&rowdata.Flag1,
		&rowdata.Flag2,
		&rowdata.Nivel,
		&rowdata.Tipo,
		&rowdata.Username,
		&rowdata.Email,
		&rowdata.Movil,
		&rowdata.Imei,
		&rowdata.Social,
		&rowdata.Googleidtoken,
		&rowdata.Password,
		&rowdata.ModoAutenticacion,
		&rowdata.ModoAcceso,
		&rowdata.Changepwd,
		&rowdata.NextChangepwdAt,
		&rowdata.LastAccessAt,
		&rowdata.SedeActual,
		&rowdata.Welcome,
		&rowdata.WelcomeAt,
		&rowdata.Notifier,
		&rowdata.NotifierAt,
		&rowdata.Mailing,
		&rowdata.MailingAt,
		&rowdata.Pdf,
		&rowdata.PdfAt,
		&rowdata.Idioma,
		&rowdata.Avatar,
		&rowdata.TokenTerminal,
		&rowdata.Manual,
		&rowdata.HasRole,
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

func (u *CredencialE) Save(token string, jsonData string, metricas string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "select * from security_credenciales_save($1, $2, $3)"

	///jsonText := fmt.Sprintf(`{"email":"%s"}`, email)
	row := db.QueryRowContext(ctx, query, token, jsonData, metricas)

	var uniqueid NullInt64

	err := row.Scan(&uniqueid)

	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return uniqueid.Int64, nil
}

/*****
 * Este procedimiento requiere la generacion de multiples tablas transaccionales.-
 *********/
func (u *CredencialE) RegisterAccount(token string, data string, auth string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Printf("RegisterAccount [data]=%s [auth]=%s [metricas]=%s\n", data, auth, metricas)

	query := `CALL auth_register_user($1, $2, $3, $4, $5)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	// _, err := db.ExecContext(ctx, "ProcName", sql.Named("Arg1", sql.Out{Dest: &outArg}))

	var tokensession string
	_, err = stmt.ExecContext(ctx, token, data, auth, metricas, &tokensession)
	//_, err = stmt.ExecContext(ctx, token, data, auth, metricas, sql.Out{Dest: &tokensession})
	if err != nil {
		return nil, err
	}
	///defer result.Close()

	retorno := make(map[string]any)
	retorno["tokensession"] = tokensession

	log.Printf("RegisterAccount [token]=%s\n", tokensession)

	return retorno, nil
}

func (u *CredencialE) ToText() string {
	jsonCredencial, _ := json.Marshal(u)
	return string(jsonCredencial)
}

func (u *CredencialE) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password.String), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func (e CredencialE) ModoAutenticacionText() string {
	switch e.ModoAutenticacion.Int32 {
	case 1:
		return "By IMEI"
	case 2:
		return "By User & Password"
	case 3:
		return "By IMEI & Password"
	case 4:
		return "By IMEI & User & Password"
	default:
		return "By Default"
	}
}

func (e CredencialE) ModoAccesoText() string {
	switch e.ModoAcceso.Int32 {
	case 1:
		return "Web only"
	case 2:
		return "App only"
	case 9:
		return "Not matter"
	default:
		return "By Default"
	}
}
