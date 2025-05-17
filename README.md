# Oiler Backup Base Repository

<div align="center">
    <img src="./images/oiler-icon.gif" alt="Logo" width="30%" height="30%" />
</div>

<br>

<div align="center">
    <img src="https://img.shields.io/badge/License-Apache 2.0-blue" alt="License" />&nbsp;&nbsp;
    <img src="https://img.shields.io/badge/Language-Go-blue" alt="Language" />&nbsp;&nbsp;
</div>

<br>

This repository contains useful entities for simplifying the writing of adapters to the [oiler-backup Kubernetes operator](https://github.com/AntonShadrinNN/oiler-backup).

# How this repository is organized

Documentation for each component could be found on [pkg.go.dev](https://pkg.go.dev/github.com/AntonShadrinNN/oiler-backup-base)

|Package|Purpose|
|-------|-------|
| [logger](./logger) | Provides default project logger |
| [metrics](./metrics) | Methods to report metrics to core |
| [proto](./proto) | All proto-files and generated code |
| [s3](./s3) | Methods to work with s3-compatible storage |
| [servers/backup](./servers/backup) | Methods to work with Kubernetes |
| [servers/backup/envgetters](./servers/backup/envgetters) | Entities to simlify work with envs |