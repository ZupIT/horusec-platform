package database

import (
	"time"

	"github.com/google/uuid"
)

type Filter struct {
	RepositoryID uuid.UUID
	WorkspaceID  uuid.UUID
	StartTime    time.Time
	EndTime      time.Time
	Page         int
	Size         int
}

func (f *Filter) GetConditionFilter() (string, []interface{}) {
	query, args := f.getWorkspaceFilter()
	query, args = f.getRepositoryFilter(query, args)
	query, args = f.getInitialDateFilter(query, args)
	query, args = f.getFinalDateFilter(query, args)

	return query, args
}

func (f *Filter) getWorkspaceFilter() (string, []interface{}) {
	return "workspace_id = ? ", []interface{}{f.WorkspaceID}
}

func (f *Filter) getRepositoryFilter(query string, args []interface{}) (string, []interface{}) {
	if f.RepositoryID != uuid.Nil {
		query += "AND repository_id = ? "
		args = append(args, f.RepositoryID)
	}

	return query, args
}

func (f *Filter) getInitialDateFilter(query string, args []interface{}) (string, []interface{}) {
	if !f.StartTime.IsZero() {
		query += "AND created_at >= ? "
		args = append(args, f.StartTime)
	}

	return query, args
}

func (f *Filter) getFinalDateFilter(query string, args []interface{}) (string, []interface{}) {
	if !f.EndTime.IsZero() {
		query += "AND created_at <= ? "
		args = append(args, f.EndTime)
	}

	return query, args
}
