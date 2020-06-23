package upgrade

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "github.com/jaegertracing/jaeger-operator/pkg/apis/jaegertracing/v1"
)

type version struct {
	v       string
	upgrade func(ctx context.Context, client client.Client, jaeger v1.Jaeger) (v1.Jaeger, error)
	next    *version
}

var (
	v1_11_0 = version{v: "1.11.0", upgrade: noop, next: &v1_12_0}
	v1_12_0 = version{v: "1.12.0", upgrade: noop, next: &v1_13_0}
	v1_13_0 = version{v: "1.13.0", upgrade: noop, next: &v1_13_1}
	v1_13_1 = version{v: "1.13.1", upgrade: noop, next: &v1_14_0}
	v1_14_0 = version{v: "1.14.0", upgrade: noop, next: &v1_15_0}
	v1_15_0 = version{v: "1.15.0", upgrade: upgrade1_15_0, next: &v1_16_0}
	v1_16_0 = version{v: "1.16.0", upgrade: noop, next: &v1_17_0}
	v1_17_0 = version{v: "1.17.0", upgrade: upgrade1_17_0, next: &v1_17_1}
	v1_17_1 = version{v: "1.17.1", upgrade: noop, next: &v1_17_2}
	v1_17_2 = version{v: "1.17.2", upgrade: noop, next: &v1_17_3}
	v1_17_3 = version{v: "1.17.3", upgrade: upgrade1_17_3, next: &v1_17_4}
	v1_17_4 = version{v: "1.17.4", upgrade: noop}

	latest = &v1_17_4

	versions = map[string]version{
		v1_11_0.v: v1_11_0,
		v1_12_0.v: v1_12_0,
		v1_13_0.v: v1_13_0,
		v1_13_1.v: v1_13_1,
		v1_14_0.v: v1_14_0,
		v1_15_0.v: v1_15_0,
		v1_16_0.v: v1_16_0,
		v1_17_0.v: v1_17_0,
		v1_17_1.v: v1_17_1,
		v1_17_2.v: v1_17_2,
		v1_17_3.v: v1_17_3,
		v1_17_4.v: v1_17_4,
	}
)

func noop(ctx context.Context, client client.Client, jaeger v1.Jaeger) (v1.Jaeger, error) {
	return jaeger, nil
}
