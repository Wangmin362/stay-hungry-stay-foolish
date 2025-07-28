#!/bin/bash
set -e

tools=(
    # 开发工具
    "sigs.k8s.io/kind@latest"
    "golang.org/x/vuln/cmd/govulncheck@latest"
    "google.golang.org/protobuf/cmd/protoc-gen-go@latest"
    "google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"
    "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest"
    "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest"
    "github.com/go-kratos/kratos/cmd/kratos/v2@latest"
    "github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest"
    "github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest"

    # 调试工具
    "github.com/go-delve/delve/cmd/dlv@latest"
    "github.com/google/wire/cmd/wire@latest"

    # 代码分析
    "honnef.co/go/tools/cmd/staticcheck@latest"
    "golang.org/x/tools/gopls@latest"
    "github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
)

# 并行安装 (最多4个并发)
printf "%s\n" "${tools[@]}" | xargs -P4 -I{} bash -c 'echo "Installing {}..." && go install -v {}'