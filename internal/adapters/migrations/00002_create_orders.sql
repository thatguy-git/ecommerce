-- +goose Up
CREATE TABLE IF not EXISTS orders (
  id bigserial PRIMARY KEY,
  customer_id bigint not null,
  created_at timestamptz default now ()
);

CREATE table if not exists order_items (
  id bigserial primary key,
  order_id bigint not null,
  product_id bigint not null,
  quantity int not null,
  price_cents int not null,
  constraint fk_order foreign key (order_id) references orders (id)
);

CREATE TABLE if not exists deez_nuts (
  id bigserial primary key,
  nut_id bigint not null,
  quantitiy int not null,
  weight bigint not null,
  created at timestamptz default now ()
);

-- +goose Down
drop table if exists orders;

drop table if exists order_items;
