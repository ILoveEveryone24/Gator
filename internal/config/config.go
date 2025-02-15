package config

import (
    "fmt"
    "os"
    "io"
    "encoding/json"
)

type Config struct{
    Db_url string `json:"db_url"`
    Current_user_name string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error{
    c.Current_user_name = username

    homePath, err := os.UserHomeDir()
    if err != nil{
        return fmt.Errorf("Failed to read homepath: %v", err);
    }
    filePath := homePath + "/.gatorconfig.json"

    file, err := os.Create(filePath)
    if err != nil{
        return fmt.Errorf("Failed to open file: %v", err);
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    err = encoder.Encode(c)
    if err != nil{
        return fmt.Errorf("Failed to marshal data: %v", err);
    }

    return nil
}

func Read() (Config, error){
    homePath, err := os.UserHomeDir()
    if err != nil{
        return Config{}, fmt.Errorf("Failed to read homepath: %v", err);
    }
    filePath := homePath + "/.gatorconfig.json"
    file, err := os.Open(filePath)
    if err != nil{
        return Config{}, fmt.Errorf("Failed to open file: %v", err);
    }
    defer file.Close()

    bytes, err :=  io.ReadAll(file)
    if err != nil{
        return Config{}, fmt.Errorf("Failed to convert to bytes: %v", err);
    }

    var config Config
    err = json.Unmarshal(bytes, &config)
    if err != nil{
        return Config{}, fmt.Errorf("Failed to unmarshal bytes: %v", err);
    }
    return config, nil
}
