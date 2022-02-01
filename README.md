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

# **Horusec Platform**

## **Table of contents**
### 1. [**About**](#about)
### 2. [**Usage**](#usage)
>#### 2.1. [**Requirements**](#requirements)
>#### 2.2. [**Installation**](#installation)
### 3. [**Features**](#features)
### 4. [**Documentation**](#documentation)
### 5. [**Issues**](#issues)
### 6. [**Contributing**](#contributing)
### 7. [**License**](#license)
### 8. [**Community**](#community)

## **About**
Horusec Platform is a set of web services that integrate with [**Horusec-CLI**](https://github.com/ZupIT/horusec) to make it easier for you to see and manage the vulnerabilities. 

[comment]: <> (@todo add a gif of manager usage)


## **Usage**

### **Requirements**
See below the requirements to install Horusec-Platform: 

- [**RabbitMQ**](https://www.rabbitmq.com/)
- [**PostgreSQL**](https://www.postgresql.org/)

### **Installation**
There are several ways to install the Horusec-Platform in your environment.
In some of them, we use a **`make`** command to simplify the process.
If you want to know everything that will be executed, take a look at the **`Makefile`** located at the project's root.

Choose what type of installation you want below, but remember to change the default environment variables values to new and secure ones.

### **1. Install with docker compose**
Follow the steps: 

**Step 1:** Run the command: 
```cmd
make install
```

**Step 2:** Start the docker compose file **`compose.yml`**. It has all services, migrations and the needed dependencies. 
- You can find the compose file in **`deployments/compose/compose.yaml`**; 
- You can find migrations in **`migrations/source`**.

**Step 3:** Now the installation is ready with all default values, the latest versions, and the user for tests, see below:

```
Username: dev@example.com
Password: Devpass0*
```

Docker compose file is configured to perform a standard installation by default.  
In the production environments' case, make sure to **change the values of the environment variables to new and secure ones**.

> :warning: We **do not recommend** using docker-compose installation in a productive environment.

For more information about Docker compose, check out [**Docker compose installation section**](https://horusec.io/docs/web/installation/install-with-docker-compose).

### **2. Install with Helm**

Each release contains its own helm files for that specific version, you can find them [**in the repository**](https://github.com/ZupIT/horusec-platform/releases) and in the folder **`deployments/helm`**.
In both cases, they will be separated by each service of the architecture.

For more information, check out [**the installing with Helm section**](https://horusec.io/docs/web/installation/install-with-helm).

### **3. Install with Horusec-Operator**

Horusec-Operator manages Horusec web services and its Kubernetes cluster. It was created based on the community’s idea to have a simpler way to install the services in an environment using Kubernetes. 

-  Check out how to install Horusec-Operator in our [**installation section**](https://horusec.io/docs/web/installation/install-with-operator/).
- For more information about Kubernetes Operators, [**check out the documentation**](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).


## **Features**

Horusec-Platform provides several features, see some of them below. 

### **MultiTenancy**

It distributes only the necessary [**permissions**](https://horusec.io/docs/web/overview/#1-multitenant) according to each user: 

<p align="center" margin="20 0"><img src="assets/horusec-invite-users-1.png" alt="multiTenancy" width="100%" style="max-width:100%;"/></p>

### **Dashboard**

The dashboard shows you several metrics about your workspaces and repositories' vulnerabilities:

<p align="center" margin="20 0"><img src="assets/horusec-dashboard-1.png" alt="dashboard" width="100%" style="max-width:100%;"/></p>

### **Vulnerability Management**

The vulnerability management screen allows you to identify false positives and accepted risks. You can modify a severity to an appropriate value to the reality of the vulnerability:

<p align="center" margin="20 0"><img src="assets/horusec-vuln-management-1.png" alt="vuln-management" width="100%" style="max-width:100%;"/></p>

### **Tokens**
It creates workspaces or repositories authentication 
[**tokens**](https://horusec.io/docs/tutorials/how-to-create-an-authorization-token) for your pipeline: 

<p align="center" margin="20 0"><img src="assets/horusec-create-token-1.png" alt="tokens" width="100%" style="max-width:100%;"/></p>

### **Authentication Types**

You can choose which form of authentication you will use with Horusec-Platform.

There are three possibilities:

- HORUSEC (native) 
- LDAP
- KEYCLOAK

For more information about authentication types, check out our [**documentation**](https://horusec.io/docs/tutorials/how-to-change-authentication-types).

[comment]: <> ([comment]: <> &#40;## Migrating From V1&#41;)

[comment]: <> (For more information on migrating from the previous version to the current one see our )

[comment]: <> ([documentation]&#40;@todo&#41;.)

## **Documentation**

For more information about Horusec, please check out the [**documentation**](https://horusec.io/docs/).

## **Issues**

To open or track an issue for this project, in order to better coordinate your discussions, we recommend that you use the [**Issues tab**](https://github.com/ZupIT/horusec/issues) in the main [**Horusec**](https://github.com/ZupIT/horusec) repository.

## **Contributing**

If you want to contribute to this repository, access our [**Contributing Guide**](https://github.com/ZupIT/horusec-platform/blob/main/CONTRIBUTING.md). 

### **Developer Certificate of Origin - DCO**

 This is a security layer for the project and for the developers. It is mandatory.
 
 Follow one of these two methods to add DCO to your commits:
 
**1. Command line**
 Follow the steps: 
 **Step 1:** Configure your local git environment adding the same name and e-mail configured at your GitHub account. It helps to sign commits manually during reviews and suggestions.

 ```
git config --global user.name “Name”
git config --global user.email “email@domain.com.br”
```
**Step 2:** Add the Signed-off-by line with the `'-s'` flag in the git commit command:

```
$ git commit -s -m "This is my commit message"
```

**2. GitHub website**
You can also manually sign your commits during GitHub reviews and suggestions, follow the steps below: 

**Step 1:** When the commit changes box opens, manually type or paste your signature in the comment box, see the example:

```
Signed-off-by: Name < e-mail address >
```

For this method, your name and e-mail must be the same registered on your GitHub account.

## **License**
[**Apache License 2.0**](https://github.com/ZupIT/horusec-platform/blob/main/LICENSE).

## **Community**
Do you have any question about Horusec? Let's chat in our [**forum**](https://forum.zup.com.br/).


This project exists thanks to all the contributors. You rock! ❤️🚀
