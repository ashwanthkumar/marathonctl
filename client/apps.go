package client

import (
  "time"
  "encoding/json"
)

type AppInfo struct {
  ID                    string              `json:"id,omitempty"`
  Cmd                   string              `json:"cmd,omitempty"`
  Args                  []string            `json:"args"`
  Constraints           [][]string          `json:"constraints"`
  // Container             *Container          `json:"container,omitempty"`
  CPUs                  float64             `json:"cpus,omitempty"`
  Disk                  float64             `json:"disk,omitempty"`
  Env                   map[string]string   `json:"env"`
  Executor              string              `json:"executor,omitempty"`
  // HealthChecks          []*HealthCheck      `json:"healthChecks"`
  Instances             int                 `json:"instances,omitempty"`
  Mem                   float64             `json:"mem,omitempty"`
  // Tasks                 []*Task             `json:"tasks,omitempty"`
  Ports                 []int               `json:"ports"`
  RequirePorts          bool                `json:"requirePorts,omitempty"`
  BackoffSeconds        float64             `json:"backoffSeconds,omitempty"`
  BackoffFactor         float64             `json:"backoffFactor,omitempty"`
  MaxLaunchDelaySeconds float64             `json:"maxLaunchDelaySeconds,omitempty"`
  // Deployments           []map[string]string `json:"deployments,omitempty"`
  Dependencies          []string            `json:"dependencies"`
  TasksRunning          int                 `json:"tasksRunning,omitempty"`
  TasksStaged           int                 `json:"tasksStaged,omitempty"`
  TasksHealthy          int                 `json:"tasksHealthy,omitempty"`
  TasksUnhealthy        int                 `json:"tasksUnhealthy,omitempty"`
  User                  string              `json:"user,omitempty"`
  // UpgradeStrategy       *UpgradeStrategy    `json:"upgradeStrategy,omitempty"`
  Uris                  []string            `json:"uris"`
  Version               string              `json:"version,omitempty"`
  // VersionInfo           *VersionInfo        `json:"versionInfo,omitempty"`
  Labels                map[string]string   `json:"labels,omitempty"`
  AcceptedResourceRoles []string            `json:"acceptedResourceRoles,omitempty"`
  // LastTaskFailure       *LastTaskFailure    `json:"lastTaskFailure,omitempty"`
}

type Apps struct {
  Apps                  []AppInfo          `json:"apps"`
}

func (m *Marathon) AppsInfo() (*Apps, error) {
  body, err := handle(httpClient.Timeout(5000 * time.Millisecond).Get(m.Url + "/v2/info").End())
  if err != nil {
    return nil, err
  }

  var apps Apps
  err = json.Unmarshal([]byte(body), &apps)
  return &apps, err
}
