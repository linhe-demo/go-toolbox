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