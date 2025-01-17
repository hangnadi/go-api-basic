// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package orgstore

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
)

const countOrgs = `-- name: CountOrgs :one
SELECT count(*) as org_count FROM org
`

func (q *Queries) CountOrgs(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countOrgs)
	var org_count int64
	err := row.Scan(&org_count)
	return org_count, err
}

const createOrg = `-- name: CreateOrg :execresult
INSERT INTO org (org_id, org_extl_id, org_name, org_description, create_app_id, create_user_id,
                 create_timestamp, update_app_id, update_user_id, update_timestamp)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
`

type CreateOrgParams struct {
	OrgID           uuid.UUID
	OrgExtlID       string
	OrgName         string
	OrgDescription  string
	CreateAppID     uuid.UUID
	CreateUserID    uuid.NullUUID
	CreateTimestamp time.Time
	UpdateAppID     uuid.UUID
	UpdateUserID    uuid.NullUUID
	UpdateTimestamp time.Time
}

func (q *Queries) CreateOrg(ctx context.Context, arg CreateOrgParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, createOrg,
		arg.OrgID,
		arg.OrgExtlID,
		arg.OrgName,
		arg.OrgDescription,
		arg.CreateAppID,
		arg.CreateUserID,
		arg.CreateTimestamp,
		arg.UpdateAppID,
		arg.UpdateUserID,
		arg.UpdateTimestamp,
	)
}

const deleteOrg = `-- name: DeleteOrg :exec
DELETE FROM org
WHERE org_id = $1
`

func (q *Queries) DeleteOrg(ctx context.Context, orgID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteOrg, orgID)
	return err
}

const findOrgByExtlID = `-- name: FindOrgByExtlID :one
SELECT org_id, org_extl_id, org_name, org_description, create_app_id, create_user_id, create_timestamp, update_app_id, update_user_id, update_timestamp FROM org
WHERE org_extl_id = $1 LIMIT 1
`

func (q *Queries) FindOrgByExtlID(ctx context.Context, orgExtlID string) (Org, error) {
	row := q.db.QueryRow(ctx, findOrgByExtlID, orgExtlID)
	var i Org
	err := row.Scan(
		&i.OrgID,
		&i.OrgExtlID,
		&i.OrgName,
		&i.OrgDescription,
		&i.CreateAppID,
		&i.CreateUserID,
		&i.CreateTimestamp,
		&i.UpdateAppID,
		&i.UpdateUserID,
		&i.UpdateTimestamp,
	)
	return i, err
}

const findOrgByID = `-- name: FindOrgByID :one
SELECT org_id, org_extl_id, org_name, org_description, create_app_id, create_user_id, create_timestamp, update_app_id, update_user_id, update_timestamp FROM org
WHERE org_id = $1 LIMIT 1
`

func (q *Queries) FindOrgByID(ctx context.Context, orgID uuid.UUID) (Org, error) {
	row := q.db.QueryRow(ctx, findOrgByID, orgID)
	var i Org
	err := row.Scan(
		&i.OrgID,
		&i.OrgExtlID,
		&i.OrgName,
		&i.OrgDescription,
		&i.CreateAppID,
		&i.CreateUserID,
		&i.CreateTimestamp,
		&i.UpdateAppID,
		&i.UpdateUserID,
		&i.UpdateTimestamp,
	)
	return i, err
}

const findOrgs = `-- name: FindOrgs :many
SELECT org_id, org_extl_id, org_name, org_description, create_app_id, create_user_id, create_timestamp, update_app_id, update_user_id, update_timestamp FROM org
ORDER BY org_name
`

func (q *Queries) FindOrgs(ctx context.Context) ([]Org, error) {
	rows, err := q.db.Query(ctx, findOrgs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Org
	for rows.Next() {
		var i Org
		if err := rows.Scan(
			&i.OrgID,
			&i.OrgExtlID,
			&i.OrgName,
			&i.OrgDescription,
			&i.CreateAppID,
			&i.CreateUserID,
			&i.CreateTimestamp,
			&i.UpdateAppID,
			&i.UpdateUserID,
			&i.UpdateTimestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrg = `-- name: UpdateOrg :exec
UPDATE org
SET org_name         = $1,
    org_description  = $2,
    update_app_id    = $3,
    update_user_id   = $4,
    update_timestamp = $5
WHERE org_id = $6
`

type UpdateOrgParams struct {
	OrgName         string
	OrgDescription  string
	UpdateAppID     uuid.UUID
	UpdateUserID    uuid.NullUUID
	UpdateTimestamp time.Time
	OrgID           uuid.UUID
}

func (q *Queries) UpdateOrg(ctx context.Context, arg UpdateOrgParams) error {
	_, err := q.db.Exec(ctx, updateOrg,
		arg.OrgName,
		arg.OrgDescription,
		arg.UpdateAppID,
		arg.UpdateUserID,
		arg.UpdateTimestamp,
		arg.OrgID,
	)
	return err
}
