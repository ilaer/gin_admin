create database gin_admin default character set  utf8 collate utf8_general_ci;

CREATE TABLE `ga_user` (
                        `id` int(11) NOT NULL AUTO_INCREMENT,
                        `username` char(255) NOT NULL COMMENT '账号',
                        `password` char(255) NOT NULL COMMENT '密码',
                        `email` varchar(255) NOT NULL COMMENT '邮箱',
                        `active` tinyint(1) NOT NULL COMMENT '账号状态,0未激活,1激活',
                        `create_ip` varchar(50) NOT NULL COMMENT '注册ip',
                        `login_ip` varchar(50) NOT NULL COMMENT '最后登录',
                        `login_time` int(11) DEFAULT NULL COMMENT '登录时间',
                        `create_time` datetime DEFAULT NULL,
                        `update_time` datetime DEFAULT NULL,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `gin_admin`.`ga_user` (`id`, `username`, `password`, `email`, `active`, `create_ip`, `login_ip`, `login_time`, `create_time`, `update_time`) VALUES ('1', 'admin', '123', 'admin@gin_admin.com', '1', '127.0.0.1', '127.0.0.1', '0', '2019-09-23 17:17:57', '2019-09-23 17:18:00');


CREATE TABLE `ga_auth` (
                             `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                             `username` varchar(50) DEFAULT '' COMMENT '账号',
                             `password` varchar(50) DEFAULT '' COMMENT '密码',
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `gin_admin`.`ga_auth` (`id`, `username`, `password`) VALUES (null, 'admin', '123');