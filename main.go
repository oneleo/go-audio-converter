package main

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

func sanitizeFileName(fileName string) string {
    // 替換不允許的字符
    sanitized := strings.NewReplacer(
        "/", "-", "\\", "-", "?", "-", ":", "-", "*", "-", "\"", "-", "<", "-", ">", "-", "|", "-",
    ).Replace(fileName)
    return sanitized
}

func convertFiles(inputFolder, outputFolder string) error {
    files, err := os.ReadDir(inputFolder)
    if err != nil {
        return fmt.Errorf("unable to read directory: %w", err)
    }

    for _, file := range files {
        if file.IsDir() {
            continue
        }

        ext := filepath.Ext(file.Name())
        if ext != ".mp4" && ext != ".avi" && ext != ".mkv" && ext != ".webm" {
            continue
        }

        inputPath := filepath.Join(inputFolder, file.Name())
        sanitizedFileName := sanitizeFileName(file.Name())
        outputFile := filepath.Join(outputFolder, strings.TrimSuffix(sanitizedFileName, ext) + ".mp3")
        cmd := exec.Command("ffmpeg", "-i", inputPath, "-q:a", "0", "-map", "a", outputFile)

        fmt.Printf("Converting %s to MP3...\n", file.Name())
        if err := cmd.Run(); err != nil {
            fmt.Printf("Error converting %s: %v\n", file.Name(), err)
        } else {
            fmt.Printf("Converted %s to MP3.\n", file.Name())
        }
    }

    return nil
}

func main() {
    inputFolder := `/path/to/your/folder`
    outputFolder := `/path/to/output/folder`

    if err := os.MkdirAll(outputFolder, os.ModePerm); err != nil {
        fmt.Printf("Error creating output directory: %v\n", err)
        return
    }

    if err := convertFiles(inputFolder, outputFolder); err != nil {
        fmt.Printf("Error converting files: %v\n", err)
    }
}
