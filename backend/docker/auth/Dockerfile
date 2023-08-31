FROM golang:1.20-alpine AS builder
LABEL authors="dmitrydvp"

WORKDIR /svc-auth/

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o main ./cmd/svc-auth/main.go; \
    rm -R ./internal ./shared ./api *.mod *.sum ./config/*go

FROM alpine AS runner
LABEL authors="dmitrydvp"

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder ./svc-auth/main /
COPY --from=builder ./svc-auth/config/envs/*.env ./config/envs/

EXPOSE 50051

CMD ["/main"]