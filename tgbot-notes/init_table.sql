CREATE TABLE notes (
                       id bigserial NOT NULL PRIMARY KEY,
                       chat_id int not null,
                       note text not null,
                       date date not null,
                       status text not null
)