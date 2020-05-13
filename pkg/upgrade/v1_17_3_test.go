package upgrade

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	v1 "github.com/jaegertracing/jaeger-operator/pkg/apis/jaegertracing/v1"
)

func TestUpgradeRemoveImagesv1_17_3(t *testing.T) {
	nsn := types.NamespacedName{Name: "my-instance"}
	existing := v1.NewJaeger(nsn)
	existing.Status.Version = "1.17.1"

	existing.Spec.Storage.EsIndexCleaner.Image = "test-image-1"
	existing.Spec.Storage.EsRollover.Image = "test-image-2"
	existing.Spec.Storage.Dependencies.Image = "test-image-3"
	objs := []runtime.Object{existing}

	s := scheme.Scheme
	s.AddKnownTypes(v1.SchemeGroupVersion, &v1.Jaeger{})
	s.AddKnownTypes(v1.SchemeGroupVersion, &v1.JaegerList{})
	cl := fake.NewFakeClient(objs...)

	assert.NoError(t, ManagedInstances(context.Background(), cl, cl))

	// verify
	persisted := &v1.Jaeger{}
	assert.NoError(t, cl.Get(context.Background(), nsn, persisted))

	assert.Equal(t, "", persisted.Spec.Storage.EsIndexCleaner.Image)
	assert.Equal(t, "", persisted.Spec.Storage.EsRollover.Image)
	assert.Equal(t, "", persisted.Spec.Storage.Dependencies.Image)
}
