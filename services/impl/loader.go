package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gogo/protobuf/jsonpb"
	"go.ketch.com/cli/ketch-cli/config"
	"go.ketch.com/cli/ketch-cli/services"
	"go.ketch.com/cli/ketch-cli/utils"
	"go.ketch.com/lib/app"
	"go.ketch.com/lib/orlop/errors"
	"io/fs"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	currentVersion    = "v1"
	defaultPlugin     = "plugin/plugin.js"
	defaultObjectsDir = "objects/"
	defaultAssetsDir  = "assets/"
)

type loader struct {
	reporter  services.Reporter
	client    *http.Client
}

func NewLoader(reporter services.Reporter, client *http.Client) services.Loader {
	return &loader{
		reporter:  reporter,
		client:    client,
	}
}

func (l *loader) Load(ctx context.Context, cfg *config.LoaderConfig) (*app.App, error) {
	var err error

	l.reporter.Report(ctx, "resolving inputs")
	if stat, err := os.Stat(cfg.AppConfigFile); err != nil || stat.IsDir() {
		return nil, errors.Errorf("file '%s' does not exist", cfg.AppConfigFile)
	}

	cfg.AppConfigFile, err = filepath.Abs(cfg.AppConfigFile)
	if err != nil {
		return nil, err
	}

	basePath := filepath.Dir(cfg.AppConfigFile)

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	if err = os.Chdir(basePath); err != nil {
		return nil, err
	}

	defer func() {
		os.Chdir(wd)
	}()

	if len(cfg.PluginFilename) > 0 {
		if stat, err := os.Stat(cfg.PluginFilename); err != nil || stat.IsDir() {
			return nil, errors.Errorf("plugin file '%s' does not exist", cfg.PluginFilename)
		}
	} else if stat, err := os.Stat(defaultPlugin); err == nil && !stat.IsDir() {
		cfg.PluginFilename = defaultPlugin
	}

	if len(cfg.PluginFilename) > 0 {
		cfg.PluginFilename, err = filepath.Abs(cfg.PluginFilename)
		if err != nil {
			return nil, err
		}
	}

	if len(cfg.ObjectsDir) > 0 {
		if stat, err := os.Stat(cfg.ObjectsDir); err != nil || !stat.IsDir() {
			return nil, errors.Errorf("objects directory '%s' does not exist", cfg.ObjectsDir)
		}
	} else if stat, err := os.Stat(defaultObjectsDir); err == nil && stat.IsDir() {
		cfg.ObjectsDir = defaultObjectsDir
	}

	if len(cfg.ObjectsDir) > 0 {
		cfg.ObjectsDir, err = filepath.Abs(cfg.ObjectsDir)
		if err != nil {
			return nil, err
		}
	}

	if len(cfg.AssetsDir) > 0 {
		if stat, err := os.Stat(cfg.AssetsDir); err != nil || !stat.IsDir() {
			return nil, errors.Errorf("assets directory '%s' does not exist", cfg.AssetsDir)
		}
	} else if stat, err := os.Stat(defaultAssetsDir); err == nil && stat.IsDir() {
		cfg.AssetsDir = defaultAssetsDir
	}

	if len(cfg.AssetsDir) > 0 {
		cfg.AssetsDir, err = filepath.Abs(cfg.AssetsDir)
		if err != nil {
			return nil, err
		}
	}

	l.reporter.Report(ctx, "loading '%s'...", cfg.AppConfigFile)

	b, err := ioutil.ReadFile(cfg.AppConfigFile)
	if err != nil {
		return nil, err
	}

	b = []byte(os.ExpandEnv(string(b)))

	manifest := &app.App{}

	if len(b) > 0 {
		if filepath.Ext(cfg.AppConfigFile) == ".yaml" || filepath.Ext(cfg.AppConfigFile) == ".yml" {
			b, err = utils.YAMLtoJSON(b)
			if err != nil {
				return nil, err
			}
		}

		if err = json.Unmarshal(b, manifest); err != nil {
			return nil, err
		}
	}

	if manifest.ApiVersion != currentVersion {
		return nil, errors.Errorf("apiVersion '%s' is invalid (expected '%s')", manifest.ApiVersion, currentVersion)
	}

	if manifest.Kind != "App" {
		return nil, errors.Errorf("kind '%s' is invalid (expected 'App')", manifest.Kind)
	}

	if manifest.Metadata == nil {
		return nil, errors.New("metadata is required")
	}

	if manifest.Data == nil {
		return nil, errors.New("data is required")
	}

	if len(cfg.Version) > 0 {
		manifest.Data.Version = cfg.Version
	}

	if len(manifest.Data.Version) == 0 {
		return nil, errors.New("app version must be specified in file or via --version")
	}
	return manifest, nil
}

