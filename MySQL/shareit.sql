-- ユーザ
DROP TABLE IF EXISTS users;

CREATE TABLE users(
  user_id INT UNSIGNED NOT NULL PRIMARY KEY,
  user_name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  interested_topics VARCHAR(255),
  created_date DATETIME NOT NULL,
  updated_date DATETIME NOT NULL,
  deleted_date DATETIME 
);

-- トピック
DROP TABLE IF EXISTS topics;

CREATE TABLE topics(
topic_id INT UNSIGNED NOT NULL PRIMARY KEY,
topic_name VARCHAR(255) NOT NULL,
proposed_user_id INT UNSIGNED NOT NULL,
created_date DATETIME NOT NULL,
updated_date DATETIME NOT NULL,
deleted_date DATETIME 
);

-- 記事
DROP TABLE IF EXISTS articles;

create table articles(
article_id int unsigned not null primary key,
article_title varchar(255) not null,
created_user_id int unsigned not null,
article_content varchar(1000) not null,
article_topics varchar(255) not null,
created_date datetime not null,
updated_date datetime not null,
deleted_date datetime 
);
