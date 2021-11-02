CREATE SCHEMA if not exists dev;

CREATE EXTENSION if not exists "uuid-ossp" SCHEMA dev;

----

CREATE TYPE dev.enum_user_type AS ENUM (
    'Administrator',
    'Regular',
    'Moderator',
    'Privileged',
    'Blocked',
    'Pending');

----

CREATE TABLE if not exists dev.users
(
    id          uuid                        default dev.uuid_generate_v4(),
    login       varchar(100)       not null unique,
    avatar_url  varchar(255) unique,
    url         varchar(255) unique,
    name        text               not null default 'Unnamed User',
    type        dev.enum_user_type not null,
    admin       bool                        default false,
    created_at  timestamptz                 default now(),
    modified_at timestamptz                 default now(),
    primary key (id)
);

----

CREATE OR REPLACE FUNCTION dev.update_modified_column() RETURNS TRIGGER AS
$$
BEGIN
    NEW.modified_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

----

CREATE TRIGGER update_users_modify_time
    BEFORE UPDATE
    ON dev.users
    FOR EACH ROW
EXECUTE PROCEDURE dev.update_modified_column();

----

SELECT dev.uuid_generate_v4();

----

INSERT INTO dev.users (login, avatar_url, url, name, type, admin)
VALUES ('login.moderator',
        'https://www.example.com/pic/1?full=true',
        'https://www.user.example.com/1',
        default,
        'Moderator',
        false)
RETURNING id;


INSERT INTO dev.users (login, avatar_url, url, name, type, admin)
VALUES ('login.administrator',
        'https://www.example.com/pic/2?full=true',
        'https://www.user.example.com/2',
        'odmin',
        'Administrator',
        true)
RETURNING id;

----

CREATE TABLE if not exists dev.photos
(
    id            uuid default dev.uuid_generate_v4(),
    is_main_photo bool default false,
    presented     bool default true,
    mime_type     varchar(100),
    size_kb       integer      not null,
    link_address  varchar(255) not null,
    primary key (id)

);

----

CREATE TYPE dev.enum_ticket_status AS ENUM (
    'Draft',
    'Active',
    'Closed');

----

CREATE TABLE if not exists dev.tickets
(
    id                uuid        default dev.uuid_generate_v4(),
    name              varchar(150)           not null,
    full_name         varchar(255),
    description       text                   not null,
    status            dev.enum_ticket_status not null,
    owner_id          uuid                   not null,
    amount            numeric(20)            not null,
    price             numeric(100, 2)        not null,
    currency          numeric(10)            not null,
    main_photo_id     uuid,
    additional_photos uuid ARRAY[3],
    created_at        timestamptz default current_timestamp,
    updated_at        timestamptz default current_timestamp,
    deleted_at        timestamptz,
    published_at      timestamptz,

    primary key (id),
    foreign key (owner_id)
        references dev.users (id)

--     foreign key (additional_photos)
--         references public.photos (id)
--         on delete cascade
);

----

CREATE OR REPLACE FUNCTION dev.update_updated_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

----

CREATE TRIGGER update_ticket_update_time
    BEFORE UPDATE
    ON dev.tickets
    FOR EACH ROW
EXECUTE PROCEDURE
    dev.update_updated_column();

----

INSERT INTO dev.tickets
(id, name, full_name, description, status, owner_id,
 amount, price, main_photo_id, additional_photos, currency,
 created_at, updated_at, deleted_at, published_at)
VALUES (default,
        'test_ticket',
        'test_ticket_full_name',
        'test_ticket_description',
        'Active',
        '3733649e-f2b6-4b26-a4ef-1f4fb486bf9f',
        22,
        333.33,
        'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11',
--         '{a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11, b0eebc99-9c0b-4ef8-bb6d-cbb9bd380a11}',
        ARRAY ['a6e34e5d-b1fb-4240-8ad9-21ddf23134bb','a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11']::uuid[],
        251,
        default,
        default,
        default,
        default);

--
-- UPDATE dev.tickets
-- SET main_photo_id     = '519d3bea-276e-4dcc-bbd2-5f3cea35d011',
--     additional_photos = '{{519d3bea-276e-4dcc-bbd2-5f3cea35d011},{519d3bea-276e-4dcc-bbd2-5f3cea35d011},{519d3bea-276e-4dcc-bbd2-5f3cea35d011},{519d3bea-276e-4dcc-bbd2-5f3cea35d011}}'
-- WHERE id = '0b877c1d-3bef-41bb-b66d-ee746a772389'

--
-- INSERT INTO public.photos (id, ticket_id, is_main_photo, presented, mime_type, size_kb,
--                            link_address)
-- VALUES (DEFAULT, '8479ba22-96f4-4b93-aeda-69c8de7a65f0', false, true, 'image/jpg', 4400,
--         'http://example.com/');
--
--
-- INSERT INTO public.photos
-- (ticket_id, is_main_photo, presented, mime_type, size_kb, link_address)
-- VALUES ('8479ba22-96f4-4b93-aeda-69c8de7a65f0', true, true, 'image/jpg', 64700,
--         'http://example.com/');

