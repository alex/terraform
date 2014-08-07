package openstack

type Config struct {
	TenantId   string `mapstructure:"tenant_id"`
	TenantName string `mapststructure:"tenant_name"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	AuthURL    string `mapstructure:"auth_url"`
	Region     string `mapstructure:"region"`
}
