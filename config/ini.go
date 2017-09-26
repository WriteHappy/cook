package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"strings"
)

const (
	INI_FILE_SUFFIX = ".ini"
)

func Ini_open_dir(dname string) (*ini.File, error) {
	dp, err := os.Open(dname)
	if err != nil {
		return nil, err
	}

	fnames, err := dp.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	fp := ini.Empty()

	for _, fname := range fnames {
		if !strings.HasSuffix(fname, INI_FILE_SUFFIX) {
			continue
		}

		fullFpath := dname + "/" + fname

		fInfo, err := os.Stat(fullFpath)
		if err != nil {
			return nil, err
		}
		if !fInfo.Mode().IsRegular() {
			continue
		}

		if err := fp.Append(fullFpath); err != nil {
			return nil, err
		}
	}

	return fp, nil
}

func Ini_direct_get_key(sec *ini.Section, node, key string) *ini.Key {
	if k, err := sec.GetKey(key); err == nil {
		return k
	}
	panic(fmt.Errorf("Config error: [%s#%s] not exists", node, key))
	return nil
}
func Ini_inherit_get_key(sec, psec *ini.Section, node, key string) *ini.Key {
	if k, err := sec.GetKey(key); err == nil {
		return k
	}
	if k, err := psec.GetKey(key); err == nil {
		return k
	}
	panic(fmt.Errorf("Config error: [%s#%s] not exists", node, key))
	return nil
}
