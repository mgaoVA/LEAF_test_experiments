# Docker Environment

## Prerequisites
- Install [LEAF docker development enviornment](https://github.com/department-of-veterans-affairs/LEAF/blob/master/docs/InstallationConfiguration.md)
- Map test DB into development environment (this needs to be automated)
  1. Import test databases from ../*.sql
  2. In the national_leaf_launchpad database, update fields for portal_database and orgchart_database to point to the imported test databases
  3. After running the test, restore the original fields in #2

## Run
```
docker compose up --remove-orphans
```


# Native Environment

## Prerequisites
- Install Deno

## Run tests
```
deno test --allow-net --unsafely-ignore-certificate-errors
```

## Run benchmarks
```
deno bench --allow-net --unsafely-ignore-certificate-errors
```
