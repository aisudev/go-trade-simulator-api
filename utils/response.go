package utils

func Response(success bool, msg string, data interface{}, err error) map[string]interface{} {
	resp := map[string]interface{}{
		"success": success,
	}

	if msg != "" {
		resp["msg"] = msg
	}

	if data != nil {
		resp["data"] = data
	}

	if err != nil {
		resp["error"] = err.Error()
	}

	return resp
}
