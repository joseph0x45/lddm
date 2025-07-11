create table if not exists products (
  id text not null primary key,
  name text not null,
  variant text not null,
  price integer not null,
  image text not null,
  description text not null
);

create table if not exists orders (
  id text not null primary key,
  issued_at text not null,
  customer_name text not null,
  customer_phone text not null,
  customer_address text not null,
  discount integer not null,
  total integer not null,
  total_with_discount integer not null
);

create table if not exists order_items (
  id text not null primary key,
  order_id text not null references orders(id),
  product_id text not null references products(id),
  product_name text not null,
  product_variant text not null,
  quantity integer not null,
  unit_price integer not null
);
