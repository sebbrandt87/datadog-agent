package sampler

// tracers with an env value of "" or agentEnv share
// the same sampler. This is required as remote is unaware
// of agentEnv and tracerEnv different values
func toSamplerEnv(tracerEnv, agentEnv string) string {
	env := tracerEnv
	if env == "" {
		env = agentEnv
	}
	return env
}

// tracers with empty env will have the same rate given
// as tracers with agentEnv
func rateWithEmptyEnv(samplerEnv, agentEnv string) bool {
	return samplerEnv == agentEnv
}
