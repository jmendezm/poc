CREATE TYPE record_status AS ENUM
    ('app.status.Active', 'app.status.Deleted', 'app.status.Inactive');

create table IF NOT EXISTS users (
    account_id BIGINT NOT NULL,
    user_id bigserial NOT NULL,
    create_date timestamp NOT NULL,
    last_updated timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    record_status "record_status" DEFAULT 'app.status.Active'::"record_status",
    active "record_status" DEFAULT 'app.status.Active'::"record_status",
    identification varchar(50) NULL,
    password varchar(250) NULL,
    company_name varchar(50),
    first_name varchar(50),
    last_name varchar(50),
    email varchar(50),
    phone varchar(50),
    emergency_phone varchar(50),
    i18n varchar(50),
    address varchar(200),
    auth_menu varchar NULL DEFAULT '[]'::text,
    auth_keys varchar NULL DEFAULT '[]'::text,
    auth_groups varchar NULL DEFAULT '[]'::text
);

create table IF NOT EXISTS connections (
    id BIGSERIAL NOT NULL,
    connection_id varchar(50) NOT NULL,
    account_id BIGINT,
    user_id BIGINT,
    create_date timestamp NOT NULL,
    last_updated timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    record_status "record_status" DEFAULT 'app.status.Active'::"record_status",
    active "record_status" DEFAULT 'app.status.Active'::"record_status",
    connected varchar(50),
    disconnected varchar(50),
    user_data varchar DEFAULT '{}'::text NULL,
    account_data varchar DEFAULT '{}'::text NULL,
    auth_menu varchar NULL DEFAULT '[]'::text,
    auth_keys varchar NULL DEFAULT '[]'::text,
    auth_groups varchar NULL DEFAULT '[]'::text
);

create table IF NOT EXISTS sites (
    account_id BIGINT NOT NULL,
    site_id BIGSERIAL NOT NULL,
    site_name varchar(50) NOT NULL,
    create_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    record_status "record_status" DEFAULT 'app.status.Active'::"record_status",
    active "record_status" DEFAULT 'app.status.Active'::"record_status",
    des varchar(200),
    description varchar,
    operate_by varchar(50),
    logo varchar,
    rules_documents varchar DEFAULT '{}'::text NULL,
    services_amenities varchar DEFAULT '{}'::text NULL,
    type varchar(50) DEFAULT 'Yard'::character varying,
    email varchar(50) NULL,
    phone varchar(50) NULL,
    address varchar(200) NULL,
    website varchar(200) NULL,
    geolocation varchar(200) NULL
);