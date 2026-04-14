-- +goose Up
create table if NOT exists products (
  id uuid primary key default gen_random_uuid (),
  name text not null,
  description text not null,
  price numeric(10, 2) not null,
  created_at timestamptz not null default now (),
  updated_at timestamptz not null default now ()
);

-- +goose Down
DROP TABLE IF EXISTS products;
