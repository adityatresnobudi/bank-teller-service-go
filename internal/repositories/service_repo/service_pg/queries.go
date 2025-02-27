package service_pg

const GET_ALL_SERVICE = `
	SELECT id, code, name, created_at, updated_at
	FROM services
`

const GET_SERVICE_BY_ID = `
	SELECT id, code, name, created_at, updated_at
	FROM services WHERE id = $1
`

const GET_SERVICE_BY_CODE = `
	SELECT id, code, name, created_at, updated_at
	FROM services WHERE code = $1
`

const GET_SERVICE_BY_NAME = `
	SELECT id, code, name, created_at, updated_at
	FROM services WHERE name = $1
`

const INSERT_SERVICE = `
	INSERT INTO services (code, name) 
	VALUES ($1, $2)
`

const UPDATE_SERVICE = `
	UPDATE services
	SET code = $1, name = $2
	WHERE id = $3
	RETURNING id, code, name, created_at, updated_at
`

const DELETE_SERVICE = `
	DELETE FROM services
	WHERE id = $1
`
