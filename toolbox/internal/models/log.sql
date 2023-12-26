CREATE
TABLE IF NOT EXISTS user_log
(
    `id`      int(10) unsigned NOT NULL AUTO_INCREMENT,
    `ip`  varchar(250)      NOT NULL default '' COMMENT '用户ip',
    `action`    varchar(60)   NOT NULL default '' COMMENT '用户动作',
    `action_user` varchar(60)      NOT NULL default '' COMMENT '操作人',
    `create_time` timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY `idx_ip` (`ip`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='用户日志表';

CREATE
TABLE IF NOT EXISTS life_config (
                               `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
                               `config_id` int(10) NOT NULL DEFAULT '0' COMMENT '相册id',
                               `img_url` varchar(500) NOT NULL DEFAULT '' COMMENT '图片路由',
                               `text` text COMMENT '文案',
                               `status` tinyint(3) NOT NULL DEFAULT '1' COMMENT '开启状态 1：未启用 2: 启用',
                               `horizontal_version` int(10) NOT NULL DEFAULT '0' COMMENT '是否为横向排版 0：否 1：是',
                               `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
                               `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

