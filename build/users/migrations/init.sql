-- v0.1.0

DROP SCHEMA public CASCADE;

CREATE SCHEMA if not exists public;

-- EXTENSIONS ------------------------------------------------------------------

CREATE EXTENSION if not exists "pgcrypto";

-- MIGRATION VERSIONS ----------------------------------------------------------

CREATE TABLE if not exists _schema_versions
(
    permissions      varchar not null,
    profiles         varchar not null,
    role_permissions varchar not null,
    roles            varchar not null,
    user_roles       varchar not null,
    user_subs        varchar not null,
    users            varchar not null
);

-- INITIAL VERSIONS ------------------------------------------------------------

INSERT INTO _schema_versions (permissions, profiles, role_permissions, roles,
                              user_roles, user_subs, users)
VALUES ('0.0.1', '0.0.1', '0.0.1', '0.0.1', '0.0.1', '0.0.1', '0.0.1');

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

--- permissions ----------------------------------------------------------------

CREATE TYPE enum_permissions_access_type AS ENUM
    ('Administrate', 'Create', 'Read', 'Update', 'Delete');

CREATE table if not exists permissions
(
    id         serial                       not null,
    resource   varchar(255)                 not null,
    action     enum_permissions_access_type not null,
    created_at timestamptz                  not null default now(),
    updated_at timestamptz                  not null default now(),
    deleted_at timestamptz,
    PRIMARY KEY (id)
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
    id          bigserial              not null,
    type        enum_roles_entity_type not null default 'Custom'::enum_roles_entity_type,
    role_name   varchar(255)           not null,
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
    role_id       serial      not null,
    permission_id serial      not null,
    created_at    timestamptz not null default now(),
    updated_at    timestamptz not null default now(),
    deleted_at    timestamptz,
    PRIMARY KEY (role_id, permission_id),

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
    id         uuid         not null  default gen_random_uuid(),
    login      varchar(255) not null unique,
    email      varchar(255) not null unique,
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
    user_id    uuid        not null,
    role_id    bigserial   not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    PRIMARY KEY (user_id, role_id),

    CONSTRAINT fk_user_roles_user_id_users
        foreign key (user_id)
            references users (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_user_roles_role_id_roles
        foreign key (role_id)
            references roles (id)
            ON DELETE CASCADE
);

CREATE trigger user_roles_update_updated_at
    before update
    on user_roles
    FOR EACH ROW
EXECUTE procedure touch_updated_at();

--- user_subs ------------------------------------------------------------------

CREATE TYPE enum_user_subs_target_type AS ENUM
    ('Ticket', 'User');

CREATE table if not exists user_subs
(
    id         uuid                       not null default gen_random_uuid(),
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
    id           uuid                       not null default gen_random_uuid(),
    active       bool,
    user_id      uuid                       not null,
    service_name enum_profiles_service_type not null,
    nickname     varchar(255),
    first_name   varchar(255),
    last_name    varchar(255),
    email        varchar(255),
    time_zone    smallint,
    mobile       varchar(255),
    phone        varchar(255),
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
