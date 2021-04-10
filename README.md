# api-rest-tickets
API con un CRUD para origenes


## Base de datos
Se usa postgresql

### Script para crear la tabla
```shell
CREATE TABLE IF NOT EXISTS ticket (
   id serial primary key not null,
   usuario varchar(20) not null,
   estatus boolean not null,
   fecha_creacion timestamp not null,
   fecha_modificacion timestamp not null
);
```
este se ejecutara automaticamente al iniciar la aplicacion

### Configurar conexion 
    En el archivo bd/conexionBD.go estan las constantes para acceder a la BD
    
    ```shell
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "prueba_tecnica"
)
```



## Rutas de acceso
* http://localhost:8080/ticket :GET obtener todos los tickets

* http://localhost:8080/ticket?usuario=usuario&estatus=false&id=0&limit=-1  parametros del Query

* http://localhost:8080/ticket/id :GET obtener ticket con alguna id

* http://localhost:8080/ticket  :POST guardar ticket 

recibe un JSON
```shell
{
    "usuario": "usuario",
    "fecha_creacion": "",
    "fecha_modificacion": "",
    "estatus": true
}
```

* http://localhost:8080/ticket  :PUT Edita un ticket 

recibe un JSON
```shell
{
    "id": 4,
    "usuario": "usuario",
    "fecha_creacion": "",
    "fecha_modificacion": "",
    "estatus": true
}
````
 * http://localhost:8080/ticket/id :DELETE Elimina ticket con alguna id
