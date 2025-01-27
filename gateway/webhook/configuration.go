package webhook

import (
	"strings"
	"time"

	"github.com/rudderlabs/rudder-server/config"
)

func loadConfig() {
	config.Initialize()

	sourceTransformerURL = strings.TrimSuffix(config.GetEnv("DEST_TRANSFORM_URL", "http://localhost:9090"), "/") + "/v0/sources"
	// Number of incoming webhooks that are batched before calling source transformer
	config.RegisterIntConfigVariable(32, &maxWebhookBatchSize, true, 1, "Gateway.webhook.maxBatchSize")
	// Timeout after which batch is formed anyway with whatever webhooks are available
	config.RegisterDurationConfigVariable(time.Duration(20), &webhookBatchTimeout, true, time.Millisecond, "Gateway.webhook.batchTimeoutInMS")
	// Multiple source transformers are used to generate rudder events from webhooks
	maxTransformerProcess = config.GetInt("Gateway.webhook.maxTransformerProcess", 64)
	// Max time till when retries to source transformer are done
	webhookRetryWaitMax = (config.GetDuration("Gateway.webhook.maxRetryTimeInS", time.Duration(10)) * time.Second)
	// Min time gap when retries to source transformer are done
	webhookRetryWaitMin = (config.GetDuration("Gateway.webhook.minRetryTimeInMS", time.Duration(100)) * time.Millisecond)
	// Max retry attempts to source transformer
	webhookRetryMax = config.GetInt("Gateway.webhook.maxRetry", 5)
}
