#!/usr/bin/env bash
git add .
git commit -m 'support os.env'
git push origin master
git tag v1.0.7
git push origin v1.0.7