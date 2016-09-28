# aws-getsessiontoken

A small utility for testing [AWS STS GetSessionToken](http://docs.aws.amazon.com/STS/latest/APIReference/API_GetSessionToken.html) requests.

## Usage

First, set your AWS credentials via environment variables.

```
$ export AWS_ACCESS_KEY_ID=...
$ export AWS_SECRET_ACCESS_KEY=...
```

Then, run the utility with an IAM Device Serial Number, and enter our MFA Token Code:

```
$ go run main.go -s "arn:aws:iam::123456789012:mfa/justin"
MFA Token Code: 123456

Request: {
  DurationSeconds: 900,
  SerialNumber: "arn:aws:iam::123456789012:mfa/justin",
  TokenCode: "123456"
}

Response: {
  Credentials: {
    AccessKeyId: "ASIAIDABCDEFGHIJKLMN",
    Expiration: 2016-09-28 15:14:37 +0000 UTC,
    SecretAccessKey: "x2AMZpgAzfnR4KmIlqwcJ2KLYG7sHu9dBH+nQkRxbpM",
    SessionToken: "GH/nEHnk846JH99JJiXYG4SyxwS18h9e+dNtMCUs9BwpuHut8KkDV9cMYUtu/RBLaFdMpX8lxWkNwH6+LRd5oN0VMGyAhs+1QWHvJdQVwGUoYi/tPU5dRt4m16QUp/9+r3KBmp5pLzQQ43MLgPgVOnIUqzccbr/YysvB0ess2sY="
  }
}

export AWS_ACCESS_KEY_ID=ASIAIDABCDEFGHIJKLMN
export AWS_SECRET_ACCESS_KEY=x2AMZpgAzfnR4KmIlqwcJ2KLYG7sHu9dBH+nQkRxbpM=
export AWS_SESSION_TOKEN=GH/nEHnk846JH99JJiXYG4SyxwS18h9e+dNtMCUs9BwpuHut8KkDV9cMYUtu/RBLaFdMpX8lxWkNwH6+LRd5oN0VMGyAhs+1QWHvJdQVwGUoYi/tPU5dRt4m16QUp/9+r3KBmp5pLzQQ43MLgPgVOnIUqzccbr/YysvB0ess2sY=
```
