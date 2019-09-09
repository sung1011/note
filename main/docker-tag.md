# docker tag

## < null >

The defacto image. Use it if unsure.

## stretch & jessie

jessie or stretch are the suite code names for releases of Debian and indicate which release the image is based on.

## alpine

Similarly, this image is based on the Alpine Linux, thus being a very small base image. It is recommended if you need an image size is as small as possible. The caveat is that it uses some unusual libs, but shouldn't be a problem for most software. In doubt, check the official docs below.

## slim

This image only contains the minimal packages needed to run Java (and is missing many of the UI-related Java libraries, for instance). Unless you are working in an environment where only the openjdk image will be deployed and you have space constraints, the default image is recommended over this one.

## windowsservercore

This image is based on Windows Server Core (microsoft/windowsservercore).

## bionic

## xenial
