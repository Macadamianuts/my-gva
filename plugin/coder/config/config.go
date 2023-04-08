package config

import "path/filepath"

type Config struct {
	Web    Web    `yaml:"web" mapstructure:"web"`
	Server Server `yaml:"server" mapstructure:"server"`
}

type Web struct {
	Root   string `yaml:"root" mapstructure:"root"`
	Api    string `yaml:"api" mapstructure:"api"`
	Form   string `yaml:"form" mapstructure:"form"`
	View   string `yaml:"view" mapstructure:"view"`
	Plugin string `yaml:"plugin" mapstructure:"plugin"`
}

func (c *Web) ApiPath(plugin string) string {
	return filepath.Join(c.Root, c.Plugin, plugin, c.Api)
}

func (c *Web) FormPath(plugin string) string {
	return filepath.Join(c.Root, c.Plugin, plugin, c.Form)
}

func (c *Web) ViewPath(plugin string) string {
	return filepath.Join(c.Root, c.Plugin, plugin, c.View)
}

type Server struct {
	Root     string `yaml:"root" mapstructure:"root"`
	Api      string `yaml:"api" mapstructure:"api"`
	Core     string `yaml:"core" mapstructure:"core"`
	Model    string `yaml:"model" mapstructure:"model"`
	Plugin   string `yaml:"plugin" mapstructure:"plugin"`
	Router   string `yaml:"router" mapstructure:"router"`
	Request  string `yaml:"request" mapstructure:"request"`
	Service  string `yaml:"service" mapstructure:"service"`
	Template string `yaml:"template" mapstructure:"template"`
}

func (c *Server) ApiPath(plugin string) string {
	return filepath.Join(c.Root, c.Plugin, plugin, c.Api)
}

func (c *Server) ModelPath(plugin string) string {
	return filepath.Join(c.Root, c.Plugin, plugin, c.Model)
}

func (c *Server) PluginRoot() string {
	return filepath.Join(c.Root, c.Plugin)
}

func (c *Server) PluginPath(plugin string) string {
	return filepath.Join(c.Root, c.Plugin, plugin)
}

func (c *Server) RouterPath(plugin string) string {
	return filepath.Join(c.Root, c.Plugin, plugin, c.Router)
}

func (c *Server) RequestPath(plugin string) string {
	return filepath.Join(c.Root, c.Plugin, plugin, c.Model, c.Request)
}

func (c *Server) ServicePath(plugin string) string {
	return filepath.Join(c.Root, c.Plugin, plugin, c.Service)
}

func (c *Server) TemplatePath() string {
	return filepath.Join(c.Root, c.Template)
}
