# biometricscloud-pi
BiometricsCloud IoT application for the Raspberry Pi


1. Register an account.

```
curl -X POST -H "Content-Type: application/json" \
     -d '{"username":"fherbert","agree_tos":true,"country":"Canada","country_tel_code":"1","email":"fherbert@dune.com","first_name":"Frank","last_name":"Herbert","password":"pleasechangeme","password_repeat":"pleasechangeme","telephone":"(123) 456-7898"}' \
     http://localhost:8000/api/v1/register
```

2. Or login into the account.

```
curl -X POST -H "Content-Type: application/json" \
     -d '{"username":"fherbert","password":"pleasechangeme"}' \
      http://localhost:8000/api/v1/login
```

3. You'll see the `access_token` and the `refresh_token`. Create the environment variables.

```shell
export BIOMETRICSCLOUD_PI_ACCESS_TOKEN=1-Zju62jRQOvCM4T7t6DqA
export BIOMETRICSCLOUD_PI_REFRESH_TOKEN=1-Zju62jRQOvCM4T7t6DqA
```


4. Create a device from an authenticated user.

```
curl -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $BIOMETRICSCLOUD_PI_ACCESS_TOKEN" \
    -d '{"manufacturer":"Apple","device_code":"iPhone1,1"}' \
    http://127.0.0.1:8000/api/v1/devices
```


5. You will see the following output:


```json
{"id":1,"uuid":"8fd752e2-6cd7-4f61-8a86-3fea00bd07b3","oauth2_client_id":"bbd8ddfa9b59e1ef","oauth2_client_secret":"44e841d8d8ce82f2ea2d104717cc95bb279e4282be9fcd5087eeb0b84c5beefaf1bdf69a9548fa214bc5b873ed66302bc24445f2d623301252bd77e2da73c7c66838c77ff0cc6812244b87ec5d155c51febd39ba32505744e7b7d7b56c1d881537fd62dcccba33ce7e9f576263847ccb68772ed539db6853a285001ed711372","oauth2_redirect_url":"http://127.0.0.1:8000/appauth/code"}
```

6. Setup the environment variables.

```shell
export BIOMETRICSCLOUD_PI_REMOTE_SERVER_ADDRESS=http://localhost:8000
export BIOMETRICSCLOUD_PI_CLIENT_ID=bbd8ddfa9b59e1ef; \
export BIOMETRICSCLOUD_PI_CLIENT_SECRET=44e841d8d8ce82f2ea2d104717cc95bb279e4282be9fcd5087eeb0b84c5beefaf1bdf69a9548fa214bc5b873ed66302bc24445f2d623301252bd77e2da73c7c66838c77ff0cc6812244b87ec5d155c51febd39ba32505744e7b7d7b56c1d881537fd62dcccba33ce7e9f576263847ccb68772ed539db6853a285001ed711372; \
export BIOMETRICSCLOUD_PI_TOKEN_URL=http://localhost:8000/token;
```

7. Run our `dry-auth` command to confirm the device can connect.

```
go run main.go dry-auth;
```

8. Confirm you can access your account.

```
go run main.go profile;
```
