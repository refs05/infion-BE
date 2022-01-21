package reports

import (
	"context"
	"infion-BE/businesses/reports"

	"gorm.io/gorm"
)

type mysqlReportsRepository struct {
	Conn *gorm.DB
}

func NewReportsRepository(conn *gorm.DB) reports.Repository {
	return &mysqlReportsRepository{
		Conn: conn,
	}
}

func (nr *mysqlReportsRepository) Store(ctx context.Context, reportsDomain *reports.Domain) (reports.Domain, error) {
	rec := fromDomain(reportsDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return reports.Domain{}, result.Error
	}

	record, err := nr.GetByID(ctx, rec.ID)
	if err != nil {
		return reports.Domain{}, err
	}

	return record, nil
}

func (nr *mysqlReportsRepository) GetByID(ctx context.Context, reportsId int) (reports.Domain, error) {
	rec := Reports{}
	err := nr.Conn.Where("reports.id = ?", reportsId).Joins("Thread").Joins("User").First(&rec).Error
	if err != nil {
		return reports.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlReportsRepository) Update(ctx context.Context, reportsDomain *reports.Domain) (reports.Domain, error) {
	rec := fromDomain(reportsDomain)

	result := nr.Conn.Save(&rec)
	if result.Error != nil {
		return reports.Domain{}, result.Error
	}

	record, err := nr.GetByID(ctx, rec.ID)
	if err != nil {
		return reports.Domain{}, err
	}

	return record, nil
}

func (nr *mysqlReportsRepository) Delete(ctx context.Context, reportsDomain *reports.Domain) (reports.Domain, error) {
	rec := fromDomain(reportsDomain)

	result := nr.Conn.Delete(&rec)
	if result.Error != nil {
		return reports.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlReportsRepository) GetReports(ctx context.Context) ([]reports.Domain, error) {
	var recordThread []Reports
	
	result := nr.Conn.Unscoped().Joins("Thread").Joins("User").Find(&recordThread)
	if result.Error != nil {
		return []reports.Domain{}, result.Error
	}

	return ToDomainArray(recordThread), nil
}

func (nr *mysqlReportsRepository) GetReportsByUserID(ctx context.Context, userID int) ([]reports.Domain, error) {
	var recordReport []Reports
	
	result := nr.Conn.Unscoped().Where("reports.user_id = ?", userID).Joins("User").Find(&recordReport)
	if result.Error != nil {
		return []reports.Domain{}, result.Error
	}

	return ToDomainArray(recordReport), nil
}