--製令單
CREATE TABLE manu_order
(
  manu_order_id    TEXT      NOT NULL,
  goods_name       TEXT      NULL    ,
  identifier       TEXT      NULL    ,
  total_quantity   TEXT      NULL    ,
  completion_date  TEXT      NULL    ,
  date_of_issuance TEXT      NULL    ,
  created_by       TEXT     NULL    ,
  created_at       TIMESTAMP NULL    ,
  updated_at       TIMESTAMP NULL    ,
  is_deleted       BOOLEAN   NULL    ,
  PRIMARY KEY (manu_order_id)
);