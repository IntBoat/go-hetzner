# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2022-01-27T08:56:24+08:00

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

Project forked from [https://github.com/appscode/go-hetzner](https://github.com/appscode/go-hetzner)
