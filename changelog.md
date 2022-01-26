# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2022-01-26

### Added

- Fully Support /boot/ Api
  - support for Rescue, Linux, Vnc, Windows, Plesk and cPanel boot.
- Fully Support /order/server/ Api
- Fully Support /order/server_market/ Api
- Fully Support /ip/ Api
- Fully Support /rdns/ Api
- Fully Support /traffic/ Api

### Changed

- Api must use {server-number} now, {server-ip} is deprecated
- Client.Call() remove needAuth argument, since every endpoint need auth
- Wol support IPv6 address

### Removed

- vServer is deprecated

## [0.3.1] - 2018-07-29

### Removed

- remove all mentions of github.com/cenkalti/backoff library

## [0.3.0] - 2018-07-29

### Changed

- custom data types ArrayOrString, ArrayOrInt replaced to their standard equivalents

## [0.2.0] - 2018-07-29

### Added

- Parial support of failover API
  - it's possible to get list of failover ips
  - it's possible to switch failover ip to another server

### Changed

- HTTP operation timeout changed from 10 to 120 seconds
- User agent identifier

### Removed

- Automatic retry of failed HTTP requests

## [0.1.0] - 2018-04-11

Project forked from [https://github.com/appscode/go-hetzner](https://github.com/appscode/go-hetzner)
