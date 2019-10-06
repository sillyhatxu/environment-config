#!/usr/bin/env bash
git add .
git commit -m 'add remind db config'
git push origin master
git tag v1.0.9
git push origin v1.0.9