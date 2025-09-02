# `/build`

All generated binary, cmake, make files go there

## Build image
```bash
cd ../
docker build -t email-attachment-verification -f ./poster/build/Dockerfile .
docker run -it email-attachment-verification
docker exec -it email-attachment-verification_1 sh
```
