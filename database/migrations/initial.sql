CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


create type account_user_status as enum ('ACTIVE');

alter type account_user_status
  owner to "user";

create type event_status as enum ('ACTIVE');

alter type event_status
  owner to "user";

create type discount_code_type as enum ('ACTIVE');

alter type discount_code_type
  owner to "user";

create type customer_status as enum ('ACTIVE');

alter type customer_status
  owner to "user";

create type attendee_status as enum ('ACTIVE');

alter type attendee_status
  owner to "user";

create type question_types as enum ('CHECKBOX', 'RADIO', 'TEXT', 'WAIVER');

alter type question_types
  owner to "user";

create table if not exists accounts
(
  id         serial not null
    constraint accounts_pkey
    primary key,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  email      text   not null,
  first_name text,
  last_name  text
);

alter table accounts
  owner to "user";

create unique index if not exists accounts_email_uindex
  on accounts (email);

create table if not exists users
(
  id         serial  not null
    constraint users_pkey
    primary key,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  email      text    not null,
  password   text    not null,
  first_name text    not null,
  last_name  text    not null,
  status     text    not null,
  account_id integer not null
    constraint users_accounts_id_fk
    references accounts
);

alter table users
  owner to "user";

create unique index if not exists users_email_uindex
  on users (email);

create table if not exists account_users
(
  id         serial              not null
    constraint pk_account_users
    primary key,
  first_name varchar(50)         not null,
  last_name  varchar(50)         not null,
  email      varchar(255)        not null,
  password   varchar(50)         not null,
  status     account_user_status not null,
  account_id integer             not null
    constraint fk_account_users_account_id
    references accounts,
  created_at timestamp           not null,
  updated_at timestamp           not null,
  deleted_at timestamp           not null
);

alter table account_users
  owner to "user";

create table if not exists events
(
  id          serial       not null
    constraint pk_events
    primary key,
  title       varchar(255) not null,
  account_id  integer      not null
    constraint fk_events_account_id
    references accounts
    on update cascade on delete cascade,
  user_id     integer      not null
    constraint fk_events_user_id
    references users
    on delete cascade,
  status      event_status,
  start_date  timestamp    not null,
  end_date    timestamp    not null,
  created_at  timestamp    not null,
  updated_at  timestamp    not null,
  deleted_at  timestamp,
  description text
);

alter table events
  owner to "user";

create table if not exists tickets
(
  id                          serial       not null
    constraint pk_tickets
    primary key,
  title                       varchar(255) not null,
  event_id                    integer      not null
    constraint fk_tickets_event_id
    references events
    on delete cascade,
  initital_quantity_available integer      not null,
  quantity_sold               integer      not null,
  created_at                  timestamp    not null,
  updated_at                  timestamp    not null,
  deleted_at                  timestamp    not null,
  sale_start_date             timestamp,
  sale_end_date               timestamp,
  max_per_transaction         integer
);

alter table tickets
  owner to "user";

create table if not exists transactions
(
  id             serial                      not null
    constraint pk_transactions
    primary key,
  event_id       integer                     not null
    constraint fk_transactions_event_id
    references events,
  total          numeric(12, 2) default 0.00 not null,
  total_tax      numeric(12, 2) default 0.00 not null,
  total_discount numeric(12, 2) default 0.00 not null,
  created_at     timestamp                   not null,
  updated_at     timestamp                   not null,
  deleted_at     timestamp                   not null,
  uuid           uuid default uuid_generate_v4(),
  metadata       jsonb,
  first_name     varchar(50),
  last_name      varchar(50),
  email          varchar(255),
  company_name   varchar(255)
);

alter table transactions
  owner to "user";

create unique index if not exists transactions_uuid_uindex
  on transactions (uuid);

create table if not exists transaction_items
(
  id             serial         not null
    constraint pk_transaction_items
    primary key,
  total          numeric(12, 2) not null,
  total_tax      numeric(12, 2) not null,
  total_discount numeric(12, 2) not null,
  quantity       integer        not null,
  transaction_id integer        not null
    constraint fk_transaction_items_transaction_id
    references transactions,
  ticket_id      integer        not null
    constraint fk_transaction_items_ticket_id
    references tickets
);

