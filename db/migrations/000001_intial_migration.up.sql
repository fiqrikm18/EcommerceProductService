CREATE TABLE brands
(
    id         serial PRIMARY KEY,
    name       varchar(100),
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp default null
);

CREATE TABLE products
(
    id       serial PRIMARY KEY,
    name     varchar(100) NOT NULL,
    price    INTEGER      NOT NULL,
    qty      INTEGER      NOT NULL,
    brand_id INTEGER      NOT NULL,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp default null,
    CONSTRAINT fk_brand_product FOREIGN KEY (brand_id) REFERENCES brands (id)
);
