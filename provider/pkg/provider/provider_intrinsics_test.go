package provider

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func Test_cidr(t *testing.T) {
	t.Run("ipv6, count=4, cidrbits=64", func(t *testing.T) {
		res, err := cidr(resource.NewPropertyMapFromMap(map[string]interface{}{
			"ipBlock":  "2600:1f16:44e:3e00::/56",
			"count":    4,
			"cidrBits": 64,
		}))
		assert.NoError(t, err)
		subnets := res["subnets"].ArrayValue()
		assert.Equal(t, []resource.PropertyValue{
			resource.NewStringProperty("2600:1f16:44e:3e00::/120"),
			resource.NewStringProperty("2600:1f16:44e:3e00::100/120"),
			resource.NewStringProperty("2600:1f16:44e:3e00::200/120"),
			resource.NewStringProperty("2600:1f16:44e:3e00::300/120"),
		}, subnets)

	})

	t.Run("ipv6, count=1, cidrbits=60", func(t *testing.T) {
		res, err := cidr(resource.NewPropertyMapFromMap(map[string]interface{}{
			"ipBlock":  "2a05:d024:d::/56",
			"count":    1,
			"cidrBits": 4,
		}))
		assert.NoError(t, err)
		subnets := res["subnets"].ArrayValue()
		assert.Equal(t, []resource.PropertyValue{
			resource.NewStringProperty("2a05:d024:d::/60"),
		}, subnets)

	})

	t.Run("ipv4, count=1, cidrbits=60", func(t *testing.T) {
		res, err := cidr(resource.NewPropertyMapFromMap(map[string]interface{}{
			"ipBlock":  "192.168.0.0/24",
			"count":    6,
			"cidrBits": 5,
		}))
		assert.NoError(t, err)
		subnets := res["subnets"].ArrayValue()
		assert.Equal(t, []resource.PropertyValue{
			resource.NewStringProperty("192.168.0.0/29"),
			resource.NewStringProperty("192.168.0.8/29"),
			resource.NewStringProperty("192.168.0.16/29"),
			resource.NewStringProperty("192.168.0.24/29"),
			resource.NewStringProperty("192.168.0.32/29"),
			resource.NewStringProperty("192.168.0.40/29"),
		}, subnets)

	})

	t.Run("ipv4 too long prefix", func(t *testing.T) {
		_, err := cidr(resource.NewPropertyMapFromMap(map[string]interface{}{
			"ipBlock":  "192.168.0.0/24",
			"count":    6,
			"cidrBits": 15,
		}))
		assert.ErrorContains(t, err, "cidrBits 15 would extend prefix to 39 bits, which is too long for an IPv4 address")
	})

	t.Run("ipv6 too long prefix", func(t *testing.T) {
		_, err := cidr(resource.NewPropertyMapFromMap(map[string]interface{}{
			"ipBlock":  "2600:1f16:44e:3e00::/56",
			"count":    6,
			"cidrBits": 80,
		}))
		assert.ErrorContains(t, err, "cidrBits 80 would extend prefix to 136 bits, which is too long for an IPv6 address")
	})

	t.Run("ipv6 not enough space", func(t *testing.T) {
		_, err := cidr(resource.NewPropertyMapFromMap(map[string]interface{}{
			"ipBlock":  "2600:1f16:44e:3e00::/56",
			"count":    3,
			"cidrBits": 1,
		}))
		assert.ErrorContains(t, err, "not enough remaining address space for a subnet with a prefix of 3 bits after 2600:1f16:44e:3e80::/57")
	})

}
