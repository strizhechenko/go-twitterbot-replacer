#!/usr/bin/env python
import requests
for line in requests.get('http://127.0.0.1:8080/tweets').json():
    print line.encode('utf-8')
    print