func (l *loader) LoadExternalFiles(ctx context.Context, cfg *config.LoaderConfig, manifest *app.App) error {
	var err error

	basePath := filepath.Dir(cfg.AppConfigFile)

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	if err = os.Chdir(basePath); err != nil {
		return err
	}

	defer func() {
		os.Chdir(wd)
	}()

	if len(cfg.PluginFilename) > 0 {
		manifest.Data.Assets = append(manifest.Data.Assets, &app.AppAsset{
			ContentType: "application/javascript",
			Link:        cfg.PluginFilename,
			Name:        filepath.Base(cfg.PluginFilename),
		})
	}

	// Add custom objects
	if len(cfg.ObjectsDir) > 0 {
		l.reporter.Report(ctx, "loading objects...")

		err = filepath.WalkDir(cfg.ObjectsDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return err
			}

			relativePath := strings.TrimPrefix(path, cfg.ObjectsDir + "/")

			b, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			if filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml" {
				b, err = utils.YAMLtoJSON(b)
				if err != nil {
					return err
				}
			}

			if len(b) == 0 {
				return errors.Errorf("custom object '%s' is empty", relativePath)
			}

			obj := &app.AppObject{}
			if err = jsonpb.Unmarshal(bytes.NewReader(b), obj); err != nil {
				return err
			}

			if obj.ApiVersion != currentVersion {
				return errors.Errorf("custom object '%s' is invalid - apiVersion must be '%s'", relativePath, currentVersion)
			}

			if len(obj.Kind) == 0 {
				return errors.Errorf("custom object '%s' is invalid - kind is empty", relativePath)
			}

			if obj.Metadata == nil {
				return errors.Errorf("custom object '%s' is invalid - metadata is empty", relativePath)
			}

			if obj.Data == nil {
				return errors.Errorf("custom object '%s' is invalid - data is empty", relativePath)
			}

			manifest.Data.CustomObjects = append(manifest.Data.CustomObjects, obj)
			return nil
		})
		if err != nil {
			return err
		}
	}

	if manifest.Data.Logo != nil {
		l.reporter.Report(ctx, "loading logo...")

		if manifest.Data.Logo, err = l.loadImage(ctx, manifest.Data.Logo); err != nil {
			return err
		}
	}

	if len(manifest.Data.Previews) > 0 {
		l.reporter.Report(ctx, "loading previews...")
		for n := range manifest.Data.Previews {
			if manifest.Data.Previews[n], err = l.loadImage(ctx, manifest.Data.Previews[n]); err != nil {
				return err
			}
		}
	}

	if len(cfg.AssetsDir) > 0 {
		l.reporter.Report(ctx, "loading assets...")
		err = filepath.WalkDir(cfg.AssetsDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return err
			}

			manifest.Data.Assets = append(manifest.Data.Assets, &app.AppAsset{
				Link: path,
				Name: strings.TrimPrefix(path, cfg.AssetsDir + "/"),
			})
			return nil
		})
		if err != nil {
			return err
		}
	}

	for n := range manifest.Data.Assets {
		if manifest.Data.Assets[n], err = l.loadAsset(ctx, manifest.Data.Assets[n]); err != nil {
			return err
		}
	}

	return nil
}

func (l *loader) loadImage(ctx context.Context, in *app.AppImage) (*app.AppImage, error) {
	var err error

	in.ContentType = mime.TypeByExtension(path.Ext(in.Link))

	if in.Contents, err = l.getFileData(ctx, in.Link); err != nil {
		return nil, err
	}

	in.Link = ""

	return in, nil
}

func (l *loader) loadAsset(ctx context.Context, in *app.AppAsset) (*app.AppAsset, error) {
	var err error

	in.ContentType = mime.TypeByExtension(path.Ext(in.Link))

	if in.Contents, err = l.getFileData(ctx, in.Link); err != nil {
		return nil, err
	}

	in.Link = ""

	return in, nil
}

func (l *loader) getFileData(ctx context.Context, link string) (b []byte, err error) {
	l.reporter.Report(ctx, "loading %s...", link)

	if l.isRemoteLink(link) {
		b, err = l.getRemoteFileData(ctx, link)
	} else {
		b, err = l.getLocalFileData(ctx, link)
	}
	if err != nil {
		return nil, err
	}

	return
}

func (l *loader) getRemoteFileData(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := l.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (l *loader) getLocalFileData(ctx context.Context, link string) ([]byte, error) {
	return ioutil.ReadFile(link)
}

func (l *loader) isRemoteLink(link string) bool {
	_, err := url.ParseRequestURI(link)
	if err != nil {
		return false
	}

	u, err := url.Parse(link)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

