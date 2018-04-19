#!/bin/bash


echo "Removing previous vendor "
echo "------------------------- "
for d in */ ; do
  if [[ $d = *"service"* ]]; then
    echo "---------------------------- "
    echo "Processing: $d"
    echo "---------------------------- "
    cd $d
    pwd 
    rm Gopkg.lock
    rm Gopkg.toml
    rm -rf vendor
    dep init && dep ensure
    cd ..
  fi
done

