# Goodix gt911 touch level configurator

Ultra basic command line tool / golang lib to configure "touch level" (sensivity)
on goodix gt911 devices via i2c, since the linux driver for goodix devices does not
allow this kind of config directly.

## Ressources

- Linux driver source code : https://github.com/torvalds/linux/blob/master/drivers/input/touchscreen/goodix.c
- GT911 Datasheet : https://www.distec.de/fileadmin/pdf/produkte/Touchcontroller/DDGroup/GT911_Datasheet.pdf
- GT911 Programming guide : `docs` folder
- NXP community : https://community.nxp.com/t5/i-MX-Processors/Lcd-touch-sensitivity-issue-goodix/m-p/1315673#M177678
