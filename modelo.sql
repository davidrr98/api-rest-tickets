CREATE TABLE ticket (
   id serial primary key not null,
   usuario varchar(20) not null,
   estatus boolean not null,
   fecha_creacion timestamp not null,
   fecha_modificacion timestamp not null
);
