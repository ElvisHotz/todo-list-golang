## init

- go mod init todo-list

## packages

- go get github.com/kataras/iris/v12@master
- go get github.com/jinzhu/gorm
- go get github.com/go-sql-driver/mysql
- go get github.com/joho/godotenv
- go get github.com/spf13/cast


## database

```
    DROP DATABASE IF EXISTS `todo_list`;
    CREATE DATABASE IF NOT EXISTS `todo_list` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_bin */;
    USE `todo_list`;

    -- Copiando estrutura para tabela todo_list.todos
    DROP TABLE IF EXISTS `todos`;
    CREATE TABLE IF NOT EXISTS `todos` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    `title` varchar(255) COLLATE utf8_bin DEFAULT NULL,
    `description` varchar(255) COLLATE utf8_bin DEFAULT NULL,
    `date_for_finished` datetime DEFAULT NULL,
    `finished` tinyint(1) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_todos_deleted_at` (`deleted_at`)
    ) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

```
