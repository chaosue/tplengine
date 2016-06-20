# tplengine
tplengine is a template renderer for echo framework.  tplengine bases on go template, provides a small set of extra pulgins, debug mode loads all template files on each rendering.

# Usage
```
	r := echo.New()
	viewDir := "../views/**/*.tpl"
	switch env {
	case "debug":
		Log.Warnf("Running in [%v] mode...", env)
		tplEngine := tplengine.NewDebugRenderer("myname")
		err = tplEngine.ParseGlob(viewDir)
		if err == nil {
			r.SetRenderer(tplEngine)
		}
	case "prod":
		tplEngine := tplengine.NewRenderer("myname")
		err = tplEngine.ParseGlob(viewDir)
		if err == nil {
			r.SetRenderer(tplEngine)
		}
	default:
		Log.Errorf("Invalid env[%v], available env are: %v, %v", env, "debug", "prod")
		os.Exit(1)
	}
```
