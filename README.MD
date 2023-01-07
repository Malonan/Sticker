# <center><strong>I Don't Like This Sticker</strong></center>

<p align="center"><strong>Self-deployable, multi-group Telegram sticker set review bot<strong></p>


<br>
<p align="center"><a href="https://github.com/Malonan/Sticker/releases">Download</a> | <a href="https://github.com/Malonan/Sticker/wiki">Wiki</a> | <a href="https://github.com/Malonan/Sticker/issues">Issues</a> | <a href="https://github.com//Malonan/Sticker#Donate">Donate</a> | <a href="https://github.com//Malonan/Sticker#Notice">Notice</a></p>

<br>

Example: <a href="https://t.me/IDontLikeThisStickerBot">@IDontLikeThisStickerBot</a>

<br>

# Requirement
1. Mysql5.7+ / MariaDB 10.4+
2. Redis6+ (Suggest Redis7+)

<br>

# Install
1. <a href="https://github.com/Malonan/Sticker/releases">Download</a> archive package. (Only Linux X86_64 builds provided)
2. Unzip it to any folder.
3. `wget https://github.com/Malonan/Sticker/raw/Config/sticker.yml` -- Download configuration file, then modify it.
4. <a href="https://github.com//Malonan/Sticker#persistent-service">Configure persistent services </a> (optional)
5. Done!

<br>

# Compile

## Prerequisites

git, wget, GNU Make, golang(must 1.19.x), then no more
<br>

## Step
1. `git clone https://github.com/Malonan/Sticker && cd Sticker`
2. `go mod tidy && go build -ldflags "-s -w"` (Remember to turn off CGO)
3. `wget https://github.com/Malonan/Sticker/raw/Config/sticker.yml` -- Download configuration file, then modify it.
4. <a href="https://github.com//Malonan/Sticker#persistent-service">Configure persistent services </a> (optional)
5. Done!

<br>

# Persistent Service
## Step

1. `wget https://github.com/Malonan/Sticker/raw/Config/sticker.service`
2. Modify the executable file directory in it
3. `mv sticker.service /usr/lib/systemd/system/sticker.service`
4. `systemed start sticker`
5. Done!

# Notice

1. Do not use Golang1.20RC version to compile, it has a String Map bug, which will cause the program to fail to run.
2. Issues only deals with Debian11 issues, and please Google to solve the rest


<br>

# Donate

USDT/ETH: `0xF1C61348613489C5007a1A7aFbfd72bfdC4a3C3A`

TON: `UQCKEwhmCtTBmFNwESxeEmb8Ij4hSAU1wFs_inPy3tBbLExf`


<br>

# License
This software is distributed under Apache-2.0 license.