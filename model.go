package auth551

import (
	"database/sql"
	"github.com/go51/model551"
	"time"
)

//--[ User Model ]--------
type UserModel struct {
	Id             int64 `db_table:"users"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Locked         bool
	Enabled        bool
	Name           string
	Email          string
	EmailCanonical string
	PasswordSalt   string
	Password       string
	Thumbnail      string
	Serial         string
	DeletedAt      time.Time `db_delete:"true"`
}

func NewUserModel() interface{} {
	return UserModel{}
}

func NewUserModelPointer() interface{} {
	return &UserModel{}
}

func (m *UserModel) SetId(id int64) {
	m.Id = id
}

func (m *UserModel) GetId() int64 {
	return m.Id
}

func (m *UserModel) Scan(rows *sql.Rows) error {
	return rows.Scan(
		&m.Id,
		&m.CreatedAt,
		&m.UpdatedAt,
		&m.Locked,
		&m.Enabled,
		&m.Name,
		&m.Email,
		&m.EmailCanonical,
		&m.PasswordSalt,
		&m.Password,
		&m.Thumbnail,
		&m.Serial,
	)
}

func (m *UserModel) SqlValues(sqlType model551.SqlType) []interface{} {
	values := make([]interface{}, 0, 13)

	if sqlType == model551.SQL_LOGICAL_DELETE {
		values = append(values, m.Id)
	}
	if sqlType == model551.SQL_INSERT {
		m.CreatedAt = time.Now()
		m.UpdatedAt = m.CreatedAt
	}
	if sqlType == model551.SQL_UPDATE {
		m.UpdatedAt = time.Now()
	}
	values = append(values, m.CreatedAt)
	values = append(values, m.UpdatedAt)
	values = append(values, m.Locked)
	values = append(values, m.Enabled)
	values = append(values, m.Name)
	values = append(values, m.Email)
	values = append(values, m.EmailCanonical)
	values = append(values, m.PasswordSalt)
	values = append(values, m.Password)
	values = append(values, m.Thumbnail)
	values = append(values, m.Serial)

	if sqlType == model551.SQL_UPDATE {
		values = append(values, m.Id)
	} else if sqlType == model551.SQL_LOGICAL_DELETE {
		m.DeletedAt = time.Now()
		values = append(values, m.DeletedAt)
	}

	return values

}

//--[/User Model ]--------

//--[ User Token Model ]--------
type UserTokenModel struct {
	Id           int64 `db_table:"user_tokens"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserId       int64
	Vendor       string
	AccessToken  string
	TokenSecret  string
	Expiry       time.Time
	RefreshToken string
	TokenType    string
	AccountId    string
	Name         string
	Thumbnail    string
	DeletedAt    time.Time `db_delete:"true"`
}

func NewUserTokenModel() interface{} {
	return UserTokenModel{}
}

func NewUserTokenModelPointer() interface{} {
	return &UserTokenModel{}
}

func (m *UserTokenModel) SetId(id int64) {
	m.Id = id
}

func (m *UserTokenModel) GetId() int64 {
	return m.Id
}

func (m *UserTokenModel) Scan(rows *sql.Rows) error {
	return rows.Scan(
		&m.Id,
		&m.CreatedAt,
		&m.UpdatedAt,
		&m.UserId,
		&m.Vendor,
		&m.AccessToken,
		&m.TokenSecret,
		&m.Expiry,
		&m.RefreshToken,
		&m.TokenType,
		&m.AccountId,
		&m.Name,
		&m.Thumbnail,
	)
}

func (m *UserTokenModel) SqlValues(sqlType model551.SqlType) []interface{} {
	values := make([]interface{}, 0, 11)

	if sqlType == model551.SQL_LOGICAL_DELETE {
		values = append(values, m.Id)
	}
	if sqlType == model551.SQL_INSERT {
		m.CreatedAt = time.Now()
		m.UpdatedAt = m.CreatedAt
	}
	if sqlType == model551.SQL_UPDATE {
		m.UpdatedAt = time.Now()
	}
	values = append(values, m.CreatedAt)
	values = append(values, m.UpdatedAt)
	values = append(values, m.UserId)
	values = append(values, m.Vendor)
	values = append(values, m.AccessToken)
	values = append(values, m.TokenSecret)
	values = append(values, m.Expiry)
	values = append(values, m.RefreshToken)
	values = append(values, m.TokenType)
	values = append(values, m.AccountId)
	values = append(values, m.Name)
	values = append(values, m.Thumbnail)

	if sqlType == model551.SQL_UPDATE {
		values = append(values, m.Id)
	} else if sqlType == model551.SQL_LOGICAL_DELETE {
		m.DeletedAt = time.Now()
		values = append(values, m.DeletedAt)
	}

	return values
}

//--[/User Token Model ]--------

//--[ Access Token Model ]--------
type AccessToken struct {
	AccessToken  string
	TokenSecret  string
	TokenType    string
	RefreshToken string
	Expiry       time.Time
}

//--[/Access Token Model ]--------

//--[ Access Token Model ]--------
type AccountInformation struct {
	Id      string
	Name    string
	Email   string
	Picture string
}

//--[/Access Token Model ]--------
