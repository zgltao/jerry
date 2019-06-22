package logging

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", viper.GetString("app.runtime_root_path"),
		viper.GetString("log.save_path"))
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		viper.GetString("log.save_name"),
		time.Now().Format(viper.GetString("log.time_format")),
		viper.GetString("log.file_ext"),
	)
}
