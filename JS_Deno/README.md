# LEAF API Testing

This test suite should exercise access rules, data accuracy, and measure performance within hotpaths for the LEAF API. ***Currently only compatible with Sprint-60-c2 and newer, due to recent containerization changes.***

This leverages an existing [LEAF docker development environment](https://github.com/department-of-veterans-affairs/LEAF/blob/master/docs/InstallationConfiguration.md), and temporarily uses a predefined test database. 

TODO:
1. Setup and configuration of the Test DB needs to be automated
1. Need a more streamlined way to visualize results from benchmarks between different releases

## Prerequisites
- Install [LEAF docker development environment](https://github.com/department-of-veterans-affairs/LEAF/blob/master/docs/InstallationConfiguration.md)
- The LEAF docker development environment must be running
- Map test DB into development environment (this needs to be automated)
  1. Import test databases from ../Go/database/*.sql
  2. In the national_leaf_launchpad database, update fields for portal_database and orgchart_database to point to the imported test databases
  3. After running the test, restore the original fields in #2

# Docker Environment

## Run
Navigate to the docker directory, then run:
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
