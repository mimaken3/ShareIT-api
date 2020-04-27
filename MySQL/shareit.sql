-- ユーザ
DROP TABLE IF EXISTS users;

CREATE TABLE users(
  user_id INT UNSIGNED NOT NULL PRIMARY KEY,
  user_name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_date DATETIME NOT NULL,
  updated_date DATETIME NOT NULL,
  deleted_date DATETIME NOT NULL, 
  is_deleted TINYINT(1) NOT NULL DEFAULT '0'
);

-- トピック
DROP TABLE IF EXISTS topics;

CREATE TABLE topics(
  topic_id INT UNSIGNED NOT NULL PRIMARY KEY,
  topic_name VARCHAR(255) NOT NULL,
  proposed_user_id INT UNSIGNED NOT NULL,
  created_date DATETIME NOT NULL,
  updated_date DATETIME NOT NULL,
  deleted_date DATETIME NOT NULL, 
  is_deleted TINYINT(1) NOT NULL DEFAULT '0'
);

-- 記事
DROP TABLE IF EXISTS articles;

create table articles(
  article_id INT UNSIGNED NOT NULL PRIMARY KEY,
  created_user_id INT unsigned NOT NULL,
  article_title VARCHAR(255) NOT NULL,
  article_content VARCHAR(1000) NOT NULL,
  created_date DATETIME NOT NULL,
  updated_date DATETIME NOT NULL,
  deleted_date DATETIME NOT NULL, 
  is_deleted TINYINT(1) NOT NULL DEFAULT '0'
);

-- 記事とトピックと紐付けるテーブル
DROP TABLE IF EXISTS article_topics;

CREATE TABLE article_topics(
  article_topic_id INT UNSIGNED NOT NULL PRIMARY KEY,
  article_id INT UNSIGNED NOT NULL,
  topic_id INT UNSIGNED NOT NULL
);

-- ユーザの興味のあるトピック
DROP TABLE IF EXISTS user_interested_topics;

CREATE TABLE user_interested_topics(
  user_interested_topics_id INT UNSIGNED NOT NULL PRIMARY KEY,
  user_id INT UNSIGNED NOT NULL,
  topic_id INT UNSIGNED NOT NULL
);


-- ユーザのプロフィール
DROP TABLE IF EXISTS profiles;

CREATE TABLE profiles(
  profile_id INT UNSIGNED NOT NULL PRIMARY KEY,
  user_id INT UNSIGNED NOT NULL,
  content VARCHAR(1000),
  is_deleted TINYINT(1) NOT NULL DEFAULT '0'
);

-- アイコン
DROP TABLE IF EXISTS icons;

CREATE TABLE icons(
  icon_id INT UNSIGNED NOT NULL PRIMARY KEY,
  user_id INT UNSIGNED NOT NULL,
  icon_name VARCHAR(255) NOT NULL
);

