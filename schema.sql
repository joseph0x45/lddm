create table if not exists products (
  id text not null primary key,
  name text not null,
  price integer not null,
  wholesale_price integer not null,
  image text not null,
  quantity_in_stock int not null,
  description text not null
);

create table if not exists orders (
  id text not null primary key,
  issued_at text not null,
  customer_name text not null,
  total int not null
);

create table if not exists order_items (
  id text not null primary key,
  order_id text not null references orders(id),
  product_id text not null references products(id),
  quantity integer not null,
  unit_price integer not null
);
