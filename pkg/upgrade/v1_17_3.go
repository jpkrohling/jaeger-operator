package upgrade

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "github.com/jaegertracing/jaeger-operator/pkg/apis/jaegertracing/v1"
)

func upgrade1_17_3(ctx context.Context, client client.Client, jaeger v1.Jaeger) (v1.Jaeger, error) {
	// Associated with https://github.com/jaegertracing/jaeger-operator/pull/1037
	if jaeger.Spec.Storage.EsIndexCleaner.Image != "" {
		jaeger.Spec.Storage.EsIndexCleaner.Image = ""
	}
	if jaeger.Spec.Storage.EsRollover.Image != "" {
		jaeger.Spec.Storage.EsRollover.Image = ""
	}
	if jaeger.Spec.Storage.Dependencies.Image != "" {
		jaeger.Spec.Storage.Dependencies.Image = ""
	}

	return jaeger, nil
}
