package config

import (
	"log"
)

type UserField struct {
	Name   string                 `yaml:"name"`
	Type   string                 `yaml:"type"`
	Params map[string]interface{} `yaml:"params"`
}

type Users struct {
	Fields []UserField `yaml:"fields"`
}

type Details struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
}

type Config struct {
	Details Details `yaml:"details"`
	Users   Users   `yaml:"users"`
}

func validateConfig() {
	// Validate the required fields for user
	if !Conf.ContainsUserFieldWithType("email", "string") {
		log.Fatal("email (string) is a required field for user")
	}
	if !Conf.ContainsUserFieldWithType("password", "string") {
		log.Fatal("password (string) is a required field for user")
	}

}

func (c *Config) ContainsUserField(fieldName string) bool {
	for _, field := range c.Users.Fields {
		if field.Name == fieldName {
			return true
		}
	}
	return false
}

func (c *Config) ContainsUserFieldWithType(fieldName, fieldType string) bool {
	for _, field := range c.Users.Fields {
		if field.Name == fieldName && field.Type == fieldType {
			return true
		}
	}
	return false
}

func (c *Config) GetUserField(fieldName string) UserField {
	for _, field := range c.Users.Fields {
		if field.Name == fieldName {
			return field
		}
	}
	return UserField{}
}

func HTMLInputTypeByUserField(field UserField) string {
	switch field.Type {
	case "string":
		switch field.Name {
		case "password":
			return "password"
		case "email", "mail":
			return "email"
		case "phone":
			return "tel"
		default:
			return "text"
		}
	case "int":
		return "number"
	case "bool":
		return "checkbox"
	case "date":
		return "date"
	case "datetime":
		return "datetime-local"
	default:
		return "text"
	}
}
