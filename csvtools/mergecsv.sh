#!/bin/bash

{ head -n1 *_0.csv; for f in file*.csv; do tail -n+2 "$f"; done; } > merged.csv
