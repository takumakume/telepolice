package telepolice

import (
	"reflect"
	"testing"

	"github.com/takumakume/telepolice/pkg/kube"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/kubernetes/scheme"
)

var ns1 = &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "ns1"}}
var ns2 = &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "ns2"}}

type k8sYamlString string

func (s k8sYamlString) toRuntimeObject() runtime.Object {
	obj, _, _ := scheme.Codecs.UniversalDeserializer().Decode([]byte(s), nil, nil)
	return obj
}

func newKubernetesForTest(objects ...runtime.Object) *kube.Kubernetes {
	return &kube.Kubernetes{
		Clientset: fake.NewSimpleClientset(objects...),
	}
}

func TestTelepolice_SetNamespaces(t *testing.T) {
	type fields struct {
		kubernetes  *kube.Kubernetes
		namespaces  []string
		concurrency int
	}
	type args struct {
		namespaces []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "success",
			fields: fields{},
			args:   args{namespaces: []string{"ns"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			te := &Telepolice{
				kubernetes:  tt.fields.kubernetes,
				namespaces:  tt.fields.namespaces,
				concurrency: tt.fields.concurrency,
			}
			te.SetNamespaces(tt.args.namespaces)

			expected := []string{
				"ns",
			}
			actual := te.namespaces
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("%v != %v", expected, actual)
			}
		})
	}
}

func TestTelepolice_SetAllNamespaces(t *testing.T) {
	type fields struct {
		kubernetes  *kube.Kubernetes
		namespaces  []string
		concurrency int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:   "success",
			fields: fields{kubernetes: newKubernetesForTest(ns1, ns2)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			te := &Telepolice{
				kubernetes:  tt.fields.kubernetes,
				namespaces:  tt.fields.namespaces,
				concurrency: tt.fields.concurrency,
			}
			if err := te.SetAllNamespaces(); (err != nil) != tt.wantErr {
				t.Errorf("Telepolice.SetAllNamespaces() error = %v, wantErr %v", err, tt.wantErr)
			}

			expected := []string{
				"ns1",
				"ns2",
			}
			actual := te.namespaces
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("%v != %v", expected, actual)
			}
		})
	}
}

func TestTelepolice_EnableVerbose(t *testing.T) {
	te := &Telepolice{verbose: false}
	te.EnableVerbose()
	if te.verbose != true {
		t.Errorf("telepolice.EnableVerbose() verbose is not enabled")
	}
}
