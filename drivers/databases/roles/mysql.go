package roles

import (
	"context"
	"infion-BE/businesses/roles"

	"gorm.io/gorm"
)

type mysqlRolesRepository struct {
	Conn *gorm.DB
}

func NewRolesRepository(conn *gorm.DB) roles.Repository {
	return &mysqlRolesRepository{
		Conn: conn,
	}
}

func (nr *mysqlRolesRepository) Store(ctx context.Context, rolesDomain *roles.Domain) (roles.Domain, error) {
	rec := fromDomain(rolesDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return roles.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlRolesRepository) GetByID(ctx context.Context, rolesId int) (roles.Domain, error) {
	rec := Roles{}
	err := nr.Conn.Where("id = ?", rolesId).First(&rec).Error
	if err != nil {
		return roles.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlRolesRepository) Update(ctx context.Context, rolesDomain *roles.Domain) (roles.Domain, error) {
	rec := fromDomain(rolesDomain)

	result := nr.Conn.Save(&rec)
	if result.Error != nil {
		return roles.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlRolesRepository) Delete(ctx context.Context, rolesDomain *roles.Domain) (roles.Domain, error) {
	rec := fromDomain(rolesDomain)

	result := nr.Conn.Unscoped().Delete(&rec)
	if result.Error != nil {
		return roles.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}