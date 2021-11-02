-- v0.1.1

CREATE SCHEMA if not exists public;

CREATE EXTENSION if not exists "uuid-ossp";

CREATE TYPE enum_user_type AS ENUM (
    'Administrator',
    'Regular',
    'Moderator',
    'Privileged',
    'Blocked',
    'Pending');

CREATE TABLE if not exists person
(
    id          uuid           not null default uuid_generate_v4(),
    login       varchar(255)   not null unique, -- RFC 5321
    avatar_url  text unique,
    profile_url text unique    not null,
    name        varchar(255),
    type        enum_user_type not null default 'Pending',
    admin       bool           not null default false,
    created_at  timestamptz    not null default now(),
    modified_at timestamptz    not null default now(),

    PRIMARY KEY (id)
);

CREATE OR REPLACE FUNCTION touch_modified() RETURNS TRIGGER AS
$$
BEGIN
    NEW.modified_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_person_modify_time
    BEFORE UPDATE
    ON person
    FOR EACH ROW
EXECUTE PROCEDURE touch_modified();

-- dummy users -----------------------------------------------------------------

INSERT INTO person (id, login, avatar_url, profile_url, name, type, admin, created_at,
                    modified_at)
VALUES (default,
        'login.moderator@gmp.tck',
        'https://www.example.com/pic/1?full=true',
        'https://www.user.example.com/1',
        default,
        default,
        false,
        default,
        default),
       (default,
        'login.administrator@gmp.tck',
        'https://www.example.com/pic/2?full=true',
        'https://www.user.example.com/2',
        'Odm1n',
        'Administrator',
        true,
        default,
        default)
RETURNING person.id;

--------------------------------------------------------------------------------

CREATE TYPE enum_ticket_status AS ENUM (
    'Draft',
    'Active',
    'Closed');

CREATE TABLE if not exists ticket
(
    id           uuid               not null default uuid_generate_v4(),
    name         varchar(150)       not null,
    full_name    varchar(255),
    description  text               not null,
    status       enum_ticket_status not null,
    owner_id     uuid               not null,
    amount       integer            not null,
    price        numeric(100, 2),
    currency     numeric(5),
    created_at   timestamptz        not null default current_timestamp,
    updated_at   timestamptz        not null default current_timestamp,
    deleted_at   timestamptz,
    published_at timestamptz,

    primary key (id),
    foreign key (owner_id)
        references person (id)
);

CREATE OR REPLACE FUNCTION touch_updated()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_ticket_update_time
    BEFORE UPDATE
    ON ticket
    FOR EACH ROW
EXECUTE PROCEDURE
    touch_updated();

--------------------------------------------------------------------------------

CREATE TABLE if not exists photo
(
    id        uuid         not null default uuid_generate_v4(),
    ticket_id uuid         not null,
    is_main   bool         not null default false,
    presented bool         not null default true,
    mime_type varchar(100) not null,
    size_kb   integer      not null,

    PRIMARY KEY (id),

    CONSTRAINT fk_ticket
        FOREIGN KEY (ticket_id)
            REFERENCES ticket (id)
            ON DELETE CASCADE
);
