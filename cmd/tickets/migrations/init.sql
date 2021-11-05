-- v0.1.0
DROP SCHEMA public CASCADE;

CREATE SCHEMA if not exists public;

-- extensions ------------------------------------------------------------------
CREATE EXTENSION if not exists "uuid-ossp";

CREATE EXTENSION if not exists "pgcrypto";

-- migration versions ----------------------------------------------------------
CREATE TABLE if not exists _schema_versions
(
    users            varchar not null,
    roles            varchar not null,
    tickets          varchar not null,
    photos           varchar not null,
    profiles         varchar not null,
    permissions      varchar not null,
    ticket_photos    varchar not null,
    user_roles       varchar not null,
    user_tickets     varchar not null,
    user_photos      varchar not null,
    user_profiles    varchar not null,
    role_permissions varchar not null
);

-- initial versions ------------------------------------------------------------
INSERT INTO _schema_versions (users, roles, tickets, photos, profiles,
                              permissions, ticket_photos, user_roles,
                              user_tickets, user_photos, user_profiles,
                              role_permissions)
VALUES ('0.0.1', '0.0.1', '0.0.1', '0.0.1', '0.0.1', '0.0.1', '0.0.1', '0.0.1', '0.0.1',
        '0.0.1', '0.0.1', '0.0.1');

-------------------------------- TABLES ----------------------------------------

--- permissions ----------------------------------------------------------------
CREATE TYPE enum_permissions_access_type AS ENUM
    ('');

CREATE table if not exists permissions
(
    id         uuid                         not null default uuid_generate_v4(),
    resource   varchar(255)                 not null,
    action     enum_permissions_access_type not null,
    created_at timestamptz                  not null default now(),
    updated_at timestamptz                  not null default now(),
    deleted_at timestamptz,
    PRIMARY KEY (id)
);

CREATE or replace function touch_permissions_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language 'plpgsql';

CREATE trigger permissions_update_updated_at
    before update
    on permissions
    FOR EACH ROW
EXECUTE procedure touch_permissions_updated_at();
--------------------------------------------------------------------------------

--- roles ----------------------------------------------------------------------
CREATE TYPE enum_roles__type AS ENUM
    ('');

CREATE table if not exists roles
(
    id          uuid         not null default uuid_generate_v4(),
    name        varchar(255) not null,
    description text         not null,
    created_at  timestamptz  not null default now(),
    updated_at  timestamptz  not null default now(),
    deleted_at  timestamptz,
    PRIMARY KEY (id)
);

CREATE or replace function touch_roles_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language 'plpgsql';

CREATE trigger roles_update_updated_at
    before update
    on roles
    FOR EACH ROW
EXECUTE procedure touch_roles_updated_at();
--------------------------------------------------------------------------------

--- role_permissions -----------------------------------------------------------
CREATE TYPE enum_role_permissions__type AS ENUM
    ('');

CREATE table if not exists role_permissions
(
    id         uuid        not null default uuid_generate_v4(),
    role_id    uuid        not null default uuid_generate_v4(),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    PRIMARY KEY (id),

    CONSTRAINT fk_role_permissions_role_id_roles
        foreign key (role_id)
            references roles (id)
            ON DELETE CASCADE
);

CREATE or replace function touch_role_permissions_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language 'plpgsql';

CREATE trigger role_permissions_update_updated_at
    before update
    on role_permissions
    FOR EACH ROW
EXECUTE procedure touch_role_permissions_updated_at();
--------------------------------------------------------------------------------

--- users ----------------------------------------------------------------------
CREATE TYPE enum_users_status_type AS ENUM
    ('Pending', 'Regular', 'Privileged');

CREATE table if not exists users
(
    id         uuid        not null   default uuid_generate_v4(),
    status     enum_users_status_type default 'Pending'::enum_users_status_type,
    created_at timestamptz not null   default now(),
    updated_at timestamptz not null   default now(),
    deleted_at timestamptz,
    PRIMARY KEY (id)
);

CREATE or replace function touch_users_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language 'plpgsql';

CREATE trigger users_update_updated_at
    before update
    on users
    FOR EACH ROW
EXECUTE procedure touch_users_updated_at();
--------------------------------------------------------------------------------

--- user_ticket ----------------------------------------------------------------
CREATE table if not exists user_tickets
(
    id         bigserial   not null,
    user_id    uuid        not null,
    ticket_id  uuid        not null unique default uuid_generate_v4(),
    created_at timestamptz not null        default now(),
    updated_at timestamptz not null        default now(),
    deleted_at timestamptz,
    PRIMARY KEY (id),

    CONSTRAINT fk_user_tickets_user_id_user
        foreign key (user_id)
            references users (id)
            ON DELETE CASCADE
);

CREATE or replace function touch_user_tickets_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language 'plpgsql';

CREATE trigger user_tickets_update_updated_at
    before update
    on user_tickets
    FOR EACH ROW
EXECUTE procedure touch_user_tickets_updated_at();
--------------------------------------------------------------------------------

