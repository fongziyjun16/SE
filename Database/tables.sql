CREATE TABLE ADMIN (
    ID INT NOT NULL AUTO_INCREMENT,
    USERNAME VARCHAR(20) NOT NULL,
    PASSWORD VARCHAR(20) NOT NULL,
    SALT VARCHAR(6) NOT NULL,
    PRIMARY KEY (ID)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE USER (
    ID INT NOT NULL AUTO_INCREMENT,
    USERNAME VARCHAR(20) NOT NULL,
    PASSWORD VARCHAR(20) NOT NULL,
    SALT VARCHAR(6) NOT NULL,
    NICKNAME VARCHAR(20) NOT NULL,
    BIRTHDAY DATE NOT NULL,
    GENDER TINYINT,
    DEPARTMENT VARCHAR(100),
    PRIMARY KEY (ID)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE SPACE (
    ID INT NOT NULL AUTO_INCREMENT,
    USERNAME VARCHAR(20) NOT NULL,
    CAPACITY DOUBLE(6,2) NOT NULL DEFAULT(10.00),
    USED DOUBLE(6,2) NOT NULL DEFAULT(0),
    PRIMARY KEY (ID)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE FOLLOW (
    FOLLOWEE VARCHAR(20) NOT NULL,
    FOLLOWER VARCHAR(20) NOT NULL,
    PRIMARY KEY (FOLLOWEE, FOLLOWER)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE COMMUNITY (
    ID INT NOT NULL AUTO_INCREMENT,
    CREATOR VARCHAR(20) NOT NULL,
    NAME VARCHAR(20) NOT NULL UNIQUE,
    DESCRIPTION VARCHAR(500),
    NUM_MEMBER INT NOT NULL DEFAULT(0),
    CREATE_TIME DATETIME NOT NULL,
    PRIMARY KEY (ID)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE COMMUNITY_MEMBER (
    COMMUNITY_ID INT NOT NULL,
    MEMBER VARCHAR(20) NOT NULL,
    PRIMARY KEY (COMMUNITY_ID, MEMBER)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE ARTICLE_TYPE (
    ID INT NOT NULL AUTO_INCREMENT,
    NAME VARCHAR(20) NOT NULL UNIQUE,
    DESCRIPTION VARCHAR(500),
    PRIMARY KEY (ID)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE ARTICLE (
    ID INT NOT NULL AUTO_INCREMENT,
    USERNAME VARCHAR(20) NOT NULL,
    TITLE VARCHAR(50) NOT NULL,
    TYPE_ID INT,
    COMMUNITY_ID INT,
    CREATE_TIME DATETIME NOT NULL,
    CONTENT TEXT NOT NULL,
    NUM_COMMENT INT NOT NULL DEFAULT(0),
    NUM_LIKE INT NOT NULL DEFAULT(0),
    NUM_FAVOURITE INT NOT NULL DEFAULT(0),
    POINT INT NOT NULL DEFAULT(0),
    PRIMARY KEY (ID)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE EMO_ARTICLE (
    ID INT NOT NULL AUTO_INCREMENT,
    ARTICLE_ID INT NOT NULL,
    USERNAME VARCHAR(20) NOT NULL,
    STATUS TINYINT NOT NULL,
    PRIMARY KEY (ID)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE COMMENT (
    ID INT NOT NULL AUTO_INCREMENT,
    USERNAME VARCHAR(20) NOT NULL,
    ARTICLE_ID INT NOT NULL,
    COMMENT_ID INT,
    CONTENT TEXT NOT NULL,
    PRIMARY KEY (ID)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


