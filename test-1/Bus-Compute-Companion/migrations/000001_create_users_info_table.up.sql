CREATE TABLE IF NOT EXISTS users_info (
  id bigserial PRIMARY KEY,
  fname varchar(255),
  lname varchar(255),
  email varchar(255),
  addres varchar(255), --name change probably
  phone_number varchar(255),
  passwrd varchar(255) --name change probably
);