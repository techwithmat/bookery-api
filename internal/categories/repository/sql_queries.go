package repository

const (
	getCategoriesQuery = `
	SELECT 
		id,
		name
	FROM
		categories
	LIMIT $1 OFFSET $2;
	`

	getCategoryByIdQuery = `
	SELECT 
		id,
		name
	FROM
		categories
	WHERE 
		id = $1;
	`

	insertCategoryQuery = `
	INSERT INTO
		categories (name)
	VALUES 
		($1)
	RETURNING 
		id;
	`

	deleteCategoryQuery = `
	DELETE FROM
		categories
	WHERE
		id = $1;
	`

	updateCategoryQuery = `
	UPDATE 
		categories
	SET
		name = :name
	WHERE
		id = :id;
	`
)
