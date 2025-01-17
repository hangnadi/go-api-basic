create table demo.org
(
    org_id           uuid                     not null,
    org_extl_id      varchar                  not null,
    org_name         varchar                  not null,
    org_description  varchar                  not null,
    genesis_org      boolean                  not null,
    create_app_id    uuid                     not null,
    create_user_id   uuid,
    create_timestamp timestamp with time zone not null,
    update_app_id    uuid                     not null,
    update_user_id   uuid,
    update_timestamp timestamp with time zone not null,
    constraint org_pk
        primary key (org_id),
    constraint org_create_user_fk
        foreign key (create_user_id) references demo.app_user
            deferrable initially deferred,
    constraint org_update_user_fk
        foreign key (update_user_id) references demo.app_user
            deferrable initially deferred,
    constraint org_create_app_fk
        foreign key (create_app_id) references demo.app
            deferrable initially deferred,
    constraint org_update_app_fk
        foreign key (update_app_id) references demo.app
            deferrable initially deferred
);

comment on column demo.org.org_id is 'Organization ID - Unique ID for table';

comment on column demo.org.org_extl_id is 'Organization Unique External ID to be given to outside callers.';

comment on column demo.org.org_name is 'Organization Name - a short name for the organization';

comment on column demo.org.org_description is 'Organization Description - several sentences to describe the organization';

comment on column demo.org.genesis_org is 'If true, the record represents the first organization created in the database and exists purely for the administrative purpose of creating other organizations, apps and users.';

comment on column demo.org.create_app_id is 'The application which created this record.';

comment on column demo.org.create_user_id is 'The user which created this record.';

comment on column demo.org.create_timestamp is 'The timestamp representing when this record was created.';

comment on column demo.org.update_app_id is 'The application which performed the most recent update to this record.';

comment on column demo.org.update_user_id is 'The user which performed the most recent update to this record.';

comment on column demo.org.update_timestamp is 'The timestamp representing when the record was updated most recently.';

alter table demo.org
    owner to demo_user;

alter table demo.app
    add constraint app_org_org_id_fk
        foreign key (org_id) references demo.org
            deferrable initially deferred;

alter table demo.app_user
    add constraint user_org_fk
        foreign key (org_id) references demo.org
            deferrable initially deferred;

create unique index org_org_id_uindex
    on demo.org (org_id);

create unique index org_org_name_uindex
    on demo.org (org_name);

create unique index org_org_extl_id_uindex
    on demo.org (org_extl_id);

