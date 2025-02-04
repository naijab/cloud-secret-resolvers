[![codecov](https://codecov.io/gh/kubeopsskills/cloud-secret-resolvers/branch/main/graph/badge.svg?token=t65R7COoaz)](https://codecov.io/gh/kubeopsskills/cloud-secret-resolvers)
# Cloud Secret Resolvers (CSR)

Cloud Secret Resolvers is a set of tools to help your applications (on Kubernetes) to retrieve any credentials from cloud managed vaults without the needed to write additional boilerplate code in your applications!.

<!-- TOC -->

- [CSR](#csr)
    - [Installation](#installation)
    - [Using on Kubernetes](#using-on-kubernetes)
    - [How it works](#how-it-works)

<!-- /TOC -->

## Installation

Cloud Secret Resolvers is available on Linux, ARM, macOS and Windows platforms.
- Binaries for Linux, ARM, Windows and Mac are available as tarballs in the [release](https://github.com/kubeopsskills/cloud-secret-resolvers/releases) page

## Using on Kubernetes

- AWS
  
  - Prerequisites:
    - Enabled the OIDC provider on your [EKS](https://aws.amazon.com/th/eks/) cluster (https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html)
    - Your application Kubernetes pod has a service account with the following privillege:
       [policy.json](assets/policy.json)
  - Update your application entrypoint as follows:
    ```bash
    #!/bin/bash
    eval $(csr)
    node ... # your application runtime command
    ```
  - Update your application Kubernetes config maps as follows:
    ```yaml
    apiVersion: v1
    kind: ConfigMap
    metadata:
      name: [your config map name]
      namespace: [your config map namespace name]
    data:
    ...
    CLOUD_TYPE: "aws"
    AWS_REGION: "[your AWS region name]"
    AWS_SECRET_NAME: "[your AWS secret name]"
    ```

- Azure
  - Coming Soon!

- Google Cloud
  - Coming Soon!

## How it works
The architecture looks like below.

Internally, the `CSR` find local environment variables in the Kubernetes Pod Container which have Cloud Vault key placeholders for example: export db_username=${db_username}, then the `CSR` will extract db_username as a key and ${db_username} as a value. Finally, the `CSR` will use ${db_username} to match cloud vault key, retrieve cloud vault value, and map the value with db_username local environment.

![Diagram](https://github.com/kubeopsskills/cloud-secret-resolvers/blob/main/assets/diagram.png)