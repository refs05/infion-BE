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
	Reporter		string
	ReportMessage	string
	Status			bool
	CreatedAt    	time.Time
	UpdatedAt    	time.Time
	DeletedAt		time.Time
}

type Usecase interface {
	Store(ctx context.Context, reportsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, reportsId int) (Domain, error)
	GetReports(ctx context.Context) ([]Domain, error)
	GetReportsByUserID(ctx context.Context, userID int) ([]Domain, error)
	Update(ctx context.Context, reportsDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, reportsDomain *Domain) (*Domain, error)
}

type Repository interface {
	Store(ctx context.Context, reportsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, reportsId int) (Domain, error)
	GetReports(ctx context.Context) ([]Domain, error)
	GetReportsByUserID(ctx context.Context, userID int) ([]Domain, error)
	GetReportsByThreadID(ctx context.Context, threadID int) ([]Domain, error)
	Update(ctx context.Context, reportsDomain *Domain) (Domain, error)
	Delete(ctx context.Context, reportsDomain *Domain) (Domain, error)
}