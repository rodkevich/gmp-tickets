-- v0.1.0

DROP SCHEMA public CASCADE;

CREATE SCHEMA if not exists public;

-- EXTENSIONS ------------------------------------------------------------------

CREATE EXTENSION if not exists "pgcrypto";

-- MIGRATION VERSIONS ----------------------------------------------------------

CREATE TABLE if not exists _schema_versions
(
    photos        varchar not null,
    ticket_photos varchar not null,
    tickets       varchar not null,
    tags          varchar not null,
    ticket_tags   varchar not null
);

-- INITIAL VERSIONS ------------------------------------------------------------

INSERT INTO _schema_versions (tickets, photos, ticket_photos, tags, ticket_tags)
VALUES ('0.0.1', '0.0.1', '0.0.1', '0.0.1', '0.0.1');

--- COMMON FUNCTIONS -----------------------------------------------------------

CREATE or replace function touch_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language 'plpgsql';

--- TABLES ---------------------------------------------------------------------

--- tickets --------------------------------------------------------------------

CREATE TYPE enum_tickets_perk_type AS ENUM
    ('Draft','Regular','Premium', 'Promoted');

CREATE table if not exists tickets
(
    id           uuid         not null default gen_random_uuid(),
    owner_id     uuid         not null,
    name_short   varchar(100) not null,
    name_ext     varchar(255),
    description  text,
    amount       smallint     not null default 1,
    price        float        not null,
    currency     smallint     not null,
    active       bool,
    perk         enum_tickets_perk_type,
    published_at timestamptz,
    created_at   timestamptz  not null default now(),
    updated_at   timestamptz  not null default now(),
    deleted_at   timestamptz,
    PRIMARY KEY (id)

--     TODO: create event type for delete from here
--     PRIMARY KEY (id),
--     CONSTRAINT fk_tickets_owner_id_users
--         foreign key (owner_id)
--             references users (id)
--             ON DELETE CASCADE

);

CREATE trigger tickets_update_updated_at
    before update
    on tickets
    FOR EACH ROW
EXECUTE procedure touch_updated_at();

--- ticket_photos --------------------------------------------------------------

CREATE table if not exists ticket_photos
(
    ticket_id  uuid        not null,
    photo_id   uuid        not null,
    main       bool        not null default false,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    PRIMARY KEY (ticket_id, photo_id),

    CONSTRAINT fk_ticket_photos_ticket_id_tickets
        foreign key (ticket_id)
            references tickets (id)
            ON DELETE CASCADE
);

CREATE trigger ticket_photos_update_updated_at
    before update
    on ticket_photos
    FOR EACH ROW
EXECUTE procedure touch_updated_at();

--- photos ---------------------------------------------------------------------

CREATE TYPE enum_photos_mime_type AS ENUM
    ('image/jpeg', 'image/png', 'image/tiff');

CREATE table if not exists photos
(
    id         uuid                  not null default gen_random_uuid(),
    type       enum_photos_mime_type not null,
    size_kb    int,
    created_at timestamptz           not null default now(),
    updated_at timestamptz           not null default now(),
    deleted_at timestamptz,
    PRIMARY KEY (id)
);

CREATE trigger photos_update_updated_at
    before update
    on photos
    FOR EACH ROW
EXECUTE procedure touch_updated_at();

--- tags -----------------------------------------------------------------------

CREATE table if not exists tags
(
    id          bigserial,
    name        varchar(255) not null,
    description text,
    created_at  timestamptz  not null default now(),
    updated_at  timestamptz  not null default now(),
    deleted_at  timestamptz,
    PRIMARY KEY (id)

);

CREATE trigger tags_update_updated_at
    before update
    on tags
    FOR EACH ROW
EXECUTE procedure touch_updated_at();

--- ticket_tags ----------------------------------------------------------------

CREATE table if not exists ticket_tags
(
    ticket_id  uuid        not null,
    tag_id     bigserial   not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    PRIMARY KEY (ticket_id, tag_id),

    CONSTRAINT fk_ticket_tags_ticket_id_tickets
        foreign key (ticket_id)
            references tickets (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_ticket_tags_tag_id_tags
        foreign key (tag_id)
            references tags (id)
            ON DELETE CASCADE
);

CREATE trigger ticket_tags_update_updated_at
    before update
    on ticket_tags
    FOR EACH ROW
EXECUTE procedure touch_updated_at();

--------------------------------------------------------------------------------
