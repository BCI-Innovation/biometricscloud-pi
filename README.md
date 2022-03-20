# biometricscloud-pi
BiometricsCloud IoT application for the Raspberry Pi

### Pre-Configuration

```shell
export BIOMETRICSCLOUD_PI_REMOTE_SERVER_ADDRESS=http://127.0.0.1:8000
```

### Register Account

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

### Register Device and Send Data

1. Register our device:

```
curl -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $BIOMETRICSCLOUD_PI_ACCESS_TOKEN" \
    -d '{"manufacturer":"BCI","device_code":"GrowBuddyPC"}' \
    http://127.0.0.1:8000/api/v1/devices
```

2. You will get a response like this:

```json
{"id":2,"uuid":"53ae8b46-7f1a-450f-91e2-71caf80f3d6c","oauth2_client_id":"4cc9fd97c45f4aa7","oauth2_client_secret":"a21f67920d903875e5a706b73442ad87afd9b4d098d5c0495f153440f0c9b9a38f9942d422c85b7af535ed49b9abfd7b2a336a96cd894e48be667741fc2ef4edf0c286e1b7508e4968edc00cde7b8bafdd6fb2ac8994e3ecec5935fc390188513dd4c94d4754d64c0e2e2d100d2827513de4980039285969b13de558a11492b","oauth2_redirect_url":"http://127.0.0.1:8000/appauth/code"}
```

3. Create our metric.

```
curl -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $BIOMETRICSCLOUD_PI_ACCESS_TOKEN" \
    -d '{"device_id":2,"name":"Camera", "sample_type":"camera","quantity_type":"image","is_continuous_data":false}' \
    http://127.0.0.1:8000/api/v1/metrics
```

4. You will see a response like this:

```json
{"id":1,"uuid":"22773ad9-c76e-4416-8f7d-43ceb895c6fe","tenant_id":1,"device_id":2,"device_uuid":"53ae8b46-7f1a-450f-91e2-71caf80f3d6c","user_id":1,"name":"Camera","sample_type":"camera","quantity_type":"image","application_id":null,"application_name":null,"avg_yesterday":null,"avg_today":null,"avg_this_week":null,"avg_last_week":null,"avg_this_month":null,"avg_last_month":null,"avg_this_year":null,"avg_last_year":null,"latest_results_str":null,"latest_value":null,"latest_timestamp":null}
```

5. Finally try submitting a photo.

```
curl -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $BIOMETRICSCLOUD_PI_ACCESS_TOKEN" \
    -d '{"metric_id":1,"start_date":"2022-01-12T06:33:26.696888Z", "end_date":"2022-01-12T06:33:26.696888Z","sample_motion_context":"Active","sample_sensor_location":2, "upload_content":"iVBORw0KGgoAAAANSUhEUgAAABAAAAAPCAYAAADtc08vAAAABGdBTUEAANbY1E9YMgAAABl0RVh0U29mdHdhcmUAQWRvYmUgSW1hZ2VSZWFkeXHJZTwAAAKjSURBVHjadJNfSFNRGMC/e8/uxnVbVtoqJW1sBlFW9F8jtyIisHqIQQ9LDAqjggIfYhpFL6GDekkqhhU1/2SYkKKRLhCcL1OSwgfptoYo2l227V51Dtmf0zmbja3yg++ee78/v/N93zmXwRhDptjtdrDZbJfPnTp6I6/QWBSU5GXfxJdlJYcUnFZX4Xa7vZnxisyPlpZW9Lq9zeFoqL1kNBrguOUa3G+wQ2mRFqSQBPNs/iYSlgVgaAU1NTUQXorA8OBAZ/MLp8XrFUD8OgLms1Wg4hDEf46DgsFw+1FX3VbDtsZYLEo2a0kRKIBlmWMFunUdd29WYVkKYc/oJ3zk0F58omwnftNUjydcz7BvyIkvWk4K6/M2sGazOZlHNdmCJocvK9FvMe0wFsHIh1YApITKit0QjKpBo9VCPJGAcGQZzldWlAgz89c3FxQ0ZbXAMAxH3kuvXjjTf+uKJX9hcQkYFsG0Xwbd2hxQqThgSIBKyYGjcxAcbT2N81KojgLYFVCU6Njb90MdQTlMsKTveAKKN+YmkwJSOAmjkqdRwoIs9f2pgM2c6FxQvvPE2TPJq5QpKoGQ+YB3SgRhRoJoLAa6/FzqCvwXQER61eWy+aZFQCjlisbisG97MZh2FSbDV+Dq1QAQi8dd7tHxJVp6elAsSx9pIA1bFUBEFueCCyz7r0uhQPDDH6AEf9ZNNBgMKRpJ+u6bBERW5q9kamMBQzCCkVqz5jQxNacBHo8nGaQlZ26rq1d7Bnt5uhtHFCGycgqgR/u4te/Xx1Hh5cED+2ezKigvL09XIIpiQpKk2qft/Q9Nh/fkzs5Mw7gwFXg3MNyLEX+vurp60mq1Zv9MgiBkVrtI9LnLM2Ed++aP8Dz/oLu7+zOxhfR6Pb10Wa39FmAAqAgNiwPnYuoAAAAASUVORK5CYII=", "upload_filename":"favicon.ico"}' \
    http://127.0.0.1:8000/api/v1/photo-samples
```

### Usage

#### Setup Environment Variables:

```shell
export BIOMETRICSCLOUD_PI_REMOTE_SERVER_ADDRESS=http://127.0.0.1:8000
export BIOMETRICSCLOUD_PI_DEVICE_CAMERA_METRIC_ID=1
export BIOMETRICSCLOUD_PI_CLIENT_ID=bbd8ddfa9b59e1ef
export BIOMETRICSCLOUD_PI_CLIENT_SECRET=44e841d8d8ce82f2ea2d104717cc95bb279e4282be9fcd5087eeb0b84c5beefaf1bdf69a9548fa214bc5b873ed66302bc24445f2d623301252bd77e2da73c7c66838c77ff0cc6812244b87ec5d155c51febd39ba32505744e7b7d7b56c1d881537fd62dcccba33ce7e9f576263847ccb68772ed539db6853a285001ed711372
export BIOMETRICSCLOUD_PI_TOKEN_URL=http://127.0.0.1:8000/token
export BIOMETRICSCLOUD_PI_WIDTH=1640
export BIOMETRICSCLOUD_PI_HEIGHT=1232
export BIOMETRICSCLOUD_PI_FORMAT=png
export BIOMETRICSCLOUD_PI_WK_GRP=./
```

#### Commands

Verify we can perform oAuth 2.0 communication:

```
go run main.go dry-auth;
```

Start the service:

```
go run main.go serve;
```
