opt43builder
============

A tool to generate DHCP option 43 strings

Usage
-----

```
./opt43builder -h
-delim=",": Specify the delimiter to seperate the options. Default is ","
-options="": DHCP Option format code:value[,code:value,...]

/opt43builder -options="0:junos.tgz,1:ex4300-01.cfg,3:http"
Decimal (Spaces): 00 09 106 117 110 111 115 46 116 103 122 01 13 101 120 52 51 48 48 45 48 49 46 99 102 103 03 04 104 116 116 112
Decimal: 00091061171101111154611610312201131011205251484845484946991021030304104116116112
Hex (Spaces): 00 09 6a 75 6e 6f 73 2e 74 67 7a 01 0d 65 78 34 33 30 30 2d 30 31 2e 63 66 67 03 04 68 74 74 70
Hex: 00096a756e6f732e74677a010d6578343330302d30312e636667030468747470
Hex Full(Spaces): 0x0 0x9 0x6a 0x75 0x6e 0x6f 0x73 0x2e 0x74 0x67 0x7a 0x1 0xd 0x65 0x78 0x34 0x33 0x30 0x30 0x2d 0x30 0x31 0x2e 0x63 0x66 0x67 0x3 0x4 0x68 0x74 0x74 0x70
Hex Full: 0x00x90x6a0x750x6e0x6f0x730x2e0x740x670x7a0x10xd0x650x780x340x330x300x300x2d0x300x310x2e0x630x660x670x30x40x680x740x740x70

```

Downloads
---------

To download please go to the [releases page](https://github.com/JNPRAutomate/opt43builder/releases)

Please select the correct download for your platform. The naming scheme is opt43builder_PLATFORM_ARCH.

-	opt43builder_darwin_386
-	opt43builder_darwin_amd64
-	opt43builder_freebsd_386
-	opt43builder_freebsd_amd64
-	opt43builder_freebsd_arm
-	opt43builder_linux_386
-	opt43builder_linux_amd64
-	opt43builder_linux_arm
-	opt43builder_netbsd_386
-	opt43builder_netbsd_amd64
-	opt43builder_netbsd_arm
-	opt43builder_openbsd_386
-	opt43builder_openbsd_amd64
-	opt43builder_plan9_386
-	opt43builder_windows_386.exe
-	opt43builder_windows_amd64.exe

### Note: darwin is Mac OS X
