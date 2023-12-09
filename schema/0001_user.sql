-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS crm_user(
                                       id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
                                       full_name VARCHAR(64) NOT NULL,
                                       nick_name varchar(64) NOT NULL UNIQUE,
                                       birthday_date TIMESTAMP NOT NULL,
                                       password TEXT NOT NULL,
                                       photo VARCHAR(32)  NULL,
                                       location TEXT NOT NULL ,
                                       created_at TIMESTAMP DEFAULT (NOW()),
                                       updated_at TIMESTAMP NULL,
                                       deleted_at TIMESTAMP NULL,
                                       created_by uuid NULL ,
                                       updated_by uuid NULL ,
                                       deleted_by uuid NULL
);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back