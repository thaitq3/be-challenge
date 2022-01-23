create table wager
(
    id                    int unsigned auto_increment,
    total_wager_value     int unsigned                        not null,
    odds                  int                                 not null,
    selling_percentage    double unsigned                     not null,
    selling_price         double unsigned                     not null,
    current_selling_price double unsigned                     not null,
    percentage_sold       int unsigned                        null,
    amount_sold           double unsigned                     null,
    placed_at             timestamp default current_timestamp not null,
    created_at          datetime                            null,
    updated_at          datetime                            null,
    deleted_at            datetime                            null,
    constraint wager_pk
        primary key (id)
);

create table purchase
(
    id           int unsigned auto_increment,
    wager_id     int unsigned                        not null,
    buying_price double unsigned                     not null,
    bought_at    timestamp default current_timestamp not null,
    created_at datetime                            null,
    updated_at datetime                            null,
    deleted_at   datetime                            null,
    constraint purchase_pk
        primary key (id),
    CONSTRAINT `fk_purchase_wager_wager_id`
        FOREIGN KEY (`wager_id`)
            REFERENCES `wager` (`id`)
            ON UPDATE NO ACTION
            ON DELETE NO ACTION
);

