-- name: FindUserByID :one
SELECT u.user_id,
       u.username,
       u.org_id,
       o.org_extl_id,
       o.org_name,
       o.org_description,
       u.person_profile_id,
       pp.name_prefix,
       pp.first_name,
       pp.middle_name,
       pp.last_name,
       pp.name_suffix,
       pp.nickname,
       pp.company_name,
       pp.company_dept,
       pp.job_title,
       pp.birth_date,
       pp.birth_year,
       pp.birth_month,
       pp.birth_day,
       pp.language_id,
       p.person_id
FROM app_user u
         inner join org o on o.org_id = u.org_id
         inner join person_profile pp on pp.person_profile_id = u.person_profile_id
         inner join person p on p.person_id = pp.person_id
WHERE u.user_id = $1
LIMIT 1;

-- name: FindUserByUsername :one
SELECT u.user_id,
       u.username,
       u.org_id,
       o.org_extl_id,
       o.org_name,
       o.org_description,
       u.person_profile_id,
       pp.name_prefix,
       pp.first_name,
       pp.middle_name,
       pp.last_name,
       pp.name_suffix,
       pp.nickname,
       pp.company_name,
       pp.company_dept,
       pp.job_title,
       pp.birth_date,
       pp.birth_year,
       pp.birth_month,
       pp.birth_day,
       pp.language_id,
       p.person_id
FROM app_user u
         inner join org o on o.org_id = u.org_id
         inner join person_profile pp on pp.person_profile_id = u.person_profile_id
         inner join person p on p.person_id = pp.person_id
WHERE u.username = $1
  AND u.org_id = $2
LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO app_user (user_id, username, org_id, person_profile_id, create_app_id, create_user_id,
                      create_timestamp, update_app_id, update_user_id, update_timestamp)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

-- name: DeleteUser :exec
DELETE
FROM app_user
WHERE user_id = $1;
