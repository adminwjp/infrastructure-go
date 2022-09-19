package bee_orms
var BeeOrmConfigInstance * BeeOrmConfig
func init()  {
	BeeOrmConfigInstance =&BeeOrmConfig{
		DriverNames: map[string]bool{},
		AliasNames: map[string]bool{},
	}
}
