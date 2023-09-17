CREATE DATABASE rental_mobil;

use rental_mobil;

CREATE TABLE users(
    id bigint auto_increment primary key,
    name varchar(100) not null,
    email varchar(100) not null ,
    no_hp varchar(16) not null,
    password varchar(50) not null
) engine = InnoDB;

CREATE table lease_types(
    id bigint auto_increment primary key,
    title varchar(100) not null,
    description text not null
) engine = InnoDB;

CREATE TABLE cars(
     id bigint auto_increment primary key,
     title varchar(100) not null,
     duration varchar(50) not null,
     image_url varchar(255) not null,
     description text not null,
     passenger tinyint not null,
    luggage tinyint not null,
    car_type varchar(20) not null,
    is_driver boolean not null
)engine = InnoDB;

CREATE TABLE car_lease_type(
   id bigint auto_increment primary key,
   lease_type_id bigint not null,
   car_id bigint not null,
   CONSTRAINT pk_car_lease_type_lease_type FOREIGN KEY (lease_type_id)
       REFERENCES lease_types(id) ON DELETE CASCADE,
    CONSTRAINT pk_car_lease_type_car FOREIGN KEY (car_id)
                           REFERENCES cars (id) ON DELETE CASCADE
)engine = InnoDB;

CREATE TABLE tours(
  id bigint auto_increment primary key,
  title varchar(150) not null ,
  price int not null,
  duration varchar(20) not null,
  description text not null
)engine = InnoDB;

CREATE TABLE car_tours(
  id bigint auto_increment primary key,
  car_id bigint not null,
  tour_id bigint not null,
  CONSTRAINT fk_car_tour_car FOREIGN KEY (car_id)
              REFERENCES cars (id) ON DELETE CASCADE,
  CONSTRAINT fk_car_tour FOREIGN KEY (tour_id)
          REFERENCES tours(id) ON DELETE CASCADE
)engine = InnoDB;

CREATE TABLE car_rating(
    id bigint auto_increment primary key,
    car_id bigint not null,
    user_id bigint not null,
    CONSTRAINT fk_car_rating_car FOREIGN KEY (car_id)
       REFERENCES cars (id) ON DELETE CASCADE,
    CONSTRAINT fk_car_rating_user FOREIGN KEY (user_id)
       REFERENCES users(id) ON DELETE CASCADE
)engine = InnoDB;

CREATE TABLE transactions(
     id bigint auto_increment primary key,
     car_id bigint not null,
     user_id bigint not null,
     price int not null,
     CONSTRAINT fk_transaction_car FOREIGN KEY (car_id)
         REFERENCES cars (id) ON DELETE CASCADE,
     CONSTRAINT fk_transaction_user FOREIGN KEY (user_id)
         REFERENCES users(id) ON DELETE CASCADE
)engine = InnoDB
