#!/usr/bin/env bash
git add .
git commit -m 'release environment-config'
git push origin master
git tag v1.0.2
git push origin v1.0.2