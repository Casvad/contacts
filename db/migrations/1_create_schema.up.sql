CREATE TABLE users
(
    id         SERIAL                 NOT NULL PRIMARY KEY,
    email      CHARACTER VARYING(255) NOT NULL,
    password   CHARACTER VARYING(255) NOT NULL,
    name       CHARACTER VARYING(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone
);

CREATE UNIQUE INDEX users_email_index ON users (email);

CREATE TABLE user_contacts
(
    id           SERIAL                 NOT NULL PRIMARY KEY,
    user_id      BIGINT                 NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    email        CHARACTER VARYING(255) NOT NULL,
    name         CHARACTER VARYING(255) NOT NULL,
    prefix_phone CHARACTER VARYING(5)   NOT NULL,
    phone        CHARACTER VARYING(20)  NOT NULL,
    created_at   timestamp without time zone NOT NULL,
    updated_at   timestamp without time zone NOT NULL,
    deleted_at   timestamp without time zone
);

CREATE UNIQUE INDEX user_contact_user_id_email_idx ON user_contacts (user_id, email);

CREATE TABLE user_contact_modifications
(
    id              SERIAL    NOT NULL PRIMARY KEY,
    user_contact_id BIGSERIAL NOT NULL REFERENCES user_contacts (id) ON DELETE CASCADE,
    modification    JSONB     NOT NULL,
    created_at      timestamp without time zone NOT NULL
);

CREATE INDEX user_contact_modifications_user_contact_id_idx ON user_contact_modifications (user_contact_id);