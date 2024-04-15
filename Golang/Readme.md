硬盘信息
```
smartctl -i /dev/sdb

smartctl 7.4 2023-08-01 r5530 [x86_64-linux-6.6.12-0-lts] (local build)
Copyright (C) 2002-23, Bruce Allen, Christian Franke, www.smartmontools.org

=== START OF INFORMATION SECTION ===
Model Family:     Seagate IronWolf
Device Model:     ST4000VN008-2DR166
Serial Number:    ZDHBGR41
LU WWN Device Id: 5 000c50 0e4995be2
Firmware Version: SC60
User Capacity:    4,000,787,030,016 bytes [4.00 TB]
Sector Sizes:     512 bytes logical, 4096 bytes physical
Rotation Rate:    5980 rpm
Form Factor:      3.5 inches
Device is:        In smartctl database 7.3/5528
ATA Version is:   ACS-3 T13/2161-D revision 5
SATA Version is:  SATA 3.1, 6.0 Gb/s (current: 6.0 Gb/s)
Local Time is:    Tue Mar 19 17:13:05 2024 CST
SMART support is: Available - device has SMART capability.
SMART support is: Enabled
```

```
smartctl -A /dev/sdb|grep -E 'Power_On_Hours|Temperature_Celsius'

  9 Power_On_Hours          0x0032   082   082   000    Old_age   Always       -       16251 (247 140 0)
194 Temperature_Celsius     0x0022   032   044   000    Old_age   Always       -       32 (0 22 0 0 0)
```


S.M.A.R.T 信息
```
smartctl -H /dev/sdb

smartctl 7.4 2023-08-01 r5530 [x86_64-linux-6.6.12-0-lts] (local build)
Copyright (C) 2002-23, Bruce Allen, Christian Franke, www.smartmontools.org

=== START OF READ SMART DATA SECTION ===
SMART overall-health self-assessment test result: PASSED
```

```
➜  ~ zpool list -o name,size,free,health
NAME    SIZE   FREE    HEALTH
dpool  10.9T  10.7T    ONLINE

➜  ~ zpool list -Ho name,size,free,health,cap
dpool	10.9T	10.7T	ONLINE
```

```
smartctl --all /dev/sdb -j
```

```
smartctl --all /dev/sdb -j|jq -r '.model_name,.smart_status.passed,.temperature.current,.serial_number,.firmware_version,.user_capacity.bytes,.power_on_time.hours,.form_factor.name'
```

```
zpool iostat -H 1|awk '{print $1,$6,$7}'
```

```
zpool status |awk '/config:/,/errors:/'|awk 'NR>4 {print last} {last=$0}'
```

```
find /etc/periodic/15min -type f -exec head -n 1 -q {} +
```