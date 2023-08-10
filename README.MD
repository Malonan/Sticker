# <center><strong>I Don't Like This Sticker [END]</strong></center>
 
<br>

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FMalonan%2FSticker.svg?type=small)](https://app.fossa.com/projects/git%2Bgithub.com%2FMalonan%2FSticker?ref=badge_smail)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Malonan/Sticker?label=Go%20Version&style=flat-square)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/Malonan/Sticker?label=Release%20Version&style=flat-square)](https://github.com/Malonan/Sticker/release)
[![Apache-2.0](https://img.shields.io/github/license/Malonan/Sticker?style=flat-square)](https://github.com/Malonan/Sticker/blob/master/LICENSE)

[![GitHub issues](https://img.shields.io/github/issues/Malonan/Sticker?label=Sticker%20Issues&style=flat-square)](https://github.com/Malonan/Sticker/issues)
[![GitHub Repo stars](https://img.shields.io/github/stars/Malonan/Sticker?label=Stars&style=flat-square)](https://github.com/Malonan/Sticker/stargazers)
![GitHub release (latest by date)](https://img.shields.io/github/downloads/Malonan/Sticker/latest/total?label=Downloads%40Latest&style=flat-square)
![GitHub repo size](https://img.shields.io/github/repo-size/Malonan/Sticker?style=flat-square)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/Malonan/Sticker?style=flat-square)

<br>

---

<p align="center"><strong>Self-deployable, a Telegram group sticker set management robot that supports multi-group use.<strong></p>


<br>
<p align="center"><a href="https://github.com/Malonan/Sticker/releases">Download</a> | <a href="https://github.com/Malonan/Sticker/tree/master/wiki">Wiki</a> | <a href="https://github.com/Malonan/Sticker/issues">Issues</a> | <a href="#notice">Notice</a> | <a href="#donate">Donate</a> </p>


<br>

Example: <a href="https://t.me/IDontLikeThisStickerBot">@IDontLikeThisStickerBot</a>


<br>

# Requirement

1. Mysql5.7+ / MariaDB 10.4+
2. Redis6+ (Suggest Redis7+)

<br>

# Install

1. [Download](https://github.com/Malonan/Sticker/releases) archive package.
2. Unzip it to any folder.
3. `wget https://github.com/Malonan/Sticker/raw/Config/sticker.yml` -- Download configuration file, then modify it.
4. [Configure persistent services](#persistent-service) (optional)
5. Done!

<br>

# Compile

## Prerequisites

git, wget, [gmake2](https://github.com/3JoB/gmake2), golang(must 1.20.x), then no more
<br>

## Step
1. `git clone https://github.com/Malonan/Sticker && cd Sticker`
2. `gmake2`
3. `wget https://github.com/Malonan/Sticker/raw/Config/sticker.yml` -- Download configuration file, then modify it. (If gmake2 is used, it will be downloaded automatically)
4. [Configure persistent services](#persistent-service) (optional)
5. Done!

<br>

# Persistent Service
## Step

1. `wget https://github.com/Malonan/Sticker/raw/Config/sticker.service` (If gmake2 is used, it will be downloaded automatically)
2. Modify the executable file directory in it
3. `mv sticker.service /usr/lib/systemd/system/sticker.service`
4. `systemectl start sticker`
5. Done!

# Notice

1. Do not use any RC version of Golang.
2. Bots require full admin privileges before being enabled for the current group.
3. Privacy Mode must be turned off for bots (at BotFather).
4. Issues only deals with Debian11 issues, and please Google to solve the rest


<br>

# Donate

USDT/ETH: `0xF1C61348613489C5007a1A7aFbfd72bfdC4a3C3A`

TON: `UQCKEwhmCtTBmFNwESxeEmb8Ij4hSAU1wFs_inPy3tBbLExf`


<br>

# License
This software is distributed under Apache-2.0 license.

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FMalonan%2FSticker.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2FMalonan%2FSticker?ref=badge_large)
