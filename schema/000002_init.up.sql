CREATE TABLE images
(
    id serial not null unique primary key
);

CREATE TABLE staffs
(
    id serial not null unique primary key,
    name varchar(255) not null,
    photo_id int not null references images(id),
    meta json
);