alter table transaction_items
  owner to "user";

create table if not exists discount_codes
(
  id         serial             not null
    constraint pk_discount_codes
    primary key,
  code       varchar(50)        not null,
  discount   numeric(12, 2)     not null,
  type       discount_code_type not null,
  created_at timestamp          not null,
  updated_at timestamp          not null,
  deleted_at timestamp          not null
);

alter table discount_codes
  owner to "user";

create table if not exists transaction_discount_codes
(
  id               serial         not null
    constraint pk_transaction_discount_codes
    primary key,
  transaction_id   integer        not null
    constraint fk_transaction_discount_codes_transaction_id
    references transactions,
  discount_code_id integer        not null
    constraint fk_transaction_discount_codes_discount_code_id
    references discount_codes,
  total_discount   numeric(12, 2) not null
);

alter table transaction_discount_codes
  owner to "user";

create table if not exists customers
(
  id             serial          not null
    constraint pk_customers
    primary key,
  first_name     varchar(50)     not null,
  last_name      varchar(50)     not null,
  email          varchar(100)    not null,
  password       varchar(50),
  status         customer_status not null,
  transaction_id integer         not null
    constraint fk_customers_transaction_id
    references transactions,
  event_id       integer         not null
    constraint fk_customers_event_id
    references events,
  created_at     timestamp       not null,
  updated_at     timestamp       not null,
  deleted_at     timestamp,
  company_name   varchar(150),
  metadata       jsonb
);

alter table customers
  owner to "user";

create table if not exists attendees
(
  id             serial          not null
    constraint attendees_pk
    primary key,
  transaction_id integer         not null
    constraint fk_attendees_transaction_id
    references transactions
    on delete cascade,
  ticket_id      integer         not null
    constraint fk_attendees_ticket_id
    references tickets
    on delete cascade,
  email          varchar(255)    not null,
  status         attendee_status not null,
  created_at     timestamp       not null,
  updated_at     timestamp       not null,
  deleted_at     timestamp       not null,
  event_id       integer         not null
    constraint attendees_events_id_fk
    references events
    on delete cascade
);

alter table attendees
  owner to "user";

create table if not exists questions
(
  id         serial                not null
    constraint pk_questions
    primary key,
  title      text                  not null,
  type       question_types        not null,
  created_at timestamp             not null,
  updated_at timestamp             not null,
  deleted_at timestamp             not null,
  required   boolean default false not null
);

alter table questions
  owner to "user";

create table if not exists question_answers
(
  id          serial    not null
    constraint pk_question_answers
    primary key,
  answer      text      not null,
  question_id integer   not null
    constraint fk_question_answers_question_id
    references questions,
  created_at  timestamp not null,
  updated_at  timestamp not null,
  deleted_at  timestamp not null
);

alter table question_answers
  owner to "user";

create table if not exists ticket_questions
(
  ticket_id   integer not null
    constraint fk_ticket_questions_ticket_id
    references tickets
    on delete cascade,
  question_id integer not null
    constraint fk_ticket_questions_question_id
    references questions
    on delete cascade,
  constraint ticket_questions_pk
  primary key (ticket_id, question_id)
);

alter table ticket_questions
  owner to "user";

create table if not exists event_questions
(
  event_id    integer not null
    constraint fk_event_questions_event_id
    references events,
  question_id integer not null
    constraint fk_event_questions_question_id
    references questions,
  constraint event_questions_pk
  primary key (event_id, question_id)
);

alter table event_questions
  owner to "user";

create table if not exists attributes
(
  id         serial      not null
    constraint pk_attributes
    primary key,
  name       varchar(50) not null,
  value      text        not null,
  type       varchar(20) not null,
  created_at timestamp   not null,
  updated_at timestamp   not null,
  deleted_at timestamp   not null
);

alter table attributes
  owner to "user";

