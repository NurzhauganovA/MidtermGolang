CREATE TABLE Users
(
    id serial not null unique,
    firstname varchar(255) not null,
    surname varchar(255),
    username varchar(255) not null unique,
    email varchar(255) unique,
    password_hash varchar(255) not null
);

CREATE TABLE UserInformation
(
    id serial not null unique,
    user_id int REFERENCES Users (id) ON DELETE CASCADE NOT NULL,
    paymentCard varchar(255),
    birthdate time,
    phone varchar(255),
    sex varchar(255)
);

CREATE TABLE Category
(
    id          serial not null unique,
    title       varchar(255)       not null,
    description text
);

CREATE TABLE UserCategory
(
    id serial not null UNIQUE,
    user_id int REFERENCES Users (id) on delete cascade not null,
    category_id int references Category (id) on delete cascade not null
);

CREATE TABLE Product
(
    id          serial not null unique,
    image text,
    title       varchar(255)       not null,
    description text,
    cost float not null,
    created_company varchar(255) not null,
    created_country varchar(255) not null,
    created_date time
);

CREATE TABLE CategoryProduct
(
    id serial not null unique,
    product_id int references Product (id) on delete cascade not null,
    category_id int references Category (id) on delete cascade not null
);

CREATE TABLE CartProduct
(
    id          serial not null unique,
    quantity int default 1,
    total float
);

CREATE TABLE CartProductMany
(
    id serial not null unique,
    cart_product_id int  REFERENCES CartProduct (id) ON DELETE CASCADE,
    product_id int REFERENCES Product (id) ON DELETE CASCADE
);

CREATE TABLE Cart
(
    id          serial not null unique,
    cart_product_id int REFERENCES CartProduct (id) ON DELETE CASCADE,
    created_at time,
    total float
);

CREATE TABLE OrderProduct
(
    id          serial not null unique,
    user_id int REFERENCES Users (id) ON DELETE CASCADE NOT NULL,
    cart_id int REFERENCES Cart (id) ON DELETE CASCADE NOT NULL,
    status varchar(255) not null,
    confirmed boolean default false
);

CREATE TABLE Review
(
    id          serial not null unique,
    order_id int REFERENCES OrderProduct (id) ON DELETE CASCADE,
    user_id int REFERENCES Users (id) ON DELETE CASCADE,
    image varchar(255),
    review_text text not null,
    rating int
);