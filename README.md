# biometricscloud-pi
BiometricsCloud IoT application for the Raspberry Pi


1. Create a device from an authenticated user.

```
curl -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer DOY5lnC2SoWLznBp_4vS4Q" \
    -d '{"manufacturer":"Apple","device_code":"iPhone1,1"}' \
    http://127.0.0.1:8000/api/v1/devices
```


2. You will see the following output:


```json
{"id":2,"uuid":"21091bf3-3a34-446a-826b-8fda224ff5da","oauth2_client_id":"e9eeb2d53f6d0469","oauth2_client_secret":"72f601fea7e6ff07d89c18f7f5fcbb360c7daf64dd407230c8f7d84377f9328df0354ed5699c4f0dfda7cc9899390b68e04cc8050693e8c52ad26ca9fa17647fda0776abbb56897d46d37c0ef0958c1a670051993c73ba0acdd63d8cdca764cf8d84183ac886b010f977bee3a21ab9d8e18c059d612bba17aa6221c396ec48a","oauth2_redirect_url":"http://127.0.0.1:8000/appauth/code"}
```

3. Setup the environment variables.

```shell
export BIOMETRICSCLOUD_PI_CLIENT_ID=e9eeb2d53f6d0469; \
export BIOMETRICSCLOUD_PI_CLIENT_SECRET=72f601fea7e6ff07d89c18f7f5fcbb360c7daf64dd407230c8f7d84377f9328df0354ed5699c4f0dfda7cc9899390b68e04cc8050693e8c52ad26ca9fa17647fda0776abbb56897d46d37c0ef0958c1a670051993c73ba0acdd63d8cdca764cf8d84183ac886b010f977bee3a21ab9d8e18c059d612bba17aa6221c396ec48a; \
export BIOMETRICSCLOUD_PI_TOKEN_URL=http://localhost:8000/token;
```

3. Run our `dry-auth` command to confirm the device can connect.

```
go run main.go dry-auth;
```
