//+build windows
package hservice

import (
	"golang.org/x/sys/windows/svc/mgr"
	"os"
	"path/filepath"
	"fmt"
	"golang.org/x/sys/windows/svc/eventlog"
	"golang.org/x/sys/windows/svc"
)

func IsSupported() bool {
	/* File will only be built on Windows, so return true */
	return true
}

type service struct {
}

/* Install attempts to install the service */
func (s *service) Install() error {
	manager, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer manager.Disconnect()

	path, err := ExePath()
	if err != nil {
		return err
	}
	serv, err := manager.CreateService(s.Name(), path, mgr.Config{DisplayName:s.Description()}, "is", "auto-started")
	if err != nil {
		return err
	}
	defer serv.Close()

	err = eventlog.InstallAsEventCreate(s.Name(), eventlog.Error|eventlog.Warning|eventlog.Info)
	if err != nil {
		serv.Delete()
		return fmt.Errorf("InstallAsEventCreate: %s", err.Error())
	}

	return nil
}

/* Uninstall attempts to remove the service */
func (s *service) Uninstall() error {
	manager, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer manager.Disconnect()

	serv, err := manager.OpenService(s.Name())
	if err != nil {
		return err
	}
	defer serv.Close()

	err = serv.Delete()
	if err != nil {
		return err
	}

	err = eventlog.Remove(s.Name())
	if err != nil {
		return err
	}

	return nil
}

/* Start attempts to manually start the service */
func (s *service) Start() error {
	manager, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer manager.Disconnect()

	serv, err := manager.OpenService(s.Name())
	if err != nil {
		return err
	}
	defer serv.Close()

	err = serv.Start("is", "manual-started")
	if err != nil {
		return err
	}

	return nil
}

/* Stop attempts to stop the running service */
func (s *service) Stop() error {
	manager, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer manager.Disconnect()

	serv, err := manager.OpenService(s.Name())
	if err != nil {
		return err
	}
	defer serv.Close()

	status, err := serv.Control(svc.Stop)
	if err != nil {
		return err
	}

	if status.State != svc.Stopped {
		return fmt.Errorf("Status after stopping was not stopped (%d), it was: %s", svc.Stopped, status.State)
	}

	return nil
}


func (s *service) Name() string {
	return daemonName
}

func (s *service) Description() string {
	return daemonDesc
}

/* ExePath attempts to return the full path to this executable */
func ExePath() (string, error) {
	var path string
	var err error

	/* First argument is always the executable name */
	fileName := os.Args[0]
	path, err = filepath.Abs(fileName)
	if err != nil {
		return path, err
	}

	fileInfo, err := os.Stat(path)
	getPath := func(f os.FileInfo, p string) (string, error) {
		if !f.Mode().IsDir() {
			return p, nil
		}

		err = fmt.Errorf("%s is a directory, not an executable", p)
		return path, err
	}
	if err == nil {
		return getPath(fileInfo, path)
	}
	if filepath.Ext(path) == "" {
		path += ".exe"
		fileInfo, err := os.Stat(path)
		if err == nil {
			return getPath(fileInfo, path)
		}
	}

	return "", err
}

func New() (Daemon, error) {
	s := new(service)

	return s, nil
}