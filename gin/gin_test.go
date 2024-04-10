package gin

import "testing"

func TestNestedGroup(t *testing.T) {
	r := New()
	v1 := r.Group("/v1")
	v2 := r.Group("/v2")
	v3 := r.Group("/v3")
	if v1.prefix != "/v1" {
		t.Fatal("v1 prefix should be /v1")
	}
	if v2.prefix != "/v1/v2" {
		t.Fatal("v2 prefix should be /v1/v2, out is", v2.prefix)
	}
	if v3.prefix != "/v1/v2/v3" {
		t.Fatal("v3 prefix should be /v1/v2/v3")

	}
}
