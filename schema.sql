PRAGMA foreign_keys = ON;

create table if not exists groups (
  id text not null primary key,
  name text not null unique,
  picture text not null
);

create table if not exists products (
  id text not null primary key,
  group_id text not null references groups(id) on delete cascade,
  name text not null,
  variant text not null,
  picture text not null,
  in_stock integer not null default 0,
  base_price integer not null
);

create table product_bundle_prices (
  id text not null primary key,
  product_id text not null references products(id) on delete cascade,
  quantity integer not null,
  bundle_price integer not null,
  UNIQUE(product_id, quantity)
);
