package announcements

import (
	"context"
	"infion-BE/businesses/announcements"

	"gorm.io/gorm"
)

type mysqlAnnouncementsRepository struct {
	Conn *gorm.DB
}

func NewAnnouncementsRepository(conn *gorm.DB) announcements.Repository {
	return &mysqlAnnouncementsRepository{
		Conn: conn,
	}
}

func (nr *mysqlAnnouncementsRepository) Store(ctx context.Context, announcementsDomain *announcements.Domain) (announcements.Domain, error) {
	rec := fromDomain(announcementsDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return announcements.Domain{}, result.Error
	}

	record, err := nr.GetByID(ctx, rec.ID)
	if err != nil {
		return announcements.Domain{}, err
	}

	return record, nil
}

func (nr *mysqlAnnouncementsRepository) GetByID(ctx context.Context, announcementsId int) (announcements.Domain, error) {
	rec := Announcements{}
	err := nr.Conn.Where("announcements.id = ?", announcementsId).Joins("User").First(&rec).Error
	if err != nil {
		return announcements.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlAnnouncementsRepository) Update(ctx context.Context, announcementsDomain *announcements.Domain) (announcements.Domain, error) {
	rec := fromDomain(announcementsDomain)

	result := nr.Conn.Save(&rec)
	if result.Error != nil {
		return announcements.Domain{}, result.Error
	}

	record, err := nr.GetByID(ctx, rec.ID)
	if err != nil {
		return announcements.Domain{}, err
	}

	return record, nil
}

func (nr *mysqlAnnouncementsRepository) Delete(ctx context.Context, announcementsDomain *announcements.Domain) (announcements.Domain, error) {
	rec := fromDomain(announcementsDomain)

	result := nr.Conn.Unscoped().Delete(&rec)
	if result.Error != nil {
		return announcements.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlAnnouncementsRepository) GetAnnouncements(ctx context.Context) ([]announcements.Domain, error) {
	var recordAnnouncement []Announcements
	
	result := nr.Conn.Unscoped().Order("created_at desc").Joins("User").Find(&recordAnnouncement)
	if result.Error != nil {
		return []announcements.Domain{}, result.Error
	}

	return ToDomainArray(recordAnnouncement), nil
}

func (nr *mysqlAnnouncementsRepository) GetAnnouncementsByUserID(ctx context.Context, userID int) ([]announcements.Domain, error) {
	var recordAnnouncement []Announcements
	
	result := nr.Conn.Unscoped().Order("created_at desc").Where("announcements.user_id = ?", userID).Joins("User").Find(&recordAnnouncement)
	if result.Error != nil {
		return []announcements.Domain{}, result.Error
	}

	return ToDomainArray(recordAnnouncement), nil
}