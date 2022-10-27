package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	GroupName string   `yaml:"groupname"`
	Ips       []string `yaml:"ips"`
	Username  string   `yaml:"username"`
	Password  string   `yaml:"password"`
	cmd       string   `yaml:"cmd"`
}

func automate(config *Config) {
	sshConfig := &ssh.ClientConfig{
		User: config.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	for _, ip := range config.Ips {
		conn, err := ssh.Dial("tcp", ip+":22", sshConfig)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		session, err := conn.NewSession()
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()
		// configure terminal mode
		// setup standard out and error
		// uses writer interface
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr

		// run single command
		fmt.Println("Working on worker - ", ip)
		err = session.Run("ls")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	f, err := os.ReadFile("cohesity_assignment/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var c Config
	if err := yaml.Unmarshal(f, &c); err != nil {
		log.Fatal(err)
	}
	fmt.Println("%+v\n", c)
	automate(&c)
}
