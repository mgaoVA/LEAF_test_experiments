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
- Install Go

## Run tests
```
go test
```

## Run benchmarks
```
go test -run="^$" -bench=. -count=3
```

## Benchmark analysis

Prerequisite:
```
go install golang.org/x/perf/cmd/benchstat
```

Store results:
```
go test -run="^$" -bench=. -count=10 > sprint61.txt
```

Compare:
```
benchstat sprint60.txt sprint63.txt
```

Example output:
```
goos: windows
goarch: amd64
pkg: LEAF/API-tester
cpu: 11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz
                           │ sprint60.txt │            sprint63.txt             │
                           │    sec/op    │   sec/op     vs base                │
Homepage_defaultQuery-8       57.46m ± 3%   58.15m ± 3%        ~ (p=0.190 n=10)
Inbox_nonAdminActionable-8   315.92m ± 3%   80.22m ± 2%  -74.61% (p=0.000 n=10)
geomean                       134.7m        68.30m       -49.31%
```
