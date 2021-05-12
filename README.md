<p align="center" margin="20 0"><a href="https://horusec.io/">
    <img src="https://github.com/ZupIT/horusec-devkit/blob/main/assets/horusec_logo.png?raw=true" 
            alt="logo_header" width="65%" style="max-width:100%;"/></a></p>

<p align="center">
    <a href="https://github.com/ZupIT/horusec-platform/pulse" alt="activity">
        <img src="https://img.shields.io/github/commit-activity/m/ZupIT/horusec-platform"/></a>
    <a href="https://github.com/ZupIT/horusec-platform/graphs/contributors" alt="contributors">
        <img src="https://img.shields.io/github/contributors/ZupIT/horusec-platform"/></a>
    <a href="https://github.com/ZupIT/horusec-platform/actions/workflows/analytic-pipeline.yml" alt="analytic">
        <img src="https://img.shields.io/github/workflow/status/ZupIT/horusec-platform/Analytic?label=analytic"/></a>
    <a href="https://github.com/ZupIT/horusec-platform/actions/workflows/api-pipeline.yml" alt="api">
        <img src="https://img.shields.io/github/workflow/status/ZupIT/horusec-platform/Api?label=api"/></a>
    <a href="https://github.com/ZupIT/horusec-platform/actions/workflows/core-pipeline.yml" alt="core">
        <img src="https://img.shields.io/github/workflow/status/ZupIT/horusec-platform/Core?label=core"/></a>
    <a href="https://github.com/ZupIT/horusec-platform/actions/workflows/manager-pipeline.yml" alt="manager">
        <img src="https://img.shields.io/github/workflow/status/ZupIT/horusec-platform/Manager?label=manager"/></a>
    <a href="https://github.com/ZupIT/horusec-platform/actions/workflows/messages-pipeline.yml" alt="messages">
        <img src="https://img.shields.io/github/workflow/status/ZupIT/horusec-platform/Messages?label=messages"/></a>
    <a href="https://github.com/ZupIT/horusec-platform/actions/workflows/migrations-pipeline.yml" alt="migrations">
        <img src="https://img.shields.io/github/workflow/status/ZupIT/horusec-platform/Migrations?label=migrations"/></a>
    <a href="https://github.com/ZupIT/horusec-platform/actions/workflows/vulnerability-pipeline.yml" alt="vulnerability">
        <img src="https://img.shields.io/github/workflow/status/ZupIT/horusec-platform/Vulnerability?label=vulnerability"/></a>
    <a href="https://github.com/ZupIT/horusec-platform/actions/workflows/webhook-pipeline.yml" alt="webhook">
        <img src="https://img.shields.io/github/workflow/status/ZupIT/horusec-platform/Webhook?label=webhook"/></a>
    <a href="https://github.com/ZupIT/horusec-platform/actions/workflows/auth-pipeline.yml" alt="auth">
        <img src="https://img.shields.io/github/workflow/status/ZupIT/horusec-platform/Auth?label=auth"/></a>
    <a href="https://opensource.org/licenses/Apache-2.0" alt="license">
        <img src="https://img.shields.io/badge/license-Apache%202-blue"/></a>
</p>

# Horusec Platform

Horusec Platform is a set of web services that integrate with the [Horusec CLI](https://github.com/ZupIT/horusec) 
to facilitate the visualization and management of vulnerabilities.

{@TODO EXAMPLE IMAGE DASHBOARD}

## Dependencies

- [RabbitMQ](https://www.rabbitmq.com/)
- [Postgresql](https://www.postgresql.org/)

## Installation

### Quick Run:

If you just want to try the web interface, we made an image that will automatically configure a ready-to-use environment.
This image is not recommended for production environments, and will not persist any data after being interrupted.

```cmd
make run-web
```

After executing, the Horusec [image](https://hub.docker.com/r/horuszup/horusec-all-in-one) will start to install 
all dependencies and services, after finished it will present the following message
`HORUSEC WEB IS UP AND CAN BE ACCESSED IN -> http://localhost:8043/auth`.

The installation will be done with all default values and latest versions
and also create the following test user:

```
Username: dev@example.com
Password: Devpass0*
```

To stop the running container just execute:

```cmd
make stop-web
```

Click [here](@TODO QUICK RUN DOCS) to check full quick run docs.

### Install with docker compose:

```cmd
make install
```

We will execute the file `compose.yml` which contains all services, migrations and the needed dependencies. 
The compose file and can be found in `deployments/compose/compose.yaml` and migrations in `migrations/source`.

The installation will be done with all default values and latest versions
and also create the following test user:

```
Username: dev@example.com
Password: Devpass0*
```

Click [here](https://horusec.io/docs/web/installation/install-with-docker-compose) 
to check full docker compose installation docs.

### Install with helm:

@TODO

The helm files for each service can be found at `deployments/helm`.

Click [here](https://horusec.io/docs/web/installation/install-with-helm) to check the helm installation docs.

### Install with horusec admin:

@TODO

### Install with horusec operator:

@TODO

## Features

@TODO

### Authentication Types
@TODO
- HORUSEC
- LDAP
- KEYCLOAK

### Tokens
@TODO

### Dashboard
@TODO

### Vulnerability Management
@TODO

## Migrating From V1
@TODO

## Contributing

Feel free to use, recommend improvements, or contribute to new implementations.

If this is our first repository that you visit, or would like to know more about Horusec,
check out some of our other projects.

- [Horusec CLI](https://github.com/ZupIT/horusec)
- [Horusec DevKit](https://github.com/ZupIT/horusec-devkit)
- [Horusec Engine](https://github.com/ZupIT/horusec-engine)
- [Horusec Operator](https://github.com/ZupIT/horusec-operator)
- [Horusec Admin](https://github.com/ZupIT/horusec-admin)
- [Horusec VsCode](https://github.com/ZupIT/horusec-vscode-plugin)

This project exists thanks to all the contributors. You rock! ❤️🚀
