create table ceps (
    id bigint not null primary key,
    city varchar(100) not null,
    state varchar(30) not null,
    uf varchar(2) not null,
    logradouro varchar(200) not null,
    neighborhood varchar(50) not null,
    address varchar(200) not null,
    complement varchar(200) not null,
    value varchar(8) not null
);
create index idx_search on ceps(value);
