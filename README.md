# block_explorer

A backend server that synchronize metaos chain data into a database

# SetUp

## Run Mongodb replica set via docker-compose
`docker-compose up`

Mongo-express(a web-based MongoDB admin interface): http://localhost:8081

## Config
1. Set `node_urls`(blockchain node RPC address) in [config/config.toml](./config/config.toml)
2. `export CONFIG_FILE_PATH=config/config.toml`
## Build And Run

- Build: `make all`
- Run: `make run`
- Cross compilation: `make build-linux`

## Config Description

### [config.toml](https://github.com/FiiLabs/block_explorer/blob/master/config/config.toml)

### Db config

- **node_uri**: `required` `string` mongodb connection uri（example: `mongodb://127.0.0.1:27017/?connect=direct&authSource=metaos`）
- database：`required` `string` database name（example：`DB_DATABASE`）

### Server config

- **ip_port**: `required` `string` API server IP address and port (example `127.0.0.1:8787`).
- **node_urls**: `required` `string`  full node uri（example: `tcp://127.0.0.1:26657, tcp://127.0.0.2:26657, ...`）
- worker_num_create_task: `required` `int` the maximum time (in seconds) that create task are allowed （default: `1`
  example: `1`）
- worker_num_execute_task: `required` `int` the maximum time (in seconds) that synchronization TX threads are allowed
  to be out of work（example: `30`）
- worker_max_sleep_time: `required` `int` num of worker to create tasks(unit: seconds)（example: `90`）
- block_num_per_worker_handle: `required` `int`  number of blocks per sync TX task（example: `50`）
- max_connection_num: `required` `int` client pool config total max connection
- init_connection_num: `required` `int` client pool config idle connection
- bech32_acc_prefix: `option` `string` block chain address prefix（default(cosmoshub): `` example(metaos): `metaos`）
- chain_block_interval: `option` `int` block interval; default `5` (example: `5`)
- behind_block_num: `option` `int` wait block num to handle when retry failed; default `0` (example: `0`)
- promethous_port: `option` `int` promethous metrics server port
- support_modules: `option` `string` setting only support module tx sync,default
  support [all module] (default: ``
  example: `bank,nft`)
- deny_modules: `option` `string` disable support module tx sync
- support_types: `option` `string` setting only support msgType tx sync,default support all types(default: ``
  example: `transfer,recv_packet`)
- ignore_ibc_header: `option` `boolean` setting update_client header info for tx collection ,default not ignore ibc header info(default: `false`
    example: `false`)
- chain_id: `option` `string` setting collection name by chain_id

Note:
> synchronizes cosmos data from specify block height(such as:17908 current time:1576208532)
> At first,stop the block_explorer and create the task. Run:

  ```
  db.sync_task.insert({
      'start_height':NumberLong(17908),
      'end_height':NumberLong(0),
      'current_height':NumberLong(0),
      'status':'unhandled',
      'worker_id' : '',
       'worker_logs' : [],
      'last_update_time' : NumberLong(0)
  })
  ```
