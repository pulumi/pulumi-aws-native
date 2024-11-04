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
			resource.NewStringProperty("2600:1f16:44e:3e00::/64"),
			resource.NewStringProperty("2600:1f16:44e:3e01::/64"),
			resource.NewStringProperty("2600:1f16:44e:3e02::/64"),
			resource.NewStringProperty("2600:1f16:44e:3e03::/64"),
		}, subnets)

	})

	t.Run("ipv6, count=2, cidrbits=64", func(t *testing.T) {
		res, err := cidr(resource.NewPropertyMapFromMap(map[string]interface{}{
			"ipBlock":  "2600:1f16:19ee:8000::/56",
			"count":    2,
			"cidrBits": 64,
		}))
		assert.NoError(t, err)
		subnets := res["subnets"].ArrayValue()
		assert.Equal(t, []resource.PropertyValue{
			resource.NewStringProperty("2600:1f16:19ee:8000::/64"),
			resource.NewStringProperty("2600:1f16:19ee:8001::/64"),
		}, subnets)
	})

	t.Run("ipv4, count=6, cidrbits=60", func(t *testing.T) {
		res, err := cidr(resource.NewPropertyMapFromMap(map[string]interface{}{
			"ipBlock":  "192.168.0.0/24",
			"count":    6,
			"cidrBits": 5,
		}))
		assert.NoError(t, err)
		subnets := res["subnets"].ArrayValue()
		assert.Equal(t, []resource.PropertyValue{
			resource.NewStringProperty("192.168.0.0/27"),
			resource.NewStringProperty("192.168.0.32/27"),
			resource.NewStringProperty("192.168.0.64/27"),
			resource.NewStringProperty("192.168.0.96/27"),
			resource.NewStringProperty("192.168.0.128/27"),
			resource.NewStringProperty("192.168.0.160/27"),
		}, subnets)

	})

	t.Run("ipv4 invalid cidrBits", func(t *testing.T) {
		_, err := cidr(resource.NewPropertyMapFromMap(map[string]interface{}{
			"ipBlock":  "192.168.0.0/24",
			"count":    6,
			"cidrBits": 36,
		}))
		assert.ErrorContains(t, err, "cidrBits 36 is more than 32 bits for an IPv4 address")
	})

	t.Run("ipv6 invalid cidrBits", func(t *testing.T) {
		_, err := cidr(resource.NewPropertyMapFromMap(map[string]interface{}{
			"ipBlock":  "2600:1f16:44e:3e00::/56",
			"count":    6,
			"cidrBits": 129,
		}))
		assert.ErrorContains(t, err, "cidrBits 129 is more than 128 bits for an IPv6 address")
	})

	t.Run("ipv6, not enough space", func(t *testing.T) {
		_, err := cidr(resource.NewPropertyMapFromMap(map[string]interface{}{
			"ipBlock": "2a05:d024:d::/56",
			"count":   3,
			// 128 - 71 = 57 (2 subnets)
			"cidrBits": 71,
		}))
		assert.ErrorContains(t, err, "not enough remaining address space for a subnet")
	})

	t.Run("ipv4, not enough space", func(t *testing.T) {
		_, err := cidr(resource.NewPropertyMapFromMap(map[string]interface{}{
			"ipBlock": "192.168.1.0/24",
			"count":   5,
			// 32 - 6 = 26 (4 subnets)
			"cidrBits": 6,
		}))
		assert.ErrorContains(t, err, "not enough remaining address space for a subnet")
	})

}
