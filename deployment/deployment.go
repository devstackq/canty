package deployment

import (
	"log"
	"os/exec"
	"time"
)

type Service struct {
	Name       string
	StartCmd   string
	StopCmd    string
	HealthCmd  string
	RestartCmd string
}

type DeploymentManager struct {
	Services []Service
}

func NewDeploymentManager(services []Service) *DeploymentManager {
	return &DeploymentManager{Services: services}
}

func (dm *DeploymentManager) StartServices() {
	for _, service := range dm.Services {
		go func(s Service) {
			log.Printf("Starting service: %s", s.Name)
			cmd := exec.Command("sh", "-c", s.StartCmd)
			if err := cmd.Run(); err != nil {
				log.Fatalf("Failed to start service %s: %v", s.Name, err)
			}
		}(service)
	}
}

func (dm *DeploymentManager) StopServices() {
	for _, service := range dm.Services {
		go func(s Service) {
			log.Printf("Stopping service: %s", s.Name)
			cmd := exec.Command("sh", "-c", s.StopCmd)
			if err := cmd.Run(); err != nil {
				log.Fatalf("Failed to stop service %s: %v", s.Name, err)
			}
		}(service)
	}
}

func (dm *DeploymentManager) RestartServices() {
	for _, service := range dm.Services {
		go func(s Service) {
			log.Printf("Restarting service: %s", s.Name)
			cmd := exec.Command("sh", "-c", s.RestartCmd)
			if err := cmd.Run(); err != nil {
				log.Fatalf("Failed to restart service %s: %v", s.Name, err)
			}
		}(service)
	}
}

func (dm *DeploymentManager) MonitorServices() {
	for {
		for _, service := range dm.Services {
			go func(s Service) {
				log.Printf("Checking health of service: %s", s.Name)
				cmd := exec.Command("sh", "-c", s.HealthCmd)
				if err := cmd.Run(); err != nil {
					log.Printf("Service %s is not healthy: %v", s.Name, err)
					log.Printf("Restarting service: %s", s.Name)
					restartCmd := exec.Command("sh", "-c", s.RestartCmd)
					if err := restartCmd.Run(); err != nil {
						log.Fatalf("Failed to restart service %s: %v", s.Name, err)
					}
				} else {
					log.Printf("Service %s is healthy", s.Name)
				}
			}(service)
		}
		time.Sleep(30 * time.Second)
	}
}
