create database restaurant;
use restaurant;

create table PRODUCTS (
  id int not null auto_increment,
  name varchar(100) not null,
  price decimal(10,2) not null,
  primary key(id)
);

insert into PRODUCTS (name, price) values
('Margherita Pizza', 12.99),
('Pepperoni Pizza', 14.99),
('Hawaiian Pizza', 13.99),
('Veggie Pizza', 11.99),
('Chicken Alfredo', 16.99),
('Spaghetti Bolognese', 15.99),
('Caesar Salad', 8.99),
('Greek Salad', 9.99),
('Grilled Chicken Sandwich', 10.99),
('Hamburger', 9.99),
('French Fries', 3.99),
('Onion Rings', 4.99),
('Garlic Bread', 4.99),
('Chocolate Brownie', 5.99),
('Tiramisu', 7.99);