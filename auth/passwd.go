package auth

import (
	"bufio"
	"os"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Passwd struct {
	path string
}

func NewPasswd() *Passwd {
	path := os.Getenv("JASAS_PASSWD_PATH")
	if path == "" {
		path = "passwd"
	}
	return &Passwd{path}
}

func (passwd *Passwd) Exists(username string) (bool, error) {
	entries, err := passwd.read()
	if err != nil {
		return false, err
	}
	_, exists := entries[username]
	return exists, nil
}

func (passwd *Passwd) Put(username, password string) error {
	entries, err := passwd.read()
	if err != nil {
		return err
	}

	data, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "failed to create bcrypt password")
	}

	entries[username] = string(data)
	return passwd.write(entries)
}

func (passwd *Passwd) Remove(username string) error {
	entries, err := passwd.read()
	if err != nil {
		return err
	}
	delete(entries, username)
	return passwd.write(entries)
}

func (passwd *Passwd) Authenticate(username, password string) error {
	passwdFile, err := os.Open(passwd.path)
	if err != nil {
		return errors.Wrap(err, "failed to open passwd")
	}

	scanner := bufio.NewScanner(passwdFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		if len(parts) == 2 && parts[0] == username {
			err = bcrypt.CompareHashAndPassword([]byte(parts[1]), []byte(password))
			if err != nil {
				return errors.Wrap(err, "password seems to be wrong")
			}
			return nil
		}
	}
	return errors.New("username not found")
}

func (passwd *Passwd) write(entries map[string]string) error {
	passwdFile, err := os.OpenFile(passwd.path, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return errors.Wrap(err, "failed to open passwd")
	}
	defer passwdFile.Close()

	if len(entries) == 0 {
		passwdFile.Truncate(0)
	} else {

		for username, password := range entries {
			_, err := passwdFile.WriteString(username + ":" + password + "\n")
			if err != nil {
				return errors.Wrap(err, "failed to write entry to passwd")
			}
		}

	}

	return nil
}

func (passwd *Passwd) read() (map[string]string, error) {
	entries := make(map[string]string)

	_, err := os.Stat(passwd.path)
	if os.IsNotExist(err) {
		return entries, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "failed stat for passwd file")
	}

	passwdFile, err := os.Open(passwd.path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open passwd")
	}
	defer passwdFile.Close()

	scanner := bufio.NewScanner(passwdFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		if len(parts) == 2 {
			entries[parts[0]] = parts[1]
		}
	}

	return entries, nil
}
