---
id: 05212668-8F09-4A0E-AC9C-06CBB39FE100
creation-date: 2023-02-04T12:14:29
type: walkthrough
aliases:
  - SSH server
tags:
  - devops
  - cli
---

## 0. Pre-requisite Installations

There are several CLI softwares required to create SSH connection: 
- `ssh-keygen`
- `net-tools`

---
## 1. Enable SSH
To enable SSH in a Linux machine, we need to: 
- allow the machine to be a *Host*
	- `ssh-keygen -A`
- enable Hosting of SSH to external network 
	- `sudo service ssh start`
	- akin to expose your address publicly
- check if the machine has a SSH hosting service running in the background
	- `sudo service --status-all | grep ssh`


> [!NOTE] 
> SSH automatically exposes `port:22` for ingress


---
## 2. Setup Public & Private Keys

- setup RSA [[0_pages/02/d-2023-08-26-12-56-05|Public & Private Keys]] 
	- `ssh-keygen -t rsa`
	- akin to getting a pair of lock and key
- authorise public key to
	- `cp <publicKey>.pub ~/.ssh/authorised_keys`


---
## ℹ️  Resources
- [Geeks for Geeks](https://www.geeksforgeeks.org/how-to-set-up-openssh-server-in-windows-subsystem-for-linux-wsl/)