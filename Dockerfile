FROM golang:1.18-alpine
WORKDIR /line

ENV TOKEN = "QrJIAEfzqCWKkPW9VsR2z/hiel2DOD6qBUzEc9XwVZgg/IaVmPHkHsOvPAdG3LewXxcTmsiDBV6cBNWPfi5y7cpH8D5RgV8yTRwQSoazsxOZoJ+0kSfIi/ei5+PwpKzxzGTTqkTlqr3VcYP4cGvU7wdB04t89/1O/w1cDnyilFU="
ENV CHANNEL_SECRET = "f42e89a44c8c70f05072a969ae46cd5d"
ENV CHANNEL_ID = '1657015218'

COPY . .
RUN go build main.go
RUN ls