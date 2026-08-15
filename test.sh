#!/bin/bash

feed_links="./feed_list.txt"
go run . reset
go run . register jenix
go run . register rejet
go run . login jenix

while IFS=" " read -r name url; do
  go run . addfeed "$name" "$url"
done <"$feed_links"

go run . following
go run . agg 10s
