CREATE SCHEMA IF NOT EXISTS sid_ci;

CREATE TABLE IF NOT EXISTS sid_ci.clients (
    id int generated always as identity primary key,
    identifier text unique,
    password text,
    created_at timestamp,
    last_communication_at timestamp,
    last_status text,
    active boolean,
    CONSTRAINT clients_pk PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS sid_ci.repos (
    id int generated always as identity primary key,
    name text not null,
    ssh_url text not null,
    enabled boolean not null,
    added_by int references sid_ci.clients(id),
    added_at timestamp,
    CONSTRAINT repos_pk PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS sid_ci.job (
    id int generated always as identity primary key,
    succeeded boolean,
    image_digest text,
    git_hexsha text,
    job_uuid text,
    run_by int references sid_ci.clients(id),
    repo int references sid_ci.repos(id),
    received_at timestamp
);

CREATE TABLE IF NOT EXISTS sid_ci.run_event (
    event_type text,
    content text,
    event_at timestamp,
    job_id int references sid_ci.job(id)
)