--- user_photos ----------------------------------------------------------------
CREATE table if not exists user_photos
(
    id         uuid        not null default uuid_generate_v4(),
    user_id    uuid        not null default uuid_generate_v4(),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    PRIMARY KEY (id),

    CONSTRAINT fk_user_photos_user_id_users
        foreign key (user_id)
            references users (id)
            ON DELETE CASCADE
);

CREATE or replace function touch_user_photos_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language 'plpgsql';

CREATE trigger user_photos_update_updated_at
    before update
    on user_photos
    FOR EACH ROW
EXECUTE procedure touch_user_photos_updated_at();
--------------------------------------------------------------------------------

--- user_profiles --------------------------------------------------------------
CREATE TYPE enum_user_profiles__type AS ENUM
    ('');

CREATE table if not exists user_profiles
(
    id         uuid        not null default uuid_generate_v4(),
    user_id    uuid        not null default uuid_generate_v4(),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    PRIMARY KEY (id),

    CONSTRAINT fk_user_profiles_user_id_users
        foreign key (user_id)
            references users (id)
            ON DELETE CASCADE
);

CREATE or replace function touch_user_profiles_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language 'plpgsql';

CREATE trigger user_profiles_update_updated_at
    before update
    on user_profiles
    FOR EACH ROW
EXECUTE procedure touch_user_profiles_updated_at();
--------------------------------------------------------------------------------

--- user_roles -----------------------------------------------------------------
CREATE TYPE enum_user_roles__type AS ENUM
    ('');

CREATE table if not exists user_roles
(
    id         uuid        not null default uuid_generate_v4(),
    user_id    uuid        not null default uuid_generate_v4(),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    PRIMARY KEY (id),

    CONSTRAINT fk_user_roles_user_id_users
        foreign key (user_id)
            references users (id)
            ON DELETE CASCADE
);

CREATE or replace function touch_user_roles_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language 'plpgsql';

CREATE trigger user_roles_update_updated_at
    before update
    on user_roles
    FOR EACH ROW
EXECUTE procedure touch_user_roles_updated_at();
--------------------------------------------------------------------------------


--- profiles -------------------------------------------------------------------------
CREATE TYPE enum_profiles__type AS ENUM
    ('');

CREATE table if not exists profiles
(
    id           uuid         not null default uuid_generate_v4(),
    user_id      uuid         not null default uuid_generate_v4(),
    service_name varchar(255) not null,
    active       bool,
    first_name   varchar(255),
    last_name    varchar(255),
    email        varchar(255),
    time_zone    varchar(255),
    avatar_url   text,
    created_at   timestamptz  not null default now(),
    updated_at   timestamptz  not null default now(),
    deleted_at   timestamptz,
    PRIMARY KEY (id),

    CONSTRAINT fk_profiles_user_id_users
        foreign key (id)
            references users (id)
            ON DELETE CASCADE
);

CREATE or replace function touch_profiles_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language 'plpgsql';

CREATE trigger profiles_update_updated_at
    before update
    on profiles
    FOR EACH ROW
EXECUTE procedure touch_profiles_updated_at();
--------------------------------------------------------------------------------

--- tickets --------------------------------------------------------------------
CREATE TYPE enum_tickets__type AS ENUM
    ('');

CREATE table if not exists tickets
(
    id         uuid        not null,
    owner_id   uuid        not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    PRIMARY KEY (id),

    CONSTRAINT fk_tickets_owner_id_users
        foreign key (owner_id)
            references users (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_tickets_id_user_tickets
        foreign key (id)
            references user_tickets (ticket_id)
            ON DELETE CASCADE
);

CREATE or replace function touch_tickets_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language 'plpgsql';

CREATE trigger tickets_update_updated_at
    before update
    on tickets
    FOR EACH ROW
EXECUTE procedure touch_tickets_updated_at();
--------------------------------------------------------------------------------

--- ticket_photos -------------------------------------------------------------------------
CREATE TYPE enum_ticket_photos__type AS ENUM
    ('');

CREATE table if not exists ticket_photos
(
    id         uuid        not null default uuid_generate_v4(),
    ticket_id  uuid        not null default uuid_generate_v4(),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    PRIMARY KEY (id),

    CONSTRAINT fk_ticket_photos_ticket_id_tickets
        foreign key (id)
            references tickets (id)
            ON DELETE CASCADE
);

CREATE or replace function touch_ticket_photos_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language 'plpgsql';

CREATE trigger ticket_photos_update_updated_at
    before update
    on ticket_photos
    FOR EACH ROW
EXECUTE procedure touch_ticket_photos_updated_at();
--------------------------------------------------------------------------------

--- photos -------------------------------------------------------------------------
CREATE TYPE enum_photos__type AS ENUM
    ('');

CREATE table if not exists photos
(
    id         uuid        not null default uuid_generate_v4(),
    owner_id   uuid        not null default uuid_generate_v4(),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    PRIMARY KEY (id),

    CONSTRAINT fk_photos_owner_id_users
        foreign key (id)
            references users (id)
            ON DELETE CASCADE
);

CREATE or replace function touch_photos_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language 'plpgsql';

CREATE trigger photos_update_updated_at
    before update
    on photos
    FOR EACH ROW
EXECUTE procedure touch_photos_updated_at();
--------------------------------------------------------------------------------
