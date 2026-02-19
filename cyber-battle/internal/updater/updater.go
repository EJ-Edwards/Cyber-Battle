package updater

import (
	"cyber-battle/internal/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const currentVersion = "1.0.0"
const repoAPI = "https://api.github.com/repos/EJ-Edwards/Cyber-Battle/releases/latest"
const updateURL = "https://github.com/EJ-Edwards/Cyber-Battle/releases/latest"

func CheckForUpdates() (string, error) {
	resp, err := http.Get(repoAPI)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if data.TagName != "" && data.TagName != currentVersion {
		return data.TagName, nil
	}
	return "", errors.New("No updates available")
}

func NotifyIfUpdate() {
	utils.Display("Checking for updates...")
	latestVersion, err := CheckForUpdates()
	if err != nil {
		if err.Error() != "No updates available" {
			utils.Display("Update check failed: " + err.Error())
		}
		return
	}
	utils.Display(fmt.Sprintf("New version available: %s", latestVersion))
	utils.Display(fmt.Sprintf("Go update: %s", updateURL))
}
