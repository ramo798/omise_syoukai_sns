DROP TABLE IF EXISTS employee;

CREATE TABLE employee
(
  id           INT(10),
  name     VARCHAR(40)
);

INSERT INTO employee (id, name) VALUES (1, "Nagaoka");
INSERT INTO employee (id, name) VALUES (2, "Tanaka");


DROP TABLE IF EXISTS restaurant_info;

CREATE TABLE restaurant_info
(
  id           VARCHAR(40) primary key,
  name     VARCHAR(40),
  address VARCHAR(80),
  logo_image_url  VARCHAR(200),
  station_name VARCHAR(40)
);


INSERT INTO restaurant_info (id, name,address,logo_image_url,station_name) VALUES ("J0000", "鳥貴族","pa-kus2121u","htttp/11111","なんば");
INSERT INTO restaurant_info (id, name,address,logo_image_url,station_name) VALUES ("J01200", "鳥貴族1212","pa-k1212usu","htttp/1111121","なん1ば");
INSERT INTO restaurant_info (id, name,address,logo_image_url,station_name) VALUES ("J0002220", "鳥貴族212","pa-k1212usu","htttp/12121111","な1111んば");


DROP TABLE IF EXISTS post_list;

CREATE TABLE post_list
(
  id INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
  users_id VARCHAR(30),
  comment VARCHAR(240),
  restaurant_id VARCHAR(100),
  price INT,
  assessment INT
);
INSERT INTO post_list (users_id,comment,restaurant_id,price,assessment) VALUES ("ramo123", "めちゃおいしいです。","J0000",4500,80);
INSERT INTO post_list (users_id,comment,restaurant_id,price,assessment) VALUES ("ramo456", "めちゃおいしいです。sasasa","J0000",4500,80);
INSERT INTO post_list (users_id,comment,restaurant_id,price,assessment) VALUES ("ramo789", "めちゃおいしいです。sasasa","J0000",4500,80);
