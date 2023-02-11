CREATE TABLE "User"
(
    id serial not null unique,
    name varchar(255) not null,
    surname varchar(255),
    username varchar(255) not null unique,
    email varchar(255) unique,
    password varchar(255) not null
);

CREATE TABLE "UserInformation"
(
    id serial not null unique,
    user_id int REFERENCES "User" (id) ON DELETE CASCADE NOT NULL,
    paymentCard varchar(255),
    birthdate time,
    phone varchar(255),
    sex varchar(255)
);

CREATE TABLE "Category"
(
    id          serial not null unique,
    title       varchar(255)       not null,
    description text
);

CREATE TABLE "Product"
(
    id          serial not null unique,
    image text,
    title       varchar(255)       not null,
    description text,
    cost float not null,
    created_company varchar(255) not null,
    created_country varchar(255) not null,
    created_date time not null,
    category_id int REFERENCES "Category" (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE "CartProduct"
(
    id          serial not null unique,
    quantity int default 1,
    total float
);

CREATE TABLE "CartProductMany"
(
    id serial not null unique,
    cart_product_id int  REFERENCES "CartProduct" (id) ON DELETE CASCADE,
    product_id int REFERENCES "Product" (id) ON DELETE CASCADE
);

CREATE TABLE "Cart"
(
    id          serial not null unique,
    cart_product_id int REFERENCES "CartProduct" (id) ON DELETE CASCADE,
    created_at time,
    total float
);

CREATE TABLE "Order"
(
    id          serial not null unique,
    user_id int REFERENCES "User" (id) ON DELETE CASCADE NOT NULL,
    cart_id int REFERENCES "Cart" (id) ON DELETE CASCADE NOT NULL,
    status varchar(255) not null,
    confirmed boolean default false
);

CREATE TABLE "Review"
(
    id          serial not null unique,
    order_id int REFERENCES "Order" (id) ON DELETE CASCADE,
    user_id int REFERENCES "User" (id) ON DELETE CASCADE,
    image varchar(255),
    review_text text not null,
    rating int
);