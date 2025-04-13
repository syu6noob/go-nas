package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	path := c.Param("path")

	info, err := GetInfo(path)

	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate).SetMeta("Failed to load the target")
		return
	}

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(200, info)
}

func Download(c *gin.Context) {
	t := c.Param("path")

	info, infoErr := GetInfo(t)
	if infoErr != nil {
		c.Error(infoErr).
			SetType(gin.ErrorTypePrivate).
			SetMeta("Failed to load the target")
		return
	}
	if info.IsDir {
		c.Error(errors.New("failed to open a directory")).
			SetType(gin.ErrorTypePrivate).
			SetMeta("Failed to open a directory")
		return
	}

	rt, rtErr := GetRawTarget(t)
	if rtErr != nil {
		c.Error(rtErr).
			SetType(gin.ErrorTypePrivate).
			SetMeta("Failed to load the target")
	}

	file, errFile := os.Open(rt)
	if errFile != nil {
		c.Error(errFile).
			SetType(gin.ErrorTypePrivate).
			SetMeta("Failed to open file")
		return
	}
	defer file.Close()

	mime := getMime(file)
	c.Header("Content-Type", mime)
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", info.Name))
	c.File(rt)
}

func Stream(c *gin.Context) {
	t := c.Param("path")

	rt, rtErr := GetRawTarget(t)
	if rtErr != nil {
		c.Error(rtErr).
			SetType(gin.ErrorTypePrivate).
			SetMeta("Failed to load the target")
		return
	}

	stat, statErr := GetStat(t)
	if statErr != nil {
		c.Error(statErr).
			SetType(gin.ErrorTypePrivate).
			SetMeta("Failed to get file stat")
		return
	}

	file, errFile := os.Open(rt)
	if errFile != nil {
		c.Error(errFile).
			SetType(gin.ErrorTypePrivate).
			SetMeta("Failed to open file")
		return
	}
	defer file.Close()

	mime := getMime(file)
	c.Header("Content-Type", mime)

	rangeHeader := c.Request.Header.Get("Range")
	if rangeHeader != "" {
		byteString := strings.TrimPrefix(rangeHeader, "bytes=")
		parts := strings.Split(byteString, "-")
		start, _ := strconv.ParseInt(parts[0], 10, 64)
		end := stat.Size - 1

		if len(parts) == 2 && parts[1] != "" {
			end, _ = strconv.ParseInt(parts[1], 10, 64)
		}

		if start > end || end >= stat.Size {
			c.Header("Content-Range", fmt.Sprintf("bytes */%d", stat.Size))
			c.AbortWithStatus(http.StatusRequestedRangeNotSatisfiable)
			return
		}

		c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, stat.Size))
		c.Header("Content-Length", fmt.Sprintf("%d", end-start+1))
		c.Status(http.StatusPartialContent)

		_, err := file.Seek(start, io.SeekStart)
		if err != nil {
			c.Error(statErr).
				SetType(gin.ErrorTypePrivate).
				SetMeta("Failed to seek file")
			return
		}
		io.CopyN(c.Writer, file, end-start+1)
	} else {
		c.Header("Content-Length", fmt.Sprintf("%d", stat.Size))
		c.Status(http.StatusOK)
		io.Copy(c.Writer, file)
	}
}
