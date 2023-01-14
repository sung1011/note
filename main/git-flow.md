# git-flow

![img](res/gitflow.svg)

## main/master

1. The `main` branch stores the official release history

## develop

1. A `develop` branch is created from `main`

## feature

1. `Feature` branches are created from `develop`
1. When a `feature` is complete it is merged into the `develop` branch

## release

1. A `release` branch is created from `develop`
1. When the `release` branch is done it is merged into `develop` and main

## hotfix

1. If an issue in `main` is detected a `hotfix` branch is created from `main`
1. Once the `hotfix` is complete it is merged to both `develop` and `main`