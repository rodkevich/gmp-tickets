-- v0.1.0
DROP SCHEMA public CASCADE;

CREATE SCHEMA if not exists public;

-- extensions ------------------------------------------------------------------
CREATE EXTENSION if not exists "uuid-ossp";

CREATE EXTENSION if not exists "pgcrypto";

-- migration versions ----------------------------------------------------------
CREATE TABLE if not exists _schema_versions
(
    permissions      varchar not null,
    photos           varchar not null,
    profiles         varchar not null,
    role_permissions varchar not null,
    roles            varchar not null,
    ticket_photos    varchar not null,
    tickets          varchar not null,
    user_roles       varchar not null,
    user_subs        varchar not null,
    users            varchar not null,
    tags             varchar not null,
    ticket_tags      varchar not null
);

-- initial versions ------------------------------------------------------------
INSERT INTO _schema_versions (users, roles, tickets, photos, profiles,
                              permissions, ticket_photos, user_roles,
                              user_subs,
                              role_permissions, tags, ticket_tags)
VALUES ('0.0.1', '0.0.1', '0.0.1', '0.0.1', '0.0.1', '0.0.1', '0.0.1', '0.0.1', '0.0.1',
        '0.0.1',
        '0.0.1', '0.0.1');

-------------------------------- FUNCTIONS -------------------------------------

CREATE or replace function touch_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language 'plpgsql';

-------------------------------- TABLES ----------------------------------------

--- permissions ----------------------------------------------------------------
CREATE TYPE enum_permissions_access_type AS ENUM
    ('All', 'Create', 'Read', 'Update', 'Delete');

CREATE table if not exists permissions
(
    id         serial                       not null,
    resource   varchar(255)                 not null,
    action     enum_permissions_access_type not null,
    created_at timestamptz                  not null default now(),
    updated_at timestamptz                  not null default now(),
    deleted_at timestamptz,
    CONSTRAINT permissions_pkey PRIMARY KEY (id)
);


CREATE trigger permissions_update_updated_at
    before update
    ON permissions
    FOR EACH ROW
EXECUTE procedure touch_updated_at();

--- roles ----------------------------------------------------------------------
CREATE TYPE enum_roles_entity_type AS ENUM
    ('Persistent', 'Custom');

CREATE table if not exists roles
(
    id          serial                 not null,
    type        enum_roles_entity_type not null default 'Custom'::enum_roles_entity_type,
    name        varchar(255)           not null,
    description text                   not null,
    created_at  timestamptz            not null default now(),
    updated_at  timestamptz            not null default now(),
    deleted_at  timestamptz,
    PRIMARY KEY (id)
);

CREATE trigger roles_update_updated_at
    before update
    ON roles
    FOR EACH ROW
EXECUTE procedure touch_updated_at();

--- role_permissions -----------------------------------------------------------
CREATE table if not exists role_permissions
(
    id            bigserial   not null,
    role_id       serial      not null,
    permission_id serial      not null,
    created_at    timestamptz not null default now(),
    updated_at    timestamptz not null default now(),
    deleted_at    timestamptz,
    PRIMARY KEY (id),

    CONSTRAINT fk_role_permissions_role_id_roles
        foreign key (role_id)
            references roles (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_role_permissions_permission_id_permissions
        foreign key (permission_id)
            references permissions (id)
            ON DELETE CASCADE
);

CREATE trigger role_permissions_update_updated_at
    before update
    on role_permissions
    FOR EACH ROW
EXECUTE procedure touch_updated_at();

--- users ----------------------------------------------------------------------
CREATE TYPE enum_users_status_type AS ENUM
    ('Active', 'Pending', 'Blocked');

CREATE table if not exists users
(
    id         uuid         not null  default uuid_generate_v4(),
    login      varchar(255) not null unique,
    password   varchar(255) not null,
    status     enum_users_status_type default 'Pending'::enum_users_status_type,
    created_at timestamptz  not null  default now(),
    updated_at timestamptz  not null  default now(),
    deleted_at timestamptz,
    PRIMARY KEY (id)
);

CREATE trigger users_update_updated_at
    before update
    on users
    FOR EACH ROW
EXECUTE procedure touch_updated_at();

--- user_roles -----------------------------------------------------------------
CREATE table if not exists user_roles
(
    id         uuid        not null default uuid_generate_v4(),
    user_id    uuid        not null,
    role_id    uuid        not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    PRIMARY KEY (id),

    CONSTRAINT fk_user_roles_user_id_users
        foreign key (user_id)
            references users (id)
            ON DELETE CASCADE
);

CREATE trigger user_roles_update_updated_at
    before update
    on user_roles
    FOR EACH ROW
EXECUTE procedure touch_updated_at();

--- user_subs ------------------------------------------------------------------
CREATE TYPE enum_user_subs_target_type AS ENUM
    ('ticket', 'user');

CREATE table if not exists user_subs
(
    id         uuid                       not null default uuid_generate_v4(),
    target     enum_user_subs_target_type not null,
    user_id    uuid                       not null,
    created_at timestamptz                not null default now(),
    updated_at timestamptz                not null default now(),
    deleted_at timestamptz,
    PRIMARY KEY (id),

    CONSTRAINT fk_user_subs_target_users
        foreign key (user_id)
            references users (id)
            ON DELETE CASCADE
);

CREATE trigger user_subs_update_updated_at
    before update
    on user_subs
    FOR EACH ROW
EXECUTE procedure touch_updated_at();

--- profiles -------------------------------------------------------------------
CREATE TYPE enum_profiles_service_type AS ENUM
    ('TicketService');

CREATE table if not exists profiles
(
    id           uuid                       not null default uuid_generate_v4(),
    active       bool,
    user_id      uuid                       not null,
    service_name enum_profiles_service_type not null,
    nickname     varchar(255),
    first_name   varchar(255),
    last_name    varchar(255),
    email        varchar(255),
    time_zone    smallint,
    avatar_url   text,
    created_at   timestamptz                not null default now(),
    updated_at   timestamptz                not null default now(),
    deleted_at   timestamptz,
    PRIMARY KEY (id),

    CONSTRAINT fk_profiles_user_id_users
        foreign key (user_id)
            references users (id)
            ON DELETE CASCADE
);


CREATE trigger profiles_update_updated_at
    before update
    on profiles
    FOR EACH ROW
EXECUTE procedure touch_updated_at();

--- tickets --------------------------------------------------------------------
CREATE TYPE enum_tickets_perk_type AS ENUM
    ('Draft','Regular','Premium', 'Promoted');

CREATE table if not exists tickets
(
    id           uuid         not null,
    owner_id     uuid         not null,
    name         varchar(100) not null,
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
    PRIMARY KEY (id),

    CONSTRAINT fk_tickets_owner_id_users
        foreign key (owner_id)
            references users (id)
            ON DELETE CASCADE
);

CREATE trigger tickets_update_updated_at
    before update
    on tickets
    FOR EACH ROW
EXECUTE procedure touch_updated_at();

--- ticket_photos --------------------------------------------------------------
CREATE table if not exists ticket_photos
(
    id         uuid        not null default uuid_generate_v4(),
    ticket_id  uuid        not null,
    photo_id   uuid        not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    PRIMARY KEY (id),

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
    id         uuid                  not null default uuid_generate_v4(),
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
    id         bigserial   not null,
    ticket_id  uuid        not null,
    tag_id     bigserial   not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    PRIMARY KEY (id),

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
