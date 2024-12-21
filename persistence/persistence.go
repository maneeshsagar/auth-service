package persistence

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/maneeshsagar/auth-service/models"
	"github.com/spf13/cast"
)

type Persistence interface {
	AddUser(user *models.User) (int64, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUserId(id int) (*models.User, error)
	AddToken(token *models.Token) (int64, error)
	GetToken(token string) (*models.Token, error)
	GetAccesTokenByRefreshToken(refreshToken string) (*models.Token, error)
	AddRefreshToken(token *models.RefreshToken) (int64, error)
	GetRefreshToken(token string) (*models.RefreshToken, error)
}

type PersistenceImpl struct {
}

func NewPersistence() Persistence {
	return &PersistenceImpl{}
}

func (p *PersistenceImpl) AddUser(user *models.User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(user)
	return id, err
}
func (p *PersistenceImpl) GetUserByEmail(email string) (*models.User, error) {
	o := orm.NewOrm()
	m := &models.User{}
	err := o.QueryTable(new(models.User)).Filter("email", email).One(m)
	if err != nil {
		return nil, err
	}
	return m, err
}
func (p *PersistenceImpl) AddToken(token *models.Token) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(token)
	return id, err
}
func (p *PersistenceImpl) GetToken(token string) (*models.Token, error) {
	o := orm.NewOrm()
	m := &models.Token{}
	err := o.QueryTable(new(models.Token)).Filter("token", token).One(m)
	if err != nil {
		return nil, err
	}
	return m, err
}

func (p *PersistenceImpl) AddRefreshToken(token *models.RefreshToken) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(token)
	return id, err
}
func (p *PersistenceImpl) GetRefreshToken(token string) (*models.RefreshToken, error) {
	o := orm.NewOrm()
	m := &models.RefreshToken{}
	err := o.QueryTable(new(models.RefreshToken)).Filter("token", token).One(m)
	if err != nil {
		return nil, err
	}
	return m, err
}

func (p *PersistenceImpl) GetAccesTokenByRefreshToken(refreshToken string) (*models.Token, error) {
	o := orm.NewOrm()
	m := &models.Token{}
	err := o.QueryTable(new(models.Token)).Filter("refresh_token", refreshToken).OrderBy("-id").One(m)
	if err != nil {
		return nil, err
	}
	return m, err
}

func (p *PersistenceImpl) GetUserByUserId(id int) (*models.User, error) {
	o := orm.NewOrm()

	u := &models.User{
		ID: cast.ToInt64(id),
	}
	err := o.Read(u)
	if err != nil {
		return nil, err
	}
	return u, err
}
