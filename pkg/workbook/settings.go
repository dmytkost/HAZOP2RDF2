package workbook

import (
    "log"

    "github.com/spf13/viper"
)

type Settings struct {
    Program Program `mapstructure:"program"`
    Common  Common  `mapstructure:"common"`
    Hazop   Hazop   `mapstructure:"hazop"`
}

type Program struct {
    Package     string `mapstructure:"package"`
    Description string `mapstructure:"description"`
    Help        string `mapstructure:"help"`
    Version     string `mapstructure:"version"`
    Author      string `mapstructure:"author"`
}

type Common struct {
    DataDir    string `mapstructure:"data_dir"`
    DataExt    string `mapstructure:"data_ext"`
    ReportDir  string `mapstructure:"report_dir"`
    ReportExt  string `mapstructure:"report_ext"`
    TempFile   string `mapstructure:"temp_file"`
    TempStdout string `mapstructure:"temp_stdout"`
}

type DataType struct {
    Metadata int `mapstructure:"metadata"`
    Analysis int `mapstructure:"analysis"`
}

type CellType struct {
    String  int `mapstructure:"string"`
    Integer int `mapstructure:"integer"`
    Float   int `mapstructure:"float"`
}

type Element struct {
    DataType int    `mapstructure:"data_type"`
    Name     string `mapstructure:"name"`
    Regex    string `mapstructure:"regex"`
    CellType int    `mapstructure:"cell_type"`
    MinLen   int    `mapstructure:"min_len"`
    MaxLen   int    `mapstructure:"max_len"`
}

type Hazop struct {
    DataType DataType  `mapstructure:"data_type"`
    CellType CellType  `mapstructure:"cell_type"`
    Element  []Element `mapstructure:"element"`
}

var settings Settings

func init() {
    viper.SetConfigName("settings")
    viper.SetConfigType("toml")
    viper.AddConfigPath(".")

    if err := viper.ReadInConfig(); err != nil {
        log.Fatal("Error reading `.toml` settings: ", err)
    }

    if err := viper.Unmarshal(&settings); err != nil {
        log.Fatal("Error parsing `.toml` settings: ", err)
    }
}

func groupElements(dtype int) map[int]Element {
    elements := map[int]Element{}
    for i, e := range settings.Hazop.Element {
        if e.DataType != dtype {
            continue
        }
        elements[i] = e
    }
    return elements
}

func getElement(idx int) Element {
    return settings.Hazop.Element[idx]
}

func GetCommon() Common {
    return settings.Common
}

func GetProgram() Program {
    return settings.Program
}
