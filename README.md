# biometricscloud-pi
BiometricsCloud IoT application for the Raspberry Pi

## Installation

1. Log into your Raspberry Pi

```
ssh -l pi 192.168.0.133
```
2. Checkout the project and build it.

```
cd ~/go/src/github.com/BCI-Innovation
git clone https://github.com/BCI-Innovation/biometricscloud-pi.git
cd biometricscloud-pi
go build
```

3. Setup environment variable(s).

```
export BIOMETRICSCLOUD_PI_REMOTE_SERVER_ADDRESS=https://biometricscloud.net
```

4. Log into your account and copy and paste your `access_token`. Example:

```
curl -X POST -H "Content-Type: application/json" \
     -d '{"username":"fherbert","password":"pleasechangeme"}' \
      https://biometricscloud.net/api/v1/login
```

5. Run the registration process when you are ready.

```
./biometricscloud-pi register --access_token=xyz
```

6. Finally run the application:

```
./biometricscloud-pi serve
```
