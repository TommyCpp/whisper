create table `group`
(
	id VARCHAR(128) not null
		primary key,
	groupname VARCHAR(128) null
)
;

create table user
(
	id INT(10) not null
		primary key,
	username VARCHAR(128) not null,
	password VARCHAR(128) not null,
	constraint user_username_uindex
	unique (username)
)
;

create table user_group
(
	user_id INT(10) not null,
	group_id VARCHAR(128) not null,
	primary key (user_id, group_id),
	constraint user_group_group_id_fk
	foreign key (group_id) references `group` (id)
		on delete cascade,
	constraint user_group_user_id_fk
	foreign key (user_id) references user (id)
		on update cascade on delete cascade
)
;

