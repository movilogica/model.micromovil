package model

import (
	"context"
	"fmt"
)

// User is the structure which holds one user from the database.
type UsuarioE struct {
	Merchantid       NullString `json:"merchantid"`
	Sede             int32      `json:"sede"`
	Flag1            string     `json:"flag1,omitempty"`
	Flag2            string     `json:"flag2,omitempty"`
	Tokendataid      NullString `json:"tokendataid,omitempty"`
	Nroperacion      NullString `json:"nroperacion"`
	Nickname         NullString `json:"nickname,omitempty"`
	Username         NullString `json:"username,omitempty"`
	Email            NullString `json:"email,omitempty"`
	Movil            NullString `json:"movil,omitempty"`
	Shortname        NullString `json:"shortname,omitempty"`
	Nombres          NullString `json:"nombres,omitempty"`
	Apaterno         NullString `json:"apaterno,omitempty"`
	Amaterno         NullString `json:"amaterno,omitempty"`
	Girotext         NullString `json:"girotext,omitempty"`
	Departamentotext NullString `json:"departamentotext,omitempty"`
	Country_iso_2    NullString `json:"country_iso_2,omitempty"`
	Country_iso_3    NullString `json:"country_iso_3,omitempty"`
	Country_iso_m49  NullString `json:"country_iso_m49,omitempty"`
	Country_prefijo  NullString `json:"country_prefijo,omitempty"`
	Paistext         NullString `json:"paistext,omitempty"`
	Avatar           NullString `json:"avatar,omitempty"`
	Idioma           NullString `json:"idioma,omitempty"`
	LastAccessAt     NullTime   `json:"flastaccess,omitempty"`
	FCreated         NullTime   `json:"fcreated,omitempty"`
	FUpdated         NullTime   `json:"fupdated,omitempty"`
	Activo           int32      `json:"activo,omitempty"`
	Estadoreg        int32      `json:"estadoreg,omitempty"`
	TotalRecords     int64      `json:"total_records"`
}

func (e UsuarioE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectUserInfo = `select * from usuario_info( $1, $2)`
const querySelectUserLabel = `select * from usuario_label( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetOne returns one user by id
func (u *UsuarioE) GetUserInfo(token string, uniqueid int) (*UsuarioE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectUserInfo

	var rowdata UsuarioE
	jsonText := fmt.Sprintf(`{"uniqueid":%d}`, uniqueid)
	row := db.QueryRowContext(ctx, query, token, jsonText)

	err := row.Scan(
		&rowdata.Merchantid,
		&rowdata.Sede,
		&rowdata.Flag1,
		&rowdata.Flag2,
		&rowdata.Tokendataid,
		&rowdata.Nroperacion,
		&rowdata.Nickname,
		&rowdata.Username,
		&rowdata.Email,
		&rowdata.Movil,
		&rowdata.Shortname,
		&rowdata.Nombres,
		&rowdata.Apaterno,
		&rowdata.Amaterno,
		&rowdata.Girotext,
		&rowdata.Departamentotext,
		&rowdata.Country_iso_2,
		&rowdata.Country_iso_3,
		&rowdata.Country_iso_m49,
		&rowdata.Country_prefijo,
		&rowdata.Paistext,
		&rowdata.Avatar,
		&rowdata.Idioma,
		&rowdata.LastAccessAt,
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

func (u *UsuarioE) GetUserLabel(token string, uniqueid int) (*UsuarioE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectUserLabel

	var rowdata UsuarioE
	jsonText := fmt.Sprintf(`{"uniqueid":%d}`, uniqueid)
	row := db.QueryRowContext(ctx, query, token, jsonText)

	err := row.Scan(
		&rowdata.Merchantid,
		&rowdata.Sede,
		&rowdata.Flag1,
		&rowdata.Flag2,
		&rowdata.Tokendataid,
		&rowdata.Nroperacion,
		&rowdata.Nickname,
		&rowdata.Shortname,
		&rowdata.Departamentotext,
		&rowdata.Country_iso_2,
		&rowdata.Country_iso_3,
		&rowdata.Country_iso_m49,
		&rowdata.Country_prefijo,
		&rowdata.Paistext,
		&rowdata.Avatar,
		&rowdata.Idioma,
		&rowdata.LastAccessAt,
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

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
/*func (u *User) Insert(user User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return 0, err
	}

	var newID int
	stmt := `insert into users (email, first_name, last_name, password, user_active, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7) returning id`

	err = db.QueryRowContext(ctx, stmt,
		user.Email,
		user.FirstName,
		user.LastName,
		hashedPassword,
		user.Active,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// ResetPassword is the method we will use to change a user's password.
func (u *User) ResetPassword(password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `update users set password = $1 where id = $2`
	_, err = db.ExecContext(ctx, stmt, hashedPassword, u.ID)
	if err != nil {
		return err
	}

	return nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
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
}*/