create table if not exists event_attributes
(
  event_id     integer not null
    constraint fk_event_attributes_event_id
    references events
    on delete cascade,
  attribute_id integer not null
    constraint fk_event_attributes_attribute_id
    references attributes
    on delete cascade,
  constraint event_attributes_pk
  primary key (event_id, attribute_id)
);

alter table event_attributes
  owner to "user";

create table if not exists ticket_attributes
(
  ticket_id    integer not null
    constraint fk_ticket_attributes_ticket_id
    references tickets
    on delete cascade,
  attribute_id integer not null
    constraint fk_ticket_attributes_attribute_id
    references attributes
    on delete cascade,
  constraint ticket_attributes_pk
  primary key (ticket_id, attribute_id)
);

alter table ticket_attributes
  owner to "user";

create table if not exists transaction_attributes
(
  transaction_id integer not null
    constraint fk_transaction_attributes_transaction_id
    references transactions,
  attribute_id   integer not null
    constraint fk_transaction_attributes_attribute_id
    references attributes,
  constraint transaction_attributes_pk
  primary key (transaction_id, attribute_id)
);

alter table transaction_attributes
  owner to "user";

create table if not exists question_options
(
  id          serial    not null
    constraint question_options_pkey
    primary key,
  title       text      not null,
  question_id integer   not null
    constraint question_options_questions_id_fk
    references questions
    on delete cascade,
  created_at  timestamp not null,
  updated_at  timestamp
);

alter table question_options
  owner to "user";

create unique index if not exists question_options_id_uindex
  on question_options (id);

create table if not exists ticket_reservations
(
  id              serial    not null
    constraint ticket_reservations_pkey
    primary key,
  ticket_id       integer   not null
    constraint ticket_reservations_tickets_id_fk
    references tickets
    on delete cascade,
  transaction_id  integer   not null
    constraint ticket_reservations_transactions_id_fk
    references transactions
    on delete cascade,
  ticket_quantity integer   not null,
  reserved_until  timestamp not null
);

alter table ticket_reservations
  owner to "user";

create unique index if not exists ticket_reservations_id_uindex
  on ticket_reservations (id);

create function uuid_nil()
immutable
strict
parallel safe
language c
as
-- missing source code for uuid_nil
;

alter function uuid_nil()
  owner to "user";

create function uuid_ns_dns()
immutable
strict
parallel safe
language c
as
-- missing source code for uuid_ns_dns
;

alter function uuid_ns_dns()
  owner to "user";

create function uuid_ns_url()
immutable
strict
parallel safe
language c
as
-- missing source code for uuid_ns_url
;

alter function uuid_ns_url()
  owner to "user";

create function uuid_ns_oid()
immutable
strict
parallel safe
language c
as
-- missing source code for uuid_ns_oid
;

alter function uuid_ns_oid()
  owner to "user";

create function uuid_ns_x500()
immutable
strict
parallel safe
language c
as
-- missing source code for uuid_ns_x500
;

alter function uuid_ns_x500()
  owner to "user";

create function uuid_generate_v1()
strict
parallel safe
language c
as
-- missing source code for uuid_generate_v1
;

alter function uuid_generate_v1()
  owner to "user";

create function uuid_generate_v1mc()
strict
parallel safe
language c
as
-- missing source code for uuid_generate_v1mc
;

alter function uuid_generate_v1mc()
  owner to "user";

create function uuid_generate_v3(namespace uuid, name text)
immutable
strict
parallel safe
language c
as
-- missing source code for uuid_generate_v3
;

alter function uuid_generate_v3(uuid, text)
  owner to "user";

create function uuid_generate_v4()
strict
parallel safe
language c
as
-- missing source code for uuid_generate_v4
;

alter function uuid_generate_v4()
  owner to "user";

create function uuid_generate_v5(namespace uuid, name text)
immutable
strict
parallel safe
language c
as
-- missing source code for uuid_generate_v5
;

alter function uuid_generate_v5(uuid, text)
  owner to "user";

