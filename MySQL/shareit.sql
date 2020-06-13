-- ユーザ
DROP TABLE IF EXISTS users;

CREATE TABLE users(
  user_id INT UNSIGNED NOT NULL PRIMARY KEY,
  user_name VARCHAR(255) character set utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  email VARCHAR(255),
  password VARCHAR(255) NOT NULL,
  created_date DATETIME NOT NULL,
  updated_date DATETIME NOT NULL,
  deleted_date DATETIME NOT NULL, 
  is_deleted TINYINT(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
;

-- 記事
DROP TABLE IF EXISTS articles;

CREATE table articles(
  article_id INT UNSIGNED NOT NULL PRIMARY KEY,
  created_user_id INT unsigned NOT NULL,
  icon_name VARCHAR(500) NOT NULL,
  article_title VARCHAR(255) character set utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  article_content VARCHAR(9999) character set utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  created_date DATETIME NOT NULL,
  updated_date DATETIME NOT NULL,
  deleted_date DATETIME NOT NULL, 
  is_private TINYINT(1) NOT NULL,
  is_deleted TINYINT(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
;

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
  content VARCHAR(999) character set utf8mb4 COLLATE utf8mb4_bin,
  is_deleted TINYINT(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 
;

-- アイコン
DROP TABLE IF EXISTS icons;

CREATE TABLE icons(
  icon_id INT UNSIGNED NOT NULL PRIMARY KEY,
  user_id INT UNSIGNED NOT NULL,
  icon_name VARCHAR(500) NOT NULL
);

-- いいね
DROP TABLE IF EXISTS likes;

CREATE TABLE likes(
  like_id INT UNSIGNED NOT NULL PRIMARY KEY,
  user_id INT UNSIGNED NOT NULL,
  article_id INT UNSIGNED NOT NULL
  notification_id INT UNSIGNED,
);

-- コメント
DROP TABLE IF EXISTS comments;

CREATE TABLE comments(
  comment_id INT UNSIGNED NOT NULL PRIMARY KEY,
  article_id INT UNSIGNED NOT NULL,
  user_id INT UNSIGNED NOT NULL,
  content VARCHAR(999) character set utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  notification_id INT UNSIGNED,
  created_date DATETIME NOT NULL,
  updated_date DATETIME NOT NULL,
  deleted_date DATETIME NOT NULL, 
  is_deleted TINYINT(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
;

-- 通知
DROP TABLE IF EXISTS notifications;

CREATE TABLE notifications(
  notification_id INT UNSIGNED NOT NULL PRIMARY KEY,
  user_id INT UNSIGNED NOT NULL,
  source_user_id INT UNSIGNED NOT NULL,
  source_user_icon_name VARCHAR(500) NOT NULL,
  is_read TINYINT(1) NOT NULL DEFAULT '0',
  destination_id INT UNSIGNED NOT NULL,
  created_date DATETIME NOT NULL
);

-- 通知元名
DROP TABLE IF EXISTS destination_types;

CREATE TABLE destination_types(
  destination_type_id INT UNSIGNED NOT NULL PRIMARY KEY,
  destination_type_name VARCHAR(255) NOT NULL
);
INSERT INTO destination_types VALUES(1, '記事詳細画面');
INSERT INTO destination_types VALUES(2, '記事一覧画面');
INSERT INTO destination_types VALUES(3, 'ユーザ詳細画面');
INSERT INTO destination_types VALUES(4, 'ユーザ一覧画面');

-- 通知元情報 
DROP TABLE IF EXISTS destinations;

CREATE TABLE destinations(
  destination_id INT UNSIGNED NOT NULL PRIMARY KEY,
  destination_type_id INT UNSIGNED NOT NULL,
  destination_type_name_id INT UNSIGNED NOT NULL,
  behavior_type_id INT UNSIGNED NOT NULL,
  behavior_type_name_id INT UNSIGNED NOT NULL
);

-- 通知することになったユーザの動作
DROP TABLE IF EXISTS behavior_types;

CREATE TABLE behavior_types(
  behavior_type_id INT UNSIGNED NOT NULL,
  behavior_type_name VARCHAR(255) NOT NULL
);
INSERT INTO behavior_types VALUES(1, 'いいね');
INSERT INTO behavior_types VALUES(2, 'コメント書き込み');

