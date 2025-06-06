package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

//go:embed templates/*
var templateFS embed.FS

func host() {
	_ = godotenv.Load()

	port := "80"
	if v := os.Getenv("PORT"); v != "" {
		port = v
	}

	// Fiber HTML engine: load all .html in /views, including subfolders (layouts, partials)
	engine := html.NewFileSystem(http.FS(templateFS), ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	for _, t := range engine.Templates.Templates() {
		fmt.Println("Loaded template:", t.Name())
	}

	// Serve static files.
	app.Static("/static", "./static")

	// Home page route.
	app.Get("/", func(c *fiber.Ctx) error {
		data := fiber.Map{
			"Title": "OpenFluke - Home",
			"User":  "Guest",
		}
		content, err := renderContent(engine, "index", data)
		if err != nil {
			return err
		}
		data["Content"] = template.HTML(content)
		return c.Render("layout", data)
	})

	// About page
	app.Get("/about", func(c *fiber.Ctx) error {
		data := fiber.Map{
			"Title": "About",
			"User":  "Guest",
		}
		content, err := renderContent(engine, "about", data)
		if err != nil {
			return err
		}
		data["Content"] = template.HTML(content)
		return c.Render("layout", data)
	})

	app.Get("/paragon", func(c *fiber.Ctx) error {
		data := fiber.Map{
			"Title": "Paragon",
			"User":  "Guest",
		}
		content, err := renderContent(engine, "paragon", data)
		if err != nil {
			return err
		}
		data["Content"] = template.HTML(content)
		return c.Render("layout", data)
	})

	app.Get("/bench", func(c *fiber.Ctx) error {
		data := fiber.Map{
			"Title": "Paragon",
			"User":  "Guest",
		}
		content, err := renderContent(engine, "bench", data)
		if err != nil {
			return err
		}
		data["Content"] = template.HTML(content)
		return c.Render("layout", data)
	})

	app.Get("/biofoundry", func(c *fiber.Ctx) error {
		data := fiber.Map{
			"Title": "biofoundry",
			"User":  "Guest",
		}
		content, err := renderContent(engine, "biofoundry", data)
		if err != nil {
			return err
		}
		data["Content"] = template.HTML(content)
		return c.Render("layout", data)
	})

	app.Get("/demo", func(c *fiber.Ctx) error {
		data := fiber.Map{
			"Title": "Demo",
			"User":  "Guest",
		}
		content, err := renderContent(engine, "demo", data)
		if err != nil {
			return err
		}
		data["Content"] = template.HTML(content)
		return c.Render("layout", data)
	})

	/*log.Printf("Starting Fiber server on port %s\n", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}*/

	// Detect SSL certificates
	certFile := "./cert.pem"
	keyFile := "./key.pem"

	if isTLSEnabled() {
		log.Printf("SSL certificates detected. Running HTTPS on port 4443...")
		log.Fatal(app.ListenTLS("0.0.0.0:443", certFile, keyFile))

	} else {
		log.Printf("No SSL certificates found. Running HTTP on port 8080...")
		log.Fatal(app.Listen("0.0.0.0:" + port))
	}
}

func renderContent(engine *html.Engine, tmplName string, data fiber.Map) (string, error) {
	tmpl := engine.Templates.Lookup(tmplName)
	if tmpl == nil {
		return "", fmt.Errorf("template %s not found", tmplName)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func isTLSEnabled() bool {
	val := strings.ToLower(strings.TrimSpace(os.Getenv("ENABLE_TLS")))
	return val == "true" || val == "1" || val == "yes"
}
