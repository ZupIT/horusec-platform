# **BUILD**

## **Table of contents** 
### 1. [**About**](#about)
### 2. [**Environment**](#environment)
### 3. [**Architecture**](#architecture)
### 4. [**Development**](#development)
>#### 4.1. [**Golang**](#golang)
>#### 4.2. [**Javascript**](#javascript)
>#### 4.3. [**Style Guide**](#style-guide)
>#### 4.4. [**Tests**](#tests)
>##### 4.4.1. [**E2E**](#e2e)
>##### 4.4.2. [**Unitary Tests**](#unitary-tests)
>#### 4.5. [**Security**](#security)
### 5. [**Development with Docker**](#development-with-docker)       
>#### 5.1. [**Live reload**](#live-reload)
>#### 5.2. [**Manager**](#manager)
### 6. [**Production**](#production)  
>#### 6.1. [**Helm**](#helm)
>#### 6.2. [**Operator**](#operator)
>#### 6.3. [**Docker compose**](#docker-compose)

## **About**

The **BUILD.md** is a file to check the environment and build specifications of **horusec-platform** project.

## **Environment**

- [**Golang:**](https://go.dev/dl/) ^1.17.X
- [**PostgresSQL:**](https://www.postgresql.org/) ^12.X
- [**RabbitMQ:**](https://www.rabbitmq.com/) ^3.9.X
- [**NodeJS:**](https://nodejs.org/en/) ^16.X
- [**Yarn:**](https://yarnpkg.com/) ^1.20.X
- [**GNU Make:**](https://www.gnu.org/software/make/) ^4.2.X

## **Architecture**

The project consists of 9 microservices, 1 database (PostgreSQL by default) and a queue manager (RabbitMQ).

The microservices are:

| Service                         | Language  | Description                                                                                                                                |
| ------------------------------- | ---------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| [Analytic](./analytic)          | Golang     | Receives the analysis via broker, and saves the necessary data in its database that will be presented in the dashboard.                     |
| [Api](./api)                    | Golang     | Responsible for receiving requests [**Horusec-CLI**](https://github.com/ZupIT/horusec) via HTTP request to start a new analysis. |
| [Auth](./auth)                  | Golang     | Responsible for managing users, authentication and platform access.                                                                |
| [Core](./core)                  | Golang     | Responsible for managing workspaces, repositories and updating accesses.                                                            |
| [Messages](./messages)          | Golang     | Responsible for sending transactional emails.                                                                                         |
| [Migrations](./migrations)      | Golang     | Responsible for performing the migration in the Horusec database.                                                                        |
| [Vulnerability](./vulnerabiliy) | Golang     | Responsible for managing the vulnerabilities found in the analyses.                                                         |
| [Webhook](./webhook)            | Golang     | Responsible for configuring HTTP destinations and triggering analytics performed for third-party services.                         |
| [Manager](./manager)            | Javascript | The project's web interface.                                                                                                       |

You can learn more about the architecture in our [**documentation**](https://docs.horusec.io/docs/web/overview/).

## **Development**

With microservices architecture, we can handle the development of each unit in particular.

### **Golang**

For development in Golang microservices (analytic, api, auth, core, messages, vulnerability and webhook) the following steps must be followed:

**1**. Access the directory corresponding to the microservice you want to work with:

_**Example**_:

```bash
cd analytic/
```

**2**. Download the dependencies using the command:

```go
go mod download
```

**3**. Run the microservice using the command:

```go
go run ./cmd/app/main.go
```

### **Javascript**

The graphical interface (manager) is developed with [**ReactJS**](https://pt-br.reactjs.org/), for development you must follow the steps below:

**1**. Access the directory corresponding to the interface:

```bash
cd manager/
```

**2**. Download the dependencies using the command:

```bash
yarn
```

**3**. Run the interface using the command:

```bash
yarn start
```

### **Style Guide**

For source code standardization, this project use [golangci-lint](https://golangci-lint.run) tool as a linter aggregator of Go.

You can perform the lint check using the `make` command available in each microservice:

```bash
make lint
```

The project has a pattern of dependency imports, the commands below organize your code in the pattern defined by the Horusec team, these commands must be run in each microservice:

```bash
make fmt
```

Then, run the command:

```bash
make fix-imports
```

All project files must have the [**license header**](./copyright.txt). You can check if all files are in agreement by running this command in project root:

```bash
make license
```

If it is necessary to add the license in any file, the command below inserts it in all files that do not have the license:

```bash
make license-fix
```

### **Tests**

Each microservice has its unit tests, and the application as a whole has E2E tests.

#### **E2E**

The e2e tests are written with the [**cypress**](https://www.cypress.io/) tool.

To run the tests, follow the steps:

**1**. Access the directory to run the test:

```bash
cd e2e/cypress/
```

**2**. Then, run the command according to your scenario:

```bash
make test-e2e-auth-horusec-without-application-admin
``` 

Or

```bash
make test-e2e-auth-keycloak-without-application-admin
```

#### **Unitary Tests**

The Golang microservices unit tests were written with the [**standard package**](https://pkg.go.dev/testing) and some mock and assert excerpts, we used the [**testify**](https://github.com/stretchr/testify). You can run the tests using the command below:

```bash
make test
```

To check test coverage, run the command below:

```bash
make coverage
```

### **Security**

We use the latest version of [Horusec-CLI](https://github.com/ZupIT/horusec) to maintain the security of our source code. Through the command below, you can perform this verification in the project:

```bash
make security
```

## **Development with Docker**

To facilitate development, the project has the option of development through `Docker` images, which simulates a complete `Horusec-Platform` environment, using all microservices.

This development mode requires previously installed:

- [**Docker:**](https://www.docker.com/) ^20.0.X
- [**Docker Compose:**](https://docs.docker.com/compose/) ^1.20.X

In the [deployments/compose](./deployments) directory, you find the `docker-compose` files for building the environment.

### **Live reload**

With all services running through `Docker` and `docker-compose`, it is possible to make changes to the source code and these changes will be reflected in the running container.

To start development mode with docker, just run the following command at the root of the project:

```bash
make compose-dev
```

This way, all services will be available for use.

### **Manager**

The [**manager**](./manager) microservice is not available for live-reload via docker image.

If you need to make changes to your source code, it is recommended to use the [**traditional method**](###reactjs).

## **Production**

For production environments, we provide the following methods:

- Kubernetes
   - Helm
   - Operator

- Docker
   - docker-compose (not recommended)

### **Helm**

[**Helm**](https://helm.sh/docs/) is a package manager that gathers in a single file, called chart, all the defined resources of Kubernetes that make up an application.

This installation is for you to use the Horusec web application linked to your Kubernetes cluster with Helm.

See how to install Horusec-Platform via Helm in our [**documentation**](https://docs.horusec.io/docs/web/installation/install-with-helm/).

### **Operator**

[**Horusec-operator**](https://github.com/ZupIT/horusec-operator) performs management between the Horusec web services and the Kubernetes cluster. The creation idea came from the community with the desire to simplify the way to install the services in a Kubernetes environment.

See how to install Horusec-Platform via Operator in our [**documentation**](https://docs.horusec.io/docs/web/installation/install-with-operator/overview/).

### **Docker compose**

[**Docker-Compose**](https://docs.docker.com/compose/) is a tool that configures your application services as well as defines and runs Docker applications in various containers. You create and start all the services in your configuration with a single command.

See how to install Horusec-Platform via Operator in our [**documentation**](https://docs.horusec.io/docs/web/installation/install-with-docker-compose/).
