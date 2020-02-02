package server

import (
	"fmt"
	"math/rand"
	"net/http"
	"path/filepath"

	"github.com/kingdom-tower/kingdom/princess"

	"github.com/labstack/echo/v4"
)

func (s *Server) showPrincess(c echo.Context) error {
	id := c.Param("id")
	path := filepath.Join(s.Dir, id)

	exists, err := princess.Exists(path)
	if err != nil {
		return err
	}
	if !exists {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Princess with id %s not found", id))
	}

	return c.File(path)
}

func (s *Server) savePrincess(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	path := filepath.Join(s.Dir, file.Filename)
	exists, err := princess.Exists(path)
	if err != nil {
		return err
	}
	if exists {
		return echo.NewHTTPError(http.StatusConflict, fmt.Sprintf("File %s already exists", file.Filename))
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if err := princess.Save(src, path); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func (s *Server) randomPrincess(c echo.Context) error {
	princesses, err := princess.All(s.Dir)
	if err != nil {
		return err
	}

	if len(princesses) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	i := rand.Intn(len(princesses))
	return c.File(filepath.Join(s.Dir, princesses[i]))
}
