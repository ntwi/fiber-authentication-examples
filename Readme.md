# Auth Examples

```bash

docker run --name local-postgres -p 5432:5432 -e POSTGRES_PASSWORD=MyP@ssw0rd -d postgres

```

Run the following query on PostgreSQL

```sql

create table users
(
    email             text not null,
    password_hash     text not null,
    id                serial
        primary key,
    username          text,
    referral_username text
);

alter table users
    owner to postgres;


```

Adjust and rename .env.example to .env and run `make`
