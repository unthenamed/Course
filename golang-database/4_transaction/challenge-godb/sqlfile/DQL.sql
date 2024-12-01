SELECT order_id FROM "order" WHERE order_id = $1;

SELECT customer_id FROM customer WHERE customer_id = $1;

SELECT service_id FROM service WHERE service_id = $1;

SELECT COUNT(order_detail_id) FROM order_detail;

INSERT INTO "order" (order_id, customer_id,order_date, completion_date,received_by, created_at, update_at) 
VALUES ($1,$2,$3,$4,$5,$6,$7);


INSERT INTO order_detail (order_detail_id,order_id,service_id,qty) 
VALUES ($1,$2,$3,$4);

DELETE FROM "order"  WHERE order_id = $1;

DELETE FROM order_detail  WHERE order_id = $1;

INSERT INTO service (service_id, service_name,unit, price, created_at, update_at)
VALUES ($1, $2, $3, $4, $5, $6);


SELECT * FROM service;

SELECT * FROM service WHERE service_id = $1

UPDATE service
    SET service_name = $2, unit = $3, price = $4, update_at = $5 
WHERE service_id = $1;
                        
DELETE FROM service WHERE service_id = $1

SELECT order_id FROM "order" WHERE completion_date = $1

SELECT a.order_id, a.customer_id, b.order_detail_id, a.order_date, a.completion_date, b.service_id, b.qty
    FROM "order" AS a
    JOIN order_detail AS b
    ON b.order_id = b.order_id 
ORDER BY a.order_id ASC;

SELECT a.order_id, a.customer_id, b.order_detail_id, a.order_date, a.completion_date, b.service_id, b.qty
    FROM "order" AS a
    JOIN order_detail AS b
    ON b.order_id = b.order_id 
WHERE a.order_id = $1;

INSERT INTO customer (customer_id, name,phone, address, created_at, update_at)
VALUES ($1, $2, $3, $4, $5, $6);
					
SELECT * FROM customer

SELECT * FROM customer WHERE customer_id = $1

UPDATE customer
    SET name = $2, phone = $3, address = $4, update_at = $5 
WHERE customer_id = $1;
                        
DELETE FROM customer WHERE customer_id = $1