# Debezium

```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    price NUMERIC(10, 2)
);

INSERT INTO products (name, price) VALUES ('Product A', 10.99);
INSERT INTO products (name, price) VALUES ('Product B', 20.50);
-- update
UPDATE products SET price = 30.50 WHERE name = 'Product B';
-- delete
DELETE FROM products WHERE name = 'Product A';
-- insert
INSERT INTO products (name, price) VALUES ('Product C', 40.50);
```

## Example update
```json
{
    "before": null,
    "after": {
        "id": 2,
        "name": "Product B",
        "price": "C+o="
    },
    "source": {
        "version": "2.6.2.Final",
        "connector": "postgresql",
        "name": "cdc_",
        "ts_ms": 1719649856386,
        "snapshot": "false",
        "db": "debezium",
        "sequence": "[\"25685616\",\"25685672\"]",
        "ts_us": 1719649856386323,
        "ts_ns": 1719649856386323000,
        "schema": "public",
        "table": "products",
        "txId": 740,
        "lsn": 25685672,
        "xmin": null
    },
    "op": "u",
    "ts_ms": 1719649856839,
    "ts_us": 1719649856839757,
    "ts_ns": 1719649856839757000,
    "transaction": null
}
```

## Example delete
```json
{
    "before": {
        "id": 1,
        "name": null,
        "price": null
    },
    "after": null,
    "source": {
        "version": "2.6.2.Final",
        "connector": "postgresql",
        "name": "cdc_",
        "ts_ms": 1719649976130,
        "snapshot": "false",
        "db": "debezium",
        "sequence": "[\"25685808\",\"25685864\"]",
        "ts_us": 1719649976130193,
        "ts_ns": 1719649976130193000,
        "schema": "public",
        "table": "products",
        "txId": 741,
        "lsn": 25685864,
        "xmin": null
    },
    "op": "d",
    "ts_ms": 1719649976332,
    "ts_us": 1719649976332573,
    "ts_ns": 1719649976332573000,
    "transaction": null
}
```

## Example insert
```json
{
    "before": null,
    "after": {
        "id": 3,
        "name": "Product C",
        "price": "D9I="
    },
    "source": {
        "version": "2.6.2.Final",
        "connector": "postgresql",
        "name": "cdc_",
        "ts_ms": 1719650034227,
        "snapshot": "false",
        "db": "debezium",
        "sequence": "[\"25685976\",\"25686368\"]",
        "ts_us": 1719650034227535,
        "ts_ns": 1719650034227535000,
        "schema": "public",
        "table": "products",
        "txId": 742,
        "lsn": 25686368,
        "xmin": null
    },
    "op": "c",
    "ts_ms": 1719650034237,
    "ts_us": 1719650034237394,
    "ts_ns": 1719650034237394000,
    "transaction": null
}
```