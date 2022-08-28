# k6 dotenv Extension

- simple extension to load env variables
- on your test script use:

`import dotenv from 'k6/x/dotenv';`

`dotenv.loadEnvFileFromEnv()` based on your ENV env variable
`dotenv.getEnvVar("VAR_NAME")` will return the value for the env variable given the loaded .env file