package interceptor

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

type contextKey string

const TraceIDKey = contextKey("x-trace-id")

func UnaryLoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	start := time.Now()

	// 获取 traceId 并放到 context 中
	traceID := getOrCreateTraceID(ctx)
	ctx = context.WithValue(ctx, TraceIDKey, traceID)

	// 将 trace ID 添加到响应 metadata
	pairs := metadata.Pairs(string(TraceIDKey), traceID)
	grpc.SetHeader(ctx, pairs)

	// 记录请求开始
	log.Printf("Unary Request: Method=%s, TraceID=%s", info.FullMethod, traceID)

	// 调用实际的处理逻辑
	resp, err = handler(ctx, req)

	// 记录请求结束和耗时
	log.Printf("Unary Response: Method=%s, TraceID=%s, Duration=%v, Error=%v",
		info.FullMethod, traceID, time.Since(start), err)

	return
}

func StreamLoggingInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	start := time.Now()
	ctx := ss.Context()

	// 获取 traceId 并放到 context 中
	traceID := getOrCreateTraceID(ctx)
	ctx = context.WithValue(ctx, TraceIDKey, traceID)

	// 将 trace ID 添加到响应 metadata
	pairs := metadata.Pairs(string(TraceIDKey), traceID)
	ss.SetHeader(pairs)

	// 记录流开始
	log.Printf("Stream Request: Method=%s, TraceID=%s", info.FullMethod, traceID)

	// ServerStream 包装上下文，调用实际的处理逻辑
	err := handler(srv, &wrappedServerStream{ServerStream: ss, ctx: ctx})

	// 记录流结束和耗时
	log.Printf("Stream Response: Method=%s, TraceID=%s, Duration=%v, Error=%v",
		info.FullMethod, traceID, time.Since(start), err)

	return err
}

type wrappedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedServerStream) Context() context.Context {
	return w.ctx
}

func getOrCreateTraceID(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if traceIDs := md.Get(string(TraceIDKey)); len(traceIDs) > 0 {
			return traceIDs[0]
		}
	}
	return uuid.New().String()
}
