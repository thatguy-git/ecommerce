-- name: ListProducts :many
SELECT * FROM products
ORDER BY name;

-- name: FindProductById :one
select * from products
where id = $1;

