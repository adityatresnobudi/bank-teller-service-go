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
	RETURNING id, name, phone_number, password, role, email, created_at, updated_at
`

const UPDATE_USER = `
	UPDATE users
	SET account_holder = $1, balance = $2
	WHERE id = $3
	RETURNING id, name, phone_number, password, role, email, created_at, updated_at
`

const DELETE_USER = `
	DELETE FROM users
	WHERE id = $1
`