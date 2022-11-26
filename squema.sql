create database inventory;
use inventory;

-- Para esta base de datos utilizaremos DOCKER para instalar una base de datos en nuestra pc
-- Lo mas recomendado es usar una imagen ene specifico
-- docker.hub buscar las imagenes de sql que vayamos a utilizar en este caso MariaDB
-- docker pull mariadb:10.7.4
-- docker run -d --name mariadb -p 3306:3306 --env MARIADB_ROOT_PASSWORD=root mariadb:10.7.5
-- necesito crear un contener de docker en base a esa configuracion
-- este comando nos permite crear un contenedor con la imagen de mariadb, los puertos se dan
-- de la siguiente manera: 3306:3306, el primero es el puerto de la pc y el segundo es el puerto del contenedor
-- docker logs mariadb
-- docker ps
-- docker stop mariadb 
-- esto detiene el contenedor de docker
-- docker start mariadb
-- esto inicia el contenedor de docker

create table USERS (
    id int not null auto_increment,
    email varchar(255) not null,
    name varchar(255) not null,
    password varchar(255) not null,
    primary key(id)
);

create table PRODUCT (
    id int not null auto_increment,
    name varchar(255) not null,
    description varchar(255) not null,
    price float not null,
    created_by int not null,
    primary key (id),
    foreign key (created_by) references USERS (id)
)

create table ROLES (
    id int not null auto_increment,
    name varchar(255) not null,
    primary key (id)
);

create table USER_ROLES (
    id int not null auto_increment,
    user_id int not null,
    role_id int not null,
    primary key (id),
    foreign key (user_id) references USERS (id),
    foreign key (role_id) references ROLES (id)
);

