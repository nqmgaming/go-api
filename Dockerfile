# Sử dụng Go image chính thức để build ứng dụng
FROM golang:1.23-alpine AS builder

# Đặt thư mục làm việc trong container
WORKDIR /app

# Copy toàn bộ mã nguồn vào container
COPY . .

# Tải các package và build ứng dụng
RUN go mod tidy
RUN go build -o main .

# Tạo image thứ hai để chạy ứng dụng
FROM alpine:latest

# Đặt thư mục làm việc trong container
WORKDIR /app

# Copy binary đã build từ giai đoạn đầu tiên vào image thứ hai
COPY --from=builder /app/main .

# Copy file .env (nếu có)
COPY .env .env

# Expose cổng 8080 (cổng mà ứng dụng Go của bạn đang lắng nghe)
EXPOSE 8080

# Chạy ứng dụng
CMD ["./main"]
