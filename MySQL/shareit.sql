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

CREATE TABLE articles(
article_id INT UNSIGNED NOT NULL PRIMARY KEY,
article_title VARCHAR(255) NOT NULL,
created_user_id INT UNSIGNED NOT NULL,
article_content VARCHAR(1000) NOT NULL,
article_topics VARCHAR(255) NOT NULL,
created_date DATETIME NOT NULL,
updated_date DATETIME NOT NULL,
deleted_date DATETIME 
);
