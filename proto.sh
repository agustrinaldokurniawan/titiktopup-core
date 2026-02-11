#!/bin/bash

echo "🚀 Starting Protobuf Generation with Buf..."

buf dep update
buf generate

echo "✅ Done! PB, gRPC, and Gateway generated in /pb"