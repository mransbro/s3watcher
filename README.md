# S3 Object Check

Searches for last modified file in an s3 bucket and checks the size and last edit timestamp.
This can be run on a cronjob to check for backup objects.

Following environment variables are available:
```
BUCKET=                 // Only the bucket name ex. system-backups
FILE_AGE_IN_HOURS=      // Defaults to 24
AWS_REGION=             // 
FILE_PATH=              // Optional path prefix
OBJECT_SIZE_MB=         // Defaults to 1
```

## Usage

`docker run --env-file .env mransbro/s3objectcheck`

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).