CREATE TABLE "user"
(
    id serial primary key unique,
    name varchar(255) not null,
    surname varchar(255),
    username varchar(255) not null unique,
    email varchar(255) unique,
    password varchar(255) not null
);

CREATE TABLE "UserInformation"
(
    id serial primary key not null unique,
    user_id int,
    FOREIGN KEY (user_id) REFERENCES "user" ON DELETE CASCADE,
    paymentCard varchar(255),
    birthdate time,
    phone varchar(255),
    sex varchar(255)
);

CREATE TABLE "Category"
(
    id          serial primary key not null unique,
    title       varchar(255)       not null,
    description text
);

CREATE TABLE "Product"
(
    id          serial primary key not null unique,
    image text,
    title       varchar(255)       not null,
    description text,
    cost float not null,
    created_company varchar(255) not null,
    created_country varchar(255) not null,
    created_date time not null,
    category_id int,
    FOREIGN KEY (category_id) REFERENCES "Category" ON DELETE CASCADE
);

CREATE TABLE "CartProduct"
(
    id          serial primary key not null unique,
    quantity int default 1,
    total float
);

CREATE TABLE "CartProductMany"
(
    id serial primary key not null unique,
    cart_product_id int,
    FOREIGN KEY (cart_product_id) REFERENCES "CartProduct" ON DELETE SET NULL,
    product_id int,
    FOREIGN KEY (product_id) REFERENCES "Product" ON DELETE SET NULL
);

CREATE TABLE "Cart"
(
    id          serial primary key not null unique,
    cart_product_id int,
    FOREIGN KEY (cart_product_id) REFERENCES "CartProduct" ON DELETE CASCADE,
    created_at time,
    total float not null
);

CREATE TABLE "Order"
(
    id          serial primary key not null unique,
    user_id int,
    FOREIGN KEY (user_id) REFERENCES "user" ON DELETE CASCADE,
    cart_id int references "Cart"(id) on delete cascade not null,
    status varchar(255) not null,
    confirmed boolean default false
);

CREATE TABLE "Review"
(
    id          serial primary key not null unique,
    order_id int,
    FOREIGN KEY (order_id) REFERENCES "Order" ON DELETE CASCADE,
    image varchar(255),
    review_text text not null,
    rating int
);