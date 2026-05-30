package initapp

import (
	"fmt"
	"time"

	"github.com/rei0721/go-scaffold/internal/config"
	"github.com/rei0721/go-scaffold/pkg/iam"
	"github.com/rei0721/go-scaffold/pkg/iam/memory"
	"github.com/rei0721/go-scaffold/pkg/logger"
)

func NewIAM(cfg *config.Config, log logger.Logger) (iam.Service, error) {
	if !cfg.IAM.Enabled {
		if log != nil {
			log.Info("iam disabled")
		}
		return nil, nil
	}
	if cfg.IAM.Mode != "" && cfg.IAM.Mode != "memory" {
		return nil, fmt.Errorf("unsupported iam mode: %s", cfg.IAM.Mode)
	}
	service, err := memory.NewService(IAMMemoryConfig(cfg))
	if err != nil {
		return nil, fmt.Errorf("create iam service: %w", err)
	}
	if log != nil {
		log.Info("iam initialized", "mode", "memory")
	}
	return service, nil
}

func IAMMemoryConfig(cfg *config.Config) memory.Config {
	principals := make(map[string]iam.Principal, len(cfg.IAM.Tokens))
	for _, token := range cfg.IAM.Tokens {
		principals[token.Token] = iam.Principal{
			ID:         token.Principal.ID,
			Name:       token.Principal.Name,
			Roles:      append([]string(nil), token.Principal.Roles...),
			Attributes: copyStringMap(token.Principal.Attributes),
		}
	}
	policies := make([]iam.Policy, 0, len(cfg.IAM.Policies))
	for _, policy := range cfg.IAM.Policies {
		policies = append(policies, iam.Policy{
			Subject:   policy.Subject,
			Action:    iam.Action(policy.Action),
			Resource:  iam.Resource(policy.Resource),
			Effect:    iam.Effect(policy.Effect),
			ExpiresAt: parsePolicyExpiry(policy.ExpiresAt),
		})
	}
	return memory.Config{
		DefaultDeny: cfg.IAM.DefaultDenyEnabled(),
		Principals:  principals,
		Policies:    policies,
	}
}

func parsePolicyExpiry(value string) time.Time {
	if value == "" {
		return time.Time{}
	}
	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Time{}
	}
	return parsed.UTC()
}

func copyStringMap(src map[string]string) map[string]string {
	if len(src) == 0 {
		return nil
	}
	dst := make(map[string]string, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
