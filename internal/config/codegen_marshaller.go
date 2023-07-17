package config


func (config *Config) StructToMap() map[string]interface{} {
	convertMap := make(map[string]interface{})
    
	convertMap["Name"] = config.Name
    
	convertMap["app_host"] = config.Host
    
	convertMap["app_port"] = config.Port
    
	convertMap["environment"] = config.Environment
    
	convertMap["volumes"] = config.Volumes
    
	return convertMap
    
}