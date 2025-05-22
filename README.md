# RESTAPI for executing `ipmitool raw`

This app might useful for your Enterprise Server and you want to control it manually via ipmitool raw command

### How to use 
1. Pull docker image `docker pull ryakadev/ipmitool-raw:latest`
2. Run docker and expose port `8080`
3. Use POST Method to url `/ipmi/raw` with request body
```
{
  "host": "<YOUR_IPMI_ADDRESS>",
  "username": "<YOUR_IPMI_USER>",
  "password": "<YOUR_IPMI_PASS>",
  "raw": "<YOUR_RAW_COMMAND>"
}
```

### Example
- iDRAC disable fan auto

```
POST /ipmi/raw
{
  "host": "192.168.7.152",
  "username": "root",
  "password": "calvin",
  "raw": "0x30 0x30 0x01 0x00"
}
```
- iDRAC enable fan auto

```
POST /ipmi/raw
{
  "host": "192.168.7.152",
  "username": "root",
  "password": "calvin",
  "raw": "0x30 0x30 0x01 0x01"
}
```