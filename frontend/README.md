# Overview

Filscan(https://filscan.io/) is the first data browser in the Filecoin ecosystem and is dedicated to building the Filecoin network data analysis center, providing the Filecoin ecosystem data services, visualization and other one-stop data analysis for the whole network, and building a data intelligence platform for 3rd party ecosystem applications. 

# Table of Contents
- [Overview](#overview)
- [Table of Contents](#table-of-contents)
  - Key Features of Filscan
  - Half year Plan
  - Get involved and Contribute
- [Front-End](#front-end)
  - [Install](#install)
    - [Environmental requirements](#environmental-requirements)
    - [Install Dependencies](#install-dependencies)
    - [Compiles and hot-reloads for development](#compiles-and-hot-reloads-for-development)
    - [API Environment Configuration](#api-environment-configuration)
    - [Build](#build)
    - [Lints and fixes files](#lints-and-fixes-files)
    - [Customize configuration](#customize-configuration)
- [Back-End](#back-end)
  - [Build and Install](#build-and-install)
    - [Environment](#environment)
    - [System Require](#system-require)
    - [Build](#build-1)
    - [Configuration](#configuration)
    - [Run](#run)
  - [API Document](#api-document)
  - [License](#license)


## Key Features of Filscan

•  Provide data applications based on various scenarios

![image](https://user-images.githubusercontent.com/86345705/129295935-15f398b2-fbe9-46fb-9d95-f6e8b5c7beb3.png)

• Provide real-time data display and API interface


![image](https://user-images.githubusercontent.com/86345705/129296005-298481b5-db14-4c79-b8c3-1e80ff7f1f61.png)

• Provide data visualization services to increase business management capabilities


![image](https://user-images.githubusercontent.com/86345705/129296094-c463333b-494c-47c7-a25a-9406eb8edcd8.png)

• Provide rich data analysis functions to drive business progress

![image](https://user-images.githubusercontent.com/86345705/129296192-7b9ae3b2-d576-44b8-b7f7-6a17b297ee42.png)

• Provide advanced services such as predictive data analysis and business evaluation

![image](https://user-images.githubusercontent.com/86345705/129296331-5e75e79e-2cc5-4533-a5df-a8b2605e414e.png)

• Monitoring panel to support multi-node data collection and monitoring

![image](https://user-images.githubusercontent.com/86345705/129296455-c5d02cab-27b4-4bf8-9803-67a99914cd25.png)

### Half year plan :(Next 6 months)

•  Upgrade the back-end architecture to improve the loading speed

•  Better UI display and more intuitive data display

•  Added order query about Filecoin real data 

•  Added the isolated block statistics and query function

•  Set private Key to control API interfaces

•  Add address collection list, convenient to view address information faster

### Get involved and Contribute

As a Filecoin ecosystem open source project, Filscan is keen to welcome your feedback and suggestions. 

Reach us@

Twitter: https://twitter.com/force_ipfs

Medium: https://ipfsforce-1751.medium.com/

Email: filscan@ipfsforce.com

# Front-End

## Install

### Environmental requirements

- required: Node.js(https://nodejs.org/en/)

- optional : Yarn (https://yarnpkg.com/)

### Install Dependencies
```
yarn install or npm install 
```
### Compiles and hot-reloads for development
```
yarn serve or npm run serve
```
### API Environment Configuration

By default, the profiles are in the root directory of the project. You can modify the value of **VUE_APP_BASE_URL** to change the server address.

**Example:**

If you want to change the server address, you can open the file .env.development. Then you will see the default value of **VUE_APP_BASE_URL** is "http://192.168.1.2:8700/v0/filscan", modify it according to your actual server address. Please notice that you should run "yarn serve" or "npm run serve" to apply this change.

### Build
```
yarn build:pro or npm run build:pro
```
### Lints and fixes files
```
yarn lint or npm run lint
```
### Customize configuration

See Configuration Reference(https://cli.vuejs.org/config/).

# [Back-End](https://github.com/ipfs-force-community/filscan-backend)

## Build and Install

### Environment

- golang >= v1.13
- mongo >= v4.2
- lotus >= v0.2.7

### System Require

- Linux or Mac OS

### Build
```
git clone (githuburl)

cd Backend

make build-lotus

go build
```
### Configuration

Edit app.conf in path /conf and set the correct parameter
```
mongoHost = "127.0.0.1:27017"

mongoUser = "root"

mongoPass = "admin"

mongoDB   = "filscan"

lotusGetWay="192.168.1.1:1234"
```
### Run

Make sure mongo and lotus is active, and run the filscan_lotus
```
./filscan_lotus
```
The application will check lotus and mongo’s status. The application will stop if got any error from them. If application start success, it will work until sync all data down from lotus. 

## API Document

Check document [here](https://github.com/ipfs-force-community/filscan-backend/blob/master/Filscan_Interface_v1.0.md)



## License
Dual-licensed under 

[MIT](https://github.com/filecoin-project/lotus/blob/master/LICENSE-MIT) 

[Apache 2.0](https://github.com/filecoin-project/lotus/blob/master/LICENSE-APACHE)
