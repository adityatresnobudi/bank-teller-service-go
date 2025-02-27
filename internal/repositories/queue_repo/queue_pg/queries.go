package queue_pg

const GET_ALL_QUEUE = `
	SELECT id, status, queue_number, created_at, updated_at, user_id, services_id
	FROM queues
`

const GET_QUEUE_BY_ID = `
	SELECT id, status, queue_number, created_at, updated_at, user_id, services_id
	FROM queues WHERE id = $1
`

const GET_QUEUE_BY_QUEUE_NUMBER = `
	SELECT id, status, queue_number, created_at, updated_at, user_id, services_id
	FROM queues WHERE queue_number = $1
`

const INSERT_QUEUE = `
	INSERT INTO queues (queue_number, user_id, service_id) 
	VALUES ($1, $2, $3)
`

const UPDATE_QUEUE = `
	UPDATE queues
	SET status = $1
	WHERE id = $2
	RETURNING id, status, queue_number, created_at, updated_at, user_id, service_id
`

const DELETE_QUEUE = `
	DELETE FROM queues
	WHERE id = $1
`
