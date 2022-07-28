#!/bin/bash

docker build -t shing1211/go-binance-scraper .
docker tag shing1211/go-binance-scraper shing1211/go-binance-scraper:0.0.1
docker push shing1211/go-binance-scraper
docker push shing1211/go-binance-scraper:0.0.1
