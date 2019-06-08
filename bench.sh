#!/usr/bin/env bash

set -x #echo on

(cd bad; go test -bench=. > ../bad.bench)
(cd good; go test -bench=. > ../good.bench)
(cd v3; go test -bench=. > ../v3.bench)
(cd v4; go test -bench=. > ../v4.bench)
(cd v5; go test -bench=. > ../v5.bench)
(cd v6; go test -bench=. > ../v6.bench)
(cd v7; go test -bench=. > ../v7.bench)
(cd v8; go test -bench=. > ../v8.bench)


benchcmp bad.bench good.bench
benchcmp good.bench v3.bench
benchcmp bad.bench v3.bench
benchcmp v3.bench v4.bench
benchcmp bad.bench v4.bench
benchcmp v4.bench v5.bench
benchcmp bad.bench v5.bench
benchcmp v5.bench v6.bench
benchcmp bad.bench v6.bench
benchcmp v6.bench v7.bench
benchcmp v7.bench v8.bench