package user_pg

const GET_ALL_USER = `
	SELECT id, name, phone_number, password, role, email, created_at, updated_at
	FROM users
`

const GET_USER_BY_ID = `
	SELECT id, name, phone_number, password, role, email, created_at, updated_at
	FROM users WHERE id = $1
`

const GET_USER_BY_EMAIL = `
	SELECT id, name, phone_number, password, role, email, created_at, updated_at
	FROM users WHERE email = $1
`

const INSERT_USER = `
	INSERT INTO users (name, phone_number, password, email) 
	VALUES ($1, $2, $3, $4)
`

const UPDATE_USER = `
	UPDATE users
	SET name = $1, phone_number = $2, password = $3, email = $4, role = $5
	WHERE id = $6
	RETURNING id, name, phone_number, password, role, email, created_at, updated_at
`

const DELETE_USER = `
	DELETE FROM users
	WHERE id = $1
`