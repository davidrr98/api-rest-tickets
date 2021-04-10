# api-rest-tickets
API con un CRUD para origenes


## Base de datos
Se usa postgresql

### Script para crear la tabla
```shell
CREATE TABLE ticket (
   id serial primary key not null,
   usuario varchar(20) not null,
   estatus boolean not null,
   fecha_creacion timestamp not null,
   fecha_modificacion timestamp not null
);
```

### Configurar conexion 
    en el archivo bd/conexionBD.go estan las constantes para acceder a la BD



## Rutas de acceso
* http://localhost:8080/ticket :GET obtener todos los tickets

* http://localhost:8080/ticket?usuario=usuario&estatus=false&id=0&limit=-1  parametros del Query

* http://localhost:8080/ticket/id :GET obtener ticket con alguna id

* http://localhost:8080/ticket  :POST guardar ticket 

recibe un JSON
{
    "usuario": "usuario",
    "fecha_creacion": "",
    "fecha_modificacion": "",
    "estatus": true
}

* http://localhost:8080/ticket  :PUT Edita un ticket 

recibe un JSON
{
    "id": 4,
    "usuario": "usuario",
    "fecha_creacion": "",
    "fecha_modificacion": "",
    "estatus": true
}
 
 * http://localhost:8080/ticket/id :DELETE Elimina ticket con alguna id
