package reports

import (
	"context"
	"time"
)

type Domain struct {
	ID           	int
	ThreadID		int
	Title			string
	UserID			int
	Moderator		string
	ReportMessage	string
	Status			string
	CreatedAt    	time.Time
	UpdatedAt    	time.Time
	DeletedAt		time.Time
}

type Usecase interface {
	Store(ctx context.Context, reportsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, reportsId int) (Domain, error)
	GetReports(ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, reportsDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, reportsDomain *Domain) (*Domain, error)
}

type Repository interface {
	Store(ctx context.Context, reportsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, reportsId int) (Domain, error)
	GetReports(ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, reportsDomain *Domain) (Domain, error)
	Delete(ctx context.Context, reportsDomain *Domain) (Domain, error)
}