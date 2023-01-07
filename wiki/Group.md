# Initialize for a group
1. You need to set the bot as a group admin and add permission to delete messages.
2. Use the `/refresh` command to refresh the list of group admins.
3. Done. The `/refresh` command is required to refresh the admin list every time an admin is added or removed.

<br>

# Add or remove stickers
1. Find that sticker!
2. Use the command `/make` to reference it!
3. When the sticker already exists in the group database, it will be deleted, otherwise it will be added.

# Sticker Whitelist Mode
The sticker whitelist mode is off by default, and currently does not support the switch. (The judging function has been completed, but the switching function has not been done)

# Group Whitelist Mode
If the instance is only for your personal use, please enable the group whitelist mode.


Please change `whitelist_mode` to `true` in `sticker.yaml`, and imitate the default example in `whitelist`, add whitelist group ID.


In whitelist mode, if the robot enters a non-whitelist group, then it will automatically exit the group.