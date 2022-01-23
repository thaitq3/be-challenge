
## Implementation:
- Go Version: 1.16.12
- Database: MySQL
- Cache: Redis


## Requirements:
- Docker: version >= 3.0
- Go runtime to run unittest ( should be 1.16.12 or higher)

## Installation
- Start: sh start.sh --start
- Unit test: sh start.sh --unit-test 
- Help: sh start.sh

## API Interface
#### Place Wager

- Method: `POST`
- URL path: `/wagers`
- Request body:
    ```
    {
        "total_wager_value": <total_wager_value>,
        "odds": <odds>,
        "selling_percentage": <selling_percentage>,
        "selling_price": <selling_price>,
    }
    ```

- Response:
    Header: `HTTP 201`
    Body:
    ```
    {
        "id": <wager_id>,
        "total_wager_value": <total_wager_value>,
        "odds": <odds>,
        "selling_percentage": <selling_percentage>,
        "selling_price": <selling_price>,
        "current_selling_price": <current_selling_price>,
        "percentage_sold": <percentage_sold>,
        "amount_sold": <amount_sold>,
        "placed_at": <placed_at>
    }
    ```
    or

    Header: `HTTP <HTTP_CODE>`
    Body:
    ```
    {
        "error": "ERROR_DESCRIPTION"
    }
    ```

- Requirements:

    - `total_wager_value` must be specified as a positive integer above 0
    - `odds` must be specified as a positive integer above 0
    - `selling_percentage` must be specified as an integer between 1 and 100
    - `selling_price` must be specified as a positive decimal value to two decimal places, it is a monetary value
    - `selling_price` must be greater than `total_wager_value` * (`selling_percentage` / 100)
    - `id` should be an auto increment field
    - `current_selling_price` should be the `selling_price` until a `Buy Wager` action is taken against this wager record
    - `percentage_sold` should be null until a `Buy Wager` action is taken against this wager record
    - `amount_sold` should be null until a `Buy Wager` action is taken against this wager record
    - `placed_at` should be a timestamp at the completion of the request


#### Buy wager

- Method: `POST`
- URL path: `/buy/:wager_id`
- Request body:
    ```
    {
        "buying_price": <buying_price>
    }
    ```

- Response:
    Header: `HTTP 201`
    Body:
    ```
    {
        "id": <purchase_id>,
        "wager_id": <wager_id>,
        "buying_price": <buying_price>,
        "bought_at": <bought_at>
    }
    ```
    or

    Header: `HTTP <HTTP_CODE>`
    Body:
    ```
    {
        "error": "ERROR_DESCRIPTION"
    }
    ```

- Requirements:
    - `buying_price` should be an positive decimal value
    - `buying_price` must be lesser or equal to `current_selling_price` of the `wager_id`
    - A successful purchase should update the wager fields `current_selling_price`, `percentage_sold`, `amount_sold`
    - `id` should be an auto increment field
    - `bought_at` should be a timestamp at completion of the request


#### Wager list

- Method: `GET`
- URL path: `/wagers?page=:page&limit=:limit`
- Response:
    Header: `HTTP 200`
    Body:
    ```
    [
        {
            "id": <wager_id>,
            "total_wager_value": <total_wager_value>,
            "odds": <odds>,
            "selling_percentage": <selling_percentage>,
            "selling_price": <selling_price>,
            "current_selling_price": <current_selling_price>,
            "percentage_sold": <percentage_sold>,
            "amount_sold": <amount_sold>,
            "placed_at": <placed_at>
        }
        ...
    ]
    ```